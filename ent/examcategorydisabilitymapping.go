// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/examcategorydisabilitymapping"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ExamCategoryDisabilityMapping is the model entity for the ExamCategoryDisabilityMapping schema.
type ExamCategoryDisabilityMapping struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// ExamCode holds the value of the "ExamCode" field.
	ExamCode int64 `json:"ExamCode,omitempty"`
	// ExamShortName holds the value of the "ExamShortName" field.
	ExamShortName string `json:"ExamShortName,omitempty"`
	// CategoryDisability holds the value of the "CategoryDisability" field.
	CategoryDisability string `json:"CategoryDisability,omitempty"`
	// CategoryDisabilityCode holds the value of the "CategoryDisabilityCode" field.
	CategoryDisabilityCode string `json:"CategoryDisabilityCode,omitempty"`
	// CategoryDisabilityDescription holds the value of the "CategoryDisabilityDescription" field.
	CategoryDisabilityDescription string `json:"CategoryDisabilityDescription,omitempty"`
	// AgeException holds the value of the "AgeException" field.
	AgeException int32 `json:"AgeException,omitempty"`
	// ServiceException holds the value of the "ServiceException" field.
	ServiceException int32 `json:"ServiceException,omitempty"`
	// DrivingLicense holds the value of the "DrivingLicense" field.
	DrivingLicense bool `json:"DrivingLicense,omitempty"`
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
func (*ExamCategoryDisabilityMapping) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case examcategorydisabilitymapping.FieldDrivingLicense:
			values[i] = new(sql.NullBool)
		case examcategorydisabilitymapping.FieldID, examcategorydisabilitymapping.FieldExamCode, examcategorydisabilitymapping.FieldAgeException, examcategorydisabilitymapping.FieldServiceException, examcategorydisabilitymapping.FieldCreatedById, examcategorydisabilitymapping.FieldVerifiedById, examcategorydisabilitymapping.FieldDeletedById:
			values[i] = new(sql.NullInt64)
		case examcategorydisabilitymapping.FieldExamShortName, examcategorydisabilitymapping.FieldCategoryDisability, examcategorydisabilitymapping.FieldCategoryDisabilityCode, examcategorydisabilitymapping.FieldCategoryDisabilityDescription, examcategorydisabilitymapping.FieldOrderNumber, examcategorydisabilitymapping.FieldStatus, examcategorydisabilitymapping.FieldCreatedByUserName, examcategorydisabilitymapping.FieldCreatedByEmployeeId, examcategorydisabilitymapping.FieldCreatedByDesignation, examcategorydisabilitymapping.FieldVerifiedByUserName, examcategorydisabilitymapping.FieldVerifiedByEmployeeId, examcategorydisabilitymapping.FieldVerifiedByDesignation, examcategorydisabilitymapping.FieldDeletedByUserName, examcategorydisabilitymapping.FieldDeletedByEmployeeId, examcategorydisabilitymapping.FieldDeletedByDesignation:
			values[i] = new(sql.NullString)
		case examcategorydisabilitymapping.FieldCreatedDate, examcategorydisabilitymapping.FieldVerifiedDate, examcategorydisabilitymapping.FieldDeletedDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExamCategoryDisabilityMapping fields.
