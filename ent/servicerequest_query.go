// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/predicate"
	"recruit/ent/servicerequest"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ServiceRequestQuery is the builder for querying ServiceRequest entities.
type ServiceRequestQuery struct {
	config
	ctx        *QueryContext
	order      []servicerequest.OrderOption
	inters     []Interceptor
	predicates []predicate.ServiceRequest
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServiceRequestQuery builder.
func (srq *ServiceRequestQuery) Where(ps ...predicate.ServiceRequest) *ServiceRequestQuery {
	srq.predicates = append(srq.predicates, ps...)
	return srq
}

// Limit the number of records to be returned by this query.
func (srq *ServiceRequestQuery) Limit(limit int) *ServiceRequestQuery {
	srq.ctx.Limit = &limit
	return srq
}

// Offset to start from.
func (srq *ServiceRequestQuery) Offset(offset int) *ServiceRequestQuery {
	srq.ctx.Offset = &offset
	return srq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (srq *ServiceRequestQuery) Unique(unique bool) *ServiceRequestQuery {
	srq.ctx.Unique = &unique
	return srq
}

// Order specifies how the records should be ordered.
func (srq *ServiceRequestQuery) Order(o ...servicerequest.OrderOption) *ServiceRequestQuery {
	srq.order = append(srq.order, o...)
	return srq
}

// First returns the first ServiceRequest entity from the query.
// Returns a *NotFoundError when no ServiceRequest was found.
func (srq *ServiceRequestQuery) First(ctx context.Context) (*ServiceRequest, error) {
	nodes, err := srq.Limit(1).All(setContextOp(ctx, srq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{servicerequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (srq *ServiceRequestQuery) FirstX(ctx context.Context) *ServiceRequest {
	node, err := srq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServiceRequest ID from the query.
// Returns a *NotFoundError when no ServiceRequest ID was found.
func (srq *ServiceRequestQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = srq.Limit(1).IDs(setContextOp(ctx, srq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{servicerequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (srq *ServiceRequestQuery) FirstIDX(ctx context.Context) int64 {
	id, err := srq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServiceRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ServiceRequest entity is found.
// Returns a *NotFoundError when no ServiceRequest entities are found.
func (srq *ServiceRequestQuery) Only(ctx context.Context) (*ServiceRequest, error) {
	nodes, err := srq.Limit(2).All(setContextOp(ctx, srq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{servicerequest.Label}
	default:
		return nil, &NotSingularError{servicerequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (srq *ServiceRequestQuery) OnlyX(ctx context.Context) *ServiceRequest {
	node, err := srq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServiceRequest ID in the query.
// Returns a *NotSingularError when more than one ServiceRequest ID is found.
// Returns a *NotFoundError when no entities are found.
func (srq *ServiceRequestQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = srq.Limit(2).IDs(setContextOp(ctx, srq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{servicerequest.Label}
	default:
		err = &NotSingularError{servicerequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (srq *ServiceRequestQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := srq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServiceRequests.
func (srq *ServiceRequestQuery) All(ctx context.Context) ([]*ServiceRequest, error) {
	ctx = setContextOp(ctx, srq.ctx, "All")
	if err := srq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ServiceRequest, *ServiceRequestQuery]()
	return withInterceptors[[]*ServiceRequest](ctx, srq, qr, srq.inters)
}

// AllX is like All, but panics if an error occurs.
func (srq *ServiceRequestQuery) AllX(ctx context.Context) []*ServiceRequest {
	nodes, err := srq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServiceRequest IDs.
func (srq *ServiceRequestQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if srq.ctx.Unique == nil && srq.path != nil {
		srq.Unique(true)
	}
	ctx = setContextOp(ctx, srq.ctx, "IDs")
	if err = srq.Select(servicerequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (srq *ServiceRequestQuery) IDsX(ctx context.Context) []int64 {
	ids, err := srq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (srq *ServiceRequestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, srq.ctx, "Count")
	if err := srq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, srq, querierCount[*ServiceRequestQuery](), srq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (srq *ServiceRequestQuery) CountX(ctx context.Context) int {
	count, err := srq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (srq *ServiceRequestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, srq.ctx, "Exist")
	switch _, err := srq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (srq *ServiceRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := srq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServiceRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (srq *ServiceRequestQuery) Clone() *ServiceRequestQuery {
	if srq == nil {
		return nil
	}
	return &ServiceRequestQuery{
		config:     srq.config,
		ctx:        srq.ctx.Clone(),
		order:      append([]servicerequest.OrderOption{}, srq.order...),
		inters:     append([]Interceptor{}, srq.inters...),
		predicates: append([]predicate.ServiceRequest{}, srq.predicates...),
		// clone intermediate query.
		sql:  srq.sql.Clone(),
		path: srq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Remarks string `json:"remarks,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ServiceRequest.Query().
//		GroupBy(servicerequest.FieldRemarks).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (srq *ServiceRequestQuery) GroupBy(field string, fields ...string) *ServiceRequestGroupBy {
	srq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ServiceRequestGroupBy{build: srq}
	grbuild.flds = &srq.ctx.Fields
	grbuild.label = servicerequest.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Remarks string `json:"remarks,omitempty"`
//	}
//
//	client.ServiceRequest.Query().
//		Select(servicerequest.FieldRemarks).
//		Scan(ctx, &v)
func (srq *ServiceRequestQuery) Select(fields ...string) *ServiceRequestSelect {
	srq.ctx.Fields = append(srq.ctx.Fields, fields...)
	sbuild := &ServiceRequestSelect{ServiceRequestQuery: srq}
	sbuild.label = servicerequest.Label
	sbuild.flds, sbuild.scan = &srq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ServiceRequestSelect configured with the given aggregations.
func (srq *ServiceRequestQuery) Aggregate(fns ...AggregateFunc) *ServiceRequestSelect {
	return srq.Select().Aggregate(fns...)
}

func (srq *ServiceRequestQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range srq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, srq); err != nil {
				return err
			}
		}
	}
	for _, f := range srq.ctx.Fields {
		if !servicerequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if srq.path != nil {
		prev, err := srq.path(ctx)
		if err != nil {
			return err
		}
		srq.sql = prev
	}
	return nil
}

func (srq *ServiceRequestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ServiceRequest, error) {
	var (
		nodes = []*ServiceRequest{}
		_spec = srq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ServiceRequest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ServiceRequest{config: srq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, srq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (srq *ServiceRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := srq.querySpec()
	_spec.Node.Columns = srq.ctx.Fields
	if len(srq.ctx.Fields) > 0 {
		_spec.Unique = srq.ctx.Unique != nil && *srq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, srq.driver, _spec)
}

func (srq *ServiceRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(servicerequest.Table, servicerequest.Columns, sqlgraph.NewFieldSpec(servicerequest.FieldID, field.TypeInt64))
	_spec.From = srq.sql
	if unique := srq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if srq.path != nil {
		_spec.Unique = true
	}
	if fields := srq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servicerequest.FieldID)
		for i := range fields {
			if fields[i] != servicerequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := srq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := srq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := srq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := srq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (srq *ServiceRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(srq.driver.Dialect())
	t1 := builder.Table(servicerequest.Table)
	columns := srq.ctx.Fields
	if len(columns) == 0 {
		columns = servicerequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if srq.sql != nil {
		selector = srq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if srq.ctx.Unique != nil && *srq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range srq.predicates {
		p(selector)
	}
	for _, p := range srq.order {
		p(selector)
	}
	if offset := srq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := srq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServiceRequestGroupBy is the group-by builder for ServiceRequest entities.
type ServiceRequestGroupBy struct {
	selector
	build *ServiceRequestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (srgb *ServiceRequestGroupBy) Aggregate(fns ...AggregateFunc) *ServiceRequestGroupBy {
	srgb.fns = append(srgb.fns, fns...)
	return srgb
}

// Scan applies the selector query and scans the result into the given value.
func (srgb *ServiceRequestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, srgb.build.ctx, "GroupBy")
	if err := srgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServiceRequestQuery, *ServiceRequestGroupBy](ctx, srgb.build, srgb, srgb.build.inters, v)
}

func (srgb *ServiceRequestGroupBy) sqlScan(ctx context.Context, root *ServiceRequestQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(srgb.fns))
	for _, fn := range srgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*srgb.flds)+len(srgb.fns))
		for _, f := range *srgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*srgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ServiceRequestSelect is the builder for selecting fields of ServiceRequest entities.
type ServiceRequestSelect struct {
	*ServiceRequestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (srs *ServiceRequestSelect) Aggregate(fns ...AggregateFunc) *ServiceRequestSelect {
	srs.fns = append(srs.fns, fns...)
	return srs
}

// Scan applies the selector query and scans the result into the given value.
func (srs *ServiceRequestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, srs.ctx, "Select")
	if err := srs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServiceRequestQuery, *ServiceRequestSelect](ctx, srs.ServiceRequestQuery, srs, srs.inters, v)
}

func (srs *ServiceRequestSelect) sqlScan(ctx context.Context, root *ServiceRequestQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(srs.fns))
	for _, fn := range srs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*srs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}