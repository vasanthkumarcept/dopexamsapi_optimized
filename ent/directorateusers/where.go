// Code generated by ent, DO NOT EDIT.

package directorateusers

import (
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLTE(FieldID, id))
}

// Role applies equality check predicate on the "Role" field. It's identical to RoleEQ.
func Role(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldRole, v))
}

// EmployeedID applies equality check predicate on the "EmployeedID" field. It's identical to EmployeedIDEQ.
func EmployeedID(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldEmployeedID, v))
}

// EmployeeName applies equality check predicate on the "EmployeeName" field. It's identical to EmployeeNameEQ.
func EmployeeName(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldEmployeeName, v))
}

// EmailId applies equality check predicate on the "EmailId" field. It's identical to EmailIdEQ.
func EmailId(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldEmailId, v))
}

// MobileNumber applies equality check predicate on the "MobileNumber" field. It's identical to MobileNumberEQ.
func MobileNumber(v int64) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldMobileNumber, v))
}

// SequenceNumber applies equality check predicate on the "SequenceNumber" field. It's identical to SequenceNumberEQ.
func SequenceNumber(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldSequenceNumber, v))
}

// Status applies equality check predicate on the "Status" field. It's identical to StatusEQ.
func Status(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldStatus, v))
}

// RoleEQ applies the EQ predicate on the "Role" field.
func RoleEQ(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "Role" field.
func RoleNEQ(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "Role" field.
func RoleIn(vs ...string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "Role" field.
func RoleNotIn(vs ...string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "Role" field.
func RoleGT(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "Role" field.
func RoleGTE(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "Role" field.
func RoleLT(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "Role" field.
func RoleLTE(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLTE(FieldRole, v))
}

// RoleContains applies the Contains predicate on the "Role" field.
func RoleContains(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldContains(FieldRole, v))
}

// RoleHasPrefix applies the HasPrefix predicate on the "Role" field.
func RoleHasPrefix(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldHasPrefix(FieldRole, v))
}

// RoleHasSuffix applies the HasSuffix predicate on the "Role" field.
func RoleHasSuffix(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldHasSuffix(FieldRole, v))
}

// RoleEqualFold applies the EqualFold predicate on the "Role" field.
func RoleEqualFold(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEqualFold(FieldRole, v))
}

// RoleContainsFold applies the ContainsFold predicate on the "Role" field.
func RoleContainsFold(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldContainsFold(FieldRole, v))
}

// EmployeedIDEQ applies the EQ predicate on the "EmployeedID" field.
func EmployeedIDEQ(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldEmployeedID, v))
}

// EmployeedIDNEQ applies the NEQ predicate on the "EmployeedID" field.
func EmployeedIDNEQ(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNEQ(FieldEmployeedID, v))
}

// EmployeedIDIn applies the In predicate on the "EmployeedID" field.
func EmployeedIDIn(vs ...int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIn(FieldEmployeedID, vs...))
}

// EmployeedIDNotIn applies the NotIn predicate on the "EmployeedID" field.
func EmployeedIDNotIn(vs ...int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotIn(FieldEmployeedID, vs...))
}

// EmployeedIDGT applies the GT predicate on the "EmployeedID" field.
func EmployeedIDGT(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGT(FieldEmployeedID, v))
}

// EmployeedIDGTE applies the GTE predicate on the "EmployeedID" field.
func EmployeedIDGTE(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGTE(FieldEmployeedID, v))
}

// EmployeedIDLT applies the LT predicate on the "EmployeedID" field.
func EmployeedIDLT(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLT(FieldEmployeedID, v))
}

// EmployeedIDLTE applies the LTE predicate on the "EmployeedID" field.
func EmployeedIDLTE(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLTE(FieldEmployeedID, v))
}

// EmployeeNameEQ applies the EQ predicate on the "EmployeeName" field.
func EmployeeNameEQ(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldEmployeeName, v))
}

// EmployeeNameNEQ applies the NEQ predicate on the "EmployeeName" field.
func EmployeeNameNEQ(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNEQ(FieldEmployeeName, v))
}

// EmployeeNameIn applies the In predicate on the "EmployeeName" field.
func EmployeeNameIn(vs ...string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIn(FieldEmployeeName, vs...))
}

