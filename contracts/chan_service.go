package contracts

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/ent/rulechan"
	"github.com/dipper-iot/dipper-engine-server/models"
)

type ChanService interface {
	List(ctx context.Context, req *models.ListChanRequest) (list []*ent.RuleChan, total int, err error)
	Get(ctx context.Context, id uint64) (detail *ent.RuleChan, err error)
	Create(ctx context.Context, data *models.Chan) (err error)
	Update(ctx context.Context, id uint64, data *models.Chan) (err error)
	UpdateStatus(ctx context.Context, id uint64, status rulechan.Status) (err error)
	UpdateRootNode(ctx context.Context, id uint64, rootNode string) (err error)
	Delete(ctx context.Context, id uint64) (err error)
}
