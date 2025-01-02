package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/center"
	"recruit/ent/exam_applications_ps"

	"recruit/ent/recommendationspsapplications"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"
	"strconv"
	"time"
)

func CreatePSApplications(client *ent.Client, newAppln *ca_reg.ApplicationGroupB) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateTempHallTicket(newAppln.TempHallTicket, newAppln.EmployeeID); err != nil {
		return nil, 422, " -STR100", false, err
	}
	if newAppln == nil {
		return nil, 422, " -STR001", false, errors.New("payload cannot be nil")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR004", false, err
	}
	defer func() {
		handleTransaction(tx, &err)
	}()

	statuses := []string{"CAVerificationPending", "ResubmitCAVerificationPending", "PendingWithCandidate", "VerifiedByCA"}
	existing, status, stgError, err := checkIfApplicationExists(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear, newAppln.ExamCode, statuses)
	if status == 500 {
		return nil, 500 + status, " -STR005 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR006 " + stgError, false, err

	}

	if existing {
		return nil, 422 + status, " -STR007" + stgError, false, errors.New("already application submitted for this candidate")
	}

	var currentTime = time.Now().UTC().Truncate(time.Second) // Ensure UTC for consistent comparison
	currentTime = currentTime.Add(5*time.Hour + 30*time.Minute)
	applicationLastDate := newAppln.ApplicationLastDate.UTC().Truncate(time.Second)

	fmt.Print("current time: ", currentTime, " application last date: ", applicationLastDate, "date from payload", newAppln.ApplicationLastDate)
	if currentTime.After(applicationLastDate) {
		return nil, 422, " -STR007", false, fmt.Errorf("application submission deadline has passed as current time is %v", currentTime)
	}

	// Generate Application number
	// Generate application number in the format "PS2023XXXXXX"
	applicationNumber, err := util.GenerateApplicationNumber(client, newAppln.EmployeeID, newAppln.ExamYear, "PS")
	if err != nil {
		return nil, 500, " -STR008", false, err
	}
	fmt.Println("Application number: ", applicationNumber)

	createdAppln, status, stgError, err := saveApplication(tx, newAppln, applicationNumber, newAppln.ExamCode, ctx)
	if err != nil {
		return nil, 500 + status, " -STR009 " + stgError, false, err
	}
	if err != nil {
		return nil, 500, " -STR010", false, err
	}
	return createdAppln, 200, "", true, nil
}

func SubGetLSGIPPSGRBApplicationsFacilityIDYear(ctx context.Context, client *ent.Client, facilityID string, year string) ([]*ent.Exam_Applications_PS, int32, error) {
	// Array of exams
	if facilityID == "" || year == "" {
		return nil, 422, errors.New("facility ID and exam year cannot be blank or null")
	}

	// Query the database for the specified conditions
	records, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ControllingOfficeFacilityIDEQ(facilityID),
			exam_applications_ps.ExamYearEQ(year),
			exam_applications_ps.StatusEQ("active"),
		).
		Order(ent.Desc(exam_applications_ps.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 400, fmt.Errorf("failed querying LSG/IP to PS Group B exams applications: %w", err)
	}

	// Check if no records are found
	if len(records) == 0 {
		return nil, 422, fmt.Errorf("no applications found for the year %s and facility ID %s", year, facilityID)
	}
	// Return the records and status code 200
	return records, 200, nil
}

// Create Applications with Circle Preferences ...

// Query PSExam Application with Emp ID.
func QueryPSExamApplicationsByEmpID(ctx context.Context, client *ent.Client, empid int64, examYear string) (*ent.Exam_Applications_PS, int32, string, bool, error) {
	newAppln, err := client.Exam_Applications_PS.
		Query().
		Where(
			(exam_applications_ps.EmployeeIDEQ(empid)),
			(exam_applications_ps.ExamYearEQ(examYear)),
			(exam_applications_ps.StatusEQ("active")),
		).
		//Order(ent.Desc(exam_applications_ps.FieldID)).
		WithCirclePrefRefPS().
		WithPSApplicationsRef().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, errors.New("no application exists ")
		} else {
			return nil, 500, " -STR002", false, err
		}
	}

	// Extract only the desired fields from the CirclePrefRef edge
	/* 	var circlePrefs []*ent.PlaceOfPreferencePS
	   	for _, edge := range newAppln.Edges.CirclePrefRefPS {
	   		circlePrefs = append(circlePrefs, &ent.PlaceOfPreferencePS{
	   			PlacePrefNo:    edge.PlacePrefNo,
	   			PlacePrefValue: edge.PlacePrefValue,
	   		})
	   	}

	   	// // Update the CirclePrefRef edge with the filtered values
	   	newAppln.Edges.CirclePrefRefPS = circlePrefs
	   	var recomondPref []*ent.RecommendationsPSApplications
	   	for _, edge := range newAppln.Edges.PSApplicationsRef {
	   		recomondPref = append(recomondPref, &ent.RecommendationsPSApplications{
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
	   	newAppln.Edges.PSApplicationsRef = recomondPref */

	newAppln.UpdatedAt = newAppln.UpdatedAt.Truncate(24 * time.Hour)

	return newAppln, 200, "", true, nil
}

// Update / Verification of PS Exam Application By CA
// Update Resubmission By Candidate.

// func UpdateApplicationRemarksPS(client *ent.Client, newAppln *ca_reg.VerifyApplicationGroupB, nonQualifyService string) (*ent.Exam_Applications_PS, int32, string, bool, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
// 	defer cancel()

// 	// Check if newAppln is not nil.
// 	if newAppln == nil {
// 		return nil, 400, " -STR001", false, errors.New("payload received in empty")
// 	}
// 	// Start a transaction.
// 	tx, err := client.Tx(ctx)
// 	if err != nil {
// 		return nil, 500, " -STR002", false, err
// 	}

// 	// Defer rollback in case anything fails.
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 			panic(r)
// 		} else if err != nil {
// 			tx.Rollback()
// 		} else {
// 			if err := tx.Commit(); err != nil {
// 				tx.Rollback()
// 			}
// 		}
// 	}()

