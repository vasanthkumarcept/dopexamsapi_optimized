// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/adminlogin"
	"recruit/ent/logs"
	"recruit/ent/predicate"
	"recruit/ent/rolemaster"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminLoginQuery is the builder for querying AdminLogin entities.
type AdminLoginQuery struct {
	config
	ctx            *QueryContext
	order          []adminlogin.OrderOption
	inters         []Interceptor
	predicates     []predicate.AdminLogin
	withRoleMaster *RoleMasterQuery
	withLogData    *LogsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdminLoginQuery builder.
func (alq *AdminLoginQuery) Where(ps ...predicate.AdminLogin) *AdminLoginQuery {
	alq.predicates = append(alq.predicates, ps...)
	return alq
}

// Limit the number of records to be returned by this query.
func (alq *AdminLoginQuery) Limit(limit int) *AdminLoginQuery {
	alq.ctx.Limit = &limit
	return alq
}

// Offset to start from.
func (alq *AdminLoginQuery) Offset(offset int) *AdminLoginQuery {
	alq.ctx.Offset = &offset
	return alq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (alq *AdminLoginQuery) Unique(unique bool) *AdminLoginQuery {
	alq.ctx.Unique = &unique
	return alq
}

// Order specifies how the records should be ordered.
func (alq *AdminLoginQuery) Order(o ...adminlogin.OrderOption) *AdminLoginQuery {
	alq.order = append(alq.order, o...)
	return alq
}

// QueryRoleMaster chains the current query on the "role_master" edge.
func (alq *AdminLoginQuery) QueryRoleMaster() *RoleMasterQuery {
	query := (&RoleMasterClient{config: alq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := alq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := alq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adminlogin.Table, adminlogin.FieldID, selector),
			sqlgraph.To(rolemaster.Table, rolemaster.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, adminlogin.RoleMasterTable, adminlogin.RoleMasterColumn),
		)
		fromU = sqlgraph.SetNeighbors(alq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLogData chains the current query on the "LogData" edge.
func (alq *AdminLoginQuery) QueryLogData() *LogsQuery {
	query := (&LogsClient{config: alq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := alq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := alq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adminlogin.Table, adminlogin.FieldID, selector),
			sqlgraph.To(logs.Table, logs.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adminlogin.LogDataTable, adminlogin.LogDataColumn),
		)
		fromU = sqlgraph.SetNeighbors(alq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AdminLogin entity from the query.
// Returns a *NotFoundError when no AdminLogin was found.
func (alq *AdminLoginQuery) First(ctx context.Context) (*AdminLogin, error) {
	nodes, err := alq.Limit(1).All(setContextOp(ctx, alq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adminlogin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (alq *AdminLoginQuery) FirstX(ctx context.Context) *AdminLogin {
	node, err := alq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdminLogin ID from the query.
// Returns a *NotFoundError when no AdminLogin ID was found.
func (alq *AdminLoginQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = alq.Limit(1).IDs(setContextOp(ctx, alq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adminlogin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (alq *AdminLoginQuery) FirstIDX(ctx context.Context) int32 {
	id, err := alq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdminLogin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AdminLogin entity is found.
// Returns a *NotFoundError when no AdminLogin entities are found.
func (alq *AdminLoginQuery) Only(ctx context.Context) (*AdminLogin, error) {
	nodes, err := alq.Limit(2).All(setContextOp(ctx, alq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adminlogin.Label}
	default:
		return nil, &NotSingularError{adminlogin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (alq *AdminLoginQuery) OnlyX(ctx context.Context) *AdminLogin {
	node, err := alq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdminLogin ID in the query.
// Returns a *NotSingularError when more than one AdminLogin ID is found.
// Returns a *NotFoundError when no entities are found.
func (alq *AdminLoginQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = alq.Limit(2).IDs(setContextOp(ctx, alq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adminlogin.Label}
	default:
		err = &NotSingularError{adminlogin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (alq *AdminLoginQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := alq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdminLogins.
func (alq *AdminLoginQuery) All(ctx context.Context) ([]*AdminLogin, error) {
	ctx = setContextOp(ctx, alq.ctx, "All")
	if err := alq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AdminLogin, *AdminLoginQuery]()
	return withInterceptors[[]*AdminLogin](ctx, alq, qr, alq.inters)
}

// AllX is like All, but panics if an error occurs.
func (alq *AdminLoginQuery) AllX(ctx context.Context) []*AdminLogin {
	nodes, err := alq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdminLogin IDs.
func (alq *AdminLoginQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if alq.ctx.Unique == nil && alq.path != nil {
		alq.Unique(true)
	}
	ctx = setContextOp(ctx, alq.ctx, "IDs")
	if err = alq.Select(adminlogin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (alq *AdminLoginQuery) IDsX(ctx context.Context) []int32 {
	ids, err := alq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (alq *AdminLoginQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, alq.ctx, "Count")
	if err := alq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, alq, querierCount[*AdminLoginQuery](), alq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (alq *AdminLoginQuery) CountX(ctx context.Context) int {
	count, err := alq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (alq *AdminLoginQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, alq.ctx, "Exist")
	switch _, err := alq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (alq *AdminLoginQuery) ExistX(ctx context.Context) bool {
	exist, err := alq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdminLoginQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (alq *AdminLoginQuery) Clone() *AdminLoginQuery {
	if alq == nil {
		return nil
	}
	return &AdminLoginQuery{
		config:         alq.config,
		ctx:            alq.ctx.Clone(),
		order:          append([]adminlogin.OrderOption{}, alq.order...),
		inters:         append([]Interceptor{}, alq.inters...),
		predicates:     append([]predicate.AdminLogin{}, alq.predicates...),
		withRoleMaster: alq.withRoleMaster.Clone(),
		withLogData:    alq.withLogData.Clone(),
		// clone intermediate query.
		sql:  alq.sql.Clone(),
		path: alq.path,
	}
}

// WithRoleMaster tells the query-builder to eager-load the nodes that are connected to
// the "role_master" edge. The optional arguments are used to configure the query builder of the edge.
func (alq *AdminLoginQuery) WithRoleMaster(opts ...func(*RoleMasterQuery)) *AdminLoginQuery {
	query := (&RoleMasterClient{config: alq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	alq.withRoleMaster = query
	return alq
}

// WithLogData tells the query-builder to eager-load the nodes that are connected to
// the "LogData" edge. The optional arguments are used to configure the query builder of the edge.
func (alq *AdminLoginQuery) WithLogData(opts ...func(*LogsQuery)) *AdminLoginQuery {
	query := (&LogsClient{config: alq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	alq.withLogData = query
	return alq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RoleUserCode int32 `json:"RoleUserCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdminLogin.Query().
//		GroupBy(adminlogin.FieldRoleUserCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (alq *AdminLoginQuery) GroupBy(field string, fields ...string) *AdminLoginGroupBy {
	alq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AdminLoginGroupBy{build: alq}
	grbuild.flds = &alq.ctx.Fields
	grbuild.label = adminlogin.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RoleUserCode int32 `json:"RoleUserCode,omitempty"`
//	}
//
//	client.AdminLogin.Query().
//		Select(adminlogin.FieldRoleUserCode).
//		Scan(ctx, &v)
func (alq *AdminLoginQuery) Select(fields ...string) *AdminLoginSelect {
	alq.ctx.Fields = append(alq.ctx.Fields, fields...)
	sbuild := &AdminLoginSelect{AdminLoginQuery: alq}
	sbuild.label = adminlogin.Label
	sbuild.flds, sbuild.scan = &alq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AdminLoginSelect configured with the given aggregations.
func (alq *AdminLoginQuery) Aggregate(fns ...AggregateFunc) *AdminLoginSelect {
	return alq.Select().Aggregate(fns...)
}

func (alq *AdminLoginQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range alq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, alq); err != nil {
				return err
			}
		}
	}
	for _, f := range alq.ctx.Fields {
		if !adminlogin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if alq.path != nil {
		prev, err := alq.path(ctx)
		if err != nil {
			return err
		}
		alq.sql = prev
	}
	return nil
}

func (alq *AdminLoginQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AdminLogin, error) {
	var (
		nodes       = []*AdminLogin{}
		_spec       = alq.querySpec()
		loadedTypes = [2]bool{
			alq.withRoleMaster != nil,
			alq.withLogData != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AdminLogin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AdminLogin{config: alq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, alq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := alq.withRoleMaster; query != nil {
		if err := alq.loadRoleMaster(ctx, query, nodes, nil,
			func(n *AdminLogin, e *RoleMaster) { n.Edges.RoleMaster = e }); err != nil {
			return nil, err
		}
	}
	if query := alq.withLogData; query != nil {
		if err := alq.loadLogData(ctx, query, nodes,
			func(n *AdminLogin) { n.Edges.LogData = []*Logs{} },
			func(n *AdminLogin, e *Logs) { n.Edges.LogData = append(n.Edges.LogData, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (alq *AdminLoginQuery) loadRoleMaster(ctx context.Context, query *RoleMasterQuery, nodes []*AdminLogin, init func(*AdminLogin), assign func(*AdminLogin, *RoleMaster)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*AdminLogin)
	for i := range nodes {
		fk := nodes[i].RoleUserCode
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(rolemaster.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "RoleUserCode" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (alq *AdminLoginQuery) loadLogData(ctx context.Context, query *LogsQuery, nodes []*AdminLogin, init func(*AdminLogin), assign func(*AdminLogin, *Logs)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*AdminLogin)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Logs(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(adminlogin.LogDataColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.admin_login_log_data
		if fk == nil {
			return fmt.Errorf(`foreign-key "admin_login_log_data" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "admin_login_log_data" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (alq *AdminLoginQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := alq.querySpec()
	_spec.Node.Columns = alq.ctx.Fields
	if len(alq.ctx.Fields) > 0 {
		_spec.Unique = alq.ctx.Unique != nil && *alq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, alq.driver, _spec)
}

func (alq *AdminLoginQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(adminlogin.Table, adminlogin.Columns, sqlgraph.NewFieldSpec(adminlogin.FieldID, field.TypeInt32))
	_spec.From = alq.sql
	if unique := alq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if alq.path != nil {
		_spec.Unique = true
	}
	if fields := alq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminlogin.FieldID)
		for i := range fields {
			if fields[i] != adminlogin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if alq.withRoleMaster != nil {
			_spec.Node.AddColumnOnce(adminlogin.FieldRoleUserCode)
		}
	}
	if ps := alq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := alq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := alq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := alq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (alq *AdminLoginQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(alq.driver.Dialect())
	t1 := builder.Table(adminlogin.Table)
	columns := alq.ctx.Fields
	if len(columns) == 0 {
		columns = adminlogin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if alq.sql != nil {
		selector = alq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if alq.ctx.Unique != nil && *alq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range alq.predicates {
		p(selector)
	}
	for _, p := range alq.order {
		p(selector)
	}
	if offset := alq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := alq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdminLoginGroupBy is the group-by builder for AdminLogin entities.
type AdminLoginGroupBy struct {
	selector
	build *AdminLoginQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (algb *AdminLoginGroupBy) Aggregate(fns ...AggregateFunc) *AdminLoginGroupBy {
	algb.fns = append(algb.fns, fns...)
	return algb
}

// Scan applies the selector query and scans the result into the given value.
func (algb *AdminLoginGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, algb.build.ctx, "GroupBy")
	if err := algb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminLoginQuery, *AdminLoginGroupBy](ctx, algb.build, algb, algb.build.inters, v)
}

func (algb *AdminLoginGroupBy) sqlScan(ctx context.Context, root *AdminLoginQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(algb.fns))
	for _, fn := range algb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*algb.flds)+len(algb.fns))
		for _, f := range *algb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*algb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := algb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AdminLoginSelect is the builder for selecting fields of AdminLogin entities.
type AdminLoginSelect struct {
	*AdminLoginQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (als *AdminLoginSelect) Aggregate(fns ...AggregateFunc) *AdminLoginSelect {
	als.fns = append(als.fns, fns...)
	return als
}

// Scan applies the selector query and scans the result into the given value.
func (als *AdminLoginSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, als.ctx, "Select")
	if err := als.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminLoginQuery, *AdminLoginSelect](ctx, als.AdminLoginQuery, als, als.inters, v)
}

func (als *AdminLoginSelect) sqlScan(ctx context.Context, root *AdminLoginQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(als.fns))
	for _, fn := range als.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*als.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := als.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
