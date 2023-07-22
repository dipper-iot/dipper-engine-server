package services

import (
	"context"
	"fmt"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/ent/rulenode"
	"github.com/dipper-iot/dipper-engine-server/ent/session"
	"github.com/dipper-iot/dipper-engine-server/mapping"
	"github.com/dipper-iot/dipper-engine-server/models"
	"github.com/dipper-iot/dipper-engine/data"
	"github.com/dipper-iot/dipper-engine/queue"
	"github.com/sirupsen/logrus"
	"time"
)

type SessionService struct {
	clientRead   *ent.Client
	clientWrite  *ent.Client
	inputSession queue.QueueEngine[*data.Session]
}

func NewSessionService(clientRead *ent.Client, clientWrite *ent.Client, inputSession queue.QueueEngine[*data.Session]) *SessionService {
	return &SessionService{clientRead: clientRead, clientWrite: clientWrite, inputSession: inputSession}
}

func (s SessionService) Get(ctx context.Context, id uint64) (detail *ent.Session, err error) {
	detail, err = s.clientRead.Session.Query().Where(session.ID(id)).First(ctx)
	return
}

func (s SessionService) Create(ctx context.Context, data *ent.Session) (err error) {
	query := s.clientWrite.Session.Create()

	query.SetID(data.ID)
	query.SetInfinite(data.Infinite)
	query.SetChainID(data.ChainID)
	query.SetData(data.Data)
	query.SetEndCount(data.EndCount)
	query.SetIsTest(data.IsTest)
	query.SetTimeout(data.Timeout)
	query.SetCreatedAt(data.CreatedAt)

	_, err = query.Save(ctx)
	return
}

func (s SessionService) Update(ctx context.Context, data *ent.Session) (err error) {
	query := s.clientWrite.Session.Update()

	query.SetEndCount(data.EndCount)
	query.SetResult(data.Result)
	query.SetUpdatedAt(time.Now())

	_, err = query.Where(session.IDEQ(data.ID)).Save(ctx)
	return
}

func (s SessionService) Delete(ctx context.Context, id uint64) (err error) {
	query := s.clientWrite.Session.Delete()
	_, err = query.Where(session.ID(id)).Exec(ctx)
	return
}

func (s SessionService) RunSession(ctx context.Context, session *models.CreateSession) error {

	chanData, err := s.clientRead.RuleChan.Get(ctx, session.ChanID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	nodes, err := s.clientRead.RuleNode.Query().Where(rulenode.ChainID(session.ChanID)).All(ctx)
	if err != nil {
		logrus.Error(err)
		return err
	}

	mapNode := make(map[string]*data.NodeRule)

	for _, node := range nodes {
		mapNode[node.NodeID] = mapping.MapNode(node)
	}

	return s.inputSession.Publish(ctx, &data.Session{
		ChanId:   fmt.Sprintf("%d", session.ChanID),
		Data:     session.Data,
		Result:   map[string]*data.OutputEngine{},
		RootNode: chanData.RootNode,
		MapNode:  mapNode,
	})
}
