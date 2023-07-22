package services

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/ent/session"
	"github.com/dipper-iot/dipper-engine-server/models"
	"github.com/dipper-iot/dipper-engine-server/rule_engine"
	"sync"
)

type ManageSessionServiceImpl struct {
	engine      *rule_engine.Engine
	clientRead  *ent.Client
	clientWrite *ent.Client
}

func NewManageSessionServiceImpl(
	clientRead *ent.Client,
	clientWrite *ent.Client,
	engine *rule_engine.Engine,
) *ManageSessionServiceImpl {
	return &ManageSessionServiceImpl{engine: engine, clientRead: clientRead, clientWrite: clientWrite}
}

func (m *ManageSessionServiceImpl) List(ctx context.Context, req *models.ListSessionRequest) (list []*ent.Session, total int, err error) {
	query := m.clientRead.Session.Query()

	if req.Infinity != nil {
		query = query.Where(session.Infinite(*req.Infinity))
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	var (
		errTotal error
		errList  error
	)

	go func() {
		total, errTotal = query.Clone().Count(ctx)
		wg.Done()
	}()

	go func() {
		list, errList = query.Clone().Limit(req.Limit).Offset(req.Skip).All(ctx)
		wg.Done()
	}()

	wg.Wait()

	if errTotal != nil {
		err = errTotal
		return
	}
	err = errList

	return
}

func (m *ManageSessionServiceImpl) GetSection(ctx context.Context, id uint64) (session *ent.Session, err error) {
	return m.clientRead.Session.Get(ctx, id)
}

func (m *ManageSessionServiceImpl) Stop(ctx context.Context, id uint64) (err error) {
	ses, err := m.GetSection(ctx, id)
	if err != nil {
		return err
	}

	if ses.Infinite {
		ruleIds := m.engine.ControlGetRule(ses.ID)
		for _, node := range ruleIds {
			m.engine.ControlStopSession(node, ses.ID)
		}
	}

	_, err = m.clientWrite.Session.Delete().Where(session.ID(id)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
