package server

import (
	"github.com/dipper-iot/dipper-engine-server/database"
	"github.com/dipper-iot/dipper-engine/core"
)

type ConfigServer struct {
	Database *database.ConfigDatabase `json:"database"`
	Engine   *core.ConfigEngine       `json:"engine"`
	Dev      bool                     `json:"dev"`
}
