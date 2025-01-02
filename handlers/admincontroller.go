package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"recruit/authentication"
	"recruit/ent"
	"recruit/mail"
	"recruit/sms"
	"recruit/start"
	"recruit/util"

	_ "github.com/lib/pq"

	ca_reg "recruit/payloadstructure/candidate_registration"
)

// CreateAdminUser creates a new admin user
// @Summary Create a new admin user
// @Description Creates a new admin user with the provided information
// @Tags Admin
// @Accept json
// @Produce json
// @Param adminRequest body ca_reg.AdminCreation true "Admin creation request"
// @Success 200 {object} start.AdminCreationResponse "Admin creation succesfully registerd"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/profileadmin [post]
func CreateAdminUser(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var adminRequest ca_reg.AdminCreation
		//ctx := context.Background()
		var mainFunction string = " main - CreateAdminUser "
		var startFunction string = " - start - CreateAdmin "

		if err := gctx.ShouldBindJSON(&adminRequest); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return

		}
		// Validate the struct fields (if necessary)

		logdata := adminRequest.Edges.LogData[0]

		admin, status, stgError, dataStatus, err := start.CreateAdmin(gctx, client, adminRequest)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		stringEmployeeID := strconv.Itoa(int(admin.EmployeeId))
		smssentstatus := sms.SendSmsNew(ctx, client, stringEmployeeID, 36, admin.EmployeeName,
			admin.RoleUserDescription, admin.UserName)

		emailsentstatus := mail.SendEMailNew(ctx, client, stringEmployeeID, 36, admin.EmployeeName,
			admin.RoleUserDescription, admin.UserName)

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Admin user created  successfully",
			"data":        admin,
			"dataexists":  dataStatus,
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"Email":       admin.EmailID,
			"Mobile":      admin.Mobile,
		})
	}
}

// VerifyAdminLoginn godoc
// @Summary Verify Admin Login
// @Description Authenticates an admin user based on OTP
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param newUser body ca_reg.AdminLoginValidateOTP true "Admin Login Validate OTP"
// @Success 200 {object} start.VerifyAdminLoginResponse "Admin user authenticated  successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/adminusers/verifyloginn [post]
func VerifyAdminLoginn(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var mainFunction string = " main - VerifyAdminLoginn "
		var startFunction string = " - start - ValidateAdminLoginn "

		/*
			} */
		var newUser ca_reg.AdminLoginValidateOTP

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]
		decryptUsername, err := util.Decrypt(newUser.UserName)
		logdata.Userid = decryptUsername
		if err != nil {
			Remarks = mainFunction + " - User Name tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -HA03", logdata, client, Remarks)
			return
		}

		logdata.Userid = decryptUsername
		// Validate the newUser credentials using the ValidateAdminLogin function
		user, status, token, stgError, dataStatus, err := start.ValidateAdminLoginn(client, &ent.AdminMaster{
			UserName: decryptUsername,
			OTP:      newUser.OTP,
			//UidToken: newUser.UidToken,
		})
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Admin user authenticated  successfully",
			"dataexists": dataStatus,
			"data":       user,
			"token":      token,
		})

	}
	return gin.HandlerFunc(fn)
}

