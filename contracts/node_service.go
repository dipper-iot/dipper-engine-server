package contracts

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/models"
)

type NodeService interface {
	All(ctx context.Context, chanId uint64) (list []*ent.RuleNode, err error)
	Get(ctx context.Context, id uint64) (detail *ent.RuleNode, err error)
	UpdateAll(ctx context.Context, chanId uint64, list []*models.Node) (err error)
	Update(ctx context.Context, chanId uint64, data *models.Node) (err error)
	Delete(ctx context.Context, id uint64) (err error)
}
