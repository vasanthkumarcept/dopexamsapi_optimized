package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"log"
	//"recruit/authentication"
	"recruit/ent"
	"recruit/mail"
	"recruit/sms"
	"recruit/start"
	"recruit/util"

	//"recruit/ent/usermaster"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	ca_reg "recruit/payloadstructure/candidate_registration"
)

// UserResetValidateUserName validates a username during the password reset process.
// @Summary Validate username during password reset
// @Description Validate the username and send OTP for password reset.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param request body ca_reg.StrucUserResetValidateUserName true "User reset validate username request"
// @Success 200 {object} start.UserResetValidateUserNameResponse "Successful response"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/reset/validateusername [post]
func UserResetValidateUserName(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - UserResetValidateUserName "
		var startFunction string = " - start - SubUserResetValidateUserName "
		var newUser ca_reg.StrucUserResetValidateUserName

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
		user, status, stgError, dataStatus, err := start.SubUserResetValidateUserName(ctx, client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		//smsotp := util.GenerateOTP()
		stringSMSOTP := strconv.Itoa(int(user.OTP))
		smssentstatus := sms.SendSmsNew(ctx, client, newUser.UserName, 31, stringSMSOTP)
		logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "For user validation OTP triggered to SMS ",
			"data": gin.H{
				"EmployeeID": user.EmployeeID,
				"UserName":   user.UserName,
				//"EmailStatus":  emailsentstatus,
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

// UserResetValidateOTP godoc
// @Summary Validate OTP for User Reset
// @Description Validates the OTP for user reset process
// @Tags Users
// @Accept  json
// @Produce  json
// @Param newUser body ca_reg.StrucUserResetValidateOTP true "User Reset Validate OTP"
// @Success 200 {object} start.UserResetValidateOTPResponse "User OTP intiated"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/reset/validateOTP [post]
func UserResetValidateOTP(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var newUser ca_reg.StrucUserResetValidateOTP
		var mainFunction string = " main - UserResetValidateOTP "
		var startFunction string = " - start - SubUserResetValidateOTP "
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
		user, status, stgError, dataStatus, err := start.SubUserResetValidateOTP(ctx, client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return

		}
		logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "User OTP intiated",
			"data":       user,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// UserResetSaveNewPassword saves a new password for a user after reset.
// @Summary Save new password after user reset
// @Description Save new password for a user after password reset.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param request body ca_reg.StrucUserResetSaveNewPassword true "User reset save new password request"
// @Success 200 {object} start.UserResetSaveNewPasswordResponse "User password reset successfully "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/reset/savenewpassword [post]
func UserResetSaveNewPassword(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		var mainFunction string = " main - UserResetSaveNewPassword "
		var startFunction string = " - start SubUserResetSaveNewPassword "
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var newUser ca_reg.StrucUserResetSaveNewPassword

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

		_, status, stgError, dataStatus, err := start.SubUserResetSaveNewPassword(ctx, client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return

		}
		logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User password reset successfully ",
			"data": gin.H{
				"UserName": newUser.UserName,
			},
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// FirstTimeUserCreation godoc
// @Summary Create a new user for the first time
// @Description Create a new user and send OTP for verification
// @Tags Users
// @Accept json
// @Produce json
// @Param newUser body ca_reg.CandidateRegistrationGenerateOTP true "New User Data"
// @Success 200 {object} start.UserResponse "New User created successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/new/submit [post]
func FirstTimeUserCreation(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - FirstTimeUserCreation "
		var startFunction string = " - start NewCreateUserByEmpId "
		var newUser ca_reg.CandidateRegistrationGenerateOTP

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]
		decryptUsername, err := util.Decrypt(newUser.UserName)
		logdata.Userid = decryptUsername
		if err != nil {
			Remarks = mainFunction + " -User Name tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN01", logdata, client, Remarks)
			return
		}
		decryptPassword, err := util.Decrypt(newUser.Password)
		if err != nil {
			Remarks = mainFunction + " - Password tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN02", logdata, client, Remarks)
			return
		}
		newUser.UserName = decryptUsername
		newUser.Password = decryptPassword
		user, status, stgError, _, err := start.NewCreateUserByEmpId(ctx, client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var smssentstatus string = ""
		//----for sending SMS
		stringSMSOTP := strconv.Itoa(int(user.OTP))
		smssentstatus = sms.SendSmsNew(ctx, client, decryptUsername, 1, stringSMSOTP)

		stringEmailOTP := strconv.Itoa(int(user.EmailOTP))
		emailsentstatus := mail.SendEMailNew(ctx, client, newUser.UserName, 1, stringEmailOTP)

		logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "OTP triggering initiated ",
			"data":        user,
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"dataexists":  true,
		})
	}
	return gin.HandlerFunc(fn)
}