func (ecdm *ExamCategoryDisabilityMapping) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case examcategorydisabilitymapping.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ecdm.ID = int32(value.Int64)
		case examcategorydisabilitymapping.FieldExamCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ExamCode", values[i])
			} else if value.Valid {
				ecdm.ExamCode = value.Int64
			}
		case examcategorydisabilitymapping.FieldExamShortName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamShortName", values[i])
			} else if value.Valid {
				ecdm.ExamShortName = value.String
			}
		case examcategorydisabilitymapping.FieldCategoryDisability:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CategoryDisability", values[i])
			} else if value.Valid {
				ecdm.CategoryDisability = value.String
			}
		case examcategorydisabilitymapping.FieldCategoryDisabilityCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CategoryDisabilityCode", values[i])
			} else if value.Valid {
				ecdm.CategoryDisabilityCode = value.String
			}
		case examcategorydisabilitymapping.FieldCategoryDisabilityDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CategoryDisabilityDescription", values[i])
			} else if value.Valid {
				ecdm.CategoryDisabilityDescription = value.String
			}
		case examcategorydisabilitymapping.FieldAgeException:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field AgeException", values[i])
			} else if value.Valid {
				ecdm.AgeException = int32(value.Int64)
			}
		case examcategorydisabilitymapping.FieldServiceException:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ServiceException", values[i])
			} else if value.Valid {
				ecdm.ServiceException = int32(value.Int64)
			}
		case examcategorydisabilitymapping.FieldDrivingLicense:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field DrivingLicense", values[i])
			} else if value.Valid {
				ecdm.DrivingLicense = value.Bool
			}
		case examcategorydisabilitymapping.FieldOrderNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field OrderNumber", values[i])
			} else if value.Valid {
				ecdm.OrderNumber = value.String
			}
		case examcategorydisabilitymapping.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				ecdm.Status = value.String
			}
		case examcategorydisabilitymapping.FieldCreatedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedById", values[i])
			} else if value.Valid {
				ecdm.CreatedById = int32(value.Int64)
			}
		case examcategorydisabilitymapping.FieldCreatedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByUserName", values[i])
			} else if value.Valid {
				ecdm.CreatedByUserName = value.String
			}
		case examcategorydisabilitymapping.FieldCreatedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByEmployeeId", values[i])
			} else if value.Valid {
				ecdm.CreatedByEmployeeId = value.String
			}
		case examcategorydisabilitymapping.FieldCreatedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByDesignation", values[i])
			} else if value.Valid {
				ecdm.CreatedByDesignation = value.String
			}
		case examcategorydisabilitymapping.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedDate", values[i])
			} else if value.Valid {
				ecdm.CreatedDate = value.Time
			}
		case examcategorydisabilitymapping.FieldVerifiedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedById", values[i])
			} else if value.Valid {
				ecdm.VerifiedById = value.Int64
			}
		case examcategorydisabilitymapping.FieldVerifiedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedByUserName", values[i])
			} else if value.Valid {
				ecdm.VerifiedByUserName = value.String
			}
		case examcategorydisabilitymapping.FieldVerifiedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedByEmployeeId", values[i])
			} else if value.Valid {
				ecdm.VerifiedByEmployeeId = value.String
			}
		case examcategorydisabilitymapping.FieldVerifiedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedByDesignation", values[i])
			} else if value.Valid {
				ecdm.VerifiedByDesignation = value.String
			}
		case examcategorydisabilitymapping.FieldVerifiedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedDate", values[i])
			} else if value.Valid {
				ecdm.VerifiedDate = value.Time
			}
		case examcategorydisabilitymapping.FieldDeletedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedById", values[i])
			} else if value.Valid {
				ecdm.DeletedById = value.Int64
			}
		case examcategorydisabilitymapping.FieldDeletedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedByUserName", values[i])
			} else if value.Valid {
				ecdm.DeletedByUserName = value.String
			}
		case examcategorydisabilitymapping.FieldDeletedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedByEmployeeId", values[i])
			} else if value.Valid {
				ecdm.DeletedByEmployeeId = value.String
			}
		case examcategorydisabilitymapping.FieldDeletedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedByDesignation", values[i])
			} else if value.Valid {
				ecdm.DeletedByDesignation = value.String
			}
		case examcategorydisabilitymapping.FieldDeletedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedDate", values[i])
			} else if value.Valid {
				ecdm.DeletedDate = value.Time
			}
		default:
			ecdm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExamCategoryDisabilityMapping.
// This includes values selected through modifiers, order, etc.
func (ecdm *ExamCategoryDisabilityMapping) Value(name string) (ent.Value, error) {
	return ecdm.selectValues.Get(name)
}

// Update returns a builder for updating this ExamCategoryDisabilityMapping.
// Note that you need to call ExamCategoryDisabilityMapping.Unwrap() before calling this method if this ExamCategoryDisabilityMapping
// was returned from a transaction, and the transaction was committed or rolled back.
func (ecdm *ExamCategoryDisabilityMapping) Update() *ExamCategoryDisabilityMappingUpdateOne {
	return NewExamCategoryDisabilityMappingClient(ecdm.config).UpdateOne(ecdm)
}

