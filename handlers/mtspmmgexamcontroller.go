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

	ca_reg "recruit/payloadstructure/candidate_registration"

	"github.com/gin-gonic/gin"
)

// MTS PM MG FUNC

// CreateNewMTSPMMGApplications godoc
// @Summary Create a new MTS to PM/MG application
// @Description Creates a new MTS to PM/MG application with the provided details
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param request body ca_reg.ApplicationMTSPM true "Request body for creating a new MTS to PM/MG application"
// @Success 200 {object} start.CreateApplicationResponse "MTS to PM/MG Application submitted successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/applications/submit [post]
func CreateNewMTSPMMGApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - CreateNewMTSPMMGApplications "
		var startFunction string = " - start - CreateMTSPMMGApplications "
		var newPAAppplns ca_reg.ApplicationGDSPM

		if err := gctx.ShouldBindJSON(&newPAAppplns); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newPAAppplns.Edges.LogData[0]

		newPAApppln, status, stgError, dataStatus, err := start.CreateMTSPMMGApplications(client, &newPAAppplns)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var Smsuserid string = fmt.Sprint(newPAApppln.EmployeeID)
		var submittedDate string = newPAApppln.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		//examshortname ,appno ,ca submittedDate
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 3, " LDCE MTS to PM/MG ", newPAApppln.ApplicationNumber, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE MTS to PM/MG", newPAApppln.ApplicationNumber, "Submitted")

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "MTS to PM/MG Application submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        newPAApppln,
			"dataexists":  dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

//verify application

// VerifyMTSPMMGApplication godoc
// @Summary Verify MTS to PM/MG Application
// @Description Verifies the MTS to PM/MG application and updates its status
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param request body ca_reg.VerifyApplicationMTSPM true "Verify Application Request"
// @Success 200 {object} start.CreateApplicationResponse "MTS to PM/MG exam application verified successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/applications/verify [put]
func VerifyMTSPMMGApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - VerifyMTSPMMGApplication "
		var startFunction string = " - start - UpdateApplicationRemarksMTSPMMG "

		var newAppln ca_reg.VerifyApplicationMTSPM
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
		application, status, stgError, dataStatus, err := start.UpdateApplicationRemarksMTSPMMG(client, &newAppln, nonQualifyService)
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
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE for MTS to PM/MG ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.CAGeneralRemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE for MTS to PM/MG ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.CAGeneralRemarks)
		}
		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByCA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE for MTS to PM/MG", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE for MTS to PM/MG", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "MTS to PM/MG exam application verified successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// ResubmitMTSPMMGApplication godoc
// @Summary Resubmit MTS to PM/MG Application
// @Description Resubmits the MTS to PM/MG application and updates its status
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param request body ca_reg.ReApplicationGDSPM true "Re-Application Request"
// @Success 200 {object} start.CreateApplicationResponse "MTS to PM/MG exam application re-submitted successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/applications/resubmit [put]
func ResubmitMTSPMMGApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - ResubmitMTSPMMGApplication "
		var startFunction string = " - start - ResubmitApplicationRemarksMTSPMMG "

		var newAppln ca_reg.ReApplicationGDSPM

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newAppln.Edges.LogData[0]

		application, status, stgError, dataStatus, err := start.ResubmitApplicationRemarksMTSPMMG(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		//examshortname ,	appno, ca, submittedDate
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 7, " LDCE MTS to PM/MG ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE MTS to PM/MG", application.ApplicationNumber, "Re-Submitted")

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "MTS to PM/MG exam application re-submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// MTS PM M G GET APPLICATIONS

// GetMTSPMMGApplicationsByEmpId godoc
// @Summary Get MTS to PM/MG Applications by Employee ID
// @Description Retrieves the MTS to PM/MG applications for a given employee ID and exam year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param employeeid path string true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetApplicationsResponse "Application details fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/applications/getbyempid/{employeeid}/{examyear} [get]
func GetMTSPMMGApplicationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetMTSPMMGApplicationsByEmpId "
		var startFunction string = " - start - QueryMTSPMMGPMExamApplicationsByEmpID "
		//employeeid/:examyear
		id := gctx.Param("employeeid")
		id1 := gctx.Param("examyear")

		empid, _ := strconv.ParseInt(id, 10, 64)
		paapplns, status, stgError, dataStatus, err := start.QueryMTSPMMGPMExamApplicationsByEmpID(ctx, client, int64(empid), string(id1))
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

