// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/exampapers"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ExamPapers is the model entity for the ExamPapers schema.
type ExamPapers struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// PaperDescription holds the value of the "PaperDescription" field.
	PaperDescription string `json:"PaperDescription,omitempty"`
	// ExamCode holds the value of the "ExamCode" field.
	ExamCode int32 `json:"ExamCode,omitempty"`
	// ExamName holds the value of the "ExamName" field.
	ExamName string `json:"ExamName,omitempty"`
	// ExamShortName holds the value of the "ExamShortName" field.
	ExamShortName string `json:"ExamShortName,omitempty"`
	// PaperTypeCode holds the value of the "PaperTypeCode" field.
	PaperTypeCode int32 `json:"PaperTypeCode,omitempty"`
	// PaperTypeName holds the value of the "PaperTypeName" field.
	PaperTypeName string `json:"PaperTypeName,omitempty"`
	// CompetitiveQualifying holds the value of the "CompetitiveQualifying" field.
	CompetitiveQualifying bool `json:"CompetitiveQualifying,omitempty"`
	// ExceptionForDisability holds the value of the "ExceptionForDisability" field.
	ExceptionForDisability bool `json:"ExceptionForDisability,omitempty"`
	// MaximumMarks holds the value of the "MaximumMarks" field.
	MaximumMarks int `json:"MaximumMarks,omitempty"`
	// Duration holds the value of the "Duration" field.
	Duration int `json:"Duration,omitempty"`
	// LocalLanguageAllowedQuestionPaper holds the value of the "localLanguageAllowedQuestionPaper" field.
	LocalLanguageAllowedQuestionPaper string `json:"localLanguageAllowedQuestionPaper,omitempty"`
	// LocalLanguageAllowedAnswerPaper holds the value of the "localLanguageAllowedAnswerPaper" field.
	LocalLanguageAllowedAnswerPaper string `json:"localLanguageAllowedAnswerPaper,omitempty"`
	// DisabilityTypeID holds the value of the "DisabilityTypeID" field.
	DisabilityTypeID int32 `json:"DisabilityTypeID,omitempty"`
	// FromTime holds the value of the "fromTime" field.
	FromTime time.Time `json:"fromTime,omitempty"`
	// ToTime holds the value of the "toTime" field.
	ToTime time.Time `json:"toTime,omitempty"`
	// OrderNumber holds the value of the "OrderNumber" field.
	OrderNumber string `json:"OrderNumber,omitempty"`
	// CreatedById holds the value of the "CreatedById" field.
	CreatedById int64 `json:"CreatedById,omitempty"`
	// CreatedByUserName holds the value of the "CreatedByUserName" field.
	CreatedByUserName string `json:"CreatedByUserName,omitempty"`
	// CreatedByEmpId holds the value of the "CreatedByEmpId" field.
	CreatedByEmpId int64 `json:"CreatedByEmpId,omitempty"`
	// CreatedByDesignation holds the value of the "CreatedByDesignation" field.
	CreatedByDesignation string `json:"CreatedByDesignation,omitempty"`
	// CreatedDate holds the value of the "CreatedDate" field.
	CreatedDate time.Time `json:"CreatedDate,omitempty"`
	// Verifiedbyid holds the value of the "verifiedbyid" field.
	Verifiedbyid int64 `json:"verifiedbyid,omitempty"`
	// Verifiedbyusername holds the value of the "verifiedbyusername" field.
	Verifiedbyusername string `json:"verifiedbyusername,omitempty"`
	// VerifiedbyEmployeeid holds the value of the "verifiedbyEmployeeid" field.
	VerifiedbyEmployeeid int64 `json:"verifiedbyEmployeeid,omitempty"`
	// VerifiedbyDesignation holds the value of the "verifiedbyDesignation" field.
	VerifiedbyDesignation string `json:"verifiedbyDesignation,omitempty"`
	// VerifiedDate holds the value of the "verifiedDate" field.
	VerifiedDate time.Time `json:"verifiedDate,omitempty"`
	// Statuss holds the value of the "Statuss" field.
	Statuss string `json:"Statuss,omitempty"`
	// Deletedbyid holds the value of the "deletedbyid" field.
	Deletedbyid int64 `json:"deletedbyid,omitempty"`
	// Deletedbyusername holds the value of the "deletedbyusername" field.
	Deletedbyusername string `json:"deletedbyusername,omitempty"`
	// DeletedbyEmployeeid holds the value of the "deletedbyEmployeeid" field.
	DeletedbyEmployeeid int64 `json:"deletedbyEmployeeid,omitempty"`
	// DeletedbyDesignation holds the value of the "deletedbyDesignation" field.
	DeletedbyDesignation string `json:"deletedbyDesignation,omitempty"`
	// DeletedDate holds the value of the "deletedDate" field.
	DeletedDate time.Time `json:"deletedDate,omitempty"`
	// PaperStatus holds the value of the "PaperStatus" field.
	PaperStatus string `json:"PaperStatus,omitempty"`
	// CalendarCode holds the value of the "CalendarCode" field.
	CalendarCode int32 `json:"CalendarCode,omitempty"`
	// ExamCodePS holds the value of the "ExamCodePS" field.
	ExamCodePS int32 `json:"ExamCodePS,omitempty"`
	// CreatedByEmployeeId holds the value of the "CreatedByEmployeeId" field.
	CreatedByEmployeeId string `json:"CreatedByEmployeeId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExamPapersQuery when eager-loading is set.
	Edges              ExamPapersEdges `json:"edges"`
	disability_dis_ref *int32
	exam_papers        *int32
	selectValues       sql.SelectValues
}

// ExamPapersEdges holds the relations/edges for other nodes in the graph.
type ExamPapersEdges struct {
	// Centers holds the value of the centers edge.
	Centers []*Center `json:"centers,omitempty"`
	// ExampapersTypes holds the value of the exampapers_types edge.
	ExampapersTypes []*PaperTypes `json:"exampapers_types,omitempty"`
	// PapersRef holds the value of the papers_ref edge.
	PapersRef []*ExamCalendar `json:"papers_ref,omitempty"`
	// ExamPaperEligibility holds the value of the ExamPaperEligibility edge.
	ExamPaperEligibility []*EligibilityMaster `json:"ExamPaperEligibility,omitempty"`
	// DisRef holds the value of the dis_ref edge.
	DisRef []*Disability `json:"dis_ref,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// CentersOrErr returns the Centers value or an error if the edge
