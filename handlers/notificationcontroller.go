package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"recruit/ent"
	"recruit/ent/examnotifications"
	"recruit/ent/facilitymasters"
	"recruit/ent/usermaster"
	"recruit/start"
	"recruit/util"
	"strconv"
	"time"

	ca_reg "recruit/payloadstructure/candidate_registration"

	"github.com/gin-gonic/gin"
)

// CreateexamNotification godoc
// @Summary      Create Exam Notification
// @Description  Create a new exam notification
// @Tags         ExamNotifications
// @Accept       json
// @Produce      json
// @Param        notification  body      ca_reg.CreateNotificationStruct  true  "Exam Notification"
// @Success 200 {object} start.EmployeeMasterResponse "Successfully created the Notification"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /deptexam/notification/create [post]
func CreateexamNotification(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//newNotification := new(ent.ExamNotifications)
		var newNotification ca_reg.CreateNotificationStruct

		var mainFunction string = " main - CreateexamNotification "
		var startFunction string = " - start - CreateExamNotification "

		if err := gctx.ShouldBindJSON(&newNotification); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newNotification.Edges.LogData[0]

		_, status, stgError, dataStatus, err := start.CreateExamNotification(client, &newNotification)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)

			return
		}
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Successfully created the Notification",
			"data":       "Successfully created the Notification",
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// Function to get Draft notification based on exam Code and exam year

// GetDraftNotification godoc
// @Summary      Get Draft Notification
// @Description  Fetch notification based on Exam code, Exam year, Office type, and facility ID
// @Tags         ExamNotifications
// @Accept       json
// @Produce      json
// @Param        examid     path    int     true  "Exam ID"
// @Param        examyear   path    int     true  "Exam Year"
// @Param        officetype path    string  true  "Office Type"
// @Param        facilityid path    string  true  "Facility ID"
// @Success 200 {object} start.EmployeeMasterResponse "Fetched notification based on Exam code, Exam year, Office type and facility ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /deptexam/notification/GetDraftNotification/{examid}/{examyear}/{officetype}/{facilityid} [get]
func GetDraftNotification(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var mainFunction string = " main - GetDraftNotification "
		var startFunction string = " - start - GetDraftNotification "
		//:examid/:examyear/:officetype/:facilityid
		intExamId, _ := strconv.ParseInt(gctx.Param("examid"), 10, 32)
		intExamYear, _ := strconv.ParseInt(gctx.Param("examyear"), 10, 32)
		officeType := gctx.Param("officetype")
		facilityId := gctx.Param("facilityid")
		getNotifications, status, stgError, dataStatus, err := start.GetDraftNotification(client, int32(intExamId), int32(intExamYear), officeType, facilityId)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Fetched notification based on Exam code, Exam year, Office type and facility ID",
			"data":       getNotifications,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// function to get all draft notification by approver name
func GetDraftexamNotificationBySuUserName(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		un := gctx.Param("id")

		// Check if the username exists in UserMaster
		exists, err := client.UserMaster.Query().Where(usermaster.UserName(un)).Exist(gctx)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking username existence"})
			return
		}

		if !exists {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
			return
		}

		getNotification := new(ent.ExamNotifications)
		fmt.Println("The username provided is : ", un)

		getNotifications, err := start.GetexamNotification(client, un, getNotification)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(getNotifications) <= 0 {
			gctx.JSON(http.StatusOK, gin.H{"Message": "No notification pending for Approval"})

		} else {
			gctx.JSON(http.StatusOK, gin.H{"data": getNotifications,
				"Message": "The pending Notifications"})
		}
	}
	return gin.HandlerFunc(fn)
}

func GetReDraftexamNotificationBySuUserName(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		un := gctx.Param("id")

		// Check if the username exists in UserMaster
		exists, err := client.UserMaster.Query().Where(usermaster.UserName(un)).Exist(gctx)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking username existence"})
			return
		}

		if !exists {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
			return
		}

		getNotification := new(ent.ExamNotifications)
		fmt.Println("The username provided is : ", un)

		getNotifications, err := start.GetReexamNotification(client, un, getNotification)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(getNotifications) <= 0 {
			gctx.JSON(http.StatusOK, gin.H{"Message": "No notification pending for Approval"})

		} else {
			gctx.JSON(http.StatusOK, gin.H{"data": getNotifications,
				"Message": "The pending resubmit Notifications"})
		}
	}
	return gin.HandlerFunc(fn)
}

