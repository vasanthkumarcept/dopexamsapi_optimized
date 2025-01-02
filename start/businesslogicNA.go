package start

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent"

	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/exam_applications_gdspa"
	"recruit/ent/exam_applications_gdspm"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_pmpa"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/recommendationsgdspaapplications"
	"recruit/ent/recommendationsgdspmapplications"
	"recruit/ent/recommendationsipapplications"
	"recruit/ent/recommendationsmtspmmgapplications"
	"recruit/ent/recommendationspmpaapplications"
	"recruit/ent/recommendationspsapplications"
	ca_reg "recruit/payloadstructure/candidate_registration"

	"time"
)

func validateInput(applicationRecord *ca_reg.NAVerifyApplicationIp) error {
	if applicationRecord == nil {
		return errors.New("payload received in empty")
	}
	if applicationRecord.EmployeeID == 0 {
		return errors.New("employee id should not be empty")
	}
	return nil
}
func validatePsInput(applicationRecord *ca_reg.NAVerifyApplicationGroupB) error {
	if applicationRecord == nil {
		return errors.New("payload received in empty")
	}
	if applicationRecord.EmployeeID == 0 {
		return errors.New("employee id should not be empty")
	}
	return nil
}
func validateMtspmInput(applicationRecord *ca_reg.NAVerifyApplicationMTSPM) error {
	if applicationRecord == nil {
		return errors.New("payload received in empty")
	}
	if applicationRecord.EmployeeID == 0 {
		return errors.New("employee id should not be empty")
	}
	return nil
}

func validatePmpaInput(applicationRecord *ca_reg.NAVerifyApplicationPMPA) error {
	if applicationRecord == nil {
		return errors.New("payload received in empty")
	}
	if applicationRecord.EmployeeID == 0 {
		return errors.New("employee id should not be empty")
	}
	return nil
}

func validateGdspaInput(applicationRecord *ca_reg.NAVerifyApplicationGDStoPA) error {
	if applicationRecord == nil {
		return errors.New("payload received in empty")
	}
	if applicationRecord.EmployeeID == 0 {
		return errors.New("employee id should not be empty")
	}
	return nil
}
func validateGdspmInput(applicationRecord *ca_reg.NAVerifyApplicationGDSPM) error {
	if applicationRecord == nil {
		return errors.New("payload received in empty")
	}
	if applicationRecord.EmployeeID == 0 {
		return errors.New("employee id should not be empty")
	}
	return nil
}
func checkApplicationExists(tx *ent.Tx, ctx context.Context, applicationRecord *ca_reg.NAVerifyApplicationIp) (bool, int32, string, error) {
	exists, err := tx.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(applicationRecord.EmployeeID),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamYearEQ(applicationRecord.ExamYear),
		).
		Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, 422, " -SUB001", fmt.Errorf("no active application available for employee ID %d", applicationRecord.EmployeeID)
		}
		return false, 500, " -SUB002", err
	}
	return exists, 200, " ", nil
}
func checkPsApplicationExists(tx *ent.Tx, ctx context.Context, applicationRecord *ca_reg.NAVerifyApplicationGroupB) (bool, int32, string, error) {
	exists, err := tx.Exam_Applications_PS.Query().
		Where(
			exam_applications_ps.EmployeeIDEQ(applicationRecord.EmployeeID),
			exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ps.StatusEQ("active"),
			exam_applications_ps.ExamYearEQ(applicationRecord.ExamYear),
		).
		Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, 422, " -SUB001", fmt.Errorf("no active application available for employee ID %d", applicationRecord.EmployeeID)
		}
		return false, 500, " -SUB002", err
	}
	return exists, 200, " ", nil
}
func checkMtspmApplicationExists(tx *ent.Tx, ctx context.Context, applicationRecord *ca_reg.NAVerifyApplicationMTSPM) (bool, int32, string, error) {
	exists, err := tx.Exam_Application_MTSPMMG.Query().
		Where(
			exam_application_mtspmmg.EmployeeIDEQ(applicationRecord.EmployeeID),
			exam_application_mtspmmg.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_application_mtspmmg.StatusEQ("active"),
			exam_application_mtspmmg.ExamYearEQ(applicationRecord.ExamYear),
		).
		Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, 422, " -SUB001", fmt.Errorf("no active application available for employee ID %d", applicationRecord.EmployeeID)
		}
		return false, 500, " -SUB002", err
	}
	return exists, 200, " ", nil
}
func checkPmpaApplicationExists(tx *ent.Tx, ctx context.Context, applicationRecord *ca_reg.NAVerifyApplicationPMPA) (bool, int32, string, error) {
	exists, err := tx.Exam_Applications_PMPA.Query().
		Where(
			exam_applications_pmpa.EmployeeIDEQ(applicationRecord.EmployeeID),
			exam_applications_pmpa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_pmpa.StatusEQ("active"),
			exam_applications_pmpa.ExamYearEQ(applicationRecord.ExamYear),
		).
		Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, 422, " -SUB001", fmt.Errorf("no active application available for employee ID %d", applicationRecord.EmployeeID)
		}
		return false, 500, " -SUB002", err
	}
	return exists, 200, " ", nil
}
func checkGdspaApplicationExists(tx *ent.Tx, ctx context.Context, applicationRecord *ca_reg.NAVerifyApplicationGDStoPA) (bool, int32, string, error) {
	exists, err := tx.Exam_Applications_GDSPA.Query().
		Where(
			exam_applications_gdspa.EmployeeIDEQ(applicationRecord.EmployeeID),
			exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.ExamYearEQ(applicationRecord.ExamYear),
		).
		Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, 422, " -SUB001", fmt.Errorf("no active application available for employee ID %d", applicationRecord.EmployeeID)
		}
		return false, 500, " -SUB002", err
	}
	return exists, 200, " ", nil
}
func checkGdspmApplicationExists(tx *ent.Tx, ctx context.Context, applicationRecord *ca_reg.NAVerifyApplicationGDSPM) (bool, int32, string, error) {
	exists, err := tx.Exam_Applications_GDSPM.Query().
		Where(
			exam_applications_gdspm.EmployeeIDEQ(applicationRecord.EmployeeID),
			exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_gdspm.StatusEQ("active"),
			exam_applications_gdspm.ExamYearEQ(applicationRecord.ExamYear),
		).
		Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, 422, " -SUB001", fmt.Errorf("no active application available for employee ID %d", applicationRecord.EmployeeID)
		}
		return false, 500, " -SUB002", err
	}
	return exists, 200, " ", nil
}
func getRecommendationsByEmpID(ctx context.Context, tx *ent.Tx, empID int64) ([]*ent.RecommendationsIPApplications, error) {
	records, err := tx.RecommendationsIPApplications.
		Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying recommendations for employee ID %d: %w", empID, err)
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}
	return records, nil
}
func getPsRecommendationsByEmpID(ctx context.Context, tx *ent.Tx, empID int64) ([]*ent.RecommendationsPSApplications, error) {
	records, err := tx.RecommendationsPSApplications.
		Query().
		Where(recommendationspsapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying recommendations for employee ID %d: %w", empID, err)
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}
	return records, nil
}
func getMtspmRecommendationsByEmpID(ctx context.Context, tx *ent.Tx, empID int64) ([]*ent.RecommendationsMTSPMMGApplications, error) {
	records, err := tx.RecommendationsMTSPMMGApplications.
		Query().
		Where(recommendationsmtspmmgapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying recommendations for employee ID %d: %w", empID, err)
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}
	return records, nil
}
func getPmpaRecommendationsByEmpID(ctx context.Context, tx *ent.Tx, empID int64) ([]*ent.RecommendationsPMPAApplications, error) {
	records, err := tx.RecommendationsPMPAApplications.
		Query().
		Where(recommendationspmpaapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying recommendations for employee ID %d: %w", empID, err)
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}
	return records, nil
}
func getGdspaRecommendationsByEmpID(ctx context.Context, tx *ent.Tx, empID int64) ([]*ent.RecommendationsGDSPAApplications, error) {
	records, err := tx.RecommendationsGDSPAApplications.
		Query().
		Where(recommendationsgdspaapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying recommendations for employee ID %d: %w", empID, err)
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}
	return records, nil
}
func getGdspmRecommendationsByEmpID(ctx context.Context, tx *ent.Tx, empID int64) ([]*ent.RecommendationsGDSPMApplications, error) {
	records, err := tx.RecommendationsGDSPMApplications.
		Query().
		Where(recommendationsgdspmapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying recommendations for employee ID %d: %w", empID, err)
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}
	return records, nil
}
func getActiveExamApplicationIP(ctx context.Context, tx *ent.Tx, empID int64, examYear string) (*ent.Exam_Applications_IP, int32, string, error) {
	updatedRecord, err := tx.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(empID),
			exam_applications_ip.ExamYear(examYear),
			exam_applications_ip.StatusEQ("active")).
		WithCirclePrefRef().
		WithIPApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", fmt.Errorf("no application exists for employee ID %d with exam year %s", empID, examYear)
		}
		return nil, 422, " -SUB002", fmt.Errorf("error querying active application for employee ID %d: %w", empID, err)
	}
	return updatedRecord, 200, " ", nil
}
func getActiveExamApplicationPS(ctx context.Context, tx *ent.Tx, empID int64, examYear string) (*ent.Exam_Applications_PS, int32, string, error) {
	updatedRecord, err := tx.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.EmployeeIDEQ(empID),
			exam_applications_ps.ExamYear(examYear),
			exam_applications_ps.StatusEQ("active")).
		WithPSApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", fmt.Errorf("no application exists for employee ID %d with exam year %s", empID, examYear)
		}
		return nil, 422, " -SUB002", fmt.Errorf("error querying active application for employee ID %d: %w", empID, err)
	}
	return updatedRecord, 200, " ", nil
}
func getActiveExamApplicationMtspm(ctx context.Context, tx *ent.Tx, empID int64, examYear string) (*ent.Exam_Application_MTSPMMG, int32, string, error) {
	updatedRecord, err := tx.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.EmployeeIDEQ(empID),
			exam_application_mtspmmg.ExamYear(examYear),
			exam_application_mtspmmg.StatusEQ("active")).
		WithMTSPMMGApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", fmt.Errorf("no application exists for employee ID %d with exam year %s", empID, examYear)
		}
		return nil, 422, " -SUB002", fmt.Errorf("error querying active application for employee ID %d: %w", empID, err)
	}
	return updatedRecord, 200, " ", nil
}
func getActiveExamApplicationPmpa(ctx context.Context, tx *ent.Tx, empID int64, examYear string) (*ent.Exam_Applications_PMPA, int32, string, error) {
	updatedRecord, err := tx.Exam_Applications_PMPA.
		Query().
		Where(
			exam_applications_pmpa.EmployeeIDEQ(empID),
			exam_applications_pmpa.ExamYear(examYear),
			exam_applications_pmpa.StatusEQ("active")).
		WithPMPAApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", fmt.Errorf("no application exists for employee ID %d with exam year %s", empID, examYear)
		}
		return nil, 422, " -SUB002", fmt.Errorf("error querying active application for employee ID %d: %w", empID, err)
	}
	return updatedRecord, 200, " ", nil
}
func getActiveExamApplicationGdspa(ctx context.Context, tx *ent.Tx, empID int64, examYear string) (*ent.Exam_Applications_GDSPA, int32, string, error) {
	updatedRecord, err := tx.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.EmployeeIDEQ(empID),
			exam_applications_gdspa.ExamYear(examYear),
			exam_applications_gdspa.StatusEQ("active")).
		WithGDSPAApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", fmt.Errorf("no application exists for employee ID %d with exam year %s", empID, examYear)
		}
		return nil, 422, " -SUB002", fmt.Errorf("error querying active application for employee ID %d: %w", empID, err)
	}
	return updatedRecord, 200, " ", nil
}
func getActiveExamApplicationGdspm(ctx context.Context, tx *ent.Tx, empID int64, examYear string) (*ent.Exam_Applications_GDSPM, int32, string, error) {
	updatedRecord, err := tx.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.EmployeeIDEQ(empID),
			exam_applications_gdspm.ExamYear(examYear),
			exam_applications_gdspm.StatusEQ("active")).
		WithGDSPMApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", fmt.Errorf("no application exists for employee ID %d with exam year %s", empID, examYear)
		}
		return nil, 422, " -SUB002", fmt.Errorf("error querying active application for employee ID %d: %w", empID, err)
	}
	return updatedRecord, 200, " ", nil
}
func checkHallTicketGenerated(applicationRecord *ca_reg.NAVerifyApplicationIp, updatedRecord *ent.Exam_Applications_IP) bool {
	if applicationRecord.GenerateHallTicketFlag {
		return updatedRecord.HallTicketNumber != ""
	}
	return false
}

