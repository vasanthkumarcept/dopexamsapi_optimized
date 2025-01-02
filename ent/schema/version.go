package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type Version struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (Version) Fields() []ent.Field {
	return []ent.Field{
		field.String("UiVersion").Optional(),
		field.String("ApiVersion").Optional(),
		field.Int32("ApiType").Optional(),
	}
}

// Edges of the ExamPapers.

func (Version) Edges() []ent.Edge {

	return []ent.Edge{}

}

func (Version) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Version"}}
}
