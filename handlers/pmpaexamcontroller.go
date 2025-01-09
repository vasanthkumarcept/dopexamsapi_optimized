package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"recruit/ent"
	"recruit/mail"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/sms"
	"recruit/start"
	"recruit/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNewPMPAApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - CreateNewPMPAApplications "
		var startFunction string = " - start - CreatePMPAApplications "
		var newPAAppplns ca_reg.ApplicationPMPA

		if err := gctx.ShouldBindJSON(&newPAAppplns); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newPAAppplns.Edges.LogData[0]

		newPAApppln, status, stgError, dataStatus, err := start.CreatePMPAApplications(client, &newPAAppplns)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		var Smsuserid string = fmt.Sprint(newPAApppln.EmployeeID)
		var submittedDate string = newPAApppln.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		//examshortname ,appno ,ca submittedDate
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 3, " LDCE MTS/PM/MG to PA/SA ", newPAApppln.ApplicationNumber, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE MTS/PM/MG to PA/SA", newPAApppln.ApplicationNumber, "Submitted")

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "PM to PA Application submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        newPAApppln,
			"dataexists":  dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

//Update Verification details

func VerifyPMPAApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - VerifyPMPAApplication "
		var startFunction string = " - start - UpdateApplicationRemarksPMPA "
		var newAppln ca_reg.VerifyApplicationPMPA
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

		application, status, stgError, dataStatus, err := start.UpdateApplicationRemarksPMPA(client, &newAppln, nonQualifyService)
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
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE MTS/PM/MG to PA/SA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.CAGeneralRemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE MTS/PM/MG to PA/SA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.CAGeneralRemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByCA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE MTS/PM/MG to PA/SA", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE MTS/PM/MG to PA/SA", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "PM to PA exam application verified successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

func ResubmitPMPAApplication(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - ResubmitPMPAApplication "
		var startFunction string = " - start - ResubmitApplicationRemarksPMPA "

		var newAppln ca_reg.ReApplicationPMPA

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newAppln.Edges.LogData[0]

		application, status, stgError, dataStatus, err := start.ResubmitApplicationRemarksPMPA(client, &newAppln)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		//examshortname ,	appno, ca, submittedDate
		var submittedDate string = application.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
		var Smsuserid = fmt.Sprint(newAppln.EmployeeID)
		emailsentstatus := mail.SendEMailNew(ctx, client, Smsuserid, 7, " LDCE MTS/PM/MG to PA/SA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate)

		smssentstatus := sms.SendSmsNew(ctx, client, Smsuserid, 3, "LDCE MTS/PM/MG to PA/SA", application.ApplicationNumber, "Re-submitted")

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "PM to PA application re-submitted successfully",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

// get ca pending verify records
func GetAllPMPACAPendingVerifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllPMPACAPendingVerifications "
		var startFunction string = " - start - QueryPMPAApplicationsByCAVerificationsPending "
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		id1 := gctx.Param("id1")

		circles, status, stgError, dataStatus, err := start.QueryPMPAApplicationsByCAVerificationsPending(ctx, client, facilityID, id1)
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
func GetPMPACAPendingDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryPMPAApplicationsByCAPendingByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// Return Previous Remarks
func GetPMPACAPendingOldRemarksByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//newAppln := new(ent.Exam_Applications_PMPA)
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.GetOldPMPACAApplicationRemarksByEmployeeID(ctx, client, int64(empid))
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
func GetAllPMPACAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllPMPACAVerified "
		var startFunction string = " - start - QueryPMPAApplicationsByCAVerified "
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		id1 := gctx.Param("id1")

		circles, status, stgError, dataStatus, err := start.QueryPMPAApplicationsByCAVerified(ctx, client, facilityID, id1)
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
func GetPMPACAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryPMPAApplicationsByCAVerifiedByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// Get all CA verified for NA..
func GetPMPAAllCAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		facilityID1 := gctx.Param("id1")

		circles, err := start.QueryPMPAApplicationsByCAVerifiedForNA(ctx, client, facilityID, facilityID1)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)

}

