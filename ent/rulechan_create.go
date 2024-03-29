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
	"github.com/dipper-iot/dipper-engine-server/ent/session"
)

// RuleChanCreate is the builder for creating a RuleChan entity.
type RuleChanCreate struct {
	config
	mutation *RuleChanMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rcc *RuleChanCreate) SetName(s string) *RuleChanCreate {
	rcc.mutation.SetName(s)
	return rcc
}

// SetDescription sets the "description" field.
func (rcc *RuleChanCreate) SetDescription(s string) *RuleChanCreate {
	rcc.mutation.SetDescription(s)
	return rcc
}

// SetRootNode sets the "root_node" field.
func (rcc *RuleChanCreate) SetRootNode(s string) *RuleChanCreate {
	rcc.mutation.SetRootNode(s)
	return rcc
}

// SetInfinite sets the "infinite" field.
func (rcc *RuleChanCreate) SetInfinite(b bool) *RuleChanCreate {
	rcc.mutation.SetInfinite(b)
	return rcc
}

// SetNillableInfinite sets the "infinite" field if the given value is not nil.
func (rcc *RuleChanCreate) SetNillableInfinite(b *bool) *RuleChanCreate {
	if b != nil {
		rcc.SetInfinite(*b)
	}
	return rcc
}

// SetStatus sets the "status" field.
func (rcc *RuleChanCreate) SetStatus(r rulechan.Status) *RuleChanCreate {
	rcc.mutation.SetStatus(r)
	return rcc
}

// SetCreatedAt sets the "created_at" field.
func (rcc *RuleChanCreate) SetCreatedAt(t time.Time) *RuleChanCreate {
	rcc.mutation.SetCreatedAt(t)
	return rcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rcc *RuleChanCreate) SetNillableCreatedAt(t *time.Time) *RuleChanCreate {
	if t != nil {
		rcc.SetCreatedAt(*t)
	}
	return rcc
}

// SetUpdatedAt sets the "updated_at" field.
func (rcc *RuleChanCreate) SetUpdatedAt(t time.Time) *RuleChanCreate {
	rcc.mutation.SetUpdatedAt(t)
	return rcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rcc *RuleChanCreate) SetNillableUpdatedAt(t *time.Time) *RuleChanCreate {
	if t != nil {
		rcc.SetUpdatedAt(*t)
	}
	return rcc
}

// SetID sets the "id" field.
func (rcc *RuleChanCreate) SetID(u uint64) *RuleChanCreate {
	rcc.mutation.SetID(u)
	return rcc
}

// AddRuleIDs adds the "rules" edge to the RuleNode entity by IDs.
func (rcc *RuleChanCreate) AddRuleIDs(ids ...uint64) *RuleChanCreate {
	rcc.mutation.AddRuleIDs(ids...)
	return rcc
}

// AddRules adds the "rules" edges to the RuleNode entity.
func (rcc *RuleChanCreate) AddRules(r ...*RuleNode) *RuleChanCreate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcc.AddRuleIDs(ids...)
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (rcc *RuleChanCreate) AddSessionIDs(ids ...uint64) *RuleChanCreate {
	rcc.mutation.AddSessionIDs(ids...)
	return rcc
}

// AddSessions adds the "sessions" edges to the Session entity.
func (rcc *RuleChanCreate) AddSessions(s ...*Session) *RuleChanCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return rcc.AddSessionIDs(ids...)
}

// Mutation returns the RuleChanMutation object of the builder.
func (rcc *RuleChanCreate) Mutation() *RuleChanMutation {
	return rcc.mutation
}