// 	// Check if the EmployeeID exists.
// 	oldAppln, err := tx.Exam_Applications_PS.
// 		Query().
// 		Where(
// 			exam_applications_ps.EmployeeIDEQ(newAppln.EmployeeID),
// 			exam_applications_ps.ExamYearEQ(newAppln.ExamYear),
// 			exam_applications_ps.StatusEQ("active"),
// 		).
// 		Only(ctx)

// 	if err != nil {
// 		if ent.IsNotFound(err) {
// 			return nil, 422, " -STR003", false, errors.New("no active application for this candidate ")
// 		} else {
// 			return nil, 500, " -STR004", false, err
// 		}
// 	}

// 	// Insert a new record with the specified conditions.

// 	var updatedAppln *ent.Exam_Applications_PS
// 	var applicationStatus string

// 	currentTime := time.Now().Truncate(time.Second)
// 	// Format the current time to "yyyymmddhhmmss"
// 	stat := "inactive_" + time.Now().Format("20060102150405")
// 	// Determine the application status and set the corresponding action.
// 	switch oldAppln.ApplicationStatus {
// 	case "VerifiedByNA", "VerifiedByCA":
// 		return nil, 422, " -STR005", false, errors.New("the Application was already verified by Nodal Authority/ Controlling Authority")
// 	case "CAVerificationPending":
// 		if newAppln.CA_Remarks == "InCorrect" {
// 			oldAppln, err = oldAppln.
// 				Update().
// 				SetStatus(stat).
// 				Save(ctx)
// 			if err != nil {
// 				return nil, 500, " -STR006", false, err
// 			}
// 			applicationStatus = "PendingWithCandidate"
// 		} else if newAppln.CA_Remarks == "Correct" {
// 			oldAppln, err = oldAppln.
// 				Update().
// 				SetStatus(stat).
// 				Save(ctx)
// 			if err != nil {
// 				return nil, 500, " -STR007", false, err
// 			}
// 			applicationStatus = "VerifiedByCA" // Correctly set the application status here.
// 		}
// 	case "ResubmitCAVerificationPending":
// 		if newAppln.CA_Remarks == "InCorrect" {
// 			oldAppln, err = oldAppln.
// 				Update().
// 				SetStatus(stat).
// 				Save(ctx)
// 			if err != nil {
// 				return nil, 500, " -STR008", false, err
// 			}
// 			applicationStatus = "PendingWithCandidate"
// 		} else if newAppln.CA_Remarks == "Correct" {
// 			oldAppln, err = oldAppln.
// 				Update().
// 				SetStatus(stat).
// 				Save(ctx)
// 			if err != nil {
// 				return nil, 500, " -STR009", false, err
// 			}
// 			applicationStatus = "VerifiedByCA" // Correctly set the application status here.
// 		}
// 	}
// 	updatedAppln, err = tx.Exam_Applications_PS.
// 		Create().
// 		SetAppliactionRemarks(newAppln.AppliactionRemarks).
// 		SetApplicationNumber(oldAppln.ApplicationNumber).
// 		SetApplicationStatus(applicationStatus).
// 		SetApplnSubmittedDate(oldAppln.ApplnSubmittedDate).
// 		SetCADate(currentTime).
// 		SetCAEmployeeDesignation(newAppln.CA_EmployeeDesignation).
// 		SetCAEmployeeID(newAppln.CA_EmployeeID).
// 		SetCAGeneralRemarks(newAppln.CA_GeneralRemarks).
// 		SetCARemarks(newAppln.CA_Remarks).
// 		SetCAUserName(newAppln.CA_UserName).
// 		SetRecommendedStatus(newAppln.RecommendedStatus).
// 		SetCadre(oldAppln.Cadre).
// 		SetCandidateRemarks(oldAppln.CandidateRemarks).
// 		SetCategoryCode(oldAppln.CategoryCode).
// 		SetCategoryDescription(oldAppln.CategoryDescription).
// 		SetCenterFacilityId(oldAppln.CenterFacilityId).
// 		SetCenterId(oldAppln.CenterId).
// 		SetCentrePreference(oldAppln.CentrePreference).
// 		SetClaimingQualifyingService(oldAppln.ClaimingQualifyingService).
// 		SetControllingOfficeFacilityID(oldAppln.ControllingOfficeFacilityID).
// 		SetControllingOfficeName(oldAppln.ControllingOfficeName).
// 		SetDCCS(oldAppln.DCCS).
// 		SetDOB(oldAppln.DOB).
// 		SetDeputationControllingOfficeID(oldAppln.DeputationControllingOfficeID).
// 		SetDeputationControllingOfficeName(oldAppln.DeputationControllingOfficeName).
// 		SetDeputationOfficeFacilityID(oldAppln.DeputationOfficeFacilityID).
// 		SetDeputationOfficeName(oldAppln.DeputationOfficeName).
// 		SetDeputationOfficePincode(oldAppln.DeputationOfficePincode).
// 		SetDeputationOfficeUniqueId(oldAppln.DeputationOfficeUniqueId).
// 		SetInDeputation(oldAppln.InDeputation).
// 		SetDeputationType(oldAppln.DeputationType).
// 		SetDesignationID(oldAppln.DesignationID).
// 		SetDisabilityPercentage(oldAppln.DisabilityPercentage).
// 		SetDisabilityTypeCode(oldAppln.DisabilityTypeCode).
// 		SetDisabilityTypeDescription(oldAppln.DisabilityTypeDescription).
// 		SetDisabilityTypeID(oldAppln.DisabilityTypeID).
// 		SetEducationCode(oldAppln.EducationCode).
// 		SetEducationDescription(oldAppln.EducationDescription).
// 		SetEmailID(oldAppln.EmailID).
// 		SetEmployeeID(oldAppln.EmployeeID).
// 		SetEmployeeName(oldAppln.EmployeeName).
// 		SetEmployeePost(oldAppln.EmployeePost).
// 		SetEntryPostCode(oldAppln.EntryPostCode).
// 		SetEntryPostDescription(oldAppln.EntryPostDescription).
// 		SetExamCode(oldAppln.ExamCode).
// 		SetExamName(oldAppln.ExamName).
// 		SetExamShortName(oldAppln.ExamShortName).
// 		SetExamYear(oldAppln.ExamYear).
// 		SetExamCityCenterCode(oldAppln.CenterId).
// 		SetFacilityUniqueID(oldAppln.FacilityUniqueID).
// 		SetFeederPostCode(oldAppln.FeederPostCode).
// 		SetFeederPostDescription(oldAppln.FeederPostDescription).
// 		SetFeederPostJoiningDate(oldAppln.FeederPostJoiningDate).
// 		SetGender(oldAppln.Gender).
// 		SetGenerateHallTicketFlag(newAppln.GenerateHallTicketFlag).
// 		SetLienControllingOfficeID(oldAppln.LienControllingOfficeID).
// 		SetLienControllingOfficeName(oldAppln.LienControllingOfficeName).
// 		SetMobileNumber(oldAppln.MobileNumber).
// 		SetNodalOfficeFacilityID(oldAppln.NodalOfficeFacilityID).
// 		SetNodalOfficeName(oldAppln.NodalOfficeName).
// 		SetNonQualifyingService(*newAppln.NonQualifyingService).
// 		SetOptionUsed(oldAppln.OptionUsed).
// 		SetPhoto(oldAppln.Photo).
// 		SetPhotoPath(oldAppln.PhotoPath).
// 		SetPresentDesignation(oldAppln.PresentDesignation).
// 		SetPresentPostCode(oldAppln.PresentPostCode).
// 		SetPresentPostDescription(oldAppln.PresentPostDescription).
// 		SetReportingOfficeFacilityID(oldAppln.ReportingOfficeFacilityID).
// 		SetReportingOfficeName(oldAppln.ReportingOfficeName).
// 		SetServiceLength(oldAppln.ServiceLength).
// 		SetSignature(oldAppln.Signature).
// 		SetSignaturePath(oldAppln.SignaturePath).
// 		SetStatus("active").
// 		SetTempHallTicket(oldAppln.TempHallTicket).
// 		SetUserID(oldAppln.UserID).
// 		SetWorkingOfficeCircleFacilityID(oldAppln.WorkingOfficeCircleFacilityID).
// 		SetWorkingOfficeCircleName(oldAppln.WorkingOfficeCircleName).
// 		SetWorkingOfficeDivisionFacilityID(oldAppln.WorkingOfficeDivisionFacilityID).
// 		SetWorkingOfficeDivisionName(oldAppln.WorkingOfficeDivisionName).
// 		SetWorkingOfficeFacilityID(oldAppln.WorkingOfficeFacilityID).
// 		SetWorkingOfficeName(oldAppln.WorkingOfficeName).
// 		SetWorkingOfficePincode(oldAppln.WorkingOfficePincode).
// 		SetWorkingOfficeRegionFacilityID(oldAppln.WorkingOfficeRegionFacilityID).
// 		SetWorkingOfficeRegionName(oldAppln.WorkingOfficeRegionName).
// 		SetPunishmentStatus(newAppln.PunishmentStatus).
// 		SetDisciplinaryCaseStatus(newAppln.DisciplinaryCaseStatus).
// 		Save(ctx)
// 	if err != nil {
// 		return nil, 500, " -STR010", false, err
// 	}
// 	if nonQualifyService == "Yes" {
// 		_, err = updatedAppln.
// 			Update().
// 			SetNonQualifyingService(*newAppln.NonQualifyingService).
// 			Save(ctx)
// 		if err != nil {
// 			return nil, 500, " -STR011", false, err
// 		}
// 	}
// 	recommendationsRef := make([]*ent.RecommendationsPSApplications, len(newAppln.Edges.ApplicationData))
// 	for i, recommendation := range newAppln.Edges.ApplicationData {
// 		if recommendation.VacancyYear == 0 {
// 			return nil, 400, " -STR015", false, fmt.Errorf("recommendations value at index %d is nil", i)
// 		}
// 		RecommendationsRefEntity, err := tx.RecommendationsPSApplications.
// 			Create().
// 			SetApplicationID(updatedAppln.ID).
// 			SetEmployeeID(updatedAppln.EmployeeID).
// 			SetExamYear(updatedAppln.ExamYear).
// 			SetVacancyYear(recommendation.VacancyYear).
// 			SetCARecommendations(recommendation.CA_Recommendations).
// 			SetNORecommendations(recommendation.CA_Recommendations).
// 			SetCAUserName(newAppln.CA_UserName).
// 			SetCARemarks(recommendation.CA_Remarks).
// 			SetCAUpdatedAt(currentTime).
// 			SetNOUpdatedAt(currentTime).
// 			SetApplicationStatus("VerifiedRecommendationsByCA").
// 			Save(ctx)
// 		if err != nil {
// 			return nil, 500, " -STR016", false, err
// 		}
// 		recommendationsRef[i] = RecommendationsRefEntity
// 	}
// 	_, err = updatedAppln.Update().
// 		ClearPSApplicationsRef().
// 		AddPSApplicationsRef(recommendationsRef...).
// 		Save(ctx)
// 	if err = tx.Commit(); err != nil {
// 		tx.Rollback()
// 		return nil, 500, " -STR017", false, err
// 	}
// 	return updatedAppln, 200, "", true, nil
// }