// function to get individual draft notification by approver name using notification ID
func GetIndividualDraftNotificationbySuUserName(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		nid, _ := strconv.ParseInt(gctx.Param("id"), 10, 32)

		getNotifications, err := start.GetIndividualDraftNotificationbySuUserName(client, int32(nid))

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{"data": getNotifications,
			"Message": "The Notification deatils for given ID"})
	}
	return gin.HandlerFunc(fn)
}

func GetRemarksofNotificationByNotificationID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		nid, _ := strconv.ParseInt(gctx.Param("id"), 10, 32)

		getNotifications, err := start.GetRemarksofNotificationByNotificationID(ctx, client, int32(nid))

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		notificationRemarks := make([]string, 0)
		for _, notification := range getNotifications {
			notificationRemarks = append(notificationRemarks, notification.NotificationRemarks)
		}
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"Notification Remarks": notificationRemarks,
			},
		})
	}
	return gin.HandlerFunc(fn)
}

// function to get all notification issued based on exam id and exam year
func GetexamNotificationbyYear(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		eyear64, _ := strconv.ParseInt(gctx.Param("id1"), 10, 32)
		eid64, _ := strconv.ParseInt(gctx.Param("id2"), 10, 32)
		year := int32(eyear64)
		eid := int32(eid64)

		getNotifications, err := start.GetexamNotificationByYear(client, year, eid)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(getNotifications) <= 0 {
			gctx.JSON(http.StatusOK, gin.H{"Message": "No Notifications Issued"})

		} else {
			gctx.JSON(http.StatusOK, gin.H{"data": getNotifications,
				"Message": "The List of Notifications for the Year and Exam ID:"})
		}
	}
	return gin.HandlerFunc(fn)
}

// function to get all notification issued based on approver name
func GetallIssuedexamNotificationBySuUserName(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		un := gctx.Param("id")

		getNotifications, err := start.GetallIssuedexamNotificationBySuUserName(client, un)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(getNotifications) <= 0 {
			gctx.JSON(http.StatusOK, gin.H{"Message": "No Notifications Issued by the User"})

		} else {
			gctx.JSON(http.StatusOK, gin.H{"data": getNotifications,
				"Message": "The List of all issued Notifications by the User:"})
		}
	}
	return gin.HandlerFunc(fn)
}

// function to get Individual notification issued based on notification number
func GetIndividualIssuedexamNotificationByNotificationNumber(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		notinum := gctx.Param("id")

		getNotifications, err := start.GetIndividualIssuedexamNotificationByNotificationNumber(client, notinum)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(getNotifications) <= 0 {
			gctx.JSON(http.StatusOK, gin.H{"Message": "No Notifications Issued by the User"})

		} else {
			gctx.JSON(http.StatusOK, gin.H{"data": getNotifications,
				"Message": "The Issued Notification Details for the notification Number :"})
		}
	}
	return gin.HandlerFunc(fn)
}

// function to get all notification issued for a particular year
func GetAllIssuedexamNotificationbyYear(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		eyear, _ := strconv.ParseInt(gctx.Param("id"), 10, 32)

		getNotifications, err := start.GetAllIssuedexamNotificationbyYear(client, int32(eyear))

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(getNotifications) <= 0 {
			gctx.JSON(http.StatusOK, gin.H{"Message": "No Notifications Issued"})

		} else {
			gctx.JSON(http.StatusOK, gin.H{"data": getNotifications,
				"Message": "The List of Notifications for the Year:"})
		}
	}
	return gin.HandlerFunc(fn)
}

func getDistinctCircleFacilities(client *ent.Client) ([]*ent.FacilityMasters, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Define a struct to hold the distinct fields
	var circleFacilities []struct {
		CircleFacilityID   string `json:"circle_facility_id"`
		CircleFacilityName string `json:"circle_facility_name"`
	}

	// Perform the distinct query using Ent
	err := client.FacilityMasters.Query().
		Where(
			facilitymasters.CircleFacilityIDNEQ("AS71000000141"),
			facilitymasters.CircleFacilityIDNEQ("CR34000000000"),
			facilitymasters.CircleFacilityIDNEQ("DT01000000000"),
			facilitymasters.CircleFacilityIDNEQ(" "),
		).
		GroupBy(facilitymasters.FieldCircleFacilityID, facilitymasters.FieldCircleFacilityName).
		Scan(ctx, &circleFacilities)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve distinct circle facilities: %w", err)
	}

	// Convert the result into the expected return type
	var results []*ent.FacilityMasters
	for _, cf := range circleFacilities {
		results = append(results, &ent.FacilityMasters{
			CircleFacilityID:   cf.CircleFacilityID,
			CircleFacilityName: cf.CircleFacilityName,
		})
	}

	return results, nil
}

