package handlers

import (
	"context"
	"fmt"
	"net/http"
	"recruit/ent"
	"recruit/start"
	"recruit/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MTS to PM/MG Exam - Start
// Get all CA verified for NA..
/* func GetMTSPMMGAllCAVerifiedForNA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")

		circles, err := start.QueryMTSPMMGApplicationsByCAVerifiedForNA(ctx, client, facilityID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)

} */

//OLD REMARKS

/* func GetMTSPMMGHallticketNumberscenter(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		//ec := gctx.Param("id")
		//CenterCode, _ := strconv.ParseInt(ec, 10, 32)
		// Call the GenerateHallticketNumberIP function to generate hall ticket numbers and get the success message
		successMessage, err := start.GenerateHallticketNumberMTSPMMGwithCenterCode(ctx, client)
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

// CENTERE UPADTE API
// func updateExamCentersInMTSPMMGApplsreturnstring(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := context.Background()

// 		// Bind the JSON data to the newappls variable
// 		var req struct {
// 			Newappls []*ent.Exam_Application_MTSPMMG `json:"newappls"`
// 		}
// 		if err := gctx.ShouldBindJSON(&req); err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Call the UpdateExamCentresIPExams function to update the exam centers
// 		status, err := start.UpdateExamCentresGDSPMExamsreturnarray(ctx, client, req.Newappls)
// 		if err != nil {
// 			gctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Display the status
// 		gctx.JSON(http.StatusOK, gin.H{"data": status})
// 	}

// 	return gin.HandlerFunc(fn)
// }

// MISC

/*
// MISC 2

	func UpdateCenterCodeForApplicationsPS(client *ent.Client) gin.HandlerFunc {
		fn := func(gctx *gin.Context) {
			ctx := context.Background()

			var requestDataArray []struct {
				CenterPreference    string `json:"centerPreference"`
				CenterCode          int    `json:"centerCode"`
				ReportingOfficeName string `json:"reportingOfficeName"`
				SeatsToAllot        int    `json:"seatsToAllot"`
			}

			if err := gctx.BindJSON(&requestDataArray); err != nil {
				gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
				return
			}

			for _, requestData := range requestDataArray {
				err := start.UpdateCenterCodeForApplicationsPS(ctx, client, requestData.CenterPreference, requestData.ReportingOfficeName, int32(requestData.CenterCode), int32(requestData.SeatsToAllot))
				if err != nil {
					// Handle error
					fmt.Println("Error updating center codes:", err)
					gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					//return
				} else {
					gctx.JSON(http.StatusOK, gin.H{"Message": "Applications Updated with center code Successfully",
						"data": gin.H{
							"CenterCode":            requestData.CenterCode,
							"CenterPreference Name": requestData.CenterPreference,
							"ReportingOfficeName":   requestData.ReportingOfficeName,
							"Seats Alloted":         requestData.SeatsToAllot,
						}})
				}
			}
		}

		return gin.HandlerFunc(fn)
	}
*/

// MTS to PM/MG Exam - End

// GDS to MTS/PM/MG Exam - Start

// Generates Hall tickets and gives count circle wise.
/* func GetHallticketNumbersGDSPM(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := gctx.Request.Context()

		// Call the GenerateHallticketNumber function to get the hall ticket numbers
		hallticketStats, err := start.GenerateHallticketNumberGDSPM(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Log the generated hall ticket numbers
		for _, stats := range hallticketStats {
			log.Printf("Circle ID: %s, Count: %d\n", stats.CircleID, stats.HallTicketCount)

		}

		// Return the hall ticket numbers as the response
		gctx.JSON(http.StatusOK, gin.H{"data": hallticketStats})
	}

	return gin.HandlerFunc(fn)
}
*/

/* func GetGDSPMMGMTSApplicationsFacilityIDYear(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		year := gctx.Param("id1")

		circles, status, err := start.SubGetGDSPMMGMTSApplicationsFacilityIDYear(ctx, client, facilityID, year)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.SubGetGDSPMMGMTSApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "422"
				Remarks = "start.SubGetGDSPMMGMTSApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			} else {
				Action = "500"
				Remarks = "start.SubGetGDSPMMGMTSApplicationsFacilityIDYear - other error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)
} */
// GDS to MTS/PM/MG Exam - End

// IP Exam - Staart
// Fetch Exam details ...
/* func GetExamDetailsIPByExamCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		username := gctx.Param("id")

		exams, status, err := start.QueryExamsIPByExamNameCode(ctx, client, username)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.QueryExamsIPByExamNameCode - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": os.Getenv("USER_ERROR_REMARKS")})
				return
			} else if status == 422 {
				Action = "422"
				Remarks = "start.QueryExamsIPByExamNameCode - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			} else {
				Action = "500"
				Remarks = "start.QueryExamsIPByExamNameCode - other error"
				gctx.JSON(http.StatusBadRequest, gin.H{"error": os.Getenv("USER_ERROR_REMARKS")})
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exams})

	}

	return gin.HandlerFunc(fn)
} */

// IP Exam - end

//  PS Group B - Start
/* func QueryPSRecommendationsByEmpId(ctx context.Context, client *ent.Client, employeeID int64) ([]*ent.RecommendationsPSApplications, error) {
	//Array of exams

	employeeExists, err := client.RecommendationsPSApplications.
		Query().
		Where(recommendationspsapplications.EmployeeIDEQ(employeeID)).
		Exist(ctx)
	if err != nil {
		log.Println("error checking employee existence: ", err)
		return nil, fmt.Errorf("failed checking employee existence: %w", err)
	}
	if !employeeExists {
		return nil, fmt.Errorf("employee not found with ID: %d", employeeID)
	}

	records, err := client.RecommendationsPSApplications.
		Query().
		Where(recommendationspsapplications.EmployeeIDEQ(employeeID)).
		All(ctx)
	if err != nil {
		log.Println("error querying PS recommendations: ", err)
		return nil, fmt.Errorf("failed to query PS recommendations: %w", err)
	}

	return records, nil
}
*/
/*

 */
/* func GetPSHallTicketWithExamCodeEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		//ctx := context.Background()

		ec := gctx.Param("id1")
		ExamCode, _ := strconv.ParseInt(ec, 10, 32)
		EmployeeID, _ := strconv.ParseInt(gctx.Param("id2"), 10, 64)

		examcenters, err := start.GetPSApplicationsWithHallTicket(client, int32(ExamCode), EmployeeID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": examcenters})

	}

	return gin.HandlerFunc(fn)
} */

/* func GetAllPSPendingWithCandidateNew(client *ent.Client) gin.HandlerFunc {

	//	circles, err := start.QueryPSApplicationsByPendingWithCandidate(ctx, client, facilityID)
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		facilityID := gctx.Param("id")
		id1 := gctx.Param("id1")
		fmt.Println("enter function")
		circles, status, err := start.QueryPSApplicationsByPendingWithCandidateNew(ctx, client, facilityID, id1)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.GetAllPSPendingWithCandidateNew - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": os.Getenv("USER_ERROR_REMARKS")})
				return
			} else if status == 422 {
				Action = "422"
				Remarks = "start.GetAllPSPendingWithCandidateNew - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			} else {
				Action = "500"
				Remarks = "start.GetAllPSPendingWithCandidateNew - other error"
				gctx.JSON(http.StatusBadRequest, gin.H{"error": os.Getenv("USER_ERROR_REMARKS")})
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)

} */

// func UpdateCenterCodeForApplicationsPS(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := context.Background()

// 		var requestDataArray []struct {
// 			CenterPreference    string `json:"centerPreference"`
// 			CenterCode          int    `json:"centerCode"`
// 			ReportingOfficeName string `json:"reportingOfficeName"`
// 			SeatsToAllot        int    `json:"seatsToAllot"`
// 		}

// 		if err := gctx.BindJSON(&requestDataArray); err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
// 			return
// 		}

// 		for _, requestData := range requestDataArray {
// 			err := start.UpdateCenterCodeForApplicationsPS(ctx, client, requestData.CenterPreference, requestData.ReportingOfficeName, int32(requestData.CenterCode), int32(requestData.SeatsToAllot))
// 			if err != nil {
// 				// Handle error
// 				fmt.Println("Error updating center codes:", err)
// 				gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				//return
// 			} else {
// 				gctx.JSON(http.StatusOK, gin.H{"Message": "Applications Updated with center code Successfully",
// 					"data": gin.H{
// 						"CenterCode":            requestData.CenterCode,
// 						"CenterPreference Name": requestData.CenterPreference,
// 						"ReportingOfficeName":   requestData.ReportingOfficeName,
// 						"Seats Alloted":         requestData.SeatsToAllot,
// 					}})
// 			}
// 		}
// 	}

// 	return gin.HandlerFunc(fn)
// }

// PS Group B - End

/// from Main - Start
// func SendtextMSGverify(ctx context.Context, client *ent.Client, application *ent.Exam_Applications_IP) (int32, error) {
// 	// Retrieve the user from the database based on the username
// 	user, err := client.Exam_Applications_IP.Query().Where(exam_applications_ip.UserIDEQ(application.UserID)).Only(ctx)
// 	if err != nil {
// 		return 400, fmt.Errorf("failed to retrieve user: %v", err)
// 	}

// 	// Retrieve the mobile number from the retrieved user
// 	mobileNumber := user.MobileNumber
// 	if mobileNumber == "" {
// 		return 400, errors.New("user's mobile number not found")
// 	}

// 	// Construct the SMS message
// 	msg := fmt.Sprintf("Dear %s - %d, your application status is %s. For more details, visit our website.", user.EmployeeName, user.EmployeeID, application.ApplicationStatus)

// 	// Trigger the SMS
// 	err = sendSMSssss(msg, mobileNumber)
// 	if err != nil {
// 		log.Printf("Failed to send SMS: %v", err)
// 		return 422, fmt.Errorf("failed to send SMS")
// 	}

// 	log.Printf("SMS sent successfully to %s", mobileNumber)

// 	return 200, nil
// }
// func sendSMSssss(msg, phone string) error {
// 	url := "https://api.cept.gov.in/sendsms/api/values/sendsms"

// 	payload := `{
//         "Msg": "` + msg + `",
//         "Phone": "` + phone + `",
//         "TemplateID": "1007453409116473296",
//         "EntityID": "1001081725895192800",
//         "AppName": "IBC"
//     }`

// 	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
// 	if err != nil {
// 		return fmt.Errorf("failed to create request: %v", err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	client := http.DefaultClient
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("failed to send request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("failed to read response body: %v", err)
// 	}
// 	log.Println("Response body:", string(body))

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("unexpected response status: %s", resp.Status)
// 	}

// 	response := string(body)
// 	if !strings.Contains(response, "SMS Pushed to NIC Successfully") {
// 		return fmt.Errorf("failed to send SMS")
// 	}

// 	return nil
// }

// ValidateAdminLogin

// Generates Hall tickets and gives count circle wise.
// func GetHallticketNumbersPS(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context()

// 		// Call the GenerateHallticketNumber function to get the hall ticket numbers
// 		hallticketStats, err := start.GenerateHallticketNumberPS(ctx, client)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Log the generated hall ticket numbers
// 		for _, stats := range hallticketStats {
// 			log.Printf("Circle ID: %s, Count: %d\n", stats.CircleID, stats.HallTicketCount)

// 		}

// 		// Return the hall ticket numbers as the response
// 		gctx.JSON(http.StatusOK, gin.H{"data": hallticketStats})
// 	}

// 	return gin.HandlerFunc(fn)
// }

// generateHallticketIPReturnStringMessage
// func GetHallticketNumbersPSreturnstring(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context()

// 		// Call the generateHallticketIPReturnStringMessage function to generate hall ticket numbers and get the string message
// 		message, err := start.GenerateHallticketPSReturnStringMessage(ctx, client)
// 		if err != nil {
// 			gctx.String(http.StatusBadRequest, "Failed to generate Hall ticket Numbers, retry")
// 			return
// 		}

// 		// Log the generated hall ticket numbers message
// 		log.Println(message)

// 		// Return the success message as the response
// 		gctx.String(http.StatusOK, "Hall tickets generated successfully for eligible candidates")
// 	}

// 	return gin.HandlerFunc(fn)
// }

// Directorate Users

/* func InsertDirectorateUsers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		Dusers := new(ent.DirectorateUsers)
		if err := gctx.ShouldBindJSON(&Dusers); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Dusers, err := start.CreateDirectorateUsers(client, Dusers)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the DirectorateUsers")
	}
	return gin.HandlerFunc(fn)
}

func GetDirectorateUsersByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		DuserID, _ := strconv.ParseInt(id, 10, 32)
		Dusers, err := start.QueryDUsersByID(ctx, client, int32(DuserID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// gctx.JSON(http.StatusOK, gin.H{"data": Dusers})
	}
	return gin.HandlerFunc(fn)
}

func getAllDusers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		Dusers, err := start.QueryeDirectorateUsers(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": Dusers})

	}

	return gin.HandlerFunc(fn)

}

func GetDusersWithEmpID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		DuserID, _ := strconv.ParseInt(id, 10, 32)
		Dusers, err := start.QueryeDirectorateUsersyEmpId(ctx, client, int32(DuserID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Get Dusers by Emp ID and retrieve only details

		// Create an array to store the PaperID and PaperDescription.
		var DuserData []map[string]interface{}
		for _, user := range Dusers {
			DuserData = append(DuserData, map[string]interface{}{
				"EmpID":        user.EmployeedID,
				"EmployeeName": user.EmployeeName,
				"Role":         user.Role,
				"MobileNumber": user.MobileNumber,
				"EmailId":      user.EmailId,
			})
		}
		gctx.JSON(http.StatusOK, gin.H{"data": DuserData})

	}
	return gin.HandlerFunc(fn)
}
*/

/* func GetPASAIPApplicationsFacilityIDYear(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		year := gctx.Param("id1")

		circles, status, err := start.SubGetPASAIPApplicationsFacilityIDYear(ctx, client, facilityID, year)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.SubGetPASAIPApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "422"
				Remarks = "start.SubGetPASAIPApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			} else {
				Action = "500"
				Remarks = "start.SubGetPASAIPApplicationsFacilityIDYear - other error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)
} */

/* func GetMTSPMMGApplicationsFacilityIDYear(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		year := gctx.Param("id1")

		circles, status, err := start.SubGetMTSPMMGApplicationsFacilityIDYear(ctx, client, facilityID, year)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.SubGetMTSPMMGApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "422"
				Remarks = "start.SubGetMTSPMMGApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			} else {
				Action = "500"
				Remarks = "start.SubGetMTSPMMGApplicationsFacilityIDYear - other error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)
} */

/* func GetGDSPASApplicationsFacilityIDYear(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		year := gctx.Param("id1")

		circles, status, err := start.SubGetGDSPASApplicationsFacilityIDYear(ctx, client, facilityID, year)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.SubGetGDSPASApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "422"
				Remarks = "start.SubGetGDSPASApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			} else {
				Action = "500"
				Remarks = "start.SubGetGDSPASApplicationsFacilityIDYear - other error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)
} */

/* func GetLSGIPPSGRBAApplicationsFacilityIDYear(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		year := gctx.Param("id1")

		circles, status, err := start.SubGetLSGIPPSGRBApplicationsFacilityIDYear(ctx, client, facilityID, year)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.SubGetLSGIPPSGRBAApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "422"
				Remarks = "start.SubGetLSGIPPSGRBAApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			} else {
				Action = "500"
				Remarks = "start.SubGetLSGIPPSGRBAApplicationsFacilityIDYear - other error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)
} */

/* func GetPMMGMTSPASAApplicationsFacilityIDYear(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		//facilityID := gctx.Query("facilityID") // Get the facilityID from the query parameter
		facilityID := gctx.Param("id")
		year := gctx.Param("id1")

		circles, status, err := start.SubGetPMMGMTSPASAApplicationsFacilityIDYear(ctx, client, facilityID, year)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.SubGetPMMGMTSPASAApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "422"
				Remarks = "start.SubGetPMMGMTSPASAApplicationsFacilityIDYear - error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			} else {
				Action = "500"
				Remarks = "start.SubGetPMMGMTSPASAApplicationsFacilityIDYear - other error"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}

	return gin.HandlerFunc(fn)
} */

// Verify AdminLoginUser

// admin newUser new dev by vasanth

// QueryFacilityMasterDetailsByID
/* func GetFacilityDetailsByFacilityOfficeID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		officeID := gctx.Param("id")

		facility, err := start.QueryFacilityByOfficeID(ctx, client, officeID)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{"data": facility})
	}

	return gin.HandlerFunc(fn)
 }*/

// func GetNewFacilityDetailsByFacilityOfficeID(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := context.Background()
// 		officeID := gctx.Param("id")

// 		facility, err := start.QueryFacilityByOfficeID(ctx, client, officeID)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		gctx.JSON(http.StatusOK, gin.H{"data": facility})
// 	}

// 	return gin.HandlerFunc(fn)
// }

// Create User and Send SMS

// Get divisions of Circle Office.

/* func QueryDivisionsByCircleOfficeID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		circleOfficeID := gctx.Param("id")
		log.Printf("Querying divisions for Circle Office ID: %s\n", circleOfficeID)

		divisions, err := start.GetDivisionsByCircleOfficeID(ctx, client, circleOfficeID)
		if err != nil {
			log.Printf("Failed to get divisions: %v\n", err)
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		divisionStrings := make([]string, len(divisions))
		for i, division := range divisions {
			divisionStrings[i] = fmt.Sprintf("Reporting Office ID: %s, Reporting Office Name: %s", division.ReportingOfficeID, division.ReportingOfficeName)
		}

		log.Printf("Found %d divisions\n", len(divisions))
		gctx.JSON(http.StatusOK, gin.H{"data": divisionStrings})
	}

	return gin.HandlerFunc(fn)
}
*/

// Middle Ware

/*
	 func util.TokenValidationMiddlewareNew(client *ent.Client) gin.HandlerFunc {
		return func(gctx *gin.Context) {
			token := gctx.Request.Header.Get("UidToken")
			username := gctx.Request.Header.Get("UserName")
			if token == "" {
				gctx.JSON(http.StatusUnauthorized, gin.H{"error": "PASS TOKEN IN HEADER"})
				gctx.Abort()
				return
			}
			if username == "" {
				gctx.JSON(http.StatusUnauthorized, gin.H{"error": "PASS USERNAME IN HEADER"})
				gctx.Abort()
				return
			}
			fmt.Println(token, username)
			// flag, err := authentication.ValidateToken(token, client)
			// if err != nil || !flag {
			// 	gctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			// 	gctx.Abort()
			// 	return
			// }
			var UID string
			adminuserdetails, err1 := client.AdminMaster.Query().Where(adminmaster.UserNameEQ(username), adminmaster.StatussEQ("active")).Only(context.Background())
			usermasteruserdetails, err2 := client.UserMaster.Query().Where(usermaster.UserNameEQ(username), usermaster.StatussEQ("active")).Only(context.Background())

			if err1 != nil && err2 != nil {
				gctx.JSON(http.StatusUnauthorized, gin.H{"error": "No User Found" + err1.Error() + err2.Error()})
				gctx.Abort()
				return
			}

			if adminuserdetails != nil {
				UID = adminuserdetails.UidToken
				fmt.Println("ADMIN USER")
			}
			if usermasteruserdetails != nil {
				UID = usermasteruserdetails.UidToken
				fmt.Println("CANDIADATE USER")
			}

			if token != UID {
				gctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
				gctx.Abort()
				return
			}

			// Pass the user information to the context if needed
			// user, err := authentication.ParseToken(token)
			// gctx.Set("user", user)

			gctx.Next()
		}
	}
*/

// func GetDetailsBasedOnExamCode(client *ent.Client) gin.HandlerFunc {
// 	return func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context()
// 		id := gctx.Param("Exam_Code")
// 		id1, err := strconv.ParseInt(id, 10, 64)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Exam_Code provided"})
// 			return
// 		}

// 		// Handle the case when the Exam_Code is not found
// 		if id1 == 0 {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Exam_Code is missing or invalid"})
// 			return
// 		}

// 		// Get the details for the given exam code
// 		details, err := start.QueryDetailsByExamCode(ctx, client, id1)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// response := gin.H{
// 		// 	"data": details,
// 		// }

// 		gctx.JSON(http.StatusOK, details)
// 	}
// }

// generateHallticketIPReturnStringMessage
/* func GetHallticketNumbersGDSPMreturnstring(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := gctx.Request.Context()

		// Call the generateHallticketIPReturnStringMessage function to generate hall ticket numbers and get the string message
		message, err := start.GenerateHallticketGDSPMReturnStringMessage(ctx, client)
		if err != nil {
			gctx.String(http.StatusBadRequest, "Failed to generate Hall ticket Numbers, retry")
			return
		}

		// Log the generated hall ticket numbers message
		log.Println(message)

		// Return the success message as the response
		gctx.String(http.StatusOK, "Hall tickets generated successfully for eligible candidates")
	}

	return gin.HandlerFunc(fn)
}
*/
// func queryActiveExamsByExamYear(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {

// 		eyear64, _ := strconv.ParseInt(gctx.Param("id"), 10, 32)
// 		year := int32(eyear64)

// 		activeExams, err := start.QueryActiveExamsByExamYear(client, year)

// 		if err != nil {

// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if len(activeExams) <= 0 {
// 			gctx.JSON(http.StatusOK, gin.H{"Message": "No Active exams for the year"})

// 		} else {

// 			// activeexamnames := make([]string, 0)
// 			// examcodes := make([]int, 0)
// 			// notificationnumbers := make([]string, 0)
// 			// for _, examnames := range activeExams {
// 			// 	activeexamnames = append(activeexamnames, examnames.ExamName)
// 			// 	examcodes = append(examcodes, int(examnames.ExamCode))
// 			// 	notificationnumbers = append(notificationnumbers, examnames.NotificationNumber)
// 			// }
// 			// gctx.JSON(http.StatusOK, gin.H{
// 			// 	"data": gin.H{
// 			// 		"Exam Name":           activeExams,
// 			// 		"Exam Code":           examcodes,
// 			// 		"Notification Number": notificationnumbers,
// 			// 	},
// 			// })

// 			examsData := make([]gin.H, 0)

// 			for _, exam := range activeExams {
// 				examsData = append(examsData, gin.H{
// 					"Exam Code":                          exam.ExamCode,
// 					"Exam Name":                          exam.ExamName,
// 					"Notification Number":                exam.NotificationNumber,
// 					"User Name":                          exam.UserName,
// 					"Exam Year ":                         exam.ExamYear,
// 					"Application Start Date":             exam.ApplicationStartDate,
// 					"Application End Date":               exam.ApplicationEndDate,
// 					"Application Correction Start Date ": exam.ApplicationCorrectionStartDate,
// 					"Application Correction End Date":    exam.ApplicationCorrectionLastDate,
// 					"Application Verifiction Last Date":  exam.ApplicationVerificationLastDate,
// 					"Center Allotment End Date ":         exam.CenterAllotmentEndDate,
// 					"Nodal Officer Approval Date ":       exam.NodalOfficerApprovalDate,
// 					"Admit Card Date ":                   exam.AdmitCardDate,
// 					"Crucial Date ":                      exam.CrucialDate,
// 					"Notification Order Number ":         exam.NotificationOrderNumber,
// 					"Order Date ":                        exam.OrderDate,
// 					"Exam Short Name":                    exam.ExamShortName,
// 					//"Calender Issued By":exam.
// 				})
// 			}

// 			gctx.JSON(http.StatusOK, gin.H{
// 				"data": examsData,
// 			})
// 		}
// 	}
// 	return gin.HandlerFunc(fn)
// }

// func queryActiveExamsByExamYear(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {

// 		eyear64, _ := strconv.ParseInt(gctx.Param("id"), 10, 32)
// 		year := int32(eyear64)

// 		activeExams, err := start.QueryActiveExamsByExamYear(client, year)

// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if len(activeExams) <= 0 {
// 			gctx.JSON(http.StatusOK, gin.H{"Message": "No Active exams for the year"})
// 		} else {
// 			examsData := make([]map[string]interface{}, 0)

// 			for _, exam := range activeExams {
// 				examsData = append(examsData, map[string]interface{}{
// 					"Exam Code":                          exam["Exam Code"],
// 					"Exam Name":                          exam["Exam Name"],
// 					"Notification Number":                exam["Notification Number"],
// 					"User Name":                          exam["User Name"],
// 					"Exam Year":                          exam["Exam Year"],
// 					"Application Start Date":             exam["Application Start Date"],
// 					"Application End Date":               exam["Application End Date"],
// 					"Application Correction Start Date":  exam["Application Correction Start Date"],
// 					"Application Correction End Date":    exam["Application Correction End Date"],
// 					"Application Verification Last Date": exam["Application Verification Last Date"],
// 					"Center Allotment End Date":          exam["Center Allotment End Date"],
// 					"Nodal Officer Approval Date":        exam["Nodal Officer Approval Date"],
// 					"Admit Card Date":                    exam["Admit Card Date"],
// 					"Crucial Date":                       exam["Crucial Date"],
// 					"Notification Order Number":          exam["Notification Order Number"],
// 					"Order Date":                         exam["Order Date"],
// 					"Exam Short Name":                    exam["Exam Short Name"],
// 					"Exam Data":                          exam["Exam Data"],
// 				})
// 			}

// 			gctx.JSON(http.StatusOK, gin.H{"data": examsData})
// 		}
// 	}
// 	return gin.HandlerFunc(fn)
// }

// Generates Hall tickets and gives count circle wise.
// func GetHallticketNumbersGDSPA(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context()

// 		// Call the GenerateHallticketNumber function to get the hall ticket numbers
// 		hallticketStats, err := start.GenerateHallticketNumberGDSPA(ctx, client)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Log the generated hall ticket numbers
// 		for _, stats := range hallticketStats {
// 			log.Printf("Circle ID: %s, Count: %d\n", stats.CircleID, stats.HallTicketCount)

// 		}

// 		// Return the hall ticket numbers as the response
// 		gctx.JSON(http.StatusOK, gin.H{"data": hallticketStats})
// 	}

// 	return gin.HandlerFunc(fn)
// }

// generateHallticketIPReturnStringMessage
// func GetHallticketNumbersGDSPAreturnstring(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context()

// 		// Call the generateHallticketIPReturnStringMessage function to generate hall ticket numbers and get the string message
// 		message, err := start.GenerateHallticketGDSPAReturnStringMessage(ctx, client)
// 		if err != nil {
// 			gctx.String(http.StatusBadRequest, "Failed to generate Hall ticket Numbers, retry")
// 			return
// 		}

// 		// Log the generated hall ticket numbers message
// 		log.Println(message)

// 		// Return the success message as the response
// 		gctx.String(http.StatusOK, "Hall tickets generated successfully for eligible candidates")
// 	}

// 	return gin.HandlerFunc(fn)
// }

// Get IP Applications with EmpID - QueryIPExamApplicationsByEmpID

// Generates Hall tickets and gives count circle wise.
// func GetHallticketNumbersPMPA(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context()

// 		// Call the GenerateHallticketNumber function to get the hall ticket numbers
// 		hallticketStats, err := start.GenerateHallticketNumberPMPA(ctx, client)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Log the generated hall ticket numbers
// 		for _, stats := range hallticketStats {
// 			log.Printf("Circle ID: %s, Count: %d\n", stats.CircleID, stats.HallTicketCount)

// 		}

// 		// Return the hall ticket numbers as the response
// 		gctx.JSON(http.StatusOK, gin.H{"data": hallticketStats})
// 	}

// 	return gin.HandlerFunc(fn)
// }

// generateHallticketIPReturnStringMessage
// func GetHallticketNumbersPMPAreturnstring(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context()

// 		// Call the generateHallticketIPReturnStringMessage function to generate hall ticket numbers and get the string message
// 		message, err := start.GenerateHallticketPMPAReturnStringMessage(ctx, client)
// 		if err != nil {
// 			gctx.String(http.StatusBadRequest, "Failed to generate Hall ticket Numbers, retry")
// 			return
// 		}

// 		// Log the generated hall ticket numbers message
// 		log.Println(message)

// 		// Return the success message as the response
// 		gctx.String(http.StatusOK, "Hall tickets generated successfully for eligible candidates")
// 	}

// 	return gin.HandlerFunc(fn)
// }

// func CreateNewGDSPMApplications(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		newPAAppplns := new(ent.Exam_Applications_GDSPM)
// 		if err := gctx.ShouldBindJSON(&newPAAppplns); err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		if len(newPAAppplns.Edges.LogData) <= 0 {
// 			fmt.Println(len(newPAAppplns.Edges.LogData))
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
// 			return
// 		}
// 		//logdata := newPAAppplns.Edges.LogData[0]

// 		fmt.Println(newPAAppplns)
// 		//	fmt.Println(newPAAppplns.Edges.CirclePrefRefMTSPMMG)

// 		fmt.Println("Logdata----------------------", newPAAppplns.Edges.LogData)
// 		newPAAppplns,status, err := start.CreateGDSPMApplicationss(client, &newPAAppplns)

// 		logdata := newPAAppplns.Edges.LogData[0]

// 		if err != nil {
// 			if status == 400 {
// 				logdata.Remarks = "400 error occured in start.IpApplication"
// 				util.LogErrorNew(client, logdata, err)
// 				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
// 				return
// 			} else if status == 422 {
// 				logdata.Remarks = "422 error occured in start.IpApplication"
// 				util.LogErrorNew(client, logdata, err)
// 				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
// 				return
// 			} else {
// 				logdata.Remarks = "start.IpApplication - Other error"
// 				util.LogErrorNew(client, logdata, err)
// 				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
// 				return
// 			}
// 		}

// 	}
// 	return gin.HandlerFunc(fn)
// }

//Update Verification details

// func VerifyGDSPMApplication(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		newAppln := new(ent.Exam_Applications_GDSPM)

// 		if err := gctx.ShouldBindJSON(&newAppln); err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		application, err := start.UpdateApplicationRemarkssGDSPM(client, newAppln)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		gctx.JSON(http.StatusOK, gin.H{
// 			"data": gin.H{
// 				"EmployeeID":          application.EmployeeID,
// 				"ApplicationStatus":   application.ApplicationStatus,
// 				"Application Remarks": application.AppliactionRemarks,
// 				//	"CAOldRemarks":        application.CAPreviousRemarks,
// 			},
// 		})
// 	}

//		return gin.HandlerFunc(fn)
//	}

// Function to create Draft notification

// function to retrieve all the office details for a particular PIN code
/* func queryOfficeByPincode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		PinID, _ := strconv.ParseInt(id, 10, 32)
		offices, err := start.QueryOfficeByPincode(ctx, client, int32(PinID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": offices})
	}
	return gin.HandlerFunc(fn)
}
*/

// Update Nodal recommendations

// QueryLatestPendingWithCandidateApplicationByEmpID

//CA Pending with EmpID -- QueryIPApplicationsByCAPendingByEmpID

// Approve Hall Ticket for NO.
// func UpdateApprovalOfHallTicketForNO(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		// Parse the JSON input into ApprovalInput struct
// 		var input ApprovalInput
// 		if err := gctx.ShouldBindJSON(&input); err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		var successMessage string
// 		var err error

// 		if input.ExamCode == 1 {
// 			successMessage, err = start.ApproveHallTicketGenerationByNOForPSExam(client, input.ExamCode, input.FacilityID, input.ApproveHallTicket)
// 		} else if input.ExamCode == 2 {
// 			successMessage, err = start.ApproveHallTicketGenerationByNOForIPExam(client, input.ExamCode, input.FacilityID, input.ApproveHallTicket)
// 		} else if input.ExamCode == 4 {
// 			successMessage, err = start.ApproveHallTicketGenerationByNOForGDSPAExam(client, input.ExamCode, input.FacilityID, input.ApproveHallTicket)
// 		} else if input.ExamCode == 3 {
// 			successMessage, err = start.ApproveHallTicketGenerationByNOForPMPAExam(client, input.ExamCode, input.FacilityID, input.ApproveHallTicket)
// 		} else if input.ExamCode == 6 {
// 			successMessage, err = start.ApproveHallTicketGenerationByNOForGDSPMExam(client, input.ExamCode, input.FacilityID, input.ApproveHallTicket)
// 		} else {
// 			log.Printf("Invalid exam code: %d", input.ExamCode)
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
// 			return
// 		}

// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		gctx.String(http.StatusOK, successMessage)
// 	}

// 	return gin.HandlerFunc(fn)
// }

/*func getLoginDetails(client *ent.Client) gin.HandlerFunc {
fn := func(gctx *gin.Context) {
	ctx := context.Background()
	id := gctx.Param("id")
	//var examID int32
	empid, _ := strconv.ParseInt(id, 10, 32)
	newUsers, err := start.QueryLoginByEmpID(ctx, client, int32(empid))

	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"data": newUsers})

}
}*/

// Generates Hall tickets and gives count circle wise.
// func GetHallticketNumbers(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context()

// 		// Call the GenerateHallticketNumber function to get the hall ticket numbers
// 		hallticketStats, err := start.GenerateHallticketNumber(ctx, client)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Log the generated hall ticket numbers
// 		for _, stats := range hallticketStats {
// 			log.Printf("Circle ID: %s, Count: %d\n", stats.CircleID, stats.HallTicketCount)

// 		}

// 		// Return the hall ticket numbers as the response
// 		gctx.JSON(http.StatusOK, gin.H{"data": hallticketStats})
// 	}

// 	return gin.HandlerFunc(fn)
// }

// generateHallticketIPReturnStringMessage
// func GetHallticketNumbersreturnstring(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context()

// 		// Call the generateHallticketIPReturnStringMessage function to generate hall ticket numbers and get the string message
// 		message, err := start.GenerateHallticketIPReturnStringMessage(ctx, client)
// 		if err != nil {
// 			gctx.String(http.StatusBadRequest, "Failed to generate Hall ticket Numbers, retry")
// 			return
// 		}

// 		// Log the generated hall ticket numbers message
// 		log.Println(message)

// 		// Return the success message as the response
// 		gctx.String(http.StatusOK, "Hall tickets generated successfully for eligible candidates")
// 	}

// 	return gin.HandlerFunc(fn)
// }

/*func updateUserWithPassword(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newemp := new(ent.UserMaster)

		id := gctx.Param("id")

		empid, _ := strconv.ParseInt(id, 10, 64)

		if err := gctx.ShouldBindJSON(&newemp); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		emps, err := start.UpdateUserByEmpId(client, empid, newemp)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": emps})
	}
	return gin.HandlerFunc(fn)
}
*/

// For Division Master
/* func CreateDivisionMasters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newDivisions := new(ent.DivisionMaster)
		if err := gctx.ShouldBindJSON(&newDivisions); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newDivisions, err := start.CreateDivisionMaster(client, newDivisions)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Divisions Master")
	}
	return gin.HandlerFunc(fn)
}

func getAllDivisions(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		divisions, err := start.QueryDivisionMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": divisions})

	}
	return gin.HandlerFunc(fn)
}

// QueryCircleMasterWithRegions
func GetDivisionsByRegionCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		DivisionID, _ := strconv.ParseInt(id, 10, 32)
		divisions, err := start.QueryDivisionMasterByRegionCode(ctx, client, int32(DivisionID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": divisions})
	}
	return gin.HandlerFunc(fn)
}
func GetDivisionsByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		DivID, _ := strconv.ParseInt(id, 10, 32)

		divisions, err := start.QueryDivisionMasterByID(ctx, client, int32(DivID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": divisions})

	}

	return gin.HandlerFunc(fn)
}

// Facilities
func CreateFacilityMasters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newfac := new(ent.Facility)
		if err := gctx.ShouldBindJSON(&newfac); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newfac, err := start.CreateFacility(client, newfac)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Facility Master")
	}
	return gin.HandlerFunc(fn)
}

func getAllFacilities(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		facs, err := start.QueryFacilitiesMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})

	}
	return gin.HandlerFunc(fn)
}
func getAllFacilitiesNew(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		facs, err := start.QueryFacilitiesNewMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})

	}
	return gin.HandlerFunc(fn)
}

func GetFacilitiesByDivisionCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		//FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, err := start.QueryFacilityMasterByDivisionCode(ctx, client, id)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}
func GetNewFacilitiesByDivisionCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		//FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, err := start.QueryNewFacilityMasterByDivisionCode(ctx, client, id)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}
func GetFacilityByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		facID, _ := strconv.ParseInt(id, 10, 32)

		facs, err := start.QueryFacilityMasterByID(ctx, client, int32(facID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})

	}

	return gin.HandlerFunc(fn)
}
func GetNewFacilityByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		facID, _ := strconv.ParseInt(id, 10, 32)

		facs, err := start.QueryFacilityNewMasterByID(ctx, client, int32(facID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})

	}

	return gin.HandlerFunc(fn)
}

func GetFacilitiesByRegionCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		//FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, err := start.QueryFacilityMasterByRegionCode(ctx, client, id)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}
func GetNewFacilitiesByRegionCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		//FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, err := start.QueryNewFacilityMasterByRegionCode(ctx, client, id)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}



func GetFacilitiesByCircleCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		//FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, err := start.QueryFacilityMasterByCircleCode(ctx, client, id)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}
func GetNewFacilitiesByCircleCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		//FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, err := start.QueryNewFacilityMasterByCircleCode(ctx, client, id)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}
*/

// QueryFacilities

/*func getAllFacilities(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		newEmployeCategories, err := start.QueryFacilities(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": newEmployeCategories})
	}

	return gin.HandlerFunc(fn)

}
*/
// GetCircles bY ID

/* func GetCircleByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		CircleID, _ := strconv.ParseInt(id, 10, 32)

		circles, err := start.QueryCircleMasterByID(ctx, client, int32(CircleID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})

	}

	return gin.HandlerFunc(fn)
}

// For Circle Master
func CreateCircleMasters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newCircles := new(ent.CircleMaster)
		if err := gctx.ShouldBindJSON(&newCircles); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newCircles, err := start.CreateCircleMaster(client, newCircles)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Circle Master")
	}
	return gin.HandlerFunc(fn)
}

func getAllCircles(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		circles, err := start.QueryCircleMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})

	}

	return gin.HandlerFunc(fn)

}
*/

// For Region Master
/* func CreateRegionMasters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newRegions := new(ent.RegionMaster)
		if err := gctx.ShouldBindJSON(&newRegions); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newRegions, err := start.CreateRegionMaster(client, newRegions)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Regions Master")
	}
	return gin.HandlerFunc(fn)
}

func getAllRegions(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		circles, err := start.QueryRegionMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})

	}

	return gin.HandlerFunc(fn)

}

// QueryCircleMasterWithRegions
func GetRegionsByCircleCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		CircleID, _ := strconv.ParseInt(id, 10, 32)
		circles, err := start.QueryRegionMasterByCircleCode(ctx, client, int32(CircleID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
}
func GetRegionsByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		CircleID, _ := strconv.ParseInt(id, 10, 32)

		regions, err := start.QueryRegionMasterByID(ctx, client, int32(CircleID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": regions})

	}

	return gin.HandlerFunc(fn)
}
*/
// update examcentres
/*
func updateExamCentersInIPAppls(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()

		// Bind the JSON data to the newappls variable
		var req struct {
			Newappls []*ent.Exam_Applications_IP `json:"newappls"`
		}
		if err := gctx.ShouldBindJSON(&req); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the UpdateExamCentresIPExams function to update the exam centers
		updatedRecords, err := start.UpdateExamCentresIPExams(ctx, client, req.Newappls)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Display the updated records
		gctx.JSON(http.StatusOK, gin.H{"data": updatedRecords})
	}

	return gin.HandlerFunc(fn)
} */

/* func updateExamCenterPS(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newCenter := new(ent.Center)
		id := gctx.Param("id")
		CenterID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newCenter); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newCenter, err := start.UpdateCenter(client, int32(CenterID), newCenter)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"ID":             newCenter.ID,
				"ExamCenterName": newCenter.ExamCenterName,
				"Address":        newCenter.Address,
				"Pincode":        newCenter.Pincode,
				"Landmark":       newCenter.Landmark,
				//"SeatCapacity":   newCenter.MaxSeats,
				//"AllotedSeats":   newCenter.NoAlloted,
				"Message": "Exam Centre Updated",
				"Status":  newCenter.Status,
			},
		})

	}

	return gin.HandlerFunc(fn)
}
*/

// Get Exam Centres based on Admin Office ID .
/* func GetExamCentresBynodalOfficeID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		username := gctx.Param("id")

		examcentres, err := start.GetExamCentresByAdminFacilityOfficeID(ctx, client, username)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": examcentres})
	}
	return gin.HandlerFunc(fn)
}
*/
// from Main - End

/* func updateExamCentersInPSAppls(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()

		// Bind the JSON data to the newappls variable
		var req struct {
			Newappls []*ent.Exam_Applications_PS `json:"newappls"`
		}
		if err := gctx.ShouldBindJSON(&req); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the UpdateExamCentresIPExams function to update the exam centers
		updatedRecords, err := start.UpdateExamCentresPSExams(ctx, client, req.Newappls)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Display the updated records
		gctx.JSON(http.StatusOK, gin.H{"data": updatedRecords})
	}

	return gin.HandlerFunc(fn)
} */

// func updateExamCenterGDSPA(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		newCenter := new(ent.Center)
// 		id := gctx.Param("id")
// 		CenterID, _ := strconv.ParseInt(id, 10, 32)

// 		if err := gctx.ShouldBindJSON(&newCenter); err != nil {

// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		newCenter, err := start.UpdateCenter(client, int32(CenterID), newCenter)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		gctx.JSON(http.StatusOK, gin.H{
// 			"data": gin.H{
// 				"ID":             newCenter.ID,
// 				"ExamCenterName": newCenter.ExamCenterName,
// 				"Address":        newCenter.Address,
// 				"Pincode":        newCenter.Pincode,
// 				"Landmark":       newCenter.Landmark,
// 				//"SeatCapacity":   newCenter.MaxSeats,
// 				//"AllotedSeats":   newCenter.NoAlloted,
// 				"Message": "Exam Centre Updated",
// 				"Status":  newCenter.Status,
// 			},
// 		})

// 	}

// 	return gin.HandlerFunc(fn)
// }

/* func updateExamCentersInGDSPAAppls(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()

		// Bind the JSON data to the newappls variable
		var req struct {
			Newappls []*ent.Exam_Applications_GDSPA `json:"newappls"`
		}
		if err := gctx.ShouldBindJSON(&req); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the UpdateExamCentresGDSPAExams function to update the exam centers
		updatedRecords, err := start.UpdateExamCentresGDSPAExams(ctx, client, req.Newappls)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Display the updated records
		gctx.JSON(http.StatusOK, gin.H{"data": updatedRecords})
	}

	return gin.HandlerFunc(fn)
} */

// func updateExamCenterGDSPM(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		newCenter := new(ent.Center)
// 		id := gctx.Param("id")
// 		CenterID, _ := strconv.ParseInt(id, 10, 32)

// 		if err := gctx.ShouldBindJSON(&newCenter); err != nil {

// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		newCenter, err := start.UpdateCenter(client, int32(CenterID), newCenter)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		gctx.JSON(http.StatusOK, gin.H{
// 			"data": gin.H{
// 				"ID":             newCenter.ID,
// 				"ExamCenterName": newCenter.ExamCenterName,
// 				"Address":        newCenter.Address,
// 				"Pincode":        newCenter.Pincode,
// 				"Landmark":       newCenter.Landmark,
// 				//"SeatCapacity":   newCenter.MaxSeats,
// 				//"AllotedSeats":   newCenter.NoAlloted,
// 				"Message": "Exam Centre Updated",
// 				"Status":  newCenter.Status,
// 			},
// 		})

// 	}

// 	return gin.HandlerFunc(fn)
// }
// func UpdateExamCenterPMPA(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {
// 		newCenter := new(ent.Center)
// 		id := gctx.Param("id")
// 		CenterID, _ := strconv.ParseInt(id, 10, 32)

// 		if err := gctx.ShouldBindJSON(&newCenter); err != nil {

// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		newCenter, err := start.UpdateCenter(client, int32(CenterID), newCenter)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		gctx.JSON(http.StatusOK, gin.H{
// 			"data": gin.H{
// 				"ID":             newCenter.ID,
// 				"ExamCenterName": newCenter.ExamCenterName,
// 				"Address":        newCenter.Address,
// 				"Pincode":        newCenter.Pincode,
// 				"Landmark":       newCenter.Landmark,
// 				//"SeatCapacity":   newCenter.MaxSeats,
// 				//"AllotedSeats":   newCenter.NoAlloted,
// 				"Message": "Exam Centre Updated",
// 				"Status":  newCenter.Status,
// 			},
// 		})

// 	}

//		return gin.HandlerFunc(fn)
//	}
func GetExamIDsandNames(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		exams, err := start.QueryExamIDAndNames(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exams})

	}

	return gin.HandlerFunc(fn)

}
func CreateExam(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newExam := new(ent.Exam)

		if err := gctx.ShouldBindJSON(&newExam); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(newExam.Edges.LogData) <= 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
			return
		}
		logdata := newExam.Edges.LogData[0]
		newExam, err := start.CreateExam(client, newExam)
		if err != nil {
			logerr := util.LogError(client, logdata, err)
			if logerr != nil {
				fmt.Println(logerr.Error() + "LOG ERROR")
			}
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Exam")
		util.Logger(client, logdata)
	}

	return gin.HandlerFunc(fn)
}

