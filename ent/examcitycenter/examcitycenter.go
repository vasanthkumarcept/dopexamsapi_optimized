// Code generated by ent, DO NOT EDIT.

package examcitycenter

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the examcitycenter type in the database.
	Label = "exam_city_center"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "ExamCityCentreCode"
	// FieldExamCode holds the string denoting the examcode field in the database.
	FieldExamCode = "exam_code"
	// FieldExamName holds the string denoting the examname field in the database.
	FieldExamName = "exam_name"
	// FieldExamShortName holds the string denoting the examshortname field in the database.
	FieldExamShortName = "exam_short_name"
	// FieldExamYear holds the string denoting the examyear field in the database.
	FieldExamYear = "exam_year"
	// FieldConductedBy holds the string denoting the conductedby field in the database.
	FieldConductedBy = "conducted_by"
	// FieldNodalOfficeFacilityID holds the string denoting the nodalofficefacilityid field in the database.
	FieldNodalOfficeFacilityID = "nodal_office_facility_id"
	// FieldNodalOfficeName holds the string denoting the nodalofficename field in the database.
	FieldNodalOfficeName = "nodal_office_name"
	// FieldNotificationCode holds the string denoting the notificationcode field in the database.
	FieldNotificationCode = "notification_code"
	// FieldNotificationNumber holds the string denoting the notificationnumber field in the database.
	FieldNotificationNumber = "notification_number"
	// FieldCenterCityName holds the string denoting the centercityname field in the database.
	FieldCenterCityName = "center_city_name"
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
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDeletedbyid holds the string denoting the deletedbyid field in the database.
	FieldDeletedbyid = "deletedbyid"
	// FieldDeletedbyusername holds the string denoting the deletedbyusername field in the database.
	FieldDeletedbyusername = "deletedbyusername"
	// FieldDeletedbyEmployeeid holds the string denoting the deletedbyemployeeid field in the database.
	FieldDeletedbyEmployeeid = "deletedby_employeeid"
	// FieldDeletedbyDesignation holds the string denoting the deletedbydesignation field in the database.
	FieldDeletedbyDesignation = "deletedby_designation"
	// FieldDeletedDate holds the string denoting the deleteddate field in the database.
	FieldDeletedDate = "deleted_date"
	// FieldCircleCityName holds the string denoting the circlecityname field in the database.
	FieldCircleCityName = "circle_city_name"
	// FieldDivisionCode holds the string denoting the divisioncode field in the database.
	FieldDivisionCode = "division_code"
	// FieldRegionCode holds the string denoting the regioncode field in the database.
	FieldRegionCode = "region_code"
	// FieldDivisionName holds the string denoting the divisionname field in the database.
	FieldDivisionName = "division_name"
	// FieldRegionID holds the string denoting the regionid field in the database.
	FieldRegionID = "region_id"
	// FieldRegionName holds the string denoting the regionname field in the database.
	FieldRegionName = "region_name"
	// FieldRegionCityName holds the string denoting the regioncityname field in the database.
	FieldRegionCityName = "region_city_name"
	// FieldCentreCityName holds the string denoting the centrecityname field in the database.
	FieldCentreCityName = "centre_city_name"
	// FieldRemarks holds the string denoting the remarks field in the database.
	FieldRemarks = "remarks"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updatedby field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCentreCode holds the string denoting the centrecode field in the database.
	FieldCentreCode = "centre_code"
	// FieldCircleID holds the string denoting the circleid field in the database.
	FieldCircleID = "circle_id"
	// EdgeExamCityCenterRef holds the string denoting the examcitycenterref edge name in mutations.
	EdgeExamCityCenterRef = "ExamCityCenterRef"
	// EdgeExamCityCenterMTSPMMGRef holds the string denoting the examcitycentermtspmmgref edge name in mutations.
	EdgeExamCityCenterMTSPMMGRef = "ExamCityCenterMTSPMMGRef"
	// EdgeExamCityCenterGDSPARef holds the string denoting the examcitycentergdsparef edge name in mutations.
	EdgeExamCityCenterGDSPARef = "ExamCityCenterGDSPARef"
	// EdgeExamCityCenterGDSPMRef holds the string denoting the examcitycentergdspmref edge name in mutations.
	EdgeExamCityCenterGDSPMRef = "ExamCityCenterGDSPMRef"
	// EdgeExamCityCenterPMPARef holds the string denoting the examcitycenterpmparef edge name in mutations.
	EdgeExamCityCenterPMPARef = "ExamCityCenterPMPARef"
	// EdgeExamCityCenterPSRef holds the string denoting the examcitycenterpsref edge name in mutations.
	EdgeExamCityCenterPSRef = "ExamCityCenterPSRef"
	// Exam_Applications_IPFieldID holds the string denoting the ID field of the Exam_Applications_IP.
	Exam_Applications_IPFieldID = "ApplicationID"
	// Exam_Application_MTSPMMGFieldID holds the string denoting the ID field of the Exam_Application_MTSPMMG.
	Exam_Application_MTSPMMGFieldID = "ApplicationID"
	// Exam_Applications_GDSPAFieldID holds the string denoting the ID field of the Exam_Applications_GDSPA.
	Exam_Applications_GDSPAFieldID = "ApplicationID"
	// Exam_Applications_GDSPMFieldID holds the string denoting the ID field of the Exam_Applications_GDSPM.
	Exam_Applications_GDSPMFieldID = "ApplicationID"
	// Exam_Applications_PMPAFieldID holds the string denoting the ID field of the Exam_Applications_PMPA.
	Exam_Applications_PMPAFieldID = "ApplicationID"
	// Exam_Applications_PSFieldID holds the string denoting the ID field of the Exam_Applications_PS.
	Exam_Applications_PSFieldID = "ApplicationID"
	// Table holds the table name of the examcitycenter in the database.
	Table = "ExamCityCenter"
	// ExamCityCenterRefTable is the table that holds the ExamCityCenterRef relation/edge.
	ExamCityCenterRefTable = "Exam_Applications_IP"
	// ExamCityCenterRefInverseTable is the table name for the Exam_Applications_IP entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_ip" package.
	ExamCityCenterRefInverseTable = "Exam_Applications_IP"
	// ExamCityCenterRefColumn is the table column denoting the ExamCityCenterRef relation/edge.
	ExamCityCenterRefColumn = "exam_city_center_code"
	// ExamCityCenterMTSPMMGRefTable is the table that holds the ExamCityCenterMTSPMMGRef relation/edge.
	ExamCityCenterMTSPMMGRefTable = "Exam_Application_MTSPMMG"
	// ExamCityCenterMTSPMMGRefInverseTable is the table name for the Exam_Application_MTSPMMG entity.
	// It exists in this package in order to avoid circular dependency with the "exam_application_mtspmmg" package.
	ExamCityCenterMTSPMMGRefInverseTable = "Exam_Application_MTSPMMG"
	// ExamCityCenterMTSPMMGRefColumn is the table column denoting the ExamCityCenterMTSPMMGRef relation/edge.
	ExamCityCenterMTSPMMGRefColumn = "exam_city_center_code"
	// ExamCityCenterGDSPARefTable is the table that holds the ExamCityCenterGDSPARef relation/edge.
	ExamCityCenterGDSPARefTable = "Exam_Applications_GDSPA"
	// ExamCityCenterGDSPARefInverseTable is the table name for the Exam_Applications_GDSPA entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_gdspa" package.
	ExamCityCenterGDSPARefInverseTable = "Exam_Applications_GDSPA"
	// ExamCityCenterGDSPARefColumn is the table column denoting the ExamCityCenterGDSPARef relation/edge.
	ExamCityCenterGDSPARefColumn = "exam_city_center_code"
	// ExamCityCenterGDSPMRefTable is the table that holds the ExamCityCenterGDSPMRef relation/edge.
	ExamCityCenterGDSPMRefTable = "Exam_Applications_GDSPM"
	// ExamCityCenterGDSPMRefInverseTable is the table name for the Exam_Applications_GDSPM entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_gdspm" package.
	ExamCityCenterGDSPMRefInverseTable = "Exam_Applications_GDSPM"
	// ExamCityCenterGDSPMRefColumn is the table column denoting the ExamCityCenterGDSPMRef relation/edge.
	ExamCityCenterGDSPMRefColumn = "exam_city_center_code"
	// ExamCityCenterPMPARefTable is the table that holds the ExamCityCenterPMPARef relation/edge.
	ExamCityCenterPMPARefTable = "Exam_Applications_PMPA"
	// ExamCityCenterPMPARefInverseTable is the table name for the Exam_Applications_PMPA entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_pmpa" package.
	ExamCityCenterPMPARefInverseTable = "Exam_Applications_PMPA"
	// ExamCityCenterPMPARefColumn is the table column denoting the ExamCityCenterPMPARef relation/edge.
	ExamCityCenterPMPARefColumn = "exam_city_center_code"
	// ExamCityCenterPSRefTable is the table that holds the ExamCityCenterPSRef relation/edge.
	ExamCityCenterPSRefTable = "Exam_Applications_PS"
	// ExamCityCenterPSRefInverseTable is the table name for the Exam_Applications_PS entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_ps" package.
	ExamCityCenterPSRefInverseTable = "Exam_Applications_PS"
	// ExamCityCenterPSRefColumn is the table column denoting the ExamCityCenterPSRef relation/edge.
	ExamCityCenterPSRefColumn = "exam_city_center_code"
)

