package handlers

import (
	"context"
	"fmt"
	"net/http"
	"recruit/ent"
	"recruit/mail"
	"recruit/sms"
	"recruit/start"
	"recruit/util"
	"strconv"

	"github.com/gin-gonic/gin"

	ca_reg "recruit/payloadstructure/candidate_registration"
)

func TriggerSMSOTP(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - TriggerSMSOTP "
		var startFunction string = " - start - SubTriggerSMSOTP "
		var empMaster ca_reg.TriggerCandidateSMSOTP

		if err := gctx.ShouldBindJSON(&empMaster); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		response, status, stgError, dataStatus, err := start.SubTriggerSMSOTP(gctx, client, empMaster)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}

		var smssentstatus string = ""
		//----for sending SMS
		stringSMSOTP := strconv.Itoa(int(response.SmsOtp))
		smssentstatus = sms.SendSmsNew(ctx, client, empMaster.EmployeeId, 1, stringSMSOTP)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "SMS OTP Triggered successfully",
			"data":       empMaster.EmployeeId,
			"dataexists": dataStatus,
			"SMSstatus":  smssentstatus,
		})
	}
}
func TriggerEmailOTP(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - TriggerEmailOTP "
		var startFunction string = " - start - SubTriggerEmailOTP "
		var empMaster ca_reg.TriggerCandidateEmailOTP

		if err := gctx.ShouldBindJSON(&empMaster); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		response, status, stgError, dataStatus, err := start.SubTriggerEmailOTP(gctx, client, empMaster)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}

		stringEmailOTP := strconv.Itoa(int(response.EmailOtp))
		emailsentstatus := mail.SendEMailNew(ctx, client, empMaster.EmployeeId, 1, stringEmailOTP)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Email OTP Triggered successfully",
			"data":        empMaster.EmployeeId,
			"dataexists":  dataStatus,
			"Emailstatus": emailsentstatus,
		})
	}
}

func VerifyTriggerSMSOTP(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		var mainFunction string = " main - VerifyTriggerSMSOTP "
		var startFunction string = " - start - SubVerifyTriggerSMSOTP "
		var empMaster ca_reg.VerifyTriggerCandidateSMSOTP

		if err := gctx.ShouldBindJSON(&empMaster); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		response, status, stgError, dataStatus, err := start.SubVerifyTriggerSMSOTP(gctx, client, empMaster)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    response,
			"data":       empMaster.EmployeeId,
			"dataexists": dataStatus,
		})
	}
}

func VerifyTriggerEmailOTP(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		var mainFunction string = " main - VerifyTriggerEmailOTP "
		var startFunction string = " - start - SubVerifyTriggerEmailOTP "
		var empMaster ca_reg.VerifyTriggerCandidateEmailOTP

		if err := gctx.ShouldBindJSON(&empMaster); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		response, status, stgError, dataStatus, err := start.SubVerifyTriggerEmailOTP(gctx, client, empMaster)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    response,
			"data":       empMaster.EmployeeId,
			"dataexists": dataStatus,
		})
	}
}

// CreateEmployeeMaster godoc
// @Summary Create a new Employee Master
// @Description Create a new employee master record in the system
// @Tags Employee
// @Accept json
// @Produce json
// @Param empMaster body ca_reg.StrucCreateEmployeeMaster true "Employee Master Data"
// @Success 200 {object} start.EmployeeMasterResponse "Employee Master created  successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/empmaster/createemployee [post]
func CreateEmployeeMaster(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		errRemarks := "500 error DB Connection erorr -DBS01"
		errStatus := start.HandleDBErrorInitial(gctx, client, errRemarks)
		if errStatus == "error" {
			return
		}

		var mainFunction string = " main - CreateEmployeeMaster "
		var startFunction string = " - start - SubCreateEmployeeMaster "
		var empMaster ca_reg.StrucCreateEmployeeMaster

		if err := gctx.ShouldBindJSON(&empMaster); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := empMaster.Edges.LogData[0]

		employeemaster1, status, stgError, _, err := start.SubCreateEmployeeMaster(gctx, client, empMaster)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		var smssentstatus, emailsentstatus string
		//----for sending SMS
		stringEmployeeID := strconv.Itoa(int(employeemaster1.EmployeeID))
		if employeemaster1.VerifyStatus {
			smssentstatus = sms.SendSmsNew(ctx, client, stringEmployeeID, 35, employeemaster1.EmployeeName,
				employeemaster1.CreatedByUserName, "Approved", employeemaster1.MobileNumber)
			emailsentstatus = mail.SendEMailNew(ctx, client, stringEmployeeID, 35, employeemaster1.EmployeeName, employeemaster1.CreatedByUserName, "Approved", "nil", employeemaster1.EmailID)
		} else {
			smssentstatus = sms.SendSmsNew(ctx, client, stringEmployeeID, 34, employeemaster1.EmployeeName)

			emailsentstatus = mail.SendEMailNew(ctx, client, stringEmployeeID, 34, employeemaster1.EmployeeName, employeemaster1.ControllingAuthorityName)
		}
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Employee Master request submitted successfully",
			"data":        emptyObject,
			"dataexists":  false,
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"Email":       employeemaster1.EmailID,
			"Mobile":      employeemaster1.MobileNumber,
		})

	}
}

