// Code generated by ent, DO NOT EDIT.

package examcalendar

import (
	"recruit/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldID, id))
}

// ExamYear applies equality check predicate on the "ExamYear" field. It's identical to ExamYearEQ.
func ExamYear(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldExamYear, v))
}

// ExamName applies equality check predicate on the "ExamName" field. It's identical to ExamNameEQ.
func ExamName(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldExamName, v))
}

// ExamCode applies equality check predicate on the "ExamCode" field. It's identical to ExamCodeEQ.
func ExamCode(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldExamCode, v))
}

// NotificationDate applies equality check predicate on the "NotificationDate" field. It's identical to NotificationDateEQ.
func NotificationDate(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldNotificationDate, v))
}

// ModelNotificationDate applies equality check predicate on the "ModelNotificationDate" field. It's identical to ModelNotificationDateEQ.
func ModelNotificationDate(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldModelNotificationDate, v))
}

// ApplicationEndDate applies equality check predicate on the "ApplicationEndDate" field. It's identical to ApplicationEndDateEQ.
func ApplicationEndDate(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldApplicationEndDate, v))
}

// ApprovedOrderDate applies equality check predicate on the "ApprovedOrderDate" field. It's identical to ApprovedOrderDateEQ.
func ApprovedOrderDate(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldApprovedOrderDate, v))
}

// TentativeResultDate applies equality check predicate on the "TentativeResultDate" field. It's identical to TentativeResultDateEQ.
func TentativeResultDate(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldTentativeResultDate, v))
}

// CreatedDate applies equality check predicate on the "CreatedDate" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldCreatedDate, v))
}

// ApprovedOrderNumber applies equality check predicate on the "ApprovedOrderNumber" field. It's identical to ApprovedOrderNumberEQ.
func ApprovedOrderNumber(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldApprovedOrderNumber, v))
}

// VacancyYearCode applies equality check predicate on the "VacancyYearCode" field. It's identical to VacancyYearCodeEQ.
func VacancyYearCode(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldVacancyYearCode, v))
}

// PaperCode applies equality check predicate on the "PaperCode" field. It's identical to PaperCodeEQ.
func PaperCode(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldPaperCode, v))
}

// ExamCodePS applies equality check predicate on the "ExamCodePS" field. It's identical to ExamCodePSEQ.
func ExamCodePS(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldExamCodePS, v))
}

// ExamYearEQ applies the EQ predicate on the "ExamYear" field.
func ExamYearEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldExamYear, v))
}

// ExamYearNEQ applies the NEQ predicate on the "ExamYear" field.
func ExamYearNEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldExamYear, v))
}

// ExamYearIn applies the In predicate on the "ExamYear" field.
func ExamYearIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldExamYear, vs...))
}

// ExamYearNotIn applies the NotIn predicate on the "ExamYear" field.
func ExamYearNotIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldExamYear, vs...))
}

// ExamYearGT applies the GT predicate on the "ExamYear" field.
func ExamYearGT(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldExamYear, v))
}

// ExamYearGTE applies the GTE predicate on the "ExamYear" field.
func ExamYearGTE(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldExamYear, v))
}

// ExamYearLT applies the LT predicate on the "ExamYear" field.
func ExamYearLT(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldExamYear, v))
}

// ExamYearLTE applies the LTE predicate on the "ExamYear" field.
func ExamYearLTE(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldExamYear, v))
}

// ExamNameEQ applies the EQ predicate on the "ExamName" field.
func ExamNameEQ(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldExamName, v))
}

// ExamNameNEQ applies the NEQ predicate on the "ExamName" field.
func ExamNameNEQ(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldExamName, v))
}

// ExamNameIn applies the In predicate on the "ExamName" field.
func ExamNameIn(vs ...string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldExamName, vs...))
}

// ExamNameNotIn applies the NotIn predicate on the "ExamName" field.
func ExamNameNotIn(vs ...string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldExamName, vs...))
}

// ExamNameGT applies the GT predicate on the "ExamName" field.
func ExamNameGT(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldExamName, v))
}

// ExamNameGTE applies the GTE predicate on the "ExamName" field.
func ExamNameGTE(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldExamName, v))
}