// Columns holds all SQL columns for examcitycenter fields.
var Columns = []string{
	FieldID,
	FieldExamCode,
	FieldExamName,
	FieldExamShortName,
	FieldExamYear,
	FieldConductedBy,
	FieldNodalOfficeFacilityID,
	FieldNodalOfficeName,
	FieldNotificationCode,
	FieldNotificationNumber,
	FieldCenterCityName,
	FieldCreatedById,
	FieldCreatedByUserName,
	FieldCreatedByEmpId,
	FieldCreatedByDesignation,
	FieldCreatedDate,
	FieldStatus,
	FieldDeletedbyid,
	FieldDeletedbyusername,
	FieldDeletedbyEmployeeid,
	FieldDeletedbyDesignation,
	FieldDeletedDate,
	FieldCircleCityName,
	FieldDivisionCode,
	FieldRegionCode,
	FieldDivisionName,
	FieldRegionID,
	FieldRegionName,
	FieldRegionCityName,
	FieldCentreCityName,
	FieldRemarks,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldCentreCode,
	FieldCircleID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ExamCityCenter"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"center_examscentres",
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
	// DefaultExamCode holds the default value on creation for the "ExamCode" field.
	DefaultExamCode int32
	// DefaultExamYear holds the default value on creation for the "ExamYear" field.
	DefaultExamYear int32
	// DefaultNotificationCode holds the default value on creation for the "NotificationCode" field.
	DefaultNotificationCode int32
	// DefaultCreatedById holds the default value on creation for the "CreatedById" field.
	DefaultCreatedById int64
	// DefaultCreatedByEmpId holds the default value on creation for the "CreatedByEmpId" field.
	DefaultCreatedByEmpId int64
	// DefaultDeletedbyid holds the default value on creation for the "deletedbyid" field.
	DefaultDeletedbyid int64
	// DefaultDeletedbyEmployeeid holds the default value on creation for the "deletedbyEmployeeid" field.
	DefaultDeletedbyEmployeeid int64
	// DefaultDivisionCode holds the default value on creation for the "DivisionCode" field.
	DefaultDivisionCode int32
	// DefaultRegionCode holds the default value on creation for the "RegionCode" field.
	DefaultRegionCode int32
	// DefaultCentreCode holds the default value on creation for the "CentreCode" field.
	DefaultCentreCode int32
	// DefaultCircleID holds the default value on creation for the "CircleID" field.
	DefaultCircleID int32
)

