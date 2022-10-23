package server

import "github.com/dipper-iot/dipper-engine-server/database"

type ConfigServer struct {
	Database *database.ConfigDatabase `json:"database"`
	Dev      bool                     `json:"dev"`
}
