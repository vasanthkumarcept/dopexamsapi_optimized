// Code generated by ent, DO NOT EDIT.

package recommendationsgdspmapplications

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the recommendationsgdspmapplications type in the database.
	Label = "recommendations_gdspm_applications"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "RecommendationId"
	// FieldApplicationID holds the string denoting the applicationid field in the database.
	FieldApplicationID = "application_id"
	// FieldEmployeeID holds the string denoting the employeeid field in the database.
	FieldEmployeeID = "employee_id"
	// FieldExamNameCode holds the string denoting the examnamecode field in the database.
	FieldExamNameCode = "exam_name_code"
	// FieldExamYear holds the string denoting the examyear field in the database.
	FieldExamYear = "exam_year"
	// FieldVacancyYear holds the string denoting the vacancyyear field in the database.
	FieldVacancyYear = "vacancy_year"
	// FieldPost holds the string denoting the post field in the database.
	FieldPost = "post"
	// FieldEligible holds the string denoting the eligible field in the database.
	FieldEligible = "eligible"
	// FieldCARecommendations holds the string denoting the ca_recommendations field in the database.
	FieldCARecommendations = "ca_recommendations"
	// FieldCAUpdatedAt holds the string denoting the ca_updatedat field in the database.
	FieldCAUpdatedAt = "ca_updated_at"
	// FieldCAUserName holds the string denoting the ca_username field in the database.
	FieldCAUserName = "ca_user_name"
	// FieldCARemarks holds the string denoting the ca_remarks field in the database.
	FieldCARemarks = "ca_remarks"
	// FieldNORecommendations holds the string denoting the no_recommendations field in the database.
	FieldNORecommendations = "no_recommendations"
	// FieldNOUpdatedAt holds the string denoting the no_updatedat field in the database.
	FieldNOUpdatedAt = "no_updated_at"
	// FieldNOUserName holds the string denoting the no_username field in the database.
	FieldNOUserName = "no_user_name"
	// FieldNORemarks holds the string denoting the no_remarks field in the database.
	FieldNORemarks = "no_remarks"
	// FieldApplicationStatus holds the string denoting the applicationstatus field in the database.
	FieldApplicationStatus = "application_status"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updatedby field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldGenerateHallTicketFlag holds the string denoting the generatehallticketflag field in the database.
	FieldGenerateHallTicketFlag = "generate_hall_ticket_flag"
	// EdgeApplnRef holds the string denoting the applnref edge name in mutations.
	EdgeApplnRef = "ApplnRef"
	// Exam_Applications_GDSPMFieldID holds the string denoting the ID field of the Exam_Applications_GDSPM.
	Exam_Applications_GDSPMFieldID = "ApplicationID"
	// Table holds the table name of the recommendationsgdspmapplications in the database.
	Table = "RecommendationsGDSPMApplications"
	// ApplnRefTable is the table that holds the ApplnRef relation/edge.
	ApplnRefTable = "RecommendationsGDSPMApplications"
	// ApplnRefInverseTable is the table name for the Exam_Applications_GDSPM entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_gdspm" package.
	ApplnRefInverseTable = "Exam_Applications_GDSPM"
	// ApplnRefColumn is the table column denoting the ApplnRef relation/edge.
	ApplnRefColumn = "exam_applications_gdspm_gdspm_applications_ref"
)

// Columns holds all SQL columns for recommendationsgdspmapplications fields.
var Columns = []string{
	FieldID,
	FieldApplicationID,
	FieldEmployeeID,
	FieldExamNameCode,
	FieldExamYear,
	FieldVacancyYear,
	FieldPost,
	FieldEligible,
	FieldCARecommendations,
	FieldCAUpdatedAt,
	FieldCAUserName,
	FieldCARemarks,
	FieldNORecommendations,
	FieldNOUpdatedAt,
	FieldNOUserName,
	FieldNORemarks,
	FieldApplicationStatus,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldGenerateHallTicketFlag,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "RecommendationsGDSPMApplications"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"exam_applications_gdspm_gdspm_applications_ref",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// DefaultUpdatedBy holds the default value on creation for the "UpdatedBy" field.
	DefaultUpdatedBy string
)

// OrderOption defines the ordering options for the RecommendationsGDSPMApplications queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByApplicationID orders the results by the ApplicationID field.
func ByApplicationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplicationID, opts...).ToFunc()
}

// ByEmployeeID orders the results by the EmployeeID field.
func ByEmployeeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeeID, opts...).ToFunc()
}

// ByExamNameCode orders the results by the ExamNameCode field.
func ByExamNameCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamNameCode, opts...).ToFunc()
}

// ByExamYear orders the results by the ExamYear field.
func ByExamYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamYear, opts...).ToFunc()
}

// ByVacancyYear orders the results by the VacancyYear field.
func ByVacancyYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVacancyYear, opts...).ToFunc()
}

// ByPost orders the results by the Post field.
func ByPost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPost, opts...).ToFunc()
}

// ByEligible orders the results by the Eligible field.
func ByEligible(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEligible, opts...).ToFunc()
}

// ByCARecommendations orders the results by the CA_Recommendations field.
func ByCARecommendations(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCARecommendations, opts...).ToFunc()
}

// ByCAUpdatedAt orders the results by the CA_UpdatedAt field.
func ByCAUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCAUpdatedAt, opts...).ToFunc()
}

// ByCAUserName orders the results by the CA_UserName field.
func ByCAUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCAUserName, opts...).ToFunc()
}

// ByCARemarks orders the results by the CA_Remarks field.
func ByCARemarks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCARemarks, opts...).ToFunc()
}

// ByNORecommendations orders the results by the NO_Recommendations field.
func ByNORecommendations(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNORecommendations, opts...).ToFunc()
}

// ByNOUpdatedAt orders the results by the NO_UpdatedAt field.
func ByNOUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNOUpdatedAt, opts...).ToFunc()
}

// ByNOUserName orders the results by the NO_UserName field.
func ByNOUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNOUserName, opts...).ToFunc()
}

// ByNORemarks orders the results by the NO_Remarks field.
func ByNORemarks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNORemarks, opts...).ToFunc()
}

// ByApplicationStatus orders the results by the ApplicationStatus field.
func ByApplicationStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplicationStatus, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the UpdatedBy field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByGenerateHallTicketFlag orders the results by the GenerateHallTicketFlag field.
func ByGenerateHallTicketFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGenerateHallTicketFlag, opts...).ToFunc()
}

// ByApplnRefField orders the results by ApplnRef field.
func ByApplnRefField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApplnRefStep(), sql.OrderByField(field, opts...))
	}
}
func newApplnRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApplnRefInverseTable, Exam_Applications_GDSPMFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ApplnRefTable, ApplnRefColumn),
	)
}
