// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/exam_applications_gdspa"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamApplicationsGDSPADelete is the builder for deleting a Exam_Applications_GDSPA entity.
type ExamApplicationsGDSPADelete struct {
	config
	hooks    []Hook
	mutation *ExamApplicationsGDSPAMutation
}

// Where appends a list predicates to the ExamApplicationsGDSPADelete builder.
func (eagd *ExamApplicationsGDSPADelete) Where(ps ...predicate.Exam_Applications_GDSPA) *ExamApplicationsGDSPADelete {
	eagd.mutation.Where(ps...)
	return eagd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eagd *ExamApplicationsGDSPADelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, eagd.sqlExec, eagd.mutation, eagd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (eagd *ExamApplicationsGDSPADelete) ExecX(ctx context.Context) int {
	n, err := eagd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eagd *ExamApplicationsGDSPADelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(exam_applications_gdspa.Table, sqlgraph.NewFieldSpec(exam_applications_gdspa.FieldID, field.TypeInt64))
	if ps := eagd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eagd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	eagd.mutation.done = true
	return affected, err
}

// ExamApplicationsGDSPADeleteOne is the builder for deleting a single Exam_Applications_GDSPA entity.
type ExamApplicationsGDSPADeleteOne struct {
	eagd *ExamApplicationsGDSPADelete
}

// Where appends a list predicates to the ExamApplicationsGDSPADelete builder.
func (eagdo *ExamApplicationsGDSPADeleteOne) Where(ps ...predicate.Exam_Applications_GDSPA) *ExamApplicationsGDSPADeleteOne {
	eagdo.eagd.mutation.Where(ps...)
	return eagdo
}

// Exec executes the deletion query.
func (eagdo *ExamApplicationsGDSPADeleteOne) Exec(ctx context.Context) error {
	n, err := eagdo.eagd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{exam_applications_gdspa.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eagdo *ExamApplicationsGDSPADeleteOne) ExecX(ctx context.Context) {
	if err := eagdo.Exec(ctx); err != nil {
		panic(err)
	}
}
