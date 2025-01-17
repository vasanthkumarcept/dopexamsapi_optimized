// Code generated by ent, DO NOT EDIT.

package examcenterhall

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the examcenterhall type in the database.
	Label = "exam_center_hall"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "ExamCenterHall"
	// FieldCenterCode holds the string denoting the centercode field in the database.
	FieldCenterCode = "center_code"
	// FieldCityID holds the string denoting the cityid field in the database.
	FieldCityID = "city_id"
	// FieldExamCenterName holds the string denoting the examcentername field in the database.
	FieldExamCenterName = "exam_center_name"
	// FieldExamYear holds the string denoting the examyear field in the database.
	FieldExamYear = "exam_year"
	// FieldExamCode holds the string denoting the examcode field in the database.
	FieldExamCode = "exam_code"
	// FieldExamName holds the string denoting the examname field in the database.
	FieldExamName = "exam_name"
	// FieldCenterCityName holds the string denoting the centercityname field in the database.
	FieldCenterCityName = "center_city_name"
	// FieldConductedByFacilityID holds the string denoting the conductedbyfacilityid field in the database.
	FieldConductedByFacilityID = "conducted_by_facility_id"
	// FieldConductedBy holds the string denoting the conductedby field in the database.
	FieldConductedBy = "conducted_by"
	// FieldHallName holds the string denoting the hallname field in the database.
	FieldHallName = "hall_name"
	// FieldAdminCircleOfficeID holds the string denoting the admincircleofficeid field in the database.
	FieldAdminCircleOfficeID = "admin_circle_office_id"
	// FieldMappingIdentificationNumber holds the string denoting the mappingidentificationnumber field in the database.
	FieldMappingIdentificationNumber = "mapping_identification_number"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedById holds the string denoting the createdbyid field in the database.
	FieldCreatedById = "created_by_id"
	// FieldCreatedByUserName holds the string denoting the createdbyusername field in the database.
	FieldCreatedByUserName = "created_by_user_name"
	// FieldCreatedByEmpId holds the string denoting the createdbyempid field in the database.
	FieldCreatedByEmpId = "created_by_emp_id"
	// FieldCreatedByDesignation holds the string denoting the createdbydesignation field in the database.
	FieldCreatedByDesignation = "created_by_designation"
	// FieldCreatedDate holds the string denoting the createddate field in the database.
	FieldCreatedDate = "created_date"
	// FieldModifiedById holds the string denoting the modifiedbyid field in the database.
	FieldModifiedById = "modified_by_id"
	// FieldModifiedByUserName holds the string denoting the modifiedbyusername field in the database.
	FieldModifiedByUserName = "modified_by_user_name"
	// FieldModifiedByEmpId holds the string denoting the modifiedbyempid field in the database.
	FieldModifiedByEmpId = "modified_by_emp_id"
	// FieldModifiedByDesignantion holds the string denoting the modifiedbydesignantion field in the database.
	FieldModifiedByDesignantion = "modified_by_designantion"
	// FieldModifiedDate holds the string denoting the modifieddate field in the database.
	FieldModifiedDate = "modified_date"
	// FieldDeletedById holds the string denoting the deletedbyid field in the database.
	FieldDeletedById = "deleted_by_id"
	// FieldDeletedByUserName holds the string denoting the deletedbyusername field in the database.
	FieldDeletedByUserName = "deleted_by_user_name"
	// FieldDeletedByEmpId holds the string denoting the deletedbyempid field in the database.
	FieldDeletedByEmpId = "deleted_by_emp_id"
	// FieldDeletedByDesignation holds the string denoting the deletedbydesignation field in the database.
	FieldDeletedByDesignation = "deleted_by_designation"
	// FieldDeletedDate holds the string denoting the deleteddate field in the database.
	FieldDeletedDate = "deleted_date"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldNoSeats holds the string denoting the noseats field in the database.
	FieldNoSeats = "no_seats"
	// EdgeExaCenterHall holds the string denoting the exacenterhall edge name in mutations.
	EdgeExaCenterHall = "ExaCenterHall"
	// CenterFieldID holds the string denoting the ID field of the Center.
	CenterFieldID = "CenterCode"
	// Table holds the table name of the examcenterhall in the database.
	Table = "ExamCenterHall"
	// ExaCenterHallTable is the table that holds the ExaCenterHall relation/edge.
	ExaCenterHallTable = "ExamCenterHall"
	// ExaCenterHallInverseTable is the table name for the Center entity.
	// It exists in this package in order to avoid circular dependency with the "center" package.
	ExaCenterHallInverseTable = "Center"
	// ExaCenterHallColumn is the table column denoting the ExaCenterHall relation/edge.
	ExaCenterHallColumn = "center_code"
)

