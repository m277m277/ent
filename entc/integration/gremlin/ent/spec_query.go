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
	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"entgo.io/ent/entc/integration/gremlin/ent/spec"
)

// SpecQuery is the builder for querying Spec entities.
type SpecQuery struct {
	config
	ctx        *QueryContext
	order      []spec.OrderOption
	inters     []Interceptor
	predicates []predicate.Spec
	withCard   *CardQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the SpecQuery builder.
func (_q *SpecQuery) Where(ps ...predicate.Spec) *SpecQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *SpecQuery) Limit(limit int) *SpecQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *SpecQuery) Offset(offset int) *SpecQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *SpecQuery) Unique(unique bool) *SpecQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *SpecQuery) Order(o ...spec.OrderOption) *SpecQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryCard chains the current query on the "card" edge.
func (_q *SpecQuery) QueryCard() *CardQuery {
	query := (&CardClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := _q.gremlinQuery(ctx)
		fromU = gremlin.OutE(spec.CardLabel).InV()
		return fromU, nil
	}
	return query
}

// First returns the first Spec entity from the query.
// Returns a *NotFoundError when no Spec was found.
func (_q *SpecQuery) First(ctx context.Context) (*Spec, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{spec.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *SpecQuery) FirstX(ctx context.Context) *Spec {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Spec ID from the query.
// Returns a *NotFoundError when no Spec ID was found.
func (_q *SpecQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{spec.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *SpecQuery) FirstIDX(ctx context.Context) string {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Spec entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Spec entity is found.
// Returns a *NotFoundError when no Spec entities are found.
func (_q *SpecQuery) Only(ctx context.Context) (*Spec, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{spec.Label}
	default:
		return nil, &NotSingularError{spec.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *SpecQuery) OnlyX(ctx context.Context) *Spec {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Spec ID in the query.
// Returns a *NotSingularError when more than one Spec ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *SpecQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{spec.Label}
	default:
		err = &NotSingularError{spec.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *SpecQuery) OnlyIDX(ctx context.Context) string {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Specs.
func (_q *SpecQuery) All(ctx context.Context) ([]*Spec, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Spec, *SpecQuery]()
	return withInterceptors[[]*Spec](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *SpecQuery) AllX(ctx context.Context) []*Spec {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Spec IDs.
func (_q *SpecQuery) IDs(ctx context.Context) (ids []string, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(spec.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *SpecQuery) IDsX(ctx context.Context) []string {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *SpecQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*SpecQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *SpecQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *SpecQuery) Exist(ctx context.Context) (bool, error) {
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
func (_q *SpecQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SpecQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *SpecQuery) Clone() *SpecQuery {
	if _q == nil {
		return nil
	}
	return &SpecQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]spec.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.Spec{}, _q.predicates...),
		withCard:   _q.withCard.Clone(),
		// clone intermediate query.
		gremlin: _q.gremlin.Clone(),
		path:    _q.path,
	}
}

// WithCard tells the query-builder to eager-load the nodes that are connected to
// the "card" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *SpecQuery) WithCard(opts ...func(*CardQuery)) *SpecQuery {
	query := (&CardClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withCard = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (_q *SpecQuery) GroupBy(field string, fields ...string) *SpecGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SpecGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = spec.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (_q *SpecQuery) Select(fields ...string) *SpecSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &SpecSelect{SpecQuery: _q}
	sbuild.label = spec.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SpecSelect configured with the given aggregations.
func (_q *SpecQuery) Aggregate(fns ...AggregateFunc) *SpecSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *SpecQuery) prepareQuery(ctx context.Context) error {
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
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.gremlin = prev
	}
	return nil
}

func (_q *SpecQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Spec, error) {
	res := &gremlin.Response{}
	traversal := _q.gremlinQuery(ctx)
	if len(_q.ctx.Fields) > 0 {
		fields := make([]any, len(_q.ctx.Fields))
		for i, f := range _q.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var _ms Specs
	if err := _ms.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range _ms {
		_ms[i].config = _q.config
	}
	return _ms, nil
}

func (_q *SpecQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _q.gremlinQuery(ctx).Count().Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (_q *SpecQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(spec.Label)
	if _q.gremlin != nil {
		v = _q.gremlin.Clone()
	}
	for _, p := range _q.predicates {
		p(v)
	}
	if len(_q.order) > 0 {
		v.Order()
		for _, p := range _q.order {
			p(v)
		}
	}
	switch limit, offset := _q.ctx.Limit, _q.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := _q.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// SpecGroupBy is the group-by builder for Spec entities.
type SpecGroupBy struct {
	selector
	build *SpecQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (_g *SpecGroupBy) Aggregate(fns ...AggregateFunc) *SpecGroupBy {
	_g.fns = append(_g.fns, fns...)
	return _g
}

// Scan applies the selector query and scans the result into the given value.
func (_g *SpecGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _g.build.ctx, ent.OpQueryGroupBy)
	if err := _g.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpecQuery, *SpecGroupBy](ctx, _g.build, _g, _g.build.inters, v)
}

func (_g *SpecGroupBy) gremlinScan(ctx context.Context, root *SpecQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range _g.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *_g.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*_g.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := _g.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*_g.flds)+len(_g.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// SpecSelect is the builder for selecting fields of Spec entities.
type SpecSelect struct {
	*SpecQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (_s *SpecSelect) Aggregate(fns ...AggregateFunc) *SpecSelect {
	_s.fns = append(_s.fns, fns...)
	return _s
}

// Scan applies the selector query and scans the result into the given value.
func (_s *SpecSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _s.ctx, ent.OpQuerySelect)
	if err := _s.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpecQuery, *SpecSelect](ctx, _s.SpecQuery, _s, _s.inters, v)
}

func (_s *SpecSelect) gremlinScan(ctx context.Context, root *SpecQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := _s.ctx.Fields; len(fields) == 1 {
		if fields[0] != spec.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(_s.ctx.Fields))
		for i, f := range _s.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := _s.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
