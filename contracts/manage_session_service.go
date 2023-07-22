package contracts

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/models"
)

type ManageSessionService interface {
	List(ctx context.Context, request *models.ListSessionRequest) (list []*ent.Session, total int, err error)
	GetSection(ctx context.Context, id uint64) (session *ent.Session, err error)
	Stop(ctx context.Context, id uint64) (err error)
}
