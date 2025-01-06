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

	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"

	"time"
)

func validateTempHallTicket(tempHallTicket string, employeeID int64) error {
	if !isNumeric(tempHallTicket) {
		return fmt.Errorf("issue for employee %d with temp hall ticket number: %s", employeeID, tempHallTicket)
	}
	if len(tempHallTicket) != 8 {
		return fmt.Errorf("issue for employee %d with temp hall ticket number length issue: %s", employeeID, tempHallTicket)
	}
	if employeeID == 0 {
		return fmt.Errorf("please enter valid employee id %d", employeeID)
	}

	return nil
}

func validateApplicationData(newAppln *ca_reg.ApplicationIp) error {
	if newAppln.EmployeeID == 0 {
		return errors.New("employee id is missing")
	}
	if newAppln.Edges.CircleData == nil || len(newAppln.Edges.CircleData) == 0 {
		return errors.New("circle preference values are missing")
	}
	if len(newAppln.Edges.CircleData) != 23 {
		return errors.New("invalid number of Circle preferences")
	}
	return nil
}
func handleTransaction(tx *ent.Tx, err *error) {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	} else if *err != nil {
		tx.Rollback()
	} else {
		if commitErr := tx.Commit(); commitErr != nil {
			tx.Rollback()
			*err = commitErr
		}
	}
}
func saveApplication(tx *ent.Tx, newAppln any, applicationNumber string, examcode int32, ctx context.Context) (*ca_reg.ApplicationsResponse, int32, string, error) {
	currentTime := time.Now().Truncate(time.Second)

	var createdAppln interface{}
	var err error

	switch examcode {
	case 2:
		applicationIp, ok := newAppln.(*ca_reg.ApplicationIp)
		if !ok {
			return nil, 422, " -SUB001", fmt.Errorf("invalid application data type for exam code 2")
		}

		if err := validateApplicationData(applicationIp); err != nil {
			return nil, 400, " -SUB002", fmt.Errorf("circle preference values are missing")
		}

		createdAppln, err = tx.Exam_Applications_IP.
			Create().
			SetApplicationNumber(applicationNumber).
			SetApplicationStatus("CAVerificationPending").
			SetApplnSubmittedDate(currentTime).
			SetCandidateRemarks(applicationIp.CandidateRemarks).
			SetCategoryCode(applicationIp.CategoryCode).
			SetCategoryDescription(applicationIp.CategoryDescription).
			SetCenterFacilityId(applicationIp.CenterFacilityId).
			SetCenterId(applicationIp.CenterId).
			SetCentrePreference(applicationIp.CentrePreference).
			SetClaimingQualifyingService(applicationIp.ClaimingQualifyingService).
			SetControllingOfficeFacilityID(applicationIp.ControllingOfficeFacilityID).
			SetControllingOfficeName(applicationIp.ControllingOfficeName).
			SetDCCS(applicationIp.DCCS).
			SetDOB(applicationIp.DOB).
			SetDeputationControllingOfficeID(applicationIp.DeputationControllingOfficeID).
			SetDeputationControllingOfficeName(applicationIp.DeputationControllingOfficeName).
			SetDeputationOfficeFacilityID(applicationIp.DeputationOfficeFacilityID).
			SetDeputationOfficeName(applicationIp.DeputationOfficeName).
			SetDeputationOfficeUniqueId(applicationIp.DeputationOfficeUniqueId).
			SetDeputationOfficePincode(applicationIp.DeputationOfficePincode).
			SetInDeputation(applicationIp.InDeputation).
			SetDeputationType(applicationIp.DeputationType).
			SetDisabilityPercentage(applicationIp.DisabilityPercentage).
			SetDisabilityTypeCode(applicationIp.DisabilityTypeCode).
			SetDisabilityTypeDescription(applicationIp.DisabilityTypeDescription).
			SetEducationDescription(applicationIp.EducationDescription).
			SetEmailID(applicationIp.EmailID).
			SetEmployeeID(applicationIp.EmployeeID).
			SetEmployeeName(applicationIp.EmployeeName).
			SetEntryPostCode(applicationIp.EntryPostCode).
			SetEntryPostDescription(applicationIp.EntryPostDescription).
			SetExamCode(applicationIp.ExamCode).
			SetExamName(applicationIp.ExamName).
			SetExamShortName(applicationIp.ExamShortName).
			SetExamYear(applicationIp.ExamYear).
			SetExamCityCenterCode(applicationIp.CenterId).
			SetFacilityUniqueID(applicationIp.FacilityUniqueID).
			SetFeederPostCode(applicationIp.FeederPostCode).
			SetFeederPostDescription(applicationIp.FeederPostDescription).
			SetFeederPostJoiningDate(applicationIp.FeederPostJoiningDate).
			SetGender(applicationIp.Gender).
			SetLienControllingOfficeID(applicationIp.LienControllingOfficeID).
			SetLienControllingOfficeName(applicationIp.LienControllingOfficeName).
			SetMobileNumber(applicationIp.MobileNumber).
			SetNodalOfficeFacilityID(applicationIp.NodalOfficeFacilityID).
			SetNodalOfficeName(applicationIp.NodalOfficeName).
			SetPhoto(applicationIp.Photo).
			SetPhotoPath(applicationIp.PhotoPath).
			SetPresentDesignation(applicationIp.PresentDesignation).
			SetPresentPostCode(applicationIp.PresentPostCode).
			SetPresentPostDescription(applicationIp.PresentPostDescription).
			SetReportingOfficeFacilityID(applicationIp.ReportingOfficeFacilityID).
			SetReportingOfficeName(applicationIp.ReportingOfficeName).
			SetServiceLength(*applicationIp.ServiceLength).
			SetSignature(applicationIp.Signature).
			SetSignaturePath(applicationIp.SignaturePath).
			SetTempHallTicket(applicationIp.TempHallTicket).
			SetUserID(applicationIp.UserID).
			SetWorkingOfficeCircleFacilityID(applicationIp.WorkingOfficeCircleFacilityID).
			SetWorkingOfficeCircleName(applicationIp.WorkingOfficeCircleName).
			SetWorkingOfficeDivisionFacilityID(applicationIp.WorkingOfficeDivisionFacilityID).
			SetWorkingOfficeDivisionName(applicationIp.WorkingOfficeDivisionName).
			SetWorkingOfficeFacilityID(applicationIp.WorkingOfficeFacilityID).
			SetWorkingOfficeName(applicationIp.WorkingOfficeName).
			SetWorkingOfficePincode(applicationIp.WorkingOfficePincode).
			SetWorkingOfficeRegionFacilityID(applicationIp.WorkingOfficeRegionFacilityID).
			SetWorkingOfficeRegionName(applicationIp.WorkingOfficeRegionName).
			SetDisabilityTypeID(applicationIp.DisabilityTypeID).
			SetDesignationID(applicationIp.DesignationID).
			SetCandidatePhoto(applicationIp.CandidatePhoto).
			SetCandidateSignature(applicationIp.CandidateSignature).
			SetEducationCode(applicationIp.EducationCode).
			Save(ctx)

		if err != nil {
			return nil, 500, " -SUB003", err
		}

		circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(applicationIp.Edges.CircleData))
		for i, circlePrefRef := range applicationIp.Edges.CircleData {
			if circlePrefRef.PlacePrefNo == 0 {
				return nil, 422, " -SUB004", fmt.Errorf("circle preference value at index %d is nil", i)
			}

			circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
				Create().
				SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
				SetApplicationID(createdAppln.(*ent.Exam_Applications_IP).ID).
				SetEmployeeID(applicationIp.EmployeeID).
				SetPlacePrefValue(circlePrefRef.PlacePrefValue).
				SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
				Save(ctx)

			if err != nil {
				return nil, 422, " -SUB005", err
			}
			circlePrefRefs[i] = circlePrefRefEntity
		}

		if _, err = createdAppln.(*ent.Exam_Applications_IP).Update().AddCirclePrefRef(circlePrefRefs...).Save(ctx); err != nil {
			return nil, 500, " -SUB006", err
		}

	case 1:
		applicationGroupB, ok := newAppln.(*ca_reg.ApplicationGroupB)
		if !ok {
			return nil, 422, " -SUB007", fmt.Errorf("invalid application data type for exam code 1")
		}

		createdAppln, err = tx.Exam_Applications_PS.
			Create().
			SetApplicationNumber(applicationNumber).
			SetApplicationStatus("CAVerificationPending").
			SetApplnSubmittedDate(currentTime).
			SetCandidateRemarks(applicationGroupB.CandidateRemarks).
			SetCategoryCode(applicationGroupB.CategoryCode).
			SetCategoryDescription(applicationGroupB.CategoryDescription).
			SetCenterFacilityId(applicationGroupB.CenterFacilityId).
			SetCenterId(applicationGroupB.CenterId).
			SetCentrePreference(applicationGroupB.CentrePreference).
			SetClaimingQualifyingService(applicationGroupB.ClaimingQualifyingService).
			SetControllingOfficeFacilityID(applicationGroupB.ControllingOfficeFacilityID).
			SetControllingOfficeName(applicationGroupB.ControllingOfficeName).
			SetDCCS(applicationGroupB.DCCS).
			SetDOB(applicationGroupB.DOB).
			SetDeputationControllingOfficeID(applicationGroupB.DeputationControllingOfficeID).
			SetDeputationControllingOfficeName(applicationGroupB.DeputationControllingOfficeName).
			SetDeputationOfficeFacilityID(applicationGroupB.DeputationOfficeFacilityID).
			SetDeputationOfficeName(applicationGroupB.DeputationOfficeName).
			SetDeputationOfficeUniqueId(applicationGroupB.DeputationOfficeUniqueId).
			SetDeputationOfficePincode(applicationGroupB.DeputationOfficePincode).
			SetInDeputation(applicationGroupB.InDeputation).
			SetDeputationType(applicationGroupB.DeputationType).
			SetDisabilityPercentage(applicationGroupB.DisabilityPercentage).
			SetDisabilityTypeCode(applicationGroupB.DisabilityTypeCode).
			SetDisabilityTypeDescription(applicationGroupB.DisabilityTypeDescription).
			SetEducationDescription(applicationGroupB.EducationDescription).
			SetEmailID(applicationGroupB.EmailID).
			SetEmployeeID(applicationGroupB.EmployeeID).
			SetEmployeeName(applicationGroupB.EmployeeName).
			SetEntryPostCode(applicationGroupB.EntryPostCode).
			SetEntryPostDescription(applicationGroupB.EntryPostDescription).
			SetExamCode(applicationGroupB.ExamCode).
			SetExamName(applicationGroupB.ExamName).
			SetExamShortName(applicationGroupB.ExamShortName).
			SetExamYear(applicationGroupB.ExamYear).
			SetExamCityCenterCode(applicationGroupB.CenterId).
			SetFacilityUniqueID(applicationGroupB.FacilityUniqueID).
			SetFeederPostCode(applicationGroupB.FeederPostCode).
			SetFeederPostDescription(applicationGroupB.FeederPostDescription).
			SetFeederPostJoiningDate(applicationGroupB.FeederPostJoiningDate).
			SetGender(applicationGroupB.Gender).
			SetLienControllingOfficeID(applicationGroupB.LienControllingOfficeID).
			SetLienControllingOfficeName(applicationGroupB.LienControllingOfficeName).
			SetMobileNumber(applicationGroupB.MobileNumber).
			SetNodalOfficeFacilityID(applicationGroupB.NodalOfficeFacilityID).
			SetNodalOfficeName(applicationGroupB.NodalOfficeName).
			SetPhoto(applicationGroupB.Photo).
			SetPhotoPath(applicationGroupB.PhotoPath).
			SetPresentDesignation(applicationGroupB.PresentDesignation).
			SetPresentPostCode(applicationGroupB.PresentPostCode).
			SetPresentPostDescription(applicationGroupB.PresentPostDescription).
			SetReportingOfficeFacilityID(applicationGroupB.ReportingOfficeFacilityID).
			SetReportingOfficeName(applicationGroupB.ReportingOfficeName).
			SetServiceLength(*applicationGroupB.ServiceLength).
			SetSignature(applicationGroupB.Signature).
			SetSignaturePath(applicationGroupB.SignaturePath).
			SetTempHallTicket(applicationGroupB.TempHallTicket).
			SetUserID(applicationGroupB.UserID).
			SetWorkingOfficeCircleFacilityID(applicationGroupB.WorkingOfficeCircleFacilityID).
			SetWorkingOfficeCircleName(applicationGroupB.WorkingOfficeCircleName).
			SetWorkingOfficeDivisionFacilityID(applicationGroupB.WorkingOfficeDivisionFacilityID).
			SetWorkingOfficeDivisionName(applicationGroupB.WorkingOfficeDivisionName).
			SetWorkingOfficeFacilityID(applicationGroupB.WorkingOfficeFacilityID).
			SetWorkingOfficeName(applicationGroupB.WorkingOfficeName).
			SetWorkingOfficePincode(applicationGroupB.WorkingOfficePincode).
			SetWorkingOfficeRegionFacilityID(applicationGroupB.WorkingOfficeRegionFacilityID).
			SetWorkingOfficeRegionName(applicationGroupB.WorkingOfficeRegionName).
			SetDisabilityTypeID(applicationGroupB.DisabilityTypeID).
			SetDesignationID(applicationGroupB.DesignationID).
			SetEducationCode(applicationGroupB.EducationCode).
			SetCandidatePhoto(applicationGroupB.CandidatePhoto).
			SetCandidateSignature(applicationGroupB.CandidateSignature).
			Save(ctx)

		if err != nil {
			return nil, 500, " -SUB008", err
		}
	case 3:
		applicationPmpa, ok := newAppln.(*ca_reg.ApplicationPMPA)
		if !ok {
			return nil, 422, " -SUB009", fmt.Errorf("invalid application data type for exam code 3")
		}

		createdAppln, err = tx.Exam_Applications_PMPA.
			Create().
			SetApplicationNumber(applicationNumber).
			SetApplicationStatus("CAVerificationPending").
			SetApplnSubmittedDate(currentTime).
			SetCadre(applicationPmpa.Cadre).
			SetCandidateRemarks(applicationPmpa.CandidateRemarks).
			SetCategoryCode(applicationPmpa.CategoryCode).
			SetCategoryDescription(applicationPmpa.CategoryDescription).
			SetCenterFacilityId(applicationPmpa.CenterFacilityId).
			SetCenterId(applicationPmpa.CenterId).
			SetCentrePreference(applicationPmpa.CentrePreference).
			SetClaimingQualifyingService(applicationPmpa.ClaimingQualifyingService).
			SetControllingOfficeFacilityID(applicationPmpa.ControllingOfficeFacilityID).
			SetControllingOfficeName(applicationPmpa.ControllingOfficeName).
			SetDCCS(applicationPmpa.DCCS).
			SetDOB(applicationPmpa.DOB).
			SetDeputationControllingOfficeID(applicationPmpa.DeputationControllingOfficeID).
			SetDeputationControllingOfficeName(applicationPmpa.DeputationControllingOfficeName).
			SetDeputationOfficeFacilityID(applicationPmpa.DeputationOfficeFacilityID).
			SetDeputationOfficeName(applicationPmpa.DeputationOfficeName).
			SetDeputationOfficePincode(applicationPmpa.DeputationOfficePincode).
			SetDeputationOfficeUniqueId(applicationPmpa.DeputationOfficeUniqueId).
			SetInDeputation(applicationPmpa.InDeputation).
			SetDeputationType(applicationPmpa.DeputationType).
			SetDesignationID(applicationPmpa.DesignationID).
			SetDisabilityPercentage(applicationPmpa.DisabilityPercentage).
			SetDisabilityTypeCode(applicationPmpa.DisabilityTypeCode).
			SetDisabilityTypeDescription(applicationPmpa.DisabilityTypeDescription).
			SetDisabilityTypeID(applicationPmpa.DisabilityTypeID).
			SetEducationCode(applicationPmpa.EducationCode).
			SetEducationDescription(applicationPmpa.EducationDescription).
			SetEmailID(applicationPmpa.EmailID).
			SetEmployeeID(applicationPmpa.EmployeeID).
			SetEmployeeName(applicationPmpa.EmployeeName).
			SetEmployeePost(applicationPmpa.EmployeePost).
			SetEntryPostCode(applicationPmpa.EntryPostCode).
			SetEntryPostDescription(applicationPmpa.EntryPostDescription).
			SetExamCode(applicationPmpa.ExamCode).
			SetExamName(applicationPmpa.ExamName).
			SetCandidatePhoto(applicationPmpa.CandidatePhoto).
			SetCandidateSignature(applicationPmpa.CandidateSignature).
			SetExamShortName(applicationPmpa.ExamShortName).
			SetExamYear(applicationPmpa.ExamYear).
			SetExamCityCenterCode(applicationPmpa.CenterId).
			SetFacilityUniqueID(applicationPmpa.FacilityUniqueID).
			SetFeederPostCode(applicationPmpa.FeederPostCode).
			SetFeederPostDescription(applicationPmpa.FeederPostDescription).
			SetFeederPostJoiningDate(applicationPmpa.FeederPostJoiningDate).
			SetGender(applicationPmpa.Gender).
			SetLienControllingOfficeID(applicationPmpa.LienControllingOfficeID).
			SetLienControllingOfficeName(applicationPmpa.LienControllingOfficeName).
			SetMobileNumber(applicationPmpa.MobileNumber).
			SetNodalOfficeFacilityID(applicationPmpa.NodalOfficeFacilityID).
			SetNodalOfficeName(applicationPmpa.NodalOfficeName).
			SetPhoto(applicationPmpa.Photo).
			SetPhotoPath(applicationPmpa.PhotoPath).
			SetPostPreferences(*applicationPmpa.PostPreferences).
			SetPresentDesignation(applicationPmpa.PresentDesignation).
			SetPresentPostCode(applicationPmpa.PresentPostCode).
			SetPresentPostDescription(applicationPmpa.PresentPostDescription).
			SetReportingOfficeFacilityID(applicationPmpa.ReportingOfficeFacilityID).
			SetReportingOfficeName(applicationPmpa.ReportingOfficeName).
			SetServiceLength(*applicationPmpa.ServiceLength).
			SetSignature(applicationPmpa.Signature).
			SetSignaturePath(applicationPmpa.SignaturePath).
			SetTempHallTicket(applicationPmpa.TempHallTicket).
			SetUnitPreferences(*applicationPmpa.UnitPreferences).
			SetUserID(applicationPmpa.UserID).
			SetWorkingOfficeCircleFacilityID(applicationPmpa.WorkingOfficeCircleFacilityID).
			SetWorkingOfficeCircleName(applicationPmpa.WorkingOfficeCircleName).
			SetWorkingOfficeDivisionFacilityID(applicationPmpa.WorkingOfficeDivisionFacilityID).
			SetWorkingOfficeDivisionName(applicationPmpa.WorkingOfficeDivisionName).
			SetWorkingOfficeFacilityID(applicationPmpa.WorkingOfficeFacilityID).
			SetWorkingOfficeName(applicationPmpa.WorkingOfficeName).
			SetWorkingOfficePincode(applicationPmpa.WorkingOfficePincode).
			SetWorkingOfficeRegionFacilityID(applicationPmpa.WorkingOfficeRegionFacilityID).
			SetWorkingOfficeRegionName(applicationPmpa.WorkingOfficeRegionName).
			SetPMMailGuardMTSEngagement(*applicationPmpa.PMMailGuardMTSEngagement).
			Save(ctx)
		if err != nil {
			return nil, 422, " -SUB010", fmt.Errorf("failed to save application: %w", err)
		}
	case 5:
		applicationMtsPmMg, ok := newAppln.(*ca_reg.ApplicationMTSPM)
		if !ok {
			return nil, 422, " -SUB011", fmt.Errorf("invalid application data type for exam code 4")
		}

		createdAppln, err = tx.Exam_Application_MTSPMMG.
			Create().
			SetApplicationNumber(applicationNumber).
			SetApplicationStatus("CAVerificationPending").
			SetApplnSubmittedDate(currentTime).
			SetCadre(applicationMtsPmMg.Cadre).
			SetCandidateRemarks(applicationMtsPmMg.CandidateRemarks).
			SetCategoryCode(applicationMtsPmMg.CategoryCode).
			SetCategoryDescription(applicationMtsPmMg.CategoryDescription).
			SetCenterFacilityId(applicationMtsPmMg.CenterFacilityId).
			SetCenterId(applicationMtsPmMg.CenterId).
			SetCentrePreference(applicationMtsPmMg.CentrePreference).
			SetClaimingQualifyingService(applicationMtsPmMg.ClaimingQualifyingService).
			SetControllingOfficeFacilityID(applicationMtsPmMg.ControllingOfficeFacilityID).
			SetControllingOfficeName(applicationMtsPmMg.ControllingOfficeName).
			SetDCCS(applicationMtsPmMg.DCCS).
			SetDOB(applicationMtsPmMg.DOB).
			SetDeputationControllingOfficeID(applicationMtsPmMg.DeputationControllingOfficeID).
			SetDeputationControllingOfficeName(applicationMtsPmMg.DeputationControllingOfficeName).
			SetDeputationOfficeFacilityID(applicationMtsPmMg.DeputationOfficeFacilityID).
			SetDeputationOfficeName(applicationMtsPmMg.DeputationOfficeName).
			SetDeputationOfficePincode(applicationMtsPmMg.DeputationOfficePincode).
			SetDeputationOfficeUniqueId(applicationMtsPmMg.DeputationOfficeUniqueId).
			SetInDeputation(applicationMtsPmMg.InDeputation).
			SetDeputationType(applicationMtsPmMg.DeputationType).
			SetDesignationID(applicationMtsPmMg.DesignationID).
			SetDisabilityPercentage(applicationMtsPmMg.DisabilityPercentage).
			SetDisabilityTypeCode(applicationMtsPmMg.DisabilityTypeCode).
			SetDisabilityTypeDescription(applicationMtsPmMg.DisabilityTypeDescription).
			SetDisabilityTypeID(applicationMtsPmMg.DisabilityTypeID).
			SetEducationCode(applicationMtsPmMg.EducationCode).
			SetEducationDescription(applicationMtsPmMg.EducationDescription).
			SetEmailID(applicationMtsPmMg.EmailID).
			SetEmployeeID(applicationMtsPmMg.EmployeeID).
			SetEmployeeName(applicationMtsPmMg.EmployeeName).
			SetEmployeePost(applicationMtsPmMg.EmployeePost).
			SetEntryPostCode(applicationMtsPmMg.EntryPostCode).
			SetEntryPostDescription(applicationMtsPmMg.EntryPostDescription).
			SetExamCode(applicationMtsPmMg.ExamCode).
			SetExamName(applicationMtsPmMg.ExamName).
			SetExamShortName(applicationMtsPmMg.ExamShortName).
			SetExamYear(applicationMtsPmMg.ExamYear).
			SetExamCityCenterCode(applicationMtsPmMg.CenterId).
			SetFacilityUniqueID(applicationMtsPmMg.FacilityUniqueID).
			SetFeederPostCode(applicationMtsPmMg.FeederPostCode).
			SetFeederPostDescription(applicationMtsPmMg.FeederPostDescription).
			SetFeederPostJoiningDate(applicationMtsPmMg.FeederPostJoiningDate).
			SetGender(applicationMtsPmMg.Gender).
			SetGDSEngagement(*applicationMtsPmMg.GDSEngagement).
			SetLienControllingOfficeID(applicationMtsPmMg.LienControllingOfficeID).
			SetLienControllingOfficeName(applicationMtsPmMg.LienControllingOfficeName).
			SetMobileNumber(applicationMtsPmMg.MobileNumber).
			SetNodalOfficeFacilityID(applicationMtsPmMg.NodalOfficeFacilityID).
			SetNodalOfficeName(applicationMtsPmMg.NodalOfficeName).
			SetPhoto(applicationMtsPmMg.Photo).
			SetPhotoPath(applicationMtsPmMg.PhotoPath).
			SetPostPreferences(*applicationMtsPmMg.PostPreferences).
			SetPresentDesignation(applicationMtsPmMg.PresentDesignation).
			SetPresentPostCode(applicationMtsPmMg.PresentPostCode).
			SetPresentPostDescription(applicationMtsPmMg.PresentPostDescription).
			SetReportingOfficeFacilityID(applicationMtsPmMg.ReportingOfficeFacilityID).
			SetReportingOfficeName(applicationMtsPmMg.ReportingOfficeName).
			SetServiceLength(*applicationMtsPmMg.ServiceLength).
			SetSignature(applicationMtsPmMg.Signature).
			SetSignaturePath(applicationMtsPmMg.SignaturePath).
			SetTempHallTicket(applicationMtsPmMg.TempHallTicket).
			SetUnitPreferences(*applicationMtsPmMg.UnitPreferences).
			SetUserID(applicationMtsPmMg.UserID).
			SetWorkingOfficeCircleFacilityID(applicationMtsPmMg.WorkingOfficeCircleFacilityID).
			SetWorkingOfficeCircleName(applicationMtsPmMg.WorkingOfficeCircleName).
			SetWorkingOfficeDivisionFacilityID(applicationMtsPmMg.WorkingOfficeDivisionFacilityID).
			SetWorkingOfficeDivisionName(applicationMtsPmMg.WorkingOfficeDivisionName).
			SetWorkingOfficeFacilityID(applicationMtsPmMg.WorkingOfficeFacilityID).
			SetWorkingOfficeName(applicationMtsPmMg.WorkingOfficeName).
			SetWorkingOfficePincode(applicationMtsPmMg.WorkingOfficePincode).
			SetWorkingOfficeRegionFacilityID(applicationMtsPmMg.WorkingOfficeRegionFacilityID).
			SetWorkingOfficeRegionName(applicationMtsPmMg.WorkingOfficeRegionName).
			SetCandidatePhoto(applicationMtsPmMg.CandidatePhoto).
			SetCandidateSignature(applicationMtsPmMg.CandidateSignature).
			Save(ctx)
		if err != nil {
			return nil, 422, " -SUB012", fmt.Errorf("failed to save application: %w", err)
		}
	case 6:
		applicationGdsPM, ok := newAppln.(*ca_reg.ApplicationGDSPM)
		if !ok {
			return nil, 422, " -SUB013", fmt.Errorf("invalid application data type for exam code 4")
		}

		createdAppln, err = tx.Exam_Applications_GDSPM.
			Create().
			SetApplicationNumber(applicationNumber).
			SetApplicationStatus("CAVerificationPending").
			SetApplnSubmittedDate(currentTime).
			SetCadre(applicationGdsPM.Cadre).
			SetCandidateRemarks(applicationGdsPM.CandidateRemarks).
			SetCategoryCode(applicationGdsPM.CategoryCode).
			SetCategoryDescription(applicationGdsPM.CategoryDescription).
			SetCenterFacilityId(applicationGdsPM.CenterFacilityId).
			SetCenterId(applicationGdsPM.CenterId).
			SetCentrePreference(applicationGdsPM.CentrePreference).
			SetClaimingQualifyingService(applicationGdsPM.ClaimingQualifyingService).
			SetControllingOfficeFacilityID(applicationGdsPM.ControllingOfficeFacilityID).
			SetControllingOfficeName(applicationGdsPM.ControllingOfficeName).
			SetDCCS(applicationGdsPM.DCCS).
			SetDOB(applicationGdsPM.DOB).
			SetDeputationControllingOfficeID(applicationGdsPM.DeputationControllingOfficeID).
			SetDeputationControllingOfficeName(applicationGdsPM.DeputationControllingOfficeName).
			SetDeputationOfficeFacilityID(applicationGdsPM.DeputationOfficeFacilityID).
			SetDeputationOfficeName(applicationGdsPM.DeputationOfficeName).
			SetDeputationOfficePincode(applicationGdsPM.DeputationOfficePincode).
			SetDeputationOfficeUniqueId(applicationGdsPM.DeputationOfficeUniqueId).
			SetInDeputation(applicationGdsPM.InDeputation).
			SetDeputationType(applicationGdsPM.DeputationType).
			SetDesignationID(applicationGdsPM.DesignationID).
			SetDisabilityPercentage(applicationGdsPM.DisabilityPercentage).
			SetDisabilityTypeCode(applicationGdsPM.DisabilityTypeCode).
			SetDisabilityTypeDescription(applicationGdsPM.DisabilityTypeDescription).
			SetDisabilityTypeID(applicationGdsPM.DisabilityTypeID).
			SetEducationCode(applicationGdsPM.EducationCode).
			SetEducationDescription(applicationGdsPM.EducationDescription).
			SetEmailID(applicationGdsPM.EmailID).
			SetEmployeeID(applicationGdsPM.EmployeeID).
			SetEmployeeName(applicationGdsPM.EmployeeName).
			SetEmployeePost(applicationGdsPM.EmployeePost).
			SetEntryPostCode(applicationGdsPM.EntryPostCode).
			SetEntryPostDescription(applicationGdsPM.EntryPostDescription).
			SetExamCode(applicationGdsPM.ExamCode).
			SetExamName(applicationGdsPM.ExamName).
			SetExamShortName(applicationGdsPM.ExamShortName).
			SetExamYear(applicationGdsPM.ExamYear).
			SetExamCityCenterCode(applicationGdsPM.CenterId).
			SetFacilityUniqueID(applicationGdsPM.FacilityUniqueID).
			SetFeederPostCode(applicationGdsPM.FeederPostCode).
			SetFeederPostDescription(applicationGdsPM.FeederPostDescription).
			SetFeederPostJoiningDate(applicationGdsPM.FeederPostJoiningDate).
			SetGender(applicationGdsPM.Gender).
			SetLienControllingOfficeID(applicationGdsPM.LienControllingOfficeID).
			SetLienControllingOfficeName(applicationGdsPM.LienControllingOfficeName).
			SetMobileNumber(applicationGdsPM.MobileNumber).
			SetNodalOfficeFacilityID(applicationGdsPM.NodalOfficeFacilityID).
			SetNodalOfficeName(applicationGdsPM.NodalOfficeName).
			SetPhoto(applicationGdsPM.Photo).
			SetPhotoPath(applicationGdsPM.PhotoPath).
			SetPostPreferences(*applicationGdsPM.PostPreferences).
			SetPresentDesignation(applicationGdsPM.PresentDesignation).
			SetPresentPostCode(applicationGdsPM.PresentPostCode).
			SetPresentPostDescription(applicationGdsPM.PresentPostDescription).
			SetReportingOfficeFacilityID(applicationGdsPM.ReportingOfficeFacilityID).
			SetReportingOfficeName(applicationGdsPM.ReportingOfficeName).
			SetServiceLength(*applicationGdsPM.ServiceLength).
			SetSignature(applicationGdsPM.Signature).
			SetSignaturePath(applicationGdsPM.SignaturePath).
			SetTempHallTicket(applicationGdsPM.TempHallTicket).
			SetUnitPreferences(*applicationGdsPM.UnitPreferences).
			SetUserID(applicationGdsPM.UserID).
			SetWorkingOfficeCircleFacilityID(applicationGdsPM.WorkingOfficeCircleFacilityID).
			SetWorkingOfficeCircleName(applicationGdsPM.WorkingOfficeCircleName).
			SetWorkingOfficeDivisionFacilityID(applicationGdsPM.WorkingOfficeDivisionFacilityID).
			SetWorkingOfficeDivisionName(applicationGdsPM.WorkingOfficeDivisionName).
			SetWorkingOfficeFacilityID(applicationGdsPM.WorkingOfficeFacilityID).
			SetWorkingOfficeName(applicationGdsPM.WorkingOfficeName).
			SetWorkingOfficePincode(applicationGdsPM.WorkingOfficePincode).
			SetWorkingOfficeRegionFacilityID(applicationGdsPM.WorkingOfficeRegionFacilityID).
			SetWorkingOfficeRegionName(applicationGdsPM.WorkingOfficeRegionName).
			SetGDSEngagement(*applicationGdsPM.GDSEngagement).
			SetPMMailGuardMTSEngagement(*applicationGdsPM.PMMailGuardMTSEngagement).
			SetCandidatePhoto(applicationGdsPM.CandidatePhoto).
			SetCandidateSignature(applicationGdsPM.CandidateSignature).
			Save(ctx)
		if err != nil {
			return nil, 422, " -SUB014", fmt.Errorf("failed to save application: %w", err)
		}
	case 4:
		applicationGdsPa, ok := newAppln.(*ca_reg.ApplicationGDStoPA)
		if !ok {
			return nil, 422, " -SUB015", fmt.Errorf("invalid application data type for exam code 4")
		}

		createdAppln, err = tx.Exam_Applications_GDSPA.
			Create().
			SetApplicationNumber(applicationNumber).
			SetApplicationStatus("CAVerificationPending").
			SetApplnSubmittedDate(currentTime).
			SetCadre(applicationGdsPa.Cadre).
			SetCandidateRemarks(applicationGdsPa.CandidateRemarks).
			SetCategoryCode(applicationGdsPa.CategoryCode).
			SetCategoryDescription(applicationGdsPa.CategoryDescription).
			SetCenterFacilityId(applicationGdsPa.CenterFacilityId).
			SetCenterId(applicationGdsPa.CenterId).
			SetCentrePreference(applicationGdsPa.CentrePreference).
			SetClaimingQualifyingService(applicationGdsPa.ClaimingQualifyingService).
			SetControllingOfficeFacilityID(applicationGdsPa.ControllingOfficeFacilityID).
			SetControllingOfficeName(applicationGdsPa.ControllingOfficeName).
			SetDCCS(applicationGdsPa.DCCS).
			SetDOB(applicationGdsPa.DOB).
			SetDeputationControllingOfficeID(applicationGdsPa.DeputationControllingOfficeID).
			SetDeputationControllingOfficeName(applicationGdsPa.DeputationControllingOfficeName).
			SetDeputationOfficeFacilityID(applicationGdsPa.DeputationOfficeFacilityID).
			SetDeputationOfficeName(applicationGdsPa.DeputationOfficeName).
			SetDeputationOfficePincode(applicationGdsPa.DeputationOfficePincode).
			SetDeputationOfficeUniqueId(applicationGdsPa.DeputationOfficeUniqueId).
			SetInDeputation(applicationGdsPa.InDeputation).
			SetDeputationType(applicationGdsPa.DeputationType).
			SetDesignationID(applicationGdsPa.DesignationID).
			SetDisabilityPercentage(applicationGdsPa.DisabilityPercentage).
			SetDisabilityTypeCode(applicationGdsPa.DisabilityTypeCode).
			SetDisabilityTypeDescription(applicationGdsPa.DisabilityTypeDescription).
			SetDisabilityTypeID(applicationGdsPa.DisabilityTypeID).
			SetEducationCode(applicationGdsPa.EducationCode).
			SetEducationDescription(applicationGdsPa.EducationDescription).
			SetEmailID(applicationGdsPa.EmailID).
			SetEmployeeID(applicationGdsPa.EmployeeID).
			SetEmployeeName(applicationGdsPa.EmployeeName).
			SetEmployeePost(applicationGdsPa.EmployeePost).
			SetEntryPostCode(applicationGdsPa.EntryPostCode).
			SetEntryPostDescription(applicationGdsPa.EntryPostDescription).
			SetExamCode(applicationGdsPa.ExamCode).
			SetExamName(applicationGdsPa.ExamName).
			SetExamShortName(applicationGdsPa.ExamShortName).
			SetExamYear(applicationGdsPa.ExamYear).
			SetExamCityCenterCode(applicationGdsPa.CenterId).
			SetFacilityUniqueID(applicationGdsPa.FacilityUniqueID).
			SetFeederPostCode(applicationGdsPa.FeederPostCode).
			SetFeederPostDescription(applicationGdsPa.FeederPostDescription).
			SetFeederPostJoiningDate(applicationGdsPa.FeederPostJoiningDate).
			SetGender(applicationGdsPa.Gender).
			SetLienControllingOfficeID(applicationGdsPa.LienControllingOfficeID).
			SetLienControllingOfficeName(applicationGdsPa.LienControllingOfficeName).
			SetMobileNumber(applicationGdsPa.MobileNumber).
			SetNodalOfficeFacilityID(applicationGdsPa.NodalOfficeFacilityID).
			SetNodalOfficeName(applicationGdsPa.NodalOfficeName).
			SetPhoto(applicationGdsPa.Photo).
			SetPhotoPath(applicationGdsPa.PhotoPath).
			SetPostPreferences(*applicationGdsPa.PostPreferences).
			SetPresentDesignation(applicationGdsPa.PresentDesignation).
			SetPresentPostCode(applicationGdsPa.PresentPostCode).
			SetPresentPostDescription(applicationGdsPa.PresentPostDescription).
			SetReportingOfficeFacilityID(applicationGdsPa.ReportingOfficeFacilityID).
			SetReportingOfficeName(applicationGdsPa.ReportingOfficeName).
			SetServiceLength(*applicationGdsPa.ServiceLength).
			SetSignature(applicationGdsPa.Signature).
			SetSignaturePath(applicationGdsPa.SignaturePath).
			SetTempHallTicket(applicationGdsPa.TempHallTicket).
			SetUnitPreferences(*applicationGdsPa.UnitPreferences).
			SetUserID(applicationGdsPa.UserID).
			SetWorkingOfficeCircleFacilityID(applicationGdsPa.WorkingOfficeCircleFacilityID).
			SetWorkingOfficeCircleName(applicationGdsPa.WorkingOfficeCircleName).
			SetWorkingOfficeDivisionFacilityID(applicationGdsPa.WorkingOfficeDivisionFacilityID).
			SetWorkingOfficeDivisionName(applicationGdsPa.WorkingOfficeDivisionName).
			SetWorkingOfficeFacilityID(applicationGdsPa.WorkingOfficeFacilityID).
			SetWorkingOfficeName(applicationGdsPa.WorkingOfficeName).
			SetWorkingOfficePincode(applicationGdsPa.WorkingOfficePincode).
			SetWorkingOfficeRegionFacilityID(applicationGdsPa.WorkingOfficeRegionFacilityID).
			SetWorkingOfficeRegionName(applicationGdsPa.WorkingOfficeRegionName).
			SetCandidatePhoto(applicationGdsPa.CandidatePhoto).
			SetCandidateSignature(applicationGdsPa.CandidateSignature).
			Save(ctx)
		if err != nil {
			return nil, 422, " -SUB016", fmt.Errorf("failed to save application: %w", err)
		}
	default:
		return nil, 400, " -SUB017", fmt.Errorf("unsupported exam code: %d", examcode)
	}

	appResponse, err := MapExamApplicationsToResponse(createdAppln)
	if err != nil {
		return nil, 500, " -SUB018", err
	}

	return appResponse, 200, " ", nil
}

