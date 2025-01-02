package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	"recruit/ent"
	"strconv"

	"recruit/ent/center"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/exam_applications_gdspm"

	"recruit/ent/recommendationsgdspmapplications"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"
	"time"
)

// Create Applications with Circle Preferences ...
func CreateGDSPMApplicationss(client *ent.Client, newAppln *ca_reg.ApplicationGDSPM) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateTempHallTicket(newAppln.TempHallTicket, newAppln.EmployeeID); err != nil {
		return nil, 422, " -STR100", false, err
	}
	if newAppln == nil {
		return nil, 422, " -STR001", false, errors.New("payload cannot be blank")
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
	// Generate Application number
	// Generate application number in the format "GDSPM2023XXXXXX"
	applicationNumber, err := util.GenerateApplicationNumber(client, newAppln.EmployeeID, newAppln.ExamYear, "GDSPM")
	if err != nil {
		return nil, 422, " -STR006", false, fmt.Errorf(" failed to generate application number: %v", err)
	}
	createdAppln, status, stgError, err := saveApplication(tx, newAppln, applicationNumber, newAppln.ExamCode, ctx)
	if err != nil {
		return nil, 500 + status, " -STR007 " + stgError, false, err
	}

	return createdAppln, 200, "", true, nil
}

// Query GDSPMExam Application with Emp ID.
func QueryGDSPMExamApplicationsByEmpID(ctx context.Context, client *ent.Client, empid int64, id1 string) (*ent.Exam_Applications_GDSPM, int32, string, bool, error) {

	newAppln, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			(exam_applications_gdspm.EmployeeIDEQ(empid)),
			(exam_applications_gdspm.ExamYearEQ(id1)),
			(exam_applications_gdspm.StatusEQ("active")),
		).
		WithCirclePrefRefGDSPM().
		WithGDSPMApplicationsRef().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, errors.New("no application exists ")
		} else {
			return nil, 500, " -STR002", false, err
		}
	}

	// Extract only the desired fields from the CirclePrefRefGDSPM edge
	var circlePrefs []*ent.Division_Choice_PM
	for _, edge := range newAppln.Edges.CirclePrefRefGDSPM {
		circlePrefs = append(circlePrefs, &ent.Division_Choice_PM{
			Group:            edge.Group,
			CadrePrefNo:      edge.CadrePrefNo,
			Cadre:            edge.Cadre,
			PostPrefNo:       edge.PostPrefNo,
			PostingPrefValue: edge.PostingPrefValue,
		})
	}

	// Update the CirclePrefRefGDSPM edge with the filtered values
	newAppln.Edges.CirclePrefRefGDSPM = circlePrefs

	// Extract only the desired fields from the CadrePrefRefGDSPM edge
	var cadrePrefs []*ent.Cadre_Choice_PM
	for _, edge := range newAppln.Edges.CadrePrefRefGDSPM {
		cadrePrefs = append(cadrePrefs, &ent.Cadre_Choice_PM{
			Group:         edge.Group,
			CadrePrefNo:   edge.CadrePrefNo,
			Cadre:         edge.Cadre,
			PostPrefNo:    edge.PostPrefNo,
			PostPrefValue: edge.PostPrefValue,
		})
	}

	// Update the CadrePrefRefGDSPM edge with the filtered values
	newAppln.Edges.CadrePrefRefGDSPM = cadrePrefs

	var recomondPref []*ent.RecommendationsGDSPMApplications
	for _, edge := range newAppln.Edges.GDSPMApplicationsRef {
		recomondPref = append(recomondPref, &ent.RecommendationsGDSPMApplications{
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
	newAppln.Edges.GDSPMApplicationsRef = recomondPref
	newAppln.UpdatedAt = newAppln.UpdatedAt.Truncate(24 * time.Hour)

	return newAppln, 200, "", true, nil
}

func UpdateApplicationRemarkssGDSPM(client *ent.Client, newAppln *ca_reg.VerifyApplicationGDSPM, nonQualifyService string) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
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

	// Defer rollback in case anything fails.
	defer handleTransaction(tx, &err)

	// Check if the EmployeeID exists.
	oldAppln, status, stgError, err := fetchExistingGdspmApplication(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear)
	if status == 500 {
		return nil, 500 + status, " -STR003 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR004 " + stgError, false, fmt.Errorf("no active application found for this candidate")

	}
	if err != nil {
		return nil, 500 + status, " -STR005 " + stgError, false, err
	}
	// Format the current time to "yyyymmddhhmmss"
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
	updatedAppln, err := createUpdategdspmApplication(ctx, tx, oldAppln, newAppln, applicationStatus, nonQualifyService)
	if err != nil {
		return nil, 500, " -STR009", false, err
	}
	err = handleGDSpmRecommendations(ctx, tx, updatedAppln, newAppln)
	if err != nil {
		return nil, 500, " -STR010", false, err
	}

	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB017", false, err
	}
	return appResponse, 200, "", true, nil
}

