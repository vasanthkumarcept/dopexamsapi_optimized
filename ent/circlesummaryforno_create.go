// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/circlesummaryforno"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/exam_applications_gdspa"
	"recruit/ent/exam_applications_gdspm"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_pmpa"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/usermaster"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CircleSummaryForNOCreate is the builder for creating a CircleSummaryForNO entity.
type CircleSummaryForNOCreate struct {
	config
	mutation *CircleSummaryForNOMutation
	hooks    []Hook
}

// SetCircleOfficeId sets the "CircleOfficeId" field.
func (csfnc *CircleSummaryForNOCreate) SetCircleOfficeId(s string) *CircleSummaryForNOCreate {
	csfnc.mutation.SetCircleOfficeId(s)
	return csfnc
}

// SetCircleOfficeName sets the "CircleOfficeName" field.
func (csfnc *CircleSummaryForNOCreate) SetCircleOfficeName(s string) *CircleSummaryForNOCreate {
	csfnc.mutation.SetCircleOfficeName(s)
	return csfnc
}

// SetApproveHallTicketGenrationIP sets the "ApproveHallTicketGenrationIP" field.
func (csfnc *CircleSummaryForNOCreate) SetApproveHallTicketGenrationIP(b bool) *CircleSummaryForNOCreate {
	csfnc.mutation.SetApproveHallTicketGenrationIP(b)
	return csfnc
}

// SetNillableApproveHallTicketGenrationIP sets the "ApproveHallTicketGenrationIP" field if the given value is not nil.
func (csfnc *CircleSummaryForNOCreate) SetNillableApproveHallTicketGenrationIP(b *bool) *CircleSummaryForNOCreate {
	if b != nil {
		csfnc.SetApproveHallTicketGenrationIP(*b)
	}
	return csfnc
}

// SetApproveHallTicketGenrationPS sets the "ApproveHallTicketGenrationPS" field.
func (csfnc *CircleSummaryForNOCreate) SetApproveHallTicketGenrationPS(b bool) *CircleSummaryForNOCreate {
	csfnc.mutation.SetApproveHallTicketGenrationPS(b)
	return csfnc
}

// SetNillableApproveHallTicketGenrationPS sets the "ApproveHallTicketGenrationPS" field if the given value is not nil.
func (csfnc *CircleSummaryForNOCreate) SetNillableApproveHallTicketGenrationPS(b *bool) *CircleSummaryForNOCreate {
	if b != nil {
		csfnc.SetApproveHallTicketGenrationPS(*b)
	}
	return csfnc
}

// SetApproveHallTicketGenrationPM sets the "ApproveHallTicketGenrationPM" field.
func (csfnc *CircleSummaryForNOCreate) SetApproveHallTicketGenrationPM(b bool) *CircleSummaryForNOCreate {
	csfnc.mutation.SetApproveHallTicketGenrationPM(b)
	return csfnc
}

// SetNillableApproveHallTicketGenrationPM sets the "ApproveHallTicketGenrationPM" field if the given value is not nil.
func (csfnc *CircleSummaryForNOCreate) SetNillableApproveHallTicketGenrationPM(b *bool) *CircleSummaryForNOCreate {
	if b != nil {
		csfnc.SetApproveHallTicketGenrationPM(*b)
	}
	return csfnc
}

// SetApproveHallTicketGenrationPA sets the "ApproveHallTicketGenrationPA" field.
func (csfnc *CircleSummaryForNOCreate) SetApproveHallTicketGenrationPA(b bool) *CircleSummaryForNOCreate {
	csfnc.mutation.SetApproveHallTicketGenrationPA(b)
	return csfnc
}

// SetNillableApproveHallTicketGenrationPA sets the "ApproveHallTicketGenrationPA" field if the given value is not nil.
func (csfnc *CircleSummaryForNOCreate) SetNillableApproveHallTicketGenrationPA(b *bool) *CircleSummaryForNOCreate {
	if b != nil {
		csfnc.SetApproveHallTicketGenrationPA(*b)
	}
	return csfnc
}

