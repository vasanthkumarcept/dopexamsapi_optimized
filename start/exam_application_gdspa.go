package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"recruit/ent"
	"recruit/ent/center"

	"recruit/ent/exam_applications_gdspa"

	"recruit/ent/recommendationsgdspaapplications"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"
	//"recruit/util"
)

func SubGetGDSPASApplicationsFacilityIDYear(ctx context.Context, client *ent.Client, facilityID string, year string) ([]*ent.Exam_Applications_GDSPA, int32, error) {
	// Array of exams

	if facilityID == "" || year == "" {
		return nil, 422, errors.New("facility ID and Examyear cannot be blank/null")
	}
	records, err := client.Exam_Applications_GDSPA.Query().
		Where(exam_applications_gdspa.ControllingOfficeFacilityIDEQ(facilityID),
			exam_applications_gdspa.ExamYearEQ(year),
			exam_applications_gdspa.StatusEQ("active"),
		).
		Order(ent.Desc(exam_applications_gdspa.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 400, fmt.Errorf("failed querying GDS to PA/SA exams Applications: %w", err)
	}
	if len(records) == 0 {
		return nil, 422, fmt.Errorf("no applications for the Year %s and facility ID  %s", year, facilityID)
	}

	return records, 200, nil
}

type Service1 struct {
	VacancyDate            *time.Time `json:"Vacancy Date" binding:"required"`
	Age                    string     `json:"Age" binding:"required"`
	AgeEligibility         string     `json:"AgeEligibility" binding:"required"`
	Service                string     `json:"Service" binding:"required"`
	ServiceEligibility     string     `json:"ServiceEligibility" binding:"required"`
	ServiceEligibilityYear int64      `json:"ServiceEligibilityYear" binding:"required"`
}

/* func convertToInterfaceSlice1(details []Service1) []interface{} {
	interfaceSlice := make([]interface{}, len(details))
	for i, d := range details {
		interfaceSlice[i] = d
	}
	return interfaceSlice
} */

/*
	 func processPreferences(preferences []interface{}) {
		fmt.Println(preferences)
	}
*/
func CreateGDSPAApplications(client *ent.Client, newAppln *ca_reg.ApplicationGDSPM) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateTempHallTicket(newAppln.TempHallTicket, newAppln.EmployeeID); err != nil {
		return nil, 422, " -STR100", false, err
	}
	if newAppln == nil {
		return nil, 422, " -STR001", false, fmt.Errorf("newAppln is nil or its ID is nil")
	}
	// Start a transaction.
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	defer func() {
		handleTransaction(tx, &err)
	}()

	statuses := []string{"CAVerificationPending", "ResubmitCAVerificationPending", "PendingWithCandidate", "VerifiedByCA"}
	existing, status, stgError, err := checkIfApplicationExists(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear, newAppln.ExamCode, statuses)
	if status == 500 {
		return nil, 500 + status, " -STR003 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR004 " + stgError, false, err

	}

	if existing {
		return nil, 422 + status, " -STR005 " + stgError, false, errors.New("already application submitted for this candidate")
	}
	var currentTime = time.Now().UTC().Truncate(time.Second) // Ensure UTC for consistent comparison
	currentTime = currentTime.Add(5*time.Hour + 30*time.Minute)
	applicationLastDate := newAppln.ApplicationLastDate.UTC().Truncate(time.Second)

	fmt.Print("current time: ", currentTime, " application last date: ", applicationLastDate, "date from payload", newAppln.ApplicationLastDate)
	if currentTime.After(applicationLastDate) {
		return nil, 422, " -STR007", false, fmt.Errorf("application submission deadline has passed as current time is %v", currentTime)
	}
	// Generate Application number
	// Generate application number in the format "GDSPA2023XXXXXX"

	applicationNumber, err := util.GenerateApplicationNumber(client, newAppln.EmployeeID, newAppln.ExamYear, "GDSPA")
	if err != nil {
		tx.Rollback()
		return nil, 422, " -STR006", false, fmt.Errorf("failed to generate application number %v", err)
	}

	createdAppln, status, stgError, err := saveApplication(tx, newAppln, applicationNumber, newAppln.ExamCode, ctx)
	if err != nil {
		return nil, 500 + status, " -STR007 " + stgError, false, err
	}

	return createdAppln, 200, "", true, nil
}

/* func generateGDSPAApplicationNumber(client *ent.Client, employeeID int64) (string, error) {
	nextApplicationNumber, err := getNextGDSPAApplicationNumberFromDatabase(client)
	if err != nil {
		return "", err
	}

	// Get the current year
	currentYear := time.Now().Year()

	// Format the application number as "GDSPAYYYYXXXXXX"
	applicationNumber := fmt.Sprintf("GDSPA%d%06d", currentYear, nextApplicationNumber)

	return applicationNumber, nil
}

func getNextGDSPAApplicationNumberFromDatabase(client *ent.Client) (int64, error) {
	ctx := context.TODO()
	lastApplication, err := client.Exam_Applications_GDSPA.
		Query().
		Order(ent.Desc(exam_applications_gdspa.FieldID)).
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
// Query GDSPAExam Application with Emp ID.
func QueryGDSPAExamApplicationsByEmpID(ctx context.Context, client *ent.Client, empid int64, examYear string) (*ent.Exam_Applications_GDSPA, error) {
	newAppln, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			(exam_applications_gdspa.EmployeeIDEQ(empid)),
			(exam_applications_gdspa.ExamYearEQ(examYear)),
			(exam_applications_gdspa.StatusEQ("active")),
		).
		//Order(ent.Desc(exam_applications_gdspa.FieldID)).
		WithGDSPAApplicationsRef().
		Only(ctx)

	if err != nil {
		log.Println("error getting Emp ID Application Details: ", err)
		return nil, fmt.Errorf("failed querying GDSPA Exam Application details: %w", err)
	}

	// Extract only the desired fields from the CirclePrefRefGDSPA edge
	var circlePrefs []*ent.Division_Choice_PA
	for _, edge := range newAppln.Edges.CirclePrefRefGDSPA {
		circlePrefs = append(circlePrefs, &ent.Division_Choice_PA{
			PlacePrefNo:    edge.PlacePrefNo,
			PlacePrefValue: edge.PlacePrefValue,
		})
	}

	// Update the CirclePrefRefGDSPA edge with the filtered values
	newAppln.Edges.CirclePrefRefGDSPA = circlePrefs

	var recomondPref []*ent.RecommendationsGDSPAApplications
	for _, edge := range newAppln.Edges.GDSPAApplicationsRef {
		recomondPref = append(recomondPref, &ent.RecommendationsGDSPAApplications{
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
	newAppln.Edges.GDSPAApplicationsRef = recomondPref
	newAppln.UpdatedAt = newAppln.UpdatedAt.Truncate(24 * time.Hour)

	return newAppln, nil
}

// Update / Verification of GDSPA Exam Application By CA
// Update Resubmission By Candidate.

func UpdateApplicationRemarksGDSPA(client *ent.Client, newAppln *ca_reg.VerifyApplicationGDStoPA, nonQualifyService string) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Check if newAppln is not nil.
	if newAppln == nil {
		return nil, 400, " -STR001", false, errors.New("payload is empty")
	}
	// Start a transaction.
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, fmt.Errorf("failed to start transaction: %v", err)
	}

	defer handleTransaction(tx, &err)
	// Fetch the existing application.
	oldAppln, status, stgError, err := fetchExistingGdspaApplication(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear)
	if status == 500 {
		return nil, 500 + status, " -STR003 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR004 " + stgError, false, fmt.Errorf("no active application found for this candidate")

	}
	if err != nil {
		return nil, 500 + status, " -STR005 " + stgError, false, err
	}
	// Update the application status based on the current status and remarks.
	var applicationStatus string
	applicationStatus, status, stgError, err = determineApplicationStatus(ctx, oldAppln, newAppln.CA_Remarks, newAppln.ExamCode)
	if status == 500 {
		return nil, 500 + status, " -STR006 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR007 " + stgError, false, fmt.Errorf("no active application found for this Remarks")

	}
	if err != nil {
		return nil, 500 + status, " -STR008 " + stgError, false, err
	}

	// Create a new application record.
	updatedAppln, err := createUpdateGdspaApplication(ctx, tx, oldAppln, newAppln, applicationStatus, nonQualifyService)
	if err != nil {
		return nil, 500, " -STR009", false, err
	}
	// Handle the recommendations.
	err = handleGdspaRecommendations(ctx, tx, updatedAppln, newAppln)
	if err != nil {
		return nil, 500, " -STR010", false, err
	}
	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB011", false, err
	}
	return appResponse, 200, "", true, nil
}

func ResubmitApplicationRemarksGDSPA(client *ent.Client, newAppln *ca_reg.ReApplicationGDStoPA) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateTempHallTicket(newAppln.TempHallTicket, newAppln.EmployeeID); err != nil {
		return nil, 422, " -STR001", false, fmt.Errorf("issue for employee %d with temp hall ticket number: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
	}
	if newAppln == nil {
		return nil, 400, " -STR002", false, fmt.Errorf("payload is empty")
	}
	//transaction implementation--------------
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, fmt.Errorf("failed to start transaction %v", err)
	}
	defer func() {
		handleTransaction(tx, &err)
	}()
	// Check if the EmployeeID exists.
	oldAppln, status, stgError, err := fetchOldApplicationGdsPa(ctx, tx, newAppln)
	if status == 500 {
		return nil, 500 + status, " -STR004 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR005 " + stgError, false, fmt.Errorf("no active application found for this candidate")
	}

	if err != nil {
		return nil, 500 + status, " -STR005 " + stgError, false, err
	}
	statuses := []string{"ResubmitCAVerificationPending"}
	existing, status, stgError, err := checkIfApplicationExists(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear, newAppln.ExamCode, statuses)
	if status == 500 {
		return nil, 500 + status, " -STR004 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR005 " + stgError, false, err

	}
	if existing {
		return nil, 422 + status, " -STR006 " + stgError, false, fmt.Errorf("already application submitted for this candidate")
	}
	var currentTime = time.Now().UTC().Truncate(time.Second) // Ensure UTC for consistent comparison
	currentTime = currentTime.Add(5*time.Hour + 30*time.Minute)
	correctionLastDate := newAppln.ApplicationCorrectionLastDate.UTC().Truncate(time.Second)

	fmt.Print("current time: ", currentTime, " correction last date: ", correctionLastDate, "date from payload", newAppln.ApplicationCorrectionLastDate)
	if currentTime.After(correctionLastDate) {
		return nil, 422, " -STR007", false, fmt.Errorf("application correction deadline has passed as current time is %v", currentTime)
	}
	// Perform the application update and resubmission
	updatedAppln, status, stgError, err := processResubmission(ctx, tx, oldAppln, newAppln, int(newAppln.ExamCode))
	if status == 500 {
		return nil, 500 + status, " -STR007 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR008 " + stgError, false, err

	}
	if err != nil {
		return nil, 500, " -STR009 ", false, err
	}
	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB018", false, err
	}

	return appResponse, 200, "", true, nil
}

// func GetGDSPAApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64) (string, error) {
// 	application, err := client.Exam_Applications_GDSPA.
// 		Query().
// 		Where(
// 			exam_applications_gdspa.EmployeeIDEQ(employeeID),
// 			exam_applications_gdspa.ApplicationStatusEQ("PendingWithCandidate"),
// 		).
// 		Order(ent.Desc(exam_applications_gdspa.FieldID)).
// 		First(ctx)

// 	if err != nil {
// 		return "", fmt.Errorf("failed to retrieve the GDS to PA Application: %v", err)
// 	}

// 	return application.AppliactionRemarks, nil
// }

// func getGDSPAInputRecordByVacancyYear(inputRecords []*ent.RecommendationsGDSPAApplications, vacancyYear int32) *ent.RecommendationsGDSPAApplications {
// 	// Find the corresponding input record based on vacancy year
// 	for _, record := range inputRecords {
// 		if record.VacancyYear == vacancyYear {
// 			return record
// 		}
// 	}
// 	return nil
// }

// UpdateNodalRecommendationsByEmpID updates the recommendations for a given employee ID
/*func UpdateNodalRecommendationsByEmpID(client *ent.Client, empID int64, newRecommendations []*ent.RecommendationsGDSPAApplications) ([]*ent.RecommendationsGDSPAApplications, error) {
	ctx := context.Background()

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		// Handle the error, such as logging or returning an error
		log.Printf("Error loading location: %v", err)
		return nil, err
	}

	currentTime := time.Now().In(loc).Truncate(time.Second)
	// Check if empID exists
	exists, err := client.RecommendationsGDSPAApplications.Query().
		Where(recommendationsgdspaapplications.EmployeeIDEQ(empID)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check if employee with ID %d exists: %v", empID, err)
	}
	if !exists {
		return nil, fmt.Errorf("employee with ID %d does not exist", empID)
	}

	// Retrieve all records for the employee ID
	records, err := client.RecommendationsGDSPAApplications.Query().
		Where(recommendationsgdspaapplications.EmployeeIDEQ(empID)).
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

	// Query the RecommendationsGDSPAApplications table for the specific employee
	record, err := client.RecommendationsGDSPAApplications.Query().
		Where(
			recommendationsgdspaapplications.EmployeeIDEQ(empID),
			//recommendationsgdspaapplications.ApplicationStatusEQ("VerifiedRecommendationsByNO"),
		).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve RecommendedByNO record: %v", err)
	}

	if record != nil {
		// Retrieve the corresponding Exam_Applications_GDSPA record using edges
		applicationRecord, err := client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.EmployeeIDEQ(empID),
				exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve Exam_Applications_GDSPA record: %v", err)
		}

		// Update the Exam_Applications_GDSPA record
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
	recordsupdated, err := client.RecommendationsGDSPAApplications.Query().
		Where(recommendationsgdspaapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records for employee with ID after updation %d: %v", empID, err)
	}

	return recordsupdated, nil
}*/

