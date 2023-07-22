package rule_engine

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/contracts"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/mapping"
	"github.com/dipper-iot/dipper-engine/data"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type StoreDb struct {
	chanService    contracts.ChanService
	nodeService    contracts.NodeService
	sessionService contracts.SessionService
}

func NewStoreDb(chanService contracts.ChanService, nodeService contracts.NodeService, sessionService contracts.SessionService) *StoreDb {
	return &StoreDb{chanService: chanService, nodeService: nodeService, sessionService: sessionService}
}

func (s StoreDb) Add(session *data.Info) error {

	timeData := time.Now()
	if session.Time != nil {
		timeData = *session.Time
	}
	chanId, err := strconv.ParseUint(session.ChanId, 10, 64)
	if err != nil {
		logrus.Error(err)
		return err
	}
	ctx := context.Background()

	detail, err := s.chanService.Get(ctx, chanId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return s.sessionService.Create(ctx, &ent.Session{
		ID:        session.Id,
		ChainID:   detail.ID,
		Data:      session.Data,
		Infinite:  session.Infinite,
		Result:    map[string]interface{}{},
		Timeout:   int(session.Timeout),
		CreatedAt: timeData,
		EndCount:  session.EndCount,
	})
}

func (s StoreDb) Get(sessionId uint64) *data.Info {

	session, err := s.sessionService.Get(context.Background(), sessionId)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	dataInfo := mapping.MapSession(session)

	chanData, err := s.chanService.Get(context.Background(), session.ChainID)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	listNode, err := s.nodeService.All(context.Background(), session.ChainID)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	for _, node := range listNode {
		dataInfo.MapNode[node.NodeID] = mapping.MapNode(node)
		if node.NodeID == chanData.RootNode {
			dataInfo.RootNode = dataInfo.MapNode[node.NodeID]
		}
	}

	return dataInfo
}

func (s StoreDb) Has(sessionId uint64) bool {
	session, err := s.sessionService.Get(context.Background(), sessionId)
	if err != nil {
		logrus.Error(err)
		return false
	}
	return session.ID == sessionId
}

func (s StoreDb) Done(sessionId uint64, result *data.OutputEngine) (session *data.ResultSession, success bool) {
	success = false

	sessionInfo := s.Get(sessionId)
	if sessionInfo == nil {
		logrus.Error("Not found ", sessionId)
		return
	}

	if sessionInfo.Infinite {
		return nil, false
	}

	sessionInfo.EndCount -= 1
	if sessionInfo.Result == nil {
		sessionInfo.Result = map[string]*data.OutputEngine{}
	}
	sessionInfo.Result[result.IdNode] = result

	if sessionInfo.EndCount > 0 {

		session := mapping.MapInfoToSession(sessionInfo)
		err := s.sessionService.Update(context.Background(), session)
		if err != nil {
			logrus.Error("Session Service Delete ", err)
			return nil, false
		}

		return nil, false
	}
	success = true

	// delete store
	err := s.sessionService.Delete(context.Background(), sessionId)
	if err != nil {
		logrus.Error("Session Service Delete ", err)
		return
	}

	// result
	session = &data.ResultSession{
		Id:     sessionInfo.Id,
		Data:   sessionInfo.Data,
		ChanId: sessionInfo.ChanId,
		Result: sessionInfo.Result,
	}

	return
}
