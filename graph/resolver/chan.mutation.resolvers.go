package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/dipper-iot/dipper-engine-server/graph/model"
)

// CreateChan is the resolver for the CreateChan field.
func (r *mutationResolver) CreateChan(ctx context.Context, input model.InputChan) (bool, error) {
	panic(fmt.Errorf("not implemented: CreateChan - CreateChan"))
}

// UpdateChan is the resolver for the UpdateChan field.
func (r *mutationResolver) UpdateChan(ctx context.Context, id uint64, input model.InputChan) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdateChan - UpdateChan"))
}

// UpdateStatusChan is the resolver for the UpdateStatusChan field.
func (r *mutationResolver) UpdateStatusChan(ctx context.Context, input model.SetStatusChan) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdateStatusChan - UpdateStatusChan"))
}

// UpdateRootNodeChan is the resolver for the UpdateRootNodeChan field.
func (r *mutationResolver) UpdateRootNodeChan(ctx context.Context, id uint64, rootNode string) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdateRootNodeChan - UpdateRootNodeChan"))
}

// DeleteChan is the resolver for the DeleteChan field.
func (r *mutationResolver) DeleteChan(ctx context.Context, id uint64) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteChan - DeleteChan"))
}
