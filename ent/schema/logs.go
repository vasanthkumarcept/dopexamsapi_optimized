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
type Logs struct {
	ent.Schema
}

// Fields of the Log.
func (Logs) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("ID"),
		field.String("userid").
			Optional(),
		field.Int64("uniqueid").
			Optional(),
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
			Optional(),
		field.Float("longitude").
			Optional(),
		field.Time("eventtime").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Default(time.Now),
	}
}

// Edges of the Log.
func (Logs) Edges() []ent.Edge {
	// return []ent.Edge{
	// 	// edge.From("adminmaster", AdminMaster.Type).
	// 	// 	Ref("admin_create").
	// 	// 	Unique().
	// 	// 	Field("EmployeeId"), // Changed Field name to "EmployeeId"

	// }
	return nil
}

func (Logs) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "LogsTable"}}
}
