// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/postexampaper"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostExamPaperQuery is the builder for querying PostExamPaper entities.
type PostExamPaperQuery struct {
	config
	ctx        *QueryContext
	order      []postexampaper.OrderOption
	inters     []Interceptor
	predicates []predicate.PostExamPaper
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PostExamPaperQuery builder.
func (pepq *PostExamPaperQuery) Where(ps ...predicate.PostExamPaper) *PostExamPaperQuery {
	pepq.predicates = append(pepq.predicates, ps...)
	return pepq
}

// Limit the number of records to be returned by this query.
func (pepq *PostExamPaperQuery) Limit(limit int) *PostExamPaperQuery {
	pepq.ctx.Limit = &limit
	return pepq
}

// Offset to start from.
func (pepq *PostExamPaperQuery) Offset(offset int) *PostExamPaperQuery {
	pepq.ctx.Offset = &offset
	return pepq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pepq *PostExamPaperQuery) Unique(unique bool) *PostExamPaperQuery {
	pepq.ctx.Unique = &unique
	return pepq
}

// Order specifies how the records should be ordered.
func (pepq *PostExamPaperQuery) Order(o ...postexampaper.OrderOption) *PostExamPaperQuery {
	pepq.order = append(pepq.order, o...)
	return pepq
}

// First returns the first PostExamPaper entity from the query.
// Returns a *NotFoundError when no PostExamPaper was found.
func (pepq *PostExamPaperQuery) First(ctx context.Context) (*PostExamPaper, error) {
	nodes, err := pepq.Limit(1).All(setContextOp(ctx, pepq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{postexampaper.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pepq *PostExamPaperQuery) FirstX(ctx context.Context) *PostExamPaper {
	node, err := pepq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PostExamPaper ID from the query.
// Returns a *NotFoundError when no PostExamPaper ID was found.
func (pepq *PostExamPaperQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = pepq.Limit(1).IDs(setContextOp(ctx, pepq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{postexampaper.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pepq *PostExamPaperQuery) FirstIDX(ctx context.Context) int32 {
	id, err := pepq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PostExamPaper entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PostExamPaper entity is found.
// Returns a *NotFoundError when no PostExamPaper entities are found.
func (pepq *PostExamPaperQuery) Only(ctx context.Context) (*PostExamPaper, error) {
	nodes, err := pepq.Limit(2).All(setContextOp(ctx, pepq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{postexampaper.Label}
	default:
		return nil, &NotSingularError{postexampaper.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pepq *PostExamPaperQuery) OnlyX(ctx context.Context) *PostExamPaper {
	node, err := pepq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PostExamPaper ID in the query.
// Returns a *NotSingularError when more than one PostExamPaper ID is found.
// Returns a *NotFoundError when no entities are found.
func (pepq *PostExamPaperQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = pepq.Limit(2).IDs(setContextOp(ctx, pepq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{postexampaper.Label}
	default:
		err = &NotSingularError{postexampaper.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pepq *PostExamPaperQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := pepq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PostExamPapers.
func (pepq *PostExamPaperQuery) All(ctx context.Context) ([]*PostExamPaper, error) {
	ctx = setContextOp(ctx, pepq.ctx, "All")
	if err := pepq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PostExamPaper, *PostExamPaperQuery]()
	return withInterceptors[[]*PostExamPaper](ctx, pepq, qr, pepq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pepq *PostExamPaperQuery) AllX(ctx context.Context) []*PostExamPaper {
	nodes, err := pepq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PostExamPaper IDs.
func (pepq *PostExamPaperQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if pepq.ctx.Unique == nil && pepq.path != nil {
		pepq.Unique(true)
	}
	ctx = setContextOp(ctx, pepq.ctx, "IDs")
	if err = pepq.Select(postexampaper.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pepq *PostExamPaperQuery) IDsX(ctx context.Context) []int32 {
	ids, err := pepq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pepq *PostExamPaperQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pepq.ctx, "Count")
	if err := pepq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pepq, querierCount[*PostExamPaperQuery](), pepq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pepq *PostExamPaperQuery) CountX(ctx context.Context) int {
	count, err := pepq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pepq *PostExamPaperQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pepq.ctx, "Exist")
	switch _, err := pepq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pepq *PostExamPaperQuery) ExistX(ctx context.Context) bool {
	exist, err := pepq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PostExamPaperQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pepq *PostExamPaperQuery) Clone() *PostExamPaperQuery {
	if pepq == nil {
		return nil
	}
	return &PostExamPaperQuery{
		config:     pepq.config,
		ctx:        pepq.ctx.Clone(),
		order:      append([]postexampaper.OrderOption{}, pepq.order...),
		inters:     append([]Interceptor{}, pepq.inters...),
		predicates: append([]predicate.PostExamPaper{}, pepq.predicates...),
		// clone intermediate query.
		sql:  pepq.sql.Clone(),
		path: pepq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ExamConfigurationExamCode int32 `json:"ExamConfigurationExamCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PostExamPaper.Query().
//		GroupBy(postexampaper.FieldExamConfigurationExamCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pepq *PostExamPaperQuery) GroupBy(field string, fields ...string) *PostExamPaperGroupBy {
	pepq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PostExamPaperGroupBy{build: pepq}
	grbuild.flds = &pepq.ctx.Fields
	grbuild.label = postexampaper.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ExamConfigurationExamCode int32 `json:"ExamConfigurationExamCode,omitempty"`
//	}
//
//	client.PostExamPaper.Query().
//		Select(postexampaper.FieldExamConfigurationExamCode).
//		Scan(ctx, &v)
func (pepq *PostExamPaperQuery) Select(fields ...string) *PostExamPaperSelect {
	pepq.ctx.Fields = append(pepq.ctx.Fields, fields...)
	sbuild := &PostExamPaperSelect{PostExamPaperQuery: pepq}
	sbuild.label = postexampaper.Label
	sbuild.flds, sbuild.scan = &pepq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PostExamPaperSelect configured with the given aggregations.
func (pepq *PostExamPaperQuery) Aggregate(fns ...AggregateFunc) *PostExamPaperSelect {
	return pepq.Select().Aggregate(fns...)
}

func (pepq *PostExamPaperQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pepq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pepq); err != nil {
				return err
			}
		}
	}
	for _, f := range pepq.ctx.Fields {
		if !postexampaper.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pepq.path != nil {
		prev, err := pepq.path(ctx)
		if err != nil {
			return err
		}
		pepq.sql = prev
	}
	return nil
}

func (pepq *PostExamPaperQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PostExamPaper, error) {
	var (
		nodes = []*PostExamPaper{}
		_spec = pepq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PostExamPaper).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PostExamPaper{config: pepq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pepq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pepq *PostExamPaperQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pepq.querySpec()
	_spec.Node.Columns = pepq.ctx.Fields
	if len(pepq.ctx.Fields) > 0 {
		_spec.Unique = pepq.ctx.Unique != nil && *pepq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pepq.driver, _spec)
}

func (pepq *PostExamPaperQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(postexampaper.Table, postexampaper.Columns, sqlgraph.NewFieldSpec(postexampaper.FieldID, field.TypeInt32))
	_spec.From = pepq.sql
	if unique := pepq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pepq.path != nil {
		_spec.Unique = true
	}
	if fields := pepq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, postexampaper.FieldID)
		for i := range fields {
			if fields[i] != postexampaper.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pepq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pepq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pepq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pepq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pepq *PostExamPaperQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pepq.driver.Dialect())
	t1 := builder.Table(postexampaper.Table)
	columns := pepq.ctx.Fields
	if len(columns) == 0 {
		columns = postexampaper.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pepq.sql != nil {
		selector = pepq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pepq.ctx.Unique != nil && *pepq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pepq.predicates {
		p(selector)
	}
	for _, p := range pepq.order {
		p(selector)
	}
	if offset := pepq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pepq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PostExamPaperGroupBy is the group-by builder for PostExamPaper entities.
type PostExamPaperGroupBy struct {
	selector
	build *PostExamPaperQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pepgb *PostExamPaperGroupBy) Aggregate(fns ...AggregateFunc) *PostExamPaperGroupBy {
	pepgb.fns = append(pepgb.fns, fns...)
	return pepgb
}

// Scan applies the selector query and scans the result into the given value.
func (pepgb *PostExamPaperGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pepgb.build.ctx, "GroupBy")
	if err := pepgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostExamPaperQuery, *PostExamPaperGroupBy](ctx, pepgb.build, pepgb, pepgb.build.inters, v)
}

func (pepgb *PostExamPaperGroupBy) sqlScan(ctx context.Context, root *PostExamPaperQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pepgb.fns))
	for _, fn := range pepgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pepgb.flds)+len(pepgb.fns))
		for _, f := range *pepgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pepgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pepgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PostExamPaperSelect is the builder for selecting fields of PostExamPaper entities.
type PostExamPaperSelect struct {
	*PostExamPaperQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (peps *PostExamPaperSelect) Aggregate(fns ...AggregateFunc) *PostExamPaperSelect {
	peps.fns = append(peps.fns, fns...)
	return peps
}

// Scan applies the selector query and scans the result into the given value.
func (peps *PostExamPaperSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, peps.ctx, "Select")
	if err := peps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PostExamPaperQuery, *PostExamPaperSelect](ctx, peps.PostExamPaperQuery, peps, peps.inters, v)
}

func (peps *PostExamPaperSelect) sqlScan(ctx context.Context, root *PostExamPaperQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(peps.fns))
	for _, fn := range peps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*peps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := peps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}