func ResubmitApplicationRemarkssGDSPM(client *ent.Client, newAppln *ca_reg.ReApplicationGDSPM) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateTempHallTicket(newAppln.TempHallTicket, newAppln.EmployeeID); err != nil {
		return nil, 422, " -STR001", false, fmt.Errorf("issue for employee %d with temp hall ticket number: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
	}
	// Check if newAppln is not nil.
	if newAppln == nil {
		return nil, 400, " -STR001", false, errors.New("payload is empty")
	}
	// Begin a transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	defer func() {
		handleTransaction(tx, &err)
	}()

	// Check if the EmployeeID exists.
	oldAppln, status, stgError, err := fetchOldApplicationGdsPm(ctx, tx, newAppln)
	if status == 500 {
		return nil, 500 + status, " -STR004 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR005 " + stgError, false, fmt.Errorf("no active application found for this candidate")
	}

	if err != nil {
		return nil, 500 + status, " -STR005 " + stgError, false, err
	}
	statuses := []string{ "ResubmitCAVerificationPending"}
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
		return nil, 500 + status, " -STR008 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR009 " + stgError, false, err

	}
	if err != nil {
		return nil, 500, " -STR010 ", false, err
	}
	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB011 ", false, err
	}

	return appResponse, 200, "", true, nil
}

// UpdateNodalRecommendationsByEmpID updates the recommendations for a given employee ID
func UpdateGDSPMNodalRecommendationsByEmpID(client *ent.Client, applicationRecord *ca_reg.NAVerifyApplicationGDSPM) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateGdspmInput(applicationRecord); err != nil {
		return nil, 422, " -STR001", false, errors.New("employee id should not be empty")
	}
	empID := applicationRecord.EmployeeID
	id1 := applicationRecord.ExamYear
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR013", false, err
	}
	defer handleTransaction(tx, &err)


	// Check if empID exists in exam_applications_gdspm and the status is "VerifiedByCA" or "VerifiedByNA"
	exists, status, stgError, err := checkGdspmApplicationExists(tx, ctx, applicationRecord)

	if status == 500 {
		return nil, 500, " -STR004 " + stgError, false, err
	}

	if status == 422 {
		return nil, 422, " -STR005 " + stgError, false, fmt.Errorf("no active application found for this candidate")
	}

	if err != nil {
		return nil, 500, " -STR006 " + stgError, false, err
	}

	if !exists {
		return nil, 422, " -STR007 ", false, fmt.Errorf("no active application found for this candidate")
	}

	// Retrieve all records for the employee ID from RecommendationsGDSPMApplications

	records, err := getGdspmRecommendationsByEmpID(ctx, tx, empID)
	if err != nil {
		return nil, 500, " -STR008 ", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR009 ", false, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}

	//currentTime := time.Now().Truncate(time.Second)

	// Format the current time to "yyyymmddhhmmss"
	stat := "inactive_" + time.Now().Format("20060102150405")

	// Update the retrieved record with the provided values
	updatedRecord, status, stgError, err := getActiveExamApplicationGdspm(ctx, tx, empID, id1)
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
		return nil, 422, " -STR009", false, fmt.Errorf(" failed to update application: %v", err)
	}
	// Hall Ticket Generated Flag
	hallticketgeneratedflag := checkGdspmHallTicketGenerated(applicationRecord, updatedRecord)
	updatedAppln, err := createUpdatedGdspmAppln(tx, applicationRecord, updatedRecord, hallticketgeneratedflag, ctx)
	if err != nil {
		return nil, 500, " -STR014 ", false, err
	}

	// Save the Recommendation records.
	recommendationsRef, err := createGdspmRecommendationsRef(ctx, tx, applicationRecord, updatedAppln)
	if err != nil {
		return nil, 500, " -STR015", false, err
	}

	updatedAppln.Update().
		//ClearIPApplicationsRef().
		AddGDSPMApplicationsRef(recommendationsRef...).
		Save(ctx)
		appResponse, err := MapExamApplicationsToResponse(updatedAppln)
		if err != nil {
			return nil, 500, " -SUB017 ", false, err
		}
		return appResponse, 200, "", true, nil
}

