// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

)

type Exam_Applications_GDSPM struct {
	ent.Schema
}

func (Exam_Applications_GDSPM) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("ApplicationID"), 
		field.String("ApplicationNumber").Optional(),
		field.Time("ApplnSubmittedDate").
		SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Default(time.Now).
		Optional(),
		field.Int32("ExamCode").Optional(),
		field.String("ExamShortName").Optional(),
		field.String("ExamName").Optional(),
		field.String("SmsExamShortName").Optional(),
		field.String("ExamYear").Optional(),
		field.Int32("UserID").Optional(),
		field.Int64("EmployeeID").Optional(),
		field.String("EmployeeName").Optional(),
		field.String("DOB").Optional(),
		field.String("Gender").Optional(),
		field.String("MobileNumber").Optional(),
		field.String("EmailID").Optional(),	
		field.String("CategoryCode").Optional(),
		field.String("CategoryDescription").Optional(),
		field.String("DisabilityTypeID").Optional(),
		field.String("DisabilityTypeCode").Optional(),
		field.String("DisabilityTypeDescription").Optional(),
		field.Int32("DisabilityPercentage").Optional(),
		field.String("DCCS").Optional(),
		field.String("EntryPostCode").Optional(),
		field.String("EntryPostDescription").Optional(),
		field.String("PresentPostCode").Optional(),
		field.String("PresentPostDescription").Optional(),
		field.String("FeederPostCode").Optional(),
		field.String("FeederPostDescription").Optional(),
		field.String("FeederPostJoiningDate").Optional(),
		field.String("DesignationID").Optional(),
		field.String("PresentDesignation").Optional(),
		field.String("EducationCode").Optional(),
		field.String("EducationDescription").Optional(),
		field.String("FacilityUniqueID").Optional(),
		field.Int32("WorkingOfficePincode").Optional(),
		field.String("WorkingOfficeFacilityID").Optional(),
		field.String("WorkingOfficeName").Optional(),
		field.String("WorkingOfficeCircleFacilityID").Optional(),
		field.String("WorkingOfficeCircleName").Optional(),
		field.String("WorkingOfficeRegionFacilityID").Optional(),
		field.String("WorkingOfficeRegionName").Optional(),
		field.String("WorkingOfficeDivisionFacilityID").Optional(),
		field.String("WorkingOfficeDivisionName").Optional(),
		field.String("ReportingOfficeFacilityID").Optional(),
		field.String("ReportingOfficeName").Optional(),
		field.String("LienControllingOfficeID").Optional(),
		field.String("LienControllingOfficeName").Optional(),
		field.String("InDeputation").Optional(),
		field.String("DeputationType").Optional(),
		field.String("DeputationOfficeUniqueId").Optional(),
		field.String("DeputationOfficeFacilityID").Optional(),
		field.String("DeputationOfficeName").Optional(),
		field.String("DeputationControllingOfficeID").Optional(),
		field.String("DeputationControllingOfficeName").Optional(),
		field.String("ControllingOfficeFacilityID").Optional(),
		field.String("ControllingOfficeName").Optional(),
		field.String("NodalOfficeFacilityID").Optional(),
		field.String("NodalOfficeName").Optional(),
		field.String("SubdivisionOfficeFacilityID").Optional(),
		field.String("SubdivisionOfficeName").Optional(),		
		field.Int32("ExamCityCenterCode").Optional().Default(0),
		field.String("CenterFacilityId").Optional(),
		field.String("CentrePreference").Optional(),
		field.String("Signature").Optional(),
		field.String("Photo").Optional(),
		field.Bytes("CandidatePhoto").Optional(),
		field.Bytes("CandidateSignature").Optional(),
		field.String("SignaturePath").Optional(),
		field.String("PhotoPath").Optional(),
		field.String("TempHallTicket").Optional(),
		field.String("CandidateRemarks").Optional(),
		field.String("VAGeneralRemarks").Optional(),
		field.String("CAGeneralRemarks").Optional(),
		field.String("NAGeneralRemarks").Optional(),
		field.String("ApplicationStatus").Optional(),
		field.String("Status").Default("active").Optional(),
		field.String("RecommendedStatus").Optional(),
		
		field.Time("ApplicationWdlDate").
		SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).
		Optional(),	
		field.Int32("VA_UserId").Optional(),
		field.String("VA_UserName").Optional(),
		field.String("VA_EmployeeID").Optional(),
		field.String("VA_EmployeeDesignation").Optional(),
		field.String("VA_Remarks").Optional(),
		field.Time("VA_Date").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Int32("CA_UserId").Optional(),
		field.String("CA_UserName").Optional(),
		field.String("CA_EmployeeID").Optional(),
		field.String("CA_EmployeeDesignation").Optional(),
		field.String("CA_Remarks").Optional(),
		field.Time("CA_Date").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Int32("NA_UserId").Optional(),
		field.String("NA_UserName").Optional(),
		field.String("NA_EmployeeID").Optional(),
		field.String("NA_EmployeeDesignation").Optional(),
		field.String("NA_Remarks").Optional(),
		field.Time("NA_Date").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.String("AppliactionRemarks").Optional(),
		field.JSON("CadrePreferences", []interface{}{}).
		Optional(),
		field.String("CAPreviousRemarks").Optional(),
		field.Bool("PunishmentStatus").Optional().Default(false),		//new coloumn
		field.Bool("DisciplinaryCaseStatus").Optional().Default(false),		//new coloumn
		field.Bool("GenerateHallTicketFlag").Optional().Nillable(),//.Default(false),
		field.String("HallTicketNumber").Optional().Default(""),
		field.Bool("HallTicketGeneratedFlag").Optional().Default(false),
		field.Bool("GenerateHallTicketFlagByNO").Optional().Nillable(),
		field.Time("HallTicketGeneratedDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.String("TemporaryHallTicket").Optional(),
		field.String("OptionUsed").Optional(),
		field.String("Remarks").Optional(),
		field.String("Cadre").Optional(),
		field.String("EmployeePost").Optional(),
		field.String("DOJInEligiblePost").Optional(),
		field.JSON("DivisionPreferences", []interface{}{}).
		Optional(),
		field.Int32("RoleUserCode").Optional(),
		field.JSON("ServiceLength", []interface{}{}).
		Optional(),
		field.JSON("NonQualifyingService", []interface{}{}).
		Optional(),
		field.String("DCInPresentCadre").Optional(),
		field.String("ReportingOfficeID").Optional(),
		field.JSON("PostPreferences", []interface{}{}).
		Optional(),
		field.JSON("UnitPreferences", []interface{}{}).
		Optional(),
		field.Int32("CenterId").Optional(),
		field.Int32("CenterCode").Optional(),
		field.String("ClaimingQualifyingService").Optional(),
		field.String("DeputationOfficePincode").Optional(),
		field.Int32("CircleSID").Optional().Nillable(),
		field.String("FacilityName").Optional(),
		field.Time("UpdatedAt").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Default(time.Now).Optional(),
		field.String("UpdatedBy").Default("API").Optional(),
		field.Int32("ExamCenterHall").Optional(),
		field.String("HallName").Optional(),
		field.JSON("GDSEngagement", []interface{}{}).
			Optional(),
		field.JSON("PMMailGuardMTSEngagement", []interface{}{}).
		Optional(),

		}
}

