package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {

	return []ent.Field{

		//field.Int32("id").StorageKey("NodalOfficerCode")

		field.String("EmployeedID").Unique(),
		field.Bool("IDVerified").Default(false),
		field.Bool("IDRemStatus").Default(false),
		field.String("IDRemarks"),

		field.String("EmployeedName"),
		field.Bool("nameVerified").Default(false),
		field.Bool("nameRemStatus").Default(false),
		field.String("nameRemarks"),

		field.Time("DOB").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Bool("DOBVerified").Default(false),
		field.Bool("DOBRemStatus").Default(false),
		field.String("DOBRemarks"),

		field.Enum("Gender").Values("Male", "Female"),
		field.Bool("genderVerified").Default(false),
		field.Bool("genderRemStatus").Default(false),
		field.String("genderRemarks"),

		field.Int32("Cadreid"),
		field.Bool("cadreidVerified").Default(false),
		field.Bool("cadreidRemStatus").Default(false),
		field.String("cadreidRemarks"),

		field.Int32("OfficeID"),
		field.Bool("officeIDVerified").Default(false),
		field.Bool("officeIDRemStatus").Default(false),
		field.String("officeIDRemarks"),

		field.Bool("PH"),
		field.Bool("PHVerified").Default(false),
		field.Bool("PHRemStatus").Default(false),
		field.String("PHRemarks"),

		field.String("PHDetails").Optional(),
		field.Bool("PHDetailsVerified").Default(false),
		field.Bool("PHDetailsRemStatus").Default(false),
		field.String("PHDetailsRemarks"),

		field.Bool("APSWorking"),
		field.Bool("APSWorkingVerified").Default(false),
		field.Bool("APSWorkingRemStatus").Default(false),
		field.String("APSWorkingRemarks"),

		field.Bool("profilestatus").Default(false),
	}

	//Employee Number,Correct/Incorrect, Remarks
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
