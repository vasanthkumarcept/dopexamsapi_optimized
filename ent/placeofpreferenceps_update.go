// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/placeofpreferenceps"
	"recruit/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaceOfPreferencePSUpdate is the builder for updating PlaceOfPreferencePS entities.
type PlaceOfPreferencePSUpdate struct {
	config
	hooks    []Hook
	mutation *PlaceOfPreferencePSMutation
}

// Where appends a list predicates to the PlaceOfPreferencePSUpdate builder.
func (poppu *PlaceOfPreferencePSUpdate) Where(ps ...predicate.PlaceOfPreferencePS) *PlaceOfPreferencePSUpdate {
	poppu.mutation.Where(ps...)
	return poppu
}

// SetApplicationID sets the "ApplicationID" field.
func (poppu *PlaceOfPreferencePSUpdate) SetApplicationID(i int64) *PlaceOfPreferencePSUpdate {
	poppu.mutation.SetApplicationID(i)
	return poppu
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (poppu *PlaceOfPreferencePSUpdate) SetNillableApplicationID(i *int64) *PlaceOfPreferencePSUpdate {
	if i != nil {
		poppu.SetApplicationID(*i)
	}
	return poppu
}

// ClearApplicationID clears the value of the "ApplicationID" field.
func (poppu *PlaceOfPreferencePSUpdate) ClearApplicationID() *PlaceOfPreferencePSUpdate {
	poppu.mutation.ClearApplicationID()
	return poppu
}

// SetPlacePrefNo sets the "PlacePrefNo" field.
func (poppu *PlaceOfPreferencePSUpdate) SetPlacePrefNo(i int32) *PlaceOfPreferencePSUpdate {
	poppu.mutation.ResetPlacePrefNo()
	poppu.mutation.SetPlacePrefNo(i)
	return poppu
}

// SetNillablePlacePrefNo sets the "PlacePrefNo" field if the given value is not nil.
func (poppu *PlaceOfPreferencePSUpdate) SetNillablePlacePrefNo(i *int32) *PlaceOfPreferencePSUpdate {
	if i != nil {
		poppu.SetPlacePrefNo(*i)
	}
	return poppu
}

// AddPlacePrefNo adds i to the "PlacePrefNo" field.
func (poppu *PlaceOfPreferencePSUpdate) AddPlacePrefNo(i int32) *PlaceOfPreferencePSUpdate {
	poppu.mutation.AddPlacePrefNo(i)
	return poppu
}

// ClearPlacePrefNo clears the value of the "PlacePrefNo" field.
func (poppu *PlaceOfPreferencePSUpdate) ClearPlacePrefNo() *PlaceOfPreferencePSUpdate {
	poppu.mutation.ClearPlacePrefNo()
	return poppu
}

// SetPlacePrefValue sets the "PlacePrefValue" field.
func (poppu *PlaceOfPreferencePSUpdate) SetPlacePrefValue(s string) *PlaceOfPreferencePSUpdate {
	poppu.mutation.SetPlacePrefValue(s)
	return poppu
}

// SetNillablePlacePrefValue sets the "PlacePrefValue" field if the given value is not nil.
func (poppu *PlaceOfPreferencePSUpdate) SetNillablePlacePrefValue(s *string) *PlaceOfPreferencePSUpdate {
	if s != nil {
		poppu.SetPlacePrefValue(*s)
	}
	return poppu
}

// ClearPlacePrefValue clears the value of the "PlacePrefValue" field.
func (poppu *PlaceOfPreferencePSUpdate) ClearPlacePrefValue() *PlaceOfPreferencePSUpdate {
	poppu.mutation.ClearPlacePrefValue()
	return poppu
}

// SetEmployeeID sets the "EmployeeID" field.
func (poppu *PlaceOfPreferencePSUpdate) SetEmployeeID(i int64) *PlaceOfPreferencePSUpdate {
	poppu.mutation.ResetEmployeeID()
	poppu.mutation.SetEmployeeID(i)
	return poppu
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (poppu *PlaceOfPreferencePSUpdate) SetNillableEmployeeID(i *int64) *PlaceOfPreferencePSUpdate {
	if i != nil {
		poppu.SetEmployeeID(*i)
	}
	return poppu
}

// AddEmployeeID adds i to the "EmployeeID" field.
func (poppu *PlaceOfPreferencePSUpdate) AddEmployeeID(i int64) *PlaceOfPreferencePSUpdate {
	poppu.mutation.AddEmployeeID(i)
	return poppu
}

// ClearEmployeeID clears the value of the "EmployeeID" field.
func (poppu *PlaceOfPreferencePSUpdate) ClearEmployeeID() *PlaceOfPreferencePSUpdate {
	poppu.mutation.ClearEmployeeID()
	return poppu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (poppu *PlaceOfPreferencePSUpdate) SetUpdatedAt(t time.Time) *PlaceOfPreferencePSUpdate {
	poppu.mutation.SetUpdatedAt(t)
	return poppu
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (poppu *PlaceOfPreferencePSUpdate) SetNillableUpdatedAt(t *time.Time) *PlaceOfPreferencePSUpdate {
	if t != nil {
		poppu.SetUpdatedAt(*t)
	}
	return poppu
}

// ClearUpdatedAt clears the value of the "UpdatedAt" field.
func (poppu *PlaceOfPreferencePSUpdate) ClearUpdatedAt() *PlaceOfPreferencePSUpdate {
	poppu.mutation.ClearUpdatedAt()
	return poppu
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (poppu *PlaceOfPreferencePSUpdate) SetUpdatedBy(s string) *PlaceOfPreferencePSUpdate {
	poppu.mutation.SetUpdatedBy(s)
	return poppu
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (poppu *PlaceOfPreferencePSUpdate) SetNillableUpdatedBy(s *string) *PlaceOfPreferencePSUpdate {
	if s != nil {
		poppu.SetUpdatedBy(*s)
	}
	return poppu
}

// ClearUpdatedBy clears the value of the "UpdatedBy" field.
func (poppu *PlaceOfPreferencePSUpdate) ClearUpdatedBy() *PlaceOfPreferencePSUpdate {
	poppu.mutation.ClearUpdatedBy()
	return poppu
}

// SetPlaceApplnPSRefID sets the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity by ID.
func (poppu *PlaceOfPreferencePSUpdate) SetPlaceApplnPSRefID(id int64) *PlaceOfPreferencePSUpdate {
	poppu.mutation.SetPlaceApplnPSRefID(id)
	return poppu
}

// SetNillablePlaceApplnPSRefID sets the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity by ID if the given value is not nil.
func (poppu *PlaceOfPreferencePSUpdate) SetNillablePlaceApplnPSRefID(id *int64) *PlaceOfPreferencePSUpdate {
	if id != nil {
		poppu = poppu.SetPlaceApplnPSRefID(*id)
	}
	return poppu
}

// SetPlaceApplnPSRef sets the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity.
func (poppu *PlaceOfPreferencePSUpdate) SetPlaceApplnPSRef(e *Exam_Applications_PS) *PlaceOfPreferencePSUpdate {
	return poppu.SetPlaceApplnPSRefID(e.ID)
}

// Mutation returns the PlaceOfPreferencePSMutation object of the builder.
func (poppu *PlaceOfPreferencePSUpdate) Mutation() *PlaceOfPreferencePSMutation {
	return poppu.mutation
}

// ClearPlaceApplnPSRef clears the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity.
func (poppu *PlaceOfPreferencePSUpdate) ClearPlaceApplnPSRef() *PlaceOfPreferencePSUpdate {
	poppu.mutation.ClearPlaceApplnPSRef()
	return poppu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (poppu *PlaceOfPreferencePSUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, poppu.sqlSave, poppu.mutation, poppu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (poppu *PlaceOfPreferencePSUpdate) SaveX(ctx context.Context) int {
	affected, err := poppu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (poppu *PlaceOfPreferencePSUpdate) Exec(ctx context.Context) error {
	_, err := poppu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (poppu *PlaceOfPreferencePSUpdate) ExecX(ctx context.Context) {
	if err := poppu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (poppu *PlaceOfPreferencePSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(placeofpreferenceps.Table, placeofpreferenceps.Columns, sqlgraph.NewFieldSpec(placeofpreferenceps.FieldID, field.TypeInt32))
	if ps := poppu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := poppu.mutation.PlacePrefNo(); ok {
		_spec.SetField(placeofpreferenceps.FieldPlacePrefNo, field.TypeInt32, value)
	}
	if value, ok := poppu.mutation.AddedPlacePrefNo(); ok {
		_spec.AddField(placeofpreferenceps.FieldPlacePrefNo, field.TypeInt32, value)
	}
	if poppu.mutation.PlacePrefNoCleared() {
		_spec.ClearField(placeofpreferenceps.FieldPlacePrefNo, field.TypeInt32)
	}
	if value, ok := poppu.mutation.PlacePrefValue(); ok {
		_spec.SetField(placeofpreferenceps.FieldPlacePrefValue, field.TypeString, value)
	}
	if poppu.mutation.PlacePrefValueCleared() {
		_spec.ClearField(placeofpreferenceps.FieldPlacePrefValue, field.TypeString)
	}
	if value, ok := poppu.mutation.EmployeeID(); ok {
		_spec.SetField(placeofpreferenceps.FieldEmployeeID, field.TypeInt64, value)
	}
	if value, ok := poppu.mutation.AddedEmployeeID(); ok {
		_spec.AddField(placeofpreferenceps.FieldEmployeeID, field.TypeInt64, value)
	}
	if poppu.mutation.EmployeeIDCleared() {
		_spec.ClearField(placeofpreferenceps.FieldEmployeeID, field.TypeInt64)
	}
	if value, ok := poppu.mutation.UpdatedAt(); ok {
		_spec.SetField(placeofpreferenceps.FieldUpdatedAt, field.TypeTime, value)
	}
	if poppu.mutation.UpdatedAtCleared() {
		_spec.ClearField(placeofpreferenceps.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := poppu.mutation.UpdatedBy(); ok {
		_spec.SetField(placeofpreferenceps.FieldUpdatedBy, field.TypeString, value)
	}
	if poppu.mutation.UpdatedByCleared() {
		_spec.ClearField(placeofpreferenceps.FieldUpdatedBy, field.TypeString)
	}
	if poppu.mutation.PlaceApplnPSRefCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := poppu.mutation.PlaceApplnPSRefIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, poppu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{placeofpreferenceps.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	poppu.mutation.done = true
	return n, nil
}

// PlaceOfPreferencePSUpdateOne is the builder for updating a single PlaceOfPreferencePS entity.
type PlaceOfPreferencePSUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlaceOfPreferencePSMutation
}

// SetApplicationID sets the "ApplicationID" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetApplicationID(i int64) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.SetApplicationID(i)
	return poppuo
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetNillableApplicationID(i *int64) *PlaceOfPreferencePSUpdateOne {
	if i != nil {
		poppuo.SetApplicationID(*i)
	}
	return poppuo
}

// ClearApplicationID clears the value of the "ApplicationID" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) ClearApplicationID() *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.ClearApplicationID()
	return poppuo
}

// SetPlacePrefNo sets the "PlacePrefNo" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetPlacePrefNo(i int32) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.ResetPlacePrefNo()
	poppuo.mutation.SetPlacePrefNo(i)
	return poppuo
}

// SetNillablePlacePrefNo sets the "PlacePrefNo" field if the given value is not nil.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetNillablePlacePrefNo(i *int32) *PlaceOfPreferencePSUpdateOne {
	if i != nil {
		poppuo.SetPlacePrefNo(*i)
	}
	return poppuo
}

// AddPlacePrefNo adds i to the "PlacePrefNo" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) AddPlacePrefNo(i int32) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.AddPlacePrefNo(i)
	return poppuo
}

// ClearPlacePrefNo clears the value of the "PlacePrefNo" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) ClearPlacePrefNo() *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.ClearPlacePrefNo()
	return poppuo
}

// SetPlacePrefValue sets the "PlacePrefValue" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetPlacePrefValue(s string) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.SetPlacePrefValue(s)
	return poppuo
}

// SetNillablePlacePrefValue sets the "PlacePrefValue" field if the given value is not nil.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetNillablePlacePrefValue(s *string) *PlaceOfPreferencePSUpdateOne {
	if s != nil {
		poppuo.SetPlacePrefValue(*s)
	}
	return poppuo
}

// ClearPlacePrefValue clears the value of the "PlacePrefValue" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) ClearPlacePrefValue() *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.ClearPlacePrefValue()
	return poppuo
}

// SetEmployeeID sets the "EmployeeID" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetEmployeeID(i int64) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.ResetEmployeeID()
	poppuo.mutation.SetEmployeeID(i)
	return poppuo
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetNillableEmployeeID(i *int64) *PlaceOfPreferencePSUpdateOne {
	if i != nil {
		poppuo.SetEmployeeID(*i)
	}
	return poppuo
}

// AddEmployeeID adds i to the "EmployeeID" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) AddEmployeeID(i int64) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.AddEmployeeID(i)
	return poppuo
}

// ClearEmployeeID clears the value of the "EmployeeID" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) ClearEmployeeID() *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.ClearEmployeeID()
	return poppuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetUpdatedAt(t time.Time) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.SetUpdatedAt(t)
	return poppuo
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetNillableUpdatedAt(t *time.Time) *PlaceOfPreferencePSUpdateOne {
	if t != nil {
		poppuo.SetUpdatedAt(*t)
	}
	return poppuo
}

// ClearUpdatedAt clears the value of the "UpdatedAt" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) ClearUpdatedAt() *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.ClearUpdatedAt()
	return poppuo
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetUpdatedBy(s string) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.SetUpdatedBy(s)
	return poppuo
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetNillableUpdatedBy(s *string) *PlaceOfPreferencePSUpdateOne {
	if s != nil {
		poppuo.SetUpdatedBy(*s)
	}
	return poppuo
}

// ClearUpdatedBy clears the value of the "UpdatedBy" field.
func (poppuo *PlaceOfPreferencePSUpdateOne) ClearUpdatedBy() *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.ClearUpdatedBy()
	return poppuo
}

// SetPlaceApplnPSRefID sets the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity by ID.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetPlaceApplnPSRefID(id int64) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.SetPlaceApplnPSRefID(id)
	return poppuo
}

// SetNillablePlaceApplnPSRefID sets the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity by ID if the given value is not nil.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetNillablePlaceApplnPSRefID(id *int64) *PlaceOfPreferencePSUpdateOne {
	if id != nil {
		poppuo = poppuo.SetPlaceApplnPSRefID(*id)
	}
	return poppuo
}

// SetPlaceApplnPSRef sets the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity.
func (poppuo *PlaceOfPreferencePSUpdateOne) SetPlaceApplnPSRef(e *Exam_Applications_PS) *PlaceOfPreferencePSUpdateOne {
	return poppuo.SetPlaceApplnPSRefID(e.ID)
}

// Mutation returns the PlaceOfPreferencePSMutation object of the builder.
func (poppuo *PlaceOfPreferencePSUpdateOne) Mutation() *PlaceOfPreferencePSMutation {
	return poppuo.mutation
}

// ClearPlaceApplnPSRef clears the "PlaceApplnPS_Ref" edge to the Exam_Applications_PS entity.
func (poppuo *PlaceOfPreferencePSUpdateOne) ClearPlaceApplnPSRef() *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.ClearPlaceApplnPSRef()
	return poppuo
}

