// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/division_choice_mtspmmg"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DivisionChoiceMTSPMMGQuery is the builder for querying Division_Choice_MTSPMMG entities.
type DivisionChoiceMTSPMMGQuery struct {
	config
	ctx                 *QueryContext
	order               []division_choice_mtspmmg.OrderOption
	inters              []Interceptor
	predicates          []predicate.Division_Choice_MTSPMMG
	withApplnMTSPMMGRef *ExamApplicationMTSPMMGQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DivisionChoiceMTSPMMGQuery builder.
func (dcmq *DivisionChoiceMTSPMMGQuery) Where(ps ...predicate.Division_Choice_MTSPMMG) *DivisionChoiceMTSPMMGQuery {
	dcmq.predicates = append(dcmq.predicates, ps...)
	return dcmq
}

// Limit the number of records to be returned by this query.
func (dcmq *DivisionChoiceMTSPMMGQuery) Limit(limit int) *DivisionChoiceMTSPMMGQuery {
	dcmq.ctx.Limit = &limit
	return dcmq
}

// Offset to start from.
func (dcmq *DivisionChoiceMTSPMMGQuery) Offset(offset int) *DivisionChoiceMTSPMMGQuery {
	dcmq.ctx.Offset = &offset
	return dcmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dcmq *DivisionChoiceMTSPMMGQuery) Unique(unique bool) *DivisionChoiceMTSPMMGQuery {
	dcmq.ctx.Unique = &unique
	return dcmq
}

// Order specifies how the records should be ordered.
func (dcmq *DivisionChoiceMTSPMMGQuery) Order(o ...division_choice_mtspmmg.OrderOption) *DivisionChoiceMTSPMMGQuery {
	dcmq.order = append(dcmq.order, o...)
	return dcmq
}

// QueryApplnMTSPMMGRef chains the current query on the "ApplnMTSPMMG_Ref" edge.
func (dcmq *DivisionChoiceMTSPMMGQuery) QueryApplnMTSPMMGRef() *ExamApplicationMTSPMMGQuery {
	query := (&ExamApplicationMTSPMMGClient{config: dcmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(division_choice_mtspmmg.Table, division_choice_mtspmmg.FieldID, selector),
			sqlgraph.To(exam_application_mtspmmg.Table, exam_application_mtspmmg.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, division_choice_mtspmmg.ApplnMTSPMMGRefTable, division_choice_mtspmmg.ApplnMTSPMMGRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Division_Choice_MTSPMMG entity from the query.
// Returns a *NotFoundError when no Division_Choice_MTSPMMG was found.
func (dcmq *DivisionChoiceMTSPMMGQuery) First(ctx context.Context) (*Division_Choice_MTSPMMG, error) {
	nodes, err := dcmq.Limit(1).All(setContextOp(ctx, dcmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{division_choice_mtspmmg.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dcmq *DivisionChoiceMTSPMMGQuery) FirstX(ctx context.Context) *Division_Choice_MTSPMMG {
	node, err := dcmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Division_Choice_MTSPMMG ID from the query.
// Returns a *NotFoundError when no Division_Choice_MTSPMMG ID was found.
func (dcmq *DivisionChoiceMTSPMMGQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = dcmq.Limit(1).IDs(setContextOp(ctx, dcmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{division_choice_mtspmmg.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dcmq *DivisionChoiceMTSPMMGQuery) FirstIDX(ctx context.Context) int32 {
	id, err := dcmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Division_Choice_MTSPMMG entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Division_Choice_MTSPMMG entity is found.
// Returns a *NotFoundError when no Division_Choice_MTSPMMG entities are found.
func (dcmq *DivisionChoiceMTSPMMGQuery) Only(ctx context.Context) (*Division_Choice_MTSPMMG, error) {
	nodes, err := dcmq.Limit(2).All(setContextOp(ctx, dcmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{division_choice_mtspmmg.Label}
	default:
		return nil, &NotSingularError{division_choice_mtspmmg.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dcmq *DivisionChoiceMTSPMMGQuery) OnlyX(ctx context.Context) *Division_Choice_MTSPMMG {
	node, err := dcmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Division_Choice_MTSPMMG ID in the query.
// Returns a *NotSingularError when more than one Division_Choice_MTSPMMG ID is found.
// Returns a *NotFoundError when no entities are found.
func (dcmq *DivisionChoiceMTSPMMGQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = dcmq.Limit(2).IDs(setContextOp(ctx, dcmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{division_choice_mtspmmg.Label}
	default:
		err = &NotSingularError{division_choice_mtspmmg.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dcmq *DivisionChoiceMTSPMMGQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := dcmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Division_Choice_MTSPMMGs.
func (dcmq *DivisionChoiceMTSPMMGQuery) All(ctx context.Context) ([]*Division_Choice_MTSPMMG, error) {
	ctx = setContextOp(ctx, dcmq.ctx, "All")
	if err := dcmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Division_Choice_MTSPMMG, *DivisionChoiceMTSPMMGQuery]()
	return withInterceptors[[]*Division_Choice_MTSPMMG](ctx, dcmq, qr, dcmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dcmq *DivisionChoiceMTSPMMGQuery) AllX(ctx context.Context) []*Division_Choice_MTSPMMG {
	nodes, err := dcmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Division_Choice_MTSPMMG IDs.
func (dcmq *DivisionChoiceMTSPMMGQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if dcmq.ctx.Unique == nil && dcmq.path != nil {
		dcmq.Unique(true)
	}
	ctx = setContextOp(ctx, dcmq.ctx, "IDs")
	if err = dcmq.Select(division_choice_mtspmmg.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dcmq *DivisionChoiceMTSPMMGQuery) IDsX(ctx context.Context) []int32 {
	ids, err := dcmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dcmq *DivisionChoiceMTSPMMGQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dcmq.ctx, "Count")
	if err := dcmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dcmq, querierCount[*DivisionChoiceMTSPMMGQuery](), dcmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dcmq *DivisionChoiceMTSPMMGQuery) CountX(ctx context.Context) int {
	count, err := dcmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dcmq *DivisionChoiceMTSPMMGQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dcmq.ctx, "Exist")
	switch _, err := dcmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dcmq *DivisionChoiceMTSPMMGQuery) ExistX(ctx context.Context) bool {
	exist, err := dcmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DivisionChoiceMTSPMMGQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dcmq *DivisionChoiceMTSPMMGQuery) Clone() *DivisionChoiceMTSPMMGQuery {
	if dcmq == nil {
		return nil
	}
	return &DivisionChoiceMTSPMMGQuery{
		config:              dcmq.config,
		ctx:                 dcmq.ctx.Clone(),
		order:               append([]division_choice_mtspmmg.OrderOption{}, dcmq.order...),
		inters:              append([]Interceptor{}, dcmq.inters...),
		predicates:          append([]predicate.Division_Choice_MTSPMMG{}, dcmq.predicates...),
		withApplnMTSPMMGRef: dcmq.withApplnMTSPMMGRef.Clone(),
		// clone intermediate query.
		sql:  dcmq.sql.Clone(),
		path: dcmq.path,
	}
}

// WithApplnMTSPMMGRef tells the query-builder to eager-load the nodes that are connected to
// the "ApplnMTSPMMG_Ref" edge. The optional arguments are used to configure the query builder of the edge.
func (dcmq *DivisionChoiceMTSPMMGQuery) WithApplnMTSPMMGRef(opts ...func(*ExamApplicationMTSPMMGQuery)) *DivisionChoiceMTSPMMGQuery {
	query := (&ExamApplicationMTSPMMGClient{config: dcmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcmq.withApplnMTSPMMGRef = query
	return dcmq
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
//	client.DivisionChoiceMTSPMMG.Query().
//		GroupBy(division_choice_mtspmmg.FieldApplicationID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dcmq *DivisionChoiceMTSPMMGQuery) GroupBy(field string, fields ...string) *DivisionChoiceMTSPMMGGroupBy {
	dcmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DivisionChoiceMTSPMMGGroupBy{build: dcmq}
	grbuild.flds = &dcmq.ctx.Fields
	grbuild.label = division_choice_mtspmmg.Label
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
//	client.DivisionChoiceMTSPMMG.Query().
//		Select(division_choice_mtspmmg.FieldApplicationID).
//		Scan(ctx, &v)
func (dcmq *DivisionChoiceMTSPMMGQuery) Select(fields ...string) *DivisionChoiceMTSPMMGSelect {
	dcmq.ctx.Fields = append(dcmq.ctx.Fields, fields...)
	sbuild := &DivisionChoiceMTSPMMGSelect{DivisionChoiceMTSPMMGQuery: dcmq}
	sbuild.label = division_choice_mtspmmg.Label
	sbuild.flds, sbuild.scan = &dcmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DivisionChoiceMTSPMMGSelect configured with the given aggregations.
func (dcmq *DivisionChoiceMTSPMMGQuery) Aggregate(fns ...AggregateFunc) *DivisionChoiceMTSPMMGSelect {
	return dcmq.Select().Aggregate(fns...)
}

func (dcmq *DivisionChoiceMTSPMMGQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dcmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dcmq); err != nil {
				return err
			}
		}
	}
	for _, f := range dcmq.ctx.Fields {
		if !division_choice_mtspmmg.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dcmq.path != nil {
		prev, err := dcmq.path(ctx)
		if err != nil {
			return err
		}
		dcmq.sql = prev
	}
	return nil
}

func (dcmq *DivisionChoiceMTSPMMGQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Division_Choice_MTSPMMG, error) {
	var (
		nodes       = []*Division_Choice_MTSPMMG{}
		_spec       = dcmq.querySpec()
		loadedTypes = [1]bool{
			dcmq.withApplnMTSPMMGRef != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Division_Choice_MTSPMMG).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Division_Choice_MTSPMMG{config: dcmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dcmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dcmq.withApplnMTSPMMGRef; query != nil {
		if err := dcmq.loadApplnMTSPMMGRef(ctx, query, nodes, nil,
			func(n *Division_Choice_MTSPMMG, e *Exam_Application_MTSPMMG) { n.Edges.ApplnMTSPMMGRef = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dcmq *DivisionChoiceMTSPMMGQuery) loadApplnMTSPMMGRef(ctx context.Context, query *ExamApplicationMTSPMMGQuery, nodes []*Division_Choice_MTSPMMG, init func(*Division_Choice_MTSPMMG), assign func(*Division_Choice_MTSPMMG, *Exam_Application_MTSPMMG)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Division_Choice_MTSPMMG)
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
	query.Where(exam_application_mtspmmg.IDIn(ids...))
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

func (dcmq *DivisionChoiceMTSPMMGQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dcmq.querySpec()
	_spec.Node.Columns = dcmq.ctx.Fields
	if len(dcmq.ctx.Fields) > 0 {
		_spec.Unique = dcmq.ctx.Unique != nil && *dcmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dcmq.driver, _spec)
}

func (dcmq *DivisionChoiceMTSPMMGQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(division_choice_mtspmmg.Table, division_choice_mtspmmg.Columns, sqlgraph.NewFieldSpec(division_choice_mtspmmg.FieldID, field.TypeInt32))
	_spec.From = dcmq.sql
	if unique := dcmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dcmq.path != nil {
		_spec.Unique = true
	}
	if fields := dcmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, division_choice_mtspmmg.FieldID)
		for i := range fields {
			if fields[i] != division_choice_mtspmmg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dcmq.withApplnMTSPMMGRef != nil {
			_spec.Node.AddColumnOnce(division_choice_mtspmmg.FieldApplicationID)
		}
	}
	if ps := dcmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dcmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dcmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dcmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dcmq *DivisionChoiceMTSPMMGQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dcmq.driver.Dialect())
	t1 := builder.Table(division_choice_mtspmmg.Table)
	columns := dcmq.ctx.Fields
	if len(columns) == 0 {
		columns = division_choice_mtspmmg.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dcmq.sql != nil {
		selector = dcmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dcmq.ctx.Unique != nil && *dcmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dcmq.predicates {
		p(selector)
	}
	for _, p := range dcmq.order {
		p(selector)
	}
	if offset := dcmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dcmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DivisionChoiceMTSPMMGGroupBy is the group-by builder for Division_Choice_MTSPMMG entities.
type DivisionChoiceMTSPMMGGroupBy struct {
	selector
	build *DivisionChoiceMTSPMMGQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dcmgb *DivisionChoiceMTSPMMGGroupBy) Aggregate(fns ...AggregateFunc) *DivisionChoiceMTSPMMGGroupBy {
	dcmgb.fns = append(dcmgb.fns, fns...)
	return dcmgb
}

// Scan applies the selector query and scans the result into the given value.
func (dcmgb *DivisionChoiceMTSPMMGGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcmgb.build.ctx, "GroupBy")
	if err := dcmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DivisionChoiceMTSPMMGQuery, *DivisionChoiceMTSPMMGGroupBy](ctx, dcmgb.build, dcmgb, dcmgb.build.inters, v)
}

func (dcmgb *DivisionChoiceMTSPMMGGroupBy) sqlScan(ctx context.Context, root *DivisionChoiceMTSPMMGQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dcmgb.fns))
	for _, fn := range dcmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dcmgb.flds)+len(dcmgb.fns))
		for _, f := range *dcmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dcmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DivisionChoiceMTSPMMGSelect is the builder for selecting fields of DivisionChoiceMTSPMMG entities.
type DivisionChoiceMTSPMMGSelect struct {
	*DivisionChoiceMTSPMMGQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dcms *DivisionChoiceMTSPMMGSelect) Aggregate(fns ...AggregateFunc) *DivisionChoiceMTSPMMGSelect {
	dcms.fns = append(dcms.fns, fns...)
	return dcms
}

// Scan applies the selector query and scans the result into the given value.
func (dcms *DivisionChoiceMTSPMMGSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcms.ctx, "Select")
	if err := dcms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DivisionChoiceMTSPMMGQuery, *DivisionChoiceMTSPMMGSelect](ctx, dcms.DivisionChoiceMTSPMMGQuery, dcms, dcms.inters, v)
}

func (dcms *DivisionChoiceMTSPMMGSelect) sqlScan(ctx context.Context, root *DivisionChoiceMTSPMMGQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dcms.fns))
	for _, fn := range dcms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dcms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
