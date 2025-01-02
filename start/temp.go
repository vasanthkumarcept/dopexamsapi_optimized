package start

import (
	"context"
	//"errors"
	"fmt"
	"log"

	//	"net/http"
	"recruit/ent"

	"recruit/ent/exam_applications_ps"
	//"recruit/ent/usermaster"
	//"github.com/gin-gonic/gin"
)

// MTS to PM/MG Start

// Query PMPAExam Application with Emp ID.
// func QueryMTSPMMGExamApplicationsByEmpID(ctx context.Context, client *ent.Client, empid int64, examYear string) (*ent.Exam_Application_MTSPMMG, error) {
// 	newAppln, err := client.Exam_Application_MTSPMMG.
// 		Query().
// 		Where(
// 			exam_application_mtspmmg.EmployeeIDEQ(empid),
// 			exam_application_mtspmmg.ExamYearEQ(examYear),
// 			exam_application_mtspmmg.StatusEQ("active"),
// 		).
// 		Order(ent.Desc(exam_application_mtspmmg.FieldID)).
// 		WithCirclePrefRefMTSPMMG().
// 		First(ctx)

// 	if err != nil {
// 		return nil, fmt.Errorf("failed querying PMPA Exam Application details: %w", err)
// 	}

// 	// Extract only the desired fields from the CirclePrefRefPMPA edge
// 	var circlePrefs []*ent.Division_Choice_MTSPMMG
// 	for _, edge := range newAppln.Edges.CirclePrefRefMTSPMMG {
// 		circlePrefs = append(circlePrefs, &ent.Division_Choice_MTSPMMG{
// 			PlacePrefNo:    edge.PlacePrefNo,
// 			PlacePrefValue: edge.PlacePrefValue,
// 		})
// 	}

// 	// Update the CirclePrefRefPMPA edge with the filtered values
// 	newAppln.Edges.CirclePrefRefMTSPMMG = circlePrefs

// 	newAppln.UpdatedAt = newAppln.UpdatedAt.Truncate(24 * time.Hour)
// 	return newAppln, nil
// }

// Get Recommendations with Emp ID ..
// func QueryMTSPMMGRecommendationsByEmpId(ctx context.Context, client *ent.Client, employeeID int64) ([]*ent.RecommendationsMTSPMMGApplications, error) {
// 	//Array of exams

// 	employeeExists, err := client.RecommendationsMTSPMMGApplications.
// 		Query().
// 		Where(recommendationsmtspmmgapplications.EmployeeIDEQ(employeeID)).
// 		Exist(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed checking employee existence: %w", err)
// 	}
// 	if !employeeExists {
// 		return nil, fmt.Errorf("employee not found with ID: %d", employeeID)
// 	}

// 	records, err := client.RecommendationsMTSPMMGApplications.
// 		Query().
// 		Where(recommendationsmtspmmgapplications.EmployeeIDEQ(employeeID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to query PMPA recommendations: %w", err)
// 	}

// 	return records, nil
// }

// func SubGetMTSPMMGApplicationsFacilityIDYear(ctx context.Context, client *ent.Client, facilityID string, year string) ([]*ent.Exam_Application_MTSPMMG, int32, error) {
// 	// Array of exams

// 	if facilityID == "" || year == "" {
// 		return nil, 422, errors.New("facility ID and Examyear cannot be blank/null")
// 	}
// 	records, err := client.Exam_Application_MTSPMMG.Query().
// 		Where(exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(facilityID),
// 			exam_application_mtspmmg.ExamYearEQ(year),
// 			exam_application_mtspmmg.StatusEQ("active"),
// 		).
// 		Order(ent.Desc(exam_application_mtspmmg.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, 400, fmt.Errorf("failed querying MTS to PM/MG exams Applications: %w", err)
// 	}
// 	if len(records) == 0 {
// 		return nil, 422, fmt.Errorf("no applications for the Year %s and facility ID  %s", year, facilityID)
// 	}

// 	return records, 200, nil
// }

// // Get All CA verified records for NA
/* func QueryMTSPMMGApplicationsByCAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID string) ([]*ent.Exam_Application_MTSPMMG, error) {
	if facilityID == "" {
		return nil, fmt.Errorf("facility ID cannot be null")
	}

	records, err := client.Exam_Application_MTSPMMG.Query().
		Where(
			exam_application_mtspmmg.And(
				exam_application_mtspmmg.ApplicationStatusEQ("VerifiedByCA"),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(facilityID),
			),
		).
		WithMTSPMMGApplicationsRef().
		WithCirclePrefRefMTSPMMG(). // Include the CirclePrefRef association
		All(ctx)

	if err != nil {
		log.Println("error at PMPA Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying MTSPMMG exams Applications for CA Verified records: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("no applications found for CA verified for the Office ID %s", facilityID)
	}

	// Update CirclePrefRefMTSPMMG with the desired fields for each record
	for _, record := range records {
		// Extract only the desired fields from the CirclePrefRefMTSPMMG edge
		var circlePrefs []*ent.Division_Choice_MTSPMMG
		for _, edge := range record.Edges.CirclePrefRefMTSPMMG {
			circlePrefs = append(circlePrefs, &ent.Division_Choice_MTSPMMG{
				Group:            edge.Group,
				CadrePrefNo:      edge.CadrePrefNo,
				Cadre:            edge.Cadre,
				PostPrefNo:       edge.PostPrefNo,
				PostingPrefValue: edge.PostingPrefValue,
			})
		}

		// Update the CirclePrefRefMTSPMMG with the extracted fields
		record.Edges.CirclePrefRefMTSPMMG = circlePrefs

		// Log the updated CirclePrefRefMTSPMMG for each record
		fmt.Printf("Record ID: %d, Circle Preference Ref: %+v\n", record.ID, record.Edges.CirclePrefRefMTSPMMG)
	}

	return records, nil
}
*/

// Generate ht with centercode
// func GenerateHallticketNumberMTSPMMGwithCenterCode(ctx context.Context, client *ent.Client) (string, error) {

// 	currentTime := time.Now().Truncate(time.Second)

// 	applications, err := client.Exam_Application_MTSPMMG.
// 		Query().
// 		Where(
// 			//exam_applications_pmpa.GenerateHallTicketFlag(true),
// 			//exam_applications_pmpa.GenerateHallTicketFlagByNO(true),
// 			exam_application_mtspmmg.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_pmpa.CenterCodeEQ(ExamCenterCode),
// 			exam_application_mtspmmg.ExamCodeNEQ(0),
// 			exam_application_mtspmmg.ExamYearNEQ(""),
// 			exam_application_mtspmmg.CategoryCodeNEQ(""),
// 			exam_application_mtspmmg.CircleIDNEQ(0),
// 			//exam_application_mtspmmg.RegionIDNEQ(0),
// 			//exam_application_mtspmmg.DivisionIDNEQ(0),
// 			exam_application_mtspmmg.EmployeeIDNEQ(0),
// 			//exam_applications_pmpa.HallTicketNumberEQ(""),
// 			exam_application_mtspmmg.HallTicketGeneratedFlagNEQ(true),
// 		).
// 		Order(ent.Asc(exam_application_mtspmmg.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		return "", err
// 	}

// 	circleStats := make(map[string]int)
// 	for _, application := range applications {
// 		key := fmt.Sprintf("%d", application.CircleID)
// 		circleStats[key]++

// 		identificationNo := circleStats[key]
// 		examYear := application.ExamYear
// 		if len(examYear) >= 2 {
// 			examYear = examYear[len(examYear)-2:]
// 		}

// 		hallticketNumber := util.GenerateHallticketNumberMTSPMMG(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.CircleID,
// 			//application.RegionID,
// 			application.DivisionID,
// 			identificationNo)

// 		// Validate if the hallticket number is of 12 digits
// 		if len(hallticketNumber) != 12 {
// 			log.Printf("Hallticket Number: %s\n", hallticketNumber)
// 			continue
// 		}

// 		log.Printf("Generated hallticket number is of 12 digits: %s\n", hallticketNumber)
// 		log.Printf("Employee ID: %d\n", application.EmployeeID)

// 		_, err = application.Update().
// 			SetHallTicketNumber(hallticketNumber).
// 			SetHallTicketGeneratedFlag(true).
// 			SetHallTicketGeneratedDate(currentTime).
// 			Save(ctx)
// 		if err != nil {
// 			errMsg := fmt.Sprintf("Unable to generate Hallticket number: %s", err.Error())
// 			log.Println(errMsg)
// 			return "", errors.New(errMsg)
// 		}
// 	}

// 	// Return the success message with the count of eligible candidates
// 	successMessage := fmt.Sprintf("Hall Ticket generated successfully for %d eligible candidates", len(applications))
// 	return successMessage, nil
// }

// MTS to PM/MG End

// GDS to MTS/PM/MG Start

// func GetGDSPMExamStatisticsDOOfficeWiseLatests(ctx context.Context, client *ent.Client, examCode int32, facilityID string, examYear string) ([]map[string]interface{}, error) {
// 	// Check if exam code is valid
// 	if examCode <= 0 {
// 		log.Println("No such valid exam code exists")
// 		return nil, fmt.Errorf("No such valid exam code exists")
// 	}

// 	// Check if facilityID is provided
// 	if facilityID == "" {
// 		log.Println("Facility ID cannot be null")
// 		return nil, fmt.Errorf("Facility ID cannot be null")
// 	}

// 	// Query to get the applications from Exam_Applications_GDSPM table matching the provided facilityID
// 	applications, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			exam_applications_gdspm.NodalOfficeFacilityIDEQ(facilityID),
// 			exam_applications_gdspm.StatusEQ("active"),
// 			exam_applications_gdspm.ExamYearEQ(examYear),
// 		).
// 		Order(ent.Asc(exam_applications_gdspm.FieldEmployeeID), ent.Desc(exam_applications_gdspm.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
// 		return nil, fmt.Errorf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
// 	}

// 	// Create a map to store the latest application details for each employee
// 	latestApplications := make(map[int64]*ent.Exam_Applications_GDSPM)

// 	// Loop through the applications to find the latest application for each employee and their reporting office
// 	for _, app := range applications {
// 		employeeID := app.EmployeeID

// 		// Check if this application is the latest for the employee
// 		if latestApp, ok := latestApplications[employeeID]; !ok || app.ID > latestApp.ID {
// 			latestApplications[employeeID] = app
// 		}
// 	}

// 	// Create a map to store reporting office-wise summaries
// 	doOfficeSummaries := make(map[string]*DOOfficeWiseSummary)

// 	// Loop through the latest applications to update counts
// 	for _, app := range latestApplications {
// 		reportingOfficeID := app.ReportingOfficeID

// 		// Create and initialize doOfficeSummary outside the condition
// 		doOfficeSummary := doOfficeSummaries[reportingOfficeID]
// 		if doOfficeSummary == nil {
// 			doOfficeSummaries[reportingOfficeID] = &DOOfficeWiseSummary{
// 				ControllingOfficeFacilityID: reportingOfficeID,
// 				ControllingOfficeName:       app.ControllingOfficeName,
// 				Permitted:                   0,
// 				NotPermitted:                0,
// 				PendingWithCA:               0,
// 				PendingWithCandidate:        0,
// 				Received:                    0,
// 				HallTicketGenerated:         0,
// 				HallTicketNotGenerated:      0,
// 				UniqueEmployees:             make(map[int64]struct{}),
// 			}
// 		}

// 		if _, ok := doOfficeSummary.UniqueEmployees[app.EmployeeID]; !ok {
// 			doOfficeSummary.UniqueEmployees[app.EmployeeID] = struct{}{}
// 			doOfficeSummary.Received++

// 			if app.GenerateHallTicketFlag == nil {
// 				if app.ApplicationStatus == "PendingWithCandidate" {
// 					// For pending, check if GenerateHallTicketFlag is nil
// 					doOfficeSummary.PendingWithCandidate++
// 				} else {
// 					doOfficeSummary.PendingWithCA++
// 				}
// 			} else if *app.GenerateHallTicketFlag {
// 				doOfficeSummary.Permitted++
// 				if *&app.HallTicketNumber == ("") {
// 					doOfficeSummary.HallTicketNotGenerated++
// 				} else {
// 					doOfficeSummary.HallTicketGenerated++
// 				}
// 			} else {
// 				doOfficeSummary.NotPermitted++
// 			}
// 		}

// 	// Create an empty slice to store the final result
// 	result := []map[string]interface{}{}
// 	serialNumber := 0

// 	// Add reportingOfficeID wise counts and names to the result
// 	for _, summary := range doOfficeSummaries {
// 		serialNumber++

// 		// Display only the latest reporting office counts
// 		result = append(result, map[string]interface{}{
// 			"S.No.":                        serialNumber,
// 			"ReportingOfficeID":            summary.ControllingOfficeFacilityID,
// 			"ReportingOfficeName":          summary.ControllingOfficeName,
// 			"No: Of Applications Received": summary.Received,
// 			"No. Permitted":                summary.Permitted,
// 			"No. Not Permitted":            summary.NotPermitted,
// 			"No. Pending":                  summary.Pending,
// 			"No. Pending With Candidate":   summary.PendingWithCandidate,
// 		})
// 	}

// 	return result, nil
// }

// func getApprovalFlagForHallTicketGDSPM(client *ent.Client, circleOfficeID string) (bool, error) {
// 	circleMaster, err := client.CircleSummaryForNO.
// 		Query().
// 		Where(circlesummaryforno.CircleOfficeIdEQ(circleOfficeID)).
// 		Only(context.Background())
// 	if err != nil {
// 		return false, fmt.Errorf("failed to get CircleMaster for CircleOfficeID %v: %v", circleOfficeID, err)
// 	}

// 	return circleMaster.ApproveHallTicketGenrationPA, nil
// }

// type HallticketStatsGDSPM struct {
// 	CircleID        int32  `json:"CircleID"`
// 	StartingNumber  int    `json:"StartingNumber"`
// 	EndingNumber    int    `json:"EndingNumber"`
// 	Count           int    `json:"Count"`
// 	StartHallTicket string `json:"StartHallTicket"`
// 	EndHallTicket   string `json:"EndHallTicket"`
// }

// func generateHallticketNumberGDSPM(examCode int32, examYear string, categoryCode string, circleID int32 /*regionID int32,*/, divisionID int32, identificationNo int) string {
// 	// Generate the Hallticket Number based on the provided formatfmt.Sprintf("%d%s%d%d%d%d%04d", examCode, examYear, getFormattedCode(circleID), regionID, getFormattedCode(divisionID), categoryCode, identificationNo)
// 	hallticketNumber := fmt.Sprintf("%d%s%s%s%d%04d", examCode, examYear, getFormattedCodeGDSPM(circleID) /* regionID,*/, getFormattedCodeGDSPM(divisionID), categoryCode, identificationNo)
// 	return hallticketNumber
// }

// func getFormattedCodeGDSPM(code int32) string {
// 	// Format the code as a string with the required number of digits
// 	lastTwoDigits := code % 100
// 	return fmt.Sprintf("%02d", lastTwoDigits)
// }

// // list of reporting offices

// type HallticketResultGDSPM struct {
// 	CircleID string `json:"circleID"`
// 	Count    int    `json:"count"`
// }
// type CircleStatsGDSPM struct {
// 	CircleID        string `json:"CircleID"`
// 	HallTicketCount int    `json:"Count"`
// }

// func GetGDSPMDivisionsByCircleOfficeID(ctx context.Context, client *ent.Client, circleOfficeID string, Examyear string) ([]*ent.Exam_Applications_GDSPM, error) {
// 	// Check if the circle office ID exists in the exam_application_gdspm table.
// 	exists, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			exam_applications_gdspm.NodalOfficeFacilityIDEQ(circleOfficeID),
// 			exam_applications_gdspm.ExamYearEQ(Examyear),
// 			exam_applications_gdspm.StatusEQ("active"),
// 		).
// 		Exist(ctx)
// 	if err != nil {
// 		log.Printf("Failed to query exam_application_gdspm: %v\n", err)
// 		return nil, fmt.Errorf("failed to query exam_application_gdspm: %v", err)
// 	}
// 	if !exists {
// 		log.Printf("Circle office ID does not exist: %s\n", circleOfficeID)
// 		return nil, fmt.Errorf("circle office ID does not exist")
// 	}

// 	// Query the exam_application_gdspm table for unique records based on the provided conditions.
// 	applications, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Select(
// 			exam_applications_gdspm.FieldReportingOfficeID,
// 			exam_applications_gdspm.FieldReportingOfficeName,
// 		).
// 		Where(
// 			exam_applications_gdspm.NodalOfficeFacilityIDEQ(circleOfficeID),
// 			exam_applications_gdspm.Not(exam_applications_gdspm.GenerateHallTicketFlag(true)),
// 			exam_applications_gdspm.ExamYearEQ(Examyear),
// 			exam_applications_gdspm.StatusEQ("active"),
// 			exam_applications_gdspm.ExamCityCenterCodeIsNil(),
// 		).
// 		All(ctx)
// 	if err != nil {
// 		log.Printf("Failed to query exam_application_gdspm: %v\n", err)
// 		return nil, fmt.Errorf("failed to query exam_application_gdspm: %v", err)
// 	}

// 	// Filter and return distinct records based on reporting office ID and name.
// 	distinctApplications := make(map[string]*ent.Exam_Applications_GDSPM)
// 	for _, app := range applications {
// 		key := app.ReportingOfficeID
// 		distinctApplications[key] = app
// 	}

// 	result := make([]*ent.Exam_Applications_GDSPM, 0, len(distinctApplications))
// 	for _, app := range distinctApplications {
// 		result = append(result, app)
// 	}

// 	log.Printf("Retrieved %d distinct divisions for Circle Office ID: %s\n", len(result), circleOfficeID)

// 	// Log the applications as an array of strings
// 	appStrings := make([]string, len(result))
// 	for i, app := range result {
// 		appStrings[i] = fmt.Sprintf("Reporting Office ID: %s, Reporting Office Name: %s", app.ReportingOfficeID, app.ReportingOfficeName)
// 	}
// 	log.Printf("Applications: %+v\n", appStrings)

// 	return result, nil
// }

// // Get Recommendations with Emp ID ..
// func QueryGDSPMRecommendationsByEmpId(ctx context.Context, client *ent.Client, employeeID int64) ([]*ent.RecommendationsGDSPMApplications, error) {
// 	//Array of exams

// 	records, err := client.RecommendationsGDSPMApplications.
// 		Query().
// 		Where(recommendationsgdspmapplications.EmployeeIDEQ(employeeID)).
// 		All(ctx)
// 	if err != nil {
// 		log.Println("error querying GDSPM recommendations: ", err)
// 		return nil, fmt.Errorf("failed to query GDSPM recommendations: %w", err)
// 	}

// 	return records, nil
// }

// // Update / Verification of GDSPM Exam Application By CA
// // Update Resubmission By Candidate.
// func UpdateApplicationRemarksGDSPM(client *ent.Client, newAppln *ent.Exam_Applications_GDSPM) (*ent.Exam_Applications_GDSPM, error) {
// 	ctx := context.Background()

// 	// Check if newAppln is not nil.
// 	if newAppln == nil {
// 		return nil, fmt.Errorf("newAppln is nil")
// 	}

// 	// Check if the EmployeeID exists.
// 	oldAppln, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			exam_applications_gdspm.EmployeeIDEQ(newAppln.EmployeeID),
// 			exam_applications_gdspm.ExamYearEQ(newAppln.ExamYear),
// 			exam_applications_gdspm.StatusEQ("active"),
// 		).
// 		Order(ent.Desc(exam_applications_gdspm.FieldID)).
// 		First(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("the Applicant's details are unavailable in the database for employee_id %d", newAppln.EmployeeID)
// 	}

// 	// Insert a new record with the specified conditions.
// 	caRemarks := newAppln.CARemarks

// 	currentTime := time.Now().Truncate(time.Second)

// 	if oldAppln != nil {
// 		// Update the existing record.
// 		if oldAppln.ApplicationStatus == "VerifiedByNA" || oldAppln.ApplicationStatus == "VerifiedByCA" {
// 			log.Println("The Application is already verified By Nodal Authority/ Controlling Authority:", err)
// 			return nil, fmt.Errorf("The Application is already verified By Nodal Authority/ Controlling Authority:")
// 		}
// 		if oldAppln.ApplicationStatus == "CAVerificationPending" {
// 			if caRemarks == "InCorrect" {