// UpdateNodalRecommendationsByEmpID updates the recommendations for a given employee ID
func UpdateGDSPANodalRecommendationsByEmpID(client *ent.Client, applicationRecord *ca_reg.NAVerifyApplicationGDStoPA) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateGdspaInput(applicationRecord); err != nil {
		return nil, 422, " -STR001", false, errors.New("employee id should not be empty")
	}
	empID := applicationRecord.EmployeeID
	id1 := applicationRecord.ExamYear

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	defer handleTransaction(tx, &err)

	// if applicationRecord.Edges.GDSPAApplicationsRef == nil || len(applicationRecord.Edges.GDSPAApplicationsRef) == 0 {
	// 	return nil, fmt.Errorf("the recommendations are mandatory")
	// }
	// Check if empID exists in exam_applications_gdspa and the status is "VerifiedByCA" or "VerifiedByNA"
	exists, status, stgError, err := checkGdspaApplicationExists(tx, ctx, applicationRecord)

	if status == 500 {
		return nil, 500, " -STR004 " + stgError, false, err
	}

	if status == 422 {
		return nil, 422, " -STR005 " + stgError, false, fmt.Errorf("no active application found for this candidate")
	}
	if !exists {
		return nil, 422, " -STR007 ", false, fmt.Errorf("no active application found for this candidate")
	}

	// Retrieve all records for the employee ID from RecommendationsGDSPAApplications
	//records, err := QueryGDSPARecommendationsByEmpId(ctx, client, empID)
	records, err := getGdspaRecommendationsByEmpID(ctx, tx, empID)
	if err != nil {
		return nil, 500, " -STR008 ", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR009 ", false, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}
	// Format the current time to "yyyymmddhhmmss"
	stat := "inactive_" + time.Now().Format("20060102150405")

	updatedRecord, status, stgError, err := getActiveExamApplicationGdspa(ctx, tx, empID, id1)
	if status == 500 {
		return nil, 500 + status, " -STR010 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR011 " + stgError, false, fmt.Errorf("no active application exists")

	}
	if err != nil {
		return nil, 500 + status, " -STR012 " + stgError, false, err
	}

	_, err = updatedRecord.
		Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		log.Println("Failed to update application:", err)
		return nil, 422, " -STR010", false, fmt.Errorf("failed to update application: %v", err)
	}

	// Hall Ticket Generated Flag
	hallticketgeneratedflag := checkGdspaHallTicketGenerated(applicationRecord, updatedRecord)
	updatedAppln, err := createUpdatedGdspaAppln(tx, applicationRecord, updatedRecord, hallticketgeneratedflag, ctx)
	if err != nil {
		return nil, 500, " -STR014 ", false, err
	}

	recommendationsRef, err := createGdspaRecommendationsRef(ctx, tx, applicationRecord, updatedAppln)
	if err != nil {
		return nil, 500, " -STR015", false, err
	}
	updatedAppln.Update().
		//ClearIPApplicationsRef().
		AddGDSPAApplicationsRef(recommendationsRef...).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR014", false, err
	}
	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB017 ", false, err
	}
	return appResponse, 200, "", true, nil
}

// Get All CA Pending records ...
func QueryGDSPAApplicationsByCAVerificationsPending(ctx context.Context, client *ent.Client, facilityID string, examYear string) ([]*ent.Exam_Applications_GDSPA, int32, string, bool, error) {
	// Array of exams

	if facilityID == "" || examYear == "" {
		return nil, 422, " -STR001", false, fmt.Errorf("facility ID and Examyear cannot be null")
	}
	log.Println("Input Facility ID:", facilityID, "Examyear:", examYear) // Log the facility ID and Examyear

	records, err := client.Exam_Applications_GDSPA.Query().
		Where(
			exam_applications_gdspa.And(
				exam_applications_gdspa.Or(
					exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				),
				exam_applications_gdspa.ControllingOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
			),
		).
		Order(ent.Desc(exam_applications_gdspa.FieldID)). // Order by descending updated_at timestamp
		//Limit(1).                                      // Limit to 1 record (the latest)
		WithCirclePrefRefGDSPA(). // Add the Where clause with multgdspale statuses using Or
		All(ctx)
	if err != nil {
		log.Println("error at GDSPA Exam Applications fetching: ", err)
		return nil, 422, " -STR002", false, fmt.Errorf("failed querying GDSPA exams Applications: %w", err)
	}
	//for _, record := range records {
	//	log.Println("Reporting Facility ID:", record.ReportingOfficeID)
	//}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf(" nil Applications for the CA pending verification for the Office ID %s", facilityID)
	} //log.Println("CA verifications pending returned: ", records)

	return records, 200, "", true, nil
}

// Get All CA verified records
func QueryGDSPAApplicationsByCAVerified(ctx context.Context, client *ent.Client, facilityID string, examYear string) ([]*ent.Exam_Applications_GDSPA, int32, string, bool, error) {
	if facilityID == "" || examYear == "" {
		return nil, 422, " -STR001", false, errors.New(" facility ID  and Exam Year cannot be null")
	}
	records, err := client.Exam_Applications_GDSPA.Query().
		Where(
			exam_applications_gdspa.And(
				exam_applications_gdspa.ApplicationStatusEQ("VerifiedByCA"),
				exam_applications_gdspa.ControllingOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
			),
		).
		//WithGDSPAApplicationsRef().
		WithCirclePrefRefGDSPA().
		All(ctx)
	if err != nil {
		log.Println("error at GDSPA Exam Applications fetching: ", err)
		return nil, 422, " -STR002", false, fmt.Errorf("failed querying GDSPA exams Applications for NA Verified records: %w", err)
	}
	//for _, record := range records {
	//	log.Println("Reporting Facility ID:", record.ReportingOfficeID)
	//}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("nil Applications for the CA verified for the Office ID %s", facilityID)
	}
	//log.Println("CA verified records returned: ", records)
	return records, 200, " ", true, nil
}

