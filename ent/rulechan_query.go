// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dipper-iot/dipper-engine-server/ent/predicate"
	"github.com/dipper-iot/dipper-engine-server/ent/rulechan"
	"github.com/dipper-iot/dipper-engine-server/ent/rulenode"
	"github.com/dipper-iot/dipper-engine-server/ent/session"
)

// RuleChanQuery is the builder for querying RuleChan entities.
type RuleChanQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.RuleChan
	withRules    *RuleNodeQuery
	withSessions *SessionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RuleChanQuery builder.
func (rcq *RuleChanQuery) Where(ps ...predicate.RuleChan) *RuleChanQuery {
	rcq.predicates = append(rcq.predicates, ps...)
	return rcq
}

// Limit adds a limit step to the query.
func (rcq *RuleChanQuery) Limit(limit int) *RuleChanQuery {
	rcq.limit = &limit
	return rcq
}

// Offset adds an offset step to the query.
func (rcq *RuleChanQuery) Offset(offset int) *RuleChanQuery {
	rcq.offset = &offset
	return rcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rcq *RuleChanQuery) Unique(unique bool) *RuleChanQuery {
	rcq.unique = &unique
	return rcq
}

// Order adds an order step to the query.
func (rcq *RuleChanQuery) Order(o ...OrderFunc) *RuleChanQuery {
	rcq.order = append(rcq.order, o...)
	return rcq
}