// 				log.Println("Hi ! Into CAVerification Pending , Incorrect remarks")
// 				upapp1, err := oldAppln.
// 					Update().
// 					SetStatus("inactive").
// 					Save(ctx)
// 				if err != nil {
// 					log.Println("Failed to update application:", err)
// 					return nil, fmt.Errorf("failed to update application: %v", err)
// 				}
// 				fmt.Println(upapp1)

// 				updatedAppln, err := client.Exam_Applications_GDSPM.
// 					Create(). //modifeid
// 					SetApplicationStatus("PendingWithCandidate").
// 					SetCARemarks(newAppln.CARemarks).
// 					SetCAUserName(newAppln.CAUserName).
// 					SetCADate(currentTime).
// 					SetCAGeneralRemarks(newAppln.CAGeneralRemarks).
// 					SetStatus("active").
// 					SetAppliactionRemarks(newAppln.AppliactionRemarks).
// 					SetCADate(currentTime).
// 					SetUserID(newAppln.UserID).
// 					SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
// 					SetCenterId(newAppln.CenterId).
// 					SetExamName(newAppln.ExamName).
// 					SetEmployeeID(newAppln.EmployeeID).
// 					SetEmployeeName(newAppln.EmployeeName).
// 					SetDOB(newAppln.DOB).
// 					SetGender(newAppln.Gender).
// 					SetMobileNumber(newAppln.MobileNumber).
// 					SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
// 					SetEmailID(newAppln.EmailID).
// 					SetCategoryCode(newAppln.CategoryCode).
// 					SetCategoryDescription(newAppln.CategoryDescription).
// 					SetCadre(newAppln.Cadre).
// 					SetEmployeePost(newAppln.EmployeePost).
// 					SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
// 					SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
// 					SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
// 					SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
// 					SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
// 					SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
// 					SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
// 					SetReportingOfficeName(newAppln.ReportingOfficeName).
// 					SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
// 					SetEntryPostDescription(newAppln.EntryPostDescription).
// 					SetPresentPostDescription(newAppln.PresentPostDescription).
// 					SetPresentDesignation(newAppln.PresentDesignation).
// 					SetFeederPostDescription(newAppln.FeederPostDescription).
// 					SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 					SetServiceLength(newAppln.ServiceLength).
// 					SetCandidateRemarks(newAppln.CandidateRemarks).
// 					SetRemarks(newAppln.Remarks).
// 					SetDCCS(newAppln.DCCS).
// 					SetDCInPresentCadre(newAppln.DCInPresentCadre).
// 					SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
// 					SetDisabilityTypeID(newAppln.DisabilityTypeID).
// 					SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
// 					SetEntryPostCode(newAppln.EntryPostCode).
// 					SetPresentPostCode(newAppln.PresentPostCode).
// 					SetFeederPostCode(newAppln.FeederPostCode).
// 					SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 					SetDesignationID(newAppln.DesignationID).
// 					SetEducationCode(newAppln.EducationCode).
// 					SetFacilityUniqueID(newAppln.FacilityUniqueID).
// 					SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
// 					SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
// 					SetDeputationType(newAppln.DeputationType).
// 					SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
// 					SetDeputationOfficeName(newAppln.DeputationOfficeName).
// 					SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
// 					SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
// 					SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
// 					SetControllingOfficeName(newAppln.ControllingOfficeName).
// 					SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 					SetNodalOfficeName(newAppln.NodalOfficeName).
// 					SetCenterFacilityId(newAppln.CenterFacilityId).
// 					SetCentrePreference(newAppln.CentrePreference).
// 					SetPhotoPath(newAppln.PhotoPath).
// 					SetNonQualifyingService(newAppln.NonQualifyingService).
// 					SetSignaturePath(newAppln.SignaturePath).
// 					SetTempHallTicket(newAppln.TempHallTicket).
// 					SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
// 					SetDisabilityPercentage(newAppln.DisabilityPercentage).
// 					SetEducationDescription(newAppln.EducationDescription).
// 					SetExamCode(newAppln.ExamCode).
// 					SetExamShortName(newAppln.ExamShortName).
// 					SetExamYear(newAppln.ExamYear).
// 					SetCentrePreference(newAppln.CentrePreference).
// 					SetSignature(newAppln.Signature).
// 					SetPhoto(newAppln.Photo).
// 					SetApplicationNumber(newAppln.ApplicationNumber).
// 					SetApplnSubmittedDate(currentTime).
// 					SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 					SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
// 					SetWorkingOfficeName(newAppln.WorkingOfficeName).
// 					SetOptionUsed(newAppln.OptionUsed).
// 					SetRemarks(newAppln.Remarks).
// 					SetPostPreferences(newAppln.PostPreferences).
// 					SetUnitPreferences(newAppln.CadrePreferences).
// 					Save(ctx)

// 				if err != nil {
// 					return nil, fmt.Errorf("failed to update application: %v", err)
// 				}

// 				return updatedAppln, nil
// 			} else if caRemarks == "Correct" {
// 				log.Println("Hi ! Into CAVerification Pending , Correct remarks")
// 				if newAppln.Edges.GDSPMApplicationsRef == nil || len(newAppln.Edges.GDSPMApplicationsRef) == 0 {
// 					return nil, fmt.Errorf("The recommendations are mandatory")
// 				}
// 				upapp, err := oldAppln.
// 					Update().
// 					SetStatus("inactive").
// 					Save(ctx)
// 				if err != nil {
// 					log.Println("Failed to update application:", err)
// 					return nil, fmt.Errorf("failed to update application: %v", err)
// 				}
// 				fmt.Println(upapp)
// 				updatedAppln, err := client.Exam_Applications_GDSPM.
// 					Create(). //modifeid
// 					SetApplicationStatus("VerifiedByCA").

// 					//SetApplnSubmittedDate(currentTime).
// 					SetCARemarks(newAppln.CARemarks).
// 					SetCAUserName(newAppln.CAUserName).
// 					SetCADate(currentTime).
// 					SetUserID(newAppln.UserID).
// 					SetExamName(newAppln.ExamName).
// 					SetEmployeeID(newAppln.EmployeeID).
// 					SetEmployeeName(newAppln.EmployeeName).
// 					SetDOB(newAppln.DOB).
// 					SetGender(newAppln.Gender).
// 					SetMobileNumber(newAppln.MobileNumber).
// 					SetEmailID(newAppln.EmailID).
// 					SetCAGeneralRemarks(newAppln.CAGeneralRemarks).
// 					SetCategoryCode(newAppln.CategoryCode).
// 					SetCategoryDescription(newAppln.CategoryDescription).
// 					SetCadre(newAppln.Cadre).
// 					SetEmployeePost(newAppln.EmployeePost).
// 					SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
// 					SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
// 					SetCenterId(newAppln.CenterId).
// 					SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
// 					SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
// 					SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
// 					SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
// 					SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
// 					SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
// 					SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
// 					SetReportingOfficeName(newAppln.ReportingOfficeName).
// 					SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
// 					SetEntryPostDescription(newAppln.EntryPostDescription).
// 					SetNonQualifyingService(newAppln.NonQualifyingService).
// 					SetPresentPostDescription(newAppln.PresentPostDescription).
// 					SetPresentDesignation(newAppln.PresentDesignation).
// 					SetFeederPostDescription(newAppln.FeederPostDescription).
// 					SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 					SetServiceLength(newAppln.ServiceLength).
// 					SetCandidateRemarks(newAppln.CandidateRemarks).
// 					SetRemarks(newAppln.Remarks).
// 					SetDCCS(newAppln.DCCS).
// 					SetDCInPresentCadre(newAppln.DCInPresentCadre).
// 					SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
// 					SetDisabilityTypeID(newAppln.DisabilityTypeID).
// 					SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
// 					SetEntryPostCode(newAppln.EntryPostCode).
// 					SetPresentPostCode(newAppln.PresentPostCode).
// 					SetFeederPostCode(newAppln.FeederPostCode).
// 					SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 					SetDesignationID(newAppln.DesignationID).
// 					SetEducationCode(newAppln.EducationCode).
// 					SetFacilityUniqueID(newAppln.FacilityUniqueID).
// 					SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
// 					SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
// 					SetDeputationType(newAppln.DeputationType).
// 					SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
// 					SetDeputationOfficeName(newAppln.DeputationOfficeName).
// 					SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
// 					SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
// 					SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
// 					SetControllingOfficeName(newAppln.ControllingOfficeName).
// 					SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 					SetNodalOfficeName(newAppln.NodalOfficeName).
// 					SetCenterFacilityId(newAppln.CenterFacilityId).
// 					SetCentrePreference(newAppln.CentrePreference).
// 					SetPhotoPath(newAppln.PhotoPath).
// 					SetSignaturePath(newAppln.SignaturePath).
// 					SetTempHallTicket(newAppln.TempHallTicket).
// 					SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
// 					SetDisabilityPercentage(newAppln.DisabilityPercentage).
// 					SetEducationDescription(newAppln.EducationDescription).
// 					SetExamCode(newAppln.ExamCode).
// 					SetExamShortName(newAppln.ExamShortName).
// 					SetExamYear(newAppln.ExamYear).
// 					SetCentrePreference(newAppln.CentrePreference).
// 					SetSignature(newAppln.Signature).
// 					SetPhoto(newAppln.Photo).
// 					SetApplicationNumber(newAppln.ApplicationNumber).
// 					SetApplnSubmittedDate(currentTime).
// 					SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 					SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
// 					SetWorkingOfficeName(newAppln.WorkingOfficeName).
// 					SetOptionUsed(newAppln.OptionUsed).
// 					SetRemarks(newAppln.Remarks).
// 					SetPostPreferences(newAppln.CadrePreferences).
// 					SetUnitPreferences(newAppln.UnitPreferences).
// 					SetStatus("active").
// 					SetAppliactionRemarks(newAppln.AppliactionRemarks).
// 					SetGenerateHallTicketFlag(*newAppln.GenerateHallTicketFlag).
// 					Save(ctx)
// 				if err != nil {
// 					return nil, fmt.Errorf("failed to update application: %v", err)
// 				}

// 				currentTime := time.Now().Truncate(time.Second)
// 				recommendationsRef := make([]*ent.RecommendationsGDSPMApplications, len(newAppln.Edges.GDSPMApplicationsRef))
// 				for i, recommendation := range newAppln.Edges.GDSPMApplicationsRef {
// 					if recommendation == nil {
// 						return nil, fmt.Errorf("Recommendations value at index %d is nil", i)
// 					}

// 					RecommendationsRefEntity, err := client.RecommendationsGDSPMApplications.
// 						Create().
// 						SetApplicationID(updatedAppln.ID).
// 						SetEmployeeID(updatedAppln.EmployeeID).
// 						SetExamYear(updatedAppln.ExamYear).
// 						SetPost(recommendation.Post).
// 						SetEligible(recommendation.Eligible).
// 						SetVacancyYear(recommendation.VacancyYear).
// 						SetCARecommendations(recommendation.CARecommendations).
// 						SetNORecommendations(recommendation.CARecommendations).
// 						SetCAUserName(newAppln.CAUserName).     // Use newAppln.CAUserName instead of updatedAppln.CAUserName
// 						SetCARemarks(recommendation.CARemarks). // Use newAppln.CARemarks instead of updatedAppln.CARemarks
// 						SetCAUpdatedAt(currentTime).
// 						SetNOUpdatedAt(currentTime).
// 						SetApplicationStatus("VerifiedRecommendationsByCA").
// 						Save(ctx)
// 					if err != nil {
// 						return nil, fmt.Errorf("failed to save Recommendation: %v", err)
// 					}

// 					recommendationsRef[i] = RecommendationsRefEntity
// 				}

// 				updatedAppln.Update().
// 					ClearGDSPMApplicationsRef().
// 					AddGDSPMApplicationsRef(recommendationsRef...).
// 					Save(ctx)
// 				return updatedAppln, nil
// 			}
// 		}
// 		if oldAppln.ApplicationStatus == "PendingWithCandidate" {

// 			if newAppln.CARemarks != "" {
// 				return nil, fmt.Errorf("The application is pending with the candidate")
// 			}

// 			applicationNumber, err := util.GenerateApplicationNumber(client, newAppln.EmployeeID, newAppln.ExamYear, "GDSPM")
// 			if err != nil {
// 				return nil, err
// 			}
// 			// Insert a new record.

// 			upapp, err := oldAppln.
// 				Update().
// 				SetStatus("inactive").
// 				Save(ctx)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to update application: %v", err)
// 			}
// 			updatedAppln, err := client.Exam_Applications_GDSPM.
// 				Create().
// 				SetUserID(newAppln.UserID).
// 				SetEmployeeID(newAppln.EmployeeID).
// 				SetStatus("active").
// 				SetEmployeeName(newAppln.EmployeeName).
// 				SetDOB(newAppln.DOB).
// 				SetGender(newAppln.Gender).
// 				SetMobileNumber(newAppln.MobileNumber).
// 				SetEmailID(newAppln.EmailID).
// 				SetCategoryCode(newAppln.CategoryCode).
// 				SetCategoryDescription(newAppln.CategoryDescription).
// 				SetCadre(newAppln.Cadre).
// 				SetEmployeePost(newAppln.EmployeePost).
// 				SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
// 				SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
// 				SetCenterId(newAppln.CenterId).
// 				SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
// 				SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
// 				SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
// 				SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
// 				SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
// 				SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
// 				SetReportingOfficeName(newAppln.ReportingOfficeName).
// 				SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
// 				SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
// 				SetEntryPostDescription(newAppln.EntryPostDescription).
// 				SetPresentPostDescription(newAppln.PresentPostDescription).
// 				SetPresentDesignation(newAppln.PresentDesignation).
// 				SetFeederPostDescription(newAppln.FeederPostDescription).
// 				SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 				SetServiceLength(newAppln.ServiceLength).
// 				SetCandidateRemarks(newAppln.CandidateRemarks).
// 				SetDCCS(newAppln.DCCS).
// 				SetDCInPresentCadre(newAppln.DCInPresentCadre).
// 				SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
// 				SetDisabilityTypeID(newAppln.DisabilityTypeID).
// 				SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
// 				SetEntryPostCode(newAppln.EntryPostCode).
// 				SetPresentPostCode(newAppln.PresentPostCode).
// 				SetFeederPostCode(newAppln.FeederPostCode).
// 				SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 				SetDesignationID(newAppln.DesignationID).
// 				SetEducationCode(newAppln.EducationCode).
// 				SetFacilityUniqueID(newAppln.FacilityUniqueID).
// 				SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
// 				SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
// 				SetDeputationType(newAppln.DeputationType).
// 				SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
// 				SetDeputationOfficeName(newAppln.DeputationOfficeName).
// 				SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
// 				SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
// 				SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
// 				SetControllingOfficeName(newAppln.ControllingOfficeName).
// 				SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 				SetNodalOfficeName(newAppln.NodalOfficeName).
// 				SetCenterFacilityId(newAppln.CenterFacilityId).
// 				SetCentrePreference(newAppln.CentrePreference).
// 				SetPhotoPath(newAppln.PhotoPath).
// 				SetSignaturePath(newAppln.SignaturePath).
// 				SetTempHallTicket(newAppln.TempHallTicket).
// 				SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
// 				SetDisabilityPercentage(newAppln.DisabilityPercentage).
// 				SetEducationDescription(newAppln.EducationDescription).
// 				SetExamCode(newAppln.ExamCode).
// 				SetExamShortName(newAppln.ExamShortName).
// 				SetExamYear(newAppln.ExamYear).
// 				SetCentrePreference(newAppln.CentrePreference).
// 				SetSignature(newAppln.Signature).
// 				SetPhoto(newAppln.Photo).
// 				SetApplicationNumber(applicationNumber).
// 				SetApplnSubmittedDate(currentTime).
// 				SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 				SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
// 				SetWorkingOfficeName(newAppln.WorkingOfficeName).
// 				SetOptionUsed(newAppln.OptionUsed).
// 				SetRemarks(newAppln.Remarks).
// 				SetPostPreferences(newAppln.CadrePreferences).
// 				SetUnitPreferences(newAppln.UnitPreferences).
// 				SetApplicationStatus("ResubmitCAVerificationPending").
// 				Save(ctx)

// 			if err != nil {
// 				return nil, fmt.Errorf("save Exam_Applications_GDSPM: %v", err)
// 			}

// 			// For Resubmission
// 			return updatedAppln, nil
// 		}

// 		// Resubmit with CA Pending

// 		if oldAppln.ApplicationStatus == "ResubmitCAVerificationPending" {
// 			if oldAppln.ApplicationStatus == "ResubmitCAVerificationPending" {
// 				if err != nil {
// 					return nil, fmt.Errorf("failed to retrieve previous remarks: %v", err)
// 				}
// 			}
// 			if caRemarks == "InCorrect" {

// 				_, err = oldAppln.
// 					Update().
// 					SetStatus("inactive").
// 					Save(ctx)
// 				if err != nil {
// 					return nil, fmt.Errorf("failed to update application: %v", err)
// 				}
// 				updatedAppln, err := client.Exam_Applications_GDSPM.
// 					Create().
// 					SetUserID(newAppln.UserID).
// 					SetEmployeeID(newAppln.EmployeeID).
// 					SetEmployeeName(newAppln.EmployeeName).
// 					SetDOB(newAppln.DOB).
// 					SetGender(newAppln.Gender).
// 					SetMobileNumber(newAppln.MobileNumber).
// 					SetEmailID(newAppln.EmailID).
// 					SetCategoryCode(newAppln.CategoryCode).
// 					SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
// 					SetCenterId(newAppln.CenterId).
// 					SetCategoryDescription(newAppln.CategoryDescription).
// 					SetCadre(newAppln.Cadre).
// 					SetEmployeePost(newAppln.EmployeePost).
// 					SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
// 					SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
// 					SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
// 					SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
// 					SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
// 					SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
// 					SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
// 					SetReportingOfficeName(newAppln.ReportingOfficeName).
// 					SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
// 					SetEntryPostDescription(newAppln.EntryPostDescription).
// 					SetPresentPostDescription(newAppln.PresentPostDescription).
// 					SetPresentDesignation(newAppln.PresentDesignation).
// 					SetFeederPostDescription(newAppln.FeederPostDescription).
// 					SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 					SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
// 					SetServiceLength(newAppln.ServiceLength).
// 					SetCandidateRemarks(newAppln.CandidateRemarks).
// 					SetDCCS(newAppln.DCCS).
// 					SetDCInPresentCadre(newAppln.DCInPresentCadre).
// 					SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
// 					SetDisabilityTypeID(newAppln.DisabilityTypeID).
// 					SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
// 					SetEntryPostCode(newAppln.EntryPostCode).
// 					SetPresentPostCode(newAppln.PresentPostCode).
// 					SetFeederPostCode(newAppln.FeederPostCode).
// 					SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 					SetDesignationID(newAppln.DesignationID).
// 					SetEducationCode(newAppln.EducationCode).
// 					SetFacilityUniqueID(newAppln.FacilityUniqueID).
// 					SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
// 					SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
// 					SetDeputationType(newAppln.DeputationType).
// 					SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
// 					SetDeputationOfficeName(newAppln.DeputationOfficeName).
// 					SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
// 					SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
// 					SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
// 					SetControllingOfficeName(newAppln.ControllingOfficeName).
// 					SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 					SetNodalOfficeName(newAppln.NodalOfficeName).
// 					SetCenterFacilityId(newAppln.CenterFacilityId).
// 					SetCentrePreference(newAppln.CentrePreference).
// 					SetPhotoPath(newAppln.PhotoPath).
// 					SetSignaturePath(newAppln.SignaturePath).
// 					SetTempHallTicket(newAppln.TempHallTicket).
// 					SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
// 					SetDisabilityPercentage(newAppln.DisabilityPercentage).
// 					SetEducationDescription(newAppln.EducationDescription).
// 					SetExamCode(newAppln.ExamCode).
// 					SetExamShortName(newAppln.ExamShortName).
// 					SetExamYear(newAppln.ExamYear).
// 					SetCentrePreference(newAppln.CentrePreference).
// 					SetSignature(newAppln.Signature).
// 					SetPhoto(newAppln.Photo).
// 					SetApplicationNumber(newAppln.ApplicationNumber).
// 					SetApplnSubmittedDate(currentTime).
// 					SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 					SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
// 					SetWorkingOfficeName(newAppln.WorkingOfficeName).
// 					SetOptionUsed(newAppln.OptionUsed).
// 					SetRemarks(newAppln.Remarks).
// 					SetPostPreferences(newAppln.CadrePreferences).
// 					SetUnitPreferences(newAppln.UnitPreferences).
// 					SetStatus("active").
// 					SetApplicationStatus("PendingWithCandidate").
// 					Save(ctx)

