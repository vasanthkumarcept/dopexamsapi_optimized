// Code generated by ent, DO NOT EDIT.

package circlesummaryforno

import (
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldLTE(FieldID, id))
}

// CircleOfficeId applies equality check predicate on the "CircleOfficeId" field. It's identical to CircleOfficeIdEQ.
func CircleOfficeId(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldCircleOfficeId, v))
}

// CircleOfficeName applies equality check predicate on the "CircleOfficeName" field. It's identical to CircleOfficeNameEQ.
func CircleOfficeName(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldCircleOfficeName, v))
}

// ApproveHallTicketGenrationIP applies equality check predicate on the "ApproveHallTicketGenrationIP" field. It's identical to ApproveHallTicketGenrationIPEQ.
func ApproveHallTicketGenrationIP(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldApproveHallTicketGenrationIP, v))
}

// ApproveHallTicketGenrationPS applies equality check predicate on the "ApproveHallTicketGenrationPS" field. It's identical to ApproveHallTicketGenrationPSEQ.
func ApproveHallTicketGenrationPS(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldApproveHallTicketGenrationPS, v))
}

// ApproveHallTicketGenrationPM applies equality check predicate on the "ApproveHallTicketGenrationPM" field. It's identical to ApproveHallTicketGenrationPMEQ.
func ApproveHallTicketGenrationPM(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldApproveHallTicketGenrationPM, v))
}

// ApproveHallTicketGenrationPA applies equality check predicate on the "ApproveHallTicketGenrationPA" field. It's identical to ApproveHallTicketGenrationPAEQ.
func ApproveHallTicketGenrationPA(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldApproveHallTicketGenrationPA, v))
}

// CircleOfficeIdEQ applies the EQ predicate on the "CircleOfficeId" field.
func CircleOfficeIdEQ(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldCircleOfficeId, v))
}

// CircleOfficeIdNEQ applies the NEQ predicate on the "CircleOfficeId" field.
func CircleOfficeIdNEQ(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNEQ(FieldCircleOfficeId, v))
}

// CircleOfficeIdIn applies the In predicate on the "CircleOfficeId" field.
func CircleOfficeIdIn(vs ...string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldIn(FieldCircleOfficeId, vs...))
}

// CircleOfficeIdNotIn applies the NotIn predicate on the "CircleOfficeId" field.
func CircleOfficeIdNotIn(vs ...string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNotIn(FieldCircleOfficeId, vs...))
}

// CircleOfficeIdGT applies the GT predicate on the "CircleOfficeId" field.
func CircleOfficeIdGT(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldGT(FieldCircleOfficeId, v))
}

// CircleOfficeIdGTE applies the GTE predicate on the "CircleOfficeId" field.
func CircleOfficeIdGTE(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldGTE(FieldCircleOfficeId, v))
}

// CircleOfficeIdLT applies the LT predicate on the "CircleOfficeId" field.
func CircleOfficeIdLT(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldLT(FieldCircleOfficeId, v))
}

// CircleOfficeIdLTE applies the LTE predicate on the "CircleOfficeId" field.
func CircleOfficeIdLTE(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldLTE(FieldCircleOfficeId, v))
}

// CircleOfficeIdContains applies the Contains predicate on the "CircleOfficeId" field.
func CircleOfficeIdContains(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldContains(FieldCircleOfficeId, v))
}

// CircleOfficeIdHasPrefix applies the HasPrefix predicate on the "CircleOfficeId" field.
func CircleOfficeIdHasPrefix(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldHasPrefix(FieldCircleOfficeId, v))
}

// CircleOfficeIdHasSuffix applies the HasSuffix predicate on the "CircleOfficeId" field.
func CircleOfficeIdHasSuffix(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldHasSuffix(FieldCircleOfficeId, v))
}

// CircleOfficeIdEqualFold applies the EqualFold predicate on the "CircleOfficeId" field.
func CircleOfficeIdEqualFold(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEqualFold(FieldCircleOfficeId, v))
}

// CircleOfficeIdContainsFold applies the ContainsFold predicate on the "CircleOfficeId" field.
func CircleOfficeIdContainsFold(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldContainsFold(FieldCircleOfficeId, v))
}

// CircleOfficeNameEQ applies the EQ predicate on the "CircleOfficeName" field.
func CircleOfficeNameEQ(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldCircleOfficeName, v))
}

// CircleOfficeNameNEQ applies the NEQ predicate on the "CircleOfficeName" field.
func CircleOfficeNameNEQ(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNEQ(FieldCircleOfficeName, v))
}

// CircleOfficeNameIn applies the In predicate on the "CircleOfficeName" field.
func CircleOfficeNameIn(vs ...string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldIn(FieldCircleOfficeName, vs...))
}

// CircleOfficeNameNotIn applies the NotIn predicate on the "CircleOfficeName" field.
func CircleOfficeNameNotIn(vs ...string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNotIn(FieldCircleOfficeName, vs...))
}