// was not loaded in eager-loading.
func (e ExamPapersEdges) CentersOrErr() ([]*Center, error) {
	if e.loadedTypes[0] {
		return e.Centers, nil
	}
	return nil, &NotLoadedError{edge: "centers"}
}

// ExampapersTypesOrErr returns the ExampapersTypes value or an error if the edge
// was not loaded in eager-loading.
func (e ExamPapersEdges) ExampapersTypesOrErr() ([]*PaperTypes, error) {
	if e.loadedTypes[1] {
		return e.ExampapersTypes, nil
	}
	return nil, &NotLoadedError{edge: "exampapers_types"}
}

// PapersRefOrErr returns the PapersRef value or an error if the edge
// was not loaded in eager-loading.
func (e ExamPapersEdges) PapersRefOrErr() ([]*ExamCalendar, error) {
	if e.loadedTypes[2] {
		return e.PapersRef, nil
	}
	return nil, &NotLoadedError{edge: "papers_ref"}
}

// ExamPaperEligibilityOrErr returns the ExamPaperEligibility value or an error if the edge
// was not loaded in eager-loading.
func (e ExamPapersEdges) ExamPaperEligibilityOrErr() ([]*EligibilityMaster, error) {
	if e.loadedTypes[3] {
		return e.ExamPaperEligibility, nil
	}
	return nil, &NotLoadedError{edge: "ExamPaperEligibility"}
}

