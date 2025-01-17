// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamEligibility holds the schema definition for the ExamEligibility entity.
type EligibilityMaster struct {
	ent.Schema
}

// Fields of the ExamEligibility.
func (EligibilityMaster) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("EligibilityCode"),
		 field.Int32("ExamCode").Optional(), 
		 field.String("ExamName"),
		 field.String("PostCode").Optional(),
		 field.Bool("gdsService").Default(false),
		 field.Int32("AgeCriteria").Optional(), 
		 field.Int32("ServiceCriteria").Optional(),
		 field.Bool("DrivingLicenseCriteria").Default(false),
		 field.Bool("ComputerKnowledge").Default(false),
		 field.Bool("LevelOfPayMatrixEligibility").Default(false),
		 field.String("Education"),
		// field.String("OrderNumber").Optional(),

		 field.Int32("NotifyCode").Optional(),
		 field.String("CategoryCode").Optional(),
		 field.Int32("PaperCode").Optional(),
		 field.String("PaperDescription"),
		 field.Int32("MinimumMarks"),

		 field.String("OrderNumber").Optional(),
		 field.String("Status").Optional(),
		 field.Int32("CreatedById").Optional(),
		  field.String("CreatedByUserName").Optional(),
			 field.String("CreatedByEmployeeId").Optional(),
			 field.String("CreatedByDesignation").Optional(),
			 field.Time("CreatedDate").Optional(),
			 field.Int64("VerifiedById").Optional(),
	 field.String("VerifiedByUserName").Optional(),
	 field.String("VerifiedByEmployeeId").Optional(),
	 field.String("VerifiedByDesignation").Optional(),
	 field.Time("VerifiedDate").Optional(),
	 field.Int64("DeletedById").Optional(),
	 field.String("DeletedByUserName").Optional(),
	 field.String("DeletedByEmployeeId").Optional(),
	 field.String("DeletedByDesignation").Optional(),
	 field.Time("DeletedDate").Optional(),
	 
		}

	// Edges of the ExamEligibility.		
	
}

func (EligibilityMaster) Edges() []ent.Edge {
	return []ent.Edge{//edge.To("ExamEligibility", Exam.Type),
	edge.To("Notifications", Notification.Type),
	edge.To("CategoryEligibility", EmployeeCategory.Type),
	edge.To("PostEligibility", EmployeePosts.Type),
	//edge.To("ExamPaperEligibility", ExamPapers.Type),
	edge.From("ExamPaper_Eligibility", ExamPapers.Type).Ref("ExamPaperEligibility").Unique().Field("PaperCode"),
	edge.From("Exam_Eligibility", Exam.Type).Ref("ExamEligibility").Unique().Field("ExamCode"),}
}
func (EligibilityMaster) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "EligibilityMaster"}}
}
