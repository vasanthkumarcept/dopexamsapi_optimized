package schema

import (
	"entgo.io/ent"
	//"entgo.io/ent/dialect"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// FacilityNewTab holds the schema definition for the Facility entity.
type FacilityMasters struct {
	ent.Schema
}

// Fields of the Facility.
func (FacilityMasters) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("uniqueid").Max(10000),
		// Add other fields of Facility entity here.
		field.String("FacilityID").Optional(),
		field.Int64("UUID").Optional(),

		field.String("FacilityType").Optional(),
		field.String("FacilityIDDescription").Optional(),
		field.String("ReportingOfficeFacilityID").Optional(),
		field.String("ReportingOfficeFacilityName").Optional(),
		field.String("HOFacilityID").Optional(),
		field.String("HOFacilityName").Optional(),
		field.String("SubDivisionFacilityID").Optional(),
		field.String("SubDivisionFacilityName").Optional(),
		field.String("DivisionFacilityID").Optional(),
		field.String("DivisionFacilityName").Optional(),
		field.String("RegionFacilityID").Optional(),
		field.String("RegionFacilityName").Optional(),

		field.String("CircleFacilityID").Optional(),
		field.String("CircleFacilityName").Optional(),
		field.String("Pincode").Optional(),

		field.String("ControllingAuthorityFacilityID").Optional(),
		field.String("ControllingAuthorityFacilityName").Optional(),

		field.String("NodalOfficerFacilityID").Optional(),
		field.String("NodalOfficerFacilityName").Optional(),
		field.String("CityName").Optional(),
		field.String("HallCircleCode").Optional(),
		field.String("DeliveryNonDeliveryOffice").Optional(),
		field.String("CreatedID").Optional(),
		field.String("CreatedBy").Optional(),
		field.String("CreatedByName").Optional(),
		field.String("CreatedByEmpID").Optional(),
		field.String("CircleCode").Optional(),
		field.String("Status").Default("Active"),
		field.Time("EventTime").
		SchemaType(map[string]string{
			dialect.Postgres: "timestamp",
			}).
		Optional(),

		

		// field.Time("ApplicationWdlDate").
		// SchemaType(map[string]string{
		// 	dialect.Postgres: "timestamp",
		// }).
		// Optional(),	
	}
}

// Edges of the Facility.
func (FacilityMasters) Edges() []ent.Edge {

	return []ent.Edge{
		// edge.From("divisions", DivisionMaster.Type).Ref("divisions_ref").Unique().Field("DivisionID"),
		// edge.From("regions", RegionMaster.Type).Ref("region_ref_ref").Unique().Field("RegionID"),
		// edge.From("circles", CircleMaster.Type).Ref("circle_ref").Unique().Field("CircleID"),
		// //edge.To("Office_PS_Ref", DivisionMaster.Type),
		// //edge.To("region_ref", RegionMaster.Type),
		// edge.To("circle_ref", CircleMaster.Type),
		// edge.To("Office_PS_Ref", Exam_Applications_PS.Type),
		// edge.To("Office_IP_Ref", Exam_Applications_IP.Type),
		// //edge.To("Office_IP1_Ref", Exam_Applications_IP1.Type),
		// //edge.To("OfficeCentres", Center.Type),
		// //emp_facility
		// edge.To("Office_GDSPA_Ref", Exam_Applications_GDSPA.Type),
		// edge.To("Office_PMPA_Ref", Exam_Applications_PMPA.Type),
		// edge.To("Office_GDSPM_Ref", Exam_Applications_GDSPM.Type),
	}
}
func (FacilityMasters) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "FacilityMasters"}}
}