func checkIfApplicationExists(ctx context.Context, tx *ent.Tx, employeeID int64, examYear string, examCode int32, statuses []string) (bool, int32, string, error) {
	var exists bool
	var err error

	// Switch based on examCode to select the appropriate table and query
	switch examCode {
	case 1:
		exists, err = tx.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.EmployeeIDEQ(employeeID),
				exam_applications_ps.ApplicationStatusIn(statuses...),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
			).
			Exist(ctx)
	case 2:
		exists, err = tx.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.EmployeeIDEQ(employeeID),
				exam_applications_ip.ApplicationStatusIn(statuses...),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
			).
			Exist(ctx)

	case 3:
		exists, err = tx.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.EmployeeIDEQ(employeeID),
				exam_applications_pmpa.ApplicationStatusIn(statuses...),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
			).
			Exist(ctx)

	case 4:
		exists, err = tx.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.EmployeeIDEQ(employeeID),
				exam_applications_gdspa.ApplicationStatusIn(statuses...),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
			).
			Exist(ctx)

	case 5:
		exists, err = tx.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.EmployeeIDEQ(employeeID),
				exam_application_mtspmmg.ApplicationStatusIn(statuses...),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
			).
			Exist(ctx)

	case 6:
		exists, err = tx.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.EmployeeIDEQ(employeeID),
				exam_applications_gdspm.ApplicationStatusIn(statuses...),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
			).
			Exist(ctx)
	default:
		return false, 422, " -SUB0001 ", fmt.Errorf("exam code %d is not supported", examCode)
	}

	if err != nil {
		return false, 500, " -SUB0002 ", err
	}
	return exists, 200, "", nil
}

