package main

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/database"
	server "github.com/dipper-iot/dipper-engine-server/server"
	"log"
)

var (
	Version = "beta"
)

func main() {
	s := server.NewServer(&server.ConfigServer{
		Database: &database.ConfigDatabase{
			Type:      database.SqliteType,
			Debug:     true,
			Migration: true,
		},
		Dev: true,
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