// 				if err != nil {
// 					return nil, fmt.Errorf("failed to update application: %v", err)
// 				}
// 				return updatedAppln, nil

// 			} else if caRemarks == "Correct" {
// 				if oldAppln.ApplicationStatus == "ResubmitCAVerificationPending" {
// 					if newAppln.Edges.GDSPMApplicationsRef == nil || len(newAppln.Edges.GDSPMApplicationsRef) == 0 {
// 						return nil, fmt.Errorf("The recommendations are mandatory")
// 					}
// 					if err != nil {
// 						return nil, fmt.Errorf("failed to retrieve previous remarks: %v", err)
// 					}
// 				}
// 				_, err = oldAppln.
// 					Update().
// 					SetStatus("inactive").
// 					Save(ctx)
// 				if err != nil {
// 					return nil, fmt.Errorf("failed to update application: %v", err)
// 				}
// 				updatedAppln, err := client.Exam_Applications_GDSPM.
// 					Create().
// 					SetUserID(newAppln.UserID).
// 					SetEmployeeID(newAppln.EmployeeID).
// 					SetEmployeeName(newAppln.EmployeeName).
// 					SetDOB(newAppln.DOB).
// 					SetGender(newAppln.Gender).
// 					SetMobileNumber(newAppln.MobileNumber).
// 					SetEmailID(newAppln.EmailID).
// 					SetCategoryCode(newAppln.CategoryCode).
// 					SetCategoryDescription(newAppln.CategoryDescription).
// 					SetCadre(newAppln.Cadre).
// 					SetEmployeePost(newAppln.EmployeePost).
// 					SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
// 					SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
// 					SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
// 					SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
// 					SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
// 					SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
// 					SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
// 					SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
// 					SetReportingOfficeName(newAppln.ReportingOfficeName).
// 					SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
// 					SetCenterId(newAppln.CenterId).
// 					SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
// 					SetEntryPostDescription(newAppln.EntryPostDescription).
// 					SetPresentPostDescription(newAppln.PresentPostDescription).
// 					SetPresentDesignation(newAppln.PresentDesignation).
// 					SetFeederPostDescription(newAppln.FeederPostDescription).
// 					SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 					SetServiceLength(newAppln.ServiceLength).
// 					SetCandidateRemarks(newAppln.CandidateRemarks).
// 					SetDCCS(newAppln.DCCS).
// 					SetDCInPresentCadre(newAppln.DCInPresentCadre).
// 					SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
// 					SetDisabilityTypeID(newAppln.DisabilityTypeID).
// 					SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
// 					SetEntryPostCode(newAppln.EntryPostCode).
// 					SetPresentPostCode(newAppln.PresentPostCode).
// 					SetFeederPostCode(newAppln.FeederPostCode).
// 					SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 					SetDesignationID(newAppln.DesignationID).
// 					SetEducationCode(newAppln.EducationCode).
// 					SetFacilityUniqueID(newAppln.FacilityUniqueID).
// 					SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
// 					SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
// 					SetDeputationType(newAppln.DeputationType).
// 					SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
// 					SetDeputationOfficeName(newAppln.DeputationOfficeName).
// 					SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
// 					SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
// 					SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
// 					SetControllingOfficeName(newAppln.ControllingOfficeName).
// 					SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 					SetNodalOfficeName(newAppln.NodalOfficeName).
// 					SetCenterFacilityId(newAppln.CenterFacilityId).
// 					SetCentrePreference(newAppln.CentrePreference).
// 					SetPhotoPath(newAppln.PhotoPath).
// 					SetSignaturePath(newAppln.SignaturePath).
// 					SetTempHallTicket(newAppln.TempHallTicket).
// 					SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
// 					SetDisabilityPercentage(newAppln.DisabilityPercentage).
// 					SetEducationDescription(newAppln.EducationDescription).
// 					SetExamCode(newAppln.ExamCode).
// 					SetExamShortName(newAppln.ExamShortName).
// 					SetExamYear(newAppln.ExamYear).
// 					SetCentrePreference(newAppln.CentrePreference).
// 					SetSignature(newAppln.Signature).
// 					SetPhoto(newAppln.Photo).
// 					SetApplicationNumber(newAppln.ApplicationNumber).
// 					SetApplnSubmittedDate(currentTime).
// 					SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 					SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
// 					SetWorkingOfficeName(newAppln.WorkingOfficeName).
// 					SetOptionUsed(newAppln.OptionUsed).
// 					SetRemarks(newAppln.Remarks).
// 					SetPostPreferences(newAppln.CadrePreferences).
// 					SetUnitPreferences(newAppln.UnitPreferences).
// 					SetStatus("active").
// 					SetApplicationStatus("VerifiedByCA").
// 					Save(ctx)

// 				if err != nil {
// 					return nil, fmt.Errorf("failed to update application: %v", err)
// 				}
// 				// Insert into recommendations.
// 				// Save the Recommendation records.
// 				recommendationsRef := make([]*ent.RecommendationsGDSPMApplications, len(newAppln.Edges.GDSPMApplicationsRef))
// 				for i, recommendation := range newAppln.Edges.GDSPMApplicationsRef {
// 					if recommendation == nil {
// 						return nil, fmt.Errorf("Recommendations value at index %d is nil", i)
// 					}

// 					RecommendationsRefEntity, err := client.RecommendationsGDSPMApplications.
// 						Create().
// 						SetApplicationID(updatedAppln.ID).
// 						SetEmployeeID(updatedAppln.EmployeeID).
// 						SetExamYear(updatedAppln.ExamYear).
// 						SetPost(recommendation.Post).
// 						SetEligible(recommendation.Eligible).
// 						SetVacancyYear(recommendation.VacancyYear).
// 						SetCARecommendations(recommendation.CARecommendations).
// 						SetNORecommendations(recommendation.CARecommendations).
// 						SetCAUserName(newAppln.CAUserName). //
// 						SetCARemarks(recommendation.CARemarks).
// 						SetCAUpdatedAt(currentTime).
// 						SetNOUpdatedAt(currentTime).
// 						SetApplicationStatus("VerifiedRecommendationsByCA").
// 						//

// 						Save(ctx)

// 					if err != nil {
// 						return nil, fmt.Errorf("failed to save Recommendation: %v", err)
// 					}

// 					recommendationsRef[i] = RecommendationsRefEntity
// 				}

// 				updatedAppln.Update().
// 					ClearGDSPMApplicationsRef().
// 					AddGDSPMApplicationsRef(recommendationsRef...).
// 					Save(ctx)
// 				return updatedAppln, nil
// 			}

// 		}

// 	}
// 	return oldAppln, nil
// }

// func SubGetGDSPMMGMTSApplicationsFacilityIDYear(ctx context.Context, client *ent.Client, facilityID string, year string) ([]*ent.Exam_Applications_GDSPM, int32, error) {
// 	// Array of exams

// 	if facilityID == "" || year == "" {
// 		return nil, 422, errors.New("facility ID and Examyear cannot be blank/null")
// 	}
// 	records, err := client.Exam_Applications_GDSPM.Query().
// 		Where(exam_applications_gdspm.ControllingOfficeFacilityIDEQ(facilityID),
// 			exam_applications_gdspm.ExamYearEQ(year),
// 			exam_applications_gdspm.StatusEQ("active"),
// 		).
// 		Order(ent.Desc(exam_applications_gdspm.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, 400, fmt.Errorf("failed querying GDS to PM/MG/MTS exams Applications: %w", err)
// 	}
// 	if len(records) == 0 {
// 		return nil, 422, fmt.Errorf("no applications for the Year %s and facility ID  %s", year, facilityID)
// 	}

// 	return records, 200, nil
// }

// func GetGDSPMApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64) (string, error) {
// 	application, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			exam_applications_gdspm.EmployeeIDEQ(employeeID),
// 			exam_applications_gdspm.ApplicationStatusEQ("PendingWithCandidate"),
// 		).
// 		Order(ent.Desc(exam_applications_gdspm.FieldID)).
// 		First(ctx)

// 	if err != nil {
// 		return "", fmt.Errorf("failed to retrieve the GDS to PA Application: %v", err)
// 	}

// 	return application.AppliactionRemarks, nil
// }

// func getGDSPMInputRecordByVacancyYear(inputRecords []*ent.RecommendationsGDSPMApplications, vacancyYear int32) *ent.RecommendationsGDSPMApplications {
// 	// Find the corresponding input record based on vacancy year
// 	for _, record := range inputRecords {
// 		if record.VacancyYear == vacancyYear {
// 			return record
// 		}
// 	}
// 	return nil
// }

// UpdateNodalRecommendationsByEmpID updates the recommendations for a given employee ID
/*func UpdateNodalRecommendationsByEmpID(client *ent.Client, empID int64, newRecommendations []*ent.RecommendationsGDSPMApplications) ([]*ent.RecommendationsGDSPMApplications, error) {
	ctx := context.Background()

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		// Handle the error, such as logging or returning an error
		log.Printf("Error loading location: %v", err)
		return nil, err
	}

	currentTime := time.Now().In(loc).Truncate(time.Second)
	// Check if empID exists
	exists, err := client.RecommendationsGDSPMApplications.Query().
		Where(recommendationsgdspmapplications.EmployeeIDEQ(empID)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check if employee with ID %d exists: %v", empID, err)
	}
	if !exists {
		return nil, fmt.Errorf("employee with ID %d does not exist", empID)
	}

	// Retrieve all records for the employee ID
	records, err := client.RecommendationsGDSPMApplications.Query().
		Where(recommendationsgdspmapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records for employee with ID %d: %v", empID, err)
	}

	if len(newRecommendations) == 0 {
		return nil, fmt.Errorf("input recommendations are empty")
	}

	// Update the records for each vacancy year
	for _, record := range records {
		vacancyYear := record.VacancyYear
		inputRecord := getInputRecordByVacancyYear(newRecommendations, int32(vacancyYear))

		//if inputRecord != nil && inputRecord.NORemarks != "" {
		// Update the NO_Recommendations field and set the ApplicationStatus to "RecommendedByNO"
		record.Update().
			SetNORecommendations(inputRecord.NORecommendations).
			SetNOUserName(inputRecord.NOUserName).
			SetNORemarks(inputRecord.NORemarks).
			SetApplicationStatus("VerifiedRecommendationsByNO").
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to save updated record for vacancy year %d: %v", vacancyYear, err)
		}

		// Log the input record
		//log.Printf("Input record for Vacancy Year %d: %+v", vacancyYear, inputRecord)
		//} else {
		// Log if there is no matching input record or NORemarks is null
		//log.Printf("No update for Vacancy Year %d", vacancyYear)
		//}
	}

	// Query the RecommendationsGDSPMApplications table for the specific employee
	record, err := client.RecommendationsGDSPMApplications.Query().
		Where(
			recommendationsgdspmapplications.EmployeeIDEQ(empID),
			//recommendationsgdspmapplications.ApplicationStatusEQ("VerifiedRecommendationsByNO"),
		).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve RecommendedByNO record: %v", err)
	}

	if record != nil {
		// Retrieve the corresponding Exam_Applications_GDSPM record using edges
		applicationRecord, err := client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.EmployeeIDEQ(empID),
				exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve Exam_Applications_GDSPM record: %v", err)
		}

		// Update the Exam_Applications_GDSPM record
		_, err = applicationRecord.Update().
			SetApplicationStatus("VerifiedByNA").
			//SetApplnSubmittedDate(currentTime).
			//SetNARemarks(record.NORemarks).
			SetNAUserName(record.NOUserName).
			SetNADate(currentTime).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update application status: %v", err)
		}
	}

	// Retrieve all records for the employee ID
	recordsupdated, err := client.RecommendationsGDSPMApplications.Query().
		Where(recommendationsgdspmapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records for employee with ID after updation %d: %v", empID, err)
	}

	return recordsupdated, nil
}*/

// Get CA Pending with EmpID
// func QueryGDSPMApplicationsByCAPendingByEmpID(ctx context.Context, client *ent.Client, empID int64) ([]*ent.Exam_Applications_GDSPM, error) {
// 	// Check if employee ID exists
// 	employeeExists, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			exam_applications_gdspm.EmployeeIDEQ(empID),
// 			exam_applications_gdspm.Or(
// 				exam_applications_gdspm.ApplicationStatusEQ("CAVerificationPending"),
// 				exam_applications_gdspm.ApplicationStatusEQ("ResubmitCAVerificationPending"),
// 			),
// 		).
// 		WithCirclePrefRefGDSPM().
// 		WithGDSPMApplicationsRef().
// 		Exist(ctx)
// 	if err != nil {
// 		log.Println("error checking employee existence: ", err)
// 		return nil, fmt.Errorf("failed checking employee existence: %w", err)
// 	}
// 	if !employeeExists {
// 		return nil, fmt.Errorf("employee not found with ID: or the verification is not pending with CA %d", empID)
// 	}

// 	// Retrieve the record
// 	record, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			exam_applications_gdspm.EmployeeIDEQ(empID),
// 			exam_applications_gdspm.Or(
// 				exam_applications_gdspm.ApplicationStatusEQ("CAVerificationPending"),
// 				exam_applications_gdspm.ApplicationStatusEQ("ResubmitCAVerificationPending"),
// 			),
// 		).
// 		WithGDSPMApplicationsRef().
// 		WithCirclePrefRefGDSPM().
// 		All(ctx)
// 	if err != nil {
// 		log.Println("error at GDSPM Exam Applications fetching: ", err)
// 		return nil, fmt.Errorf("failed querying GDSPM exams Applications: %w", err)
// 	}

// 	//log.Println("CA pending records returned: ", record)
// 	return record, nil
// }

// Get Exams by Exam Code.
/* func QueryExamsGDSPMByExamNameCode(ctx context.Context, client *ent.Client, examNameCode string) (*ent.Exam_PM, error) {
	// Check if examNameCode is empty
	if examNameCode == "" {
		return nil, fmt.Errorf("Please provide exam name code")
	}

	u, err := client.Exam_PM.Query().
		Where(exam_pm.ExamNameCode(examNameCode)).
		Only(ctx)
	if err != nil {
		log.Println("error at getting Exam_PM: ", err)
		return nil, fmt.Errorf("failed querying Exam_PM: %w", err)
	}
	log.Println("Exam_PM details returned: ", u)
	return u, nil
} */

// Generate Hall Ticket Numbers return array with stng & eng nos.
/*func GenerateHallticketNumber(ctx context.Context, client *ent.Client) ([]HallticketStats, error) {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Printf("Error loading location: %v", err)
		return nil, err
	}

	currentTime := time.Now().In(loc).Truncate(time.Second)

	applications, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.GenerateHallTicketFlag(true),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.CenterCodeNEQ(0),
			exam_applications_gdspm.ExamCodeNEQ(0),
			exam_applications_gdspm.ExamYearNEQ(""),
			exam_applications_gdspm.CategoryCodeNEQ(0),
			exam_applications_gdspm.CircleIDNEQ(0),
			exam_applications_gdspm.RegionIDNEQ(0),
			exam_applications_gdspm.DivisionIDNEQ(0),
			exam_applications_gdspm.EmployeeIDNEQ(0),
		).
		Order(ent.Desc(exam_applications_gdspm.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	circleStats := make(map[string]HallticketStats)
	for _, application := range applications {
		circleID := application.CircleID
		regionID := application.RegionID
		divisionID := application.DivisionID
		key := fmt.Sprintf("%d-%d-%d", circleID, regionID, divisionID)

		stats, exists := circleStats[key]
		if !exists {
			// Reset the serial number for each unique combination of circle, region, and division
			stats.StartingNumber = 1
		}

		identificationNo := stats.StartingNumber + stats.Count
		examYear := application.ExamYear
		if len(examYear) >= 2 {
			examYear = examYear[len(examYear)-2:]
		}

		hallticketNumber := generateHallticketNumber(
			application.ExamCode,
			examYear,
			application.CategoryCode,
			circleID,
			regionID,
			divisionID,
			identificationNo)

		// Validate if the hallticket number is of 13 digits
		if len(hallticketNumber) != 13 {
			log.Printf("Hallticket Number: %s\n", hallticketNumber)
			continue
		}

		log.Printf("Generated hallticket number is of 13 digits: %s\n", hallticketNumber)
		log.Printf("Employee ID: %d\n", application.EmployeeID)

		_, err = application.Update().
			SetHallTicketNumber(hallticketNumber).
			SetHallTicketGeneratedFlag(true).
			SetHallTicketGeneratedDate(currentTime).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		 // Update the CircleID field in HallticketStats
		 stats.CircleID = circleID

		stats.Count++
		stats.EndingNumber = stats.StartingNumber + stats.Count - 1
		stats.StartHallTicket = generateHallticketNumber(
			application.ExamCode,
			examYear,
			application.CategoryCode,
			circleID,
			regionID,
			divisionID,
			stats.StartingNumber)
		stats.EndHallTicket = generateHallticketNumber(
			application.ExamCode,
			examYear,
			application.CategoryCode,
			circleID,
			regionID,
			divisionID,
			stats.EndingNumber)

		circleStats[key] = stats
	}

	statsSlice := make([]HallticketStats, 0, len(circleStats))
	for _, stats := range circleStats {
		statsSlice = append(statsSlice, stats)
	}

	return statsSlice, nil
}*/

// Generate Hall Ticket Numbers and return JSON array of CircleID and count.

// func GenerateHallticketNumberGDSPM(ctx context.Context, client *ent.Client) ([]CircleStatsGDSPM, error) {
// 	// loc, err := time.LoadLocation("Asia/Kolkata")
// 	// if err != nil {
// 	// 	log.Printf("Error loading location: %v", err)
// 	// 	return nil, err
// 	// }

// 	// currentTime := time.Now().In(loc).Truncate(time.Second)
// 	currentTime := time.Now().Truncate(time.Second)

// 	applications, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			exam_applications_gdspm.GenerateHallTicketFlag(true),
// 			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_gdspm.ExamCityCenterCode(0),
// 			//exam_applications_gdspm.CenterCodeEQ(ExamCenterCode),
// 			exam_applications_gdspm.ExamCodeNEQ(0),
// 			exam_applications_gdspm.ExamYearNEQ(""),
// 			exam_applications_gdspm.CategoryCodeNEQ(""),
// 			/* 			exam_applications_gdspm.CircleIDNEQ(0),
// 			   			exam_applications_gdspm.RegionIDNEQ(0),
// 			   			exam_applications_gdspm.DivisionIDNEQ(0), */
// 			exam_applications_gdspm.EmployeeIDNEQ(0),
// 		).
// 		Order(ent.Desc(exam_applications_gdspm.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	circleStats := make(map[string]int)
// 	for _, application := range applications {
// 		key := fmt.Sprintf("%d", application.CircleID)
// 		circleStats[key]++

// 		identificationNo := circleStats[key]
// 		examYear := application.ExamYear
// 		if len(examYear) >= 2 {
// 			examYear = examYear[len(examYear)-2:]
// 		}

// 		hallticketNumber := generateHallticketNumberGDSPM(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.CircleID,
// 			//application.RegionID,
// 			application.DivisionID,
// 			identificationNo)

// 		// Validate if the hallticket number is of 12 digits
// 		if len(hallticketNumber) != 12 {
// 			log.Printf("Hallticket Number: %s\n", hallticketNumber)
// 			continue
// 		}

