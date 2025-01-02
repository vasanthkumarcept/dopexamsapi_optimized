package schema

import (
	"entgo.io/ent"
	//"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	//"entgo.io/ent/schema/index"
	//"entgo.io/ent/schema/mixin"
)

// DirectorateUsers holds the schema definition for the DirectorateUsers entity.
type DirectorateUsers struct {
	ent.Schema
}

// Fields of the DirectorateUsers.
func (DirectorateUsers) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("RoleUserCode"),
		field.String("Role"),
		field.Int32("EmployeedID"),
		field.String("EmployeeName"),
		field.String("EmailId"),
		field.Int64("MobileNumber"),
		field.Int32("SequenceNumber").Optional(),
		field.String("Status").Optional().Default("active"),
		/*Validate(func(s string) error {
			// Add custom validation logic if required.
			return nil
		}).*/

	}
}

// Edges of the DirectorateUsers.
func (DirectorateUsers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employee_user", Employees.Type),
	}

}

/*
// Indexes of the DirectorateUsers.

	func (DirectorateUsers) Indexes() []ent.Index {
		return []ent.Index{
			index.Fields("employeeCode").
				Unique(),
			index.Fields("emailId").
				Unique(),
		}
	}
*/
func (DirectorateUsers) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "DirectorateUsers"}}
}
