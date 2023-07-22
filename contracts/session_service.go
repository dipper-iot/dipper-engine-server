package contracts

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/models"
)

type SessionService interface {
	Get(ctx context.Context, id uint64) (detail *ent.Session, err error)
	Create(ctx context.Context, data *ent.Session) (err error)
	Update(ctx context.Context, data *ent.Session) (err error)
	Delete(ctx context.Context, id uint64) (err error)
	RunSession(ctx context.Context, data *models.CreateSession) error
}