func PutIssueNotification(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		var issueNotification ca_reg.IssueNotificationStruct

		var mainFunction string = " main - PutIssueNotification "
		var startFunction string = " - start - IssueExamNotification "
		if err := gctx.ShouldBindJSON(issueNotification); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := issueNotification.Edges.LogData[0]

		u := util.AdminDatas(client, logdata.Userid)
		if u == nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID No Admin Data Found"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		result, status, stgError, dataStatus, err := start.IssueExamNotification(ctx, client, &issueNotification, u)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Nodal officer successfully verified this IP application ",
			"data":       result,
			"dataexists": dataStatus,
		})

	}

}

// PutIssueNotificationSingle godoc
// @Summary      Issue a Single Notification
// @Description  Issue a single exam notification
// @Tags         ExamNotifications
// @Accept       json
// @Produce      json
// @Param        notification  body      ca_reg.IssueNotificationStruct  true  "Issue Notification"
// @Success 200 {object} start.EmployeeMasterResponse "Notification issued successfully "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /deptexam/notification/PutIssueNotificationSingle [put]
func PutIssueNotificationSingle(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		var issueNotification ca_reg.IssueNotificationStruct

		var mainFunction string = " main - PutIssueNotificationSingle "
		var startFunction string = " - start - IssueExamNotificationSingle "

		if err := gctx.ShouldBindJSON(&issueNotification); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := issueNotification.Edges.LogData[0]

		/* 		var EmployeeName string = "xyz"
		   		var ID int64 = 1
		*/ // Issue the notification

		issuenot, status, stgError, dataStatus, err := start.IssueExamNotificationSingle(client, issueNotification)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    " Notification issued successfully ",
			"data":       issuenot,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// modified
func PutReIssueNotification(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		reissueNotification := new(ent.ExamNotifications)

		// Parse the JSON body into the reissueNotification object
		if err := gctx.ShouldBindJSON(reissueNotification); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the log data
		if len(reissueNotification.Edges.LogData) <= 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
			return
		}

		logdata := reissueNotification.Edges.LogData[0]

		// Fetch admin data using the user ID from log data
		u := util.AdminDatas(client, logdata.Userid)
		if u == nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID. No Admin Data Found"})
			return
		}

		// Use the context from the gin request
		ctx := gctx.Request.Context()

		// Retrieve distinct circle facilities
		circleFacilities, err := getDistinctCircleFacilities(client)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Generate a new notification number
		notificationNumber, err := start.GenerateNotificationNumber(client, reissueNotification.ExamCode)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Query the existing active notifications for the given exam code and year
		oldAppln, err := client.ExamNotifications.
			Query().
			Where(
				examnotifications.ExamCodeEQ(reissueNotification.ExamCode),
				examnotifications.ExamYearEQ(reissueNotification.ExamYear),
				examnotifications.StatusEQ("active"),
			).
			Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				gctx.JSON(http.StatusNotFound, gin.H{"error": "No active notification found for this Exam"})
				return
			}
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error in fetching data from the Exam table: %v", err)})
			return
		}

		// Update the status of the existing active notification to inactive
		stat := "inactive_" + time.Now().Format("20060102150405")
		_, err = oldAppln.
			Update().
			SetStatus(stat).
			Save(ctx)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update notification: %v", err)})
			return
		}

		// Iterate over each circle facility and issue a notification
		for _, facility := range circleFacilities {
			// Create a new notification for each circle facility
			copyNotification := *reissueNotification
			copyNotification.CircleOfficeFacilityId = facility.CircleFacilityID
			copyNotification.CircleOfficeName = facility.CircleFacilityName
			copyNotification.NotificationNumber = notificationNumber

			reissueNotifications, err := start.PutReIssueNotification(client, &copyNotification)
			if err != nil {
				util.LogError(client, logdata, err)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			_, logerr3 := client.Logs.Create().
				SetUserid(logdata.Userid).
				SetUsertype(logdata.Usertype).
				SetRemarks("ReIssue Notification Successfully").
				SetAction(logdata.Action).
				SetIpaddress(logdata.Ipaddress).
				SetDevicetype(logdata.Devicetype).
				SetOs(logdata.Os).
				SetBrowser(logdata.Browser).
				SetLatitude(logdata.Latitude).
				SetLongitude(logdata.Longitude).
				SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
				SetUserdetails(u.EmployeeName).
				SetUniqueid(u.ID).
				Save(context.Background())

			if logerr3 != nil {
				log.Println(logerr3.Error())
			}

			gctx.String(http.StatusOK, reissueNotifications)
		}
	}
}

