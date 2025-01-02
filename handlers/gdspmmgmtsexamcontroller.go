package handlers

import (
	"context"
	"fmt"
	"net/http"
	"recruit/ent"
	"strconv"

	"recruit/mail"
	"recruit/sms"
	"recruit/start"
	"recruit/util"

	ca_reg "recruit/payloadstructure/candidate_registration"

	"github.com/gin-gonic/gin"
)

// CreateNewGDSPMApplicationss godoc
// @Summary Create new GDS to PM/MG/MTS applications
// @Description Submit new GDS to PM/MG/MTS applications
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param newPAAppplns body ca_reg.ApplicationGDSPM true "New GDS to PM/MG/MTS Application Data"
// @Success 200 {object} start.CreateGDSPMResponse "GDS to PM/MG/MTS application submitted successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/Applications/Submit [post]
func CreateNewGDSPMApplicationss(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - CreateNewGDSPMApplicationss "
		var startFunction string = " - start - CreateGDSPMApplicationss "

		var newPAAppplns ca_reg.ApplicationGDSPM
		if err := gctx.ShouldBindJSON(&newPAAppplns); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newPAAppplns.Edges.LogData[0]
		newPAApppln, status, stgError, dataStatus, err := start.CreateGDSPMApplicationss(client, &newPAAppplns)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var submittedDate string = newPAApppln.ApplnSubmittedDate.Format("02-01-2006 15:04:05")

		var Smsuserid = fmt.Sprint(newPAAppplns.EmployeeID)
		//For 3 examshortname ,appno ,ca submittedDate
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 3, " CE GDS to PM/MG/MTS ", newPAApppln.ApplicationNumber, submittedDate)

		//var smssentstatus string
		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "CE GDS to PM/MG/MTS", newPAApppln.ApplicationNumber, "Submitted")

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "GDS to PM/MG/MTS submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        newPAApppln,
			"dataexists":  dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// Get IP Applications with EmpID - QueryIPExamApplicationsByEmpID

// GetGDSPMApplicationsByEmpId godoc
// @Summary Get GDS to PM/MG/MTS applications by Employee ID
// @Description Fetch GDS to PM/MG/MTS applications for a given employee ID and exam year
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Application details fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/applications/getbyempid/{employeeid}/{examyear} [get]
func GetGDSPMApplicationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetGDSPMApplicationsByEmpId "
		var startFunction string = " - start - QueryGDSPMExamApplicationsByEmpID "
		//employeeid/:examyear
		id := gctx.Param("employeeid")
		id1 := gctx.Param("examyear")

		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		paapplns, status, stgError, dataStatus, err := start.QueryGDSPMExamApplicationsByEmpID(ctx, client, int64(empid), string(id1))
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Application details fetched successfully",
			"data":       paapplns,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// VerifyGDSPMApplications godoc