// 		log.Printf("Generated hallticket number is of 12 digits: %s\n", hallticketNumber)
// 		log.Printf("Employee ID: %d\n", application.EmployeeID)

// 		_, err = application.Update().
// 			SetHallTicketNumber(hallticketNumber).
// 			SetHallTicketGeneratedFlag(true).
// 			SetHallTicketGeneratedDate(currentTime).
// 			Save(ctx)
// 		if err != nil {
// 			errMsg := fmt.Sprintf("Unable to generate Hallticket number: %s", err.Error())
// 			log.Println(errMsg)
// 			return nil, errors.New(errMsg)
// 		}
// 	}

// 	// Convert the circleStats map to the desired JSON output
// 	statsSlice := make([]CircleStatsGDSPM, 0, len(circleStats))
// 	for key, count := range circleStats {
// 		statsSlice = append(statsSlice, CircleStatsGDSPM{CircleID: key, HallTicketCount: count})
// 	}

// 	return statsSlice, nil
// }

// generatew ht's and return as string
// func GenerateHallticketGDSPMReturnStringMessage(ctx context.Context, client *ent.Client) (string, error) {
// func GenerateHallticketGDSPMReturnStringMessage(ctx context.Context, client *ent.Client) ([]HallticketStatsGDSPM, error) {
// 	// loc, err := time.LoadLocation("Asia/Kolkata")
// 	// if err != nil {
// 	// 	log.Printf("Error loading location: %v", err)
// 	// 	return nil, err
// 	// }

// 	// currentTime := time.Now().In(loc).Truncate(time.Second)
// 	currentTime := time.Now().Truncate(time.Second)

// 	applications, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			exam_applications_gdspm.GenerateHallTicketFlag(true),
// 			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_gdspm.ExamCityCenterCodeNEQ(0),
// 			exam_applications_gdspm.ExamCodeNEQ(0),
// 			exam_applications_gdspm.ExamYearNEQ(""),
// 			exam_applications_gdspm.CategoryCodeNEQ(""),
// 			/* 			exam_applications_gdspm.CircleIDNEQ(0),
// 			   			exam_applications_gdspm.RegionIDNEQ(0),
// 			   			exam_applications_gdspm.DivisionIDNEQ(0), */
// 			exam_applications_gdspm.EmployeeIDNEQ(0),
// 		).
// 		Order(ent.Desc(exam_applications_gdspm.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	circleStats := make(map[string]HallticketStatsGDSPM)
// 	for _, application := range applications {
// 		circleID := strconv.Itoa(int(application.CircleID))
// 		regionID := strconv.Itoa(int(application.RegionID))
// 		divisionID := strconv.Itoa(int(application.DivisionID))
// 		key := circleID + "-" + regionID + "-" + divisionID

// 		stats, exists := circleStats[key]
// 		if !exists {
// 			// Reset the serial number for each unique combination of CircleID, RegionID, and DivisionID
// 			stats.StartingNumber = 1
// 		}

// 		identificationNo := stats.StartingNumber + stats.Count
// 		examYear := application.ExamYear
// 		if len(examYear) >= 2 {
// 			examYear = examYear[len(examYear)-2:]
// 		}

// 		hallticketNumber := generateHallticketNumberGDSPM(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.CircleID,
// 			//application.RegionID,
// 			application.DivisionID,
// 			identificationNo)

// 		// Validate if the hallticket number is of 12 digits
// 		if len(hallticketNumber) != 12 {
// 			log.Printf("Generated hallticket number is not of 12 digits. Skipping application with ID: %d\n", application.EmployeeID)
// 			continue
// 		}

// 		log.Printf("Generated hallticket number is of 12 digits: %s\n", hallticketNumber)
// 		//log.Printf("Application Details: ExamCode: %d, ExamYear: %s, CategoryCode: %d, CircleID: %d, RegionID: %d, DivisionID: %d\n",
// 		//	application.ExamCode, application.ExamYear, application.CategoryCode,
// 		//	application.CircleID, application.RegionID, application.DivisionID)
// 		//log.Printf("Employee ID: %d\n", application.EmployeeID)

// 		_, err = application.Update().
// 			SetHallTicketNumber(hallticketNumber).
// 			SetHallTicketGeneratedFlag(true).
// 			SetHallTicketGeneratedDate(currentTime).
// 			Save(ctx)
// 		if err != nil {
// 			return nil, err
// 		}

// 		stats.Count++
// 		stats.EndingNumber = stats.StartingNumber + stats.Count - 1
// 		stats.StartHallTicket = generateHallticketNumberGDSPM(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.CircleID,
// 			//application.RegionID,
// 			application.DivisionID,
// 			stats.StartingNumber)
// 		stats.EndHallTicket = generateHallticketNumberGDSPM(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.CircleID,
// 			//application.RegionID,
// 			application.DivisionID,
// 			stats.EndingNumber)

// 		circleStats[key] = stats
// 	}

// 	statsSlice := make([]HallticketStatsGDSPM, 0, len(circleStats))
// 	for _, stats := range circleStats {
// 		statsSlice = append(statsSlice, stats)
// 	}

// 	return statsSlice, nil
// }

// Generate ht with centercode
/* func GenerateHallticketNumberGDSPMwithCenterCode(ctx context.Context, client *ent.Client) (string, error) {
/*if ExamCenterCode == 0 {
	return "", errors.New("Please provide a valid input Exam Center")
}

// Check if the Exam Center exists in the database
exists, err := client.Exam_Applications_GDSPM.
	Query().
	Where(exam_applications_gdspm.CenterCodeEQ(ExamCenterCode)).
	Exist(ctx)
if err != nil {
	return "", err
}
if !exists {
	return "", fmt.Errorf("Exam Center with code %d does not exist", ExamCenterCode)
}*/
// loc, err := time.LoadLocation("Asia/Kolkata")
// if err != nil {
// 	log.Printf("Error loading location: %v", err)
// 	return "", err
// }

// currentTime := time.Now().In(loc).Truncate(time.Second)
// 	currentTime := time.Now().Truncate(time.Second)

// 	applications, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			//exam_applications_gdspm.GenerateHallTicketFlag(true),
// 			//exam_applications_gdspm.GenerateHallTicketFlagByNO(true),
// 			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_gdspm.CenterCodeEQ(ExamCenterCode),
// 			exam_applications_gdspm.ExamCodeNEQ(0),
// 			exam_applications_gdspm.ExamYearNEQ(""),
// 			exam_applications_gdspm.CategoryCodeNEQ(""),
// 			exam_applications_gdspm.CircleIDNEQ(0),
// 			//exam_applications_gdspm.RegionIDNEQ(0),
// 			//exam_applications_gdspm.DivisionIDNEQ(0),
// 			exam_applications_gdspm.EmployeeIDNEQ(0),
// 			//exam_applications_gdspm.HallTicketNumberEQ(""),
// 			exam_applications_gdspm.HallTicketGeneratedFlagNEQ(true),
// 		).
// 		Order(ent.Asc(exam_applications_gdspm.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		return "", err
// 	}

// 	circleStats := make(map[string]int)
// 	for _, application := range applications {
// 		key := fmt.Sprintf("%d", application.CircleID)
// 		circleStats[key]++

// 		identificationNo := circleStats[key]
// 		examYear := application.ExamYear
// 		if len(examYear) >= 2 {
// 			examYear = examYear[len(examYear)-2:]
// 		}

// 		hallticketNumber := generateHallticketNumberGDSPM(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.CircleID,
// 			//application.RegionID,
// 			application.DivisionID,
// 			identificationNo)

// 		// Validate if the hallticket number is of 12 digits
// 		if len(hallticketNumber) != 12 {
// 			log.Printf("Hallticket Number: %s\n", hallticketNumber)
// 			continue
// 		}

// 		log.Printf("Generated hallticket number is of 12 digits: %s\n", hallticketNumber)
// 		log.Printf("Employee ID: %d\n", application.EmployeeID)

// 		_, err = application.Update().
// 			SetHallTicketNumber(hallticketNumber).
// 			SetHallTicketGeneratedFlag(true).
// 			SetHallTicketGeneratedDate(currentTime).
// 			Save(ctx)
// 		if err != nil {
// 			errMsg := fmt.Sprintf("Unable to generate Hallticket number: %s", err.Error())
// 			log.Println(errMsg)
// 			return "", errors.New(errMsg)
// 		}
// 	} */

// 	// Return the success message with the count of eligible candidates
// 	successMessage := fmt.Sprintf("Hall Ticket generated successfully for %d eligible candidates", len(applications))
// 	return successMessage, nil
// }

// Count of details based on Reporting Offices

// Generate Hall Ticket Flag .../*
/*func ApproveHallTicketGenerationByNO(client *ent.Client, applicationRecord *ent.Exam_Applications_GDSPM) (string, error) {
	ctx := context.Background()

	// Check if applicationRecord is nil
	if applicationRecord == nil {
		return "", fmt.Errorf("please provide exam code, Facility ID, and Approval Flag for approving. They are mandatory")
	}

	// Check if ExamCode is nil or invalid
	if applicationRecord.ExamCode <= 0 {
		log.Println("No such valid exam code exists")
		return "", fmt.Errorf("No such valid exam code exists")
	}

	if applicationRecord.NodalOfficeID == "" {
		return "", fmt.Errorf("Please provide Nodal Officer's office ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_GDSPM table
	count, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.NodalOfficeIDEQ(applicationRecord.NodalOfficeID),
			exam_applications_gdspm.GenerateHallTicketFlag(true),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.CenterCodeNEQ(0),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", applicationRecord.NodalOfficeID)
		return "", fmt.Errorf("No valid applications available for the circle")
	}
	// Perform the update to set GenerateHallTicketFlagByNO for eligible candidates
	_, err = client.Exam_Applications_GDSPM.
		Update().
		Where(
			exam_applications_gdspm.NodalOfficeIDEQ(applicationRecord.NodalOfficeID),
			exam_applications_gdspm.GenerateHallTicketFlag(true),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.CenterCodeNEQ(0),
		).
		SetGenerateHallTicketFlagByNO(true).
		Save(ctx)
	if err != nil {
		log.Printf("Failed to update applications: %v", err)
		return "", fmt.Errorf("Failed to update applications: %v", err)
	}

	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", applicationRecord.CircleName), nil
}*/

//Here it sets in Exam appl table..
/*func ApproveHallTicketGenerationByNOForGDSPMExam(client *ent.Client, examCode int32, facilityID string, approveHallTicket bool) (string, error) {
	ctx := context.Background()

	// Check if ExamCode is nil or invalid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return "", fmt.Errorf("No such valid exam code exists")
	}

	if facilityID == "" {
		return "", fmt.Errorf("Please provide Facility ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_PS table
	count, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.NodalOfficeIDEQ(facilityID),
			exam_applications_gdspm.GenerateHallTicketFlag(true),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.CenterCodeNEQ(0),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", facilityID)
		return "", fmt.Errorf("No valid applications available for the circle")
	}
	// Perform the update to set GenerateHallTicketFlagByNO for eligible candidates
	_, err = client.Exam_Applications_GDSPM.
		Update().
		Where(
			exam_applications_gdspm.NodalOfficeIDEQ(facilityID),
			exam_applications_gdspm.GenerateHallTicketFlag(true),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.CenterCodeNEQ(0),
		).
		//SetGenerateHallTicketFlagByNO(approveHallTicket).
		Save(ctx)
	if err != nil {
		log.Printf("Failed to update applications: %v", err)
		return "", fmt.Errorf("Failed to update applications: %v", err)
	}

	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", facilityID), nil
}*/

// Here it sets in circlemaster table
// func ApproveHallTicketGenerationByNOForGDSPMExam(client *ent.Client, examCode int32, facilityID string, approveHallTicket bool) (string, error) {
// 	ctx := context.Background()

// 	// Check if ExamCode is nil or invalid
// 	if examCode <= 0 {
// 		log.Println("No such valid exam code exists")
// 		return "", fmt.Errorf("No such valid exam code exists")
// 	}

// 	if facilityID == "" {
// 		return "", fmt.Errorf("Please provide Facility ID")
// 	}

// 	// Check if circleOfficeID exists in Exam_Applications_PS table
// 	count, err := client.Exam_Applications_GDSPM.
// 		Query().
// 		Where(
// 			exam_applications_gdspm.NodalOfficeIDEQ(facilityID),
// 			exam_applications_gdspm.GenerateHallTicketFlag(true),
// 			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_gdspm.CenterCodeNEQ(0),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
// 		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
// 	}

// 	if count == 0 {
// 		log.Printf("No valid applications available for the circle: %s", facilityID)
// 		return "", fmt.Errorf("No valid applications available for the circle")
// 	}

// 	// Get the CircleMaster entity based on the Facility ID
// 	circleMaster, err := client.CircleSummaryForNO.
// 		Query().
// 		Where(circlesummaryforno.CircleOfficeIdEQ(facilityID)).
// 		Only(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve CircleMaster: %v", err)
// 		return "", fmt.Errorf("Failed to retrieve CircleMaster: %v", err)
// 	}

// 	// Update the GenerateHallTicketFlagByNO field in the CircleMaster entity
// 	circleMaster = circleMaster.
// 		Update().
// 		SetApproveHallTicketGenrationPA(approveHallTicket).
// 		SaveX(ctx)

// 	// Update the GenerateHallTicketFlagByNO in Exam_Applications_GDSPM table
// 	_, err = client.Exam_Applications_GDSPM.
// 		Update().
// 		Where(
// 			exam_applications_gdspm.NodalOfficeIDEQ(facilityID),
// 			exam_applications_gdspm.GenerateHallTicketFlag(true),
// 			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_gdspm.CenterCodeNEQ(0),
// 		).
// 		SetGenerateHallTicketFlagByNO(approveHallTicket).
// 		Save(ctx)
// 	if err != nil {
// 		log.Printf("Failed to update applications: %v", err)
// 		return "", fmt.Errorf("Failed to update applications: %v", err)
// 	}
// 	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", facilityID), nil
// }

/*

func GetExamApplicatonsPreferenenceCityWiseStatssss(ctx context.Context, client *ent.Client, cityPreference string) ([]ExamStats, error) {
	var result []ExamStats

	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.CentrePreferenceEQ(cityPreference),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.HallTicketNumberNEQ(""),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	groupedApplications := make(map[string]map[string]ExamStats)

	for _, examApplication := range applications {
		circleName := examApplication.CircleName
		reportingOfficeName := examApplication.ReportingOfficeName
		centerCode := examApplication.CenterCode

		if _, ok := groupedApplications[circleName]; !ok {
			groupedApplications[circleName] = make(map[string]ExamStats)
		}

		// Check if ExamStats for the reporting office already exists
		stats, ok := groupedApplications[circleName][reportingOfficeName]
		if !ok {
			stats = ExamStats{
				CircleName:               circleName,
				ReportingOfficeName:      reportingOfficeName,
				NoOfCandidatesChosenCity: 0,
				NoOfCandidatesAlloted:    0,
			}
		}

		//Increment counts based on center code
		if examApplication.CentrePreference == cityPreference {

			//if centerCode == 0 {
			stats.NoOfCandidatesChosenCity++
			//} else
		}
		if centerCode >= 0 {
			stats.NoOfCandidatesAlloted++
		}

		// Update the stats in the map
		groupedApplications[circleName][reportingOfficeName] = stats
	}

	// Convert the grouped data to the desired struct
	for _, reportingOffices := range groupedApplications {
		for _, stats := range reportingOffices {
			result = append(result, stats)
		}
	}

	return result, nil
}
*/

// revathi
/*
func GetExamApplicatonsPreferenenceCityWiseStatsrevathi(ctx context.Context, client *ent.Client, cityPreference string) ([]ExamStats, error) {
	var result []ExamStats

	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.CentrePreferenceEQ(cityPreference),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.HallTicketNumberNEQ(""),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	groupedApplications := make(map[string]ExamStats)

	for _, examApplication := range applications {
		// Check if the application's preference city matches the filter
		if examApplication.CentrePreference == cityPreference {
			circleName := examApplication.CircleName
			reportingOfficeName := examApplication.ReportingOfficeName
			centerCode := examApplication.CenterCode

			// Check if ExamStats for the reporting office already exists
			stats, ok := groupedApplications[reportingOfficeName]
			if !ok {
				stats = ExamStats{
					CircleName:               circleName,
					ReportingOfficeName:      reportingOfficeName,
					NoOfCandidatesChosenCity: 0,
					NoOfCandidatesAlloted:    0,
				}
			}

			// Increment counts based on center code
			if centerCode == 0 {
				stats.NoOfCandidatesChosenCity++
			}

			// Check if center code is greater than 0
			if centerCode > 0 {
				stats.NoOfCandidatesAlloted++
			}

			// Update the stats in the map
			groupedApplications[reportingOfficeName] = stats
		}
	}

	// Convert the grouped data to the desired struct
	for _, stats := range groupedApplications {
		result = append(result, stats)
	}

	return result, nil
}
*/

// GDS to MTS/PM/MG End