// function to Resubmit Draft notification
func PutResubmitDraftNotification(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//resubmitNotification := new(ent.ExamNotifications)
		var resubmitNotification ca_reg.ResubitNotificationStruct
		var mainFunction string = " main - PutResubmitDraftNotification "
		var startFunction string = " - start - PutResubmitDraftNotification "

		if err := gctx.ShouldBindJSON(&resubmitNotification); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := resubmitNotification.Edges.LogData[0]

		resubmitNotifications, status, stgError, dataStatus, err := start.PutResubmitDraftNotification(client, &resubmitNotification)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Notification re-submitted Sucessfully ",
			"data":       resubmitNotifications,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// UpdateResubmitDraftNotificationNew godoc
// @Summary      Update and Resubmit Draft Notification
// @Description  Update and resubmit a draft notification
// @Tags         ExamNotifications
// @Accept       json
// @Produce      json
// @Param        notification  body      ca_reg.UpdateNotificationStruct  true  "Update Notification"
// @Success 200 {object} start.EmployeeMasterResponse "Notification Updated Sucessfully "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /deptexam/notification/UpdateResubmitDraftNotification [put]
func UpdateResubmitDraftNotificationNew(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		var resubmitNotification ca_reg.UpdateNotificationStruct

		var mainFunction string = " main - UpdateResubmitDraftNotificationNew "
		var startFunction string = " - start - UpdateResubmitDraftNotificationNew "

		if err := gctx.ShouldBindJSON(&resubmitNotification); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := resubmitNotification.Edges.LogData[0]

		resubmitNotifications, status, stgError, dataStatus, err := start.UpdateResubmitDraftNotificationNew(client, resubmitNotification)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		util.LoggerNew(client, logdata)

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Notification Updated Sucessfully ",
			"data":       resubmitNotifications,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// modified
func CancelDraftNotification(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		// Parse the JSON body into the resubmitNotification object
		resubmitNotification := new(ent.ExamNotifications)
		if err := gctx.ShouldBindJSON(&resubmitNotification); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the log data
		if len(resubmitNotification.Edges.LogData) <= 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
			return
		}

		logdata := resubmitNotification.Edges.LogData[0]

		// Fetch admin data using the user ID from log data
		u := util.AdminDatas(client, logdata.Userid)
		if u == nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID. No Admin Data Found"})
			return
		}

		// Use the context from the gin request
		ctx := gctx.Request.Context()

		// Retrieve distinct circle facilities
		circleFacilities, err := getDistinctCircleFacilities(client)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Query the existing active notifications for the given exam code and year
		oldAppln, err := client.ExamNotifications.
			Query().
			Where(
				examnotifications.ExamCodeEQ(resubmitNotification.ExamCode),
				examnotifications.ExamYearEQ(resubmitNotification.ExamYear),
				examnotifications.StatusEQ("active"),
			).
			All(ctx)

		if err != nil {

			gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error in fetching data from the Exam table: %v", err)})
			return
		}

		if len(oldAppln) <= 0 {
			gctx.JSON(http.StatusNotFound, gin.H{"error": "No active notification found for this Exam"})
			return
		}
		// Update the status of existing active notifications to inactive
		for _, notif := range oldAppln {
			stat := "inactive_" + time.Now().Format("20060102150405")
			_, err = notif.Update().
				SetStatus(stat).
				Save(ctx)
			if err != nil {
				gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update application: %v", err)})
				return
			}
		}

		// Create a new notification for each circle facility
		for _, facility := range circleFacilities {
			copyNotification := *resubmitNotification
			copyNotification.CircleOfficeFacilityId = facility.CircleFacilityID
			copyNotification.CircleOfficeName = facility.CircleFacilityName

			resubmitNotifications, err := start.UpdateCancelDraftNotification(client, &copyNotification, u)
			if err != nil {
				util.LogError(client, logdata, err)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			util.Logger(client, logdata)
			gctx.String(http.StatusOK, resubmitNotifications)
		}
	}

	return gin.HandlerFunc(fn)
}

