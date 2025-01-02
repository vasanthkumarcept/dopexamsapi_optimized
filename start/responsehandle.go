package start

import (
	"os"
	"recruit/ent"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func getEnvVar() string {
	return os.Getenv("USER_ERROR_REMARKS")
}

var Action string = ""
var errRemark string = ""

// err, status, stgError, logdata, client, Remarks
func StartErrorHandlerWithLog(ctx *gin.Context, err error, status int32, stgError string, logdata ca_reg.LogData, client *ent.Client, Remarks string) {
	UserErrorRemarks := getEnvVar()
	if status == 400 {
		logdata.Remarks = "400 error occured in " + Remarks + stgError + " " + err.Error()
		util.LogErrorNew(client, logdata, err)
		HandleError(ctx, 400, UserErrorRemarks+stgError)
		return
	} else if status == 422 {
		logdata.Remarks = "422 error occured in " + Remarks + stgError + " " + err.Error()
		util.LogErrorNew(client, logdata, err)
		HandleError(ctx, 422, err.Error()+stgError)
		return
	} else {
		errRemarks := "500 error occured in error occured in " + Remarks + " " + err.Error()
		HandleDatabaseErrorWithoutLog(ctx, err, status, stgError, client, errRemarks)
		return
	}
}

func StartErrorHandlerWithoutLog(ctx *gin.Context, err error, status int32, stgError string, client *ent.Client, Remarks string) {
	UserErrorRemarks := getEnvVar()
	if status == 400 {
		Action = "400"
		errRemark = "400 error occured in " + Remarks + stgError + " " + err.Error()
		util.SystemLogError(client, Action, errRemark)
		HandleError(ctx, 400, UserErrorRemarks+stgError)
		return
	} else if status == 422 {
		Action = "422"
		errRemark = "422 error occured in " + Remarks + stgError + " " + err.Error()
		util.SystemLogError(client, Action, errRemark)
		HandleError(ctx, 422, err.Error()+stgError)
		return
	} else {
		errRemarks := "500 error occured in" + Remarks + " " + err.Error()
		HandleDatabaseErrorWithoutLog(ctx, err, status, stgError, client, errRemarks)
		return
	}
}

func ValidateEmail(email string) string {
	var emailError string = ""
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		emailError = " invalid email address"
		return emailError
	}
	// Extract the part before the "@" symbol
	localPart := email[:atIndex]

	if len(localPart) < 6 {
		emailError = " update with valid email ID / email ID Username should be minimum six characters "
		return emailError
	}
	return emailError
}

func MainHandleDBError(ctx *gin.Context, client *ent.Client, Remarks string, handleError string, username string) {
	UserErrorRemarks := getEnvVar()
	Action = "500"
	errRemark = Remarks + handleError
	util.SystemLogErrorNew(client, Action, errRemark, username)
	handleString := UserErrorRemarks + handleError
	HandleError(ctx, 500, handleString)
	//return
}

// gctx, client,Remarks, " -HA01", gctx.GetHeader("UserName")
func MainHandleError(ctx *gin.Context, client *ent.Client, Remarks string, handleError string, username string) {
	UserErrorRemarks := getEnvVar()
	Action = "400"
	errRemark = Remarks + handleError
	util.SystemLogErrorNew(client, Action, errRemark, username)
	handleString := UserErrorRemarks + handleError
	HandleError(ctx, 400, handleString)
	//return
}

// handleError(ctx *gin.Context, err error)
func HandleError(ctx *gin.Context, status int, handleString string) {
	emptyObject := struct{}{}
	{
		ctx.JSON(status, gin.H{
			"success":    true,
			"message":    handleString,
			"data":       emptyObject,
			"dataexists": false,
		})
	}
}

func HandleDbError(ctx *gin.Context, status int, handleString string) {
	emptyObject := struct{}{}
	{
		ctx.JSON(status, gin.H{
			"success":    false,
			"message":    handleString,
			"data":       emptyObject,
			"dataexists": false,
		})
	}
}

type RollmasterResponse struct {
	RoleUserCode int       `json:"RoleUserCode"`
	RoleName     string    `json:"role_name"`
	CreatedDate  time.Time `json:"created_date"`
	Status       bool      `json:"status"`
}

// ErrorResponse defines the error response format
type ErrorResponse struct {
	Error string `json:"error"`
}

// func newErrordbResponse(message []string, err error) errordbResponse {
// 	return errordbResponse{

// 		Success: false,
// 		Message: err.Error(),
// 	}

// }

// type errordbResponse struct {
// 	Success bool   `json:"success" example:"false"`
// 	Message string `json:"error" example:"Error message"`
// }

/* func newErrordbResponse(err error) gin.H {
	emptyObject := struct{}{}
	return gin.H{
		"success":    false,
		"message":    err.Error(),
		"data":       emptyObject,
		"dataexists": false,
	}
} */

