// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"recruit/ent/eligibilitycadrepaymatrix"
	"recruit/ent/logs"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EligibilityCadrePayMatrixCreate is the builder for creating a EligibilityCadrePayMatrix entity.
type EligibilityCadrePayMatrixCreate struct {
	config
	mutation *EligibilityCadrePayMatrixMutation
	hooks    []Hook
}

// SetCadreEligibleConfigurationCadreEligibleCode sets the "cadreEligibleConfiguration_cadreEligibleCode" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetCadreEligibleConfigurationCadreEligibleCode(i int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetCadreEligibleConfigurationCadreEligibleCode(i)
	return ecpmc
}

// SetNillableCadreEligibleConfigurationCadreEligibleCode sets the "cadreEligibleConfiguration_cadreEligibleCode" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableCadreEligibleConfigurationCadreEligibleCode(i *int64) *EligibilityCadrePayMatrixCreate {
	if i != nil {
		ecpmc.SetCadreEligibleConfigurationCadreEligibleCode(*i)
	}
	return ecpmc
}

// SetPostId sets the "PostId" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetPostId(i int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetPostId(i)
	return ecpmc
}

// SetNillablePostId sets the "PostId" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillablePostId(i *int64) *EligibilityCadrePayMatrixCreate {
	if i != nil {
		ecpmc.SetPostId(*i)
	}
	return ecpmc
}

// SetPostCode sets the "PostCode" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetPostCode(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetPostCode(s)
	return ecpmc
}

// SetNillablePostCode sets the "PostCode" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillablePostCode(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetPostCode(*s)
	}
	return ecpmc
}

// SetPostDescription sets the "PostDescription" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetPostDescription(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetPostDescription(s)
	return ecpmc
}

// SetNillablePostDescription sets the "PostDescription" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillablePostDescription(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetPostDescription(*s)
	}
	return ecpmc
}

// SetOrderNumber sets the "OrderNumber" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetOrderNumber(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetOrderNumber(s)
	return ecpmc
}

// SetNillableOrderNumber sets the "OrderNumber" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableOrderNumber(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetOrderNumber(*s)
	}
	return ecpmc
}

// SetCreatedById sets the "CreatedById" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetCreatedById(i int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetCreatedById(i)
	return ecpmc
}

// SetNillableCreatedById sets the "CreatedById" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableCreatedById(i *int64) *EligibilityCadrePayMatrixCreate {
	if i != nil {
		ecpmc.SetCreatedById(*i)
	}
	return ecpmc
}

// SetCreatedByUserName sets the "CreatedByUserName" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetCreatedByUserName(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetCreatedByUserName(s)
	return ecpmc
}

// SetNillableCreatedByUserName sets the "CreatedByUserName" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableCreatedByUserName(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetCreatedByUserName(*s)
	}
	return ecpmc
}

// SetCreatedByEmpId sets the "CreatedByEmpId" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetCreatedByEmpId(i int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetCreatedByEmpId(i)
	return ecpmc
}

// SetNillableCreatedByEmpId sets the "CreatedByEmpId" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableCreatedByEmpId(i *int64) *EligibilityCadrePayMatrixCreate {
	if i != nil {
		ecpmc.SetCreatedByEmpId(*i)
	}
	return ecpmc
}

// SetCreatedByDesignation sets the "CreatedByDesignation" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetCreatedByDesignation(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetCreatedByDesignation(s)
	return ecpmc
}

// SetNillableCreatedByDesignation sets the "CreatedByDesignation" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableCreatedByDesignation(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetCreatedByDesignation(*s)
	}
	return ecpmc
}

// SetCreatedDate sets the "CreatedDate" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetCreatedDate(t time.Time) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetCreatedDate(t)
	return ecpmc
}

// SetNillableCreatedDate sets the "CreatedDate" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableCreatedDate(t *time.Time) *EligibilityCadrePayMatrixCreate {
	if t != nil {
		ecpmc.SetCreatedDate(*t)
	}
	return ecpmc
}

// SetVerifiedbyid sets the "verifiedbyid" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetVerifiedbyid(i int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetVerifiedbyid(i)
	return ecpmc
}

// SetNillableVerifiedbyid sets the "verifiedbyid" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableVerifiedbyid(i *int64) *EligibilityCadrePayMatrixCreate {
	if i != nil {
		ecpmc.SetVerifiedbyid(*i)
	}
	return ecpmc
}

