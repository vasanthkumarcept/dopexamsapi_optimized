package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"

	//`"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Employees struct {
	ent.Schema
}

// Fields of the User.
func (Employees) Fields() []ent.Field {

	return []ent.Field{
		field.Int32("id").StorageKey("RegistrationID"),

		field.Int32("EmployeedID"),
		field.Bool("IDVerified").Default(false),
		field.Bool("IDRemStatus").Default(false),
		field.String("IDRemarks").Optional(),

		field.String("EmployeeName"),
		field.Bool("nameVerified").Default(false),
		field.Bool("nameRemStatus").Default(false),
		field.String("nameRemarks").Optional(),

		field.String("EmployeeFathersName"),
		field.Bool("FathersNameVerified").Default(false),
		field.Bool("FathersNameRemStatus").Default(false),
		field.String("FathersNameRemarks").Optional(),

		field.Time("DOB").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Bool("DOBVerified").Default(false),
		field.Bool("DOBRemStatus").Default(false),
		field.String("DOBRemarks").Optional(),

		field.Enum("Gender").Values("Male", "Female"),
		field.Bool("genderVerified").Default(false),
		field.Bool("genderRemStatus").Default(false),
		field.String("genderRemarks").Optional(),

		field.Int64("MobileNumber").Optional(),
		field.Bool("MobileNumberVerified").Default(false),
		field.Bool("MobileNumberRemStatus").Default(false),
		field.String("MobileNumberRemarks").Optional(),

		field.String("EmailID").Optional(),
		field.Bool("EmailIDVerified").Default(false),
		field.Bool("EmailIDRemStatus").Default(false),
		field.String("EmailIDRemarks").Optional(),

		field.Int32("Categoryid").Optional(),
		field.String("EmployeeCategoryCode").Optional(),
		field.String("EmployeeCategory"),
		field.Bool("EmployeeCategoryCodeVerified").Default(false),
		field.Bool("EmployeeCategoryCodeRemStatus").Default(false),
		field.String("EmployeeCategoryCodeRemarks").Optional(),

		field.String("WithDisability"),
		field.Bool("WithDisabilityVerified").Default(false),
		field.Bool("WithDisabilityRemStatus").Default(false),
		field.Bool("WithDisabilityRemarks").Optional(),

		field.String("DisabilityType").Optional(),
		field.Bool("DisabilityTypeVerified").Default(false),
		field.Bool("DisabilityTypeRemStatus").Default(false),
		field.String("DisabilityTypeRemarks").Optional(),

		field.Int32("DisabilityPercentage").Optional(),
		field.Bool("DisabilityPercentageVerified").Default(false),
		field.Bool("DisabilityPercentageRemStatus").Default(false),
		field.String("DisabilityPercentageRemarks").Optional(),

		field.String("Signature"),
		field.Bool("SignatureVerified").Default(false),
		field.Bool("SignatureRemStatus").Default(false),
		field.String("SignatureRemarks").Optional(),

		field.String("Photo"),
		field.Bool("PhotoVerified").Default(false),
		field.Bool("PhotoRemStatus").Default(false),
		field.String("PhotoRemarks").Optional(),

		/*field.Int32("Cadreid").Optional(),
		field.String("EmployeeCadre"),
		field.Bool("EmployeeCadreVerified").Default(false),
		field.Bool("EmployeeCadreRemStatus").Default(false),
		field.String("EmployeeCadreRemarks").Optional(),*/

		field.Int32("PostID").Optional(),
		field.String("PostCode").Optional(),
		field.String("EmployeePost"),
		field.Bool("EmployeePostVerified").Default(false),
		field.Bool("EmployeePostRemStatus").Default(false),
		field.String("EmployeePostRemarks").Optional(),

		field.Int32("DesignationID").Optional(),
		field.String("EmployeeDesignation"),
		field.Bool("EmployeeDesignationVerified").Default(false),
		field.Bool("EmployeeDesignationRemStatus").Default(false),
		field.String("EmployeeDesignationRemarks").Optional(),

		field.Int32("CircleID").Optional(),
		field.String("CircleName"),
		field.Bool("CircleVerified").Default(false),
		field.Bool("CircleRemStatus").Default(false),
		field.String("CircleRemarks").Optional(),

		field.Int32("RegionID").Optional(),
		field.String("RegionName").Optional(),
		field.Bool("RegionVerified").Default(false),
		field.Bool("RegionRemStatus").Default(false),
		field.String("RegionRemarks").Optional(),

		field.Int32("DivisionID").Optional(),
		field.String("DivisionName").Optional(),
		field.Bool("DivisionVerified").Default(false),
		field.Bool("DivisionRemStatus").Default(false),
		field.String("DivisionRemarks"),

		field.Int32("OfficeID").Optional(),
		field.String("OfficeName"),
		field.Bool("OfficeVerified").Default(false),
		field.Bool("OfficeRemStatus").Default(false),
		field.String("OfficeRemarks").Optional(),

		field.String("Role"),
		field.Bool("RoleVerified").Default(false),
		field.Bool("RoleRemStatus").Default(false),
		field.String("RoleRemarks"),

		field.Time("DCCS").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Bool("DCCSVerified").Default(false),
		field.Bool("DCCSRemStatus").Default(false),
		field.String("DCCSRemarks").Optional(),

		field.Time("DCInPresentCadre").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Bool("DCInPresentCadreVerified").Default(false),
		field.Bool("DCInPresentCadreRemStatus").Default(false),
		field.String("DCInPresentCadreRemarks").Optional(),

		field.Bool("APSWorking").Optional(),
		field.Bool("APSWorkingVerified").Default(false),
		field.Bool("APSWorkingRemStatus").Default(false),
		field.String("APSWorkingRemarks").Optional(),

		field.Bool("profilestatus").Default(false),
		field.Int32("RoleUserCode").Optional(),
		//field.Int32("RegistrationID").Unique(),
	}

	//Employee Number,Correct/Incorrect, Remarks
}

// Edges of the User.
func (Employees) Edges() []ent.Edge {
	return nil
	//dedge.To("employee_user", DirectorateUsers.Type),
	//edge.To("emp_login", Login.Type),
}

func (Employees) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Employees"}}
}
