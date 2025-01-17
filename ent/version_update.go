// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/predicate"
	"recruit/ent/version"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VersionUpdate is the builder for updating Version entities.
type VersionUpdate struct {
	config
	hooks    []Hook
	mutation *VersionMutation
}

// Where appends a list predicates to the VersionUpdate builder.
func (vu *VersionUpdate) Where(ps ...predicate.Version) *VersionUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetUiVersion sets the "UiVersion" field.
func (vu *VersionUpdate) SetUiVersion(s string) *VersionUpdate {
	vu.mutation.SetUiVersion(s)
	return vu
}

// SetNillableUiVersion sets the "UiVersion" field if the given value is not nil.
func (vu *VersionUpdate) SetNillableUiVersion(s *string) *VersionUpdate {
	if s != nil {
		vu.SetUiVersion(*s)
	}
	return vu
}

// ClearUiVersion clears the value of the "UiVersion" field.
func (vu *VersionUpdate) ClearUiVersion() *VersionUpdate {
	vu.mutation.ClearUiVersion()
	return vu
}

// SetApiVersion sets the "ApiVersion" field.
func (vu *VersionUpdate) SetApiVersion(s string) *VersionUpdate {
	vu.mutation.SetApiVersion(s)
	return vu
}

// SetNillableApiVersion sets the "ApiVersion" field if the given value is not nil.
func (vu *VersionUpdate) SetNillableApiVersion(s *string) *VersionUpdate {
	if s != nil {
		vu.SetApiVersion(*s)
	}
	return vu
}

// ClearApiVersion clears the value of the "ApiVersion" field.
func (vu *VersionUpdate) ClearApiVersion() *VersionUpdate {
	vu.mutation.ClearApiVersion()
	return vu
}

// SetApiType sets the "ApiType" field.
func (vu *VersionUpdate) SetApiType(i int32) *VersionUpdate {
	vu.mutation.ResetApiType()
	vu.mutation.SetApiType(i)
	return vu
}

// SetNillableApiType sets the "ApiType" field if the given value is not nil.
func (vu *VersionUpdate) SetNillableApiType(i *int32) *VersionUpdate {
	if i != nil {
		vu.SetApiType(*i)
	}
	return vu
}

// AddApiType adds i to the "ApiType" field.
func (vu *VersionUpdate) AddApiType(i int32) *VersionUpdate {
	vu.mutation.AddApiType(i)
	return vu
}

// ClearApiType clears the value of the "ApiType" field.
func (vu *VersionUpdate) ClearApiType() *VersionUpdate {
	vu.mutation.ClearApiType()
	return vu
}