// SetVerifiedbyusername sets the "verifiedbyusername" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetVerifiedbyusername(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetVerifiedbyusername(s)
	return ecpmc
}

// SetNillableVerifiedbyusername sets the "verifiedbyusername" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableVerifiedbyusername(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetVerifiedbyusername(*s)
	}
	return ecpmc
}

// SetVerifiedbyEmployeeid sets the "verifiedbyEmployeeid" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetVerifiedbyEmployeeid(i int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetVerifiedbyEmployeeid(i)
	return ecpmc
}

// SetNillableVerifiedbyEmployeeid sets the "verifiedbyEmployeeid" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableVerifiedbyEmployeeid(i *int64) *EligibilityCadrePayMatrixCreate {
	if i != nil {
		ecpmc.SetVerifiedbyEmployeeid(*i)
	}
	return ecpmc
}

// SetVerifiedbyDesignation sets the "verifiedbyDesignation" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetVerifiedbyDesignation(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetVerifiedbyDesignation(s)
	return ecpmc
}

// SetNillableVerifiedbyDesignation sets the "verifiedbyDesignation" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableVerifiedbyDesignation(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetVerifiedbyDesignation(*s)
	}
	return ecpmc
}

// SetVerifiedDate sets the "verifiedDate" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetVerifiedDate(t time.Time) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetVerifiedDate(t)
	return ecpmc
}

// SetNillableVerifiedDate sets the "verifiedDate" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableVerifiedDate(t *time.Time) *EligibilityCadrePayMatrixCreate {
	if t != nil {
		ecpmc.SetVerifiedDate(*t)
	}
	return ecpmc
}

// SetStatuss sets the "Statuss" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetStatuss(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetStatuss(s)
	return ecpmc
}

// SetNillableStatuss sets the "Statuss" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableStatuss(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetStatuss(*s)
	}
	return ecpmc
}

// SetDeletedbyid sets the "deletedbyid" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetDeletedbyid(i int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetDeletedbyid(i)
	return ecpmc
}

// SetNillableDeletedbyid sets the "deletedbyid" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableDeletedbyid(i *int64) *EligibilityCadrePayMatrixCreate {
	if i != nil {
		ecpmc.SetDeletedbyid(*i)
	}
	return ecpmc
}

// SetDeletedbyusername sets the "deletedbyusername" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetDeletedbyusername(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetDeletedbyusername(s)
	return ecpmc
}

// SetNillableDeletedbyusername sets the "deletedbyusername" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableDeletedbyusername(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetDeletedbyusername(*s)
	}
	return ecpmc
}

// SetDeletedbyEmployeeid sets the "deletedbyEmployeeid" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetDeletedbyEmployeeid(i int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetDeletedbyEmployeeid(i)
	return ecpmc
}

// SetNillableDeletedbyEmployeeid sets the "deletedbyEmployeeid" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableDeletedbyEmployeeid(i *int64) *EligibilityCadrePayMatrixCreate {
	if i != nil {
		ecpmc.SetDeletedbyEmployeeid(*i)
	}
	return ecpmc
}

// SetDeletedbyDesignation sets the "deletedbyDesignation" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetDeletedbyDesignation(s string) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetDeletedbyDesignation(s)
	return ecpmc
}

// SetNillableDeletedbyDesignation sets the "deletedbyDesignation" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableDeletedbyDesignation(s *string) *EligibilityCadrePayMatrixCreate {
	if s != nil {
		ecpmc.SetDeletedbyDesignation(*s)
	}
	return ecpmc
}

// SetDeletedDate sets the "deletedDate" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetDeletedDate(t time.Time) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetDeletedDate(t)
	return ecpmc
}

// SetNillableDeletedDate sets the "deletedDate" field if the given value is not nil.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetNillableDeletedDate(t *time.Time) *EligibilityCadrePayMatrixCreate {
	if t != nil {
		ecpmc.SetDeletedDate(*t)
	}
	return ecpmc
}

// SetID sets the "id" field.
func (ecpmc *EligibilityCadrePayMatrixCreate) SetID(i int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.SetID(i)
	return ecpmc
}