// ExamNameLT applies the LT predicate on the "ExamName" field.
func ExamNameLT(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldExamName, v))
}

// ExamNameLTE applies the LTE predicate on the "ExamName" field.
func ExamNameLTE(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldExamName, v))
}

// ExamNameContains applies the Contains predicate on the "ExamName" field.
func ExamNameContains(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldContains(FieldExamName, v))
}

// ExamNameHasPrefix applies the HasPrefix predicate on the "ExamName" field.
func ExamNameHasPrefix(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldHasPrefix(FieldExamName, v))
}

// ExamNameHasSuffix applies the HasSuffix predicate on the "ExamName" field.
func ExamNameHasSuffix(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldHasSuffix(FieldExamName, v))
}

// ExamNameEqualFold applies the EqualFold predicate on the "ExamName" field.
func ExamNameEqualFold(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEqualFold(FieldExamName, v))
}

// ExamNameContainsFold applies the ContainsFold predicate on the "ExamName" field.
func ExamNameContainsFold(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldContainsFold(FieldExamName, v))
}

// ExamCodeEQ applies the EQ predicate on the "ExamCode" field.
func ExamCodeEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldExamCode, v))
}

// ExamCodeNEQ applies the NEQ predicate on the "ExamCode" field.
func ExamCodeNEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldExamCode, v))
}

// ExamCodeIn applies the In predicate on the "ExamCode" field.
func ExamCodeIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldExamCode, vs...))
}

// ExamCodeNotIn applies the NotIn predicate on the "ExamCode" field.
func ExamCodeNotIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldExamCode, vs...))
}

// ExamCodeIsNil applies the IsNil predicate on the "ExamCode" field.
func ExamCodeIsNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIsNull(FieldExamCode))
}

// ExamCodeNotNil applies the NotNil predicate on the "ExamCode" field.
func ExamCodeNotNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotNull(FieldExamCode))
}

// NotificationDateEQ applies the EQ predicate on the "NotificationDate" field.
func NotificationDateEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldNotificationDate, v))
}

// NotificationDateNEQ applies the NEQ predicate on the "NotificationDate" field.
func NotificationDateNEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldNotificationDate, v))
}

// NotificationDateIn applies the In predicate on the "NotificationDate" field.
func NotificationDateIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldNotificationDate, vs...))
}

// NotificationDateNotIn applies the NotIn predicate on the "NotificationDate" field.
func NotificationDateNotIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldNotificationDate, vs...))
}

// NotificationDateGT applies the GT predicate on the "NotificationDate" field.
func NotificationDateGT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldNotificationDate, v))
}

// NotificationDateGTE applies the GTE predicate on the "NotificationDate" field.
func NotificationDateGTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldNotificationDate, v))
}

// NotificationDateLT applies the LT predicate on the "NotificationDate" field.
func NotificationDateLT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldNotificationDate, v))
}

// NotificationDateLTE applies the LTE predicate on the "NotificationDate" field.
func NotificationDateLTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldNotificationDate, v))
}

// ModelNotificationDateEQ applies the EQ predicate on the "ModelNotificationDate" field.
func ModelNotificationDateEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldModelNotificationDate, v))
}

// ModelNotificationDateNEQ applies the NEQ predicate on the "ModelNotificationDate" field.
func ModelNotificationDateNEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldModelNotificationDate, v))
}

// ModelNotificationDateIn applies the In predicate on the "ModelNotificationDate" field.
func ModelNotificationDateIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldModelNotificationDate, vs...))
}

// ModelNotificationDateNotIn applies the NotIn predicate on the "ModelNotificationDate" field.
func ModelNotificationDateNotIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldModelNotificationDate, vs...))
}

// ModelNotificationDateGT applies the GT predicate on the "ModelNotificationDate" field.
func ModelNotificationDateGT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldModelNotificationDate, v))
}

// ModelNotificationDateGTE applies the GTE predicate on the "ModelNotificationDate" field.
func ModelNotificationDateGTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldModelNotificationDate, v))
}

// ModelNotificationDateLT applies the LT predicate on the "ModelNotificationDate" field.
func ModelNotificationDateLT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldModelNotificationDate, v))
}

