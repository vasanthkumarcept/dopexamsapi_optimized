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

// RecommendationsMTSPMMGApplications holds the schema definition for the RecommendationsMTSPMMGApplications entity.
type RecommendationsMTSPMMGApplications struct {
	ent.Schema
}

// Fields of the RecommendationsMTSPMMGApplications.
func (RecommendationsMTSPMMGApplications) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("RecommendationId"),
		field.Int64("ApplicationID").Unique().Optional(),
		field.Int64("EmployeeID").Optional(),
		field.String("ExamNameCode").Optional(),
		field.String("ExamYear").Optional(),
		field.String("ExamName").Optional(),
		field.Int32("VacancyYear").Optional(),

		field.String("CA_Recommendations").Optional(),
		//field.String("CA_UpdatedAt").Optional(),
		field.Time("CA_UpdatedAt").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.String("CA_UserName").Optional(),
		field.String("CA_Remarks").Optional(),
		field.String("NO_Recommendations").Optional(),
		//field.String("NO_UpdatedAt").Optional(),
		field.Time("NO_UpdatedAt").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.String("NO_UserName").Optional(),
		field.String("NO_Remarks").Optional(),
		field.String("ApplicationStatus").Optional(),

		//field.String("AppliactionRemarks").Optional(),
		field.Time("UpdatedAt").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Default(time.Now).Optional(),
		field.String("UpdatedBy").Default("API").Optional(),
		// Boolean Value to add the hall ticket generation .
		field.Bool("GenerateHallTicketFlag").Optional(),
	}

}

// Edges of the RecommendationsMTSPMMGApplications.
func (RecommendationsMTSPMMGApplications) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ApplnRef", Exam_Application_MTSPMMG.Type).Ref("MTSPMMGApplicationsRef").Unique().Field("ApplicationID"),
	}

}

func (RecommendationsMTSPMMGApplications) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "RecommendationsMTSPMMGApplications"}}
}