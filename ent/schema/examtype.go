package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type ExamType struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (ExamType) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("ExamTypeCode"),
		field.String("ExamType"),
		field.String("Status"),
		field.Int32("ExamCode").Optional(),
	}
}

// Edges of the ExamPapers.

func (ExamType) Edges() []ent.Edge {

	return []ent.Edge{edge.From("exam", Exam.Type).Ref("exams_type").Unique().Field("ExamCode")} //	edge.From("exam", Exam.Type).Ref("exams_type").Unique().Field("ExamCode"),

}

func (ExamType) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "ExamType"}}
}
