package services

import (
	"context"
	"github.com/dipper-iot/dipper-engine-server/database"
	"github.com/dipper-iot/dipper-engine-server/ent"
	"github.com/dipper-iot/dipper-engine-server/ent/rulenode"
	"github.com/dipper-iot/dipper-engine-server/models"
	"log"
)

type NodeServiceImpl struct {
	clientRead  *ent.Client
	clientWrite *ent.Client
}

func NewNodeServiceImpl(clientRead *ent.Client, clientWrite *ent.Client) *NodeServiceImpl {
	return &NodeServiceImpl{clientRead: clientRead, clientWrite: clientWrite}
}

func (n NodeServiceImpl) All(ctx context.Context, chanId uint64) (list []*ent.RuleNode, err error) {
	list, err = n.clientRead.RuleNode.Query().Where(rulenode.ChainID(chanId)).All(ctx)
	return
}

func (n NodeServiceImpl) Get(ctx context.Context, id uint64) (detail *ent.RuleNode, err error) {
	detail, err = n.clientRead.RuleNode.Query().Where(rulenode.ID(id)).First(ctx)
	return
}

func (n NodeServiceImpl) createNode(chanId uint64, node *models.Node) (create *ent.RuleNodeCreate, err error) {
	id, err := database.NextID()
	if err != nil {
		return nil, err
	}

	create = n.clientWrite.RuleNode.Create()
	create.SetID(id)
	create.SetChainID(chanId)
	create.SetNodeID(node.NodeID)
	create.SetRuleID(node.RuleID)
	create.SetDebug(node.Debug)
	create.SetEnd(node.End)
	create.SetOption(node.Option)

	return create, nil
}

func (n NodeServiceImpl) UpdateAll(ctx context.Context, chanId uint64, list []*models.Node) (err error) {
	var (
		count int
	)

	count, err = n.clientWrite.RuleNode.Query().Where(rulenode.ChainID(chanId)).Count(ctx)
	if err != nil {
		return
	}
	if count == 0 {

		listCreate := make([]*ent.RuleNodeCreate, 0)

		for _, node := range list {
			create, err := n.createNode(chanId, node)
			if err != nil {
				return err
			}
			listCreate = append(listCreate, create)
		}

		_, err = n.clientWrite.RuleNode.CreateBulk(listCreate...).Save(ctx)

	} else {

		tx, err := n.clientWrite.Tx(ctx)
		if err != nil {
			return err
		}

		var (
			listExits    []*ent.RuleNode
			mapListExits = map[string]*ent.RuleNode{}
			mapList      = map[string]*models.Node{}
		)
		listExits, err = n.All(ctx, chanId)

		listCreate := make([]*ent.RuleNodeCreate, 0)
		listDelete := make([]uint64, 0)

		for _, node := range list {
			mapList[node.NodeID] = node
		}

		for _, node := range listExits {
			_, exits := mapList[node.NodeID]
			if exits {
				mapListExits[node.NodeID] = node
			} else {
				listDelete = append(listDelete, node.ID)
			}
		}

		for _, node := range list {
			_, exits := mapListExits[node.NodeID]
			if exits {
				// update
				q := n.clientWrite.RuleNode.Update()
				q.SetDebug(node.Debug)
				q.SetEnd(node.End)
				q.SetOption(node.Option)
				q.SetRuleID(node.RuleID)

				err = q.Where(rulenode.ChainID(chanId)).Where(rulenode.NodeID(node.NodeID)).Exec(ctx)
				if err != nil {
					return err
				}

			} else {
				create, err := n.createNode(chanId, node)
				if err != nil {
					e := tx.Rollback()
					if e != nil {
						log.Print(e)
					}
					return err
				}
				listCreate = append(listCreate, create)
			}
		}

		// create
		if len(listCreate) > 0 {
			_, err = tx.RuleNode.CreateBulk(listCreate...).Save(ctx)
			if err != nil {
				e := tx.Rollback()
				if e != nil {
					log.Print(e)
				}
				return err
			}
		}

		// delete
		if len(listDelete) > 0 {
			_, err = tx.RuleNode.Delete().Where(rulenode.IDIn(listDelete...)).Exec(ctx)
			if err != nil {
				e := tx.Rollback()
				if e != nil {
					log.Print(e)
				}
				return err
			}
		}

		return tx.Commit()
	}

	if err != nil {
		return
	}

	return
}

func (n NodeServiceImpl) Update(ctx context.Context, chanId uint64, data *models.Node) (err error) {
	var count int
	count, err = n.clientWrite.RuleNode.Query().
		Where(rulenode.ChainID(chanId)).Where(rulenode.NodeID(data.NodeID)).Count(ctx)
	if err != nil {

		return err
	}

	if count == 0 {
		create, err := n.createNode(chanId, data)
		if err != nil {
			return err
		}
		_, err = create.Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	q := n.clientWrite.RuleNode.Update()
	q.SetDebug(data.Debug)
	q.SetEnd(data.End)
	q.SetOption(data.Option)
	q.SetRuleID(data.RuleID)

	err = q.Where(rulenode.ChainID(chanId)).Where(rulenode.NodeID(data.NodeID)).Exec(ctx)
	if err != nil {
		return
	}

	if err != nil {
		return
	}
	return
}

func (n NodeServiceImpl) Delete(ctx context.Context, id uint64) (err error) {
	query := n.clientWrite.RuleNode.Delete()
	_, err = query.Where(rulenode.ID(id)).Exec(ctx)
	return
}
