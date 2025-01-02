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

// Cadre_Choice_MTSPMMG holds the schema definition for the Cadre_Choice_MTSPMMG entity.
type Cadre_Choice_MTSPMMG struct {
	ent.Schema
}

// Fields of the Cadre_Choice_MTSPMMG.
func (Cadre_Choice_MTSPMMG) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("CadrePrefId"),
		field.Int64("ApplicationID").Optional(),
		field.Int64("PlacePrefNo"),
		//field.Int32("NumOfPapers"),
		field.String("PlacePrefValue"),
		field.String("FeederCader").Optional(),
		field.String("FeederCaderDateOfJoining").Optional(),
		field.Int64("EmployeeID").Optional(),
		field.Time("UpdatedAt").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Default(time.Now).Optional(),
		field.String("UpdatedBy").Default("API").Optional()}
}

// Edges of the Cadre_Choice_MTSPMMG.
func (Cadre_Choice_MTSPMMG) Edges() []ent.Edge {
	return []ent.Edge{edge.From("ApplnMTSPMMG_Ref", Exam_Application_MTSPMMG.Type).Ref("CadrePref_Ref").Unique().Field("ApplicationID")}
}

func (Cadre_Choice_MTSPMMG) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Cadre_Choice_MTSPMMG"}}
}