// OrderOption defines the ordering options for the ExamCityCenter queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByExamCode orders the results by the ExamCode field.
func ByExamCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamCode, opts...).ToFunc()
}

// ByExamName orders the results by the ExamName field.
func ByExamName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamName, opts...).ToFunc()
}

// ByExamShortName orders the results by the ExamShortName field.
func ByExamShortName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamShortName, opts...).ToFunc()
}

// ByExamYear orders the results by the ExamYear field.
func ByExamYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamYear, opts...).ToFunc()
}

// ByConductedBy orders the results by the ConductedBy field.
func ByConductedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConductedBy, opts...).ToFunc()
}

// ByNodalOfficeFacilityID orders the results by the NodalOfficeFacilityID field.
func ByNodalOfficeFacilityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodalOfficeFacilityID, opts...).ToFunc()
}

// ByNodalOfficeName orders the results by the NodalOfficeName field.
func ByNodalOfficeName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodalOfficeName, opts...).ToFunc()
}

// ByNotificationCode orders the results by the NotificationCode field.
func ByNotificationCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotificationCode, opts...).ToFunc()
}

// ByNotificationNumber orders the results by the NotificationNumber field.
func ByNotificationNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotificationNumber, opts...).ToFunc()
}

// ByCenterCityName orders the results by the CenterCityName field.
func ByCenterCityName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCenterCityName, opts...).ToFunc()
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

// ByStatus orders the results by the Status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDeletedbyid orders the results by the deletedbyid field.
func ByDeletedbyid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedbyid, opts...).ToFunc()
}

// ByDeletedbyusername orders the results by the deletedbyusername field.
func ByDeletedbyusername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedbyusername, opts...).ToFunc()
}

// ByDeletedbyEmployeeid orders the results by the deletedbyEmployeeid field.
func ByDeletedbyEmployeeid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedbyEmployeeid, opts...).ToFunc()
}

// ByDeletedbyDesignation orders the results by the deletedbyDesignation field.
func ByDeletedbyDesignation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedbyDesignation, opts...).ToFunc()
}

// ByDeletedDate orders the results by the deletedDate field.
func ByDeletedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedDate, opts...).ToFunc()
}