// Get all pending

// GetAllMTSPMMGCAPendingVerifications godoc
// @Summary Get all MTS to PM/MG CA pending verifications
// @Description Retrieves all MTS to PM/MG applications pending CA verifications for a given facility ID and selected year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetched pending applications with CA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/getallcapendingapplications/{facilityid}/{selectedyear} [get]
func GetAllMTSPMMGCAPendingVerifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetAllCAPendingVerifications "
		var startFunction string = " - start - QueryIPApplicationsByCAVerificationsPending "
		//facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryMTSPMMGApplicationsByCAVerificationsPending(ctx, client, facilityID, id1)
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

// CA PENDING BY EMP ID

// GetMTSPMMGCAPendingDetailsByEmpId godoc
// @Summary Get MTS to PM/MG CA pending details by employee ID
// @Description Retrieves details of MTS to PM/MG applications pending CA verification for a given employee ID and selected year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetched pending details with CA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Details not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/getallcapending/{employeeid}{examyear} [get]
func GetMTSPMMGCAPendingDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetMTSPMMGCAPendingDetailsByEmpId "
		var startFunction string = " - start - QueryPMTSPMMGApplicationsByCAPendingByEmpID "
		id := gctx.Param("employeeid")

		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("selectedyear")
		data, status, stgError, dataStatus, err := start.QueryPMTSPMMGApplicationsByCAPendingByEmpID(ctx, client, int64(empid), examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetching application pending with CA",
			"data":       data,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// Get all CA verified for NA..

// MtspmAllCAVerifiedForNA godoc
// @Summary Get all MTS to PM/MG applications verified by CA for NA
// @Description Retrieves all MTS to PM/MG applications that are verified by CA for NA for a given facility ID and selected year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetched CA Verified applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Details not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/getallcaverifiedapplicationsforna/{facilityid}/{selectedyear} [get]
func MtspmAllCAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		var mainFunction string = " main - MtspmAllCAVerifiedForNA "
		var startFunction string = " - start - QueryMTSPMApplicationsByCAVerifiedForNA "

		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		facilityID1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryMTSPMApplicationsByCAVerifiedForNA(ctx, client, facilityID, facilityID1)
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

// GetAllMTSPMMGCAVerified godoc
// @Summary Get all MTS to PM/MG applications verified by CA
// @Description Retrieves all MTS to PM/MG applications that are verified by CA for a given facility ID and selected year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetched CA Verified applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Details not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/getallcaverifiedapplications/{facilityid}/{selectedyear} [get]
func GetAllMTSPMMGCAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//facilityid/:selectedyear

		var mainFunction string = " main - GetAllMTSPMMGCAVerified "
		var startFunction string = " - start - QueryMTSPMMGApplicationsByCAVerified "
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryMTSPMMGApplicationsByCAVerified(ctx, client, facilityID, id1)
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

// GetMTSPMMGCAVerifiedDetailsByEmpId godoc
// @Summary Get details of MTS to PM/MG applications verified by CA based on employee ID
// @Description Retrieves details of MTS to PM/MG applications that are verified by CA for a given employee ID and selected year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching applicatins pending with candiate with remarks"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Details not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/getallcaverified/{employeeid}/{selectedyear} [get]
func GetMTSPMMGCAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetMTSPMMGCAVerifiedDetailsByEmpId "
		var startFunction string = " - start - QueryMTSPMMGApplicationsByCAVerifiedByEmpID "
		//employeeid
		id := gctx.Param("employeeid")

		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("selectedyear")
		circles, status, stgError, dataStatus, err := start.QueryMTSPMMGApplicationsByCAVerifiedByEmpID(ctx, client, int64(empid), examYear)
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

// GetMTSPMMGCAPendingOldRemarksByEmpId godoc
// @Summary Get old CA application remarks by employee ID
// @Description Retrieves old CA application remarks for MTS to PM/MG applications pending with a candidate based on employee ID and selected year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching applicatins pending with candiate with remarks"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Details not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/caprevremarks/{employeeid}/{examyear} [get]
func GetMTSPMMGCAPendingOldRemarksByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetCAPendingOldRemarksByEmpId "
		var startFunction string = " - start - GetOldCAApplicationRemarksByEmployeeID "
		id := gctx.Param("employeeid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		examYear := gctx.Param("selectedyear")
		circles, status, stgError, dataStatus, err := start.GetOldMTSPMMGCAApplicationRemarksByEmployeeID(ctx, client, int64(empid), examYear)
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

// Update Nodal recommendations

// UpdateNodalRecommendationsMTSPMMGByEmpID godoc
// @Summary Update nodal recommendations for MTS to PM/MG by employee ID
// @Description Updates nodal recommendations for the MTS to PM/MG application and sends notifications
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param request body ca_reg.NAVerifyApplicationMTSPM true "Nodal Verification Request"
// @Success 200 {object} start.CreateApplicationResponse "Nodal officer successfully verified this MTS to PM/MG application "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Details not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/noverify [put]
func UpdateNodalRecommendationsMTSPMMGByEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateNodalRecommendationsMTSPMMGByEmpID "
		var startFunction string = " - start - UpdateMTSPMMGNodalRecommendationsByEmpID "
		var newAppln ca_reg.NAVerifyApplicationMTSPM
		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newAppln.Edges.LogData[0]

		application, status, stgError, dataStatus, err := start.UpdateMTSPMMGNodalRecommendationsByEmpID(client, &newAppln)
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
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE for MTS to PM/MG ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.NARemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE for MTS to PM/MG ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.NARemarks)
		}
		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByNA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE for MTS to PM/MG", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE for MTS to PM/MG", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Nodal officer successfully verified this MTS to PM/MG application ",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetRecommendationsByEmpID

