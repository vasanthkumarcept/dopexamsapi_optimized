// Code generated by entimport, DO NOT EDIT.

package schema

import (
	
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent"
	//"time"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	//"encoding/json"
)

// Exam holds the schema definition for the Exam entity.
type PlaceOfPreferenceIP struct {
	ent.Schema
}

/// Fields of the Exam.
func (PlaceOfPreferenceIP) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("PlacePrefId"),
		field.Int64("ApplicationID").Optional(),
		field.Int32("PlacePrefNo").Optional(),
		field.String("PlacePrefValue").Optional(),
		field.Int64("EmployeeID").Optional(),
		field.Time("UpdatedAt").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}).
			Optional(),
		field.String("UpdatedBy").Default("API").Optional(),
	}
}

// Edges of the Exam.
func (PlaceOfPreferenceIP) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ApplnIP_Ref", Exam_Applications_IP.Type).
			Ref("CirclePrefRef").
			Unique(),
			//Field("ApplicationID"),
	}
}

func (PlaceOfPreferenceIP) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "PlaceOfPreferenceIP"}}
}


