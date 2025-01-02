package start

import (
	"fmt"
	"recruit/ent"
	"recruit/util"

	//"recruit/util"
	"regexp"
	//"strconv"
	"strings"

	ca_reg "recruit/payloadstructure/candidate_registration"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

var ErrordbMap map[string]string
var Remarks string

func init() {
	ErrordbMap = map[string]string{
		"03000": "03—SQL Statement Not Yet Complete—422",
		"08000": "08—Connection Exception—503",
		"08003": "08—Connection Exception—503",
		"08006": "08—Connection Exception—503",
		"08001": "08—Connection Exception—503",
		"08004": "08—Connection Exception—503",
		"08007": "08—Connection Exception—503",
		"08P01": "08—Connection Exception—400",
		"09000": "09—Triggered Action Exception—422",
		"0A000": "0A—Feature Not Supported—501",
		"0B000": "0B—Invalid Transaction Initiation—422",
		"0F000": "0F—Locator Exception—422",
		"0F001": "0F—Locator Exception—422",
		"0L000": "0L—Invalid Grantor—422",
		"0LP01": "0L—Invalid Grantor—422",
		"0P000": "0P—Invalid Role Specification—422",
		"0Z000": "0Z—Diagnostics Exception—422",
		"0Z002": "0Z—Diagnostics Exception—422",
		"20000": "20—Case Not Found—422",
		"21000": "21—Cardinality Violation—422",
		"22000": "22—Data Exception—422",
		"2202E": "22—Data Exception—422",
		"22021": "22—Data Exception—422",
		"22008": "22—Data Exception—422",
		"22012": "22—Data Exception—422",
		"22005": "22—Data Exception—422",
		"2200B": "22—Data Exception—422",
		"22022": "22—Data Exception—422",
		"22015": "22—Data Exception—422",
		"2201E": "22—Data Exception—422",
		"22014": "22—Data Exception—422",
		"22016": "22—Data Exception—422",
		"2201F": "22—Data Exception—422",
		"2201G": "22—Data Exception—422",
		"22018": "22—Data Exception—422",
		"22007": "22—Data Exception—422",
		"22019": "22—Data Exception—422",
		"2200D": "22—Data Exception—422",
		"22025": "22—Data Exception—422",
		"22P06": "22—Data Exception—422",
		"22010": "22—Data Exception—422",
		"22023": "22—Data Exception—422",
		"22013": "22—Data Exception—422",
		"2201B": "22—Data Exception—422",
		"2201W": "22—Data Exception—422",
		"2201X": "22—Data Exception—422",
		"2202H": "22—Data Exception—422",
		"2202G": "22—Data Exception—422",
		"22009": "22—Data Exception—422",
		"2200C": "22—Data Exception—422",
		"2200G": "22—Data Exception—422",
		"22004": "22—Data Exception—422",
		"22002": "22—Data Exception—422",
		"22003": "22—Data Exception—422",
		"2200H": "22—Data Exception—422",
		"22026": "22—Data Exception—422",
		"22001": "22—Data Exception—422",
		"22011": "22—Data Exception—422",
		"22027": "22—Data Exception—422",
		"22024": "22—Data Exception—422",
		"2200F": "22—Data Exception—422",
		"22P01": "22—Data Exception—422",
		"22P02": "22—Data Exception—422",
		"22P03": "22—Data Exception—422",
		"22P04": "22—Data Exception—422",
		"22P05": "22—Data Exception—422",
		"2200L": "22—Data Exception—422",
		"2200M": "22—Data Exception—422",
		"2200N": "22—Data Exception—422",
		"2200S": "22—Data Exception—422",
		"2200T": "22—Data Exception—422",
		"22030": "22—Data Exception—422",
		"22031": "22—Data Exception—422",
		"22032": "22—Data Exception—422",
		"22033": "22—Data Exception—422",
		"22034": "22—Data Exception—422",
		"22035": "22—Data Exception—422",
		"22036": "22—Data Exception—422",
		"22037": "22—Data Exception—422",
		"22038": "22—Data Exception—422",
		"22039": "22—Data Exception—422",
		"2203A": "22—Data Exception—422",
		"2203B": "22—Data Exception—422",
		"2203C": "22—Data Exception—422",
		"2203D": "22—Data Exception—422",
		"2203E": "22—Data Exception—422",
		"2203F": "22—Data Exception—422",
		"2203G": "22—Data Exception—500",
		"23000": "23—Integrity Constraint Violation—409",
		"23001": "23—Integrity Constraint Violation—422",
		"23502": "23—Integrity Constraint Violation—422",
		"23503": "23—Integrity Constraint Violation—422",
		"23505": "23—Integrity Constraint Violation—409",
		"23514": "23—Integrity Constraint Violation—422",
		"23P01": "23—Integrity Constraint Violation—422",
		"24000": "24—Invalid Cursor State—422",
		"25000": "25—Invalid Transaction State—409",
		"25001": "25—Invalid Transaction State—409",
		"25002": "25—Invalid Transaction State—409",
		"25008": "25—Invalid Transaction State—409",
		"25003": "25—Invalid Transaction State—409",
		"25004": "25—Invalid Transaction State—409",
		"25005": "25—Invalid Transaction State—409",
		"25006": "25—Invalid Transaction State—409",
		"25007": "25—Invalid Transaction State—422",
		"25P01": "25—Invalid Transaction State—409",
		"25P02": "25—Invalid Transaction State—409",
		"25P03": "25—Invalid Transaction State—409",
		"26000": "26—Invalid SQL Statement Name—422",
		"27000": "27—Triggered Data Change Violation—422",
		"28000": "28—Invalid Authorization Specification—401",
		"28P01": "28—Invalid Authorization Specification—401",
		"2B000": "2B—Dependent Privilege Descriptors Still Exist—422",
		"2BP01": "2B—Dependent Privilege Descriptors Still Exist—422",
		"2D000": "2D—Invalid Transaction Termination—422",
		"2F000": "2F—SQL Routine Exception—500",
		"2F005": "2F—SQL Routine Exception—500",
		"2F002": "2F—SQL Routine Exception—500",
		"2F003": "2F—SQL Routine Exception—500",
		"2F004": "2F—SQL Routine Exception—500",
		"34000": "34—Invalid Cursor Name—500",
		"38000": "38—External Routine Exception—500",
		"38001": "38—External Routine Exception—500",
		"38002": "38—External Routine Exception—500",
		"38003": "38—External Routine Exception—500",
		"38004": "38—External Routine Exception—500",
		"39000": "39—External Routine Invocation Exception—422",
		"39001": "39—External Routine Invocation Exception—422",
		"39004": "39—External Routine Invocation Exception—422",
		"39P01": "39—External Routine Invocation Exception—422",
		"39P02": "39—External Routine Invocation Exception—422",
		"39P03": "39—External Routine Invocation Exception—422",
		"3B000": "3B—Savepoint Exception—422",
		"3B001": "3B—Savepoint Exception—422",
		"3D000": "3D—Invalid Catalog Name—422",
		"3F000": "3F—Invalid Schema Name—422",
		"40000": "40—Transaction Rollback—409",
		"40002": "40—Transaction Rollback—409",
		"40001": "40—Transaction Rollback—409",
		"40003": "40—Transaction Rollback—409",
		"40P01": "40—Transaction Rollback—409",
		"42000": "42—Syntax Error or Access Rule Violation—400",
		"42601": "42—Syntax Error or Access Rule Violation—400",
		"42501": "42—Syntax Error or Access Rule Violation—403",
		"42846": "42—Syntax Error or Access Rule Violation—422",
		"42803": "42—Syntax Error or Access Rule Violation—422",
		"42P20": "42—Syntax Error or Access Rule Violation—422",
		"42P19": "42—Syntax Error or Access Rule Violation—422",
		"42830": "42—Syntax Error or Access Rule Violation—422",
		"42602": "42—Syntax Error or Access Rule Violation—422",
		"42622": "42—Syntax Error or Access Rule Violation—422",
		"42939": "42—Syntax Error or Access Rule Violation—422",
		"42804": "42—Syntax Error or Access Rule Violation—422",
		"42P18": "42—Syntax Error or Access Rule Violation—422",
		"42P21": "42—Syntax Error or Access Rule Violation—422",
		"42P22": "42—Syntax Error or Access Rule Violation—422",
		"42809": "42—Syntax Error or Access Rule Violation—422",
		"428C9": "42—Syntax Error or Access Rule Violation—422",
		"42703": "42—Syntax Error or Access Rule Violation—422",
		"42883": "42—Syntax Error or Access Rule Violation—422",
		"42P01": "42—Syntax Error or Access Rule Violation—422",
		"42P02": "42—Syntax Error or Access Rule Violation—422",
		"42704": "42—Syntax Error or Access Rule Violation—422",
		"42701": "42—Syntax Error or Access Rule Violation—422",
		"42P03": "42—Syntax Error or Access Rule Violation—422",
		"42P04": "42—Syntax Error or Access Rule Violation—422",
		"42723": "42—Syntax Error or Access Rule Violation—422",
		"42P05": "42—Syntax Error or Access Rule Violation—422",
		"42P06": "42—Syntax Error or Access Rule Violation—422",
		"42P07": "42—Syntax Error or Access Rule Violation—422",
		"42712": "42—Syntax Error or Access Rule Violation—422",
		"42710": "42—Syntax Error or Access Rule Violation—422",
		"42702": "42—Syntax Error or Access Rule Violation—422",
		"42725": "42—Syntax Error or Access Rule Violation—422",
		"42P08": "42—Syntax Error or Access Rule Violation—422",
		"42P09": "42—Syntax Error or Access Rule Violation—422",
		"42P10": "42—Syntax Error or Access Rule Violation—422",
		"42611": "42—Syntax Error or Access Rule Violation—422",
		"42P11": "42—Syntax Error or Access Rule Violation—422",
		"42P12": "42—Syntax Error or Access Rule Violation—422",
		"42P13": "42—Syntax Error or Access Rule Violation—422",
		"42P14": "42—Syntax Error or Access Rule Violation—422",
		"42P15": "42—Syntax Error or Access Rule Violation—422",
		"42P16": "42—Syntax Error or Access Rule Violation—422",
		"42P17": "42—Syntax Error or Access Rule Violation—422",
		"44000": "44—WITH CHECK OPTION Violation—409",
		"53000": "53—Insufficient Resources—503",
		"53100": "53—Insufficient Resources—503",
		"53200": "53—Insufficient Resources—503",
		"53300": "53—Insufficient Resources—503",
		"53400": "53—Insufficient Resources—503",
		"54000": "54—Program Limit Exceeded—503",
		"54001": "54—Program Limit Exceeded—400",
		"54011": "54—Program Limit Exceeded—400",
		"54023": "54—Program Limit Exceeded—400",
		"55000": "55—Object Not In Prerequisite State—422",
		"55006": "55—Object Not In Prerequisite State—422",
		"55P02": "55—Object Not In Prerequisite State—409",
		"55P03": "55—Object Not In Prerequisite State—409",
		"55P04": "55—Object Not In Prerequisite State—400",
		"57000": "57—Operator Intervention—503",
		"57014": "57—Operator Intervention—499",
		"57P01": "57—Operator Intervention—503",
		"57P02": "57—Operator Intervention—503",
		"57P03": "57—Operator Intervention—503",
		"57P04": "57—Operator Intervention—503",
		"57P05": "57—Operator Intervention—503",
		"58000": "58—System Error (errors external to PostgreSQL itself)—500",
		"58030": "58—System Error (errors external to PostgreSQL itself)—500",
		"58P01": "58—System Error (errors external to PostgreSQL itself)—404",
		"58P02": "58—System Error (errors external to PostgreSQL itself)—409",
		"72000": "72—Snapshot Failure—409",
		"F0000": "F0—Configuration File Error—500",
		"F0001": "F0—Configuration File Error—409",
		"HV000": "HV—Foreign Data Wrapper Error (SQL/MED)—500",
		"HV005": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV002": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV010": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV021": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV024": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV007": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV008": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV004": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV006": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV091": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV00B": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV00C": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV00D": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV090": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV00A": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV009": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV014": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV001": "HV—Foreign Data Wrapper Error (SQL/MED)—500",
		"HV00P": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV00J": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV00K": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV00Q": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV00R": "HV—Foreign Data Wrapper Error (SQL/MED)—422",
		"HV00L": "HV—Foreign Data Wrapper Error (SQL/MED)—500",
		"HV00M": "HV—Foreign Data Wrapper Error (SQL/MED)—500",
		"HV00N": "HV—Foreign Data Wrapper Error (SQL/MED)—500",
		"P0000": "P0—PL/pgSQL Error—500",
		"P0001": "P0—PL/pgSQL Error—422",
		"P0002": "P0—PL/pgSQL Error—404",
		"P0003": "P0—PL/pgSQL Error—422",
		"P0004": "P0—PL/pgSQL Error—500",
	}
}

// handleDatabaseError handles database errors and sends appropriate responses
// func HandleDatabaseError(ctx *gin.Context, err error, status int32, stgError string, logdata ca_reg.LogData, client *ent.Client) {
// 	statusCode := 500
// 	sqlStateRegex := regexp.MustCompile(`SQLSTATE (\d+)`)
// 	e := err.Error()
// 	matches := sqlStateRegex.FindStringSubmatch(e)
// 	var sqlState string
// 	if len(matches) >= 2 {
// 		sqlState = matches[1]
// 	} else {
// 		logdata.Remarks = "Unknown database error" + stgError
// 		estring := err.Error()
// 		errRsps := newErrordbResponse([]string{estring},  err)
// 		ctx.JSON(statusCode, errRsps)
// 		return
// 	}

// 	errordbClass1, ok := ErrordbMap[sqlState]
// 	if !ok {
// 		errRsps := newErrordbResponse([]string{"Unknown database error"},  err) //[]string{"POTH04"},
// 		ctx.JSON(statusCode, errRsps)
// 		return
// 	}
// 	dberror := strings.Split(errordbClass1, "—")
// 	errRsp := newErrordbResponse([]string{dberror[1]},  err)
// 	code, _ := strconv.Atoi(dberror[2])
// 	ctx.JSON(code, errRsp)
// }

func HandleDatabaseError(ctx *gin.Context, err error, status int32, stgError string, logdata ca_reg.LogData, client *ent.Client, errRemarks string) {
	sqlStateRegex := regexp.MustCompile(`SQLSTATE (\d+)`)
	//e := err
	matches := sqlStateRegex.FindStringSubmatch(err.Error())
	var sqlState string
	if len(matches) >= 2 {
		sqlState = matches[1]
	} else {
		logdata.Remarks = errRemarks + " Unknown database error " + stgError + " -DB01 "
		util.LogErrorNew(client, logdata, err)
		HandleError(ctx, 500, " Unknown database error "+stgError+" -DB01 ")
		return
	}

	errordbClass1, ok := ErrordbMap[sqlState]
	if !ok {
		logdata.Remarks = errRemarks + " Unknown database error " + stgError + " -DB02 "
		util.LogErrorNew(client, logdata, err)
		HandleError(ctx, 500, " Unknown database error "+stgError+" -DB02 ")
		return
	}
	dberror := strings.Split(errordbClass1, "—")
	//message := dberror[1]
	//errRsp := newErrordbResponse(e)
	//code, _ := strconv.Atoi(dberror[2])
	logdata.Remarks = errRemarks + " Database error " + stgError + " -DB03 " + dberror[2]
	util.LogErrorNew(client, logdata, err)
	HandleError(ctx, 500, " Unknown database error "+stgError+" -DB03 ")

}

func HandleDBErrorInitial(ctx *gin.Context, client *ent.Client, errRemarks string) string {
	err := util.CheckDatabaseConnection(client)
	if err != nil {
		HandleDbError(ctx, 500, errRemarks)
		return "error"
	}
	return "noerror"
}

func HandleDatabaseErrorWithoutLog(ctx *gin.Context, err error, status int32, stgError string, client *ent.Client, errRemarks string) {
	sqlStateRegex := regexp.MustCompile(`SQLSTATE (\d+)`)
	//e := err
	//var Remarks string = ""
	matches := sqlStateRegex.FindStringSubmatch(err.Error())
	var sqlState string
	if len(matches) >= 2 {
		sqlState = matches[1]
	} else {
		Remarks = errRemarks + " Unknown database error " + err.Error() + " " + stgError + " -DB01 "
		util.SystemLogError(client, "500", Remarks)
		HandleError(ctx, 500, " Unknown database error "+stgError+" -DB01 ")
		return
	}

	errordbClass1, ok := ErrordbMap[sqlState]
	if !ok {
		Remarks = errRemarks + " Unknown database error " + err.Error() + " " + stgError + " -DB02 "
		util.SystemLogError(client, "500", Remarks)
		HandleError(ctx, 500, " Unknown database error "+stgError+" -DB02 ")
		return
	}
	dberror := strings.Split(errordbClass1, "—")
	//message := dberror[1]
	//errRsp := newErrordbResponse(e)
	//code, _ := strconv.Atoi(dberror[2])
	Remarks = errRemarks + " Database error " + err.Error() + " " + stgError + " -DB03 " + dberror[2]
	util.SystemLogError(client, "500", Remarks)
	HandleError(ctx, 500, " Unknown database error "+stgError+" -DB03 ")
}

func PrintDBError(err error) {
	// Type assert to *pq.Error to access PostgreSQL-specific error details

	if pqErr, ok := err.(*pq.Error); ok {
		fmt.Printf("Database error code: %s\n", pqErr.Code)
		fmt.Printf("Database error message: %s\n", pqErr.Message)
		fmt.Printf("Database error detail: %s\n", pqErr.Detail)
		fmt.Printf("Database error hint: %s\n", pqErr.Hint)
	} else {
		fmt.Printf("General error: %s\n", err.Error())
	}
}