// EditMobileorEmailCredentials godoc
// @Summary Edit Mobile or Email Credentials
// @Description Initiate OTP triggering for changing mobile or email credentials
// @Tags Users
// @Accept  json
// @Produce  json
// @Param newUser body ca_reg.CandidateChangeEmailMobileGenerateOTP true "Candidate Change Email Mobile Generate OTP"
// @Success 200 {object} start.UserResponse "OTP triggering initiated",
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/users/trigotp/change/emailmobile [post]
func EditMobileorEmailCredentials(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = "main - EditMobileorEmailCredentials "
		var startFunction string = " - start MobileEmailchangeUserByUsername "
		var newUser ca_reg.CandidateChangeEmailMobileGenerateOTP

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]

		user, status, stgError, _, err := start.MobileEmailchangeUserByUsername(ctx, client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var smssentstatus, emailsentstatus, emailsentstatusnew, smssentstatusnew string

		newemailotp := user.EmailOTPNew
		newmobileotp := user.OTPNew
		oldMobileOTP := strconv.Itoa(int(user.OTP))
		oldEmailOTP := strconv.Itoa(int(user.EmailOTP))
		// Send OTP based on MobileEmail field
		if newUser.MobileEmail == "EmailID" {
			//smssentstatus = sms.SendSmsNew(ctx, client, newUser.OldMobile, 1, oldMobileOTP)
			smssentstatus = sms.SendSmsNew(ctx, client, newUser.UserName, 39, newUser.OldMobile, oldEmailOTP)

			emailsentstatus = mail.SendEMailNew(ctx, client, newUser.UserName, 37, newUser.OldEmailID, oldEmailOTP)
			emailsentstatusnew = mail.SendEMailNew(ctx, client, newUser.UserName, 37, newUser.NewEmailID, strconv.Itoa(int(newemailotp)))
		} else if newUser.MobileEmail == "Mobile" {
			emailsentstatus = mail.SendEMailNew(ctx, client, newUser.UserName, 37, newUser.OldEmailID, oldMobileOTP)
			smssentstatus = sms.SendSmsNew(ctx, client, newUser.UserName, 39, newUser.OldMobile, oldMobileOTP)
			smssentstatusnew = sms.SendSmsNew(ctx, client, newUser.UserName, 39, newUser.NewMobile, strconv.Itoa(int(newmobileotp)))
		}

		// logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "OTP triggering initiated",
			"data": gin.H{
				"EmployeeID":         user.EmployeeID,
				"UserName":           user.UserName,
				"EmailStatus":        emailsentstatus,
				"SMSStatus":          smssentstatus,
				"RoleUserCode":       user.RoleUserCode,
				"Email":              user.EmailID,
				"Mobile":             user.Mobile,
				"NewEmailOTPStatus":  emailsentstatusnew,
				"NewMobileOTPStatus": smssentstatusnew,
			},
			"dataexists": true,
		})
	}
	return gin.HandlerFunc(fn)
}

// UpdateFirstTimeUserDetails godoc
// @Summary Update first-time user details
// @Description Update details of a first-time user based on provided information and OTP validation
// @Tags Users
// @Accept json
// @Produce json
// @Param newUser body ca_reg.CandidateRegistrationOTP true "New User Details"
// @Success 200 {object} start.UserResponse "New User updated successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/new/update [post]
func UpdateFirstTimeUserDetails(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		// Parse the request body into an AdminLogin entity
		var newUser ca_reg.CandidateRegistrationOTP

		var mainFunction string = " main - UpdateFirstTimeUserDetails "
		var startFunction string = " - start - NewValidateLoginUser "

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]
		decryptUsername, err := util.Decrypt(newUser.UserName)
		logdata.Userid = decryptUsername
		if err != nil {
			Remarks = mainFunction + " -User Name tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN01", logdata, client, Remarks)
			return
		}
		decryptPassword, err := util.Decrypt(newUser.Password)
		if err != nil {
			Remarks = mainFunction + " - Password tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN02", logdata, client, Remarks)
			return
		}
		newUser.UserName = decryptUsername
		newUser.Password = decryptPassword

		// Validate the newUser credentials using the ValidateLoginUser function
		user, status, stgError, _, err := start.NewValidateLoginUser(client, &ent.UserMaster{
			UserName: newUser.UserName,
			OTP:      newUser.OTP,
			Password: newUser.Password,
			EmailOTP: newUser.EmailOTP,
		})
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		emailsentstatus := mail.SendEMailNew(ctx, client, decryptUsername, 2, decryptPassword)

		smssentstatus := sms.SendSmsNew(ctx, client, newUser.UserName, 2)

		logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"Message":     "User registered successfully, now you can able to Login !!",
			"data":        user,
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"dataexists":  true,
		})

	}
	return gin.HandlerFunc(fn)
}

