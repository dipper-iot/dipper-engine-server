// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dipper-iot/dipper-engine-server/ent/rulechan"
	"github.com/dipper-iot/dipper-engine-server/ent/rulenode"
)

// RuleNodeCreate is the builder for creating a RuleNode entity.
type RuleNodeCreate struct {
	config
	mutation *RuleNodeMutation
	hooks    []Hook
}

// SetChainID sets the "chain_id" field.
func (rnc *RuleNodeCreate) SetChainID(u uint64) *RuleNodeCreate {
	rnc.mutation.SetChainID(u)
	return rnc
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (rnc *RuleNodeCreate) SetNillableChainID(u *uint64) *RuleNodeCreate {
	if u != nil {
		rnc.SetChainID(*u)
	}
	return rnc
}

// SetNodeID sets the "node_id" field.
func (rnc *RuleNodeCreate) SetNodeID(s string) *RuleNodeCreate {
	rnc.mutation.SetNodeID(s)
	return rnc
}

// SetOption sets the "option" field.
func (rnc *RuleNodeCreate) SetOption(m map[string]interface{}) *RuleNodeCreate {
	rnc.mutation.SetOption(m)
	return rnc
}

// SetInfinite sets the "infinite" field.
func (rnc *RuleNodeCreate) SetInfinite(b bool) *RuleNodeCreate {
	rnc.mutation.SetInfinite(b)
	return rnc
}

// SetNillableInfinite sets the "infinite" field if the given value is not nil.
func (rnc *RuleNodeCreate) SetNillableInfinite(b *bool) *RuleNodeCreate {
	if b != nil {
		rnc.SetInfinite(*b)
	}
	return rnc
}

// SetDebug sets the "debug" field.
func (rnc *RuleNodeCreate) SetDebug(b bool) *RuleNodeCreate {
	rnc.mutation.SetDebug(b)
	return rnc
}

// SetNillableDebug sets the "debug" field if the given value is not nil.
func (rnc *RuleNodeCreate) SetNillableDebug(b *bool) *RuleNodeCreate {
	if b != nil {
		rnc.SetDebug(*b)
	}
	return rnc
}

// SetEnd sets the "end" field.
func (rnc *RuleNodeCreate) SetEnd(b bool) *RuleNodeCreate {
	rnc.mutation.SetEnd(b)
	return rnc
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (rnc *RuleNodeCreate) SetNillableEnd(b *bool) *RuleNodeCreate {
	if b != nil {
		rnc.SetEnd(*b)
	}
	return rnc
}

// SetCreatedAt sets the "created_at" field.
func (rnc *RuleNodeCreate) SetCreatedAt(t time.Time) *RuleNodeCreate {
	rnc.mutation.SetCreatedAt(t)
	return rnc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rnc *RuleNodeCreate) SetNillableCreatedAt(t *time.Time) *RuleNodeCreate {
	if t != nil {
		rnc.SetCreatedAt(*t)
	}
	return rnc
}

// SetUpdatedAt sets the "updated_at" field.
func (rnc *RuleNodeCreate) SetUpdatedAt(t time.Time) *RuleNodeCreate {
	rnc.mutation.SetUpdatedAt(t)
	return rnc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rnc *RuleNodeCreate) SetNillableUpdatedAt(t *time.Time) *RuleNodeCreate {
	if t != nil {
		rnc.SetUpdatedAt(*t)
	}
	return rnc
}

// SetID sets the "id" field.
func (rnc *RuleNodeCreate) SetID(u uint64) *RuleNodeCreate {
	rnc.mutation.SetID(u)
	return rnc
}

// SetChain sets the "chain" edge to the RuleChan entity.
func (rnc *RuleNodeCreate) SetChain(r *RuleChan) *RuleNodeCreate {
	return rnc.SetChainID(r.ID)
}

// Mutation returns the RuleNodeMutation object of the builder.
func (rnc *RuleNodeCreate) Mutation() *RuleNodeMutation {
	return rnc.mutation
}

// Save creates the RuleNode in the database.
func (rnc *RuleNodeCreate) Save(ctx context.Context) (*RuleNode, error) {
	var (
		err  error
		node *RuleNode
	)
	rnc.defaults()
	if len(rnc.hooks) == 0 {
		if err = rnc.check(); err != nil {
			return nil, err
		}
		node, err = rnc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RuleNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rnc.check(); err != nil {
				return nil, err
			}
			rnc.mutation = mutation
			if node, err = rnc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rnc.hooks) - 1; i >= 0; i-- {
			if rnc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rnc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rnc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RuleNode)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RuleNodeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rnc *RuleNodeCreate) SaveX(ctx context.Context) *RuleNode {
	v, err := rnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rnc *RuleNodeCreate) Exec(ctx context.Context) error {
	_, err := rnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rnc *RuleNodeCreate) ExecX(ctx context.Context) {
	if err := rnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rnc *RuleNodeCreate) defaults() {
	if _, ok := rnc.mutation.Option(); !ok {
		v := rulenode.DefaultOption
		rnc.mutation.SetOption(v)
	}
	if _, ok := rnc.mutation.Infinite(); !ok {
		v := rulenode.DefaultInfinite
		rnc.mutation.SetInfinite(v)
	}
	if _, ok := rnc.mutation.Debug(); !ok {
		v := rulenode.DefaultDebug
		rnc.mutation.SetDebug(v)
	}
	if _, ok := rnc.mutation.End(); !ok {
		v := rulenode.DefaultEnd
		rnc.mutation.SetEnd(v)
	}
	if _, ok := rnc.mutation.CreatedAt(); !ok {
		v := rulenode.DefaultCreatedAt()
		rnc.mutation.SetCreatedAt(v)
	}
	if _, ok := rnc.mutation.UpdatedAt(); !ok {
		v := rulenode.DefaultUpdatedAt()
		rnc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rnc *RuleNodeCreate) check() error {
	if _, ok := rnc.mutation.NodeID(); !ok {
		return &ValidationError{Name: "node_id", err: errors.New(`ent: missing required field "RuleNode.node_id"`)}
	}
	if v, ok := rnc.mutation.NodeID(); ok {
		if err := rulenode.NodeIDValidator(v); err != nil {
			return &ValidationError{Name: "node_id", err: fmt.Errorf(`ent: validator failed for field "RuleNode.node_id": %w`, err)}
		}
	}
	if _, ok := rnc.mutation.Option(); !ok {
		return &ValidationError{Name: "option", err: errors.New(`ent: missing required field "RuleNode.option"`)}
	}
	if _, ok := rnc.mutation.Infinite(); !ok {
		return &ValidationError{Name: "infinite", err: errors.New(`ent: missing required field "RuleNode.infinite"`)}
	}
	if _, ok := rnc.mutation.Debug(); !ok {
		return &ValidationError{Name: "debug", err: errors.New(`ent: missing required field "RuleNode.debug"`)}
	}
	if _, ok := rnc.mutation.End(); !ok {
		return &ValidationError{Name: "end", err: errors.New(`ent: missing required field "RuleNode.end"`)}
	}
	if _, ok := rnc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RuleNode.created_at"`)}
	}
	if _, ok := rnc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RuleNode.updated_at"`)}
	}
	return nil
}

func (rnc *RuleNodeCreate) sqlSave(ctx context.Context) (*RuleNode, error) {
	_node, _spec := rnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (rnc *RuleNodeCreate) createSpec() (*RuleNode, *sqlgraph.CreateSpec) {
	var (
		_node = &RuleNode{config: rnc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rulenode.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: rulenode.FieldID,
			},
		}
	)
	if id, ok := rnc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rnc.mutation.NodeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rulenode.FieldNodeID,
		})
		_node.NodeID = value
	}
	if value, ok := rnc.mutation.Option(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: rulenode.FieldOption,
		})
		_node.Option = value
	}
	if value, ok := rnc.mutation.Infinite(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rulenode.FieldInfinite,
		})
		_node.Infinite = value
	}
	if value, ok := rnc.mutation.Debug(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rulenode.FieldDebug,
		})
		_node.Debug = value
	}
	if value, ok := rnc.mutation.End(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rulenode.FieldEnd,
		})
		_node.End = value
	}
	if value, ok := rnc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rulenode.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rnc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rulenode.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := rnc.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rulenode.ChainTable,
			Columns: []string{rulenode.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: rulechan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ChainID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RuleNodeCreateBulk is the builder for creating many RuleNode entities in bulk.
type RuleNodeCreateBulk struct {
	config
	builders []*RuleNodeCreate
}

// Save creates the RuleNode entities in the database.
func (rncb *RuleNodeCreateBulk) Save(ctx context.Context) ([]*RuleNode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rncb.builders))
	nodes := make([]*RuleNode, len(rncb.builders))
	mutators := make([]Mutator, len(rncb.builders))
	for i := range rncb.builders {
		func(i int, root context.Context) {
			builder := rncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RuleNodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rncb *RuleNodeCreateBulk) SaveX(ctx context.Context) []*RuleNode {
	v, err := rncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rncb *RuleNodeCreateBulk) Exec(ctx context.Context) error {
	_, err := rncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rncb *RuleNodeCreateBulk) ExecX(ctx context.Context) {
	if err := rncb.Exec(ctx); err != nil {
		panic(err)
	}
}