// @Summary Verify Admin User Login
// @Description Authenticates an admin user based on username and password, and sends an OTP to the registered mobile number
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param newUser body ca_reg.AdminLoginGenerateOTP true "Admin Login Generate OTP"
// @Success 200 {object} start.VerifyAdminUserLoginResponse "OTP sent to registered mobile number"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/adminusers/verifyuserr [post]
func VerifyAdminUserLoginn(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		/* 		errRemarks := "500 error DB Connection erorr -DBS01"
		   		errStatus := start.HandleDBErrorInitial(gctx, client, errRemarks)
		   		if errStatus == "error" {
		   			return
		   		} */
		var mainFunction string = " main - VerifyAdminUserLoginn "
		var startFunction string = " - start - ValidateAdminUserLoginn "

		// Parse the request body into an AdminLogin entity
		var newUser ca_reg.AdminLoginGenerateOTP
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		logdata := newUser.Edges.LogData[0]

		decryptUsername, err := util.Decrypt(newUser.UserName)
		logdata.Userid = decryptUsername
		if err != nil {
			Remarks = mainFunction + " - User Name tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -HA03", logdata, client, Remarks)
			return
		}

		decryptPassword, error := util.Decrypt(newUser.Password)
		if error != nil {
			Remarks = mainFunction + " - Password  tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -HA04", logdata, client, Remarks)
			return
		}
		logdata.Userid = decryptUsername
		// Validate the newUser credentials using the ValidateLoginUser function
		user, status, stgError, dataStatus, otp, err := start.ValidateAdminUserLoginn(client, &ent.AdminMaster{
			UserName: decryptUsername,
			Password: decryptPassword,
		})
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		stringSMSOTP := strconv.Itoa(int(otp))
		smssentstatus := sms.SendSmsNew(ctx, client, decryptUsername, 31, stringSMSOTP)

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "OTP send to registered mobile number",
			"data":       user,
			"SMSStatus":  smssentstatus,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// @Summary Update Admin User
// @Description Update an existing admin user's information
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param AdminUser body ca_reg.UpdateAdminMasterStruc true "Admin user information"
// @Success 200 {object} start.AdminUserResponse "Admin Users updated Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/profileadminupdate [put]
func UpdateAdminUser(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var newUser ca_reg.UpdateAdminMasterStruc
		var mainFunction string = " main - UpdateAdminUser "
		var startFunction string = " - start.UpdateadminUser "
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]
		admin, status, stgError, dataStatus, err := start.UpdateadminUser(client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
		}

		stringEmployeeID := strconv.Itoa(int(admin.EmployeeId))
		smssentstatus := sms.SendSmsNew(ctx, client, stringEmployeeID, 36, admin.EmployeeName,
			admin.RoleUserDescription, admin.UserName)

		emailsentstatus := mail.SendEMailNew(ctx, client, stringEmployeeID, 36, admin.EmployeeName,
			admin.RoleUserDescription, admin.UserName)

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Admin user updated  successfully",
			"data":        admin,
			"dataexists":  dataStatus,
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"Email":       admin.EmailID,
			"Mobile":      admin.Mobile,
		})
	}
}

// @Summary Delete Admin User
// @Description Delete an admin user by username
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param username path string true "Admin Username"
// @Param AdminUser body ca_reg.DeleteAdminMasterStruc true "Admin user deletion information"
// @Success 200 {object} start.AdminUserResponse "Admin Users deleted Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/admin/deleteadminuserbyempid/{username} [delete]
func DeleteAdminUsersByUsername(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		/*
			} */
		var newUser ca_reg.DeleteAdminMasterStruc
		id := gctx.Param("username")
		var mainFunction string = " main - DeleteAdminUsersByUsername "
		var startFunction string = " - start - DeleteadminUser "
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			//ctx *gin.Context, client *ent.Client, Remarks string, handleError string, username string
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]

		//Log
		user, status, stgError, dataStatus, err := start.DeleteadminUser(client, id, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
		}
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Admin user deleted successfully",
			"data":       user,
			"dataexists": dataStatus,
		})

	}
}