// Mutation returns the VersionMutation object of the builder.
func (vu *VersionUpdate) Mutation() *VersionMutation {
	return vu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VersionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VersionUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VersionUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VersionUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vu *VersionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(version.Table, version.Columns, sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.UiVersion(); ok {
		_spec.SetField(version.FieldUiVersion, field.TypeString, value)
	}
	if vu.mutation.UiVersionCleared() {
		_spec.ClearField(version.FieldUiVersion, field.TypeString)
	}
	if value, ok := vu.mutation.ApiVersion(); ok {
		_spec.SetField(version.FieldApiVersion, field.TypeString, value)
	}
	if vu.mutation.ApiVersionCleared() {
		_spec.ClearField(version.FieldApiVersion, field.TypeString)
	}
	if value, ok := vu.mutation.ApiType(); ok {
		_spec.SetField(version.FieldApiType, field.TypeInt32, value)
	}
	if value, ok := vu.mutation.AddedApiType(); ok {
		_spec.AddField(version.FieldApiType, field.TypeInt32, value)
	}
	if vu.mutation.ApiTypeCleared() {
		_spec.ClearField(version.FieldApiType, field.TypeInt32)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{version.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VersionUpdateOne is the builder for updating a single Version entity.
type VersionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VersionMutation
}

// SetUiVersion sets the "UiVersion" field.
func (vuo *VersionUpdateOne) SetUiVersion(s string) *VersionUpdateOne {
	vuo.mutation.SetUiVersion(s)
	return vuo
}

// SetNillableUiVersion sets the "UiVersion" field if the given value is not nil.
func (vuo *VersionUpdateOne) SetNillableUiVersion(s *string) *VersionUpdateOne {
	if s != nil {
		vuo.SetUiVersion(*s)
	}
	return vuo
}

// ClearUiVersion clears the value of the "UiVersion" field.
func (vuo *VersionUpdateOne) ClearUiVersion() *VersionUpdateOne {
	vuo.mutation.ClearUiVersion()
	return vuo
}

// SetApiVersion sets the "ApiVersion" field.
func (vuo *VersionUpdateOne) SetApiVersion(s string) *VersionUpdateOne {
	vuo.mutation.SetApiVersion(s)
	return vuo
}

// SetNillableApiVersion sets the "ApiVersion" field if the given value is not nil.
func (vuo *VersionUpdateOne) SetNillableApiVersion(s *string) *VersionUpdateOne {
	if s != nil {
		vuo.SetApiVersion(*s)
	}
	return vuo
}

// ClearApiVersion clears the value of the "ApiVersion" field.
func (vuo *VersionUpdateOne) ClearApiVersion() *VersionUpdateOne {
	vuo.mutation.ClearApiVersion()
	return vuo
}

// SetApiType sets the "ApiType" field.
func (vuo *VersionUpdateOne) SetApiType(i int32) *VersionUpdateOne {
	vuo.mutation.ResetApiType()
	vuo.mutation.SetApiType(i)
	return vuo
}

// SetNillableApiType sets the "ApiType" field if the given value is not nil.
func (vuo *VersionUpdateOne) SetNillableApiType(i *int32) *VersionUpdateOne {
	if i != nil {
		vuo.SetApiType(*i)
	}
	return vuo
}

// AddApiType adds i to the "ApiType" field.
func (vuo *VersionUpdateOne) AddApiType(i int32) *VersionUpdateOne {
	vuo.mutation.AddApiType(i)
	return vuo
}

// ClearApiType clears the value of the "ApiType" field.
func (vuo *VersionUpdateOne) ClearApiType() *VersionUpdateOne {
	vuo.mutation.ClearApiType()
	return vuo
}

// Mutation returns the VersionMutation object of the builder.
func (vuo *VersionUpdateOne) Mutation() *VersionMutation {
	return vuo.mutation
}

// Where appends a list predicates to the VersionUpdate builder.
func (vuo *VersionUpdateOne) Where(ps ...predicate.Version) *VersionUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VersionUpdateOne) Select(field string, fields ...string) *VersionUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Version entity.
func (vuo *VersionUpdateOne) Save(ctx context.Context) (*Version, error) {
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VersionUpdateOne) SaveX(ctx context.Context) *Version {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VersionUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VersionUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vuo *VersionUpdateOne) sqlSave(ctx context.Context) (_node *Version, err error) {
	_spec := sqlgraph.NewUpdateSpec(version.Table, version.Columns, sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Version.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, version.FieldID)
		for _, f := range fields {
			if !version.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != version.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.UiVersion(); ok {
		_spec.SetField(version.FieldUiVersion, field.TypeString, value)
	}
	if vuo.mutation.UiVersionCleared() {
		_spec.ClearField(version.FieldUiVersion, field.TypeString)
	}
	if value, ok := vuo.mutation.ApiVersion(); ok {
		_spec.SetField(version.FieldApiVersion, field.TypeString, value)
	}
	if vuo.mutation.ApiVersionCleared() {
		_spec.ClearField(version.FieldApiVersion, field.TypeString)
	}
	if value, ok := vuo.mutation.ApiType(); ok {
		_spec.SetField(version.FieldApiType, field.TypeInt32, value)
	}
	if value, ok := vuo.mutation.AddedApiType(); ok {
		_spec.AddField(version.FieldApiType, field.TypeInt32, value)
	}
	if vuo.mutation.ApiTypeCleared() {
		_spec.ClearField(version.FieldApiType, field.TypeInt32)
	}
	_node = &Version{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{version.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