// ModelNotificationDateLTE applies the LTE predicate on the "ModelNotificationDate" field.
func ModelNotificationDateLTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldModelNotificationDate, v))
}

// ApplicationEndDateEQ applies the EQ predicate on the "ApplicationEndDate" field.
func ApplicationEndDateEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldApplicationEndDate, v))
}

// ApplicationEndDateNEQ applies the NEQ predicate on the "ApplicationEndDate" field.
func ApplicationEndDateNEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldApplicationEndDate, v))
}

// ApplicationEndDateIn applies the In predicate on the "ApplicationEndDate" field.
func ApplicationEndDateIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldApplicationEndDate, vs...))
}

// ApplicationEndDateNotIn applies the NotIn predicate on the "ApplicationEndDate" field.
func ApplicationEndDateNotIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldApplicationEndDate, vs...))
}

// ApplicationEndDateGT applies the GT predicate on the "ApplicationEndDate" field.
func ApplicationEndDateGT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldApplicationEndDate, v))
}

// ApplicationEndDateGTE applies the GTE predicate on the "ApplicationEndDate" field.
func ApplicationEndDateGTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldApplicationEndDate, v))
}

// ApplicationEndDateLT applies the LT predicate on the "ApplicationEndDate" field.
func ApplicationEndDateLT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldApplicationEndDate, v))
}

// ApplicationEndDateLTE applies the LTE predicate on the "ApplicationEndDate" field.
func ApplicationEndDateLTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldApplicationEndDate, v))
}

// ApprovedOrderDateEQ applies the EQ predicate on the "ApprovedOrderDate" field.
func ApprovedOrderDateEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldApprovedOrderDate, v))
}

// ApprovedOrderDateNEQ applies the NEQ predicate on the "ApprovedOrderDate" field.
func ApprovedOrderDateNEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldApprovedOrderDate, v))
}

// ApprovedOrderDateIn applies the In predicate on the "ApprovedOrderDate" field.
func ApprovedOrderDateIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldApprovedOrderDate, vs...))
}

// ApprovedOrderDateNotIn applies the NotIn predicate on the "ApprovedOrderDate" field.
func ApprovedOrderDateNotIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldApprovedOrderDate, vs...))
}

// ApprovedOrderDateGT applies the GT predicate on the "ApprovedOrderDate" field.
func ApprovedOrderDateGT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldApprovedOrderDate, v))
}

// ApprovedOrderDateGTE applies the GTE predicate on the "ApprovedOrderDate" field.
func ApprovedOrderDateGTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldApprovedOrderDate, v))
}

// ApprovedOrderDateLT applies the LT predicate on the "ApprovedOrderDate" field.
func ApprovedOrderDateLT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldApprovedOrderDate, v))
}

// ApprovedOrderDateLTE applies the LTE predicate on the "ApprovedOrderDate" field.
func ApprovedOrderDateLTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldApprovedOrderDate, v))
}

// TentativeResultDateEQ applies the EQ predicate on the "TentativeResultDate" field.
func TentativeResultDateEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldTentativeResultDate, v))
}

// TentativeResultDateNEQ applies the NEQ predicate on the "TentativeResultDate" field.
func TentativeResultDateNEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldTentativeResultDate, v))
}

// TentativeResultDateIn applies the In predicate on the "TentativeResultDate" field.
func TentativeResultDateIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldTentativeResultDate, vs...))
}

// TentativeResultDateNotIn applies the NotIn predicate on the "TentativeResultDate" field.
func TentativeResultDateNotIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldTentativeResultDate, vs...))
}

// TentativeResultDateGT applies the GT predicate on the "TentativeResultDate" field.
func TentativeResultDateGT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldTentativeResultDate, v))
}

// TentativeResultDateGTE applies the GTE predicate on the "TentativeResultDate" field.
func TentativeResultDateGTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldTentativeResultDate, v))
}

// TentativeResultDateLT applies the LT predicate on the "TentativeResultDate" field.
func TentativeResultDateLT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldTentativeResultDate, v))
}

// TentativeResultDateLTE applies the LTE predicate on the "TentativeResultDate" field.
func TentativeResultDateLTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldTentativeResultDate, v))
}

