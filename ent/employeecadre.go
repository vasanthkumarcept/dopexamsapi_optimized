// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/employeecadre"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// EmployeeCadre is the model entity for the EmployeeCadre schema.
type EmployeeCadre struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Cadrecode holds the value of the "cadrecode" field.
	Cadrecode string `json:"cadrecode,omitempty"`
	// Cadredescription holds the value of the "cadredescription" field.
	Cadredescription string `json:"cadredescription,omitempty"`
	// PayLevel holds the value of the "PayLevel" field.
	PayLevel string `json:"PayLevel,omitempty"`
	// Scale holds the value of the "Scale" field.
	Scale string `json:"Scale,omitempty"`
	// ExamconfigurationExamcode holds the value of the "ExamconfigurationExamcode" field.
	ExamconfigurationExamcode int32 `json:"ExamconfigurationExamcode,omitempty"`
	// ExamShortDescription holds the value of the "ExamShortDescription" field.
	ExamShortDescription string `json:"ExamShortDescription,omitempty"`
	// ExamLongDescription holds the value of the "ExamLongDescription" field.
	ExamLongDescription string `json:"ExamLongDescription,omitempty"`
	// EmployeePostPostId holds the value of the "EmployeePost_postId" field.
	EmployeePostPostId int32 `json:"EmployeePost_postId,omitempty"`
	// EmployeeGroupGroupId holds the value of the "EmployeeGroup_groupId" field.
	EmployeeGroupGroupId int32 `json:"EmployeeGroup_groupId,omitempty"`
	// GroupDescription holds the value of the "GroupDescription" field.
	GroupDescription string `json:"GroupDescription,omitempty"`
	// PostCode holds the value of the "PostCode" field.
	PostCode string `json:"PostCode,omitempty"`
	// PostDescription holds the value of the "PostDescription" field.
	PostDescription string `json:"PostDescription,omitempty"`
	// BaseCadre holds the value of the "BaseCadre" field.
	BaseCadre int32 `json:"BaseCadre,omitempty"`
	// GdsService holds the value of the "GdsService" field.
	GdsService int32 `json:"GdsService,omitempty"`
	// AgeCriteria holds the value of the "ageCriteria" field.
	AgeCriteria int32 `json:"ageCriteria,omitempty"`
	// ServiceCriteria holds the value of the "ServiceCriteria" field.
	ServiceCriteria int32 `json:"ServiceCriteria,omitempty"`
	// DrivingLicenceCriteria holds the value of the "DrivingLicenceCriteria" field.
	DrivingLicenceCriteria int32 `json:"DrivingLicenceCriteria,omitempty"`
	// ComputerKnowledge holds the value of the "ComputerKnowledge" field.
	ComputerKnowledge int32 `json:"ComputerKnowledge,omitempty"`
	// EligibiltyBasedOnLevelOfPaymatrix holds the value of the "EligibiltyBasedOnLevelOfPaymatrix" field.
	EligibiltyBasedOnLevelOfPaymatrix int32 `json:"EligibiltyBasedOnLevelOfPaymatrix,omitempty"`
	// EducationDetailsEducationCode holds the value of the "EducationDetails_educationCode" field.
	EducationDetailsEducationCode int32 `json:"EducationDetails_educationCode,omitempty"`
	// EducationDescription holds the value of the "EducationDescription" field.
	EducationDescription string `json:"EducationDescription,omitempty"`
	// OrderNumber holds the value of the "OrderNumber" field.
	OrderNumber string `json:"OrderNumber,omitempty"`
	// Status holds the value of the "Status" field.
	Status string `json:"Status,omitempty"`
	// CreatedById holds the value of the "CreatedById" field.
	CreatedById int32 `json:"CreatedById,omitempty"`
	// CreatedByUserName holds the value of the "CreatedByUserName" field.
	CreatedByUserName string `json:"CreatedByUserName,omitempty"`
	// CreatedByEmployeeId holds the value of the "CreatedByEmployeeId" field.
	CreatedByEmployeeId string `json:"CreatedByEmployeeId,omitempty"`
	// CreatedByDesignation holds the value of the "CreatedByDesignation" field.
	CreatedByDesignation string `json:"CreatedByDesignation,omitempty"`
	// CreatedDate holds the value of the "CreatedDate" field.
	CreatedDate time.Time `json:"CreatedDate,omitempty"`
	// VerifiedById holds the value of the "VerifiedById" field.
	VerifiedById int64 `json:"VerifiedById,omitempty"`
	// VerifiedByUserName holds the value of the "VerifiedByUserName" field.
	VerifiedByUserName string `json:"VerifiedByUserName,omitempty"`
	// VerifiedByEmployeeId holds the value of the "VerifiedByEmployeeId" field.
	VerifiedByEmployeeId string `json:"VerifiedByEmployeeId,omitempty"`
	// VerifiedByDesignation holds the value of the "VerifiedByDesignation" field.
	VerifiedByDesignation string `json:"VerifiedByDesignation,omitempty"`
	// VerifiedDate holds the value of the "VerifiedDate" field.
	VerifiedDate time.Time `json:"VerifiedDate,omitempty"`
	// DeletedById holds the value of the "DeletedById" field.
	DeletedById int64 `json:"DeletedById,omitempty"`
	// DeletedByUserName holds the value of the "DeletedByUserName" field.
	DeletedByUserName string `json:"DeletedByUserName,omitempty"`
	// DeletedByEmployeeId holds the value of the "DeletedByEmployeeId" field.
	DeletedByEmployeeId string `json:"DeletedByEmployeeId,omitempty"`
	// DeletedByDesignation holds the value of the "DeletedByDesignation" field.
	DeletedByDesignation string `json:"DeletedByDesignation,omitempty"`
	// DeletedDate holds the value of the "DeletedDate" field.
	DeletedDate  time.Time `json:"DeletedDate,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EmployeeCadre) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case employeecadre.FieldID, employeecadre.FieldExamconfigurationExamcode, employeecadre.FieldEmployeePostPostId, employeecadre.FieldEmployeeGroupGroupId, employeecadre.FieldBaseCadre, employeecadre.FieldGdsService, employeecadre.FieldAgeCriteria, employeecadre.FieldServiceCriteria, employeecadre.FieldDrivingLicenceCriteria, employeecadre.FieldComputerKnowledge, employeecadre.FieldEligibiltyBasedOnLevelOfPaymatrix, employeecadre.FieldEducationDetailsEducationCode, employeecadre.FieldCreatedById, employeecadre.FieldVerifiedById, employeecadre.FieldDeletedById:
			values[i] = new(sql.NullInt64)
		case employeecadre.FieldCadrecode, employeecadre.FieldCadredescription, employeecadre.FieldPayLevel, employeecadre.FieldScale, employeecadre.FieldExamShortDescription, employeecadre.FieldExamLongDescription, employeecadre.FieldGroupDescription, employeecadre.FieldPostCode, employeecadre.FieldPostDescription, employeecadre.FieldEducationDescription, employeecadre.FieldOrderNumber, employeecadre.FieldStatus, employeecadre.FieldCreatedByUserName, employeecadre.FieldCreatedByEmployeeId, employeecadre.FieldCreatedByDesignation, employeecadre.FieldVerifiedByUserName, employeecadre.FieldVerifiedByEmployeeId, employeecadre.FieldVerifiedByDesignation, employeecadre.FieldDeletedByUserName, employeecadre.FieldDeletedByEmployeeId, employeecadre.FieldDeletedByDesignation:
			values[i] = new(sql.NullString)
		case employeecadre.FieldCreatedDate, employeecadre.FieldVerifiedDate, employeecadre.FieldDeletedDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EmployeeCadre fields.
func (ec *EmployeeCadre) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employeecadre.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = int32(value.Int64)
		case employeecadre.FieldCadrecode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cadrecode", values[i])
			} else if value.Valid {
				ec.Cadrecode = value.String
			}
		case employeecadre.FieldCadredescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cadredescription", values[i])
			} else if value.Valid {
				ec.Cadredescription = value.String
			}
		case employeecadre.FieldPayLevel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PayLevel", values[i])
			} else if value.Valid {
				ec.PayLevel = value.String
			}
		case employeecadre.FieldScale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Scale", values[i])
			} else if value.Valid {
				ec.Scale = value.String
			}
		case employeecadre.FieldExamconfigurationExamcode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ExamconfigurationExamcode", values[i])
			} else if value.Valid {
				ec.ExamconfigurationExamcode = int32(value.Int64)
			}
		case employeecadre.FieldExamShortDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamShortDescription", values[i])
			} else if value.Valid {
				ec.ExamShortDescription = value.String
			}
		case employeecadre.FieldExamLongDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamLongDescription", values[i])
			} else if value.Valid {
				ec.ExamLongDescription = value.String
			}
		case employeecadre.FieldEmployeePostPostId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeePost_postId", values[i])
			} else if value.Valid {
				ec.EmployeePostPostId = int32(value.Int64)
			}
		case employeecadre.FieldEmployeeGroupGroupId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeeGroup_groupId", values[i])
			} else if value.Valid {
				ec.EmployeeGroupGroupId = int32(value.Int64)
			}
		case employeecadre.FieldGroupDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field GroupDescription", values[i])
			} else if value.Valid {
				ec.GroupDescription = value.String
			}
		case employeecadre.FieldPostCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PostCode", values[i])
			} else if value.Valid {
				ec.PostCode = value.String
			}
		case employeecadre.FieldPostDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PostDescription", values[i])
			} else if value.Valid {
				ec.PostDescription = value.String
			}
		case employeecadre.FieldBaseCadre:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field BaseCadre", values[i])
			} else if value.Valid {
				ec.BaseCadre = int32(value.Int64)
			}
		case employeecadre.FieldGdsService:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field GdsService", values[i])
			} else if value.Valid {
				ec.GdsService = int32(value.Int64)
			}
		case employeecadre.FieldAgeCriteria:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ageCriteria", values[i])
			} else if value.Valid {
				ec.AgeCriteria = int32(value.Int64)
			}
		case employeecadre.FieldServiceCriteria:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ServiceCriteria", values[i])
			} else if value.Valid {
				ec.ServiceCriteria = int32(value.Int64)
			}
		case employeecadre.FieldDrivingLicenceCriteria:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field DrivingLicenceCriteria", values[i])
			} else if value.Valid {
				ec.DrivingLicenceCriteria = int32(value.Int64)
			}
		case employeecadre.FieldComputerKnowledge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ComputerKnowledge", values[i])
			} else if value.Valid {
				ec.ComputerKnowledge = int32(value.Int64)
			}
		case employeecadre.FieldEligibiltyBasedOnLevelOfPaymatrix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EligibiltyBasedOnLevelOfPaymatrix", values[i])
			} else if value.Valid {
				ec.EligibiltyBasedOnLevelOfPaymatrix = int32(value.Int64)
			}
		case employeecadre.FieldEducationDetailsEducationCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EducationDetails_educationCode", values[i])
			} else if value.Valid {
				ec.EducationDetailsEducationCode = int32(value.Int64)
			}
		case employeecadre.FieldEducationDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EducationDescription", values[i])
			} else if value.Valid {
				ec.EducationDescription = value.String
			}
		case employeecadre.FieldOrderNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field OrderNumber", values[i])
			} else if value.Valid {
				ec.OrderNumber = value.String
			}
		case employeecadre.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				ec.Status = value.String
			}
		case employeecadre.FieldCreatedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedById", values[i])
			} else if value.Valid {
				ec.CreatedById = int32(value.Int64)
			}
		case employeecadre.FieldCreatedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByUserName", values[i])
			} else if value.Valid {
				ec.CreatedByUserName = value.String
			}
		case employeecadre.FieldCreatedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByEmployeeId", values[i])
			} else if value.Valid {
				ec.CreatedByEmployeeId = value.String
			}
		case employeecadre.FieldCreatedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByDesignation", values[i])
			} else if value.Valid {
				ec.CreatedByDesignation = value.String
			}
		case employeecadre.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedDate", values[i])
			} else if value.Valid {
				ec.CreatedDate = value.Time
			}
		case employeecadre.FieldVerifiedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedById", values[i])
			} else if value.Valid {
				ec.VerifiedById = value.Int64
			}
		case employeecadre.FieldVerifiedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedByUserName", values[i])
			} else if value.Valid {
				ec.VerifiedByUserName = value.String
			}
		case employeecadre.FieldVerifiedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedByEmployeeId", values[i])
			} else if value.Valid {
				ec.VerifiedByEmployeeId = value.String
			}
		case employeecadre.FieldVerifiedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedByDesignation", values[i])
			} else if value.Valid {
				ec.VerifiedByDesignation = value.String
			}
		case employeecadre.FieldVerifiedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedDate", values[i])
			} else if value.Valid {
				ec.VerifiedDate = value.Time
			}
		case employeecadre.FieldDeletedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedById", values[i])
			} else if value.Valid {
				ec.DeletedById = value.Int64
			}
		case employeecadre.FieldDeletedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedByUserName", values[i])
			} else if value.Valid {
				ec.DeletedByUserName = value.String
			}
		case employeecadre.FieldDeletedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedByEmployeeId", values[i])
			} else if value.Valid {
				ec.DeletedByEmployeeId = value.String
			}
		case employeecadre.FieldDeletedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedByDesignation", values[i])
			} else if value.Valid {
				ec.DeletedByDesignation = value.String
			}
		case employeecadre.FieldDeletedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedDate", values[i])
			} else if value.Valid {
				ec.DeletedDate = value.Time
			}
		default:
			ec.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EmployeeCadre.
// This includes values selected through modifiers, order, etc.
func (ec *EmployeeCadre) Value(name string) (ent.Value, error) {
	return ec.selectValues.Get(name)
}

// Update returns a builder for updating this EmployeeCadre.
// Note that you need to call EmployeeCadre.Unwrap() before calling this method if this EmployeeCadre
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *EmployeeCadre) Update() *EmployeeCadreUpdateOne {
	return NewEmployeeCadreClient(ec.config).UpdateOne(ec)
}

// Unwrap unwraps the EmployeeCadre entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *EmployeeCadre) Unwrap() *EmployeeCadre {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: EmployeeCadre is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *EmployeeCadre) String() string {
	var builder strings.Builder
	builder.WriteString("EmployeeCadre(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ec.ID))
	builder.WriteString("cadrecode=")
	builder.WriteString(ec.Cadrecode)
	builder.WriteString(", ")
	builder.WriteString("cadredescription=")
	builder.WriteString(ec.Cadredescription)
	builder.WriteString(", ")
	builder.WriteString("PayLevel=")
	builder.WriteString(ec.PayLevel)
	builder.WriteString(", ")
	builder.WriteString("Scale=")
	builder.WriteString(ec.Scale)
	builder.WriteString(", ")
	builder.WriteString("ExamconfigurationExamcode=")
	builder.WriteString(fmt.Sprintf("%v", ec.ExamconfigurationExamcode))
	builder.WriteString(", ")
	builder.WriteString("ExamShortDescription=")
	builder.WriteString(ec.ExamShortDescription)
	builder.WriteString(", ")
	builder.WriteString("ExamLongDescription=")
	builder.WriteString(ec.ExamLongDescription)
	builder.WriteString(", ")
	builder.WriteString("EmployeePost_postId=")
	builder.WriteString(fmt.Sprintf("%v", ec.EmployeePostPostId))
	builder.WriteString(", ")
	builder.WriteString("EmployeeGroup_groupId=")
	builder.WriteString(fmt.Sprintf("%v", ec.EmployeeGroupGroupId))
	builder.WriteString(", ")
	builder.WriteString("GroupDescription=")
	builder.WriteString(ec.GroupDescription)
	builder.WriteString(", ")
	builder.WriteString("PostCode=")
	builder.WriteString(ec.PostCode)
	builder.WriteString(", ")
	builder.WriteString("PostDescription=")
	builder.WriteString(ec.PostDescription)
	builder.WriteString(", ")
	builder.WriteString("BaseCadre=")
	builder.WriteString(fmt.Sprintf("%v", ec.BaseCadre))
	builder.WriteString(", ")
	builder.WriteString("GdsService=")
	builder.WriteString(fmt.Sprintf("%v", ec.GdsService))
	builder.WriteString(", ")
	builder.WriteString("ageCriteria=")
	builder.WriteString(fmt.Sprintf("%v", ec.AgeCriteria))
	builder.WriteString(", ")
	builder.WriteString("ServiceCriteria=")
	builder.WriteString(fmt.Sprintf("%v", ec.ServiceCriteria))
	builder.WriteString(", ")
	builder.WriteString("DrivingLicenceCriteria=")
	builder.WriteString(fmt.Sprintf("%v", ec.DrivingLicenceCriteria))
	builder.WriteString(", ")
	builder.WriteString("ComputerKnowledge=")
	builder.WriteString(fmt.Sprintf("%v", ec.ComputerKnowledge))
	builder.WriteString(", ")
	builder.WriteString("EligibiltyBasedOnLevelOfPaymatrix=")
	builder.WriteString(fmt.Sprintf("%v", ec.EligibiltyBasedOnLevelOfPaymatrix))
	builder.WriteString(", ")
	builder.WriteString("EducationDetails_educationCode=")
	builder.WriteString(fmt.Sprintf("%v", ec.EducationDetailsEducationCode))
	builder.WriteString(", ")
	builder.WriteString("EducationDescription=")
	builder.WriteString(ec.EducationDescription)
	builder.WriteString(", ")
	builder.WriteString("OrderNumber=")
	builder.WriteString(ec.OrderNumber)
	builder.WriteString(", ")
	builder.WriteString("Status=")
	builder.WriteString(ec.Status)
	builder.WriteString(", ")
	builder.WriteString("CreatedById=")
	builder.WriteString(fmt.Sprintf("%v", ec.CreatedById))
	builder.WriteString(", ")
	builder.WriteString("CreatedByUserName=")
	builder.WriteString(ec.CreatedByUserName)
	builder.WriteString(", ")
	builder.WriteString("CreatedByEmployeeId=")
	builder.WriteString(ec.CreatedByEmployeeId)
	builder.WriteString(", ")
	builder.WriteString("CreatedByDesignation=")
	builder.WriteString(ec.CreatedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("CreatedDate=")
	builder.WriteString(ec.CreatedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("VerifiedById=")
	builder.WriteString(fmt.Sprintf("%v", ec.VerifiedById))
	builder.WriteString(", ")
	builder.WriteString("VerifiedByUserName=")
	builder.WriteString(ec.VerifiedByUserName)
	builder.WriteString(", ")
	builder.WriteString("VerifiedByEmployeeId=")
	builder.WriteString(ec.VerifiedByEmployeeId)
	builder.WriteString(", ")
	builder.WriteString("VerifiedByDesignation=")
	builder.WriteString(ec.VerifiedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("VerifiedDate=")
	builder.WriteString(ec.VerifiedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("DeletedById=")
	builder.WriteString(fmt.Sprintf("%v", ec.DeletedById))
	builder.WriteString(", ")
	builder.WriteString("DeletedByUserName=")
	builder.WriteString(ec.DeletedByUserName)
	builder.WriteString(", ")
	builder.WriteString("DeletedByEmployeeId=")
	builder.WriteString(ec.DeletedByEmployeeId)
	builder.WriteString(", ")
	builder.WriteString("DeletedByDesignation=")
	builder.WriteString(ec.DeletedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("DeletedDate=")
	builder.WriteString(ec.DeletedDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// EmployeeCadres is a parsable slice of EmployeeCadre.
type EmployeeCadres []*EmployeeCadre
