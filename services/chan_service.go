package services

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/database"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/ent/rulechan"
	"github.com/dipper-iot/dipper-engine-server/models"
	"sync"
)

type ChanServiceImpl struct {
	clientRead  *ent.Client
	clientWrite *ent.Client
}

func NewChanServiceImpl(
	clientRead *ent.Client,
	clientWrite *ent.Client,
) *ChanServiceImpl {
	return &ChanServiceImpl{
		clientRead:  clientRead,
		clientWrite: clientWrite,
	}
}

func (c ChanServiceImpl) List(ctx context.Context, req *models.ListChanRequest) (list []*ent.RuleChan, total int, err error) {
	query := c.clientRead.RuleChan.Query()

	wg := sync.WaitGroup{}
	wg.Add(2)

	var (
		errTotal error
		errList  error
	)

	go func() {
		total, errTotal = query.Clone().Count(ctx)
		wg.Done()
	}()

	go func() {
		list, errList = query.Clone().Limit(req.Limit).Offset(req.Skip).All(ctx)
		wg.Done()
	}()

	wg.Wait()

	if errTotal != nil {
		err = errTotal
		return
	}
	err = errList

	return
}

func (c ChanServiceImpl) Get(ctx context.Context, id uint64) (detail *ent.RuleChan, err error) {
	detail, err = c.clientRead.RuleChan.Get(ctx, id)
	return
}

func (c ChanServiceImpl) Create(ctx context.Context, data *models.Chan) (err error) {
	query := c.clientWrite.RuleChan.Create()

	id, err := database.NextID()
	if err != nil {
		return err
	}

	query.SetID(id)
	query.SetName(data.Name)
	if data.NodeRoot == "" {
		query.SetRootNode("1")
	} else {
		query.SetRootNode(data.NodeRoot)
	}
	query.SetDescription(data.Description)
	query.SetStatus(rulechan.StatusActivated)

	_, err = query.Save(ctx)
	return
}

func (c ChanServiceImpl) Update(ctx context.Context, id uint64, data *models.Chan) (err error) {
	query := c.clientWrite.RuleChan.Update()

	query.SetName(data.Name)
	query.SetDescription(data.Description)

	_, err = query.Where(rulechan.IDEQ(id)).Save(ctx)
	return
}
func (c ChanServiceImpl) UpdateRootNode(ctx context.Context, id uint64, rootNode string) (err error) {
	query := c.clientWrite.RuleChan.Update()
	query.SetRootNode(rootNode)
	_, err = query.Where(rulechan.IDEQ(id)).Save(ctx)
	return
}

func (c ChanServiceImpl) UpdateStatus(ctx context.Context, id uint64, status rulechan.Status) (err error) {
	query := c.clientWrite.RuleChan.Update()
	query.SetStatus(status)
	_, err = query.Where(rulechan.IDEQ(id)).Save(ctx)
	return
}

func (c ChanServiceImpl) Delete(ctx context.Context, id uint64) (err error) {
	query := c.clientWrite.RuleChan.Delete()
	_, err = query.Where(rulechan.IDEQ(id)).Exec(ctx)
	return
}
