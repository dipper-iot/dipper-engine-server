package test

import (
	"context"
	"encoding/json"
	"github.com/dipper-iot/dipper-engine-server/rule_engine/debug"
	"github.com/go-redis/redis/v9"
	"io"
	"testing"
)

func TestInput(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	err := client.Ping(context.TODO()).Err()
	if err != nil {
		t.Error(err)
		return
	}
	var transferData = map[string]interface{}{
		"default": map[string]interface{}{
			"a": 10,
			"b": 20,
			"d": 5,
		},
	}
	dataByte, err := json.Marshal(&transferData)
	if err != nil {
		t.Error(err)
		return
	}

	err = client.LPush(context.TODO(), "test_queue", dataByte).Err()
	if err != nil {
		t.Error(err)
		return
	}
Loop:
	for {
		dataRaw, err := client.RPop(context.TODO(), "test_output_queue").Bytes()
		if err == io.EOF {
			return
		}
		if err == redis.Nil {
			continue
		}

		if err != nil {
			t.Error(err)
			continue
		}

		var transferData map[string]interface{}
		err = json.Unmarshal(dataRaw, &transferData)
		if err != nil {
			t.Error(err)
			continue
		}
		debug.PrintJson(transferData, "OutPut: ")
		break Loop
	}

}