// Update Nodal recommendations

func UpdateNodalRecommendationsPMPAByEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - UpdateNodalRecommendationsPMPAByEmpID "
		var startFunction string = " - start - UpdatePMPANodalRecommendationsByEmpID "
		var newAppln ca_reg.NAVerifyApplicationPMPA

		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newAppln.Edges.LogData[0]
		fmt.Println(logdata)

		application, status, stgError, dataStatus, err := start.UpdatePMPANodalRecommendationsByEmpID(client, &newAppln)
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
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 4, " LDCE MTS/PM/MG to PA/SA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.RecommendedStatus, application.NARemarks)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			//examshortname ,	appno, ca, submittedDate, comments
			emailsentstatus = mail.SendEMailNew(ctx, client, Smsuserid, 6, " LDCE MTS/PM/MG to PA/SA ", application.ApplicationNumber, application.ControllingOfficeName, submittedDate, application.NARemarks)
		}

		var smssentstatus string = ""
		if application.ApplicationStatus == "VerifiedByNA" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE MTS/PM/MG to PA/SA", application.ApplicationNumber, application.RecommendedStatus)
		} else if application.ApplicationStatus == "PendingWithCandidate" {
			smssentstatus = sms.SendSmsNew(ctx, client, Smsuserid, 4, "LDCE MTS/PM/MG to PA/SA", application.ApplicationNumber, "Returned for correction")
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Nodal officer successfully verified this PMMG to PA  application ",
			"EmailStatus": emailsentstatus,
			"SMSStatus":   smssentstatus,
			"data":        application,
			"dataexists":  dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetRecommendationsByEmpID
func GetPMPAExamRecommendationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		records, err := start.GetPMPARecommendationsByEmpID(client, empid)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": records})
	}
	return gin.HandlerFunc(fn)
}

// Get All NA Verified records by CA ...
func GetPMPAAllNAVerified(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetPMPAAllNAVerified "
		var startFunction string = " - start - QueryPMPAApplicationsByNAVerified "
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		id1 := gctx.Param("id1")

		circles, status, stgError, dataStatus, err := start.QueryPMPAApplicationsByNAVerified(ctx, client, facilityID, id1)
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
func GetPMPANAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryPMPAApplicationsByNAVerifiedByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}