// Get CA Verified with Emp ID
func QueryGDSPAApplicationsByCAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_GDSPA, error) {

	record, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ApplicationStatusEQ("VerifiedByCA"), // Check for "CAVerified" status
			exam_applications_gdspa.EmployeeIDEQ(employeeID),
			exam_applications_gdspa.ExamYearEQ(examYear),
			exam_applications_gdspa.StatusEQ("active"),
		).
		WithGDSPAApplicationsRef().
		WithCirclePrefRefGDSPA().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("record not found for employee ID: %d with 'CAVerified' status", employeeID)
		}
		log.Println("error at GDSPA Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying GDSPA exams Applications: %w", err)
	}

	//log.Println("CA verified record returned: ", record)
	return record, nil
}

// Get CA Pending with EmpID
// func QueryGDSPAApplicationsByCAPendingByEmpID(ctx context.Context, client *ent.Client, empID int64) ([]*ent.Exam_Applications_GDSPA, error) {
// 	// Check if employee ID exists
// 	employeeExists, err := client.Exam_Applications_GDSPA.
// 		Query().
// 		Where(
// 			exam_applications_gdspa.EmployeeIDEQ(empID),
// 			exam_applications_gdspa.Or(
// 				exam_applications_gdspa.ApplicationStatusEQ("CAVerificationPending"),
// 				exam_applications_gdspa.ApplicationStatusEQ("ResubmitCAVerificationPending"),
// 			),
// 		).
// 		WithCirclePrefRefGDSPA().
// 		WithGDSPAApplicationsRef().
// 		Exist(ctx)
// 	if err != nil {
// 		log.Println("error checking employee existence: ", err)
// 		return nil, fmt.Errorf("failed checking employee existence: %w", err)
// 	}
// 	if !employeeExists {
// 		return nil, fmt.Errorf("employee not found with ID: or the verification is not pending with CA %d", empID)
// 	}

// 	// Retrieve the record
// 	record, err := client.Exam_Applications_GDSPA.
// 		Query().
// 		Where(
// 			exam_applications_gdspa.EmployeeIDEQ(empID),
// 			exam_applications_gdspa.Or(
// 				exam_applications_gdspa.ApplicationStatusEQ("CAVerificationPending"),
// 				exam_applications_gdspa.ApplicationStatusEQ("ResubmitCAVerificationPending"),
// 			),
// 		).
// 		WithGDSPAApplicationsRef().
// 		WithCirclePrefRefGDSPA().
// 		All(ctx)
// 	if err != nil {
// 		log.Println("error at GDSPA Exam Applications fetching: ", err)
// 		return nil, fmt.Errorf("failed querying GDSPA exams Applications: %w", err)
// 	}

// 	//log.Println("CA pending records returned: ", record)
// 	return record, nil
// }

func QueryGDSPAApplicationsByCAPendingByEmpID(ctx context.Context, client *ent.Client, empID int64, examYear string) (*ent.Exam_Applications_GDSPA, error) {

	// Retrieve the latest record based on UpdatedAt timestamp
	record, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.EmployeeIDEQ(empID),
			exam_applications_gdspa.ExamYearEQ(examYear),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.Or(
				exam_applications_gdspa.ApplicationStatusEQ("CAVerificationPending"),
				exam_applications_gdspa.ApplicationStatusEQ("ResubmitCAVerificationPending"),
			),
		).
		Order(ent.Desc("updated_at")). // Order by UpdatedAt in descending order
		WithGDSPAApplicationsRef().
		WithCirclePrefRefGDSPA().
		First(ctx)
	if err != nil {
		log.Println("error at GDSPA Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying GDSPA exams Applications: %w", err)
	}

	return record, nil
}

// Get latest old Application Remarks given to Candidate for CA Verification
func GetOldGDSPACAApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64) (*ent.Exam_Applications_GDSPA, error) {
	employeeExists, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.EmployeeIDEQ(employeeID),
			exam_applications_gdspa.Status("active"),
		).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check employee existence: %v", err)
	}
	if !employeeExists {
		return nil, fmt.Errorf("employee not found with ID: %d", employeeID)
	}

	application, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.EmployeeIDEQ(employeeID),
			exam_applications_gdspa.ApplicationStatusEQ("PendingWithCandidate"),
			exam_applications_gdspa.Status("active"),
		).
		//Order(ent.Desc(exam_applications_gdspa.FieldID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("application not found for employee ID: %d with 'PendingWithCandidate' status", employeeID)
		}
		return nil, fmt.Errorf("failed to retrieve application: %v", err)
	}

	return application, nil
}

// Get Recommendations/ Remarks with Emp ID
func GetGDSPARecommendationsByEmpID(client *ent.Client, empID int64) ([]*ent.RecommendationsGDSPAApplications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	// Check if empID is null
	if empID == 0 {
		return nil, fmt.Errorf("no employee ID provided to process")
	}
	// Check if empID exists
	exists, err := client.RecommendationsGDSPAApplications.Query().
		Where(recommendationsgdspaapplications.EmployeeIDEQ(empID)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check if employee with ID %d exists: %v", empID, err)
	}
	if !exists {
		return nil, fmt.Errorf("employee with ID %d does not exist", empID)
	}

	// Retrieve all records for the employee ID
	records, err := client.RecommendationsGDSPAApplications.Query().
		Where(recommendationsgdspaapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records for employee with ID %d: %v", empID, err)
	}

	return records, nil
}

// Get All NA Verified Records
func QueryGDSPAApplicationsByNAVerified(ctx context.Context, client *ent.Client, facilityID string, examYear string) ([]*ent.Exam_Applications_GDSPA, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || examYear == "" {
		return nil, 422, " -STR001", false, fmt.Errorf("facility ID cannot be null")
	}
	records, err := client.Exam_Applications_GDSPA.Query().
		Where(
			exam_applications_gdspa.And(
				exam_applications_gdspa.ApplicationStatusEQ("VerifiedByNA"),
				exam_applications_gdspa.ControllingOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ExamYearEQ(examYear),
			),
		).
		WithCirclePrefRefGDSPA().
		All(ctx)
	if err != nil {
		log.Println("error at GDSPA Exam Applications fetching: ", err)
		return nil, 422, " -STR002", false, fmt.Errorf("failed querying GDSPA exams Applications for NA Verified records: %w", err)
	}
	//for _, record := range records {
	//	log.Println("Reporting Facility ID:", record.ReportingOfficeID)
	//}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("nil Applications for the NA verified for the Office ID %s", facilityID)
	}
	//log.Println("CA verified records returned: ", records)
	return records, 200, "", true, nil
}

// Get All NA Verified Records with Emp ID
func QueryGDSPAApplicationsByNAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64) (*ent.Exam_Applications_GDSPA, error) {
	/* employeeExists, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.EmployeeIDEQ(employeeID),
			exam_applications_gdspa.Status("active"),
		).
		Exist(ctx)
	if err != nil {
		log.Println("error checking employee existence: ", err)
		return nil, fmt.Errorf("failed checking employee existence: %w", err)
	}
	if !employeeExists {
		return nil, fmt.Errorf("employee not found with ID: %d", employeeID)
	} */

	record, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ApplicationStatusEQ("VerifiedByNA"), // Check for "CAVerified" status
			exam_applications_gdspa.EmployeeIDEQ(employeeID),
			exam_applications_gdspa.Status("active"),
		).
		WithGDSPAApplicationsRef().
		WithCirclePrefRefGDSPA().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("record not found for employee ID: %d with 'CAVerified' status", employeeID)
		}
		log.Println("error at GDSPA Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying GDSPA exams Applications: %w", err)
	}

	//log.Println("CA verified record returned: ", record)
	return record, nil
}

func QueryGDSPAApplicationsByNAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID, examYear string) ([]*ent.Exam_Applications_GDSPA, error) {
	// Array of exams
	if facilityID == "" || examYear == "" {
		return nil, errors.New(" facility ID and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_GDSPA.Query().
		Where(
			exam_applications_gdspa.And(
				exam_applications_gdspa.ApplicationStatusEQ("VerifiedByNA"),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
			),
		).
		WithCirclePrefRefGDSPA().
		WithGDSPAApplicationsRef().
		All(ctx)
	if err != nil {
		log.Println("error at GDSPA Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying GDSPA exams Applications for NA Verified records: %w", err)
	}
	//for _, record := range records {
	//	log.Println("Reporting Facility ID:", record.ReportingOfficeID)
	//}
	if len(records) == 0 {
		return nil, fmt.Errorf("nil Applications for the NA verified status for view by Nodal Officer of the Office ID %s", facilityID)
	}
	//log.Println("CA verified records returned: ", records)
	return records, nil
}

// // Get All CA verified records for NA
func QueryGDSPAApplicationsByCAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID, examYear string) ([]*ent.Exam_Applications_GDSPA, error) {
	if facilityID == "" || examYear == "" {
		return nil, fmt.Errorf("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_GDSPA.Query().
		Where(
			exam_applications_gdspa.And(
				exam_applications_gdspa.ApplicationStatusEQ("VerifiedByCA"),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ExamYearEQ(examYear),
			),
		).
		WithGDSPAApplicationsRef().
		WithCirclePrefRefGDSPA().
		All(ctx)
	if err != nil {
		log.Println("error at GDSPA Exam Applications fetching: ", err)
		return nil, fmt.Errorf("failed querying GDSPA exams Applications for CA Verified records: %w", err)
	}
	//for _, record := range records {
	//	log.Println("Reporting Facility ID:", record.ReportingOfficeID)
	//}
	if len(records) == 0 {
		return nil, fmt.Errorf(" nil Applications for the CA verified for the Office ID %s", facilityID)
	}
	//log.Println("CA verified records returned: ", records)
	return records, nil
}

// Get Recommendations with Emp ID ..
func QueryGDSPARecommendationsByEmpId(ctx context.Context, client *ent.Client, employeeID int64) ([]*ent.RecommendationsGDSPAApplications, error) {
	//Array of exams

	employeeExists, err := client.RecommendationsGDSPAApplications.
		Query().
		Where(recommendationsgdspaapplications.EmployeeIDEQ(employeeID)).
		Exist(ctx)
	if err != nil {
		log.Println("error checking employee existence: ", err)
		return nil, fmt.Errorf("failed checking employee existence: %w", err)
	}
	if !employeeExists {
		return nil, fmt.Errorf("employee not found with ID: %d", employeeID)
	}

	records, err := client.RecommendationsGDSPAApplications.
		Query().
		Where(recommendationsgdspaapplications.EmployeeIDEQ(employeeID)).
		All(ctx)
	if err != nil {
		log.Println("error querying GDSPA recommendations: ", err)
		return nil, fmt.Errorf("failed to query GDSPA recommendations: %w", err)
	}

	return records, nil
}

// // Get Exams by Exam Code.
// func QueryExamsGDSPAByExamNameCode(ctx context.Context, client *ent.Client, examNameCode string) (*ent.Exam_PA, error) {
// 	// Check if examNameCode is empty
// 	if examNameCode == "" {
// 		return nil, fmt.Errorf("Please provide exam name code")
// 	}

// 	u, err := client.Exam_PA.Query().
// 		Where(exam_pa.ExamNameCode(examNameCode)).
// 		Only(ctx)
// 	if err != nil {
// 		log.Println("error at getting Exam_GDSPA: ", err)
// 		return nil, fmt.Errorf("failed querying Exam_GDSPA: %w", err)
// 	}
// 	log.Println("Exam_GDSPA details returned: ", u)
// 	return u, nil
// }

// list of reporting offices

func GetGDSPADivisionsByCircleOfficeID(ctx context.Context, client *ent.Client, circleOfficeID string, examYear string) ([]*ent.Exam_Applications_GDSPA, error) {
	// Check if the circle office ID exists in the exam_application_gdspa table.
	exists, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_gdspa.ExamYearEQ(examYear),
			exam_applications_gdspa.StatusEQ("active"),
		).
		Exist(ctx)
	if err != nil {
		log.Printf("Failed to query exam_application_gdspa: %v\n", err)
		return nil, fmt.Errorf("failed to query exam_application_gdspa: %v", err)
	}
	if !exists {
		log.Printf("Circle office ID does not exist: %s\n", circleOfficeID)
		return nil, fmt.Errorf("circle office ID does not exist")
	}

	// Query the exam_application_gdspa table for unique records based on the provided conditions.
	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Select(
			exam_applications_gdspa.FieldReportingOfficeID,
			exam_applications_gdspa.FieldReportingOfficeName,
		).
		Where(
			exam_applications_gdspa.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_gdspa.Not(exam_applications_gdspa.GenerateHallTicketFlag(true)),
			exam_applications_gdspa.ExamCityCenterCodeIsNil(),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.ExamYearEQ(examYear),
		).
		All(ctx)
	if err != nil {
		log.Printf("Failed to query exam_application_gdspa: %v\n", err)
		return nil, fmt.Errorf("failed to query exam_application_gdspa: %v", err)
	}

	// Filter and return distinct records based on reporting office ID and name.
	distinctApplications := make(map[string]*ent.Exam_Applications_GDSPA)
	for _, app := range applications {
		key := app.ReportingOfficeID
		distinctApplications[key] = app
	}

	result := make([]*ent.Exam_Applications_GDSPA, 0, len(distinctApplications))
	for _, app := range distinctApplications {
		result = append(result, app)
	}

	log.Printf("Retrieved %d distinct divisions for Circle Office ID: %s\n", len(result), circleOfficeID)

	// Log the applications as an array of strings
	appStrings := make([]string, len(result))
	for i, app := range result {
		appStrings[i] = fmt.Sprintf("Reporting Office ID: %s, Reporting Office Name: %s", app.ReportingOfficeID, app.ReportingOfficeName)
	}
	log.Printf("Applications: %+v\n", appStrings)

	return result, nil
}

// type HallticketStatsGDSPA struct {
// 	CircleID        int32  `json:"CircleID"`
// 	StartingNumber  int    `json:"StartingNumber"`
// 	EndingNumber    int    `json:"EndingNumber"`
// 	Count           int    `json:"Count"`
// 	StartHallTicket string `json:"StartHallTicket"`
// 	EndHallTicket   string `json:"EndHallTicket"`
// }

// Generate Hall Ticket Numbers return array with stng & eng nos.
/*func GenerateHallticketNumber(ctx context.Context, client *ent.Client) ([]HallticketStats, error) {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Printf("Error loading location: %v", err)
		return nil, err
	}

	currentTime := time.Now().In(loc).Truncate(time.Second)

	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.GenerateHallTicketFlag(true),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.CenterCodeNEQ(0),
			exam_applications_gdspa.ExamCodeNEQ(0),
			exam_applications_gdspa.ExamYearNEQ(""),
			exam_applications_gdspa.CategoryCodeNEQ(0),
			exam_applications_gdspa.CircleIDNEQ(0),
			exam_applications_gdspa.RegionIDNEQ(0),
			exam_applications_gdspa.DivisionIDNEQ(0),
			exam_applications_gdspa.EmployeeIDNEQ(0),
		).
		Order(ent.Desc(exam_applications_gdspa.FieldID)).
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

// type HallticketResultGDSPA struct {
// 	CircleID string `json:"circleID"`
// 	Count    int    `json:"count"`
// }
// type CircleStatsGDSPA struct {
// 	CircleID        string `json:"CircleID"`
// 	HallTicketCount int    `json:"Count"`
// }

// Generate Hall Ticket Numbers and return JSON array of CircleID and count.

// func GenerateHallticketNumberGDSPA(ctx context.Context, client *ent.Client) ([]CircleStatsGDSPA, error) {
// 	// loc, err := time.LoadLocation("Asia/Kolkata")
// 	// if err != nil {
// 	// 	log.Printf("Error loading location: %v", err)
// 	// 	return nil, err
// 	// }

// 	// currentTime := time.Now().In(loc).Truncate(time.Second)
// 	currentTime := time.Now().Truncate(time.Second)

// 	applications, err := client.Exam_Applications_GDSPA.
// 		Query().
// 		Where(
// 			exam_applications_gdspa.GenerateHallTicketFlag(true),
// 			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_gdspa.ExamCityCenterCodeNEQ(0),
// 			//exam_applications_gdspa.CenterCodeEQ(ExamCenterCode),
// 			exam_applications_gdspa.ExamCodeNEQ(0),
// 			exam_applications_gdspa.ExamYearNEQ(""),
// 			exam_applications_gdspa.CategoryCodeNEQ(""),
// 			exam_applications_gdspa.CircleIDNEQ(0),
// 			exam_applications_gdspa.RegionIDNEQ(0),
// 			exam_applications_gdspa.DivisionIDNEQ(0),
// 			exam_applications_gdspa.EmployeeIDNEQ(0),
// 		).
// 		Order(ent.Desc(exam_applications_gdspa.FieldID)).
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

// 		hallticketNumber := util.GenerateHallticketNumberGDSPA(
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
// 	statsSlice := make([]CircleStatsGDSPA, 0, len(circleStats))
// 	for key, count := range circleStats {
// 		statsSlice = append(statsSlice, CircleStatsGDSPA{CircleID: key, HallTicketCount: count})
// 	}

// 	return statsSlice, nil
// }

// generatew ht's and return as string
// func GenerateHallticketGDSPAReturnStringMessage(ctx context.Context, client *ent.Client) (string, error) {
// func GenerateHallticketGDSPAReturnStringMessage(ctx context.Context, client *ent.Client) ([]HallticketStatsGDSPA, error) {
// 	// loc, err := time.LoadLocation("Asia/Kolkata")
// 	// if err != nil {
// 	// 	log.Printf("Error loading location: %v", err)
// 	// 	return nil, err
// 	// }

// 	// currentTime := time.Now().In(loc).Truncate(time.Second)
// 	currentTime := time.Now().Truncate(time.Second)

// 	applications, err := client.Exam_Applications_GDSPA.
// 		Query().
// 		Where(
// 			exam_applications_gdspa.GenerateHallTicketFlag(true),
// 			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			exam_applications_gdspa.ExamCityCenterCodeNEQ(0),
// 			exam_applications_gdspa.ExamCodeNEQ(0),
// 			exam_applications_gdspa.ExamYearNEQ(""),
// 			exam_applications_gdspa.CategoryCodeNEQ(""),
// 			exam_applications_gdspa.CircleIDNEQ(0),
// 			exam_applications_gdspa.RegionIDNEQ(0),
// 			exam_applications_gdspa.DivisionIDNEQ(0),
// 			exam_applications_gdspa.EmployeeIDNEQ(0),
// 		).
// 		Order(ent.Desc(exam_applications_gdspa.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	circleStats := make(map[string]HallticketStatsGDSPA)
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

// 		hallticketNumber := util.GenerateHallticketNumberGDSPA(
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
// 		stats.StartHallTicket = util.GenerateHallticketNumberGDSPA(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.CircleID,
// 			//application.RegionID,
// 			application.DivisionID,
// 			stats.StartingNumber)
// 		stats.EndHallTicket =util.GenerateHallticketNumberGDSPA(
// 			application.ExamCode,
// 			examYear,
// 			application.CategoryCode,
// 			application.CircleID,
// 			//application.RegionID,
// 			application.DivisionID,
// 			stats.EndingNumber)

// 		circleStats[key] = stats
// 	}

// 	statsSlice := make([]HallticketStatsGDSPA, 0, len(circleStats))
// 	for _, stats := range circleStats {
// 		statsSlice = append(statsSlice, stats)
// 	}

// 	return statsSlice, nil
// }

/* // Generate ht with centercode
func GenerateHallticketNumberGDSPAwithCenterCode(ctx context.Context, client *ent.Client) (string, error) {

	currentTime := time.Now().Truncate(time.Second)

	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.ExamCodeNEQ(0),
			exam_applications_gdspa.ExamYearNEQ(""),
			exam_applications_gdspa.CategoryCodeNEQ(""),
			exam_applications_gdspa.EmployeeIDNEQ(0),
			exam_applications_gdspa.HallTicketGeneratedFlagNEQ(true),
		).
		Order(ent.Asc(exam_applications_gdspa.FieldID)).
		All(ctx)
	if err != nil {
		return "", err
	}

	circleStats := make(map[string]int)
	for _, application := range applications {
		key := fmt.Sprintf("%d", application.WorkingOfficeCircleFacilityID)
		circleStats[key]++

		identificationNo := circleStats[key]
		examYear := application.ExamYear
		if len(examYear) >= 2 {
			examYear = examYear[len(examYear)-2:]
		}

		hallticketNumber := util.GenerateHallticketNumberGDSPA(
			application.ExamCode,
			examYear,
			application.CategoryCode,
			identificationNo)

		// Validate if the hallticket number is of 12 digits
		if len(hallticketNumber) != 12 {
			log.Printf("Hallticket Number: %s\n", hallticketNumber)
			continue
		}

		log.Printf("Generated hallticket number is of 12 digits: %s\n", hallticketNumber)
		log.Printf("Employee ID: %d\n", application.EmployeeID)

		_, err = application.Update().
			SetHallTicketNumber(hallticketNumber).
			SetHallTicketGeneratedFlag(true).
			SetHallTicketGeneratedDate(currentTime).
			Save(ctx)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to generate Hallticket number: %s", err.Error())
			log.Println(errMsg)
			return "", errors.New(errMsg)
		}
	}

	// Return the success message with the count of eligible candidates
	successMessage := fmt.Sprintf("Hall Ticket generated successfully for %d eligible candidates", len(applications))
	return successMessage, nil
} */

// Get Admit card details .
func GetGDSPAApplicationsWithHallTicket(client *ent.Client, examCode int32, employeeID int64) (*ent.Exam_Applications_GDSPA, *ent.Exam_Applications_GDSPA, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Check if exam code is valid
	if examCode == 0 {
		return nil, nil, errors.New(" please provide a valid exam code")
	}

	if examCode == 4 {
		// Check if the employee_ID exists in the Exam_Applications_GDSPA table
		exists, err := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.EmployeeIDEQ(employeeID),
				exam_applications_gdspa.StatusEQ("active"),
			).
			Exist(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf(" failed to check employee ID in GDSPA Applications: %v", err)
		}
		if !exists {
			return nil, nil, fmt.Errorf(" no applications are found for the employee in GDSPA Applications: %d", employeeID)
		}

		// Query the Exam_Applications_GDSPA table to retrieve the applicant details
		application, err := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.EmployeeIDEQ(employeeID),
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.HallTicketNumberNEQ(""),
				exam_applications_gdspa.ExamCityCenterCodeNEQ(0),
				exam_applications_gdspa.StatusEQ("active"),
			).
			//WithGDSPAExamCentres().
			Only(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf(" no Admit card details available for the applicant: %v", err)
		}

		// Fetch the associated RecommendationsGDSPA records matching the employee ID
		recommendations, err := client.RecommendationsGDSPAApplications.Query().
			Where(
				recommendationsgdspaapplications.EmployeeIDEQ(employeeID),
				recommendationsgdspaapplications.ApplicationIDEQ(application.ID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf(" failed to retrieve recommendations: %v", err)
		}

		// Assign the fetched recommendations to the application entity
		application.Edges.GDSPAApplicationsRef = recommendations

		return application, nil, nil
	} else {
		return nil, nil, errors.New("invalid exam code")
	}
}

// Get Circle details summary ofExam Applications for the Nodal Officer Office ID. - For GDSPA ALone
func GetEligibleGDSPAApplicationsForCircleDetails(ctx context.Context, client *ent.Client, examCode int32, circleOfficeID string, Examyear string) ([]map[string]interface{}, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return nil, errors.New(" no such valid exam code exists")
	}

	if circleOfficeID == "" {
		return nil, errors.New(" please provide Nodal Officer's office ID ")
	}

	// Check if circleOfficeID exists in Exam_Applications_PS table
	count, err := client.Exam_Applications_GDSPA.
		Query().
		Where(exam_applications_gdspa.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.ExamYearEQ(Examyear),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
		return nil, fmt.Errorf(" failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", circleOfficeID)
		return nil, fmt.Errorf(" no valid applications available for the circle")
	}

	// Query to get the applications matching the circleOfficeID
	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(exam_applications_gdspa.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.ExamYearEQ(Examyear),
		).
		Order(ent.Asc(exam_applications_gdspa.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications: %v", err)
		return nil, fmt.Errorf("failed to retrieve applications: %v", err)
	}

	uniqueEmployees := make(map[int64]struct{})
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_GDSPA) // Map to store the latest application for each employee
	circleName := ""

	for _, app := range applications {
		// If this employee's latest application is not yet stored, or if this application is newer
		if latestApp, exists := employeeLatestApplication[app.EmployeeID]; !exists || app.ID > latestApp.ID {
			employeeLatestApplication[app.EmployeeID] = app
			uniqueEmployees[app.EmployeeID] = struct{}{}
			circleName = app.WorkingOfficeCircleName
		}
	}

	permittedCount := 0
	notPermittedCount := 0
	pendingCount := 0
	pendingWithCandidateCount := 0
	HallticketGeneratedCount := 0
	HallticetNotGeneratedCount := 0

	for _, app := range employeeLatestApplication {
		if app.GenerateHallTicketFlag != nil {
			if *app.GenerateHallTicketFlag {
				permittedCount++
			} else {
				notPermittedCount++
			}
		}
		if app.GenerateHallTicketFlag == nil {
			if app.ApplicationStatus == "PendingWithCandidate" {
				// For pending, check if GenerateHallTicketFlag is nil
				pendingWithCandidateCount++
			} else {
				pendingCount++
			}
		}
	}
	for _, app := range employeeLatestApplication {
		if app.HallTicketGeneratedFlag {
			HallticketGeneratedCount++
		} else {
			HallticetNotGeneratedCount++
		}
	}

	employeeCount := len(uniqueEmployees)

	result := []map[string]interface{}{
		{
			"CircleID":                    circleOfficeID,
			"CircleName":                  circleName,
			"NoOfApplications Received":   employeeCount,
			"NoPermitted":                 permittedCount,
			"NoNotPermitted":              notPermittedCount,
			"NoPending":                   pendingCount,
			"NoPendingwithCandidate":      pendingWithCandidateCount,
			"NoHallticketGenerated Count": HallticketGeneratedCount,
		},
	}

	return result, nil
}