/* func updateExamCentersInGDSPMAppls(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()

		// Bind the JSON data to the newappls variable
		var req struct {
			Newappls []*ent.Exam_Applications_GDSPM `json:"newappls"`
		}
		if err := gctx.ShouldBindJSON(&req); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the UpdateExamCentresGDSPMExams function to update the exam centers
		updatedRecords, err := start.UpdateExamCentresGDSPMExams(ctx, client, req.Newappls)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Display the updated records
		gctx.JSON(http.StatusOK, gin.H{"data": updatedRecords})
	}

	return gin.HandlerFunc(fn)
} */

/* func getAllAdminLogins(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		Dusers, err := start.QueryAdminLogin(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": Dusers})

	}

	return gin.HandlerFunc(fn)

} */

/* func getAllExamPapers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		exampapers, err := start.QueryExamPapers(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exampapers})

	}

	return gin.HandlerFunc(fn)

} */

// Exam Papers
func GetExamPapers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		exampaps, err := start.QueryExamPapers(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exampaps})

	}

	return gin.HandlerFunc(fn)

}

// QueryExamPapersByExamCode
// QueryCircleMasterWithRegions
func GetExamPapersWithExamCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		ExamPaperID, _ := strconv.ParseInt(id, 10, 32)
		exampapers, err := start.QueryExamPapersByExamCode(ctx, client, int32(ExamPaperID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Get ExamPapers by ExamCode and retrieve only ID and PaperDescription.

		// Create an array to store the PaperID and PaperDescription.
		var paperData []map[string]interface{}
		for _, paper := range exampapers {
			paperData = append(paperData, map[string]interface{}{
				"PaperID":          paper.ID,
				"PaperDescription": paper.PaperDescription,
			})
		}
		gctx.JSON(http.StatusOK, gin.H{"data": paperData})
	}
	return gin.HandlerFunc(fn)
}

