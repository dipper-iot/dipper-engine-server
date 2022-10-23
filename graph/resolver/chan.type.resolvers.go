package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/graph/generated"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
)

// Status is the resolver for the status field.
func (r *chanResolver) Status(ctx context.Context, obj *ent.RuleChan) (*model.ChanStatus, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// CreatedAt is the resolver for the created_at field.
func (r *chanResolver) CreatedAt(ctx context.Context, obj *ent.RuleChan) (*string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// UpdatedAt is the resolver for the updated_at field.
func (r *chanResolver) UpdatedAt(ctx context.Context, obj *ent.RuleChan) (*string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updated_at"))
}

// Chan returns generated.ChanResolver implementation.
func (r *Resolver) Chan() generated.ChanResolver { return &chanResolver{r} }

type chanResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *chanResolver) ID(ctx context.Context, obj *ent.RuleChan) (uint64, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}
