package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Log holds the schema definition for the Log entity.
type ErrorLogs struct {
	ent.Schema
}

// Fields of the Log.
func (ErrorLogs) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("ID"),
		field.String("userid").
			Optional(),
		field.Int64("uniqueid").
			Optional().Default(0),
		field.String("usertype").
			Optional(),
		field.String("userdetails").
			Optional(),
		field.String("remarks").
			Optional(),
		field.String("action").
			Optional(),
		field.String("ipaddress").
			Optional(),
		field.String("devicetype").
			Optional(),
		field.String("os").
			Optional(),
		field.String("browser").
			Optional(),
		field.Float("latitude").
			Optional().Default(0),
		field.Float("longitude").
			Optional().Default(0),
		field.Time("eventtime").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Default(time.Now),
		field.String("UpdatedBy").
			Optional(),
		field.Time("UpdatedTime").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.String("AssignedTo").
			Optional(),
		field.String("RemarksNew").
			Optional(),
		field.String("Status").
			Optional(),
		field.Time("ClosedOn").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
	}
}

// Edges of the Log.
func (ErrorLogs) Edges() []ent.Edge {
	// return []ent.Edge{
	// 	// edge.From("adminmaster", AdminMaster.Type).
	// 	// 	Ref("admin_create").
	// 	// 	Unique().
	// 	// 	Field("EmployeeId"), // Changed Field name to "EmployeeId"

	// }
	return nil
}

func (ErrorLogs) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "ErrorLogsTable"}}
}
