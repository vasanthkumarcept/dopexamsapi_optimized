package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ExamNotifications struct {
	ent.Schema
}

func (ExamNotifications) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("NotificationCode"),
		//field.Int32("ExamID").Optional(),
		field.String("UserName").Optional(),
		field.Int32("ExamYear").Optional(),
		field.Time("EmployeeMasterRequestLastDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("EmployeeMasterRequestApprovalLastDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("ExamRegisterLastDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("ApplicationStartDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("ApplicationEndDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("ApplicationCorrectionStartDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("ApplicationCorrectionLastDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("ApplicationVerificationLastDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("CenterAllotmentEndDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("NodalOfficerApprovalDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),

		//field.Time("AdmitCardDate").SchemaType(map[string]string{
		//	dialect.Postgres: "date",
		//}).Optional(),
		field.Time("AdmitCardDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("UpdatedAt").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).Optional(),
		field.String("UpdatedBy").Optional(),
		field.JSON("CrucialDate", []interface{}{}).
			Optional(),
		field.String("Designation").Optional(),
		field.String("OfficerName").Optional(),
		field.String("NotificationOrderNumber").Optional(),
		field.String("NotesheetScannedCopy").Optional(),
		field.String("NotificationNumber").Optional(),
		field.Bool("Flag").Default(false).Optional(),
		field.String("ExamShortName").Optional(),
		field.String("CircleOfficeFacilityId").Optional(),
		field.String("CircleOfficeName").Optional(),
		field.String("IssuedBy").Optional(),
		// field.Time("OrderDate").SchemaType(map[string]string{
		// 	dialect.Postgres: "date",
		// }).Optional(),
		field.Time("OrderDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),

		field.String("CreatedBy").Optional(),
		field.Int64("CreatedById").Optional(),
		field.String("CreatedByName").Optional(),
		field.String("CreatedByDesignation").Optional(),
		field.String("ApprovedBy").Optional(),
		field.Int64("ApprovedById").Optional(),
		field.String("ApprovedByName").Optional(),
		field.String("ApprovedByDesignation").Optional(),
		// field.Time("ResubmittedApplicationVerificationDate").Optional().SchemaType(map[string]string{
		// 	dialect.Postgres: "date",
		// }).Optional(),
		field.Time("ResubmittedApplicationVerificationDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),

		field.JSON("Papers", []interface{}{}).
			Optional(),
		field.String("NotificationStatus").Optional(),
		field.String("Status").Optional(),

		//field.String("NotificationIssuedStatus").Optional(),
		field.Bool("NotificationReIssueStatus").Optional().Default(false),
		field.Bool("EditFlagStatus").Default(false).Optional(),
		field.Int32("ExamCode").Optional(),
		field.String("ExamName").Optional(),
		field.Int64("UserID").Optional(),
		field.String("NotificationRemarks").Optional(),
		field.String("SmsExamShortName").Optional(),
	}
}

func (ExamNotifications) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("ExamNotificationref", Exam.Type).Ref("Exam_Notifications_Ref").Unique().Field("ExamCode"),
		edge.From("UserIDref", UserMaster.Type).Ref("User_ID_Ref").Unique().Field("UserID"),
		edge.To("LogData", Logs.Type),
	}
}

func (ExamNotifications) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "ExamNotifications"}}
}