// QueryRules chains the current query on the "rules" edge.
func (rcq *RuleChanQuery) QueryRules() *RuleNodeQuery {
	query := &RuleNodeQuery{config: rcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rulechan.Table, rulechan.FieldID, selector),
			sqlgraph.To(rulenode.Table, rulenode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, rulechan.RulesTable, rulechan.RulesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySessions chains the current query on the "sessions" edge.
func (rcq *RuleChanQuery) QuerySessions() *SessionQuery {
	query := &SessionQuery{config: rcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rulechan.Table, rulechan.FieldID, selector),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, rulechan.SessionsTable, rulechan.SessionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RuleChan entity from the query.
// Returns a *NotFoundError when no RuleChan was found.
func (rcq *RuleChanQuery) First(ctx context.Context) (*RuleChan, error) {
	nodes, err := rcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rulechan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rcq *RuleChanQuery) FirstX(ctx context.Context) *RuleChan {
	node, err := rcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RuleChan ID from the query.
// Returns a *NotFoundError when no RuleChan ID was found.
func (rcq *RuleChanQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = rcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rulechan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rcq *RuleChanQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := rcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RuleChan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RuleChan entity is found.
// Returns a *NotFoundError when no RuleChan entities are found.
func (rcq *RuleChanQuery) Only(ctx context.Context) (*RuleChan, error) {
	nodes, err := rcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rulechan.Label}
	default:
		return nil, &NotSingularError{rulechan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rcq *RuleChanQuery) OnlyX(ctx context.Context) *RuleChan {
	node, err := rcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RuleChan ID in the query.
// Returns a *NotSingularError when more than one RuleChan ID is found.
// Returns a *NotFoundError when no entities are found.
func (rcq *RuleChanQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = rcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rulechan.Label}
	default:
		err = &NotSingularError{rulechan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rcq *RuleChanQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := rcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RuleChans.
func (rcq *RuleChanQuery) All(ctx context.Context) ([]*RuleChan, error) {
	if err := rcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rcq *RuleChanQuery) AllX(ctx context.Context) []*RuleChan {
	nodes, err := rcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RuleChan IDs.
func (rcq *RuleChanQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := rcq.Select(rulechan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rcq *RuleChanQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := rcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rcq *RuleChanQuery) Count(ctx context.Context) (int, error) {
	if err := rcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rcq *RuleChanQuery) CountX(ctx context.Context) int {
	count, err := rcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rcq *RuleChanQuery) Exist(ctx context.Context) (bool, error) {
	if err := rcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rcq *RuleChanQuery) ExistX(ctx context.Context) bool {
	exist, err := rcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RuleChanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rcq *RuleChanQuery) Clone() *RuleChanQuery {
	if rcq == nil {
		return nil
	}
	return &RuleChanQuery{
		config:       rcq.config,
		limit:        rcq.limit,
		offset:       rcq.offset,
		order:        append([]OrderFunc{}, rcq.order...),
		predicates:   append([]predicate.RuleChan{}, rcq.predicates...),
		withRules:    rcq.withRules.Clone(),
		withSessions: rcq.withSessions.Clone(),
		// clone intermediate query.
		sql:    rcq.sql.Clone(),
		path:   rcq.path,
		unique: rcq.unique,
	}
}

// WithRules tells the query-builder to eager-load the nodes that are connected to
// the "rules" edge. The optional arguments are used to configure the query builder of the edge.
func (rcq *RuleChanQuery) WithRules(opts ...func(*RuleNodeQuery)) *RuleChanQuery {
	query := &RuleNodeQuery{config: rcq.config}
	for _, opt := range opts {
		opt(query)
	}
	rcq.withRules = query
	return rcq
}

// WithSessions tells the query-builder to eager-load the nodes that are connected to
// the "sessions" edge. The optional arguments are used to configure the query builder of the edge.
func (rcq *RuleChanQuery) WithSessions(opts ...func(*SessionQuery)) *RuleChanQuery {
	query := &SessionQuery{config: rcq.config}
	for _, opt := range opts {
		opt(query)
	}
	rcq.withSessions = query
	return rcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RuleChan.Query().
//		GroupBy(rulechan.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rcq *RuleChanQuery) GroupBy(field string, fields ...string) *RuleChanGroupBy {
	grbuild := &RuleChanGroupBy{config: rcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rcq.sqlQuery(ctx), nil
	}
	grbuild.label = rulechan.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.RuleChan.Query().
//		Select(rulechan.FieldName).
//		Scan(ctx, &v)
func (rcq *RuleChanQuery) Select(fields ...string) *RuleChanSelect {
	rcq.fields = append(rcq.fields, fields...)
	selbuild := &RuleChanSelect{RuleChanQuery: rcq}
	selbuild.label = rulechan.Label
	selbuild.flds, selbuild.scan = &rcq.fields, selbuild.Scan
	return selbuild
}

func (rcq *RuleChanQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rcq.fields {
		if !rulechan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rcq.path != nil {
		prev, err := rcq.path(ctx)
		if err != nil {
			return err
		}
		rcq.sql = prev
	}
	return nil
}

func (rcq *RuleChanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RuleChan, error) {
	var (
		nodes       = []*RuleChan{}
		_spec       = rcq.querySpec()
		loadedTypes = [2]bool{
			rcq.withRules != nil,
			rcq.withSessions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RuleChan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RuleChan{config: rcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rcq.withRules; query != nil {
		if err := rcq.loadRules(ctx, query, nodes,
			func(n *RuleChan) { n.Edges.Rules = []*RuleNode{} },
			func(n *RuleChan, e *RuleNode) { n.Edges.Rules = append(n.Edges.Rules, e) }); err != nil {
			return nil, err
		}
	}
	if query := rcq.withSessions; query != nil {
		if err := rcq.loadSessions(ctx, query, nodes,
			func(n *RuleChan) { n.Edges.Sessions = []*Session{} },
			func(n *RuleChan, e *Session) { n.Edges.Sessions = append(n.Edges.Sessions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rcq *RuleChanQuery) loadRules(ctx context.Context, query *RuleNodeQuery, nodes []*RuleChan, init func(*RuleChan), assign func(*RuleChan, *RuleNode)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*RuleChan)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.RuleNode(func(s *sql.Selector) {
		s.Where(sql.InValues(rulechan.RulesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ChainID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chain_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rcq *RuleChanQuery) loadSessions(ctx context.Context, query *SessionQuery, nodes []*RuleChan, init func(*RuleChan), assign func(*RuleChan, *Session)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*RuleChan)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Session(func(s *sql.Selector) {
		s.Where(sql.InValues(rulechan.SessionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ChainID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chain_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rcq *RuleChanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rcq.querySpec()
	_spec.Node.Columns = rcq.fields
	if len(rcq.fields) > 0 {
		_spec.Unique = rcq.unique != nil && *rcq.unique
	}
	return sqlgraph.CountNodes(ctx, rcq.driver, _spec)
}

func (rcq *RuleChanQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := rcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (rcq *RuleChanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rulechan.Table,
			Columns: rulechan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: rulechan.FieldID,
			},
		},
		From:   rcq.sql,
		Unique: true,
	}
	if unique := rcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rulechan.FieldID)
		for i := range fields {
			if fields[i] != rulechan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rcq *RuleChanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rcq.driver.Dialect())
	t1 := builder.Table(rulechan.Table)
	columns := rcq.fields
	if len(columns) == 0 {
		columns = rulechan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rcq.sql != nil {
		selector = rcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rcq.unique != nil && *rcq.unique {
		selector.Distinct()
	}
	for _, p := range rcq.predicates {
		p(selector)
	}
	for _, p := range rcq.order {
		p(selector)
	}
	if offset := rcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RuleChanGroupBy is the group-by builder for RuleChan entities.
type RuleChanGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rcgb *RuleChanGroupBy) Aggregate(fns ...AggregateFunc) *RuleChanGroupBy {
	rcgb.fns = append(rcgb.fns, fns...)
	return rcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rcgb *RuleChanGroupBy) Scan(ctx context.Context, v any) error {
	query, err := rcgb.path(ctx)
	if err != nil {
		return err
	}
	rcgb.sql = query
	return rcgb.sqlScan(ctx, v)
}

func (rcgb *RuleChanGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range rcgb.fields {
		if !rulechan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rcgb *RuleChanGroupBy) sqlQuery() *sql.Selector {
	selector := rcgb.sql.Select()
	aggregation := make([]string, 0, len(rcgb.fns))
	for _, fn := range rcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rcgb.fields)+len(rcgb.fns))
		for _, f := range rcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rcgb.fields...)...)
}

// RuleChanSelect is the builder for selecting fields of RuleChan entities.
type RuleChanSelect struct {
	*RuleChanQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rcs *RuleChanSelect) Scan(ctx context.Context, v any) error {
	if err := rcs.prepareQuery(ctx); err != nil {
		return err
	}
	rcs.sql = rcs.RuleChanQuery.sqlQuery(ctx)
	return rcs.sqlScan(ctx, v)
}

func (rcs *RuleChanSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := rcs.sql.Query()
	if err := rcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
