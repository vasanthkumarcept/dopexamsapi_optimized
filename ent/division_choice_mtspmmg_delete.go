// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/division_choice_mtspmmg"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DivisionChoiceMTSPMMGDelete is the builder for deleting a Division_Choice_MTSPMMG entity.
type DivisionChoiceMTSPMMGDelete struct {
	config
	hooks    []Hook
	mutation *DivisionChoiceMTSPMMGMutation
}

// Where appends a list predicates to the DivisionChoiceMTSPMMGDelete builder.
func (dcmd *DivisionChoiceMTSPMMGDelete) Where(ps ...predicate.Division_Choice_MTSPMMG) *DivisionChoiceMTSPMMGDelete {
	dcmd.mutation.Where(ps...)
	return dcmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dcmd *DivisionChoiceMTSPMMGDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dcmd.sqlExec, dcmd.mutation, dcmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dcmd *DivisionChoiceMTSPMMGDelete) ExecX(ctx context.Context) int {
	n, err := dcmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dcmd *DivisionChoiceMTSPMMGDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(division_choice_mtspmmg.Table, sqlgraph.NewFieldSpec(division_choice_mtspmmg.FieldID, field.TypeInt32))
	if ps := dcmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dcmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dcmd.mutation.done = true
	return affected, err
}

// DivisionChoiceMTSPMMGDeleteOne is the builder for deleting a single Division_Choice_MTSPMMG entity.
type DivisionChoiceMTSPMMGDeleteOne struct {
	dcmd *DivisionChoiceMTSPMMGDelete
}

// Where appends a list predicates to the DivisionChoiceMTSPMMGDelete builder.
func (dcmdo *DivisionChoiceMTSPMMGDeleteOne) Where(ps ...predicate.Division_Choice_MTSPMMG) *DivisionChoiceMTSPMMGDeleteOne {
	dcmdo.dcmd.mutation.Where(ps...)
	return dcmdo
}

// Exec executes the deletion query.
func (dcmdo *DivisionChoiceMTSPMMGDeleteOne) Exec(ctx context.Context) error {
	n, err := dcmdo.dcmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{division_choice_mtspmmg.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dcmdo *DivisionChoiceMTSPMMGDeleteOne) ExecX(ctx context.Context) {
	if err := dcmdo.Exec(ctx); err != nil {
		panic(err)
	}
}