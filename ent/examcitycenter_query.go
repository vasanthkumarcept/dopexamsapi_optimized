// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/exam_applications_gdspa"
	"recruit/ent/exam_applications_gdspm"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_pmpa"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/examcitycenter"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamCityCenterQuery is the builder for querying ExamCityCenter entities.
type ExamCityCenterQuery struct {
	config
	ctx                          *QueryContext
	order                        []examcitycenter.OrderOption
	inters                       []Interceptor
	predicates                   []predicate.ExamCityCenter
	withExamCityCenterRef        *ExamApplicationsIPQuery
	withExamCityCenterMTSPMMGRef *ExamApplicationMTSPMMGQuery
	withExamCityCenterGDSPARef   *ExamApplicationsGDSPAQuery
	withExamCityCenterGDSPMRef   *ExamApplicationsGDSPMQuery
	withExamCityCenterPMPARef    *ExamApplicationsPMPAQuery
	withExamCityCenterPSRef      *ExamApplicationsPSQuery
	withFKs                      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExamCityCenterQuery builder.
func (eccq *ExamCityCenterQuery) Where(ps ...predicate.ExamCityCenter) *ExamCityCenterQuery {
	eccq.predicates = append(eccq.predicates, ps...)
	return eccq
}

// Limit the number of records to be returned by this query.
func (eccq *ExamCityCenterQuery) Limit(limit int) *ExamCityCenterQuery {
	eccq.ctx.Limit = &limit
	return eccq
}

// Offset to start from.
func (eccq *ExamCityCenterQuery) Offset(offset int) *ExamCityCenterQuery {
	eccq.ctx.Offset = &offset
	return eccq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eccq *ExamCityCenterQuery) Unique(unique bool) *ExamCityCenterQuery {
	eccq.ctx.Unique = &unique
	return eccq
}

// Order specifies how the records should be ordered.
func (eccq *ExamCityCenterQuery) Order(o ...examcitycenter.OrderOption) *ExamCityCenterQuery {
	eccq.order = append(eccq.order, o...)
	return eccq
}

// QueryExamCityCenterRef chains the current query on the "ExamCityCenterRef" edge.
func (eccq *ExamCityCenterQuery) QueryExamCityCenterRef() *ExamApplicationsIPQuery {
	query := (&ExamApplicationsIPClient{config: eccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcitycenter.Table, examcitycenter.FieldID, selector),
			sqlgraph.To(exam_applications_ip.Table, exam_applications_ip.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, examcitycenter.ExamCityCenterRefTable, examcitycenter.ExamCityCenterRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExamCityCenterMTSPMMGRef chains the current query on the "ExamCityCenterMTSPMMGRef" edge.
func (eccq *ExamCityCenterQuery) QueryExamCityCenterMTSPMMGRef() *ExamApplicationMTSPMMGQuery {
	query := (&ExamApplicationMTSPMMGClient{config: eccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcitycenter.Table, examcitycenter.FieldID, selector),
			sqlgraph.To(exam_application_mtspmmg.Table, exam_application_mtspmmg.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, examcitycenter.ExamCityCenterMTSPMMGRefTable, examcitycenter.ExamCityCenterMTSPMMGRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExamCityCenterGDSPARef chains the current query on the "ExamCityCenterGDSPARef" edge.
func (eccq *ExamCityCenterQuery) QueryExamCityCenterGDSPARef() *ExamApplicationsGDSPAQuery {
	query := (&ExamApplicationsGDSPAClient{config: eccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcitycenter.Table, examcitycenter.FieldID, selector),
			sqlgraph.To(exam_applications_gdspa.Table, exam_applications_gdspa.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, examcitycenter.ExamCityCenterGDSPARefTable, examcitycenter.ExamCityCenterGDSPARefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExamCityCenterGDSPMRef chains the current query on the "ExamCityCenterGDSPMRef" edge.
func (eccq *ExamCityCenterQuery) QueryExamCityCenterGDSPMRef() *ExamApplicationsGDSPMQuery {
	query := (&ExamApplicationsGDSPMClient{config: eccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcitycenter.Table, examcitycenter.FieldID, selector),
			sqlgraph.To(exam_applications_gdspm.Table, exam_applications_gdspm.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, examcitycenter.ExamCityCenterGDSPMRefTable, examcitycenter.ExamCityCenterGDSPMRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExamCityCenterPMPARef chains the current query on the "ExamCityCenterPMPARef" edge.
func (eccq *ExamCityCenterQuery) QueryExamCityCenterPMPARef() *ExamApplicationsPMPAQuery {
	query := (&ExamApplicationsPMPAClient{config: eccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcitycenter.Table, examcitycenter.FieldID, selector),
			sqlgraph.To(exam_applications_pmpa.Table, exam_applications_pmpa.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, examcitycenter.ExamCityCenterPMPARefTable, examcitycenter.ExamCityCenterPMPARefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExamCityCenterPSRef chains the current query on the "ExamCityCenterPSRef" edge.
func (eccq *ExamCityCenterQuery) QueryExamCityCenterPSRef() *ExamApplicationsPSQuery {
	query := (&ExamApplicationsPSClient{config: eccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcitycenter.Table, examcitycenter.FieldID, selector),
			sqlgraph.To(exam_applications_ps.Table, exam_applications_ps.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, examcitycenter.ExamCityCenterPSRefTable, examcitycenter.ExamCityCenterPSRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ExamCityCenter entity from the query.
// Returns a *NotFoundError when no ExamCityCenter was found.
func (eccq *ExamCityCenterQuery) First(ctx context.Context) (*ExamCityCenter, error) {
	nodes, err := eccq.Limit(1).All(setContextOp(ctx, eccq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{examcitycenter.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eccq *ExamCityCenterQuery) FirstX(ctx context.Context) *ExamCityCenter {
	node, err := eccq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExamCityCenter ID from the query.
// Returns a *NotFoundError when no ExamCityCenter ID was found.
func (eccq *ExamCityCenterQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = eccq.Limit(1).IDs(setContextOp(ctx, eccq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{examcitycenter.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eccq *ExamCityCenterQuery) FirstIDX(ctx context.Context) int32 {
	id, err := eccq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExamCityCenter entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExamCityCenter entity is found.
// Returns a *NotFoundError when no ExamCityCenter entities are found.
func (eccq *ExamCityCenterQuery) Only(ctx context.Context) (*ExamCityCenter, error) {
	nodes, err := eccq.Limit(2).All(setContextOp(ctx, eccq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{examcitycenter.Label}
	default:
		return nil, &NotSingularError{examcitycenter.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eccq *ExamCityCenterQuery) OnlyX(ctx context.Context) *ExamCityCenter {
	node, err := eccq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExamCityCenter ID in the query.
// Returns a *NotSingularError when more than one ExamCityCenter ID is found.
// Returns a *NotFoundError when no entities are found.
func (eccq *ExamCityCenterQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = eccq.Limit(2).IDs(setContextOp(ctx, eccq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{examcitycenter.Label}
	default:
		err = &NotSingularError{examcitycenter.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eccq *ExamCityCenterQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := eccq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExamCityCenters.
func (eccq *ExamCityCenterQuery) All(ctx context.Context) ([]*ExamCityCenter, error) {
	ctx = setContextOp(ctx, eccq.ctx, "All")
	if err := eccq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExamCityCenter, *ExamCityCenterQuery]()
	return withInterceptors[[]*ExamCityCenter](ctx, eccq, qr, eccq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eccq *ExamCityCenterQuery) AllX(ctx context.Context) []*ExamCityCenter {
	nodes, err := eccq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExamCityCenter IDs.
func (eccq *ExamCityCenterQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if eccq.ctx.Unique == nil && eccq.path != nil {
		eccq.Unique(true)
	}
	ctx = setContextOp(ctx, eccq.ctx, "IDs")
	if err = eccq.Select(examcitycenter.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eccq *ExamCityCenterQuery) IDsX(ctx context.Context) []int32 {
	ids, err := eccq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eccq *ExamCityCenterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eccq.ctx, "Count")
	if err := eccq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eccq, querierCount[*ExamCityCenterQuery](), eccq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eccq *ExamCityCenterQuery) CountX(ctx context.Context) int {
	count, err := eccq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eccq *ExamCityCenterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eccq.ctx, "Exist")
	switch _, err := eccq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eccq *ExamCityCenterQuery) ExistX(ctx context.Context) bool {
	exist, err := eccq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExamCityCenterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eccq *ExamCityCenterQuery) Clone() *ExamCityCenterQuery {
	if eccq == nil {
		return nil
	}
	return &ExamCityCenterQuery{
		config:                       eccq.config,
		ctx:                          eccq.ctx.Clone(),
		order:                        append([]examcitycenter.OrderOption{}, eccq.order...),
		inters:                       append([]Interceptor{}, eccq.inters...),
		predicates:                   append([]predicate.ExamCityCenter{}, eccq.predicates...),
		withExamCityCenterRef:        eccq.withExamCityCenterRef.Clone(),
		withExamCityCenterMTSPMMGRef: eccq.withExamCityCenterMTSPMMGRef.Clone(),
		withExamCityCenterGDSPARef:   eccq.withExamCityCenterGDSPARef.Clone(),
		withExamCityCenterGDSPMRef:   eccq.withExamCityCenterGDSPMRef.Clone(),
		withExamCityCenterPMPARef:    eccq.withExamCityCenterPMPARef.Clone(),
		withExamCityCenterPSRef:      eccq.withExamCityCenterPSRef.Clone(),
		// clone intermediate query.
		sql:  eccq.sql.Clone(),
		path: eccq.path,
	}
}

// WithExamCityCenterRef tells the query-builder to eager-load the nodes that are connected to
// the "ExamCityCenterRef" edge. The optional arguments are used to configure the query builder of the edge.
func (eccq *ExamCityCenterQuery) WithExamCityCenterRef(opts ...func(*ExamApplicationsIPQuery)) *ExamCityCenterQuery {
	query := (&ExamApplicationsIPClient{config: eccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eccq.withExamCityCenterRef = query
	return eccq
}

// WithExamCityCenterMTSPMMGRef tells the query-builder to eager-load the nodes that are connected to
// the "ExamCityCenterMTSPMMGRef" edge. The optional arguments are used to configure the query builder of the edge.
func (eccq *ExamCityCenterQuery) WithExamCityCenterMTSPMMGRef(opts ...func(*ExamApplicationMTSPMMGQuery)) *ExamCityCenterQuery {
	query := (&ExamApplicationMTSPMMGClient{config: eccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eccq.withExamCityCenterMTSPMMGRef = query
	return eccq
}

// WithExamCityCenterGDSPARef tells the query-builder to eager-load the nodes that are connected to
// the "ExamCityCenterGDSPARef" edge. The optional arguments are used to configure the query builder of the edge.
func (eccq *ExamCityCenterQuery) WithExamCityCenterGDSPARef(opts ...func(*ExamApplicationsGDSPAQuery)) *ExamCityCenterQuery {
	query := (&ExamApplicationsGDSPAClient{config: eccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eccq.withExamCityCenterGDSPARef = query
	return eccq
}

// WithExamCityCenterGDSPMRef tells the query-builder to eager-load the nodes that are connected to
// the "ExamCityCenterGDSPMRef" edge. The optional arguments are used to configure the query builder of the edge.
func (eccq *ExamCityCenterQuery) WithExamCityCenterGDSPMRef(opts ...func(*ExamApplicationsGDSPMQuery)) *ExamCityCenterQuery {
	query := (&ExamApplicationsGDSPMClient{config: eccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eccq.withExamCityCenterGDSPMRef = query
	return eccq
}

// WithExamCityCenterPMPARef tells the query-builder to eager-load the nodes that are connected to
// the "ExamCityCenterPMPARef" edge. The optional arguments are used to configure the query builder of the edge.
func (eccq *ExamCityCenterQuery) WithExamCityCenterPMPARef(opts ...func(*ExamApplicationsPMPAQuery)) *ExamCityCenterQuery {
	query := (&ExamApplicationsPMPAClient{config: eccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eccq.withExamCityCenterPMPARef = query
	return eccq
}

// WithExamCityCenterPSRef tells the query-builder to eager-load the nodes that are connected to
// the "ExamCityCenterPSRef" edge. The optional arguments are used to configure the query builder of the edge.
func (eccq *ExamCityCenterQuery) WithExamCityCenterPSRef(opts ...func(*ExamApplicationsPSQuery)) *ExamCityCenterQuery {
	query := (&ExamApplicationsPSClient{config: eccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eccq.withExamCityCenterPSRef = query
	return eccq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ExamCode int32 `json:"ExamCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExamCityCenter.Query().
//		GroupBy(examcitycenter.FieldExamCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eccq *ExamCityCenterQuery) GroupBy(field string, fields ...string) *ExamCityCenterGroupBy {
	eccq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExamCityCenterGroupBy{build: eccq}
	grbuild.flds = &eccq.ctx.Fields
	grbuild.label = examcitycenter.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ExamCode int32 `json:"ExamCode,omitempty"`
//	}
//
//	client.ExamCityCenter.Query().
//		Select(examcitycenter.FieldExamCode).
//		Scan(ctx, &v)
func (eccq *ExamCityCenterQuery) Select(fields ...string) *ExamCityCenterSelect {
	eccq.ctx.Fields = append(eccq.ctx.Fields, fields...)
	sbuild := &ExamCityCenterSelect{ExamCityCenterQuery: eccq}
	sbuild.label = examcitycenter.Label
	sbuild.flds, sbuild.scan = &eccq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExamCityCenterSelect configured with the given aggregations.
func (eccq *ExamCityCenterQuery) Aggregate(fns ...AggregateFunc) *ExamCityCenterSelect {
	return eccq.Select().Aggregate(fns...)
}

func (eccq *ExamCityCenterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eccq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eccq); err != nil {
				return err
			}
		}
	}
	for _, f := range eccq.ctx.Fields {
		if !examcitycenter.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eccq.path != nil {
		prev, err := eccq.path(ctx)
		if err != nil {
			return err
		}
		eccq.sql = prev
	}
	return nil
}

func (eccq *ExamCityCenterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExamCityCenter, error) {
	var (
		nodes       = []*ExamCityCenter{}
		withFKs     = eccq.withFKs
		_spec       = eccq.querySpec()
		loadedTypes = [6]bool{
			eccq.withExamCityCenterRef != nil,
			eccq.withExamCityCenterMTSPMMGRef != nil,
			eccq.withExamCityCenterGDSPARef != nil,
			eccq.withExamCityCenterGDSPMRef != nil,
			eccq.withExamCityCenterPMPARef != nil,
			eccq.withExamCityCenterPSRef != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, examcitycenter.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExamCityCenter).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExamCityCenter{config: eccq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eccq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eccq.withExamCityCenterRef; query != nil {
		if err := eccq.loadExamCityCenterRef(ctx, query, nodes,
			func(n *ExamCityCenter) { n.Edges.ExamCityCenterRef = []*Exam_Applications_IP{} },
			func(n *ExamCityCenter, e *Exam_Applications_IP) {
				n.Edges.ExamCityCenterRef = append(n.Edges.ExamCityCenterRef, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := eccq.withExamCityCenterMTSPMMGRef; query != nil {
		if err := eccq.loadExamCityCenterMTSPMMGRef(ctx, query, nodes,
			func(n *ExamCityCenter) { n.Edges.ExamCityCenterMTSPMMGRef = []*Exam_Application_MTSPMMG{} },
			func(n *ExamCityCenter, e *Exam_Application_MTSPMMG) {
				n.Edges.ExamCityCenterMTSPMMGRef = append(n.Edges.ExamCityCenterMTSPMMGRef, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := eccq.withExamCityCenterGDSPARef; query != nil {
		if err := eccq.loadExamCityCenterGDSPARef(ctx, query, nodes,
			func(n *ExamCityCenter) { n.Edges.ExamCityCenterGDSPARef = []*Exam_Applications_GDSPA{} },
			func(n *ExamCityCenter, e *Exam_Applications_GDSPA) {
				n.Edges.ExamCityCenterGDSPARef = append(n.Edges.ExamCityCenterGDSPARef, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := eccq.withExamCityCenterGDSPMRef; query != nil {
		if err := eccq.loadExamCityCenterGDSPMRef(ctx, query, nodes,
			func(n *ExamCityCenter) { n.Edges.ExamCityCenterGDSPMRef = []*Exam_Applications_GDSPM{} },
			func(n *ExamCityCenter, e *Exam_Applications_GDSPM) {
				n.Edges.ExamCityCenterGDSPMRef = append(n.Edges.ExamCityCenterGDSPMRef, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := eccq.withExamCityCenterPMPARef; query != nil {
		if err := eccq.loadExamCityCenterPMPARef(ctx, query, nodes,
			func(n *ExamCityCenter) { n.Edges.ExamCityCenterPMPARef = []*Exam_Applications_PMPA{} },
			func(n *ExamCityCenter, e *Exam_Applications_PMPA) {
				n.Edges.ExamCityCenterPMPARef = append(n.Edges.ExamCityCenterPMPARef, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := eccq.withExamCityCenterPSRef; query != nil {
		if err := eccq.loadExamCityCenterPSRef(ctx, query, nodes,
			func(n *ExamCityCenter) { n.Edges.ExamCityCenterPSRef = []*Exam_Applications_PS{} },
			func(n *ExamCityCenter, e *Exam_Applications_PS) {
				n.Edges.ExamCityCenterPSRef = append(n.Edges.ExamCityCenterPSRef, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eccq *ExamCityCenterQuery) loadExamCityCenterRef(ctx context.Context, query *ExamApplicationsIPQuery, nodes []*ExamCityCenter, init func(*ExamCityCenter), assign func(*ExamCityCenter, *Exam_Applications_IP)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*ExamCityCenter)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(exam_applications_ip.FieldExamCityCenterCode)
	}
	query.Where(predicate.Exam_Applications_IP(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(examcitycenter.ExamCityCenterRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ExamCityCenterCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "ExamCityCenterCode" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eccq *ExamCityCenterQuery) loadExamCityCenterMTSPMMGRef(ctx context.Context, query *ExamApplicationMTSPMMGQuery, nodes []*ExamCityCenter, init func(*ExamCityCenter), assign func(*ExamCityCenter, *Exam_Application_MTSPMMG)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*ExamCityCenter)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(exam_application_mtspmmg.FieldExamCityCenterCode)
	}
	query.Where(predicate.Exam_Application_MTSPMMG(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(examcitycenter.ExamCityCenterMTSPMMGRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ExamCityCenterCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "ExamCityCenterCode" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eccq *ExamCityCenterQuery) loadExamCityCenterGDSPARef(ctx context.Context, query *ExamApplicationsGDSPAQuery, nodes []*ExamCityCenter, init func(*ExamCityCenter), assign func(*ExamCityCenter, *Exam_Applications_GDSPA)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*ExamCityCenter)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(exam_applications_gdspa.FieldExamCityCenterCode)
	}
	query.Where(predicate.Exam_Applications_GDSPA(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(examcitycenter.ExamCityCenterGDSPARefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ExamCityCenterCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "ExamCityCenterCode" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eccq *ExamCityCenterQuery) loadExamCityCenterGDSPMRef(ctx context.Context, query *ExamApplicationsGDSPMQuery, nodes []*ExamCityCenter, init func(*ExamCityCenter), assign func(*ExamCityCenter, *Exam_Applications_GDSPM)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*ExamCityCenter)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(exam_applications_gdspm.FieldExamCityCenterCode)
	}
	query.Where(predicate.Exam_Applications_GDSPM(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(examcitycenter.ExamCityCenterGDSPMRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ExamCityCenterCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "ExamCityCenterCode" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eccq *ExamCityCenterQuery) loadExamCityCenterPMPARef(ctx context.Context, query *ExamApplicationsPMPAQuery, nodes []*ExamCityCenter, init func(*ExamCityCenter), assign func(*ExamCityCenter, *Exam_Applications_PMPA)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*ExamCityCenter)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(exam_applications_pmpa.FieldExamCityCenterCode)
	}
	query.Where(predicate.Exam_Applications_PMPA(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(examcitycenter.ExamCityCenterPMPARefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ExamCityCenterCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "ExamCityCenterCode" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eccq *ExamCityCenterQuery) loadExamCityCenterPSRef(ctx context.Context, query *ExamApplicationsPSQuery, nodes []*ExamCityCenter, init func(*ExamCityCenter), assign func(*ExamCityCenter, *Exam_Applications_PS)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*ExamCityCenter)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(exam_applications_ps.FieldExamCityCenterCode)
	}
	query.Where(predicate.Exam_Applications_PS(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(examcitycenter.ExamCityCenterPSRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ExamCityCenterCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "ExamCityCenterCode" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eccq *ExamCityCenterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eccq.querySpec()
	_spec.Node.Columns = eccq.ctx.Fields
	if len(eccq.ctx.Fields) > 0 {
		_spec.Unique = eccq.ctx.Unique != nil && *eccq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eccq.driver, _spec)
}

func (eccq *ExamCityCenterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(examcitycenter.Table, examcitycenter.Columns, sqlgraph.NewFieldSpec(examcitycenter.FieldID, field.TypeInt32))
	_spec.From = eccq.sql
	if unique := eccq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eccq.path != nil {
		_spec.Unique = true
	}
	if fields := eccq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examcitycenter.FieldID)
		for i := range fields {
			if fields[i] != examcitycenter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eccq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eccq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eccq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eccq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eccq *ExamCityCenterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eccq.driver.Dialect())
	t1 := builder.Table(examcitycenter.Table)
	columns := eccq.ctx.Fields
	if len(columns) == 0 {
		columns = examcitycenter.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eccq.sql != nil {
		selector = eccq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eccq.ctx.Unique != nil && *eccq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eccq.predicates {
		p(selector)
	}
	for _, p := range eccq.order {
		p(selector)
	}
	if offset := eccq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eccq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExamCityCenterGroupBy is the group-by builder for ExamCityCenter entities.
type ExamCityCenterGroupBy struct {
	selector
	build *ExamCityCenterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eccgb *ExamCityCenterGroupBy) Aggregate(fns ...AggregateFunc) *ExamCityCenterGroupBy {
	eccgb.fns = append(eccgb.fns, fns...)
	return eccgb
}

// Scan applies the selector query and scans the result into the given value.
func (eccgb *ExamCityCenterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eccgb.build.ctx, "GroupBy")
	if err := eccgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamCityCenterQuery, *ExamCityCenterGroupBy](ctx, eccgb.build, eccgb, eccgb.build.inters, v)
}

func (eccgb *ExamCityCenterGroupBy) sqlScan(ctx context.Context, root *ExamCityCenterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eccgb.fns))
	for _, fn := range eccgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eccgb.flds)+len(eccgb.fns))
		for _, f := range *eccgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eccgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eccgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExamCityCenterSelect is the builder for selecting fields of ExamCityCenter entities.
type ExamCityCenterSelect struct {
	*ExamCityCenterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eccs *ExamCityCenterSelect) Aggregate(fns ...AggregateFunc) *ExamCityCenterSelect {
	eccs.fns = append(eccs.fns, fns...)
	return eccs
}

// Scan applies the selector query and scans the result into the given value.
func (eccs *ExamCityCenterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eccs.ctx, "Select")
	if err := eccs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamCityCenterQuery, *ExamCityCenterSelect](ctx, eccs.ExamCityCenterQuery, eccs, eccs.inters, v)
}

func (eccs *ExamCityCenterSelect) sqlScan(ctx context.Context, root *ExamCityCenterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eccs.fns))
	for _, fn := range eccs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eccs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eccs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}