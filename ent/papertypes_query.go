// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/exampapers"
	"recruit/ent/papertypes"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaperTypesQuery is the builder for querying PaperTypes entities.
type PaperTypesQuery struct {
	config
	ctx           *QueryContext
	order         []papertypes.OrderOption
	inters        []Interceptor
	predicates    []predicate.PaperTypes
	withPapercode *ExamPapersQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PaperTypesQuery builder.
func (ptq *PaperTypesQuery) Where(ps ...predicate.PaperTypes) *PaperTypesQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit the number of records to be returned by this query.
func (ptq *PaperTypesQuery) Limit(limit int) *PaperTypesQuery {
	ptq.ctx.Limit = &limit
	return ptq
}

// Offset to start from.
func (ptq *PaperTypesQuery) Offset(offset int) *PaperTypesQuery {
	ptq.ctx.Offset = &offset
	return ptq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptq *PaperTypesQuery) Unique(unique bool) *PaperTypesQuery {
	ptq.ctx.Unique = &unique
	return ptq
}

// Order specifies how the records should be ordered.
func (ptq *PaperTypesQuery) Order(o ...papertypes.OrderOption) *PaperTypesQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryPapercode chains the current query on the "papercode" edge.
func (ptq *PaperTypesQuery) QueryPapercode() *ExamPapersQuery {
	query := (&ExamPapersClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(papertypes.Table, papertypes.FieldID, selector),
			sqlgraph.To(exampapers.Table, exampapers.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, papertypes.PapercodeTable, papertypes.PapercodeColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PaperTypes entity from the query.
// Returns a *NotFoundError when no PaperTypes was found.
func (ptq *PaperTypesQuery) First(ctx context.Context) (*PaperTypes, error) {
	nodes, err := ptq.Limit(1).All(setContextOp(ctx, ptq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{papertypes.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *PaperTypesQuery) FirstX(ctx context.Context) *PaperTypes {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PaperTypes ID from the query.
// Returns a *NotFoundError when no PaperTypes ID was found.
func (ptq *PaperTypesQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ptq.Limit(1).IDs(setContextOp(ctx, ptq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{papertypes.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *PaperTypesQuery) FirstIDX(ctx context.Context) int32 {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PaperTypes entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PaperTypes entity is found.
// Returns a *NotFoundError when no PaperTypes entities are found.
func (ptq *PaperTypesQuery) Only(ctx context.Context) (*PaperTypes, error) {
	nodes, err := ptq.Limit(2).All(setContextOp(ctx, ptq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{papertypes.Label}
	default:
		return nil, &NotSingularError{papertypes.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *PaperTypesQuery) OnlyX(ctx context.Context) *PaperTypes {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PaperTypes ID in the query.
// Returns a *NotSingularError when more than one PaperTypes ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptq *PaperTypesQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ptq.Limit(2).IDs(setContextOp(ctx, ptq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{papertypes.Label}
	default:
		err = &NotSingularError{papertypes.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *PaperTypesQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PaperTypesSlice.
func (ptq *PaperTypesQuery) All(ctx context.Context) ([]*PaperTypes, error) {
	ctx = setContextOp(ctx, ptq.ctx, "All")
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PaperTypes, *PaperTypesQuery]()
	return withInterceptors[[]*PaperTypes](ctx, ptq, qr, ptq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptq *PaperTypesQuery) AllX(ctx context.Context) []*PaperTypes {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PaperTypes IDs.
func (ptq *PaperTypesQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if ptq.ctx.Unique == nil && ptq.path != nil {
		ptq.Unique(true)
	}
	ctx = setContextOp(ctx, ptq.ctx, "IDs")
	if err = ptq.Select(papertypes.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *PaperTypesQuery) IDsX(ctx context.Context) []int32 {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *PaperTypesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Count")
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptq, querierCount[*PaperTypesQuery](), ptq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *PaperTypesQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *PaperTypesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptq.ctx, "Exist")
	switch _, err := ptq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *PaperTypesQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PaperTypesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *PaperTypesQuery) Clone() *PaperTypesQuery {
	if ptq == nil {
		return nil
	}
	return &PaperTypesQuery{
		config:        ptq.config,
		ctx:           ptq.ctx.Clone(),
		order:         append([]papertypes.OrderOption{}, ptq.order...),
		inters:        append([]Interceptor{}, ptq.inters...),
		predicates:    append([]predicate.PaperTypes{}, ptq.predicates...),
		withPapercode: ptq.withPapercode.Clone(),
		// clone intermediate query.
		sql:  ptq.sql.Clone(),
		path: ptq.path,
	}
}

// WithPapercode tells the query-builder to eager-load the nodes that are connected to
// the "papercode" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *PaperTypesQuery) WithPapercode(opts ...func(*ExamPapersQuery)) *PaperTypesQuery {
	query := (&ExamPapersClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withPapercode = query
	return ptq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PaperTypeDescription string `json:"PaperTypeDescription,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PaperTypes.Query().
//		GroupBy(papertypes.FieldPaperTypeDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ptq *PaperTypesQuery) GroupBy(field string, fields ...string) *PaperTypesGroupBy {
	ptq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PaperTypesGroupBy{build: ptq}
	grbuild.flds = &ptq.ctx.Fields
	grbuild.label = papertypes.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PaperTypeDescription string `json:"PaperTypeDescription,omitempty"`
//	}
//
//	client.PaperTypes.Query().
//		Select(papertypes.FieldPaperTypeDescription).
//		Scan(ctx, &v)
func (ptq *PaperTypesQuery) Select(fields ...string) *PaperTypesSelect {
	ptq.ctx.Fields = append(ptq.ctx.Fields, fields...)
	sbuild := &PaperTypesSelect{PaperTypesQuery: ptq}
	sbuild.label = papertypes.Label
	sbuild.flds, sbuild.scan = &ptq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PaperTypesSelect configured with the given aggregations.
func (ptq *PaperTypesQuery) Aggregate(fns ...AggregateFunc) *PaperTypesSelect {
	return ptq.Select().Aggregate(fns...)
}

func (ptq *PaperTypesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptq.ctx.Fields {
		if !papertypes.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *PaperTypesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PaperTypes, error) {
	var (
		nodes       = []*PaperTypes{}
		_spec       = ptq.querySpec()
		loadedTypes = [1]bool{
			ptq.withPapercode != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PaperTypes).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PaperTypes{config: ptq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ptq.withPapercode; query != nil {
		if err := ptq.loadPapercode(ctx, query, nodes, nil,
			func(n *PaperTypes, e *ExamPapers) { n.Edges.Papercode = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ptq *PaperTypesQuery) loadPapercode(ctx context.Context, query *ExamPapersQuery, nodes []*PaperTypes, init func(*PaperTypes), assign func(*PaperTypes, *ExamPapers)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*PaperTypes)
	for i := range nodes {
		fk := nodes[i].PaperCode
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(exampapers.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "PaperCode" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ptq *PaperTypesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	_spec.Node.Columns = ptq.ctx.Fields
	if len(ptq.ctx.Fields) > 0 {
		_spec.Unique = ptq.ctx.Unique != nil && *ptq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *PaperTypesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(papertypes.Table, papertypes.Columns, sqlgraph.NewFieldSpec(papertypes.FieldID, field.TypeInt32))
	_spec.From = ptq.sql
	if unique := ptq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptq.path != nil {
		_spec.Unique = true
	}
	if fields := ptq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, papertypes.FieldID)
		for i := range fields {
			if fields[i] != papertypes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ptq.withPapercode != nil {
			_spec.Node.AddColumnOnce(papertypes.FieldPaperCode)
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *PaperTypesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(papertypes.Table)
	columns := ptq.ctx.Fields
	if len(columns) == 0 {
		columns = papertypes.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptq.ctx.Unique != nil && *ptq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PaperTypesGroupBy is the group-by builder for PaperTypes entities.
type PaperTypesGroupBy struct {
	selector
	build *PaperTypesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *PaperTypesGroupBy) Aggregate(fns ...AggregateFunc) *PaperTypesGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptgb *PaperTypesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptgb.build.ctx, "GroupBy")
	if err := ptgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaperTypesQuery, *PaperTypesGroupBy](ctx, ptgb.build, ptgb, ptgb.build.inters, v)
}

func (ptgb *PaperTypesGroupBy) sqlScan(ctx context.Context, root *PaperTypesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptgb.fns))
	for _, fn := range ptgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptgb.flds)+len(ptgb.fns))
		for _, f := range *ptgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PaperTypesSelect is the builder for selecting fields of PaperTypes entities.
type PaperTypesSelect struct {
	*PaperTypesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pts *PaperTypesSelect) Aggregate(fns ...AggregateFunc) *PaperTypesSelect {
	pts.fns = append(pts.fns, fns...)
	return pts
}

// Scan applies the selector query and scans the result into the given value.
func (pts *PaperTypesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pts.ctx, "Select")
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaperTypesQuery, *PaperTypesSelect](ctx, pts.PaperTypesQuery, pts, pts.inters, v)
}

func (pts *PaperTypesSelect) sqlScan(ctx context.Context, root *PaperTypesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pts.fns))
	for _, fn := range pts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}