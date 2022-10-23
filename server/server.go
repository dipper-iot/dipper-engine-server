package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dipper-iot/dipper-engine-server/database"
	"github.com/dipper-iot/dipper-engine-server/graph/generated"
	"github.com/dipper-iot/dipper-engine-server/graph/resolver"
	"github.com/dipper-iot/dipper-engine-server/services"
	"github.com/go-chi/chi/v5"
	"log"
	"net"
	"net/http"
)

type Server struct {
	config      *ConfigServer
	serverGraph *handler.Server
	db          *database.Database
	router      chi.Router
	listen      net.Listener
	version     string
}

func NewServer(config *ConfigServer) *Server {
	return &Server{config: config}
}

func (s *Server) SetVersion(version string) {
	s.version = version
}

func (s *Server) Init(ctx context.Context) error {
	s.db = database.NewDatabase(s.config.Database)
	if err := s.db.Connect(ctx); err != nil {
		return err
	}
	read, write := s.db.Client()
	chanService := services.NewChanServiceImpl(read, write)
	nodeService := services.NewNodeServiceImpl(read, write)

	s.router = chi.NewRouter()
	s.serverGraph = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.NewResolver(s.version, nodeService, chanService, nil)}))

	// router
	s.router.Handle("/graph/playground", playground.Handler("GraphQL playground of Dipper Engine", "/graph/query"))
	s.router.Handle("/graph/query", s.serverGraph)

	return nil
}

func (s *Server) Run(address string) error {
	var err error
	s.listen, err = net.Listen("tcp", address)
	if err != nil {
		return err
	}
	log.Printf("Server Listen On %s", s.listen.Addr().String())
	return http.Serve(s.listen, s.router)
}

func (s *Server) Stop() {
	if s.listen != nil {
		s.listen.Close()
	}
}
