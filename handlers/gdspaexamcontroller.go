package handlers

import (
	"context"
	"fmt"
	"log"

	//"os"
	"strconv"

	"net/http"
	"recruit/ent"
	"recruit/mail"
	"recruit/sms"
	"recruit/start"
	"recruit/util"

	ca_reg "recruit/payloadstructure/candidate_registration"

	"github.com/gin-gonic/gin"
)

func CreateNewGDSPAApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - CreateNewGDSPAApplications "
		var startFunction string = " - start - CreateGDSPAApplications "
		var newPAAppplns ca_reg.ApplicationGDSPM
		if err := gctx.ShouldBindJSON(&newPAAppplns); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newPAAppplns.Edges.LogData[0]

		newPAApppln, status, stgError, dataStatus, err := start.CreateGDSPAApplications(client, &newPAAppplns)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		var employeeIDSMS string = fmt.Sprint(newPAApppln.EmployeeID)
		var submittedDate string = newPAApppln.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		//examshortname ,appno ,ca submittedDate
		emailsentstatus := mail.SendEMailNew(ctx, client, employeeIDSMS, 3, " CE GDS to PA ", newPAApppln.ApplicationNumber, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, employeeIDSMS, 3, " CE GDS to PA ", newPAApppln.ApplicationNumber, submittedDate)

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "GDS to PA Application submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        newPAApppln,
			"dataexists":  dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}
func ResubmitGDSPAApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var newAppln ca_reg.ReApplicationGDStoPA
		var mainFunction string = " main - ResubmitGDSPAApplication "
		var startFunction string = " - start - ResubmitApplicationRemarksGDSPA "

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newAppln.Edges.LogData[0]

		application, status, stgError, dataStatus, err := start.ResubmitApplicationRemarksGDSPA(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		//examshortname ,	appno, ca, submittedDate
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 7, " CE GDS to PA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "CE GDS to PA", application.ApplicationNumber, "Re-submitted")

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "GDS to PA application re-submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

//Update Verification details

func VerifyGDSPAApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - VerifyGDSPAApplication "
		var startFunction string = " - start - UpdateApplicationRemarksGDSPA "
		var newAppln ca_reg.VerifyApplicationGDStoPA
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
		application, status, stgError, dataStatus, err := start.UpdateApplicationRemarksGDSPA(client, &newAppln, nonQualifyService)
		logdata := newAppln.Edges.LogData[0]
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		var emailsentstatus string
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		if application.ApplicationStatus == "VerifiedByCA" {
			//examshortname ,appno ,ca,submittedDate ,recommended, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " CE GDS to PA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.CAGeneralRemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " CE GDS to PA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.CAGeneralRemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByCA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "CE GDS to PA", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "CE GDS to PA", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "GDS to PA exam application verified successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

/* // get GDSPM Sub Division pending verification
func GetAllGDSPMVAPendingVerifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		facilityID := gctx.Param("id")
		id1 := gctx.Param("id1")

		circles, status, err := start.QueryGDSPMApplicationsByVAVerificationsPending(ctx, client, facilityID, id1)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "main - QueryGDSPMApplicationsByVAVerificationsPending - ShouldBindJSON error" + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": os.Getenv("USER_ERROR_REMARKS")})
				return
			} else if status == 422 {
				Action = "422"
				Remarks = "main - QueryGDSPMApplicationsByVAVerificationsPending  user error" + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			} else {
				Action = "500"
				Remarks = "main - QueryGDSPMApplicationsByVAVerificationsPending other errors" + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": os.Getenv("USER_ERROR_REMARKS")})
				return
			}
		} else {
			gctx.JSON(http.StatusOK, gin.H{"data": circles})
		}
	}
	return gin.HandlerFunc(fn)
}
*/
// get ca pending verify records
func GetAllGDSPACAPendingVerifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllPMPACAPendingVerifications "
		var startFunction string = " - start - QueryPMPAApplicationsByCAVerificationsPending "
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPAApplicationsByCAVerificationsPending(ctx, client, facilityID, examYear)
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
func GetGDSPACAPendingDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("employeeid")
		examYear := gctx.Param("examyear")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryGDSPAApplicationsByCAPendingByEmpID(ctx, client, int64(empid), examYear)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// Return Previous Remarks
func GetGDSPACAPendingOldRemarksByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//newAppln := new(ent.Exam_Applications_GDSPA)
		id := gctx.Param("empid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.GetOldGDSPACAApplicationRemarksByEmployeeID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//gctx.JSON(http.StatusOK, gin.H{"data": circles})

		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"EmployeeID":          circles.EmployeeID,
				"ApplicationStatus":   circles.ApplicationStatus,
				"Application Remarks": circles.AppliactionRemarks,
				//	"CAOldRemarks":        application.CAPreviousRemarks,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// Get All CA Verified records
func GetAllGDSPACAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllPSCAVerified "
		var startFunction string = " - start - QueryPSApplicationsByCAVerified "
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPAApplicationsByCAVerified(ctx, client, facilityID, examYear)
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
func GetGDSPACAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("employeeid")
		examyear := gctx.Param("examyear")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryGDSPAApplicationsByCAVerifiedByEmpID(ctx, client, int64(empid), examyear)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// Get all CA verified for NA..
func GetGDSPAAllCAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, err := start.QueryGDSPAApplicationsByCAVerifiedForNA(ctx, client, facilityID, examYear)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)

}
func UpdateNodalRecommendationsGDSPAByEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var newAppln ca_reg.NAVerifyApplicationGDStoPA

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - UpdateNodalRecommendationsGDSPAByEmpID "
		var startFunction string = " - start - UpdateGDSPANodalRecommendationsByEmpID "

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		if len(newAppln.Edges.LogData) <= 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "NO LOG DATA"})
			return
		}
		logdata := newAppln.Edges.LogData[0]

		application, status, stgError, dataStatus, err := start.UpdateGDSPANodalRecommendationsByEmpID(client, &newAppln)
		//logdata := newAppln.Edges.LogData[0]
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var emailsentstatus string
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		if application.ApplicationStatus == "VerifiedByNA" {
			//examshortname ,appno ,ca,submittedDate ,recommended, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " CE GDS to PA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.NARemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " CE GDS to PA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.NARemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByNA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "CE GDS to PA", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "CE GDS to PA", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Nodal officer successfully verified this GDSPA application ",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// GetRecommendationsByEmpID
func GetGDSPAExamRecommendationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		id := gctx.Param("empid")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		records, err := start.GetGDSPARecommendationsByEmpID(client, empid)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": records})
	}
	return gin.HandlerFunc(fn)
}