// IP Exam Start
/* func GetIPExamStatisticsDOOfficeWiseLatests(ctx context.Context, client *ent.Client, examCode int32, facilityID string, examYear string) ([]map[string]interface{}, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return nil, fmt.Errorf("No such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		log.Println("Facility ID cannot be null")
		return nil, fmt.Errorf("Facility ID cannot be null")
	}

	// Query to get the applications from Exam_Applications_IP table matching the provided facilityID
	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_ip.FieldEmployeeID), ent.Desc(exam_applications_ip.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
		return nil, fmt.Errorf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
	}

	// Create a map to store the latest application details for each employee
	latestApplications := make(map[int64]*ent.Exam_Applications_IP)

	// Loop through the applications to find the latest application for each employee and their reporting office
	for _, app := range applications {
		employeeID := app.EmployeeID

		// Check if this application is the latest for the employee
		if latestApp, ok := latestApplications[employeeID]; !ok || app.ID > latestApp.ID {
			latestApplications[employeeID] = app
		}
	}

	// Create a map to store reporting office-wise summaries
	doOfficeSummaries := make(map[string]*DOOfficeWiseSummary)

	// Loop through the latest applications to update counts
	for _, app := range latestApplications {
		reportingOfficeID := app.ReportingOfficeFacilityID

		// Create and initialize doOfficeSummary outside the condition
		doOfficeSummary := doOfficeSummaries[reportingOfficeID]
		if doOfficeSummary == nil {
			doOfficeSummaries[reportingOfficeID] = &DOOfficeWiseSummary{
				ControllingOfficeFacilityID: reportingOfficeID,
				ControllingOfficeName:       app.ControllingOfficeName,
				Permitted:                   0,
				NotPermitted:                0,
				PendingWithCA:               0,
				PendingWithCandidate:        0,
				Received:                    0,
				HallTicketGenerated:         0,
				HallTicketNotGenerated:      0,
				UniqueEmployees:             make(map[int64]struct{}),
			}
		}

		if _, ok := doOfficeSummary.UniqueEmployees[app.EmployeeID]; !ok {
			doOfficeSummary.UniqueEmployees[app.EmployeeID] = struct{}{}
			doOfficeSummary.Received++

			// Update counts based on GenerateHallTicketFlag
			if app.GenerateHallTicketFlag == nil {
				if app.ApplicationStatus == "PendingWithCandidate" {
					// For pending, check if GenerateHallTicketFlag is nil
					doOfficeSummary.PendingWithCandidate++
				} else {
					doOfficeSummary.Pending++
				}
			} else if *app.GenerateHallTicketFlag {
				doOfficeSummary.Permitted++
			} else {
				doOfficeSummary.NotPermitted++
			}
		}
	}

	// Create an empty slice to store the final result
	result := []map[string]interface{}{}
	serialNumber := 0

	// Add reportingOfficeID wise counts and names to the result
	for _, summary := range doOfficeSummaries {
		serialNumber++

		// Display only the latest reporting office counts
		result = append(result, map[string]interface{}{
			"S.No.":                        serialNumber,
			"ControllingOfficeFacilityID":  summary.ControllingOfficeFacilityID,
			"ControllingOfficeName":        summary.ControllingOfficeName,
			"No: Of Applications Received": summary.Received,
			"No. Permitted":                summary.Permitted,
			"No. Not Permitted":            summary.NotPermitted,
			"No. Pending":                  summary.Pending,
			"No. Pending With Candidate":   summary.PendingWithCandidate,
		})
	}

	return result, nil
}

func GetIPExamStatisticsDOOfficeWise(ctx context.Context, client *ent.Client, examCode int32, facilityID string, examYear string) ([]map[string]interface{}, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return nil, fmt.Errorf("No such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		log.Println("Facility ID cannot be null")
		return nil, fmt.Errorf("Facility ID cannot be null")
	}

	// Query to get the applications from Exam_Applications_IP table matching the provided facilityID
	applications, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_ip.ExamCodeEQ(examCode),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active")).
		Order(ent.Desc(exam_applications_ip.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
		return nil, fmt.Errorf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
	}

	// Create a map to store reporting office-wise summaries
	doOfficeSummaries := make(map[string]*DOOfficeWiseSummary)

	// Loop through the applications to group by reporting office-wise and update counts
	for _, app := range applications {
		reportingOfficeID := app.ControllingOfficeFacilityID

		if doOfficeSummaries[reportingOfficeID] == nil {
			doOfficeSummaries[reportingOfficeID] = &DOOfficeWiseSummary{
				ControllingOfficeFacilityID: reportingOfficeID,
				ControllingOfficeName:       app.ControllingOfficeName,
				Permitted:                   0,
				NotPermitted:                0,
				Pending:                     0,
				PendingWithCandidate:        0,
				Received:                    0,
				UniqueEmployees:             make(map[int64]struct{}),
			}
		}

		// Check if the employee is unique for this reporting office
		doOfficeSummary := doOfficeSummaries[reportingOfficeID]
		if _, ok := doOfficeSummary.UniqueEmployees[app.EmployeeID]; !ok {
			doOfficeSummary.UniqueEmployees[app.EmployeeID] = struct{}{}
			doOfficeSummary.Received++

			// Update counts based on GenerateHallTicketFlag
			if app.GenerateHallTicketFlag == nil {
				if app.ApplicationStatus == "PendingWithCandidate" {
					// For pending, check if GenerateHallTicketFlag is nil
					doOfficeSummary.PendingWithCandidate++
				} else {
					doOfficeSummary.Pending++
				}
			} else if *app.GenerateHallTicketFlag {
				doOfficeSummary.Permitted++
			} else {
				doOfficeSummary.NotPermitted++
			}
		}
	}

	// Create an empty slice to store the final result
	result := []map[string]interface{}{}
	serialNumber := 0

	// Add reportingOfficeID wise counts and names to the result
	for _, summary := range doOfficeSummaries {
		serialNumber++

		result = append(result, map[string]interface{}{
			"S.No.":                        serialNumber,
			"ControllingOfficeFacilityID":  summary.ControllingOfficeFacilityID,
			"ControllingOfficeName":        summary.ControllingOfficeName,
			"No: Of Applications Received": summary.Received,
			"No. Permitted":                summary.Permitted,
			"No. Not Permitted":            summary.NotPermitted,
			"No. Pending":                  summary.Pending,
			"No. Pending With Candidate":   summary.PendingWithCandidate,
		})
	}

	return result, nil
}

// Get Recommendations with Emp ID ..
func QueryIPRecommendationsByEmpId(ctx context.Context, client *ent.Client, employeeID int64) ([]*ent.RecommendationsIPApplications, error) {
	//Array of exams

	records, err := client.RecommendationsIPApplications.
		Query().
		Where(recommendationsipapplications.EmployeeIDEQ(employeeID)).
		All(ctx)
	if err != nil {
		log.Println("error querying IP recommendations: ", err)
		return nil, fmt.Errorf("failed to query IP recommendations: %w", err)
	}

	return records, nil
} */

// func GenerateHallticketNumberrIP(ctx context.Context, client *ent.Client, year string, examCode int32, nodalOfficerFacilityID string) (string, error) {
// 	//currentTime := time.Now().Truncate(time.Second)

// 	// Retrieve the last hall ticket number and extract its last four digits
// 	lastFourDigitsMap := make(map[int]bool)
// 	lastHallTicketNumber, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamYearEQ(year),
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),

// 			exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
// 			exam_applications_ip.GenerateHallTicketFlagEQ(true),
// 			exam_applications_ip.HallTicketNumberNEQ(""),
// 			//	exam_applications_ip.GenerateHallTicketFlagByNOEQ(true),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		Order(ent.Desc(exam_applications_ip.FieldHallTicketNumber)).
// 		First(ctx)

// 	fmt.Println(lastHallTicketNumber, "last ")

// 	if err != nil && !ent.IsNotFound(err) {
// 		return "", err
// 	}

// 	if lastHallTicketNumber.HallTicketNumber == "" {
// 		lastHallTicketNumber.HallTicketNumber = "100000000"
// 	}

// 	if lastHallTicketNumber.HallTicketNumber != "" {
// 		lastFourDigitsStr := lastHallTicketNumber.HallTicketNumber[len(lastHallTicketNumber.HallTicketNumber)-4:]
// 		lastFourDigits, err := strconv.Atoi(lastFourDigitsStr)
// 		if err != nil {
// 			return "", err
// 		}
// 		lastFourDigitsMap[lastFourDigits] = true
// 	}

// 	// Retrieve all eligible applications
// 	applications, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamYearEQ(year),
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
// 			exam_applications_ip.GenerateHallTicketFlagEQ(true),
// 			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_ip.StatusEQ("active"),
// 			exam_applications_ip.Or(
// 				//exam_applications_ip.HallTicketNumberIsNull(),
// 				exam_applications_ip.HallTicketNumberEQ(""),
// 			),
// 			//exam_applications_ip.GenerateHallTicketFlagByNOEQ(true),
// 		).
// 		Order(ent.Asc(exam_applications_ip.FieldTempHallTicket)).
// 		All(ctx)

// 	if err != nil {
// 		return "", err
// 	}

// 	// If no data, set the start number to 1, else set it to the maximum number found + 1
// 	startNumber := 1
// 	if len(lastFourDigitsMap) > 0 {
// 		for lastFourDigits := range lastFourDigitsMap {
// 			startNumber = lastFourDigits + 1
// 		}
// 	}

// 	// Generate hall tickets
// 	var successCount int
// 	for _, application := range applications {
// 		hallTicketNumber := fmt.Sprintf("%s%04d", application.TempHallTicket, startNumber)
// 		_, err := application.Update().
// 			SetHallTicketNumber(hallTicketNumber).
// 			SetGenerateHallTicketFlagByNO(true).
// 			Save(ctx)
// 		if err != nil {
// 			return "", err
// 		}
// 		startNumber++
// 		successCount++
// 	}

// 	// Return success message
// 	return fmt.Sprintf("Generated hall tickets successfully for %d eligible candidates", successCount), nil
// }

// func GenerateHallticketNumberrIP(ctx context.Context, client *ent.Client, year string, examCode int32, nodalOfficerFacilityID string) (string, error) {
// 	// Retrieve the last hall ticket number and extract its last four digits
// 	lastFourDigitsMap := make(map[int]bool)
// 	lastHallTicketNumber, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamYearEQ(year),
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
// 			exam_applications_ip.GenerateHallTicketFlagEQ(true),
// 			exam_applications_ip.HallTicketNumberNEQ(""),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		Order(ent.Desc(exam_applications_ip.FieldHallTicketNumber)).
// 		First(ctx)

// 	// Check for error, handle not found errors explicitly
// 	if err != nil {
// 		if ent.IsNotFound(err) {
// 			lastHallTicketNumber = nil
// 		} else {
// 			return "", err
// 		}
// 	}

// 	// Initialize lastHallTicketNumber if not found
// 	if lastHallTicketNumber.HallTicketNumber == "" {
// 		lastHallTicketNumber.HallTicketNumber = "100000000"
// 	}

// 	// Extract last four digits
// 	lastFourDigitsStr := lastHallTicketNumber.HallTicketNumber[len(lastHallTicketNumber.HallTicketNumber)-4:]
// 	lastFourDigits, err := strconv.Atoi(lastFourDigitsStr)
// 	if err != nil {
// 		return "", err
// 	}
// 	lastFourDigitsMap[lastFourDigits] = true

// 	// Retrieve all eligible applications
// 	applications, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamYearEQ(year),
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
// 			exam_applications_ip.GenerateHallTicketFlagEQ(true),
// 			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_ip.StatusEQ("active"),
// 			exam_applications_ip.Or(
// 				exam_applications_ip.HallTicketNumberNEQ(""),
// 			),
// 		).
// 		Order(ent.Asc(exam_applications_ip.FieldTempHallTicket)).
// 		All(ctx)

// 	if err != nil {
// 		return "", err
// 	}

// 	// If no data, set the start number to 1, else set it to the maximum number found + 1
// 	startNumber := 1
// 	if len(lastFourDigitsMap) > 0 {
// 		for lastFourDigits := range lastFourDigitsMap {
// 			startNumber = lastFourDigits + 1
// 		}
// 	}

// 	// Generate hall tickets
// 	var successCount int
// 	for _, application := range applications {
// 		hallTicketNumber := fmt.Sprintf("%s%04d", application.TempHallTicket, startNumber)
// 		_, err := application.Update().
// 			SetHallTicketNumber(hallTicketNumber).
// 			SetGenerateHallTicketFlagByNO(true).
// 			Save(ctx)
// 		if err != nil {
// 			return "", err
// 		}
// 		startNumber++
// 		successCount++
// 	}

// 	// Return success message
// 	return fmt.Sprintf("Generated hall tickets successfully for %d eligible candidates", successCount), nil
// }

// revathi
/*
func GetExamApplicatonsPreferenenceCityWiseStatsrevathi(ctx context.Context, client *ent.Client, cityPreference string) ([]ExamStats, error) {
	var result []ExamStats

	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.CentrePreferenceEQ(cityPreference),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.HallTicketNumberNEQ(""),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	groupedApplications := make(map[string]ExamStats)

	for _, examApplication := range applications {
		// Check if the application's preference city matches the filter
		if examApplication.CentrePreference == cityPreference {
			circleName := examApplication.CircleName
			reportingOfficeName := examApplication.ReportingOfficeName
			centerCode := examApplication.CenterCode

			// Check if ExamStats for the reporting office already exists
			stats, ok := groupedApplications[reportingOfficeName]
			if !ok {
				stats = ExamStats{
					CircleName:               circleName,
					ReportingOfficeName:      reportingOfficeName,
					NoOfCandidatesChosenCity: 0,
					NoOfCandidatesAlloted:    0,
				}
			}

			// Increment counts based on center code
			if centerCode == 0 {
				stats.NoOfCandidatesChosenCity++
			}

			// Check if center code is greater than 0
			if centerCode > 0 {
				stats.NoOfCandidatesAlloted++
			}

			// Update the stats in the map
			groupedApplications[reportingOfficeName] = stats
		}
	}

	// Convert the grouped data to the desired struct
	for _, stats := range groupedApplications {
		result = append(result, stats)
	}

	return result, nil
}
*/

/* func GetExamApplicatonsPreferenenceCityWiseStatssss(ctx context.Context, client *ent.Client, cityPreference string) ([]ExamStats, error) {
	var result []ExamStats

	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.CentrePreferenceEQ(cityPreference),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.HallTicketNumberNEQ(""),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	groupedApplications := make(map[string]map[string]ExamStats)

	for _, examApplication := range applications {
		circleName := examApplication.CircleName
		reportingOfficeName := examApplication.ReportingOfficeName
		centerCode := examApplication.CenterCode

		if _, ok := groupedApplications[circleName]; !ok {
			groupedApplications[circleName] = make(map[string]ExamStats)
		}

		// Check if ExamStats for the reporting office already exists
		stats, ok := groupedApplications[circleName][reportingOfficeName]
		if !ok {
			stats = ExamStats{
				CircleName:               circleName,
				ReportingOfficeName:      reportingOfficeName,
				NoOfCandidatesChosenCity: 0,
				NoOfCandidatesAlloted:    0,
			}
		}

		//Increment counts based on center code
		if examApplication.CentrePreference == cityPreference {

			//if centerCode == 0 {
			stats.NoOfCandidatesChosenCity++
			//} else
		}
		if centerCode >= 0 {
			stats.NoOfCandidatesAlloted++
		}

		// Update the stats in the map
		groupedApplications[circleName][reportingOfficeName] = stats
	}

	// Convert the grouped data to the desired struct
	for _, reportingOffices := range groupedApplications {
		for _, stats := range reportingOffices {
			result = append(result, stats)
		}
	}

	return result, nil
}
*/
// func GetExamApplicatonsPreferenenceCityWiseStats(ctx context.Context, client *ent.Client, cityPreference string) (map[string]map[string]int, error) {
// 	result := make(map[string]map[string]int)

// 	// Fetch all applications for the given cityPreference
// 	applications, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.CentrePreferenceEQ(cityPreference),
// 			//exam_applications_ip.HallTicketNumberNEQ(0),
// 			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_ip.CenterCodeIsNil(),
// 		).
// 		All(ctx)

// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, examApplication := range applications {
// 		circleName := examApplication.CircleName
// 		reportingOfficeName := examApplication.ReportingOfficeName

// 		// Create the map if not exists
// 		if _, ok := result[circleName]; !ok {
// 			result[circleName] = make(map[string]int)
// 		}

// 		// Increment the count for the corresponding reporting office
// 		result[circleName][reportingOfficeName]++
// 	}

// 	return result, nil
// }

// Count of details based on Reporting Offices

// Generate Hall Ticket Flag .../*
/*func ApproveHallTicketGenerationByNO(client *ent.Client, applicationRecord *ent.Exam_Applications_IP) (string, error) {
	ctx := context.Background()

	// Check if applicationRecord is nil
	if applicationRecord == nil {
		return "", fmt.Errorf("please provide exam code, Facility ID, and Approval Flag for approving. They are mandatory")
	}

	// Check if ExamCode is nil or invalid
	if applicationRecord.ExamCode <= 0 {
		log.Println("No such valid exam code exists")
		return "", fmt.Errorf("No such valid exam code exists")
	}

	if applicationRecord.NodalOfficeID == "" {
		return "", fmt.Errorf("Please provide Nodal Officer's office ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_IP table
	count, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.NodalOfficeIDEQ(applicationRecord.NodalOfficeID),
			exam_applications_ip.GenerateHallTicketFlag(true),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.CenterCodeNEQ(0),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", applicationRecord.NodalOfficeID)
		return "", fmt.Errorf("No valid applications available for the circle")
	}
	// Perform the update to set GenerateHallTicketFlagByNO for eligible candidates
	_, err = client.Exam_Applications_IP.
		Update().
		Where(
			exam_applications_ip.NodalOfficeIDEQ(applicationRecord.NodalOfficeID),
			exam_applications_ip.GenerateHallTicketFlag(true),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.CenterCodeNEQ(0),
		).
		SetGenerateHallTicketFlagByNO(true).
		Save(ctx)
	if err != nil {
		log.Printf("Failed to update applications: %v", err)
		return "", fmt.Errorf("Failed to update applications: %v", err)
	}

	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", applicationRecord.CircleName), nil
}*/

//Here it sets in Exam appl table..
/*func ApproveHallTicketGenerationByNOForIPExam(client *ent.Client, examCode int32, facilityID string, approveHallTicket bool) (string, error) {
	ctx := context.Background()

	// Check if ExamCode is nil or invalid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return "", fmt.Errorf("No such valid exam code exists")
	}

	if facilityID == "" {
		return "", fmt.Errorf("Please provide Facility ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_PS table
	count, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.NodalOfficeIDEQ(facilityID),
			exam_applications_ip.GenerateHallTicketFlag(true),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.CenterCodeNEQ(0),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", facilityID)
		return "", fmt.Errorf("No valid applications available for the circle")
	}
	// Perform the update to set GenerateHallTicketFlagByNO for eligible candidates
	_, err = client.Exam_Applications_IP.
		Update().
		Where(
			exam_applications_ip.NodalOfficeIDEQ(facilityID),
			exam_applications_ip.GenerateHallTicketFlag(true),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.CenterCodeNEQ(0),
		).
		//SetGenerateHallTicketFlagByNO(approveHallTicket).
		Save(ctx)
	if err != nil {
		log.Printf("Failed to update applications: %v", err)
		return "", fmt.Errorf("Failed to update applications: %v", err)
	}

	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", facilityID), nil
}*/

// Here it sets in circlemaster table
// func ApproveHallTicketGenerationByNOForIPExam(client *ent.Client, examCode int32, facilityID string, approveHallTicket bool) (string, error) {
// 	ctx := context.Background()

// 	// Check if ExamCode is nil or invalid
// 	if examCode <= 0 {
// 		log.Println("No such valid exam code exists")
// 		return "", fmt.Errorf("No such valid exam code exists")
// 	}

// 	if facilityID == "" {
// 		return "", fmt.Errorf("Please provide Facility ID")
// 	}

// 	// Check if circleOfficeID exists in Exam_Applications_PS table
// 	count, err := client.Exam_Applications_IP.
// 		Query()./*  */
// 		Where(
// 			exam_applications_ip.NodalOfficeIDEQ(facilityID),
// 			exam_applications_ip.GenerateHallTicketFlag(true),
// 			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_ip.CenterCodeNEQ(0),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
// 		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
// 	}

// 	if count == 0 {
// 		log.Printf("No valid applications available for the circle: %s", facilityID)
// 		return "", fmt.Errorf("No valid applications available for the circle")
// 	}

// 	// Get the CircleMaster entity based on the Facility ID
// 	circleMaster, err := client.CircleSummaryForNO.
// 		Query().
// 		Where(circlesummaryforno.CircleOfficeIdEQ(facilityID)).
// 		Only(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve CircleMaster: %v", err)
// 		return "", fmt.Errorf("Failed to retrieve CircleMaster: %v", err)
// 	}

// 	// Update the GenerateHallTicketFlagByNO field in the CircleMaster entity
// 	circleMaster = circleMaster.
// 		Update().
// 		SetApproveHallTicketGenrationIP(approveHallTicket).
// 		SaveX(ctx)

// 	// Update the GenerateHallTicketFlagByNO in Exam_Applications_IP table
// 	_, err = client.Exam_Applications_IP.
// 		Update().
// 		Where(
// 			exam_applications_ip.NodalOfficeIDEQ(facilityID),
// 			exam_applications_ip.GenerateHallTicketFlag(true),
// 			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_ip.CenterCodeNEQ(0),
// 		).
// 		//SetGenerateHallTicketFlagByNO(approveHallTicket).
// 		Save(ctx)
// 	if err != nil {
// 		log.Printf("Failed to update applications: %v", err)
// 		return "", fmt.Errorf("Failed to update applications: %v", err)
// 	}
// 	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", facilityID), nil
// }

// Get Exams by Exam Code.
/* func QueryExamsIPByExamNameCode(ctx context.Context, client *ent.Client, examNameCode string) (*ent.Exam_IP, int32, error) {
	// Check if examNameCode is empty
	if examNameCode == "" {
		return nil, 422, fmt.Errorf("Please provide exam name code")
	}

	u, err := client.Exam_IP.Query().
		Where(exam_ip.ExamNameCode(examNameCode)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			fmt.Println("No active employee found with the given ID.")
			return nil, 422, fmt.Errorf("not found Exam_IP: %w", err)
		} else {
			log.Fatalf("failed querying Exam_IP: %v", err)
			return nil, 400, fmt.Errorf("failed querying Exam_IP: %w", err)
		}
	}

	return u, 200, nil
}
*/