// @Summary Verify GDS to PM/MG/MTS applications
// @Description Verify GDS to PM/MG/MTS applications by updating remarks
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param request body ca_reg.VerifyApplicationGDSPM true "Verify GDS PM Application"
// @Success 200 {object} start.CreateGDSPMResponse "GDS to PM/MG/MTS application verified successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/applications/Verify [put]
func VerifyGDSPMApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - VerifyGDSPMApplications "
		var startFunction string = " - start - UpdateApplicationRemarkssGDSPM "

		var newAppln ca_reg.VerifyApplicationGDSPM
		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newAppln.Edges.LogData[0]
		var nonQualifyService string
		if newAppln.NonQualifyingService == nil || len(*newAppln.NonQualifyingService) == 0 {
			nonQualifyService = "No"
		} else {
			nonQualifyService = "Yes"
		}
		application, status, stgError, dataStatus, err := start.UpdateApplicationRemarkssGDSPM(client, &newAppln, nonQualifyService)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var emailsentstatus string = ""
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		if application.ApplicationStatus == "VerifiedByCA" {
			//examshortname ,appno ,ca,submittedDate ,recommended, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " CE GDS to PM/MG/MTS ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.CAGeneralRemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " CE GDS to PM/MG/MTS ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.CAGeneralRemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByCA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "CE GDS to PM/MG/MTS", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "CE GDS to PM/MG/MTS", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "GDS to PM/MG/MTS application verified successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// ResubmitGDSPMApplications godoc
// @Summary Resubmit GDS to PM/MG/MTS applications
// @Description Resubmit GDS to PM/MG/MTS applications with updated remarks
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param request body ca_reg.ReApplicationGDSPM true "Resubmit GDS PM Application"
// @Success 200 {object} start.CreateGDSPMResponse "GDS to PM/MG/MTS application re-submitted successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/Applications/resubmit [put]
func ResubmitGDSPMApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - ResubmitGDSPMApplications "
		var startFunction string = " - start - ResubmitApplicationRemarkssGDSPM "

		var newAppln ca_reg.ReApplicationGDSPM
		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newAppln.Edges.LogData[0]
		application, status, stgError, dataStatus, err := start.ResubmitApplicationRemarkssGDSPM(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		//examshortname ,	appno, ca, submittedDate
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 7, " CE GDS to PM/MG/MTS ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "CE GDS to PM/MG/MTS", application.ApplicationNumber, "Re-submitted")

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "GDS to PM/MG/MTS application re-submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// VerifyGDSPMVAApplications godoc
// @Summary Verify GDS to PM/MG/MTS applications by Sub Verifying Authority
// @Description Verify GDS to PM/MG/MTS applications with updated remarks by Sub Verifying Authority
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param request body ca_reg.VerifyApplicationVAGDSPM true "Verify GDS PM Application by Sub Verifying Authority"
// @Success 200 {object} start.VerifyGDSPMVAApplicationsResponse "GDS to PM/MG/MTS application verified successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/applications/vaVerify [put]
func VerifyGDSPMVAApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		var mainFunction string = " main - VerifyGDSPMVAApplications "
		var startFunction string = " - start - UpdateApplicationRemarksVAGDSPM "
		var newAppln ca_reg.VerifyApplicationVAGDSPM
		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		if len(newAppln.Edges.LogData) <= 0 {
			Remarks = "main - VerifyGDSPMVAApplications - Log data received blank"
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newAppln.Edges.LogData[0]

		application, status, stgError, dataStatus, err := start.UpdateApplicationRemarksVAGDSPM(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Sub Verifying authority verified Successfully  ",
			"data": gin.H{
				"EmployeeID":          application.EmployeeID,
				"ApplicationStatus":   application.ApplicationStatus,
				"Application Remarks": "Sub Verifying authority verified Successfully  ",
				"RoleUserCode":        application.RoleUserCode,
			},
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

// get ca pending verify records

// GetAllGDSPMCAPendingVerifications godoc
// @Summary Get all GDS to PM/MG/MTS applications pending CA verifications by facility ID and selected year
// @Description Fetches all pending GDS to PM/MG/MTS applications that are awaiting CA verifications
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Successfully fetched pending applications with CA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facility or applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getallcapendingapplications/{facilityid}/{selectedyear} [get]
func GetAllGDSPMCAPendingVerifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetAllGDSPMCAPendingVerifications "
		var startFunction string = " - start - QueryGDSPMApplicationsByCAVerificationsPending "

		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByCAVerificationsPending(ctx, client, facilityID, id1)
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

// CA Verified with Emp ID

// GetGDSPMCAPendingDetailsByEmpId godoc
// @Summary Get GDS to PM/MG/MTS application pending details by employee ID and exam year
// @Description Fetches GDS to PM/MG/MTS applications that are pending with CA verifications based on employee ID and exam year
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse  "Successfully fetched pending application details with CA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee or applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getallcapending/{employeeid}/{examyear} [get]
func GetGDSPMCAPendingDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetGDSPMCAPendingDetailsByEmpId "
		var startFunction string = " - start - QueryGDSPMApplicationsByCAPendingByEmpID "

		id := gctx.Param("employeeid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("examyear")
		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByCAPendingByEmpID(ctx, client, int64(empid), examYear)
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

// GetGDSPMCAPendingOldRemarksByEmpId godoc
// @Summary Get old remarks for GDS to PM/MG/MTS applications pending with candidate by employee ID and exam year
// @Description Fetches old remarks for GDS to PM/MG/MTS applications pending with candidate based on employee ID and exam year
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetGDSPMCAPendingOldRemarksByEmpIdResponse "Successfully fetched old remarks for pending applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee or applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/caprevremarks/{employeeid}/{examyear} [get]
func GetGDSPMCAPendingOldRemarksByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetGDSPMCAPendingOldRemarksByEmpId "
		var startFunction string = " - start - GetOldGDSPMCAApplicationRemarksByEmployeeID "

		id := gctx.Param("employeeid")
		examYear := gctx.Param("examyear")

		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, status, stgError, dataStatus, err := start.GetOldGDSPMCAApplicationRemarksByEmployeeID(ctx, client, int64(empid), examYear)
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

// Get All CA Verified records

// GetAllGDSPMCAVerified godoc
// @Summary Get all CA verified GDS to PM/MG/MTS applications
// @Description Fetches all GDS to PM/MG/MTS applications that have been verified by CA based on facility ID and selected year
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Successfully fetched CA verified applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facility or applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getallcaverifiedapplications/{facilityid}/{selectedyear} [get]
func GetAllGDSPMCAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetAllGDSPMCAVerified "
		var startFunction string = " - start - QueryGDSPMApplicationsByCAVerified "
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByCAVerified(ctx, client, facilityID, id1)
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

// CA Verified with Emp ID

// GetGDSPMCAVerifiedDetailsByEmpId godoc
// @Summary Get CA verified GDS to PM/MG/MTS application details by Employee ID
// @Description Fetches the GDS to PM/MG/MTS application details that have been verified by CA based on Employee ID and exam year
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Successfully fetched application details with CA verified status"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee or application not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getallcaverified/{employeeid}/{examyear} [get]
func GetGDSPMCAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetGDSPMCAVerifiedDetailsByEmpId "
		var startFunction string = " - start - QueryGDSPMApplicationsByCAVerifiedByEmpID "
		//employeeid
		id := gctx.Param("employeeid")
		examYear := gctx.Param("examyear")

		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByCAVerifiedByEmpID(ctx, client, int64(empid), examYear)
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

// Get all CA verified for NA..

// GetGDSPMAllCAVerifiedForNA godoc
// @Summary Get all GDS to PM/MG/MTS applications verified by CA for NA
// @Description Fetches all GDS to PM/MG/MTS applications that have been verified by CA for a given facility and year
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Successfully fetched applications verified by CA for NA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facility or application not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getallcaverifiedapplicationsforna/{facilityid}/{selectedyear} [get]
func GetGDSPMAllCAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetGDSPMAllCAVerifiedForNA "
		var startFunction string = " - start - QueryGDSPMApplicationsByCAVerifiedForNA "

		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		facilityID1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByCAVerifiedForNA(ctx, client, facilityID, facilityID1)
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

// Update Nodal recommendations

// UpdateNodalRecommendationsGDSPMByEmpID godoc
// @Summary Update Nodal Recommendations for GDSPM by Employee ID
// @Description Updates the nodal recommendations for a given GDSPM application based on Employee ID
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param request body ca_reg.NAVerifyApplicationGDSPM true "GDSPM Application"
// @Success 200 {object} start.CreateGDSPMResponse "Nodal officer successfully verified the GDSPMMG application"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee or application not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/noverify [put]
func UpdateNodalRecommendationsGDSPMByEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - UpdateNodalRecommendationsGDSPMByEmpID "
		var startFunction string = " - start - UpdateGDSPMNodalRecommendationsByEmpID "
		var newAppln ca_reg.NAVerifyApplicationGDSPM
		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newAppln.Edges.LogData[0]
		fmt.Println(logdata)

		application, status, stgError, dataStatus, err := start.UpdateGDSPMNodalRecommendationsByEmpID(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		var emailsentstatus string = ""
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		if application.ApplicationStatus == "VerifiedByNA" {
			//examshortname ,appno ,ca,submittedDate ,recommended, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " CE GDS to PM/MG/MTS ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.NARemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " CE GDS to PM/MG/MTS ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.NARemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByNA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "CE GDS to PM/MG/MTS", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "CE GDS to PM/MG/MTS", application.ApplicationNumber, "Returned for correction")
		}
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Nodal officer successfully verified this GDSPMMG application ",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetRecommendationsByEmpID

// GetGDSPMExamRecommendationsByEmpId godoc
// @Summary Get GDSPM Exam Recommendations by Employee ID
// @Description Fetches GDSPM exam recommendations based on Employee ID
// @Tags GDS PM Applications
// @Produce json
// @Param employeeid path string true "Employee ID"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Fetching recommendations by employee ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Employee or recommendations not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/recommendations/{employeeid} [get]
func GetGDSPMExamRecommendationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		id := gctx.Param("employeeid")
		var mainFunction string = " main - GetGDSPMExamRecommendationsByEmpId "
		var startFunction string = " - start - GetGDSPMRecommendationsByEmpID "
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		records, status, stgError, dataStatus, err := start.GetGDSPMRecommendationsByEmpID(client, empid)
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

// GetGDSPMAllNAVerified godoc
// @Summary Get GDSPM Applications Verified by NA
// @Description Fetches GDSPM applications verified by NA based on facility ID and selected year
// @Tags GDS PM Applications
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse  "Fetching applications verified by NA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getallnaverifiedapplications/{facilityid}/{selectedyear} [get]
func GetGDSPMAllNAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetGDSPMAllNAVerified "
		var startFunction string = " - start - QueryGDSPMApplicationsByNAVerified "
		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		selectedYear := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByNAVerified(ctx, client, facilityID, selectedYear)
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

// Get na verified record by Emp ID

// GetGDSPMNAVerifiedDetailsByEmpId godoc
// @Summary Get GDSPM Applications NA Verified Details by Employee ID
// @Description Fetches GDSPM applications NA verified details by employee ID and exam year
// @Tags GDS PM Applications
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Fetching NA verified application by employee ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getallnaverified/{employeeid}/{examyear} [get]
func GetGDSPMNAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetGDSPMNAVerifiedDetailsByEmpId "
		var startFunction string = " - start - QueryGDSPMApplicationsByNAVerifiedByEmpID "
		id := gctx.Param("employeeid")
		examYear := gctx.Param("examyear")

		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByNAVerifiedByEmpID(ctx, client, int64(empid), examYear)
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

// GetGDSPMAllNAVerifiedForNA godoc
// @Summary Get GDSPM Applications NA Verified Details for NA
// @Description Fetches GDSPM applications NA verified details for a specific facility and selected year
// @Tags GDS PM Applications
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse  "Fetching applications verified by NA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getallnaverifiedapplicationsforna/{facilityid}/{selectedyear} [get]
func GetGDSPMAllNAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetGDSPMAllNAVerifiedForNA "
		var startFunction string = " - start - QueryGDSPMApplicationsByNAVerifiedForNA "
		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		facilityID1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByNAVerifiedForNA(ctx, client, facilityID, facilityID1)
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

// UpdateExamCentersInGDSPMApplsreturnstring godoc
// @Summary Update Exam Centers in GDSPM Applications and return string
// @Description Updates exam centers for GDSPM applications and returns string response
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param request body ca_reg.ApplicationGDSPMforUpdateExamCenters  true "Request body to update exam centers"
// @Success 200 {object} start.GetGDSPMApplicationsResponse  "Exam centers updated successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/center/updatecenters [put]
func UpdateExamCentersInGDSPMApplsreturnstring(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateExamCentersInGDSPMApplsreturnstring "
		var startFunction string = " - start - UpdateExamCentresGDSPMExamsreturnarray "
		// Bind the JSON data to the newappls variable
		var req struct {
			Newappls []*ent.Exam_Applications_GDSPM `json:"newappls"`
		}
		if err := gctx.ShouldBindJSON(&req); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		// Call the UpdateExamCentresIPExams function to update the exam centers
		examCenters, status, stgError, dataStatus, err := start.UpdateExamCentresGDSPMExamsreturnarray(ctx, client, req.Newappls)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Updating Exam centers for IP Applications",
			"data":       examCenters,
			"dataexists": dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// GetExamApplicatonsPreferenenceCityWiseStatsGDSPM godoc
// @Summary Get Exam Applications Preference City Wise Statistics for GDSPM
// @Description Fetches statistics for GDSPM exam applications based on exam year, exam code, and city ID
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param examYear path string true "Exam Year"
// @Param examcode path string true "Exam Code"
// @Param cityid path string true "City ID"
// @Success 200 {object} start.GetGDSPMApplicationsResponse  "Statistics fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Data not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/getExamApplicationsByCityPrefGDSPM/{examYear}/{examcode}/{cityid} [get]
func GetExamApplicatonsPreferenenceCityWiseStatsGDSPM(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetExamApplicatonsPreferenenceCityWiseStatsGDSPM "
		var startFunction string = " - start - GetExamApplicatonsPreferenenceCityWiseStatsGDSPM "

		//:examYear/:examcode/:cityid
		ExamYear := gctx.Param("examYear")
		Examcode := gctx.Param("examcode")
		Cityid := gctx.Param("cityid")

		var statistics []start.ExamStatsGDSPM
		var err error
		statistics, status, stgError, dataStatus, err := start.GetExamApplicatonsPreferenenceCityWiseStatsGDSPM(ctx, client, ExamYear, Examcode, Cityid)

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

// UpdateCenterCodeForApplicationsGDSPM godoc
// @Summary Update Center Code for Applications in GDSPM
// @Description Updates center codes for GDSPM applications based on provided data array
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param request body []requestDataArray true "Array of request data to update center codes"
// @Success 200 {object} start.GetGDSPMApplicationsResponse  "Applications updated successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/updateCenterCodeForApplicationsGDSPM [put]
func UpdateCenterCodeForApplicationsGDSPM(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateCenterCodeForApplicationsGDSPM "
		var startFunction string = " - start - UpdateCenterCodeForApplicationsGDSPM "

		var reqData requestDataArray

		if err := gctx.BindJSON(&reqData); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		var totalUpdatedCount int
		var updatedApplications []*ent.Exam_Applications_GDSPM

		var smsStatus, emailSentStatus string
		var dataStatus bool

		for _, requestData := range reqData {
			count, updatedApps, status, stgError, currentDataStatus, err := start.UpdateCenterCodeForApplicationsGDSPM(ctx, client, requestData.ControllingOfficeFacilityID, int32(requestData.ExamCenterID), int32(requestData.SeatsToAllot), int32(requestData.ExamCityID))
			if err != nil {
				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
				return
			}
			totalUpdatedCount += count
			totalUpdatedCount += count
			updatedApplications = append(updatedApplications, updatedApps...)
			dataStatus = currentDataStatus
			for _, app := range updatedApps {
				employeeIDStr := strconv.FormatInt(app.EmployeeID, 10)
				smsStatus = sms.SendSmsNew(ctx, client, employeeIDStr, 12, app.ApplicationNumber, app.HallTicketNumber, "CE GDS to PM/MG/MTS")
				var submittedDate string = app.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
				var hallticketDate string = app.HallTicketGeneratedDate.Format("02-01-2006 15:04:05")
				emailSentStatus = mail.SendEMailNew(ctx, client, employeeIDStr, 12, "CE GDS to PM/MG/MTS", app.ApplicationNumber, app.ControllingOfficeName, submittedDate, app.NodalOfficeName, app.HallTicketNumber, hallticketDate)
			}
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":             true,
			"message":             fmt.Sprintf("Total Number of Applications %d Updated with center code Successfully", totalUpdatedCount),
			"data":                struct{}{},
			"dataexists":          dataStatus,
			"SMSStatus":           smsStatus,
			"EmailStatus":         emailSentStatus,
			"updatedApplications": updatedApplications,
		})
	}

	return gin.HandlerFunc(fn)
}

// GetGDSPMHallTicketWithExamCodeEmpID godoc
// @Summary Get GDSPM Hall Ticket by Exam Code and Employee ID
// @Description Fetches GDSPM applications with hall ticket based on exam code, employee ID, and exam year
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param examcode path int true "Exam Code"
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Applications fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/hallticket/get/{examcode}/{employeeid}/{examyear} [get]
func GetGDSPMHallTicketWithExamCodeEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		var mainFunction string = " main - GetGDSPMHallTicketWithExamCodeEmpID "
		var startFunction string = " - start - GetGDSPMApplicationsWithHallTicket "
		ec := gctx.Param("examcode")
		ExamCode, _ := strconv.ParseInt(ec, 10, 32)
		EmployeeID, _ := strconv.ParseInt(gctx.Param("employeeid"), 10, 64)
		examYear := gctx.Param("examyear")
		examcenters, status, stgError, dataStatus, err := start.GetGDSPMApplicationsWithHallTicket(client, int32(ExamCode), EmployeeID, examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching application based on Exam year",
			"data":       examcenters,
			"dataexists": dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// GetAllGDSPMPendingWithCandidate godoc
// @Summary Get all GDSPM applications pending with candidate
// @Description Fetches GDSPM applications that are pending with the candidate based on facility ID and selected year
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Applications fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getAllPendingWithCandidate/{facilityid}/{selectedyear} [get]
func GetAllGDSPMPendingWithCandidate(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllGDSPMPendingWithCandidate "
		var startFunction string = " - start - QueryGDSPMApplicationsByPendingWithCandidate "
		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		facilityID1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByPendingWithCandidate(ctx, client, facilityID, facilityID1)
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

type requestBodyy struct {
	ExamYear               string `json:"ExamYear"`
	ExamCode               int32  `json:"ExamCode"`
	NodalOfficerFacilityID string `json:"nodalOfficerFacilityID"`
}

// GenerateHallticketNumbersGDSPM godoc
// @Summary Generate Hall Ticket Numbers for GDSPM
// @Description Generates hall ticket numbers for GDSPM based on the provided exam year, exam code, and nodal officer facility ID
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param request body requestBodyy true "Request body to generate hall ticket numbers"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Hall Ticket generation done successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/Halltickets [put]
func GenerateHallticketNumbersGDSPM(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		// Define a struct to represent the request body

		var mainFunction string = " main - GenerateHallticketNumbersGDSPM "
		var startFunction string = " - start - GenerateHallticketNumberrGdsPm "
		// Parse the request body into the struct
		var reqBody requestBodyy
		if err := gctx.BindJSON(&reqBody); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		// Call GenerateHallticketNumberrIP function
		successMessage, status, stgError, dataStatus, err := start.GenerateHallticketNumberrGdsPm(ctx, client, reqBody.ExamYear, reqBody.ExamCode, reqBody.NodalOfficerFacilityID)
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

// GetAllGDSPMVAPendingVerifications godoc
// @Summary Get all GDSPM VA pending verifications
// @Description Fetches GDSPM applications that are pending verification by the VA (Verification Authority) based on facility ID and selected year
// @Tags GDS PM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetGDSPMApplicationsResponse "Applications fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GDSPMexams/getallvapendingapplications/{facilityid}/{selectedyear} [get]
func GetAllGDSPMVAPendingVerifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")
		var mainFunction string = " main - GetAllGDSPMVAPendingVerifications "
		var startFunction string = " - start - QueryGDSPMApplicationsByVAVerificationsPending "
		circles, status, stgError, dataStatus, err := start.QueryGDSPMApplicationsByVAVerificationsPending(ctx, client, facilityID, id1)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Sub Division authority verified this application",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}
