// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/placeofpreferenceip"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaceOfPreferenceIPCreate is the builder for creating a PlaceOfPreferenceIP entity.
type PlaceOfPreferenceIPCreate struct {
	config
	mutation *PlaceOfPreferenceIPMutation
	hooks    []Hook
}

// SetApplicationID sets the "ApplicationID" field.
func (popic *PlaceOfPreferenceIPCreate) SetApplicationID(i int64) *PlaceOfPreferenceIPCreate {
	popic.mutation.SetApplicationID(i)
	return popic
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (popic *PlaceOfPreferenceIPCreate) SetNillableApplicationID(i *int64) *PlaceOfPreferenceIPCreate {
	if i != nil {
		popic.SetApplicationID(*i)
	}
	return popic
}

// SetPlacePrefNo sets the "PlacePrefNo" field.
func (popic *PlaceOfPreferenceIPCreate) SetPlacePrefNo(i int32) *PlaceOfPreferenceIPCreate {
	popic.mutation.SetPlacePrefNo(i)
	return popic
}

// SetNillablePlacePrefNo sets the "PlacePrefNo" field if the given value is not nil.
func (popic *PlaceOfPreferenceIPCreate) SetNillablePlacePrefNo(i *int32) *PlaceOfPreferenceIPCreate {
	if i != nil {
		popic.SetPlacePrefNo(*i)
	}
	return popic
}

// SetPlacePrefValue sets the "PlacePrefValue" field.
func (popic *PlaceOfPreferenceIPCreate) SetPlacePrefValue(s string) *PlaceOfPreferenceIPCreate {
	popic.mutation.SetPlacePrefValue(s)
	return popic
}

// SetNillablePlacePrefValue sets the "PlacePrefValue" field if the given value is not nil.
func (popic *PlaceOfPreferenceIPCreate) SetNillablePlacePrefValue(s *string) *PlaceOfPreferenceIPCreate {
	if s != nil {
		popic.SetPlacePrefValue(*s)
	}
	return popic
}

// SetEmployeeID sets the "EmployeeID" field.
func (popic *PlaceOfPreferenceIPCreate) SetEmployeeID(i int64) *PlaceOfPreferenceIPCreate {
	popic.mutation.SetEmployeeID(i)
	return popic
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (popic *PlaceOfPreferenceIPCreate) SetNillableEmployeeID(i *int64) *PlaceOfPreferenceIPCreate {
	if i != nil {
		popic.SetEmployeeID(*i)
	}
	return popic
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (popic *PlaceOfPreferenceIPCreate) SetUpdatedAt(t time.Time) *PlaceOfPreferenceIPCreate {
	popic.mutation.SetUpdatedAt(t)
	return popic
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (popic *PlaceOfPreferenceIPCreate) SetNillableUpdatedAt(t *time.Time) *PlaceOfPreferenceIPCreate {
	if t != nil {
		popic.SetUpdatedAt(*t)
	}
	return popic
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (popic *PlaceOfPreferenceIPCreate) SetUpdatedBy(s string) *PlaceOfPreferenceIPCreate {
	popic.mutation.SetUpdatedBy(s)
	return popic
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (popic *PlaceOfPreferenceIPCreate) SetNillableUpdatedBy(s *string) *PlaceOfPreferenceIPCreate {
	if s != nil {
		popic.SetUpdatedBy(*s)
	}
	return popic
}

// SetID sets the "id" field.
func (popic *PlaceOfPreferenceIPCreate) SetID(i int32) *PlaceOfPreferenceIPCreate {
	popic.mutation.SetID(i)
	return popic
}

// SetApplnIPRefID sets the "ApplnIP_Ref" edge to the Exam_Applications_IP entity by ID.
func (popic *PlaceOfPreferenceIPCreate) SetApplnIPRefID(id int64) *PlaceOfPreferenceIPCreate {
	popic.mutation.SetApplnIPRefID(id)
	return popic
}

// SetNillableApplnIPRefID sets the "ApplnIP_Ref" edge to the Exam_Applications_IP entity by ID if the given value is not nil.
func (popic *PlaceOfPreferenceIPCreate) SetNillableApplnIPRefID(id *int64) *PlaceOfPreferenceIPCreate {
	if id != nil {
		popic = popic.SetApplnIPRefID(*id)
	}
	return popic
}

// SetApplnIPRef sets the "ApplnIP_Ref" edge to the Exam_Applications_IP entity.
func (popic *PlaceOfPreferenceIPCreate) SetApplnIPRef(e *Exam_Applications_IP) *PlaceOfPreferenceIPCreate {
	return popic.SetApplnIPRefID(e.ID)
}

// Mutation returns the PlaceOfPreferenceIPMutation object of the builder.
func (popic *PlaceOfPreferenceIPCreate) Mutation() *PlaceOfPreferenceIPMutation {
	return popic.mutation
}

// Save creates the PlaceOfPreferenceIP in the database.
func (popic *PlaceOfPreferenceIPCreate) Save(ctx context.Context) (*PlaceOfPreferenceIP, error) {
	popic.defaults()
	return withHooks(ctx, popic.sqlSave, popic.mutation, popic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (popic *PlaceOfPreferenceIPCreate) SaveX(ctx context.Context) *PlaceOfPreferenceIP {
	v, err := popic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (popic *PlaceOfPreferenceIPCreate) Exec(ctx context.Context) error {
	_, err := popic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (popic *PlaceOfPreferenceIPCreate) ExecX(ctx context.Context) {
	if err := popic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (popic *PlaceOfPreferenceIPCreate) defaults() {
	if _, ok := popic.mutation.UpdatedBy(); !ok {
		v := placeofpreferenceip.DefaultUpdatedBy
		popic.mutation.SetUpdatedBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (popic *PlaceOfPreferenceIPCreate) check() error {
	return nil
}

func (popic *PlaceOfPreferenceIPCreate) sqlSave(ctx context.Context) (*PlaceOfPreferenceIP, error) {
	if err := popic.check(); err != nil {
		return nil, err
	}
	_node, _spec := popic.createSpec()
	if err := sqlgraph.CreateNode(ctx, popic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	popic.mutation.id = &_node.ID
	popic.mutation.done = true
	return _node, nil
}

func (popic *PlaceOfPreferenceIPCreate) createSpec() (*PlaceOfPreferenceIP, *sqlgraph.CreateSpec) {
	var (
		_node = &PlaceOfPreferenceIP{config: popic.config}
		_spec = sqlgraph.NewCreateSpec(placeofpreferenceip.Table, sqlgraph.NewFieldSpec(placeofpreferenceip.FieldID, field.TypeInt32))
	)
	if id, ok := popic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := popic.mutation.ApplicationID(); ok {
		_spec.SetField(placeofpreferenceip.FieldApplicationID, field.TypeInt64, value)
		_node.ApplicationID = value
	}
	if value, ok := popic.mutation.PlacePrefNo(); ok {
		_spec.SetField(placeofpreferenceip.FieldPlacePrefNo, field.TypeInt32, value)
		_node.PlacePrefNo = value
	}
	if value, ok := popic.mutation.PlacePrefValue(); ok {
		_spec.SetField(placeofpreferenceip.FieldPlacePrefValue, field.TypeString, value)
		_node.PlacePrefValue = value
	}
	if value, ok := popic.mutation.EmployeeID(); ok {
		_spec.SetField(placeofpreferenceip.FieldEmployeeID, field.TypeInt64, value)
		_node.EmployeeID = value
	}
	if value, ok := popic.mutation.UpdatedAt(); ok {
		_spec.SetField(placeofpreferenceip.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := popic.mutation.UpdatedBy(); ok {
		_spec.SetField(placeofpreferenceip.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if nodes := popic.mutation.ApplnIPRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placeofpreferenceip.ApplnIPRefTable,
			Columns: []string{placeofpreferenceip.ApplnIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.exam_applications_ip_circle_pref_ref = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlaceOfPreferenceIPCreateBulk is the builder for creating many PlaceOfPreferenceIP entities in bulk.
type PlaceOfPreferenceIPCreateBulk struct {
	config
	builders []*PlaceOfPreferenceIPCreate
}

// Save creates the PlaceOfPreferenceIP entities in the database.
func (popicb *PlaceOfPreferenceIPCreateBulk) Save(ctx context.Context) ([]*PlaceOfPreferenceIP, error) {
	specs := make([]*sqlgraph.CreateSpec, len(popicb.builders))
	nodes := make([]*PlaceOfPreferenceIP, len(popicb.builders))
	mutators := make([]Mutator, len(popicb.builders))
	for i := range popicb.builders {
		func(i int, root context.Context) {
			builder := popicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaceOfPreferenceIPMutation)
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
					_, err = mutators[i+1].Mutate(root, popicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, popicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, popicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (popicb *PlaceOfPreferenceIPCreateBulk) SaveX(ctx context.Context) []*PlaceOfPreferenceIP {
	v, err := popicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (popicb *PlaceOfPreferenceIPCreateBulk) Exec(ctx context.Context) error {
	_, err := popicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (popicb *PlaceOfPreferenceIPCreateBulk) ExecX(ctx context.Context) {
	if err := popicb.Exec(ctx); err != nil {
		panic(err)
	}
}