// ChangeAdminPassword godoc
// @Summary Change the admin password
// @Description Changes the admin password with the provided data
// @Tags Admin
// @Accept json
// @Produce json
// @Param username path string true "Admin Username"
// @Param requestBody body ca_reg.RequestBody true "Request Body containing current and new password"
// @Success 200 {object} start.AdminUserResponse "New Password updated Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/admin/ChangeAdminPassword/{username} [put]
func ChangeAdminPassword(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var requestBody ca_reg.RequestBody
		var mainFunction string = " main - ChangeAdminPassword "
		var startFunction string = " - start -ChangeAdminUserPasswordByUsername "
		if err := gctx.ShouldBind(&requestBody); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := requestBody.Edges.LogData[0]
		//username, err := util.Decrypt(requestBody.UserName)
		decryptUsername, err := util.Decrypt(requestBody.UserName)
		if err != nil {
			Remarks = mainFunction + " - User Name tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN02", logdata, client, Remarks)
			return
		}

		decryptCurrentPassword, err := util.Decrypt(requestBody.CurrentPassword)
		if err != nil {
			Remarks = mainFunction + " - Current password tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN02", logdata, client, Remarks)
			return
		}

		decryptNewPassword, err := util.Decrypt(requestBody.NewPassword)
		if err != nil {
			Remarks = mainFunction + " - New password tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN02", logdata, client, Remarks)
			return
		}

		logdata.Userid = decryptUsername

		user, status, stgError, dataStatus, err := start.ChangeAdminUserPasswordByUsername(client, decryptUsername, decryptCurrentPassword, decryptNewPassword)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		//var smssentstatus string
		//startFunction = " - start - SendSMSGetAdminPassword "
		//----for sending SMS
		//msg := "Dear " + usernam + ", Password is " + password + "-DOPExam-INDIAPOST"
		statussms := sms.SendSmsNew(ctx, client, user.UserName, 32, user.Password)

		/* statussms, err := start.SendSMSGetAdminPassword(gctx.Request.Context(), client, user)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, statussms, stgError, client, Remarks)
			return
		} else {
			smssentstatus = "Success"
		} */
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Admin user password reset successfully",
			"data": gin.H{
				"EmployeeID": user.UserName,
				"Mobile":     user.Mobile,
				"SMSStatus":  statussms,
			},
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// AdminResetValidateOTP godoc
// @Summary Admin Reset Validate OTP
// @Description Validates the OTP for admin user reset
// @Tags Admin
// @Accept json
// @Produce json
// @Param newUser body ca_reg.StrucAdminResetValidateOTP true "Request Body containing the username and OTP"
// @Success 200 {object} start.AdminUserResponse "Admin user OTP Verified Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/admins/reset/validateOTP [post]
func AdminResetValidateOTP(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		/*
			} */
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - AdminResetValidateOTP "
		var startFunction string = " - start SubAdminResetValidateOTP "
		var newUser ca_reg.StrucAdminResetValidateOTP
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]
		decryptUsername, err := util.Decrypt(newUser.UserName)
		newUser.UserName = decryptUsername
		if err != nil {
			Remarks = mainFunction + " -  User Name tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN02", logdata, client, Remarks)
			return
		}
		logdata.Userid = decryptUsername
		newUser.UserName = decryptUsername

		user, status, stgError, dataStatus, err := start.SubAdminResetValidateOTP(ctx, client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return

		}
		logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Admin user OTP Verified Successfully",
			"data": gin.H{
				"UserName": user.UserName,
			},
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// AdminResetSaveNewPassword godoc
// @Summary Admin Reset Save New Password
// @Description Saves a new password for the admin user after validation
// @Tags Admin
// @Accept json
// @Produce json
// @Param newUser body ca_reg.StrucAdminResetSaveNewPassword true "Request Body containing the username and new password"
// @Success 200 {object} start.AdminUserResponse "Admin user password reset  successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/admins/reset/savenewpassword [post]
func AdminResetSaveNewPassword(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		/*
			} */
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var newUser ca_reg.StrucAdminResetSaveNewPassword
		var mainFunction string = " main - AdminResetSaveNewPassword "
		var startFunction string = " - start - SubAdminResetSaveNewPassword "
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]
		decryptUsername, error := util.Decrypt(newUser.UserName)
		if error != nil {
			Remarks = "400 error from " + mainFunction + " - API Struct validation "
			start.MainHandleError(gctx, client, Remarks, " -HA02", gctx.GetHeader("UserName"))
			return

		}
		logdata.Userid = decryptUsername
		newUser.UserName = decryptUsername
		decryptNewPassword, err := util.Decrypt(newUser.NewPassword)
		if err != nil {
			Remarks = mainFunction + " - Password tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN02", logdata, client, Remarks)
			return
		}
		newUser.NewPassword = decryptNewPassword

		_, status, stgError, dataStatus, err := start.SubAdminResetSaveNewPassword(ctx, client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Admin user password reset  successfully",
			"data": gin.H{
				"UserName": newUser.UserName,
			},
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// GetAdminUsersByEmpId godoc
// @Summary Get Admin Users by Employee ID
// @Description Retrieves admin user details based on the provided employee ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} start.AdminUserResponse "Admin users fetched by EMployeeID successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/admin/getadminuserbyempid/{id} [get]
func GetAdminUsersByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, status, err := start.QueryAdminUsersByEmpId(ctx, client, empid)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.QueryAdminUsersByEmpId " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "400"
				Remarks = "start.QueryAdminUsersByEmpId " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return
			} else {
				Action = "500"
				Remarks = "start.QueryAdminUsersByEmpId " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// @Summary Get All Admin Users by Facility and Role
// @Description Get a list of admin users by facility ID and role
// @Tags Admin
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param role path int true "Role"
// @Success 200 {object} start.AdminUserResponse "Admin users fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/admin/getadminuserbyfacilityrole/{facilityid}/{role} [get]
func GetAllAdminUsersByFacilityandRole(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		/*
			} */
		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		facilityID := gctx.Param("facilityid")
		role := gctx.Param("role")
		var startFunction string = " main - GetAllAdminUsersByFacilityandRole - start.QueryAdminUsersByfacility "
		roleNum, _ := strconv.ParseInt(role, 10, 32)
		adminUsers, status, stgError, dataStatus, err := start.QueryAdminUsersByfacility(ctx, client, facilityID, int32(roleNum))
		if err != nil {
			// err error, status int32, stgError string, client *ent.Client, Remarks string
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       adminUsers,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

func GetAllAdminUsersByRole(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		/*
			} */
		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		role := gctx.Param("role")
		var startFunction string = " main - GetAllAdminUsersByRole - start.QueryAdminUsersByRole "
		roleNum, _ := strconv.ParseInt(role, 10, 32)
		adminUsers, status, stgError, dataStatus, err := start.QueryAdminUsersByRole(ctx, client, int32(roleNum))
		if err != nil {
			// err error, status int32, stgError string, client *ent.Client, Remarks string
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       adminUsers,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetAdminUsersByUsername godoc
// @Summary Get Admin Users by Username
// @Description Retrieves admin user details based on the provided username
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "Username"
// @Success 200 {object} start.AdminUserResponse "Admin users fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/admin/getadminuserbyusername/{id} [get]
func GetAdminUsersByUsername(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		empid := gctx.Param("id")
		//var examID int32
		//empid, _ := strconv.ParseInt(id, 10, 64)
		circles, status, err := start.QueryAdminUsersByusername(ctx, client, empid)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.QueryAdminUsersByusername " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "400"
				Remarks = "start.QueryAdminUsersByusername " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return
			} else {
				Action = "500"
				Remarks = "start.QueryAdminUsersByusername " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// AdminResetValidateUserName godoc
// @Summary Validate Admin User by Username
// @Description Validates admin user details based on the provided username and triggers OTP for verification
// @Tags Admin
// @Accept json
// @Produce json
// @Param request body ca_reg.StrucAdminResetValidateUserName true "Admin user validation request"
// @Success 200 {object} start.AdminUserResponse "Admin users fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/admins/reset/validateusername [post]
func AdminResetValidateUserName(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		/*
			} */
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - AdminResetValidateUserName "
		var startFunction string = " - start - SubAdminResetValidateUserName "
		var newUser ca_reg.StrucAdminResetValidateUserName

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]
		decryptUsername, err := util.Decrypt(newUser.UserName)
		if err != nil {
			Remarks = mainFunction + " - User Name tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN02", logdata, client, Remarks)
			return
		}
		logdata.Userid = decryptUsername
		newUser.UserName = decryptUsername

		user, status, stgError, dataStatus, otp, err := start.SubAdminResetValidateUserName(ctx, client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		stringSMSOTP := strconv.Itoa(int(otp))
		smssentstatus := sms.SendSmsNew(ctx, client, newUser.UserName, 31, stringSMSOTP)

		logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "OTP triggered for admin user Validation",
			"data": gin.H{
				"UserName": user.UserName,

				"SMSStatus":    smssentstatus,
				"Email":        user.EmailID,
				"Mobile":       user.Mobile,
				"RoleUserCode": user.RoleUserCode,
			},
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

func VerifyAdminLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newUser := new(ent.UserMaster)
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(newUser.Edges.LogData) <= 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
			return
		}

		logdata := newUser.Edges.LogData[0]

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		_, logerr := client.Logs.Create().
			SetUserid(logdata.Userid).
			SetUsertype(logdata.Usertype).
			SetRemarks("Admin Login with otp Attempted").
			SetAction(logdata.Action).
			SetIpaddress(logdata.Ipaddress).
			SetDevicetype(logdata.Devicetype).
			SetOs(logdata.Os).
			SetBrowser(logdata.Browser).
			SetLatitude(logdata.Latitude).
			SetLongitude(logdata.Longitude).
			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
			Save(ctx)

		// Validate the newUser credentials using the ValidateAdminLogin function
		user, circleOfficeName, err := start.ValidateAdminLogin(client, &ent.UserMaster{
			UserName: newUser.UserName,
			OTP:      newUser.OTP,
		})
		if err != nil {
			ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
			defer cancel()
			_, logerr1 := client.Logs.Create().
				SetUserid(logdata.Userid).
				SetUsertype(logdata.Usertype).
				SetRemarks(err.Error()).
				SetAction(logdata.Action).
				SetIpaddress(logdata.Ipaddress).
				SetDevicetype(logdata.Devicetype).
				SetOs(logdata.Os).
				SetBrowser(logdata.Browser).
				SetLatitude(logdata.Latitude).
				SetLongitude(logdata.Longitude).
				SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
				Save(ctx)

			if logerr1 != nil {
				log.Println(logerr.Error())
			}
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err1 := client.Logs.Create().
			SetUserid(logdata.Userid).
			SetUsertype(logdata.Usertype).
			SetRemarks("Admin Logged With OTP Sucessfully").
			SetAction(logdata.Action).
			SetIpaddress(logdata.Ipaddress).
			SetDevicetype(logdata.Devicetype).
			SetOs(logdata.Os).
			SetBrowser(logdata.Browser).
			SetLatitude(logdata.Latitude).
			SetLongitude(logdata.Longitude).
			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
			Save(context.Background())
		if err1 != nil {
			log.Println("NO LOG ")
		}

		token := authentication.CreateToken(newUser.EmployeeName, client)
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"RoleUserCode": user.RoleUserCode,
				"UserName":     user.UserName,
				"FacilityID":   user.FacilityID,
				"LoginStatus":  "Verified successfully",
				"OfficeName":   circleOfficeName,
				"EmployeeName": user.EmployeeName,
				"Token":        token,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

//For Admin Login

func InsertAdminLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		adminnewUsers := new(ent.AdminLogin)
		if err := gctx.ShouldBindJSON(&adminnewUsers); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		adminnewUsers, err := start.CreateAdminLogin(client, adminnewUsers)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Admin Login")
	}
	return gin.HandlerFunc(fn)
}

func VerifyLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newUsers := new(ent.AdminLogin)
		if err := gctx.ShouldBindJSON(&newUsers); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUsers, err := start.ValidateAdminLoginUser(client, newUsers)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			//"message": "User " + newUsers.Username + " verified successfully " + "for the role " + newUsers.RoleName,

			"data": gin.H{
				"roleUserCode":  newUsers.RoleUserCode,
				"verifyRemarks": newUsers.VerifyRemarks,
			},
		})
	}
	return gin.HandlerFunc(fn)
}
