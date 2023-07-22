package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/dipper-iot/dipper-engine-server/graph/model"
	"github.com/dipper-iot/dipper-engine-server/models"
	"github.com/sirupsen/logrus"
)

// AddSession is the resolver for the AddSession field.
func (r *mutationResolver) AddSession(ctx context.Context, input model.SessionRequest) (bool, error) {
	chanId, err := strconv.ParseUint(input.ChanID, 10, 64)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	err = r.sessionService.RunSession(ctx, &models.CreateSession{
		Data:   input.Data,
		ChanID: chanId,
		IsTest: input.IsTest,
	})
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return true, nil
}

// StopSession is the resolver for the StopSession field.
func (r *mutationResolver) StopSession(ctx context.Context, id uint64) (bool, error) {
	err := r.manageSessionService.Stop(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
