// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/errorlogs"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ErrorLogsDelete is the builder for deleting a ErrorLogs entity.
type ErrorLogsDelete struct {
	config
	hooks    []Hook
	mutation *ErrorLogsMutation
}

// Where appends a list predicates to the ErrorLogsDelete builder.
func (eld *ErrorLogsDelete) Where(ps ...predicate.ErrorLogs) *ErrorLogsDelete {
	eld.mutation.Where(ps...)
	return eld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eld *ErrorLogsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, eld.sqlExec, eld.mutation, eld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (eld *ErrorLogsDelete) ExecX(ctx context.Context) int {
	n, err := eld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eld *ErrorLogsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(errorlogs.Table, sqlgraph.NewFieldSpec(errorlogs.FieldID, field.TypeInt64))
	if ps := eld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	eld.mutation.done = true
	return affected, err
}

// ErrorLogsDeleteOne is the builder for deleting a single ErrorLogs entity.
type ErrorLogsDeleteOne struct {
	eld *ErrorLogsDelete
}

// Where appends a list predicates to the ErrorLogsDelete builder.
func (eldo *ErrorLogsDeleteOne) Where(ps ...predicate.ErrorLogs) *ErrorLogsDeleteOne {
	eldo.eld.mutation.Where(ps...)
	return eldo
}

// Exec executes the deletion query.
func (eldo *ErrorLogsDeleteOne) Exec(ctx context.Context) error {
	n, err := eldo.eld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{errorlogs.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eldo *ErrorLogsDeleteOne) ExecX(ctx context.Context) {
	if err := eldo.Exec(ctx); err != nil {
		panic(err)
	}
}
