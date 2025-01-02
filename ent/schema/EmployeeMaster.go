package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type EmployeeMaster struct {
	ent.Schema
}

// Fields of the User.
func (EmployeeMaster) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("EmpID"),
		field.Int64("EmployeeID").Optional().Default(0),
		field.String("EmployeeName").Optional(),
		field.String("DOB").Optional(),
		field.String("Gender").Optional(),
		field.String("MobileNumber").Optional(),
		field.String("EmailID").Optional(),
		field.String("EmployeeCategoryCode").Optional(),
		field.String("EmployeeCategory").Optional(),
		field.String("PostCode").Optional(),
		field.String("EmployeePost").Optional(),
		field.String("FacilityID").Optional(),
		field.String("OfficeName").Optional(),
		field.String("ControllingAuthorityFacilityId").Optional(),
		field.String("ControllingAuthorityName").Optional(),
		field.String("NodalAuthorityFaciliyId").Optional(),
		field.String("NodalAuthorityName").Optional(),
		field.String("Pincode").Optional(),
		field.String("CircleFacilityID").Optional(),
		field.String("Statuss").Default("active").Optional(),
		field.Bool("VerifyStatus").Default(false),
		field.String("UidToken").Optional(),
		field.String("Createdby").Optional(),
		field.String("DCCS").Optional(),
		field.Int64("CreatedById").Optional().Default(0),
		field.String("CreatedByUserName").Optional(),
		field.Int64("CreatedByEmpId").Optional().Default(0),
		field.String("CreatedByDesignation").Optional(),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Int64("ModifiedById").Optional().Default(0),
		field.String("ModifiedByUserName").Optional(),
		field.Int64("ModifiedByEmpId").Optional().Default(0),
		field.String("ModifiedByDesignantion").Optional(),
		field.Time("ModifiedDate").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Int64("DeletedById").Optional().Default(0),
		field.String("DeletedByUserName").Optional(),
		field.Int64("DeletedByEmpId").Optional().Default(0),
		field.String("DeletedByDesignation").Optional(),
		field.Time("DeletedDate").Optional().SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Time("UpdatedAt").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.String("UpdatedBy").Default("API").Optional(),
		field.Int64("SmsOtp").Optional().Default(0),
		field.Time("SmsTriggeredTime").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Bool("SmsVerifyStatus").Default(false),
		field.Int64("EmailOtp").Optional().Default(0),
		field.Time("EmailTriggeredTime").SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
		}).Optional(),
		field.Bool("EmailVerifyStatus").Default(false),
		field.Bool("FinalSubmitStatus").Default(false),
		field.String("DCInPresentCadre").Optional(),
		field.String("Cadre").Optional(),
	}

	//Employee Number,Correct/Incorrect, Remarks
}

func (EmployeeMaster) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("EmployeeID", "Statuss").Unique(),
		index.Fields("ControllingAuthorityFacilityId", "Statuss"),
	}
}

// Edges of the User.
func (EmployeeMaster) Edges() []ent.Edge {
	return []ent.Edge{edge.To("UsermasterRef", UserMaster.Type),
		edge.To("Emp_Ref", Exam_Applications_PS.Type),
		edge.To("LogData", Logs.Type),
	}
}

func (EmployeeMaster) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "EmployeeMaster"}}
}