// Count of details based on Reporting Offices

// Generate Hall Ticket Flag .../*
/*func ApproveHallTicketGenerationByNO(client *ent.Client, applicationRecord *ent.Exam_Applications_GDSPA) (string, error) {
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

	// Check if circleOfficeID exists in Exam_Applications_GDSPA table
	count, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.NodalOfficeIDEQ(applicationRecord.NodalOfficeID),
			exam_applications_gdspa.GenerateHallTicketFlag(true),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.CenterCodeNEQ(0),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", applicationRecord.NodalOfficeID)
		return "", fmt.Errorf("No valid applications available for the circle")
	}
	// Perform the update to set GenerateHallTicketFlagByNO for eligible candidates
	_, err = client.Exam_Applications_GDSPA.
		Update().
		Where(
			exam_applications_gdspa.NodalOfficeIDEQ(applicationRecord.NodalOfficeID),
			exam_applications_gdspa.GenerateHallTicketFlag(true),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.CenterCodeNEQ(0),
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
/*func ApproveHallTicketGenerationByNOForGDSPAExam(client *ent.Client, examCode int32, facilityID string, approveHallTicket bool) (string, error) {
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
	count, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.NodalOfficeIDEQ(facilityID),
			exam_applications_gdspa.GenerateHallTicketFlag(true),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.CenterCodeNEQ(0),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", facilityID)
		return "", fmt.Errorf("No valid applications available for the circle")
	}
	// Perform the update to set GenerateHallTicketFlagByNO for eligible candidates
	_, err = client.Exam_Applications_GDSPA.
		Update().
		Where(
			exam_applications_gdspa.NodalOfficeIDEQ(facilityID),
			exam_applications_gdspa.GenerateHallTicketFlag(true),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.CenterCodeNEQ(0),
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
// func ApproveHallTicketGenerationByNOForGDSPAExam(client *ent.Client, examCode int32, facilityID string, approveHallTicket bool) (string, error) {
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
// 	count, err := client.Exam_Applications_GDSPA.
// 		Query().
// 		Where(
// 			exam_applications_gdspa.NodalOfficeIDEQ(facilityID),
// 			exam_applications_gdspa.GenerateHallTicketFlag(true),
// 			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_gdspa.CenterCodeNEQ(0),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
// 		return "", fmt.Errorf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
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

// 	// Update the GenerateHallTicketFlagByNO in Exam_Applications_GDSPA table
// 	_, err = client.Exam_Applications_GDSPA.
// 		Update().
// 		Where(
// 			exam_applications_gdspa.NodalOfficeIDEQ(facilityID),
// 			exam_applications_gdspa.GenerateHallTicketFlag(true),
// 			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 			//exam_applications_gdspa.CenterCodeNEQ(0),
// 		).
// 		SetGenerateHallTicketFlagByNO(approveHallTicket).
// 		Save(ctx)
// 	if err != nil {
// 		log.Printf("Failed to update applications: %v", err)
// 		return "", fmt.Errorf("Failed to update applications: %v", err)
// 	}
// 	return fmt.Sprintf("Approved successfully for eligible candidates in the Circle %s", facilityID), nil
// }

// Get GDSPA Exam statistics
func GetGDSPAExamStatistics(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Query to get the applications from Exam_Applications_GDSPA table
	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ExamCodeEQ(examCode),
			exam_applications_gdspa.ExamYearEQ(examYear),
			exam_applications_gdspa.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_gdspa.FieldEmployeeID), ent.Desc(exam_applications_gdspa.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	// Create a map to store the latest application for each employee
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_GDSPA)

	// Loop through the applications and store the latest application for each employee
	for _, app := range applications {
		if _, found := employeeLatestApplication[app.EmployeeID]; !found {
			employeeLatestApplication[app.EmployeeID] = app
		}
	}

	permittedCount := 0
	notPermittedCount := 0
	pendingCount := 0
	pendingWithCandidateCount := 0

	for _, app := range employeeLatestApplication {
		if app.GenerateHallTicketFlag != nil {
			if *app.GenerateHallTicketFlag {
				permittedCount++
			} else {
				notPermittedCount++
			}
		}
		if app.GenerateHallTicketFlag == nil {
			if app.ApplicationStatus == "PendingWithCandidate" {
				// For pending, check if GenerateHallTicketFlag is nil
				pendingWithCandidateCount++
			} else {
				pendingCount++
			}
		}
	}
	employeeCount := len(employeeLatestApplication)

	result := []map[string]interface{}{
		{
			"Total No. Of Applications Received": employeeCount,
			"No. Permitted":                      permittedCount,
			"No. Not Permitted":                  notPermittedCount,
			"No. Pending":                        pendingCount,
			"No. Pending with Candidate":         pendingWithCandidateCount,
		},
	}

	return result, 200, "", true, nil
}

// Directorate statistics for Circle wise Applications Summary
type CircleWiseSummaryGDSPA struct {
	CircleOfficeID            string
	CircleOfficeName          string
	Permitted                 int
	NotPermitted              int
	Pending                   int
	PendingWithCandidate      int
	Received                  int
	ApprovalFlagForHallTicket bool
	UniqueEmployees           map[int64]struct{}
}

// func getApprovalFlagForHallTicketGDSPA(client *ent.Client, circleOfficeID string) (bool, error) {
// 	circleMaster, err := client.CircleSummaryForNO.
// 		Query().
// 		Where(circlesummaryforno.CircleOfficeIdEQ(circleOfficeID)).
// 		Only(context.Background())
// 	if err != nil {
// 		return false, fmt.Errorf("failed to get CircleMaster for CircleOfficeID %v: %v", circleOfficeID, err)
// 	}

// 	return circleMaster.ApproveHallTicketGenrationPA, nil
// }

func GetGDSPAExamStatisticsCircleWise(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Query to get the applications from Exam_Applications_GDSPA table
	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ExamCodeEQ(examCode),
			exam_applications_gdspa.ExamYearEQ(examYear),
			exam_applications_gdspa.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_gdspa.FieldEmployeeID), ent.Desc(exam_applications_gdspa.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	// Create a map to store employee-wise latest applications
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_GDSPA)

	// Loop through the applications and store the latest application for each employee
	for _, app := range applications {
		if _, found := employeeLatestApplication[app.EmployeeID]; !found {
			employeeLatestApplication[app.EmployeeID] = app
		}
	}

	// Create a map to store circle-wise summaries
	circleSummaries := make(map[string]*CircleWiseSummaryGDSPA)

	// Loop through the latest applications to update counts
	for _, app := range employeeLatestApplication {
		circleOfficeID := app.NodalOfficeFacilityID
		if circleSummaries[circleOfficeID] == nil {
			approvalFlag, err := getApprovalFlagForHallTicket(client, circleOfficeID)
			if err != nil {
				log.Printf("Failed to get ApprovalFlagForHallTicket for CircleOfficeID %v: %v", circleOfficeID, err)
				continue
			}

			circleSummaries[circleOfficeID] = &CircleWiseSummaryGDSPA{
				CircleOfficeID:            circleOfficeID,
				CircleOfficeName:          app.WorkingOfficeCircleName,
				Permitted:                 0,
				NotPermitted:              0,
				Pending:                   0,
				PendingWithCandidate:      0,
				Received:                  0,
				ApprovalFlagForHallTicket: approvalFlag,
			}
		}

		circleSummary := circleSummaries[circleOfficeID]
		circleSummary.Received++

		if app.GenerateHallTicketFlag == nil {
			if app.ApplicationStatus == "PendingWithCandidate" {
				// For pending, check if GenerateHallTicketFlag is nil
				circleSummary.PendingWithCandidate++
			} else {
				circleSummary.Pending++
			}
		} else if *app.GenerateHallTicketFlag {
			circleSummary.Permitted++
		} else {
			circleSummary.NotPermitted++
		}
	}

	// Create an empty slice to store the final result
	result := []map[string]interface{}{}
	var serialNumber int

	// Add circleOfficeID wise counts and names to the result
	for circleOfficeID, summary := range circleSummaries {
		serialNumber++

		result = append(result, map[string]interface{}{
			"S.No.":                        serialNumber,
			"CircleOfficeID":               circleOfficeID,
			"CircleOfficeName":             summary.CircleOfficeName,
			"No: Of Applications Received": summary.Received,
			"No. Permitted":                summary.Permitted,
			"No. Not Permitted":            summary.NotPermitted,
			"No. Pending":                  summary.Pending,
			"No. Pending With Candidate":   summary.PendingWithCandidate,
			"ApprovalFlagForHallTicket":    summary.ApprovalFlagForHallTicket,
		})
	}

	return result, 200, "", true, nil
}

// GDSPA office wise stats for NO
// type DOOfficeWiseSummaryGDSPA struct {
// 	ReportingOfficeID    string
// 	ReportingOfficeName  string
// 	Permitted            int
// 	NotPermitted         int
// 	Pending              int
// 	PendingWithCandidate int
// 	Received             int
// 	UniqueEmployees      map[int64]struct{}
// }

// func GetGDSPAExamStatisticsDOOfficeWise(ctx context.Context, client *ent.Client, examCode int32, facilityID string) ([]map[string]interface{}, error) {
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

// 	// Query to get the applications from Exam_Applications_GDSPA table matching the provided facilityID
// 	applications, err := client.Exam_Applications_GDSPA.
// 		Query().
// 		Where(exam_applications_gdspa.NodalOfficeFacilityIDEQ(facilityID)).
// 		Order(ent.Desc(exam_applications_gdspa.FieldID)).
// 		All(ctx)
// 	if err != nil {
// 		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
// 		return nil, fmt.Errorf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
// 	}

// 	// Create a map to store reporting office-wise summaries
// 	doOfficeSummaries := make(map[string]*DOOfficeWiseSummaryGDSPA)

// 	// Loop through the applications to group by reporting office-wise and update counts
// 	for _, app := range applications {
// 		reportingOfficeID := app.ReportingOfficeID

// 		if doOfficeSummaries[reportingOfficeID] == nil {
// 			doOfficeSummaries[reportingOfficeID] = &DOOfficeWiseSummaryGDSPA{
// 				ReportingOfficeID:    reportingOfficeID,
// 				ReportingOfficeName:  app.ReportingOfficeName,
// 				Permitted:            0,
// 				NotPermitted:         0,
// 				Pending:              0,
// 				PendingWithCandidate: 0,
// 				Received:             0,
// 				UniqueEmployees:      make(map[int64]struct{}),
// 			}
// 		}

// 		// Check if the employee is unique for this reporting office
// 		doOfficeSummary := doOfficeSummaries[reportingOfficeID]
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

// 		result = append(result, map[string]interface{}{
// 			"S.No.":                        serialNumber,
// 			"ReportingOfficeID":            summary.ReportingOfficeID,
// 			"ReportingOfficeName":          summary.ReportingOfficeName,
// 			"No: Of Applications Received": summary.Received,
// 			"No. Permitted":                summary.Permitted,
// 			"No. Not Permitted":            summary.NotPermitted,
// 			"No. Pending":                  summary.Pending,
// 			"No. Pending With Candidate":   summary.PendingWithCandidate,
// 		})
// 	}

// 	return result, nil
// }

// // Get All Pending with Candidate
// Assuming Exam_Applications_GDSPA has a field named EmployeeID, you might adapt the code like this:
func QueryGDSPAApplicationsByPendingWithCandidate(ctx context.Context, client *ent.Client, facilityID string, examYear string) ([]*ent.Exam_Applications_GDSPA, int32, string, bool, error) {
	if facilityID == "" || examYear == "" {
		return nil, 422, " -STR001", false, errors.New(" facility ID and ExamYear cannot be empty")
	}

	// Fetch all applications matching the criteria
	records, err := client.Exam_Applications_GDSPA.Query().
		Where(
			exam_applications_gdspa.ControllingOfficeFacilityIDEQ(facilityID),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.ExamYearEQ(examYear),
			exam_applications_gdspa.ApplicationStatusEQ("PendingWithCandidate"),
		).
		Order(ent.Asc("employee_id")). /*, ent.Desc("application_number"))*/ // Order by employee_id and application_number
		All(ctx)
	if err != nil {
		return nil, 422, " -STR002", false, fmt.Errorf("failed querying GDSPA exams Applications: %w", err)
	}

	// Create a map to store the latest applications for each employee
	// latestApplications := make(map[int64]*ent.Exam_Applications_GDSPA)

	// // Iterate through the records and update the latest application
	// for _, record := range records {
	// 	employeeID := record.EmployeeID

	// 	// Check if the application is the latest for this employee
	// 	latestApp, exists := latestApplications[employeeID]
	// 	if !exists || record.ApplicationNumber > latestApp.ApplicationNumber {
	// 		if record.ApplicationStatus == "PendingWithCandidate" {
	// 			latestApplications[employeeID] = record
	// 		} else {
	// 			// If latest status is not "PendingWithCandidate," exclude employee
	// 			// Exclude even if employeeID was added previously
	// 			delete(latestApplications, employeeID)
	// 		}
	// 	} else if record.ApplicationStatus != "PendingWithCandidate" {
	// 		// If the current record is not the latest and has status other than "PendingWithCandidate,"
	// 		// exclude employee
	// 		delete(latestApplications, employeeID)
	// 	}
	// }

	// // Create a slice to store the result
	// var result []*ent.Exam_Applications_GDSPA
	// for _, application := range latestApplications {
	// 	result = append(result, application)
	// }

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no Applications matching criteria for the Office ID %s", facilityID)
	}

	return records, 200, " ", true, nil
}