// Unwrap unwraps the ExamCategoryDisabilityMapping entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ecdm *ExamCategoryDisabilityMapping) Unwrap() *ExamCategoryDisabilityMapping {
	_tx, ok := ecdm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExamCategoryDisabilityMapping is not a transactional entity")
	}
	ecdm.config.driver = _tx.drv
	return ecdm
}

// String implements the fmt.Stringer.
func (ecdm *ExamCategoryDisabilityMapping) String() string {
	var builder strings.Builder
	builder.WriteString("ExamCategoryDisabilityMapping(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ecdm.ID))
	builder.WriteString("ExamCode=")
	builder.WriteString(fmt.Sprintf("%v", ecdm.ExamCode))
	builder.WriteString(", ")
	builder.WriteString("ExamShortName=")
	builder.WriteString(ecdm.ExamShortName)
	builder.WriteString(", ")
	builder.WriteString("CategoryDisability=")
	builder.WriteString(ecdm.CategoryDisability)
	builder.WriteString(", ")
	builder.WriteString("CategoryDisabilityCode=")
	builder.WriteString(ecdm.CategoryDisabilityCode)
	builder.WriteString(", ")
	builder.WriteString("CategoryDisabilityDescription=")
	builder.WriteString(ecdm.CategoryDisabilityDescription)
	builder.WriteString(", ")
	builder.WriteString("AgeException=")
	builder.WriteString(fmt.Sprintf("%v", ecdm.AgeException))
	builder.WriteString(", ")
	builder.WriteString("ServiceException=")
	builder.WriteString(fmt.Sprintf("%v", ecdm.ServiceException))
	builder.WriteString(", ")
	builder.WriteString("DrivingLicense=")
	builder.WriteString(fmt.Sprintf("%v", ecdm.DrivingLicense))
	builder.WriteString(", ")
	builder.WriteString("OrderNumber=")
	builder.WriteString(ecdm.OrderNumber)
	builder.WriteString(", ")
	builder.WriteString("Status=")
	builder.WriteString(ecdm.Status)
	builder.WriteString(", ")
	builder.WriteString("CreatedById=")
	builder.WriteString(fmt.Sprintf("%v", ecdm.CreatedById))
	builder.WriteString(", ")
	builder.WriteString("CreatedByUserName=")
	builder.WriteString(ecdm.CreatedByUserName)
	builder.WriteString(", ")
	builder.WriteString("CreatedByEmployeeId=")
	builder.WriteString(ecdm.CreatedByEmployeeId)
	builder.WriteString(", ")
	builder.WriteString("CreatedByDesignation=")
	builder.WriteString(ecdm.CreatedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("CreatedDate=")
	builder.WriteString(ecdm.CreatedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("VerifiedById=")
	builder.WriteString(fmt.Sprintf("%v", ecdm.VerifiedById))
	builder.WriteString(", ")
	builder.WriteString("VerifiedByUserName=")
	builder.WriteString(ecdm.VerifiedByUserName)
	builder.WriteString(", ")
	builder.WriteString("VerifiedByEmployeeId=")
	builder.WriteString(ecdm.VerifiedByEmployeeId)
	builder.WriteString(", ")
	builder.WriteString("VerifiedByDesignation=")
	builder.WriteString(ecdm.VerifiedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("VerifiedDate=")
	builder.WriteString(ecdm.VerifiedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("DeletedById=")
	builder.WriteString(fmt.Sprintf("%v", ecdm.DeletedById))
	builder.WriteString(", ")
	builder.WriteString("DeletedByUserName=")
	builder.WriteString(ecdm.DeletedByUserName)
	builder.WriteString(", ")
	builder.WriteString("DeletedByEmployeeId=")
	builder.WriteString(ecdm.DeletedByEmployeeId)
	builder.WriteString(", ")
	builder.WriteString("DeletedByDesignation=")
	builder.WriteString(ecdm.DeletedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("DeletedDate=")
	builder.WriteString(ecdm.DeletedDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ExamCategoryDisabilityMappings is a parsable slice of ExamCategoryDisabilityMapping.
type ExamCategoryDisabilityMappings []*ExamCategoryDisabilityMapping