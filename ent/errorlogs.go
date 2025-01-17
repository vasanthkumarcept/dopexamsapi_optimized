// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/errorlogs"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ErrorLogs is the model entity for the ErrorLogs schema.
type ErrorLogs struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Userid holds the value of the "userid" field.
	Userid string `json:"userid,omitempty"`
	// Uniqueid holds the value of the "uniqueid" field.
	Uniqueid int64 `json:"uniqueid,omitempty"`
	// Usertype holds the value of the "usertype" field.
	Usertype string `json:"usertype,omitempty"`
	// Userdetails holds the value of the "userdetails" field.
	Userdetails string `json:"userdetails,omitempty"`
	// Remarks holds the value of the "remarks" field.
	Remarks string `json:"remarks,omitempty"`
	// Action holds the value of the "action" field.
	Action string `json:"action,omitempty"`
	// Ipaddress holds the value of the "ipaddress" field.
	Ipaddress string `json:"ipaddress,omitempty"`
	// Devicetype holds the value of the "devicetype" field.
	Devicetype string `json:"devicetype,omitempty"`
	// Os holds the value of the "os" field.
	Os string `json:"os,omitempty"`
	// Browser holds the value of the "browser" field.
	Browser string `json:"browser,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude float64 `json:"longitude,omitempty"`
	// Eventtime holds the value of the "eventtime" field.
	Eventtime time.Time `json:"eventtime,omitempty"`
	// UpdatedBy holds the value of the "UpdatedBy" field.
	UpdatedBy string `json:"UpdatedBy,omitempty"`
	// UpdatedTime holds the value of the "UpdatedTime" field.
	UpdatedTime time.Time `json:"UpdatedTime,omitempty"`
	// AssignedTo holds the value of the "AssignedTo" field.
	AssignedTo string `json:"AssignedTo,omitempty"`
	// RemarksNew holds the value of the "RemarksNew" field.
	RemarksNew string `json:"RemarksNew,omitempty"`
	// Status holds the value of the "Status" field.
	Status string `json:"Status,omitempty"`
	// ClosedOn holds the value of the "ClosedOn" field.
	ClosedOn     time.Time `json:"ClosedOn,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ErrorLogs) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case errorlogs.FieldLatitude, errorlogs.FieldLongitude:
			values[i] = new(sql.NullFloat64)
		case errorlogs.FieldID, errorlogs.FieldUniqueid:
			values[i] = new(sql.NullInt64)
		case errorlogs.FieldUserid, errorlogs.FieldUsertype, errorlogs.FieldUserdetails, errorlogs.FieldRemarks, errorlogs.FieldAction, errorlogs.FieldIpaddress, errorlogs.FieldDevicetype, errorlogs.FieldOs, errorlogs.FieldBrowser, errorlogs.FieldUpdatedBy, errorlogs.FieldAssignedTo, errorlogs.FieldRemarksNew, errorlogs.FieldStatus:
			values[i] = new(sql.NullString)
		case errorlogs.FieldEventtime, errorlogs.FieldUpdatedTime, errorlogs.FieldClosedOn:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ErrorLogs fields.
