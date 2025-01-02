package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Log holds the schema definition for the Log entity.
type ServiceRequest struct {
	ent.Schema
}

// Fields of the Log.
func (ServiceRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("ID"),
		field.String("remarks").
			Optional(),
		field.String("action").
			Optional(),
		field.Time("PushedTime").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
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
func (ServiceRequest) Edges() []ent.Edge {
	// return []ent.Edge{
	// 	// edge.From("adminmaster", AdminMaster.Type).
	// 	// 	Ref("admin_create").
	// 	// 	Unique().
	// 	// 	Field("EmployeeId"), // Changed Field name to "EmployeeId"

	// }
	return nil
}

func (ServiceRequest) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "ServiceRequest"}}
}
