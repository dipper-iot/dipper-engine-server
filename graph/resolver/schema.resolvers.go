package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dipper-iot/dipper-engine-server/graph/generated"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
)

// Version is the resolver for the version field.
func (r *mutationResolver) Version(ctx context.Context) (*model.Version, error) {
	return &model.Version{
		Version: r.version,
	}, nil
}

// Version is the resolver for the version field.
func (r *queryResolver) Version(ctx context.Context) (*model.Version, error) {
	return &model.Version{
		Version: r.version,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
