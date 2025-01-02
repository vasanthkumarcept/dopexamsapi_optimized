package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type ExamPostMapping struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (ExamPostMapping) Fields() []ent.Field {
	return []ent.Field{field.Int("id").StorageKey("UniqueID"),

		field.Int64("ExamCode").Optional(),
		field.String("ExamShortDescription").Optional(),

		//	field.String("ExamType").Optional(),
		//field.Int("ExamCode").Optional(),
		//	field.String("ExamName").Optional(),
		//	field.String("ExamShortName").Optional(),
		field.Int("PostType").Optional(),
		field.String("PostTypeDescription").Optional(),
		field.String("FromPostCode").Optional(),        //employepost referring
		field.String("FromPostDescription").Optional(), //employepost referring
		field.String("ToPostCode").Optional(),          //employepost referring
		field.String("ToPostDescription").Optional(),   //employepost referring
		field.Int("AgeCriteria").Optional(),
		field.Int("ServiceCriteria").Optional(),
		field.Int32("EducationCode").Optional(),
		field.String("EducationDescription").Optional(),

		//field.String("PostCode").Optional(),
		//field.String("PostDescription"),
		field.String("OrderNumber").Optional(),
		field.Int64("CreatedById").Optional(),
		field.String("CreatedByUserName").Optional(),
		field.String("CreatedByEmployeeId").Optional(),
		field.String("CreatedByDesignation").Optional(),
		field.String("CreatedDate").Optional(),
		field.Int64("VerifiedById").Optional(),
		field.String("VerifiedByUserName").Optional(),
		field.String("VerifiedByEmployeeId").Optional(),
		field.String("VerifiedByDesignation").Optional(),
		field.String("VerifiedDate").Optional(),
		field.String("Status").Optional(),
		field.Int64("DeletedById").Optional(),
		field.String("DeletedByUserName").Optional(),
		field.String("DeletedByEmployeeId").Optional(),
		field.String("DeletedByDesignation").Optional(),
		field.String("DeletedDate").Optional(),
	}
}

// Edges of the ExamPostMapping

func (ExamPostMapping) Edges() []ent.Edge {

	return nil
}

func (ExamPostMapping) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "ExamPostMapping"}}
}