// EmployeeNameNotIn applies the NotIn predicate on the "EmployeeName" field.
func EmployeeNameNotIn(vs ...string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotIn(FieldEmployeeName, vs...))
}

// EmployeeNameGT applies the GT predicate on the "EmployeeName" field.
func EmployeeNameGT(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGT(FieldEmployeeName, v))
}

// EmployeeNameGTE applies the GTE predicate on the "EmployeeName" field.
func EmployeeNameGTE(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGTE(FieldEmployeeName, v))
}

// EmployeeNameLT applies the LT predicate on the "EmployeeName" field.
func EmployeeNameLT(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLT(FieldEmployeeName, v))
}

// EmployeeNameLTE applies the LTE predicate on the "EmployeeName" field.
func EmployeeNameLTE(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLTE(FieldEmployeeName, v))
}

// EmployeeNameContains applies the Contains predicate on the "EmployeeName" field.
func EmployeeNameContains(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldContains(FieldEmployeeName, v))
}

// EmployeeNameHasPrefix applies the HasPrefix predicate on the "EmployeeName" field.
func EmployeeNameHasPrefix(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldHasPrefix(FieldEmployeeName, v))
}

// EmployeeNameHasSuffix applies the HasSuffix predicate on the "EmployeeName" field.
func EmployeeNameHasSuffix(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldHasSuffix(FieldEmployeeName, v))
}

// EmployeeNameEqualFold applies the EqualFold predicate on the "EmployeeName" field.
func EmployeeNameEqualFold(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEqualFold(FieldEmployeeName, v))
}

// EmployeeNameContainsFold applies the ContainsFold predicate on the "EmployeeName" field.
func EmployeeNameContainsFold(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldContainsFold(FieldEmployeeName, v))
}

// EmailIdEQ applies the EQ predicate on the "EmailId" field.
func EmailIdEQ(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldEmailId, v))
}

// EmailIdNEQ applies the NEQ predicate on the "EmailId" field.
func EmailIdNEQ(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNEQ(FieldEmailId, v))
}

// EmailIdIn applies the In predicate on the "EmailId" field.
func EmailIdIn(vs ...string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIn(FieldEmailId, vs...))
}

// EmailIdNotIn applies the NotIn predicate on the "EmailId" field.
func EmailIdNotIn(vs ...string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotIn(FieldEmailId, vs...))
}

// EmailIdGT applies the GT predicate on the "EmailId" field.
func EmailIdGT(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGT(FieldEmailId, v))
}

// EmailIdGTE applies the GTE predicate on the "EmailId" field.
func EmailIdGTE(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGTE(FieldEmailId, v))
}

// EmailIdLT applies the LT predicate on the "EmailId" field.
func EmailIdLT(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLT(FieldEmailId, v))
}

// EmailIdLTE applies the LTE predicate on the "EmailId" field.
func EmailIdLTE(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLTE(FieldEmailId, v))
}

// EmailIdContains applies the Contains predicate on the "EmailId" field.
func EmailIdContains(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldContains(FieldEmailId, v))
}

// EmailIdHasPrefix applies the HasPrefix predicate on the "EmailId" field.
func EmailIdHasPrefix(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldHasPrefix(FieldEmailId, v))
}

// EmailIdHasSuffix applies the HasSuffix predicate on the "EmailId" field.
func EmailIdHasSuffix(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldHasSuffix(FieldEmailId, v))
}

// EmailIdEqualFold applies the EqualFold predicate on the "EmailId" field.
func EmailIdEqualFold(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEqualFold(FieldEmailId, v))
}

// EmailIdContainsFold applies the ContainsFold predicate on the "EmailId" field.
func EmailIdContainsFold(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldContainsFold(FieldEmailId, v))
}

// MobileNumberEQ applies the EQ predicate on the "MobileNumber" field.
func MobileNumberEQ(v int64) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldMobileNumber, v))
}

// MobileNumberNEQ applies the NEQ predicate on the "MobileNumber" field.
func MobileNumberNEQ(v int64) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNEQ(FieldMobileNumber, v))
}