// SetID sets the "id" field.
func (csfnc *CircleSummaryForNOCreate) SetID(i int32) *CircleSummaryForNOCreate {
	csfnc.mutation.SetID(i)
	return csfnc
}

// AddCircleuserIDs adds the "circleusers" edge to the UserMaster entity by IDs.
func (csfnc *CircleSummaryForNOCreate) AddCircleuserIDs(ids ...int64) *CircleSummaryForNOCreate {
	csfnc.mutation.AddCircleuserIDs(ids...)
	return csfnc
}

// AddCircleusers adds the "circleusers" edges to the UserMaster entity.
func (csfnc *CircleSummaryForNOCreate) AddCircleusers(u ...*UserMaster) *CircleSummaryForNOCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return csfnc.AddCircleuserIDs(ids...)
}

// AddCircleRefsForHallTicketIPIDs adds the "CircleRefsForHallTicketIP" edge to the Exam_Applications_IP entity by IDs.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketIPIDs(ids ...int64) *CircleSummaryForNOCreate {
	csfnc.mutation.AddCircleRefsForHallTicketIPIDs(ids...)
	return csfnc
}

// AddCircleRefsForHallTicketIP adds the "CircleRefsForHallTicketIP" edges to the Exam_Applications_IP entity.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketIP(e ...*Exam_Applications_IP) *CircleSummaryForNOCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csfnc.AddCircleRefsForHallTicketIPIDs(ids...)
}

// AddCircleRefsForHallTicketPSIDs adds the "CircleRefsForHallTicketPS" edge to the Exam_Applications_PS entity by IDs.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketPSIDs(ids ...int64) *CircleSummaryForNOCreate {
	csfnc.mutation.AddCircleRefsForHallTicketPSIDs(ids...)
	return csfnc
}

// AddCircleRefsForHallTicketPS adds the "CircleRefsForHallTicketPS" edges to the Exam_Applications_PS entity.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketPS(e ...*Exam_Applications_PS) *CircleSummaryForNOCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csfnc.AddCircleRefsForHallTicketPSIDs(ids...)
}

// AddCircleRefsForHallTicketGDSPAIDs adds the "CircleRefsForHallTicketGDSPA" edge to the Exam_Applications_GDSPA entity by IDs.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketGDSPAIDs(ids ...int64) *CircleSummaryForNOCreate {
	csfnc.mutation.AddCircleRefsForHallTicketGDSPAIDs(ids...)
	return csfnc
}

// AddCircleRefsForHallTicketGDSPA adds the "CircleRefsForHallTicketGDSPA" edges to the Exam_Applications_GDSPA entity.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketGDSPA(e ...*Exam_Applications_GDSPA) *CircleSummaryForNOCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csfnc.AddCircleRefsForHallTicketGDSPAIDs(ids...)
}

// AddCircleRefsForHallTicketGDSPMIDs adds the "CircleRefsForHallTicketGDSPM" edge to the Exam_Applications_GDSPM entity by IDs.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketGDSPMIDs(ids ...int64) *CircleSummaryForNOCreate {
	csfnc.mutation.AddCircleRefsForHallTicketGDSPMIDs(ids...)
	return csfnc
}

// AddCircleRefsForHallTicketGDSPM adds the "CircleRefsForHallTicketGDSPM" edges to the Exam_Applications_GDSPM entity.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketGDSPM(e ...*Exam_Applications_GDSPM) *CircleSummaryForNOCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csfnc.AddCircleRefsForHallTicketGDSPMIDs(ids...)
}

// AddCircleRefsForHallTicketPMPAIDs adds the "CircleRefsForHallTicketPMPA" edge to the Exam_Applications_PMPA entity by IDs.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketPMPAIDs(ids ...int64) *CircleSummaryForNOCreate {
	csfnc.mutation.AddCircleRefsForHallTicketPMPAIDs(ids...)
	return csfnc
}

