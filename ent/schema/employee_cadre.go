// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	//"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type EmployeeCadre struct {
	ent.Schema
}

func (EmployeeCadre) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("cadreid"),
	 field.String("cadrecode"), field.String("cadredescription"),
	field.String("PayLevel"), field.String("Scale"),
    field.Int32("ExamconfigurationExamcode").Optional(),
	 field.String("ExamShortDescription").Optional(),
	 field.String("ExamLongDescription").Optional(),
	 field.Int32("EmployeePost_postId").Optional(),
	 field.Int32("EmployeeGroup_groupId").Optional(),
	 field.String("GroupDescription").Optional(),
	 field.String("PostCode").Optional(),
	 field.String("PostDescription").Optional(),
	 field.Int32("BaseCadre").Optional(),
	 field.Int32("GdsService").Optional(),
	 field.Int32("ageCriteria").Optional(),
	 field.Int32("ServiceCriteria").Optional(),
	 field.Int32("DrivingLicenceCriteria").Optional(),
	 field.Int32("ComputerKnowledge").Optional(),
	 field.Int32("EligibiltyBasedOnLevelOfPaymatrix").Optional(),
	 field.Int32("EducationDetails_educationCode").Optional(),
	 field.String("EducationDescription").Optional(),
	 field.String("OrderNumber").Optional(),
	 //field.String("OrderNumber").Optional(),
	 field.String("Status").Optional(),
	 field.Int32("CreatedById").Optional(),
	  field.String("CreatedByUserName").Optional(),
		 field.String("CreatedByEmployeeId").Optional(),
		 field.String("CreatedByDesignation").Optional(),
		 field.Time("CreatedDate").Optional(),
		 field.Int64("VerifiedById").Optional(),
 field.String("VerifiedByUserName").Optional(),
 field.String("VerifiedByEmployeeId").Optional(),
 field.String("VerifiedByDesignation").Optional(),
 field.Time("VerifiedDate").Optional(),
 field.Int64("DeletedById").Optional(),
 field.String("DeletedByUserName").Optional(),
 field.String("DeletedByEmployeeId").Optional(),
 field.String("DeletedByDesignation").Optional(),
 field.Time("DeletedDate").Optional(),
 
}
}
func (EmployeeCadre) Edges() []ent.Edge {
	return nil
	// []ent.Edge{
		//edge.To("posts_ref", Employees.Type),}
}
func (EmployeeCadre) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "EmployeeCadre"}}
}