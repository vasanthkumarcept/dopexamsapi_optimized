// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)
 
	type CircleSummaryForNO struct {
		ent.Schema
	}

	func (CircleSummaryForNO) Fields() []ent.Field {
		return []ent.Field{
			field.Int32("id").StorageKey("CircleSID"), 
			//field.Int32("CircleCode").Unique(),
			field.String("CircleOfficeId"),
			field.String("CircleOfficeName"), 
			field.Bool("ApproveHallTicketGenrationIP").Optional(),
			field.Bool("ApproveHallTicketGenrationPS").Optional(),
			field.Bool("ApproveHallTicketGenrationPM").Optional(),
			field.Bool("ApproveHallTicketGenrationPA").Optional(),
		}
	}
	func (CircleSummaryForNO) Edges() []ent.Edge {
		return []ent.Edge{
		//edge.To("region_ref",RegionMaster.Type),
		//edge.To("circle_ref", Facility.Type),
		//edge.To("CircleExamCentres", Center.Type),
		//edge.To("emp_Region", "RegionMaster.Type"),
		//edge.To("emp_Division"), "DivisionMaster.Type"),
		//edge.To("emp_facility", Facility.Type),
		edge.To("circleusers", UserMaster.Type),	
		edge.To("CircleRefsForHallTicketIP", Exam_Applications_IP.Type),
		edge.To("CircleRefsForHallTicketPS", Exam_Applications_PS.Type),
		edge.To("CircleRefsForHallTicketGDSPA", Exam_Applications_GDSPA.Type),
		edge.To("CircleRefsForHallTicketGDSPM", Exam_Applications_GDSPM.Type),
		edge.To("CircleRefsForHallTicketPMPA", Exam_Applications_PMPA.Type),
		edge.To("CircleRefsForHallTicketMTSPMMG", Exam_Application_MTSPMMG.Type),
	} 
	}
	func (CircleSummaryForNO) Annotations() []schema.Annotation {
		return []schema.Annotation{entsql.Annotation{Table: "CircleSummaryForNO"}}
	}
