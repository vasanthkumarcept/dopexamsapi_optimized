package handlers

import (
	"context"
	"fmt"
	"log"
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

// CreateNewPSApplications godoc
// @Summary Create new PS Applications
// @Description Create a new PS Group B exam application
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param data body ca_reg.ApplicationGroupB true "Application data"
// @Success 200 {object} start.PsGroupBResponse "PS Group B Exam Application submitted successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/applications/submit [post]

func CreateNewPSApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - CreateNewPSApplications "
		var startFunction string = " - start - CreatePSApplications "
		var newIPAppplns ca_reg.ApplicationGroupB
		if err := gctx.ShouldBindJSON(&newIPAppplns); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newIPAppplns.Edges.LogData[0]

		newIPApppln, status, stgError, dataStatus, err := start.CreatePSApplications(client, &newIPAppplns)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var Smsuserid string = fmt.Sprint(newIPApppln.EmployeeID)
		var submittedDate string = newIPApppln.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		//examshortname ,appno ,ca submittedDate
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 3, " LDCE LSG/IP to PSGRB ", newIPApppln.ApplicationNumber, submittedDate)

		//----for sending SMS
		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE LSG/IP to PSGRB", newIPApppln.ApplicationNumber, "Submitted")

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "PS Group B Application submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        newIPApppln,
			"dataexists":  dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// to be checked