func GetGDSPAExamStatisticsDOOfficeWiseLatests(ctx context.Context, client *ent.Client, examCode int32, facilityID string, examYear string) ([]map[string]interface{}, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return nil, errors.New(" no such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		log.Println("Facility ID cannot be null")
		return nil, errors.New(" facility ID cannot be null")
	}

	// Query to get the applications from Exam_Applications_GDSPA table matching the provided facilityID
	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_gdspa.ExamYearEQ(examYear),
			exam_applications_gdspa.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_gdspa.FieldEmployeeID), ent.Desc(exam_applications_gdspa.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
		return nil, fmt.Errorf(" failed to retrieve applications from Exam_Applications_GDSPA: %v", err)
	}

	// Create a map to store the latest application details for each employee
	latestApplications := make(map[int64]*ent.Exam_Applications_GDSPA)

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
		reportingOfficeID := app.ReportingOfficeID

		// Create and initialize doOfficeSummary outside the condition
		if _, exists := doOfficeSummaries[reportingOfficeID]; !exists {
			//doOfficeSummary := doOfficeSummaries[reportingOfficeID]
			//if doOfficeSummary == nil {
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

		doOfficeSummary := doOfficeSummaries[reportingOfficeID]
		if _, ok := doOfficeSummary.UniqueEmployees[app.EmployeeID]; !ok {
			doOfficeSummary.UniqueEmployees[app.EmployeeID] = struct{}{}
			doOfficeSummary.Received++

			if app.GenerateHallTicketFlag == nil {
				if app.ApplicationStatus == "PendingWithCandidate" {
					// For pending, check if GenerateHallTicketFlag is nil
					doOfficeSummary.PendingWithCandidate++
				} else {
					doOfficeSummary.PendingWithCA++
				}
			} else if *app.GenerateHallTicketFlag {
				doOfficeSummary.Permitted++
				//if *&app.HallTicketNumber == ("") {
				//if app.HallTicketNumber == ("") {
				if app.HallTicketNumber == "" {
					doOfficeSummary.HallTicketNotGenerated++
				} else {
					doOfficeSummary.HallTicketGenerated++
				}
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
			"SNo":                         serialNumber,
			"ControllingOfficeFacilityID": summary.ControllingOfficeFacilityID,
			"ControllingOfficeName":       summary.ControllingOfficeName,
			"NoOfApplicationsReceived":    summary.Received,
			"NoPermitted":                 summary.Permitted,
			"NoNotPermitted":              summary.NotPermitted,
			"NoPending":                   summary.PendingWithCA,
			"NoPendingWithCandidate":      summary.PendingWithCandidate,
			"NoHallTicketGenerated":       summary.HallTicketGenerated,
			"NoHallTicketNotGenerated":    summary.HallTicketNotGenerated,
		})
	}

	return result, nil
}

