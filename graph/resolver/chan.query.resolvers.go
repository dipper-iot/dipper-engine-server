package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/dipper-iot/dipper-engine-server/models"
)

// ListChan is the resolver for the ListChan field.
func (r *queryResolver) ListChan(ctx context.Context, input models.ListChanRequest) (*model.ListChan, error) {
	panic(fmt.Errorf("not implemented: ListChan - ListChan"))
}

// GetChan is the resolver for the GetChan field.
func (r *queryResolver) GetChan(ctx context.Context, id uint64) (*ent.RuleChan, error) {
	panic(fmt.Errorf("not implemented: GetChan - GetChan"))
}