// Columns holds all SQL columns for examcenterhall fields.
var Columns = []string{
	FieldID,
	FieldCenterCode,
	FieldCityID,
	FieldExamCenterName,
	FieldExamYear,
	FieldExamCode,
	FieldExamName,
	FieldCenterCityName,
	FieldConductedByFacilityID,
	FieldConductedBy,
	FieldHallName,
	FieldAdminCircleOfficeID,
	FieldMappingIdentificationNumber,
	FieldStatus,
	FieldCreatedById,
	FieldCreatedByUserName,
	FieldCreatedByEmpId,
	FieldCreatedByDesignation,
	FieldCreatedDate,
	FieldModifiedById,
	FieldModifiedByUserName,
	FieldModifiedByEmpId,
	FieldModifiedByDesignantion,
	FieldModifiedDate,
	FieldDeletedById,
	FieldDeletedByUserName,
	FieldDeletedByEmpId,
	FieldDeletedByDesignation,
	FieldDeletedDate,
	FieldUpdatedAt,
	FieldNoSeats,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCenterCode holds the default value on creation for the "CenterCode" field.
	DefaultCenterCode int32
	// DefaultCityID holds the default value on creation for the "CityID" field.
	DefaultCityID int32
	// DefaultExamCode holds the default value on creation for the "ExamCode" field.
	DefaultExamCode int32
	// DefaultStatus holds the default value on creation for the "Status" field.
	DefaultStatus string
	// DefaultCreatedById holds the default value on creation for the "CreatedById" field.
	DefaultCreatedById int64
	// DefaultCreatedByEmpId holds the default value on creation for the "CreatedByEmpId" field.
	DefaultCreatedByEmpId int64
	// DefaultCreatedDate holds the default value on creation for the "CreatedDate" field.
	DefaultCreatedDate func() time.Time
	// DefaultModifiedById holds the default value on creation for the "ModifiedById" field.
	DefaultModifiedById int64
	// DefaultModifiedByEmpId holds the default value on creation for the "ModifiedByEmpId" field.
	DefaultModifiedByEmpId int64
	// DefaultModifiedDate holds the default value on creation for the "ModifiedDate" field.
	DefaultModifiedDate func() time.Time
	// DefaultDeletedById holds the default value on creation for the "DeletedById" field.
	DefaultDeletedById int64
	// DefaultDeletedByEmpId holds the default value on creation for the "DeletedByEmpId" field.
	DefaultDeletedByEmpId int64
	// DefaultDeletedDate holds the default value on creation for the "DeletedDate" field.
	DefaultDeletedDate func() time.Time
	// DefaultNoSeats holds the default value on creation for the "NoSeats" field.
	DefaultNoSeats int32
)

// OrderOption defines the ordering options for the ExamCenterHall queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCenterCode orders the results by the CenterCode field.
func ByCenterCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCenterCode, opts...).ToFunc()
}

// ByCityID orders the results by the CityID field.
func ByCityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCityID, opts...).ToFunc()
}

// ByExamCenterName orders the results by the ExamCenterName field.
func ByExamCenterName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamCenterName, opts...).ToFunc()
}

// ByExamYear orders the results by the ExamYear field.
func ByExamYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamYear, opts...).ToFunc()
}

// ByExamCode orders the results by the ExamCode field.
func ByExamCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamCode, opts...).ToFunc()
}

// ByExamName orders the results by the ExamName field.
func ByExamName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamName, opts...).ToFunc()
}

// ByCenterCityName orders the results by the CenterCityName field.
func ByCenterCityName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCenterCityName, opts...).ToFunc()
}

// ByConductedByFacilityID orders the results by the ConductedByFacilityID field.
func ByConductedByFacilityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConductedByFacilityID, opts...).ToFunc()
}

// ByConductedBy orders the results by the ConductedBy field.
func ByConductedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConductedBy, opts...).ToFunc()
}

// ByHallName orders the results by the HallName field.
func ByHallName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHallName, opts...).ToFunc()
}

// ByAdminCircleOfficeID orders the results by the AdminCircleOfficeID field.
func ByAdminCircleOfficeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdminCircleOfficeID, opts...).ToFunc()
}

// ByStatus orders the results by the Status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedById orders the results by the CreatedById field.
func ByCreatedById(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedById, opts...).ToFunc()
}

// ByCreatedByUserName orders the results by the CreatedByUserName field.
func ByCreatedByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedByUserName, opts...).ToFunc()
}

// ByCreatedByEmpId orders the results by the CreatedByEmpId field.
func ByCreatedByEmpId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedByEmpId, opts...).ToFunc()
}

// ByCreatedByDesignation orders the results by the CreatedByDesignation field.
func ByCreatedByDesignation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedByDesignation, opts...).ToFunc()
}

// ByCreatedDate orders the results by the CreatedDate field.
func ByCreatedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedDate, opts...).ToFunc()
}

// ByModifiedById orders the results by the ModifiedById field.
func ByModifiedById(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedById, opts...).ToFunc()
}

// ByModifiedByUserName orders the results by the ModifiedByUserName field.
func ByModifiedByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedByUserName, opts...).ToFunc()
}

// ByModifiedByEmpId orders the results by the ModifiedByEmpId field.
func ByModifiedByEmpId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedByEmpId, opts...).ToFunc()
}

// ByModifiedByDesignantion orders the results by the ModifiedByDesignantion field.
func ByModifiedByDesignantion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedByDesignantion, opts...).ToFunc()
}

// ByModifiedDate orders the results by the ModifiedDate field.
func ByModifiedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedDate, opts...).ToFunc()
}

// ByDeletedById orders the results by the DeletedById field.
func ByDeletedById(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedById, opts...).ToFunc()
}

// ByDeletedByUserName orders the results by the DeletedByUserName field.
func ByDeletedByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedByUserName, opts...).ToFunc()
}

// ByDeletedByEmpId orders the results by the DeletedByEmpId field.
func ByDeletedByEmpId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedByEmpId, opts...).ToFunc()
}

// ByDeletedByDesignation orders the results by the DeletedByDesignation field.
func ByDeletedByDesignation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedByDesignation, opts...).ToFunc()
}

// ByDeletedDate orders the results by the DeletedDate field.
func ByDeletedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedDate, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByNoSeats orders the results by the NoSeats field.
func ByNoSeats(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNoSeats, opts...).ToFunc()
}

// ByExaCenterHallField orders the results by ExaCenterHall field.
func ByExaCenterHallField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExaCenterHallStep(), sql.OrderByField(field, opts...))
	}
}
func newExaCenterHallStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExaCenterHallInverseTable, CenterFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ExaCenterHallTable, ExaCenterHallColumn),
	)
}