func checkPsHallTicketGenerated(applicationRecord *ca_reg.NAVerifyApplicationGroupB, updatedRecord *ent.Exam_Applications_PS) bool {
	if applicationRecord.GenerateHallTicketFlag {
		return updatedRecord.HallTicketNumber != ""
	}
	return false
}
func checkMtspmHallTicketGenerated(applicationRecord *ca_reg.NAVerifyApplicationMTSPM, updatedRecord *ent.Exam_Application_MTSPMMG) bool {
	if applicationRecord.GenerateHallTicketFlag {
		return updatedRecord.HallTicketNumber != ""
	}
	return false
}
func checkPmpaHallTicketGenerated(applicationRecord *ca_reg.NAVerifyApplicationPMPA, updatedRecord *ent.Exam_Applications_PMPA) bool {
	if applicationRecord.GenerateHallTicketFlag {
		return updatedRecord.HallTicketNumber != ""
	}
	return false
}
func checkGdspaHallTicketGenerated(applicationRecord *ca_reg.NAVerifyApplicationGDStoPA, updatedRecord *ent.Exam_Applications_GDSPA) bool {
	if applicationRecord.GenerateHallTicketFlag {
		return updatedRecord.HallTicketNumber != ""
	}
	return false
}
func checkGdspmHallTicketGenerated(applicationRecord *ca_reg.NAVerifyApplicationGDSPM, updatedRecord *ent.Exam_Applications_GDSPM) bool {
	if applicationRecord.GenerateHallTicketFlag {
		return updatedRecord.HallTicketNumber != ""
	}
	return false
}
func updateRecordEdges(updatedRecord *ent.Exam_Applications_IP) (*ent.Exam_Applications_IP, error) {
	// Extract only the desired fields from the CirclePrefRef edge
	var circlePrefs []*ent.PlaceOfPreferenceIP
	for _, edge := range updatedRecord.Edges.CirclePrefRef {
		circlePrefs = append(circlePrefs, &ent.PlaceOfPreferenceIP{
			PlacePrefNo:    edge.PlacePrefNo,
			PlacePrefValue: edge.PlacePrefValue,
		})
	}

	// Update the CirclePrefRef edge with the filtered values
	updatedRecord.Edges.CirclePrefRef = circlePrefs

	// Extract and update IPApplicationsRef edges
	var recomondPref []*ent.RecommendationsIPApplications
	for _, edge := range updatedRecord.Edges.IPApplicationsRef {
		recomondPref = append(recomondPref, &ent.RecommendationsIPApplications{
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
	updatedRecord.Edges.IPApplicationsRef = recomondPref
	updatedRecord.UpdatedAt = updatedRecord.UpdatedAt.Truncate(24 * time.Hour)

	return updatedRecord, nil
}

func createUpdatedAppln(tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationIp, updatedRecord *ent.Exam_Applications_IP, hallticketgeneratedflag bool, ctx context.Context) (*ent.Exam_Applications_IP, error) {

	// Create a new Exam_Applications_IP record with the provided values

	currentTime := time.Now().Truncate(time.Second)
	updatedAppln, err := tx.Exam_Applications_IP.
		Create().
		SetApplicationStatus("VerifiedByNA").
		SetCARemarks(updatedRecord.CARemarks).
		SetCAUserName(updatedRecord.CAUserName).
		SetCADate(currentTime).
		SetCAGeneralRemarks(updatedRecord.CAGeneralRemarks).
		SetStatus("active").
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetCADate(currentTime).
		SetUserID(updatedRecord.UserID).
		SetExamName(updatedRecord.ExamName).
		SetEmployeeID(updatedRecord.EmployeeID).
		SetEmployeeName(updatedRecord.EmployeeName).
		SetCAEmployeeID(updatedRecord.CAEmployeeID).
		SetGenerateHallTicketFlag(applicationRecord.GenerateHallTicketFlag).
		SetHallTicketNumber(updatedRecord.HallTicketNumber).
		SetCAEmployeeDesignation(updatedRecord.CAEmployeeDesignation).
		SetDOB(updatedRecord.DOB).
		SetGender(updatedRecord.Gender).
		SetMobileNumber(updatedRecord.MobileNumber).
		SetEmailID(updatedRecord.EmailID).
		SetCategoryCode(updatedRecord.CategoryCode).
		SetCategoryDescription(updatedRecord.CategoryDescription).
		SetCadre(updatedRecord.Cadre).
		SetEmployeePost(updatedRecord.EmployeePost).
		SetWorkingOfficeFacilityID(updatedRecord.WorkingOfficeFacilityID).
		SetWorkingOfficeCircleFacilityID(updatedRecord.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(updatedRecord.WorkingOfficeCircleName).
		SetWorkingOfficeRegionFacilityID(updatedRecord.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(updatedRecord.WorkingOfficeRegionName).
		SetWorkingOfficeDivisionFacilityID(updatedRecord.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(updatedRecord.WorkingOfficeDivisionName).
		SetReportingOfficeName(updatedRecord.ReportingOfficeName).
		SetReportingOfficeFacilityID(updatedRecord.ReportingOfficeFacilityID).
		SetEntryPostDescription(updatedRecord.EntryPostDescription).
		SetPresentPostDescription(updatedRecord.PresentPostDescription).
		SetPresentDesignation(updatedRecord.PresentDesignation).
		SetFeederPostDescription(updatedRecord.FeederPostDescription).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetServiceLength(updatedRecord.ServiceLength).
		SetCandidateRemarks(updatedRecord.CandidateRemarks).
		SetRemarks(updatedRecord.Remarks).
		SetDCCS(updatedRecord.DCCS).
		SetDCInPresentCadre(updatedRecord.DCInPresentCadre).
		SetDeputationOfficeFacilityID(updatedRecord.DeputationOfficeFacilityID).
		SetDisabilityTypeID(updatedRecord.DisabilityTypeID).
		SetDisabilityTypeCode(updatedRecord.DisabilityTypeCode).
		SetEntryPostCode(updatedRecord.EntryPostCode).
		SetPresentPostCode(updatedRecord.PresentPostCode).
		SetFeederPostCode(updatedRecord.FeederPostCode).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetDesignationID(updatedRecord.DesignationID).
		SetEducationCode(updatedRecord.EducationCode).
		SetFacilityUniqueID(updatedRecord.FacilityUniqueID).
		SetLienControllingOfficeID(updatedRecord.LienControllingOfficeID).
		SetLienControllingOfficeName(updatedRecord.LienControllingOfficeName).
		SetInDeputation(updatedRecord.InDeputation).
		SetDeputationType(updatedRecord.DeputationType).
		SetDeputationOfficeUniqueId(updatedRecord.DeputationOfficeUniqueId).
		SetDeputationOfficeName(updatedRecord.DeputationOfficeName).
		SetDeputationControllingOfficeID(updatedRecord.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(updatedRecord.DeputationControllingOfficeName).
		SetControllingOfficeFacilityID(updatedRecord.ControllingOfficeFacilityID).
		SetControllingOfficeName(updatedRecord.ControllingOfficeName).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetNodalOfficeName(updatedRecord.NodalOfficeName).
		SetCenterFacilityId(updatedRecord.CenterFacilityId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetPhotoPath(updatedRecord.PhotoPath).
		SetNonQualifyingService(updatedRecord.NonQualifyingService).
		SetSignaturePath(updatedRecord.SignaturePath).
		SetTempHallTicket(updatedRecord.TempHallTicket).
		SetDisabilityTypeDescription(updatedRecord.DisabilityTypeDescription).
		SetDisabilityPercentage(updatedRecord.DisabilityPercentage).
		SetEducationDescription(updatedRecord.EducationDescription).
		SetExamCode(updatedRecord.ExamCode).
		SetExamShortName(updatedRecord.ExamShortName).
		SetExamYear(updatedRecord.ExamYear).
		SetExamCityCenterCode(updatedRecord.CenterId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetSignature(updatedRecord.Signature).
		SetPhoto(updatedRecord.Photo).
		SetNAEmployeeID(applicationRecord.NA_EmployeeID).
		SetNAEmployeeDesignation(applicationRecord.NA_EmployeeDesignation).
		SetNAUserName(applicationRecord.NA_UserName).
		SetApplicationNumber(updatedRecord.ApplicationNumber).
		SetApplnSubmittedDate(updatedRecord.ApplnSubmittedDate).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetWorkingOfficePincode(updatedRecord.WorkingOfficePincode).
		SetWorkingOfficeName(updatedRecord.WorkingOfficeName).
		SetOptionUsed(updatedRecord.OptionUsed).
		SetRemarks(updatedRecord.Remarks).
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetCenterId(updatedRecord.CenterId).
		SetClaimingQualifyingService(updatedRecord.ClaimingQualifyingService).
		SetDeputationOfficePincode(updatedRecord.DeputationOfficePincode).
		SetNADate(currentTime).
		SetRecommendedStatus(applicationRecord.RecommendedStatus).
		SetPunishmentStatus(updatedRecord.PunishmentStatus).
		SetDisciplinaryCaseStatus(updatedRecord.DisciplinaryCaseStatus).
		SetHallTicketGeneratedFlag(hallticketgeneratedflag).
		SetHallIdentificationNumber(updatedRecord.HallIdentificationNumber).
		SetHallTicketGeneratedDate(updatedRecord.HallTicketGeneratedDate).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createUpdatedPsAppln(tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationGroupB, updatedRecord *ent.Exam_Applications_PS, hallticketgeneratedflag bool, ctx context.Context) (*ent.Exam_Applications_PS, error) {

	// Create a new Exam_Applications_IP record with the provided values

	currentTime := time.Now().Truncate(time.Second)
	updatedAppln, err := tx.Exam_Applications_PS.
		Create().
		SetApplicationStatus("VerifiedByNA").
		SetCARemarks(updatedRecord.CARemarks).
		SetCAUserName(updatedRecord.CAUserName).
		SetGenerateHallTicketFlag(applicationRecord.GenerateHallTicketFlag).
		SetCADate(currentTime).
		SetCAGeneralRemarks(updatedRecord.CAGeneralRemarks).
		SetStatus("active").
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetRecommendedStatus(applicationRecord.RecommendedStatus).
		SetCADate(currentTime).
		SetUserID(updatedRecord.UserID).
		SetExamName(updatedRecord.ExamName).
		SetEmployeeID(updatedRecord.EmployeeID).
		SetEmployeeName(updatedRecord.EmployeeName).
		SetNAEmployeeID(applicationRecord.NA_EmployeeID).
		SetNAEmployeeDesignation(applicationRecord.NA_EmployeeDesignation).
		SetNAUserName(applicationRecord.NA_UserName).
		SetCAEmployeeID(updatedRecord.CAEmployeeID).
		SetCAEmployeeDesignation(updatedRecord.CAEmployeeDesignation).
		SetDOB(updatedRecord.DOB).
		SetCenterId(updatedRecord.CenterId).
		SetClaimingQualifyingService(updatedRecord.ClaimingQualifyingService).
		SetDeputationOfficePincode(updatedRecord.DeputationOfficePincode).
		SetGender(updatedRecord.Gender).
		SetMobileNumber(updatedRecord.MobileNumber).
		SetEmailID(updatedRecord.EmailID).
		SetCategoryCode(updatedRecord.CategoryCode).
		SetCategoryDescription(updatedRecord.CategoryDescription).
		SetCadre(updatedRecord.Cadre).
		SetEmployeePost(updatedRecord.EmployeePost).
		SetWorkingOfficeFacilityID(updatedRecord.WorkingOfficeFacilityID).
		SetWorkingOfficeCircleFacilityID(updatedRecord.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(updatedRecord.WorkingOfficeCircleName).
		SetWorkingOfficeRegionFacilityID(updatedRecord.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(updatedRecord.WorkingOfficeRegionName).
		SetWorkingOfficeDivisionFacilityID(updatedRecord.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(updatedRecord.WorkingOfficeDivisionName).
		SetReportingOfficeName(updatedRecord.ReportingOfficeName).
		SetReportingOfficeFacilityID(updatedRecord.ReportingOfficeFacilityID).
		SetEntryPostDescription(updatedRecord.EntryPostDescription).
		SetPresentPostDescription(updatedRecord.PresentPostDescription).
		SetPresentDesignation(updatedRecord.PresentDesignation).
		SetFeederPostDescription(updatedRecord.FeederPostDescription).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetServiceLength(updatedRecord.ServiceLength).
		SetCandidateRemarks(updatedRecord.CandidateRemarks).
		SetRemarks(updatedRecord.Remarks).
		SetDCCS(updatedRecord.DCCS).
		SetDCInPresentCadre(updatedRecord.DCInPresentCadre).
		SetDeputationOfficeFacilityID(updatedRecord.DeputationOfficeFacilityID).
		SetDisabilityTypeID(updatedRecord.DisabilityTypeID).
		SetDisabilityTypeCode(updatedRecord.DisabilityTypeCode).
		SetEntryPostCode(updatedRecord.EntryPostCode).
		SetPresentPostCode(updatedRecord.PresentPostCode).
		SetFeederPostCode(updatedRecord.FeederPostCode).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetDesignationID(updatedRecord.DesignationID).
		SetEducationCode(updatedRecord.EducationCode).
		SetFacilityUniqueID(updatedRecord.FacilityUniqueID).
		SetLienControllingOfficeID(updatedRecord.LienControllingOfficeID).
		SetLienControllingOfficeName(updatedRecord.LienControllingOfficeName).
		SetInDeputation(updatedRecord.InDeputation).
		SetDeputationType(updatedRecord.DeputationType).
		SetDeputationOfficeUniqueId(updatedRecord.DeputationOfficeUniqueId).
		SetDeputationOfficeName(updatedRecord.DeputationOfficeName).
		SetDeputationControllingOfficeID(updatedRecord.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(updatedRecord.DeputationControllingOfficeName).
		SetControllingOfficeFacilityID(updatedRecord.ControllingOfficeFacilityID).
		SetControllingOfficeName(updatedRecord.ControllingOfficeName).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetNodalOfficeName(updatedRecord.NodalOfficeName).
		SetCenterFacilityId(updatedRecord.CenterFacilityId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetPhotoPath(updatedRecord.PhotoPath).
		SetNonQualifyingService(updatedRecord.NonQualifyingService).
		SetSignaturePath(updatedRecord.SignaturePath).
		SetTempHallTicket(updatedRecord.TempHallTicket).
		SetDisabilityTypeDescription(updatedRecord.DisabilityTypeDescription).
		SetDisabilityPercentage(updatedRecord.DisabilityPercentage).
		SetEducationDescription(updatedRecord.EducationDescription).
		SetExamCode(updatedRecord.ExamCode).
		SetExamShortName(updatedRecord.ExamShortName).
		SetExamYear(updatedRecord.ExamYear).
		SetExamCityCenterCode(updatedRecord.CenterId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetSignature(updatedRecord.Signature).
		SetPhoto(updatedRecord.Photo).
		SetApplicationNumber(updatedRecord.ApplicationNumber).
		SetApplnSubmittedDate(updatedRecord.ApplnSubmittedDate).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetWorkingOfficePincode(updatedRecord.WorkingOfficePincode).
		SetWorkingOfficeName(updatedRecord.WorkingOfficeName).
		SetOptionUsed(updatedRecord.OptionUsed).
		SetRemarks(updatedRecord.Remarks).
		SetNAUserName(applicationRecord.NA_UserName).
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetPunishmentStatus(updatedRecord.PunishmentStatus).
		SetDisciplinaryCaseStatus(updatedRecord.DisciplinaryCaseStatus).
		SetNADate(currentTime).
		SetHallTicketGeneratedFlag(hallticketgeneratedflag).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createUpdatedMtspmAppln(tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationMTSPM, updatedRecord *ent.Exam_Application_MTSPMMG, hallticketgeneratedflag bool, ctx context.Context) (*ent.Exam_Application_MTSPMMG, error) {

	// Create a new Exam_Applications_IP record with the provided values

	currentTime := time.Now().Truncate(time.Second)
	updatedAppln,err := tx.Exam_Application_MTSPMMG.
		Create().
		SetApplicationStatus("VerifiedByNA").
		SetCARemarks(updatedRecord.CARemarks).
		SetCAUserName(updatedRecord.CAUserName).
		SetCADate(currentTime).
		SetCAGeneralRemarks(updatedRecord.CAGeneralRemarks).
		SetStatus("active").
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetCADate(currentTime).
		SetRecommendedStatus(applicationRecord.RecommendedStatus).
		SetUserID(updatedRecord.UserID).
		SetExamName(updatedRecord.ExamName).
		SetEmployeeID(updatedRecord.EmployeeID).
		SetGenerateHallTicketFlag(applicationRecord.GenerateHallTicketFlag).
		SetHallTicketNumber(updatedRecord.HallTicketNumber).
		SetEmployeeName(updatedRecord.EmployeeName).
		SetCAEmployeeID(updatedRecord.CAEmployeeID).
		SetPostPreferences(updatedRecord.PostPreferences).
		SetUnitPreferences(updatedRecord.UnitPreferences).
		SetCAEmployeeDesignation(updatedRecord.CAEmployeeDesignation).
		SetDOB(updatedRecord.DOB).
		SetNAEmployeeID(applicationRecord.NA_EmployeeID).
		SetNAEmployeeDesignation(applicationRecord.NA_EmployeeDesignation).
		SetNAUserName(applicationRecord.NA_UserName).
		SetGender(updatedRecord.Gender).
		SetMobileNumber(updatedRecord.MobileNumber).
		SetEmailID(updatedRecord.EmailID).
		SetCategoryCode(updatedRecord.CategoryCode).
		SetCategoryDescription(updatedRecord.CategoryDescription).
		SetCadre(updatedRecord.Cadre).
		SetCenterId(updatedRecord.CenterId).
		SetClaimingQualifyingService(updatedRecord.ClaimingQualifyingService).
		SetDeputationOfficePincode(updatedRecord.DeputationOfficePincode).
		SetEmployeePost(updatedRecord.EmployeePost).
		SetWorkingOfficeFacilityID(updatedRecord.WorkingOfficeFacilityID).
		SetWorkingOfficeCircleFacilityID(updatedRecord.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(updatedRecord.WorkingOfficeCircleName).
		SetWorkingOfficeRegionFacilityID(updatedRecord.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(updatedRecord.WorkingOfficeRegionName).
		SetWorkingOfficeDivisionFacilityID(updatedRecord.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(updatedRecord.WorkingOfficeDivisionName).
		SetReportingOfficeName(updatedRecord.ReportingOfficeName).
		SetReportingOfficeFacilityID(updatedRecord.ReportingOfficeFacilityID).
		SetEntryPostDescription(updatedRecord.EntryPostDescription).
		SetPresentPostDescription(updatedRecord.PresentPostDescription).
		SetPresentDesignation(updatedRecord.PresentDesignation).
		SetFeederPostDescription(updatedRecord.FeederPostDescription).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetServiceLength(updatedRecord.ServiceLength).
		SetCandidateRemarks(updatedRecord.CandidateRemarks).
		SetRemarks(updatedRecord.Remarks).
		SetDCCS(updatedRecord.DCCS).
		SetDCInPresentCadre(updatedRecord.DCInPresentCadre).
		SetDeputationOfficeFacilityID(updatedRecord.DeputationOfficeFacilityID).
		SetDisabilityTypeID(updatedRecord.DisabilityTypeID).
		SetDisabilityTypeCode(updatedRecord.DisabilityTypeCode).
		SetEntryPostCode(updatedRecord.EntryPostCode).
		SetPresentPostCode(updatedRecord.PresentPostCode).
		SetFeederPostCode(updatedRecord.FeederPostCode).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetDesignationID(updatedRecord.DesignationID).
		SetEducationCode(updatedRecord.EducationCode).
		SetFacilityUniqueID(updatedRecord.FacilityUniqueID).
		SetLienControllingOfficeID(updatedRecord.LienControllingOfficeID).
		SetLienControllingOfficeName(updatedRecord.LienControllingOfficeName).
		SetInDeputation(updatedRecord.InDeputation).
		SetDeputationType(updatedRecord.DeputationType).
		SetDeputationOfficeUniqueId(updatedRecord.DeputationOfficeUniqueId).
		SetDeputationOfficeName(updatedRecord.DeputationOfficeName).
		SetDeputationControllingOfficeID(updatedRecord.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(updatedRecord.DeputationControllingOfficeName).
		SetControllingOfficeFacilityID(updatedRecord.ControllingOfficeFacilityID).
		SetControllingOfficeName(updatedRecord.ControllingOfficeName).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetNodalOfficeName(updatedRecord.NodalOfficeName).
		SetCenterFacilityId(updatedRecord.CenterFacilityId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetPhotoPath(updatedRecord.PhotoPath).
		SetNonQualifyingService(updatedRecord.NonQualifyingService).
		SetSignaturePath(updatedRecord.SignaturePath).
		SetTempHallTicket(updatedRecord.TempHallTicket).
		SetDisabilityTypeDescription(updatedRecord.DisabilityTypeDescription).
		SetDisabilityPercentage(updatedRecord.DisabilityPercentage).
		SetEducationDescription(updatedRecord.EducationDescription).
		SetExamCode(updatedRecord.ExamCode).
		SetExamShortName(updatedRecord.ExamShortName).
		SetExamYear(updatedRecord.ExamYear).
		SetExamCityCenterCode(updatedRecord.CenterId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetSignature(updatedRecord.Signature).
		SetPhoto(updatedRecord.Photo).
		SetApplicationNumber(updatedRecord.ApplicationNumber).
		SetApplnSubmittedDate(updatedRecord.ApplnSubmittedDate).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetWorkingOfficePincode(updatedRecord.WorkingOfficePincode).
		SetWorkingOfficeName(updatedRecord.WorkingOfficeName).
		SetOptionUsed(updatedRecord.OptionUsed).
		SetRemarks(updatedRecord.Remarks).
		SetNAUserName(applicationRecord.NA_UserName).
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetNADate(currentTime).
		SetPunishmentStatus(updatedRecord.PunishmentStatus).             //here added punishmentstatus
		SetDisciplinaryCaseStatus(updatedRecord.DisciplinaryCaseStatus). //here added DisciplinaryCaseStatus
		SetHallTicketGeneratedFlag(hallticketgeneratedflag).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createUpdatedGdspaAppln(tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationGDStoPA, updatedRecord *ent.Exam_Applications_GDSPA, hallticketgeneratedflag bool, ctx context.Context) (*ent.Exam_Applications_GDSPA, error) {

	// Create a new Exam_Applications_IP record with the provided values

	currentTime := time.Now().Truncate(time.Second)
	updatedAppln,err := tx.Exam_Applications_GDSPA.
		Create().
		SetApplicationStatus("VerifiedByNA").
		SetCARemarks(updatedRecord.CARemarks).
		SetCAUserName(updatedRecord.CAUserName).
		SetCADate(currentTime).
		SetCAGeneralRemarks(updatedRecord.CAGeneralRemarks).
		SetStatus("active").
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetCADate(currentTime).
		SetUserID(updatedRecord.UserID).
		SetExamName(updatedRecord.ExamName).
		SetEmployeeID(updatedRecord.EmployeeID).
		SetGenerateHallTicketFlag(applicationRecord.GenerateHallTicketFlag).
		SetHallTicketNumber(updatedRecord.HallTicketNumber).
		SetEmployeeName(updatedRecord.EmployeeName).
		SetCAEmployeeID(updatedRecord.CAEmployeeID).
		SetNAEmployeeID(applicationRecord.NA_EmployeeID).
		SetNAEmployeeDesignation(applicationRecord.NA_EmployeeDesignation).
		SetNAUserName(applicationRecord.NA_UserName).
		SetCAEmployeeDesignation(updatedRecord.CAEmployeeDesignation).
		SetDOB(updatedRecord.DOB).
		SetCenterId(updatedRecord.CenterId).
		SetClaimingQualifyingService(updatedRecord.ClaimingQualifyingService).
		SetDeputationOfficePincode(updatedRecord.DeputationOfficePincode).
		SetGender(updatedRecord.Gender).
		SetMobileNumber(updatedRecord.MobileNumber).
		SetPostPreferences(updatedRecord.PostPreferences).
		SetUnitPreferences(updatedRecord.UnitPreferences).
		SetEmailID(updatedRecord.EmailID).
		SetCategoryCode(updatedRecord.CategoryCode).
		SetCategoryDescription(updatedRecord.CategoryDescription).
		SetCadre(updatedRecord.Cadre).
		SetPunishmentStatus(updatedRecord.PunishmentStatus). //here added punishmentstatus and DisciplinaryCaseStatus
		SetDisciplinaryCaseStatus(updatedRecord.DisciplinaryCaseStatus).
		SetEmployeePost(updatedRecord.EmployeePost).
		SetWorkingOfficeFacilityID(updatedRecord.WorkingOfficeFacilityID).
		SetWorkingOfficeCircleFacilityID(updatedRecord.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(updatedRecord.WorkingOfficeCircleName).
		SetWorkingOfficeRegionFacilityID(updatedRecord.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(updatedRecord.WorkingOfficeRegionName).
		SetWorkingOfficeDivisionFacilityID(updatedRecord.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(updatedRecord.WorkingOfficeDivisionName).
		SetReportingOfficeName(updatedRecord.ReportingOfficeName).
		SetReportingOfficeFacilityID(updatedRecord.ReportingOfficeFacilityID).
		SetEntryPostDescription(updatedRecord.EntryPostDescription).
		SetPresentPostDescription(updatedRecord.PresentPostDescription).
		SetPresentDesignation(updatedRecord.PresentDesignation).
		SetFeederPostDescription(updatedRecord.FeederPostDescription).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetServiceLength(updatedRecord.ServiceLength).
		SetCandidateRemarks(updatedRecord.CandidateRemarks).
		SetDCCS(updatedRecord.DCCS).
		SetDeputationOfficeFacilityID(updatedRecord.DeputationOfficeFacilityID).
		SetDisabilityTypeID(updatedRecord.DisabilityTypeID).
		SetDisabilityTypeCode(updatedRecord.DisabilityTypeCode).
		SetEntryPostCode(updatedRecord.EntryPostCode).
		SetPresentPostCode(updatedRecord.PresentPostCode).
		SetFeederPostCode(updatedRecord.FeederPostCode).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetDesignationID(updatedRecord.DesignationID).
		SetEducationCode(updatedRecord.EducationCode).
		SetFacilityUniqueID(updatedRecord.FacilityUniqueID).
		SetLienControllingOfficeID(updatedRecord.LienControllingOfficeID).
		SetLienControllingOfficeName(updatedRecord.LienControllingOfficeName).
		SetInDeputation(updatedRecord.InDeputation).
		SetDeputationType(updatedRecord.DeputationType).
		SetDeputationOfficeUniqueId(updatedRecord.DeputationOfficeUniqueId).
		SetDeputationOfficeName(updatedRecord.DeputationOfficeName).
		SetDeputationControllingOfficeID(updatedRecord.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(updatedRecord.DeputationControllingOfficeName).
		SetControllingOfficeFacilityID(updatedRecord.ControllingOfficeFacilityID).
		SetControllingOfficeName(updatedRecord.ControllingOfficeName).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetNodalOfficeName(updatedRecord.NodalOfficeName).
		SetCenterFacilityId(updatedRecord.CenterFacilityId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetPhotoPath(updatedRecord.PhotoPath).
		SetNonQualifyingService(updatedRecord.NonQualifyingService).
		SetSignaturePath(updatedRecord.SignaturePath).
		SetSubdivisionOfficeFacilityID(updatedRecord.SubdivisionOfficeFacilityID).
		SetSubdivisionOfficeName(updatedRecord.SubdivisionOfficeName).
		SetTempHallTicket(updatedRecord.TempHallTicket).
		SetDisabilityTypeDescription(updatedRecord.DisabilityTypeDescription).
		SetDisabilityPercentage(updatedRecord.DisabilityPercentage).
		SetEducationDescription(updatedRecord.EducationDescription).
		SetExamCode(updatedRecord.ExamCode).
		SetExamShortName(updatedRecord.ExamShortName).
		SetExamYear(updatedRecord.ExamYear).
		SetExamCityCenterCode(updatedRecord.CenterId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetSignature(updatedRecord.Signature).
		SetPhoto(updatedRecord.Photo).
		SetApplicationNumber(updatedRecord.ApplicationNumber).
		SetApplnSubmittedDate(updatedRecord.ApplnSubmittedDate).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetWorkingOfficePincode(updatedRecord.WorkingOfficePincode).
		SetWorkingOfficeName(updatedRecord.WorkingOfficeName).
		SetOptionUsed(updatedRecord.OptionUsed).
		SetNAUserName(applicationRecord.NA_UserName).
		SetRecommendedStatus(applicationRecord.RecommendedStatus).
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetNADate(currentTime).
		SetHallTicketGeneratedFlag(hallticketgeneratedflag).
		//	SetGenerateHallTicketFlag(*updatedRecord.GenerateHallTicketFlag).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createUpdatedPmpaAppln(tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationPMPA, updatedRecord *ent.Exam_Applications_PMPA, hallticketgeneratedflag bool, ctx context.Context) (*ent.Exam_Applications_PMPA, error) {

	// Create a new Exam_Applications_IP record with the provided values

	currentTime := time.Now().Truncate(time.Second)
	updatedAppln,err := tx.Exam_Applications_PMPA.
		Create().
		SetApplicationStatus("VerifiedByNA").
		SetCARemarks(updatedRecord.CARemarks).
		SetCAUserName(updatedRecord.CAUserName).
		SetCADate(currentTime).
		SetCAGeneralRemarks(updatedRecord.CAGeneralRemarks).
		SetStatus("active").
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetCADate(currentTime).
		SetUserID(updatedRecord.UserID).
		SetExamName(updatedRecord.ExamName).
		SetRecommendedStatus(applicationRecord.RecommendedStatus).
		SetEmployeeID(updatedRecord.EmployeeID).
		SetPostPreferences(updatedRecord.PostPreferences).
		SetUnitPreferences(updatedRecord.UnitPreferences).
		SetGenerateHallTicketFlag(applicationRecord.GenerateHallTicketFlag).
		SetHallTicketNumber(updatedRecord.HallTicketNumber).
		SetEmployeeName(updatedRecord.EmployeeName).
		SetCAEmployeeID(updatedRecord.CAEmployeeID).
		SetCAEmployeeDesignation(updatedRecord.CAEmployeeDesignation).
		SetDOB(updatedRecord.DOB).
		SetNAEmployeeID(applicationRecord.NA_EmployeeID).
		SetNAEmployeeDesignation(applicationRecord.NA_EmployeeDesignation).
		SetNAUserName(applicationRecord.NA_UserName).
		SetGender(updatedRecord.Gender).
		SetMobileNumber(updatedRecord.MobileNumber).
		SetEmailID(updatedRecord.EmailID).
		SetCategoryCode(updatedRecord.CategoryCode).
		SetCategoryDescription(updatedRecord.CategoryDescription).
		SetCadre(updatedRecord.Cadre).
		SetCenterId(updatedRecord.CenterId).
		SetClaimingQualifyingService(updatedRecord.ClaimingQualifyingService).
		SetDeputationOfficePincode(updatedRecord.DeputationOfficePincode).
		SetEmployeePost(updatedRecord.EmployeePost).
		SetWorkingOfficeFacilityID(updatedRecord.WorkingOfficeFacilityID).
		SetWorkingOfficeCircleFacilityID(updatedRecord.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(updatedRecord.WorkingOfficeCircleName).
		SetWorkingOfficeRegionFacilityID(updatedRecord.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(updatedRecord.WorkingOfficeRegionName).
		SetWorkingOfficeDivisionFacilityID(updatedRecord.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(updatedRecord.WorkingOfficeDivisionName).
		SetReportingOfficeName(updatedRecord.ReportingOfficeName).
		SetReportingOfficeFacilityID(updatedRecord.ReportingOfficeFacilityID).
		SetEntryPostDescription(updatedRecord.EntryPostDescription).
		SetPresentPostDescription(updatedRecord.PresentPostDescription).
		SetPresentDesignation(updatedRecord.PresentDesignation).
		SetFeederPostDescription(updatedRecord.FeederPostDescription).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetServiceLength(updatedRecord.ServiceLength).
		SetCandidateRemarks(updatedRecord.CandidateRemarks).
		SetRemarks(updatedRecord.Remarks).
		SetDCCS(updatedRecord.DCCS).
		SetDCInPresentCadre(updatedRecord.DCInPresentCadre).
		SetDeputationOfficeFacilityID(updatedRecord.DeputationOfficeFacilityID).
		SetDisabilityTypeID(updatedRecord.DisabilityTypeID).
		SetDisabilityTypeCode(updatedRecord.DisabilityTypeCode).
		SetEntryPostCode(updatedRecord.EntryPostCode).
		SetPresentPostCode(updatedRecord.PresentPostCode).
		SetFeederPostCode(updatedRecord.FeederPostCode).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetDesignationID(updatedRecord.DesignationID).
		SetEducationCode(updatedRecord.EducationCode).
		SetFacilityUniqueID(updatedRecord.FacilityUniqueID).
		SetLienControllingOfficeID(updatedRecord.LienControllingOfficeID).
		SetLienControllingOfficeName(updatedRecord.LienControllingOfficeName).
		SetInDeputation(updatedRecord.InDeputation).
		SetDeputationType(updatedRecord.DeputationType).
		SetDeputationOfficeUniqueId(updatedRecord.DeputationOfficeUniqueId).
		SetDeputationOfficeName(updatedRecord.DeputationOfficeName).
		SetDeputationControllingOfficeID(updatedRecord.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(updatedRecord.DeputationControllingOfficeName).
		SetControllingOfficeFacilityID(updatedRecord.ControllingOfficeFacilityID).
		SetControllingOfficeName(updatedRecord.ControllingOfficeName).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetNodalOfficeName(updatedRecord.NodalOfficeName).
		SetCenterFacilityId(updatedRecord.CenterFacilityId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetPhotoPath(updatedRecord.PhotoPath).
		SetNonQualifyingService(updatedRecord.NonQualifyingService).
		SetSignaturePath(updatedRecord.SignaturePath).
		SetTempHallTicket(updatedRecord.TempHallTicket).
		SetDisabilityTypeDescription(updatedRecord.DisabilityTypeDescription).
		SetDisabilityPercentage(updatedRecord.DisabilityPercentage).
		SetEducationDescription(updatedRecord.EducationDescription).
		SetExamCode(updatedRecord.ExamCode).
		SetExamShortName(updatedRecord.ExamShortName).
		SetExamCityCenterCode(updatedRecord.CenterId).
		SetExamYear(updatedRecord.ExamYear).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetSignature(updatedRecord.Signature).
		SetPhoto(updatedRecord.Photo).
		SetApplicationNumber(updatedRecord.ApplicationNumber).
		SetApplnSubmittedDate(updatedRecord.ApplnSubmittedDate).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetNodalOfficeName(updatedRecord.NodalOfficeName).
		SetWorkingOfficePincode(updatedRecord.WorkingOfficePincode).
		SetWorkingOfficeName(updatedRecord.WorkingOfficeName).
		SetOptionUsed(updatedRecord.OptionUsed).
		SetRemarks(updatedRecord.Remarks).
		SetNAUserName(applicationRecord.NA_UserName).
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetPunishmentStatus(updatedRecord.PunishmentStatus).             //here added punishmentstatus
		SetDisciplinaryCaseStatus(updatedRecord.DisciplinaryCaseStatus). //here added DisciplinaryCaseStatus
		SetHallTicketGeneratedFlag(hallticketgeneratedflag).
		SetNADate(currentTime).
		//SetGenerateHallTicketFlag(*updatedRecord.GenerateHallTicketFlag).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createUpdatedGdspmAppln(tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationGDSPM, updatedRecord *ent.Exam_Applications_GDSPM, hallticketgeneratedflag bool, ctx context.Context) (*ent.Exam_Applications_GDSPM, error) {

	// Create a new Exam_Applications_IP record with the provided values

	currentTime := time.Now().Truncate(time.Second)
	updatedAppln,err := tx.Exam_Applications_GDSPM.
		Create().
		SetApplicationStatus("VerifiedByNA").
		SetCARemarks(updatedRecord.CARemarks).
		SetCAUserName(updatedRecord.CAUserName).
		SetCADate(currentTime).
		SetCAGeneralRemarks(updatedRecord.CAGeneralRemarks).
		SetStatus("active").
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetCADate(currentTime).
		SetUserID(updatedRecord.UserID).
		SetExamName(updatedRecord.ExamName).
		SetEmployeeID(updatedRecord.EmployeeID).
		SetGenerateHallTicketFlag(applicationRecord.GenerateHallTicketFlag).
		SetHallTicketNumber(updatedRecord.HallTicketNumber).
		SetEmployeeName(updatedRecord.EmployeeName).
		SetCAEmployeeID(updatedRecord.CAEmployeeID).
		SetPostPreferences(updatedRecord.PostPreferences).
		SetUnitPreferences(updatedRecord.UnitPreferences).
		SetNAEmployeeID(applicationRecord.NA_EmployeeID).
		SetNAEmployeeDesignation(applicationRecord.NA_EmployeeDesignation).
		SetNAUserName(applicationRecord.NA_UserName).
		SetCAEmployeeDesignation(updatedRecord.CAEmployeeDesignation).
		SetDOB(updatedRecord.DOB).
		SetGender(updatedRecord.Gender).
		SetMobileNumber(updatedRecord.MobileNumber).
		SetEmailID(updatedRecord.EmailID).
		SetCenterId(updatedRecord.CenterId).
		SetClaimingQualifyingService(updatedRecord.ClaimingQualifyingService).
		SetDeputationOfficePincode(updatedRecord.DeputationOfficePincode).
		SetCategoryCode(updatedRecord.CategoryCode).
		SetCategoryDescription(updatedRecord.CategoryDescription).
		SetCadre(updatedRecord.Cadre).
		SetEmployeePost(updatedRecord.EmployeePost).
		SetWorkingOfficeFacilityID(updatedRecord.WorkingOfficeFacilityID).
		SetWorkingOfficeCircleFacilityID(updatedRecord.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(updatedRecord.WorkingOfficeCircleName).
		SetWorkingOfficeRegionFacilityID(updatedRecord.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(updatedRecord.WorkingOfficeRegionName).
		SetWorkingOfficeDivisionFacilityID(updatedRecord.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(updatedRecord.WorkingOfficeDivisionName).
		SetReportingOfficeName(updatedRecord.ReportingOfficeName).
		SetReportingOfficeFacilityID(updatedRecord.ReportingOfficeFacilityID).
		SetEntryPostDescription(updatedRecord.EntryPostDescription).
		SetPresentPostDescription(updatedRecord.PresentPostDescription).
		SetPresentDesignation(updatedRecord.PresentDesignation).
		SetFeederPostDescription(updatedRecord.FeederPostDescription).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetServiceLength(updatedRecord.ServiceLength).
		SetCandidateRemarks(updatedRecord.CandidateRemarks).
		SetRemarks(updatedRecord.Remarks).
		SetDCCS(updatedRecord.DCCS).
		SetDCInPresentCadre(updatedRecord.DCInPresentCadre).
		SetDeputationOfficeFacilityID(updatedRecord.DeputationOfficeFacilityID).
		SetDisabilityTypeID(updatedRecord.DisabilityTypeID).
		SetDisabilityTypeCode(updatedRecord.DisabilityTypeCode).
		SetEntryPostCode(updatedRecord.EntryPostCode).
		SetPresentPostCode(updatedRecord.PresentPostCode).
		SetFeederPostCode(updatedRecord.FeederPostCode).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetDesignationID(updatedRecord.DesignationID).
		SetEducationCode(updatedRecord.EducationCode).
		SetFacilityUniqueID(updatedRecord.FacilityUniqueID).
		SetLienControllingOfficeID(updatedRecord.LienControllingOfficeID).
		SetLienControllingOfficeName(updatedRecord.LienControllingOfficeName).
		SetInDeputation(updatedRecord.InDeputation).
		SetDeputationType(updatedRecord.DeputationType).
		SetDeputationOfficeUniqueId(updatedRecord.DeputationOfficeUniqueId).
		SetDeputationOfficeName(updatedRecord.DeputationOfficeName).
		SetDeputationControllingOfficeID(updatedRecord.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(updatedRecord.DeputationControllingOfficeName).
		SetControllingOfficeFacilityID(updatedRecord.ControllingOfficeFacilityID).
		SetControllingOfficeName(updatedRecord.ControllingOfficeName).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetNodalOfficeName(updatedRecord.NodalOfficeName).
		SetCenterFacilityId(updatedRecord.CenterFacilityId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetPhotoPath(updatedRecord.PhotoPath).
		SetNonQualifyingService(updatedRecord.NonQualifyingService).
		SetSignaturePath(updatedRecord.SignaturePath).
		SetTempHallTicket(updatedRecord.TempHallTicket).
		SetDisabilityTypeDescription(updatedRecord.DisabilityTypeDescription).
		SetDisabilityPercentage(updatedRecord.DisabilityPercentage).
		SetEducationDescription(updatedRecord.EducationDescription).
		SetExamCode(updatedRecord.ExamCode).
		SetExamShortName(updatedRecord.ExamShortName).
		SetExamYear(updatedRecord.ExamYear).
		SetExamCityCenterCode(updatedRecord.CenterId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetSignature(updatedRecord.Signature).
		SetPhoto(updatedRecord.Photo).
		SetApplicationNumber(updatedRecord.ApplicationNumber).
		SetApplnSubmittedDate(updatedRecord.ApplnSubmittedDate).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetWorkingOfficePincode(updatedRecord.WorkingOfficePincode).
		SetWorkingOfficeName(updatedRecord.WorkingOfficeName).
		SetOptionUsed(updatedRecord.OptionUsed).
		SetRemarks(updatedRecord.Remarks).
		SetNAUserName(applicationRecord.NA_UserName).
		SetRecommendedStatus(applicationRecord.RecommendedStatus).
		SetAppliactionRemarks(updatedRecord.AppliactionRemarks).
		SetPunishmentStatus(updatedRecord.PunishmentStatus).             //here added punishmentstatus
		SetDisciplinaryCaseStatus(updatedRecord.DisciplinaryCaseStatus). //here added DisciplinaryCaseStatus
		SetNADate(currentTime).
		SetHallTicketGeneratedFlag(hallticketgeneratedflag).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createRecommendationsRef(ctx context.Context, tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationIp, updatedAppln *ent.Exam_Applications_IP) ([]*ent.RecommendationsIPApplications, error) {

	currentTime := time.Now().Truncate(time.Second)
	recommendationsRef := make([]*ent.RecommendationsIPApplications, len(applicationRecord.Edges.ApplicationDataN))
	for i, recommendation := range applicationRecord.Edges.ApplicationDataN {
		if recommendation.VacancyYear == 0 {
			return nil, fmt.Errorf("recommendations value at index %d is nil", i)
		}
		prevRecommendation, err := tx.RecommendationsIPApplications.
			Query().
			Where(
				recommendationsipapplications.And(
					recommendationsipapplications.EmployeeID(updatedAppln.EmployeeID),
					recommendationsipapplications.ApplicationID(applicationRecord.ID),
				),
			).
			First(ctx)
		if err != nil {
			return nil, err
		}
		RecommendationsRefEntity, err := tx.RecommendationsIPApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamNameCode(updatedAppln.ExamShortName).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(prevRecommendation.CARecommendations).
			SetNORecommendations(recommendation.NO_Recommendations).
			SetCAUserName(prevRecommendation.CAUserName).
			SetCARemarks(prevRecommendation.CARemarks).
			SetCAUpdatedAt(prevRecommendation.CAUpdatedAt).
			SetNOUpdatedAt(currentTime).
			SetNOUserName(applicationRecord.NA_UserName).
			SetNORemarks(recommendation.NO_Remarks).
			SetApplicationStatus("VerifiedRecommendationsByNA").
			Save(ctx)
		if err != nil {
			return nil, err
		}

		recommendationsRef[i] = RecommendationsRefEntity
	}
	return recommendationsRef, nil
}
func createPsRecommendationsRef(ctx context.Context, tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationGroupB, updatedAppln *ent.Exam_Applications_PS) ([]*ent.RecommendationsPSApplications, error) {

	currentTime := time.Now().Truncate(time.Second)
	recommendationsRef := make([]*ent.RecommendationsPSApplications, len(applicationRecord.Edges.ApplicationDataN))
	for i, recommendation := range applicationRecord.Edges.ApplicationDataN {
		if recommendation.VacancyYear == 0 {
			return nil, fmt.Errorf("recommendations value at index %d is nil", i)
		}
		prevRecommendation, err := tx.RecommendationsPSApplications.
			Query().
			Where(
				recommendationspsapplications.And(
					recommendationspsapplications.EmployeeID(updatedAppln.EmployeeID),
					recommendationspsapplications.ApplicationID(applicationRecord.ID),
				),
			).
			First(ctx)
		if err != nil {
			return nil, err
		}
		RecommendationsRefEntity, err := tx.RecommendationsPSApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamNameCode(updatedAppln.ExamShortName).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(prevRecommendation.CARecommendations).
			SetNORecommendations(recommendation.NO_Recommendations).
			SetCAUserName(prevRecommendation.CAUserName).
			SetCARemarks(prevRecommendation.CARemarks).
			SetCAUpdatedAt(prevRecommendation.CAUpdatedAt).
			SetNOUpdatedAt(currentTime).
			SetNOUserName(applicationRecord.NA_UserName).
			SetNORemarks(recommendation.NO_Remarks).
			SetApplicationStatus("VerifiedRecommendationsByNA").
			Save(ctx)
		if err != nil {
			return nil, err
		}

		recommendationsRef[i] = RecommendationsRefEntity
	}
	return recommendationsRef, nil
}
func createMtspmRecommendationsRef(ctx context.Context, tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationMTSPM, updatedAppln *ent.Exam_Application_MTSPMMG) ([]*ent.RecommendationsMTSPMMGApplications, error) {

	currentTime := time.Now().Truncate(time.Second)
	recommendationsRef := make([]*ent.RecommendationsMTSPMMGApplications, len(applicationRecord.Edges.ApplicationDataN))
	for i, recommendation := range applicationRecord.Edges.ApplicationDataN {
		if recommendation.VacancyYear == 0 {
			return nil, fmt.Errorf("recommendations value at index %d is nil", i)
		}
		prevRecommendation, err := tx.RecommendationsMTSPMMGApplications.
			Query().
			Where(
				recommendationsmtspmmgapplications.And(
					recommendationsmtspmmgapplications.EmployeeID(updatedAppln.EmployeeID),
					recommendationsmtspmmgapplications.ApplicationID(applicationRecord.ID),
				),
			).
			First(ctx)
		if err != nil {
			return nil, err
		}
		RecommendationsRefEntity, err := tx.RecommendationsMTSPMMGApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamNameCode(updatedAppln.ExamShortName).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(prevRecommendation.CARecommendations).
			SetNORecommendations(recommendation.NO_Recommendations).
			SetCAUserName(prevRecommendation.CAUserName).
			SetCARemarks(prevRecommendation.CARemarks).
			SetCAUpdatedAt(prevRecommendation.CAUpdatedAt).
			SetNOUpdatedAt(currentTime).
			SetNOUserName(applicationRecord.NA_UserName).
			SetNORemarks(recommendation.NO_Remarks).
			SetApplicationStatus("VerifiedRecommendationsByNA").
			Save(ctx)
		if err != nil {
			return nil, err
		}

		recommendationsRef[i] = RecommendationsRefEntity
	}
	return recommendationsRef, nil
}
func createPmpaRecommendationsRef(ctx context.Context, tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationPMPA, updatedAppln *ent.Exam_Applications_PMPA) ([]*ent.RecommendationsPMPAApplications, error) {

	currentTime := time.Now().Truncate(time.Second)
	recommendationsRef := make([]*ent.RecommendationsPMPAApplications, len(applicationRecord.Edges.ApplicationDataN))
	for i, recommendation := range applicationRecord.Edges.ApplicationDataN {
		if recommendation.VacancyYear == 0 {
			return nil, fmt.Errorf("recommendations value at index %d is nil", i)
		}
		prevRecommendation, err := tx.RecommendationsPMPAApplications.
			Query().
			Where(
				recommendationspmpaapplications.And(
					recommendationspmpaapplications.EmployeeID(updatedAppln.EmployeeID),
					recommendationspmpaapplications.ApplicationID(applicationRecord.ID),
				),
			).
			First(ctx)
		if err != nil {
			return nil, err
		}
		RecommendationsRefEntity, err := tx.RecommendationsPMPAApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamNameCode(updatedAppln.ExamShortName).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(prevRecommendation.CARecommendations).
			SetNORecommendations(recommendation.NO_Recommendations).
			SetCAUserName(prevRecommendation.CAUserName).
			SetCARemarks(prevRecommendation.CARemarks).
			SetCAUpdatedAt(prevRecommendation.CAUpdatedAt).
			SetNOUpdatedAt(currentTime).
			SetNOUserName(applicationRecord.NA_UserName).
			SetNORemarks(recommendation.NO_Remarks).
			SetApplicationStatus("VerifiedRecommendationsByNA").
			Save(ctx)
		if err != nil {
			return nil, err
		}

		recommendationsRef[i] = RecommendationsRefEntity
	}
	return recommendationsRef, nil
}
func createGdspaRecommendationsRef(ctx context.Context, tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationGDStoPA, updatedAppln *ent.Exam_Applications_GDSPA) ([]*ent.RecommendationsGDSPAApplications, error) {

	currentTime := time.Now().Truncate(time.Second)
	recommendationsRef := make([]*ent.RecommendationsGDSPAApplications, len(applicationRecord.Edges.ApplicationDataN))
	for i, recommendation := range applicationRecord.Edges.ApplicationDataN {
		if recommendation.VacancyYear == 0 {
			return nil, fmt.Errorf("recommendations value at index %d is nil", i)
		}
		prevRecommendation, err := tx.RecommendationsGDSPAApplications.
			Query().
			Where(
				recommendationsgdspaapplications.And(
					recommendationsgdspaapplications.EmployeeID(updatedAppln.EmployeeID),
					recommendationsgdspaapplications.ApplicationID(applicationRecord.ID),
				),
			).
			First(ctx)
		if err != nil {
			return nil, err
		}
		RecommendationsRefEntity, err := tx.RecommendationsGDSPAApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamNameCode(updatedAppln.ExamShortName).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(prevRecommendation.CARecommendations).
			SetNORecommendations(recommendation.NO_Recommendations).
			SetCAUserName(prevRecommendation.CAUserName).
			SetCARemarks(prevRecommendation.CARemarks).
			SetCAUpdatedAt(prevRecommendation.CAUpdatedAt).
			SetNOUpdatedAt(currentTime).
			SetNOUserName(applicationRecord.NA_UserName).
			SetNORemarks(recommendation.NO_Remarks).
			SetApplicationStatus("VerifiedRecommendationsByNA").
			Save(ctx)
		if err != nil {
			return nil, err
		}

		recommendationsRef[i] = RecommendationsRefEntity
	}
	return recommendationsRef, nil
}
func createGdspmRecommendationsRef(ctx context.Context, tx *ent.Tx, applicationRecord *ca_reg.NAVerifyApplicationGDSPM, updatedAppln *ent.Exam_Applications_GDSPM) ([]*ent.RecommendationsGDSPMApplications, error) {

	currentTime := time.Now().Truncate(time.Second)
	recommendationsRef := make([]*ent.RecommendationsGDSPMApplications, len(applicationRecord.Edges.ApplicationDataNAv))
	for i, recommendation := range applicationRecord.Edges.ApplicationDataNAv {
		if recommendation.VacancyYear == 0 {
			return nil, fmt.Errorf("recommendations value at index %d is nil", i)
		}
		prevRecommendation, err := tx.RecommendationsGDSPMApplications.
			Query().
			Where(
				recommendationsgdspmapplications.And(
					recommendationsgdspmapplications.EmployeeID(updatedAppln.EmployeeID),
					recommendationsgdspmapplications.ApplicationID(applicationRecord.ID),
				),
			).
			First(ctx)
		if err != nil {
			return nil, err
		}
		RecommendationsRefEntity, err := tx.RecommendationsGDSPMApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamNameCode(updatedAppln.ExamShortName).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(prevRecommendation.CARecommendations).
			SetNORecommendations(recommendation.NO_Recommendations).
			SetCAUserName(prevRecommendation.CAUserName).
			SetCARemarks(prevRecommendation.CARemarks).
			SetCAUpdatedAt(prevRecommendation.CAUpdatedAt).
			SetNOUpdatedAt(currentTime).
			SetNOUserName(applicationRecord.NA_UserName).
			SetNORemarks(recommendation.NO_Remarks).
			SetApplicationStatus("VerifiedRecommendationsByNA").
			Save(ctx)
		if err != nil {
			return nil, err
		}

		recommendationsRef[i] = RecommendationsRefEntity
	}
	return recommendationsRef, nil
}
