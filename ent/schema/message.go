package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type Message struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("ID"),
		field.String("Description").Optional(),
		field.Int32("Priority").Default(0),
		field.Bool("Status").Default(false),
	}
}

// Edges of the ExamPapers.

func (Message) Edges() []ent.Edge {

	return []ent.Edge{}

}

func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Message"}}
}