func GetGDSPAExamStatisticsDOOfficeWiseL(ctx context.Context, client *ent.Client, examCode int32, facilityID string, Examyear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		return nil, 422, " -STR002", false, errors.New("facility ID cannot be null")
	}

	// Fetch all applications for the given facilityID
	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(exam_applications_gdspa.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.ExamYearEQ(Examyear),
		).
		Order(ent.Asc(exam_applications_gdspa.FieldEmployeeID), ent.Desc(exam_applications_gdspa.FieldUpdatedAt)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	// Create a map to store the greatest ApplicationID for each unique employee_id
	greatestAppIDs := make(map[int64]int32)

	// Loop through applications to find the greatest ApplicationID for each employee_id
	for _, app := range applications {
		if greatestAppID, exists := greatestAppIDs[app.EmployeeID]; !exists || int32(app.ID) > greatestAppID {
			greatestAppIDs[app.EmployeeID] = int32(app.ID)
		}
	}

	// Create a map to store reporting office-wise summaries
	doOfficeSummaries := make(map[string]*DOOfficeWiseSummary)

	// Loop through the greatest ApplicationIDs to get the corresponding applications
	for employeeID, greatestAppID := range greatestAppIDs {
		for _, app := range applications {
			if app.EmployeeID == employeeID && int32(app.ID) == greatestAppID {
				// Process the application here
				reportingOfficeID := app.ControllingOfficeFacilityID

				if doOfficeSummaries[reportingOfficeID] == nil {
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

				// Check if the employee is unique for this reporting office
				doOfficeSummary := doOfficeSummaries[reportingOfficeID]
				if _, ok := doOfficeSummary.UniqueEmployees[app.EmployeeID]; !ok {
					doOfficeSummary.UniqueEmployees[app.EmployeeID] = struct{}{}
					doOfficeSummary.Received++

					if app.GenerateHallTicketFlag == nil {
						if app.ApplicationStatus == "PendingWithCandidate" {
							// For pending, check if GenerateHallTicketFlag is nil
							doOfficeSummary.PendingWithCandidate++
						} else {
							doOfficeSummary.PendingWithCA++
						}
					} else if *app.GenerateHallTicketFlag {
						doOfficeSummary.Permitted++
						//if *&app.HallTicketNumber == ("") {
						if app.HallTicketGeneratedFlag {
							doOfficeSummary.HallTicketGenerated++
						} else {
							doOfficeSummary.HallTicketNotGenerated++
						}
					} else {
						doOfficeSummary.NotPermitted++
					}
				}
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
			"SNo":                         serialNumber,
			"ControllingOfficeFacilityID": summary.ControllingOfficeFacilityID,
			"ControllingOfficeName":       summary.ControllingOfficeName,
			"NoOfApplicationsReceived":    summary.Received,
			"NoPermitted":                 summary.Permitted,
			"NoNotPermitted":              summary.NotPermitted,
			"NoPending":                   summary.PendingWithCA,
			"NoPendingWithCandidate":      summary.PendingWithCandidate,
			"NoHallTicketGenerated":       summary.HallTicketGenerated,
			"NoHallTicketNotGenerated":    summary.HallTicketNotGenerated,
		})
	}

	return result, 200, "", true, nil
}

type ExamStatsGDSPA struct {
	CircleName                  string
	ControllingOfficeName       string
	ControllingOfficeFacilityID string
	NodalOfficeFacilityID       string
	NoOfCandidatesChosenCity    int
	NoOfCandidatesAlloted       int
}

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

func GetExamApplicatonsPreferenenceCityWiseStatsGDSPA(ctx context.Context, client *ent.Client, ExamYear string, SubExamcode string, SubCityid string) ([]ExamStatsGDSPA, int32, string, bool, error) {
	var result []ExamStatsGDSPA

	// Convert strings to int64
	StrExamcode, err1 := strconv.ParseInt(SubExamcode, 10, 32)
	if err1 != nil {
		return nil, 422, " -STR001", false, errors.New("invalid exam code")
	}

	StrCityid, err2 := strconv.ParseInt(SubCityid, 10, 32)
	if err2 != nil {
		return nil, 422, " -STR002", false, errors.New("invalid City Id")
	}

	// Convert int64 to int32
	Examcode := int32(StrExamcode)
	Cityid := int32(StrCityid)

	if ExamYear == " " {
		return nil, 422, " -STR003", false, errors.New("exam year should not null")
	}

	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ExamYearEQ(ExamYear),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.HallTicketNumberNEQ(""),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.ExamCodeEQ(Examcode),
			exam_applications_gdspa.CenterIdEQ(Cityid),
		).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR004", false, err
	} else {
		if len(applications) == 0 {
			return nil, 422, " -STR005", false, errors.New("no matching data with Exam City")
		}
	}

	groupedApplications := make(map[string]ExamStatsGDSPA)

	for _, examApplication := range applications {
		nodalOfficeFacilityID := examApplication.NodalOfficeFacilityID
		nodalOfficeName := examApplication.NodalOfficeName

		controllingOfficeFacilityID := examApplication.ControllingOfficeFacilityID
		controllingOfficeName := examApplication.ControllingOfficeName
		centerCode := examApplication.ExamCityCenterCode

		// Check if ExamStats for the reporting office already exists
		stats, ok := groupedApplications[controllingOfficeName]
		if !ok {
			stats = ExamStatsGDSPA{
				CircleName:                  nodalOfficeName,
				ControllingOfficeName:       controllingOfficeName,
				NodalOfficeFacilityID:       nodalOfficeFacilityID,
				ControllingOfficeFacilityID: controllingOfficeFacilityID,
				NoOfCandidatesChosenCity:    0,
				NoOfCandidatesAlloted:       0,
			}
		}
		stats.NoOfCandidatesChosenCity++

		if centerCode > 0 {
			stats.NoOfCandidatesAlloted++
		}

		// Update the stats in the map
		groupedApplications[controllingOfficeName] = stats
	}

	// Convert the grouped data to the desired struct
	for _, stats := range groupedApplications {
		result = append(result, stats)
	}
	if len(result) == 0 {
		return nil, 422, " -STR006", false, errors.New("no matching data with Exam City")
	}
	return result, 200, "", true, nil
}

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

