// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/cadreeligibleconfiguration"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CadreEligibleConfigurationDelete is the builder for deleting a CadreEligibleConfiguration entity.
type CadreEligibleConfigurationDelete struct {
	config
	hooks    []Hook
	mutation *CadreEligibleConfigurationMutation
}

// Where appends a list predicates to the CadreEligibleConfigurationDelete builder.
func (cecd *CadreEligibleConfigurationDelete) Where(ps ...predicate.CadreEligibleConfiguration) *CadreEligibleConfigurationDelete {
	cecd.mutation.Where(ps...)
	return cecd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cecd *CadreEligibleConfigurationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cecd.sqlExec, cecd.mutation, cecd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cecd *CadreEligibleConfigurationDelete) ExecX(ctx context.Context) int {
	n, err := cecd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cecd *CadreEligibleConfigurationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cadreeligibleconfiguration.Table, sqlgraph.NewFieldSpec(cadreeligibleconfiguration.FieldID, field.TypeInt64))
	if ps := cecd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cecd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cecd.mutation.done = true
	return affected, err
}

// CadreEligibleConfigurationDeleteOne is the builder for deleting a single CadreEligibleConfiguration entity.
type CadreEligibleConfigurationDeleteOne struct {
	cecd *CadreEligibleConfigurationDelete
}

// Where appends a list predicates to the CadreEligibleConfigurationDelete builder.
func (cecdo *CadreEligibleConfigurationDeleteOne) Where(ps ...predicate.CadreEligibleConfiguration) *CadreEligibleConfigurationDeleteOne {
	cecdo.cecd.mutation.Where(ps...)
	return cecdo
}

// Exec executes the deletion query.
func (cecdo *CadreEligibleConfigurationDeleteOne) Exec(ctx context.Context) error {
	n, err := cecdo.cecd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cadreeligibleconfiguration.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cecdo *CadreEligibleConfigurationDeleteOne) ExecX(ctx context.Context) {
	if err := cecdo.Exec(ctx); err != nil {
		panic(err)
	}
}