// Query IPExam Application with Emp ID.
/* func QueryIPExamApplicationsByEmpID(ctx context.Context, client *ent.Client, empid int64, id1 string) (*ent.Exam_Applications_IP, error) {

	newAppln, err := client.Exam_Applications_IP.
		Query().
		Where((exam_applications_ip.EmployeeIDEQ(empid)), exam_applications_ip.ExamYear(id1), exam_applications_ip.StatusEQ("active")).
		//Order(ent.Desc(exam_applications_ip.FieldID)).
		WithCirclePrefRef().
		WithIPApplicationsRef().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("no application exists ")
		} else {
			return nil, err
		}
	}

	// Extract only the desired fields from the CirclePrefRef edge
	var circlePrefs []*ent.PlaceOfPreferenceIP
	for _, edge := range newAppln.Edges.CirclePrefRef {
		circlePrefs = append(circlePrefs, &ent.PlaceOfPreferenceIP{
			PlacePrefNo:    edge.PlacePrefNo,
			PlacePrefValue: edge.PlacePrefValue,
		})
	}

	// Update the CirclePrefRef edge with the filtered values
	newAppln.Edges.CirclePrefRef = circlePrefs

	//currentTime := time.Now().Truncate(time.Second)
	var recomondPref []*ent.RecommendationsIPApplications
	for _, edge := range newAppln.Edges.IPApplicationsRef {
		recomondPref = append(recomondPref, &ent.RecommendationsIPApplications{
			//RecommendationId:            edge.RecommendationId,
			ApplicationID:     edge.ApplicationID,
			EmployeeID:        edge.EmployeeID,
			CARecommendations: edge.CARecommendations,
			NORecommendations: edge.NORecommendations,
			ApplicationStatus: edge.ApplicationStatus,
			ExamNameCode:      edge.ExamNameCode,
			CAUserName:        edge.CAUserName,
			CARemarks:         edge.CARemarks,
			CAUpdatedAt:       edge.CAUpdatedAt,
			NOUpdatedAt:       edge.NOUpdatedAt,
			NORemarks:         edge.NORemarks,
			NOUserName:        edge.NOUserName,
			VacancyYear:       edge.VacancyYear,
		})
	}
	newAppln.Edges.IPApplicationsRef = recomondPref
	newAppln.UpdatedAt = newAppln.UpdatedAt.Truncate(24 * time.Hour)

	// log.Println("details returned by IP Exam Applications for the Employee: ", newAppln)
	return newAppln, nil
} */

// IP Exam End

// PS Group B Start

func GetPSDivisionsByCircleOfficeID(ctx context.Context, client *ent.Client, circleOfficeID string, examYear string) ([]*ent.Exam_Applications_PS, error) {
	// Check if the circle office ID exists in the exam_application_ps table.
	exists, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.StatusEQ("active"),
		).
		Exist(ctx)
	if err != nil {
		log.Printf("Failed to query exam_application_ps: %v\n", err)
		return nil, fmt.Errorf("failed to query exam_application_ps: %v", err)
	}
	if !exists {
		log.Printf("Circle office ID does not exist: %s\n", circleOfficeID)
		return nil, fmt.Errorf("circle office ID does not exist")
	}

	// Query the exam_application_ps table for unique records based on the provided conditions.
	applications, err := client.Exam_Applications_PS.
		Query().
		Select(
			exam_applications_ps.FieldReportingOfficeFacilityID,
			exam_applications_ps.FieldReportingOfficeName,
		).
		Where(
			exam_applications_ps.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_ps.Not(exam_applications_ps.GenerateHallTicketFlag(true)),
			exam_applications_ps.ExamCityCenterCodeIsNil(),
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		log.Printf("Failed to query exam_application_ps: %v\n", err)
		return nil, fmt.Errorf("failed to query exam_application_ps: %v", err)
	}

	// Filter and return distinct records based on reporting office ID and name.
	distinctApplications := make(map[string]*ent.Exam_Applications_PS)
	for _, app := range applications {
		key := app.ReportingOfficeFacilityID
		distinctApplications[key] = app
	}

	result := make([]*ent.Exam_Applications_PS, 0, len(distinctApplications))
	for _, app := range distinctApplications {
		result = append(result, app)
	}

	log.Printf("Retrieved %d distinct divisions for Circle Office ID: %s\n", len(result), circleOfficeID)

	// Log the applications as an array of strings
	appStrings := make([]string, len(result))
	for i, app := range result {
		appStrings[i] = fmt.Sprintf("Reporting Office ID: %s, Reporting Office Name: %s", app.ReportingOfficeFacilityID, app.ReportingOfficeName)
	}
	log.Printf("Applications: %+v\n", appStrings)

	return result, nil
}

/* func GetPSApplicationsWithHallTicket(client *ent.Client, examCode int32, employeeID int64) (*ent.Exam_Applications_PS, error) {
	ctx := context.Background()

	// Check if exam code is valid
	if examCode == 0 {
		return nil, errors.New("Please provide a valid exam code")
	}

	if examCode == 2 {
		// Check if the employee_ID exists in the Exam_Applications_PS table
		exists, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.EmployeeIDEQ(employeeID),
			).
			Exist(ctx)
		if err != nil {
			return nil, fmt.Errorf("Failed to check employee ID in PS Applications: %v", err)
		}
		if !exists {
			return nil, fmt.Errorf("No applications are found for the employee in PS Applications: %d", employeeID)
		}

		// Query the Exam_Applications_PS table to retrieve the applicant details
		application, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.EmployeeIDEQ(employeeID),
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.HallTicketNumberNEQ(""),
				exam_applications_ps.ExamCityCenterCodeNEQ(0),
			).
			WithPSExamCentres().
			First(ctx)
		if err != nil {
			return nil, fmt.Errorf("No Admit card details available for the applicant: %v", err)
		}

		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := client.RecommendationsPSApplications.Query().
			Where(recommendationspsapplications.EmployeeIDEQ(employeeID)).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("Failed to retrieve recommendations: %v", err)
		}

		// Assign the fetched recommendations to the application entity
		application.Edges.PSApplicationsRef = recommendations

		return application, nil
	} else if examCode == 1 {
		// Check if the employee_ID exists in the Exam_Applications_PS table
		exists, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.EmployeeIDEQ(employeeID),
			).
			Exist(ctx)
		if err != nil {
			return nil, fmt.Errorf("Failed to check employee ID in PS Group B Applications: %v", err)
		}
		if !exists {
			return nil, fmt.Errorf("No applications are found for the employee in PS Group B Applications: %d", employeeID)
		}

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationps, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.EmployeeIDEQ(employeeID),
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.HallTicketNumberNEQ(""),
				exam_applications_ps.ExamCityCenterCodeNEQ(0),
			).
			WithPSExamCentres().
			First(ctx)
		if err != nil {
			return nil, fmt.Errorf("No Admit card details available for the applicant: %v", err)
		}

		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := client.RecommendationsPSApplications.Query().
			Where(recommendationspsapplications.EmployeeIDEQ(employeeID)).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("Failed to retrieve recommendations: %v", err)
		}

		// Assign the fetched recommendations to the application entity
		applicationps.Edges.PSApplicationsRef = recommendations

		return applicationps, nil
	}

	return nil, errors.New("Invalid exam code")
}
*/
/* func generatePSApplicationNumber(client *ent.Client, employeeID int64) (string, error) {
	nextApplicationNumber, err := getNextPSApplicationNumberFromDatabase(client)
	if err != nil {
		return "", err
	}

	// Get the current year
	currentYear := time.Now().Year()

	// Format the application number as "PSYYYYXXXXXX"
	applicationNumber := fmt.Sprintf("PS%d%06d", currentYear, nextApplicationNumber)

	return applicationNumber, nil
}

func getNextPSApplicationNumberFromDatabase(client *ent.Client) (int64, error) {
	ctx := context.TODO()
	lastApplication, err := client.Exam_Applications_PS.
		Query().
		Order(ent.Desc(exam_applications_ps.FieldID)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// No existing applications, start from 100001
			return 100001, nil
		}
		return 0, fmt.Errorf("failed to retrieve last application: %v", err)
	}

	return lastApplication.ID + 1, nil
}
*/

// Get CA Pending with EmpID
// func QueryPSApplicationsByCAPendingByEmpID(ctx context.Context, client *ent.Client, empID int64) ([]*ent.Exam_Applications_PS, error) {
// 	// Check if employee ID exists
// 	employeeExists, err := client.Exam_Applications_PS.
// 		Query().
// 		Where(
// 			exam_applications_ps.EmployeeIDEQ(empID),
// 			exam_applications_ps.Or(
// 				exam_applications_ps.ApplicationStatusEQ("CAVerificationPending"),
// 				exam_applications_ps.ApplicationStatusEQ("ResubmitCAVerificationPending"),
// 			),
// 		).
// 		WithCirclePrefRefPS().
// 		WithPSApplicationsRef().
// 		Exist(ctx)
// 	if err != nil {
// 		log.Println("error checking employee existence: ", err)
// 		return nil, fmt.Errorf("failed checking employee existence: %w", err)
// 	}
// 	if !employeeExists {
// 		return nil, fmt.Errorf("employee not found with ID: or the verification is not pending with CA %d", empID)
// 	}

// 	// Retrieve the record
// 	record, err := client.Exam_Applications_PS.
// 		Query().
// 		Where(
// 			exam_applications_ps.EmployeeIDEQ(empID),
// 			exam_applications_ps.Or(
// 				exam_applications_ps.ApplicationStatusEQ("CAVerificationPending"),
// 				exam_applications_ps.ApplicationStatusEQ("ResubmitCAVerificationPending"),
// 			),
// 		).
// 		WithPSApplicationsRef().
// 		WithCirclePrefRefPS().
// 		All(ctx)
// 	if err != nil {
// 		log.Println("error at PS Exam Applications fetching: ", err)
// 		return nil, fmt.Errorf("failed querying PS exams Applications: %w", err)
// 	}

// 	//log.Println("CA pending records returned: ", record)
// 	return record, nil
// }

// Get Exams by Exam Code.
/* func QueryExamsPSByExamNameCode(ctx context.Context, client *ent.Client, examNameCode string) (*ent.Exam_PS, error) {
	// Check if examNameCode is empty
	if examNameCode == "" {
		return nil, fmt.Errorf("Please provide exam name code")
	}

	u, err := client.Exam_PS.Query().
		Where(exam_ps.ExamNameCode(examNameCode)).
		Only(ctx)
	if err != nil {
		log.Println("error at getting Exam_PS: ", err)
		return nil, fmt.Errorf("failed querying Exam_PS: %w", err)
	}
	log.Println("Exam_PS details returned: ", u)
	return u, nil
}
*/
// list of reporting offices

// type HallticketStatsPS struct {
// 	CircleID        int32  `json:"CircleID"`
// 	StartingNumber  int    `json:"StartingNumber"`
// 	EndingNumber    int    `json:"EndingNumber"`
// 	Count           int    `json:"Count"`
// 	StartHallTicket string `json:"StartHallTicket"`
// 	EndHallTicket   string `json:"EndHallTicket"`
// }

// Generate Hall Ticket Numbers return array with stng & eng nos.
/*func GenerateHallticketNumber(ctx context.Context, client *ent.Client) ([]HallticketStatsPS, error) {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Printf("Error loading location: %v", err)
		return nil, err
	}

	currentTime := time.Now().In(loc).Truncate(time.Second)

	applications, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.GenerateHallTicketFlag(true),
			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ps.CenterCodeNEQ(0),
			exam_applications_ps.ExamCodeNEQ(0),
			exam_applications_ps.ExamYearNEQ(""),
			exam_applications_ps.CategoryCodeNEQ(0),
			exam_applications_ps.CircleIDNEQ(0),
			exam_applications_ps.RegionIDNEQ(0),
			exam_applications_ps.DivisionIDNEQ(0),
			exam_applications_ps.EmployeeIDNEQ(0),
		).
		Order(ent.Desc(exam_applications_ps.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	circleStats := make(map[string]HallticketStats)
	for _, application := range applications {
		circleID := application.CircleID
		regionID := application.RegionID
		divisionID := application.DivisionID
		key := fmt.Sprintf("%d-%d-%d", circleID, regionID, divisionID)

		stats, exists := circleStats[key]
		if !exists {
			// Reset the serial number for each unique combination of circle, region, and division
			stats.StartingNumber = 1
		}

		identificationNo := stats.StartingNumber + stats.Count
		examYear := application.ExamYear
		if len(examYear) >= 2 {
			examYear = examYear[len(examYear)-2:]
		}

		hallticketNumber := generateHallticketNumber(
			application.ExamCode,
			examYear,
			application.CategoryCode,
			circleID,
			regionID,
			divisionID,
			identificationNo)

		// Validate if the hallticket number is of 13 digits
		if len(hallticketNumber) != 13 {
			log.Printf("Hallticket Number: %s\n", hallticketNumber)
			continue
		}

		log.Printf("Generated hallticket number is of 13 digits: %s\n", hallticketNumber)
		log.Printf("Employee ID: %d\n", application.EmployeeID)

		_, err = application.Update().
			SetHallTicketNumber(hallticketNumber).
			SetHallTicketGeneratedFlag(true).
			SetHallTicketGeneratedDate(currentTime).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		 // Update the CircleID field in HallticketStats
		 stats.CircleID = circleID

		stats.Count++
		stats.EndingNumber = stats.StartingNumber + stats.Count - 1
		stats.StartHallTicket = generateHallticketNumber(
			application.ExamCode,
			examYear,
			application.CategoryCode,
			circleID,
			regionID,
			divisionID,
			stats.StartingNumber)
		stats.EndHallTicket = generateHallticketNumber(
			application.ExamCode,
			examYear,
			application.CategoryCode,
			circleID,
			regionID,
			divisionID,
			stats.EndingNumber)

		circleStats[key] = stats
	}

	statsSlice := make([]HallticketStats, 0, len(circleStats))
	for _, stats := range circleStats {
		statsSlice = append(statsSlice, stats)
	}

	return statsSlice, nil
}*/

// type HallticketResultPS struct {
// 	CircleID string `json:"circleID"`
// 	Count    int    `json:"count"`
// }
// type CircleStatsPS struct {
// 	CircleID        string `json:"CircleID"`
// 	HallTicketCount int    `json:"Count"`
// }

// // Generate Hall Ticket Numbers and return JSON array of CircleID and count.

// func GenerateHallticketNumberPS(ctx context.Context, client *ent.Client) ([]CircleStatsPS, error) {
// 	// loc, err := time.LoadLocation("Asia/Kolkata")
// 	// if err != nil {
// 	// 	log.Printf("Error loading location: %v", err)
// 	// 	return nil, err
// 	// }

// 	// currentTime := time.Now().In(loc).Truncate(time.Second)
// 	currentTime := time.Now().Truncate(time.Second)

// 	applications, err := client.Exam_Applications_PS.
// 		Query().
// 		Where(
// 			exam_applications_ps.GenerateHallTicketFlag(true),
// 			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_ps.ExamCityCenterCodeNEQ(0),
// 			//exam_applications_ps.CenterCodeEQ(ExamCenterCode),
// 			exam_applications_ps.ExamCodeNEQ(0),
// 			exam_applications_ps.ExamYearNEQ(""),
// 			exam_applications_ps.CategoryCodeNEQ(""),
// 			exam_applications_ps.WorkingOfficeCircleFacilityIDNEQ(""),
// 			exam_applications_ps.WorkingOfficeRegionFacilityIDNEQ(""),
// 			exam_applications_ps.WorkingOfficeDivisionFacilityIDNEQ(""),
// 			exam_applications_ps.EmployeeIDNEQ(0),
// 		).
// 		Order(ent.Desc(exam_applications_ps.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	circleStatsPS := make(map[string]int)
// 	for _, application := range applications {
// 		key := fmt.Sprintf("%d", application.WorkingOfficeCircleFacilityID)
// 		circleStatsPS[key]++

// 		identificationNo := circleStatsPS[key]
// 		examYear := application.ExamYear
// 		if len(examYear) >= 2 {
// 			examYear = examYear[len(examYear)-2:]
// 		}

// 		hallticketNumber := generateHallticketNumberPS(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.WorkingOfficeCircleFacilityID,
// 			//application.RegionID,
// 			application.WorkingOfficeDivisionFacilityID,
// 			identificationNo)

// 		// Validate if the hallticket number is of 12 digits
// 		if len(hallticketNumber) != 12 {
// 			log.Printf("Hallticket Number: %s\n", hallticketNumber)
// 			continue
// 		}

// 		log.Printf("Generated hallticket number is of 12 digits: %s\n", hallticketNumber)
// 		log.Printf("Employee ID: %d\n", application.EmployeeID)

// 		_, err = application.Update().
// 			SetHallTicketNumber(hallticketNumber).
// 			SetHallTicketGeneratedFlag(true).
// 			SetHallTicketGeneratedDate(currentTime).
// 			Save(ctx)
// 		if err != nil {
// 			errMsg := fmt.Sprintf("Unable to generate Hallticket number: %s", err.Error())
// 			log.Println(errMsg)
// 			return nil, errors.New(errMsg)
// 		}
// 	}

// 	// Convert the circleStats map to the desired JSON output
// 	statsSlice := make([]CircleStatsPS, 0, len(circleStatsPS))
// 	for key, count := range circleStatsPS {
// 		statsSlice = append(statsSlice, CircleStatsPS{CircleID: key, HallTicketCount: count})
// 	}

// 	return statsSlice, nil
// }

// generatew ht's and return as string
// func GenerateHallticketPSReturnStringMessage(ctx context.Context, client *ent.Client) (string, error) {
// func GenerateHallticketPSReturnStringMessage(ctx context.Context, client *ent.Client) ([]HallticketStatsPS, error) {
// 	// loc, err := time.LoadLocation("Asia/Kolkata")
// 	// if err != nil {
// 	// 	log.Printf("Error loading location: %v", err)
// 	// 	return nil, err
// 	// }

// 	// currentTime := time.Now().In(loc).Truncate(time.Second)
// 	currentTime := time.Now().Truncate(time.Second)

// 	applications, err := client.Exam_Applications_PS.
// 		Query().
// 		Where(
// 			exam_applications_ps.GenerateHallTicketFlag(true),
// 			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_ps.ExamCityCenterCodeNEQ(0),
// 			exam_applications_ps.ExamCodeNEQ(0),
// 			exam_applications_ps.ExamYearNEQ(""),
// 			exam_applications_ps.CategoryCodeNEQ(""),
// 			exam_applications_ps.WorkingOfficeCircleFacilityIDNEQ(""),
// 			exam_applications_ps.WorkingOfficeRegionFacilityIDNEQ(""),
// 			exam_applications_ps.WorkingOfficeDivisionFacilityIDNEQ(""),
// 			exam_applications_ps.EmployeeIDNEQ(0),
// 		).
// 		Order(ent.Desc(exam_applications_ps.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	circleStatsPS := make(map[string]HallticketStatsPS)
// 	for _, application := range applications {
// 		circleID := (application.WorkingOfficeCircleFacilityID)
// 		regionID := (application.WorkingOfficeRegionFacilityID)
// 		divisionID := (application.WorkingOfficeDivisionFacilityID)
// 		key := circleID + "-" + regionID + "-" + divisionID

// 		stats, exists := circleStatsPS[key]
// 		if !exists {
// 			// Reset the serial number for each unique combination of CircleID, RegionID, and DivisionID
// 			stats.StartingNumber = 1
// 		}

// 		identificationNo := stats.StartingNumber + stats.Count
// 		examYear := application.ExamYear
// 		if len(examYear) >= 2 {
// 			examYear = examYear[len(examYear)-2:]
// 		}

// 		hallticketNumber := generateHallticketNumberPS(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.WorkingOfficeCircleFacilityID,
// 			//application.RegionID,
// 			application.WorkingOfficeDivisionFacilityID,
// 			identificationNo)

// 		// Validate if the hallticket number is of 12 digits
// 		if len(hallticketNumber) != 12 {
// 			log.Printf("Generated hallticket number is not of 12 digits. Skipping application with ID: %d\n", application.EmployeeID)
// 			continue
// 		}

// 		log.Printf("Generated hallticket number is of 12 digits: %s\n", hallticketNumber)
// 		//log.Printf("Application Details: ExamCode: %d, ExamYear: %s, CategoryCode: %d, CircleID: %d, RegionID: %d, DivisionID: %d\n",
// 		//	application.ExamCode, application.ExamYear, application.CategoryCode,
// 		//	application.CircleID, application.RegionID, application.DivisionID)
// 		//log.Printf("Employee ID: %d\n", application.EmployeeID)

