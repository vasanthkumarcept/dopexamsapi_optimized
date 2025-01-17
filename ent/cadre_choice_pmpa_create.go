// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/cadre_choice_pmpa"
	"recruit/ent/exam_applications_pmpa"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CadreChoicePMPACreate is the builder for creating a Cadre_Choice_PMPA entity.
type CadreChoicePMPACreate struct {
	config
	mutation *CadreChoicePMPAMutation
	hooks    []Hook
}

// SetApplicationID sets the "ApplicationID" field.
func (ccpc *CadreChoicePMPACreate) SetApplicationID(i int64) *CadreChoicePMPACreate {
	ccpc.mutation.SetApplicationID(i)
	return ccpc
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (ccpc *CadreChoicePMPACreate) SetNillableApplicationID(i *int64) *CadreChoicePMPACreate {
	if i != nil {
		ccpc.SetApplicationID(*i)
	}
	return ccpc
}

// SetPlacePrefNo sets the "PlacePrefNo" field.
func (ccpc *CadreChoicePMPACreate) SetPlacePrefNo(i int64) *CadreChoicePMPACreate {
	ccpc.mutation.SetPlacePrefNo(i)
	return ccpc
}

// SetPlacePrefValue sets the "PlacePrefValue" field.
func (ccpc *CadreChoicePMPACreate) SetPlacePrefValue(s string) *CadreChoicePMPACreate {
	ccpc.mutation.SetPlacePrefValue(s)
	return ccpc
}

// SetEmployeeID sets the "EmployeeID" field.
func (ccpc *CadreChoicePMPACreate) SetEmployeeID(i int64) *CadreChoicePMPACreate {
	ccpc.mutation.SetEmployeeID(i)
	return ccpc
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (ccpc *CadreChoicePMPACreate) SetNillableEmployeeID(i *int64) *CadreChoicePMPACreate {
	if i != nil {
		ccpc.SetEmployeeID(*i)
	}
	return ccpc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (ccpc *CadreChoicePMPACreate) SetUpdatedAt(t time.Time) *CadreChoicePMPACreate {
	ccpc.mutation.SetUpdatedAt(t)
	return ccpc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (ccpc *CadreChoicePMPACreate) SetNillableUpdatedAt(t *time.Time) *CadreChoicePMPACreate {
	if t != nil {
		ccpc.SetUpdatedAt(*t)
	}
	return ccpc
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (ccpc *CadreChoicePMPACreate) SetUpdatedBy(s string) *CadreChoicePMPACreate {
	ccpc.mutation.SetUpdatedBy(s)
	return ccpc
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (ccpc *CadreChoicePMPACreate) SetNillableUpdatedBy(s *string) *CadreChoicePMPACreate {
	if s != nil {
		ccpc.SetUpdatedBy(*s)
	}
	return ccpc
}

// SetID sets the "id" field.
func (ccpc *CadreChoicePMPACreate) SetID(i int32) *CadreChoicePMPACreate {
	ccpc.mutation.SetID(i)
	return ccpc
}

// SetApplnPMPARefID sets the "ApplnPMPA_Ref" edge to the Exam_Applications_PMPA entity by ID.
func (ccpc *CadreChoicePMPACreate) SetApplnPMPARefID(id int64) *CadreChoicePMPACreate {
	ccpc.mutation.SetApplnPMPARefID(id)
	return ccpc
}

// SetNillableApplnPMPARefID sets the "ApplnPMPA_Ref" edge to the Exam_Applications_PMPA entity by ID if the given value is not nil.
func (ccpc *CadreChoicePMPACreate) SetNillableApplnPMPARefID(id *int64) *CadreChoicePMPACreate {
	if id != nil {
		ccpc = ccpc.SetApplnPMPARefID(*id)
	}
	return ccpc
}

// SetApplnPMPARef sets the "ApplnPMPA_Ref" edge to the Exam_Applications_PMPA entity.
func (ccpc *CadreChoicePMPACreate) SetApplnPMPARef(e *Exam_Applications_PMPA) *CadreChoicePMPACreate {
	return ccpc.SetApplnPMPARefID(e.ID)
}

// Mutation returns the CadreChoicePMPAMutation object of the builder.
func (ccpc *CadreChoicePMPACreate) Mutation() *CadreChoicePMPAMutation {
	return ccpc.mutation
}

// Save creates the Cadre_Choice_PMPA in the database.
func (ccpc *CadreChoicePMPACreate) Save(ctx context.Context) (*Cadre_Choice_PMPA, error) {
	ccpc.defaults()
	return withHooks(ctx, ccpc.sqlSave, ccpc.mutation, ccpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ccpc *CadreChoicePMPACreate) SaveX(ctx context.Context) *Cadre_Choice_PMPA {
	v, err := ccpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccpc *CadreChoicePMPACreate) Exec(ctx context.Context) error {
	_, err := ccpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccpc *CadreChoicePMPACreate) ExecX(ctx context.Context) {
	if err := ccpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccpc *CadreChoicePMPACreate) defaults() {
	if _, ok := ccpc.mutation.UpdatedAt(); !ok {
		v := cadre_choice_pmpa.DefaultUpdatedAt()
		ccpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ccpc.mutation.UpdatedBy(); !ok {
		v := cadre_choice_pmpa.DefaultUpdatedBy
		ccpc.mutation.SetUpdatedBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccpc *CadreChoicePMPACreate) check() error {
	if _, ok := ccpc.mutation.PlacePrefNo(); !ok {
		return &ValidationError{Name: "PlacePrefNo", err: errors.New(`ent: missing required field "Cadre_Choice_PMPA.PlacePrefNo"`)}
	}
	if _, ok := ccpc.mutation.PlacePrefValue(); !ok {
		return &ValidationError{Name: "PlacePrefValue", err: errors.New(`ent: missing required field "Cadre_Choice_PMPA.PlacePrefValue"`)}
	}
	return nil
}

func (ccpc *CadreChoicePMPACreate) sqlSave(ctx context.Context) (*Cadre_Choice_PMPA, error) {
	if err := ccpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ccpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	ccpc.mutation.id = &_node.ID
	ccpc.mutation.done = true
	return _node, nil
}

func (ccpc *CadreChoicePMPACreate) createSpec() (*Cadre_Choice_PMPA, *sqlgraph.CreateSpec) {
	var (
		_node = &Cadre_Choice_PMPA{config: ccpc.config}
		_spec = sqlgraph.NewCreateSpec(cadre_choice_pmpa.Table, sqlgraph.NewFieldSpec(cadre_choice_pmpa.FieldID, field.TypeInt32))
	)
	if id, ok := ccpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ccpc.mutation.PlacePrefNo(); ok {
		_spec.SetField(cadre_choice_pmpa.FieldPlacePrefNo, field.TypeInt64, value)
		_node.PlacePrefNo = value
	}
	if value, ok := ccpc.mutation.PlacePrefValue(); ok {
		_spec.SetField(cadre_choice_pmpa.FieldPlacePrefValue, field.TypeString, value)
		_node.PlacePrefValue = value
	}
	if value, ok := ccpc.mutation.EmployeeID(); ok {
		_spec.SetField(cadre_choice_pmpa.FieldEmployeeID, field.TypeInt64, value)
		_node.EmployeeID = value
	}
	if value, ok := ccpc.mutation.UpdatedAt(); ok {
		_spec.SetField(cadre_choice_pmpa.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ccpc.mutation.UpdatedBy(); ok {
		_spec.SetField(cadre_choice_pmpa.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if nodes := ccpc.mutation.ApplnPMPARefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cadre_choice_pmpa.ApplnPMPARefTable,
			Columns: []string{cadre_choice_pmpa.ApplnPMPARefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_pmpa.FieldID, field.TypeInt64),
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

// CadreChoicePMPACreateBulk is the builder for creating many Cadre_Choice_PMPA entities in bulk.
type CadreChoicePMPACreateBulk struct {
	config
	builders []*CadreChoicePMPACreate
}

// Save creates the Cadre_Choice_PMPA entities in the database.
func (ccpcb *CadreChoicePMPACreateBulk) Save(ctx context.Context) ([]*Cadre_Choice_PMPA, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccpcb.builders))
	nodes := make([]*Cadre_Choice_PMPA, len(ccpcb.builders))
	mutators := make([]Mutator, len(ccpcb.builders))
	for i := range ccpcb.builders {
		func(i int, root context.Context) {
			builder := ccpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CadreChoicePMPAMutation)
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
					_, err = mutators[i+1].Mutate(root, ccpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccpcb *CadreChoicePMPACreateBulk) SaveX(ctx context.Context) []*Cadre_Choice_PMPA {
	v, err := ccpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccpcb *CadreChoicePMPACreateBulk) Exec(ctx context.Context) error {
	_, err := ccpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccpcb *CadreChoicePMPACreateBulk) ExecX(ctx context.Context) {
	if err := ccpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
