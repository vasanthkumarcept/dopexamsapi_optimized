package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type PaperTypes struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (PaperTypes) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("PaperTypeCode"),

		field.String("PaperTypeDescription").MaxLen(100).NotEmpty(),
		field.Int32("SequenceNumber").Optional(),
		field.String("OrderNumber").Optional(),
		field.Int64("CreatedById").Optional(),
		field.String("CreatedByUserName").Optional(),
		field.String("CreatedByEmployeeId").Optional(),
		field.String("CreatedByDesignation").Optional(),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.Int64("VerifiedById").Optional(),
		field.String("VerifiedByUserName").Optional(),
		field.String("VerifiedByEmployeeId").Optional(),
		field.String("VerifiedByDesignation").Optional(),
		field.Time("VerifiedDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.String("Status"),
		field.Int64("DeletedById").Optional(),
		field.String("DeletedByUserName").Optional(),
		field.String("DeletedByEmployeeId").Optional(),
		field.String("DeletedByDesignation").Optional(),
		field.Time("DeletedDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),

		field.Int32("PaperCode").Optional(),
	}
}

// Edges of the ExamPapers.
func (PaperTypes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("papercode", ExamPapers.Type).Ref("exampapers_types").Unique().Field("PaperCode"),
	}
}

func (PaperTypes) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "PaperTypes"}}
}