// Get All CA Pending records ...
func QueryGDSPMApplicationsByCAVerificationsPending(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID and Examyear cannot be null")
	}

	records, err := client.Exam_Applications_GDSPM.Query().
		Where(
			exam_applications_gdspm.And(
				exam_applications_gdspm.Or(
					exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				),
				exam_applications_gdspm.ControllingOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspm.ExamYearEQ(id1),
				exam_applications_gdspm.StatusEQ("active"),
			),
		).
		Order(ent.Desc(exam_applications_gdspm.FieldID)). // Order by descending updated_at timestamp
		WithCirclePrefRefGDSPM().                         // Add the Where clause with multgdspmle statuses using Or
		All(ctx)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR006", false, fmt.Errorf("no Applications is pending for CA pending verification for the Office ID %s", facilityID)
	}
	return records, 200, "", true, nil
}

func QueryGDSPMApplicationsByVAVerificationsPending(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID and Examyear cannot be null")
	}
	records, err := client.Exam_Applications_GDSPM.Query().
		Where(
			exam_applications_gdspm.And(
				exam_applications_gdspm.Or(
					exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				),
				exam_applications_gdspm.SubdivisionOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspm.ExamYearEQ(id1),
				exam_applications_gdspm.StatusEQ("active"),
			),
		).
		Order(ent.Desc(exam_applications_gdspm.FieldID)). // Order by descending updated_at timestamp
		WithCirclePrefRefGDSPM().                         // Add the Where clause with multgdspmle statuses using Or
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, errors.New("no pending application found for this Sub Division")
	}
	return records, 200, "", true, nil
}

// Get All CA verified records
func QueryGDSPMApplicationsByCAVerified(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_GDSPM.Query().
		Where(
			exam_applications_gdspm.And(
				exam_applications_gdspm.ApplicationStatusEQ("VerifiedByCA"),
				exam_applications_gdspm.ControllingOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(id1),
			),
		).
		WithCirclePrefRefGDSPM().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application pending for  CA verification for the Office ID %s", facilityID)
	}
	return records, 200, "", true, nil
}

// Get CA Verified with Emp ID
func QueryGDSPMApplicationsByCAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_GDSPM, int32, string, bool, error) {

	record, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.ApplicationStatusEQ("VerifiedByCA"), // Check for "CAVerified" status
			exam_applications_gdspm.EmployeeIDEQ(employeeID),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		WithGDSPMApplicationsRef().
		WithCirclePrefRefGDSPM().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf("no application found for this employee ID: %d with 'CAVerified' status", employeeID)
		}
		return nil, 500, " -STR002", false, err
	}

	return record, 200, "", true, nil
}

