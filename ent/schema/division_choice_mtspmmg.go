package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Division_Choice_MTSPMMG holds the schema definition for the Division_Choice_MTSPMMG entity.
type Division_Choice_MTSPMMG struct {
	ent.Schema
}

// Fields of the Division_Choice_MTSPMMG.
func (Division_Choice_MTSPMMG) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("CadrePrefId"),
		field.Int64("ApplicationID").Optional(),
		field.String("Group").Optional(),
		field.Int64("CadrePrefNo").Optional(),
		field.String("Cadre").Optional(),
		field.Int64("PlacePrefNo").Optional(),
		field.String("PlacePrefValue").Optional(),
		field.Int64("PostPrefNo").Optional(),
		//field.Int32("NumOfPapers"),
		field.String("PostingPrefValue").Optional(),
		field.Int64("EmployeeID").Optional(),
		field.Time("UpdatedAt").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Default(time.Now).Optional(),
		field.String("UpdatedBy").Default("API").Optional()}
}

// Edges of the Division_Choice_MTSPMMG.
func (Division_Choice_MTSPMMG) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ApplnMTSPMMG_Ref", Exam_Application_MTSPMMG.Type).Ref("CirclePrefRefMTSPMMG").Unique().Field("ApplicationID"),
	}
}

func (Division_Choice_MTSPMMG) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Division_Choice_MTSPMMG"}}
}