func UpdateApplicationRemarksPS(client *ent.Client, newAppln *ca_reg.VerifyApplicationGroupB, nonQualifyService string) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	if newAppln == nil {
		return nil, 400, " -STR001", false, errors.New("payload received in empty")
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	defer handleTransaction(tx, &err)

	// Fetch the existing application.
	oldAppln, status, stgError, err := fetchExistingPsApplication(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear)
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
	updatedAppln, err := createUpdateApplication(ctx, tx, oldAppln, newAppln, applicationStatus, nonQualifyService)
	if err != nil {
		return nil, 500, " -STR009", false, err
	}

	// Handle the recommendations.
	err = handleRecommendations(ctx, tx, updatedAppln, newAppln)
	if err != nil {
		return nil, 500, " -STR016", false, err
	}

	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB017", false, err
	}
	return appResponse, 200, "", true, nil
}

// func ResubmitApplicationRemarksPS(client *ent.Client, newAppln *ca_reg.ReApplicationGroupB) (*ent.Exam_Applications_PS, int32, string, bool, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
// 	defer cancel()
// 	if !isNumeric(newAppln.TempHallTicket) {
// 		return nil, 400, " -STR100", false, fmt.Errorf("issue for employee %d with temp hall ticket number: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
// 	}
// 	if len(newAppln.TempHallTicket) != 8 {
// 		return nil, 400, " -STR101", false, fmt.Errorf("issue for employee %d with temp hall ticket number length issue: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
// 	}
// 	// Check if newAppln is not nil.
// 	if newAppln == nil {
// 		return nil, 400, " -STR001", false, errors.New("payload is empty")
// 	}
// 	tx, err := client.Tx(ctx)
// 	if err != nil {
// 		return nil, 500, " -STR002", false, err
// 	}

// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 			panic(r)
// 		} else if err != nil {
// 			tx.Rollback()
// 		} else {
// 			if err := tx.Commit(); err != nil {
// 				tx.Rollback()
// 			}
// 		}
// 	}()
// 	// Check if the EmployeeID exists.
// 	oldAppln, err := tx.Exam_Applications_PS.
// 		Query().
// 		Where(
// 			exam_applications_ps.EmployeeIDEQ(newAppln.EmployeeID),
// 			exam_applications_ps.ExamYearEQ(newAppln.ExamYear),
// 			exam_applications_ps.StatusEQ("active"),
// 		).
// 		Only(ctx)
// 	if err != nil {
// 		if ent.IsNotFound(err) {
// 			return nil, 422, " -STR003", false, errors.New("no active application found for this candidate")
// 		} else {
// 			return nil, 500, " -STR004", false, err
// 		}
// 	}
// 	exists, err := tx.Exam_Applications_PS.
// 		Query().
// 		Where(
// 			exam_applications_ps.EmployeeIDEQ(newAppln.EmployeeID),
// 			exam_applications_ps.ApplicationStatusIn("ResubmitCAVerificationPending"),
// 			exam_applications_ps.ExamYearEQ(newAppln.ExamYear),
// 			exam_applications_ps.StatusEQ("active"),
// 		).
// 		Exist(ctx)

// 	if err != nil {
// 		return nil, 500, " -STR005", false, err
// 	}

// 	if exists {
// 		return nil, 422, " -STR006", false, errors.New("already application resubmitted for this candidate")
// 	}
// 	currentTime := time.Now().Truncate(time.Second)
// 	// Format the current time to "yyyymmddhhmmss"
// 	stat := "inactive_" + time.Now().Format("20060102150405")

// 	if oldAppln == nil {
// 		return nil, 422, " -STR007", false, errors.New("no active application found for this candidate")
// 	} else {
// 		// Update the existing record.
// 		if oldAppln.ApplicationStatus == "VerifiedByNA" || oldAppln.ApplicationStatus == "VerifiedByCA" {
// 			return nil, 422, " -STR008", false, errors.New("the Application was already verified By Nodal Authority/ Controlling Authority")
// 		}

// 		if oldAppln.ApplicationStatus != "PendingWithCandidate" {
// 			return nil, 422, " -STR009", false, errors.New("this application was not in pending with candidate status")
// 		} else {
// 			applicationNumber, err := util.GenerateApplicationNumber(client, newAppln.EmployeeID, newAppln.ExamYear, "PS")
// 			if err != nil {
// 				return nil, 400, " -STR010", false, fmt.Errorf("failed to generate application number: %v", err)
// 			}
// 			// Insert a new record.

// 			_, err = oldAppln.
// 				Update().
// 				SetStatus(stat).
// 				Save(ctx)
// 			if err != nil {
// 				return nil, 500, " -STR011", false, err
// 			}
// 			updatedAppln, err := tx.Exam_Applications_PS.
// 				Create().
// 				SetApplicationNumber(applicationNumber).
// 				SetApplicationStatus("ResubmitCAVerificationPending").
// 				SetApplnSubmittedDate(currentTime).
// 				SetCAEmployeeDesignation(oldAppln.CAEmployeeDesignation).
// 				SetCADate(oldAppln.CADate).
// 				SetCAEmployeeID(oldAppln.CAEmployeeID).
// 				SetCAGeneralRemarks(oldAppln.CAGeneralRemarks).
// 				SetCAUserId(oldAppln.CAUserId).
// 				SetCAUserName(oldAppln.CAUserName).
// 				SetCandidateRemarks(newAppln.CandidateRemarks).
// 				SetCategoryCode(newAppln.CategoryCode).
// 				SetCategoryDescription(newAppln.CategoryDescription).
// 				SetCenterFacilityId(newAppln.CenterFacilityId).
// 				SetCenterId(newAppln.CenterId).
// 				SetCentrePreference(newAppln.CentrePreference).
// 				SetCentrePreference(newAppln.CentrePreference).
// 				SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
// 				SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
// 				SetControllingOfficeName(newAppln.ControllingOfficeName).
// 				SetDCCS(newAppln.DCCS).
// 				SetDOB(newAppln.DOB).
// 				SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
// 				SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
// 				SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
// 				SetDeputationOfficeName(newAppln.DeputationOfficeName).
// 				SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
// 				SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
// 				SetInDeputation(newAppln.InDeputation).
// 				SetDeputationType(newAppln.DeputationType).
// 				SetDesignationID(newAppln.DesignationID).
// 				SetDisabilityPercentage(newAppln.DisabilityPercentage).
// 				SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
// 				SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
// 				SetDisabilityTypeID(newAppln.DisabilityTypeID).
// 				SetEducationCode(newAppln.EducationCode).
// 				SetEducationDescription(newAppln.EducationDescription).
// 				SetEmailID(newAppln.EmailID).
// 				SetEmployeeID(newAppln.EmployeeID).
// 				SetEmployeeName(newAppln.EmployeeName).
// 				SetEntryPostCode(newAppln.EntryPostCode).
// 				SetEntryPostDescription(newAppln.EntryPostDescription).
// 				SetExamCode(newAppln.ExamCode).
// 				SetExamName(newAppln.ExamName).
// 				SetExamShortName(newAppln.ExamShortName).
// 				SetExamYear(newAppln.ExamYear).
// 				SetExamCityCenterCode(newAppln.CenterId).
// 				SetFacilityUniqueID(newAppln.FacilityUniqueID).
// 				SetFeederPostCode(newAppln.FeederPostCode).
// 				SetFeederPostDescription(newAppln.FeederPostDescription).
// 				SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
// 				SetGender(newAppln.Gender).
// 				SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
// 				SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
// 				SetMobileNumber(newAppln.MobileNumber).
// 				SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
// 				SetNodalOfficeName(newAppln.NodalOfficeName).
// 				SetPhoto(newAppln.Photo).
// 				SetPhotoPath(newAppln.PhotoPath).
// 				SetPresentDesignation(newAppln.PresentDesignation).
// 				SetPresentPostCode(newAppln.PresentPostCode).
// 				SetPresentPostDescription(newAppln.PresentPostDescription).
// 				SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
// 				SetReportingOfficeName(newAppln.ReportingOfficeName).
// 				SetServiceLength(*newAppln.ServiceLength).
// 				SetSignature(newAppln.Signature).
// 				SetSignaturePath(newAppln.SignaturePath).
// 				SetStatus("active").
// 				SetTempHallTicket(newAppln.TempHallTicket).
// 				SetUserID(newAppln.UserID).
// 				SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
// 				SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
// 				SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
// 				SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
// 				SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
// 				SetWorkingOfficeName(newAppln.WorkingOfficeName).
// 				SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
// 				SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
// 				SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
// 				Save(ctx)