// Save creates the RuleChan in the database.
func (rcc *RuleChanCreate) Save(ctx context.Context) (*RuleChan, error) {
	var (
		err  error
		node *RuleChan
	)
	rcc.defaults()
	if len(rcc.hooks) == 0 {
		if err = rcc.check(); err != nil {
			return nil, err
		}
		node, err = rcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RuleChanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rcc.check(); err != nil {
				return nil, err
			}
			rcc.mutation = mutation
			if node, err = rcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rcc.hooks) - 1; i >= 0; i-- {
			if rcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rcc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rcc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RuleChan)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RuleChanMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rcc *RuleChanCreate) SaveX(ctx context.Context) *RuleChan {
	v, err := rcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcc *RuleChanCreate) Exec(ctx context.Context) error {
	_, err := rcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcc *RuleChanCreate) ExecX(ctx context.Context) {
	if err := rcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcc *RuleChanCreate) defaults() {
	if _, ok := rcc.mutation.Infinite(); !ok {
		v := rulechan.DefaultInfinite
		rcc.mutation.SetInfinite(v)
	}
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		v := rulechan.DefaultCreatedAt()
		rcc.mutation.SetCreatedAt(v)
	}
	if _, ok := rcc.mutation.UpdatedAt(); !ok {
		v := rulechan.DefaultUpdatedAt()
		rcc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rcc *RuleChanCreate) check() error {
	if _, ok := rcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "RuleChan.name"`)}
	}
	if v, ok := rcc.mutation.Name(); ok {
		if err := rulechan.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "RuleChan.name": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "RuleChan.description"`)}
	}
	if v, ok := rcc.mutation.Description(); ok {
		if err := rulechan.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "RuleChan.description": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.RootNode(); !ok {
		return &ValidationError{Name: "root_node", err: errors.New(`ent: missing required field "RuleChan.root_node"`)}
	}
	if v, ok := rcc.mutation.RootNode(); ok {
		if err := rulechan.RootNodeValidator(v); err != nil {
			return &ValidationError{Name: "root_node", err: fmt.Errorf(`ent: validator failed for field "RuleChan.root_node": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.Infinite(); !ok {
		return &ValidationError{Name: "infinite", err: errors.New(`ent: missing required field "RuleChan.infinite"`)}
	}
	if _, ok := rcc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "RuleChan.status"`)}
	}
	if v, ok := rcc.mutation.Status(); ok {
		if err := rulechan.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "RuleChan.status": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RuleChan.created_at"`)}
	}
	if _, ok := rcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RuleChan.updated_at"`)}
	}
	return nil
}

func (rcc *RuleChanCreate) sqlSave(ctx context.Context) (*RuleChan, error) {
	_node, _spec := rcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rcc.driver, _spec); err != nil {
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

func (rcc *RuleChanCreate) createSpec() (*RuleChan, *sqlgraph.CreateSpec) {
	var (
		_node = &RuleChan{config: rcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rulechan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: rulechan.FieldID,
			},
		}
	)
	if id, ok := rcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rcc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rulechan.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rcc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rulechan.FieldDescription,
		})
		_node.Description = &value
	}
	if value, ok := rcc.mutation.RootNode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rulechan.FieldRootNode,
		})
		_node.RootNode = value
	}
	if value, ok := rcc.mutation.Infinite(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rulechan.FieldInfinite,
		})
		_node.Infinite = value
	}
	if value, ok := rcc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: rulechan.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := rcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rulechan.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rulechan.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := rcc.mutation.RulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rulechan.RulesTable,
			Columns: []string{rulechan.RulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: rulenode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rcc.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rulechan.SessionsTable,
			Columns: []string{rulechan.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RuleChanCreateBulk is the builder for creating many RuleChan entities in bulk.
type RuleChanCreateBulk struct {
	config
	builders []*RuleChanCreate
}

// Save creates the RuleChan entities in the database.
func (rccb *RuleChanCreateBulk) Save(ctx context.Context) ([]*RuleChan, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rccb.builders))
	nodes := make([]*RuleChan, len(rccb.builders))
	mutators := make([]Mutator, len(rccb.builders))
	for i := range rccb.builders {
		func(i int, root context.Context) {
			builder := rccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RuleChanMutation)
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
					_, err = mutators[i+1].Mutate(root, rccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rccb *RuleChanCreateBulk) SaveX(ctx context.Context) []*RuleChan {
	v, err := rccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rccb *RuleChanCreateBulk) Exec(ctx context.Context) error {
	_, err := rccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rccb *RuleChanCreateBulk) ExecX(ctx context.Context) {
	if err := rccb.Exec(ctx); err != nil {
		panic(err)
	}
}