// GetMTSPMMGExamRecommendationsByEmpId godoc
// @Summary Get MTS to PM/MG exam recommendations by employee ID
// @Description Fetches MTS to PM/MG exam recommendations based on employee ID
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param employeeid path int true "Employee ID"
// @Success 200 {object} start.GetApplicationsResponse "Fetching recommendations by employee ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Details not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/recommendations/{employeeid} [get]
func GetMTSPMMGExamRecommendationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		id := gctx.Param("employeeid")
		var mainFunction string = " main - GetMTSPMMGExamRecommendationsByEmpId "
		var startFunction string = " - start - GetMTSPMMGRecommendationsByEmpID "

		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		records, status, stgError, dataStatus, err := start.GetMTSPMMGRecommendationsByEmpID(client, empid)
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

// Get all NA Verified for NA ..

// GetMTSPMMGAllNAVerifiedForNA godoc
// @Summary Get all MTS to PM/MG applications verified by NA for NA
// @Description Fetches MTS to PM/MG applications verified by National Authority (NA) for National Authority (NA) based on facility ID and selected year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching applications verified by NA"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Details not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/getallnaverifiedapplicationsforna/{facilityid}/{selectedyear} [get]
func GetMTSPMMGAllNAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		var mainFunction string = " main - GetMTSPMMGAllNAVerifiedForNA "
		var startFunction string = " - start - QueryMTSPMMGApplicationsByNAVerifiedForNA "
		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		facilityID1 := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryMTSPMMGApplicationsByNAVerifiedForNA(ctx, client, facilityID, facilityID1)
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

