// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/postexampaper"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostExamPaperDelete is the builder for deleting a PostExamPaper entity.
type PostExamPaperDelete struct {
	config
	hooks    []Hook
	mutation *PostExamPaperMutation
}

// Where appends a list predicates to the PostExamPaperDelete builder.
func (pepd *PostExamPaperDelete) Where(ps ...predicate.PostExamPaper) *PostExamPaperDelete {
	pepd.mutation.Where(ps...)
	return pepd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pepd *PostExamPaperDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pepd.sqlExec, pepd.mutation, pepd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pepd *PostExamPaperDelete) ExecX(ctx context.Context) int {
	n, err := pepd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pepd *PostExamPaperDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(postexampaper.Table, sqlgraph.NewFieldSpec(postexampaper.FieldID, field.TypeInt32))
	if ps := pepd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pepd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pepd.mutation.done = true
	return affected, err
}

// PostExamPaperDeleteOne is the builder for deleting a single PostExamPaper entity.
type PostExamPaperDeleteOne struct {
	pepd *PostExamPaperDelete
}

// Where appends a list predicates to the PostExamPaperDelete builder.
func (pepdo *PostExamPaperDeleteOne) Where(ps ...predicate.PostExamPaper) *PostExamPaperDeleteOne {
	pepdo.pepd.mutation.Where(ps...)
	return pepdo
}

// Exec executes the deletion query.
func (pepdo *PostExamPaperDeleteOne) Exec(ctx context.Context) error {
	n, err := pepdo.pepd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{postexampaper.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pepdo *PostExamPaperDeleteOne) ExecX(ctx context.Context) {
	if err := pepdo.Exec(ctx); err != nil {
		panic(err)
	}
}
