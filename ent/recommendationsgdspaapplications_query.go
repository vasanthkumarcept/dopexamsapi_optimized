// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/exam_applications_gdspa"
	"recruit/ent/predicate"
	"recruit/ent/recommendationsgdspaapplications"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecommendationsGDSPAApplicationsQuery is the builder for querying RecommendationsGDSPAApplications entities.
type RecommendationsGDSPAApplicationsQuery struct {
	config
	ctx          *QueryContext
	order        []recommendationsgdspaapplications.OrderOption
	inters       []Interceptor
	predicates   []predicate.RecommendationsGDSPAApplications
	withApplnRef *ExamApplicationsGDSPAQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RecommendationsGDSPAApplicationsQuery builder.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Where(ps ...predicate.RecommendationsGDSPAApplications) *RecommendationsGDSPAApplicationsQuery {
	rgaq.predicates = append(rgaq.predicates, ps...)
	return rgaq
}

// Limit the number of records to be returned by this query.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Limit(limit int) *RecommendationsGDSPAApplicationsQuery {
	rgaq.ctx.Limit = &limit
	return rgaq
}

// Offset to start from.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Offset(offset int) *RecommendationsGDSPAApplicationsQuery {
	rgaq.ctx.Offset = &offset
	return rgaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Unique(unique bool) *RecommendationsGDSPAApplicationsQuery {
	rgaq.ctx.Unique = &unique
	return rgaq
}

// Order specifies how the records should be ordered.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Order(o ...recommendationsgdspaapplications.OrderOption) *RecommendationsGDSPAApplicationsQuery {
	rgaq.order = append(rgaq.order, o...)
	return rgaq
}