type ApplicationMTSPMUpdatecenter struct {
	Cadre                           string              `json:"Cadre"`
	CandidateRemarks                string              `json:"CandidateRemarks" `
	CategoryCode                    string              `json:"CategoryCode"`
	CategoryDescription             string              `json:"CategoryDescription"`
	CenterFacilityId                string              `json:"CenterFacilityId" `
	CenterId                        int32               `json:"CenterId" `
	CentrePreference                string              `json:"CentrePreference"`
	ClaimingQualifyingService       string              `json:"ClaimingQualifyingService"`
	ControllingOfficeFacilityID     string              `json:"ControllingOfficeFacilityID" `
	ControllingOfficeName           string              `json:"ControllingOfficeName" `
	DCCS                            string              `json:"DCCS"`
	DOB                             string              `json:"DOB" `
	DeputationControllingOfficeID   string              `json:"DeputationControllingOfficeID" `
	DeputationControllingOfficeName string              `json:"DeputationControllingOfficeName" `
	DeputationOfficeFacilityID      string              `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string              `json:"DeputationOfficeName" `
	DeputationOfficePincode         string              `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string              `json:"DeputationOfficeUniqueId" `
	DeputationType                  string              `json:"DeputationType"`
	DesignationID                   string              `json:"DesignationID"`
	DisabilityPercentage            int32               `json:"DisabilityPercentage"`
	DisabilityTypeCode              string              `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string              `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string              `json:"DisabilityTypeID"`
	Edges                           ca_reg.EdgesLogdata `json:"edges" `
	EducationCode                   string              `json:"EducationCode"`
	EducationDescription            string              `json:"EducationDescription"`
	EmailID                         string              `json:"EmailID"`
	EmployeeID                      int64               `json:"EmployeeID" `
	EmployeeName                    string              `json:"EmployeeName"`
	EmployeePost                    string              `json:"EmployeePost" `
	EntryPostCode                   string              `json:"EntryPostCode"`
	EntryPostDescription            string              `json:"EntryPostDescription"`
	ExamCode                        int32               `json:"ExamCode"`
	ExamName                        string              `json:"ExamName"`
	ExamShortName                   string              `json:"ExamShortName"`
	ExamYear                        string              `json:"ExamYear"`
	FacilityName                    string              `json:"FacilityName"`
	FacilityUniqueID                string              `json:"FacilityUniqueID"`
	FeederPostCode                  string              `json:"FeederPostCode"`
	FeederPostDescription           string              `json:"FeederPostDescription"`
	FeederPostJoiningDate           string              `json:"FeederPostJoiningDate"`
	Gender                          string              `json:"Gender" `
	GDSEngagement                   *[]interface{}      `json:"GDSEngagement" `
	LienControllingOfficeID         string              `json:"LienControllingOfficeID"`
	LienControllingOfficeName       string              `json:"LienControllingOfficeName"`
	MobileNumber                    string              `json:"MobileNumber" `
	NodalOfficeFacilityID           string              `json:"NodalOfficeFacilityID"`
	NodalOfficeName                 string              `json:"NodalOfficeName"`
	Photo                           string              `json:"Photo"`
	PhotoPath                       string              `json:"PhotoPath"`
	PostPreferences                 *[]interface{}      `json:"PostPreferences"`
	PresentDesignation              string              `json:"PresentDesignation"`
	PresentPostCode                 string              `json:"PresentPostCode"`
	PresentPostDescription          string              `json:"PresentPostDescription"`
	ReportingOfficeFacilityID       string              `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string              `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{}      `json:"ServiceLength"`
	Signature                       string              `json:"Signature"`
	SignaturePath                   string              `json:"SignaturePath"`
	TempHallTicket                  string              `json:"TempHallTicket"`
	UnitPreferences                 *[]interface{}      `json:"UnitPreferences"`
	UserID                          int32               `json:"UserID"`
	WorkingOfficeCircleFacilityID   string              `json:"WorkingOfficeCircleFacilityID"`
	WorkingOfficeCircleName         string              `json:"WorkingOfficeCircleName"`
	WorkingOfficeDivisionFacilityID string              `json:"WorkingOfficeDivisionFacilityID" `
	WorkingOfficeDivisionName       string              `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string              `json:"WorkingOfficeFacilityID"`
	WorkingOfficeName               string              `json:"WorkingOfficeName"`
	WorkingOfficePincode            int32               `json:"WorkingOfficePincode"`
	WorkingOfficeRegionFacilityID   string              `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string              `json:"WorkingOfficeRegionName" `
}

