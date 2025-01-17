// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/employeedesignation"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeDesignationCreate is the builder for creating a EmployeeDesignation entity.
type EmployeeDesignationCreate struct {
	config
	mutation *EmployeeDesignationMutation
	hooks    []Hook
}

// SetPostID sets the "PostID" field.
func (edc *EmployeeDesignationCreate) SetPostID(i int32) *EmployeeDesignationCreate {
	edc.mutation.SetPostID(i)
	return edc
}

// SetNillablePostID sets the "PostID" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillablePostID(i *int32) *EmployeeDesignationCreate {
	if i != nil {
		edc.SetPostID(*i)
	}
	return edc
}

// SetPostCode sets the "PostCode" field.
func (edc *EmployeeDesignationCreate) SetPostCode(s string) *EmployeeDesignationCreate {
	edc.mutation.SetPostCode(s)
	return edc
}

// SetNillablePostCode sets the "PostCode" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillablePostCode(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetPostCode(*s)
	}
	return edc
}

// SetPostDescription sets the "PostDescription" field.
func (edc *EmployeeDesignationCreate) SetPostDescription(s string) *EmployeeDesignationCreate {
	edc.mutation.SetPostDescription(s)
	return edc
}

// SetNillablePostDescription sets the "PostDescription" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillablePostDescription(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetPostDescription(*s)
	}
	return edc
}

// SetDesignationCode sets the "DesignationCode" field.
func (edc *EmployeeDesignationCreate) SetDesignationCode(s string) *EmployeeDesignationCreate {
	edc.mutation.SetDesignationCode(s)
	return edc
}

// SetNillableDesignationCode sets the "DesignationCode" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableDesignationCode(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetDesignationCode(*s)
	}
	return edc
}

// SetDesignationDescription sets the "DesignationDescription" field.
func (edc *EmployeeDesignationCreate) SetDesignationDescription(s string) *EmployeeDesignationCreate {
	edc.mutation.SetDesignationDescription(s)
	return edc
}

// SetOrderNumber sets the "OrderNumber" field.
func (edc *EmployeeDesignationCreate) SetOrderNumber(s string) *EmployeeDesignationCreate {
	edc.mutation.SetOrderNumber(s)
	return edc
}

// SetNillableOrderNumber sets the "OrderNumber" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableOrderNumber(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetOrderNumber(*s)
	}
	return edc
}

// SetCreatedById sets the "CreatedById" field.
func (edc *EmployeeDesignationCreate) SetCreatedById(i int64) *EmployeeDesignationCreate {
	edc.mutation.SetCreatedById(i)
	return edc
}

// SetNillableCreatedById sets the "CreatedById" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableCreatedById(i *int64) *EmployeeDesignationCreate {
	if i != nil {
		edc.SetCreatedById(*i)
	}
	return edc
}

// SetCreatedByUserName sets the "CreatedByUserName" field.
func (edc *EmployeeDesignationCreate) SetCreatedByUserName(s string) *EmployeeDesignationCreate {
	edc.mutation.SetCreatedByUserName(s)
	return edc
}

// SetNillableCreatedByUserName sets the "CreatedByUserName" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableCreatedByUserName(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetCreatedByUserName(*s)
	}
	return edc
}

// SetCreatedByEmpId sets the "CreatedByEmpId" field.
func (edc *EmployeeDesignationCreate) SetCreatedByEmpId(i int64) *EmployeeDesignationCreate {
	edc.mutation.SetCreatedByEmpId(i)
	return edc
}

// SetNillableCreatedByEmpId sets the "CreatedByEmpId" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableCreatedByEmpId(i *int64) *EmployeeDesignationCreate {
	if i != nil {
		edc.SetCreatedByEmpId(*i)
	}
	return edc
}

// SetCreatedByDesignation sets the "CreatedByDesignation" field.
func (edc *EmployeeDesignationCreate) SetCreatedByDesignation(s string) *EmployeeDesignationCreate {
	edc.mutation.SetCreatedByDesignation(s)
	return edc
}

// SetNillableCreatedByDesignation sets the "CreatedByDesignation" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableCreatedByDesignation(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetCreatedByDesignation(*s)
	}
	return edc
}

