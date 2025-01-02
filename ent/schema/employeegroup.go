package schema

import (
	"entgo.io/ent"
	//"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// FacilityNewTab holds the schema definition for the Facility entity.
type EmployeeGroup struct {
	ent.Schema
}

// Fields of the Facility.
func (EmployeeGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("groupid"),
		// Add other fields of Facility entity here.
		field.String("GroupCode").Optional(),
		field.String("GroupDescription").Optional(),
		field.String("OrderNumber").Optional(),
		field.String("Status").Optional(),
		field.Int32("CreatedById").Optional(),
		field.String("CreatedByUserName").Optional(),
		field.String("CreatedByEmployeeId").Optional(),
		field.String("CreatedByDesignation").Optional(),
		field.Time("CreatedDate").Optional(),
		field.Int64("VerifiedById").Optional(),
		field.String("VerifiedByUserName").Optional(),
		field.String("VerifiedByEmployeeId").Optional(),
		field.String("VerifiedByDesignation").Optional(),
		field.Time("VerifiedDate").Optional(),
		field.Int64("DeletedById").Optional(),
		field.String("DeletedByUserName").Optional(),
		field.String("DeletedByEmployeeId").Optional(),
		field.String("DeletedByDesignation").Optional(),
		field.Time("DeletedDate").Optional(),
	}
}

// Edges of the Facility.
func (EmployeeGroup) Edges() []ent.Edge {

	return []ent.Edge{}
}
func (EmployeeGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "EmployeeGroup"}}
}
