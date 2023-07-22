package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/ent/rulechan"
	"github.com/dipper-iot/dipper-engine-server/graph/generated"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
)

// Nodes is the resolver for the nodes field.
func (r *chanResolver) Nodes(ctx context.Context, obj *ent.RuleChan) ([]*ent.RuleNode, error) {
	allNode, err := r.nodeService.All(ctx, obj.ID)
	if err == nil {
		return allNode, nil
	}
	return nil, nil
}

// Status is the resolver for the status field.
func (r *chanResolver) Status(ctx context.Context, obj *ent.RuleChan) (*model.ChanStatus, error) {
	status := model.ChanStatusActivated
	if obj.Status == rulechan.StatusDeactivated {
		status = model.ChanStatusDeactivated
	}
	return &status, nil
}

// Chan returns generated.ChanResolver implementation.
func (r *Resolver) Chan() generated.ChanResolver { return &chanResolver{r} }

type chanResolver struct{ *Resolver }