func QueryGDSPMApplicationsByCAPendingByEmpID(ctx context.Context, client *ent.Client, empID int64, examYear string) (*ent.Exam_Applications_GDSPM, int32, string, bool, error) {

	// Retrieve the latest record based on UpdatedAt timestamp
	record, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.EmployeeIDEQ(empID),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
			exam_applications_gdspm.Or(
				exam_applications_gdspm.ApplicationStatusEQ("CAVerificationPending"),
				exam_applications_gdspm.ApplicationStatusEQ("ResubmitCAVerificationPending"),
			),
		).
		WithGDSPMApplicationsRef().
		WithCirclePrefRefGDSPM().
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

// Get latest old Application Remarks given to Candidate for CA Verification
func GetOldGDSPMCAApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_GDSPM, int32, string, bool, error) {

	application, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.EmployeeIDEQ(employeeID),
			exam_applications_gdspm.ApplicationStatusEQ("PendingWithCandidate"),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
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

// Get Recommendations/ Remarks with Emp ID
func GetGDSPMRecommendationsByEmpID(client *ent.Client, empID int64) ([]*ent.RecommendationsGDSPMApplications, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	// Check if empID is null
	if empID == 0 {
		return nil, 422, " -STR001", false, fmt.Errorf("no employee ID provided to process")
	}

	// Retrieve all records for the employee ID
	records, err := client.RecommendationsGDSPMApplications.Query().
		Where(recommendationsgdspmapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no recommendations found for this employee ID: %d", empID)
	}

	return records, 200, "", true, nil
}

// Get All NA Verified Records
func QueryGDSPMApplicationsByNAVerified(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New(" facility ID and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_GDSPM.Query().
		Where(
			exam_applications_gdspm.And(
				exam_applications_gdspm.ApplicationStatusEQ("VerifiedByNA"),
				exam_applications_gdspm.ControllingOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(id1),
			),
		).
		WithCirclePrefRefGDSPM().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application pending with NA for facility Id %s", facilityID)
	}

	return records, 200, "", true, nil
}

// Get All NA Verified Records with Emp ID
func QueryGDSPMApplicationsByNAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_GDSPM, int32, string, bool, error) {

	record, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.ApplicationStatusEQ("VerifiedByNA"), // Check for "CAVerified" status
			exam_applications_gdspm.EmployeeIDEQ(employeeID),
			exam_applications_gdspm.StatusEQ("active"),
			exam_applications_gdspm.ExamYearEQ(examYear),
		).
		WithGDSPMApplicationsRef().
		WithCirclePrefRefGDSPM().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf("no application pending with CA Verification for  employee ID: %d ", employeeID)
		}
		return nil, 500, " -STR002", false, err
	}

	//log.Println("CA verified record returned: ", record)
	return record, 200, "", true, nil
}

func QueryGDSPMApplicationsByNAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID, facilityID1 string) ([]*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || facilityID1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_GDSPM.Query().
		Where(
			exam_applications_gdspm.And(
				exam_applications_gdspm.ApplicationStatusEQ("VerifiedByNA"),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspm.ExamYearEQ(facilityID1),
				exam_applications_gdspm.StatusEQ("active"),
			),
		).
		WithCirclePrefRefGDSPM().
		WithGDSPMApplicationsRef().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application pending with NA for facility Id %s", facilityID)
	}

	return records, 200, "", true, nil
}

