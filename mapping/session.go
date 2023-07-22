package mapping

import (
	"fmt"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine/data"
	"github.com/sirupsen/logrus"
	"time"
)

func MapSession(session *ent.Session) *data.Info {
	dataInfo := new(data.Info)

	dataInfo.Data = session.Data
	dataInfo.Time = &session.CreatedAt
	dataInfo.ChanId = fmt.Sprintf("%d", session.ID)
	dataInfo.Result = map[string]*data.OutputEngine{}
	dataInfo.EndCount = session.EndCount
	dataInfo.Timeout = time.Duration(session.Timeout)
	dataInfo.Id = session.ID
	dataInfo.Infinite = session.Infinite
	dataInfo.MapNode = map[string]*data.NodeRule{}

	err := data.MapToStruct(session.Result, &dataInfo.Result)
	if err != nil {
		logrus.Error(err)
		return dataInfo
	}

	return dataInfo
}

func MapInfoToSession(data *data.Info) *ent.Session {
	session := new(ent.Session)

	mapResult := make(map[string]interface{})
	for key, engine := range data.Result {
		mapResult[key] = engine
	}

	session.ID = data.Id
	session.Result = mapResult
	session.EndCount = data.EndCount
	session.UpdatedAt = time.Now()

	return session
}
