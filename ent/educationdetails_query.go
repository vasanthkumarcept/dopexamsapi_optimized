// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/educationdetails"
	"recruit/ent/logs"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EducationDetailsQuery is the builder for querying EducationDetails entities.
type EducationDetailsQuery struct {
	config
	ctx         *QueryContext
	order       []educationdetails.OrderOption
	inters      []Interceptor
	predicates  []predicate.EducationDetails
	withLogData *LogsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EducationDetailsQuery builder.
func (edq *EducationDetailsQuery) Where(ps ...predicate.EducationDetails) *EducationDetailsQuery {
	edq.predicates = append(edq.predicates, ps...)
	return edq
}

// Limit the number of records to be returned by this query.
func (edq *EducationDetailsQuery) Limit(limit int) *EducationDetailsQuery {
	edq.ctx.Limit = &limit
	return edq
}

// Offset to start from.
func (edq *EducationDetailsQuery) Offset(offset int) *EducationDetailsQuery {
	edq.ctx.Offset = &offset
	return edq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (edq *EducationDetailsQuery) Unique(unique bool) *EducationDetailsQuery {
	edq.ctx.Unique = &unique
	return edq
}

// Order specifies how the records should be ordered.
func (edq *EducationDetailsQuery) Order(o ...educationdetails.OrderOption) *EducationDetailsQuery {
	edq.order = append(edq.order, o...)
	return edq
}

// QueryLogData chains the current query on the "LogData" edge.
func (edq *EducationDetailsQuery) QueryLogData() *LogsQuery {
	query := (&LogsClient{config: edq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := edq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := edq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(educationdetails.Table, educationdetails.FieldID, selector),
			sqlgraph.To(logs.Table, logs.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, educationdetails.LogDataTable, educationdetails.LogDataColumn),
		)
		fromU = sqlgraph.SetNeighbors(edq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EducationDetails entity from the query.
// Returns a *NotFoundError when no EducationDetails was found.
func (edq *EducationDetailsQuery) First(ctx context.Context) (*EducationDetails, error) {
	nodes, err := edq.Limit(1).All(setContextOp(ctx, edq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{educationdetails.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (edq *EducationDetailsQuery) FirstX(ctx context.Context) *EducationDetails {
	node, err := edq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EducationDetails ID from the query.
// Returns a *NotFoundError when no EducationDetails ID was found.
func (edq *EducationDetailsQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = edq.Limit(1).IDs(setContextOp(ctx, edq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{educationdetails.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (edq *EducationDetailsQuery) FirstIDX(ctx context.Context) int64 {
	id, err := edq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EducationDetails entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EducationDetails entity is found.
// Returns a *NotFoundError when no EducationDetails entities are found.
func (edq *EducationDetailsQuery) Only(ctx context.Context) (*EducationDetails, error) {
	nodes, err := edq.Limit(2).All(setContextOp(ctx, edq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{educationdetails.Label}
	default:
		return nil, &NotSingularError{educationdetails.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (edq *EducationDetailsQuery) OnlyX(ctx context.Context) *EducationDetails {
	node, err := edq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EducationDetails ID in the query.
// Returns a *NotSingularError when more than one EducationDetails ID is found.
// Returns a *NotFoundError when no entities are found.
func (edq *EducationDetailsQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = edq.Limit(2).IDs(setContextOp(ctx, edq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{educationdetails.Label}
	default:
		err = &NotSingularError{educationdetails.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (edq *EducationDetailsQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := edq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EducationDetailsSlice.
func (edq *EducationDetailsQuery) All(ctx context.Context) ([]*EducationDetails, error) {
	ctx = setContextOp(ctx, edq.ctx, "All")
	if err := edq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EducationDetails, *EducationDetailsQuery]()
	return withInterceptors[[]*EducationDetails](ctx, edq, qr, edq.inters)
}

// AllX is like All, but panics if an error occurs.
func (edq *EducationDetailsQuery) AllX(ctx context.Context) []*EducationDetails {
	nodes, err := edq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EducationDetails IDs.
func (edq *EducationDetailsQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if edq.ctx.Unique == nil && edq.path != nil {
		edq.Unique(true)
	}
	ctx = setContextOp(ctx, edq.ctx, "IDs")
	if err = edq.Select(educationdetails.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (edq *EducationDetailsQuery) IDsX(ctx context.Context) []int64 {
	ids, err := edq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (edq *EducationDetailsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, edq.ctx, "Count")
	if err := edq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, edq, querierCount[*EducationDetailsQuery](), edq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (edq *EducationDetailsQuery) CountX(ctx context.Context) int {
	count, err := edq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (edq *EducationDetailsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, edq.ctx, "Exist")
	switch _, err := edq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (edq *EducationDetailsQuery) ExistX(ctx context.Context) bool {
	exist, err := edq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EducationDetailsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (edq *EducationDetailsQuery) Clone() *EducationDetailsQuery {
	if edq == nil {
		return nil
	}
	return &EducationDetailsQuery{
		config:      edq.config,
		ctx:         edq.ctx.Clone(),
		order:       append([]educationdetails.OrderOption{}, edq.order...),
		inters:      append([]Interceptor{}, edq.inters...),
		predicates:  append([]predicate.EducationDetails{}, edq.predicates...),
		withLogData: edq.withLogData.Clone(),
		// clone intermediate query.
		sql:  edq.sql.Clone(),
		path: edq.path,
	}
}

// WithLogData tells the query-builder to eager-load the nodes that are connected to
// the "LogData" edge. The optional arguments are used to configure the query builder of the edge.
func (edq *EducationDetailsQuery) WithLogData(opts ...func(*LogsQuery)) *EducationDetailsQuery {
	query := (&LogsClient{config: edq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	edq.withLogData = query
	return edq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EducationDescription string `json:"educationDescription,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EducationDetails.Query().
//		GroupBy(educationdetails.FieldEducationDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (edq *EducationDetailsQuery) GroupBy(field string, fields ...string) *EducationDetailsGroupBy {
	edq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EducationDetailsGroupBy{build: edq}
	grbuild.flds = &edq.ctx.Fields
	grbuild.label = educationdetails.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EducationDescription string `json:"educationDescription,omitempty"`
//	}
//
//	client.EducationDetails.Query().
//		Select(educationdetails.FieldEducationDescription).
//		Scan(ctx, &v)
func (edq *EducationDetailsQuery) Select(fields ...string) *EducationDetailsSelect {
	edq.ctx.Fields = append(edq.ctx.Fields, fields...)
	sbuild := &EducationDetailsSelect{EducationDetailsQuery: edq}
	sbuild.label = educationdetails.Label
	sbuild.flds, sbuild.scan = &edq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EducationDetailsSelect configured with the given aggregations.
func (edq *EducationDetailsQuery) Aggregate(fns ...AggregateFunc) *EducationDetailsSelect {
	return edq.Select().Aggregate(fns...)
}

func (edq *EducationDetailsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range edq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, edq); err != nil {
				return err
			}
		}
	}
	for _, f := range edq.ctx.Fields {
		if !educationdetails.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if edq.path != nil {
		prev, err := edq.path(ctx)
		if err != nil {
			return err
		}
		edq.sql = prev
	}
	return nil
}

func (edq *EducationDetailsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EducationDetails, error) {
	var (
		nodes       = []*EducationDetails{}
		_spec       = edq.querySpec()
		loadedTypes = [1]bool{
			edq.withLogData != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EducationDetails).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EducationDetails{config: edq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, edq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := edq.withLogData; query != nil {
		if err := edq.loadLogData(ctx, query, nodes,
			func(n *EducationDetails) { n.Edges.LogData = []*Logs{} },
			func(n *EducationDetails, e *Logs) { n.Edges.LogData = append(n.Edges.LogData, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (edq *EducationDetailsQuery) loadLogData(ctx context.Context, query *LogsQuery, nodes []*EducationDetails, init func(*EducationDetails), assign func(*EducationDetails, *Logs)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*EducationDetails)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Logs(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(educationdetails.LogDataColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.education_details_log_data
		if fk == nil {
			return fmt.Errorf(`foreign-key "education_details_log_data" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "education_details_log_data" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (edq *EducationDetailsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := edq.querySpec()
	_spec.Node.Columns = edq.ctx.Fields
	if len(edq.ctx.Fields) > 0 {
		_spec.Unique = edq.ctx.Unique != nil && *edq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, edq.driver, _spec)
}

func (edq *EducationDetailsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(educationdetails.Table, educationdetails.Columns, sqlgraph.NewFieldSpec(educationdetails.FieldID, field.TypeInt64))
	_spec.From = edq.sql
	if unique := edq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if edq.path != nil {
		_spec.Unique = true
	}
	if fields := edq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, educationdetails.FieldID)
		for i := range fields {
			if fields[i] != educationdetails.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := edq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := edq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := edq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := edq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (edq *EducationDetailsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(edq.driver.Dialect())
	t1 := builder.Table(educationdetails.Table)
	columns := edq.ctx.Fields
	if len(columns) == 0 {
		columns = educationdetails.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if edq.sql != nil {
		selector = edq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if edq.ctx.Unique != nil && *edq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range edq.predicates {
		p(selector)
	}
	for _, p := range edq.order {
		p(selector)
	}
	if offset := edq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := edq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EducationDetailsGroupBy is the group-by builder for EducationDetails entities.
type EducationDetailsGroupBy struct {
	selector
	build *EducationDetailsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (edgb *EducationDetailsGroupBy) Aggregate(fns ...AggregateFunc) *EducationDetailsGroupBy {
	edgb.fns = append(edgb.fns, fns...)
	return edgb
}

// Scan applies the selector query and scans the result into the given value.
func (edgb *EducationDetailsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, edgb.build.ctx, "GroupBy")
	if err := edgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EducationDetailsQuery, *EducationDetailsGroupBy](ctx, edgb.build, edgb, edgb.build.inters, v)
}

func (edgb *EducationDetailsGroupBy) sqlScan(ctx context.Context, root *EducationDetailsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(edgb.fns))
	for _, fn := range edgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*edgb.flds)+len(edgb.fns))
		for _, f := range *edgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*edgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := edgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EducationDetailsSelect is the builder for selecting fields of EducationDetails entities.
type EducationDetailsSelect struct {
	*EducationDetailsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eds *EducationDetailsSelect) Aggregate(fns ...AggregateFunc) *EducationDetailsSelect {
	eds.fns = append(eds.fns, fns...)
	return eds
}

// Scan applies the selector query and scans the result into the given value.
func (eds *EducationDetailsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eds.ctx, "Select")
	if err := eds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EducationDetailsQuery, *EducationDetailsSelect](ctx, eds.EducationDetailsQuery, eds, eds.inters, v)
}

func (eds *EducationDetailsSelect) sqlScan(ctx context.Context, root *EducationDetailsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eds.fns))
	for _, fn := range eds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