// ByCircleCityName orders the results by the CircleCityName field.
func ByCircleCityName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCircleCityName, opts...).ToFunc()
}

// ByDivisionCode orders the results by the DivisionCode field.
func ByDivisionCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDivisionCode, opts...).ToFunc()
}

// ByRegionCode orders the results by the RegionCode field.
func ByRegionCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionCode, opts...).ToFunc()
}

// ByDivisionName orders the results by the DivisionName field.
func ByDivisionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDivisionName, opts...).ToFunc()
}

// ByRegionID orders the results by the RegionID field.
func ByRegionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionID, opts...).ToFunc()
}

// ByRegionName orders the results by the RegionName field.
func ByRegionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionName, opts...).ToFunc()
}

// ByRegionCityName orders the results by the RegionCityName field.
func ByRegionCityName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionCityName, opts...).ToFunc()
}

// ByCentreCityName orders the results by the CentreCityName field.
func ByCentreCityName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCentreCityName, opts...).ToFunc()
}

// ByRemarks orders the results by the Remarks field.
func ByRemarks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemarks, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the UpdatedBy field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCentreCode orders the results by the CentreCode field.
func ByCentreCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCentreCode, opts...).ToFunc()
}

// ByCircleID orders the results by the CircleID field.
func ByCircleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCircleID, opts...).ToFunc()
}

// ByExamCityCenterRefCount orders the results by ExamCityCenterRef count.
func ByExamCityCenterRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamCityCenterRefStep(), opts...)
	}
}

// ByExamCityCenterRef orders the results by ExamCityCenterRef terms.
func ByExamCityCenterRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamCityCenterRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExamCityCenterMTSPMMGRefCount orders the results by ExamCityCenterMTSPMMGRef count.
func ByExamCityCenterMTSPMMGRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamCityCenterMTSPMMGRefStep(), opts...)
	}
}

// ByExamCityCenterMTSPMMGRef orders the results by ExamCityCenterMTSPMMGRef terms.
func ByExamCityCenterMTSPMMGRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamCityCenterMTSPMMGRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExamCityCenterGDSPARefCount orders the results by ExamCityCenterGDSPARef count.
func ByExamCityCenterGDSPARefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamCityCenterGDSPARefStep(), opts...)
	}
}

// ByExamCityCenterGDSPARef orders the results by ExamCityCenterGDSPARef terms.
func ByExamCityCenterGDSPARef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamCityCenterGDSPARefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExamCityCenterGDSPMRefCount orders the results by ExamCityCenterGDSPMRef count.
func ByExamCityCenterGDSPMRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamCityCenterGDSPMRefStep(), opts...)
	}
}

// ByExamCityCenterGDSPMRef orders the results by ExamCityCenterGDSPMRef terms.
func ByExamCityCenterGDSPMRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamCityCenterGDSPMRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExamCityCenterPMPARefCount orders the results by ExamCityCenterPMPARef count.
func ByExamCityCenterPMPARefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamCityCenterPMPARefStep(), opts...)
	}
}

// ByExamCityCenterPMPARef orders the results by ExamCityCenterPMPARef terms.
func ByExamCityCenterPMPARef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamCityCenterPMPARefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExamCityCenterPSRefCount orders the results by ExamCityCenterPSRef count.
func ByExamCityCenterPSRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamCityCenterPSRefStep(), opts...)
	}
}

// ByExamCityCenterPSRef orders the results by ExamCityCenterPSRef terms.
func ByExamCityCenterPSRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamCityCenterPSRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newExamCityCenterRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamCityCenterRefInverseTable, Exam_Applications_IPFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamCityCenterRefTable, ExamCityCenterRefColumn),
	)
}
func newExamCityCenterMTSPMMGRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamCityCenterMTSPMMGRefInverseTable, Exam_Application_MTSPMMGFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamCityCenterMTSPMMGRefTable, ExamCityCenterMTSPMMGRefColumn),
	)
}
func newExamCityCenterGDSPARefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamCityCenterGDSPARefInverseTable, Exam_Applications_GDSPAFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamCityCenterGDSPARefTable, ExamCityCenterGDSPARefColumn),
	)
}
func newExamCityCenterGDSPMRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamCityCenterGDSPMRefInverseTable, Exam_Applications_GDSPMFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamCityCenterGDSPMRefTable, ExamCityCenterGDSPMRefColumn),
	)
}
func newExamCityCenterPMPARefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamCityCenterPMPARefInverseTable, Exam_Applications_PMPAFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamCityCenterPMPARefTable, ExamCityCenterPMPARefColumn),
	)
}
func newExamCityCenterPSRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamCityCenterPSRefInverseTable, Exam_Applications_PSFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamCityCenterPSRefTable, ExamCityCenterPSRefColumn),
	)
}
