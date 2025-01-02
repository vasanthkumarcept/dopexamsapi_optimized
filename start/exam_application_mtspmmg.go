package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/center"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/recommendationsmtspmmgapplications"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"
	"strconv"
	"time"
)

func CreateMTSPMMGApplications(client *ent.Client, newAppln *ca_reg.ApplicationGDSPM) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateTempHallTicket(newAppln.TempHallTicket, newAppln.EmployeeID); err != nil {
		return nil, 422, " -STR100", false, err
	}
	if newAppln == nil {
		return nil, 422, " -STR001", false, errors.New("employee id is missing")
	}
	// Begin a transaction
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

	applicationNumber, err := util.GenerateApplicationNumber(client, newAppln.EmployeeID, newAppln.ExamYear, "MTSPM")
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	createdAppln, status, stgError, err := saveApplication(tx, newAppln, applicationNumber, newAppln.ExamCode, ctx)
	if err != nil {
		return nil, 500 + status, " -STR007 " + stgError, false, err
	}
	return createdAppln, 200, "", true, nil
}

func QueryMTSPMMGPMExamApplicationsByEmpID(ctx context.Context, client *ent.Client, empid int64, id1 string) (*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	// Add query hook for logging queries

	// Query for the Exam Application by Employee ID
	newAppln, err := client.Exam_Application_MTSPMMG.
		Query().
		Where(
			(exam_application_mtspmmg.EmployeeIDEQ(empid)),
			(exam_application_mtspmmg.ExamYearEQ(id1)),
			(exam_application_mtspmmg.StatusEQ("active")),
		).
		WithCirclePrefRefMTSPMMG().
		WithMTSPMMGApplicationsRef().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, errors.New("no application exists ")
		} else {
			return nil, 500, " -STR002", false, err
		}
	}

	// Extract only the desired fields from the CirclePrefRefMTSPMMG edge
	var circlePrefs []*ent.Division_Choice_MTSPMMG
	for _, edge := range newAppln.Edges.CirclePrefRefMTSPMMG {
		circlePrefs = append(circlePrefs, &ent.Division_Choice_MTSPMMG{
			Group:            edge.Group,
			CadrePrefNo:      edge.CadrePrefNo,
			Cadre:            edge.Cadre,
			PostPrefNo:       edge.PostPrefNo,
			PostingPrefValue: edge.PostingPrefValue,
		})
	}

	// Update the CirclePrefRefMTSPMMG with the extracted fields
	newAppln.Edges.CirclePrefRefMTSPMMG = circlePrefs
	var recomondPref []*ent.RecommendationsMTSPMMGApplications
	for _, edge := range newAppln.Edges.MTSPMMGApplicationsRef {
		recomondPref = append(recomondPref, &ent.RecommendationsMTSPMMGApplications{
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
	newAppln.Edges.MTSPMMGApplicationsRef = recomondPref

	return newAppln, 200, "", true, nil
}

// Verify Fuanctions

func UpdateApplicationRemarksMTSPMMG(client *ent.Client, newAppln *ca_reg.VerifyApplicationMTSPM, nonQualifyService string) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	if newAppln == nil {
		return nil, 400, " -STR001", false, errors.New("payload received in empty")
	}
	// Start a transaction.
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	defer handleTransaction(tx, &err)

	// Check if the EmployeeID exists.
		
		oldAppln, status, stgError, err := fetchExistingMtspmApplication(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear)
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
	updatedAppln, err := createUpdatemtspmApplication(ctx, tx, oldAppln, newAppln, applicationStatus, nonQualifyService)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	err = handleMtspmRecommendations(ctx, tx, updatedAppln, newAppln)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}
	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB017", false, err
	}
	return appResponse, 200, "", true, nil
}

