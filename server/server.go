package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dipper-iot/dipper-engine-server/database"
	"github.com/dipper-iot/dipper-engine-server/graph/generated"
	"github.com/dipper-iot/dipper-engine-server/graph/resolver"
	"github.com/dipper-iot/dipper-engine-server/models"
	"github.com/dipper-iot/dipper-engine-server/rule_engine"
	"github.com/dipper-iot/dipper-engine-server/services"
	"github.com/dipper-iot/dipper-engine/bus"
	"github.com/dipper-iot/dipper-engine/data"
	"github.com/dipper-iot/dipper-engine/queue"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	stop         bool
	config       *ConfigServer
	serverGraph  *handler.Server
	db           *database.Database
	engine       *rule_engine.Engine
	router       chi.Router
	listen       net.Listener
	version      string
	signalStop   chan os.Signal
	inputSession queue.QueueEngine[*data.Session]
	busData      bus.Bus
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewServer(config *ConfigServer) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		config: config,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Server) SetVersion(version string) {
	s.version = version
}

func (s *Server) Init(ctx context.Context) error {
	s.signalStop = make(chan os.Signal)
	signal.Notify(s.signalStop, os.Interrupt, os.Kill)

	s.db = database.NewDatabase(s.config.Database)
	if err := s.db.Connect(ctx); err != nil {
		return err
	}

	s.inputSession = queue.NewDefaultQueue[*data.Session]("input_session")
	s.busData = bus.NewDefaultBus()

	read, write := s.db.Client()
	chanService := services.NewChanServiceImpl(read, write)
	nodeService := services.NewNodeServiceImpl(read, write)
	sessionService := services.NewSessionService(read, write, s.inputSession)

	storeDb := rule_engine.NewStoreDb(
		chanService,
		nodeService,
		sessionService,
	)

	s.engine = rule_engine.NewEngine(
		s.config.Engine,
		storeDb,
		s.inputSession,
		s.busData,
		s.ctx,
	)

	manageSessionService := services.NewManageSessionServiceImpl(read, write, s.engine)

	s.router = chi.NewRouter()
	s.serverGraph = handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver.NewResolver(
				s.version,
				nodeService,
				chanService,
				sessionService,
				manageSessionService,
				s.busData,
			)}))

	s.router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	s.serverGraph.AddTransport(transport.POST{})
	s.serverGraph.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	s.serverGraph.Use(extension.Introspection{})

	// router
	if s.config.Dev {
		s.router.Handle("/graph/playground", playground.Handler("GraphQL playground of Dipper Engine", "/graph/query"))
	}
	s.router.Handle("/graph/query", s.serverGraph)

	// start session infinity
	go func() {
		for {
			limit := 50
			skip := 0

			list, _, err := manageSessionService.List(ctx, &models.ListSessionRequest{
				Limit: limit,
				Skip:  skip,
			})
			if err != nil {
				logrus.Error(err)
				return
			}
			for _, session := range list {
				if session.Infinite {
					s.engine.StartSession(session.ID)
				} else {
					err = sessionService.Delete(ctx, session.ID)
					if err != nil {
						logrus.Error(err)
						continue
					}
				}
			}

			if len(list) < limit {
				return
			}
			skip = skip + limit
		}

	}()

	return nil
}

func (s *Server) Run(address string) error {
	s.stop = false
	err := s.engine.Start()
	if err != nil {
		return err
	}

	s.listen, err = net.Listen("tcp", address)
	if err != nil {
		return err
	}
	log.Printf("Server Listen On %s", s.listen.Addr().String())
	go func() {
		err := http.Serve(s.listen, s.router)
		if err != nil {
			log.Fatal(err)
		}
		s.signalStop <- os.Interrupt
	}()

	<-s.signalStop
	log.Print("Stop Server")
	if !s.stop {
		s.Stop()
	}
	return nil
}

func (s *Server) Stop() {
	s.cancel()
	if s.listen != nil {
		s.listen.Close()
	}
	if s.engine != nil {
		s.engine.Stop()
	}
	if s.db != nil {
		s.db.Close()
	}
	s.stop = true
}