// TentativeResultDateIsNil applies the IsNil predicate on the "TentativeResultDate" field.
func TentativeResultDateIsNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIsNull(FieldTentativeResultDate))
}

// TentativeResultDateNotNil applies the NotNil predicate on the "TentativeResultDate" field.
func TentativeResultDateNotNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotNull(FieldTentativeResultDate))
}

// CreatedDateEQ applies the EQ predicate on the "CreatedDate" field.
func CreatedDateEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldCreatedDate, v))
}

// CreatedDateNEQ applies the NEQ predicate on the "CreatedDate" field.
func CreatedDateNEQ(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldCreatedDate, v))
}

// CreatedDateIn applies the In predicate on the "CreatedDate" field.
func CreatedDateIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldCreatedDate, vs...))
}

// CreatedDateNotIn applies the NotIn predicate on the "CreatedDate" field.
func CreatedDateNotIn(vs ...time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldCreatedDate, vs...))
}

// CreatedDateGT applies the GT predicate on the "CreatedDate" field.
func CreatedDateGT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldCreatedDate, v))
}

// CreatedDateGTE applies the GTE predicate on the "CreatedDate" field.
func CreatedDateGTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldCreatedDate, v))
}

// CreatedDateLT applies the LT predicate on the "CreatedDate" field.
func CreatedDateLT(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldCreatedDate, v))
}

// CreatedDateLTE applies the LTE predicate on the "CreatedDate" field.
func CreatedDateLTE(v time.Time) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldCreatedDate, v))
}

// ApprovedOrderNumberEQ applies the EQ predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberEQ(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberNEQ applies the NEQ predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberNEQ(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberIn applies the In predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberIn(vs ...string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldApprovedOrderNumber, vs...))
}

// ApprovedOrderNumberNotIn applies the NotIn predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberNotIn(vs ...string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldApprovedOrderNumber, vs...))
}

// ApprovedOrderNumberGT applies the GT predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberGT(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberGTE applies the GTE predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberGTE(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberLT applies the LT predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberLT(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberLTE applies the LTE predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberLTE(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberContains applies the Contains predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberContains(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldContains(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberHasPrefix applies the HasPrefix predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberHasPrefix(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldHasPrefix(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberHasSuffix applies the HasSuffix predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberHasSuffix(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldHasSuffix(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberEqualFold applies the EqualFold predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberEqualFold(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEqualFold(FieldApprovedOrderNumber, v))
}

// ApprovedOrderNumberContainsFold applies the ContainsFold predicate on the "ApprovedOrderNumber" field.
func ApprovedOrderNumberContainsFold(v string) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldContainsFold(FieldApprovedOrderNumber, v))
}

// VacancyYearsIsNil applies the IsNil predicate on the "VacancyYears" field.
func VacancyYearsIsNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIsNull(FieldVacancyYears))
}

// VacancyYearsNotNil applies the NotNil predicate on the "VacancyYears" field.
func VacancyYearsNotNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotNull(FieldVacancyYears))
}

// ExamPapersIsNil applies the IsNil predicate on the "ExamPapers" field.
func ExamPapersIsNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIsNull(FieldExamPapers))
}

// ExamPapersNotNil applies the NotNil predicate on the "ExamPapers" field.
func ExamPapersNotNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotNull(FieldExamPapers))
}

// VacancyYearCodeEQ applies the EQ predicate on the "VacancyYearCode" field.
func VacancyYearCodeEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldVacancyYearCode, v))
}

// VacancyYearCodeNEQ applies the NEQ predicate on the "VacancyYearCode" field.
func VacancyYearCodeNEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldVacancyYearCode, v))
}

// VacancyYearCodeIn applies the In predicate on the "VacancyYearCode" field.
func VacancyYearCodeIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldVacancyYearCode, vs...))
}

// VacancyYearCodeNotIn applies the NotIn predicate on the "VacancyYearCode" field.
func VacancyYearCodeNotIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldVacancyYearCode, vs...))
}

// VacancyYearCodeIsNil applies the IsNil predicate on the "VacancyYearCode" field.
func VacancyYearCodeIsNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIsNull(FieldVacancyYearCode))
}

// VacancyYearCodeNotNil applies the NotNil predicate on the "VacancyYearCode" field.
func VacancyYearCodeNotNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotNull(FieldVacancyYearCode))
}

