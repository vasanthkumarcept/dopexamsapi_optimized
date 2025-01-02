package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type PostExamPaper struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (PostExamPaper) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("Id"),
		field.Int32("ExamConfigurationExamCode").Optional(),
		field.String("ExamShortDescription").Optional(),
		field.String("ExamLongDescription").Optional(),
		field.String("ExamPaperCode").Optional(),
		field.String("PaperDescription").Optional(),
		field.String("EmployeePost_postId").Optional(),
		field.String("EmployeeGroup_GroupId").Optional(),
		field.String("GroupDescription").Optional(),
		field.String("PostCode").Optional(),
		field.String("PostDescription").Optional(),
		field.Int("BaseCadre").Optional(),
		field.String("PayLevel").Optional(),
		field.String("Scale").Optional(),
		field.String("OrderNumber").Optional(),
		field.String("Status").Optional(),
		field.Int64("CreatedById").Optional(),
		field.String("CreatedByUserName").Optional(),
		field.String("CreatedByEmployeeId").Optional(),
		field.String("CreatedByDesignation").Optional(),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Int64("VerifiedById").Optional(),
		field.String("VerifiedByUserName").Optional(),
		field.String("VerifiedByEmployeeId").Optional(),
		field.String("VerifiedByDesignation").Optional(),
		field.Time("VerifiedDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Int64("DeletedById").Optional(),
		field.String("DeletedByUserName").Optional(),
		field.String("DeletedByEmployeeId").Optional(),
		field.String("DeletedByDesignation").Optional(),
		field.Time("DeletedDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
	}
}

// // Edges of the ExamPapers.
// func (PostExamPaper) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.From("papercode", ExamPapers.Type).Ref("exampapers_types").Unique().Field("PaperCode"),
// 	}
// }
