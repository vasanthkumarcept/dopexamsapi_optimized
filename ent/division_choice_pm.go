// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/division_choice_pm"
	"recruit/ent/exam_applications_gdspm"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Division_Choice_PM is the model entity for the Division_Choice_PM schema.
type Division_Choice_PM struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// ApplicationID holds the value of the "ApplicationID" field.
	ApplicationID int64 `json:"ApplicationID,omitempty"`
	// Group holds the value of the "Group" field.
	Group string `json:"Group,omitempty"`
	// CadrePrefNo holds the value of the "CadrePrefNo" field.
	CadrePrefNo int64 `json:"CadrePrefNo,omitempty"`
	// Cadre holds the value of the "Cadre" field.
	Cadre string `json:"Cadre,omitempty"`
	// PostPrefNo holds the value of the "PostPrefNo" field.
	PostPrefNo int64 `json:"PostPrefNo,omitempty"`
	// PostingPrefValue holds the value of the "PostingPrefValue" field.
	PostingPrefValue string `json:"PostingPrefValue,omitempty"`
	// EmployeeID holds the value of the "EmployeeID" field.
	EmployeeID int64 `json:"EmployeeID,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// UpdatedBy holds the value of the "UpdatedBy" field.
	UpdatedBy string `json:"UpdatedBy,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Division_Choice_PMQuery when eager-loading is set.
	Edges        Division_Choice_PMEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Division_Choice_PMEdges holds the relations/edges for other nodes in the graph.
type Division_Choice_PMEdges struct {
	// ApplnGDSPMRef holds the value of the ApplnGDSPM_Ref edge.
	ApplnGDSPMRef *Exam_Applications_GDSPM `json:"ApplnGDSPM_Ref,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ApplnGDSPMRefOrErr returns the ApplnGDSPMRef value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Division_Choice_PMEdges) ApplnGDSPMRefOrErr() (*Exam_Applications_GDSPM, error) {
	if e.loadedTypes[0] {
		if e.ApplnGDSPMRef == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: exam_applications_gdspm.Label}
		}
		return e.ApplnGDSPMRef, nil
	}
	return nil, &NotLoadedError{edge: "ApplnGDSPM_Ref"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Division_Choice_PM) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case division_choice_pm.FieldID, division_choice_pm.FieldApplicationID, division_choice_pm.FieldCadrePrefNo, division_choice_pm.FieldPostPrefNo, division_choice_pm.FieldEmployeeID:
			values[i] = new(sql.NullInt64)
		case division_choice_pm.FieldGroup, division_choice_pm.FieldCadre, division_choice_pm.FieldPostingPrefValue, division_choice_pm.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case division_choice_pm.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Division_Choice_PM fields.
func (dcp *Division_Choice_PM) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case division_choice_pm.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dcp.ID = int32(value.Int64)
		case division_choice_pm.FieldApplicationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ApplicationID", values[i])
			} else if value.Valid {
				dcp.ApplicationID = value.Int64
			}
		case division_choice_pm.FieldGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Group", values[i])
			} else if value.Valid {
				dcp.Group = value.String
			}
		case division_choice_pm.FieldCadrePrefNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CadrePrefNo", values[i])
			} else if value.Valid {
				dcp.CadrePrefNo = value.Int64
			}
		case division_choice_pm.FieldCadre:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Cadre", values[i])
			} else if value.Valid {
				dcp.Cadre = value.String
			}
		case division_choice_pm.FieldPostPrefNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field PostPrefNo", values[i])
			} else if value.Valid {
				dcp.PostPrefNo = value.Int64
			}
		case division_choice_pm.FieldPostingPrefValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PostingPrefValue", values[i])
			} else if value.Valid {
				dcp.PostingPrefValue = value.String
			}
		case division_choice_pm.FieldEmployeeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeeID", values[i])
			} else if value.Valid {
				dcp.EmployeeID = value.Int64
			}
		case division_choice_pm.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				dcp.UpdatedAt = value.Time
			}
		case division_choice_pm.FieldUpdatedBy:
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

// Value returns the ent.Value that was dynamically selected and assigned to the Division_Choice_PM.
// This includes values selected through modifiers, order, etc.
func (dcp *Division_Choice_PM) Value(name string) (ent.Value, error) {
	return dcp.selectValues.Get(name)
}

// QueryApplnGDSPMRef queries the "ApplnGDSPM_Ref" edge of the Division_Choice_PM entity.
func (dcp *Division_Choice_PM) QueryApplnGDSPMRef() *ExamApplicationsGDSPMQuery {
	return NewDivisionChoicePMClient(dcp.config).QueryApplnGDSPMRef(dcp)
}

// Update returns a builder for updating this Division_Choice_PM.
// Note that you need to call Division_Choice_PM.Unwrap() before calling this method if this Division_Choice_PM
// was returned from a transaction, and the transaction was committed or rolled back.
func (dcp *Division_Choice_PM) Update() *DivisionChoicePMUpdateOne {
	return NewDivisionChoicePMClient(dcp.config).UpdateOne(dcp)
}

// Unwrap unwraps the Division_Choice_PM entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dcp *Division_Choice_PM) Unwrap() *Division_Choice_PM {
	_tx, ok := dcp.config.driver.(*txDriver)
	if !ok {
		panic("ent: Division_Choice_PM is not a transactional entity")
	}
	dcp.config.driver = _tx.drv
	return dcp
}

// String implements the fmt.Stringer.
func (dcp *Division_Choice_PM) String() string {
	var builder strings.Builder
	builder.WriteString("Division_Choice_PM(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dcp.ID))
	builder.WriteString("ApplicationID=")
	builder.WriteString(fmt.Sprintf("%v", dcp.ApplicationID))
	builder.WriteString(", ")
	builder.WriteString("Group=")
	builder.WriteString(dcp.Group)
	builder.WriteString(", ")
	builder.WriteString("CadrePrefNo=")
	builder.WriteString(fmt.Sprintf("%v", dcp.CadrePrefNo))
	builder.WriteString(", ")
	builder.WriteString("Cadre=")
	builder.WriteString(dcp.Cadre)
	builder.WriteString(", ")
	builder.WriteString("PostPrefNo=")
	builder.WriteString(fmt.Sprintf("%v", dcp.PostPrefNo))
	builder.WriteString(", ")
	builder.WriteString("PostingPrefValue=")
	builder.WriteString(dcp.PostingPrefValue)
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

// Division_Choice_PMs is a parsable slice of Division_Choice_PM.
type Division_Choice_PMs []*Division_Choice_PM
