// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/cadre_choice_mtspmmg"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CadreChoiceMTSPMMGQuery is the builder for querying Cadre_Choice_MTSPMMG entities.
type CadreChoiceMTSPMMGQuery struct {
	config
	ctx                 *QueryContext
	order               []cadre_choice_mtspmmg.OrderOption
	inters              []Interceptor
	predicates          []predicate.Cadre_Choice_MTSPMMG
	withApplnMTSPMMGRef *ExamApplicationMTSPMMGQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CadreChoiceMTSPMMGQuery builder.
func (ccmq *CadreChoiceMTSPMMGQuery) Where(ps ...predicate.Cadre_Choice_MTSPMMG) *CadreChoiceMTSPMMGQuery {
	ccmq.predicates = append(ccmq.predicates, ps...)
	return ccmq
}

// Limit the number of records to be returned by this query.
func (ccmq *CadreChoiceMTSPMMGQuery) Limit(limit int) *CadreChoiceMTSPMMGQuery {
	ccmq.ctx.Limit = &limit
	return ccmq
}

// Offset to start from.
func (ccmq *CadreChoiceMTSPMMGQuery) Offset(offset int) *CadreChoiceMTSPMMGQuery {
	ccmq.ctx.Offset = &offset
	return ccmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccmq *CadreChoiceMTSPMMGQuery) Unique(unique bool) *CadreChoiceMTSPMMGQuery {
	ccmq.ctx.Unique = &unique
	return ccmq
}

// Order specifies how the records should be ordered.
func (ccmq *CadreChoiceMTSPMMGQuery) Order(o ...cadre_choice_mtspmmg.OrderOption) *CadreChoiceMTSPMMGQuery {
	ccmq.order = append(ccmq.order, o...)
	return ccmq
}

// QueryApplnMTSPMMGRef chains the current query on the "ApplnMTSPMMG_Ref" edge.
func (ccmq *CadreChoiceMTSPMMGQuery) QueryApplnMTSPMMGRef() *ExamApplicationMTSPMMGQuery {
	query := (&ExamApplicationMTSPMMGClient{config: ccmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ccmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ccmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cadre_choice_mtspmmg.Table, cadre_choice_mtspmmg.FieldID, selector),
			sqlgraph.To(exam_application_mtspmmg.Table, exam_application_mtspmmg.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cadre_choice_mtspmmg.ApplnMTSPMMGRefTable, cadre_choice_mtspmmg.ApplnMTSPMMGRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(ccmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Cadre_Choice_MTSPMMG entity from the query.
// Returns a *NotFoundError when no Cadre_Choice_MTSPMMG was found.
func (ccmq *CadreChoiceMTSPMMGQuery) First(ctx context.Context) (*Cadre_Choice_MTSPMMG, error) {
	nodes, err := ccmq.Limit(1).All(setContextOp(ctx, ccmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cadre_choice_mtspmmg.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccmq *CadreChoiceMTSPMMGQuery) FirstX(ctx context.Context) *Cadre_Choice_MTSPMMG {
	node, err := ccmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Cadre_Choice_MTSPMMG ID from the query.
// Returns a *NotFoundError when no Cadre_Choice_MTSPMMG ID was found.
func (ccmq *CadreChoiceMTSPMMGQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ccmq.Limit(1).IDs(setContextOp(ctx, ccmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cadre_choice_mtspmmg.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccmq *CadreChoiceMTSPMMGQuery) FirstIDX(ctx context.Context) int32 {
	id, err := ccmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Cadre_Choice_MTSPMMG entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Cadre_Choice_MTSPMMG entity is found.
// Returns a *NotFoundError when no Cadre_Choice_MTSPMMG entities are found.
func (ccmq *CadreChoiceMTSPMMGQuery) Only(ctx context.Context) (*Cadre_Choice_MTSPMMG, error) {
	nodes, err := ccmq.Limit(2).All(setContextOp(ctx, ccmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cadre_choice_mtspmmg.Label}
	default:
		return nil, &NotSingularError{cadre_choice_mtspmmg.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccmq *CadreChoiceMTSPMMGQuery) OnlyX(ctx context.Context) *Cadre_Choice_MTSPMMG {
	node, err := ccmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Cadre_Choice_MTSPMMG ID in the query.
// Returns a *NotSingularError when more than one Cadre_Choice_MTSPMMG ID is found.
// Returns a *NotFoundError when no entities are found.
func (ccmq *CadreChoiceMTSPMMGQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ccmq.Limit(2).IDs(setContextOp(ctx, ccmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cadre_choice_mtspmmg.Label}
	default:
		err = &NotSingularError{cadre_choice_mtspmmg.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccmq *CadreChoiceMTSPMMGQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := ccmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cadre_Choice_MTSPMMGs.
func (ccmq *CadreChoiceMTSPMMGQuery) All(ctx context.Context) ([]*Cadre_Choice_MTSPMMG, error) {
	ctx = setContextOp(ctx, ccmq.ctx, "All")
	if err := ccmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Cadre_Choice_MTSPMMG, *CadreChoiceMTSPMMGQuery]()
	return withInterceptors[[]*Cadre_Choice_MTSPMMG](ctx, ccmq, qr, ccmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ccmq *CadreChoiceMTSPMMGQuery) AllX(ctx context.Context) []*Cadre_Choice_MTSPMMG {
	nodes, err := ccmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Cadre_Choice_MTSPMMG IDs.
func (ccmq *CadreChoiceMTSPMMGQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if ccmq.ctx.Unique == nil && ccmq.path != nil {
		ccmq.Unique(true)
	}
	ctx = setContextOp(ctx, ccmq.ctx, "IDs")
	if err = ccmq.Select(cadre_choice_mtspmmg.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccmq *CadreChoiceMTSPMMGQuery) IDsX(ctx context.Context) []int32 {
	ids, err := ccmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccmq *CadreChoiceMTSPMMGQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ccmq.ctx, "Count")
	if err := ccmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ccmq, querierCount[*CadreChoiceMTSPMMGQuery](), ccmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ccmq *CadreChoiceMTSPMMGQuery) CountX(ctx context.Context) int {
	count, err := ccmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccmq *CadreChoiceMTSPMMGQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ccmq.ctx, "Exist")
	switch _, err := ccmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ccmq *CadreChoiceMTSPMMGQuery) ExistX(ctx context.Context) bool {
	exist, err := ccmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CadreChoiceMTSPMMGQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccmq *CadreChoiceMTSPMMGQuery) Clone() *CadreChoiceMTSPMMGQuery {
	if ccmq == nil {
		return nil
	}
	return &CadreChoiceMTSPMMGQuery{
		config:              ccmq.config,
		ctx:                 ccmq.ctx.Clone(),
		order:               append([]cadre_choice_mtspmmg.OrderOption{}, ccmq.order...),
		inters:              append([]Interceptor{}, ccmq.inters...),
		predicates:          append([]predicate.Cadre_Choice_MTSPMMG{}, ccmq.predicates...),
		withApplnMTSPMMGRef: ccmq.withApplnMTSPMMGRef.Clone(),
		// clone intermediate query.
		sql:  ccmq.sql.Clone(),
		path: ccmq.path,
	}
}

// WithApplnMTSPMMGRef tells the query-builder to eager-load the nodes that are connected to
// the "ApplnMTSPMMG_Ref" edge. The optional arguments are used to configure the query builder of the edge.
func (ccmq *CadreChoiceMTSPMMGQuery) WithApplnMTSPMMGRef(opts ...func(*ExamApplicationMTSPMMGQuery)) *CadreChoiceMTSPMMGQuery {
	query := (&ExamApplicationMTSPMMGClient{config: ccmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ccmq.withApplnMTSPMMGRef = query
	return ccmq
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
//	client.CadreChoiceMTSPMMG.Query().
//		GroupBy(cadre_choice_mtspmmg.FieldApplicationID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ccmq *CadreChoiceMTSPMMGQuery) GroupBy(field string, fields ...string) *CadreChoiceMTSPMMGGroupBy {
	ccmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CadreChoiceMTSPMMGGroupBy{build: ccmq}
	grbuild.flds = &ccmq.ctx.Fields
	grbuild.label = cadre_choice_mtspmmg.Label
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
//	client.CadreChoiceMTSPMMG.Query().
//		Select(cadre_choice_mtspmmg.FieldApplicationID).
//		Scan(ctx, &v)
func (ccmq *CadreChoiceMTSPMMGQuery) Select(fields ...string) *CadreChoiceMTSPMMGSelect {
	ccmq.ctx.Fields = append(ccmq.ctx.Fields, fields...)
	sbuild := &CadreChoiceMTSPMMGSelect{CadreChoiceMTSPMMGQuery: ccmq}
	sbuild.label = cadre_choice_mtspmmg.Label
	sbuild.flds, sbuild.scan = &ccmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CadreChoiceMTSPMMGSelect configured with the given aggregations.
func (ccmq *CadreChoiceMTSPMMGQuery) Aggregate(fns ...AggregateFunc) *CadreChoiceMTSPMMGSelect {
	return ccmq.Select().Aggregate(fns...)
}

func (ccmq *CadreChoiceMTSPMMGQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ccmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ccmq); err != nil {
				return err
			}
		}
	}
	for _, f := range ccmq.ctx.Fields {
		if !cadre_choice_mtspmmg.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ccmq.path != nil {
		prev, err := ccmq.path(ctx)
		if err != nil {
			return err
		}
		ccmq.sql = prev
	}
	return nil
}

func (ccmq *CadreChoiceMTSPMMGQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Cadre_Choice_MTSPMMG, error) {
	var (
		nodes       = []*Cadre_Choice_MTSPMMG{}
		_spec       = ccmq.querySpec()
		loadedTypes = [1]bool{
			ccmq.withApplnMTSPMMGRef != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Cadre_Choice_MTSPMMG).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Cadre_Choice_MTSPMMG{config: ccmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ccmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ccmq.withApplnMTSPMMGRef; query != nil {
		if err := ccmq.loadApplnMTSPMMGRef(ctx, query, nodes, nil,
			func(n *Cadre_Choice_MTSPMMG, e *Exam_Application_MTSPMMG) { n.Edges.ApplnMTSPMMGRef = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ccmq *CadreChoiceMTSPMMGQuery) loadApplnMTSPMMGRef(ctx context.Context, query *ExamApplicationMTSPMMGQuery, nodes []*Cadre_Choice_MTSPMMG, init func(*Cadre_Choice_MTSPMMG), assign func(*Cadre_Choice_MTSPMMG, *Exam_Application_MTSPMMG)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Cadre_Choice_MTSPMMG)
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

func (ccmq *CadreChoiceMTSPMMGQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccmq.querySpec()
	_spec.Node.Columns = ccmq.ctx.Fields
	if len(ccmq.ctx.Fields) > 0 {
		_spec.Unique = ccmq.ctx.Unique != nil && *ccmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ccmq.driver, _spec)
}

func (ccmq *CadreChoiceMTSPMMGQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cadre_choice_mtspmmg.Table, cadre_choice_mtspmmg.Columns, sqlgraph.NewFieldSpec(cadre_choice_mtspmmg.FieldID, field.TypeInt32))
	_spec.From = ccmq.sql
	if unique := ccmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ccmq.path != nil {
		_spec.Unique = true
	}
	if fields := ccmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cadre_choice_mtspmmg.FieldID)
		for i := range fields {
			if fields[i] != cadre_choice_mtspmmg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ccmq.withApplnMTSPMMGRef != nil {
			_spec.Node.AddColumnOnce(cadre_choice_mtspmmg.FieldApplicationID)
		}
	}
	if ps := ccmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccmq *CadreChoiceMTSPMMGQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccmq.driver.Dialect())
	t1 := builder.Table(cadre_choice_mtspmmg.Table)
	columns := ccmq.ctx.Fields
	if len(columns) == 0 {
		columns = cadre_choice_mtspmmg.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccmq.sql != nil {
		selector = ccmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ccmq.ctx.Unique != nil && *ccmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ccmq.predicates {
		p(selector)
	}
	for _, p := range ccmq.order {
		p(selector)
	}
	if offset := ccmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CadreChoiceMTSPMMGGroupBy is the group-by builder for Cadre_Choice_MTSPMMG entities.
type CadreChoiceMTSPMMGGroupBy struct {
	selector
	build *CadreChoiceMTSPMMGQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccmgb *CadreChoiceMTSPMMGGroupBy) Aggregate(fns ...AggregateFunc) *CadreChoiceMTSPMMGGroupBy {
	ccmgb.fns = append(ccmgb.fns, fns...)
	return ccmgb
}

// Scan applies the selector query and scans the result into the given value.
func (ccmgb *CadreChoiceMTSPMMGGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccmgb.build.ctx, "GroupBy")
	if err := ccmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CadreChoiceMTSPMMGQuery, *CadreChoiceMTSPMMGGroupBy](ctx, ccmgb.build, ccmgb, ccmgb.build.inters, v)
}

func (ccmgb *CadreChoiceMTSPMMGGroupBy) sqlScan(ctx context.Context, root *CadreChoiceMTSPMMGQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ccmgb.fns))
	for _, fn := range ccmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ccmgb.flds)+len(ccmgb.fns))
		for _, f := range *ccmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ccmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CadreChoiceMTSPMMGSelect is the builder for selecting fields of CadreChoiceMTSPMMG entities.
type CadreChoiceMTSPMMGSelect struct {
	*CadreChoiceMTSPMMGQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ccms *CadreChoiceMTSPMMGSelect) Aggregate(fns ...AggregateFunc) *CadreChoiceMTSPMMGSelect {
	ccms.fns = append(ccms.fns, fns...)
	return ccms
}

// Scan applies the selector query and scans the result into the given value.
func (ccms *CadreChoiceMTSPMMGSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccms.ctx, "Select")
	if err := ccms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CadreChoiceMTSPMMGQuery, *CadreChoiceMTSPMMGSelect](ctx, ccms.CadreChoiceMTSPMMGQuery, ccms, ccms.inters, v)
}

func (ccms *CadreChoiceMTSPMMGSelect) sqlScan(ctx context.Context, root *CadreChoiceMTSPMMGQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ccms.fns))
	for _, fn := range ccms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ccms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}