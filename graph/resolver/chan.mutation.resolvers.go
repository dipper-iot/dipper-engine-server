package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dipper-iot/dipper-engine-server/ent/rulechan"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/dipper-iot/dipper-engine-server/models"
)

// CreateChan is the resolver for the CreateChan field.
func (r *mutationResolver) CreateChan(ctx context.Context, input model.InputChan) (bool, error) {
	description := ""
	if input.Description != nil {
		description = *input.Description
	}
	err := r.chanService.Create(ctx, &models.Chan{
		Name:        input.Name,
		Description: description,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateChan is the resolver for the UpdateChan field.
func (r *mutationResolver) UpdateChan(ctx context.Context, id uint64, input model.InputChan) (bool, error) {
	description := ""
	if input.Description != nil {
		description = *input.Description
	}
	err := r.chanService.Update(ctx, id, &models.Chan{
		Name:        input.Name,
		Description: description,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateStatusChan is the resolver for the UpdateStatusChan field.
func (r *mutationResolver) UpdateStatusChan(ctx context.Context, input model.SetStatusChan) (bool, error) {
	status := rulechan.StatusActivated
	if input.Status == model.ChanStatusDeactivated {
		status = rulechan.StatusDeactivated
	}
	err := r.chanService.UpdateStatus(ctx, input.ID, status)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateRootNodeChan is the resolver for the UpdateRootNodeChan field.
func (r *mutationResolver) UpdateRootNodeChan(ctx context.Context, id uint64, rootNode string) (bool, error) {
	err := r.chanService.UpdateRootNode(ctx, id, rootNode)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteChan is the resolver for the DeleteChan field.
func (r *mutationResolver) DeleteChan(ctx context.Context, id uint64) (bool, error) {
	err := r.chanService.Delete(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