/* func updateNotification(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newNotification := new(ent.Notification)

		id := gctx.Param("id")

		NotificationID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newNotification); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		exam, err := start.UpdateNotification(client, int32(NotificationID), newNotification)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exam})

	}

	return gin.HandlerFunc(fn)
} */

/* func getNotifications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		notfications, err := start.QueryNotification(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": notfications})

	}

	return gin.HandlerFunc(fn)

} */

func CreateExamCalendarData(client *ent.Client) gin.HandlerFunc {
	// Parse the request body and extract the necessary data
	fn := func(gctx *gin.Context) {
		//var requestData ExamCalendar
		newExamCalendar := new(ent.ExamCalendar)
		if err := gctx.ShouldBindJSON(&newExamCalendar); err != nil {
			// Handle invalid request data error
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		// Start a new transaction
		newExamCalendar, err := start.CreateExamCalendar(client, newExamCalendar)
		if err != nil {
			// Handle transaction error
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Notification")
	}
	return gin.HandlerFunc(fn)
}
func CreateNotification(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newNotification := new(ent.Notification)

		if err := gctx.ShouldBindJSON(&newNotification); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newNotification, err := start.CreateNotification(client, newNotification)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Notification")

	}
	return gin.HandlerFunc(fn)
}
func GetNotificationID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		NotificationID, _ := strconv.ParseInt(id, 10, 32)

		notification, err := start.QueryNotificationID(ctx, client, int32(NotificationID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": notification})

	}

	return gin.HandlerFunc(fn)
}

// function to Reissue notification
// func PutReIssueNotification(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {

// 		reissueNotification := new(ent.ExamNotifications)

// 		if err := gctx.ShouldBindJSON(&reissueNotification); err != nil {

// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		if len(reissueNotification.Edges.LogData) <= 0 {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
// 			return
// 		}

// 		logdata := reissueNotification.Edges.LogData[0]

// 		u := start.AdminData(client, logdata.Userid)
// 		if u == nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID No Admin Data Found"})
// 			return
// 		}
// 		ctx := context.Background()
// 		_, logerr1 := client.Logs.Create().
// 			SetUserid(logdata.Userid).
// 			SetUsertype(logdata.Usertype).
// 			SetRemarks(logdata.Remarks).
// 			SetAction(logdata.Action).
// 			SetIpaddress(logdata.Ipaddress).
// 			SetDevicetype(logdata.Devicetype).
// 			SetOs(logdata.Os).
// 			SetBrowser(logdata.Browser).
// 			SetLatitude(logdata.Latitude).
// 			SetLongitude(logdata.Longitude).
// 			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
// 			SetUserdetails(u.EmployeeName).
// 			SetUniqueid(u.ID).
// 			Save(ctx)

// 		if logerr1 != nil {
// 			log.Println(logerr1.Error())
// 		}

// 		reissueNotifications, err := start.PutReIssueNotification(client, reissueNotification)

// 		if err != nil {
// 			_, logerr2 := client.Logs.Create().
// 				SetUserid(logdata.Userid).
// 				SetUsertype(logdata.Usertype).
// 				SetRemarks(err.Error()).
// 				SetAction(logdata.Action).
// 				SetIpaddress(logdata.Ipaddress).
// 				SetDevicetype(logdata.Devicetype).
// 				SetOs(logdata.Os).
// 				SetBrowser(logdata.Browser).
// 				SetLatitude(logdata.Latitude).
// 				SetLongitude(logdata.Longitude).
// 				SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
// 				SetUserdetails(u.EmployeeName).
// 				SetUniqueid(u.ID).
// 				Save(context.Background())

// 			if logerr2 != nil {
// 				log.Println(logerr2.Error())
// 			}
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		_, logerr3 := client.Logs.Create().
// 			SetUserid(logdata.Userid).
// 			SetUsertype(logdata.Usertype).
// 			SetRemarks("ReIssuse Notification  Sucessfully").
// 			SetAction(logdata.Action).
// 			SetIpaddress(logdata.Ipaddress).
// 			SetDevicetype(logdata.Devicetype).
// 			SetOs(logdata.Os).
// 			SetBrowser(logdata.Browser).
// 			SetLatitude(logdata.Latitude).
// 			SetLongitude(logdata.Longitude).
// 			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
// 			SetUserdetails(u.EmployeeName).
// 			SetUniqueid(u.ID).
// 			Save(context.Background())

// 		if logerr3 != nil {
// 			log.Println(logerr3.Error())
// 		}

// 		gctx.String(http.StatusOK, reissueNotifications)

// 	}
// 	return gin.HandlerFunc(fn)
// }

// function to cancel the draft notification
// func CancelDraftNotification(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {

// 		resubmitNotification := new(ent.ExamNotifications)

// 		if err := gctx.ShouldBindJSON(&resubmitNotification); err != nil {

// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		if len(resubmitNotification.Edges.LogData) <= 0 {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
// 			return
// 		}

// 		logdata := resubmitNotification.Edges.LogData[0]

// 		u := start.AdminData(client, logdata.Userid)
// 		if u == nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID No Admin Data Found"})
// 			return
// 		}
// 		ctx := gctx.Request.Context()

// 		oldAppln, err := client.ExamNotifications.
// 			Query().
// 			Where(
// 				examnotifications.ExamCodeEQ(resubmitNotification.ExamCode),
// 				examnotifications.ExamYearEQ(resubmitNotification.ExamYear),
// 				examnotifications.StatusEQ("active"),
// 			).
// 			All(ctx)

// 		if err != nil {
// 			if ent.IsNotFound(err) {
// 				gctx.JSON(http.StatusNotFound, gin.H{"error": "No active notification found for this Exam"})
// 				return
// 			}
// 			gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error in fetching data from the Exam table: %v", err)})
// 			return
// 		}

// 		for _, notif := range oldAppln {
// 			stat := "inactive_" + time.Now().Format("20060102150405")
// 			_, err = notif.Update().
// 				SetStatus(stat).
// 				Save(ctx)
// 			if err != nil {
// 				gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update application: %v", err)})
// 				return
// 			}
// 		}

// 		resubmitNotifications, err := start.UpdateCancelDraftNotification(client, resubmitNotification, u)

// 		if err != nil {
// 			util.LogError(client, logdata, err)

// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		util.Logger(client, logdata)
// 		gctx.String(http.StatusOK, resubmitNotifications)
// 	}
// 	return gin.HandlerFunc(fn)
// }
// function to issue notification
// func PutIssueNotification(client *ent.Client) gin.HandlerFunc {
// 	fn := func(gctx *gin.Context) {

// 		issueNotification := new(ent.ExamNotifications)

// 		if err := gctx.ShouldBindJSON(&issueNotification); err != nil {

// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if len(issueNotification.Edges.LogData) <= 0 {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
// 			return
// 		}

// 		logdata := issueNotification.Edges.LogData[0]

// 		u := start.AdminData(client, logdata.Userid)
// 		if u == nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID No Admin Data Found"})
// 			return
// 		}
// 		ctx := context.Background()
// 		_, logerr1 := client.Logs.Create().
// 			SetUserid(logdata.Userid).
// 			SetUsertype(logdata.Usertype).
// 			SetRemarks(logdata.Remarks).
// 			SetAction(logdata.Action).
// 			SetIpaddress(logdata.Ipaddress).
// 			SetDevicetype(logdata.Devicetype).
// 			SetOs(logdata.Os).
// 			SetBrowser(logdata.Browser).
// 			SetLatitude(logdata.Latitude).
// 			SetLongitude(logdata.Longitude).
// 			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
// 			SetUserdetails(u.EmployeeName).
// 			SetUniqueid(u.ID).
// 			Save(ctx)

// 		if logerr1 != nil {
// 			log.Println(logerr1.Error())
// 		}

// 		issueNotifications, err := start.IssueExamNotification(client, issueNotification, u)

// 		if err != nil {
// 			ctx := context.Background()
// 			_, logerr2 := client.Logs.Create().
// 				SetUserid(logdata.Userid).
// 				SetUsertype(logdata.Usertype).
// 				SetRemarks(err.Error()).
// 				SetAction(logdata.Action).
// 				SetIpaddress(logdata.Ipaddress).
// 				SetDevicetype(logdata.Devicetype).
// 				SetOs(logdata.Os).
// 				SetBrowser(logdata.Browser).
// 				SetLatitude(logdata.Latitude).
// 				SetLongitude(logdata.Longitude).
// 				SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
// 				SetUserdetails(u.EmployeeName).
// 				SetUniqueid(u.ID).
// 				Save(ctx)

// 			if logerr2 != nil {
// 				log.Println(logerr2.Error())
// 			}

// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		_, logerr3 := client.Logs.Create().
// 			SetUserid(logdata.Userid).
// 			SetUsertype(logdata.Usertype).
// 			SetRemarks("ISSUE UPDATE DONE").
// 			SetAction(logdata.Action).
// 			SetIpaddress(logdata.Ipaddress).
// 			SetDevicetype(logdata.Devicetype).
// 			SetOs(logdata.Os).
// 			SetBrowser(logdata.Browser).
// 			SetLatitude(logdata.Latitude).
// 			SetLongitude(logdata.Longitude).
// 			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
// 			SetUserdetails(u.EmployeeName).
// 			SetUniqueid(u.ID).
// 			Save(context.Background())

// 		if logerr3 != nil {
// 			log.Println(logerr3.Error())
// 		}
// 		gctx.String(http.StatusOK, issueNotifications)

// 	}
// 	return gin.HandlerFunc(fn)
// }

//modified
// func getDistinctCircleFacilities(client *ent.Client) ([]*ent.FacilityMasters, error) {
// 	ctx := context.Background()

// 	// Define a struct to hold the distinct fields
// 	var circleFacilities []struct {
// 		CircleFacilityID   int    `json:"circle_facility_id"`
// 		CircleFacilityName string `json:"circle_facility_name"`
// 	}

// 	// Perform the distinct query using Ent
// 	err := client.FacilityMasters.Query().
// 		Where(facilitymasters.CircleFacilityIDIsNotNil()).
// 		GroupBy(facilitymasters.FieldCircleFacilityID, facilitymasters.FieldCircleFacilityName).
// 		Scan(ctx, &circleFacilities)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve distinct circle facilities: %w", err)
// 	}

// 	return circleFacilities, nil
// }

/* func GetAdminPasswordByUsername(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := gctx.Request.Context()
		id := gctx.Param("username")

		// Get the password for the given employee ID
		password, err := start.QueryAdminPasswordByEmpId(ctx, client, id)
		fmt.Println(password)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Retrieve user information to send SMS
		user, err := client.AdminMaster.Query().
			Where(adminmaster.UserNameEQ(id)).
			Only(ctx)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var smssentstatus string
		// Send the password through SMS
		statussms, err := start.SendSMSGetAdminPassword(ctx, client, user)
		if err != nil {
			if statussms == 400 {
				smssentstatus = "Failed"
				Action = "400"
				Remarks = "400 error from start.QueryEmpUsersByEmpId"
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if statussms == 422 {
				Action = "422"
				Remarks = "422 error from start.QueryEmpUsersByEmpId"
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return
			} else {
				Action = "500"
				Remarks = "500 error from start.QueryEmpUsersByEmpId"
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			}
		} else {
			smssentstatus = "Success"
		}
		gctx.JSON(http.StatusOK, gin.H{
			"SMS sent status is ": smssentstatus,
		})

	}
}
*/

/* func GenerateOTPAndSendSMSAndSave(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := gctx.Request.Context()
		newUser := new(ent.UserMaster)
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//user, err := start.CreateUserByEmpId(ctx, client, empID)
		_, _, err := start.SendSMSCandidateOTP(ctx, client, newUser)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//gctx.JSON(http.StatusOK, gin.H{"message": message})
		gctx.JSON(http.StatusOK, gin.H{"message": "OTP Generated and sent successfully"})
	}
	return gin.HandlerFunc(fn)
} */

// function to insert bulk users
/* func InsertBulkUsersHandler(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		var nbu []*ent.UserMaster
		if err := gctx.BindJSON(&nbu); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		//insertedUsers, err :=
		start.InsertBulkUsers(gctx, client, nbu)

	}
}
*/

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags Users
// @Accept json
// @Produce json
// @Param user body ca_reg.UserRequest true "User  data"
// @Success 200 {object} start.EmployeeMasterResponse "Successfully created the User"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/profile [post]
/* func CreateUser(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newUser := new(ent.User)
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUser, err := start.CreateUser(client, newUser)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the User")

	}
	return gin.HandlerFunc(fn)
} */

//Query for existing User from user master by user name

/* func CheckAdminUserByUserName(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		username := gctx.Param("id")

		var mainFunction string = " main - CheckAdminUserByUserName "
		var startFunction string = " - start - QueryUserMasterByUserName "
		user, status, stgError, _, err := start.QueryUserMasterByUserName(ctx, client, username)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": gin.H{
			"Username":     user.UserName,
			"FacilityID":   user.FacilityID,
			"RoleUserCode": user.RoleUserCode,
		}})
	}
	return gin.HandlerFunc(fn)
} */

// main.go
//r.PUT(apiSubPath+"users/updatepassword/:empid", h.UpdateUserPassword(client))
// 	r.GET(apiSubPath+"gdspmmgmts/getapplicationsfacilityidyear/:id/:id1", util.TokenValidationMiddlewareNew(client), GetGDSPMMGMTSApplicationsFacilityIDYear(client))
// 	r.GET(apiSubPath+"lsgippsgrb/getapplicationsfacilityidyear/:id/:id1", util.TokenValidationMiddlewareNew(client), GetLSGIPPSGRBAApplicationsFacilityIDYear(client))
// 	r.GET(apiSubPath+"mtspmmg/getapplicationsfacilityidyear/:id/:id1", util.TokenValidationMiddlewareNew(client), GetMTSPMMGApplicationsFacilityIDYear(client))
// 	r.PUT(apiSubPath+"GDSPAexams/halltickets", util.TokenValidationMiddlewareNew(client), h.GetGDSPAHallticketNumberscenter(client)) // just returns a string message generated successfully.
// 	r.PUT(apiSubPath+"GDSPMexams/halltickets", util.TokenValidationMiddlewareNew(client), h.GetGDSPMHallticketNumberscenter(client)) // just returns a string message generated successfully.
//  r.POST(apiSubPath+"regions", util.TokenValidationMiddlewareNew(client), CreateRegionMasters(client))
// r.GET(apiSubPath+"GDSPAexams/getallvapendingapplications/:id/:id1", util.TokenValidationMiddlewareNew(client), h.GetAllGDSPAVAPendingVerifications(client))
// r.GET(apiSubPath+"directorateusers", getAllDusers(client))
// r.GET(apiSubPath+"directorateusers/:id", GetDirectorateUsersByID(client))
// r.GET(apiSubPath+"directorateusers/byemployeeid/:id", GetDusersWithEmpID(client))
// r.GET(apiSubPath+"divisions", getAllDivisions(client))
// r.GET(apiSubPath+"divisions/:id", GetDivisionsByID(client))
// r.GET(apiSubPath+"divisions/byregioncode/:id", GetDivisionsByRegionCode(client))
// r.GET(apiSubPath+"facilities", util.TokenValidationMiddlewareNew(client), getAllFacilities(client))
// r.GET(apiSubPath+"facilities/:id", util.TokenValidationMiddlewareNew(client), GetFacilityByID(client))
// r.GET(apiSubPath+"facilities/bycirclecode/:id", util.TokenValidationMiddlewareNew(client), GetFacilitiesByCircleCode(client))
// r.GET(apiSubPath+"facilities/bycircleofficeid/:id", util.TokenValidationMiddlewareNew(client), QueryDivisionsByCircleOfficeID(client))
// r.GET(apiSubPath+"facilities/bydivisioncode/:id", util.TokenValidationMiddlewareNew(client), GetFacilitiesByDivisionCode(client))
// r.GET(apiSubPath+"facilities/byregioncode/:id", util.TokenValidationMiddlewareNew(client), GetFacilitiesByRegionCode(client))
// r.GET(apiSubPath+"gdspasa/getapplicationsfacilityidyear/:id/:id1", util.TokenValidationMiddlewareNew(client), GetGDSPASApplicationsFacilityIDYear(client))
// r.GET(apiSubPath+"notification/:id", GetNotificationID(client))
// r.GET(apiSubPath+"notifications", getNotifications(client))
// r.GET(apiSubPath+"profile/:id", GetUserID(client))
// r.GET(apiSubPath+"profiles", util.TokenValidationMiddlewareNew(client), getUsers(client))
// r.GET(apiSubPath+"regions", util.TokenValidationMiddlewareNew(client), getAllRegions(client))
// r.GET(apiSubPath+"regions/:id", util.TokenValidationMiddlewareNew(client), GetRegionsByID(client))
// r.GET(apiSubPath+"regions/bycirclecode/:id", util.TokenValidationMiddlewareNew(client), GetRegionsByCircleCode(client))
// r.POST(apiSubPath+"directorateusers", InsertDirectorateUsers(client))
// r.POST(apiSubPath+"divisions", util.TokenValidationMiddlewareNew(client), CreateDivisionMasters(client))
// r.POST(apiSubPath+"facilities", util.TokenValidationMiddlewareNew(client), CreateFacilityMasters(client))
// r.POST(apiSubPath+"notification", CreateNotification(client))
// r.POST(apiSubPath+"notifyupdate/:id", updateNotification(client))
// r.POST(apiSubPath+"profile/:id", util.TokenValidationMiddlewareNew(client), updateUser(client)) //NL
// r.PUT(apiSubPath+"ipexams/halltickets", GetHallticketNumbers(client))  // returns circle wise count.
// r.PUT(apiSubPath+"ipexams/halltickets", GetHallticketNumbers(client))  // returns circle wise count.
//r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//r.GET("/gettoken/:id", gettoken(client))
//r.GET("/verify", validateToken)
//r.GET("/verify", verifytoken)
//r.GET(apiSubPath+"ActiveExamsByYear/:selectedyear", util.TokenValidationMiddlewareNew(client), h.QueryActiveExamsByExamYearWithoutCAFacilityID(client))
//r.GET(apiSubPath+"ActiveExamsByYear/:selectedyear/:circleidcontext", util.TokenValidationMiddlewareNew(client), h.QueryActiveExamsByExamYear(client))
//r.GET(apiSubPath+"IpExams/GetDetailsByOnExamCode/:Exam_Code", util.TokenValidationMiddlewareNew(client), GetDetailsBasedOnExamCode(client)) //----vasanth
//r.GET(apiSubPath+"admin/ForgotAdminPassword/:username", GetAdminPasswordByUsername(client))                                            //----vasanth dev
//r.GET(apiSubPath+"admin/ForgotAdminPassword/:username", GetAdminPasswordByUsername(client))                                            //----vasanth dev
//r.GET(apiSubPath+"centers", getCenters(client))
//r.GET(apiSubPath+"centers/:id1/:id2",
//r.GET(apiSubPath+"centers/:id1/:id2",
//r.GET(apiSubPath+"circles", util.TokenValidationMiddlewareNew(client), getAllCircles(client))
//r.GET(apiSubPath+"circles/:id", util.TokenValidationMiddlewareNew(client), GetCircleByID(client))
//r.GET(apiSubPath+"eligibilities", h.GetAllEligibilitiyCriteria(client))
//r.GET(apiSubPath+"eligibilities/:id", h.GetEligibilityByID(client))
//r.GET(apiSubPath+"employeedisbilitytypes", h.GetAllDisabilityTypes(client))
//r.GET(apiSubPath+"employeedisbilitytypes/:id", h.GetDisabilitiesByID(client))
//r.GET(apiSubPath+"employees/:id", util.TokenValidationMiddlewareNew(client), GetEmployeesByID(client))
//r.GET(apiSubPath+"exam/:id", handlers.GetExamID(client))
//r.GET(apiSubPath+"examcalendars/Details/:id", getExamCalendarsWithDetails(client))
//r.GET(apiSubPath+"exampapers/:id", GetExamPapersByID(client))
//r.GET(apiSubPath+"exams/ListExamIDsandNames", h.GetExamIDsandNames(client))
//r.GET(apiSubPath+"exams/statistics/fordirectorate/circlewise/get/:examcode/:examyear", h.GetExamStatisticsForDirectorateCircleWise(client))
//r.GET(apiSubPath+"exams/statistics/fordirectorate/get/:examcode/:selectedyear", h.GetExamApplicationsStatisticsForDirectorate(client))
//r.GET(apiSubPath+"exams/witheligibilities/byexamcode/:id", h.GetExamWithEligibility(client))
//r.GET(apiSubPath+"facilities/byPincodeNoUID/:pincode", util.TokenValidationMiddlewareUID(client), h.QueryOfficeByPincodehandle(client))
//r.GET(apiSubPath+"facilities/bycirclecodenew/:id", util.TokenValidationMiddlewareNew(client), GetNewFacilitiesByCircleCode(client))     //----dev by vasanth
//r.GET(apiSubPath+"facilities/bydivisioncodenew/:id", util.TokenValidationMiddlewareNew(client), GetNewFacilitiesByDivisionCode(client)) //---dev by vasanth
//r.GET(apiSubPath+"facilities/bynewfacilityofficeid/:id", GetNewFacilityDetailsByFacilityOfficeID(client)) //--dev by Vasanth
//r.GET(apiSubPath+"facilities/byregioncodenew/:id", util.TokenValidationMiddlewareNew(client), GetNewFacilitiesByRegionCode(client))     //--dev by Vasnth
//r.GET(apiSubPath+"getAllCircles", util.TokenValidationMiddlewareUID(client), h.GetAllCircles(client))
//r.GET(apiSubPath+"getfacilities/:id", util.TokenValidationMiddlewareNew(client), GetNewFacilityByID(client))                            //----dev by vasanth
//r.GET(apiSubPath+"getfacilitiesnew", util.TokenValidationMiddlewareNew(client), getAllFacilitiesNew(client))          //---vasanth dev
//r.GET(apiSubPath+"ipexams/get/:id", util.TokenValidationMiddlewareNew(client), h.GetExamDetailsIPByExamCode(client))
//r.GET(apiSubPath+"nodalofficer/:id", util.TokenValidationMiddlewareNew(client), GetNodalOfficerID(client))
//r.GET(apiSubPath+"nodalofficers", util.TokenValidationMiddlewareNew(client), getNodalOfficers(client))
//r.GET(apiSubPath+"pasaip/getapplicationsfacilityidyear/:id/:id1", util.TokenValidationMiddlewareNew(client), GetPASAIPApplicationsFacilityIDYear(client))
//r.GET(apiSubPath+"pmmgmtspasa/getapplicationsfacilityidyear/:id/:id1", util.TokenValidationMiddlewareNew(client), GetPMMGMTSPASAApplicationsFacilityIDYear(client))
//r.GET(apiSubPath+"psexams/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllPSPendingWithCandidateNew(client))
//r.GET(apiSubPath+"psexams/get/:id", GetExamDetailsPSByExamCode(client))
//r.GET(apiSubPath+"psexams/hallticket/get/:id1/:id2", GetHallTicketWithExamCodeEmpID(client))
//r.GET(apiSubPath+"report/getAllExamSummaryByRecommendedStatus/:examcode/:examyear/:facilityid/:type", h.GetExamSummaryByRecommendedStatus(client))
//r.GET(apiSubPath+"secret/admin/newUser/all", getAllAdminLogins(client))
//r.GET(apiSubPath+"testing/emails", SendTestEmails(client))
//r.GET(apiSubPath+"testing/emails", SendTestEmails(client))
//r.GET(apiSubPath+"totp", topgenerate)profile
//r.GET(apiSubPath+"totp", topgenerate)profile
//r.GET(apiSubPath+"users/checkbyempid/:id", CheckUserByEmpId(client))
//r.GET(apiSubPath+"users/checkbyusername/:id", CheckAdminUserByUserName(client))
//r.GET(apiSubPath+"users/getpassword/:id1", h.GetPasswordByEmpId(client))                                          //----vasanth
//r.GET(apiSubPath+"users/getpassword/:id1", h.GetPasswordByEmpId(client))                                          //----vasanth
//r.GET(apiSubPath+"users/getuserbyempid/:id", GetUsersByEmpId(client))
//r.GET(apiSubPath+"users/insertuser/:id", CreateUserfromEmpMaster(client))
//r.POST("/token", validateRefreshToken)
//r.POST(apiSubPath+"PostPreferences", util.TokenValidationMiddlewareNew(client), postPreferences(client))
//r.POST(apiSubPath+"admin/newUser/verify", VerifyLogin(client))
//r.POST(apiSubPath+"admin/newUser/verify", VerifyLogin(client))
//r.POST(apiSubPath+"admin/roles", InsertRoles(client))
//r.POST(apiSubPath+"admin/roles", InsertRoles(client))
//r.POST(apiSubPath+"admins/reset/savenewpassword", util.TokenValidationMiddlewareUID(client), h.AdminResetSaveNewPassword(client))
//r.POST(apiSubPath+"admins/reset/validateOTP", util.TokenValidationMiddlewareUID(client), h.AdminResetValidateOTP(client))
//r.POST(apiSubPath+"admins/reset/validateusername", util.TokenValidationMiddlewareUID(client), h.AdminResetValidateUserName(client))
//r.POST(apiSubPath+"adminusers/verifyloginn", util.TokenValidationMiddlewareUID(client), h.VerifyAdminLoginn(client)) //--vasanth dev
//r.POST(apiSubPath+"adminusers/verifynewUser", h.VerifyAdminLogin(client))
//r.POST(apiSubPath+"adminusers/verifyuser", h.VerifyAdminUserLogin(client))
//r.POST(apiSubPath+"adminusers/verifyuser", h.VerifyAdminUserLogin(client))
//r.POST(apiSubPath+"adminusers/verifyuserr", util.TokenValidationMiddlewareUID(client), h.VerifyAdminUserLoginn(client)) //--vasanth dev
//r.POST(apiSubPath+"circles", util.TokenValidationMiddlewareNew(client), CreateCircleMasters(client))
//r.POST(apiSubPath+"createfacilitiesnew", util.TokenValidationMiddlewareNew(client), CreateFacilitynewMasters(client)) //----vasanth dev
//r.POST(apiSubPath+"eligibilities", h.InsertEligibilityCriteria(client))
//r.POST(apiSubPath+"empmaster/submitemployee", util.TokenValidationMiddlewareUID(client), h.CreateEmployeeMaster(client)) //---K.Mohandoss
//r.POST(apiSubPath+"exam", h.CreateExam(client))
//r.POST(apiSubPath+"exams/no/approvehallticket", UpdateApprovalOfHallTicketForNO(client))
//r.POST(apiSubPath+"examupdate/:id", util.TokenValidationMiddlewareNew(client), h.Updateexam(client))
//r.POST(apiSubPath+"nodalofficer", util.TokenValidationMiddlewareNew(client), CreateNodalOfficer(client))
//r.POST(apiSubPath+"nodalofficer/:id", util.TokenValidationMiddlewareNew(client), UpdateNodalOfficer(client))
//r.POST(apiSubPath+"notification/create", util.TokenValidationMiddlewareNew(client), h.CreateexamNotification(client))
//r.POST(apiSubPath+"notification/createDraftResubmission", CreateexamNotification(client))
//r.POST(apiSubPath+"profileadmin", util.TokenValidationMiddlewareNew(client), h.CreateAdminUser(client)) //---Vasanth
//r.POST(apiSubPath+"sendsms", generateOTPAndSendSMS)
//r.POST(apiSubPath+"users/login", util.TokenValidationMiddlewareUID(client), h.VerifyUserLogin(client))
//r.POST(apiSubPath+"users/new/submit", util.TokenValidationMiddlewareUID(client), h.FirstTimeUserCreation(client))
//r.POST(apiSubPath+"users/new/update", util.TokenValidationMiddlewareUID(client), h.UpdateFirstTimeUserDetails(client))
//r.POST(apiSubPath+"users/reset/savenewpassword", util.TokenValidationMiddlewareUID(client), h.UserResetSaveNewPassword(client))
//r.POST(apiSubPath+"users/reset/validateOTP", util.TokenValidationMiddlewareUID(client), h.UserResetValidateOTP(client))
//r.POST(apiSubPath+"users/reset/validateusername", util.TokenValidationMiddlewareUID(client), h.UserResetValidateUserName(client))
//r.POST(apiSubPath+"vacancyyears/:id", UpdateVacancyYearID(client))
//r.PUT(apiSubPath+"GDSPMexams/applications/verify", util.TokenValidationMiddlewareNew(client), VerifyGDSPMApplication(client))
//r.PUT(apiSubPath+"MTSPMMGexams/halltickets", util.TokenValidationMiddlewareNew(client), h.GetMTSPMMGHallticketNumberscenter(client))                    //NL
//r.PUT(apiSubPath+"admin/ChangeAdminPassword/:username", util.TokenValidationMiddlewareNew(client), h.ChangeAdminPassword(client))                       //----vasanth dev
//r.PUT(apiSubPath+"notification/PutIssueNotification", util.TokenValidationMiddlewareNew(client), h.PutIssueNotification(client))
//r.PUT(apiSubPath+"psexams/halltickets/:id", GetHallticketNumberscenter(client)) // just returns a string message generated successfully.