type AdminCreationResponse struct {
	EmployeeId                  int64  `json:"EmployeeId"`
	EmployeeName                string `json:"EmployeeName"`
	Designation                 string `json:"Designation"`
	RoleUserCode                int32  `json:"RoleUserCode"`
	RoleUserDescription         string `json:"RoleUserDescription"`
	Mobile                      string `json:"Mobile"`
	EmailID                     string `json:"EmailID"`
	UserName                    string `json:"UserName"`
	FacilityIDUniqueid          int64  `json:"FacilityIdUniqueid"`
	FacilityID                  string `json:"FacilityID"`
	AuthorityFacilityName       string `json:"AuthorityFacilityName"`
	FacilityType                string `json:"FacilityType"`
	ReportingOfficeFacilityId   string `json:"ReportingOfficeFacilityID"`
	ReportingOfficeFacilityName string `json:"ReportingOfficeFacilityName"`
	CreatedById                 int32  `json:"CreatedById"`
	CreatedByUserName           string `json:"CreatedByUserName"`
	CreatedByEmpId              int32  `json:"CreatedByEmpId"`
	CreatedByDesignation        string `json:"CreatedByDesignation"`
	CircleOfficeFacilityId      string `json:"CircleOfficeFacilityId"`
	CircleOfficeName            string `json:"CircleOfficeName"`
}
type AdminUserResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	DataExists bool        `json:"dataexists"`
}

type EmployeeMasterResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	DataeExists bool        `json:"dataexists"`
}

// type FirstTimeUserResponse struct {
// 	Success   bool   `json:"success"`
// 	Message   string `json:"message"`
// 	Data      UserData `json:"data"`
// 	DataeExists bool   `json:"dataexists"`
// }

// type UserData struct {
// 	EmployeeID   int    `json:"EmployeeID"`
// 	UserName     string `json:"UserName"`
// 	EmailStatus  string `json:"EmailStatus"`
// 	SMSStatus    string `json:"SMSStatus"`
// 	RoleUserCode string `json:"RoleUserCode"`
// 	Email        string `json:"Email"`
// 	Mobile       string `json:"Mobile"`
// }

type UserResponse struct {
	Success     bool              `json:"success"`
	Message     string            `json:"message"`
	UserDetail  FirstTimeUserData `json:"data"`
	DataeExists bool              `json:"dataexists"`
}

type FirstTimeUserData struct {
	EmployeeID   int    `json:"EmployeeID"`
	UserName     string `json:"UserName"`
	EmailStatus  string `json:"EmailStatus"`
	SMSStatus    string `json:"SMSStatus"`
	RoleUserCode string `json:"RoleUserCode"`
	Email        string `json:"Email"`
	Mobile       string `json:"Mobile"`
}

type UserResetSaveNewPasswordResponse struct {
	UserName string `json:"UserName"`
	Message  string `json:"message"`
}

// UserResetValidateUserNameResponse represents the structure of the response for UserResetValidateUserName API.
type UserResetValidateUserNameData struct {
	EmployeeID   string `json:"EmployeeID"`
	UserName     string `json:"UserName"`
	SMSStatus    string `json:"SMSStatus"`
	Email        string `json:"Email"`
	Mobile       string `json:"Mobile"`
	RoleUserCode string `json:"RoleUserCode"`
}
type UserResetValidateUserNameResponse struct {
	Success    bool                          `json:"success"`
	Message    string                        `json:"message"`
	Data       UserResetValidateUserNameData `json:"data"`
	DataExists bool                          `json:"dataexists"`
}

type UserResetValidateOTPResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	DataeExists bool        `json:"dataexists"`
}

// AdminUser represents the structure of the admin user data in the response.
type AdminUser struct {
	ID                          int    `json:"id"`
	EmployeeID                  string `json:"EmployeeID"`
	EmployeeName                string `json:"EmployeeName"`
	RoleUserCode                string `json:"RoleUserCode"`
	RoleUserDescription         string `json:"RoleUserDescription"`
	FacilityID                  string `json:"FacilityID"`
	AuthorityFacilityName       string `json:"AuthorityFacilityName"`
	FacilityType                string `json:"FacilityType"`
	ReportingOfficeFacilityID   string `json:"ReportingOfficeFacilityID"`
	ReportingOfficeFacilityName string `json:"ReportingOfficeFacilityName"`
	CircleOfficeFacilityID      string `json:"CircleOfficeFacilityID"`
	CircleOfficeName            string `json:"CircleOfficeName"`
	Designation                 string `json:"Designation"`
	Token                       string `json:"Token"`
}

// VerifyAdminLoginResponse represents the structure of the response for VerifyAdminLoginn API.
type VerifyAdminLoginResponse struct {
	Success    bool      `json:"success"`
	Message    string    `json:"message"`
	DataExists bool      `json:"dataexists"`
	Data       AdminUser `json:"data"`
}
type AdminUserr struct {
	RoleUserCode          string `json:"RoleUserCode"`
	UserName              string `json:"UserName"`
	FacilityID            string `json:"FacilityID"`
	EmployeeName          string `json:"EmployeeName"`
	EmployeeID            string `json:"EmployeeID"`
	Designation           string `json:"Designation"`
	Mobile                string `json:"Mobile"`
	Email                 string `json:"Email"`
	AuthorityFacilityName string `json:"AuthorityFacilityName"`
	CAUsername            string `json:"CAUsername"`
	SMSStatus             string `json:"SMSStatus"`
}