// ModifyEmployeeMaster godoc
// @Summary Modify an existing Employee Master
// @Description Modify an employee master record in the system
// @Tags Employee
// @Accept json
// @Produce json
// @Param modifyEmpMaster body ca_reg.StrucModifyEmployeeMaster true "Modify Employee Master Data"
// @Success 200 {object} start.EmployeeMasterResponse "Employee Master modified  successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/empmaster/modifyemployee [put]
func ModifyEmployeeMaster(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var modifyEmpMaster ca_reg.StrucModifyEmployeeMaster

		var mainFunction string = " main - UpdateAdminUser "
		var startFunction string = " - start.UpdateadminUser "
		if err := gctx.ShouldBindJSON(&modifyEmpMaster); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := modifyEmpMaster.Edges.LogData[0]
		employeemaster, status, stgError, _, err := start.SubModifyEmployeeMaster(client, modifyEmpMaster)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		fmt.Println("employeemaster1", employeemaster)
		//----for sending SMS
		if employeemaster.VerifyStatus {
			stringEmployeeID := strconv.Itoa(int(employeemaster.EmployeeID))
			smssentstatus := sms.SendSmsNew(ctx, client, stringEmployeeID, 35, employeemaster.EmployeeName,
				employeemaster.ModifiedByUserName, "Approved", employeemaster.MobileNumber)
			emailsentstatus := mail.SendEMailNew(ctx, client, stringEmployeeID, 35, employeemaster.EmployeeName, employeemaster.ModifiedByUserName, "Approved", "nil", employeemaster.EmailID)
			util.LoggerNew(client, logdata)
			gctx.JSON(http.StatusOK, gin.H{
				"success":     true,
				"message":     "Employee details modified / approved successfully",
				"data":        emptyObject,
				"dataexists":  false,
				"EmailStatus": emailsentstatus,
				"SMSStatus":   smssentstatus,
				"Email":       employeemaster.EmailID,
				"Mobile":      employeemaster.MobileNumber,
			})
		} else {
			stringEmployeeID := strconv.Itoa(int(employeemaster.EmployeeID))
			smssentstatus := sms.SendSmsNew(ctx, client, stringEmployeeID, 35, employeemaster.EmployeeName,
				employeemaster.ModifiedByUserName, "Rejected", employeemaster.MobileNumber)
			emailsentstatus := mail.SendEMailNew(ctx, client, stringEmployeeID, 35, employeemaster.EmployeeName, employeemaster.ModifiedByUserName, "Rejected", employeemaster.Cadre, employeemaster.EmailID)
			util.LoggerNew(client, logdata)
			gctx.JSON(http.StatusOK, gin.H{
				"success":     true,
				"message":     "Employee details rejected successfully",
				"data":        emptyObject,
				"dataexists":  false,
				"EmailStatus": emailsentstatus,
				"SMSStatus":   smssentstatus,
				"Email":       employeemaster.EmailID,
				"Mobile":      employeemaster.MobileNumber,
			})
		}
	}
}

// GetEmployeesBasedOnCA godoc
// @Summary Get employees based on CA Facility ID
// @Description Get details of employees based on CA Facility ID
// @Tags Employee
// @Accept json
// @Produce json
// @Param cafacilityid path string true "CA Facility ID"
// @Success 200 {object} start.EmployeeMasterResponse "Employee retrieved successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/empmaster/GetEmployeesBasedOnCA/{cafacilityid} [get]
func GetEmployeesBasedOnCA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - ViewEmployeeMaster - start - SubViewEmployeeMaster "
		var caFacilityID = gctx.Param("cafacilityid")

		employees, status, stgError, dataStatus, err := start.SubGetEmployeesBasedOnCA(ctx, client, caFacilityID)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       employees,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// ViewEmployeeMaster godoc
