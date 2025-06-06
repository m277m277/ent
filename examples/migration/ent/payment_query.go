// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/migration/ent/card"
	"entgo.io/ent/examples/migration/ent/payment"
	"entgo.io/ent/examples/migration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// PaymentQuery is the builder for querying Payment entities.
type PaymentQuery struct {
	config
	ctx        *QueryContext
	order      []payment.OrderOption
	inters     []Interceptor
	predicates []predicate.Payment
	withCard   *CardQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PaymentQuery builder.
func (_q *PaymentQuery) Where(ps ...predicate.Payment) *PaymentQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *PaymentQuery) Limit(limit int) *PaymentQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *PaymentQuery) Offset(offset int) *PaymentQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *PaymentQuery) Unique(unique bool) *PaymentQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *PaymentQuery) Order(o ...payment.OrderOption) *PaymentQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryCard chains the current query on the "card" edge.
func (_q *PaymentQuery) QueryCard() *CardQuery {
	query := (&CardClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payment.Table, payment.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payment.CardTable, payment.CardColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Payment entity from the query.
// Returns a *NotFoundError when no Payment was found.
func (_q *PaymentQuery) First(ctx context.Context) (*Payment, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{payment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *PaymentQuery) FirstX(ctx context.Context) *Payment {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Payment ID from the query.
// Returns a *NotFoundError when no Payment ID was found.
func (_q *PaymentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{payment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *PaymentQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Payment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Payment entity is found.
// Returns a *NotFoundError when no Payment entities are found.
func (_q *PaymentQuery) Only(ctx context.Context) (*Payment, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{payment.Label}
	default:
		return nil, &NotSingularError{payment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *PaymentQuery) OnlyX(ctx context.Context) *Payment {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Payment ID in the query.
// Returns a *NotSingularError when more than one Payment ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *PaymentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{payment.Label}
	default:
		err = &NotSingularError{payment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *PaymentQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Payments.
func (_q *PaymentQuery) All(ctx context.Context) ([]*Payment, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Payment, *PaymentQuery]()
	return withInterceptors[[]*Payment](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *PaymentQuery) AllX(ctx context.Context) []*Payment {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Payment IDs.
func (_q *PaymentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(payment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *PaymentQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *PaymentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*PaymentQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *PaymentQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *PaymentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *PaymentQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PaymentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *PaymentQuery) Clone() *PaymentQuery {
	if _q == nil {
		return nil
	}
	return &PaymentQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]payment.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.Payment{}, _q.predicates...),
		withCard:   _q.withCard.Clone(),
		// clone intermediate query.
		sql:  _q.sql.Clone(),
		path: _q.path,
	}
}

// WithCard tells the query-builder to eager-load the nodes that are connected to
// the "card" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *PaymentQuery) WithCard(opts ...func(*CardQuery)) *PaymentQuery {
	query := (&CardClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withCard = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CardID int `json:"card_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Payment.Query().
//		GroupBy(payment.FieldCardID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *PaymentQuery) GroupBy(field string, fields ...string) *PaymentGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PaymentGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = payment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CardID int `json:"card_id,omitempty"`
//	}
//
//	client.Payment.Query().
//		Select(payment.FieldCardID).
//		Scan(ctx, &v)
func (_q *PaymentQuery) Select(fields ...string) *PaymentSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &PaymentSelect{PaymentQuery: _q}
	sbuild.label = payment.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PaymentSelect configured with the given aggregations.
func (_q *PaymentQuery) Aggregate(fns ...AggregateFunc) *PaymentSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *PaymentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range _q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, _q); err != nil {
				return err
			}
		}
	}
	for _, f := range _q.ctx.Fields {
		if !payment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.sql = prev
	}
	return nil
}

func (_q *PaymentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Payment, error) {
	var (
		nodes       = []*Payment{}
		_spec       = _q.querySpec()
		loadedTypes = [1]bool{
			_q.withCard != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Payment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Payment{config: _q.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, _q.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := _q.withCard; query != nil {
		if err := _q.loadCard(ctx, query, nodes, nil,
			func(n *Payment, e *Card) { n.Edges.Card = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (_q *PaymentQuery) loadCard(ctx context.Context, query *CardQuery, nodes []*Payment, init func(*Payment), assign func(*Payment, *Card)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Payment)
	for i := range nodes {
		fk := nodes[i].CardID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(card.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "card_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (_q *PaymentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *PaymentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for i := range fields {
			if fields[i] != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if _q.withCard != nil {
			_spec.Node.AddColumnOnce(payment.FieldCardID)
		}
	}
	if ps := _q.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := _q.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := _q.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := _q.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (_q *PaymentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(payment.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = payment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if _q.sql != nil {
		selector = _q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if _q.ctx.Unique != nil && *_q.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range _q.predicates {
		p(selector)
	}
	for _, p := range _q.order {
		p(selector)
	}
	if offset := _q.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := _q.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PaymentGroupBy is the group-by builder for Payment entities.
type PaymentGroupBy struct {
	selector
	build *PaymentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (_g *PaymentGroupBy) Aggregate(fns ...AggregateFunc) *PaymentGroupBy {
	_g.fns = append(_g.fns, fns...)
	return _g
}

// Scan applies the selector query and scans the result into the given value.
func (_g *PaymentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _g.build.ctx, ent.OpQueryGroupBy)
	if err := _g.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaymentQuery, *PaymentGroupBy](ctx, _g.build, _g, _g.build.inters, v)
}

func (_g *PaymentGroupBy) sqlScan(ctx context.Context, root *PaymentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(_g.fns))
	for _, fn := range _g.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*_g.flds)+len(_g.fns))
		for _, f := range *_g.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*_g.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := _g.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PaymentSelect is the builder for selecting fields of Payment entities.
type PaymentSelect struct {
	*PaymentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (_s *PaymentSelect) Aggregate(fns ...AggregateFunc) *PaymentSelect {
	_s.fns = append(_s.fns, fns...)
	return _s
}

// Scan applies the selector query and scans the result into the given value.
func (_s *PaymentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _s.ctx, ent.OpQuerySelect)
	if err := _s.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaymentQuery, *PaymentSelect](ctx, _s.PaymentQuery, _s, _s.inters, v)
}

func (_s *PaymentSelect) sqlScan(ctx context.Context, root *PaymentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(_s.fns))
	for _, fn := range _s.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*_s.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := _s.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
