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
type RoleMaster struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (RoleMaster) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("RoleUserCode"),
		field.String("RoleName"),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.Bool("Status").Default(true),
		field.Int32("RoleCode").Optional(),
	}
}

// Edges of the ExamPapers.

func (RoleMaster) Edges() []ent.Edge {

	return []ent.Edge{edge.To("roles", AdminLogin.Type),
		edge.To("Roles_Ref", UserMaster.Type),
		edge.To("Roles_PS_Ref", Exam_Applications_PS.Type),
		edge.To("Roles_IP_Ref", Exam_Applications_IP.Type),
		edge.To("Roles_GDSPA_Ref", Exam_Applications_GDSPA.Type),
		edge.To("Roles_GDSPM_Ref", Exam_Applications_GDSPM.Type),
		edge.To("Roles_PMPA_Ref", Exam_Applications_PMPA.Type),
		edge.To("Roles_MTSPMMG_Ref", Exam_Application_MTSPMMG.Type),
		//edge.To("Roles_IP1_Ref", Exam_Applications_IP1.Type)
	}

}

func (RoleMaster) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "RoleMaster"}}
}