// 		_, err = application.Update().
// 			SetHallTicketNumber(hallticketNumber).
// 			SetHallTicketGeneratedFlag(true).
// 			SetHallTicketGeneratedDate(currentTime).
// 			Save(ctx)
// 		if err != nil {
// 			return nil, err
// 		}

// 		stats.Count++
// 		stats.EndingNumber = stats.StartingNumber + stats.Count - 1
// 		stats.StartHallTicket = generateHallticketNumberPS(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.WorkingOfficeCircleFacilityID,
// 			//application.RegionID,
// 			application.WorkingOfficeDivisionFacilityID,
// 			stats.StartingNumber)
// 		stats.EndHallTicket = generateHallticketNumberPS(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.WorkingOfficeCircleFacilityID,
// 			//application.RegionID,
// 			application.WorkingOfficeDivisionFacilityID,
// 			stats.EndingNumber)

// 		circleStatsPS[key] = stats
// 	}

// 	statsSlice := make([]HallticketStatsPS, 0, len(circleStatsPS))
// 	for _, stats := range circleStatsPS {
// 		statsSlice = append(statsSlice, stats)
// 	}

// 	return statsSlice, nil
// }

// Return the success message with the count of eligible candidates
//successMessage := fmt.Sprintf("Hall Ticket generated successfully for %d eligible candidates", len(applications))
//return successMessage, nil
//}

//	func generateHallticketNumberPS(examCode int32, examYear string, categoryCode string, circleID string /*regionID int32,*/, divisionID string, identificationNo int) string {
//		// Generate the Hallticket Number based on the provided formatfmt.Sprintf("%d%s%d%d%d%d%04d", examCode, examYear, getFormattedCode(circleID), regionID, getFormattedCode(divisionID), categoryCode, identificationNo)
//		hallticketNumber := fmt.Sprintf("%d%s%s%s%d%04d", examCode, examYear, getFormattedCodePS(circleID) /*regionID,*/, getFormattedCodePS(divisionID), categoryCode, identificationNo)
//		return hallticketNumber
//	}

// func getFormattedCodePS(code string) string {
// 	// Format the code as a string with the required number of digits
// 	lastTwoDigits := code % 100
// 	return fmt.Sprintf("%02d", lastTwoDigits)
// }

// Count of details based on Reporting Offices

// Generate Hall Ticket Flag .../*
/*func ApproveHallTicketGenerationByNO(client *ent.Client, applicationRecord *ent.Exam_Applications_IP) (string, error) {
	ctx := context.Background()

	// Check if applicationRecord is nil
	if applicationRecord == nil {
		return "", fmt.Errorf("please provide exam code, Facility ID, and Approval Flag for approving. They are mandatory")
	}

	// Check if ExamCode is nil or invalid
	if applicationRecord.ExamCode <= 0 {
		log.Println("No such valid exam code exists")
		return "", fmt.Errorf("No such valid exam code exists")
	}

	if applicationRecord.NodalOfficeID == "" {
		return "", fmt.Errorf("Please provide Nodal Officer's office ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_IP table
	count, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.NodalOfficeIDEQ(applicationRecord.NodalOfficeID),
			exam_applications_ip.GenerateHallTicketFlag(true),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.CenterCodeNEQ(0),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", applicationRecord.NodalOfficeID)
		return "", fmt.Errorf("No valid applications available for the circle")
	}
	// Perform the update to set GenerateHallTicketFlagByNO for eligible candidates
	_, err = client.Exam_Applications_IP.
		Update().
		Where(
			exam_applications_ip.NodalOfficeIDEQ(applicationRecord.NodalOfficeID),
			exam_applications_ip.GenerateHallTicketFlag(true),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.CenterCodeNEQ(0),
		).
		SetGenerateHallTicketFlagByNO(true).
		Save(ctx)
	if err != nil {
		log.Printf("Failed to update applications: %v", err)
		return "", fmt.Errorf("Failed to update applications: %v", err)
	}

	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", applicationRecord.CircleName), nil
}*/

//Here it sets in Exam appl table..
/*func ApproveHallTicketGenerationByNOForIPExam(client *ent.Client, examCode int32, facilityID string, approveHallTicket bool) (string, error) {
	ctx := context.Background()

	// Check if ExamCode is nil or invalid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return "", fmt.Errorf("No such valid exam code exists")
	}

	if facilityID == "" {
		return "", fmt.Errorf("Please provide Facility ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_PS table
	count, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.NodalOfficeIDEQ(facilityID),
			exam_applications_ip.GenerateHallTicketFlag(true),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.CenterCodeNEQ(0),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_IP: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", facilityID)
		return "", fmt.Errorf("No valid applications available for the circle")
	}
	// Perform the update to set GenerateHallTicketFlagByNO for eligible candidates
	_, err = client.Exam_Applications_IP.
		Update().
		Where(
			exam_applications_ip.NodalOfficeIDEQ(facilityID),
			exam_applications_ip.GenerateHallTicketFlag(true),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.CenterCodeNEQ(0),
		).
		//SetGenerateHallTicketFlagByNO(approveHallTicket).
		Save(ctx)
	if err != nil {
		log.Printf("Failed to update applications: %v", err)
		return "", fmt.Errorf("Failed to update applications: %v", err)
	}

	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", facilityID), nil
}*/

// Here it sets in circlemaster table
// func ApproveHallTicketGenerationByNOForPSExam(client *ent.Client, examCode int32, facilityID string, approveHallTicket bool) (string, error) {
// 	ctx := context.Background()

// 	// Check if ExamCode is nil or invalid
// 	if examCode <= 0 {
// 		log.Println("No such valid exam code exists")
// 		return "", fmt.Errorf("No such valid exam code exists")
// 	}

// 	if facilityID == "" {
// 		return "", fmt.Errorf("Please provide Facility ID")
// 	}

// 	// Check if circleOfficeID exists in Exam_Applications_PS table
// 	count, err := client.Exam_Applications_PS.
// 		Query().
// 		Where(
// 			exam_applications_ps.NodalOfficeIDEQ(facilityID),
// 			exam_applications_ps.GenerateHallTicketFlag(true),
// 			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_ps.CenterCodeNEQ(0),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve applications from Exam_Applications_PS: %v", err)
// 		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_PS: %v", err)
// 	}

// 	if count == 0 {
// 		log.Printf("No valid applications available for the circle: %s", facilityID)
// 		return "", fmt.Errorf("No valid applications available for the circle")
// 	}

// 	// Get the CircleMaster entity based on the Facility ID
// 	circleMaster, err := client.CircleSummaryForNO.
// 		Query().
// 		Where(circlesummaryforno.CircleOfficeIdEQ(facilityID)).
// 		Only(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve CircleMaster: %v", err)
// 		return "", fmt.Errorf("Failed to retrieve CircleMaster: %v", err)
// 	}

// 	// Update the GenerateHallTicketFlagByNO field in the CircleMaster entity
// 	circleMaster = circleMaster.
// 		Update().
// 		SetApproveHallTicketGenrationIP(approveHallTicket).
// 		SaveX(ctx)

// 	// Update the GenerateHallTicketFlagByNO in Exam_Applications_PS table
// 	_, err = client.Exam_Applications_PS.
// 		Update().
// 		Where(
// 			exam_applications_ps.NodalOfficeIDEQ(facilityID),
// 			exam_applications_ps.GenerateHallTicketFlag(true),
// 			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_ps.CenterCodeNEQ(0),
// 		).
// 		SetGenerateHallTicketFlagByNO(approveHallTicket).
// 		Save(ctx)
// 	if err != nil {
// 		log.Printf("Failed to update applications: %v", err)
// 		return "", fmt.Errorf("Failed to update applications: %v", err)
// 	}
// 	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", facilityID), nil
// }

// Get PS Exam statistics
// func getApprovalFlagForHallTicketPS(client *ent.Client, circleOfficeID string) (bool, error) {
// 	circleMaster, err := client.CircleSummaryForNO.
// 		Query().
// 		Where(circlesummaryforno.CircleOfficeIdEQ(circleOfficeID)).
// 		Only(context.Background())
// 	if err != nil {
// 		return false, fmt.Errorf("failed to get CircleMaster for CircleOfficeID %v: %v", circleOfficeID, err)
// 	}

// 	return circleMaster.ApproveHallTicketGenrationPS, nil
// }

// func UpdateCenterCodeForApplicationsPS(ctx context.Context, client *ent.Client, cityPreference, reportingOfficeName string, centerCodeToUpdate, seatsToAllot int32) error {
// 	// Input Validation
// 	if cityPreference == " " {
// 		return errors.New("City Preference cannot be nil")
// 	}

// 	if reportingOfficeName == " " {
// 		return errors.New("reporting Office Name cannot be nil")
// 	}
// 	// Querying Applications
// 	applications, err := client.Exam_Applications_PS.
// 		Query().
// 		Where(
// 			exam_applications_ps.CentrePreferenceEQ(cityPreference),
// 			//exam_applications_ps.CentreIDEQ(ExamCityId),

// 			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),

// 			exam_applications_ps.ControllingOfficeFacilityIDEQ(reportingOfficeName),
// 			//exam_applications_ps.ControllingOfficeFacilityIDEQ(ControllingAuthorityFacilityID),

// 			exam_applications_ps.HallTicketNumberNEQ(""),
// 			exam_applications_ps.ExamCityCenterCodeIsNil(),
// 			exam_applications_ps.StatusEQ("active"),
// 		).
// 		Order(ent.Asc(exam_applications_ps.FieldApplnSubmittedDate)).
// 		Limit(int(seatsToAllot)). // Limit the number of records to be updated
// 		All(ctx)

// 	if err != nil {
// 		return err
// 	}

// 	// Handling No Applications - UI to handle
// 	/*if len(applications) == 0 {
// 		return errors.New("No applications are available for allocation of center for reporting office name: " + reportingOfficeName)
// 	}

// 	// Handling Insufficient Applications
// 	if len(applications) < int(seatsToAllot) {
// 		return errors.New("The number of Applications are less than the seats selected for allocation")
// 	}
// 	*/
// 	// Updating Applications
// 	for _, application := range applications {
// 		_, err := client.Exam_Applications_PS.
// 			UpdateOne(application).
// 			SetExamCityCenterCode(centerCodeToUpdate).
// 			//SetExamCityCenterCode(ExamCenterID).
// 			Save(ctx)

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// Counting Updated Applications
// 	updateCount, err := client.Exam_Applications_PS.
// 		Query().
// 		Where(exam_applications_ps.ExamCityCenterCodeEQ(centerCodeToUpdate)).
// 		//Where(exam_applications_ps.ExamCityCenterCodeEQ(ExamCenterID))).

// 		Count(ctx)

// 	if err != nil {
// 		return err
// 	}

// 	// Updating Center Table
// 	centerDet, err := client.Center.
// 		Query().
// 		Where(center.IDEQ(centerCodeToUpdate)).
// 		//Where(center.IDEQ(ExamCenterID)).

// 		Only(ctx)

// 	if err != nil {
// 		return err
// 	}

// 	maxSeats := centerDet.MaxSeats
// 	pendingSeats := int32(maxSeats) - int32(updateCount)

// 	_, err = client.Center.
// 		UpdateOne(centerDet).
// 		SetNoAlloted(int32(updateCount)).
// 		SetPendingSeats(pendingSeats).
// 		Save(ctx)

// 	if err != nil {
// 		return err
// 	}

//		// Success
//		return nil
//	}

/*

func GetExamApplicatonsPreferenenceCityWiseStatssss(ctx context.Context, client *ent.Client, cityPreference string) ([]ExamStats, error) {
	var result []ExamStats

	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.CentrePreferenceEQ(cityPreference),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.HallTicketNumberNEQ(""),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	groupedApplications := make(map[string]map[string]ExamStats)

	for _, examApplication := range applications {
		circleName := examApplication.CircleName
		reportingOfficeName := examApplication.ReportingOfficeName
		centerCode := examApplication.CenterCode

		if _, ok := groupedApplications[circleName]; !ok {
			groupedApplications[circleName] = make(map[string]ExamStats)
		}

		// Check if ExamStats for the reporting office already exists
		stats, ok := groupedApplications[circleName][reportingOfficeName]
		if !ok {
			stats = ExamStats{
				CircleName:               circleName,
				ReportingOfficeName:      reportingOfficeName,
				NoOfCandidatesChosenCity: 0,
				NoOfCandidatesAlloted:    0,
			}
		}

		//Increment counts based on center code
		if examApplication.CentrePreference == cityPreference {

			//if centerCode == 0 {
			stats.NoOfCandidatesChosenCity++
			//} else
		}
		if centerCode >= 0 {
			stats.NoOfCandidatesAlloted++
		}

		// Update the stats in the map
		groupedApplications[circleName][reportingOfficeName] = stats
	}

	// Convert the grouped data to the desired struct
	for _, reportingOffices := range groupedApplications {
		for _, stats := range reportingOffices {
			result = append(result, stats)
		}
	}

	return result, nil
}
*/

// revathi
/*
func GetExamApplicatonsPreferenenceCityWiseStatsrevathi(ctx context.Context, client *ent.Client, cityPreference string) ([]ExamStats, error) {
	var result []ExamStats

	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.CentrePreferenceEQ(cityPreference),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.HallTicketNumberNEQ(""),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	groupedApplications := make(map[string]ExamStats)

	for _, examApplication := range applications {
		// Check if the application's preference city matches the filter
		if examApplication.CentrePreference == cityPreference {
			circleName := examApplication.CircleName
			reportingOfficeName := examApplication.ReportingOfficeName
			centerCode := examApplication.CenterCode

			// Check if ExamStats for the reporting office already exists
			stats, ok := groupedApplications[reportingOfficeName]
			if !ok {
				stats = ExamStats{
					CircleName:               circleName,
					ReportingOfficeName:      reportingOfficeName,
					NoOfCandidatesChosenCity: 0,
					NoOfCandidatesAlloted:    0,
				}
			}

			// Increment counts based on center code
			if centerCode == 0 {
				stats.NoOfCandidatesChosenCity++
			}

			// Check if center code is greater than 0
			if centerCode > 0 {
				stats.NoOfCandidatesAlloted++
			}

			// Update the stats in the map
			groupedApplications[reportingOfficeName] = stats
		}
	}

	// Convert the grouped data to the desired struct
	for _, stats := range groupedApplications {
		result = append(result, stats)
	}

	return result, nil
}
*/

/* func QueryPSApplicationsByPendingWithCandidateNew(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_PS, int32, error) {
	fmt.Println("entered sub function")
	if facilityID == "" || id1 == "" {
		return nil, 422, errors.New("Facility ID  and Exam Year cannot be empty")
	}

	// Fetch all applications matching the criteria
	records, err := client.Exam_Applications_PS.Query().
		Where(
			exam_applications_ps.ControllingOfficeFacilityIDEQ(facilityID),
			exam_applications_ps.StatusEQ("active"),
			exam_applications_ps.ExamYearEQ(id1),
		).
		Order(ent.Asc("employee_id")). /*, ent.Desc("application_number")) // Order by employee_id and application_number
		All(ctx)

	if err != nil {
		return nil, 500, err
	}

	if len(records) == 0 {
		return nil, 422, fmt.Errorf("no application pending with candidate  %s", facilityID)
	}

	// Create a map to store the latest applications for each employee
	latestApplications := make(map[int64]*ent.Exam_Applications_PS)

	// Iterate through the records and update the latest application
	for _, record := range records {
		employeeID := record.EmployeeID

		// Check if the application is the latest for this employee
		latestApp, exists := latestApplications[employeeID]
		if !exists || record.ApplicationNumber > latestApp.ApplicationNumber {
			if record.ApplicationStatus == "PendingWithCandidate" {
				latestApplications[employeeID] = record
			} else {
				// If latest status is not "PendingWithCandidate," exclude employee
				// Exclude even if employeeID was added previously
				delete(latestApplications, employeeID)
			}
		} else if record.ApplicationStatus != "PendingWithCandidate" {
			// If the current record is not the latest and has status other than "PendingWithCandidate,"
			// exclude employee
			delete(latestApplications, employeeID)
		}
	}

	// Create a slice to store the result
	var result []*ent.Exam_Applications_PS
	for _, application := range latestApplications {
		result = append(result, application)
	}

	if len(result) == 0 {
		return nil, 422, fmt.Errorf("no Applications matching criteria for the Office ID %s", facilityID)
	}

	return result, 200, nil
} */

// func GetPSExamStatisticsDOOfficeWiseLatests(ctx context.Context, client *ent.Client, examCode int32, facilityID string) ([]map[string]interface{}, error) {
// 	// Check if exam code is valid
// 	if examCode <= 0 {
// 		log.Println("No such valid exam code exists")
// 		return nil, fmt.Errorf("No such valid exam code exists")
// 	}

// 	// Check if facilityID is provided
// 	if facilityID == "" {
// 		log.Println("Facility ID cannot be null")
// 		return nil, fmt.Errorf("Facility ID cannot be null")
// 	}

// 	// Query to get the applications from Exam_Applications_PS table matching the provided facilityID
// 	applications, err := client.Exam_Applications_PS.
// 		Query().
// 		Where(exam_applications_ps.NodalOfficeFacilityIDEQ(facilityID)).
// 		Order(ent.Asc(exam_applications_ps.FieldEmployeeID), ent.Desc(exam_applications_ps.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve applications from Exam_Applications_PS: %v", err)
// 		return nil, fmt.Errorf("Failed to retrieve applications from Exam_Applications_PS: %v", err)
// 	}

// 	// Create a map to store the latest application details for each employee
// 	latestApplications := make(map[int64]*ent.Exam_Applications_PS)

// 	// Loop through the applications to find the latest application for each employee and their reporting office
// 	for _, app := range applications {
// 		employeeID := app.EmployeeID

// 		// Check if this application is the latest for the employee
// 		if latestApp, ok := latestApplications[employeeID]; !ok || app.ID > latestApp.ID {
// 			latestApplications[employeeID] = app
// 		}
// 	}

// 	// Create a map to store reporting office-wise summaries
// 	doOfficeSummaries := make(map[string]*DOOfficeWiseSummaryPS)

// 	// Loop through the latest applications to update counts
// 	for _, app := range latestApplications {
// 		reportingOfficeID := app.ReportingOfficeFacilityID

// 		// Create and initialize doOfficeSummary outside the condition
// 		doOfficeSummary := doOfficeSummaries[reportingOfficeID]
// 		if doOfficeSummary == nil {
// 			doOfficeSummaries[reportingOfficeID] = &DOOfficeWiseSummaryPS{
// 				ControllingOfficeFacilityID: reportingOfficeID,
// 				ControllingOfficeName:       app.ReportingOfficeName,
// 				Permitted:                   0,
// 				NotPermitted:                0,
// 				Pending:                     0,
// 				PendingWithCandidate:        0,
// 				Received:                    0,
// 				UniqueEmployees:             make(map[int64]struct{}),
// 			}
// 		}

// 		if _, ok := doOfficeSummary.UniqueEmployees[app.EmployeeID]; !ok {
// 			doOfficeSummary.UniqueEmployees[app.EmployeeID] = struct{}{}
// 			doOfficeSummary.Received++

// 			// Update counts based on GenerateHallTicketFlag
// 			if app.GenerateHallTicketFlag == nil {
// 				if app.ApplicationStatus == "PendingWithCandidate" {
// 					// For pending, check if GenerateHallTicketFlag is nil
// 					doOfficeSummary.PendingWithCandidate++
// 				} else {
// 					doOfficeSummary.Pending++
// 				}
// 			} else if *app.GenerateHallTicketFlag {
// 				doOfficeSummary.Permitted++
// 			} else {
// 				doOfficeSummary.NotPermitted++
// 			}
// 		}
// 	}

// 	// Create an empty slice to store the final result
// 	result := []map[string]interface{}{}
// 	serialNumber := 0

// 	// Add reportingOfficeID wise counts and names to the result
// 	for _, summary := range doOfficeSummaries {
// 		serialNumber++

