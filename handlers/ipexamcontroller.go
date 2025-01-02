package handlers

import (
	"context"
	"fmt"

	//"os"

	"net/http"
	"recruit/ent"
	"recruit/mail"
	"recruit/sms"
	"recruit/start"
	"recruit/util"
	"strconv"

	ca_reg "recruit/payloadstructure/candidate_registration"

	"github.com/gin-gonic/gin"
)

// CreateIPApplication godoc
// @Summary Create IP Application
// @Description Submit an new IP application based on provided by Applicant.
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param application body ca_reg.ApplicationIp true "IP Application Data"
// @Success 200 {object} start.VerifyIPApplicationResponse "IP Application submitted successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// // @Router /rect/ipexams/applications/submit [post]
// func CreateNewIPApplications(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {

// 		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
// 		defer cancel()
// 		var mainFunction string = " main - CreateNewIPApplications "
// 		var startFunction string = " - start CreateIPApplications "
// 		var newIPAppplns ca_reg.ApplicationIp
// 		if err := gctx.ShouldBindJSON(&newIPAppplns); err != nil {
// 			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
// 			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
// 			return
// 		}
// 		logdata := newIPAppplns.Edges.LogData[0]

// 		newIPApppln, status, stgError, dataStatus, err := start.CreateIPApplications(client, &newIPAppplns)

// 		if err != nil {
// 			Remarks = mainFunction + startFunction
// 			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
// 			return
// 		}

// 		var Smsuserid string = fmt.Sprint(newIPApppln.EmployeeID)
// 		var submittedDate string = newIPApppln.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
// 		//examshortname ,appno ,ca submittedDate
// 		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 3, " LDCE PA/SA to IP ", newIPApppln.ApplicationNumber, submittedDate)

// 		//----for sending SMS
// 		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE PA/SA to IP", newIPApppln.ApplicationNumber, "Submitted")

// 		util.LoggerNew(client, logdata)

// 		gctx.JSON(http.StatusOK, gin.H{
// 			gctx.JSON(http.StatusOK, gin.H{
// 				"success":       true,
// 				"message":       "IP Application submitted successfully",
// 				"EmailStatus":   emailsentstatus,
// 				"SMSStatus":     smssentstatus,
// 				"data":          appResponse, // Use the variable containing the response data
// 				"dataexists":    dataStatus,
// 			})
// 	return gin.HandlerFunc(fn)
// }

func CreateNewIPApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - CreateNewIPApplications "
		var startFunction string = " - start CreateIPApplications "
		var newIPAppplns ca_reg.ApplicationIp

		if err := gctx.ShouldBindJSON(&newIPAppplns); err != nil {
			Remarks := "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newIPAppplns.Edges.LogData[0]

		newIPApppln, status, stgError, dataStatus, err := start.CreateIPApplications(client, &newIPAppplns)
		if err != nil {
			fmt.Println(err)
			Remarks := mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var Smsuserid = fmt.Sprint(newIPApppln.EmployeeID)
		var submittedDate = newIPApppln.ApplnSubmittedDate.Format("02-01-2006 15:04:05")

		// Send email
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 3, " LDCE PA/SA to IP ", newIPApppln.ApplicationNumber, submittedDate)

		// Send SMS
		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE PA/SA to IP", newIPApppln.ApplicationNumber, "Submitted")

		// Log data
		util.LoggerNew(client, logdata)

		// Send JSON response
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "IP Application submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        newIPApppln,
			"dataexists":  dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

func CreateExamCenterHall(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		//defer cancel()
		var mainFunction string = " main - CreateExamCenterHall "
		var startFunction string = " - start SubCreateExamCenterHall "
		var examCenterHall ca_reg.StruExamCenterHall
		if err := gctx.ShouldBindJSON(&examCenterHall); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := examCenterHall.Edges.LogData[0]

		result, status, stgError, dataStatus, err := start.SubCreateExamCenterHall(client, &examCenterHall)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "IP Application submitted successfully",
			"data":       result,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

func ResetExamCenterHall(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		//defer cancel()
		var mainFunction string = " main - ResetExamCenterHall "
		var startFunction string = " - start SubResetExamCenterHall "
		var examCenterHall ca_reg.StruExamCenterHallReset
		if err := gctx.ShouldBindJSON(&examCenterHall); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := examCenterHall.Edges.LogData[0]

		result, status, stgError, dataStatus, err := start.SubResetExamCenterHall(client, &examCenterHall)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "IP Application ExamHall Reseted successfully",
			"data":       result,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// VerifyIPApplication godoc
// @Summary Verify IP Application By CA
// @Description Verifies an IP application based on provided details By CA.
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param application body ca_reg.VerifyApplicationIp true "IP Application Data"
// @Success 200 {object} start.VerifyIPApplicationResponse "IP application verified successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/applications/verify [put]
func VerifyIPApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - VerifyIPApplication "
		var startFunction string = " - start - UpdateIPApplicationRemarks "
		var newAppln ca_reg.VerifyApplicationIp
		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		var nonQualifyService string
		if newAppln.NonQualifyingService == nil || len(*newAppln.NonQualifyingService) == 0 {
			nonQualifyService = "No"
		} else {
			nonQualifyService = "Yes"
		}

		application, status, stgError, dataStatus, err := start.UpdateIPApplicationRemarks(client, &newAppln, nonQualifyService)
		logdata := newAppln.Edges.LogData[0]
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		var emailsentstatus string = ""
		if application.ApplicationStatus == "VerifiedByCA" {
			//examshortname ,appno ,ca,submittedDate ,recommended, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE PA/SA to IP ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.CAGeneralRemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE PA/SA to IP ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.CAGeneralRemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByCA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "IP application verified successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

type CandidateInfo struct {
	EmployeeID  string `json:"employeeID"`
	EmailStatus string `json:"emailStatus"`
	SMSStatus   string `json:"smsStatus"`
}

func EmailSmsTriggeringPenidngWithCandidate(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//:nodalofficeid:examyear
		examYear := gctx.Param("examyear")
		nodalOfficeId := gctx.Param("nodalofficeid")
		var mainFunction string = " main - EmailSmsTriggeringPenidngWithCandidate "
		var startFunction string = " - start - SubEmailSmsTriggeringPenidngWithCandidate "

		application, status, stgError, dataStatus, err := start.SubEmailSmsTriggeringPenidngWithCandidate(ctx, client, examYear, nodalOfficeId)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		var emailSentStatus string = ""
		var smsStatus string = ""
		var totalUpdatedCount int = 0
		var candidateData []CandidateInfo
		for _, app := range application {
			var submittedDate string = app.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
			var Smsuserid = fmt.Sprint(app.EmployeeID)
			emailSentStatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE PA/SA to IP ", app.ApplicationNumber, app.ControllingOfficeName, submittedDate, app.CAGeneralRemarks)
			smsStatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", app.ApplicationNumber, "Returned for correction")
			candidateData = append(candidateData, CandidateInfo{
				EmployeeID:  Smsuserid,
				EmailStatus: emailSentStatus,
				SMSStatus:   smsStatus,
			})
			totalUpdatedCount = totalUpdatedCount + 1
		}

		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": fmt.Sprintf("SMS and Email sent for %d candidates successfully", totalUpdatedCount),
			//			"updatedApplications": updatedApplications,
			"dataexists": dataStatus,
			"data":       candidateData,
		})
	}

	return gin.HandlerFunc(fn)
}
func ResetCenterIPApplicationNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		/* ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		*/
		var mainFunction string = " main - ResetCenterIPApplicationNA "
		var startFunction string = " - start - SubResetCenterIPApplicationNA "
		var newAppln ca_reg.NAApplicationIpCenterChange

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newAppln.Edges.LogData[0]

		application, status, stgError, dataStatus, err := start.SubResetCenterIPApplicationNA(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		/*
			var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
			var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
			var emailsentstatus string = ""
			if application.ApplicationStatus == "VerifiedByNA" {
				//examshortname ,appno ,ca,submittedDate ,recommended, comments
				emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE PA/SA to IP ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.NARemarks)
			} else if application.ApplicationStatus == "PendingWithCandidate" {
				//examshortname ,	appno, ca, submittedDate, comments
				emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE PA/SA to IP ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.NARemarks)
			}

			var smssentstatus string = ""
			if application.ApplicationStatus == "VerifiedByNA" {
				smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", application.ApplicationNumber, application.RecommendedStatus)
			} else if application.ApplicationStatus == "PendingWithCandidate" {
				smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", application.ApplicationNumber, "Returned for correction")
			}
		*/
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Nodal officer successfully resetted this IP application ",
			"data": gin.H{
				"EmployeeID":          application.EmployeeID,
				"ApplicationStatus":   application.ApplicationStatus,
				"Application Remarks": application.AppliactionRemarks,
				//"EmailStatus":         emailsentstatus,
				//	"SMSStatus":           smssentstatus,
				"RoleUserCode": application.RoleUserCode,
				//"Email":               application.EmailID,
				//"Mobile":              application.MobileNumber,
			},
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}
func IPApplicationCAEdit(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - IPApplicationCAEdit "
		var startFunction string = " - start - SubIPApplicationCAEdit "
		var newAppln ca_reg.ApplicationIpEditCA

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newAppln.Edges.LogData[0]

		application, status, stgError, dataStatus, err := start.SubIPApplicationCAEdit(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		//examshortname ,	appno, ca, submittedDate
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 7, " LDCE PA/SA to IP ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE PA/SA to IP", application.ApplicationNumber, "Re-submitted")

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "IP application re-submitted successfully",
			"data": gin.H{
				"EmployeeID":          application.EmployeeID,
				"ApplicationStatus":   application.ApplicationStatus,
				"Application Remarks": application.AppliactionRemarks,
				"EmailStatus":         emailsentstatus,
				"SMSStatus":           smssentstatus,
				"RoleUserCode":        application.RoleUserCode,
				"Email":               application.EmailID,
				"Mobile":              application.MobileNumber,
			},
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

func ResetIPApplicationNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		/* ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		*/
		var mainFunction string = " main - ResetIPApplicationNA "
		var startFunction string = " - start - SubResetIPApplicationNA "
		var newAppln ca_reg.NAVerifyApplicationIp

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newAppln.Edges.LogData[0]

		application, status, stgError, dataStatus, err := start.SubResetIPApplicationNA(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		/*
			var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
			var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
			var emailsentstatus string = ""
			if application.ApplicationStatus == "VerifiedByNA" {
				//examshortname ,appno ,ca,submittedDate ,recommended, comments
				emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE PA/SA to IP ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.NARemarks)
			} else if application.ApplicationStatus == "PendingWithCandidate" {
				//examshortname ,	appno, ca, submittedDate, comments
				emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE PA/SA to IP ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.NARemarks)
			}

			var smssentstatus string = ""
			if application.ApplicationStatus == "VerifiedByNA" {
				smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", application.ApplicationNumber, application.RecommendedStatus)
			} else if application.ApplicationStatus == "PendingWithCandidate" {
				smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", application.ApplicationNumber, "Returned for correction")
			}
		*/
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Nodal officer successfully resetted this IP application ",
			"data": gin.H{
				"EmployeeID":          application.EmployeeID,
				"ApplicationStatus":   application.ApplicationStatus,
				"Application Remarks": application.AppliactionRemarks,
				//"EmailStatus":         emailsentstatus,
				//	"SMSStatus":           smssentstatus,
				"RoleUserCode": application.RoleUserCode,
				//"Email":               application.EmailID,
				//"Mobile":              application.MobileNumber,
			},
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// Re-SubmitIPApplication godoc
// @Summary Re-Submit IP Application
// @Description Re-Submit an IP application based on provided by Applicant.
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param application body ca_reg.ResubmitApplicationIp true "IP Application Data"
// @Success 200 {object} start.VerifyIPApplicationResponse "IP application re-submitted successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/applications/resubmit [put]
func ResubmitApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - ResubmitApplication "
		var startFunction string = " - start - ResubmitApplication "
		var newAppln ca_reg.ResubmitApplicationIp
		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newAppln.Edges.LogData[0]
		application, status, stgError, dataStatus, err := start.ResubmitApplication(client, &newAppln)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		//examshortname ,	appno, ca, submittedDate
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 7, " LDCE PA/SA to IP ", application.ApplicationNumber, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE PA/SA to IP", application.ApplicationNumber, "Re-submitted")

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "IP application re-submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// Get IP Applications with EmpID - QueryIPExamApplicationsByEmpID

// GetIPApplicationsByEmpId godoc
// @Summary Get Applications details by employee ID
// @Description Fetches IP applications details for a given employee ID and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Application details fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/applications/getbyempid/{employeeid}/{examyear} [get]
func GetIPApplicationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//employeeid/:examyear
		id := gctx.Param("employeeid")
		examYear := gctx.Param("examyear")
		//var examID int32
		employeeID, _ := strconv.ParseInt(id, 10, 64)

		var mainFunction string = " main - GetIPApplicationsByEmpId "
		var startFunction string = " - start - QueryIPExamApplicationsByEmpID "

		ipapplns, status, stgError, dataStatus, err := start.QueryIPExamApplicationsByEmpIDNew(ctx, client, int64(employeeID), string(examYear))
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Application details fetched successfully",
			"data":       ipapplns,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetEmployeeDetailsByEmpId godoc
// @Summary Get Employee details by employee ID
// @Description Fetches Employee details for a given employee ID
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Employee details fetched"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/employees/search/byempid/{employeeid} [get]
func GetEmployeeDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("employeeid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		var mainFunction string = " main - GetEmployeeDetailsByEmpId "
		var startFunction string = " - start - QueryEmployeeMasterByEmpID "
		circles, status, stgError, dataStatus, err := start.QueryEmployeeMasterByEmpID(ctx, client, int64(empid))
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Employee details fetched",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetAllCAPendingVerifications godoc
// @Summary Get all pending CA verifications
// @Description Fetches all IP applications pending CA verifications for a given facility and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetched CA Verified applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/getallcapendingapplications/{facilityid}/{selectedyear} [get]
func GetAllCAPendingVerifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetAllCAPendingVerifications "
		var startFunction string = " - start - QueryIPApplicationsByCAVerificationsPending "
		//facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryIPApplicationsByCAVerificationsPending(ctx, client, facilityID, id1)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetched pending application with CA",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

func GetAllPendingApplicationsOnDeputation(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetAllPendingApplicationsOnDeputation "
		var startFunction string = " - start - SubGetAllPendingApplicationsOnDeputation "
		//facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.SubGetAllPendingApplicationsOnDeputation(ctx, client, facilityID, id1)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fethced application on deputation",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetAllCAVerifiedApplications godoc
// @Summary Get all CA verified applications for a given facility and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetched CA Verified applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/getallcaverifiedapplications/{facilityid}/{selectedyear} [get]
func GetAllCAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		//facilityid/:selectedyear

		var mainFunction string = " main - GetAllCAVerified "
		var startFunction string = " - start - QueryIPApplicationsByCAVerified "
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryIPApplicationsByCAVerified(ctx, client, facilityID, id1)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetched CA Verified applications",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)

}

// GetCAVerifiedDetailsByEmpId godoc
// @Summary Get CA verified details by employee ID
// @Description Fetches IP applications verified by CA for a given employee ID and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetched data based on employeeid with Verified status"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/getallcaverified/{employeeid}/{selectedyear} [get]
func GetCAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetCAVerifiedDetailsByEmpId "
		var startFunction string = " - start - QueryIPApplicationsByCAVerifiedByEmpID "
		//employeeid
		id := gctx.Param("employeeid")
		examYear := gctx.Param("selectedyear")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, status, stgError, dataStatus, err := start.QueryIPApplicationsByCAVerifiedByEmpID(ctx, client, int64(empid), examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetched data based on employeeid with Verified status",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetCAPendingDetailsByEmpId godoc
// @Summary Get CA Pending details by employee ID
// @Description Fetches IP applications CA pending for a given employee ID and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching application pending with CA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/getallcapending/{employeeid}/{selectedyear} [get]
func GetCAPendingDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetCAPendingDetailsByEmpId "
		var startFunction string = " - start - QueryIPApplicationsByCAPendingByEmpID "
		id := gctx.Param("employeeid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("selectedyear")
		circles, status, stgError, dataStatus, err := start.QueryIPApplicationsByCAPendingByEmpID(ctx, client, int64(empid), examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching application pending with CA",
			"data":       circles,
			"dataexists": dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// Return Previous Remarks

// GetCAPendingOldRemarksByEmpId godoc
// @Summary Get CA Pending with candiate with remarks details by employee ID
// @Description Fetches IP applications CA pending with candiate with remarks for a given employee ID and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching applicatins pending with candiate with remarks"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/caprevremarks/{employeeid}/{selectedyear} [get]
func GetCAPendingOldRemarksByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetCAPendingOldRemarksByEmpId "
		var startFunction string = " - start - GetOldCAApplicationRemarksByEmployeeID "
		id := gctx.Param("employeeid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("selectedyear")
		circles, status, stgError, dataStatus, err := start.GetOldCAApplicationRemarksByEmployeeID(ctx, client, int64(empid), examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Fetching applicatins pending with candiate with remarks",
			"data": gin.H{
				"EmployeeID":          circles.EmployeeID,
				"ApplicationStatus":   circles.ApplicationStatus,
				"Application Remarks": circles.AppliactionRemarks,
			},
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// GetRecommendationsByEmpID

// GetRecommendationsByEmpId godoc
// @Summary Get Application Recommendations by employee ID
// @Description Fetches IP applications Recommendations for a given employee ID
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching recommendations by employee ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/recommendations/{employeeid} [get]
func GetIPExamRecommendationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		id := gctx.Param("employeeid")
		var mainFunction string = " main - GetIPExamRecommendationsByEmpId "
		var startFunction string = " - start - GetRecommendationsByEmpID "
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		records, status, stgError, dataStatus, err := start.GetRecommendationsByEmpID(client, empid)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching recommendations by employee ID",
			"data":       records,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// Get All NA Verified records by CA ...

// GetAllNAVerified godoc
// @Summary Get all NA Verified by facilityID and year
// @Description Fetches IP applications all NA Verified by facilityID and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching all NA Verified application"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/getallnaverifiedapplications/{facilityid}/{selectedyear} [get]
func GetAllNAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllNAVerified "
		var startFunction string = " - start - QueryIPApplicationsByNAVerified "
		//facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryIPApplicationsByNAVerified(ctx, client, facilityID, id1)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching all NA Verified application",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)

}

// Get na verified record by Emp ID

// GetNAVerifiedDetailsByEmpId godoc
// @Summary Get NA Verified by Employee ID and year
// @Description Fetches IP applications NA Verified by Employee ID and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching NA verified application by employee ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/getallnaverified/{employeeid}/{examyear} [get]
func GetNAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetNAVerifiedDetailsByEmpId "
		var startFunction string = " - start - QueryIPApplicationsByNAVerifiedByEmpID "
		id := gctx.Param("employeeid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("selectedyear")
		circles, status, stgError, dataStatus, err := start.QueryIPApplicationsByNAVerifiedByEmpID(ctx, client, int64(empid), examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching NA verified application by employee ID",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// Get all NA Verified for NA ..

// GetAllNAVerifiedForNA godoc
// @Summary Get All NA Verified by Facility ID and year
// @Description Fetches IP applications  all NA Verified by facility ID and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching applications verified by NA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/getallnaverifiedapplicationsforna/{facilityid}/{selectedyear} [get]
func GetAllNAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetAllNAVerifiedForNA "
		var startFunction string = " - start - GetAllNAVerifiedForNA "
		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		facilityID1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryIPApplicationsByNAVerifiedForNA(ctx, client, facilityID, facilityID1)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching applications verified by NA",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// Get all CA verified for NA..

// GetAllCAVerifiedForNA godoc
// @Summary Get All CA Verified by Facility ID and year
// @Description Fetches IP applications  all CA Verified by facility ID and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching applications verified by CA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/getallcaverifiedapplicationsforna/{facilityid}/{selectedyear} [get]
func GetAllCAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetAllCAVerifiedForNA "
		var startFunction string = " - start - QueryIPApplicationsByCAVerifiedForNA "

		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		facilityID1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryIPApplicationsByCAVerifiedForNA(ctx, client, facilityID, facilityID1)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching application verified by CA for NA",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// Get All Pending With Candidate

// GetAllPendingWithCandidate godoc
// @Summary Get All Applications pending with Candidate by Facility ID and year
// @Description Fetches IP applications pending with Candidate by facility ID and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching application pending with candidate"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/getAllPendingWithCandidate/{facilityid}/{selectedyear} [get]
func GetAllPendingWithCandidate(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllPendingWithCandidate "
		var startFunction string = " - start - QueryIPApplicationsByPendingWithCandidate "
		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryIPApplicationsByPendingWithCandidate(ctx, client, facilityID, id1)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching application pending with candidate ",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetHallTicketWithExamCodeEmpID godoc
// @Summary Get All Exam HallTicket  by Examcode employeeid and year
// @Description Fetches All Exam HallTicket  by Examcode employeeid and year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param examname path int true "Exam Code"
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching applications generated with Hall ticket"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/hallticket/get/{examname}/{employeeid}/{examyear} [get]
func GetHallTicketWithExamCodeEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		var mainFunction string = " main - GetHallTicketWithExamCodeEmpID "
		var startFunction string = " - util - GetApplicationsWithHallTicket "
		//:examname/:employeeid/:examyear
		ec := gctx.Param("examname")
		ExamCode, _ := strconv.ParseInt(ec, 10, 32)
		EmployeeID, _ := strconv.ParseInt(gctx.Param("employeeid"), 10, 64)
		examYear := gctx.Param("examyear")

		ipApplication, psApplication, gdspaApplication, pmpaApplication, gdspmApplication, mtspmApplication, status, stgError, dataStatus, err := util.GetApplicationsWithHallTicket(client, int32(ExamCode), EmployeeID, string(examYear))
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		var data interface{}
		if ipApplication != nil {
			data = ipApplication
		} else if psApplication != nil {
			data = psApplication
		} else if gdspaApplication != nil {
			data = gdspaApplication
		} else if pmpaApplication != nil {
			data = pmpaApplication
		} else if gdspmApplication != nil {
			data = gdspmApplication
		} else if mtspmApplication != nil {
			data = mtspmApplication
		} else {
			data = nil
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching applications generated with Hall ticket",
			"data":       data,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

type GenerateHallticketRequestBody struct {
	ExamYear               string `json:"ExamYear"`
	ExamCode               int32  `json:"ExamCode"`
	NodalOfficerFacilityID string `json:"nodalOfficerFacilityID"`
}

// GenerateHallticketNumbers godoc
// @Summary Generate Hall Ticket Numbers
// @Description Generates hall ticket numbers for a given exam year, exam code, and nodal officer facility ID
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param requestBody body GenerateHallticketRequestBody true "Request Body"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Hall Ticket generation done successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /generate-hallticket-numbers [post]
func GenerateHallticketNumbers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		// Define a struct to represent the request body

		var mainFunction string = " main - GenerateHallticketNumbers "
		var startFunction string = " - start - GenerateHallticketNumberrIP "
		// Parse the request body into the struct
		var reqBody GenerateHallticketRequestBody
		if err := gctx.BindJSON(&reqBody); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		fmt.Println(reqBody.ExamYear)

		// Call GenerateHallticketNumberrIP function
		successMessage, status, stgError, dataStatus, err := start.GenerateHallticketNumberrIP(ctx, client, reqBody.ExamYear, reqBody.ExamCode, reqBody.NodalOfficerFacilityID)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		// Return the success message as the response
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Hall Ticket generation done successfully",
			"data":       successMessage,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// GetExamApplicatonsPreferenenceCityWiseStats godoc
// @Summary Get Exam Applications Preference City Wise Stats
// @Description Fetch the application count based on Exam city
// @Tags IP Exam Application
// @Param examYear path string true "Exam Year"
// @Param examcode path string true "Exam Code"
// @Param cityid path string true "City ID"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Fetching application count based on Exam city"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/getExamApplicationsByCityPref/{examYear}/{examcode}/{cityid} [get]
func GetExamApplicatonsPreferenenceCityWiseStats(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		//:examYear/:examcode/:cityid
		ExamYear := gctx.Param("examYear")
		Examcode := gctx.Param("examcode")
		Cityid := gctx.Param("cityid")
		var mainFunction string = " main - GetExamApplicatonsPreferenenceCityWiseStats "
		var startFunction string = " - start - GetExamApplicatonsPreferenenceCityWiseStats "

		// Convert strings to int64

		var statistics []start.ExamStats
		var err error
		statistics, status, stgError, dataStatus, err := start.GetExamApplicatonsPreferenenceCityWiseStats(ctx, client, ExamYear, Examcode, Cityid)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching application count based on Exam city ",
			"data":       statistics,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

func GetExamApplicationsByCenterIP(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		//:examYear/:examcode/:centerid/:startno/:entno
		ExamYear := gctx.Param("examYear")
		Examcode := gctx.Param("examcode")
		CenterId := gctx.Param("centerid")
		StartNo := gctx.Param("startno")
		EndNo := gctx.Param("endno")
		var mainFunction string = " main - GetExamApplicationsByCenterIP "
		var startFunction string = " - start - SubGetExamApplicationsByCenterIP "

		// Convert strings to int64

		//		var statistics []start.ExamStats
		var err error
		applications, status, stgError, dataStatus, err := start.SubGetExamApplicationsByCenterIP(ctx, client, ExamYear, Examcode, CenterId, StartNo, EndNo)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching application based on exam center ",
			"data":       applications,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

type requestDataArray []struct {
	ExamCityID                  int    `json:"examCityID"`
	ExamCenterID                int    `json:"examCenterID"`
	ControllingOfficeFacilityID string `json:"controllingOfficeFacilityID"`
	SeatsToAllot                int    `json:"seatsToAllot"`
}

// UpdateCenterCodeForApplications godoc
// @Summary Update Center Code For Applications
// @Description Update the center code for multiple applications
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param requestData body requestDataArray true "Request Data Array"
// @Success 200 {object} start.UpdateCenterCodeResponse "Total Number of Applications Updated with center code Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/updateCenterCodeForApplications [put]
// func UpdateCenterCodeForApplicationsOld(client *ent.Client) gin.HandlerFunc {
// 	return func(gctx *gin.Context) {

// 		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
// 		defer cancel()
// 		var mainFunction string = " main - UpdateCenterCodeForApplications "
// 		var startFunction string = " - start - UpdateCenterCodeForApplications "
// 		var reqData requestDataArray
// 		if err := gctx.ShouldBindJSON(&reqData); err != nil {
// 			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
// 			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
// 			return
// 		}

// 		var totalUpdatedCount, count int
// 		var status int32
// 		var dataStatus bool
// 		var err error
// 		var stgError string
// 		for _, requestData := range reqData {
// 			count, status, stgError, dataStatus, err = start.UpdateCenterCodeForApplications(ctx, client, requestData.ControllingOfficeFacilityID, int32(requestData.ExamCenterID), int32(requestData.SeatsToAllot), int32(requestData.ExamCityID))
// 			if err != nil {
// 				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
// 			}
// 			totalUpdatedCount += count
// 		}
// 		gctx.JSON(http.StatusOK, gin.H{
// 			"success":    true,
// 			"message":    fmt.Sprintf("Total Number of Applications %d Updated with center code Successfully", totalUpdatedCount),
// 			"data":       struct{}{},
// 			"dataexists": dataStatus,
// 		})
// 	}

// }
// func UpdateCenterCodeForApplications(client *ent.Client) gin.HandlerFunc {
// 	return func(gctx *gin.Context) {

// 		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
// 		defer cancel()
// 		var mainFunction string = " main - UpdateCenterCodeForApplications "
// 		var startFunction string = " - start - UpdateCenterCodeForApplications "
// 		var reqData requestDataArray
// 		if err := gctx.ShouldBindJSON(&reqData); err != nil {
// 			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
// 			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
// 			return
// 		}
// 		type ApplicationData struct {
// 			EmployeeID         int64  `json:"employee_id"`
// 			ApplicationNumber  string `json:"application_number"`
// 			HallTicketNumber   string `json:"hall_ticket_number"`
// 		}
// 		var totalUpdatedCount, count int
// 		var status int32

// 		var smsStatus string
// 		var emailsentstatus string
// 		var dataStatus bool
// 		var updatedApplications []ApplicationData

// 		 var err error
// 		 var stgError string
// 		for _, requestData := range reqData {
// 			count, updatedApps, status, stgError, dataStatus, err = start.UpdateCenterCodeForApplications(ctx, client, requestData.ControllingOfficeFacilityID, int32(requestData.ExamCenterID), int32(requestData.SeatsToAllot), int32(requestData.ExamCityID))
// 			if err != nil {
// 				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
// 				return // Exit the handler if an error occurs
// 			}
// 			totalUpdatedCount += count
// 			updatedApplications = append(updatedApplications, updatedApps...)

// 			var employeeIDInt = int(EmployeeID)
// 			EmployeeIDStr := strconv.Itoa(employeeIDInt)
// 			smsStatus = sms.SendSmsNew(ctx, client, EmployeeIDStr, 12, ApplicationNumber, hallticketnumber)

// 	//  <Application Number>, <Application Submitted date and time> , <Exam Short Name>,  <Nodal Officer Officer Name> , <Hall ticket Number> , <Hall ticket generated date and time>, <URL> , <Controlling authority Office Name>

// 			emailsentstatus = mail.SendEMailNew(ctx, client, EmployeeIDStr, 12, " LDCE PA/SA to IP ", ApplicationNumber, hallticketnumber)
// 		}

// 		gctx.JSON(http.StatusOK, gin.H{
// 			"success":     true,
// 			"message":     fmt.Sprintf("Total Number of Applications %d Updated with center code Successfully", totalUpdatedCount),
// 			"data":        struct{}{},
// 			"dataexists":  dataStatus,
// 			"SMSStatus":   smsStatus,
// 			"EmailStatus": emailsentstatus,
// 		})
// 	}

// }
func UpdateCenterCodeForApplications(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - UpdateCenterCodeForApplications "
		var startFunction string = " - start - UpdateCenterCodeForApplications "
		var reqData requestDataArray
		if err := gctx.ShouldBindJSON(&reqData); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		var totalUpdatedCount int
		var updatedApplications []*ent.Exam_Applications_IP

		var smsStatus, emailSentStatus string
		var dataStatus bool
		for _, requestData := range reqData {
			count, updatedApps, status, stgError, currentDataStatus, err := start.UpdateCenterCodeForApplications(ctx, client, requestData.ControllingOfficeFacilityID, int32(requestData.ExamCenterID), int32(requestData.SeatsToAllot), int32(requestData.ExamCityID))

			if err != nil {
				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
				return
			}

			totalUpdatedCount += count
			updatedApplications = append(updatedApplications, updatedApps...)
			dataStatus = currentDataStatus
			for _, app := range updatedApps {
				employeeIDStr := strconv.FormatInt(app.EmployeeID, 10)
				smsStatus = sms.SendSmsNew(ctx, client, employeeIDStr, 12, app.ApplicationNumber, app.HallTicketNumber, "LDCE PA/SA to IP")
				var submittedDate string = app.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
				var hallticketDate string = app.HallTicketGeneratedDate.Format("02-01-2006 15:04:05")
				emailSentStatus = mail.SendEMailNew(ctx, client, employeeIDStr, 12, "LDCE PA/SA to IP", app.ApplicationNumber, app.ControllingOfficeName, submittedDate, app.NodalOfficeName, app.HallTicketNumber, hallticketDate)
			}
		}

		gctx.JSON(http.StatusOK, gin.H{
			"success":             true,
			"message":             fmt.Sprintf("Total Number of Applications %d Updated with center code Successfully", totalUpdatedCount),
			"updatedApplications": updatedApplications,
			"dataexists":          dataStatus,
			"data":                struct{}{},
			"SMSStatus":           smsStatus,
			"EmailStatus":         emailSentStatus,
		})
	}
}

type UpdateExamCentersRequest struct {
	Newappls []*ent.Exam_Applications_IP `json:"newappls"`
}

// UpdateExamCentersInIPApplsreturnstring godoc
// @Summary Update Exam Centers in IP Applications
// @Description Update the exam centers in IP applications with the provided data
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param requestBody body ca_reg.UpdateExamCentersInIP true "Request Body"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Updating Exam centers for IP Applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/center/updatecenters/IP [put]
func UpdateExamCentersInIPApplsreturnstring(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateExamCentersInIPApplsreturnstring "
		var startFunction string = " - start - UpdateExamCentresIPExamsreturnarray "

		// Bind the JSON data to the newappls variable
		var req UpdateExamCentersRequest
		if err := gctx.ShouldBindJSON(&req); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		// Call the UpdateExamCentresIPExams function to update the exam centers
		examCenters, status, stgError, dataStatus, err := start.UpdateExamCentresIPExamsreturnarray(ctx, client, req.Newappls)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		// Display the status
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Updating Exam centers for IP Applications",
			"data":       examCenters,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

// GenerateHallticketNumberscenter godoc
// @Summary Generate Hall Ticket Numbers
// @Description Generates hall ticket numbers for a given exam year
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetAllCAPendingVerificationsResponse "Generation of Hall ticket number"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/halltickets/{examyear} [put]
func GenerateHallticketNumberscenter(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GenerateHallticketNumberscenter "
		var startFunction string = " - start - GenerateHallticketNumberIP "
		examYear := gctx.Param("year")

		// Call the GenerateHallticketNumberIP function to generate hall ticket numbers and get the success message
		successMessage, status, stgError, dataStatus, err := start.GenerateHallticketNumberIP(ctx, client, examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		// Return the success message as the response
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Generation of Hall ticket number",
			"data":       successMessage,
			"dataexists": dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// UpdateNodalRecommendationsIPByEmpID godoc
// @Summary Update Nodal Recommendations for IP by Employee ID
// @Description Update nodal recommendations for IP applications by employee ID
// @Tags IP Exam Application
// @Accept json
// @Produce json
// @Param requestBody body ca_reg.NAVerifyApplicationIp true "Request Body"
// @Success 200 {object} start.VerifyIPApplicationResponse "Nodal officer successfully verified this IP application"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/ipexams/noverify [put]
func UpdateNodalRecommendationsIPByEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateNodalRecommendationsIPByEmpID "
		var startFunction string = " - start - UpdateNodalRecommendationsByEmpID "
		var newAppln ca_reg.NAVerifyApplicationIp

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newAppln.Edges.LogData[0]
		fmt.Println(logdata)

		application, status, stgError, dataStatus, err := start.UpdateNodalRecommendationsByEmpID(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		var emailsentstatus string = ""
		if application.ApplicationStatus == "VerifiedByNA" {
			//examshortname ,appno ,ca,submittedDate ,recommended, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE PA/SA to IP ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.NARemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE PA/SA to IP ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.NARemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByNA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Nodal officer successfully verified this IP application ",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}
func CheckIPExamcenterhall(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//:examyear/:examcode/:nafacilityid/:cofacilityid/:hallname
		examyear := gctx.Param("examyear")
		id2 := gctx.Param("examcode")

		naFacilityId := gctx.Param("nafacilityid")
		coFacilityId := gctx.Param("cofacilityid")
		examHall := gctx.Param("hallname")

		var mainFunction string = " main - CheckIPExamcenterhall "
		var startFunction string = " - start - SubCheckIPExamcenterhall "

		examcode, _ := strconv.ParseInt(id2, 10, 32)

		records, status, stgError, dataStatus, err := start.SubCheckIPExamcenterhall(client, examyear, int32(examcode), naFacilityId, coFacilityId, examHall)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching ExamcenterHall by CityId",
			"data":       records,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}
func GetIPExamcenterhallBycityid(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		id := gctx.Param("cityid")
		examyear := gctx.Param("examyear")
		id2 := gctx.Param("examcode")
		id3 := gctx.Param("centerid")

		var mainFunction string = " main - GetIPExamcenterhallBycityid "
		var startFunction string = " - start - GetExamCenterHallbyCityID "

		cityid, _ := strconv.ParseInt(id, 10, 32)
		examcode, _ := strconv.ParseInt(id2, 10, 32)
		centerid, _ := strconv.ParseInt(id3, 10, 32)

		records, status, stgError, dataStatus, err := start.GetExamCenterHallbyCityID(client, int32(cityid), examyear, int32(examcode), int32(centerid))
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching ExamcenterHall by CityId",
			"data":       records,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}
func GetIPCandiadatesExamcenterhallBycityid(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		id := gctx.Param("cityid")
		examyear := gctx.Param("examyear")
		id2 := gctx.Param("examcode")
		id3 := gctx.Param("centerid")
		hallname := gctx.Param("hallname")

		var mainFunction string = " main - GetIPCandiadatesExamcenterhallBycityid "
		var startFunction string = " - start - GetCandidateExamCenterHallbyCityID "

		cityid, _ := strconv.ParseInt(id, 10, 32)
		examcode, _ := strconv.ParseInt(id2, 10, 32)
		centerid, _ := strconv.ParseInt(id3, 10, 32)

		records, status, stgError, dataStatus, err := start.GetCandidateExamCenterHallbyCityID(client, int32(cityid), examyear, int32(examcode), int32(centerid), hallname)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching Candidates in ExamcenterHall by CityId",
			"data":       records,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}