// AddCircleRefsForHallTicketPMPA adds the "CircleRefsForHallTicketPMPA" edges to the Exam_Applications_PMPA entity.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketPMPA(e ...*Exam_Applications_PMPA) *CircleSummaryForNOCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csfnc.AddCircleRefsForHallTicketPMPAIDs(ids...)
}

// AddCircleRefsForHallTicketMTSPMMGIDs adds the "CircleRefsForHallTicketMTSPMMG" edge to the Exam_Application_MTSPMMG entity by IDs.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketMTSPMMGIDs(ids ...int64) *CircleSummaryForNOCreate {
	csfnc.mutation.AddCircleRefsForHallTicketMTSPMMGIDs(ids...)
	return csfnc
}

// AddCircleRefsForHallTicketMTSPMMG adds the "CircleRefsForHallTicketMTSPMMG" edges to the Exam_Application_MTSPMMG entity.
func (csfnc *CircleSummaryForNOCreate) AddCircleRefsForHallTicketMTSPMMG(e ...*Exam_Application_MTSPMMG) *CircleSummaryForNOCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return csfnc.AddCircleRefsForHallTicketMTSPMMGIDs(ids...)
}

// Mutation returns the CircleSummaryForNOMutation object of the builder.
func (csfnc *CircleSummaryForNOCreate) Mutation() *CircleSummaryForNOMutation {
	return csfnc.mutation
}

