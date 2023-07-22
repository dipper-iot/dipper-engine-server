package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/dipper-iot/dipper-engine-server/graph/util"
	"github.com/dipper-iot/dipper-engine-server/models"
)

// UpdateNode is the resolver for the UpdateNode field.
func (r *mutationResolver) UpdateNode(ctx context.Context, chanID uint64, input model.InputNode) (bool, error) {
	data := util.InputNode(&input)
	err := r.nodeService.Update(ctx, chanID, data)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateNodes is the resolver for the UpdateNodes field.
func (r *mutationResolver) UpdateNodes(ctx context.Context, chanID uint64, input []*model.InputNode) (bool, error) {
	list := make([]*models.Node, 0)
	for _, node := range input {
		data := util.InputNode(node)
		list = append(list, data)
	}

	err := r.nodeService.UpdateAll(ctx, chanID, list)
	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteNode is the resolver for the DeleteNode field.
func (r *mutationResolver) DeleteNode(ctx context.Context, id uint64) (bool, error) {
	err := r.nodeService.Delete(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
