// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	//"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/dialect"

)

type ExamCategoryDisabilityMapping struct {
	ent.Schema
}

func (ExamCategoryDisabilityMapping) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("UniqueID"), 
		field.Int64("ExamCode").Optional().Default(0),
		field.String("ExamShortName").Optional(),
		field.String("CategoryDisability").Optional(),
		field.String("CategoryDisabilityCode").Optional(),
		field.String("CategoryDisabilityDescription").Optional(),
		field.Int32("AgeException").Optional().Default(0),
		field.Int32("ServiceException").Optional().Default(0),
		field.Bool("DrivingLicense").Optional(),
		field.String("OrderNumber").Optional(),
		field.String("Status").Optional(),
		field.Int32("CreatedById").Optional().Default(0),
		 field.String("CreatedByUserName").Optional(),
			field.String("CreatedByEmployeeId").Optional(),
			field.String("CreatedByDesignation").Optional(),
			field.Time("CreatedDate").SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).Optional(),
			field.Int64("VerifiedById").Optional().Default(0),
	field.String("VerifiedByUserName").Optional(),
	field.String("VerifiedByEmployeeId").Optional(),
	field.String("VerifiedByDesignation").Optional(),
	field.Time("VerifiedDate").SchemaType(map[string]string{
		dialect.Postgres: "timestamp",
	}).Optional(),
	field.Int64("DeletedById").Optional().Default(0),
	field.String("DeletedByUserName").Optional(),
	field.String("DeletedByEmployeeId").Optional(),
	field.String("DeletedByDesignation").Optional(),
	field.Time("DeletedDate").SchemaType(map[string]string{
		dialect.Postgres: "timestamp",
	}).Optional(),
	
	}
}
func (ExamCategoryDisabilityMapping) Edges() []ent.Edge {
	//edge.To("emp_designation","Employees.Type")
	return nil
}
func (ExamCategoryDisabilityMapping) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "ExamCategoryDisabilityMapping"}}
}