// QueryApplnRef chains the current query on the "ApplnRef" edge.
func (rgaq *RecommendationsGDSPAApplicationsQuery) QueryApplnRef() *ExamApplicationsGDSPAQuery {
	query := (&ExamApplicationsGDSPAClient{config: rgaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rgaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rgaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recommendationsgdspaapplications.Table, recommendationsgdspaapplications.FieldID, selector),
			sqlgraph.To(exam_applications_gdspa.Table, exam_applications_gdspa.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recommendationsgdspaapplications.ApplnRefTable, recommendationsgdspaapplications.ApplnRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(rgaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RecommendationsGDSPAApplications entity from the query.
// Returns a *NotFoundError when no RecommendationsGDSPAApplications was found.
func (rgaq *RecommendationsGDSPAApplicationsQuery) First(ctx context.Context) (*RecommendationsGDSPAApplications, error) {
	nodes, err := rgaq.Limit(1).All(setContextOp(ctx, rgaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{recommendationsgdspaapplications.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rgaq *RecommendationsGDSPAApplicationsQuery) FirstX(ctx context.Context) *RecommendationsGDSPAApplications {
	node, err := rgaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RecommendationsGDSPAApplications ID from the query.
// Returns a *NotFoundError when no RecommendationsGDSPAApplications ID was found.
func (rgaq *RecommendationsGDSPAApplicationsQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rgaq.Limit(1).IDs(setContextOp(ctx, rgaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{recommendationsgdspaapplications.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rgaq *RecommendationsGDSPAApplicationsQuery) FirstIDX(ctx context.Context) int64 {
	id, err := rgaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RecommendationsGDSPAApplications entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RecommendationsGDSPAApplications entity is found.
// Returns a *NotFoundError when no RecommendationsGDSPAApplications entities are found.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Only(ctx context.Context) (*RecommendationsGDSPAApplications, error) {
	nodes, err := rgaq.Limit(2).All(setContextOp(ctx, rgaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{recommendationsgdspaapplications.Label}
	default:
		return nil, &NotSingularError{recommendationsgdspaapplications.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rgaq *RecommendationsGDSPAApplicationsQuery) OnlyX(ctx context.Context) *RecommendationsGDSPAApplications {
	node, err := rgaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RecommendationsGDSPAApplications ID in the query.
// Returns a *NotSingularError when more than one RecommendationsGDSPAApplications ID is found.
// Returns a *NotFoundError when no entities are found.
func (rgaq *RecommendationsGDSPAApplicationsQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rgaq.Limit(2).IDs(setContextOp(ctx, rgaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{recommendationsgdspaapplications.Label}
	default:
		err = &NotSingularError{recommendationsgdspaapplications.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rgaq *RecommendationsGDSPAApplicationsQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := rgaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RecommendationsGDSPAApplicationsSlice.
func (rgaq *RecommendationsGDSPAApplicationsQuery) All(ctx context.Context) ([]*RecommendationsGDSPAApplications, error) {
	ctx = setContextOp(ctx, rgaq.ctx, "All")
	if err := rgaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RecommendationsGDSPAApplications, *RecommendationsGDSPAApplicationsQuery]()
	return withInterceptors[[]*RecommendationsGDSPAApplications](ctx, rgaq, qr, rgaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rgaq *RecommendationsGDSPAApplicationsQuery) AllX(ctx context.Context) []*RecommendationsGDSPAApplications {
	nodes, err := rgaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RecommendationsGDSPAApplications IDs.
func (rgaq *RecommendationsGDSPAApplicationsQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if rgaq.ctx.Unique == nil && rgaq.path != nil {
		rgaq.Unique(true)
	}
	ctx = setContextOp(ctx, rgaq.ctx, "IDs")
	if err = rgaq.Select(recommendationsgdspaapplications.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rgaq *RecommendationsGDSPAApplicationsQuery) IDsX(ctx context.Context) []int64 {
	ids, err := rgaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rgaq.ctx, "Count")
	if err := rgaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rgaq, querierCount[*RecommendationsGDSPAApplicationsQuery](), rgaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rgaq *RecommendationsGDSPAApplicationsQuery) CountX(ctx context.Context) int {
	count, err := rgaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rgaq.ctx, "Exist")
	switch _, err := rgaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rgaq *RecommendationsGDSPAApplicationsQuery) ExistX(ctx context.Context) bool {
	exist, err := rgaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RecommendationsGDSPAApplicationsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Clone() *RecommendationsGDSPAApplicationsQuery {
	if rgaq == nil {
		return nil
	}
	return &RecommendationsGDSPAApplicationsQuery{
		config:       rgaq.config,
		ctx:          rgaq.ctx.Clone(),
		order:        append([]recommendationsgdspaapplications.OrderOption{}, rgaq.order...),
		inters:       append([]Interceptor{}, rgaq.inters...),
		predicates:   append([]predicate.RecommendationsGDSPAApplications{}, rgaq.predicates...),
		withApplnRef: rgaq.withApplnRef.Clone(),
		// clone intermediate query.
		sql:  rgaq.sql.Clone(),
		path: rgaq.path,
	}
}

// WithApplnRef tells the query-builder to eager-load the nodes that are connected to
// the "ApplnRef" edge. The optional arguments are used to configure the query builder of the edge.
func (rgaq *RecommendationsGDSPAApplicationsQuery) WithApplnRef(opts ...func(*ExamApplicationsGDSPAQuery)) *RecommendationsGDSPAApplicationsQuery {
	query := (&ExamApplicationsGDSPAClient{config: rgaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rgaq.withApplnRef = query
	return rgaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ApplicationID int64 `json:"ApplicationID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RecommendationsGDSPAApplications.Query().
//		GroupBy(recommendationsgdspaapplications.FieldApplicationID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rgaq *RecommendationsGDSPAApplicationsQuery) GroupBy(field string, fields ...string) *RecommendationsGDSPAApplicationsGroupBy {
	rgaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RecommendationsGDSPAApplicationsGroupBy{build: rgaq}
	grbuild.flds = &rgaq.ctx.Fields
	grbuild.label = recommendationsgdspaapplications.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ApplicationID int64 `json:"ApplicationID,omitempty"`
//	}
//
//	client.RecommendationsGDSPAApplications.Query().
//		Select(recommendationsgdspaapplications.FieldApplicationID).
//		Scan(ctx, &v)
func (rgaq *RecommendationsGDSPAApplicationsQuery) Select(fields ...string) *RecommendationsGDSPAApplicationsSelect {
	rgaq.ctx.Fields = append(rgaq.ctx.Fields, fields...)
	sbuild := &RecommendationsGDSPAApplicationsSelect{RecommendationsGDSPAApplicationsQuery: rgaq}
	sbuild.label = recommendationsgdspaapplications.Label
	sbuild.flds, sbuild.scan = &rgaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RecommendationsGDSPAApplicationsSelect configured with the given aggregations.
func (rgaq *RecommendationsGDSPAApplicationsQuery) Aggregate(fns ...AggregateFunc) *RecommendationsGDSPAApplicationsSelect {
	return rgaq.Select().Aggregate(fns...)
}

func (rgaq *RecommendationsGDSPAApplicationsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rgaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rgaq); err != nil {
				return err
			}
		}
	}
	for _, f := range rgaq.ctx.Fields {
		if !recommendationsgdspaapplications.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rgaq.path != nil {
		prev, err := rgaq.path(ctx)
		if err != nil {
			return err
		}
		rgaq.sql = prev
	}
	return nil
}

func (rgaq *RecommendationsGDSPAApplicationsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RecommendationsGDSPAApplications, error) {
	var (
		nodes       = []*RecommendationsGDSPAApplications{}
		_spec       = rgaq.querySpec()
		loadedTypes = [1]bool{
			rgaq.withApplnRef != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RecommendationsGDSPAApplications).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RecommendationsGDSPAApplications{config: rgaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rgaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rgaq.withApplnRef; query != nil {
		if err := rgaq.loadApplnRef(ctx, query, nodes, nil,
			func(n *RecommendationsGDSPAApplications, e *Exam_Applications_GDSPA) { n.Edges.ApplnRef = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rgaq *RecommendationsGDSPAApplicationsQuery) loadApplnRef(ctx context.Context, query *ExamApplicationsGDSPAQuery, nodes []*RecommendationsGDSPAApplications, init func(*RecommendationsGDSPAApplications), assign func(*RecommendationsGDSPAApplications, *Exam_Applications_GDSPA)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*RecommendationsGDSPAApplications)
	for i := range nodes {
		fk := nodes[i].ApplicationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(exam_applications_gdspa.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ApplicationID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rgaq *RecommendationsGDSPAApplicationsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rgaq.querySpec()
	_spec.Node.Columns = rgaq.ctx.Fields
	if len(rgaq.ctx.Fields) > 0 {
		_spec.Unique = rgaq.ctx.Unique != nil && *rgaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rgaq.driver, _spec)
}

func (rgaq *RecommendationsGDSPAApplicationsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(recommendationsgdspaapplications.Table, recommendationsgdspaapplications.Columns, sqlgraph.NewFieldSpec(recommendationsgdspaapplications.FieldID, field.TypeInt64))
	_spec.From = rgaq.sql
	if unique := rgaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rgaq.path != nil {
		_spec.Unique = true
	}
	if fields := rgaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recommendationsgdspaapplications.FieldID)
		for i := range fields {
			if fields[i] != recommendationsgdspaapplications.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if rgaq.withApplnRef != nil {
			_spec.Node.AddColumnOnce(recommendationsgdspaapplications.FieldApplicationID)
		}
	}
	if ps := rgaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rgaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rgaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rgaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rgaq *RecommendationsGDSPAApplicationsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rgaq.driver.Dialect())
	t1 := builder.Table(recommendationsgdspaapplications.Table)
	columns := rgaq.ctx.Fields
	if len(columns) == 0 {
		columns = recommendationsgdspaapplications.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rgaq.sql != nil {
		selector = rgaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rgaq.ctx.Unique != nil && *rgaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rgaq.predicates {
		p(selector)
	}
	for _, p := range rgaq.order {
		p(selector)
	}
	if offset := rgaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rgaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RecommendationsGDSPAApplicationsGroupBy is the group-by builder for RecommendationsGDSPAApplications entities.
type RecommendationsGDSPAApplicationsGroupBy struct {
	selector
	build *RecommendationsGDSPAApplicationsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgagb *RecommendationsGDSPAApplicationsGroupBy) Aggregate(fns ...AggregateFunc) *RecommendationsGDSPAApplicationsGroupBy {
	rgagb.fns = append(rgagb.fns, fns...)
	return rgagb
}

// Scan applies the selector query and scans the result into the given value.
func (rgagb *RecommendationsGDSPAApplicationsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgagb.build.ctx, "GroupBy")
	if err := rgagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecommendationsGDSPAApplicationsQuery, *RecommendationsGDSPAApplicationsGroupBy](ctx, rgagb.build, rgagb, rgagb.build.inters, v)
}

func (rgagb *RecommendationsGDSPAApplicationsGroupBy) sqlScan(ctx context.Context, root *RecommendationsGDSPAApplicationsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgagb.fns))
	for _, fn := range rgagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgagb.flds)+len(rgagb.fns))
		for _, f := range *rgagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RecommendationsGDSPAApplicationsSelect is the builder for selecting fields of RecommendationsGDSPAApplications entities.
type RecommendationsGDSPAApplicationsSelect struct {
	*RecommendationsGDSPAApplicationsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rgas *RecommendationsGDSPAApplicationsSelect) Aggregate(fns ...AggregateFunc) *RecommendationsGDSPAApplicationsSelect {
	rgas.fns = append(rgas.fns, fns...)
	return rgas
}

// Scan applies the selector query and scans the result into the given value.
func (rgas *RecommendationsGDSPAApplicationsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgas.ctx, "Select")
	if err := rgas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecommendationsGDSPAApplicationsQuery, *RecommendationsGDSPAApplicationsSelect](ctx, rgas.RecommendationsGDSPAApplicationsQuery, rgas, rgas.inters, v)
}

func (rgas *RecommendationsGDSPAApplicationsSelect) sqlScan(ctx context.Context, root *RecommendationsGDSPAApplicationsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rgas.fns))
	for _, fn := range rgas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rgas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}