func CreateExamCenterHallPS(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		//defer cancel()
		var mainFunction string = " main - CreateExamCenterHallPS "
		var startFunction string = " - start SubCreateExamCenterHallPS "
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

// VerifyPSApplication godoc
// @Summary Verify PS Application By CA
// @Description Verify a PS Group B exam application
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param data body ca_reg.VerifyApplicationGroupB true "Application data"
// @Success 200 {object} start.PsGroupBResponse "PS Group B exam application verified successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/applications/verify [put]
func VerifyPSApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - VerifyPSApplication "
		var startFunction string = " - start - UpdateApplicationRemarksPS "
		var newAppln ca_reg.VerifyApplicationGroupB
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

		application, status, stgError, dataStatus, err := start.UpdateApplicationRemarksPS(client, &newAppln, nonQualifyService)
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
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE LSG/IP to PSGRB ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.CAGeneralRemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE LSG/IP to PSGRB ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.CAGeneralRemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByCA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE LSG/IP to PSGRB", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE LSG/IP to PSGRB", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "PS Group B exam application verified successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

// to be checked
type CandidateInfoPS struct {
	EmployeeID  string `json:"employeeID"`
	EmailStatus string `json:"emailStatus"`
	SMSStatus   string `json:"smsStatus"`
}

// to be checked
func EmailSmsTriggeringPenidngWithCandidatePS(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//:nodalofficeid:examyear
		examYear := gctx.Param("examyear")
		nodalOfficeId := gctx.Param("nodalofficeid")
		var mainFunction string = " main - EmailSmsTriggeringPenidngWithCandidatePS "
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
		var candidateData []CandidateInfoPS
		for _, app := range application {
			var submittedDate string = app.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
			var Smsuserid = fmt.Sprint(app.EmployeeID)
			emailSentStatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE PA/SA to IP ", app.ApplicationNumber, app.ControllingOfficeName, submittedDate, app.CAGeneralRemarks)
			smsStatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE PA/SA to IP", app.ApplicationNumber, "Returned for correction")
			candidateData = append(candidateData, CandidateInfoPS{
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

// to be checked
func ResetPSApplicationNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		/* ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		*/
		var mainFunction string = " main - ResetPSApplicationNA "
		var startFunction string = " - start - SubResetPSApplicationNA "
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

// ResubmitPSApplication godoc
// @Summary Resubmit PS Application
// @Description Resubmit a PS Group B exam application
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param data body ca_reg.ReApplicationGroupB true "Application data"
// @Success 200 {object} start.PsGroupBResponse "PS Group B Exam application re-submitted successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/applications/resubmit [put]
func ResubmitPSApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - ResubmitPSApplication "
		var startFunction string = " - start - ResubmitApplicationRemarksPS "
		var newAppln ca_reg.ReApplicationGroupB
		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newAppln.Edges.LogData[0]
		application, status, stgError, dataStatus, err := start.ResubmitApplicationRemarksPSs(client, &newAppln)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		//examshortname ,	appno, ca, submittedDate
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 7, " LDCE LSG/IP to PSGRB ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE LSG/IP to PSGRB", application.ApplicationNumber, "Re-submitted")

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "PS Group B Exam application re-submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// Get IP Applications with EmpID - QueryIPExamApplicationsByEmpID

// GetPSApplicationsByEmpId godoc
// @Summary Get PS Applications by Employee ID
// @Description Fetch PS Group B exam applications by employee ID and exam year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetApplicationsResponse "Application details fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/applications/getbyempid/{employeeid}/{examyear} [get]
func GetPSApplicationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//employeeid/:examyear
		id := gctx.Param("employeeid")
		examYear := gctx.Param("examyear")
		//var examID int32
		employeeID, _ := strconv.ParseInt(id, 10, 64)

		var mainFunction string = " main - GetPSApplicationsByEmpId "
		var startFunction string = " - start - QueryPSExamApplicationsByEmpID"

		ipapplns, status, stgError, dataStatus, err := start.QueryPSExamApplicationsByEmpID(ctx, client, int64(employeeID), string(examYear))
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

// GetAllPSCAPendingVerifications godoc
// @Summary Get all PS Group A pending verifications by CA
// @Description Fetches all pending PS Group A applications pending verification by CA for a specific facility and selected year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetched pending application with CA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/getallcapendingapplications/{facilityid}/{selectedyear} [get]
func GetAllPSCAPendingVerifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetAllPSCAPendingVerifications "
		var startFunction string = " - start - QueryPSApplicationsByCAVerificationsPending "
		//:facilityid/:examyear
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, status, stgError, dataStatus, err := start.QueryPSApplicationsByCAVerificationsPending(ctx, client, facilityID, examYear)
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

// to be checked
func GetAllPendingApplicationsOnDeputationPS(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetAllPendingApplicationsOnDeputationPS "
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

// GetPSCAVerifiedDetailsByEmpId godoc
// @Summary Get PS Group A verified applications by Employee ID
// @Description Fetches PS Group A applications verified by CA for a specific Employee ID and exam year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetched data based on employeeid with Verified status"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/getallcaverified/{employeeid}/{selectedyear} [get]
func GetPSCAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetPSCAVerifiedDetailsByEmpId "
		var startFunction string = " - start - QueryPSApplicationsByCAVerifiedByEmpID "
		//:employeeid/:examyear
		id := gctx.Param("employeeid")
		examYear := gctx.Param("examyear")

		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, status, stgError, dataStatus, err := start.QueryPSApplicationsByCAVerifiedByEmpID(ctx, client, int64(empid), examYear)
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

// Return Previous Remarks

// GetPSCAPendingOldRemarksByEmpId godoc
// @Summary Get PS Group A pending old remarks by Employee ID
// @Description Fetches PS Group A applications pending with candidate along with old remarks by employee ID and exam year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching applicatins pending with candiate with remarks"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/caprevremarks/{employeeid}/{examyear} [get]
func GetPSCAPendingOldRemarksByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetPSCAPendingOldRemarksByEmpId "
		var startFunction string = " - start - GetOldPSCAApplicationRemarksByEmployeeID "

		//:employeeid/:examyear
		id := gctx.Param("employeeid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("examyear")
		circles, status, stgError, dataStatus, err := start.GetOldPSCAApplicationRemarksByEmployeeID(ctx, client, int64(empid), examYear)
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

// GetPSExamRecommendationsByEmpId godoc
// @Summary Get PS exam recommendations by Employee ID
// @Description Fetches PS exam recommendations based on employee ID
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Success 200 {object} start.GetApplicationsResponse "Fetching recommendations by employee ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/recommendations/{employeeid} [get]
func GetPSExamRecommendationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		id := gctx.Param("employeeid")
		var mainFunction string = " main - GetPSExamRecommendationsByEmpId "
		var startFunction string = " - start - GetPSRecommendationsByEmpID "
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		records, status, stgError, dataStatus, err := start.GetPSRecommendationsByEmpID(client, empid)
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

// GetPSAllNAVerified godoc
// @Summary Get PS applications verified by NA
// @Description Fetches PS applications verified by NA based on facility ID and selected year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching all NA Verified application"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/getallnaverifiedapplications/{facilityid}/{selectedyear} [get]
func GetPSAllNAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetPSAllNAVerified "
		var startFunction string = " - start - QueryPSApplicationsByNAVerified "
		//:facilityid/:examyear
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, status, stgError, dataStatus, err := start.QueryPSApplicationsByNAVerified(ctx, client, facilityID, examYear)
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

// to be checked <examyear added)
// Get na verified record by Emp ID

// GetPSNAVerifiedDetailsByEmpId godoc
// @Summary Get PS applications verified by NA for a specific employee ID
// @Description Fetches PS applications verified by NA based on employee ID
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching NA verified application by employee ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/getallnaverified/{employeeid} [get]
func GetPSNAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetPSNAVerifiedDetailsByEmpId "
		var startFunction string = " - start - QueryPSApplicationsByNAVerifiedByEmpID "
		id := gctx.Param("employeeid")
		//var examID int32
		employeeID, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("examyear")
		circles, status, stgError, dataStatus, err := start.QueryPSApplicationsByNAVerifiedByEmpID(ctx, client, int64(employeeID), examYear)
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

// Get all CA verified for NA..

// GetPSAllCAVerifiedForNA godoc
// @Summary Get PS applications verified by CA for NA
// @Description Fetches PS applications verified by CA for a specific facility ID and selected year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching application verified by CA for NA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/getallcaverifiedapplicationsforna/{facilityid}/{selectedyear} [get]
func GetPSAllCAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetPSAllCAVerifiedForNA "
		var startFunction string = " - start - QueryPSApplicationsByCAVerifiedForNA "
		//:facilityid/:examyear
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, status, stgError, dataStatus, err := start.QueryPSApplicationsByCAVerifiedForNA(ctx, client, facilityID, examYear)
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

// Get all NA Verified for NA ..

// GetPSAllNAVerifiedForNA godoc
// @Summary Get PS applications verified by NA for a specific facility and year
// @Description Fetches PS applications verified by NA based on facility ID and selected year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Updating Exam centers for IP Applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/getallnaverifiedapplicationsforna/{facilityid}/{selectedyear} [get]
func GetPSAllNAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetPSAllNAVerifiedForNA "
		var startFunction string = " - start - QueryPSApplicationsByNAVerifiedForNA "

		//:facilityid/:examyear
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryPSApplicationsByNAVerifiedForNA(ctx, client, facilityID, examYear)
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

// Get All Pending With Candidate

// GetAllPSPendingWithCandidate godoc
// @Summary Get PS applications pending with candidate for a specific facility and year
// @Description Fetches PS applications that are pending with candidates based on facility ID and selected year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching application pending with candidate "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/getAllPSPendingWithCandidate/{facilityid}/{selectedyear} [get]
func GetAllPSPendingWithCandidate(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllPSPendingWithCandidate "
		var startFunction string = " - start - QueryPSApplicationsByPendingWithCandidate "
		//:facilityid/:examyear
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, status, stgError, dataStatus, err := start.QueryPSApplicationsByPendingWithCandidate(ctx, client, facilityID, examYear)
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

// GenerateHallticketNumbersPs godoc
// @Summary Get PS applications pending with candidate for a specific facility and year
// @Description Fetches PS applications that are pending with candidates based on facility ID and selected year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param data body requestBody true "Application data"
// @Success 200 {object} start.GetApplicationsResponse "Hall Ticket generation done successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/PSexams/Halltickets [put]
func GenerateHallticketNumbersPs(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		// Define a struct to represent the request body

		var mainFunction string = " main - GenerateHallticketNumbersPs "
		var startFunction string = " - start - GenerateHallticketNumberrPS "
		// Parse the request body into the struct
		var reqBody GenerateHallticketRequestBodyPS
		if err := gctx.BindJSON(&reqBody); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		fmt.Println(reqBody.ExamYear)

		// Call GenerateHallticketNumberrIP function
		successMessage, status, stgError, dataStatus, err := start.GenerateHallticketNumberrPS(ctx, client, reqBody.ExamYear, reqBody.ExamCode, reqBody.NodalOfficerFacilityID)
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

// GetExamApplicatonsPreferenenceCityWiseStatsPS godoc
// @Summary Get Exam Applications Preference City-Wise Stats for PS
// @Description Fetches the application count based on exam city for a specific exam year, exam code, and city ID
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param examYear path string true "Exam Year"
// @Param examcode path string true "Exam Code"
// @Param cityid path string true "City ID"
// @Success 200 {object} start.GetApplicationsResponse "Fetching application count based on Exam city"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/getExamApplicationsByCityPrefPS/{examYear}/{examcode}/{cityid} [get]
func GetExamApplicatonsPreferenenceCityWiseStatsPS(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		//:examYear/:examcode/:cityid
		ExamYear := gctx.Param("examYear")
		Examcode := gctx.Param("examcode")
		Cityid := gctx.Param("cityid")
		var mainFunction string = " main - GetExamApplicatonsPreferenenceCityWiseStatsPS "
		var startFunction string = " - start - GetExamApplicatonsPreferenenceCityWiseStatsPS "

		// Convert strings to int64

		var statistics []start.ExamStatsPS
		var err error
		statistics, status, stgError, dataStatus, err := start.GetExamApplicatonsPreferenenceCityWiseStatsPS(ctx, client, ExamYear, Examcode, Cityid)

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

// UpdateExamCentersInPSApplsreturnstring godoc
// @Summary Update exam centers in PS applications
// @Description Update exam centers in PS applications with new data
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param request body []ApplicationGroupBCenter true "JSON array of PS exam applications"
// @Success 200 {object} start.GetApplicationsResponse "Updating Exam centers for IP Applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/center/updatecenters [post]
func UpdateExamCentersInPSApplsreturnstring(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateExamCentersInPSApplsreturnstring "
		var startFunction string = " - start - UpdateExamCentresPSExamsreturnarray "

		// Bind the JSON data to the newappls variable
		var req UpdateExamCentersRequestPS
		if err := gctx.ShouldBindJSON(&req); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		// Call the UpdateExamCentresIPExams function to update the exam centers
		examCenters, status, stgError, dataStatus, err := start.UpdateExamCentresPSExamsreturnarray(ctx, client, req.Newappls)
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

// UpdateCenterCodeForApplicationsPS godoc
// @Summary Update Center Code for Applications
// @Description Updates the center code for PS applications based on provided data
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param requestDataArray body []requestDataArray true "Request Data Array"
// @Success 200 {object} start.GetApplicationsResponse "All Applications Updated with center code Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/updateCenterCodeForApplicationsPS [put]
func UpdateCenterCodeForApplicationsPS(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - UpdateCenterCodeForApplicationsPS "
		var startFunction string = " - start - UpdateCenterCodeForApplicationsPS "
		var reqData requestDataArrayPS
		if err := gctx.ShouldBindJSON(&reqData); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		var totalUpdatedCount int
		var updatedApplications []*ent.Exam_Applications_PS

		var smsStatus, emailSentStatus string
		var dataStatus bool
		for _, requestData := range reqData {
			count, updatedApps, status, stgError, currentDataStatus, err := start.UpdateCenterCodeForApplicationsPS(ctx, client, requestData.ControllingOfficeFacilityID, int32(requestData.ExamCenterID), int32(requestData.SeatsToAllot), int32(requestData.ExamCityID))

			if err != nil {
				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
				return
			}

			totalUpdatedCount += count
			updatedApplications = append(updatedApplications, updatedApps...)
			dataStatus = currentDataStatus
			for _, app := range updatedApps {
				employeeIDStr := strconv.FormatInt(app.EmployeeID, 10)
				smsStatus = sms.SendSmsNew(ctx, client, employeeIDStr, 12, app.ApplicationNumber, app.HallTicketNumber, "LDCE LSG/IP to PSGRB")
				var submittedDate string = app.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
				var hallticketDate string = app.HallTicketGeneratedDate.Format("02-01-2006 15:04:05")
				emailSentStatus = mail.SendEMailNew(ctx, client, employeeIDStr, 12, "LDCE LSG/IP to PSGRB", app.ApplicationNumber, app.ControllingOfficeName, submittedDate, app.NodalOfficeName, app.HallTicketNumber, hallticketDate)
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

// to be checked
func GetExamApplicationsByCenterPS(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		//:examYear/:examcode/:centerid/:startno/:entno
		ExamYear := gctx.Param("examYear")
		Examcode := gctx.Param("examcode")
		CenterId := gctx.Param("centerid")
		StartNo := gctx.Param("startno")
		EndNo := gctx.Param("endno")
		var mainFunction string = " main - GetExamApplicationsByCenterPS "
		var startFunction string = " - start - SubGetExamApplicationsByCenterPS "

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

type requestDataArrayPS []struct {
	ExamCityID                  int    `json:"examCityID"`
	ExamCenterID                int    `json:"examCenterID"`
	ControllingOfficeFacilityID string `json:"controllingOfficeFacilityID"`
	SeatsToAllot                int    `json:"seatsToAllot"`
}

type UpdateExamCentersRequestPS struct {
	Newappls []*ent.Exam_Applications_PS `json:"newappls"`
}

// UpdateNodalRecommendationsPSByEmpID godoc
// @Summary Update Nodal Recommendations for PS by Employee ID
// @Description Updates nodal recommendations for PS GroupB applications based on Employee ID
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param newAppln body ca_reg.NAVerifyApplicationGroupB true "Application GroupB data"
// @Success 200 {object} start.PsGroupBResponse "Nodal officer successfully verified this PS GroupB application "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/noverify [put]
func UpdateNodalRecommendationsPSByEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateNodalRecommendationsPSByEmpID "
		var startFunction string = " - start - UpdateNodalRecommendationsPSByEmpID"
		var newAppln ca_reg.NAVerifyApplicationGroupB

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newAppln.Edges.LogData[0]
		fmt.Println(logdata)

		application, status, stgError, dataStatus, err := start.UpdateNodalRecommendationsPSByEmpID(client, &newAppln)
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
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE LSG/IP to PSGRB ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.NARemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE LSG/IP to PSGRB ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.NARemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByNA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE LSG/IP to PSGRB", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE LSG/IP to PSGRB", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Nodal officer successfully verified this PS GroupB application ",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// Get All CA Verified records
// GetAllPSCAVerified godoc
// @Summary Get PS Group B verified applications
// @Description Fetches PS Group B applications verified by CA for a specific facility ID and selected year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetched CA Verified applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/getallcaverifiedapplications/{facilityid}/{selectedyear} [get]
func GetAllPSCAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		//:facilityid/:examyear

		var mainFunction string = " main - GetAllPSCAVerified "
		var startFunction string = " - start - QueryPSApplicationsByCAVerified "
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, status, stgError, dataStatus, err := start.QueryPSApplicationsByCAVerified(ctx, client, facilityID, examYear)
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

// to be checked
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
func GenerateHallticketNumberscenterPS(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GenerateHallticketNumberscenter "
		var startFunction string = " - start - GenerateHallticketNumberPS "
		examYear := gctx.Param("year")

		// Call the GenerateHallticketNumberPS function to generate hall ticket numbers and get the success message
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

// to be checked
func GetEmployeeDetailsByEmpIdPS(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("employeeid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		var mainFunction string = " main - GetEmployeeDetailsByEmpIdPS "
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

// GetPSCAPendingDetailsByEmpId godoc
// @Summary Get PS Group A pending details by Employee ID
// @Description Fetches PS Group A applications pending verification by CA for a specific employee ID and exam year
// @Tags PS GroupB Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching application pending with CA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/psexams/getallcapending/{employeeid}/{selectedyear} [get]
func GetPSCAPendingDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//:employeeid/:examyear
		var mainFunction string = " main - GetPSCAPendingDetailsByEmpId "
		var startFunction string = " - start - QueryPSApplicationsByCAPendingByEmpID "
		id := gctx.Param("employeeid")

		empid, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("examyear")
		circles, status, stgError, dataStatus, err := start.QueryPSApplicationsByCAPendingByEmpID(ctx, client, int64(empid), examYear)
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

type GenerateHallticketRequestBodyPS struct {
	ExamYear               string `json:"ExamYear"`
	ExamCode               int32  `json:"ExamCode"`
	NodalOfficerFacilityID string `json:"nodalOfficerFacilityID"`
}

func GetPSHallticketNumberscenter(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//ec := gctx.Param("id")
		//CenterCode, _ := strconv.ParseInt(ec, 10, 32)
		// Call the GenerateHallticketNumberIP function to generate hall ticket numbers and get the success message
		successMessage, err := start.GenerateHallticketNumberPSwithCenterCode(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println(err.Error())
			return
		}

		// Log the success message
		log.Println(successMessage)

		// Return the success message as the response
		gctx.String(http.StatusOK, successMessage)
	}

	return gin.HandlerFunc(fn)
}

type ApplicationGroupBCenter struct {
	CandidateRemarks                string                        `json:"CandidateRemarks"`
	CategoryCode                    string                        `json:"CategoryCode"`
	CategoryDescription             string                        `json:"CategoryDescription"`
	CenterFacilityId                string                        `json:"CenterFacilityId"`
	CenterId                        int32                         `json:"CenterId"`
	CentrePreference                string                        `json:"CentrePreference"`
	ClaimingQualifyingService       string                        `json:"ClaimingQualifyingService"`
	ControllingOfficeFacilityID     string                        `json:"ControllingOfficeFacilityID"`
	ControllingOfficeName           string                        `json:"ControllingOfficeName"`
	DCCS                            string                        `json:"DCCS"`
	DOB                             string                        `json:"DOB"`
	DeputationControllingOfficeID   string                        `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string                        `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string                        `json:"DeputationOfficeFacilityID"`
	DeputationOfficeName            string                        `json:"DeputationOfficeName"`
	DeputationOfficePincode         string                        `json:"DeputationOfficePincode"`
	DeputationOfficeUniqueId        string                        `json:"DeputationOfficeUniqueId"`
	DeputationType                  string                        `json:"DeputationType"`
	DesignationID                   string                        `json:"DesignationID"`
	DisabilityPercentage            int32                         `json:"DisabilityPercentage"`
	DisabilityTypeCode              string                        `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string                        `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string                        `json:"DisabilityTypeID"`
	Edges                           ca_reg.EdgesPSApplicationData `json:"edges"`
	EducationCode                   string                        `json:"EducationCode"`
	EducationDescription            string                        `json:"EducationDescription"`
	EmailID                         string                        `json:"EmailID"`
	EmployeeID                      int64                         `json:"EmployeeID"`
	EmployeeName                    string                        `json:"EmployeeName"`
	EmployeePost                    string                        `json:"EmployeePost"`
	EntryPostCode                   string                        `json:"EntryPostCode"`
	EntryPostDescription            string                        `json:"EntryPostDescription"`
	ExamCode                        int32                         `json:"ExamCode"`
	ExamName                        string                        `json:"ExamName"`
	ExamShortName                   string                        `json:"ExamShortName"`
	ExamYear                        string                        `json:"ExamYear"`
	FacilityName                    string                        `json:"FacilityName"`
	FacilityUniqueID                string                        `json:"FacilityUniqueID"`
	FeederPostCode                  string                        `json:"FeederPostCode"`
	FeederPostDescription           string                        `json:"FeederPostDescription"`
	FeederPostJoiningDate           string                        `json:"FeederPostJoiningDate"`
	Gender                          string                        `json:"Gender"`
	LienControllingOfficeID         string                        `json:"LienControllingOfficeID"`
	LienControllingOfficeName       string                        `json:"LienControllingOfficeName"`
	MobileNumber                    string                        `json:"MobileNumber"`
	NodalOfficeFacilityID           string                        `json:"NodalOfficeFacilityID"`
	NodalOfficeName                 string                        `json:"NodalOfficeName"`
	Photo                           string                        `json:"Photo"`
	PhotoPath                       string                        `json:"PhotoPath"`
	PresentDesignation              string                        `json:"PresentDesignation"`
	PresentPostCode                 string                        `json:"PresentPostCode"`
	PresentPostDescription          string                        `json:"PresentPostDescription"`
	ReportingOfficeFacilityID       string                        `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string                        `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{}                `json:"ServiceLength"`
	Signature                       string                        `json:"Signature"`
	SignaturePath                   string                        `json:"SignaturePath"`
	TempHallTicket                  string                        `json:"TempHallTicket"`
	UserID                          int32                         `json:"UserID"`
	WorkingOfficeCircleFacilityID   string                        `json:"WorkingOfficeCircleFacilityID"`
	WorkingOfficeCircleName         string                        `json:"WorkingOfficeCircleName"`
	WorkingOfficeDivisionFacilityID string                        `json:"WorkingOfficeDivisionFacilityID"`
	WorkingOfficeDivisionName       string                        `json:"WorkingOfficeDivisionName"`
	WorkingOfficeFacilityID         string                        `json:"WorkingOfficeFacilityID"`
	WorkingOfficeName               string                        `json:"WorkingOfficeName"`
	WorkingOfficePincode            int32                         `json:"WorkingOfficePincode"`
	WorkingOfficeRegionFacilityID   string                        `json:"WorkingOfficeRegionFacilityID"`
	WorkingOfficeRegionName         string                        `json:"WorkingOfficeRegionName"`
}

// func UpdateCenterCodeForApplicationsPSold(client *ent.Client) gin.HandlerFunc {
// 	return func(gctx *gin.Context) {

// 		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
// 		defer cancel()
// 		var mainFunction string = " main - UpdateCenterCodeForApplicationsPS "
// 		var startFunction string = " - start - UpdateCenterCodeForApplicationsPS "

// 		var reqData requestDataArray
// 		if err := gctx.ShouldBindJSON(&reqData); err != nil {
// 			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
// 			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
// 			return
// 		}

// 		var status int32
// 		var totalUpdatedCount, count int
// 		var dataStatus bool
// 		var err error
// 		var stgError string

// 		for _, requestData := range reqData {
// 			count, status, stgError, dataStatus, err = start.UpdateCenterCodeForApplicationsPS(ctx, client, requestData.ControllingOfficeFacilityID, int32(requestData.ExamCenterID), int32(requestData.SeatsToAllot), int32(requestData.ExamCityID))
// 			if err != nil {
// 				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
// 				return
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
