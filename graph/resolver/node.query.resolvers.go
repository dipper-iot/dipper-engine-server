package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dipper-iot/dipper-engine-server/ent"
)

// GetNode is the resolver for the GetNode field.
func (r *queryResolver) GetNode(ctx context.Context, id uint64) (*ent.RuleNode, error) {
	return r.nodeService.Get(ctx, id)
}
