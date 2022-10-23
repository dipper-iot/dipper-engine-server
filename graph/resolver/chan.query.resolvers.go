package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/dipper-iot/dipper-engine-server/models"
)

// ListChan is the resolver for the ListChan field.
func (r *queryResolver) ListChan(ctx context.Context, input models.ListChanRequest) (*model.ListChan, error) {
	list, total, err := r.chanService.List(ctx, &input)
	if err != nil {
		return nil, err
	}

	return &model.ListChan{
		List:  list,
		Total: total,
	}, nil
}

// GetChan is the resolver for the GetChan field.
func (r *queryResolver) GetChan(ctx context.Context, id uint64) (*ent.RuleChan, error) {
	return r.chanService.Get(ctx, id)
}
