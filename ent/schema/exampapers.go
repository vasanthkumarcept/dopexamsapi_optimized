package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	//"entgo.io/ent/dialect/entsql"
	//"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type ExamPapers struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (ExamPapers) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("PaperCode"),
		field.String("PaperDescription"),
		field.Int32("ExamCode").Optional(),
		field.String("ExamName").Optional(),
		field.String("ExamShortName").Optional(),
		field.Int32("PaperTypeCode").Optional(),
		field.String("PaperTypeName").Optional(),
		field.Bool("CompetitiveQualifying").Default(false),
		field.Bool("ExceptionForDisability").Default(false),
		field.Int("MaximumMarks").Positive(),
		field.Int("Duration").Positive(),
		field.String("localLanguageAllowedQuestionPaper").MaxLen(10),
		field.String("localLanguageAllowedAnswerPaper").MaxLen(10),
		field.Int32("DisabilityTypeID").Optional(),
		field.Time("fromTime").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.Time("toTime").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),

		field.String("OrderNumber").Optional(),
		field.Int64("CreatedById").Optional(),
		field.String("CreatedByUserName").Optional(),
		field.Int64("CreatedByEmpId").Optional(),
		field.String("CreatedByDesignation").Optional(),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.Int64("verifiedbyid").Optional(),
		field.String("verifiedbyusername").Optional(),
		field.Int64("verifiedbyEmployeeid").Optional(),
		field.String("verifiedbyDesignation").Optional(),
		field.Time("verifiedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.String("Statuss").Default("active").Optional(),
		field.Int64("deletedbyid").Optional(),
		field.String("deletedbyusername").Optional(),
		field.Int64("deletedbyEmployeeid").Optional(),
		field.String("deletedbyDesignation").Optional(),
		field.Time("deletedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.String("PaperStatus").MaxLen(10).NotEmpty(),
		field.Int32("CalendarCode").Optional(),
		field.Int32("ExamCodePS").Optional(),
		field.String("CreatedByEmployeeId").Optional(),
	}
}

// Edges of the ExamPapers.

func (ExamPapers) Edges() []ent.Edge {

	return []ent.Edge{edge.To("centers", Center.Type),
		//edge.From("exam", Exam.Type).Ref("papers").Unique().Field("ExamCode"),
		edge.To("exampapers_types", PaperTypes.Type),
		edge.To("papers_ref", ExamCalendar.Type),
		edge.To("ExamPaperEligibility", EligibilityMaster.Type),
		//edge.From("disabilitiesreference", Disability.Type).Ref("dis_ref").Unique().Field("DisabilityTypeID"),
		edge.To("dis_ref", Disability.Type),
		//edge.From("exam_papers_ps", Exam_PS.Type).Ref("papers_ps_ref").Unique().Field("ExamCodePS"),
		//edge.To("papers_ps_ref", Exam_PS.Type),
		//edge.To("papers_ip_ref", Exam_IP.Type),
	}

}