// DisRefOrErr returns the DisRef value or an error if the edge
// was not loaded in eager-loading.
func (e ExamPapersEdges) DisRefOrErr() ([]*Disability, error) {
	if e.loadedTypes[4] {
		return e.DisRef, nil
	}
	return nil, &NotLoadedError{edge: "dis_ref"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExamPapers) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exampapers.FieldCompetitiveQualifying, exampapers.FieldExceptionForDisability:
			values[i] = new(sql.NullBool)
		case exampapers.FieldID, exampapers.FieldExamCode, exampapers.FieldPaperTypeCode, exampapers.FieldMaximumMarks, exampapers.FieldDuration, exampapers.FieldDisabilityTypeID, exampapers.FieldCreatedById, exampapers.FieldCreatedByEmpId, exampapers.FieldVerifiedbyid, exampapers.FieldVerifiedbyEmployeeid, exampapers.FieldDeletedbyid, exampapers.FieldDeletedbyEmployeeid, exampapers.FieldCalendarCode, exampapers.FieldExamCodePS:
			values[i] = new(sql.NullInt64)
		case exampapers.FieldPaperDescription, exampapers.FieldExamName, exampapers.FieldExamShortName, exampapers.FieldPaperTypeName, exampapers.FieldLocalLanguageAllowedQuestionPaper, exampapers.FieldLocalLanguageAllowedAnswerPaper, exampapers.FieldOrderNumber, exampapers.FieldCreatedByUserName, exampapers.FieldCreatedByDesignation, exampapers.FieldVerifiedbyusername, exampapers.FieldVerifiedbyDesignation, exampapers.FieldStatuss, exampapers.FieldDeletedbyusername, exampapers.FieldDeletedbyDesignation, exampapers.FieldPaperStatus, exampapers.FieldCreatedByEmployeeId:
			values[i] = new(sql.NullString)
		case exampapers.FieldFromTime, exampapers.FieldToTime, exampapers.FieldCreatedDate, exampapers.FieldVerifiedDate, exampapers.FieldDeletedDate:
			values[i] = new(sql.NullTime)
		case exampapers.ForeignKeys[0]: // disability_dis_ref
			values[i] = new(sql.NullInt64)
		case exampapers.ForeignKeys[1]: // exam_papers
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExamPapers fields.
func (ep *ExamPapers) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exampapers.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ep.ID = int32(value.Int64)
		case exampapers.FieldPaperDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PaperDescription", values[i])
			} else if value.Valid {
				ep.PaperDescription = value.String
			}
		case exampapers.FieldExamCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ExamCode", values[i])
			} else if value.Valid {
				ep.ExamCode = int32(value.Int64)
			}
		case exampapers.FieldExamName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamName", values[i])
			} else if value.Valid {
				ep.ExamName = value.String
			}
		case exampapers.FieldExamShortName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamShortName", values[i])
			} else if value.Valid {
				ep.ExamShortName = value.String
			}
		case exampapers.FieldPaperTypeCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field PaperTypeCode", values[i])
			} else if value.Valid {
				ep.PaperTypeCode = int32(value.Int64)
			}
		case exampapers.FieldPaperTypeName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PaperTypeName", values[i])
			} else if value.Valid {
				ep.PaperTypeName = value.String
			}
		case exampapers.FieldCompetitiveQualifying:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field CompetitiveQualifying", values[i])
			} else if value.Valid {
				ep.CompetitiveQualifying = value.Bool
			}
		case exampapers.FieldExceptionForDisability:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ExceptionForDisability", values[i])
			} else if value.Valid {
				ep.ExceptionForDisability = value.Bool
			}
		case exampapers.FieldMaximumMarks:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field MaximumMarks", values[i])
			} else if value.Valid {
				ep.MaximumMarks = int(value.Int64)
			}
		case exampapers.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Duration", values[i])
			} else if value.Valid {
				ep.Duration = int(value.Int64)
			}
		case exampapers.FieldLocalLanguageAllowedQuestionPaper:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field localLanguageAllowedQuestionPaper", values[i])
			} else if value.Valid {
				ep.LocalLanguageAllowedQuestionPaper = value.String
			}
		case exampapers.FieldLocalLanguageAllowedAnswerPaper:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field localLanguageAllowedAnswerPaper", values[i])
			} else if value.Valid {
				ep.LocalLanguageAllowedAnswerPaper = value.String
			}
		case exampapers.FieldDisabilityTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field DisabilityTypeID", values[i])
			} else if value.Valid {
				ep.DisabilityTypeID = int32(value.Int64)
			}
		case exampapers.FieldFromTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field fromTime", values[i])
			} else if value.Valid {
				ep.FromTime = value.Time
			}
		case exampapers.FieldToTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field toTime", values[i])
			} else if value.Valid {
				ep.ToTime = value.Time
			}
		case exampapers.FieldOrderNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field OrderNumber", values[i])
			} else if value.Valid {
				ep.OrderNumber = value.String
			}
		case exampapers.FieldCreatedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedById", values[i])
			} else if value.Valid {
				ep.CreatedById = value.Int64
			}
		case exampapers.FieldCreatedByUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByUserName", values[i])
			} else if value.Valid {
				ep.CreatedByUserName = value.String
			}
		case exampapers.FieldCreatedByEmpId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByEmpId", values[i])
			} else if value.Valid {
				ep.CreatedByEmpId = value.Int64
			}
		case exampapers.FieldCreatedByDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByDesignation", values[i])
			} else if value.Valid {
				ep.CreatedByDesignation = value.String
			}
		case exampapers.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedDate", values[i])
			} else if value.Valid {
				ep.CreatedDate = value.Time
			}
		case exampapers.FieldVerifiedbyid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field verifiedbyid", values[i])
			} else if value.Valid {
				ep.Verifiedbyid = value.Int64
			}
		case exampapers.FieldVerifiedbyusername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field verifiedbyusername", values[i])
			} else if value.Valid {
				ep.Verifiedbyusername = value.String
			}
		case exampapers.FieldVerifiedbyEmployeeid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field verifiedbyEmployeeid", values[i])
			} else if value.Valid {
				ep.VerifiedbyEmployeeid = value.Int64
			}
		case exampapers.FieldVerifiedbyDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field verifiedbyDesignation", values[i])
			} else if value.Valid {
				ep.VerifiedbyDesignation = value.String
			}
		case exampapers.FieldVerifiedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field verifiedDate", values[i])
			} else if value.Valid {
				ep.VerifiedDate = value.Time
			}
		case exampapers.FieldStatuss:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Statuss", values[i])
			} else if value.Valid {
				ep.Statuss = value.String
			}
		case exampapers.FieldDeletedbyid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deletedbyid", values[i])
			} else if value.Valid {
				ep.Deletedbyid = value.Int64
			}
		case exampapers.FieldDeletedbyusername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deletedbyusername", values[i])
			} else if value.Valid {
				ep.Deletedbyusername = value.String
			}
		case exampapers.FieldDeletedbyEmployeeid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deletedbyEmployeeid", values[i])
			} else if value.Valid {
				ep.DeletedbyEmployeeid = value.Int64
			}
		case exampapers.FieldDeletedbyDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deletedbyDesignation", values[i])
			} else if value.Valid {
				ep.DeletedbyDesignation = value.String
			}
		case exampapers.FieldDeletedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedDate", values[i])
			} else if value.Valid {
				ep.DeletedDate = value.Time
			}
		case exampapers.FieldPaperStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PaperStatus", values[i])
			} else if value.Valid {
				ep.PaperStatus = value.String
			}
		case exampapers.FieldCalendarCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CalendarCode", values[i])
			} else if value.Valid {
				ep.CalendarCode = int32(value.Int64)
			}
		case exampapers.FieldExamCodePS:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ExamCodePS", values[i])
			} else if value.Valid {
				ep.ExamCodePS = int32(value.Int64)
			}
		case exampapers.FieldCreatedByEmployeeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedByEmployeeId", values[i])
			} else if value.Valid {
				ep.CreatedByEmployeeId = value.String
			}
		case exampapers.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field disability_dis_ref", value)
			} else if value.Valid {
				ep.disability_dis_ref = new(int32)
				*ep.disability_dis_ref = int32(value.Int64)
			}
		case exampapers.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_papers", value)
			} else if value.Valid {
				ep.exam_papers = new(int32)
				*ep.exam_papers = int32(value.Int64)
			}
		default:
			ep.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExamPapers.
