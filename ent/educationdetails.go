// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/educationdetails"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// EducationDetails is the model entity for the EducationDetails schema.
type EducationDetails struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// EducationDescription holds the value of the "educationDescription" field.
	EducationDescription string `json:"educationDescription,omitempty"`
	// OrderNumber holds the value of the "OrderNumber" field.
	OrderNumber string `json:"OrderNumber,omitempty"`
	// CreatedById holds the value of the "CreatedById" field.
	CreatedById int64 `json:"CreatedById,omitempty"`
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
	// Status holds the value of the "Status" field.
	Status string `json:"Status,omitempty"`
	// DeletedById holds the value of the "DeletedById" field.
	DeletedById int64 `json:"DeletedById,omitempty"`
	// DeletedByUserName holds the value of the "DeletedByUserName" field.
	DeletedByUserName string `json:"DeletedByUserName,omitempty"`
	// DeletedByEmployeeId holds the value of the "DeletedByEmployeeId" field.
	DeletedByEmployeeId string `json:"DeletedByEmployeeId,omitempty"`
	// DeletedByDesignation holds the value of the "DeletedByDesignation" field.
	DeletedByDesignation string `json:"DeletedByDesignation,omitempty"`
	// DeletedDate holds the value of the "DeletedDate" field.
	DeletedDate time.Time `json:"DeletedDate,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EducationDetailsQuery when eager-loading is set.
	Edges        EducationDetailsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EducationDetailsEdges holds the relations/edges for other nodes in the graph.
type EducationDetailsEdges struct {
	// LogData holds the value of the LogData edge.
	LogData []*Logs `json:"LogData,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LogDataOrErr returns the LogData value or an error if the edge
// was not loaded in eager-loading.
func (e EducationDetailsEdges) LogDataOrErr() ([]*Logs, error) {
	if e.loadedTypes[0] {
		return e.LogData, nil
	}
	return nil, &NotLoadedError{edge: "LogData"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EducationDetails) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case educationdetails.FieldID, educationdetails.FieldCreatedById, educationdetails.FieldVerifiedById, educationdetails.FieldDeletedById:
			values[i] = new(sql.NullInt64)
		case educationdetails.FieldEducationDescription, educationdetails.FieldOrderNumber, educationdetails.FieldCreatedByUserName, educationdetails.FieldCreatedByEmployeeId, educationdetails.FieldCreatedByDesignation, educationdetails.FieldVerifiedByUserName, educationdetails.FieldVerifiedByEmployeeId, educationdetails.FieldVerifiedByDesignation, educationdetails.FieldStatus, educationdetails.FieldDeletedByUserName, educationdetails.FieldDeletedByEmployeeId, educationdetails.FieldDeletedByDesignation:
			values[i] = new(sql.NullString)
		case educationdetails.FieldCreatedDate, educationdetails.FieldVerifiedDate, educationdetails.FieldDeletedDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EducationDetails fields.
func (ed *EducationDetails) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case educationdetails.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ed.ID = int64(value.Int64)
		case educationdetails.FieldEducationDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field educationDescription", values[i])
			} else if value.Valid {
				ed.EducationDescription = value.String
			}
		case educationdetails.FieldOrderNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field OrderNumber", values[i])
			} else if value.Valid {
				ed.OrderNumber = value.String
			}
		case educationdetails.FieldCreatedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedById", values[i])
			} else if value.Valid {
				ed.CreatedById = value.Int64
			}
		case educationdetails.FieldCreatedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByUserName", values[i])
			} else if value.Valid {
				ed.CreatedByUserName = value.String
			}
		case educationdetails.FieldCreatedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByEmployeeId", values[i])
			} else if value.Valid {
				ed.CreatedByEmployeeId = value.String
			}
		case educationdetails.FieldCreatedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByDesignation", values[i])
			} else if value.Valid {
				ed.CreatedByDesignation = value.String
			}
		case educationdetails.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedDate", values[i])
			} else if value.Valid {
				ed.CreatedDate = value.Time
			}
		case educationdetails.FieldVerifiedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedById", values[i])
			} else if value.Valid {
				ed.VerifiedById = value.Int64
			}
		case educationdetails.FieldVerifiedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedByUserName", values[i])
			} else if value.Valid {
				ed.VerifiedByUserName = value.String
			}
		case educationdetails.FieldVerifiedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedByEmployeeId", values[i])
			} else if value.Valid {
				ed.VerifiedByEmployeeId = value.String
			}
		case educationdetails.FieldVerifiedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedByDesignation", values[i])
			} else if value.Valid {
				ed.VerifiedByDesignation = value.String
			}
		case educationdetails.FieldVerifiedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field VerifiedDate", values[i])
			} else if value.Valid {
				ed.VerifiedDate = value.Time
			}
		case educationdetails.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				ed.Status = value.String
			}
		case educationdetails.FieldDeletedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedById", values[i])
			} else if value.Valid {
				ed.DeletedById = value.Int64
			}
		case educationdetails.FieldDeletedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedByUserName", values[i])
			} else if value.Valid {
				ed.DeletedByUserName = value.String
			}
		case educationdetails.FieldDeletedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedByEmployeeId", values[i])
			} else if value.Valid {
				ed.DeletedByEmployeeId = value.String
			}
		case educationdetails.FieldDeletedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedByDesignation", values[i])
			} else if value.Valid {
				ed.DeletedByDesignation = value.String
			}
		case educationdetails.FieldDeletedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedDate", values[i])
			} else if value.Valid {
				ed.DeletedDate = value.Time
			}
		default:
			ed.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EducationDetails.