func CancelDraftNotificationSingle(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		resubmitNotification := new(ent.ExamNotifications)
		if err := gctx.ShouldBindJSON(&resubmitNotification); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//		var mainFunction string = " main - CancelDraftNotificationSingle "
		//		var startFunction string = " - start - SubUpdateCancelDraftNotificationSingle "

		// Validate the log data
		if len(resubmitNotification.Edges.LogData) <= 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
			return
		}

		logdata := resubmitNotification.Edges.LogData[0]

		// Fetch admin data using the user ID from log data
		u := util.AdminDatas(client, logdata.Userid)
		if u == nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID. No Admin Data Found"})
			return
		}

		// issuenot, status, stgError, dataStatus, err := start.SubUpdateCancelDraftNotificationSingle(client, resubmitNotification, u)

		// if err != nil {
		// 	Remarks = mainFunction + startFunction
		// 	start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)

		//return
		//}
	}
	return gin.HandlerFunc(fn)
}

func CancelDraftNotificationForSpecificCircle(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		// Parse the JSON body into the resubmitNotification object
		resubmitNotification := new(ent.ExamNotifications)
		if err := gctx.ShouldBindJSON(&resubmitNotification); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the log data
		if len(resubmitNotification.Edges.LogData) <= 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
			return
		}

		logdata := resubmitNotification.Edges.LogData[0]

		// Fetch admin data using the user ID from log data
		u := util.AdminDatas(client, logdata.Userid)
		if u == nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID. No Admin Data Found"})
			return
		}

		// Use the context from the gin request
		ctx := gctx.Request.Context()

		// Query the existing active notifications for the given exam code, year, and circle office facility ID
		oldAppln, err := client.ExamNotifications.
			Query().
			Where(
				examnotifications.ExamCodeEQ(resubmitNotification.ExamCode),
				examnotifications.ExamYearEQ(resubmitNotification.ExamYear),
				examnotifications.CircleOfficeFacilityIdEQ(resubmitNotification.CircleOfficeFacilityId),
				examnotifications.StatusEQ("active"),
			).
			All(ctx)

		if err != nil {

			gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error in fetching data from the Exam table: %v", err)})
			return
		}

		if len(oldAppln) <= 0 {
			gctx.JSON(http.StatusNotFound, gin.H{"error": "No active notification found for this Exam and Circle Office Facility"})
			return
		}
		// Update the status of existing active notifications to inactive
		for _, notif := range oldAppln {
			stat := "inactive_" + time.Now().Format("20060102150405")
			_, err = client.ExamNotifications.UpdateOneID(notif.ID).
				SetStatus(stat).
				Where(examnotifications.CircleOfficeFacilityIdEQ(resubmitNotification.CircleOfficeFacilityId)).
				Save(ctx)
			if err != nil {
				gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update notification: %v", err)})
				return
			}
		}

		// Create a new notification for the specific circle facility
		resubmitNotifications, err := start.UpdateCancelDraftNotificationForSpecificCircle(client, resubmitNotification, u)
		if err != nil {
			util.LogError(client, logdata, err)
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		util.Logger(client, logdata)
		gctx.String(http.StatusOK, resubmitNotifications)
	}
}

// New API for gettimg all year exam status

// GetPNDNotifications godoc
// @Summary      Get PND Notifications
// @Description  Get all PND notifications for a specific circle
// @Tags         ExamNotifications
// @Accept       json
// @Produce      json
// @Param        circleid  path      string  true  "Circle ID"
// @Success 200 {object} start.EmployeeMasterResponse "Notification details fetched "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /deptexam/notification/getPNDnotifications/{circleid} [get]
func GetPNDNotifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		circleid := gctx.Param("circleid")
		var mainFunction string = " main - GetPNDNotifications "
		var startFunction string = " - start - GetAllexamNotificationUpdated "

		getNotifications, status, stgError, dataStatus, err := start.GetAllexamNotificationUpdated(client, circleid)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Notification details fetched ",
			"data":       getNotifications,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

//Func that will return the only pending exam list

// GetAllNotificationsMax godoc
// @Summary      Get All Notifications by Year and Circle ID
// @Description  Get all notifications for a specific exam year and circle ID
// @Tags         ExamNotifications
// @Accept       json
// @Produce      json
// @Param        examyear  path      int     true  "Exam Year"
// @Param        circleid  path      string  true  "Circle ID"
// @Success 200 {object} start.EmployeeMasterResponse "Notification details fetched "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /notification/getallnotificationsmax/{examyear}/{circleid} [get]
func GetAllNotificationsMax(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		eyear, _ := strconv.ParseInt(gctx.Param("examyear"), 10, 32)
		circleid := gctx.Param("circleid")
		var mainFunction string = " main - GetAllNotificationsMax "
		var startFunction string = " - start - GetAllexamNotificationbyYearUpdated "

		getNotifications, status, stgError, dataStatus, err := start.GetAllexamNotificationbyYearUpdated(client, int32(eyear), circleid)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Notification details fetched ",
			"data":       getNotifications,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}
