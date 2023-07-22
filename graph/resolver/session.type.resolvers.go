package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/graph/generated"
)

// ChanID is the resolver for the chan_id field.
func (r *sessionInfoResolver) ChanID(ctx context.Context, obj *ent.Session) (uint64, error) {
	return obj.ChainID, nil
}

// Infinity is the resolver for the infinity field.
func (r *sessionInfoResolver) Infinity(ctx context.Context, obj *ent.Session) (bool, error) {
	return obj.Infinite, nil
}

// Chan is the resolver for the chan field.
func (r *sessionInfoResolver) Chan(ctx context.Context, obj *ent.Session) (*ent.RuleChan, error) {
	chanData, err := r.chanService.Get(ctx, obj.ChainID)
	if err != nil {
		return nil, err
	}
	return chanData, nil
}

// SessionInfo returns generated.SessionInfoResolver implementation.
func (r *Resolver) SessionInfo() generated.SessionInfoResolver { return &sessionInfoResolver{r} }

type sessionInfoResolver struct{ *Resolver }