// This includes values selected through modifiers, order, etc.
func (ed *EducationDetails) Value(name string) (ent.Value, error) {
	return ed.selectValues.Get(name)
}

// QueryLogData queries the "LogData" edge of the EducationDetails entity.
func (ed *EducationDetails) QueryLogData() *LogsQuery {
	return NewEducationDetailsClient(ed.config).QueryLogData(ed)
}

// Update returns a builder for updating this EducationDetails.
// Note that you need to call EducationDetails.Unwrap() before calling this method if this EducationDetails
// was returned from a transaction, and the transaction was committed or rolled back.
func (ed *EducationDetails) Update() *EducationDetailsUpdateOne {
	return NewEducationDetailsClient(ed.config).UpdateOne(ed)
}

// Unwrap unwraps the EducationDetails entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ed *EducationDetails) Unwrap() *EducationDetails {
	_tx, ok := ed.config.driver.(*txDriver)
	if !ok {
		panic("ent: EducationDetails is not a transactional entity")
	}
	ed.config.driver = _tx.drv
	return ed
}

// String implements the fmt.Stringer.
func (ed *EducationDetails) String() string {
	var builder strings.Builder
	builder.WriteString("EducationDetails(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ed.ID))
	builder.WriteString("educationDescription=")
	builder.WriteString(ed.EducationDescription)
	builder.WriteString(", ")
	builder.WriteString("OrderNumber=")
	builder.WriteString(ed.OrderNumber)
	builder.WriteString(", ")
	builder.WriteString("CreatedById=")
	builder.WriteString(fmt.Sprintf("%v", ed.CreatedById))
	builder.WriteString(", ")
	builder.WriteString("CreatedByUserName=")
	builder.WriteString(ed.CreatedByUserName)
	builder.WriteString(", ")
	builder.WriteString("CreatedByEmployeeId=")
	builder.WriteString(ed.CreatedByEmployeeId)
	builder.WriteString(", ")
	builder.WriteString("CreatedByDesignation=")
	builder.WriteString(ed.CreatedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("CreatedDate=")
	builder.WriteString(ed.CreatedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("VerifiedById=")
	builder.WriteString(fmt.Sprintf("%v", ed.VerifiedById))
	builder.WriteString(", ")
	builder.WriteString("VerifiedByUserName=")
	builder.WriteString(ed.VerifiedByUserName)
	builder.WriteString(", ")
	builder.WriteString("VerifiedByEmployeeId=")
	builder.WriteString(ed.VerifiedByEmployeeId)
	builder.WriteString(", ")
	builder.WriteString("VerifiedByDesignation=")
	builder.WriteString(ed.VerifiedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("VerifiedDate=")
	builder.WriteString(ed.VerifiedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("Status=")
	builder.WriteString(ed.Status)
	builder.WriteString(", ")
	builder.WriteString("DeletedById=")
	builder.WriteString(fmt.Sprintf("%v", ed.DeletedById))
	builder.WriteString(", ")
	builder.WriteString("DeletedByUserName=")
	builder.WriteString(ed.DeletedByUserName)
	builder.WriteString(", ")
	builder.WriteString("DeletedByEmployeeId=")
	builder.WriteString(ed.DeletedByEmployeeId)
	builder.WriteString(", ")
	builder.WriteString("DeletedByDesignation=")
	builder.WriteString(ed.DeletedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("DeletedDate=")
	builder.WriteString(ed.DeletedDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// EducationDetailsSlice is a parsable slice of EducationDetails.
type EducationDetailsSlice []*EducationDetails