// Get all NA Verified for NA ..
func GetPMPAAllNAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		facilityID1 := gctx.Param("id1")

		circles, err := start.QueryPMPAApplicationsByNAVerifiedForNA(ctx, client, facilityID, facilityID1)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)

}
func UpdateExamCentersInPMPAApplsreturnstring(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		// Bind the JSON data to the newappls variable
		var req struct {
			Newappls []*ent.Exam_Applications_PMPA `json:"newappls"`
		}
		if err := gctx.ShouldBindJSON(&req); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the UpdateExamCentresIPExams function to update the exam centers
		status, err := start.UpdateExamCentresPMPAExamsreturnarray(ctx, client, req.Newappls)
		if err != nil {
			gctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		// Display the status
		gctx.JSON(http.StatusOK, gin.H{"data": status})
	}

	return gin.HandlerFunc(fn)
}
func GetPMPAHallticketNumberscenter(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		//ec := gctx.Param("id")
		//CenterCode, _ := strconv.ParseInt(ec, 10, 32)
		// Call the GenerateHallticketNumberIP function to generate hall ticket numbers and get the success message
		successMessage, err := start.GenerateHallticketNumberPMPAwithCenterCode(ctx, client)
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

func GetPMPAHallTicketWithExamCodeEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		//ctx := context.Background()

		ec := gctx.Param("id1")
		ExamCode, _ := strconv.ParseInt(ec, 10, 32)
		EmployeeID, _ := strconv.ParseInt(gctx.Param("id2"), 10, 64)

		_, examcenters, err := start.GetPMPAApplicationsWithHallTicket(client, int32(ExamCode), EmployeeID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": examcenters})

	}

	return gin.HandlerFunc(fn)
}
func GetAllPMPAPendingWithCandidate(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetAllPMPAPendingWithCandidate "
		var startFunction string = " - start - QueryPMPAApplicationsByPendingWithCandidate "
		facilityID := gctx.Param("id")
		id1 := gctx.Param("id1")

		circles, status, stgError, dataStatus, err := start.QueryPMPAApplicationsByPendingWithCandidate(ctx, client, facilityID, id1)
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
func GenerateHallticketNumbersPmPa(client *ent.Client) gin.HandlerFunc {
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
		successMessage, err := start.GenerateHallticketNumberrPmPa(ctx, client, reqBody.ExamYear, reqBody.ExamCode, reqBody.NodalOfficerFacilityID)
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

func GetExamApplicatonsPreferenenceCityWiseStatsPMPA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetExamApplicatonsPreferenenceCityWiseStatsPMPA "
		var startFunction string = " - start - GetExamApplicatonsPreferenenceCityWiseStatsPMPA "

		//:examYear/:examcode/:cityid
		ExamYear := gctx.Param("examYear")
		Examcode := gctx.Param("examcode")
		Cityid := gctx.Param("cityid")

		var statistics []start.ExamStatsPMPA
		var err error
		statistics, status, stgError, dataStatus, err := start.GetExamApplicatonsPreferenenceCityWiseStatsPMPA(ctx, client, ExamYear, Examcode, Cityid)
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

func UpdateCenterCodeForApplicationsPMPA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
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
		var updatedApplications []*ent.Exam_Applications_PMPA

		var smsStatus, emailSentStatus string
		var dataStatus bool
		for _, requestData := range reqData {
			count, updatedApps, status, stgError, currentDataStatus, err := start.UpdateCenterCodeForApplicationsPMPA(ctx, client, requestData.ControllingOfficeFacilityID, int32(requestData.ExamCenterID), int32(requestData.SeatsToAllot), int32(requestData.ExamCityID))
			if err != nil {
				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			}
			totalUpdatedCount += count
			updatedApplications = append(updatedApplications, updatedApps...)
			dataStatus = currentDataStatus
			for _, app := range updatedApps {
				employeeIDStr := strconv.FormatInt(app.EmployeeID, 10)
				smsStatus = sms.SendSmsNew(ctx, client, employeeIDStr, 12, app.ApplicationNumber, app.HallTicketNumber, "LDCE MTS/PM/MG to PA/SA")
				var submittedDate string = app.ApplnSubmittedDate.Format("02-01-2006 15:04:05")
				var hallticketDate string = app.HallTicketGeneratedDate.Format("02-01-2006 15:04:05")
				emailSentStatus = mail.SendEMailNew(ctx, client, employeeIDStr, 12, "LDCE MTS/PM/MG to PA/SA", app.ApplicationNumber, app.ControllingOfficeName, submittedDate, app.NodalOfficeName, app.HallTicketNumber, hallticketDate)
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

func GetPMPAApplicationsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		id1 := gctx.Param("id1")

		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		paapplns, err := start.QueryPMPAExamApplicationsByEmpID(ctx, client, int64(empid), string(id1))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": paapplns})
	}
	return gin.HandlerFunc(fn)
}

func UpdateExamCentersInPMPAAppls(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		// Bind the JSON data to the newappls variable
		var req struct {
			Newappls []*ent.Exam_Applications_PMPA `json:"newappls"`
		}
		if err := gctx.ShouldBindJSON(&req); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the UpdateExamCentresPMPAExams function to update the exam centers
		updatedRecords, err := start.UpdateExamCentresPMPAExams(ctx, client, req.Newappls)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Display the updated records
		gctx.JSON(http.StatusOK, gin.H{"data": updatedRecords})
	}

	return gin.HandlerFunc(fn)
}