// @Summary View an existing Employee Master
// @Description Get details of an employee master record by employee ID
// @Tags Employee
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Success 200 {object} start.EmployeeMasterResponse "Employee master user retrieved successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/empmaster/viewemployee/{employeeid} [get]
func ViewEmployeeMaster(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - ViewEmployeeMaster - start - SubViewEmployeeMaster "
		empid, _ := strconv.ParseInt(gctx.Param("employeeid"), 10, 64)
		fmt.Println(empid)
		employee, status, stgError, dataStatus, err := start.SubViewEmployeeMaster(ctx, client, empid)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       employee,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetEmpMasterUsersByEmpId godoc
// @Summary Get employee master user by employee ID
// @Description Get details of an employee master user by their employee ID
// @Tags Employee
// @Accept  json
// @Produce  json
// @Param employeeid path int true "Employee ID"
// @Success 200 {object} start.EmployeeMasterResponse "Employee master user retrieved successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/admin/getempmasteruserbyempid/{employeeid} [get]
func GetEmpMasterUsersByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetEmpMasterUsersByEmpId - start.QueryEmpUsersByEmpId "
		id := gctx.Param("employeeid")
		empid, _ := strconv.ParseInt(id, 10, 64)
		adminmaster, status, stgError, dataStatus, err := start.QueryEmpUsersByEmpId(ctx, client, empid)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       adminmaster,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// create employee master

// CreatecandidateAUser godoc
// @Summary Create a candidate user
// @Description Create a new candidate user with the provided details
// @Tags Employee
// @Accept json
// @Produce json
// @Param candidate body ca_reg.CandidateCreation true "Candidate Details"
// @Success 200 {object} start.EmployeeMasterResponse "Employee master user retrieved successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/candidatecreate [post]
func CreatecandidateAUser(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		//candidateRequest := new(ent.EmployeeMaster)

		var candidateRequest ca_reg.CandidateCreation

		if err := gctx.ShouldBindJSON(&candidateRequest); err != nil {
			// gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			// return

			Action = "404"
			Remarks = "main - UpdateFirstTimeUserDetails - ShouldBindJSON error" + err.Error()
			util.SystemLogError(client, Action, Remarks)
			gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
			return
		}

		logdata := candidateRequest.Edges.LogData[0]
		candidate, status, err := start.CreateCandidate(gctx, client, candidateRequest)
		if err != nil {
			if status == 400 {
				logdata.Remarks = "400 error occured in start.CreateCandidate"
				util.LogErrorNew(client, logdata, err)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				logdata.Remarks = "422 error occured in start.CreateCandidate"
				util.LogErrorNew(client, logdata, err)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return
			} else {
				logdata.Remarks = "start.CreateCandidate - Other error"
				util.LogErrorNew(client, logdata, err)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			}
		}

		gctx.JSON(http.StatusOK, gin.H{"message": "Successfully created the CandidateUser", "candidate": candidate})
	}
}

// Query by Emp ID

type NewUserResponse struct {
	NewUser bool `json:"NewUser"`
}

func (r NewUserResponse) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`{"NewUser": %v}`, r.NewUser)
	return []byte(str), nil
}

func CheckUserByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		empID, _ := strconv.ParseInt(id, 10, 64)
		user, err := start.QueryUserMasterByEmpId(ctx, client, empID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response := NewUserResponse{
			NewUser: user,
		}
		gctx.JSON(http.StatusOK, gin.H{"data": response})
	}
	return gin.HandlerFunc(fn)
}

//CreateUserByEmpId

func CreateUserfromEmpMaster(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := gctx.Request.Context()
		id := gctx.Param("id")
		empID, _ := strconv.ParseInt(id, 10, 64)

		user, err := start.CreateUserByEmpId(ctx, client, empID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return the specified fields
		/*gctx.JSON(http.StatusOK, gin.H{
			"employeeID":    user.EmployeeID,
			"employeeName":  user.EmployeeName,
			"role":          user.Role,
			"mobile":        user.Mobile,
		})*/
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":         user.EmployeeID,
				"NewPasswordRequest": user.NewPasswordRequest,
				"RoleUserCode":       user.RoleUserCode,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// QueryUsersByEmpId
func GetUsersByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryUsersByEmpId(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

/* func UpdateUser(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newUser := new(ent.UserMaster)

		id := gctx.Param("id")

		userID, _ := strconv.ParseInt(id, 10, 64)

		if err := gctx.ShouldBindJSON(&newUser); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := start.UpdateUserByEmpID(client, int64(userID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":     user.EmployeeID,
				"PasswordStatus": "Password updated Successfully",
			},
		})
	}
	return gin.HandlerFunc(fn)
} */

// Employee Designation
func CreateEmployeeDesignations(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newEmployeeDesignations := new(ent.EmployeeDesignation)
		if err := gctx.ShouldBindJSON(&newEmployeeDesignations); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newEmployeeDesignations, err := start.CreateEmployeeDesignation(client, newEmployeeDesignations)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Employee Category")
	}
	return gin.HandlerFunc(fn)
}

func GetAllEmployeeDesignations(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		newEmployeeDesignations, err := start.QueryEmployeeDesignations(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": newEmployeeDesignations})
	}

	return gin.HandlerFunc(fn)

}

// For Employees
func CreateEmployee(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newEmployees := new(ent.Employees)
		if err := gctx.ShouldBindJSON(&newEmployees); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newEmployees, err := start.CreateEmployeeProfile(client, newEmployees)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Employee Profile")
	}
	return gin.HandlerFunc(fn)
}

func GetAllEmployees(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		newEmployees, err := start.QueryEmployees(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": newEmployees})
	}

	return gin.HandlerFunc(fn)

}

func GetEmployeesByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		EmpID, _ := strconv.ParseInt(id, 10, 32)
		emps, err := start.QueryEmployeesWithID(ctx, client, int32(EmpID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": emps})
	}
	return gin.HandlerFunc(fn)
}
func Updateemployeeverifydetails(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newEmp := new(ent.Employees)
		id := gctx.Param("id")
		empID, _ := strconv.ParseInt(id, 10, 32)
		if err := gctx.ShouldBindJSON(&newEmp); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		empupdated, err := start.UpdateVerificationDetails(client, int32(empID), newEmp)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": empupdated})

	}

	return gin.HandlerFunc(fn)
}
