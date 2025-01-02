package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type AdminLogin struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (AdminLogin) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("LoginId"),
		field.Int32("RoleUserCode").Optional(),
		field.String("RoleName").Optional(),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.String("Status").Optional(),
		field.Int32("EmployeedID").Optional(),
		field.String("EmployeeName").Optional(),
		field.String("Emailid").Optional(),
		field.Int64("MobileNumber").Optional(),
		field.String("Username"),
		field.Int32("OTP").Optional(),
		field.String("Password"),
		field.String("VerifyRemarks").
			Nillable().
			Optional(),
	}
}

// Edges of the ExamPapers.

func (AdminLogin) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("role_master", RoleMaster.Type).Ref("roles").Unique().Field("RoleUserCode"),
		edge.To("LogData", Logs.Type),
	}

}

func (AdminLogin) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "AdminLogin"}}
}
