// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/cadre_choice_pm"
	"recruit/ent/exam_applications_gdspm"
	"recruit/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CadreChoicePMUpdate is the builder for updating Cadre_Choice_PM entities.
type CadreChoicePMUpdate struct {
	config
	hooks    []Hook
	mutation *CadreChoicePMMutation
}

// Where appends a list predicates to the CadreChoicePMUpdate builder.
func (ccpu *CadreChoicePMUpdate) Where(ps ...predicate.Cadre_Choice_PM) *CadreChoicePMUpdate {
	ccpu.mutation.Where(ps...)
	return ccpu
}

// SetApplicationID sets the "ApplicationID" field.
func (ccpu *CadreChoicePMUpdate) SetApplicationID(i int64) *CadreChoicePMUpdate {
	ccpu.mutation.SetApplicationID(i)
	return ccpu
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (ccpu *CadreChoicePMUpdate) SetNillableApplicationID(i *int64) *CadreChoicePMUpdate {
	if i != nil {
		ccpu.SetApplicationID(*i)
	}
	return ccpu
}

// ClearApplicationID clears the value of the "ApplicationID" field.
func (ccpu *CadreChoicePMUpdate) ClearApplicationID() *CadreChoicePMUpdate {
	ccpu.mutation.ClearApplicationID()
	return ccpu
}

// SetGroup sets the "Group" field.
func (ccpu *CadreChoicePMUpdate) SetGroup(s string) *CadreChoicePMUpdate {
	ccpu.mutation.SetGroup(s)
	return ccpu
}

// SetCadrePrefNo sets the "CadrePrefNo" field.
func (ccpu *CadreChoicePMUpdate) SetCadrePrefNo(i int64) *CadreChoicePMUpdate {
	ccpu.mutation.ResetCadrePrefNo()
	ccpu.mutation.SetCadrePrefNo(i)
	return ccpu
}

// AddCadrePrefNo adds i to the "CadrePrefNo" field.
func (ccpu *CadreChoicePMUpdate) AddCadrePrefNo(i int64) *CadreChoicePMUpdate {
	ccpu.mutation.AddCadrePrefNo(i)
	return ccpu
}

// SetCadre sets the "Cadre" field.
func (ccpu *CadreChoicePMUpdate) SetCadre(s string) *CadreChoicePMUpdate {
	ccpu.mutation.SetCadre(s)
	return ccpu
}

// SetPostPrefNo sets the "PostPrefNo" field.
func (ccpu *CadreChoicePMUpdate) SetPostPrefNo(i int64) *CadreChoicePMUpdate {
	ccpu.mutation.ResetPostPrefNo()
	ccpu.mutation.SetPostPrefNo(i)
	return ccpu
}

// AddPostPrefNo adds i to the "PostPrefNo" field.
func (ccpu *CadreChoicePMUpdate) AddPostPrefNo(i int64) *CadreChoicePMUpdate {
	ccpu.mutation.AddPostPrefNo(i)
	return ccpu
}

// SetPostPrefValue sets the "PostPrefValue" field.
func (ccpu *CadreChoicePMUpdate) SetPostPrefValue(s string) *CadreChoicePMUpdate {
	ccpu.mutation.SetPostPrefValue(s)
	return ccpu
}

// SetEmployeeID sets the "EmployeeID" field.
func (ccpu *CadreChoicePMUpdate) SetEmployeeID(i int64) *CadreChoicePMUpdate {
	ccpu.mutation.ResetEmployeeID()
	ccpu.mutation.SetEmployeeID(i)
	return ccpu
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (ccpu *CadreChoicePMUpdate) SetNillableEmployeeID(i *int64) *CadreChoicePMUpdate {
	if i != nil {
		ccpu.SetEmployeeID(*i)
	}
	return ccpu
}

// AddEmployeeID adds i to the "EmployeeID" field.
func (ccpu *CadreChoicePMUpdate) AddEmployeeID(i int64) *CadreChoicePMUpdate {
	ccpu.mutation.AddEmployeeID(i)
	return ccpu
}

// ClearEmployeeID clears the value of the "EmployeeID" field.
func (ccpu *CadreChoicePMUpdate) ClearEmployeeID() *CadreChoicePMUpdate {
	ccpu.mutation.ClearEmployeeID()
	return ccpu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (ccpu *CadreChoicePMUpdate) SetUpdatedAt(t time.Time) *CadreChoicePMUpdate {
	ccpu.mutation.SetUpdatedAt(t)
	return ccpu
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (ccpu *CadreChoicePMUpdate) SetNillableUpdatedAt(t *time.Time) *CadreChoicePMUpdate {
	if t != nil {
		ccpu.SetUpdatedAt(*t)
	}
	return ccpu
}

// ClearUpdatedAt clears the value of the "UpdatedAt" field.
func (ccpu *CadreChoicePMUpdate) ClearUpdatedAt() *CadreChoicePMUpdate {
	ccpu.mutation.ClearUpdatedAt()
	return ccpu
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (ccpu *CadreChoicePMUpdate) SetUpdatedBy(s string) *CadreChoicePMUpdate {
	ccpu.mutation.SetUpdatedBy(s)
	return ccpu
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (ccpu *CadreChoicePMUpdate) SetNillableUpdatedBy(s *string) *CadreChoicePMUpdate {
	if s != nil {
		ccpu.SetUpdatedBy(*s)
	}
	return ccpu
}

// ClearUpdatedBy clears the value of the "UpdatedBy" field.
func (ccpu *CadreChoicePMUpdate) ClearUpdatedBy() *CadreChoicePMUpdate {
	ccpu.mutation.ClearUpdatedBy()
	return ccpu
}

// SetApplnGDSPMRefID sets the "ApplnGDSPM_Ref" edge to the Exam_Applications_GDSPM entity by ID.
func (ccpu *CadreChoicePMUpdate) SetApplnGDSPMRefID(id int64) *CadreChoicePMUpdate {
	ccpu.mutation.SetApplnGDSPMRefID(id)
	return ccpu
}

// SetNillableApplnGDSPMRefID sets the "ApplnGDSPM_Ref" edge to the Exam_Applications_GDSPM entity by ID if the given value is not nil.
func (ccpu *CadreChoicePMUpdate) SetNillableApplnGDSPMRefID(id *int64) *CadreChoicePMUpdate {
	if id != nil {
		ccpu = ccpu.SetApplnGDSPMRefID(*id)
	}
	return ccpu
}

// SetApplnGDSPMRef sets the "ApplnGDSPM_Ref" edge to the Exam_Applications_GDSPM entity.
func (ccpu *CadreChoicePMUpdate) SetApplnGDSPMRef(e *Exam_Applications_GDSPM) *CadreChoicePMUpdate {
	return ccpu.SetApplnGDSPMRefID(e.ID)
}

// Mutation returns the CadreChoicePMMutation object of the builder.
func (ccpu *CadreChoicePMUpdate) Mutation() *CadreChoicePMMutation {
	return ccpu.mutation
}

// ClearApplnGDSPMRef clears the "ApplnGDSPM_Ref" edge to the Exam_Applications_GDSPM entity.
func (ccpu *CadreChoicePMUpdate) ClearApplnGDSPMRef() *CadreChoicePMUpdate {
	ccpu.mutation.ClearApplnGDSPMRef()
	return ccpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccpu *CadreChoicePMUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ccpu.sqlSave, ccpu.mutation, ccpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccpu *CadreChoicePMUpdate) SaveX(ctx context.Context) int {
	affected, err := ccpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccpu *CadreChoicePMUpdate) Exec(ctx context.Context) error {
	_, err := ccpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccpu *CadreChoicePMUpdate) ExecX(ctx context.Context) {
	if err := ccpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ccpu *CadreChoicePMUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cadre_choice_pm.Table, cadre_choice_pm.Columns, sqlgraph.NewFieldSpec(cadre_choice_pm.FieldID, field.TypeInt32))
	if ps := ccpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccpu.mutation.Group(); ok {
		_spec.SetField(cadre_choice_pm.FieldGroup, field.TypeString, value)
	}
	if value, ok := ccpu.mutation.CadrePrefNo(); ok {
		_spec.SetField(cadre_choice_pm.FieldCadrePrefNo, field.TypeInt64, value)
	}
	if value, ok := ccpu.mutation.AddedCadrePrefNo(); ok {
		_spec.AddField(cadre_choice_pm.FieldCadrePrefNo, field.TypeInt64, value)
	}
	if value, ok := ccpu.mutation.Cadre(); ok {
		_spec.SetField(cadre_choice_pm.FieldCadre, field.TypeString, value)
	}
	if value, ok := ccpu.mutation.PostPrefNo(); ok {
		_spec.SetField(cadre_choice_pm.FieldPostPrefNo, field.TypeInt64, value)
	}
	if value, ok := ccpu.mutation.AddedPostPrefNo(); ok {
		_spec.AddField(cadre_choice_pm.FieldPostPrefNo, field.TypeInt64, value)
	}
	if value, ok := ccpu.mutation.PostPrefValue(); ok {
		_spec.SetField(cadre_choice_pm.FieldPostPrefValue, field.TypeString, value)
	}
	if value, ok := ccpu.mutation.EmployeeID(); ok {
		_spec.SetField(cadre_choice_pm.FieldEmployeeID, field.TypeInt64, value)
	}
	if value, ok := ccpu.mutation.AddedEmployeeID(); ok {
		_spec.AddField(cadre_choice_pm.FieldEmployeeID, field.TypeInt64, value)
	}
	if ccpu.mutation.EmployeeIDCleared() {
		_spec.ClearField(cadre_choice_pm.FieldEmployeeID, field.TypeInt64)
	}
	if value, ok := ccpu.mutation.UpdatedAt(); ok {
		_spec.SetField(cadre_choice_pm.FieldUpdatedAt, field.TypeTime, value)
	}
	if ccpu.mutation.UpdatedAtCleared() {
		_spec.ClearField(cadre_choice_pm.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ccpu.mutation.UpdatedBy(); ok {
		_spec.SetField(cadre_choice_pm.FieldUpdatedBy, field.TypeString, value)
	}
	if ccpu.mutation.UpdatedByCleared() {
		_spec.ClearField(cadre_choice_pm.FieldUpdatedBy, field.TypeString)
	}
	if ccpu.mutation.ApplnGDSPMRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cadre_choice_pm.ApplnGDSPMRefTable,
			Columns: []string{cadre_choice_pm.ApplnGDSPMRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_gdspm.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccpu.mutation.ApplnGDSPMRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cadre_choice_pm.ApplnGDSPMRefTable,
			Columns: []string{cadre_choice_pm.ApplnGDSPMRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_gdspm.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cadre_choice_pm.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ccpu.mutation.done = true
	return n, nil
}

// CadreChoicePMUpdateOne is the builder for updating a single Cadre_Choice_PM entity.
type CadreChoicePMUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CadreChoicePMMutation
}

// SetApplicationID sets the "ApplicationID" field.
func (ccpuo *CadreChoicePMUpdateOne) SetApplicationID(i int64) *CadreChoicePMUpdateOne {
	ccpuo.mutation.SetApplicationID(i)
	return ccpuo
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (ccpuo *CadreChoicePMUpdateOne) SetNillableApplicationID(i *int64) *CadreChoicePMUpdateOne {
	if i != nil {
		ccpuo.SetApplicationID(*i)
	}
	return ccpuo
}

// ClearApplicationID clears the value of the "ApplicationID" field.
func (ccpuo *CadreChoicePMUpdateOne) ClearApplicationID() *CadreChoicePMUpdateOne {
	ccpuo.mutation.ClearApplicationID()
	return ccpuo
}

// SetGroup sets the "Group" field.
func (ccpuo *CadreChoicePMUpdateOne) SetGroup(s string) *CadreChoicePMUpdateOne {
	ccpuo.mutation.SetGroup(s)
	return ccpuo
}

// SetCadrePrefNo sets the "CadrePrefNo" field.
func (ccpuo *CadreChoicePMUpdateOne) SetCadrePrefNo(i int64) *CadreChoicePMUpdateOne {
	ccpuo.mutation.ResetCadrePrefNo()
	ccpuo.mutation.SetCadrePrefNo(i)
	return ccpuo
}

// AddCadrePrefNo adds i to the "CadrePrefNo" field.
func (ccpuo *CadreChoicePMUpdateOne) AddCadrePrefNo(i int64) *CadreChoicePMUpdateOne {
	ccpuo.mutation.AddCadrePrefNo(i)
	return ccpuo
}

// SetCadre sets the "Cadre" field.
func (ccpuo *CadreChoicePMUpdateOne) SetCadre(s string) *CadreChoicePMUpdateOne {
	ccpuo.mutation.SetCadre(s)
	return ccpuo
}

// SetPostPrefNo sets the "PostPrefNo" field.
func (ccpuo *CadreChoicePMUpdateOne) SetPostPrefNo(i int64) *CadreChoicePMUpdateOne {
	ccpuo.mutation.ResetPostPrefNo()
	ccpuo.mutation.SetPostPrefNo(i)
	return ccpuo
}

// AddPostPrefNo adds i to the "PostPrefNo" field.
func (ccpuo *CadreChoicePMUpdateOne) AddPostPrefNo(i int64) *CadreChoicePMUpdateOne {
	ccpuo.mutation.AddPostPrefNo(i)
	return ccpuo
}

// SetPostPrefValue sets the "PostPrefValue" field.
func (ccpuo *CadreChoicePMUpdateOne) SetPostPrefValue(s string) *CadreChoicePMUpdateOne {
	ccpuo.mutation.SetPostPrefValue(s)
	return ccpuo
}

// SetEmployeeID sets the "EmployeeID" field.
func (ccpuo *CadreChoicePMUpdateOne) SetEmployeeID(i int64) *CadreChoicePMUpdateOne {
	ccpuo.mutation.ResetEmployeeID()
	ccpuo.mutation.SetEmployeeID(i)
	return ccpuo
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (ccpuo *CadreChoicePMUpdateOne) SetNillableEmployeeID(i *int64) *CadreChoicePMUpdateOne {
	if i != nil {
		ccpuo.SetEmployeeID(*i)
	}
	return ccpuo
}

// AddEmployeeID adds i to the "EmployeeID" field.
func (ccpuo *CadreChoicePMUpdateOne) AddEmployeeID(i int64) *CadreChoicePMUpdateOne {
	ccpuo.mutation.AddEmployeeID(i)
	return ccpuo
}

// ClearEmployeeID clears the value of the "EmployeeID" field.
func (ccpuo *CadreChoicePMUpdateOne) ClearEmployeeID() *CadreChoicePMUpdateOne {
	ccpuo.mutation.ClearEmployeeID()
	return ccpuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (ccpuo *CadreChoicePMUpdateOne) SetUpdatedAt(t time.Time) *CadreChoicePMUpdateOne {
	ccpuo.mutation.SetUpdatedAt(t)
	return ccpuo
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (ccpuo *CadreChoicePMUpdateOne) SetNillableUpdatedAt(t *time.Time) *CadreChoicePMUpdateOne {
	if t != nil {
		ccpuo.SetUpdatedAt(*t)
	}
	return ccpuo
}

// ClearUpdatedAt clears the value of the "UpdatedAt" field.
func (ccpuo *CadreChoicePMUpdateOne) ClearUpdatedAt() *CadreChoicePMUpdateOne {
	ccpuo.mutation.ClearUpdatedAt()
	return ccpuo
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (ccpuo *CadreChoicePMUpdateOne) SetUpdatedBy(s string) *CadreChoicePMUpdateOne {
	ccpuo.mutation.SetUpdatedBy(s)
	return ccpuo
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (ccpuo *CadreChoicePMUpdateOne) SetNillableUpdatedBy(s *string) *CadreChoicePMUpdateOne {
	if s != nil {
		ccpuo.SetUpdatedBy(*s)
	}
	return ccpuo
}

// ClearUpdatedBy clears the value of the "UpdatedBy" field.
func (ccpuo *CadreChoicePMUpdateOne) ClearUpdatedBy() *CadreChoicePMUpdateOne {
	ccpuo.mutation.ClearUpdatedBy()
	return ccpuo
}

// SetApplnGDSPMRefID sets the "ApplnGDSPM_Ref" edge to the Exam_Applications_GDSPM entity by ID.
func (ccpuo *CadreChoicePMUpdateOne) SetApplnGDSPMRefID(id int64) *CadreChoicePMUpdateOne {
	ccpuo.mutation.SetApplnGDSPMRefID(id)
	return ccpuo
}

// SetNillableApplnGDSPMRefID sets the "ApplnGDSPM_Ref" edge to the Exam_Applications_GDSPM entity by ID if the given value is not nil.
func (ccpuo *CadreChoicePMUpdateOne) SetNillableApplnGDSPMRefID(id *int64) *CadreChoicePMUpdateOne {
	if id != nil {
		ccpuo = ccpuo.SetApplnGDSPMRefID(*id)
	}
	return ccpuo
}

// SetApplnGDSPMRef sets the "ApplnGDSPM_Ref" edge to the Exam_Applications_GDSPM entity.
func (ccpuo *CadreChoicePMUpdateOne) SetApplnGDSPMRef(e *Exam_Applications_GDSPM) *CadreChoicePMUpdateOne {
	return ccpuo.SetApplnGDSPMRefID(e.ID)
}

// Mutation returns the CadreChoicePMMutation object of the builder.
func (ccpuo *CadreChoicePMUpdateOne) Mutation() *CadreChoicePMMutation {
	return ccpuo.mutation
}

// ClearApplnGDSPMRef clears the "ApplnGDSPM_Ref" edge to the Exam_Applications_GDSPM entity.
func (ccpuo *CadreChoicePMUpdateOne) ClearApplnGDSPMRef() *CadreChoicePMUpdateOne {
	ccpuo.mutation.ClearApplnGDSPMRef()
	return ccpuo
}

// Where appends a list predicates to the CadreChoicePMUpdate builder.
func (ccpuo *CadreChoicePMUpdateOne) Where(ps ...predicate.Cadre_Choice_PM) *CadreChoicePMUpdateOne {
	ccpuo.mutation.Where(ps...)
	return ccpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccpuo *CadreChoicePMUpdateOne) Select(field string, fields ...string) *CadreChoicePMUpdateOne {
	ccpuo.fields = append([]string{field}, fields...)
	return ccpuo
}

// Save executes the query and returns the updated Cadre_Choice_PM entity.
func (ccpuo *CadreChoicePMUpdateOne) Save(ctx context.Context) (*Cadre_Choice_PM, error) {
	return withHooks(ctx, ccpuo.sqlSave, ccpuo.mutation, ccpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccpuo *CadreChoicePMUpdateOne) SaveX(ctx context.Context) *Cadre_Choice_PM {
	node, err := ccpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccpuo *CadreChoicePMUpdateOne) Exec(ctx context.Context) error {
	_, err := ccpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccpuo *CadreChoicePMUpdateOne) ExecX(ctx context.Context) {
	if err := ccpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ccpuo *CadreChoicePMUpdateOne) sqlSave(ctx context.Context) (_node *Cadre_Choice_PM, err error) {
	_spec := sqlgraph.NewUpdateSpec(cadre_choice_pm.Table, cadre_choice_pm.Columns, sqlgraph.NewFieldSpec(cadre_choice_pm.FieldID, field.TypeInt32))
	id, ok := ccpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cadre_Choice_PM.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cadre_choice_pm.FieldID)
		for _, f := range fields {
			if !cadre_choice_pm.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cadre_choice_pm.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccpuo.mutation.Group(); ok {
		_spec.SetField(cadre_choice_pm.FieldGroup, field.TypeString, value)
	}
	if value, ok := ccpuo.mutation.CadrePrefNo(); ok {
		_spec.SetField(cadre_choice_pm.FieldCadrePrefNo, field.TypeInt64, value)
	}
	if value, ok := ccpuo.mutation.AddedCadrePrefNo(); ok {
		_spec.AddField(cadre_choice_pm.FieldCadrePrefNo, field.TypeInt64, value)
	}
	if value, ok := ccpuo.mutation.Cadre(); ok {
		_spec.SetField(cadre_choice_pm.FieldCadre, field.TypeString, value)
	}
	if value, ok := ccpuo.mutation.PostPrefNo(); ok {
		_spec.SetField(cadre_choice_pm.FieldPostPrefNo, field.TypeInt64, value)
	}
	if value, ok := ccpuo.mutation.AddedPostPrefNo(); ok {
		_spec.AddField(cadre_choice_pm.FieldPostPrefNo, field.TypeInt64, value)
	}
	if value, ok := ccpuo.mutation.PostPrefValue(); ok {
		_spec.SetField(cadre_choice_pm.FieldPostPrefValue, field.TypeString, value)
	}
	if value, ok := ccpuo.mutation.EmployeeID(); ok {
		_spec.SetField(cadre_choice_pm.FieldEmployeeID, field.TypeInt64, value)
	}
	if value, ok := ccpuo.mutation.AddedEmployeeID(); ok {
		_spec.AddField(cadre_choice_pm.FieldEmployeeID, field.TypeInt64, value)
	}
	if ccpuo.mutation.EmployeeIDCleared() {
		_spec.ClearField(cadre_choice_pm.FieldEmployeeID, field.TypeInt64)
	}
	if value, ok := ccpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cadre_choice_pm.FieldUpdatedAt, field.TypeTime, value)
	}
	if ccpuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(cadre_choice_pm.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ccpuo.mutation.UpdatedBy(); ok {
		_spec.SetField(cadre_choice_pm.FieldUpdatedBy, field.TypeString, value)
	}
	if ccpuo.mutation.UpdatedByCleared() {
		_spec.ClearField(cadre_choice_pm.FieldUpdatedBy, field.TypeString)
	}
	if ccpuo.mutation.ApplnGDSPMRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cadre_choice_pm.ApplnGDSPMRefTable,
			Columns: []string{cadre_choice_pm.ApplnGDSPMRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_gdspm.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ccpuo.mutation.ApplnGDSPMRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cadre_choice_pm.ApplnGDSPMRefTable,
			Columns: []string{cadre_choice_pm.ApplnGDSPMRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_gdspm.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cadre_Choice_PM{config: ccpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cadre_choice_pm.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ccpuo.mutation.done = true
	return _node, nil
}