// UpdateExamCentersInMTSPMMGApplsreturnstring godoc
// @Summary Update exam centers in MTS to PM/MG applications and return status
// @Description Updates exam centers in MTS to PM/MG applications and returns the updated status
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param request body []ApplicationMTSPMUpdatecenter true "Array of MTS to PM/MG applications with updated exam centers"
// @Success 200 {object} start.GetApplicationsResponse "Updating Exam centers for MTSPM Applications"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/center/updatecenters [put]
func UpdateExamCentersInMTSPMMGApplsreturnstring(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateExamCentersInMTSPMMGApplsreturnstring "
		var startFunction string = " - start - UpdateExamCentresMTSPMMGExamsreturnarray "

		// Bind the JSON data to the newappls variable
		var req struct {
			Newappls []*ent.Exam_Application_MTSPMMG `json:"newappls"`
		}
		if err := gctx.ShouldBindJSON(&req); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		// Call the UpdateExamCentresIPExams function to update the exam centers
		examCenters, status, stgError, dataStatus, err := start.UpdateExamCentresMTSPMMGExamsreturnarray(ctx, client, req.Newappls)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		// Display the status
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Updating Exam centers for MTSPM Applications",
			"data":       examCenters,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

// LAST SECOND API

// GetMTSPMMGHallTicketWithExamCodeEmpID godoc
// @Summary Get MTSPMMG hall ticket details by exam code and employee ID
// @Description Retrieves MTSPMMG hall ticket details based on exam code, employee ID, and selected year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param examcode path int true "Exam code"
// @Param employeeid path int true "Employee ID"
// @Param selectedyear query string false "Selected year"
// @Success 200 {object} start.GetApplicationsResponse "Generate application with Hallticket Number"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/hallticket/get/{examcode}/{employeeid}/{selectedyear} [get]
func GetMTSPMMGHallTicketWithExamCodeEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		//ctx := context.Background()

		//:examname/:employeeid/:examyear
		var mainFunction string = " main - GetMTSPMMGHallTicketWithExamCodeEmpID "
		var startFunction string = " - start - GetMTSPMMGApplicationsWithHallTicket "

		ec := gctx.Param("examcode")
		ExamCode, _ := strconv.ParseInt(ec, 10, 32)
		EmployeeID, _ := strconv.ParseInt(gctx.Param("employeeid"), 10, 64)
		examYear := gctx.Query("selectedyear")
		examcenters, _, status, stgError, dataStatus, err := start.GetMTSPMMGApplicationsWithHallTicket(client, int32(ExamCode), EmployeeID, examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Generate application with Hallticket Number",
			"data":       examcenters,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)

}

type requestBody struct {
	ExamYear               string `json:"ExamYear"`
	ExamCode               int32  `json:"ExamCode"`
	NodalOfficerFacilityID string `json:"nodalOfficerFacilityID"`
}