// 			if err != nil {
// 				return nil, 500, " -STR012", false, err
// 			}

// 			if err := tx.Commit(); err != nil {
// 				tx.Rollback()
// 				return nil, 500, " -STR013", false, err
// 			}
// 			return updatedAppln, 200, "", true, nil
// 		}
// 	}
// }

func ResubmitApplicationRemarksPSs(client *ent.Client, newAppln *ca_reg.ReApplicationGroupB) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	// Set up context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Validate temporary hall ticket
	if err := validateTempHallTicket(newAppln.TempHallTicket, newAppln.EmployeeID); err != nil {
		return nil, 422, " -STR001", false, fmt.Errorf("issue for employee %d with temp hall ticket number: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
	}

	// Begin transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	// Ensure transaction is handled properly
	defer func() {
		handleTransaction(tx, &err)
	}()

	// Fetch old application if it exists
	oldAppln, status, stgError, err := fetchOldApplicationPs(ctx, tx, newAppln)
	if err != nil {
		if status == 500 {
			return nil, 500 + status, " -STR004 " + stgError, false, err
		}
		if status == 422 {
			return nil, 422 + status, " -STR005 " + stgError, false, fmt.Errorf("no active application found for this candidate")
		}
	}

	// Check if a resubmission already exists
	statuses := []string{"ResubmitCAVerificationPending"}
	existing, status, stgError, err := checkIfApplicationExists(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear, newAppln.ExamCode, statuses)
	if err != nil {
		if status == 500 {
			return nil, 500 + status, " -STR004 " + stgError, false, err
		}
		if status == 422 {
			return nil, 422 + status, " -STR005 " + stgError, false, err
		}
	}
	if existing {
		return nil, 422 + status, " -STR006 " + stgError, false, fmt.Errorf("an application has already been submitted for this candidate")
	}
	var currentTime = time.Now().UTC().Truncate(time.Second) // Ensure UTC for consistent comparison
	currentTime = currentTime.Add(5*time.Hour + 30*time.Minute)
	correctionLastDate := newAppln.ApplicationCorrectionLastDate.UTC().Truncate(time.Second)

	fmt.Print("current time: ", currentTime, " correction last date: ", correctionLastDate, "date from payload", newAppln.ApplicationCorrectionLastDate)
	if currentTime.After(correctionLastDate) {
		return nil, 422, " -STR007", false, fmt.Errorf("application correction deadline has passed as current time is %v", currentTime)
	}

	// Process resubmission
	updatedAppln, status, stgError, err := processResubmission(ctx, tx, oldAppln, newAppln, int(newAppln.ExamCode))
	if err != nil {
		if status == 500 {
			return nil, 500 + status, " -STR008 " + stgError, false, err
		}
		if status == 422 {
			return nil, 422 + status, " -STR009 " + stgError, false, err
		}
	}

	// Map the updated application to the response
	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB010 ", false, err
	}

	// Return the successful response
	return appResponse, 200, "", true, nil
}

// UpdateNodalRecommendationsByEmpID updates the recommendations for a given employee ID
func UpdateNodalRecommendationsPSByEmpID(client *ent.Client, applicationRecord *ca_reg.NAVerifyApplicationGroupB) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validatePsInput(applicationRecord); err != nil {
		return nil, 422, " -STR001", false, errors.New("employee id should not be empty")
	}
	empID := applicationRecord.EmployeeID
	id1 := applicationRecord.ExamYear

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	defer handleTransaction(tx, &err)

	// Check if empID exists in exam_applications_ps and the status is "VerifiedByCA" or "VerifiedByNA"
	exists, status, stgError, err := checkPsApplicationExists(tx, ctx, applicationRecord)

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
	records, err := getPsRecommendationsByEmpID(ctx, tx, empID)
	if err != nil {
		return nil, 500, " -STR008 ", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR009 ", false, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}

	// Format the current time to "yyyymmddhhmmss"
	stat := "inactive_" + time.Now().Format("20060102150405")
	// Update the retrieved record with the provided values
	updatedRecord, status, stgError, err := getActiveExamApplicationPS(ctx, tx, empID, id1)
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
		return nil, 422, " -STR013 ", false, fmt.Errorf("failed to update application: %v", err)
	}
	// Hall Ticket Generated Flag
	hallticketgeneratedflag := checkPsHallTicketGenerated(applicationRecord, updatedRecord)

	updatedAppln, err := createUpdatedPsAppln(tx, applicationRecord, updatedRecord, hallticketgeneratedflag, ctx)
	if err != nil {
		return nil, 500, " -STR014 ", false, err
	}

	// Save the Recommendation records.
	recommendationsRef, err := createPsRecommendationsRef(ctx, tx, applicationRecord, updatedAppln)
	if err != nil {
		return nil, 500, " -STR015", false, err
	}

	updatedAppln.Update().
		AddPSApplicationsRef(recommendationsRef...).
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