func (Exam_Applications_GDSPM) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("EmployeeID", "ExamYear", "Status").Unique(),
		index.Fields("ApplicationNumber", "ExamYear", "Status").Unique(),
		index.Fields("ControllingOfficeFacilityID", "ExamYear", "Status"),
		index.Fields("NodalOfficeFacilityID", "ExamYear", "Status"),
	}
}

func (Exam_Applications_GDSPM) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("UsersGDSPMRef", UserMaster.Type),

		edge.To("CadrePrefRefGDSPM", Cadre_Choice_PM.Type),
		edge.To("CirclePrefRefGDSPM", Division_Choice_PM.Type),
		edge.To("GDSPMApplicationsRef", RecommendationsGDSPMApplications.Type),
		edge.To("LogData", Logs.Type),
		//edge.From("GDSPMExamCentres", Center.Type). Ref("ExamCentresRefGDSPM").Unique().Field("ExamCityCenterCode"),
		edge.From("Exams", Exam.Type).Ref("ExamMasterRefGDSPM").Unique().Field("ExamCode"),
		edge.From("GDSPMExamCentres", Center.Type). Ref("ExamCentresRefGDSPM").Unique().Field("CenterCode"),
		edge.From("CircleRefsGDSPM", CircleSummaryForNO.Type).Ref("CircleRefsForHallTicketGDSPM").Unique().Field("CircleSID"),
		edge.From("roleusers", RoleMaster.Type).Ref("Roles_GDSPM_Ref").Unique().Field("RoleUserCode"),
		edge.From("examcitycenter", ExamCityCenter.Type).Ref("ExamCityCenterGDSPMRef").Unique().Field("ExamCityCenterCode"),
		//edge.From("ExamCenterHalGDSPMSRefs", ExamCenterHall.Type).Ref("ExamCenterHallGDSPMRef").Unique().Field("ExamCenterHall"),
	}
}

func (Exam_Applications_GDSPM) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Exam_Applications_GDSPM"}}
}
