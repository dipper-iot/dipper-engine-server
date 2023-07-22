package util

import (
	"context"
	"encoding/json"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/dipper-iot/dipper-engine/bus"
	"github.com/dipper-iot/dipper-engine/data"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

func QueryLog(ctx context.Context, busData bus.Bus, input *model.DebugFilter, chanLog chan []*model.SessionOutput) {

	lock := sync.Mutex{}
	list := make([]*model.SessionOutput, 0)

	busData.Subscribe(ctx, "debug-output", func(dataByte []byte) {
		var dataOutput data.OutputEngine

		err := json.Unmarshal(dataByte, &dataOutput)
		if err != nil {
			logrus.Fatal(err)
			return
		}

		if input != nil && input.ChanID != nil {
			if dataOutput.ChanId != *input.ChanID {
				return
			}
		}

		chanId, err := strconv.ParseUint(dataOutput.ChanId, 10, 64)
		if err != nil {
			logrus.Error(err)
			return
		}

		item := &model.SessionOutput{
			Data:       dataOutput.Data,
			ChanID:     chanId,
			FormEngine: dataOutput.FromEngine,
			SessionID:  dataOutput.SessionId,
			Type:       model.OutputTypeSuccess,
		}

		if dataOutput.Type == data.TypeOutputEngineError {
			if dataOutput.Error != nil {
				msg := dataOutput.Error.ErrorDetail.Error()
				item.Error = &model.ErrorEngine{
					Message:     &dataOutput.Error.Message,
					Code:        &dataOutput.Error.Code,
					ErrorDetail: &msg,
				}
			}
		}

		lock.Lock()
		list = append(list, item)
		lock.Unlock()
	})

	t := time.Tick(time.Duration(500) * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			{
				return
			}
		case <-t:
			lock.Lock()
			if len(list) > 0 {
				go func(list []*model.SessionOutput) {
					select {
					case <-ctx.Done():
						return
					case chanLog <- list:
						return
					}
				}(list)
				list = make([]*model.SessionOutput, 0)
			}

			lock.Unlock()
		}
	}
}