// UpdateUserDetails godoc
// @Summary Update user details
// @Description Update user details with the given information
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body ca_reg.StrucProfileEmployeeMaster true "User details"
// @Success 200 {object} start.UserResponse " User Profile  updated successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/employee/updateprofile [post]
func UpdateUserDetails(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		// Parse the request body into an AdminLogin entity
		var newUser ca_reg.StrucProfileEmployeeMaster

		var mainFunction string = " main - UpdateUserDetails "
		var startFunction string = " - start - ModifyEmployeeMasterUserMaster "

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]

		// Validate the newUser credentials using the ValidateLoginUser function
		_, status, stgError, dataexist, err := start.ModifyEmployeeMasterUserMaster(client, newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		// logdata.Userid = decryptUsername
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"Message":    "User Details Updated successfully !!",
			"data":       "",
			"dataexists": dataexist,
		})

	}
	return gin.HandlerFunc(fn)
}

// ValidateOTPUserEmailMobile godoc
// @Summary Validate OTP for User's Email or Mobile
// @Description Verify OTP sent to user's email or mobile number
// @Tags Users
// @Accept  json
// @Produce  json
// @Param newUser body ca_reg.CandidateEditProfileOTP true "Candidate Edit Profile OTP"
// @Success 200 {object} start.UserResponse "OTP verified successfully, now you can update the details!"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/users/validate/otp [post]
func ValidateOTPUserEmailMobile(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var newUser ca_reg.CandidateEditProfileOTP

		var mainFunction string = "main - ValidateOTPUserEmailMobile"
		var startFunction string = " - start - ValidateOTPUserEmailorMobile"

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks := "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newUser.Edges.LogData[0]

		_, status, stgError, _, err := start.ValidateOTPUserEmailorMobile(client, newUser)
		if err != nil {
			Remarks := mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"Message":    "OTP verified successfully, now you can update the details!",
			"data":       "",
			"dataexists": true,
		})
	}
	return gin.HandlerFunc(fn)
}