// 		// Display only the latest reporting office counts
// 		result = append(result, map[string]interface{}{
// 			"S.No.":                        serialNumber,
// 			"ControllingOfficeFacilityID":  summary.ControllingOfficeFacilityID,
// 			"ControllingOfficeName":        summary.ControllingOfficeName,
// 			"No: Of Applications Received": summary.Received,
// 			"No. Permitted":                summary.Permitted,
// 			"No. Not Permitted":            summary.NotPermitted,
// 			"No. Pending":                  summary.Pending,
// 			"No. Pending With Candidate":   summary.PendingWithCandidate,
// 		})
// 	}

// 	return result, nil
// }

// PS Group B End

// obsolete
/* func UpdateExamCentresIPExams(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_IP) ([]string, error) {
	var updatedRecords []string

	for _, newappl := range newappls {
		applications, err := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ReportingOfficeFacilityIDEQ(newappl.ReportingOfficeFacilityID),
				exam_applications_ip.GenerateHallTicketFlag(false),
				exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Order(ent.Desc(exam_applications_ip.FieldID)).
			All(ctx)
		if err != nil {
			log.Printf("Failed to query applications: %v\n", err)
			return nil, fmt.Errorf("failed to query applications: %v", err)
		}

		for _, application := range applications {
			application.ExamCityCenterCode = newappl.ExamCityCenterCode

			_, err = application.Update().Save(ctx)
			if err != nil {
				log.Printf("Failed to update application: %v\n", err)
				return nil, fmt.Errorf("failed to update application: %v", err)
			}

			record := fmt.Sprintf("Employee ID: %d, Application Number: %s, Center Code: %d",
				application.EmployeeID, application.ApplicationNumber, application.ExamCityCenterCode)
			updatedRecords = append(updatedRecords, record)
			log.Printf("Application updated. %s\n", record)
		}
	}

	return updatedRecords, nil
}
*/

/* func GetExamCentresByAdminFacilityOfficeID(ctx context.Context, client *ent.Client, adminFacilityOfficeID string) ([]*ent.Center, error) {
	// Check if adminFacilityOfficeID is empty
	if adminFacilityOfficeID == "" {
		return nil, fmt.Errorf("please provide a valid AdminFacilityOfficeId")
	}

	// Query exam centres based on AdminFacilityOfficeId
	centers, err := client.Center.
		Query().
		Where(center.NodalOfficeFacilityId(adminFacilityOfficeID)).
		All(ctx)
	if err != nil {
		log.Println("error querying exam centres: ", err)
		return nil, fmt.Errorf("failed to get exam centres: %w", err)
	}

	// Check if no exam centres found
	if len(centers) == 0 {
		return nil, fmt.Errorf("no exam centres allotted for the Circle Office ID")
	}

	return centers, nil
}
*/

/* func getNodalOfficers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		NodalOfficers, err := start.QueryNodalOfficer(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": NodalOfficers})

	}

	return gin.HandlerFunc(fn)

} */

/* func CreateNodalOfficer(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newNodalOfficer := new(ent.NodalOfficer)

		if err := gctx.ShouldBindJSON(&newNodalOfficer); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newNodalOfficer, err := start.CreateNodalOfficer(client, newNodalOfficer)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Nodal Officer")

	}
	return gin.HandlerFunc(fn)
}

func GetNodalOfficerID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx := context.Background()

		id := gctx.Param("id")
		//var examID int32
		NodalOfficerID, _ := strconv.ParseInt(id, 10, 32)

		nodalofficer, err := start.QueryNodalOfficerID(ctx, client, int32(NodalOfficerID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": nodalofficer})

	}

	return gin.HandlerFunc(fn)
}
func UpdateNodalOfficer(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newNodalOfficer := new(ent.NodalOfficer)
		id := gctx.Param("id")
		NodalOfficerID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newNodalOfficer); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		nodalofficer, err := start.UpdateNodalOfficer(client, int32(NodalOfficerID), newNodalOfficer)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": nodalofficer})

	}

	return gin.HandlerFunc(fn)
}
*/
/*
	 func GetUserID(client *ent.Client) gin.HandlerFunc {
		fn := func(gctx *gin.Context) {

			ctx := context.Background()

			id := gctx.Param("id")
			//var examID int32
			//UserID, _ := strconv.ParseInt(id, 10, 32)

			center, err := start.QueryUserID(ctx, client, id)
			if err != nil {
				gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			gctx.JSON(http.StatusOK, gin.H{"data": center})

		}

		return gin.HandlerFunc(fn)
	}
*/
/* func updateUser(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newUser := new(ent.User)

		id := gctx.Param("id")

		//CenterID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newUser); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := start.UpdateUser(client, id, newUser)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": user})

	}

	return gin.HandlerFunc(fn)
} */

/* func getUsers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		users, err := start.QueryUser(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": users})

	}

	return gin.HandlerFunc(fn)

} */

/*func getExamCalendarsWithDetails(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		ExamCode, _ := strconv.ParseInt(id, 10, 32)
		examcals, err := start.QueryExamCalendarsWithVacancyAndPapers(ctx, client, int32(ExamCode))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": examcals})
	}
	return gin.HandlerFunc(fn)
}*/

/*
	 func postPreferences(client *ent.Client) gin.HandlerFunc {
		fn := func(gctx *gin.Context) {

			newExamCity := new(ent.ExamCityCenter)

			if err := gctx.ShouldBindJSON(&newExamCity); err != nil {

				gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			newExamCity, err := start.PostPreferences(client, newExamCity)

			if err != nil {

				gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			gctx.JSON(http.StatusOK, gin.H{"data": "Successfully Inserted the data"})

		}
		return gin.HandlerFunc(fn)
	}
*/

//kept in env
// config.MaxConns = 50 // Set the maximum number of connections
// config.MinConns = 10
// config.MaxConnLifetime = 15 * time.Minute
// config.MaxConnIdleTime = 15 * time.Minute

//connstr kept in env
//dev-
// connStr := "postgres://postgres:secret@172.28.12.202:2000/recruitment?sslmode=disable"

// uat
// connStr := "postgres://postgres:secret@172.24.19.181:2000/recruitment?sslmode=disable"
// connStr := "postgres://postgres:VK@123@localhost:5432/recruitment?sslmode=disable"
// connStr := "postgres://postgres:VK@123@localhost:5432/recruitment?sslmode=disable"
func CreateExamCalendar(client *ent.Client, newExamCalendar *ent.ExamCalendar) (*ent.ExamCalendar, error) {
	ctx := context.Background()
	u, err := client.ExamCalendar.Create().
		SetExamYear(newExamCalendar.ExamYear).
		SetExamName(newExamCalendar.ExamName).
		SetNotificationDate(newExamCalendar.NotificationDate).
		SetModelNotificationDate(newExamCalendar.ModelNotificationDate).
		SetApplicationEndDate(newExamCalendar.ApplicationEndDate).
		SetTentativeResultDate(newExamCalendar.TentativeResultDate).
		SetApprovedOrderDate(newExamCalendar.ApprovedOrderDate).
		SetApprovedOrderNumber(newExamCalendar.ApprovedOrderNumber).
		SetCreatedDate(newExamCalendar.CreatedDate).
		SetVacancyYears(newExamCalendar.VacancyYears).
		SetExamPapers(newExamCalendar.ExamPapers).
		SetExamCode(newExamCalendar.ExamCode).
		//SetPaperCode(newExamCalendar.PaperCode).
		//SetVacancyYearCode(newExamCalendar.VacancyYearCode).
		Save(ctx)

	if err != nil {
		log.Println("error at Creating Exam Calendar: ", newExamCalendar)
		return nil, fmt.Errorf("failed creating Exam Calendar: %w", err)
	}
	log.Println("Exam Calendar was created: ", u)

	return u, nil
}

// Update First Time username password old one
/* func oldNewValidateLoginUser(client *ent.Client, newuser *ent.UserMaster) (*ent.UserMaster, error) {
	ctx := context.Background()

	// Check if the username exists
	username := newuser.UserName
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if newuser.Password == "" {
		return nil, errors.New("password cannot be empty")
	}

	// Check if the username exists
	exists, err := client.UserMaster.
		Query().
		Where(
			usermaster.UserName(username)).
		//usermaster.PasswordNEQ("")).
		Exist(ctx)

	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("no such username: %s", username)
	}

	// Retrieve the user record with the provided username
	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserName(username)).
		Only(ctx)
	if err != nil {
		return nil, errors.New(" failed to retrieve user")
	}

	// Validate the input OTP with the stored OTP
	if user.OTP != newuser.OTP {
		return nil, errors.New(" invalid OTP")
	}

	// Update the user's password and set the status as active
	user.Password = newuser.Password
	user.Status = true

	// Save the changes to the database
	_, err = client.UserMaster.
		Update().
		Where(usermaster.EmployeeIDEQ(user.EmployeeID)).
		SetPassword(user.Password).
		SetStatus(user.Status).
		Save(ctx)

	if err != nil {
		return nil, errors.New(" failed to update user")
	}

	return user, nil
} */

/* func getUserMobileNumber(empID int64) string {
	ctx := context.Background()
	client := ent.NewClient() // Initialize your ent client here

	// Query the UserMaster entity based on the employee ID
	userMaster, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empID)).
		Only(ctx)
	if err != nil {
		// Handle the error
		return ""
	}

	// Retrieve the mobile number from the UserMaster entity
	mobileNumber := userMaster.Mobile

	return mobileNumber
} */

// Admin Login
// Validate username  pswd against db , then send sms if matches. Return PasswordMatch status : true/false.
// Validate OTP and return Login status  : true/ false.
/* func ValidateAdminUserLogin(client *ent.Client, newUser ca_reg.StruUserGenerateOTP) (*ent.UserMaster, int32, error) {
	// Check if the newUser and password are not nil

	if len(newUser.UserName) == 0 || len(newUser.NewPassword) == 0 {
		return nil, 422, errors.New("username or password cannot be nil")
	}

	// Trim the username and password
	newUser.UserName = strings.TrimSpace(newUser.UserName)
	newUser.NewPassword = strings.TrimSpace(newUser.NewPassword)

	// Check if the username exists
	ctx := context.Background()

	// Retrieve the user record with the provided username
	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(newUser.UserName),
			usermaster.RoleUserCodeIn(2, 3, 4, 5)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, errors.New("no user exists for this Role ")
		} else {
			return nil, 400, fmt.Errorf("error in fetching user: %v", err)
		}
	}
	// Compare the password from the input with the user's password stored in the database
	if user.Password != newUser.NewPassword {
		return nil, 422, errors.New("incorrect Password")
	}

	/// Send SMS and save OTP
	_, _, err = SendSMSCandidateOTP(ctx, client, user)
	if err != nil {
		return nil, 422, errors.New("failed to send SMS")
	}

	return user, 200, nil
} */

// validate login  status after pswd matcches , with otp. old

// func oldValidateAdminLogin(client *ent.Client, newUser *ent.UserMaster) (*ent.UserMaster, string, error) {
// 	// Check if the newUser and password are not nil
// 	if newUser == nil {
// 		return nil, "", errors.New(" userMaster cannot be nil")
// 	}

// 	if newUser.OTP <= 0 {
// 		return nil, "", errors.New(" otp cannot be  nil")
// 	}

// 	// Trim the username and password
// 	newUser.UserName = strings.TrimSpace(newUser.UserName)

// 	// Check if the username exists
// 	ctx := context.Background()
// 	exists, err := client.UserMaster.
// 		Query().
// 		Where(usermaster.UserNameEQ(newUser.UserName)).
// 		Exist(ctx)

// 	if err != nil {
// 		return nil, "", fmt.Errorf("failed to check username existence: %v", err)
// 	}

// 	if !exists {
// 		return nil, "", fmt.Errorf(" invalid Username or username not found: %s", newUser.UserName)
// 	}

// 	// Retrieve the user record with the provided username
// 	user, err := client.UserMaster.
// 		Query().
// 		Where(usermaster.UserNameEQ(newUser.UserName)).
// 		//WithCircleUsersRef().
// 		Only(ctx)

// 	if err != nil {
// 		return nil, "", fmt.Errorf("failed to retrieve user: %v", err)
// 	}

// 	// Retrieve the OTP saved in the database
// 	dbOTP := user.OTP

// 	// Compare the input OTP with the OTP from the database
// 	if newUser.OTP != dbOTP {
// 		// Log the mismatched OTPs
// 		log.Printf("Input OTP: %d, Database OTP: %d, Mobile Number: %s, Username: %s", newUser.OTP, dbOTP, user.Mobile, user.UserName)
// 		return nil, "", errors.New(" incorrect OTP")
// 	}

// 	// Check if FacilityID is null
// 	if user.FacilityID == "" {
// 		return nil, "", errors.New(" no valid associated OfficeID for the AdminUser: " + newUser.UserName)
// 	}
// 	var officeName string

// 	// Check if role_user_code is 3
// 	/* 	if user.RoleUserCode == 3 {
// 	   		// Get the associated CircleMaster record via the Facility edge
// 	   		circleMaster, err := client.CircleMaster.
// 	   			Query().
// 	   			Where(circlemaster.CircleOfficeIdEQ(user.FacilityID)).
// 	   			Only(ctx)

// 	   		if err != nil {
// 	   			return nil, "", fmt.Errorf("failed to retrieve CircleMaster: %v", err)
// 	   		}

// 	   		// Fetch CircleOfficeName
// 	   		officeName = circleMaster.CircleOfficeName

// 	   		log.Printf("CircleOfficeName: %s", officeName)
// 	   	} else if user.RoleUserCode == 2 {
// 	   		// Get the associated DivisionMaster record via the Facility edge
// 	   		divisionMaster, err := client.DivisionMaster.
// 	   			Query().
// 	   			Where(divisionmaster.DivisionOfficeIDEQ(user.FacilityID)).
// 	   			Only(ctx)

// 	   		if err != nil {
// 	   			return nil, "", fmt.Errorf("failed to retrieve DivisionMaster: %v", err)
// 	   		}

// 	   		// Fetch DivisionOfficeName
// 	   		officeName = divisionMaster.DivisionOfficeName

// 	   		log.Printf("DivisionOfficeName: %s", officeName)

// 	   	if user.RoleUserCode == 5 {
// 	   		circleMaster, err := client.CircleMaster.
// 	   			Query().
// 	   			Where(circlemaster.CircleOfficeIdEQ(user.FacilityID)).
// 	   			Only(ctx)

// 	   		if err != nil {
// 	   			return nil, "", fmt.Errorf("failed to retrieve CircleMaster: %v", err)
// 	   		}

// 	   		// Fetch CircleOfficeName

// 	   		officeName = circleMaster.CircleOfficeName

// 	   		log.Printf("CircleOfficeName: %s", officeName)
// 	   	}
// 	*/
// 	return user, officeName, nil
// }

// User creation & OTP sending

/*
	 func OldNewCreateUserByEmpId(ctx context.Context, client *ent.Client, newUsers *ent.UserMaster) (*ent.UserMaster, error) {
		if newUsers.UserName <= "" {
			return nil, errors.New("Please input a nonzero Emp ID as an input parameter")
		}

		exists, err := client.UserMaster.Query().
			Where(usermaster.UserNameEQ(newUsers.UserName)).Exist(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to check empid existence: %v", err)
		}
		if exists {
			user, err := client.UserMaster.Query().
				Where(
					usermaster.UserNameEQ(newUsers.UserName),
					//usermaster.PasswordNEQ(""),
					//usermaster.StatusEQ(true),
				).
				Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to query existing user: %v", err)
			}

			if user != nil {
				return nil, errors.New("The user already exists")
			}
		}

		employee, err := client.EmployeeMaster.Query().
			Where(employeemaster.EmployeeIDEQ(newUsers.EmployeeID)).
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("The Employee Id is not available in Employee Master: %v", err)
		}

		if employee.FacilityID == "" {
			return nil, errors.New("Please update your Office ID with your parent office")
		}

		if employee.MobileNumber == "" {
			return nil, errors.New("Please get your mobile number updated with your parent Office")
		}
		if employee.EmailID == "" {
			return nil, errors.New("Please get your Email ID updated with your parent Office")
		}

		// Trigger the SMS OTP.
		if employee.FacilityID != "" && employee.MobileNumber != "" && employee.EmailID != "" && employee.EmployeeName != "" {
			user, err := client.UserMaster.
				Create().
				SetEmployeeID(newUsers.EmployeeID).
				SetRoleUserCode(1).
				SetUserName(strconv.Itoa(int(newUsers.EmployeeID))).
				SetStatus(false).
				SetMobile(employee.MobileNumber).
				SetEmailID(employee.EmailID).
				SetEmployeeName(employee.EmployeeName).
				SetFacilityID(employee.FacilityID).
				//SetNewPasswordRequest(true).
				//SetPassword(newUsers.Password).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create user: %v", err)
			}

			// Call trigger SMS function
			_, _, err = SendSMSCandidateOTP(ctx, client, newUsers)
			if err != nil {
				return nil, fmt.Errorf("failed to send SMS and save OTP: %v", err)
			}

			return user, nil
		}

		return nil, nil
	}
*/

// Inserting bulk users data
/*
func InsertBulkUsers(gctx *gin.Context, client *ent.Client, nbu []*ent.UserMaster) {([]*ent.UserMaster, error)
	insertedUsers := make([]*ent.UserMaster, 0)
	var failedUsernames []string

	for _, user := range nbu {
		insertedUser, err := client.UserMaster.
			Create().
			SetEmployeeID(user.EmployeeID).
			SetEmployeeName(user.EmployeeName).
			SetFacilityID(user.FacilityID).
			SetMobile(user.Mobile).
			SetEmailID(user.EmailID).
			SetUserName(user.UserName).
			SetPassword(user.Password).
			SetRoleUserCode(user.RoleUserCode).
			Save(gctx)

		if err != nil {
			// Store the failed username and continue the loop
			failedUsernames = append(failedUsernames, user.UserName)
			continue
		}

		// Append inserted user to the result
		insertedUsers = append(insertedUsers, insertedUser)

		fmt.Printf("The user %s inserted successfully\n", insertedUser.UserName)

	}

	//print successful insertion data
	if len(insertedUsers) > 0 {
		gctx.JSON(http.StatusOK, gin.H{
			"Message": "User Details Inserted successfully",
			"data":    insertedUsers})
	}

	//Error Handling
	if len(failedUsernames) > 0 {
		fmt.Printf("The user insertion Failed \n")
		HandleValidationFailed(gctx, failedUsernames)
	}

	//return insertedUsers, nil
}
*/

// function to delete the user from UserMaster
/* func DeleteUserbyEmployeeID(ctx context.Context, client *ent.Client, eid int64) (string, error) {

	if eid == 0 {
		return "", fmt.Errorf(" employee ID cannot be Nil")
	}
	_, err := client.UserMaster.Delete().Where(usermaster.EmployeeID(eid)).Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return "", fmt.Errorf(" no such Employee ID: %d", eid)
		}
		return "", fmt.Errorf(" failed to query user: %v", err)
	}
	return "The employee details is deleted", nil
} */

/* func CreateUser(client *ent.Client, newUser *ent.User) (*ent.User, error) {
	ctx := context.Background()
	u, err := client.User.Create().
		SetEmployeedID(newUser.EmployeedID).
		SetIDRemarks(newUser.IDRemarks).
		SetEmployeedName(newUser.EmployeedName).
		SetNameRemarks(newUser.NameRemarks).
		//newUser.DOB=newUser.DOB+"T00:00:00Z"
		SetDOB(newUser.DOB).
		SetDOBRemarks(newUser.DOBRemarks).
		SetGender(newUser.Gender).
		SetGenderRemarks(newUser.GenderRemarks).
		SetCadreid(newUser.Cadreid).
		SetCadreidRemarks(newUser.CadreidRemarks).
		SetOfficeID(newUser.OfficeID).
		SetOfficeIDRemarks(newUser.OfficeIDRemarks).
		SetPH(newUser.PH).
		SetPHRemarks(newUser.PHRemarks).
		SetPHDetails(newUser.PHDetails).
		SetPHDetailsRemarks(newUser.PHDetailsRemarks).
		SetAPSWorking(newUser.APSWorking).
		SetAPSWorkingRemarks(newUser.APSWorkingRemarks).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating User: ", newUser)
		return nil, fmt.Errorf("failed creating User: %w", err)
	}
	log.Println("User was created: ", u)

	return u, nil
} */
