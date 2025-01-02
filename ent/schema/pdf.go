package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Log holds the schema definition for the Log entity.
type PDF struct {
	ent.Schema
}

// Fields of the Log.
func (PDF) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").
			Optional(),
		field.String("filename").
			Optional(),
		field.Time("eventtime").
		SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
			}).
		Optional().
		Default(time.Now),
		field.Int("examcode").
			Optional(),
		field.Int("year").
			Optional(),
	}
}

// Edges of the Log.
func (PDF) Edges() []ent.Edge {
	return nil
}

func (PDF) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "PDF"}}
}
