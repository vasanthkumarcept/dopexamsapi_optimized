// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/exam"
	"recruit/ent/examcalendar"
	"recruit/ent/vacancyyear"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VacancyYearCreate is the builder for creating a VacancyYear entity.
type VacancyYearCreate struct {
	config
	mutation *VacancyYearMutation
	hooks    []Hook
}

// SetFromDate sets the "FromDate" field.
func (vyc *VacancyYearCreate) SetFromDate(t time.Time) *VacancyYearCreate {
	vyc.mutation.SetFromDate(t)
	return vyc
}

// SetToDate sets the "ToDate" field.
func (vyc *VacancyYearCreate) SetToDate(t time.Time) *VacancyYearCreate {
	vyc.mutation.SetToDate(t)
	return vyc
}

// SetNotifyCode sets the "NotifyCode" field.
func (vyc *VacancyYearCreate) SetNotifyCode(i int32) *VacancyYearCreate {
	vyc.mutation.SetNotifyCode(i)
	return vyc
}

// SetNillableNotifyCode sets the "NotifyCode" field if the given value is not nil.
func (vyc *VacancyYearCreate) SetNillableNotifyCode(i *int32) *VacancyYearCreate {
	if i != nil {
		vyc.SetNotifyCode(*i)
	}
	return vyc
}

// SetVacancyYear sets the "VacancyYear" field.
func (vyc *VacancyYearCreate) SetVacancyYear(s string) *VacancyYearCreate {
	vyc.mutation.SetVacancyYear(s)
	return vyc
}

// SetCalendarCode sets the "CalendarCode" field.
func (vyc *VacancyYearCreate) SetCalendarCode(i int32) *VacancyYearCreate {
	vyc.mutation.SetCalendarCode(i)
	return vyc
}

// SetNillableCalendarCode sets the "CalendarCode" field if the given value is not nil.
func (vyc *VacancyYearCreate) SetNillableCalendarCode(i *int32) *VacancyYearCreate {
	if i != nil {
		vyc.SetCalendarCode(*i)
	}
	return vyc
}

// SetID sets the "id" field.
func (vyc *VacancyYearCreate) SetID(i int32) *VacancyYearCreate {
	vyc.mutation.SetID(i)
	return vyc
}

// AddVacancyRefIDs adds the "vacancy_ref" edge to the ExamCalendar entity by IDs.
func (vyc *VacancyYearCreate) AddVacancyRefIDs(ids ...int32) *VacancyYearCreate {
	vyc.mutation.AddVacancyRefIDs(ids...)
	return vyc
}

// AddVacancyRef adds the "vacancy_ref" edges to the ExamCalendar entity.
func (vyc *VacancyYearCreate) AddVacancyRef(e ...*ExamCalendar) *VacancyYearCreate {
	ids := make([]int32, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return vyc.AddVacancyRefIDs(ids...)
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (vyc *VacancyYearCreate) AddExamIDs(ids ...int32) *VacancyYearCreate {
	vyc.mutation.AddExamIDs(ids...)
	return vyc
}

// AddExams adds the "exams" edges to the Exam entity.
func (vyc *VacancyYearCreate) AddExams(e ...*Exam) *VacancyYearCreate {
	ids := make([]int32, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return vyc.AddExamIDs(ids...)
}

// Mutation returns the VacancyYearMutation object of the builder.
func (vyc *VacancyYearCreate) Mutation() *VacancyYearMutation {
	return vyc.mutation
}

// Save creates the VacancyYear in the database.
func (vyc *VacancyYearCreate) Save(ctx context.Context) (*VacancyYear, error) {
	return withHooks(ctx, vyc.sqlSave, vyc.mutation, vyc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vyc *VacancyYearCreate) SaveX(ctx context.Context) *VacancyYear {
	v, err := vyc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vyc *VacancyYearCreate) Exec(ctx context.Context) error {
	_, err := vyc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vyc *VacancyYearCreate) ExecX(ctx context.Context) {
	if err := vyc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vyc *VacancyYearCreate) check() error {
	if _, ok := vyc.mutation.FromDate(); !ok {
		return &ValidationError{Name: "FromDate", err: errors.New(`ent: missing required field "VacancyYear.FromDate"`)}
	}
	if _, ok := vyc.mutation.ToDate(); !ok {
		return &ValidationError{Name: "ToDate", err: errors.New(`ent: missing required field "VacancyYear.ToDate"`)}
	}
	if _, ok := vyc.mutation.VacancyYear(); !ok {
		return &ValidationError{Name: "VacancyYear", err: errors.New(`ent: missing required field "VacancyYear.VacancyYear"`)}
	}
	return nil
}

func (vyc *VacancyYearCreate) sqlSave(ctx context.Context) (*VacancyYear, error) {
	if err := vyc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vyc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vyc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	vyc.mutation.id = &_node.ID
	vyc.mutation.done = true
	return _node, nil
}

func (vyc *VacancyYearCreate) createSpec() (*VacancyYear, *sqlgraph.CreateSpec) {
	var (
		_node = &VacancyYear{config: vyc.config}
		_spec = sqlgraph.NewCreateSpec(vacancyyear.Table, sqlgraph.NewFieldSpec(vacancyyear.FieldID, field.TypeInt32))
	)
	if id, ok := vyc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vyc.mutation.FromDate(); ok {
		_spec.SetField(vacancyyear.FieldFromDate, field.TypeTime, value)
		_node.FromDate = value
	}
	if value, ok := vyc.mutation.ToDate(); ok {
		_spec.SetField(vacancyyear.FieldToDate, field.TypeTime, value)
		_node.ToDate = value
	}
	if value, ok := vyc.mutation.NotifyCode(); ok {
		_spec.SetField(vacancyyear.FieldNotifyCode, field.TypeInt32, value)
		_node.NotifyCode = value
	}
	if value, ok := vyc.mutation.VacancyYear(); ok {
		_spec.SetField(vacancyyear.FieldVacancyYear, field.TypeString, value)
		_node.VacancyYear = value
	}
	if value, ok := vyc.mutation.CalendarCode(); ok {
		_spec.SetField(vacancyyear.FieldCalendarCode, field.TypeInt32, value)
		_node.CalendarCode = value
	}
	if nodes := vyc.mutation.VacancyRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vacancyyear.VacancyRefTable,
			Columns: []string{vacancyyear.VacancyRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examcalendar.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vyc.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vacancyyear.ExamsTable,
			Columns: []string{vacancyyear.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VacancyYearCreateBulk is the builder for creating many VacancyYear entities in bulk.
type VacancyYearCreateBulk struct {
	config
	builders []*VacancyYearCreate
}

// Save creates the VacancyYear entities in the database.
func (vycb *VacancyYearCreateBulk) Save(ctx context.Context) ([]*VacancyYear, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vycb.builders))
	nodes := make([]*VacancyYear, len(vycb.builders))
	mutators := make([]Mutator, len(vycb.builders))
	for i := range vycb.builders {
		func(i int, root context.Context) {
			builder := vycb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VacancyYearMutation)
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
					_, err = mutators[i+1].Mutate(root, vycb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vycb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vycb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vycb *VacancyYearCreateBulk) SaveX(ctx context.Context) []*VacancyYear {
	v, err := vycb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vycb *VacancyYearCreateBulk) Exec(ctx context.Context) error {
	_, err := vycb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vycb *VacancyYearCreateBulk) ExecX(ctx context.Context) {
	if err := vycb.Exec(ctx); err != nil {
		panic(err)
	}
}