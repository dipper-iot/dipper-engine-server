package resolver

import (
	"github.com/dipper-iot/dipper-engine-server/contracts"
	"github.com/dipper-iot/dipper-engine/bus"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	nodeService          contracts.NodeService
	chanService          contracts.ChanService
	sessionService       contracts.SessionService
	manageSessionService contracts.ManageSessionService
	busData              bus.Bus
	version              string
}

func NewResolver(
	version string,
	nodeService contracts.NodeService,
	chanService contracts.ChanService,
	sessionService contracts.SessionService,
	manageSessionService contracts.ManageSessionService,
	busData bus.Bus,
) *Resolver {
	return &Resolver{
		version:              version,
		nodeService:          nodeService,
		chanService:          chanService,
		sessionService:       sessionService,
		manageSessionService: manageSessionService,
		busData:              busData,
	}
}