// // Get All CA verified records for NA
func QueryGDSPMApplicationsByCAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID, facilityID1 string) ([]*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	if facilityID == "" || facilityID1 == "" {
		return nil, 422, " -STR001", false, fmt.Errorf(" facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_GDSPM.Query().
		Where(
			exam_applications_gdspm.And(
				exam_applications_gdspm.ApplicationStatusEQ("VerifiedByCA"),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(facilityID1),
			),
		).
		WithGDSPMApplicationsRef().
		WithCirclePrefRefGDSPM().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application verified by CA/NA for the Office ID %s", facilityID)
	}
	return records, 200, "", true, nil
}

// Get Admit card details .
func GetGDSPMApplicationsWithHallTicket(client *ent.Client, examCode int32, employeeID int64, examYear string) (*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Check if exam code is valid
	if examCode == 0 {
		return nil, 422, " -STR001", false, errors.New("please provide a valid exam code")
	}

	// Check if the employee_ID exists in the Exam_Applications_GDSPM table
	exists, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.EmployeeIDEQ(employeeID),
			exam_applications_gdspm.ExamCodeEQ(examCode),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Exist(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if !exists {
		return nil, 422, " -STR003", false, fmt.Errorf(" no applications are found for the employee in GDSPM Applications: %d", employeeID)
	}

	// Query the Exam_Applications_GDSPM table to retrieve the applicant details
	application, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.EmployeeIDEQ(employeeID),
			exam_applications_gdspm.ExamCodeEQ(examCode),
			exam_applications_gdspm.HallTicketNumberNEQ(""),
			exam_applications_gdspm.ExamCityCenterCodeNEQ(0),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		//WithGDSPMExamCentres().
		First(ctx)
	if err != nil {
		return nil, 500, " -STR004", false, err
	}

	// Fetch the associated RecommendationsGDSPM records matching the employee ID
	recommendations, err := client.RecommendationsGDSPMApplications.Query().
		Where(recommendationsgdspmapplications.EmployeeIDEQ(employeeID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	// Assign the fetched recommendations to the application entity
	application.Edges.GDSPMApplicationsRef = recommendations
	return application, 200, "", true, nil
}

// Get Circle details summary ofExam Applications for the Nodal Officer Office ID. - For GDSPM ALone
func GetEligibleGDSPMApplicationsForCircleDetails(ctx context.Context, client *ent.Client, examCode int32, circleOfficeID string, examYear string) ([]map[string]interface{}, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return nil, errors.New(" no such valid exam code exists")
	}

	if circleOfficeID == "" {
		return nil, errors.New(" please provide Nodal Officer's office ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_PS table
	count, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
		return nil, fmt.Errorf(" failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", circleOfficeID)
		return nil, errors.New(" no valid applications available for the circle")
	}

	// Query to get the applications matching the circleOfficeID
	applications, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_gdspm.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications: %v", err)
		return nil, fmt.Errorf(" failed to retrieve applications: %v", err)
	}

	uniqueEmployees := make(map[int64]struct{})
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_GDSPM) // Map to store the latest application for each employee
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

// Get GDSPM Exam statistics
func GetGDSPMExamStatistics(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Query to get the applications from Exam_Applications_GDSPM table
	applications, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.ExamCodeEQ(examCode),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_gdspm.FieldEmployeeID), ent.Desc(exam_applications_gdspm.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	// Create a map to store the latest application for each employee
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_GDSPM)

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
type CircleWiseSummaryGDSPM struct {
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

func GetGDSPMExamStatisticsCircleWise(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Query to get the applications from Exam_Applications_GDSPM table
	applications, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.ExamCodeEQ(examCode),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_gdspm.FieldEmployeeID), ent.Desc(exam_applications_gdspm.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	// Create a map to store employee-wise latest applications
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_GDSPM)

	// Loop through the applications and store the latest application for each employee
	for _, app := range applications {
		if _, found := employeeLatestApplication[app.EmployeeID]; !found {
			employeeLatestApplication[app.EmployeeID] = app
		}
	}

	// Create a map to store circle-wise summaries
	circleSummaries := make(map[string]*CircleWiseSummaryGDSPM)

	// Loop through the latest applications to update counts
	for _, app := range employeeLatestApplication {
		circleOfficeID := app.NodalOfficeFacilityID
		if circleSummaries[circleOfficeID] == nil {
			approvalFlag, err := getApprovalFlagForHallTicket(client, circleOfficeID)
			if err != nil {
				log.Printf("Failed to get ApprovalFlagForHallTicket for CircleOfficeID %v: %v", circleOfficeID, err)
				continue
			}

			circleSummaries[circleOfficeID] = &CircleWiseSummaryGDSPM{
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

// GDSPM office wise stats for NO
type DOOfficeWiseSummaryGDSPM struct {
	ControllingOfficeFacilityID string
	ControllingOfficeName       string
	Permitted                   int
	NotPermitted                int
	Pending                     int
	PendingWithCandidate        int
	Received                    int
	UniqueEmployees             map[int64]struct{}
}

func GetGDSPMExamStatisticsDOOfficeWise(ctx context.Context, client *ent.Client, examCode int32, facilityID string, examYear string) ([]map[string]interface{}, error) {
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

	// Query to get the applications from Exam_Applications_GDSPM table matching the provided facilityID
	applications, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Order(ent.Desc(exam_applications_gdspm.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
		return nil, fmt.Errorf(" failed to retrieve applications from Exam_Applications_GDSPM: %v", err)
	}

	// Create a map to store reporting office-wise summaries
	doOfficeSummaries := make(map[string]*DOOfficeWiseSummaryGDSPM)

	// Loop through the applications to group by reporting office-wise and update counts
	for _, app := range applications {
		reportingOfficeID := app.ControllingOfficeFacilityID

		if doOfficeSummaries[reportingOfficeID] == nil {
			doOfficeSummaries[reportingOfficeID] = &DOOfficeWiseSummaryGDSPM{
				ControllingOfficeFacilityID: reportingOfficeID,
				ControllingOfficeName:       app.ReportingOfficeName,
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
			"ReportingOfficeID":            summary.ControllingOfficeFacilityID,
			"ReportingOfficeName":          summary.ControllingOfficeName,
			"No: Of Applications Received": summary.Received,
			"No. Permitted":                summary.Permitted,
			"No. Not Permitted":            summary.NotPermitted,
			"No. Pending":                  summary.Pending,
			"No. Pending With Candidate":   summary.PendingWithCandidate,
		})
	}

	return result, nil
}

// // Get All Pending with Candidate
// Assuming Exam_Applications_GDSPM has a field named EmployeeID, you might adapt the code like this:
func QueryGDSPMApplicationsByPendingWithCandidate(ctx context.Context, client *ent.Client, facilityID string, facilityID1 string) ([]*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	if facilityID == "" || facilityID1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and Exam Year cannot be empty")
	}

	// Fetch all applications matching the criteria
	records, err := client.Exam_Applications_GDSPM.Query().
		Where(
			exam_applications_gdspm.ControllingOfficeFacilityIDEQ(facilityID),
			exam_applications_gdspm.ExamYearEQ(facilityID1),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Order(ent.Asc("employee_id")). /*, ent.Desc("application_number"))*/ // Order by employee_id and application_number
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf(" no application found for this CA Office ID %s", facilityID)
	}
	// Create a map to store the latest applications for each employee
	latestApplications := make(map[int64]*ent.Exam_Applications_GDSPM)

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
	var result []*ent.Exam_Applications_GDSPM
	for _, application := range latestApplications {
		result = append(result, application)
	}

	if len(result) == 0 {
		return nil, 422, " -STR004", false, fmt.Errorf(" no Applicationsf found pending with candiadte under this Office ID %s", facilityID)
	}

	return result, 200, "", true, nil
}

func GetGDSPMExamStatisticsDOOfficeWiseL(ctx context.Context, client *ent.Client, examCode int32, facilityID string, Examyear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		return nil, 422, " -STR002", false, errors.New("facility ID cannot be null")
	}

	// Fetch all applications for the given facilityID
	applications, err := client.Exam_Applications_GDSPM.
		Query().
		Where(exam_applications_gdspm.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_gdspm.StatusEQ("active"),
			exam_applications_gdspm.ExamYearEQ(Examyear),
		).
		Order(ent.Asc(exam_applications_gdspm.FieldEmployeeID), ent.Desc(exam_applications_gdspm.FieldUpdatedAt)).
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

						// if app.HallTicketNumber == ("") {
						// 	doOfficeSummary.HallTicketNotGenerated++
						// } else {
						// 	doOfficeSummary.HallTicketGenerated++
						// }

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

func GetMTSPMExamStatisticsDOOfficeWiseL(ctx context.Context, client *ent.Client, examCode int32, facilityID string, Examyear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		return nil, 422, " -STR002", false, errors.New("facility ID cannot be null")
	}

	// Fetch all applications for the given facilityID
	applications, err := client.Exam_Application_MTSPMMG.
		Query().
		Where(exam_application_mtspmmg.NodalOfficeFacilityIDEQ(facilityID),
			exam_application_mtspmmg.StatusEQ("active"),
			exam_application_mtspmmg.ExamYearEQ(Examyear),
		).
		Order(ent.Asc(exam_application_mtspmmg.FieldEmployeeID), ent.Desc(exam_application_mtspmmg.FieldUpdatedAt)).
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

type ExamStatsGDSPM struct {
	CircleName                  string
	ControllingOfficeName       string
	ControllingOfficeFacilityID string
	NodalOfficeFacilityID       string
	NoOfCandidatesChosenCity    int
	NoOfCandidatesAlloted       int
}

func GetExamApplicatonsPreferenenceCityWiseStatsGDSPM(ctx context.Context, client *ent.Client, ExamYear string, SubExamcode string, SubCityid string) ([]ExamStatsGDSPM, int32, string, bool, error) {

	var result []ExamStatsGDSPM

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

	applications, err := client.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.ExamYearEQ(ExamYear),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.HallTicketNumberNEQ(""),
			exam_applications_gdspm.StatusEQ("active"),
			exam_applications_gdspm.ExamCodeEQ(Examcode),
			exam_applications_gdspm.CenterIdEQ(Cityid),
		).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR004", false, err
	} else {
		if len(applications) == 0 {
			return nil, 422, " -STR005", false, errors.New("no matching data with Exam City")
		}
	}

	groupedApplications := make(map[string]ExamStatsGDSPM)

	for _, examApplication := range applications {
		nodalOfficeFacilityID := examApplication.NodalOfficeFacilityID
		nodalOfficeName := examApplication.NodalOfficeName

		controllingOfficeFacilityID := examApplication.ControllingOfficeFacilityID
		controllingOfficeName := examApplication.ControllingOfficeName
		centerCode := examApplication.ExamCityCenterCode

		// Check if ExamStats for the reporting office already exists
		stats, ok := groupedApplications[controllingOfficeName]
		if !ok {
			stats = ExamStatsGDSPM{
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

func UpdateCenterCodeForApplicationsGDSPM(ctx context.Context, client *ent.Client, controllingOfficeFacilityID string, examCenterID, seatsToAllot, examCityID int32) (int, []*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	// Input Validation
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
	// Querying Applications
	applications, err := tx.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.ExamCityCenterCodeIsNil(),
			exam_applications_gdspm.CenterIdEQ(examCityID),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.ControllingOfficeFacilityIDEQ(controllingOfficeFacilityID),
			exam_applications_gdspm.HallTicketNumberNEQ(""),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_gdspm.FieldApplnSubmittedDate)).
		Limit(int(seatsToAllot)). // Limit the number of records to be updated
		All(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR002", false, err
	}
	var updatedCount int

	// Updating Applications
	for _, application := range applications {
		_, err := tx.Exam_Applications_GDSPM.
			UpdateOne(application).
			SetExamCityCenterCode(examCenterID).
			Save(ctx)

		if err != nil {
			return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR003", false, err
		}
		updatedCount++
	}

	// Counting Updated Applications
	updateCount, err := tx.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.ExamCityCenterCodeEQ(examCenterID),
			exam_applications_gdspm.StatusEQ("active"),
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
func UpdateApplicationRemarksVAGDSPM(client *ent.Client, newAppln *ca_reg.VerifyApplicationVAGDSPM) (*ent.Exam_Applications_GDSPM, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	if newAppln == nil {
		return nil, 400, " -STR001", false, errors.New("payload is empty")
	}
	// Start a transaction.
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	// Defer rollback in case anything fails.
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			if err := tx.Commit(); err != nil {
				tx.Rollback()
			}
		}
	}()

	// Check if the EmployeeID exists.
	oldAppln, err := tx.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.EmployeeIDEQ(newAppln.EmployeeID),
			exam_applications_gdspm.ExamYearEQ(newAppln.ExamYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Only(ctx)
	currentTime := time.Now().Truncate(time.Second)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR003", false, errors.New("no active application found for this candidate")
		} else {
			return nil, 500, " -STR004", false, err
		}
	}

	switch oldAppln.ApplicationStatus {
	case "VerifiedByNA", "VerifiedByCA":
		return nil, 422, " -STR005", false, errors.New("this Application was already verified By Nodal / Controlling Authority")
	case "CAVerificationPending":
		oldAppln, err = oldAppln.
			Update().
			SetVAGeneralRemarks(newAppln.VA_GeneralRemarks).
			SetVADate(currentTime).
			SetVAEmployeeDesignation(newAppln.VA_EmployeeDesignation).
			SetVAEmployeeID(newAppln.VA_EmployeeID).
			SetVAUserName(newAppln.VA_UserName).
			Save(ctx)
		if err != nil {
			return nil, 500, " -STR005", false, err
		}

	case "ResubmitCAVerificationPending":
		oldAppln, err = oldAppln.
			Update().
			SetVAGeneralRemarks(newAppln.VA_GeneralRemarks).
			SetVADate(currentTime).
			SetVAEmployeeDesignation(newAppln.VA_EmployeeDesignation).
			SetVAEmployeeID(newAppln.VA_EmployeeID).
			SetVAUserName(newAppln.VA_UserName).
			Save(ctx)
		if err != nil {
			return nil, 500, " -STR006", false, err
		}
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR007", false, err
	}
	return oldAppln, 200, "", true, nil
}
func GenerateHallticketNumberrGdsPm(ctx context.Context, client *ent.Client, year string, examCode int32, nodalOfficerFacilityID string) (string, int32, string, bool, error) {
	//currentTime := time.Now().Truncate(time.Second)

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
	lastHallTicketNumber, err := tx.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.ExamYearEQ(year),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			exam_applications_gdspm.GenerateHallTicketFlagEQ(true),

			exam_applications_gdspm.HallTicketNumberNEQ(""),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Order(ent.Desc(exam_applications_gdspm.FieldHallTicketNumber)).
		First(ctx)

	fmt.Println(lastHallTicketNumber, "last ")

	if err != nil {
		if ent.IsNotFound(err) {
			lastHallTicketNumber = nil
		} else {
			return "", 500, " -STR001", false, err
		}
	}

	if lastHallTicketNumber == nil {
		lastHallTicketNumber = &ent.Exam_Applications_GDSPM{HallTicketNumber: "100000000"} // Assuming ExamApplicationsIP struct
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
	applications, err := tx.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.ExamYearEQ(year),
			exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			exam_applications_gdspm.GenerateHallTicketFlagEQ(true),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.StatusEQ("active"),
			exam_applications_gdspm.Or(
				exam_applications_gdspm.HallTicketNumberEQ(""),
			),
		).
		Order(ent.Asc(exam_applications_gdspm.FieldTempHallTicket)).
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
			//SetGenerateHallTicketFlagByNO(true).
			SetHallTicketGeneratedFlag(true).
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