func ResubmitApplicationRemarksMTSPMMG(client *ent.Client, newAppln *ca_reg.ReApplicationGDSPM) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateTempHallTicket(newAppln.TempHallTicket, newAppln.EmployeeID); err != nil {
		return nil, 422, " -STR001", false, fmt.Errorf("issue for employee %d with temp hall ticket number: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
	}
	if newAppln == nil {
		return nil, 400, " -STR001", false, errors.New("payload is empty")
	}
	//transaction implementation--------------
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	defer func() {
		handleTransaction(tx, &err)
	}()

	// Check if the EmployeeID exists.
	oldAppln, status, stgError, err := fetchOldApplicationMtsPm(ctx, tx, newAppln)
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

// CA PEMNDING

func QueryMTSPMMGApplicationsByCAVerificationsPending(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	// Array of exams

	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID and Examyear cannot be null")
	}

	records, err := client.Exam_Application_MTSPMMG.Query().
		Where(
			exam_application_mtspmmg.And(
				exam_application_mtspmmg.Or(
					exam_application_mtspmmg.ApplicationStatusEQ("CAVerificationPending"),
					exam_application_mtspmmg.ApplicationStatusEQ("ResubmitCAVerificationPending"),
				),
				exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(facilityID),
				exam_application_mtspmmg.ExamYearEQ(id1),
				exam_application_mtspmmg.StatusEQ("active"),
			),
		).
		Order(ent.Desc(exam_application_mtspmmg.FieldID)). // Order by descending updated_at timestamp
		WithCirclePrefRefMTSPMMG().                        // Add the Where clause with multpmpale statuses using Or
		All(ctx)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR007", false, fmt.Errorf("no Applications is pending for CA pending verification for the Office ID %s", facilityID)
	}

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

	}

	//return records, nil

	for _, record1 := range records {
		// Extract only the desired fields from the CirclePrefRefMTSPMMG edge
		var recommendationsPrefs []*ent.RecommendationsMTSPMMGApplications
		for _, edge := range record1.Edges.MTSPMMGApplicationsRef {
			recommendationsPrefs = append(recommendationsPrefs, &ent.RecommendationsMTSPMMGApplications{

				VacancyYear:       edge.VacancyYear,
				CARecommendations: edge.CARecommendations,
				CAUserName:        edge.CAUserName,
				CARemarks:         edge.CARemarks,
				NORecommendations: edge.NORecommendations,
				NORemarks:         edge.NORemarks,
				NOUserName:        edge.NOUserName,
			})
		}
		// Update the CirclePrefRefMTSPMMG with the extracted fields
		record1.Edges.MTSPMMGApplicationsRef = recommendationsPrefs

		// Log the updated CirclePrefRefMTSPMMG for each record

	}

	return records, 200, "", true, nil
}

// CA PENDING BY EMP ID
func QueryPMTSPMMGApplicationsByCAPendingByEmpID(ctx context.Context, client *ent.Client, empID int64, examYear string) (*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {

	// Retrieve the latest record based on UpdatedAt timestamp
	record, err := client.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.EmployeeIDEQ(empID),
			exam_application_mtspmmg.ExamYearEQ(examYear),
			exam_application_mtspmmg.StatusEQ("active"),
			exam_application_mtspmmg.Or(
				exam_application_mtspmmg.ApplicationStatusEQ("CAVerificationPending"),
				exam_application_mtspmmg.ApplicationStatusEQ("ResubmitCAVerificationPending"),
			),
		).
		WithMTSPMMGApplicationsRef().
		WithCirclePrefRefMTSPMMG().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, errors.New("no  ip application pending for CA verification ")
		} else {
			return nil, 500, " -STR002", false, err
		}
	}

	return record, 200, "", true, nil
}

// all CA VERIFIED

func QueryMTSPMMGApplicationsByCAVerified(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Application_MTSPMMG.Query().
		Where(
			exam_application_mtspmmg.And(
				exam_application_mtspmmg.ApplicationStatusEQ("VerifiedByCA"),
				exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(id1),
			),
		).
		WithCirclePrefRefMTSPMMG().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application pending for  CA verification for the Office ID %s", facilityID)
	}
	//log.Println("CA verified records returned: ", records)
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

	}
	return records, 200, "", true, nil
}
func QueryMTSPMApplicationsByCAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID, facilityID1 string) ([]*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	if facilityID == "" || facilityID1 == "" {
		return nil, 422, " -STR001", false, fmt.Errorf("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.And(
				exam_application_mtspmmg.ApplicationStatusEQ("VerifiedByCA"),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(facilityID1),
			),
		).
		WithMTSPMMGApplicationsRef().
		WithCirclePrefRefMTSPMMG().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application verified by CA/NA for the Office ID %s", facilityID)
	}
	return records, 200, "", true, nil
}

