// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/division_choice_pa"
	"recruit/ent/exam_applications_gdspa"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Division_Choice_PA is the model entity for the Division_Choice_PA schema.
type Division_Choice_PA struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// ApplicationID holds the value of the "ApplicationID" field.
	ApplicationID int64 `json:"ApplicationID,omitempty"`
	// PlacePrefNo holds the value of the "PlacePrefNo" field.
	PlacePrefNo int64 `json:"PlacePrefNo,omitempty"`
	// PlacePrefValue holds the value of the "PlacePrefValue" field.
	PlacePrefValue string `json:"PlacePrefValue,omitempty"`
	// EmployeeID holds the value of the "EmployeeID" field.
	EmployeeID int64 `json:"EmployeeID,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// UpdatedBy holds the value of the "UpdatedBy" field.
	UpdatedBy string `json:"UpdatedBy,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Division_Choice_PAQuery when eager-loading is set.
	Edges        Division_Choice_PAEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Division_Choice_PAEdges holds the relations/edges for other nodes in the graph.
type Division_Choice_PAEdges struct {
	// ApplnGDSPARef holds the value of the ApplnGDSPA_Ref edge.
	ApplnGDSPARef *Exam_Applications_GDSPA `json:"ApplnGDSPA_Ref,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ApplnGDSPARefOrErr returns the ApplnGDSPARef value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Division_Choice_PAEdges) ApplnGDSPARefOrErr() (*Exam_Applications_GDSPA, error) {
	if e.loadedTypes[0] {
		if e.ApplnGDSPARef == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: exam_applications_gdspa.Label}
		}
		return e.ApplnGDSPARef, nil
	}
	return nil, &NotLoadedError{edge: "ApplnGDSPA_Ref"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Division_Choice_PA) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case division_choice_pa.FieldID, division_choice_pa.FieldApplicationID, division_choice_pa.FieldPlacePrefNo, division_choice_pa.FieldEmployeeID:
			values[i] = new(sql.NullInt64)
		case division_choice_pa.FieldPlacePrefValue, division_choice_pa.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case division_choice_pa.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Division_Choice_PA fields.
func (dcp *Division_Choice_PA) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case division_choice_pa.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dcp.ID = int32(value.Int64)
		case division_choice_pa.FieldApplicationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ApplicationID", values[i])
			} else if value.Valid {
				dcp.ApplicationID = value.Int64
			}
		case division_choice_pa.FieldPlacePrefNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field PlacePrefNo", values[i])
			} else if value.Valid {
				dcp.PlacePrefNo = value.Int64
			}
		case division_choice_pa.FieldPlacePrefValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PlacePrefValue", values[i])
			} else if value.Valid {
				dcp.PlacePrefValue = value.String
			}
		case division_choice_pa.FieldEmployeeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeeID", values[i])
			} else if value.Valid {
				dcp.EmployeeID = value.Int64
			}
		case division_choice_pa.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				dcp.UpdatedAt = value.Time
			}
		case division_choice_pa.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedBy", values[i])
			} else if value.Valid {
				dcp.UpdatedBy = value.String
			}
		default:
			dcp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Division_Choice_PA.
// This includes values selected through modifiers, order, etc.
func (dcp *Division_Choice_PA) Value(name string) (ent.Value, error) {
	return dcp.selectValues.Get(name)
}

// QueryApplnGDSPARef queries the "ApplnGDSPA_Ref" edge of the Division_Choice_PA entity.
func (dcp *Division_Choice_PA) QueryApplnGDSPARef() *ExamApplicationsGDSPAQuery {
	return NewDivisionChoicePAClient(dcp.config).QueryApplnGDSPARef(dcp)
}

// Update returns a builder for updating this Division_Choice_PA.
// Note that you need to call Division_Choice_PA.Unwrap() before calling this method if this Division_Choice_PA
// was returned from a transaction, and the transaction was committed or rolled back.
func (dcp *Division_Choice_PA) Update() *DivisionChoicePAUpdateOne {
	return NewDivisionChoicePAClient(dcp.config).UpdateOne(dcp)
}

// Unwrap unwraps the Division_Choice_PA entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dcp *Division_Choice_PA) Unwrap() *Division_Choice_PA {
	_tx, ok := dcp.config.driver.(*txDriver)
	if !ok {
		panic("ent: Division_Choice_PA is not a transactional entity")
	}
	dcp.config.driver = _tx.drv
	return dcp
}

// String implements the fmt.Stringer.
func (dcp *Division_Choice_PA) String() string {
	var builder strings.Builder
	builder.WriteString("Division_Choice_PA(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dcp.ID))
	builder.WriteString("ApplicationID=")
	builder.WriteString(fmt.Sprintf("%v", dcp.ApplicationID))
	builder.WriteString(", ")
	builder.WriteString("PlacePrefNo=")
	builder.WriteString(fmt.Sprintf("%v", dcp.PlacePrefNo))
	builder.WriteString(", ")
	builder.WriteString("PlacePrefValue=")
	builder.WriteString(dcp.PlacePrefValue)
	builder.WriteString(", ")
	builder.WriteString("EmployeeID=")
	builder.WriteString(fmt.Sprintf("%v", dcp.EmployeeID))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(dcp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedBy=")
	builder.WriteString(dcp.UpdatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// Division_Choice_PAs is a parsable slice of Division_Choice_PA.
type Division_Choice_PAs []*Division_Choice_PA