// SetCreatedDate sets the "CreatedDate" field.
func (edc *EmployeeDesignationCreate) SetCreatedDate(t time.Time) *EmployeeDesignationCreate {
	edc.mutation.SetCreatedDate(t)
	return edc
}

// SetNillableCreatedDate sets the "CreatedDate" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableCreatedDate(t *time.Time) *EmployeeDesignationCreate {
	if t != nil {
		edc.SetCreatedDate(*t)
	}
	return edc
}

// SetVerifiedbyid sets the "verifiedbyid" field.
func (edc *EmployeeDesignationCreate) SetVerifiedbyid(i int64) *EmployeeDesignationCreate {
	edc.mutation.SetVerifiedbyid(i)
	return edc
}

// SetNillableVerifiedbyid sets the "verifiedbyid" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableVerifiedbyid(i *int64) *EmployeeDesignationCreate {
	if i != nil {
		edc.SetVerifiedbyid(*i)
	}
	return edc
}

// SetVerifiedbyusername sets the "verifiedbyusername" field.
func (edc *EmployeeDesignationCreate) SetVerifiedbyusername(s string) *EmployeeDesignationCreate {
	edc.mutation.SetVerifiedbyusername(s)
	return edc
}

// SetNillableVerifiedbyusername sets the "verifiedbyusername" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableVerifiedbyusername(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetVerifiedbyusername(*s)
	}
	return edc
}

// SetVerifiedbyEmployeeid sets the "verifiedbyEmployeeid" field.
func (edc *EmployeeDesignationCreate) SetVerifiedbyEmployeeid(i int64) *EmployeeDesignationCreate {
	edc.mutation.SetVerifiedbyEmployeeid(i)
	return edc
}

// SetNillableVerifiedbyEmployeeid sets the "verifiedbyEmployeeid" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableVerifiedbyEmployeeid(i *int64) *EmployeeDesignationCreate {
	if i != nil {
		edc.SetVerifiedbyEmployeeid(*i)
	}
	return edc
}

// SetVerifiedbyDesignation sets the "verifiedbyDesignation" field.
func (edc *EmployeeDesignationCreate) SetVerifiedbyDesignation(s string) *EmployeeDesignationCreate {
	edc.mutation.SetVerifiedbyDesignation(s)
	return edc
}

// SetNillableVerifiedbyDesignation sets the "verifiedbyDesignation" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableVerifiedbyDesignation(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetVerifiedbyDesignation(*s)
	}
	return edc
}

// SetVerifiedDate sets the "verifiedDate" field.
func (edc *EmployeeDesignationCreate) SetVerifiedDate(t time.Time) *EmployeeDesignationCreate {
	edc.mutation.SetVerifiedDate(t)
	return edc
}

// SetNillableVerifiedDate sets the "verifiedDate" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableVerifiedDate(t *time.Time) *EmployeeDesignationCreate {
	if t != nil {
		edc.SetVerifiedDate(*t)
	}
	return edc
}

// SetStatuss sets the "Statuss" field.
func (edc *EmployeeDesignationCreate) SetStatuss(s string) *EmployeeDesignationCreate {
	edc.mutation.SetStatuss(s)
	return edc
}

// SetNillableStatuss sets the "Statuss" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableStatuss(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetStatuss(*s)
	}
	return edc
}

// SetDeletedbyid sets the "deletedbyid" field.
func (edc *EmployeeDesignationCreate) SetDeletedbyid(i int64) *EmployeeDesignationCreate {
	edc.mutation.SetDeletedbyid(i)
	return edc
}

// SetNillableDeletedbyid sets the "deletedbyid" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableDeletedbyid(i *int64) *EmployeeDesignationCreate {
	if i != nil {
		edc.SetDeletedbyid(*i)
	}
	return edc
}

// SetDeletedbyusername sets the "deletedbyusername" field.
func (edc *EmployeeDesignationCreate) SetDeletedbyusername(s string) *EmployeeDesignationCreate {
	edc.mutation.SetDeletedbyusername(s)
	return edc
}

// SetNillableDeletedbyusername sets the "deletedbyusername" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableDeletedbyusername(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetDeletedbyusername(*s)
	}
	return edc
}