// Get CA Verified with Emp ID
func QueryMTSPMMGApplicationsByCAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {

	record, err := client.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.ApplicationStatusEQ("VerifiedByCA"), // Check for "CAVerified" status
			exam_application_mtspmmg.EmployeeIDEQ(employeeID),
			exam_application_mtspmmg.StatusEQ("active"),
			exam_application_mtspmmg.ExamYearEQ(examYear),
		).
		WithMTSPMMGApplicationsRef().
		WithCirclePrefRefMTSPMMG().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf("no application found for this employee ID: %d with 'CAVerified' status", employeeID)
		}
		return nil, 500, " -STR002", false, err
	}

	return record, 200, "", true, nil
}

// Get latest old Application Remarks given to Candidate for CA Verification
func GetOldMTSPMMGCAApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {

	application, err := client.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.EmployeeIDEQ(employeeID),
			exam_application_mtspmmg.ApplicationStatusEQ("PendingWithCandidate"),
			exam_application_mtspmmg.ExamYearEQ(examYear),
			exam_application_mtspmmg.StatusEQ("active"),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf("application not found for employee ID: %d with 'PendingWithCandidate' status", employeeID)
		}
		return nil, 500, " -STR002", false, err
	}

	return application, 200, "", true, nil
}

// UpdateNodalRecommendationsByEmpID updates the recommendations for a given employee ID
func UpdateMTSPMMGNodalRecommendationsByEmpID(client *ent.Client, applicationRecord *ca_reg.NAVerifyApplicationMTSPM) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateMtspmInput(applicationRecord); err != nil {
		return nil, 422, " -STR001", false, errors.New("employee id should not be empty")
	}
	empID := applicationRecord.EmployeeID
	id1 := applicationRecord.ExamYear


	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	defer handleTransaction(tx, &err)


	// Check if empID exists in exam_applications_pmpa and the status is "VerifiedByCA" or "VerifiedByNA"
	exists, status, stgError, err := checkMtspmApplicationExists(tx, ctx, applicationRecord)

	if status == 500 {
		return nil, 500, " -STR004 " + stgError, false, err
	}

	if status == 422 {
		return nil, 422, " -STR005 " + stgError, false, fmt.Errorf("no active application found for this candidate")
	}

	if !exists {
		return nil, 422, " -STR007 ", false, fmt.Errorf("no active application found for this candidate")
	}

	// Retrieve all records for the employee ID from RecommendationsPMPAApplications
	records, err := getMtspmRecommendationsByEmpID(ctx, tx, empID)
	if err != nil {
		return nil, 500, " -STR008 ", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR009 ", false, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}

	//currentTime := time.Now().Truncate(time.Second)
	// Format the current time to "yyyymmddhhmmss"
	stat := "inactive_" + time.Now().Format("20060102150405")

	// Retrieve the updated record for the employee ID with the highest fieldID

	// Update the retrieved record with the provided values

	updatedRecord, status, stgError, err := getActiveExamApplicationMtspm(ctx, tx, empID, id1)
	if status == 500 {
		return nil, 500 + status, " -STR010 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR011 " + stgError, false, fmt.Errorf("no active application exists")

	}
	if err != nil {
		return nil, 500 + status, " -STR012 " + stgError, false, err
	}

	// Update the retrieved record with the provided values
	_, err = updatedRecord.
		Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return nil, 422, " -STR0010", false, fmt.Errorf("failed to update application: %v", err)
	}
	// Hall Ticket Generated Flag
	hallticketgeneratedflag := checkMtspmHallTicketGenerated(applicationRecord, updatedRecord)

	updatedAppln, err := createUpdatedMtspmAppln(tx, applicationRecord, updatedRecord, hallticketgeneratedflag, ctx)
	if err != nil {
		return nil, 500, " -STR014 ", false, err
	}
	// Save the Recommendation records.
	recommendationsRef, err := createMtspmRecommendationsRef(ctx, tx, applicationRecord, updatedAppln)
	if err != nil {
		return nil, 500, " -STR015", false, err
	}

	updatedAppln.Update().
		//ClearIPApplicationsRef().
		AddMTSPMMGApplicationsRef(recommendationsRef...).
		Save(ctx)
		if err != nil {
			return nil, 500, " -STR016 ", false, err
		}
		appResponse, err := MapExamApplicationsToResponse(updatedAppln)
		if err != nil {
			return nil, 500, " -SUB017 ", false, err
		}
		return appResponse, 200, "", true, nil
}