// AddLogDatumIDs adds the "LogData" edge to the Logs entity by IDs.
func (ecpmc *EligibilityCadrePayMatrixCreate) AddLogDatumIDs(ids ...int64) *EligibilityCadrePayMatrixCreate {
	ecpmc.mutation.AddLogDatumIDs(ids...)
	return ecpmc
}

// AddLogData adds the "LogData" edges to the Logs entity.
func (ecpmc *EligibilityCadrePayMatrixCreate) AddLogData(l ...*Logs) *EligibilityCadrePayMatrixCreate {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ecpmc.AddLogDatumIDs(ids...)
}

// Mutation returns the EligibilityCadrePayMatrixMutation object of the builder.
func (ecpmc *EligibilityCadrePayMatrixCreate) Mutation() *EligibilityCadrePayMatrixMutation {
	return ecpmc.mutation
}

// Save creates the EligibilityCadrePayMatrix in the database.
func (ecpmc *EligibilityCadrePayMatrixCreate) Save(ctx context.Context) (*EligibilityCadrePayMatrix, error) {
	ecpmc.defaults()
	return withHooks(ctx, ecpmc.sqlSave, ecpmc.mutation, ecpmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ecpmc *EligibilityCadrePayMatrixCreate) SaveX(ctx context.Context) *EligibilityCadrePayMatrix {
	v, err := ecpmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecpmc *EligibilityCadrePayMatrixCreate) Exec(ctx context.Context) error {
	_, err := ecpmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecpmc *EligibilityCadrePayMatrixCreate) ExecX(ctx context.Context) {
	if err := ecpmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecpmc *EligibilityCadrePayMatrixCreate) defaults() {
	if _, ok := ecpmc.mutation.CreatedDate(); !ok {
		v := eligibilitycadrepaymatrix.DefaultCreatedDate()
		ecpmc.mutation.SetCreatedDate(v)
	}
	if _, ok := ecpmc.mutation.VerifiedDate(); !ok {
		v := eligibilitycadrepaymatrix.DefaultVerifiedDate()
		ecpmc.mutation.SetVerifiedDate(v)
	}
	if _, ok := ecpmc.mutation.Statuss(); !ok {
		v := eligibilitycadrepaymatrix.DefaultStatuss
		ecpmc.mutation.SetStatuss(v)
	}
	if _, ok := ecpmc.mutation.DeletedDate(); !ok {
		v := eligibilitycadrepaymatrix.DefaultDeletedDate()
		ecpmc.mutation.SetDeletedDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecpmc *EligibilityCadrePayMatrixCreate) check() error {
	return nil
}

func (ecpmc *EligibilityCadrePayMatrixCreate) sqlSave(ctx context.Context) (*EligibilityCadrePayMatrix, error) {
	if err := ecpmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ecpmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ecpmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ecpmc.mutation.id = &_node.ID
	ecpmc.mutation.done = true
	return _node, nil
}

func (ecpmc *EligibilityCadrePayMatrixCreate) createSpec() (*EligibilityCadrePayMatrix, *sqlgraph.CreateSpec) {
	var (
		_node = &EligibilityCadrePayMatrix{config: ecpmc.config}
		_spec = sqlgraph.NewCreateSpec(eligibilitycadrepaymatrix.Table, sqlgraph.NewFieldSpec(eligibilitycadrepaymatrix.FieldID, field.TypeInt64))
	)
	if id, ok := ecpmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ecpmc.mutation.CadreEligibleConfigurationCadreEligibleCode(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldCadreEligibleConfigurationCadreEligibleCode, field.TypeInt64, value)
		_node.CadreEligibleConfigurationCadreEligibleCode = value
	}
	if value, ok := ecpmc.mutation.PostId(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldPostId, field.TypeInt64, value)
		_node.PostId = value
	}
	if value, ok := ecpmc.mutation.PostCode(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldPostCode, field.TypeString, value)
		_node.PostCode = value
	}
	if value, ok := ecpmc.mutation.PostDescription(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldPostDescription, field.TypeString, value)
		_node.PostDescription = value
	}
	if value, ok := ecpmc.mutation.OrderNumber(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldOrderNumber, field.TypeString, value)
		_node.OrderNumber = value
	}
	if value, ok := ecpmc.mutation.CreatedById(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldCreatedById, field.TypeInt64, value)
		_node.CreatedById = value
	}
	if value, ok := ecpmc.mutation.CreatedByUserName(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldCreatedByUserName, field.TypeString, value)
		_node.CreatedByUserName = value
	}
	if value, ok := ecpmc.mutation.CreatedByEmpId(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldCreatedByEmpId, field.TypeInt64, value)
		_node.CreatedByEmpId = value
	}
	if value, ok := ecpmc.mutation.CreatedByDesignation(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldCreatedByDesignation, field.TypeString, value)
		_node.CreatedByDesignation = value
	}
	if value, ok := ecpmc.mutation.CreatedDate(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldCreatedDate, field.TypeTime, value)
		_node.CreatedDate = value
	}
	if value, ok := ecpmc.mutation.Verifiedbyid(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldVerifiedbyid, field.TypeInt64, value)
		_node.Verifiedbyid = value
	}
	if value, ok := ecpmc.mutation.Verifiedbyusername(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldVerifiedbyusername, field.TypeString, value)
		_node.Verifiedbyusername = value
	}
	if value, ok := ecpmc.mutation.VerifiedbyEmployeeid(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldVerifiedbyEmployeeid, field.TypeInt64, value)
		_node.VerifiedbyEmployeeid = value
	}
	if value, ok := ecpmc.mutation.VerifiedbyDesignation(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldVerifiedbyDesignation, field.TypeString, value)
		_node.VerifiedbyDesignation = value
	}
	if value, ok := ecpmc.mutation.VerifiedDate(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldVerifiedDate, field.TypeTime, value)
		_node.VerifiedDate = value
	}
	if value, ok := ecpmc.mutation.Statuss(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldStatuss, field.TypeString, value)
		_node.Statuss = value
	}
	if value, ok := ecpmc.mutation.Deletedbyid(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldDeletedbyid, field.TypeInt64, value)
		_node.Deletedbyid = value
	}
	if value, ok := ecpmc.mutation.Deletedbyusername(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldDeletedbyusername, field.TypeString, value)
		_node.Deletedbyusername = value
	}
	if value, ok := ecpmc.mutation.DeletedbyEmployeeid(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldDeletedbyEmployeeid, field.TypeInt64, value)
		_node.DeletedbyEmployeeid = value
	}
	if value, ok := ecpmc.mutation.DeletedbyDesignation(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldDeletedbyDesignation, field.TypeString, value)
		_node.DeletedbyDesignation = value
	}
	if value, ok := ecpmc.mutation.DeletedDate(); ok {
		_spec.SetField(eligibilitycadrepaymatrix.FieldDeletedDate, field.TypeTime, value)
		_node.DeletedDate = value
	}
	if nodes := ecpmc.mutation.LogDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   eligibilitycadrepaymatrix.LogDataTable,
			Columns: []string{eligibilitycadrepaymatrix.LogDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(logs.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EligibilityCadrePayMatrixCreateBulk is the builder for creating many EligibilityCadrePayMatrix entities in bulk.
type EligibilityCadrePayMatrixCreateBulk struct {
	config
	builders []*EligibilityCadrePayMatrixCreate
}

// Save creates the EligibilityCadrePayMatrix entities in the database.
func (ecpmcb *EligibilityCadrePayMatrixCreateBulk) Save(ctx context.Context) ([]*EligibilityCadrePayMatrix, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecpmcb.builders))
	nodes := make([]*EligibilityCadrePayMatrix, len(ecpmcb.builders))
	mutators := make([]Mutator, len(ecpmcb.builders))
	for i := range ecpmcb.builders {
		func(i int, root context.Context) {
			builder := ecpmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EligibilityCadrePayMatrixMutation)
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
					_, err = mutators[i+1].Mutate(root, ecpmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecpmcb.driver, spec); err != nil {
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
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, ecpmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecpmcb *EligibilityCadrePayMatrixCreateBulk) SaveX(ctx context.Context) []*EligibilityCadrePayMatrix {
	v, err := ecpmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecpmcb *EligibilityCadrePayMatrixCreateBulk) Exec(ctx context.Context) error {
	_, err := ecpmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecpmcb *EligibilityCadrePayMatrixCreateBulk) ExecX(ctx context.Context) {
	if err := ecpmcb.Exec(ctx); err != nil {
		panic(err)
	}
}