func fetchExistingApplication(ctx context.Context, tx *ent.Tx, employeeID int64, examYear string) (*ent.Exam_Applications_IP, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(employeeID),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active"),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application for this candidate")
		}
		return nil, 500, " -SUB002", err
	}
	return oldAppln, 200, "", nil
}
func fetchExistingPsApplication(ctx context.Context, tx *ent.Tx, employeeID int64, examYear string) (*ent.Exam_Applications_PS, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.EmployeeIDEQ(employeeID),
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.StatusEQ("active"),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application for this candidate")
		}
		return nil, 500, " -SUB002", err
	}
	return oldAppln, 200, "", nil
}
func fetchExistingPmpaApplication(ctx context.Context, tx *ent.Tx, employeeID int64, examYear string) (*ent.Exam_Applications_PMPA, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_PMPA.
		Query().
		Where(
			exam_applications_pmpa.EmployeeIDEQ(employeeID),
			exam_applications_pmpa.ExamYearEQ(examYear),
			exam_applications_pmpa.StatusEQ("active"),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application for this candidate")
		}
		return nil, 500, " -SUB002", err
	}
	return oldAppln, 200, "", nil
}
func fetchExistingGdspaApplication(ctx context.Context, tx *ent.Tx, employeeID int64, examYear string) (*ent.Exam_Applications_GDSPA, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.EmployeeIDEQ(employeeID),
			exam_applications_gdspa.ExamYearEQ(examYear),
			exam_applications_gdspa.StatusEQ("active"),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application for this candidate")
		}
		return nil, 500, " -SUB002", err
	}
	return oldAppln, 200, "", nil
}
func fetchExistingGdspmApplication(ctx context.Context, tx *ent.Tx, employeeID int64, examYear string) (*ent.Exam_Applications_GDSPM, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.EmployeeIDEQ(employeeID),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application for this candidate")
		}
		return nil, 500, " -SUB002", err
	}
	return oldAppln, 200, "", nil
}
func fetchExistingMtspmApplication(ctx context.Context, tx *ent.Tx, employeeID int64, examYear string) (*ent.Exam_Application_MTSPMMG, int32, string, error) {
	oldAppln, err := tx.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.EmployeeIDEQ(employeeID),
			exam_application_mtspmmg.ExamYearEQ(examYear),
			exam_application_mtspmmg.StatusEQ("active"),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application for this candidate")
		}
		return nil, 500, " -SUB002", err
	}
	return oldAppln, 200, "", nil
}

/*func determineApplicationStatus(ctx context.Context, oldAppln *ent.Exam_Applications_IP, remarks string) (string, error) {
	stat := "inactive_" + time.Now().Format("20060102150405")

	switch oldAppln.ApplicationStatus {
	case "VerifiedByNA", "VerifiedByCA":
		return "", errors.New("the Application was already verified by Nodal Authority/ Controlling Authority")
	case "CAVerificationPending", "ResubmitCAVerificationPending":
		if remarks == "InCorrect" {
			if _, err := oldAppln.Update().SetStatus(stat).Save(ctx); err != nil {
				return "", err
			}
			return "PendingWithCandidate", nil
		}
		if remarks == "Correct" {
			if _, err := oldAppln.Update().SetStatus(stat).Save(ctx); err != nil {
				return "", err
			}
			return "VerifiedByCA", nil
		}
		return "", errors.New("invalid CA_Remarks value")
	default:
		return "", errors.New("invalid ApplicationStatus")
	}
} */

func determineApplicationStatus(ctx context.Context, oldAppln interface{}, remarks string, examCode int32) (string, int32, string, error) {
	stat := "inactive_" + time.Now().Format("20060102150405")

	var applicationStatus string
	var err error

	switch examCode {
	case 1: // examapplicationPS
		appln, ok := oldAppln.(*ent.Exam_Applications_PS)
		if !ok {
			return "", 422, " -SUB001", errors.New("invalid application type for examCode 1")
		}
		appln = appln.Unwrap()
		applicationStatus, err = func() (string, error) {
			var updateFunc func() (any, error) = func() (any, error) {
				return appln.Update().SetStatus(stat).Save(ctx)
			}
			return processApplicationStatus(appln.ApplicationStatus, remarks, updateFunc)
		}()
	case 2: // examapplicationIP
		appln, ok := oldAppln.(*ent.Exam_Applications_IP)
		if !ok {
			return "", 422, " -SUB002", errors.New("invalid application type for examCode 2")
		}
		applicationStatus, err = func() (string, error) {
			var updateFunc func() (any, error) = func() (any, error) {
				return appln.Update().SetStatus(stat).Save(ctx)
			}
			return processApplicationStatus(appln.ApplicationStatus, remarks, updateFunc)
		}()
	case 5: // examapplicationIP
		appln, ok := oldAppln.(*ent.Exam_Application_MTSPMMG)
		if !ok {
			return "", 422, " -SUB002", errors.New("invalid application type for examCode 2")
		}
		applicationStatus, err = func() (string, error) {
			var updateFunc func() (any, error) = func() (any, error) {
				return appln.Update().SetStatus(stat).Save(ctx)
			}
			return processApplicationStatus(appln.ApplicationStatus, remarks, updateFunc)
		}()
	case 6: // examapplicationIP
		appln, ok := oldAppln.(*ent.Exam_Applications_GDSPM)
		if !ok {
			return "", 422, " -SUB002", errors.New("invalid application type for examCode 2")
		}
		applicationStatus, err = func() (string, error) {
			var updateFunc func() (any, error) = func() (any, error) {
				return appln.Update().SetStatus(stat).Save(ctx)
			}
			return processApplicationStatus(appln.ApplicationStatus, remarks, updateFunc)
		}()
	case 4: // examapplicationIP
		appln, ok := oldAppln.(*ent.Exam_Applications_GDSPA)
		if !ok {
			return "", 422, " -SUB002", errors.New("invalid application type for examCode 2")
		}
		applicationStatus, err = func() (string, error) {
			var updateFunc func() (any, error) = func() (any, error) {
				return appln.Update().SetStatus(stat).Save(ctx)
			}
			return processApplicationStatus(appln.ApplicationStatus, remarks, updateFunc)
		}()
	case 3: // examapplicationIP
		appln, ok := oldAppln.(*ent.Exam_Applications_PMPA)
		if !ok {
			return "", 422, " -SUB002", errors.New("invalid application type for examCode 2")
		}
		applicationStatus, err = func() (string, error) {
			var updateFunc func() (any, error) = func() (any, error) {
				return appln.Update().SetStatus(stat).Save(ctx)
			}
			return processApplicationStatus(appln.ApplicationStatus, remarks, updateFunc)
		}()
	default:
		return "", 422, " -SUB003", errors.New("invalid ExamCode")
	}

	if err != nil {
		return "", 500, " -SUB004", err
	}

	return applicationStatus, 200, "", nil
}

func processApplicationStatus(appStatus, remarks string, updateFunc func() (any, error)) (string, error) {
	switch appStatus {
	case "VerifiedByNA", "VerifiedByCA":
		return "", errors.New("the Application was already verified by Nodal Authority/Controlling Authority")
	case "CAVerificationPending", "ResubmitCAVerificationPending":
		if remarks == "InCorrect" {
			if _, err := updateFunc(); err != nil {
				return "", err
			}
			return "PendingWithCandidate", nil
		}
		if remarks == "Correct" {
			if _, err := updateFunc(); err != nil {
				return "", err
			}
			return "VerifiedByCA", nil
		}
		return "", errors.New("invalid CA_Remarks value")
	default:
		return "", errors.New("invalid ApplicationStatus")
	}
}

func createUpdatedApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_IP, newAppln *ca_reg.VerifyApplicationIp, status string) (*ent.Exam_Applications_IP, error) {
	return tx.Exam_Applications_IP.
		Create().
		SetAppliactionRemarks(newAppln.AppliactionRemarks).
		SetApplicationNumber(oldAppln.ApplicationNumber).
		SetApplicationStatus(status).
		SetApplnSubmittedDate(oldAppln.ApplnSubmittedDate).
		SetCADate(time.Now().Truncate(time.Second)).
		SetCAEmployeeDesignation(newAppln.CA_EmployeeDesignation).
		SetCAEmployeeID(newAppln.CA_EmployeeID).
		SetCAGeneralRemarks(newAppln.CA_GeneralRemarks).
		SetCARemarks(newAppln.CA_Remarks).
		SetCAUserName(newAppln.CA_UserName).
		SetCadre(oldAppln.Cadre).
		SetCandidateRemarks(oldAppln.CandidateRemarks).
		SetCategoryCode(oldAppln.CategoryCode).
		SetCategoryDescription(oldAppln.CategoryDescription).
		SetCenterFacilityId(oldAppln.CenterFacilityId).
		SetCenterId(oldAppln.CenterId).
		SetCentrePreference(oldAppln.CentrePreference).
		SetClaimingQualifyingService(oldAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(oldAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(oldAppln.ControllingOfficeName).
		SetDCCS(oldAppln.DCCS).
		SetDCInPresentCadre(oldAppln.DCInPresentCadre).
		SetDOB(oldAppln.DOB).
		SetDeputationControllingOfficeID(oldAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(oldAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(oldAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(oldAppln.DeputationOfficeName).
		SetDeputationOfficePincode(oldAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(oldAppln.DeputationOfficeUniqueId).
		SetInDeputation(oldAppln.InDeputation).
		SetDeputationType(oldAppln.DeputationType).
		SetDesignationID(oldAppln.DesignationID).
		SetDisabilityPercentage(oldAppln.DisabilityPercentage).
		SetDisabilityTypeCode(oldAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(oldAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(oldAppln.DisabilityTypeID).
		SetEducationCode(oldAppln.EducationCode).
		SetEducationDescription(oldAppln.EducationDescription).
		SetEmailID(oldAppln.EmailID).
		SetEmployeeID(oldAppln.EmployeeID).
		SetEmployeeName(oldAppln.EmployeeName).
		SetEmployeePost(oldAppln.EmployeePost).
		SetEntryPostCode(oldAppln.EntryPostCode).
		SetEntryPostDescription(oldAppln.EntryPostDescription).
		SetExamCode(oldAppln.ExamCode).
		SetExamName(oldAppln.ExamName).
		SetExamShortName(oldAppln.ExamShortName).
		SetExamYear(oldAppln.ExamYear).
		SetExamCityCenterCode(oldAppln.ExamCityCenterCode).
		SetFacilityUniqueID(oldAppln.FacilityUniqueID).
		SetFeederPostCode(oldAppln.FeederPostCode).
		SetFeederPostDescription(oldAppln.FeederPostDescription).
		SetFeederPostJoiningDate(oldAppln.FeederPostJoiningDate).
		SetGender(oldAppln.Gender).
		SetGenerateHallTicketFlag(newAppln.GenerateHallTicketFlag).
		SetLienControllingOfficeID(oldAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(oldAppln.LienControllingOfficeName).
		SetMobileNumber(oldAppln.MobileNumber).
		SetNodalOfficeFacilityID(oldAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(oldAppln.NodalOfficeName).
		SetOptionUsed(oldAppln.OptionUsed).
		SetPhoto(oldAppln.Photo).
		SetPhotoPath(oldAppln.PhotoPath).
		SetPresentDesignation(oldAppln.PresentDesignation).
		SetPresentPostCode(oldAppln.PresentPostCode).
		SetPresentPostDescription(oldAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(oldAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(oldAppln.ReportingOfficeName).
		SetServiceLength(*newAppln.ServiceLength).
		SetSignature(oldAppln.Signature).
		SetSignaturePath(oldAppln.SignaturePath).
		SetStatus("active").
		SetTempHallTicket(oldAppln.TempHallTicket).
		SetUserID(oldAppln.UserID).
		SetWorkingOfficeCircleFacilityID(oldAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(oldAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(oldAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(oldAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(oldAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(oldAppln.WorkingOfficeName).
		SetWorkingOfficePincode(oldAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(oldAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(oldAppln.WorkingOfficeRegionName).
		SetRecommendedStatus(newAppln.RecommendedStatus).
		SetPunishmentStatus(newAppln.PunishmentStatus).
		SetDisciplinaryCaseStatus(newAppln.DisciplinaryCaseStatus).
		Save(ctx)
}

// func MapExamApplicationsToResponse(appIP *ent.Exam_Applications_IP) *ca_reg.ApplicationsResponse {
// 	return &ca_reg.ApplicationsResponse{
// 		EmployeeID:         appIP.EmployeeID,
// 		ApplicationNumber:  appIP.ApplicationNumber,
// 		ApplicationStatus:  appIP.ApplicationStatus,
// 		MobileNumber:       appIP.MobileNumber,
// 		EmailID:            appIP.EmailID,
// 		RoleUserCode:       appIP.RoleUserCode,
// 		ApplnSubmittedDate: appIP.ApplnSubmittedDate,
// 	}
// }

/*
	func MapExamApplicationsToResponse(appIP interface{}) (*ca_reg.ApplicationsResponse, error) {
		switch application := appIP.(type) {
		case *ent.Exam_Applications_IP:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            application.EmployeeID,
				ApplicationNumber:     application.ApplicationNumber,
				ApplicationStatus:     application.ApplicationStatus,
				MobileNumber:          application.MobileNumber,
				EmailID:               application.EmailID,
				RoleUserCode:          application.RoleUserCode,
				ApplnSubmittedDate:    application.ApplnSubmittedDate,
				ControllingOfficeName: application.ControllingOfficeName,
				RecommendedStatus:     application.RecommendedStatus,
				ApplicationRemarks:    application.AppliactionRemarks,
				CAGeneralRemarks:      application.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Applications_PS:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            application.EmployeeID,
				ApplicationNumber:     application.ApplicationNumber,
				ApplicationStatus:     application.ApplicationStatus,
				MobileNumber:          application.MobileNumber,
				EmailID:               application.EmailID,
				RoleUserCode:          application.RoleUserCode,
				ApplnSubmittedDate:    application.ApplnSubmittedDate,
				ControllingOfficeName: application.ControllingOfficeName,
				RecommendedStatus:     application.RecommendedStatus,
				ApplicationRemarks:    application.AppliactionRemarks,
				CAGeneralRemarks:      application.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Applications_PMPA:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            application.EmployeeID,
				ApplicationNumber:     application.ApplicationNumber,
				ApplicationStatus:     application.ApplicationStatus,
				MobileNumber:          application.MobileNumber,
				EmailID:               application.EmailID,
				RoleUserCode:          application.RoleUserCode,
				ApplnSubmittedDate:    application.ApplnSubmittedDate,
				ControllingOfficeName: application.ControllingOfficeName,
				RecommendedStatus:     application.RecommendedStatus,
				ApplicationRemarks:    application.AppliactionRemarks,
				CAGeneralRemarks:      application.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Applications_GDSPM:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            application.EmployeeID,
				ApplicationNumber:     application.ApplicationNumber,
				ApplicationStatus:     application.ApplicationStatus,
				MobileNumber:          application.MobileNumber,
				EmailID:               application.EmailID,
				RoleUserCode:          application.RoleUserCode,
				ApplnSubmittedDate:    application.ApplnSubmittedDate,
				ControllingOfficeName: application.ControllingOfficeName,
				RecommendedStatus:     application.RecommendedStatus,
				ApplicationRemarks:    application.AppliactionRemarks,
				CAGeneralRemarks:      application.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Application_MTSPMMG:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            application.EmployeeID,
				ApplicationNumber:     application.ApplicationNumber,
				ApplicationStatus:     application.ApplicationStatus,
				MobileNumber:          application.MobileNumber,
				EmailID:               application.EmailID,
				RoleUserCode:          application.RoleUserCode,
				ApplnSubmittedDate:    application.ApplnSubmittedDate,
				ControllingOfficeName: application.ControllingOfficeName,
				RecommendedStatus:     application.RecommendedStatus,
				ApplicationRemarks:    application.AppliactionRemarks,
				CAGeneralRemarks:      application.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Applications_GDSPA:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            application.EmployeeID,
				ApplicationNumber:     application.ApplicationNumber,
				ApplicationStatus:     application.ApplicationStatus,
				MobileNumber:          application.MobileNumber,
				EmailID:               application.EmailID,
				RoleUserCode:          application.RoleUserCode,
				ApplnSubmittedDate:    application.ApplnSubmittedDate,
				ControllingOfficeName: application.ControllingOfficeName,
				RecommendedStatus:     application.RecommendedStatus,
				ApplicationRemarks:    application.AppliactionRemarks,
				CAGeneralRemarks:      application.CAGeneralRemarks,
			}, nil
		default:
			return nil, fmt.Errorf("unexpected type %T, expected *ca_reg.ApplicationIp", appIP)
		}
	}
*/
func MapExamApplicationsToResponse(appIP interface{}) (*ca_reg.ApplicationsResponse, error) {
	// Helper function to map common fields
	mapCommonFields := func(application interface{}) (*ca_reg.ApplicationsResponse, error) {
		switch app := application.(type) {
		case *ent.Exam_Applications_IP:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            app.EmployeeID,
				ApplicationNumber:     app.ApplicationNumber,
				ApplicationStatus:     app.ApplicationStatus,
				MobileNumber:          app.MobileNumber,
				EmailID:               app.EmailID,
				RoleUserCode:          app.RoleUserCode,
				ApplnSubmittedDate:    app.ApplnSubmittedDate,
				ControllingOfficeName: app.ControllingOfficeName,
				RecommendedStatus:     app.RecommendedStatus,
				ApplicationRemarks:    app.AppliactionRemarks,
				CAGeneralRemarks:      app.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Applications_PMPA:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            app.EmployeeID,
				ApplicationNumber:     app.ApplicationNumber,
				ApplicationStatus:     app.ApplicationStatus,
				MobileNumber:          app.MobileNumber,
				EmailID:               app.EmailID,
				RoleUserCode:          app.RoleUserCode,
				ApplnSubmittedDate:    app.ApplnSubmittedDate,
				ControllingOfficeName: app.ControllingOfficeName,
				RecommendedStatus:     app.RecommendedStatus,
				ApplicationRemarks:    app.AppliactionRemarks,
				CAGeneralRemarks:      app.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Applications_GDSPM:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            app.EmployeeID,
				ApplicationNumber:     app.ApplicationNumber,
				ApplicationStatus:     app.ApplicationStatus,
				MobileNumber:          app.MobileNumber,
				EmailID:               app.EmailID,
				RoleUserCode:          app.RoleUserCode,
				ApplnSubmittedDate:    app.ApplnSubmittedDate,
				ControllingOfficeName: app.ControllingOfficeName,
				RecommendedStatus:     app.RecommendedStatus,
				ApplicationRemarks:    app.AppliactionRemarks,
				CAGeneralRemarks:      app.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Applications_GDSPA:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            app.EmployeeID,
				ApplicationNumber:     app.ApplicationNumber,
				ApplicationStatus:     app.ApplicationStatus,
				MobileNumber:          app.MobileNumber,
				EmailID:               app.EmailID,
				RoleUserCode:          app.RoleUserCode,
				ApplnSubmittedDate:    app.ApplnSubmittedDate,
				ControllingOfficeName: app.ControllingOfficeName,
				RecommendedStatus:     app.RecommendedStatus,
				ApplicationRemarks:    app.AppliactionRemarks,
				CAGeneralRemarks:      app.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Application_MTSPMMG:
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            app.EmployeeID,
				ApplicationNumber:     app.ApplicationNumber,
				ApplicationStatus:     app.ApplicationStatus,
				MobileNumber:          app.MobileNumber,
				EmailID:               app.EmailID,
				RoleUserCode:          app.RoleUserCode,
				ApplnSubmittedDate:    app.ApplnSubmittedDate,
				ControllingOfficeName: app.ControllingOfficeName,
				RecommendedStatus:     app.RecommendedStatus,
				ApplicationRemarks:    app.AppliactionRemarks,
				CAGeneralRemarks:      app.CAGeneralRemarks,
			}, nil
		case *ent.Exam_Applications_PS:
			// Reuse the same logic for all other types
			return &ca_reg.ApplicationsResponse{
				EmployeeID:            app.EmployeeID, // Accessing via type assertion
				ApplicationNumber:     app.ApplicationNumber,
				ApplicationStatus:     app.ApplicationStatus,
				MobileNumber:          app.MobileNumber,
				EmailID:               app.EmailID,
				RoleUserCode:          app.RoleUserCode,
				ApplnSubmittedDate:    app.ApplnSubmittedDate,
				ControllingOfficeName: app.ControllingOfficeName,
				RecommendedStatus:     app.RecommendedStatus,
				ApplicationRemarks:    app.AppliactionRemarks,
				CAGeneralRemarks:      app.CAGeneralRemarks,
			}, nil
		default:
			return nil, fmt.Errorf("invalid application type: %T", app)
		}
	}

	// Switch case to handle different types
	switch application := appIP.(type) {
	case *ent.Exam_Applications_IP, *ent.Exam_Applications_PS, *ent.Exam_Applications_PMPA,
		*ent.Exam_Applications_GDSPM, *ent.Exam_Application_MTSPMMG, *ent.Exam_Applications_GDSPA:
		return mapCommonFields(application)
	default:
		return nil, fmt.Errorf("unexpected type %T, expected application type", appIP)
	}
}

func MapExamApplicationsGdsToResponse(appIP *ent.Exam_Applications_GDSPM) *ca_reg.ApplicationsResponse {
	return &ca_reg.ApplicationsResponse{
		EmployeeID:         appIP.EmployeeID,
		ApplicationNumber:  appIP.ApplicationNumber,
		ApplicationStatus:  appIP.ApplicationStatus,
		MobileNumber:       appIP.MobileNumber,
		EmailID:            appIP.EmailID,
		RoleUserCode:       appIP.RoleUserCode,
		ApplnSubmittedDate: appIP.ApplnSubmittedDate,
	}
}
func MapExamApplicationsPmPaToResponse(appIP *ent.Exam_Applications_PMPA) *ca_reg.ApplicationsResponse {
	return &ca_reg.ApplicationsResponse{
		EmployeeID:         appIP.EmployeeID,
		ApplicationNumber:  appIP.ApplicationNumber,
		ApplicationStatus:  appIP.ApplicationStatus,
		MobileNumber:       appIP.MobileNumber,
		EmailID:            appIP.EmailID,
		RoleUserCode:       appIP.RoleUserCode,
		ApplnSubmittedDate: appIP.ApplnSubmittedDate,
	}
}
func MapExamApplicationsMTSPmToResponse(appIP *ent.Exam_Application_MTSPMMG) *ca_reg.ApplicationsResponse {
	return &ca_reg.ApplicationsResponse{
		EmployeeID:         appIP.EmployeeID,
		ApplicationNumber:  appIP.ApplicationNumber,
		ApplicationStatus:  appIP.ApplicationStatus,
		MobileNumber:       appIP.MobileNumber,
		EmailID:            appIP.EmailID,
		RoleUserCode:       appIP.RoleUserCode,
		ApplnSubmittedDate: appIP.ApplnSubmittedDate,
	}
}
func MapExamApplicationsGdsPAToResponse(appIP *ent.Exam_Applications_GDSPA) *ca_reg.ApplicationsResponse {
	return &ca_reg.ApplicationsResponse{
		EmployeeID:         appIP.EmployeeID,
		ApplicationNumber:  appIP.ApplicationNumber,
		ApplicationStatus:  appIP.ApplicationStatus,
		MobileNumber:       appIP.MobileNumber,
		EmailID:            appIP.EmailID,
		RoleUserCode:       appIP.RoleUserCode,
		ApplnSubmittedDate: appIP.ApplnSubmittedDate,
	}
}
func MapExamApplicationsPSToResponse(appIP *ent.Exam_Applications_PS) *ca_reg.ApplicationsResponse {
	return &ca_reg.ApplicationsResponse{
		EmployeeID:         appIP.EmployeeID,
		ApplicationNumber:  appIP.ApplicationNumber,
		ApplicationStatus:  appIP.ApplicationStatus,
		MobileNumber:       appIP.MobileNumber,
		EmailID:            appIP.EmailID,
		RoleUserCode:       appIP.RoleUserCode,
		ApplnSubmittedDate: appIP.ApplnSubmittedDate,
	}
}

func createUpdateApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_PS, newAppln *ca_reg.VerifyApplicationGroupB, applicationStatus, nonQualifyService string) (*ent.Exam_Applications_PS, error) {
	currentTime := time.Now().Truncate(time.Second)

	updatedAppln, err := tx.Exam_Applications_PS.
		Create().
		SetAppliactionRemarks(newAppln.AppliactionRemarks).
		SetApplicationNumber(oldAppln.ApplicationNumber).
		SetApplicationStatus(applicationStatus).
		SetApplnSubmittedDate(oldAppln.ApplnSubmittedDate).
		SetCADate(currentTime).
		SetCAEmployeeDesignation(newAppln.CA_EmployeeDesignation).
		SetCAEmployeeID(newAppln.CA_EmployeeID).
		SetCAGeneralRemarks(newAppln.CA_GeneralRemarks).
		SetCARemarks(newAppln.CA_Remarks).
		SetCAUserName(newAppln.CA_UserName).
		SetRecommendedStatus(newAppln.RecommendedStatus).
		SetCadre(oldAppln.Cadre).
		SetCandidateRemarks(oldAppln.CandidateRemarks).
		SetCategoryCode(oldAppln.CategoryCode).
		SetCategoryDescription(oldAppln.CategoryDescription).
		SetCenterFacilityId(oldAppln.CenterFacilityId).
		SetCenterId(oldAppln.CenterId).
		SetCentrePreference(oldAppln.CentrePreference).
		SetClaimingQualifyingService(oldAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(oldAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(oldAppln.ControllingOfficeName).
		SetDCCS(oldAppln.DCCS).
		SetDOB(oldAppln.DOB).
		SetDeputationControllingOfficeID(oldAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(oldAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(oldAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(oldAppln.DeputationOfficeName).
		SetDeputationOfficePincode(oldAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(oldAppln.DeputationOfficeUniqueId).
		SetInDeputation(oldAppln.InDeputation).
		SetDeputationType(oldAppln.DeputationType).
		SetDesignationID(oldAppln.DesignationID).
		SetDisabilityPercentage(oldAppln.DisabilityPercentage).
		SetDisabilityTypeCode(oldAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(oldAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(oldAppln.DisabilityTypeID).
		SetEducationCode(oldAppln.EducationCode).
		SetEducationDescription(oldAppln.EducationDescription).
		SetEmailID(oldAppln.EmailID).
		SetEmployeeID(oldAppln.EmployeeID).
		SetEmployeeName(oldAppln.EmployeeName).
		SetEmployeePost(oldAppln.EmployeePost).
		SetEntryPostCode(oldAppln.EntryPostCode).
		SetEntryPostDescription(oldAppln.EntryPostDescription).
		SetExamCode(oldAppln.ExamCode).
		SetExamName(oldAppln.ExamName).
		SetExamShortName(oldAppln.ExamShortName).
		SetExamYear(oldAppln.ExamYear).
		SetExamCityCenterCode(oldAppln.CenterId).
		SetFacilityUniqueID(oldAppln.FacilityUniqueID).
		SetFeederPostCode(oldAppln.FeederPostCode).
		SetFeederPostDescription(oldAppln.FeederPostDescription).
		SetFeederPostJoiningDate(oldAppln.FeederPostJoiningDate).
		SetGender(oldAppln.Gender).
		SetGenerateHallTicketFlag(newAppln.GenerateHallTicketFlag).
		SetLienControllingOfficeID(oldAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(oldAppln.LienControllingOfficeName).
		SetMobileNumber(oldAppln.MobileNumber).
		SetNodalOfficeFacilityID(oldAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(oldAppln.NodalOfficeName).
		SetNonQualifyingService(*newAppln.NonQualifyingService).
		SetOptionUsed(oldAppln.OptionUsed).
		SetPhoto(oldAppln.Photo).
		SetPhotoPath(oldAppln.PhotoPath).
		SetPresentDesignation(oldAppln.PresentDesignation).
		SetPresentPostCode(oldAppln.PresentPostCode).
		SetPresentPostDescription(oldAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(oldAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(oldAppln.ReportingOfficeName).
		SetServiceLength(oldAppln.ServiceLength).
		SetSignature(oldAppln.Signature).
		SetSignaturePath(oldAppln.SignaturePath).
		SetStatus("active").
		SetTempHallTicket(oldAppln.TempHallTicket).
		SetUserID(oldAppln.UserID).
		SetWorkingOfficeCircleFacilityID(oldAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(oldAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(oldAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(oldAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(oldAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(oldAppln.WorkingOfficeName).
		SetWorkingOfficePincode(oldAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(oldAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(oldAppln.WorkingOfficeRegionName).
		SetPunishmentStatus(newAppln.PunishmentStatus).
		SetDisciplinaryCaseStatus(newAppln.DisciplinaryCaseStatus).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if nonQualifyService == "Yes" {
		_, err := updatedAppln.Update().SetNonQualifyingService(*newAppln.NonQualifyingService).Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	return updatedAppln, nil
}
func createUpdatePmpaApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_PMPA, newAppln *ca_reg.VerifyApplicationPMPA, applicationStatus, nonQualifyService string) (*ent.Exam_Applications_PMPA, error) {
	currentTime := time.Now().Truncate(time.Second)

	updatedAppln, err := tx.Exam_Applications_PMPA.
		Create().
		SetAppliactionRemarks(newAppln.AppliactionRemarks).
		SetApplicationNumber(oldAppln.ApplicationNumber).
		SetApplicationStatus(applicationStatus).
		SetApplnSubmittedDate(oldAppln.ApplnSubmittedDate).
		SetCADate(currentTime).
		SetCAEmployeeDesignation(newAppln.CA_EmployeeDesignation).
		SetCAEmployeeID(newAppln.CA_EmployeeID).
		SetCAGeneralRemarks(newAppln.CA_GeneralRemarks).
		SetCARemarks(newAppln.CA_Remarks).
		SetCAUserName(newAppln.CA_UserName).
		SetRecommendedStatus(newAppln.RecommendedStatus).
		SetCadre(oldAppln.Cadre).
		SetCandidatePhoto(oldAppln.CandidatePhoto).
		SetCandidateSignature(oldAppln.CandidateSignature).
		SetCandidateRemarks(oldAppln.CandidateRemarks).
		SetCategoryCode(oldAppln.CategoryCode).
		SetCategoryDescription(oldAppln.CategoryDescription).
		SetCenterFacilityId(oldAppln.CenterFacilityId).
		SetCenterId(oldAppln.CenterId).
		SetCentrePreference(oldAppln.CentrePreference).
		SetCentrePreference(oldAppln.CentrePreference).
		SetClaimingQualifyingService(oldAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(oldAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(oldAppln.ControllingOfficeName).
		SetDCCS(oldAppln.DCCS).
		SetDOB(oldAppln.DOB).
		SetDeputationControllingOfficeID(oldAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(oldAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(oldAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(oldAppln.DeputationOfficeName).
		SetDeputationOfficePincode(oldAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(oldAppln.DeputationOfficeUniqueId).
		SetInDeputation(oldAppln.InDeputation).
		SetDeputationType(oldAppln.DeputationType).
		SetDesignationID(oldAppln.DesignationID).
		SetDisabilityPercentage(oldAppln.DisabilityPercentage).
		SetDisabilityTypeCode(oldAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(oldAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(oldAppln.DisabilityTypeID).
		SetEducationCode(oldAppln.EducationCode).
		SetEducationDescription(oldAppln.EducationDescription).
		SetEmailID(oldAppln.EmailID).
		SetEmployeeID(oldAppln.EmployeeID).
		SetEmployeeName(oldAppln.EmployeeName).
		SetEmployeePost(oldAppln.EmployeePost).
		SetEntryPostCode(oldAppln.EntryPostCode).
		SetEntryPostDescription(oldAppln.EntryPostDescription).
		SetExamCode(oldAppln.ExamCode).
		SetExamName(oldAppln.ExamName).
		SetExamShortName(oldAppln.ExamShortName).
		SetExamYear(oldAppln.ExamYear).
		SetExamCityCenterCode(oldAppln.CenterId).
		SetFacilityUniqueID(oldAppln.FacilityUniqueID).
		SetFeederPostCode(oldAppln.FeederPostCode).
		SetFeederPostDescription(oldAppln.FeederPostDescription).
		SetFeederPostJoiningDate(oldAppln.FeederPostJoiningDate).
		SetGender(oldAppln.Gender).
		SetGenerateHallTicketFlag(newAppln.GenerateHallTicketFlag).
		SetLienControllingOfficeID(oldAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(oldAppln.LienControllingOfficeName).
		SetMobileNumber(oldAppln.MobileNumber).
		SetNodalOfficeFacilityID(oldAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(oldAppln.NodalOfficeName).
		SetOptionUsed(oldAppln.OptionUsed).
		SetPMMailGuardMTSEngagement(oldAppln.PMMailGuardMTSEngagement).
		SetPhoto(oldAppln.Photo).
		SetPhotoPath(oldAppln.PhotoPath).
		SetPostPreferences(oldAppln.PostPreferences).
		SetPresentDesignation(oldAppln.PresentDesignation).
		SetPresentPostCode(oldAppln.PresentPostCode).
		SetPresentPostDescription(oldAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(oldAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(oldAppln.ReportingOfficeName).
		SetServiceLength(oldAppln.ServiceLength).
		SetSignature(oldAppln.Signature).
		SetSignaturePath(oldAppln.SignaturePath).
		SetStatus("active").
		SetTempHallTicket(oldAppln.TempHallTicket).
		SetUnitPreferences(oldAppln.UnitPreferences).
		SetUserID(oldAppln.UserID).
		SetWorkingOfficeCircleFacilityID(oldAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(oldAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(oldAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(oldAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(oldAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(oldAppln.WorkingOfficeName).
		SetWorkingOfficePincode(oldAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(oldAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(oldAppln.WorkingOfficeRegionName).
		SetPunishmentStatus(newAppln.PunishmentStatus).             //here added punishmentstatus
		SetDisciplinaryCaseStatus(newAppln.DisciplinaryCaseStatus). //here added DisciplinaryCaseStatus

		Save(ctx)

	if err != nil {
		return nil, err
	}

	if nonQualifyService == "Yes" {
		_, err := updatedAppln.Update().SetNonQualifyingService(*newAppln.NonQualifyingService).Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	return updatedAppln, nil
}
func createUpdateGdspaApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_GDSPA, newAppln *ca_reg.VerifyApplicationGDStoPA, applicationStatus, nonQualifyService string) (*ent.Exam_Applications_GDSPA, error) {
	currentTime := time.Now().Truncate(time.Second)

	updatedAppln, err := tx.Exam_Applications_GDSPA.
		Create().
		SetAppliactionRemarks(newAppln.AppliactionRemarks).
		SetApplicationNumber(oldAppln.ApplicationNumber).
		SetApplicationStatus(applicationStatus).
		SetApplnSubmittedDate(oldAppln.ApplnSubmittedDate).
		SetCADate(currentTime).
		SetCAEmployeeDesignation(newAppln.CA_EmployeeDesignation).
		SetCAEmployeeID(newAppln.CA_EmployeeID).
		SetCAGeneralRemarks(newAppln.CA_GeneralRemarks).
		SetCARemarks(newAppln.CA_Remarks).
		SetCAUserName(newAppln.CA_UserName).
		SetCadre(oldAppln.Cadre).
		SetCandidateRemarks(oldAppln.CandidateRemarks).
		SetCategoryCode(oldAppln.CategoryCode).
		SetCategoryDescription(oldAppln.CategoryDescription).
		SetCenterFacilityId(oldAppln.CenterFacilityId).
		SetCenterId(oldAppln.CenterId).
		SetCentrePreference(oldAppln.CentrePreference).
		SetCandidatePhoto(oldAppln.CandidatePhoto).
		SetCandidateSignature(oldAppln.CandidateSignature).
		SetClaimingQualifyingService(oldAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(oldAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(oldAppln.ControllingOfficeName).
		SetDCCS(oldAppln.DCCS).
		SetDOB(oldAppln.DOB).
		SetDeputationControllingOfficeID(oldAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(oldAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(oldAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(oldAppln.DeputationOfficeName).
		SetDeputationOfficePincode(oldAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(oldAppln.DeputationOfficeUniqueId).
		SetInDeputation(oldAppln.InDeputation).
		SetDeputationType(oldAppln.DeputationType).
		SetDesignationID(oldAppln.DesignationID).
		SetDisabilityPercentage(oldAppln.DisabilityPercentage).
		SetDisabilityTypeCode(oldAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(oldAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(oldAppln.DisabilityTypeID).
		SetEducationCode(oldAppln.EducationCode).
		SetEducationDescription(oldAppln.EducationDescription).
		SetEmailID(oldAppln.EmailID).
		SetEmployeeID(oldAppln.EmployeeID).
		SetEmployeeName(oldAppln.EmployeeName).
		SetEmployeePost(oldAppln.EmployeePost).
		SetEntryPostCode(oldAppln.EntryPostCode).
		SetEntryPostDescription(oldAppln.EntryPostDescription).
		SetExamCode(oldAppln.ExamCode).
		SetExamName(oldAppln.ExamName).
		SetExamShortName(oldAppln.ExamShortName).
		SetExamYear(oldAppln.ExamYear).
		SetExamCityCenterCode(oldAppln.CenterId).
		SetFacilityUniqueID(oldAppln.FacilityUniqueID).
		SetFeederPostCode(oldAppln.FeederPostCode).
		SetFeederPostDescription(oldAppln.FeederPostDescription).
		SetFeederPostJoiningDate(oldAppln.FeederPostJoiningDate).
		SetGender(oldAppln.Gender).
		SetGenerateHallTicketFlag(newAppln.GenerateHallTicketFlag).
		SetLienControllingOfficeID(oldAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(oldAppln.LienControllingOfficeName).
		SetMobileNumber(oldAppln.MobileNumber).
		SetNodalOfficeFacilityID(oldAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(oldAppln.NodalOfficeName).
		SetOptionUsed(oldAppln.OptionUsed).
		SetPhoto(oldAppln.Photo).
		SetPhotoPath(oldAppln.PhotoPath).
		SetPresentDesignation(oldAppln.PresentDesignation).
		SetPresentPostCode(oldAppln.PresentPostCode).
		SetPresentPostDescription(oldAppln.PresentPostDescription).
		SetPostPreferences(oldAppln.PostPreferences).
		SetReportingOfficeFacilityID(oldAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(oldAppln.ReportingOfficeName).
		SetServiceLength(oldAppln.ServiceLength).
		SetSignature(oldAppln.Signature).
		SetSignaturePath(oldAppln.SignaturePath).
		SetStatus("active").
		SetSubdivisionOfficeFacilityID(oldAppln.SubdivisionOfficeFacilityID).
		SetSubdivisionOfficeName(oldAppln.SubdivisionOfficeName).
		SetTempHallTicket(oldAppln.TempHallTicket).
		SetUnitPreferences(oldAppln.UnitPreferences).
		SetUserID(oldAppln.UserID).
		SetWorkingOfficeCircleFacilityID(oldAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(oldAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(oldAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(oldAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(oldAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(oldAppln.WorkingOfficeName).
		SetWorkingOfficePincode(oldAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(oldAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(oldAppln.WorkingOfficeRegionName).
		SetRecommendedStatus(newAppln.RecommendedStatus).
		SetPunishmentStatus(newAppln.PunishmentStatus).             //here added punishmentstatus
		SetDisciplinaryCaseStatus(newAppln.DisciplinaryCaseStatus). //here added DisciplinaryCaseStatus
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if nonQualifyService == "Yes" {
		_, err := updatedAppln.Update().SetNonQualifyingService(*newAppln.NonQualifyingService).Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	return updatedAppln, nil
}
func createUpdatemtspmApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Application_MTSPMMG, newAppln *ca_reg.VerifyApplicationMTSPM, applicationStatus, nonQualifyService string) (*ent.Exam_Application_MTSPMMG, error) {
	currentTime := time.Now().Truncate(time.Second)

	updatedAppln, err := tx.Exam_Application_MTSPMMG.
		Create().
		SetAppliactionRemarks(newAppln.AppliactionRemarks).
		SetRecommendedStatus(newAppln.RecommendedStatus).
		SetApplicationNumber(oldAppln.ApplicationNumber).
		SetApplicationStatus(applicationStatus).
		SetApplnSubmittedDate(oldAppln.ApplnSubmittedDate).
		SetCandidatePhoto(oldAppln.CandidatePhoto).
		SetCandidateSignature(oldAppln.CandidateSignature).
		SetCADate(currentTime).
		SetCAEmployeeDesignation(newAppln.CA_EmployeeDesignation).
		SetCAEmployeeID(newAppln.CA_EmployeeID).
		SetCAGeneralRemarks(newAppln.CA_GeneralRemarks).
		SetCARemarks(newAppln.CA_Remarks).
		SetCAUserName(newAppln.CA_UserName).
		SetCadre(oldAppln.Cadre).
		SetCandidateRemarks(oldAppln.CandidateRemarks).
		SetCategoryCode(oldAppln.CategoryCode).
		SetCategoryDescription(oldAppln.CategoryDescription).
		SetCenterFacilityId(oldAppln.CenterFacilityId).
		SetCenterId(oldAppln.CenterId).
		SetCentrePreference(oldAppln.CentrePreference).
		SetClaimingQualifyingService(oldAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(oldAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(oldAppln.ControllingOfficeName).
		SetDCCS(oldAppln.DCCS).
		SetDOB(oldAppln.DOB).
		SetDeputationControllingOfficeID(oldAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(oldAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(oldAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(oldAppln.DeputationOfficeName).
		SetDeputationOfficePincode(oldAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(oldAppln.DeputationOfficeUniqueId).
		SetInDeputation(oldAppln.InDeputation).
		SetDeputationType(oldAppln.DeputationType).
		SetDesignationID(oldAppln.DesignationID).
		SetDisabilityPercentage(oldAppln.DisabilityPercentage).
		SetDisabilityTypeCode(oldAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(oldAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(oldAppln.DisabilityTypeID).
		SetEducationCode(oldAppln.EducationCode).
		SetEducationDescription(oldAppln.EducationDescription).
		SetEmailID(oldAppln.EmailID).
		SetEmployeeID(oldAppln.EmployeeID).
		SetEmployeeName(oldAppln.EmployeeName).
		SetEmployeePost(oldAppln.EmployeePost).
		SetEntryPostCode(oldAppln.EntryPostCode).
		SetEntryPostDescription(oldAppln.EntryPostDescription).
		SetExamCode(oldAppln.ExamCode).
		SetExamName(oldAppln.ExamName).
		SetExamShortName(oldAppln.ExamShortName).
		SetExamYear(oldAppln.ExamYear).
		SetExamCityCenterCode(oldAppln.CenterId).
		SetFacilityUniqueID(oldAppln.FacilityUniqueID).
		SetFeederPostCode(oldAppln.FeederPostCode).
		SetFeederPostDescription(oldAppln.FeederPostDescription).
		SetFeederPostJoiningDate(oldAppln.FeederPostJoiningDate).
		SetGender(oldAppln.Gender).
		SetGDSEngagement(oldAppln.GDSEngagement).
		SetGenerateHallTicketFlag(newAppln.GenerateHallTicketFlag).
		SetLienControllingOfficeID(oldAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(oldAppln.LienControllingOfficeName).
		SetMobileNumber(oldAppln.MobileNumber).
		SetNodalOfficeFacilityID(oldAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(oldAppln.NodalOfficeName).
		SetOptionUsed(oldAppln.OptionUsed).
		SetPhoto(oldAppln.Photo).
		SetPhotoPath(oldAppln.PhotoPath).
		SetPostPreferences(oldAppln.PostPreferences).
		SetPresentDesignation(oldAppln.PresentDesignation).
		SetPresentPostCode(oldAppln.PresentPostCode).
		SetPresentPostDescription(oldAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(oldAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(oldAppln.ReportingOfficeName).
		SetServiceLength(oldAppln.ServiceLength).
		SetSignature(oldAppln.Signature).
		SetSignaturePath(oldAppln.SignaturePath).
		SetStatus("active").
		SetTempHallTicket(oldAppln.TempHallTicket).
		SetUnitPreferences(oldAppln.UnitPreferences).
		SetUserID(oldAppln.UserID).
		SetWorkingOfficeCircleFacilityID(oldAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(oldAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(oldAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(oldAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(oldAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(oldAppln.WorkingOfficeName).
		SetWorkingOfficePincode(oldAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(oldAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(oldAppln.WorkingOfficeRegionName).
		SetPunishmentStatus(newAppln.PunishmentStatus).             //here added punishmentstatus
		SetDisciplinaryCaseStatus(newAppln.DisciplinaryCaseStatus). //here added DisciplinaryCaseStatus
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if nonQualifyService == "Yes" {
		_, err := updatedAppln.Update().
			SetNonQualifyingService(*newAppln.NonQualifyingService).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	return updatedAppln, nil
}
func createUpdategdspmApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_GDSPM, newAppln *ca_reg.VerifyApplicationGDSPM, applicationStatus, nonQualifyService string) (*ent.Exam_Applications_GDSPM, error) {
	currentTime := time.Now().Truncate(time.Second)

	updatedAppln, err := tx.Exam_Applications_GDSPM.
		Create().
		SetAppliactionRemarks(newAppln.AppliactionRemarks).
		SetRecommendedStatus(newAppln.RecommendedStatus).
		SetApplicationNumber(oldAppln.ApplicationNumber).
		SetApplicationStatus(applicationStatus).
		SetApplnSubmittedDate(oldAppln.ApplnSubmittedDate).
		SetCADate(currentTime).
		SetCAEmployeeDesignation(newAppln.CA_EmployeeDesignation).
		SetCAEmployeeID(newAppln.CA_EmployeeID).
		SetCAGeneralRemarks(newAppln.CA_GeneralRemarks).
		SetCARemarks(newAppln.CA_Remarks).
		SetCAUserName(newAppln.CA_UserName).
		SetCandidatePhoto(oldAppln.CandidatePhoto).
		SetCandidateSignature(oldAppln.CandidateSignature).
		SetCadre(oldAppln.Cadre).
		SetCandidateRemarks(oldAppln.CandidateRemarks).
		SetCategoryCode(oldAppln.CategoryCode).
		SetCategoryDescription(oldAppln.CategoryDescription).
		SetCenterFacilityId(oldAppln.CenterFacilityId).
		SetCenterId(oldAppln.CenterId).
		SetCentrePreference(oldAppln.CentrePreference).
		SetClaimingQualifyingService(oldAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(oldAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(oldAppln.ControllingOfficeName).
		SetDCCS(oldAppln.DCCS).
		SetDOB(oldAppln.DOB).
		SetDeputationControllingOfficeID(oldAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(oldAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(oldAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(oldAppln.DeputationOfficeName).
		SetDeputationOfficePincode(oldAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(oldAppln.DeputationOfficeUniqueId).
		SetInDeputation(oldAppln.InDeputation).
		SetDeputationType(oldAppln.DeputationType).
		SetDesignationID(oldAppln.DesignationID).
		SetDisabilityPercentage(oldAppln.DisabilityPercentage).
		SetDisabilityTypeCode(oldAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(oldAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(oldAppln.DisabilityTypeID).
		SetEducationCode(oldAppln.EducationCode).
		SetEducationDescription(oldAppln.EducationDescription).
		SetEmailID(oldAppln.EmailID).
		SetEmployeeID(oldAppln.EmployeeID).
		SetEmployeeName(oldAppln.EmployeeName).
		SetEmployeePost(oldAppln.EmployeePost).
		SetEntryPostCode(oldAppln.EntryPostCode).
		SetEntryPostDescription(oldAppln.EntryPostDescription).
		SetExamCode(oldAppln.ExamCode).
		SetExamName(oldAppln.ExamName).
		SetExamShortName(oldAppln.ExamShortName).
		SetExamYear(oldAppln.ExamYear).
		SetExamCityCenterCode(oldAppln.CenterId).
		SetFacilityUniqueID(oldAppln.FacilityUniqueID).
		SetFeederPostCode(oldAppln.FeederPostCode).
		SetFeederPostDescription(oldAppln.FeederPostDescription).
		SetFeederPostJoiningDate(oldAppln.FeederPostJoiningDate).
		SetGender(oldAppln.Gender).
		SetGDSEngagement(oldAppln.GDSEngagement).
		SetGenerateHallTicketFlag(newAppln.GenerateHallTicketFlag).
		SetLienControllingOfficeID(oldAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(oldAppln.LienControllingOfficeName).
		SetMobileNumber(oldAppln.MobileNumber).
		SetNodalOfficeFacilityID(oldAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(oldAppln.NodalOfficeName).
		SetOptionUsed(oldAppln.OptionUsed).
		SetPhoto(oldAppln.Photo).
		SetPhotoPath(oldAppln.PhotoPath).
		SetPostPreferences(oldAppln.PostPreferences).
		SetPresentDesignation(oldAppln.PresentDesignation).
		SetPresentPostCode(oldAppln.PresentPostCode).
		SetPresentPostDescription(oldAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(oldAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(oldAppln.ReportingOfficeName).
		SetServiceLength(oldAppln.ServiceLength).
		SetSignature(oldAppln.Signature).
		SetSignaturePath(oldAppln.SignaturePath).
		SetStatus("active").
		SetTempHallTicket(oldAppln.TempHallTicket).
		SetUnitPreferences(oldAppln.UnitPreferences).
		SetUserID(oldAppln.UserID).
		SetWorkingOfficeCircleFacilityID(oldAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(oldAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(oldAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(oldAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(oldAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(oldAppln.WorkingOfficeName).
		SetWorkingOfficePincode(oldAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(oldAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(oldAppln.WorkingOfficeRegionName).
		SetPunishmentStatus(newAppln.PunishmentStatus).             //here added punishmentstatus
		SetDisciplinaryCaseStatus(newAppln.DisciplinaryCaseStatus). //here added DisciplinaryCaseStatus
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if nonQualifyService == "Yes" {
		_, err := updatedAppln.Update().
			SetNonQualifyingService(*newAppln.NonQualifyingService).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	return updatedAppln, nil
}
func handleRecommendations(ctx context.Context, tx *ent.Tx, updatedAppln *ent.Exam_Applications_PS, newAppln *ca_reg.VerifyApplicationGroupB) error {
	currentTime := time.Now().Truncate(time.Second)

	recommendationsRef := make([]*ent.RecommendationsPSApplications, len(newAppln.Edges.ApplicationData))
	for i, recommendation := range newAppln.Edges.ApplicationData {
		if recommendation.VacancyYear == 0 {
			return fmt.Errorf("recommendations value at index %d is nil", i)
		}
		RecommendationsRefEntity, err := tx.RecommendationsPSApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(recommendation.CA_Recommendations).
			SetNORecommendations(recommendation.CA_Recommendations).
			SetCAUserName(newAppln.CA_UserName).
			SetCARemarks(recommendation.CA_Remarks).
			SetCAUpdatedAt(currentTime).
			SetNOUpdatedAt(currentTime).
			SetApplicationStatus("VerifiedRecommendationsByCA").
			Save(ctx)
		if err != nil {
			return err
		}
		recommendationsRef[i] = RecommendationsRefEntity
	}

	_, err := updatedAppln.Update().
		ClearPSApplicationsRef().
		AddPSApplicationsRef(recommendationsRef...).
		Save(ctx)
	return err
}
func handlePmpaRecommendations(ctx context.Context, tx *ent.Tx, updatedAppln *ent.Exam_Applications_PMPA, newAppln *ca_reg.VerifyApplicationPMPA) error {
	currentTime := time.Now().Truncate(time.Second)

	recommendationsRef := make([]*ent.RecommendationsPMPAApplications, len(newAppln.Edges.ApplicationData))
	for i, recommendation := range newAppln.Edges.ApplicationData {
		if recommendation.VacancyYear == 0 {
			return fmt.Errorf("recommendations value at index %d is nil", i)
		}

		RecommendationsRefEntity, err := tx.RecommendationsPMPAApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(recommendation.CA_Recommendations).
			SetNORecommendations(recommendation.CA_Recommendations).
			SetCAUserName(newAppln.CA_UserName).
			SetCARemarks(recommendation.CA_Remarks).
			SetCAUpdatedAt(currentTime).
			SetNOUpdatedAt(currentTime).
			SetApplicationStatus("VerifiedRecommendationsByCA").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to save Recommendation: %v", err)
		}
		recommendationsRef[i] = RecommendationsRefEntity
	}

	_, err := updatedAppln.Update().
		ClearPMPAApplicationsRef().
		AddPMPAApplicationsRef(recommendationsRef...).
		Save(ctx)
	return err
}
func handleGdspaRecommendations(ctx context.Context, tx *ent.Tx, updatedAppln *ent.Exam_Applications_GDSPA, newAppln *ca_reg.VerifyApplicationGDStoPA) error {
	currentTime := time.Now().Truncate(time.Second)

	recommendationsRef := make([]*ent.RecommendationsGDSPAApplications, len(newAppln.Edges.ApplicationData))
	for i, recommendation := range newAppln.Edges.ApplicationData {
		if recommendation.VacancyYear == 0 {
			return fmt.Errorf("recommendations value at index %d is nil", i)
		}

		RecommendationsRefEntity, err := tx.RecommendationsGDSPAApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(recommendation.CA_Recommendations).
			SetNORecommendations(recommendation.CA_Recommendations).
			SetCAUserName(newAppln.CA_UserName).
			SetCARemarks(recommendation.CA_Remarks).
			SetCAUpdatedAt(currentTime).
			SetNOUpdatedAt(currentTime).
			SetApplicationStatus("VerifiedRecommendationsByCA").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to save Recommendation: %v", err)
		}
		recommendationsRef[i] = RecommendationsRefEntity
	}
	_, err := updatedAppln.Update().
		ClearGDSPAApplicationsRef().
		AddGDSPAApplicationsRef(recommendationsRef...).
		Save(ctx)
	return err
}

func handleMtspmRecommendations(ctx context.Context, tx *ent.Tx, updatedAppln *ent.Exam_Application_MTSPMMG, newAppln *ca_reg.VerifyApplicationMTSPM) error {
	currentTime := time.Now().Truncate(time.Second)

	recommendationsRef := make([]*ent.RecommendationsMTSPMMGApplications, len(newAppln.Edges.ApplicationData))
	for i, recommendation := range newAppln.Edges.ApplicationData {
		if recommendation.VacancyYear == 0 {
			return fmt.Errorf("recommendations value at index %d is nil", i)
		}
		RecommendationsRefEntity, err := tx.RecommendationsMTSPMMGApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(recommendation.CA_Recommendations).
			SetNORecommendations(recommendation.CA_Recommendations).
			SetCAUserName(newAppln.CA_UserName).
			SetCARemarks(recommendation.CA_Remarks).
			SetCAUpdatedAt(currentTime).
			SetNOUpdatedAt(currentTime).
			SetApplicationStatus("VerifiedRecommendationsByCA").
			Save(ctx)
		if err != nil {
			return err
		}
		recommendationsRef[i] = RecommendationsRefEntity
	}

	_, err := updatedAppln.Update().
		ClearMTSPMMGApplicationsRef().
		AddMTSPMMGApplicationsRef(recommendationsRef...).
		Save(ctx)
	return err
}
func handleGDSpmRecommendations(ctx context.Context, tx *ent.Tx, updatedAppln *ent.Exam_Applications_GDSPM, newAppln *ca_reg.VerifyApplicationGDSPM) error {
	currentTime := time.Now().Truncate(time.Second)

	recommendationsRef := make([]*ent.RecommendationsGDSPMApplications, len(newAppln.Edges.ApplicationDataV))
	for i, recommendation := range newAppln.Edges.ApplicationDataV {
		if recommendation.VacancyYear == 0 {
			return fmt.Errorf("recommendations value at index %d is nil", i)
		}
		RecommendationsRefEntity, err := tx.RecommendationsGDSPMApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamYear(updatedAppln.ExamYear).
			SetVacancyYear(recommendation.VacancyYear).
			SetCARecommendations(recommendation.CA_Recommendations).
			SetNORecommendations(recommendation.CA_Recommendations).
			SetCAUserName(newAppln.CA_UserName).
			SetCARemarks(recommendation.CA_Remarks).
			SetPost(recommendation.Post). // Set the Post field
			SetCAUpdatedAt(currentTime).
			SetNOUpdatedAt(currentTime).
			SetApplicationStatus("VerifiedRecommendationsByCA").
			Save(ctx)
		if err != nil {
			return err
		}
		recommendationsRef[i] = RecommendationsRefEntity
	}
	_, err := updatedAppln.Update().
		ClearGDSPMApplicationsRef().
		AddGDSPMApplicationsRef(recommendationsRef...).
		Save(ctx)
	return err
}
func validateApplicationInputs(newAppln *ca_reg.ResubmitApplicationIp) error {
	if !isNumeric(newAppln.TempHallTicket) {
		return fmt.Errorf("issue for employee %d with temp hall ticket number: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
	}
	if len(newAppln.TempHallTicket) != 8 {
		return fmt.Errorf("issue for employee %d with temp hall ticket number length issue: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
	}
	if newAppln.EmployeeID == 0 {
		return errors.New("please enter valid employee id ")
	}
	return nil
}
func fetchOldApplication(ctx context.Context, tx *ent.Tx, newAppln *ca_reg.ResubmitApplicationIp) (*ent.Exam_Applications_IP, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(newAppln.EmployeeID),
			exam_applications_ip.ExamYearEQ(newAppln.ExamYear),
			exam_applications_ip.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application found for this candidate")
		}
		return nil, 500, " -SUB002", err
	}

	return oldAppln, 200, "", nil
}
func fetchOldApplicationGdsPm(ctx context.Context, tx *ent.Tx, newAppln *ca_reg.ReApplicationGDSPM) (*ent.Exam_Applications_GDSPM, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_GDSPM.
		Query().
		Where(
			exam_applications_gdspm.EmployeeIDEQ(newAppln.EmployeeID),
			exam_applications_gdspm.ExamYearEQ(newAppln.ExamYear),
			exam_applications_gdspm.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application found for this candidate")
		}
		return nil, 500, " -SUB002", err
	}

	return oldAppln, 200, "", nil
}
func fetchOldApplicationPmPa(ctx context.Context, tx *ent.Tx, newAppln *ca_reg.ReApplicationPMPA) (*ent.Exam_Applications_PMPA, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_PMPA.
		Query().
		Where(
			exam_applications_pmpa.EmployeeIDEQ(newAppln.EmployeeID),
			exam_applications_pmpa.ExamYearEQ(newAppln.ExamYear),
			exam_applications_pmpa.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application found for this candidate")
		}
		return nil, 500, " -SUB002", err
	}

	return oldAppln, 200, "", nil
}
func fetchOldApplicationGdsPa(ctx context.Context, tx *ent.Tx, newAppln *ca_reg.ReApplicationGDStoPA) (*ent.Exam_Applications_GDSPA, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_GDSPA.
		Query().
		Where(
			exam_applications_gdspa.EmployeeIDEQ(newAppln.EmployeeID),
			exam_applications_gdspa.ExamYearEQ(newAppln.ExamYear),
			exam_applications_gdspa.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application found for this candidate")
		}
		return nil, 500, " -SUB002", err
	}

	return oldAppln, 200, "", nil
}
func fetchOldApplicationMtsPm(ctx context.Context, tx *ent.Tx, newAppln *ca_reg.ReApplicationGDSPM) (*ent.Exam_Application_MTSPMMG, int32, string, error) {
	oldAppln, err := tx.Exam_Application_MTSPMMG.
		Query().
		Where(
			exam_application_mtspmmg.EmployeeIDEQ(newAppln.EmployeeID),
			exam_application_mtspmmg.ExamYearEQ(newAppln.ExamYear),
			exam_application_mtspmmg.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application found for this candidate")
		}
		return nil, 500, " -SUB002", err
	}

	return oldAppln, 200, "", nil
}

func fetchOldApplicationPs(ctx context.Context, tx *ent.Tx, newAppln *ca_reg.ReApplicationGroupB) (*ent.Exam_Applications_PS, int32, string, error) {
	oldAppln, err := tx.Exam_Applications_PS.
		Query().
		Where(
			exam_applications_ps.EmployeeIDEQ(newAppln.EmployeeID),
			exam_applications_ps.ExamYearEQ(newAppln.ExamYear),
			exam_applications_ps.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -SUB001", errors.New("no active application found for this candidate")
		}
		return nil, 500, " -SUB002", err
	}

	return oldAppln, 200, "", nil
}

func processResubmission(ctx context.Context, tx *ent.Tx, oldAppln interface{}, newAppln interface{}, examCode int) (interface{}, int32, string, error) {
	switch examCode {
	case 2:
		// Handle examapplicationip
		oldApplnIP := oldAppln.(*ent.Exam_Applications_IP)
		newApplnIP := newAppln.(*ca_reg.ResubmitApplicationIp)

		if oldApplnIP.ApplicationStatus == "VerifiedByNA" || oldApplnIP.ApplicationStatus == "VerifiedByCA" {
			return nil, 422, " -SUB001", fmt.Errorf("the Application was already verified By Nodal Authority/ Controlling Authority")
		}

		if oldApplnIP.ApplicationStatus != "PendingWithCandidate" {
			return nil, 422, " -SUB002", fmt.Errorf("this application was not in pending with candidate status")
		}

		// Validate circle data
		if err := validateCircleData(newApplnIP); err != nil {
			return nil, 422, " -SUB003", fmt.Errorf("circle preference values are missing! Please provide Circle preferences")
		}

		// Generate new application number
		applicationNumber, err := util.GenerateApplicationNumber(tx.Client(), newApplnIP.EmployeeID, newApplnIP.ExamYear, "IP")
		if err != nil {
			return nil, 500, " -SUB004", err
		}

		// Update old application status to inactive
		if err := updateOldApplicationStatus(ctx, oldApplnIP); err != nil {
			return nil, 500, " -SUB005", err
		}
		updatedAppln, err := createNewResubmitApplication(ctx, tx, oldApplnIP, newApplnIP, applicationNumber)
		if err != nil {
			return nil, 500, " -SUB006", err
		}
		return updatedAppln, 200, "", nil

	case 1:
		// Handle examapplicationps
		oldApplnPS := oldAppln.(*ent.Exam_Applications_PS)
		newApplnPS := newAppln.(*ca_reg.ReApplicationGroupB)
		if oldApplnPS.ApplicationStatus == "VerifiedByNA" || oldApplnPS.ApplicationStatus == "VerifiedByCA" {
			return nil, 422, " -SUB001", fmt.Errorf("the Application was already verified By Nodal Authority/ Controlling Authority")
		}

		if oldApplnPS.ApplicationStatus != "PendingWithCandidate" {
			return nil, 422, " -SUB002", fmt.Errorf("this application was not in pending with candidate status")
		}
		// Generate new application number
		applicationNumber, err := util.GenerateApplicationNumber(tx.Client(), newApplnPS.EmployeeID, newApplnPS.ExamYear, "PS")
		if err != nil {
			return nil, 500, " -SUB004", err
		}

		if err := updateOldPSApplicationStatus(ctx, oldApplnPS); err != nil {
			return nil, 500, " -SUB005", err
		}
		updatedApplnPS, err := createNewResubmitPsApplication(ctx, tx, oldApplnPS, newApplnPS, applicationNumber)
		if err != nil {
			return nil, 500, " -SUB006", err
		}

		return updatedApplnPS, 200, "", nil
	case 6:
		// Handle examapplicationip
		oldApplnIP := oldAppln.(*ent.Exam_Applications_GDSPM)
		newApplnIP := newAppln.(*ca_reg.ReApplicationGDSPM)

		if oldApplnIP.ApplicationStatus == "VerifiedByNA" || oldApplnIP.ApplicationStatus == "VerifiedByCA" {
			return nil, 422, " -SUB001", fmt.Errorf("the Application was already verified By Nodal Authority/ Controlling Authority")
		}

		if oldApplnIP.ApplicationStatus != "PendingWithCandidate" {
			return nil, 422, " -SUB002", fmt.Errorf("this application was not in pending with candidate status")
		}

		// Generate new application number
		applicationNumber, err := util.GenerateApplicationNumber(tx.Client(), newApplnIP.EmployeeID, newApplnIP.ExamYear, "GDSPM")
		if err != nil {
			return nil, 500, " -SUB004", err
		}

		// Update old application status to inactive
		if err := updateOldGdsPmApplicationStatus(ctx, oldApplnIP); err != nil {
			return nil, 500, " -SUB005", err
		}
		updatedAppln, err := createGdsPmResubmitApplication(ctx, tx, oldApplnIP, newApplnIP, applicationNumber)
		if err != nil {
			return nil, 500, " -SUB006", err
		}
		return updatedAppln, 200, "", nil
	case 5:
		// Handle examapplicationip
		oldApplnIP := oldAppln.(*ent.Exam_Application_MTSPMMG)
		newApplnIP := newAppln.(*ca_reg.ReApplicationGDSPM)

		if oldApplnIP.ApplicationStatus == "VerifiedByNA" || oldApplnIP.ApplicationStatus == "VerifiedByCA" {
			return nil, 422, " -SUB001", fmt.Errorf("the Application was already verified By Nodal Authority/ Controlling Authority")
		}

		if oldApplnIP.ApplicationStatus != "PendingWithCandidate" {
			return nil, 422, " -SUB002", fmt.Errorf("this application was not in pending with candidate status")
		}

		// Generate new application number
		applicationNumber, err := util.GenerateApplicationNumber(tx.Client(), newApplnIP.EmployeeID, newApplnIP.ExamYear, "MTSPM")
		if err != nil {
			return nil, 500, " -SUB004", err
		}

		// Update old application status to inactive
		if err := updateOldMtsPmApplicationStatus(ctx, oldApplnIP); err != nil {
			return nil, 500, " -SUB005", err
		}
		updatedAppln, err := createMtsPmResubmitApplication(ctx, tx, oldApplnIP, newApplnIP, applicationNumber)
		if err != nil {
			return nil, 500, " -SUB006", err
		}
		return updatedAppln, 200, "", nil
	case 4:
		// Handle examapplicationip
		oldApplnIP := oldAppln.(*ent.Exam_Applications_GDSPA)
		newApplnIP := newAppln.(*ca_reg.ReApplicationGDStoPA)

		if oldApplnIP.ApplicationStatus == "VerifiedByNA" || oldApplnIP.ApplicationStatus == "VerifiedByCA" {
			return nil, 422, " -SUB001", fmt.Errorf("the Application was already verified By Nodal Authority/ Controlling Authority")
		}

		if oldApplnIP.ApplicationStatus != "PendingWithCandidate" {
			return nil, 422, " -SUB002", fmt.Errorf("this application was not in pending with candidate status")
		}

		// Generate new application number
		applicationNumber, err := util.GenerateApplicationNumber(tx.Client(), newApplnIP.EmployeeID, newApplnIP.ExamYear, "GDSPA")
		if err != nil {
			return nil, 500, " -SUB004", err
		}

		// Update old application status to inactive
		if err := updateOldGdsPaApplicationStatus(ctx, oldApplnIP); err != nil {
			return nil, 500, " -SUB005", err
		}
		updatedAppln, err := createGdsPaResubmitApplication(ctx, tx, oldApplnIP, newApplnIP, applicationNumber)
		if err != nil {
			return nil, 500, " -SUB006", err
		}
		return updatedAppln, 200, "", nil

	case 3:
		// Handle examapplicationip
		oldApplnIP := oldAppln.(*ent.Exam_Applications_PMPA)
		newApplnIP := newAppln.(*ca_reg.ReApplicationPMPA)

		if oldApplnIP.ApplicationStatus == "VerifiedByNA" || oldApplnIP.ApplicationStatus == "VerifiedByCA" {
			return nil, 422, " -SUB001", fmt.Errorf("the Application was already verified By Nodal Authority/ Controlling Authority")
		}

		if oldApplnIP.ApplicationStatus != "PendingWithCandidate" {
			return nil, 422, " -SUB002", fmt.Errorf("this application was not in pending with candidate status")
		}

		// Generate new application number
		applicationNumber, err := util.GenerateApplicationNumber(tx.Client(), newApplnIP.EmployeeID, newApplnIP.ExamYear, "PMPA")
		if err != nil {
			return nil, 500, " -SUB004", err
		}

		// Update old application status to inactive
		if err := updateOldPmPaApplicationStatus(ctx, oldApplnIP); err != nil {
			return nil, 500, " -SUB005", err
		}
		updatedAppln, err := createPmPaResubmitApplication(ctx, tx, oldApplnIP, newApplnIP, applicationNumber)
		if err != nil {
			return nil, 500, " -SUB006", err
		}
		return updatedAppln, 200, "", nil

	default:
		return nil, 400, " -SUB007", fmt.Errorf("invalid examCode")
	}
}

func validateCircleData(newAppln *ca_reg.ResubmitApplicationIp) error {
	if newAppln.Edges.CircleData == nil || len(newAppln.Edges.CircleData) == 0 {
		return errors.New("circle preference values are missing! Please provide Circle preferences")
	}
	if len(newAppln.Edges.CircleData) != 23 {
		return errors.New("invalid number of Circle preferences. Must provide preferences for all 23 circles")
	}
	return nil
}

func updateOldApplicationStatus(ctx context.Context, oldAppln *ent.Exam_Applications_IP) error {
	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err := oldAppln.
		Update().
		SetStatus(stat).
		Save(ctx)
	return err
}
func updateOldGdsPmApplicationStatus(ctx context.Context, oldAppln *ent.Exam_Applications_GDSPM) error {
	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err := oldAppln.
		Update().
		SetStatus(stat).
		Save(ctx)
	return err
}
func updateOldGdsPaApplicationStatus(ctx context.Context, oldAppln *ent.Exam_Applications_GDSPA) error {
	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err := oldAppln.
		Update().
		SetStatus(stat).
		Save(ctx)
	return err
}
func updateOldPmPaApplicationStatus(ctx context.Context, oldAppln *ent.Exam_Applications_PMPA) error {
	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err := oldAppln.
		Update().
		SetStatus(stat).
		Save(ctx)
	return err
}
func updateOldMtsPmApplicationStatus(ctx context.Context, oldAppln *ent.Exam_Application_MTSPMMG) error {
	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err := oldAppln.
		Update().
		SetStatus(stat).
		Save(ctx)
	return err
}
func updateOldPSApplicationStatus(ctx context.Context, oldAppln *ent.Exam_Applications_PS) error {
	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err := oldAppln.
		Update().
		SetStatus(stat).
		Save(ctx)
	return err
}

func createNewResubmitApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_IP, newAppln *ca_reg.ResubmitApplicationIp, applicationNumber string) (*ent.Exam_Applications_IP, error) {
	currentTime := time.Now().UTC().Add(5*time.Hour + 30*time.Minute).Truncate(time.Second)
	updatedAppln, err := tx.Exam_Applications_IP.
		Create().
		SetApplicationNumber(applicationNumber).
		SetApplicationStatus("ResubmitCAVerificationPending").
		SetApplnSubmittedDate(currentTime).
		SetCAEmployeeDesignation(oldAppln.CAEmployeeDesignation).
		SetCADate(oldAppln.CADate).
		SetCAEmployeeID(oldAppln.CAEmployeeID).
		SetCAGeneralRemarks(oldAppln.CAGeneralRemarks).
		SetCAUserId(oldAppln.CAUserId).
		SetCAUserName(oldAppln.CAUserName).
		SetCadre(newAppln.Cadre).
		SetCandidateRemarks(newAppln.CandidateRemarks).
		SetCategoryCode(newAppln.CategoryCode).
		SetCategoryDescription(newAppln.CategoryDescription).
		SetCenterFacilityId(newAppln.CenterFacilityId).
		SetCenterId(newAppln.CenterId).
		SetCentrePreference(newAppln.CentrePreference).
		SetCentrePreference(newAppln.CentrePreference).
		SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(newAppln.ControllingOfficeName).
		SetDCCS(newAppln.DCCS).
		SetDOB(newAppln.DOB).
		SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(newAppln.DeputationOfficeName).
		SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
		SetInDeputation(newAppln.InDeputation).
		SetDeputationType(newAppln.DeputationType).
		SetDesignationID(newAppln.DesignationID).
		SetDisabilityPercentage(newAppln.DisabilityPercentage).
		SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(newAppln.DisabilityTypeID).
		SetEducationCode(newAppln.EducationCode).
		SetEducationDescription(newAppln.EducationDescription).
		SetEmailID(newAppln.EmailID).
		SetEmployeeID(newAppln.EmployeeID).
		SetEmployeeName(newAppln.EmployeeName).
		SetEntryPostCode(newAppln.EntryPostCode).
		SetEntryPostDescription(newAppln.EntryPostDescription).
		SetExamCode(newAppln.ExamCode).
		SetExamName(newAppln.ExamName).
		SetExamShortName(newAppln.ExamShortName).
		SetExamYear(newAppln.ExamYear).
		SetExamCityCenterCode(newAppln.CenterId).
		SetFacilityUniqueID(newAppln.FacilityUniqueID).
		SetFeederPostCode(newAppln.FeederPostCode).
		SetFeederPostDescription(newAppln.FeederPostDescription).
		SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
		SetGender(newAppln.Gender).
		SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
		SetMobileNumber(newAppln.MobileNumber).
		SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(newAppln.NodalOfficeName).
		SetPhoto(newAppln.Photo).
		SetPhotoPath(newAppln.PhotoPath).
		SetPresentDesignation(newAppln.PresentDesignation).
		SetPresentPostCode(newAppln.PresentPostCode).
		SetPresentPostDescription(newAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(newAppln.ReportingOfficeName).
		SetServiceLength(*newAppln.ServiceLength).
		SetSignature(newAppln.Signature).
		SetSignaturePath(newAppln.SignaturePath).
		SetTempHallTicket(newAppln.TempHallTicket).
		SetCandidatePhoto(newAppln.CandidatePhoto).
		SetCandidateSignature(newAppln.CandidateSignature).
		SetUserID(newAppln.UserID).
		SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(newAppln.WorkingOfficeName).
		SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	// Save the PlaceOfPreferenceIP records
	if err := saveCirclePreferences(ctx, tx, updatedAppln.ID, newAppln); err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createGdsPmResubmitApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_GDSPM, newAppln *ca_reg.ReApplicationGDSPM, applicationNumber string) (*ent.Exam_Applications_GDSPM, error) {
	currentTime := time.Now().UTC().Add(5*time.Hour + 30*time.Minute).Truncate(time.Second)
	updatedAppln, err := tx.Exam_Applications_GDSPM.
		Create().
		SetApplicationNumber(applicationNumber).
		SetApplicationStatus("ResubmitCAVerificationPending").
		SetApplnSubmittedDate(currentTime).
		SetCAEmployeeDesignation(oldAppln.CAEmployeeDesignation).
		SetCADate(oldAppln.CADate).
		SetCAEmployeeID(oldAppln.CAEmployeeID).
		SetCAGeneralRemarks(oldAppln.CAGeneralRemarks).
		SetCAUserId(oldAppln.CAUserId).
		SetCAUserName(oldAppln.CAUserName).
		SetCadre(newAppln.Cadre).
		SetCandidateRemarks(newAppln.CandidateRemarks).
		SetCategoryCode(newAppln.CategoryCode).
		SetCategoryDescription(newAppln.CategoryDescription).
		SetCenterFacilityId(newAppln.CenterFacilityId).
		SetCenterId(newAppln.CenterId).
		SetCentrePreference(newAppln.CentrePreference).
		SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(newAppln.ControllingOfficeName).
		SetDCCS(newAppln.DCCS).
		SetDOB(newAppln.DOB).
		SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(newAppln.DeputationOfficeName).
		SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
		SetInDeputation(newAppln.InDeputation).
		SetDeputationType(newAppln.DeputationType).
		SetDesignationID(newAppln.DesignationID).
		SetDisabilityPercentage(newAppln.DisabilityPercentage).
		SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(newAppln.DisabilityTypeID).
		SetEducationCode(newAppln.EducationCode).
		SetEducationDescription(newAppln.EducationDescription).
		SetEmailID(newAppln.EmailID).
		SetEmployeeID(newAppln.EmployeeID).
		SetEmployeeName(newAppln.EmployeeName).
		SetEmployeePost(newAppln.EmployeePost).
		SetEntryPostCode(newAppln.EntryPostCode).
		SetEntryPostDescription(newAppln.EntryPostDescription).
		SetExamCode(newAppln.ExamCode).
		SetExamShortName(newAppln.ExamShortName).
		SetExamName(newAppln.ExamName).
		SetExamYear(newAppln.ExamYear).
		SetExamCityCenterCode(newAppln.CenterId).
		SetFacilityUniqueID(newAppln.FacilityUniqueID).
		SetFeederPostCode(newAppln.FeederPostCode).
		SetFeederPostDescription(newAppln.FeederPostDescription).
		SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
		SetGender(newAppln.Gender).
		SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
		SetMobileNumber(newAppln.MobileNumber).
		SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(newAppln.NodalOfficeName).
		SetPhoto(newAppln.Photo).
		SetPhotoPath(newAppln.PhotoPath).
		SetPostPreferences(*newAppln.PostPreferences).
		SetPresentDesignation(newAppln.PresentDesignation).
		SetPresentPostCode(newAppln.PresentPostCode).
		SetPresentPostDescription(newAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(newAppln.ReportingOfficeName).
		SetServiceLength(*newAppln.ServiceLength).
		SetSignature(newAppln.Signature).
		SetSignaturePath(newAppln.SignaturePath).
		SetTempHallTicket(newAppln.TempHallTicket).
		SetUnitPreferences(*newAppln.UnitPreferences).
		SetUserID(newAppln.UserID).
		SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(newAppln.WorkingOfficeName).
		SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
		SetCandidatePhoto(newAppln.CandidatePhoto).
		SetCandidateSignature(newAppln.CandidateSignature).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createGdsPaResubmitApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_GDSPA, newAppln *ca_reg.ReApplicationGDStoPA, applicationNumber string) (*ent.Exam_Applications_GDSPA, error) {
	currentTime := time.Now().UTC().Add(5*time.Hour + 30*time.Minute).Truncate(time.Second)
	updatedAppln, err := tx.Exam_Applications_GDSPA.
		Create(). //modifeid
		SetApplicationNumber(applicationNumber).
		SetApplicationStatus("ResubmitCAVerificationPending").
		SetApplnSubmittedDate(currentTime).
		SetCAEmployeeDesignation(oldAppln.CAEmployeeDesignation).
		SetCADate(oldAppln.CADate).
		SetCAEmployeeID(oldAppln.CAEmployeeID).
		SetCAGeneralRemarks(oldAppln.CAGeneralRemarks).
		SetCAUserId(oldAppln.CAUserId).
		SetCAUserName(oldAppln.CAUserName).
		SetCadre(newAppln.Cadre).
		SetCandidateRemarks(newAppln.CandidateRemarks).
		SetCategoryCode(newAppln.CategoryCode).
		SetCategoryDescription(newAppln.CategoryDescription).
		SetCenterFacilityId(newAppln.CenterFacilityId).
		SetCenterId(newAppln.CenterId).
		SetCentrePreference(newAppln.CentrePreference).
		SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(newAppln.ControllingOfficeName).
		SetDCCS(newAppln.DCCS).
		SetDOB(newAppln.DOB).
		SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(newAppln.DeputationOfficeName).
		SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
		SetInDeputation(newAppln.InDeputation).
		SetDeputationType(newAppln.DeputationType).
		SetDesignationID(newAppln.DesignationID).
		SetDisabilityPercentage(newAppln.DisabilityPercentage).
		SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(newAppln.DisabilityTypeID).
		SetEducationCode(newAppln.EducationCode).
		SetEducationDescription(newAppln.EducationDescription).
		SetEmailID(newAppln.EmailID).
		SetEmployeeID(newAppln.EmployeeID).
		SetEmployeeName(newAppln.EmployeeName).
		SetEmployeePost(newAppln.EmployeePost).
		SetEntryPostCode(newAppln.EntryPostCode).
		SetEntryPostDescription(newAppln.EntryPostDescription).
		SetExamCode(newAppln.ExamCode).
		SetExamName(newAppln.ExamName).
		SetExamShortName(newAppln.ExamShortName).
		SetExamYear(newAppln.ExamYear).
		SetExamCityCenterCode(newAppln.CenterId).
		SetFacilityUniqueID(newAppln.FacilityUniqueID).
		SetFeederPostCode(newAppln.FeederPostCode).
		SetFeederPostDescription(newAppln.FeederPostDescription).
		SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
		SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
		SetGender(newAppln.Gender).
		SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
		SetMobileNumber(newAppln.MobileNumber).
		SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
		SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(newAppln.NodalOfficeName).
		SetPhoto(newAppln.Photo).
		SetPhotoPath(newAppln.PhotoPath).
		SetPostPreferences(*newAppln.PostPreferences).
		SetPresentDesignation(newAppln.PresentDesignation).
		SetPresentPostCode(newAppln.PresentPostCode).
		SetPresentPostDescription(newAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(newAppln.ReportingOfficeName).
		SetServiceLength(*newAppln.ServiceLength).
		SetSignature(newAppln.Signature).
		SetSignaturePath(newAppln.SignaturePath).
		SetStatus("active").
		SetSubdivisionOfficeFacilityID(newAppln.SubdivisionOfficeFacilityID).
		SetSubdivisionOfficeName(newAppln.SubdivisionOfficeName).
		SetTempHallTicket(newAppln.TempHallTicket).
		SetUnitPreferences(*newAppln.UnitPreferences).
		SetUserID(newAppln.UserID).
		SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(newAppln.WorkingOfficeName).
		SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
		SetCandidatePhoto(newAppln.CandidatePhoto).
		SetCandidateSignature(newAppln.CandidateSignature).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createPmPaResubmitApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_PMPA, newAppln *ca_reg.ReApplicationPMPA, applicationNumber string) (*ent.Exam_Applications_PMPA, error) {
	currentTime := time.Now().UTC().Add(5*time.Hour + 30*time.Minute).Truncate(time.Second)
	updatedAppln, err := tx.Exam_Applications_PMPA.
		Create().
		SetApplicationNumber(applicationNumber).
		SetApplicationStatus("ResubmitCAVerificationPending").
		SetApplnSubmittedDate(currentTime).
		SetCAEmployeeDesignation(oldAppln.CAEmployeeDesignation).
		SetCADate(oldAppln.CADate).
		SetCAEmployeeID(oldAppln.CAEmployeeID).
		SetCAGeneralRemarks(oldAppln.CAGeneralRemarks).
		SetCAUserId(oldAppln.CAUserId).
		SetCAUserName(oldAppln.CAUserName).
		SetCadre(newAppln.Cadre).
		SetCandidateRemarks(newAppln.CandidateRemarks).
		SetCategoryCode(newAppln.CategoryCode).
		SetCategoryDescription(newAppln.CategoryDescription).
		SetCenterFacilityId(newAppln.CenterFacilityId).
		SetCenterId(newAppln.CenterId).
		SetCentrePreference(newAppln.CentrePreference).
		SetCentrePreference(newAppln.CentrePreference).
		SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(newAppln.ControllingOfficeName).
		SetDCCS(newAppln.DCCS).
		SetDOB(newAppln.DOB).
		SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(newAppln.DeputationOfficeName).
		SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
		SetInDeputation(newAppln.InDeputation).
		SetDeputationType(newAppln.DeputationType).
		SetDesignationID(newAppln.DesignationID).
		SetDisabilityPercentage(newAppln.DisabilityPercentage).
		SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(newAppln.DisabilityTypeID).
		SetEducationCode(newAppln.EducationCode).
		SetEducationDescription(newAppln.EducationDescription).
		SetEmailID(newAppln.EmailID).
		SetEmployeeID(newAppln.EmployeeID).
		SetEmployeeName(newAppln.EmployeeName).
		SetEmployeePost(newAppln.EmployeePost).
		SetEntryPostCode(newAppln.EntryPostCode).
		SetEntryPostDescription(newAppln.EntryPostDescription).
		SetExamCode(newAppln.ExamCode).
		SetExamName(newAppln.ExamName).
		SetExamShortName(newAppln.ExamShortName).
		SetExamYear(newAppln.ExamYear).
		SetExamCityCenterCode(newAppln.CenterId).
		SetFacilityUniqueID(newAppln.FacilityUniqueID).
		SetFeederPostCode(newAppln.FeederPostCode).
		SetFeederPostDescription(newAppln.FeederPostDescription).
		SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
		SetGender(newAppln.Gender).
		SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
		SetMobileNumber(newAppln.MobileNumber).
		SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(newAppln.NodalOfficeName).
		SetPMMailGuardMTSEngagement(*newAppln.PMMailGuardMTSEngagement).
		SetPhoto(newAppln.Photo).
		SetPhotoPath(newAppln.PhotoPath).
		SetCandidatePhoto(newAppln.CandidatePhoto).
		SetCandidateSignature(newAppln.CandidateSignature).
		SetPostPreferences(*newAppln.PostPreferences).
		SetPresentDesignation(newAppln.PresentDesignation).
		SetPresentPostCode(newAppln.PresentPostCode).
		SetPresentPostDescription(newAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(newAppln.ReportingOfficeName).
		SetServiceLength(*newAppln.ServiceLength).
		SetSignature(newAppln.Signature).
		SetSignaturePath(newAppln.SignaturePath).
		SetStatus("active").
		SetTempHallTicket(newAppln.TempHallTicket).
		SetUnitPreferences(*newAppln.UnitPreferences).
		SetUserID(newAppln.UserID).
		SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(newAppln.WorkingOfficeName).
		SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
		SetCandidatePhoto(newAppln.CandidatePhoto).
		SetCandidateSignature(newAppln.CandidateSignature).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createMtsPmResubmitApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Application_MTSPMMG, newAppln *ca_reg.ReApplicationGDSPM, applicationNumber string) (*ent.Exam_Application_MTSPMMG, error) {
	currentTime := time.Now().UTC().Add(5*time.Hour + 30*time.Minute).Truncate(time.Second)
	updatedAppln, err := tx.Exam_Application_MTSPMMG.
		Create().
		SetApplicationNumber(applicationNumber).
		SetApplicationStatus("ResubmitCAVerificationPending").
		SetApplnSubmittedDate(currentTime).
		SetCAEmployeeDesignation(oldAppln.CAEmployeeDesignation).
		SetCADate(oldAppln.CADate).
		SetCAEmployeeID(oldAppln.CAEmployeeID).
		SetCAGeneralRemarks(oldAppln.CAGeneralRemarks).
		SetCAUserId(oldAppln.CAUserId).
		SetCAUserName(oldAppln.CAUserName).
		SetCadre(newAppln.Cadre).
		SetCandidateRemarks(newAppln.CandidateRemarks).
		SetCategoryCode(newAppln.CategoryCode).
		SetCategoryDescription(newAppln.CategoryDescription).
		SetCenterFacilityId(newAppln.CenterFacilityId).
		SetCenterId(newAppln.CenterId).
		SetCentrePreference(newAppln.CentrePreference).
		SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(newAppln.ControllingOfficeName).
		SetDCCS(newAppln.DCCS).
		SetDCInPresentCadre(newAppln.DCInPresentCadre).
		SetDOB(newAppln.DOB).
		SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(newAppln.DeputationOfficeName).
		SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
		SetInDeputation(newAppln.InDeputation).
		SetDeputationType(newAppln.DeputationType).
		SetDesignationID(newAppln.DesignationID).
		SetDisabilityPercentage(newAppln.DisabilityPercentage).
		SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(newAppln.DisabilityTypeID).
		SetEducationCode(newAppln.EducationCode).
		SetEducationDescription(newAppln.EducationDescription).
		SetEmailID(newAppln.EmailID).
		SetEmployeeID(newAppln.EmployeeID).
		SetEmployeeName(newAppln.EmployeeName).
		SetEmployeePost(newAppln.EmployeePost).
		SetEntryPostCode(newAppln.EntryPostCode).
		SetEntryPostDescription(newAppln.EntryPostDescription).
		SetExamCode(newAppln.ExamCode).
		SetExamName(newAppln.ExamName).
		SetExamShortName(newAppln.ExamShortName).
		SetExamYear(newAppln.ExamYear).
		SetExamCityCenterCode(newAppln.CenterId).
		SetFacilityUniqueID(newAppln.FacilityUniqueID).
		SetFeederPostCode(newAppln.FeederPostCode).
		SetFeederPostDescription(newAppln.FeederPostDescription).
		SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
		SetGender(newAppln.Gender).
		SetGDSEngagement(*newAppln.GDSEngagement).
		SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
		SetMobileNumber(newAppln.MobileNumber).
		SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(newAppln.NodalOfficeName).
		SetPhoto(newAppln.Photo).
		SetPhotoPath(newAppln.PhotoPath).
		SetPostPreferences(*newAppln.PostPreferences).
		SetPresentDesignation(newAppln.PresentDesignation).
		SetPresentPostCode(newAppln.PresentPostCode).
		SetPresentPostDescription(newAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(newAppln.ReportingOfficeName).
		SetServiceLength(*newAppln.ServiceLength).
		SetSignature(newAppln.Signature).
		SetSignaturePath(newAppln.SignaturePath).
		SetStatus("active").
		SetTempHallTicket(newAppln.TempHallTicket).
		SetUnitPreferences(*newAppln.UnitPreferences).
		SetUserID(newAppln.UserID).
		SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(newAppln.WorkingOfficeName).
		SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
		SetCandidatePhoto(newAppln.CandidatePhoto).
		SetCandidateSignature(newAppln.CandidateSignature).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}
func createNewResubmitPsApplication(ctx context.Context, tx *ent.Tx, oldAppln *ent.Exam_Applications_PS, newAppln *ca_reg.ReApplicationGroupB, applicationNumber string) (*ent.Exam_Applications_PS, error) {
	currentTime := time.Now().UTC().Add(5*time.Hour + 30*time.Minute).Truncate(time.Second)
	updatedAppln, err := tx.Exam_Applications_PS.
		Create().
		SetApplicationNumber(applicationNumber).
		SetApplicationStatus("ResubmitCAVerificationPending").
		SetApplnSubmittedDate(currentTime).
		SetCAEmployeeDesignation(oldAppln.CAEmployeeDesignation).
		SetCADate(oldAppln.CADate).
		SetCAEmployeeID(oldAppln.CAEmployeeID).
		SetCAGeneralRemarks(oldAppln.CAGeneralRemarks).
		SetCAUserId(oldAppln.CAUserId).
		SetCAUserName(oldAppln.CAUserName).
		SetCandidateRemarks(newAppln.CandidateRemarks).
		SetCategoryCode(newAppln.CategoryCode).
		SetCategoryDescription(newAppln.CategoryDescription).
		SetCenterFacilityId(newAppln.CenterFacilityId).
		SetCenterId(newAppln.CenterId).
		SetCentrePreference(newAppln.CentrePreference).
		SetCentrePreference(newAppln.CentrePreference).
		SetClaimingQualifyingService(newAppln.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(newAppln.ControllingOfficeFacilityID).
		SetControllingOfficeName(newAppln.ControllingOfficeName).
		SetDCCS(newAppln.DCCS).
		SetDOB(newAppln.DOB).
		SetDeputationControllingOfficeID(newAppln.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(newAppln.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(newAppln.DeputationOfficeFacilityID).
		SetDeputationOfficeName(newAppln.DeputationOfficeName).
		SetDeputationOfficePincode(newAppln.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(newAppln.DeputationOfficeUniqueId).
		SetInDeputation(newAppln.InDeputation).
		SetDeputationType(newAppln.DeputationType).
		SetDesignationID(newAppln.DesignationID).
		SetDisabilityPercentage(newAppln.DisabilityPercentage).
		SetDisabilityTypeCode(newAppln.DisabilityTypeCode).
		SetDisabilityTypeDescription(newAppln.DisabilityTypeDescription).
		SetDisabilityTypeID(newAppln.DisabilityTypeID).
		SetEducationCode(newAppln.EducationCode).
		SetEducationDescription(newAppln.EducationDescription).
		SetEmailID(newAppln.EmailID).
		SetEmployeeID(newAppln.EmployeeID).
		SetEmployeeName(newAppln.EmployeeName).
		SetEntryPostCode(newAppln.EntryPostCode).
		SetEntryPostDescription(newAppln.EntryPostDescription).
		SetExamCode(newAppln.ExamCode).
		SetExamName(newAppln.ExamName).
		SetExamShortName(newAppln.ExamShortName).
		SetExamYear(newAppln.ExamYear).
		SetExamCityCenterCode(newAppln.CenterId).
		SetFacilityUniqueID(newAppln.FacilityUniqueID).
		SetFeederPostCode(newAppln.FeederPostCode).
		SetFeederPostDescription(newAppln.FeederPostDescription).
		SetFeederPostJoiningDate(newAppln.FeederPostJoiningDate).
		SetGender(newAppln.Gender).
		SetLienControllingOfficeID(newAppln.LienControllingOfficeID).
		SetLienControllingOfficeName(newAppln.LienControllingOfficeName).
		SetMobileNumber(newAppln.MobileNumber).
		SetNodalOfficeFacilityID(newAppln.NodalOfficeFacilityID).
		SetNodalOfficeName(newAppln.NodalOfficeName).
		SetPhoto(newAppln.Photo).
		SetPhotoPath(newAppln.PhotoPath).
		SetPresentDesignation(newAppln.PresentDesignation).
		SetPresentPostCode(newAppln.PresentPostCode).
		SetPresentPostDescription(newAppln.PresentPostDescription).
		SetReportingOfficeFacilityID(newAppln.ReportingOfficeFacilityID).
		SetReportingOfficeName(newAppln.ReportingOfficeName).
		SetServiceLength(*newAppln.ServiceLength).
		SetSignature(newAppln.Signature).
		SetSignaturePath(newAppln.SignaturePath).
		SetTempHallTicket(newAppln.TempHallTicket).
		SetUserID(newAppln.UserID).
		SetWorkingOfficeCircleFacilityID(newAppln.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(newAppln.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(newAppln.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(newAppln.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(newAppln.WorkingOfficeFacilityID).
		SetWorkingOfficeName(newAppln.WorkingOfficeName).
		SetWorkingOfficePincode(newAppln.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(newAppln.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(newAppln.WorkingOfficeRegionName).
		SetCandidatePhoto(newAppln.CandidatePhoto).
		SetCandidateSignature(newAppln.CandidateSignature).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return updatedAppln, nil
}

func saveCirclePreferences(ctx context.Context, tx *ent.Tx, applicationID int64, newAppln *ca_reg.ResubmitApplicationIp) error {
	circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(newAppln.Edges.CircleData))
	currentTime := time.Now().UTC().Add(5*time.Hour + 30*time.Minute).Truncate(time.Second)

	for i, circlePrefRef := range newAppln.Edges.CircleData {
		if circlePrefRef.PlacePrefNo == 0 {
			return fmt.Errorf("circle preference value at index %d is nil", i)
		}

		circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
			Create().
			SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
			SetApplicationID(applicationID).
			SetEmployeeID(newAppln.EmployeeID).
			SetPlacePrefValue(circlePrefRef.PlacePrefValue).
			SetUpdatedAt(currentTime).
			Save(ctx)

		if err != nil {
			return err
		}
		circlePrefRefs[i] = circlePrefRefEntity
	}

	_, err := tx.Exam_Applications_IP.
		UpdateOneID(applicationID).
		AddCirclePrefRef(circlePrefRefs...).
		Save(ctx)
	return err
}