// This includes values selected through modifiers, order, etc.
func (ep *ExamPapers) Value(name string) (ent.Value, error) {
	return ep.selectValues.Get(name)
}

// QueryCenters queries the "centers" edge of the ExamPapers entity.
func (ep *ExamPapers) QueryCenters() *CenterQuery {
	return NewExamPapersClient(ep.config).QueryCenters(ep)
}

// QueryExampapersTypes queries the "exampapers_types" edge of the ExamPapers entity.
func (ep *ExamPapers) QueryExampapersTypes() *PaperTypesQuery {
	return NewExamPapersClient(ep.config).QueryExampapersTypes(ep)
}

// QueryPapersRef queries the "papers_ref" edge of the ExamPapers entity.
func (ep *ExamPapers) QueryPapersRef() *ExamCalendarQuery {
	return NewExamPapersClient(ep.config).QueryPapersRef(ep)
}

// QueryExamPaperEligibility queries the "ExamPaperEligibility" edge of the ExamPapers entity.
func (ep *ExamPapers) QueryExamPaperEligibility() *EligibilityMasterQuery {
	return NewExamPapersClient(ep.config).QueryExamPaperEligibility(ep)
}

// QueryDisRef queries the "dis_ref" edge of the ExamPapers entity.
func (ep *ExamPapers) QueryDisRef() *DisabilityQuery {
	return NewExamPapersClient(ep.config).QueryDisRef(ep)
}