// UpdateUserPassword godoc
// @Summary Update user password
// @Description Update the password of a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param newUser body ca_reg.CreateUserRequest true "User Details"
// @Success 200 {object} start.UserResponse "Password updated Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/updatepassword/{id} [put]
/* func UpdateUserPassword(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newUser := new(ent.UserMaster)

		id := gctx.Param("empid")

		userID, _ := strconv.ParseInt(id, 10, 64)

		var mainFunction string = " main - UpdateUserPassword "
		var startFunction string = " - start - UpdateUserPasswordByEmpID "

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		user, status, stgError, _, err := start.UpdateUserPasswordByEmpID(client, int64(userID), newUser)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
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

// ChangeUserPassword godoc
// @Summary Change user password
// @Description Change the password of a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param requestBody body RequestBody true "Change Password Request"
// @Success 200 {object} start.UserResponse "New Password updated Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/changepassword/{id} [put]
func ChangeUserPassword(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var requestBody ca_reg.RequestBody
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - ChangeUserPassword "
		var startFunction string = " - start - ChangeUserPasswordByEmpID "
		if err := gctx.ShouldBind(&requestBody); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := requestBody.Edges.LogData[0]

		//fmt.Println(requestBody.CurrentPassword)
		empID, _ := strconv.ParseInt(gctx.Param("id"), 10, 64)

		user, status, stgError, _, err := start.ChangeUserPasswordByEmpID(client, empID, requestBody.CurrentPassword, requestBody.NewPassword)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		//body := "Dear Customer,your new password is : " + newPassword + ", please do not share it with anyone - INDPOST"
		statusemail := mail.SendEMailNew(ctx, client, user.UserName, 14, user.Password)

		statussms := sms.SendSmsNew(ctx, client, user.UserName, 14, user.Password)

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":     user.EmployeeID,
				"PasswordStatus": "New Password updated Successfully",
				"EmailStatus":    statusemail,
				"SMSStatus":      statussms,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// @Summary Change user password
// @Description Change the password of a user by their username
// @Tags Users
// @Accept json
// @Produce json
// @Param requestBody body ca_reg.RequestBody true "Change Password Request"
// @Success 200 {object} start.UserResponse "New Password updated Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/ChangeUserPassword [put]
func ChangeUserPasswordNew(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var requestBody ca_reg.RequestBody
		var mainFunction string = " main - ChangeUserPassword "
		var startFunction string = " - start -ChangeUserPasswordByUsername "
		if err := gctx.ShouldBind(&requestBody); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := requestBody.Edges.LogData[0]
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

		user, status, stgError, dataStatus, err := start.ChangeUserPasswordByUsername(client, decryptUsername, decryptCurrentPassword, decryptNewPassword)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		startFunction = " - start - SendSmsNew "
		//----for sending SMS
		//		msg = "Dear " + empid + " Password is " + password + "-DOPExam-INDIAPOST"

		statussms := sms.SendSmsNew(ctx, client, user.UserName, 32, user.Password)

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User password reset successfully",
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

// VerifyUserLogin godoc
// @Summary Verify user login
// @Description Verify a user's login credentials
// @Tags Users
// @Accept json
// @Produce json
// @Param user body ca_reg.CandidateRegistrationGenerateOTP true "User login data"
// @Success 200 {object} start.EmployeeMasterResponse "User login verified successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/login [post]
func VerifyUserLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		// Parse the request body into an AdminLogin entity
		var newUser ca_reg.CandidateRegistrationGenerateOTP
		var mainFunction string = " main - VerifyUserLogin "
		var startFunction string = " - start - ValidateLoginUser"

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		// Validate the struct fields (if necessary)
		if newUser.UserName == "" || newUser.Password == "" || newUser.Edges.LogData == nil || newUser.Edges.LogData[0].Userid == "" || newUser.Edges.LogData[0].Action == "" || newUser.Edges.LogData[0].Usertype == "" {
			Remarks = "400 error from " + mainFunction + " - API Struct validation "
			start.MainHandleError(gctx, client, Remarks, " -HA02", gctx.GetHeader("UserName"))
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
		decryptPassword, err := util.Decrypt(newUser.Password)
		if err != nil {
			Remarks = mainFunction + " - Password tampered"
			start.StartErrorHandlerWithLog(gctx, err, 400, " -EN02", logdata, client, Remarks)
			return
		}

		user, status, stgError, dataStatus, err := start.ValidateLoginUser(client, &ent.UserMaster{
			UserName: decryptUsername,
			Password: decryptPassword,
		})
		logdata.Userid = decryptUsername
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "User login verified successfully",
			"data":       user,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// get data from employee master
func GetCandidateUsersByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("employeeid")
		var mainFunction string = " main - GetCandidateUsersByEmpId "
		var startFunction string = " - start - QueryCandidateUsersByEmpId "
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, status, stgError, _, err := start.QueryCandidateUsersByEmpId(ctx, client, empid)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}
func UpdateCandidateUsersByEmpId(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		var newUser ent.UserMaster

		id := gctx.Param("id")
		id1, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		logdata := newUser.Edges.LogData[0]
		fmt.Println(logdata)
		ctx := context.Background()
		_, logerr := client.Logs.Create().
			SetUsertype(logdata.Usertype).
			SetRemarks("Candidate Updation  Attempted").
			SetAction(logdata.Action).
			SetIpaddress(logdata.Ipaddress).
			SetDevicetype(logdata.Devicetype).
			SetOs(logdata.Os).
			SetBrowser(logdata.Browser).
			SetLatitude(logdata.Latitude).
			SetLongitude(logdata.Longitude).
			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
			Save(ctx)

		if logerr != nil {
			log.Println(logerr.Error())
		}

		user, err := start.UpdateCandidateUser(client, id1, &newUser)
		if err != nil {
			ctx := context.Background()
			_, logerr := client.Logs.Create().
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

			if logerr != nil {
				log.Println(logerr.Error())
			}

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, logerr1 := client.Logs.Create().
			SetUsertype(logdata.Usertype).
			SetRemarks("CandidateUser Updation is successfully Done").
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

		gctx.JSON(http.StatusOK, gin.H{"data": user})
	}
}