// MobileNumberIn applies the In predicate on the "MobileNumber" field.
func MobileNumberIn(vs ...int64) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIn(FieldMobileNumber, vs...))
}

// MobileNumberNotIn applies the NotIn predicate on the "MobileNumber" field.
func MobileNumberNotIn(vs ...int64) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotIn(FieldMobileNumber, vs...))
}

// MobileNumberGT applies the GT predicate on the "MobileNumber" field.
func MobileNumberGT(v int64) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGT(FieldMobileNumber, v))
}

// MobileNumberGTE applies the GTE predicate on the "MobileNumber" field.
func MobileNumberGTE(v int64) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGTE(FieldMobileNumber, v))
}

// MobileNumberLT applies the LT predicate on the "MobileNumber" field.
func MobileNumberLT(v int64) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLT(FieldMobileNumber, v))
}

// MobileNumberLTE applies the LTE predicate on the "MobileNumber" field.
func MobileNumberLTE(v int64) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLTE(FieldMobileNumber, v))
}

// SequenceNumberEQ applies the EQ predicate on the "SequenceNumber" field.
func SequenceNumberEQ(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldSequenceNumber, v))
}

// SequenceNumberNEQ applies the NEQ predicate on the "SequenceNumber" field.
func SequenceNumberNEQ(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNEQ(FieldSequenceNumber, v))
}

// SequenceNumberIn applies the In predicate on the "SequenceNumber" field.
func SequenceNumberIn(vs ...int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIn(FieldSequenceNumber, vs...))
}

// SequenceNumberNotIn applies the NotIn predicate on the "SequenceNumber" field.
func SequenceNumberNotIn(vs ...int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotIn(FieldSequenceNumber, vs...))
}

// SequenceNumberGT applies the GT predicate on the "SequenceNumber" field.
func SequenceNumberGT(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGT(FieldSequenceNumber, v))
}

// SequenceNumberGTE applies the GTE predicate on the "SequenceNumber" field.
func SequenceNumberGTE(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGTE(FieldSequenceNumber, v))
}

// SequenceNumberLT applies the LT predicate on the "SequenceNumber" field.
func SequenceNumberLT(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLT(FieldSequenceNumber, v))
}

// SequenceNumberLTE applies the LTE predicate on the "SequenceNumber" field.
func SequenceNumberLTE(v int32) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLTE(FieldSequenceNumber, v))
}

// SequenceNumberIsNil applies the IsNil predicate on the "SequenceNumber" field.
func SequenceNumberIsNil() predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIsNull(FieldSequenceNumber))
}

// SequenceNumberNotNil applies the NotNil predicate on the "SequenceNumber" field.
func SequenceNumberNotNil() predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotNull(FieldSequenceNumber))
}

// StatusEQ applies the EQ predicate on the "Status" field.
func StatusEQ(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "Status" field.
func StatusNEQ(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "Status" field.
func StatusIn(vs ...string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "Status" field.
func StatusNotIn(vs ...string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "Status" field.
func StatusGT(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "Status" field.
func StatusGTE(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "Status" field.
func StatusLT(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "Status" field.
func StatusLTE(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "Status" field.
func StatusContains(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "Status" field.
func StatusHasPrefix(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "Status" field.
func StatusHasSuffix(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "Status" field.
func StatusIsNil() predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "Status" field.
func StatusNotNil() predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "Status" field.
func StatusEqualFold(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "Status" field.
func StatusContainsFold(v string) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(sql.FieldContainsFold(FieldStatus, v))
}

// HasEmployeeUser applies the HasEdge predicate on the "employee_user" edge.
func HasEmployeeUser() predicate.DirectorateUsers {
	return predicate.DirectorateUsers(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmployeeUserTable, EmployeeUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeUserWith applies the HasEdge predicate on the "employee_user" edge with a given conditions (other predicates).
func HasEmployeeUserWith(preds ...predicate.Employees) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(func(s *sql.Selector) {
		step := newEmployeeUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DirectorateUsers) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DirectorateUsers) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(func(s *sql.Selector) {
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
func Not(p predicate.DirectorateUsers) predicate.DirectorateUsers {
	return predicate.DirectorateUsers(func(s *sql.Selector) {
		p(s.Not())
	})
}