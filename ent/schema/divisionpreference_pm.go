// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"

	//"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Exam holds the schema definition for the Exam entity.
type Division_Choice_PM struct {
	ent.Schema
}

/// Fields of the Exam.
func (Division_Choice_PM) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("CadrePrefId"), 
	field.Int64("ApplicationID").Optional(),
	field.String("Group"),
	field.Int64("CadrePrefNo"), 
	field.String("Cadre"),
	field.Int64("PostPrefNo"), 
	//field.Int32("NumOfPapers"),
	field.String("PostingPrefValue"), 	
	field.Int64("EmployeeID").Optional(),	
	field.Time("UpdatedAt").SchemaType(map[string]string{
		dialect.Postgres: "date",
	}).SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Default(time.Now).Optional(),	
	field.String("UpdatedBy").Default("API").Optional(),}
}
// Edges of the Exam.
func (Division_Choice_PM) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ApplnGDSPM_Ref", Exam_Applications_GDSPM.Type).Ref("CirclePrefRefGDSPM").Unique().Field("ApplicationID"),
	}
	//[]ent.Edge{		
	//edge.To("exam", Exam.Type).Ref("papers").Unique().Field("ExamCode"),
	//edge.To("examcal_ps_ref", ExamCalendar.Type),	
}
func (Division_Choice_PM) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Division_Choice_PM"}}
}