// Get All CA Pending records ...
func QueryPSApplicationsByCAVerificationsPending(ctx context.Context, client *ent.Client, facilityID string, examYear string) ([]*ent.Exam_Applications_PS, int32, string, bool, error) {
	// Array of exams

	if facilityID == "" || examYear == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID and Examyear cannot be null")
	}

	records, err := client.Exam_Applications_PS.Query().
		Where(
			exam_applications_ps.And(
				exam_applications_ps.Or(
					exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				),
				exam_applications_ps.ControllingOfficeFacilityIDEQ(facilityID),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
			),
		).
		Order(ent.Desc(exam_applications_ps.FieldID)). // Order by descending updated_at timestamp
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no Applications is pending for CA pending verification for the Office ID %s", facilityID)
	}

	return records, 200, "", true, nil
}

// Get All CA verified records
func QueryPSApplicationsByCAVerified(ctx context.Context, client *ent.Client, facilityID string, examYear string) ([]*ent.Exam_Applications_PS, int32, string, bool, error) {
	if facilityID == "" || examYear == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_PS.Query().
		Where(
			exam_applications_ps.And(
				exam_applications_ps.ApplicationStatusEQ("VerifiedByCA"),
				exam_applications_ps.ControllingOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
			),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application pending for  CA verification for the Office ID %s", facilityID)
	}
	return records, 200, "", true, nil
}

// Generate ht with centercode
func GenerateHallticketNumberPSwithCenterCode(ctx context.Context, client *ent.Client) (string, error) {

	currentTime := time.Now().Truncate(time.Second)
	tx, err := client.Tx(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to start transaction: %w", err)
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

	applications, err := tx.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ps.ExamCodeNEQ(0),
			exam_applications_ps.ExamYearNEQ(""),
			exam_applications_ps.CategoryCodeNEQ(""),
			exam_applications_ps.WorkingOfficeCircleFacilityIDNEQ(""),
			exam_applications_ps.EmployeeIDNEQ(0),
			exam_applications_ps.HallTicketGeneratedFlagNEQ(true),
		).
		Order(ent.Asc(exam_applications_ps.FieldID)).
		All(ctx)
	if err != nil {
		return "", err
	}

	circleStatsPS := make(map[string]int)
	for _, application := range applications {
		key := application.WorkingOfficeCircleFacilityID
		circleStatsPS[key]++

		identificationNo := circleStatsPS[key]
		examYear := application.ExamYear
		if len(examYear) >= 2 {
			examYear = examYear[len(examYear)-2:]
		}
		hallticketNumber := util.GenerateHallticketNumberPS(
			application.ExamCode,
			examYear,
			application.CategoryCode,
			application.WorkingOfficeCircleFacilityID,
			application.WorkingOfficeDivisionFacilityID,
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

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return "", err
	}
	return successMessage, nil
}

// Get CA Verified with Emp ID
func QueryPSApplicationsByCAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_PS, int32, string, bool, error) {

	record, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ApplicationStatusEQ("VerifiedByCA"), // Check for "CAVerified" status
			exam_applications_ps.EmployeeIDEQ(employeeID),
			exam_applications_ps.StatusEQ("active"),
		).
		WithPSApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf("no application found for this employee ID: %d with 'CAVerified' status", employeeID)
		}
		return nil, 500, " -STR002", false, err
	}
	return record, 200, "", true, nil
}

func QueryPSApplicationsByCAPendingByEmpID(ctx context.Context, client *ent.Client, empID int64, examYear string) (*ent.Exam_Applications_PS, int32, string, bool, error) {
	// Check if employee ID exists

	// Retrieve the latest record based on UpdatedAt timestamp
	record, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.EmployeeIDEQ(empID),
			exam_applications_ps.Or(
				exam_applications_ps.ApplicationStatusEQ("CAVerificationPending"),
				exam_applications_ps.ApplicationStatusEQ("ResubmitCAVerificationPending"),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
			),
		).
		WithPSApplicationsRef().
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
func GetOldPSCAApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_PS, int32, string, bool, error) {

	application, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.EmployeeIDEQ(employeeID),
			exam_applications_ps.ApplicationStatusEQ("PendingWithCandidate"),
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.StatusEQ("active"),
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
func GetPSRecommendationsByEmpID(client *ent.Client, empID int64) ([]*ent.RecommendationsPSApplications, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	// Check if empID is null
	if empID == 0 {
		return nil, 422, " -STR001", false, fmt.Errorf("no employee ID provided to process")
	}

	// Retrieve all records for the employee ID
	records, err := client.RecommendationsPSApplications.Query().
		Where(recommendationspsapplications.EmployeeIDEQ(empID)).
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
func QueryPSApplicationsByNAVerified(ctx context.Context, client *ent.Client, facilityID string, examYear string) ([]*ent.Exam_Applications_PS, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || examYear == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_PS.Query().
		Where(
			exam_applications_ps.And(
				exam_applications_ps.ApplicationStatusEQ("VerifiedByNA"),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.ControllingOfficeFacilityIDEQ(facilityID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application pending for NA verification for the Office ID %s", facilityID)
	}
	return records, 200, "", true, nil
}

// Get All NA Verified Records with Emp ID
func QueryPSApplicationsByNAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_PS, int32, string, bool, error) {

	record, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ApplicationStatusEQ("VerifiedByNA"), // Check for "CAVerified" status
			exam_applications_ps.EmployeeIDEQ(employeeID),
			exam_applications_ps.StatusEQ("active"),
		).
		WithPSApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf("no application pending with CA Verification for  employee ID: %d ", employeeID)
		}
		return nil, 500, " -STR002", false, err
	}
	return record, 200, "", true, nil
}

func QueryPSApplicationsByNAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID, examYear string) ([]*ent.Exam_Applications_PS, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || examYear == "" {
		return nil, 422, " -STR001", false, fmt.Errorf("facility ID and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_PS.Query().
		Where(
			exam_applications_ps.And(
				exam_applications_ps.ApplicationStatusEQ("VerifiedByNA"),
				exam_applications_ps.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
			),
		).
		//WithCirclePrefRefPS().
		WithPSApplicationsRef().
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
func QueryPSApplicationsByCAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID, examYear string) ([]*ent.Exam_Applications_PS, int32, string, bool, error) {
	if facilityID == "" || examYear == "" {
		return nil, 422, " -STR001", false, fmt.Errorf("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_PS.Query().
		Where(
			exam_applications_ps.And(
				exam_applications_ps.ApplicationStatus("VerifiedByCA"),
				exam_applications_ps.NodalOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_applications_ps.ExamYearEQ(examYear),                // Add the Where clause
				exam_applications_ps.StatusEQ("active"),                  // Add the Where clause
			),
		).
		WithPSApplicationsRef().
		//WithCirclePrefRefPS().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application verified by CA/NA for the Office ID %s", facilityID)
	}
	return records, 200, "", true, nil
}

// Get Recommendations with Emp ID ..

// Get Circle details summary ofExam Applications for the Nodal Officer Office ID. - For IP ALone
func GetEligiblePSApplicationsForCircleDetails(ctx context.Context, client *ent.Client, examCode int32, circleOfficeID string, Examyear string) ([]map[string]interface{}, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return nil, fmt.Errorf(" no such valid exam code exists")
	}

	if circleOfficeID == "" {
		return nil, fmt.Errorf("please provide Nodal Officer's office ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_PS table
	count, err := client.Exam_Applications_PS.
		Query().
		Where(exam_applications_ps.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_ps.StatusEQ("active"),
			exam_applications_ps.ExamYearEQ(Examyear),
		).
		Count(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_PS: %v", err)
		return nil, fmt.Errorf(" failed to retrieve applications from Exam_Applications_PS: %v", err)
	}

	if count == 0 {
		log.Printf("No valid applications available for the circle: %s", circleOfficeID)
		return nil, fmt.Errorf(" no valid applications available for the circle")
	}

	// Query to get the applications matching the circleOfficeID
	applications, err := client.Exam_Applications_PS.
		Query().
		Where(exam_applications_ps.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_ps.StatusEQ("active"),
			exam_applications_ps.ExamYearEQ(Examyear),
		).
		Order(ent.Asc(exam_applications_ps.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications: %v", err)
		return nil, fmt.Errorf(" failed to retrieve applications: %v", err)
	}

	uniqueEmployees := make(map[int64]struct{})
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_PS) // Map to store the latest application for each employee
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

// Get PS Exam statistics
func GetPSExamStatistics(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Query to get the applications from Exam_Applications_IP table
	applications, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ExamCodeEQ(examCode),
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_ps.FieldEmployeeID), ent.Desc(exam_applications_ps.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	// Create a map to store the latest application for each employee
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_PS)

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
			} else if !*app.GenerateHallTicketFlag {
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
type CircleWiseSummaryPS struct {
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

func GetPSExamStatisticsCircleWise(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Query to get the applications from Exam_Applications_IP table
	applications, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ExamCodeEQ(examCode),
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_ps.FieldEmployeeID), ent.Desc(exam_applications_ps.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	// Create a map to store employee-wise latest applications
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_PS)

	// Loop through the applications and store the latest application for each employee
	for _, app := range applications {
		if _, found := employeeLatestApplication[app.EmployeeID]; !found {
			employeeLatestApplication[app.EmployeeID] = app
		}
	}

	// Create a map to store circle-wise summaries
	circleSummaries := make(map[string]*CircleWiseSummaryPS)

	// Loop through the latest applications to update counts
	for _, app := range employeeLatestApplication {
		circleOfficeID := app.NodalOfficeFacilityID
		if circleSummaries[circleOfficeID] == nil {
			approvalFlag, err := getApprovalFlagForHallTicket(client, circleOfficeID)
			if err != nil {
				log.Printf("Failed to get ApprovalFlagForHallTicket for CircleOfficeID %v: %v", circleOfficeID, err)
				continue
			}

			circleSummaries[circleOfficeID] = &CircleWiseSummaryPS{
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

// IP office wise stats for NO
type DOOfficeWiseSummaryPS struct {
	ControllingOfficeFacilityID string
	ControllingOfficeName       string
	Permitted                   int
	NotPermitted                int
	PendingWithCA               int
	PendingWithCandidate        int
	Received                    int
	HallTicketGenerated         int
	HallTicketNotGenerated      int
	UniqueEmployees             map[int64]struct{}
}

func GetPSExamStatisticsDOOfficeWise(ctx context.Context, client *ent.Client, examCode int32, facilityID string, examYear string) ([]map[string]interface{}, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return nil, fmt.Errorf(" no such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		log.Println("Facility ID cannot be null")
		return nil, fmt.Errorf(" facility ID cannot be null")
	}

	// Query to get the applications from Exam_Applications_PS table matching the provided facilityID
	applications, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.StatusEQ("active"),
		).
		Order(ent.Desc(exam_applications_ps.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications from Exam_Applications_PS: %v", err)
		return nil, fmt.Errorf(" failed to retrieve applications from Exam_Applications_PS: %v", err)
	}

	// Create a map to store reporting office-wise summaries
	doOfficeSummaries := make(map[string]*DOOfficeWiseSummaryPS)

	// Loop through the applications to group by reporting office-wise and update counts
	for _, app := range applications {
		reportingOfficeID := app.ReportingOfficeFacilityID

		if doOfficeSummaries[reportingOfficeID] == nil {
			doOfficeSummaries[reportingOfficeID] = &DOOfficeWiseSummaryPS{
				ControllingOfficeFacilityID: reportingOfficeID,
				ControllingOfficeName:       app.ReportingOfficeName,
				Permitted:                   0,
				NotPermitted:                0,
				PendingWithCA:               0,
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
					doOfficeSummary.PendingWithCA++
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
			"SNo.":                        serialNumber,
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

// // Get All Pending with Candidate
// Assuming Exam_Applications_PS has a field named EmployeeID, you might adapt the code like this:
func QueryPSApplicationsByPendingWithCandidate(ctx context.Context, client *ent.Client, facilityID string, examYear string) ([]*ent.Exam_Applications_PS, int32, string, bool, error) {
	if facilityID == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and Exam Year cannot be empty")
	}

	// Fetch all applications matching the criteria
	records, err := client.Exam_Applications_PS.Query().
		Where(
			exam_applications_ps.ReportingOfficeFacilityIDEQ(facilityID),
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.ApplicationStatusEQ("PendingWithCandidate"),
			exam_applications_ps.StatusEQ("active"),
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
	// latestApplications := make(map[int64]*ent.Exam_Applications_PS)

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
	// var result []*ent.Exam_Applications_PS
	// for _, application := range latestApplications {
	// 	result = append(result, application)
	// }

	if len(records) == 0 {
		return nil, 422, " -STR004", false, fmt.Errorf("no Applicationsf found pending with candiadte under this Office ID %s", facilityID)
	}

	return records, 200, "", true, nil
}

func GetPSExamStatisticsDOOfficeWiseL(ctx context.Context, client *ent.Client, examCode int32, facilityID string, Examyear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		return nil, 422, " -STR002", false, errors.New("facility ID cannot be null")
	}

	// Fetch all applications for the given facilityID
	applications, err := client.Exam_Applications_PS.
		Query().
		Where(exam_applications_ps.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_ps.StatusEQ("active"),
			exam_applications_ps.ExamYearEQ(Examyear),
		).
		Order(ent.Asc(exam_applications_ps.FieldEmployeeID), ent.Desc(exam_applications_ps.FieldUpdatedAt)).
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
	doOfficeSummaries := make(map[string]*DOOfficeWiseSummaryPS)

	// Loop through the greatest ApplicationIDs to get the corresponding applications
	for employeeID, greatestAppID := range greatestAppIDs {
		for _, app := range applications {
			if app.EmployeeID == employeeID && int32(app.ID) == greatestAppID {
				// Process the application here
				reportingOfficeID := app.ControllingOfficeFacilityID

				if doOfficeSummaries[reportingOfficeID] == nil {
					doOfficeSummaries[reportingOfficeID] = &DOOfficeWiseSummaryPS{
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

type ExamStatsPS struct {
	CircleName                  string
	ControllingOfficeName       string
	ControllingOfficeFacilityID string
	NodalOfficeFacilityID       string
	NoOfCandidatesChosenCity    int
	NoOfCandidatesAlloted       int
}

func GetExamApplicatonsPreferenenceCityWiseStatsPS(ctx context.Context, client *ent.Client, ExamYear string, SubExamcode string, SubCityid string) ([]ExamStatsPS, int32, string, bool, error) {
	var result []ExamStatsPS

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

	applications, err := client.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ExamYearEQ(ExamYear),
			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ps.HallTicketNumberNEQ(""),
			exam_applications_ps.StatusEQ("active"),
			exam_applications_ps.ExamCodeEQ(Examcode),
			exam_applications_ps.CenterIdEQ(Cityid),
		).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR004", false, err
	} else {
		if len(applications) == 0 {
			return nil, 422, " -STR005", false, errors.New("no matching data with Exam City")
		}
	}

	groupedApplications := make(map[string]ExamStatsPS)

	for _, examApplication := range applications {
		nodalOfficeFacilityID := examApplication.NodalOfficeFacilityID
		nodalOfficeName := examApplication.NodalOfficeName

		controllingOfficeFacilityID := examApplication.ControllingOfficeFacilityID
		controllingOfficeName := examApplication.ControllingOfficeName

		centerCode := examApplication.ExamCityCenterCode

		// Check if ExamStats for the reporting office already exists
		stats, ok := groupedApplications[controllingOfficeName]
		if !ok {
			stats = ExamStatsPS{
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

func GenerateHallticketNumberrPS(ctx context.Context, client *ent.Client, examYear string, examCode int32, nodalOfficerFacilityID string) (string, int32, string, bool, error) {

	// Retrieve the last hall ticket number and extract its last four digits

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
	lastFourDigitsMap := make(map[int]bool)
	lastHallTicketNumber, err := tx.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			exam_applications_ps.GenerateHallTicketFlagEQ(true),
			exam_applications_ps.HallTicketNumberNEQ(""),
			exam_applications_ps.StatusEQ("active"),
		).
		Order(ent.Desc(exam_applications_ps.FieldHallTicketNumber)).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			lastHallTicketNumber = nil
		} else {
			return "", 500, " -STR001", false, err
		}
	}

	if lastHallTicketNumber == nil {
		lastHallTicketNumber = &ent.Exam_Applications_PS{HallTicketNumber: "100000000"} // Assuming ExamApplicationsIP struct
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
	applications, err := tx.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			exam_applications_ps.GenerateHallTicketFlagEQ(true),
			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ps.StatusEQ("active"),
			exam_applications_ps.Or(
				exam_applications_ps.HallTicketNumberEQ(""),
			),
		).
		Order(ent.Asc(exam_applications_ps.FieldTempHallTicket)).
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

	// Return success message
	return fmt.Sprintf("Generated hall tickets successfully for %d eligible candidates", successCount), 200, "", true, nil
}

func UpdateCenterCodeForApplicationsPS(ctx context.Context, client *ent.Client, controllingOfficeFacilityID string, examCenterID, seatsToAllot, examCityID int32) (int, []*ent.Exam_Applications_PS, int32, string, bool, error) {
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

	applications, err := tx.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.ExamCityCenterCodeIsNil(),
			exam_applications_ps.CenterIdEQ(examCityID),
			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ps.ControllingOfficeFacilityIDEQ(controllingOfficeFacilityID),
			exam_applications_ps.HallTicketNumberNEQ(""),
			exam_applications_ps.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_ps.FieldApplnSubmittedDate)).
		Limit(int(seatsToAllot)).
		All(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR002", false, err
	}
	var updatedCount int

	for _, application := range applications {
		_, err := tx.Exam_Applications_PS.
			UpdateOne(application).
			SetExamCityCenterCode(examCenterID).
			Save(ctx)

		if err != nil {
			return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR003", false, err
		}
		updatedCount++
	}

	updateCount, err := tx.Exam_Applications_PS.
		Query().
		Where(exam_applications_ps.ExamCityCenterCodeEQ(examCenterID)).
		Count(ctx)

	if err != nil {
		fmt.Println("im here")
		//return 0, err

		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR004", false, err
	}

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

	return updatedCount, applications, 200, "", true, nil
}