// SetDeletedbyEmployeeid sets the "deletedbyEmployeeid" field.
func (edc *EmployeeDesignationCreate) SetDeletedbyEmployeeid(i int64) *EmployeeDesignationCreate {
	edc.mutation.SetDeletedbyEmployeeid(i)
	return edc
}

// SetNillableDeletedbyEmployeeid sets the "deletedbyEmployeeid" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableDeletedbyEmployeeid(i *int64) *EmployeeDesignationCreate {
	if i != nil {
		edc.SetDeletedbyEmployeeid(*i)
	}
	return edc
}

// SetDeletedbyDesignation sets the "deletedbyDesignation" field.
func (edc *EmployeeDesignationCreate) SetDeletedbyDesignation(s string) *EmployeeDesignationCreate {
	edc.mutation.SetDeletedbyDesignation(s)
	return edc
}

// SetNillableDeletedbyDesignation sets the "deletedbyDesignation" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableDeletedbyDesignation(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetDeletedbyDesignation(*s)
	}
	return edc
}

// SetDeletedDate sets the "deletedDate" field.
func (edc *EmployeeDesignationCreate) SetDeletedDate(t time.Time) *EmployeeDesignationCreate {
	edc.mutation.SetDeletedDate(t)
	return edc
}

// SetNillableDeletedDate sets the "deletedDate" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableDeletedDate(t *time.Time) *EmployeeDesignationCreate {
	if t != nil {
		edc.SetDeletedDate(*t)
	}
	return edc
}

// SetPaperStatus sets the "PaperStatus" field.
func (edc *EmployeeDesignationCreate) SetPaperStatus(s string) *EmployeeDesignationCreate {
	edc.mutation.SetPaperStatus(s)
	return edc
}

// SetCalendarCode sets the "CalendarCode" field.
func (edc *EmployeeDesignationCreate) SetCalendarCode(i int32) *EmployeeDesignationCreate {
	edc.mutation.SetCalendarCode(i)
	return edc
}

// SetNillableCalendarCode sets the "CalendarCode" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableCalendarCode(i *int32) *EmployeeDesignationCreate {
	if i != nil {
		edc.SetCalendarCode(*i)
	}
	return edc
}

// SetExamCodePS sets the "ExamCodePS" field.
func (edc *EmployeeDesignationCreate) SetExamCodePS(i int32) *EmployeeDesignationCreate {
	edc.mutation.SetExamCodePS(i)
	return edc
}

// SetNillableExamCodePS sets the "ExamCodePS" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableExamCodePS(i *int32) *EmployeeDesignationCreate {
	if i != nil {
		edc.SetExamCodePS(*i)
	}
	return edc
}

// SetCreatedByEmployeeId sets the "CreatedByEmployeeId" field.
func (edc *EmployeeDesignationCreate) SetCreatedByEmployeeId(s string) *EmployeeDesignationCreate {
	edc.mutation.SetCreatedByEmployeeId(s)
	return edc
}

// SetNillableCreatedByEmployeeId sets the "CreatedByEmployeeId" field if the given value is not nil.
func (edc *EmployeeDesignationCreate) SetNillableCreatedByEmployeeId(s *string) *EmployeeDesignationCreate {
	if s != nil {
		edc.SetCreatedByEmployeeId(*s)
	}
	return edc
}

// SetID sets the "id" field.
func (edc *EmployeeDesignationCreate) SetID(i int32) *EmployeeDesignationCreate {
	edc.mutation.SetID(i)
	return edc
}

// Mutation returns the EmployeeDesignationMutation object of the builder.
func (edc *EmployeeDesignationCreate) Mutation() *EmployeeDesignationMutation {
	return edc.mutation
}