// Update returns a builder for updating this ExamPapers.
// Note that you need to call ExamPapers.Unwrap() before calling this method if this ExamPapers
// was returned from a transaction, and the transaction was committed or rolled back.
func (ep *ExamPapers) Update() *ExamPapersUpdateOne {
	return NewExamPapersClient(ep.config).UpdateOne(ep)
}

// Unwrap unwraps the ExamPapers entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ep *ExamPapers) Unwrap() *ExamPapers {
	_tx, ok := ep.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExamPapers is not a transactional entity")
	}
	ep.config.driver = _tx.drv
	return ep
}

// String implements the fmt.Stringer.
func (ep *ExamPapers) String() string {
	var builder strings.Builder
	builder.WriteString("ExamPapers(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ep.ID))
	builder.WriteString("PaperDescription=")
	builder.WriteString(ep.PaperDescription)
	builder.WriteString(", ")
	builder.WriteString("ExamCode=")
	builder.WriteString(fmt.Sprintf("%v", ep.ExamCode))
	builder.WriteString(", ")
	builder.WriteString("ExamName=")
	builder.WriteString(ep.ExamName)
	builder.WriteString(", ")
	builder.WriteString("ExamShortName=")
	builder.WriteString(ep.ExamShortName)
	builder.WriteString(", ")
	builder.WriteString("PaperTypeCode=")
	builder.WriteString(fmt.Sprintf("%v", ep.PaperTypeCode))
	builder.WriteString(", ")
	builder.WriteString("PaperTypeName=")
	builder.WriteString(ep.PaperTypeName)
	builder.WriteString(", ")
	builder.WriteString("CompetitiveQualifying=")
	builder.WriteString(fmt.Sprintf("%v", ep.CompetitiveQualifying))
	builder.WriteString(", ")
	builder.WriteString("ExceptionForDisability=")
	builder.WriteString(fmt.Sprintf("%v", ep.ExceptionForDisability))
	builder.WriteString(", ")
	builder.WriteString("MaximumMarks=")
	builder.WriteString(fmt.Sprintf("%v", ep.MaximumMarks))
	builder.WriteString(", ")
	builder.WriteString("Duration=")
	builder.WriteString(fmt.Sprintf("%v", ep.Duration))
	builder.WriteString(", ")
	builder.WriteString("localLanguageAllowedQuestionPaper=")
	builder.WriteString(ep.LocalLanguageAllowedQuestionPaper)
	builder.WriteString(", ")
	builder.WriteString("localLanguageAllowedAnswerPaper=")
	builder.WriteString(ep.LocalLanguageAllowedAnswerPaper)
	builder.WriteString(", ")
	builder.WriteString("DisabilityTypeID=")
	builder.WriteString(fmt.Sprintf("%v", ep.DisabilityTypeID))
	builder.WriteString(", ")
	builder.WriteString("fromTime=")
	builder.WriteString(ep.FromTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("toTime=")
	builder.WriteString(ep.ToTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("OrderNumber=")
	builder.WriteString(ep.OrderNumber)
	builder.WriteString(", ")
	builder.WriteString("CreatedById=")
	builder.WriteString(fmt.Sprintf("%v", ep.CreatedById))
	builder.WriteString(", ")
	builder.WriteString("CreatedByUserName=")
	builder.WriteString(ep.CreatedByUserName)
	builder.WriteString(", ")
	builder.WriteString("CreatedByEmpId=")
	builder.WriteString(fmt.Sprintf("%v", ep.CreatedByEmpId))
	builder.WriteString(", ")
	builder.WriteString("CreatedByDesignation=")
	builder.WriteString(ep.CreatedByDesignation)
	builder.WriteString(", ")
	builder.WriteString("CreatedDate=")
	builder.WriteString(ep.CreatedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("verifiedbyid=")
	builder.WriteString(fmt.Sprintf("%v", ep.Verifiedbyid))
	builder.WriteString(", ")
	builder.WriteString("verifiedbyusername=")
	builder.WriteString(ep.Verifiedbyusername)
	builder.WriteString(", ")
	builder.WriteString("verifiedbyEmployeeid=")
	builder.WriteString(fmt.Sprintf("%v", ep.VerifiedbyEmployeeid))
	builder.WriteString(", ")
	builder.WriteString("verifiedbyDesignation=")
	builder.WriteString(ep.VerifiedbyDesignation)
	builder.WriteString(", ")
	builder.WriteString("verifiedDate=")
	builder.WriteString(ep.VerifiedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("Statuss=")
	builder.WriteString(ep.Statuss)
	builder.WriteString(", ")
	builder.WriteString("deletedbyid=")
	builder.WriteString(fmt.Sprintf("%v", ep.Deletedbyid))
	builder.WriteString(", ")
	builder.WriteString("deletedbyusername=")
	builder.WriteString(ep.Deletedbyusername)
	builder.WriteString(", ")
	builder.WriteString("deletedbyEmployeeid=")
	builder.WriteString(fmt.Sprintf("%v", ep.DeletedbyEmployeeid))
	builder.WriteString(", ")
	builder.WriteString("deletedbyDesignation=")
	builder.WriteString(ep.DeletedbyDesignation)
	builder.WriteString(", ")
	builder.WriteString("deletedDate=")
	builder.WriteString(ep.DeletedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("PaperStatus=")
	builder.WriteString(ep.PaperStatus)
	builder.WriteString(", ")
	builder.WriteString("CalendarCode=")
	builder.WriteString(fmt.Sprintf("%v", ep.CalendarCode))
	builder.WriteString(", ")
	builder.WriteString("ExamCodePS=")
	builder.WriteString(fmt.Sprintf("%v", ep.ExamCodePS))
	builder.WriteString(", ")
	builder.WriteString("CreatedByEmployeeId=")
	builder.WriteString(ep.CreatedByEmployeeId)
	builder.WriteByte(')')
	return builder.String()
}

// ExamPapersSlice is a parsable slice of ExamPapers.
type ExamPapersSlice []*ExamPapers