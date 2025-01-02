package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Notification struct {
	ent.Schema
}

func (Notification) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("NotifyCode"),
		field.Int32("ExamCode").Optional(), field.Int32("ExamYear"),
		field.Time("ApplicationStartDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Time("ApplicationEndDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Time("VerificationDateByController").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Time("CorrectionDateByCandidate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Time("CorrectionVeriyDateByController").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Time("HallTicketAllotmentDateByNodalOfficer").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Time("HallTicketDownloadDate").Optional().SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.String("NotifyFile").Optional(),
		field.String("SyllabusFile").Optional(),
		field.String("VacanciesFile").Optional(),
		field.Int32("ExamCodePS").Optional(),
		field.Time("CenterAllotmentEndDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional()}
}
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{edge.To("applications", Application.Type),
		//edge.To("centers", Center.Type), edge.To("nodal_officers", NodalOfficer.Type),
		edge.From("exam", Exam.Type).Ref("notifications").Unique().Field("ExamCode"),
		edge.To("vacancy_years", VacancyYear.Type),
		edge.To("notify_ref", Notification.Type),
		//edge.From("notificationsPS", Exam_PS.Type).Ref("notifications_ps").Unique().Field("ExamCodePS"),
		//edge.To("notifications_ps", Exam_PS.Type),
		//edge.To("notifications_ip", Exam_IP.Type),
		edge.To("LogData", Logs.Type),
	}

}
func (Notification) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Notification"}}
}
