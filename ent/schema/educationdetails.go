// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	//"time"
)

type EducationDetails struct {
	ent.Schema
}


func (EducationDetails) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("educationCode"),
		field.String("educationDescription").Optional(),
		
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
		
		}
}
func (EducationDetails) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("LogData", Logs.Type),
	}
	
}


func (EducationDetails) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "EducationDetails"}}
}