// CircleOfficeNameGT applies the GT predicate on the "CircleOfficeName" field.
func CircleOfficeNameGT(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldGT(FieldCircleOfficeName, v))
}

// CircleOfficeNameGTE applies the GTE predicate on the "CircleOfficeName" field.
func CircleOfficeNameGTE(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldGTE(FieldCircleOfficeName, v))
}

// CircleOfficeNameLT applies the LT predicate on the "CircleOfficeName" field.
func CircleOfficeNameLT(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldLT(FieldCircleOfficeName, v))
}

// CircleOfficeNameLTE applies the LTE predicate on the "CircleOfficeName" field.
func CircleOfficeNameLTE(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldLTE(FieldCircleOfficeName, v))
}

// CircleOfficeNameContains applies the Contains predicate on the "CircleOfficeName" field.
func CircleOfficeNameContains(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldContains(FieldCircleOfficeName, v))
}

// CircleOfficeNameHasPrefix applies the HasPrefix predicate on the "CircleOfficeName" field.
func CircleOfficeNameHasPrefix(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldHasPrefix(FieldCircleOfficeName, v))
}

// CircleOfficeNameHasSuffix applies the HasSuffix predicate on the "CircleOfficeName" field.
func CircleOfficeNameHasSuffix(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldHasSuffix(FieldCircleOfficeName, v))
}

// CircleOfficeNameEqualFold applies the EqualFold predicate on the "CircleOfficeName" field.
func CircleOfficeNameEqualFold(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEqualFold(FieldCircleOfficeName, v))
}

// CircleOfficeNameContainsFold applies the ContainsFold predicate on the "CircleOfficeName" field.
func CircleOfficeNameContainsFold(v string) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldContainsFold(FieldCircleOfficeName, v))
}

// ApproveHallTicketGenrationIPEQ applies the EQ predicate on the "ApproveHallTicketGenrationIP" field.
func ApproveHallTicketGenrationIPEQ(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldApproveHallTicketGenrationIP, v))
}

// ApproveHallTicketGenrationIPNEQ applies the NEQ predicate on the "ApproveHallTicketGenrationIP" field.
func ApproveHallTicketGenrationIPNEQ(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNEQ(FieldApproveHallTicketGenrationIP, v))
}

// ApproveHallTicketGenrationIPIsNil applies the IsNil predicate on the "ApproveHallTicketGenrationIP" field.
func ApproveHallTicketGenrationIPIsNil() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldIsNull(FieldApproveHallTicketGenrationIP))
}

// ApproveHallTicketGenrationIPNotNil applies the NotNil predicate on the "ApproveHallTicketGenrationIP" field.
func ApproveHallTicketGenrationIPNotNil() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNotNull(FieldApproveHallTicketGenrationIP))
}

// ApproveHallTicketGenrationPSEQ applies the EQ predicate on the "ApproveHallTicketGenrationPS" field.
func ApproveHallTicketGenrationPSEQ(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldApproveHallTicketGenrationPS, v))
}

// ApproveHallTicketGenrationPSNEQ applies the NEQ predicate on the "ApproveHallTicketGenrationPS" field.
func ApproveHallTicketGenrationPSNEQ(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNEQ(FieldApproveHallTicketGenrationPS, v))
}

// ApproveHallTicketGenrationPSIsNil applies the IsNil predicate on the "ApproveHallTicketGenrationPS" field.
func ApproveHallTicketGenrationPSIsNil() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldIsNull(FieldApproveHallTicketGenrationPS))
}

// ApproveHallTicketGenrationPSNotNil applies the NotNil predicate on the "ApproveHallTicketGenrationPS" field.
func ApproveHallTicketGenrationPSNotNil() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNotNull(FieldApproveHallTicketGenrationPS))
}

// ApproveHallTicketGenrationPMEQ applies the EQ predicate on the "ApproveHallTicketGenrationPM" field.
func ApproveHallTicketGenrationPMEQ(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldApproveHallTicketGenrationPM, v))
}

// ApproveHallTicketGenrationPMNEQ applies the NEQ predicate on the "ApproveHallTicketGenrationPM" field.
func ApproveHallTicketGenrationPMNEQ(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNEQ(FieldApproveHallTicketGenrationPM, v))
}

// ApproveHallTicketGenrationPMIsNil applies the IsNil predicate on the "ApproveHallTicketGenrationPM" field.
func ApproveHallTicketGenrationPMIsNil() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldIsNull(FieldApproveHallTicketGenrationPM))
}

// ApproveHallTicketGenrationPMNotNil applies the NotNil predicate on the "ApproveHallTicketGenrationPM" field.
func ApproveHallTicketGenrationPMNotNil() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNotNull(FieldApproveHallTicketGenrationPM))
}

// ApproveHallTicketGenrationPAEQ applies the EQ predicate on the "ApproveHallTicketGenrationPA" field.
func ApproveHallTicketGenrationPAEQ(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldEQ(FieldApproveHallTicketGenrationPA, v))
}