// Where appends a list predicates to the PlaceOfPreferencePSUpdate builder.
func (poppuo *PlaceOfPreferencePSUpdateOne) Where(ps ...predicate.PlaceOfPreferencePS) *PlaceOfPreferencePSUpdateOne {
	poppuo.mutation.Where(ps...)
	return poppuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (poppuo *PlaceOfPreferencePSUpdateOne) Select(field string, fields ...string) *PlaceOfPreferencePSUpdateOne {
	poppuo.fields = append([]string{field}, fields...)
	return poppuo
}

// Save executes the query and returns the updated PlaceOfPreferencePS entity.
func (poppuo *PlaceOfPreferencePSUpdateOne) Save(ctx context.Context) (*PlaceOfPreferencePS, error) {
	return withHooks(ctx, poppuo.sqlSave, poppuo.mutation, poppuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (poppuo *PlaceOfPreferencePSUpdateOne) SaveX(ctx context.Context) *PlaceOfPreferencePS {
	node, err := poppuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (poppuo *PlaceOfPreferencePSUpdateOne) Exec(ctx context.Context) error {
	_, err := poppuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (poppuo *PlaceOfPreferencePSUpdateOne) ExecX(ctx context.Context) {
	if err := poppuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (poppuo *PlaceOfPreferencePSUpdateOne) sqlSave(ctx context.Context) (_node *PlaceOfPreferencePS, err error) {
	_spec := sqlgraph.NewUpdateSpec(placeofpreferenceps.Table, placeofpreferenceps.Columns, sqlgraph.NewFieldSpec(placeofpreferenceps.FieldID, field.TypeInt32))
	id, ok := poppuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PlaceOfPreferencePS.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := poppuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, placeofpreferenceps.FieldID)
		for _, f := range fields {
			if !placeofpreferenceps.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != placeofpreferenceps.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := poppuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := poppuo.mutation.PlacePrefNo(); ok {
		_spec.SetField(placeofpreferenceps.FieldPlacePrefNo, field.TypeInt32, value)
	}
	if value, ok := poppuo.mutation.AddedPlacePrefNo(); ok {
		_spec.AddField(placeofpreferenceps.FieldPlacePrefNo, field.TypeInt32, value)
	}
	if poppuo.mutation.PlacePrefNoCleared() {
		_spec.ClearField(placeofpreferenceps.FieldPlacePrefNo, field.TypeInt32)
	}
	if value, ok := poppuo.mutation.PlacePrefValue(); ok {
		_spec.SetField(placeofpreferenceps.FieldPlacePrefValue, field.TypeString, value)
	}
	if poppuo.mutation.PlacePrefValueCleared() {
		_spec.ClearField(placeofpreferenceps.FieldPlacePrefValue, field.TypeString)
	}
	if value, ok := poppuo.mutation.EmployeeID(); ok {
		_spec.SetField(placeofpreferenceps.FieldEmployeeID, field.TypeInt64, value)
	}
	if value, ok := poppuo.mutation.AddedEmployeeID(); ok {
		_spec.AddField(placeofpreferenceps.FieldEmployeeID, field.TypeInt64, value)
	}
	if poppuo.mutation.EmployeeIDCleared() {
		_spec.ClearField(placeofpreferenceps.FieldEmployeeID, field.TypeInt64)
	}
	if value, ok := poppuo.mutation.UpdatedAt(); ok {
		_spec.SetField(placeofpreferenceps.FieldUpdatedAt, field.TypeTime, value)
	}
	if poppuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(placeofpreferenceps.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := poppuo.mutation.UpdatedBy(); ok {
		_spec.SetField(placeofpreferenceps.FieldUpdatedBy, field.TypeString, value)
	}
	if poppuo.mutation.UpdatedByCleared() {
		_spec.ClearField(placeofpreferenceps.FieldUpdatedBy, field.TypeString)
	}
	if poppuo.mutation.PlaceApplnPSRefCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := poppuo.mutation.PlaceApplnPSRefIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PlaceOfPreferencePS{config: poppuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, poppuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{placeofpreferenceps.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	poppuo.mutation.done = true
	return _node, nil
}