///MTS PM MG RECOMMENDATIONS

func GetMTSPMMGRecommendationsByEmpID(client *ent.Client, empID int64) ([]*ent.RecommendationsMTSPMMGApplications, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	// Check if empID is null
	if empID == 0 {
		return nil, 422, " -STR001", false, fmt.Errorf("no employee ID provided to process")
	}

	// Retrieve all records for the employee ID
	records, err := client.RecommendationsMTSPMMGApplications.Query().
		Where(recommendationsmtspmmgapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no recommendations found for this employee ID: %d", empID)
	}

	return records, 200, "", true, nil
}

func QueryMTSPMMGApplicationsByNAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID, facilityID1 string) ([]*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || facilityID1 == "" {
		return nil, 422, " -STR001", false, errors.New(" facility ID and ExamYear cannot be null")
	}
	records, err := client.Exam_Application_MTSPMMG.Query().
		Where(
			exam_application_mtspmmg.And(
				exam_application_mtspmmg.ApplicationStatusEQ("VerifiedByNA"),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(facilityID),
				exam_application_mtspmmg.ExamYearEQ(facilityID1),
				exam_application_mtspmmg.StatusEQ("active"),
			),
		).
		WithCirclePrefRefMTSPMMG().
		WithMTSPMMGApplicationsRef().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application pending with NA for facility Id %s", facilityID)
	}
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

	}
	return records, 200, "", true, nil
}

// Get Admit card details .
func GetMTSPMMGApplicationsWithHallTicket(client *ent.Client, examCode int32, employeeID int64, examYear string) (*ent.Exam_Application_MTSPMMG, *ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Check if exam code is valid
	if examCode == 0 {
		return nil, nil, 422, " -STR001", false, errors.New("please provide a valid exam code")
	}
	if examCode == 5 {

		// Query the Exam_Applications_PMPA table to retrieve the applicant details
		application, err := client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.EmployeeIDEQ(employeeID),
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.HallTicketNumberNEQ(""),
				exam_application_mtspmmg.ExamCityCenterCodeNEQ(0),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
			).
			//WithMTSPMMGExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, 422, " -STR002", false, errors.New("no admit card available for this application ")
			} else {
				return nil, nil, 500, " -STR003", false, err
			}
		}

		// Fetch the associated RecommendationsPMPA records matching the employee ID
		recommendations, err := client.RecommendationsMTSPMMGApplications.Query().
			Where(recommendationsmtspmmgapplications.EmployeeIDEQ(employeeID)).
			All(ctx)
		if err != nil {
			return nil, nil, 500, " -STR004", false, err
		}

		// Assign the fetched recommendations to the application entity
		application.Edges.MTSPMMGApplicationsRef = recommendations

		return application, nil, 200, "", true, nil

	}
	return nil, nil, 422, " -STR005", false, errors.New("invalid exam code")
}

// // Get All Pending with Candidate
// Assuming Exam_Applications_PMPA has a field named EmployeeID, you might adapt the code like this:
func QueryMTSPMMGApplicationsByPendingWithCandidate(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and Exam Year cannot be empty")
	}

	// Fetch all applications matching the criteria
	records, err := client.Exam_Application_MTSPMMG.Query().
		Where(
			exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(facilityID),
			exam_application_mtspmmg.ExamYearEQ(id1),
			exam_application_mtspmmg.StatusEQ("active"),
		).
		Order(ent.Asc("employee_id")). /*, ent.Desc("application_number"))*/ // Order by employee_id and application_number
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application found for this CA Office ID %s", facilityID)
	}

	// Create a map to store the latest applications for each employee
	latestApplications := make(map[int64]*ent.Exam_Application_MTSPMMG)

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
	var result []*ent.Exam_Application_MTSPMMG
	for _, application := range latestApplications {
		result = append(result, application)
	}

	if len(result) == 0 {
		return nil, 422, " -STR004", false, fmt.Errorf("no Applicationsf found pending with candiadte under this Office ID %s", facilityID)
	}

	return result, 200, "", true, nil
}