// GenerateHallticketNumbersMtsPm godoc
// @Summary Generate hall ticket numbers for MTS to PM/MG applicants
// @Description Generates hall ticket numbers for MTS to PM/MG applicants based on exam year, exam code, and nodal officer facility ID
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param request body requestBody true "Request Body"
// @Success 200 {object} start.GetApplicationsResponse "Hall Ticket generation done successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/Halltickets [put]
func GenerateHallticketNumbersMtsPm(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GenerateHallticketNumbersMtsPm "
		var startFunction string = " - start - GenerateHallticketNumberrMtsPm "
		// Parse the request body into the struct
		var reqBody requestBody
		if err := gctx.BindJSON(&reqBody); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		// Call GenerateHallticketNumberrIP function
		successMessage, status, stgError, dataStatus, err := start.GenerateHallticketNumberrMtsPm(ctx, client, reqBody.ExamYear, reqBody.ExamCode, reqBody.NodalOfficerFacilityID)
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

// ALL CANDIDATES TO BE VERIFIED

// GetAllMTSPMMGPendingWithCandidate godoc
// @Summary Get all MTS to PM/MG applications pending with candidate
// @Description Retrieves all MTS to PM/MG applications that are pending with the candidate based on facility ID and selected year
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param facilityid path string true "Facility ID"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} start.GetApplicationsResponse "Fetching application pending with candidate"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/MTSPMMGexams/getAllPendingWithCandidate/{facilityid}/{selectedyear} [get]
func GetAllMTSPMMGPendingWithCandidate(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllMTSPMMGPendingWithCandidate "
		var startFunction string = " - start - QueryMTSPMMGApplicationsByPendingWithCandidate "
		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		id1 := gctx.Param("selectedyear")
		circles, status, stgError, dataStatus, err := start.QueryMTSPMMGApplicationsByPendingWithCandidate(ctx, client, facilityID, id1)
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

// GetExamApplicatonsPreferenenceCityWiseStatsMTSPMMG godoc
// @Summary Get exam application preference statistics city-wise for MTS to PM/MG exams
// @Description Retrieves statistics on exam application preferences city-wise based on exam year, exam code, and city ID
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param examYear path string true "Exam Year"
// @Param examcode path string true "Exam Code"
// @Param cityid path string true "City ID"
// @Success 200 {object} start.GetApplicationsResponse "Fetching application count based on Exam city "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/getExamApplicationsByCityPrefMTSPMMG/{examYear}/{examcode}/{cityid} [get]
func GetExamApplicatonsPreferenenceCityWiseStatsMTSPMMG(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetExamApplicatonsPreferenenceCityWiseStatsMTSPMMG "
		var startFunction string = " - start - GetExamApplicatonsPreferenenceCityWiseStatsMTSPMMG "

		//:examYear/:examcode/:cityid
		ExamYear := gctx.Param("examYear")
		Examcode := gctx.Param("examcode")
		Cityid := gctx.Param("cityid")

		var statistics []start.ExamStatsMTSPMMG
		var err error
		statistics, status, stgError, dataStatus, err := start.GetExamApplicatonsPreferenenceCityWiseStatsMTSPMMG(ctx, client, ExamYear, Examcode, Cityid)

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

// MIC 2

// type requestDataArrayyy []struct {
// 	CenterPreference    string `json:"centerPreference"`
// 	CenterCode          int    `json:"centerCode"`
// 	ReportingOfficeName string `json:"reportingOfficeName"`
// 	SeatsToAllot        int    `json:"seatsToAllot"`
// }

// UpdateCenterCodeForApplicationsMTSPMMG godoc
// @Summary Update center code for applications in MTSPMMG
// @Description Update center code for applications based on preferences, center code, reporting office name, and seats to allot
// @Tags MTSPM Applications
// @Accept json
// @Produce json
// @Param request body []requestDataArray true "Array of request data to update center codes"
// @Success 200 {object} start.GetApplicationsResponse "Fetching application count based on Exam city "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Applications not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/updateCenterCodeForApplicationsMTSPMMG [put]
func UpdateCenterCodeForApplicationsMTSPMMG(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - UpdateCenterCodeForApplicationsMTSPMMG "
		var startFunction string = " - start - UpdateCenterCodeForApplicationsMTSPMMG "

		// var requestDataArray requestDataArrayyy
		// if err := gctx.BindJSON(&requestDataArray); err != nil {
		var reqData requestDataArray

		if err := gctx.BindJSON(&reqData); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		var totalUpdatedCount int
		var updatedApplications []*ent.Exam_Application_MTSPMMG

		var smsStatus, emailSentStatus string
		var dataStatus bool

		for _, requestData := range reqData {
			count, updatedApps, status, stgError, currentDataStatus, err := start.UpdateCenterCodeForApplicationsMTSPMMG(ctx, client, requestData.ControllingOfficeFacilityID, int32(requestData.ExamCenterID), int32(requestData.SeatsToAllot), int32(requestData.ExamCityID))
			if err != nil {
				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			}
			totalUpdatedCount += count
			updatedApplications = append(updatedApplications, updatedApps...)
			dataStatus = currentDataStatus
			for _, app := range updatedApps {
				employeeIDStr := strconv.FormatInt(app.EmployeeID, 10)
				smsStatus = sms.SendSmsNew(ctx, client, employeeIDStr, 12, app.ApplicationNumber, app.HallTicketNumber, "LDCE for MTS to PM/MG")
				var submittedDate string = app.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
				var hallticketDate string = app.HallTicketGeneratedDate.Format("02-01-2006 15:04:05")
				emailSentStatus = mail.SendEMailNew(ctx, client, employeeIDStr, 12, "LDCE for MTS to PM/MG", app.ApplicationNumber, app.ControllingOfficeName, submittedDate, app.NodalOfficeName, app.HallTicketNumber, hallticketDate)
			}
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":             true,
			"message":             fmt.Sprintf("Total Number of Applications %d Updated with center code Successfully", totalUpdatedCount),
			"data":                struct{}{},
			"dataexists":          dataStatus,
			"updatedApplications": updatedApplications,
			"SMSStatus":           smsStatus,
			"EmailStatus":         emailSentStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

func GetMTSPMMGAllNAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetMTSPMMGAllNAVerified "
		var startFunction string = " - start - QueryMTSPMMGApplicationsByNAVerified "
		//:facilityid/:selectedyear
		facilityID := gctx.Param("facilityid")
		selectedYear := gctx.Param("selectedyear")

		circles, status, stgError, dataStatus, err := start.QueryMTSPMMGApplicationsByNAVerified(ctx, client, facilityID, selectedYear)
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
