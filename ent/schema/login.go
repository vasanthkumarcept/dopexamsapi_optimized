package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Login struct {
	ent.Schema
}

func (Login) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("loginID", uuid.UUID{}).
			Default(uuid.New).      // Set a default value for the field.
			StorageKey("login_id"). // Set the storage key for the field.
			Immutable(),            // Make the field immutable.
		field.String("username").MaxLen(100).Unique(),
		field.String("password"),
		field.Int32("EmployeedID").Optional(),
		field.Int32("expiremins_token"),
		field.Int32("expiremins_refresh_token"),
		field.Int32("role"),
		field.String("token"),
		field.String("VerifyRemarks").Optional(),
	}
}
func (Login) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("emplogin", Employees.Type).Ref("emp_login").Unique().Field("EmployeedID"),
		//edge.To("emp_Region", "RegionMaster.Type"),
		//edge.To("emp_Division"), "DivisionMaster.Type"),
		//edge.To("emp_facility"), "Facility.Type"
		edge.To("LogData", Logs.Type),
	}
}
func (Login) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Login"}}
}
