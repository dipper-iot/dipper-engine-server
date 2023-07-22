package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dipper-iot/dipper-engine-server/graph/generated"
	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/dipper-iot/dipper-engine-server/graph/util"
)

// Debug is the resolver for the debug field.
func (r *subscriptionResolver) Debug(ctx context.Context, input *model.DebugFilter) (<-chan []*model.SessionOutput, error) {
	chanLog := make(chan []*model.SessionOutput)
	go util.QueryLog(ctx, r.busData, input, chanLog)
	return chanLog, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
