// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/adminmaster"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminMasterDelete is the builder for deleting a AdminMaster entity.
type AdminMasterDelete struct {
	config
	hooks    []Hook
	mutation *AdminMasterMutation
}

// Where appends a list predicates to the AdminMasterDelete builder.
func (amd *AdminMasterDelete) Where(ps ...predicate.AdminMaster) *AdminMasterDelete {
	amd.mutation.Where(ps...)
	return amd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (amd *AdminMasterDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, amd.sqlExec, amd.mutation, amd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (amd *AdminMasterDelete) ExecX(ctx context.Context) int {
	n, err := amd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (amd *AdminMasterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(adminmaster.Table, sqlgraph.NewFieldSpec(adminmaster.FieldID, field.TypeInt64))
	if ps := amd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, amd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	amd.mutation.done = true
	return affected, err
}

// AdminMasterDeleteOne is the builder for deleting a single AdminMaster entity.
type AdminMasterDeleteOne struct {
	amd *AdminMasterDelete
}

// Where appends a list predicates to the AdminMasterDelete builder.
func (amdo *AdminMasterDeleteOne) Where(ps ...predicate.AdminMaster) *AdminMasterDeleteOne {
	amdo.amd.mutation.Where(ps...)
	return amdo
}

// Exec executes the deletion query.
func (amdo *AdminMasterDeleteOne) Exec(ctx context.Context) error {
	n, err := amdo.amd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{adminmaster.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (amdo *AdminMasterDeleteOne) ExecX(ctx context.Context) {
	if err := amdo.Exec(ctx); err != nil {
		panic(err)
	}
}
