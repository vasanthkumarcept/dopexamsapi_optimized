// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/educationdetails"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EducationDetailsDelete is the builder for deleting a EducationDetails entity.
type EducationDetailsDelete struct {
	config
	hooks    []Hook
	mutation *EducationDetailsMutation
}

// Where appends a list predicates to the EducationDetailsDelete builder.
func (edd *EducationDetailsDelete) Where(ps ...predicate.EducationDetails) *EducationDetailsDelete {
	edd.mutation.Where(ps...)
	return edd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (edd *EducationDetailsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, edd.sqlExec, edd.mutation, edd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (edd *EducationDetailsDelete) ExecX(ctx context.Context) int {
	n, err := edd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (edd *EducationDetailsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(educationdetails.Table, sqlgraph.NewFieldSpec(educationdetails.FieldID, field.TypeInt64))
	if ps := edd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, edd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	edd.mutation.done = true
	return affected, err
}

// EducationDetailsDeleteOne is the builder for deleting a single EducationDetails entity.
type EducationDetailsDeleteOne struct {
	edd *EducationDetailsDelete
}

// Where appends a list predicates to the EducationDetailsDelete builder.
func (eddo *EducationDetailsDeleteOne) Where(ps ...predicate.EducationDetails) *EducationDetailsDeleteOne {
	eddo.edd.mutation.Where(ps...)
	return eddo
}

// Exec executes the deletion query.
func (eddo *EducationDetailsDeleteOne) Exec(ctx context.Context) error {
	n, err := eddo.edd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{educationdetails.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eddo *EducationDetailsDeleteOne) ExecX(ctx context.Context) {
	if err := eddo.Exec(ctx); err != nil {
		panic(err)
	}
}