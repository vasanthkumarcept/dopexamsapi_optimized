// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/logs"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Logs is the model entity for the Logs schema.
type Logs struct {
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
	Eventtime                              time.Time `json:"eventtime,omitempty"`
	admin_login_log_data                   *int32
	admin_master_log_data                  *int64
	cadre_eligible_configuration_log_data  *int64
	category_mininum_mark_mapping_log_data *int64
	center_log_data                        *int32
	education_details_log_data             *int64
	eligibility_cadre_pay_matrix_log_data  *int64
	employee_master_log_data               *int64
	exam_log_data                          *int32
	exam_notifications_log_data            *int32
	exam_application_mtspmmg_log_data      *int64
	exam_applications_gdspa_log_data       *int64
	exam_applications_gdspm_log_data       *int64
	exam_applications_ip_log_data          *int64
	exam_applications_pmpa_log_data        *int64
	exam_applications_ps_log_data          *int64
	login_log_data                         *int
	notification_log_data                  *int32
	user_master_log_data                   *int64
	selectValues                           sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Logs) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case logs.FieldLatitude, logs.FieldLongitude:
			values[i] = new(sql.NullFloat64)
		case logs.FieldID, logs.FieldUniqueid:
			values[i] = new(sql.NullInt64)
		case logs.FieldUserid, logs.FieldUsertype, logs.FieldUserdetails, logs.FieldRemarks, logs.FieldAction, logs.FieldIpaddress, logs.FieldDevicetype, logs.FieldOs, logs.FieldBrowser:
			values[i] = new(sql.NullString)
		case logs.FieldEventtime:
			values[i] = new(sql.NullTime)
		case logs.ForeignKeys[0]: // admin_login_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[1]: // admin_master_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[2]: // cadre_eligible_configuration_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[3]: // category_mininum_mark_mapping_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[4]: // center_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[5]: // education_details_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[6]: // eligibility_cadre_pay_matrix_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[7]: // employee_master_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[8]: // exam_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[9]: // exam_notifications_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[10]: // exam_application_mtspmmg_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[11]: // exam_applications_gdspa_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[12]: // exam_applications_gdspm_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[13]: // exam_applications_ip_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[14]: // exam_applications_pmpa_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[15]: // exam_applications_ps_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[16]: // login_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[17]: // notification_log_data
			values[i] = new(sql.NullInt64)
		case logs.ForeignKeys[18]: // user_master_log_data
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Logs fields.
func (l *Logs) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case logs.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int64(value.Int64)
		case logs.FieldUserid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userid", values[i])
			} else if value.Valid {
				l.Userid = value.String
			}
		case logs.FieldUniqueid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uniqueid", values[i])
			} else if value.Valid {
				l.Uniqueid = value.Int64
			}
		case logs.FieldUsertype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field usertype", values[i])
			} else if value.Valid {
				l.Usertype = value.String
			}
		case logs.FieldUserdetails:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userdetails", values[i])
			} else if value.Valid {
				l.Userdetails = value.String
			}
		case logs.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				l.Remarks = value.String
			}
		case logs.FieldAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action", values[i])
			} else if value.Valid {
				l.Action = value.String
			}
		case logs.FieldIpaddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ipaddress", values[i])
			} else if value.Valid {
				l.Ipaddress = value.String
			}
		case logs.FieldDevicetype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field devicetype", values[i])
			} else if value.Valid {
				l.Devicetype = value.String
			}
		case logs.FieldOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os", values[i])
			} else if value.Valid {
				l.Os = value.String
			}
		case logs.FieldBrowser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser", values[i])
			} else if value.Valid {
				l.Browser = value.String
			}
		case logs.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				l.Latitude = value.Float64
			}
		case logs.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				l.Longitude = value.Float64
			}
		case logs.FieldEventtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field eventtime", values[i])
			} else if value.Valid {
				l.Eventtime = value.Time
			}
		case logs.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field admin_login_log_data", value)
			} else if value.Valid {
				l.admin_login_log_data = new(int32)
				*l.admin_login_log_data = int32(value.Int64)
			}
		case logs.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field admin_master_log_data", value)
			} else if value.Valid {
				l.admin_master_log_data = new(int64)
				*l.admin_master_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field cadre_eligible_configuration_log_data", value)
			} else if value.Valid {
				l.cadre_eligible_configuration_log_data = new(int64)
				*l.cadre_eligible_configuration_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category_mininum_mark_mapping_log_data", value)
			} else if value.Valid {
				l.category_mininum_mark_mapping_log_data = new(int64)
				*l.category_mininum_mark_mapping_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field center_log_data", value)
			} else if value.Valid {
				l.center_log_data = new(int32)
				*l.center_log_data = int32(value.Int64)
			}
		case logs.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field education_details_log_data", value)
			} else if value.Valid {
				l.education_details_log_data = new(int64)
				*l.education_details_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[6]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field eligibility_cadre_pay_matrix_log_data", value)
			} else if value.Valid {
				l.eligibility_cadre_pay_matrix_log_data = new(int64)
				*l.eligibility_cadre_pay_matrix_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[7]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field employee_master_log_data", value)
			} else if value.Valid {
				l.employee_master_log_data = new(int64)
				*l.employee_master_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[8]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_log_data", value)
			} else if value.Valid {
				l.exam_log_data = new(int32)
				*l.exam_log_data = int32(value.Int64)
			}
		case logs.ForeignKeys[9]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_notifications_log_data", value)
			} else if value.Valid {
				l.exam_notifications_log_data = new(int32)
				*l.exam_notifications_log_data = int32(value.Int64)
			}
		case logs.ForeignKeys[10]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_application_mtspmmg_log_data", value)
			} else if value.Valid {
				l.exam_application_mtspmmg_log_data = new(int64)
				*l.exam_application_mtspmmg_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[11]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_applications_gdspa_log_data", value)
			} else if value.Valid {
				l.exam_applications_gdspa_log_data = new(int64)
				*l.exam_applications_gdspa_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[12]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_applications_gdspm_log_data", value)
			} else if value.Valid {
				l.exam_applications_gdspm_log_data = new(int64)
				*l.exam_applications_gdspm_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[13]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_applications_ip_log_data", value)
			} else if value.Valid {
				l.exam_applications_ip_log_data = new(int64)
				*l.exam_applications_ip_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[14]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_applications_pmpa_log_data", value)
			} else if value.Valid {
				l.exam_applications_pmpa_log_data = new(int64)
				*l.exam_applications_pmpa_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[15]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_applications_ps_log_data", value)
			} else if value.Valid {
				l.exam_applications_ps_log_data = new(int64)
				*l.exam_applications_ps_log_data = int64(value.Int64)
			}
		case logs.ForeignKeys[16]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field login_log_data", value)
			} else if value.Valid {
				l.login_log_data = new(int)
				*l.login_log_data = int(value.Int64)
			}
		case logs.ForeignKeys[17]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field notification_log_data", value)
			} else if value.Valid {
				l.notification_log_data = new(int32)
				*l.notification_log_data = int32(value.Int64)
			}
		case logs.ForeignKeys[18]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_master_log_data", value)
			} else if value.Valid {
				l.user_master_log_data = new(int64)
				*l.user_master_log_data = int64(value.Int64)
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Logs.
// This includes values selected through modifiers, order, etc.
func (l *Logs) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// Update returns a builder for updating this Logs.
// Note that you need to call Logs.Unwrap() before calling this method if this Logs
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Logs) Update() *LogsUpdateOne {
	return NewLogsClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Logs entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Logs) Unwrap() *Logs {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Logs is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Logs) String() string {
	var builder strings.Builder
	builder.WriteString("Logs(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("userid=")
	builder.WriteString(l.Userid)
	builder.WriteString(", ")
	builder.WriteString("uniqueid=")
	builder.WriteString(fmt.Sprintf("%v", l.Uniqueid))
	builder.WriteString(", ")
	builder.WriteString("usertype=")
	builder.WriteString(l.Usertype)
	builder.WriteString(", ")
	builder.WriteString("userdetails=")
	builder.WriteString(l.Userdetails)
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(l.Remarks)
	builder.WriteString(", ")
	builder.WriteString("action=")
	builder.WriteString(l.Action)
	builder.WriteString(", ")
	builder.WriteString("ipaddress=")
	builder.WriteString(l.Ipaddress)
	builder.WriteString(", ")
	builder.WriteString("devicetype=")
	builder.WriteString(l.Devicetype)
	builder.WriteString(", ")
	builder.WriteString("os=")
	builder.WriteString(l.Os)
	builder.WriteString(", ")
	builder.WriteString("browser=")
	builder.WriteString(l.Browser)
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(fmt.Sprintf("%v", l.Latitude))
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(fmt.Sprintf("%v", l.Longitude))
	builder.WriteString(", ")
	builder.WriteString("eventtime=")
	builder.WriteString(l.Eventtime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// LogsSlice is a parsable slice of Logs.
type LogsSlice []*Logs