// Save creates the CircleSummaryForNO in the database.
func (csfnc *CircleSummaryForNOCreate) Save(ctx context.Context) (*CircleSummaryForNO, error) {
	return withHooks(ctx, csfnc.sqlSave, csfnc.mutation, csfnc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (csfnc *CircleSummaryForNOCreate) SaveX(ctx context.Context) *CircleSummaryForNO {
	v, err := csfnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (csfnc *CircleSummaryForNOCreate) Exec(ctx context.Context) error {
	_, err := csfnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csfnc *CircleSummaryForNOCreate) ExecX(ctx context.Context) {
	if err := csfnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csfnc *CircleSummaryForNOCreate) check() error {
	if _, ok := csfnc.mutation.CircleOfficeId(); !ok {
		return &ValidationError{Name: "CircleOfficeId", err: errors.New(`ent: missing required field "CircleSummaryForNO.CircleOfficeId"`)}
	}
	if _, ok := csfnc.mutation.CircleOfficeName(); !ok {
		return &ValidationError{Name: "CircleOfficeName", err: errors.New(`ent: missing required field "CircleSummaryForNO.CircleOfficeName"`)}
	}
	return nil
}

func (csfnc *CircleSummaryForNOCreate) sqlSave(ctx context.Context) (*CircleSummaryForNO, error) {
	if err := csfnc.check(); err != nil {
		return nil, err
	}
	_node, _spec := csfnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, csfnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	csfnc.mutation.id = &_node.ID
	csfnc.mutation.done = true
	return _node, nil
}

func (csfnc *CircleSummaryForNOCreate) createSpec() (*CircleSummaryForNO, *sqlgraph.CreateSpec) {
	var (
		_node = &CircleSummaryForNO{config: csfnc.config}
		_spec = sqlgraph.NewCreateSpec(circlesummaryforno.Table, sqlgraph.NewFieldSpec(circlesummaryforno.FieldID, field.TypeInt32))
	)
	if id, ok := csfnc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := csfnc.mutation.CircleOfficeId(); ok {
		_spec.SetField(circlesummaryforno.FieldCircleOfficeId, field.TypeString, value)
		_node.CircleOfficeId = value
	}
	if value, ok := csfnc.mutation.CircleOfficeName(); ok {
		_spec.SetField(circlesummaryforno.FieldCircleOfficeName, field.TypeString, value)
		_node.CircleOfficeName = value
	}
	if value, ok := csfnc.mutation.ApproveHallTicketGenrationIP(); ok {
		_spec.SetField(circlesummaryforno.FieldApproveHallTicketGenrationIP, field.TypeBool, value)
		_node.ApproveHallTicketGenrationIP = value
	}
	if value, ok := csfnc.mutation.ApproveHallTicketGenrationPS(); ok {
		_spec.SetField(circlesummaryforno.FieldApproveHallTicketGenrationPS, field.TypeBool, value)
		_node.ApproveHallTicketGenrationPS = value
	}
	if value, ok := csfnc.mutation.ApproveHallTicketGenrationPM(); ok {
		_spec.SetField(circlesummaryforno.FieldApproveHallTicketGenrationPM, field.TypeBool, value)
		_node.ApproveHallTicketGenrationPM = value
	}
	if value, ok := csfnc.mutation.ApproveHallTicketGenrationPA(); ok {
		_spec.SetField(circlesummaryforno.FieldApproveHallTicketGenrationPA, field.TypeBool, value)
		_node.ApproveHallTicketGenrationPA = value
	}
	if nodes := csfnc.mutation.CircleusersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   circlesummaryforno.CircleusersTable,
			Columns: []string{circlesummaryforno.CircleusersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csfnc.mutation.CircleRefsForHallTicketIPIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   circlesummaryforno.CircleRefsForHallTicketIPTable,
			Columns: []string{circlesummaryforno.CircleRefsForHallTicketIPColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csfnc.mutation.CircleRefsForHallTicketPSIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   circlesummaryforno.CircleRefsForHallTicketPSTable,
			Columns: []string{circlesummaryforno.CircleRefsForHallTicketPSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csfnc.mutation.CircleRefsForHallTicketGDSPAIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   circlesummaryforno.CircleRefsForHallTicketGDSPATable,
			Columns: []string{circlesummaryforno.CircleRefsForHallTicketGDSPAColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_gdspa.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csfnc.mutation.CircleRefsForHallTicketGDSPMIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   circlesummaryforno.CircleRefsForHallTicketGDSPMTable,
			Columns: []string{circlesummaryforno.CircleRefsForHallTicketGDSPMColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_gdspm.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csfnc.mutation.CircleRefsForHallTicketPMPAIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   circlesummaryforno.CircleRefsForHallTicketPMPATable,
			Columns: []string{circlesummaryforno.CircleRefsForHallTicketPMPAColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_pmpa.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csfnc.mutation.CircleRefsForHallTicketMTSPMMGIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   circlesummaryforno.CircleRefsForHallTicketMTSPMMGTable,
			Columns: []string{circlesummaryforno.CircleRefsForHallTicketMTSPMMGColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_application_mtspmmg.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CircleSummaryForNOCreateBulk is the builder for creating many CircleSummaryForNO entities in bulk.
type CircleSummaryForNOCreateBulk struct {
	config
	builders []*CircleSummaryForNOCreate
}

// Save creates the CircleSummaryForNO entities in the database.
func (csfncb *CircleSummaryForNOCreateBulk) Save(ctx context.Context) ([]*CircleSummaryForNO, error) {
	specs := make([]*sqlgraph.CreateSpec, len(csfncb.builders))
	nodes := make([]*CircleSummaryForNO, len(csfncb.builders))
	mutators := make([]Mutator, len(csfncb.builders))
	for i := range csfncb.builders {
		func(i int, root context.Context) {
			builder := csfncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CircleSummaryForNOMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, csfncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, csfncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, csfncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (csfncb *CircleSummaryForNOCreateBulk) SaveX(ctx context.Context) []*CircleSummaryForNO {
	v, err := csfncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (csfncb *CircleSummaryForNOCreateBulk) Exec(ctx context.Context) error {
	_, err := csfncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csfncb *CircleSummaryForNOCreateBulk) ExecX(ctx context.Context) {
	if err := csfncb.Exec(ctx); err != nil {
		panic(err)
	}
}