// Save creates the EmployeeDesignation in the database.
func (edc *EmployeeDesignationCreate) Save(ctx context.Context) (*EmployeeDesignation, error) {
	edc.defaults()
	return withHooks(ctx, edc.sqlSave, edc.mutation, edc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (edc *EmployeeDesignationCreate) SaveX(ctx context.Context) *EmployeeDesignation {
	v, err := edc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (edc *EmployeeDesignationCreate) Exec(ctx context.Context) error {
	_, err := edc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (edc *EmployeeDesignationCreate) ExecX(ctx context.Context) {
	if err := edc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (edc *EmployeeDesignationCreate) defaults() {
	if _, ok := edc.mutation.Statuss(); !ok {
		v := employeedesignation.DefaultStatuss
		edc.mutation.SetStatuss(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (edc *EmployeeDesignationCreate) check() error {
	if _, ok := edc.mutation.DesignationDescription(); !ok {
		return &ValidationError{Name: "DesignationDescription", err: errors.New(`ent: missing required field "EmployeeDesignation.DesignationDescription"`)}
	}
	if _, ok := edc.mutation.PaperStatus(); !ok {
		return &ValidationError{Name: "PaperStatus", err: errors.New(`ent: missing required field "EmployeeDesignation.PaperStatus"`)}
	}
	if v, ok := edc.mutation.PaperStatus(); ok {
		if err := employeedesignation.PaperStatusValidator(v); err != nil {
			return &ValidationError{Name: "PaperStatus", err: fmt.Errorf(`ent: validator failed for field "EmployeeDesignation.PaperStatus": %w`, err)}
		}
	}
	return nil
}

func (edc *EmployeeDesignationCreate) sqlSave(ctx context.Context) (*EmployeeDesignation, error) {
	if err := edc.check(); err != nil {
		return nil, err
	}
	_node, _spec := edc.createSpec()
	if err := sqlgraph.CreateNode(ctx, edc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	edc.mutation.id = &_node.ID
	edc.mutation.done = true
	return _node, nil
}

func (edc *EmployeeDesignationCreate) createSpec() (*EmployeeDesignation, *sqlgraph.CreateSpec) {
	var (
		_node = &EmployeeDesignation{config: edc.config}
		_spec = sqlgraph.NewCreateSpec(employeedesignation.Table, sqlgraph.NewFieldSpec(employeedesignation.FieldID, field.TypeInt32))
	)
	if id, ok := edc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := edc.mutation.PostID(); ok {
		_spec.SetField(employeedesignation.FieldPostID, field.TypeInt32, value)
		_node.PostID = value
	}
	if value, ok := edc.mutation.PostCode(); ok {
		_spec.SetField(employeedesignation.FieldPostCode, field.TypeString, value)
		_node.PostCode = value
	}
	if value, ok := edc.mutation.PostDescription(); ok {
		_spec.SetField(employeedesignation.FieldPostDescription, field.TypeString, value)
		_node.PostDescription = value
	}
	if value, ok := edc.mutation.DesignationCode(); ok {
		_spec.SetField(employeedesignation.FieldDesignationCode, field.TypeString, value)
		_node.DesignationCode = value
	}
	if value, ok := edc.mutation.DesignationDescription(); ok {
		_spec.SetField(employeedesignation.FieldDesignationDescription, field.TypeString, value)
		_node.DesignationDescription = value
	}
	if value, ok := edc.mutation.OrderNumber(); ok {
		_spec.SetField(employeedesignation.FieldOrderNumber, field.TypeString, value)
		_node.OrderNumber = value
	}
	if value, ok := edc.mutation.CreatedById(); ok {
		_spec.SetField(employeedesignation.FieldCreatedById, field.TypeInt64, value)
		_node.CreatedById = value
	}
	if value, ok := edc.mutation.CreatedByUserName(); ok {
		_spec.SetField(employeedesignation.FieldCreatedByUserName, field.TypeString, value)
		_node.CreatedByUserName = value
	}
	if value, ok := edc.mutation.CreatedByEmpId(); ok {
		_spec.SetField(employeedesignation.FieldCreatedByEmpId, field.TypeInt64, value)
		_node.CreatedByEmpId = value
	}
	if value, ok := edc.mutation.CreatedByDesignation(); ok {
		_spec.SetField(employeedesignation.FieldCreatedByDesignation, field.TypeString, value)
		_node.CreatedByDesignation = value
	}
	if value, ok := edc.mutation.CreatedDate(); ok {
		_spec.SetField(employeedesignation.FieldCreatedDate, field.TypeTime, value)
		_node.CreatedDate = value
	}
	if value, ok := edc.mutation.Verifiedbyid(); ok {
		_spec.SetField(employeedesignation.FieldVerifiedbyid, field.TypeInt64, value)
		_node.Verifiedbyid = value
	}
	if value, ok := edc.mutation.Verifiedbyusername(); ok {
		_spec.SetField(employeedesignation.FieldVerifiedbyusername, field.TypeString, value)
		_node.Verifiedbyusername = value
	}
	if value, ok := edc.mutation.VerifiedbyEmployeeid(); ok {
		_spec.SetField(employeedesignation.FieldVerifiedbyEmployeeid, field.TypeInt64, value)
		_node.VerifiedbyEmployeeid = value
	}
	if value, ok := edc.mutation.VerifiedbyDesignation(); ok {
		_spec.SetField(employeedesignation.FieldVerifiedbyDesignation, field.TypeString, value)
		_node.VerifiedbyDesignation = value
	}
	if value, ok := edc.mutation.VerifiedDate(); ok {
		_spec.SetField(employeedesignation.FieldVerifiedDate, field.TypeTime, value)
		_node.VerifiedDate = value
	}
	if value, ok := edc.mutation.Statuss(); ok {
		_spec.SetField(employeedesignation.FieldStatuss, field.TypeString, value)
		_node.Statuss = value
	}
	if value, ok := edc.mutation.Deletedbyid(); ok {
		_spec.SetField(employeedesignation.FieldDeletedbyid, field.TypeInt64, value)
		_node.Deletedbyid = value
	}
	if value, ok := edc.mutation.Deletedbyusername(); ok {
		_spec.SetField(employeedesignation.FieldDeletedbyusername, field.TypeString, value)
		_node.Deletedbyusername = value
	}
	if value, ok := edc.mutation.DeletedbyEmployeeid(); ok {
		_spec.SetField(employeedesignation.FieldDeletedbyEmployeeid, field.TypeInt64, value)
		_node.DeletedbyEmployeeid = value
	}
	if value, ok := edc.mutation.DeletedbyDesignation(); ok {
		_spec.SetField(employeedesignation.FieldDeletedbyDesignation, field.TypeString, value)
		_node.DeletedbyDesignation = value
	}
	if value, ok := edc.mutation.DeletedDate(); ok {
		_spec.SetField(employeedesignation.FieldDeletedDate, field.TypeTime, value)
		_node.DeletedDate = value
	}
	if value, ok := edc.mutation.PaperStatus(); ok {
		_spec.SetField(employeedesignation.FieldPaperStatus, field.TypeString, value)
		_node.PaperStatus = value
	}
	if value, ok := edc.mutation.CalendarCode(); ok {
		_spec.SetField(employeedesignation.FieldCalendarCode, field.TypeInt32, value)
		_node.CalendarCode = value
	}
	if value, ok := edc.mutation.ExamCodePS(); ok {
		_spec.SetField(employeedesignation.FieldExamCodePS, field.TypeInt32, value)
		_node.ExamCodePS = value
	}
	if value, ok := edc.mutation.CreatedByEmployeeId(); ok {
		_spec.SetField(employeedesignation.FieldCreatedByEmployeeId, field.TypeString, value)
		_node.CreatedByEmployeeId = value
	}
	return _node, _spec
}

// EmployeeDesignationCreateBulk is the builder for creating many EmployeeDesignation entities in bulk.
type EmployeeDesignationCreateBulk struct {
	config
	builders []*EmployeeDesignationCreate
}

// Save creates the EmployeeDesignation entities in the database.
func (edcb *EmployeeDesignationCreateBulk) Save(ctx context.Context) ([]*EmployeeDesignation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(edcb.builders))
	nodes := make([]*EmployeeDesignation, len(edcb.builders))
	mutators := make([]Mutator, len(edcb.builders))
	for i := range edcb.builders {
		func(i int, root context.Context) {
			builder := edcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmployeeDesignationMutation)
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
					_, err = mutators[i+1].Mutate(root, edcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, edcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, edcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (edcb *EmployeeDesignationCreateBulk) SaveX(ctx context.Context) []*EmployeeDesignation {
	v, err := edcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (edcb *EmployeeDesignationCreateBulk) Exec(ctx context.Context) error {
	_, err := edcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (edcb *EmployeeDesignationCreateBulk) ExecX(ctx context.Context) {
	if err := edcb.Exec(ctx); err != nil {
		panic(err)
	}
}