// mts STRUCT
type ExamStatsMTSPMMG struct {
	CircleName                  string
	ControllingOfficeName       string
	ControllingOfficeFacilityID string
	NodalOfficeFacilityID       string
	NoOfCandidatesChosenCity    int
	NoOfCandidatesAlloted       int
}

func GetExamApplicatonsPreferenenceCityWiseStatsMTSPMMG(ctx context.Context, client *ent.Client, ExamYear string, SubExamcode string, SubCityid string) ([]ExamStatsMTSPMMG, int32, string, bool, error) {
	var result []ExamStatsMTSPMMG

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

	applications, err := client.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.ExamYearEQ(ExamYear),
			exam_application_mtspmmg.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_application_mtspmmg.HallTicketNumberNEQ(""),
			exam_application_mtspmmg.StatusEQ("active"),
			exam_application_mtspmmg.ExamCodeEQ(Examcode),
			exam_application_mtspmmg.CenterIdEQ(Cityid),
		).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR004", false, err
	} else {
		if len(applications) == 0 {
			return nil, 422, " -STR005", false, errors.New("no matching data with Exam City")
		}
	}

	groupedApplications := make(map[string]ExamStatsMTSPMMG)

	for _, examApplication := range applications {
		nodalOfficeFacilityID := examApplication.NodalOfficeFacilityID
		nodalOfficeName := examApplication.NodalOfficeName

		controllingOfficeFacilityID := examApplication.ControllingOfficeFacilityID
		controllingOfficeName := examApplication.ControllingOfficeName
		centerCode := examApplication.ExamCityCenterCode

		// Check if ExamStats for the reporting office already exists
		stats, ok := groupedApplications[controllingOfficeName]
		if !ok {
			stats = ExamStatsMTSPMMG{
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

func UpdateCenterCodeForApplicationsMTSPMMG(ctx context.Context, client *ent.Client, controllingOfficeFacilityID string, examCenterID, seatsToAllot, examCityID int32) (int, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	strExamCenterID := strconv.FormatInt(int64(examCenterID), 10)
	if controllingOfficeFacilityID == "" {
		return 0, nil, 422, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR001", false, errors.New("controlling Office Facility ID cannot be nil")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return 0, nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	applications, err := tx.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.ExamCityCenterCodeIsNil(),
			exam_application_mtspmmg.CenterIdEQ(examCityID),
			exam_application_mtspmmg.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(controllingOfficeFacilityID),
			exam_application_mtspmmg.HallTicketNumberNEQ(""),
			exam_application_mtspmmg.StatusEQ("active"),
		).
		Order(ent.Asc(exam_application_mtspmmg.FieldApplnSubmittedDate)).
		Limit(int(seatsToAllot)). // Limit the number of records to be updated
		All(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR002", false, err
	}

	var updatedCount int
	for _, application := range applications {
		_, err := tx.Exam_Application_MTSPMMG.
			UpdateOne(application).
			SetExamCityCenterCode(examCenterID).
			Save(ctx)
		if err != nil {
			return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR003", false, err
		}
		updatedCount++
	}
	updateCount, err := tx.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.ExamCityCenterCodeEQ(examCenterID),
			exam_application_mtspmmg.StatusEQ("active"),
		).
		Count(ctx)
	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR004", false, err
	}

	// Updating Center Table
	centerDet, err := tx.Center.
		Query().
		Where(center.IDEQ(examCenterID)).
		Only(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR005", false, err
	}

	maxSeats := centerDet.MaxSeats
	pendingSeats := int32(maxSeats) - int32(updateCount)

	_, err = tx.Center.
		UpdateOne(centerDet).
		SetNoAlloted(int32(updateCount)).
		SetPendingSeats(pendingSeats).
		Save(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR006", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return 0, nil, 500, " -STR007", false, err
	}
	// Success
	return updatedCount, applications, 200, "", true, nil
}
func GenerateHallticketNumberrMtsPm(ctx context.Context, client *ent.Client, year string, examCode int32, nodalOfficerFacilityID string) (string, int32, string, bool, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return "", 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	// Retrieve the last hall ticket number and extract its last four digits
	lastFourDigitsMap := make(map[int]bool)
	lastHallTicketNumber, err := tx.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.ExamYearEQ(year),
			exam_application_mtspmmg.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			exam_application_mtspmmg.GenerateHallTicketFlagEQ(true),
			exam_application_mtspmmg.HallTicketNumberNEQ(""),
			exam_application_mtspmmg.StatusEQ("active"),
		).
		Order(ent.Desc(exam_application_mtspmmg.FieldHallTicketNumber)).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			lastHallTicketNumber = nil
		} else {
			return "", 500, " -STR001", false, err
		}
	}

	// Initialize lastHallTicketNumber if not found
	if lastHallTicketNumber == nil {
		lastHallTicketNumber = &ent.Exam_Application_MTSPMMG{HallTicketNumber: "100000000"} // Assuming ExamApplicationsIP struct
	}

	if lastHallTicketNumber.HallTicketNumber != "" {
		lastFourDigitsStr := lastHallTicketNumber.HallTicketNumber[len(lastHallTicketNumber.HallTicketNumber)-4:]
		lastFourDigits, err := strconv.Atoi(lastFourDigitsStr)
		if err != nil {
			return "", 400, " -STR002", false, errors.New("unable to get last hall ticket number")
		}
		lastFourDigitsMap[lastFourDigits] = true
	}

	// Retrieve all eligible applications
	applications, err := tx.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.ExamYearEQ(year),
			exam_application_mtspmmg.ExamCodeEQ(examCode),
			exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			exam_application_mtspmmg.GenerateHallTicketFlagEQ(true),
			exam_application_mtspmmg.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_application_mtspmmg.StatusEQ("active"),
			exam_application_mtspmmg.Or(
				exam_application_mtspmmg.HallTicketNumberEQ(""),
			),
		).
		Order(ent.Asc(exam_application_mtspmmg.FieldTempHallTicket)).
		All(ctx)
	if err != nil {
		return "", 500, " -STR003", false, err
	} else {
		if len(applications) == 0 {
			return "", 422, " -STR004", false, errors.New("no application pending for hallticket generation")
		}
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
			SetHallTicketGeneratedFlag(true).
			SetHallTicketGeneratedDate(currentTime).
			Save(ctx)
		if err != nil {
			return "", 500, " -STR005", false, err
		}
		startNumber++
		successCount++
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return "", 500, " -STR008", false, err
	}
	// Return success message
	return fmt.Sprintf("Generated hall tickets successfully for %d eligible candidates", successCount), 200, "", true, nil
}

