// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type EmployeePosts struct {
	ent.Schema
}

func (EmployeePosts) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("PostID"), 
		field.String("PostCode"),
		field.String("PostDescription").Optional(),
		field.String("Group").Optional(),
		field.String("PayLevel").Optional(),
		field.String("Scale").Optional(),
		field.Bool("BaseCadreFlag").Optional(),
		//field.String("OrderNumber").Optional(),

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

field.String("GroupDescription").Optional(),

	
	}
}
func (EmployeePosts) Edges() []ent.Edge {
	return []ent.Edge{edge.To("emp_posts",Employees.Type),
	edge.To("PostEligibility", EligibilityMaster.Type),
}
}
func (EmployeePosts) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "EmployeePosts"}}
}