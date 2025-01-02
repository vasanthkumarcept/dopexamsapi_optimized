package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Log holds the schema definition for the Log entity.
type SmsEmailLog struct {
	ent.Schema
}

// Fields of the Log.
func (SmsEmailLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("UniqueID"),
		field.String("Type").
			Optional(),
		field.String("MobileEmail").
			Optional(),
		field.String("UserName").
			Optional(),
		field.String("EventCode").
			Optional(),
		field.String("EventDescription").
			Optional(),
		field.String("ApiResponse").
			Optional(),
		field.String("ApiResponseDescription").
			Optional(),
		field.Time("eventtime").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
	}
}

// Edges of the Log.
func (SmsEmailLog) Edges() []ent.Edge {
	// return []ent.Edge{
	// 	// edge.From("adminmaster", AdminMaster.Type).
	// 	// 	Ref("admin_create").
	// 	// 	Unique().
	// 	// 	Field("EmployeeId"), // Changed Field name to "EmployeeId"

	// }
	return nil
}

func (SmsEmailLog) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "SmsEmailLogTable"}}
}