// Get All NA Verified records by CA ...
func GetGDSPAAllNAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetGDSPAAllNAVerified "
		var startFunction string = " - start - QueryGDSPAApplicationsByNAVerified "
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPAApplicationsByNAVerified(ctx, client, facilityID, examYear)
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
func GetGDSPANAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("empid")
		employeeID, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryGDSPAApplicationsByNAVerifiedByEmpID(ctx, client, int64(employeeID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// Get all NA Verified for NA ..
func GetGDSPAAllNAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, err := start.QueryGDSPAApplicationsByNAVerifiedForNA(ctx, client, facilityID, examYear)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)

}
func UpdateExamCentersInGDSPAApplsreturnstring(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		// Bind the JSON data to the newappls variable
		var req struct {
			Newappls []*ent.Exam_Applications_GDSPA `json:"newappls"`
		}
		if err := gctx.ShouldBindJSON(&req); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the UpdateExamCentresIPExams function to update the exam centers
		status, err := start.UpdateExamCentresGDSPAExamsreturnarray(ctx, client, req.Newappls)
		if err != nil {
			gctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		// Display the status
		gctx.JSON(http.StatusOK, gin.H{"data": status})
	}

	return gin.HandlerFunc(fn)
}

/*
	 func GetGDSPAHallticketNumberscenter(client *ent.Client) gin.HandlerFunc {
		fn := func(gctx *gin.Context) {
			ctx := context.Background()
			//ec := gctx.Param("id")
			//CenterCode, _ := strconv.ParseInt(ec, 10, 32)
			// Call the GenerateHallticketNumberIP function to generate hall ticket numbers and get the success message
			successMessage, err := start.GenerateHallticketNumberGDSPAwithCenterCode(ctx, client)
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
*/
func GetGDSPAHallTicketWithExamCodeEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		//ctx := context.Background()
		ec := gctx.Param("examcode")
		ExamCode, _ := strconv.ParseInt(ec, 10, 32)
		EmployeeID, _ := strconv.ParseInt(gctx.Param("empid"), 10, 64)

		_, examcenters, err := start.GetGDSPAApplicationsWithHallTicket(client, int32(ExamCode), EmployeeID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": examcenters})

	}

	return gin.HandlerFunc(fn)
}
func GetAllGDSPAPendingWithCandidate(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllGDSPAPendingWithCandidate "
		var startFunction string = " - start - QueryGDSPAApplicationsByPendingWithCandidate "
		facilityID := gctx.Param("facilityid")
		examYear := gctx.Param("examyear")

		circles, status, stgError, dataStatus, err := start.QueryGDSPAApplicationsByPendingWithCandidate(ctx, client, facilityID, examYear)
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
func GenerateHallticketNumbersGdspa(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := gctx.Request.Context()

		//ctx := gctx.Request.Context()

		// Define a struct to represent the request body
		type requestBody struct {
			ExamYear               string `json:"ExamYear"`
			ExamCode               int32  `json:"ExamCode"`
			NodalOfficerFacilityID string `json:"nodalOfficerFacilityID"`
		}

		// Parse the request body into the struct
		var reqBody requestBody
		if err := gctx.BindJSON(&reqBody); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		fmt.Println(reqBody.ExamYear)

		// Call GenerateHallticketNumberrIP function
		successMessage, err := start.GenerateHallticketNumberrGDSPA(ctx, client, reqBody.ExamYear, reqBody.ExamCode, reqBody.NodalOfficerFacilityID)
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

func UpdateCenterCodeForApplicationsGDSPA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateCenterCodeForApplicationsGDSPA "
		var startFunction string = " - start - UpdateCenterCodeForApplicationsGDSPA "

		var reqData requestDataArray

		if err := gctx.BindJSON(&reqData); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		var totalUpdatedCount int
		var updatedApplications []*ent.Exam_Applications_GDSPA
		var smsStatus, emailSentStatus string

		var dataStatus bool

		for _, requestData := range reqData {
			count, updatedApps, status, stgError, currentDataStatus, err := start.UpdateCenterCodeForApplicationsGDSPA(ctx, client, requestData.ControllingOfficeFacilityID, int32(requestData.ExamCenterID), int32(requestData.SeatsToAllot), int32(requestData.ExamCityID))
			if err != nil {
				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			}
			totalUpdatedCount += count
			updatedApplications = append(updatedApplications, updatedApps...)
			dataStatus = currentDataStatus
			for _, app := range updatedApps {
				employeeIDStr := strconv.FormatInt(app.EmployeeID, 10)
				smsStatus = sms.SendSmsNew(ctx, client, employeeIDStr, 12, app.ApplicationNumber, app.HallTicketNumber, " CE GDS to PA ")
				var submittedDate string = app.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
				var hallticketDate string = app.HallTicketGeneratedDate.Format("02-01-2006 15:04:05")
				emailSentStatus = mail.SendEMailNew(ctx, client, employeeIDStr, 12, " CE GDS to PA ", app.ApplicationNumber, app.ControllingOfficeName, submittedDate, app.NodalOfficeName, app.HallTicketNumber, hallticketDate)
			}
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":             true,
			"message":             fmt.Sprintf("Total Number of Applications %d Updated with center code Successfully", totalUpdatedCount),
			"data":                struct{}{},
			"updatedApplications": updatedApplications,
			"dataexists":          dataStatus,
			"SMSStatus":           smsStatus,
			"EmailStatus":         emailSentStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

func GetExamApplicatonsPreferenenceCityWiseStatsGDSPA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetExamApplicatonsPreferenenceCityWiseStatsGDSPA "
		var startFunction string = " - start - GetExamApplicatonsPreferenenceCityWiseStatsGDSPA "
		//:examYear/:examcode/:cityid
		ExamYear := gctx.Param("examYear")
		Examcode := gctx.Param("examcode")
		Cityid := gctx.Param("cityid")

		/* 		// Convert strings to int64
		   		Examcode64, err1 := strconv.ParseInt(examStringCode, 10, 32)
		   		if err1 != nil {
		   			log.Printf("Error parsing Examcode: %v", err1)
		   			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Examcode"})
		   			return
		   		}

		   		Cityid64, err2 := strconv.ParseInt(cityStringId, 10, 32)
		   		if err2 != nil {
		   			log.Printf("Error parsing Cityid: %v", err2)
		   			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Cityid"})
		   			return
		   		}

		   		// Convert int64 to int32
		   		Examcode := int32(Examcode64)
		   		Cityid := int32(Cityid64) */

		var statistics []start.ExamStatsGDSPA
		var err error
		statistics, status, stgError, dataStatus, err := start.GetExamApplicatonsPreferenenceCityWiseStatsGDSPA(ctx, client, ExamYear, Examcode, Cityid)

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

// Get IP Applications with EmpID - QueryIPExamApplicationsByEmpID
func GetGDSPAApplicationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		id := gctx.Param("empid")
		examYear := gctx.Param("examyear")

		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		paapplns, err := start.QueryGDSPAExamApplicationsByEmpID(ctx, client, int64(empid), string(examYear))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": paapplns})
	}
	return gin.HandlerFunc(fn)
}