// ApproveHallTicketGenrationPANEQ applies the NEQ predicate on the "ApproveHallTicketGenrationPA" field.
func ApproveHallTicketGenrationPANEQ(v bool) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNEQ(FieldApproveHallTicketGenrationPA, v))
}

// ApproveHallTicketGenrationPAIsNil applies the IsNil predicate on the "ApproveHallTicketGenrationPA" field.
func ApproveHallTicketGenrationPAIsNil() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldIsNull(FieldApproveHallTicketGenrationPA))
}

// ApproveHallTicketGenrationPANotNil applies the NotNil predicate on the "ApproveHallTicketGenrationPA" field.
func ApproveHallTicketGenrationPANotNil() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(sql.FieldNotNull(FieldApproveHallTicketGenrationPA))
}

// HasCircleusers applies the HasEdge predicate on the "circleusers" edge.
func HasCircleusers() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CircleusersTable, CircleusersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCircleusersWith applies the HasEdge predicate on the "circleusers" edge with a given conditions (other predicates).
func HasCircleusersWith(preds ...predicate.UserMaster) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := newCircleusersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCircleRefsForHallTicketIP applies the HasEdge predicate on the "CircleRefsForHallTicketIP" edge.
func HasCircleRefsForHallTicketIP() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CircleRefsForHallTicketIPTable, CircleRefsForHallTicketIPColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCircleRefsForHallTicketIPWith applies the HasEdge predicate on the "CircleRefsForHallTicketIP" edge with a given conditions (other predicates).
func HasCircleRefsForHallTicketIPWith(preds ...predicate.Exam_Applications_IP) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := newCircleRefsForHallTicketIPStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCircleRefsForHallTicketPS applies the HasEdge predicate on the "CircleRefsForHallTicketPS" edge.
func HasCircleRefsForHallTicketPS() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CircleRefsForHallTicketPSTable, CircleRefsForHallTicketPSColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCircleRefsForHallTicketPSWith applies the HasEdge predicate on the "CircleRefsForHallTicketPS" edge with a given conditions (other predicates).
func HasCircleRefsForHallTicketPSWith(preds ...predicate.Exam_Applications_PS) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := newCircleRefsForHallTicketPSStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCircleRefsForHallTicketGDSPA applies the HasEdge predicate on the "CircleRefsForHallTicketGDSPA" edge.
func HasCircleRefsForHallTicketGDSPA() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CircleRefsForHallTicketGDSPATable, CircleRefsForHallTicketGDSPAColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCircleRefsForHallTicketGDSPAWith applies the HasEdge predicate on the "CircleRefsForHallTicketGDSPA" edge with a given conditions (other predicates).
func HasCircleRefsForHallTicketGDSPAWith(preds ...predicate.Exam_Applications_GDSPA) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := newCircleRefsForHallTicketGDSPAStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCircleRefsForHallTicketGDSPM applies the HasEdge predicate on the "CircleRefsForHallTicketGDSPM" edge.
func HasCircleRefsForHallTicketGDSPM() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CircleRefsForHallTicketGDSPMTable, CircleRefsForHallTicketGDSPMColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCircleRefsForHallTicketGDSPMWith applies the HasEdge predicate on the "CircleRefsForHallTicketGDSPM" edge with a given conditions (other predicates).
func HasCircleRefsForHallTicketGDSPMWith(preds ...predicate.Exam_Applications_GDSPM) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := newCircleRefsForHallTicketGDSPMStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCircleRefsForHallTicketPMPA applies the HasEdge predicate on the "CircleRefsForHallTicketPMPA" edge.
func HasCircleRefsForHallTicketPMPA() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CircleRefsForHallTicketPMPATable, CircleRefsForHallTicketPMPAColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCircleRefsForHallTicketPMPAWith applies the HasEdge predicate on the "CircleRefsForHallTicketPMPA" edge with a given conditions (other predicates).
func HasCircleRefsForHallTicketPMPAWith(preds ...predicate.Exam_Applications_PMPA) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := newCircleRefsForHallTicketPMPAStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCircleRefsForHallTicketMTSPMMG applies the HasEdge predicate on the "CircleRefsForHallTicketMTSPMMG" edge.
func HasCircleRefsForHallTicketMTSPMMG() predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CircleRefsForHallTicketMTSPMMGTable, CircleRefsForHallTicketMTSPMMGColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCircleRefsForHallTicketMTSPMMGWith applies the HasEdge predicate on the "CircleRefsForHallTicketMTSPMMG" edge with a given conditions (other predicates).
func HasCircleRefsForHallTicketMTSPMMGWith(preds ...predicate.Exam_Application_MTSPMMG) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		step := newCircleRefsForHallTicketMTSPMMGStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CircleSummaryForNO) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CircleSummaryForNO) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CircleSummaryForNO) predicate.CircleSummaryForNO {
	return predicate.CircleSummaryForNO(func(s *sql.Selector) {
		p(s.Not())
	})
}