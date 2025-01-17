// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/placeofpreferenceps"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaceOfPreferencePSCreate is the builder for creating a PlaceOfPreferencePS entity.
type PlaceOfPreferencePSCreate struct {
	config
	mutation *PlaceOfPreferencePSMutation
	hooks    []Hook
}

// SetApplicationID sets the "ApplicationID" field.
func (poppc *PlaceOfPreferencePSCreate) SetApplicationID(i int64) *PlaceOfPreferencePSCreate {
	poppc.mutation.SetApplicationID(i)
	return poppc
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (poppc *PlaceOfPreferencePSCreate) SetNillableApplicationID(i *int64) *PlaceOfPreferencePSCreate {
	if i != nil {
		poppc.SetApplicationID(*i)
	}
	return poppc
}

// SetPlacePrefNo sets the "PlacePrefNo" field.
func (poppc *PlaceOfPreferencePSCreate) SetPlacePrefNo(i int32) *PlaceOfPreferencePSCreate {
	poppc.mutation.SetPlacePrefNo(i)
	return poppc
}

// SetNillablePlacePrefNo sets the "PlacePrefNo" field if the given value is not nil.
func (poppc *PlaceOfPreferencePSCreate) SetNillablePlacePrefNo(i *int32) *PlaceOfPreferencePSCreate {
	if i != nil {
		poppc.SetPlacePrefNo(*i)
	}
	return poppc
}

// SetPlacePrefValue sets the "PlacePrefValue" field.
func (poppc *PlaceOfPreferencePSCreate) SetPlacePrefValue(s string) *PlaceOfPreferencePSCreate {
	poppc.mutation.SetPlacePrefValue(s)
	return poppc
}

// SetNillablePlacePrefValue sets the "PlacePrefValue" field if the given value is not nil.
func (poppc *PlaceOfPreferencePSCreate) SetNillablePlacePrefValue(s *string) *PlaceOfPreferencePSCreate {
	if s != nil {
		poppc.SetPlacePrefValue(*s)
	}
	return poppc
}

// SetEmployeeID sets the "EmployeeID" field.
func (poppc *PlaceOfPreferencePSCreate) SetEmployeeID(i int64) *PlaceOfPreferencePSCreate {
	poppc.mutation.SetEmployeeID(i)
	return poppc
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (poppc *PlaceOfPreferencePSCreate) SetNillableEmployeeID(i *int64) *PlaceOfPreferencePSCreate {
	if i != nil {
		poppc.SetEmployeeID(*i)
	}
	return poppc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (poppc *PlaceOfPreferencePSCreate) SetUpdatedAt(t time.Time) *PlaceOfPreferencePSCreate {
	poppc.mutation.SetUpdatedAt(t)
	return poppc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (poppc *PlaceOfPreferencePSCreate) SetNillableUpdatedAt(t *time.Time) *PlaceOfPreferencePSCreate {
	if t != nil {
		poppc.SetUpdatedAt(*t)
	}
	return poppc
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (poppc *PlaceOfPreferencePSCreate) SetUpdatedBy(s string) *PlaceOfPreferencePSCreate {
	poppc.mutation.SetUpdatedBy(s)
	return poppc
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (poppc *PlaceOfPreferencePSCreate) SetNillableUpdatedBy(s *string) *PlaceOfPreferencePSCreate {
	if s != nil {
		poppc.SetUpdatedBy(*s)
	}
	return poppc
}

// SetID sets the "id" field.
func (poppc *PlaceOfPreferencePSCreate) SetID(i int32) *PlaceOfPreferencePSCreate {
	poppc.mutation.SetID(i)
	return poppc
}

// SetPlaceApplnPSRefID sets the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity by ID.
func (poppc *PlaceOfPreferencePSCreate) SetPlaceApplnPSRefID(id int64) *PlaceOfPreferencePSCreate {
	poppc.mutation.SetPlaceApplnPSRefID(id)
	return poppc
}

// SetNillablePlaceApplnPSRefID sets the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity by ID if the given value is not nil.
func (poppc *PlaceOfPreferencePSCreate) SetNillablePlaceApplnPSRefID(id *int64) *PlaceOfPreferencePSCreate {
	if id != nil {
		poppc = poppc.SetPlaceApplnPSRefID(*id)
	}
	return poppc
}

// SetPlaceApplnPSRef sets the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity.
func (poppc *PlaceOfPreferencePSCreate) SetPlaceApplnPSRef(e *Exam_Applications_PS) *PlaceOfPreferencePSCreate {
	return poppc.SetPlaceApplnPSRefID(e.ID)
}

// Mutation returns the PlaceOfPreferencePSMutation object of the builder.
func (poppc *PlaceOfPreferencePSCreate) Mutation() *PlaceOfPreferencePSMutation {
	return poppc.mutation
}

// Save creates the PlaceOfPreferencePS in the database.
func (poppc *PlaceOfPreferencePSCreate) Save(ctx context.Context) (*PlaceOfPreferencePS, error) {
	poppc.defaults()
	return withHooks(ctx, poppc.sqlSave, poppc.mutation, poppc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (poppc *PlaceOfPreferencePSCreate) SaveX(ctx context.Context) *PlaceOfPreferencePS {
	v, err := poppc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (poppc *PlaceOfPreferencePSCreate) Exec(ctx context.Context) error {
	_, err := poppc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (poppc *PlaceOfPreferencePSCreate) ExecX(ctx context.Context) {
	if err := poppc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (poppc *PlaceOfPreferencePSCreate) defaults() {
	if _, ok := poppc.mutation.UpdatedBy(); !ok {
		v := placeofpreferenceps.DefaultUpdatedBy
		poppc.mutation.SetUpdatedBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (poppc *PlaceOfPreferencePSCreate) check() error {
	return nil
}

func (poppc *PlaceOfPreferencePSCreate) sqlSave(ctx context.Context) (*PlaceOfPreferencePS, error) {
	if err := poppc.check(); err != nil {
		return nil, err
	}
	_node, _spec := poppc.createSpec()
	if err := sqlgraph.CreateNode(ctx, poppc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	poppc.mutation.id = &_node.ID
	poppc.mutation.done = true
	return _node, nil
}

func (poppc *PlaceOfPreferencePSCreate) createSpec() (*PlaceOfPreferencePS, *sqlgraph.CreateSpec) {
	var (
		_node = &PlaceOfPreferencePS{config: poppc.config}
		_spec = sqlgraph.NewCreateSpec(placeofpreferenceps.Table, sqlgraph.NewFieldSpec(placeofpreferenceps.FieldID, field.TypeInt32))
	)
	if id, ok := poppc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := poppc.mutation.PlacePrefNo(); ok {
		_spec.SetField(placeofpreferenceps.FieldPlacePrefNo, field.TypeInt32, value)
		_node.PlacePrefNo = value
	}
	if value, ok := poppc.mutation.PlacePrefValue(); ok {
		_spec.SetField(placeofpreferenceps.FieldPlacePrefValue, field.TypeString, value)
		_node.PlacePrefValue = value
	}
	if value, ok := poppc.mutation.EmployeeID(); ok {
		_spec.SetField(placeofpreferenceps.FieldEmployeeID, field.TypeInt64, value)
		_node.EmployeeID = value
	}
	if value, ok := poppc.mutation.UpdatedAt(); ok {
		_spec.SetField(placeofpreferenceps.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := poppc.mutation.UpdatedBy(); ok {
		_spec.SetField(placeofpreferenceps.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if nodes := poppc.mutation.PlaceApplnPSRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placeofpreferenceps.PlaceApplnPSRefTable,
			Columns: []string{placeofpreferenceps.PlaceApplnPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ApplicationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlaceOfPreferencePSCreateBulk is the builder for creating many PlaceOfPreferencePS entities in bulk.
type PlaceOfPreferencePSCreateBulk struct {
	config
	builders []*PlaceOfPreferencePSCreate
}

// Save creates the PlaceOfPreferencePS entities in the database.
func (poppcb *PlaceOfPreferencePSCreateBulk) Save(ctx context.Context) ([]*PlaceOfPreferencePS, error) {
	specs := make([]*sqlgraph.CreateSpec, len(poppcb.builders))
	nodes := make([]*PlaceOfPreferencePS, len(poppcb.builders))
	mutators := make([]Mutator, len(poppcb.builders))
	for i := range poppcb.builders {
		func(i int, root context.Context) {
			builder := poppcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaceOfPreferencePSMutation)
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
					_, err = mutators[i+1].Mutate(root, poppcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, poppcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, poppcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (poppcb *PlaceOfPreferencePSCreateBulk) SaveX(ctx context.Context) []*PlaceOfPreferencePS {
	v, err := poppcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (poppcb *PlaceOfPreferencePSCreateBulk) Exec(ctx context.Context) error {
	_, err := poppcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (poppcb *PlaceOfPreferencePSCreateBulk) ExecX(ctx context.Context) {
	if err := poppcb.Exec(ctx); err != nil {
		panic(err)
	}
}