func UpdateCenterCodeForApplicationsGDSPA(ctx context.Context, client *ent.Client, controllingOfficeFacilityID string, examCenterID, seatsToAllot, examCityID int32) (int, []*ent.Exam_Applications_GDSPA, int32, string, bool, error) {
	// Input Validation
	strExamCenterID := strconv.FormatInt(int64(examCenterID), 10)
	if controllingOfficeFacilityID == "" {
		return 0, nil, 422, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR001", false, errors.New("controlling Office Facility ID cannot be nil")
	}
	// Querying Applications
	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ExamCityCenterCodeIsNil(),
			exam_applications_gdspa.CenterIdEQ(examCityID),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.ControllingOfficeFacilityIDEQ(controllingOfficeFacilityID),
			exam_applications_gdspa.HallTicketNumberNEQ(""),
			exam_applications_gdspa.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_gdspa.FieldApplnSubmittedDate)).
		Limit(int(seatsToAllot)). // Limit the number of records to be updated
		All(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR002", false, err
	}

	var updatedCount int
	// Updating Applications
	for _, application := range applications {
		_, err := client.Exam_Applications_GDSPA.
			UpdateOne(application).
			SetExamCityCenterCode(examCenterID).
			Save(ctx)

		if err != nil {
			return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR003", false, err
		}
		updatedCount++
	}

	// Counting Updated Applications
	updateCount, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ExamCityCenterCodeEQ(examCenterID),
			exam_applications_gdspa.StatusEQ("active")).
		Count(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR004", false, err
	}

	// Updating Center Table
	centerDet, err := client.Center.
		Query().
		Where(center.IDEQ(examCenterID)).
		Only(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR005", false, err
	}

	maxSeats := centerDet.MaxSeats
	pendingSeats := int32(maxSeats) - int32(updateCount)

	_, err = client.Center.
		UpdateOne(centerDet).
		SetNoAlloted(int32(updateCount)).
		SetPendingSeats(pendingSeats).
		Save(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR006", false, err
	}

	// Success
	return updatedCount, applications, 200, "", true, nil
}

func GenerateHallticketNumberrGDSPA(ctx context.Context, client *ent.Client, year string, examCode int32, nodalOfficerFacilityID string) (string, error) {
	//currentTime := time.Now().Truncate(time.Second)

	// Retrieve the last hall ticket number and extract its last four digits
	lastFourDigitsMap := make(map[int]bool)
	lastHallTicketNumber, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ExamYearEQ(year),
			//exam_applications_gdspa.ExamCodeEQ(examCode),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			exam_applications_gdspa.GenerateHallTicketFlagEQ(true),
			exam_applications_gdspa.HallTicketNumberNEQ(""),
			exam_applications_gdspa.StatusEQ("active"),
		).
		Order(ent.Desc(exam_applications_gdspa.FieldHallTicketNumber)).
		First(ctx)

	fmt.Println(lastHallTicketNumber, "last ")

	if err != nil && !ent.IsNotFound(err) {
		return "", err
	}

	if lastHallTicketNumber == nil {
		lastHallTicketNumber = &ent.Exam_Applications_GDSPA{HallTicketNumber: "100000000"} // Assuming ExamApplicationsIP struct
	}

	if lastHallTicketNumber.HallTicketNumber != "" {
		lastFourDigitsStr := lastHallTicketNumber.HallTicketNumber[len(lastHallTicketNumber.HallTicketNumber)-4:]
		lastFourDigits, err := strconv.Atoi(lastFourDigitsStr)
		if err != nil {
			return "", err
		}
		lastFourDigitsMap[lastFourDigits] = true
	}

	// Retrieve all eligible applications
	applications, err := client.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.ExamYearEQ(year),
			//exam_applications_gdspa.ExamCodeEQ(examCode),
			exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			exam_applications_gdspa.GenerateHallTicketFlagEQ(true),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.Or(
				exam_applications_gdspa.HallTicketNumberEQ(""),
			),
		).
		Order(ent.Asc(exam_applications_gdspa.FieldTempHallTicket)).
		All(ctx)

	if err != nil {
		return "", err
	}

	// If no data, set the start number to 1, else set it to the maximum number found + 1
	startNumber := 1
	if len(lastFourDigitsMap) > 0 {
		for lastFourDigits := range lastFourDigitsMap {
			startNumber = lastFourDigits + 1
		}
	}

	currentTime := time.Now().Truncate(time.Second)
	// Generate hall tickets
	var successCount int
	for _, application := range applications {
		hallTicketNumber := fmt.Sprintf("%s%04d", application.TempHallTicket, startNumber)
		_, err := application.Update().
			SetHallTicketNumber(hallTicketNumber).
			//SetGenerateHallTicketFlagByNO(true).
			SetHallTicketGeneratedFlag(true).
			SetHallTicketGeneratedDate(currentTime).
			Save(ctx)
		if err != nil {
			return "", err
		}
		startNumber++
		successCount++
	}

	// Return success message
	return fmt.Sprintf("Generated hall tickets successfully for %d eligible candidates", successCount), nil
}

func UpdateExamCentresGDSPAExams(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_GDSPA) ([]string, error) {
	var updatedRecords []string

	for _, newappl := range newappls {
		applications, err := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ReportingOfficeIDEQ(newappl.ReportingOfficeID),
				exam_applications_gdspa.GenerateHallTicketFlag(false),
				exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Order(ent.Desc(exam_applications_gdspa.FieldID)).
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

func UpdateExamCentresGDSPAExamsreturnarray(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_GDSPA) ([]UpdateResult, error) {
	var updateResults []UpdateResult

	for _, newappl := range newappls {
		// Check if ReportingOfficeID exists in the table
		exists, err := client.Exam_Applications_GDSPA.
			Query().
			Where(exam_applications_gdspa.ReportingOfficeIDEQ(newappl.ReportingOfficeID)).
			Exist(ctx)
		if err != nil {
			log.Printf("Failed to check ReportingOfficeID existence: %v\n", err)
			return nil, fmt.Errorf("failed to check ReportingOfficeID existence: %v", err)
		}
		if !exists {
			log.Printf("The ReportingOfficeID %s does not exist in the Applications table. Skipping to the next value in the loop.\n", newappl.ReportingOfficeID)
			updateResult := UpdateResult{
				ReportingOfficeID: newappl.ReportingOfficeID,
				CenterCode:        newappl.ExamCityCenterCode,
				Message:           "Reporting Office ID Does not exist and skipped",
				//RecordCount:       0,
			}
			updateResults = append(updateResults, updateResult)
			continue
		}

		// Check if CenterCode exists in the center.master table
		centerExists, err := client.Center.
			Query().
			Where(center.IDEQ(newappl.ExamCityCenterCode)).
			Exist(ctx)
		if err != nil {
			log.Printf("Failed to check CenterCode existence: %v\n", err)
			return nil, fmt.Errorf("failed to check CenterCode existence: %v", err)
		}
		if !centerExists {
			log.Printf("The CenterCode %d does not exist in the center.master table. Skipping to the next value in the loop.\n", newappl.ExamCityCenterCode)
			updateResult := UpdateResult{
				ReportingOfficeID: newappl.ReportingOfficeID,
				CenterCode:        newappl.ExamCityCenterCode,
				Message:           "CenterCode does not exist and skipped",
				//RecordCount:       0,
			}
			updateResults = append(updateResults, updateResult)
			continue
		}

		applications, err := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ReportingOfficeIDEQ(newappl.ReportingOfficeID),
				//exam_applications_gdspa.GenerateHallTicketFlag(true),
				//exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Order(ent.Desc(exam_applications_gdspa.FieldID)).
			All(ctx)
		if err != nil {
			log.Printf("Failed to query applications: %v\n", err)
			return nil, fmt.Errorf("failed to query applications: %v", err)
		}

		count := len(applications)
		if count > 0 {
			// Update the CenterCode for each application record
			for _, application := range applications {
				application.ExamCityCenterCode = newappl.ExamCityCenterCode
				_, err = application.Update().SetExamCityCenterCode(newappl.ExamCityCenterCode).Save(ctx)
				if err != nil {
					//log.Printf("Failed to update application with CenterCode %d: %v\n", newappl.CenterCode, err)
					return nil, fmt.Errorf("failed to update application: %v", err)
				}
				//log.Printf("Application with CenterCode %d updated successfully\n", application.CenterCode)
				// You can access the updated application values using the 'updatedApplication' variable
				// For example: updatedApplication.EmployeeID, updatedApplication.ApplicationNumber, etc.
			}

			updateResult := UpdateResult{
				ReportingOfficeID: newappl.ReportingOfficeID,
				CenterCode:        newappl.ExamCityCenterCode,
				Message:           "Updated Successfully",
				//RecordCount:       count,
			}
			updateResults = append(updateResults, updateResult)
		}
	}

	return updateResults, nil
}
