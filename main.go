package main

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/database"
	server "github.com/dipper-iot/dipper-engine-server/server"
	"github.com/dipper-iot/dipper-engine/core"
	"log"
)

var (
	Version = "beta"
)

func main() {
	s := server.NewServer(&server.ConfigServer{
		Database: &database.ConfigDatabase{
			Type:      database.SqliteType,
			Debug:     false,
			Migration: true,
		},
		Dev: true,
		Engine: &core.ConfigEngine{
			Rules: map[string]*core.RuleOption{
				"log-core": {
					Options: map[string]interface{}{},
					Worker:  1,
					Enable:  true,
				},
				"arithmetic": {
					Options: map[string]interface{}{},
					Worker:  1,
					Enable:  true,
				},
				"fork": {
					Options: map[string]interface{}{},
					Worker:  1,
					Enable:  true,
				},
				"switch": {
					Options: map[string]interface{}{},
					Worker:  1,
					Enable:  true,
				},
				"conditional": {
					Options: map[string]interface{}{},
					Worker:  1,
					Enable:  true,
				},
				"input-redis-queue": {
					Options: map[string]interface{}{
						"redis_address": "127.0.0.1:6379",
						"redis_db":      0,
					},
					Worker: 1,
					Enable: true,
				},
				"output-redis-queue": {
					Options: map[string]interface{}{
						"redis_address": "127.0.0.1:6379",
						"redis_db":      0,
					},
					Worker: 1,
					Enable: true,
				},
			},
			Log: core.ConfigLog{
				FileName:      "engine.log",
				Output:        "console",
				Level:         "debug",
				LogMethodName: true,
			},
			Plugins:        []string{},
			BusMap:         map[string]string{},
			TimeoutSession: 30,
		},
	})
	s.SetVersion(Version)

	err := s.Init(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := s.Run(":8029"); err != nil {
		log.Fatal(err)
	}
}