// VerifyAdminUserLoginResponse represents the structure of the response for VerifyAdminUserLoginn API.
type VerifyAdminUserLoginResponse struct {
	Success    bool       `json:"success"`
	Message    string     `json:"message"`
	Data       AdminUserr `json:"data"`
	DataExists bool       `json:"dataexists"`
}

//ip

type VerifyIPApplicationResponse struct {
	Success    bool                    `json:"success"`
	Message    string                  `json:"message"`
	Data       VerifyIPApplicationData `json:"data"`
	DataExists bool                    `json:"dataexists"`
}

type VerifyIPApplicationData struct {
	EmployeeID         string `json:"EmployeeID"`
	ApplicationStatus  string `json:"ApplicationStatus"`
	ApplicationNumber  string `json:"ApplicationNumber"`
	ApplicationRemarks string `json:"ApplicationRemarks"`
	EmailStatus        string `json:"EmailStatus"`
	SMSStatus          string `json:"SMSStatus"`
	RoleUserCode       string `json:"RoleUserCode"`
	Email              string `json:"Email"`
	Mobile             string `json:"Mobile"`
}
type GetAllCAPendingVerificationsResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	DataExists bool        `json:"dataexists"`
}
type UpdateCenterCodeResponse struct {
	Success    bool     `json:"success"`
	Message    string   `json:"message"`
	Data       struct{} `json:"data"`
	DataExists bool     `json:"dataexists"`
}
type ResponseData struct {
	EmployeeID        int    `json:"EmployeeID"`
	ApplicationStatus string `json:"ApplicationStatus"`
	ApplicationNumber string `json:"ApplicationNumber"`
	EmailStatus       string `json:"EmailStatus"`
	SMSStatus         string `json:"SMSStatus"`
	Email             string `json:"Email"`
	Mobile            string `json:"Mobile"`
}
type CreateGDSPMResponse struct {
	Success    bool         `json:"success"`
	Message    string       `json:"message"`
	Data       ResponseData `json:"data"`
	DataExists bool         `json:"dataexists"`
}
type GetGDSPMApplicationsResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	DataExists bool        `json:"dataexists"`
}
type VerifyGDSPMVAApplicationsResponse struct {
	Success    bool                          `json:"success"`
	Message    string                        `json:"message"`
	Data       VerifyGDSPMVAApplicationsData `json:"data"`
	DataExists bool                          `json:"dataexists"`
}

// VerifyGDSPMVAApplicationsData represents the data field in the response structure
type VerifyGDSPMVAApplicationsData struct {
	EmployeeID         int64  `json:"EmployeeID"`
	ApplicationStatus  string `json:"ApplicationStatus"`
	ApplicationRemarks string `json:"ApplicationRemarks"`
	RoleUserCode       string `json:"RoleUserCode"`
}
type GetGDSPMCAPendingOldRemarksByEmpIdResponse struct {
	Success    bool                                   `json:"success"`
	Message    string                                 `json:"message"`
	Data       GetGDSPMCAPendingOldRemarksByEmpIdData `json:"data"`
	DataExists bool                                   `json:"dataexists"`
}

// GetGDSPMCAPendingOldRemarksByEmpIdData represents the data structure for the response data
type GetGDSPMCAPendingOldRemarksByEmpIdData struct {
	EmployeeID         int64  `json:"employeeid"`
	ApplicationStatus  string `json:"applicationstatus"`
	ApplicationRemarks string `json:"applicationremarks"`
}

type CreateApplicationResponse struct {
	Success     bool          `json:"success"`
	Message     string        `json:"message"`
	Data        ApplicationData `json:"data"`
	DataExists  bool          `json:"dataexists"`
}

type ApplicationData struct {
	EmployeeID         int64  `json:"EmployeeID"`
	ApplicationStatus  string `json:"Application Status"`
	ApplicationNumber  string `json:"ApplicationNumber"`
	EmailStatus        string `json:"EmailStatus"`
	SMSStatus          string `json:"SMSStatus"`
	Email              string `json:"Email"`
	Mobile             string `json:"Mobile"`
}
type GetApplicationsResponse struct {
	Success    bool                   `json:"success"`
	Message    string                 `json:"message"`
	Data       interface{}            `json:"data"`
	DataExists bool                   `json:"dataexists"`
}

type PsgroupBData struct {
    EmployeeID          int    `json:"EmployeeID"`
    ApplicationStatus   string `json:"Application Status"`
    ApplicationNumber   string `json:"ApplicationNumber"`
    Mobile              string `json:"Mobile"`
    EmailID             string `json:"Email ID"`
    EmailDeliveryStatus bool   `json:"EmailDeliveryStatus"`
    SMSDeliveryStatus   bool   `json:"SMSDeliveryStatus"`
}

type PsGroupBResponse struct {
    Success    bool   `json:"success"`
    Message    string `json:"message"`
    Data       PsgroupBData   `json:"data"`
    DataExists bool   `json:"dataexists"`
}