package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/dipper-iot/dipper-engine-server/models"
)

// ListSession is the resolver for the ListSession field.
func (r *queryResolver) ListSession(ctx context.Context, input model.SessionListRequest) (*model.ListSessionInfo, error) {
	list, total, err := r.manageSessionService.List(ctx, &models.ListSessionRequest{
		Infinity: &input.Infinity,
		Skip:     input.Skip,
		Limit:    input.Limit,
	})
	if err != nil {
		return nil, err
	}

	return &model.ListSessionInfo{
		Sessions: list,
		Total:    total,
	}, nil
}

// GetSession is the resolver for the GetSession field.
func (r *queryResolver) GetSession(ctx context.Context, id uint64) (*ent.Session, error) {
	return r.manageSessionService.GetSection(ctx, id)
}