func (el *ErrorLogs) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case errorlogs.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			el.ID = int64(value.Int64)
		case errorlogs.FieldUserid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userid", values[i])
			} else if value.Valid {
				el.Userid = value.String
			}
		case errorlogs.FieldUniqueid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uniqueid", values[i])
			} else if value.Valid {
				el.Uniqueid = value.Int64
			}
		case errorlogs.FieldUsertype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field usertype", values[i])
			} else if value.Valid {
				el.Usertype = value.String
			}
		case errorlogs.FieldUserdetails:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userdetails", values[i])
			} else if value.Valid {
				el.Userdetails = value.String
			}
		case errorlogs.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				el.Remarks = value.String
			}
		case errorlogs.FieldAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action", values[i])
			} else if value.Valid {
				el.Action = value.String
			}
		case errorlogs.FieldIpaddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ipaddress", values[i])
			} else if value.Valid {
				el.Ipaddress = value.String
			}
		case errorlogs.FieldDevicetype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field devicetype", values[i])
			} else if value.Valid {
				el.Devicetype = value.String
			}
		case errorlogs.FieldOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os", values[i])
			} else if value.Valid {
				el.Os = value.String
			}
		case errorlogs.FieldBrowser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser", values[i])
			} else if value.Valid {
				el.Browser = value.String
			}
		case errorlogs.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				el.Latitude = value.Float64
			}
		case errorlogs.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				el.Longitude = value.Float64
			}
		case errorlogs.FieldEventtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field eventtime", values[i])
			} else if value.Valid {
				el.Eventtime = value.Time
			}
		case errorlogs.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedBy", values[i])
			} else if value.Valid {
				el.UpdatedBy = value.String
			}
		case errorlogs.FieldUpdatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedTime", values[i])
			} else if value.Valid {
				el.UpdatedTime = value.Time
			}
		case errorlogs.FieldAssignedTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field AssignedTo", values[i])
			} else if value.Valid {
				el.AssignedTo = value.String
			}
		case errorlogs.FieldRemarksNew:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field RemarksNew", values[i])
			} else if value.Valid {
				el.RemarksNew = value.String
			}
		case errorlogs.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				el.Status = value.String
			}
		case errorlogs.FieldClosedOn:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ClosedOn", values[i])
			} else if value.Valid {
				el.ClosedOn = value.Time
			}
		default:
			el.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ErrorLogs.
// This includes values selected through modifiers, order, etc.
func (el *ErrorLogs) Value(name string) (ent.Value, error) {
	return el.selectValues.Get(name)
}

// Update returns a builder for updating this ErrorLogs.
// Note that you need to call ErrorLogs.Unwrap() before calling this method if this ErrorLogs
// was returned from a transaction, and the transaction was committed or rolled back.
func (el *ErrorLogs) Update() *ErrorLogsUpdateOne {
	return NewErrorLogsClient(el.config).UpdateOne(el)
}

// Unwrap unwraps the ErrorLogs entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (el *ErrorLogs) Unwrap() *ErrorLogs {
	_tx, ok := el.config.driver.(*txDriver)
	if !ok {
		panic("ent: ErrorLogs is not a transactional entity")
	}
	el.config.driver = _tx.drv
	return el
}

// String implements the fmt.Stringer.
func (el *ErrorLogs) String() string {
	var builder strings.Builder
	builder.WriteString("ErrorLogs(")
	builder.WriteString(fmt.Sprintf("id=%v, ", el.ID))
	builder.WriteString("userid=")
	builder.WriteString(el.Userid)
	builder.WriteString(", ")
	builder.WriteString("uniqueid=")
	builder.WriteString(fmt.Sprintf("%v", el.Uniqueid))
	builder.WriteString(", ")
	builder.WriteString("usertype=")
	builder.WriteString(el.Usertype)
	builder.WriteString(", ")
	builder.WriteString("userdetails=")
	builder.WriteString(el.Userdetails)
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(el.Remarks)
	builder.WriteString(", ")
	builder.WriteString("action=")
	builder.WriteString(el.Action)
	builder.WriteString(", ")
	builder.WriteString("ipaddress=")
	builder.WriteString(el.Ipaddress)
	builder.WriteString(", ")
	builder.WriteString("devicetype=")
	builder.WriteString(el.Devicetype)
	builder.WriteString(", ")
	builder.WriteString("os=")
	builder.WriteString(el.Os)
	builder.WriteString(", ")
	builder.WriteString("browser=")
	builder.WriteString(el.Browser)
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(fmt.Sprintf("%v", el.Latitude))
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(fmt.Sprintf("%v", el.Longitude))
	builder.WriteString(", ")
	builder.WriteString("eventtime=")
	builder.WriteString(el.Eventtime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedBy=")
	builder.WriteString(el.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("UpdatedTime=")
	builder.WriteString(el.UpdatedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("AssignedTo=")
	builder.WriteString(el.AssignedTo)
	builder.WriteString(", ")
	builder.WriteString("RemarksNew=")
	builder.WriteString(el.RemarksNew)
	builder.WriteString(", ")
	builder.WriteString("Status=")
	builder.WriteString(el.Status)
	builder.WriteString(", ")
	builder.WriteString("ClosedOn=")
	builder.WriteString(el.ClosedOn.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ErrorLogsSlice is a parsable slice of ErrorLogs.
type ErrorLogsSlice []*ErrorLogs