// Get All NA Verified Records
func QueryMTSPMMGApplicationsByNAVerified(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New(" facility ID and ExamYear cannot be null")
	}
	records, err := client.Exam_Application_MTSPMMG.Query().
		Where(
			exam_application_mtspmmg.And(
				exam_application_mtspmmg.ApplicationStatusEQ("VerifiedByNA"),
				exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(facilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(id1),
			),
		).
		WithCirclePrefRefMTSPMMG().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application pending with NA for facility Id %s", facilityID)
	}

	return records, 200, "", true, nil
}

func GetEligibleMTSPMApplicationsForCircleDetails(ctx context.Context, client *ent.Client, examCode int32, circleOfficeID string, Examyear string) ([]map[string]interface{}, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return nil, fmt.Errorf("no such valid exam code exists")
	}

	if circleOfficeID == "" {
		return nil, fmt.Errorf("please provide Nodal Officer's office ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_PS table
	count, err := client.Exam_Application_MTSPMMG.
		Query().
		Where(exam_application_mtspmmg.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_application_mtspmmg.StatusEQ("active"),
			exam_application_mtspmmg.ExamYearEQ(Examyear),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_PMPA: %v", err)
		return nil, fmt.Errorf("failed to retrieve applications from Exam_Applications_PMPA: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", circleOfficeID)
		return nil, fmt.Errorf("no valid applications available for the circle")
	}

	// Query to get the applications matching the circleOfficeID
	applications, err := client.Exam_Application_MTSPMMG.
		Query().
		Where(exam_application_mtspmmg.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_application_mtspmmg.StatusEQ("active"),
			exam_application_mtspmmg.ExamYearEQ(Examyear),
		).
		Order(ent.Asc(exam_application_mtspmmg.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications: %v", err)
		return nil, fmt.Errorf("failed to retrieve applications: %v", err)
	}

	uniqueEmployees := make(map[int64]struct{})
	employeeLatestApplication := make(map[int64]*ent.Exam_Application_MTSPMMG) // Map to store the latest application for each employee
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
