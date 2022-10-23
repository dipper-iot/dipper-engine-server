// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dipper-iot/dipper-engine-server/ent/predicate"
	"github.com/dipper-iot/dipper-engine-server/ent/rulechan"
	"github.com/dipper-iot/dipper-engine-server/ent/session"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetChainID sets the "chain_id" field.
func (su *SessionUpdate) SetChainID(u uint64) *SessionUpdate {
	su.mutation.SetChainID(u)
	return su
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableChainID(u *uint64) *SessionUpdate {
	if u != nil {
		su.SetChainID(*u)
	}
	return su
}

// ClearChainID clears the value of the "chain_id" field.
func (su *SessionUpdate) ClearChainID() *SessionUpdate {
	su.mutation.ClearChainID()
	return su
}

// SetInfinite sets the "infinite" field.
func (su *SessionUpdate) SetInfinite(b bool) *SessionUpdate {
	su.mutation.SetInfinite(b)
	return su
}

// SetNillableInfinite sets the "infinite" field if the given value is not nil.
func (su *SessionUpdate) SetNillableInfinite(b *bool) *SessionUpdate {
	if b != nil {
		su.SetInfinite(*b)
	}
	return su
}

// SetData sets the "data" field.
func (su *SessionUpdate) SetData(m map[string]interface{}) *SessionUpdate {
	su.mutation.SetData(m)
	return su
}

// SetResult sets the "result" field.
func (su *SessionUpdate) SetResult(m map[string]interface{}) *SessionUpdate {
	su.mutation.SetResult(m)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SessionUpdate) SetCreatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableCreatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SessionUpdate) SetUpdatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUpdatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetUpdatedAt(*t)
	}
	return su
}

// SetChain sets the "chain" edge to the RuleChan entity.
func (su *SessionUpdate) SetChain(r *RuleChan) *SessionUpdate {
	return su.SetChainID(r.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearChain clears the "chain" edge to the RuleChan entity.
func (su *SessionUpdate) ClearChain() *SessionUpdate {
	su.mutation.ClearChain()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: session.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Infinite(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldInfinite,
		})
	}
	if value, ok := su.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: session.FieldData,
		})
	}
	if value, ok := su.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: session.FieldResult,
		})
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldUpdatedAt,
		})
	}
	if su.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.ChainTable,
			Columns: []string{session.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: rulechan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.ChainTable,
			Columns: []string{session.ChainColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetChainID sets the "chain_id" field.
func (suo *SessionUpdateOne) SetChainID(u uint64) *SessionUpdateOne {
	suo.mutation.SetChainID(u)
	return suo
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableChainID(u *uint64) *SessionUpdateOne {
	if u != nil {
		suo.SetChainID(*u)
	}
	return suo
}

// ClearChainID clears the value of the "chain_id" field.
func (suo *SessionUpdateOne) ClearChainID() *SessionUpdateOne {
	suo.mutation.ClearChainID()
	return suo
}

// SetInfinite sets the "infinite" field.
func (suo *SessionUpdateOne) SetInfinite(b bool) *SessionUpdateOne {
	suo.mutation.SetInfinite(b)
	return suo
}

// SetNillableInfinite sets the "infinite" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableInfinite(b *bool) *SessionUpdateOne {
	if b != nil {
		suo.SetInfinite(*b)
	}
	return suo
}

// SetData sets the "data" field.
func (suo *SessionUpdateOne) SetData(m map[string]interface{}) *SessionUpdateOne {
	suo.mutation.SetData(m)
	return suo
}

// SetResult sets the "result" field.
func (suo *SessionUpdateOne) SetResult(m map[string]interface{}) *SessionUpdateOne {
	suo.mutation.SetResult(m)
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SessionUpdateOne) SetCreatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableCreatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SessionUpdateOne) SetUpdatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUpdatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetUpdatedAt(*t)
	}
	return suo
}

// SetChain sets the "chain" edge to the RuleChan entity.
func (suo *SessionUpdateOne) SetChain(r *RuleChan) *SessionUpdateOne {
	return suo.SetChainID(r.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearChain clears the "chain" edge to the RuleChan entity.
func (suo *SessionUpdateOne) ClearChain() *SessionUpdateOne {
	suo.mutation.ClearChain()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Session)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SessionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: session.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Infinite(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldInfinite,
		})
	}
	if value, ok := suo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: session.FieldData,
		})
	}
	if value, ok := suo.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: session.FieldResult,
		})
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldUpdatedAt,
		})
	}
	if suo.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.ChainTable,
			Columns: []string{session.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: rulechan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.ChainTable,
			Columns: []string{session.ChainColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