// PaperCodeEQ applies the EQ predicate on the "PaperCode" field.
func PaperCodeEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldPaperCode, v))
}

// PaperCodeNEQ applies the NEQ predicate on the "PaperCode" field.
func PaperCodeNEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldPaperCode, v))
}

// PaperCodeIn applies the In predicate on the "PaperCode" field.
func PaperCodeIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldPaperCode, vs...))
}

// PaperCodeNotIn applies the NotIn predicate on the "PaperCode" field.
func PaperCodeNotIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldPaperCode, vs...))
}

// PaperCodeIsNil applies the IsNil predicate on the "PaperCode" field.
func PaperCodeIsNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIsNull(FieldPaperCode))
}

// PaperCodeNotNil applies the NotNil predicate on the "PaperCode" field.
func PaperCodeNotNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotNull(FieldPaperCode))
}

// ExamCodePSEQ applies the EQ predicate on the "ExamCodePS" field.
func ExamCodePSEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldEQ(FieldExamCodePS, v))
}

// ExamCodePSNEQ applies the NEQ predicate on the "ExamCodePS" field.
func ExamCodePSNEQ(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNEQ(FieldExamCodePS, v))
}

// ExamCodePSIn applies the In predicate on the "ExamCodePS" field.
func ExamCodePSIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIn(FieldExamCodePS, vs...))
}

// ExamCodePSNotIn applies the NotIn predicate on the "ExamCodePS" field.
func ExamCodePSNotIn(vs ...int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotIn(FieldExamCodePS, vs...))
}

// ExamCodePSGT applies the GT predicate on the "ExamCodePS" field.
func ExamCodePSGT(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGT(FieldExamCodePS, v))
}

// ExamCodePSGTE applies the GTE predicate on the "ExamCodePS" field.
func ExamCodePSGTE(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldGTE(FieldExamCodePS, v))
}

// ExamCodePSLT applies the LT predicate on the "ExamCodePS" field.
func ExamCodePSLT(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLT(FieldExamCodePS, v))
}

// ExamCodePSLTE applies the LTE predicate on the "ExamCodePS" field.
func ExamCodePSLTE(v int32) predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldLTE(FieldExamCodePS, v))
}

// ExamCodePSIsNil applies the IsNil predicate on the "ExamCodePS" field.
func ExamCodePSIsNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldIsNull(FieldExamCodePS))
}

// ExamCodePSNotNil applies the NotNil predicate on the "ExamCodePS" field.
func ExamCodePSNotNil() predicate.ExamCalendar {
	return predicate.ExamCalendar(sql.FieldNotNull(FieldExamCodePS))
}

// HasVcyYears applies the HasEdge predicate on the "vcy_years" edge.
func HasVcyYears() predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VcyYearsTable, VcyYearsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVcyYearsWith applies the HasEdge predicate on the "vcy_years" edge with a given conditions (other predicates).
func HasVcyYearsWith(preds ...predicate.VacancyYear) predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		step := newVcyYearsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExams applies the HasEdge predicate on the "exams" edge.
func HasExams() predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExamsTable, ExamsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExamsWith applies the HasEdge predicate on the "exams" edge with a given conditions (other predicates).
func HasExamsWith(preds ...predicate.Exam) predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		step := newExamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPapers applies the HasEdge predicate on the "papers" edge.
func HasPapers() predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PapersTable, PapersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPapersWith applies the HasEdge predicate on the "papers" edge with a given conditions (other predicates).
func HasPapersWith(preds ...predicate.ExamPapers) predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		step := newPapersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifyRef applies the HasEdge predicate on the "Notify_ref" edge.
func HasNotifyRef() predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotifyRefTable, NotifyRefColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotifyRefWith applies the HasEdge predicate on the "Notify_ref" edge with a given conditions (other predicates).
func HasNotifyRefWith(preds ...predicate.Notification) predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		step := newNotifyRefStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExamCalendar) predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExamCalendar) predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
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
func Not(p predicate.ExamCalendar) predicate.ExamCalendar {
	return predicate.ExamCalendar(func(s *sql.Selector) {
		p(s.Not())
	})
}