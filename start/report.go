package start

import (
	"context"
	"errors"
	"fmt"
	"log"

	//"reflect"
	"time"

	"recruit/ent"
	"recruit/ent/center"
	"recruit/ent/employeemaster"
	"recruit/ent/errorlogs"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/exam_applications_gdspa"
	"recruit/ent/exam_applications_gdspm"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_pmpa"
	"recruit/ent/exam_applications_ps"

	//"recruit/util"

	"recruit/ent/facilitymasters"
	ca_reg "recruit/payloadstructure/candidate_registration"
)

func GetExamDetailsWithExamCode(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	// ctx := context.Background()
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}

	switch examCode {
	case 2:
		applications, err := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityID(circleFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no applications found for the provided details")
		}
		return applications, nil, nil, nil, nil, nil, 200, "", true, nil

	case 1:
		applications, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityID(circleFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR004", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR005", false, errors.New("no applications found for the provided details")
		}
		return nil, applications, nil, nil, nil, nil, 200, "", true, nil

	case 4:
		applications, err := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityID(circleFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR006", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR007", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, applications, nil, nil, nil, 200, "", true, nil

	case 3:
		applications, err := client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityID(circleFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR008", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR009", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, applications, nil, nil, 200, "", true, nil

	case 6:
		applications, err := client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityID(circleFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR010", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR011", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, applications, nil, 200, "", true, nil

	case 5:
		applications, err := client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityID(circleFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR012", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR013", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, nil, applications, 200, "", true, nil

	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, errors.New("invalid exam code")
	}
}

func SubGetExamDetailsWithExamCodeExamYearCA(ctx context.Context, client *ent.Client, examCode int32, examYear string, CAFacilityID string) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	// ctx := context.Background()
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}

	switch examCode {
	case 2:
		applications, err := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.ControllingOfficeFacilityIDEQ(CAFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no applications found for the provided details")
		}
		return applications, nil, nil, nil, nil, nil, 200, "", true, nil

	case 1:
		applications, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.ControllingOfficeFacilityIDEQ(CAFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR004", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR005", false, errors.New("no applications found for the provided details")
		}
		return nil, applications, nil, nil, nil, nil, 200, "", true, nil

	case 4:
		applications, err := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.ControllingOfficeFacilityIDEQ(CAFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR006", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR007", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, applications, nil, nil, nil, 200, "", true, nil

	case 3:
		applications, err := client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.ControllingOfficeFacilityIDEQ(CAFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR008", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR009", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, applications, nil, nil, 200, "", true, nil

	case 6:
		applications, err := client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.ControllingOfficeFacilityIDEQ(CAFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR010", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR011", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, applications, nil, 200, "", true, nil

	case 5:
		applications, err := client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(CAFacilityID),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR012", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR013", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, nil, applications, 200, "", true, nil

	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, errors.New("invalid exam code")
	}
}

func GetExamDetailsWithExamCodeAndStatus(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, statusCode int32) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}

	var status string
	switch statusCode {
	case 1:
		status = "Recommended"
	case 2:
		status = "Provisionally Recommended"
	case 3:
		status = "Not Recommended"
	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR015", false, errors.New("invalid Recommended status code")
	}

	switch examCode {
	case 2:
		applications, err := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ip.RecommendedStatusEQ(status),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no applications found for the provided details")
		}
		return applications, nil, nil, nil, nil, nil, 200, "", true, nil

	case 1:
		applications, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ps.RecommendedStatusEQ(status),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR004", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR005", false, errors.New("no applications found for the provided details")
		}
		return nil, applications, nil, nil, nil, nil, 200, "", true, nil

	case 4:
		applications, err := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_gdspa.RecommendedStatusEQ(status),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR006", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR007", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, applications, nil, nil, nil, 200, "", true, nil

	case 3:
		applications, err := client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_pmpa.RecommendedStatusEQ(status),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR008", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR009", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, applications, nil, nil, 200, "", true, nil

	case 6:
		applications, err := client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_gdspm.RecommendedStatusEQ(status),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR010", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR011", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, applications, nil, 200, "", true, nil

	case 5:
		applications, err := client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityID(circleFacilityID),
				exam_application_mtspmmg.RecommendedStatusEQ(status),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR012", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR013", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, nil, applications, 200, "", true, nil

	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, errors.New("invalid exam code")
	}
}

func GetExamDetailsWithExamCodeAndStatusNew(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, statusCode int32) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}

	var statuses []string
	switch statusCode {
	case 1:
		statuses = []string{"Recommended"}
	case 2:
		statuses = []string{"Provisionally Recommended"}
	case 3:
		statuses = []string{"Not Recommended"}
	case 4:
		statuses = []string{"Recommended", "Provisionally Recommended", "Not Recommended"}
	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR015", false, errors.New("invalid Recommended status code")
	}

	switch examCode {
	case 2:
		query := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityID(circleFacilityID),
			)
		if statusCode != 4 {
			query = query.Where(exam_applications_ip.RecommendedStatusIn(statuses...))
		} else {
			query = query.Where(exam_applications_ip.RecommendedStatusIn(statuses...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no applications found for the provided details")
		}
		return applications, nil, nil, nil, nil, nil, 200, "", true, nil

	case 1:
		query := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityID(circleFacilityID),
			)
		if statusCode != 4 {
			query = query.Where(exam_applications_ps.RecommendedStatusIn(statuses...))
		} else {
			query = query.Where(exam_applications_ps.RecommendedStatusIn(statuses...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR004", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR005", false, errors.New("no applications found for the provided details")
		}
		return nil, applications, nil, nil, nil, nil, 200, "", true, nil

	case 4:
		query := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityID(circleFacilityID),
			)
		if statusCode != 4 {
			query = query.Where(exam_applications_gdspa.RecommendedStatusIn(statuses...))
		} else {
			query = query.Where(exam_applications_gdspa.RecommendedStatusIn(statuses...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR006", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR007", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, applications, nil, nil, nil, 200, "", true, nil

	case 3:
		query := client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityID(circleFacilityID),
			)
		if statusCode != 4 {
			query = query.Where(exam_applications_pmpa.RecommendedStatusIn(statuses...))
		} else {
			query = query.Where(exam_applications_pmpa.RecommendedStatusIn(statuses...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR008", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR009", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, applications, nil, nil, 200, "", true, nil

	case 6:
		query := client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityID(circleFacilityID),
			)
		if statusCode != 4 {
			query = query.Where(exam_applications_gdspm.RecommendedStatusIn(statuses...))
		} else {
			query = query.Where(exam_applications_gdspm.RecommendedStatusIn(statuses...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR010", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR011", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, applications, nil, 200, "", true, nil

	case 5:
		query := client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityID(circleFacilityID),
			)
		if statusCode != 4 {
			query = query.Where(exam_application_mtspmmg.RecommendedStatusIn(statuses...))
		} else {
			query = query.Where(exam_application_mtspmmg.RecommendedStatusIn(statuses...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR012", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR013", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, nil, applications, 200, "", true, nil

	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, errors.New("invalid exam code")
	}
}
func GetExamDetailsWithExamCodeAndApplnStatus(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, statusCode int32) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}

	var statuses []string
	switch statusCode {
	case 1:
		statuses = []string{"CAVerificationPending", "ResubmitCAVerificationPending"}
	case 2:
		statuses = []string{"PendingWithCandidate"}
	case 3:
		statuses = []string{"VerifiedByCA", "VerifiedByNA"}
	case 4:
		statuses = []string{"VerifiedByCA"}
	case 5:
		statuses = []string{"VerifiedByNA"}
	case 6:
		statuses = []string{"CAVerificationPending", "ResubmitCAVerificationPending", "PendingWithCandidate", "VerifiedByCA", "VerifiedByNA"}
	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR015", false, errors.New("invalid Application status code")
	}

	switch examCode {
	case 2:
		applications, err := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ip.ApplicationStatusIn(statuses...),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no applications found for the provided details")
		}
		return applications, nil, nil, nil, nil, nil, 200, "", true, nil

	case 1:
		applications, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ps.ApplicationStatusIn(statuses...),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR004", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR005", false, errors.New("no applications found for the provided details")
		}
		return nil, applications, nil, nil, nil, nil, 200, "", true, nil

	case 4:
		applications, err := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_gdspa.ApplicationStatusIn(statuses...),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR006", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR007", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, applications, nil, nil, nil, 200, "", true, nil

	case 3:
		applications, err := client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_pmpa.ApplicationStatusIn(statuses...),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR008", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR009", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, applications, nil, nil, 200, "", true, nil

	case 6:
		applications, err := client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_gdspm.ApplicationStatusIn(statuses...),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR010", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR011", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, applications, nil, 200, "", true, nil

	case 5:
		applications, err := client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityID(circleFacilityID),
				exam_application_mtspmmg.ApplicationStatusIn(statuses...),
			).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR012", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR013", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, nil, applications, 200, "", true, nil

	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, errors.New("invalid exam code")
	}
}

func GetExamDetailsWithHallticktdetOld(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, divisionFacilityID string) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	/* ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel() */

	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}
	var centerCodes []int32
	if divisionFacilityID != "all" {
		centers, err := client.Center.Query().
			Where(center.ConductedByFacilityIDEQ(divisionFacilityID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}

		if len(centers) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no centers found for the provided details")
		}

		centerCodes = make([]int32, len(centers))
		for i, center := range centers {
			centerCodes[i] = center.ID
		}
	}
	switch examCode {
	case 2:
		query := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ip.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
				exam_applications_ip.HallTicketNumberNEQ(""),
				exam_applications_ip.ExamCityCenterCodeNEQ(0),
			).
			WithExamCentres()

		if divisionFacilityID != "all" {
			query = query.Where(exam_applications_ip.CenterCodeIn(centerCodes...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no applications found for the provided details")
		}
		return applications, nil, nil, nil, nil, nil, 200, "", true, nil

	case 1:
		query := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ps.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
				exam_applications_ps.HallTicketNumberNEQ(""),
				exam_applications_ps.ExamCityCenterCodeNEQ(0),
			).
			WithPSExamCentres()
		if divisionFacilityID != "all" {
			query = query.Where(exam_applications_ps.CenterCodeIn(centerCodes...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR004", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR005", false, errors.New("no applications found for the provided details")
		}
		return nil, applications, nil, nil, nil, nil, 200, "", true, nil

	case 4:
		query := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_gdspa.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
				exam_applications_gdspa.HallTicketNumberNEQ(""),
				exam_applications_gdspa.ExamCityCenterCodeNEQ(0),
			).
			WithGDSPAExamCentres()
		if divisionFacilityID != "all" {
			query = query.Where(exam_applications_gdspa.CenterCodeIn(centerCodes...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR006", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR007", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, applications, nil, nil, nil, 200, "", true, nil

	case 3:
		query := client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_pmpa.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
				exam_applications_pmpa.HallTicketNumberNEQ(""),
				exam_applications_pmpa.ExamCityCenterCodeNEQ(0),
			).
			WithPMPAExamCentres()
		if divisionFacilityID != "all" {
			query = query.Where(exam_applications_pmpa.CenterCodeIn(centerCodes...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR008", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR009", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, applications, nil, nil, 200, "", true, nil

	case 6:
		query := client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_gdspm.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
				exam_applications_gdspm.HallTicketNumberNEQ(""),
				exam_applications_gdspm.ExamCityCenterCodeNEQ(0),
			).
			WithGDSPMExamCentres()
		if divisionFacilityID != "all" {
			query = query.Where(exam_applications_gdspm.CenterCodeIn(centerCodes...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR010", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR011", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, applications, nil, 200, "", true, nil

	case 5:
		query := client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityID(circleFacilityID),
				exam_application_mtspmmg.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
				exam_application_mtspmmg.HallTicketNumberNEQ(""),
				exam_application_mtspmmg.ExamCityCenterCodeNEQ(0),
			).
			WithMTSPMMGExamCentres()
		if divisionFacilityID != "all" {
			query = query.Where(exam_application_mtspmmg.CenterCodeIn(centerCodes...))
		}
		applications, err := query.All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 422, " -STR012", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR013", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, nil, applications, 200, "", true, nil

	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, errors.New("invalid exam code")
	}
}

func SubGetExamDetailsAttendanceView(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, divisionFacilityID string) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	/* ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	*/
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}

	var centerCodes []int32
	if divisionFacilityID != "all" {
		centers, err := client.Center.Query().
			Where(center.ConductedByFacilityIDEQ(divisionFacilityID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}

		if len(centers) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no centers found for the provided details")
		}

		centerCodes = make([]int32, len(centers))
		for i, center := range centers {
			centerCodes[i] = center.ID
		}
	}

	var (
		ipApplications    []*ent.Exam_Applications_IP
		psApplications    []*ent.Exam_Applications_PS
		gdspaApplications []*ent.Exam_Applications_GDSPA
		pmpaApplications  []*ent.Exam_Applications_PMPA
		gdspmApplications []*ent.Exam_Applications_GDSPM
		mtspmApplications []*ent.Exam_Application_MTSPMMG
		err               error
	)

	switch examCode {
	case 2:
		ipApplications, err = getIPApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 1:
		psApplications, err = getPSApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 4:
		gdspaApplications, err = getGDSPAApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 3:
		pmpaApplications, err = getPMPAApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 6:
		gdspmApplications, err = getGDSPMApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 5:
		mtspmApplications, err = getMTSPMMGApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, errors.New("invalid exam code")
	}

	if err != nil {
		return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
	}

	return ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, 200, "", true, nil
}

func GetExamDetailsWithHallticktdet(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, divisionFacilityID string) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	/* ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	*/
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}

	var centerCodes []int32
	if divisionFacilityID != "all" {
		centers, err := client.Center.Query().
			Where(center.ConductedByFacilityIDEQ(divisionFacilityID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}

		if len(centers) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no centers found for the provided details")
		}

		centerCodes = make([]int32, len(centers))
		for i, center := range centers {
			centerCodes[i] = center.ID
		}
	}

	var (
		ipApplications    []*ent.Exam_Applications_IP
		psApplications    []*ent.Exam_Applications_PS
		gdspaApplications []*ent.Exam_Applications_GDSPA
		pmpaApplications  []*ent.Exam_Applications_PMPA
		gdspmApplications []*ent.Exam_Applications_GDSPM
		mtspmApplications []*ent.Exam_Application_MTSPMMG
		err               error
	)

	switch examCode {
	case 2:
		ipApplications, err = getIPApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 1:
		psApplications, err = getPSApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 4:
		gdspaApplications, err = getGDSPAApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 3:
		pmpaApplications, err = getPMPAApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 6:
		gdspmApplications, err = getGDSPMApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	case 5:
		mtspmApplications, err = getMTSPMMGApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR004", false, errors.New("invalid exam code")
	}

	if err != nil {
		return nil, nil, nil, nil, nil, nil, 500, " -STR005", false, err
	}

	return ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, 200, "", true, nil
}

func getIPApplications(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, divisionFacilityID string, centerCodes []int32) ([]*ent.Exam_Applications_IP, error) {
	var applications []*ent.Exam_Applications_IP
	var err error
	if divisionFacilityID != "all" {
		applications, err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ip.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
				exam_applications_ip.HallTicketNumberNEQ(""),
				exam_applications_ip.ExamCityCenterCodeNEQ(0),
				exam_applications_ip.CenterCodeIn(centerCodes...),
			).All(ctx)
	} else {
		applications, err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ip.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
				exam_applications_ip.HallTicketNumberNEQ(""),
				exam_applications_ip.ExamCityCenterCodeNEQ(0),
			).All(ctx)
	}
	if err != nil {
		return nil, err
	}

	if len(applications) == 0 {
		return nil, errors.New("no applications found for the provided details")
	}

	return applications, nil
}
func getPSApplications(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, divisionFacilityID string, centerCodes []int32) ([]*ent.Exam_Applications_PS, error) {
	query := client.Exam_Applications_PS.Query().
		Where(
			exam_applications_ps.ExamCodeEQ(examCode),
			exam_applications_ps.StatusEQ("active"),
			exam_applications_ps.ExamYearEQ(examYear),
			exam_applications_ps.NodalOfficeFacilityID(circleFacilityID),
			exam_applications_ps.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			exam_applications_ps.HallTicketNumberNEQ(""),
			exam_applications_ps.ExamCityCenterCodeNEQ(0),
		)

	if divisionFacilityID != "all" {
		query = query.Where(exam_applications_ps.ExamCityCenterCodeIn(centerCodes...))
	}

	applications, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	if len(applications) == 0 {
		return nil, errors.New("no applications found for the provided details")
	}

	return applications, nil
}
func getGDSPAApplications(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, divisionFacilityID string, centerCodes []int32) ([]*ent.Exam_Applications_GDSPA, error) {
	query := client.Exam_Applications_GDSPA.Query().
		Where(
			exam_applications_gdspa.ExamCodeEQ(examCode),
			exam_applications_gdspa.StatusEQ("active"),
			exam_applications_gdspa.ExamYearEQ(examYear),
			exam_applications_gdspa.NodalOfficeFacilityID(circleFacilityID),
			exam_applications_gdspa.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			exam_applications_gdspa.HallTicketNumberNEQ(""),
			exam_applications_gdspa.ExamCityCenterCodeNEQ(0),
		)

	if divisionFacilityID != "all" {
		query = query.Where(exam_applications_gdspa.ExamCityCenterCodeIn(centerCodes...))
	}

	applications, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	if len(applications) == 0 {
		return nil, errors.New("no applications found for the provided details")
	}

	return applications, nil
}
func getPMPAApplications(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, divisionFacilityID string, centerCodes []int32) ([]*ent.Exam_Applications_PMPA, error) {
	query := client.Exam_Applications_PMPA.Query().
		Where(
			exam_applications_pmpa.ExamCodeEQ(examCode),
			exam_applications_pmpa.StatusEQ("active"),
			exam_applications_pmpa.ExamYearEQ(examYear),
			exam_applications_pmpa.NodalOfficeFacilityID(circleFacilityID),
			exam_applications_pmpa.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			exam_applications_pmpa.HallTicketNumberNEQ(""),
			exam_applications_pmpa.ExamCityCenterCodeNEQ(0),
		)

	if divisionFacilityID != "all" {
		query = query.Where(exam_applications_pmpa.ExamCityCenterCodeIn(centerCodes...))
	}

	applications, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	if len(applications) == 0 {
		return nil, errors.New("no applications found for the provided details")
	}

	return applications, nil
}

func getGDSPMApplications(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, divisionFacilityID string, centerCodes []int32) ([]*ent.Exam_Applications_GDSPM, error) {
	query := client.Exam_Applications_GDSPM.Query().
		Where(
			exam_applications_gdspm.ExamCodeEQ(examCode),
			exam_applications_gdspm.StatusEQ("active"),
			exam_applications_gdspm.ExamYearEQ(examYear),
			exam_applications_gdspm.NodalOfficeFacilityID(circleFacilityID),
			exam_applications_gdspm.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			exam_applications_gdspm.HallTicketNumberNEQ(""),
			exam_applications_gdspm.ExamCityCenterCodeNEQ(0),
		)

	if divisionFacilityID != "all" {
		query = query.Where(exam_applications_gdspm.ExamCityCenterCodeIn(centerCodes...))
	}

	applications, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	if len(applications) == 0 {
		return nil, errors.New("no applications found for the provided details")
	}

	return applications, nil
}
func getMTSPMMGApplications(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string, divisionFacilityID string, centerCodes []int32) ([]*ent.Exam_Application_MTSPMMG, error) {
	query := client.Exam_Application_MTSPMMG.Query().
		Where(
			exam_application_mtspmmg.ExamCodeEQ(examCode),
			exam_application_mtspmmg.StatusEQ("active"),
			exam_application_mtspmmg.ExamYearEQ(examYear),
			exam_application_mtspmmg.NodalOfficeFacilityID(circleFacilityID),
			exam_application_mtspmmg.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			exam_application_mtspmmg.HallTicketNumberNEQ(""),
			exam_application_mtspmmg.ExamCityCenterCodeNEQ(0),
		)

	if divisionFacilityID != "all" {
		query = query.Where(exam_application_mtspmmg.ExamCityCenterCodeIn(centerCodes...))
	}

	applications, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	if len(applications) == 0 {
		return nil, errors.New("no applications found for the provided details")
	}

	return applications, nil
}

// Similar functions for PS, GDSPA, PMPA, GDSPM, MTSPMMG applications can be created as `getIPApplications`

// func startErrorHandlerWithoutLog(gctx *gin.Context, err error, status int32, stgError string, client *ent.Client, startFunction string) {
//     gctx.JSON(int(status), gin.H{"error": err.Error(), "stgError": stgError, "startFunction": startFunction})
// }

func SubGetExamDetailsHallticketviewDR(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	/* ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	*/
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}

	var (
		ipApplications    []*ent.Exam_Applications_IP
		psApplications    []*ent.Exam_Applications_PS
		gdspaApplications []*ent.Exam_Applications_GDSPA
		pmpaApplications  []*ent.Exam_Applications_PMPA
		gdspmApplications []*ent.Exam_Applications_GDSPM
		mtspmApplications []*ent.Exam_Application_MTSPMMG
		err               error
	)

	switch examCode {
	case 2:
		ipApplications, err = getIPApplicationsDR(ctx, client, examCode, examYear)
	// case 1:
	// 	psApplications, err = getPSApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	// case 4:
	// 	gdspaApplications, err = getGDSPAApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	// case 3:
	// 	pmpaApplications, err = getPMPAApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	// case 6:
	// 	gdspmApplications, err = getGDSPMApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	// case 5:
	// 	mtspmApplications, err = getMTSPMMGApplications(ctx, client, examCode, examYear, circleFacilityID, divisionFacilityID, centerCodes)
	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, errors.New("invalid exam code")
	}

	if err != nil {
		return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
	}

	return ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, 200, "", true, nil
}

func getIPApplicationsDR(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]*ent.Exam_Applications_IP, error) {
	query := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.ExamCodeEQ(examCode),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			exam_applications_ip.HallTicketNumberNEQ(""),
			exam_applications_ip.ExamCityCenterCodeNEQ(0),
		)
	applications, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	if len(applications) == 0 {
		return nil, errors.New("no applications found for the provided details")
	}

	return applications, nil
}

func GetCirclesDetails(ctx context.Context, client *ent.Client, examCode int32, examYear string, directorateID string) ([]FacilityDetails, error) {
	var circles []FacilityDetails

	circleFacilities, err := client.FacilityMasters.
		Query().
		Where(facilitymasters.FacilityIDEQ(directorateID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	for _, circle := range circleFacilities {
		recommendedCount, notRecommendedCount, provisionallyRecommendedCount, err := GetCircleApplicationCounts(ctx, client, examCode, examYear, circle.FacilityID)
		if err != nil {
			return nil, err
		}

		circles = append(circles, FacilityDetails{
			FacilityID:                    circle.FacilityID,
			FacilityName:                  circle.FacilityIDDescription,
			RecommendedCount:              recommendedCount,
			NotRecommendedCount:           notRecommendedCount,
			ProvisionallyRecommendedCount: provisionallyRecommendedCount,
		})
	}

	return circles, nil
}

func GetCircleApplicationCounts(ctx context.Context, client *ent.Client, examCode int32, examYear string, facilityID string) (int, int, int, error) {
	recommendedCount, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamCodeEQ(examCode),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_ip.RecommendedStatusEQ("Recommended"),
		).
		Count(ctx)
	if err != nil {
		return 0, 0, 0, err
	}

	notRecommendedCount, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamCodeEQ(examCode),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_ip.RecommendedStatusEQ("Not Recommended"),
		).
		Count(ctx)
	if err != nil {
		return 0, 0, 0, err
	}

	provisionallyRecommendedCount, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamCodeEQ(examCode),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_ip.RecommendedStatusEQ("Provisionally Recommended"),
		).
		Count(ctx)
	if err != nil {
		return 0, 0, 0, err
	}

	return recommendedCount, notRecommendedCount, provisionallyRecommendedCount, nil
}

// func GetCAFacilityDetails(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleID string) ([]FacilityDetails, error) {
// 	var caFacilities []FacilityDetails

// 	caFacilityList, err := client.FacilityMasters.
// 		Query().
// 		Where(facilitymasters.FacilityIDEQ(circleID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, caFacility := range caFacilityList {
// 		recommendedCount, notRecommendedCount, provisionallyRecommendedCount, err := GetCAApplicationCounts(ctx, client, examCode, examYear, caFacility.FacilityID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		caFacilities = append(caFacilities, FacilityDetails{
// 			FacilityID:                    caFacility.FacilityID,
// 			FacilityName:                  caFacility.FacilityIDDescription,
// 			RecommendedCount:              recommendedCount,
// 			NotRecommendedCount:           notRecommendedCount,
// 			ProvisionallyRecommendedCount: provisionallyRecommendedCount,
// 		})
// 	}

// 	return caFacilities, nil
// }

// func GetCAFacilityDetails(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleID string) ([]FacilityDetails, error) {
// 	var caFacilities []FacilityDetails

// 	caFacilityList, err := client.FacilityMasters.
// 		Query().
// 		Where(facilitymasters.FacilityIDEQ(circleID)).
// 		All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, caFacility := range caFacilityList {
// 		// Query for distinct controlling office IDs and names using GroupBy
// 		groupedResults, err := client.Exam_Applications_IP.
// 			Query().
// 			Where(exam_applications_ip.NodalOfficeFacilityIDEQ(caFacility.FacilityID)).
// 			GroupBy(exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
// 			All(ctx)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for _, group := range groupedResults {
// 			controllingOfficeID := group.ControllingOfficeFacilityID
// 			controllingOfficeName := group.ControllingOfficeName

// 			recommendedCount, notRecommendedCount, provisionallyRecommendedCount, err := GetCAApplicationCounts(ctx, client, examCode, examYear, controllingOfficeID)
// 			if err != nil {
// 				return nil, err
// 			}

// 			caFacilities = append(caFacilities, FacilityDetails{
// 				FacilityID:                    controllingOfficeID,
// 				FacilityName:                  controllingOfficeName,
// 				RecommendedCount:              recommendedCount,
// 				NotRecommendedCount:           notRecommendedCount,
// 				ProvisionallyRecommendedCount: provisionallyRecommendedCount,
// 			})
// 		}
// 	}

// return caFacilities, nil
type FacilityDetails struct {
	FacilityID                    string `json:"facility_id"`
	FacilityName                  string `json:"facility_name"`
	RecommendedCount              int    `json:"recommended_count"`
	NotRecommendedCount           int    `json:"not_recommended_count"`
	ProvisionallyRecommendedCount int    `json:"provisionally_recommended_count"`
}

type ControllingOfficeDetails struct {
	FacilityID                    string `json:"facility_id"`
	FacilityName                  string `json:"facility_name"`
	RecommendedCount              int    `json:"recommended_count"`
	NotRecommendedCount           int    `json:"not_recommended_count"`
	ProvisionallyRecommendedCount int    `json:"provisionally_recommended_count"`
}

type CircleOfficeDetails struct {
	FacilityID           string                     `json:"facility_id"`
	FacilityName         string                     `json:"facility_name"`
	ControllingOffices   []ControllingOfficeDetails `json:"controlling_offices"`
	RecommendedCount     int                        `json:"recommended_count"`
	NotRecommendedCount  int                        `json:"not_recommended_count"`
	ProvRecommendedCount int                        `json:"provisionally_recommended_count"`
}

func GetCAFacilityDetailsBetter(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleID string) (*CircleOfficeDetails, error) {
	circleDetails, err := client.FacilityMasters.
		Query().
		Where(facilitymasters.FacilityIDEQ(circleID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	var caFacilities []FacilityDetails

	caFacilityList, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.NodalOfficeFacilityIDEQ(circleID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var controllingOffices []ControllingOfficeDetails
	for _, caFacility := range caFacilityList {
		recommendedCount, notRecommendedCount, provisionallyRecommendedCount, err := GetCAApplicationCounts(ctx, client, examCode, examYear, caFacility.ControllingOfficeFacilityID)
		if err != nil {
			return nil, err
		}

		controllingOffices = append(controllingOffices, ControllingOfficeDetails{
			FacilityID:   caFacility.ControllingOfficeFacilityID,
			FacilityName: caFacility.ControllingOfficeName,
		})

		caFacilities = append(caFacilities, FacilityDetails{
			FacilityID:                    caFacility.ControllingOfficeFacilityID,
			FacilityName:                  caFacility.ControllingOfficeName,
			RecommendedCount:              recommendedCount,
			NotRecommendedCount:           notRecommendedCount,
			ProvisionallyRecommendedCount: provisionallyRecommendedCount,
		})
	}

	circleOfficeDetails := &CircleOfficeDetails{
		FacilityID:           circleDetails.FacilityID,
		FacilityName:         circleDetails.FacilityIDDescription,
		ControllingOffices:   controllingOffices,
		RecommendedCount:     calculateSum(caFacilities, "recommended"),
		NotRecommendedCount:  calculateSum(caFacilities, "not_recommended"),
		ProvRecommendedCount: calculateSum(caFacilities, "provisionally_recommended"),
	}

	return circleOfficeDetails, nil
}

// test
func GetCAFacilityDetails(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleID string) (*CircleOfficeDetails, error) {
	circleDetails, err := client.FacilityMasters.
		Query().
		Where(facilitymasters.FacilityIDEQ(circleID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	caFacilityList, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.NodalOfficeFacilityIDEQ(circleID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var controllingOffices []ControllingOfficeDetails
	for _, caFacility := range caFacilityList {
		recommendedCount, notRecommendedCount, provisionallyRecommendedCount, err := GetCAApplicationCounts(ctx, client, examCode, examYear, caFacility.ControllingOfficeFacilityID)
		if err != nil {
			return nil, err
		}

		controllingOffices = append(controllingOffices, ControllingOfficeDetails{
			FacilityID:                    caFacility.ControllingOfficeFacilityID,
			FacilityName:                  caFacility.ControllingOfficeName,
			RecommendedCount:              recommendedCount,
			NotRecommendedCount:           notRecommendedCount,
			ProvisionallyRecommendedCount: provisionallyRecommendedCount,
		})
	}

	// Calculate total counts for the circle
	var totalRecommendedCount, totalNotRecommendedCount, totalProvisionallyRecommendedCount int
	for _, office := range controllingOffices {
		totalRecommendedCount += office.RecommendedCount
		totalNotRecommendedCount += office.NotRecommendedCount
		totalProvisionallyRecommendedCount += office.ProvisionallyRecommendedCount
	}

	circleOfficeDetails := &CircleOfficeDetails{
		FacilityID:           circleDetails.FacilityID,
		FacilityName:         circleDetails.FacilityIDDescription,
		ControllingOffices:   controllingOffices,
		RecommendedCount:     totalRecommendedCount,
		NotRecommendedCount:  totalNotRecommendedCount,
		ProvRecommendedCount: totalProvisionallyRecommendedCount,
	}

	return circleOfficeDetails, nil
}
func calculateSum(facilities []FacilityDetails, status string) int {
	sum := 0
	for _, facility := range facilities {
		switch status {
		case "recommended":
			sum += facility.RecommendedCount
		case "not_recommended":
			sum += facility.NotRecommendedCount
		case "provisionally_recommended":
			sum += facility.ProvisionallyRecommendedCount
		}
	}
	return sum
}

// func GetCAApplicationCounts(ctx context.Context, client *ent.Client, examCode int32, examYear string, facilityID string) (int, int, int, error) {
// 	recommendedCount, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.ExamYearEQ(examYear),
// 			exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
// 			exam_applications_ip.RecommendedStatusEQ("Recommended"),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}

// 	notRecommendedCount, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.ExamYearEQ(examYear),
// 			exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
// 			exam_applications_ip.RecommendedStatusEQ("Not Recommended"),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}

// 	provisionallyRecommendedCount, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.ExamYearEQ(examYear),
// 			exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
// 			exam_applications_ip.RecommendedStatusEQ("Provisionally Recommended"),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}

//		return recommendedCount, notRecommendedCount, provisionallyRecommendedCount, nil
//	}
// func GetCAApplicationCounts(ctx context.Context, client *ent.Client, examCode int32, examYear string, facilityID string) (int, int, int, error) {
// 	recommendedCount, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.ExamYearEQ(examYear),
// 			//	exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
// 			exam_applications_ip.RecommendedStatusEQ("Recommended"),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}

// 	notRecommendedCount, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.ExamYearEQ(examYear),
// 			//	exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
// 			exam_applications_ip.RecommendedStatusEQ("Not Recommended"),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}

// 	provisionallyRecommendedCount, err := client.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.ExamCodeEQ(examCode),
// 			exam_applications_ip.ExamYearEQ(examYear),
// 			//exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
// 			exam_applications_ip.RecommendedStatusEQ("Provisionally Recommended"),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		Count(ctx)
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}

// 	return recommendedCount, notRecommendedCount, provisionallyRecommendedCount, nil
// }

// 27-06
func GetCAApplicationCounts(ctx context.Context, client *ent.Client, examCode int32, examYear string, controllingOfficeID string) (int, int, int, error) {
	// Log the parameters to ensure they are as expected
	log.Printf("Querying application counts for ExamCode: %d, ExamYear: %s, ControllingOfficeID: %s", examCode, examYear, controllingOfficeID)

	var recommendedCount, notRecommendedCount, provisionallyRecommendedCount int
	var err error

	switch examCode {
	case 2:
		recommendedCount, err = client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_ip.RecommendedStatusEQ("Recommended"),
				exam_applications_ip.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Recommended count: %d", recommendedCount)

		notRecommendedCount, err = client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_ip.RecommendedStatusEQ("Not Recommended"),
				exam_applications_ip.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying not recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Not recommended count: %d", notRecommendedCount)

		provisionallyRecommendedCount, err = client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_ip.RecommendedStatusEQ("Provisionally Recommended"),
				exam_applications_ip.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying provisionally recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Provisionally recommended count: %d", provisionallyRecommendedCount)

	case 1:
		recommendedCount, err = client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_ps.RecommendedStatusEQ("Recommended"),
				exam_applications_ps.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Recommended count: %d", recommendedCount)

		notRecommendedCount, err = client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_ps.RecommendedStatusEQ("Not Recommended"),
				exam_applications_ps.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying not recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Not recommended count: %d", notRecommendedCount)

		provisionallyRecommendedCount, err = client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_ps.RecommendedStatusEQ("Provisionally Recommended"),
				exam_applications_ps.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying provisionally recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Provisionally recommended count: %d", provisionallyRecommendedCount)

	case 4:
		recommendedCount, err = client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_gdspa.RecommendedStatusEQ("Recommended"),
				exam_applications_gdspa.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Recommended count: %d", recommendedCount)

		notRecommendedCount, err = client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_gdspa.RecommendedStatusEQ("Not Recommended"),
				exam_applications_gdspa.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying not recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Not recommended count: %d", notRecommendedCount)

		provisionallyRecommendedCount, err = client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_gdspa.RecommendedStatusEQ("Provisionally Recommended"),
				exam_applications_gdspa.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying provisionally recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Provisionally recommended count: %d", provisionallyRecommendedCount)
	case 3:
		recommendedCount, err = client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_pmpa.RecommendedStatusEQ("Recommended"),
				exam_applications_pmpa.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Recommended count: %d", recommendedCount)

		notRecommendedCount, err = client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_pmpa.RecommendedStatusEQ("Not Recommended"),
				exam_applications_pmpa.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying not recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Not recommended count: %d", notRecommendedCount)

		provisionallyRecommendedCount, err = client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_pmpa.RecommendedStatusEQ("Provisionally Recommended"),
				exam_applications_pmpa.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying provisionally recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Provisionally recommended count: %d", provisionallyRecommendedCount)
	case 5:
		recommendedCount, err = client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_application_mtspmmg.RecommendedStatusEQ("Recommended"),
				exam_application_mtspmmg.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Recommended count: %d", recommendedCount)

		notRecommendedCount, err = client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_application_mtspmmg.RecommendedStatusEQ("Not Recommended"),
				exam_application_mtspmmg.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying not recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Not recommended count: %d", notRecommendedCount)

		provisionallyRecommendedCount, err = client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_application_mtspmmg.RecommendedStatusEQ("Provisionally Recommended"),
				exam_application_mtspmmg.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying provisionally recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Provisionally recommended count: %d", provisionallyRecommendedCount)
	case 6:
		recommendedCount, err = client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_gdspm.RecommendedStatusEQ("Recommended"),
				exam_applications_gdspm.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Recommended count: %d", recommendedCount)

		notRecommendedCount, err = client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_gdspm.RecommendedStatusEQ("Not Recommended"),
				exam_applications_gdspm.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying not recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Not recommended count: %d", notRecommendedCount)

		provisionallyRecommendedCount, err = client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_gdspm.RecommendedStatusEQ("Provisionally Recommended"),
				exam_applications_gdspm.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying provisionally recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Provisionally recommended count: %d", provisionallyRecommendedCount)

	default:
		log.Printf("invalid exam code: %d", examCode)
		return 0, 0, 0, fmt.Errorf("invalid exam code: %d", examCode)
	}

	return recommendedCount, notRecommendedCount, provisionallyRecommendedCount, nil
}

type Result struct {
	ControllingOfficeID string `json:"controlling_office_id"`
	RecommendedStatus   string `json:"recommended_status"`
	Count               int    `json:"count"`
}

func GetCAApplicationCountsNew(ctx context.Context, client *ent.Client, examCode int32, examYear string, controllingOfficeID string) (int, int, int, error) {
	// Log the parameters to ensure they are as expected
	log.Printf("Querying application counts for ExamCode: %d, ExamYear: %s, ControllingOfficeID: %s", examCode, examYear, controllingOfficeID)

	var recommendedCount, notRecommendedCount, provisionallyRecommendedCount int

	if examCode == 2 {
		// Define a struct to hold the scan resultsvar results []Result

		// Querying the counts grouped by RecommendedStatus
		var results []Result
		err := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.ControllingOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_ip.StatusEQ("active"),
			).
			GroupBy(exam_applications_ip.FieldRecommendedStatus).
			Aggregate(ent.Count()).
			Scan(ctx, &results)
		if err != nil {
			log.Printf("Error querying counts grouped by recommended status: %v", err)
			return 0, 0, 0, err
		}

		// Parse the results
		for _, result := range results {
			switch result.RecommendedStatus {
			case "Recommended":
				recommendedCount = result.Count
			case "Not Recommended":
				notRecommendedCount = result.Count
			case "Provisionally Recommended":
				provisionallyRecommendedCount = result.Count
			}
		}
	} else {
		log.Printf("Invalid exam code: %d", examCode)
		return 0, 0, 0, fmt.Errorf("invalid exam code: %d", examCode)
	}

	return recommendedCount, notRecommendedCount, provisionallyRecommendedCount, nil
}

type Resultt struct {
	ControllingOfficeFacilityID string `json:"controlling_office_facility_id"`
	ControllingOfficeName       string `json:"controlling_office_name"` //controlling_office_name
	RecommendedStatus           string `json:"recommended_status"`
	Count                       int    `json:"count"`
}

func GetApplicationRecomenddedCountsSummaryForCAandDT(ctx context.Context, client *ent.Client, examCode int32, examYear string, controllingOfficeID string, entityType string) ([]ca_reg.CircleOfficeDetailss, int32, string, bool, error) {
	log.Printf("Querying application counts for ExamCode: %d, ExamYear: %s, ControllingOfficeID: %s, EntityType: %s", examCode, examYear, controllingOfficeID, entityType)

	type Result struct {
		NodalOfficeFacilityID       string `json:"nodal_office_facility_id"`
		NodalOfficeName             string `json:"nodal_office_name"`
		ControllingOfficeFacilityID string `json:"controlling_office_facility_id"`
		ControllingOfficeName       string `json:"controlling_office_name"`
		RecommendedStatus           string `json:"recommended_status"`
		Count                       int    `json:"count"`
	}

	var results []Result

	switch examCode {
	case 2:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_IP.
				Query().
				Where(
					exam_applications_ip.ExamCodeEQ(examCode),
					exam_applications_ip.ExamYearEQ(examYear),
					exam_applications_ip.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_ip.FieldNodalOfficeFacilityID,
					exam_applications_ip.FieldNodalOfficeName,
					exam_applications_ip.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR001", false, err
			}

		case "CR":
			err := client.Exam_Applications_IP.
				Query().
				Where(
					exam_applications_ip.ExamCodeEQ(examCode),
					exam_applications_ip.ExamYearEQ(examYear),
					exam_applications_ip.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_ip.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_ip.FieldControllingOfficeFacilityID,
					exam_applications_ip.FieldControllingOfficeName,
					exam_applications_ip.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR002", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR001", false, fmt.Errorf("invalid entity type: %s", entityType)
		}

	case 1:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_PS.
				Query().
				Where(
					exam_applications_ps.ExamCodeEQ(examCode),
					exam_applications_ps.ExamYearEQ(examYear),
					exam_applications_ps.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_ps.FieldNodalOfficeFacilityID,
					exam_applications_ps.FieldNodalOfficeName,
					exam_applications_ps.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR003", false, err
			}

		case "CR":
			err := client.Exam_Applications_PS.
				Query().
				Where(
					exam_applications_ps.ExamCodeEQ(examCode),
					exam_applications_ps.ExamYearEQ(examYear),
					exam_applications_ps.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_ps.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_ip.FieldControllingOfficeFacilityID,
					exam_applications_ip.FieldControllingOfficeName,
					exam_applications_ip.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR004", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR005", false, fmt.Errorf("invalid entity type: %s", entityType)
		}
	case 3:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_PMPA.
				Query().
				Where(
					exam_applications_pmpa.ExamCodeEQ(examCode),
					exam_applications_pmpa.ExamYearEQ(examYear),
					exam_applications_pmpa.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_pmpa.FieldNodalOfficeFacilityID,
					exam_applications_pmpa.FieldNodalOfficeName,
					exam_applications_pmpa.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR006", false, err
			}

		case "CR":
			err := client.Exam_Applications_PMPA.
				Query().
				Where(
					exam_applications_pmpa.ExamCodeEQ(examCode),
					exam_applications_pmpa.ExamYearEQ(examYear),
					exam_applications_pmpa.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_pmpa.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_pmpa.FieldControllingOfficeFacilityID,
					exam_applications_pmpa.FieldControllingOfficeName,
					exam_applications_pmpa.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR007", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR008", false, fmt.Errorf("invalid entity type: %s", entityType)
		}
	case 4:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_GDSPA.
				Query().
				Where(
					exam_applications_gdspa.ExamCodeEQ(examCode),
					exam_applications_gdspa.ExamYearEQ(examYear),
					exam_applications_gdspa.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_gdspa.FieldNodalOfficeFacilityID,
					exam_applications_gdspa.FieldNodalOfficeName,
					exam_applications_gdspa.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR009", false, err
			}

		case "CR":
			err := client.Exam_Applications_GDSPA.
				Query().
				Where(
					exam_applications_gdspa.ExamCodeEQ(examCode),
					exam_applications_gdspa.ExamYearEQ(examYear),
					exam_applications_gdspa.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_gdspa.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_gdspa.FieldControllingOfficeFacilityID,
					exam_applications_gdspa.FieldControllingOfficeName,
					exam_applications_gdspa.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR010", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR011", false, fmt.Errorf("invalid entity type: %s", entityType)
		}

	case 5:
		switch entityType {
		case "DT":
			err := client.Exam_Application_MTSPMMG.
				Query().
				Where(
					exam_application_mtspmmg.ExamCodeEQ(examCode),
					exam_application_mtspmmg.ExamYearEQ(examYear),
					exam_application_mtspmmg.StatusEQ("active"),
				).
				GroupBy(
					exam_application_mtspmmg.FieldNodalOfficeFacilityID,
					exam_application_mtspmmg.FieldNodalOfficeName,
					exam_application_mtspmmg.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR012", false, err
			}

		case "CR":
			err := client.Exam_Application_MTSPMMG.
				Query().
				Where(
					exam_application_mtspmmg.ExamCodeEQ(examCode),
					exam_application_mtspmmg.ExamYearEQ(examYear),
					exam_application_mtspmmg.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_application_mtspmmg.StatusEQ("active"),
				).
				GroupBy(
					exam_application_mtspmmg.FieldControllingOfficeFacilityID,
					exam_application_mtspmmg.FieldControllingOfficeName,
					exam_application_mtspmmg.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR014", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR015", false, fmt.Errorf("invalid entity type: %s", entityType)
		}

	case 6:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_GDSPM.
				Query().
				Where(
					exam_applications_gdspm.ExamCodeEQ(examCode),
					exam_applications_gdspm.ExamYearEQ(examYear),
					exam_applications_gdspm.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_gdspm.FieldNodalOfficeFacilityID,
					exam_applications_gdspm.FieldNodalOfficeName,
					exam_applications_gdspm.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR016", false, err
			}

		case "CR":
			err := client.Exam_Applications_GDSPM.
				Query().
				Where(
					exam_applications_gdspm.ExamCodeEQ(examCode),
					exam_applications_gdspm.ExamYearEQ(examYear),
					exam_applications_gdspm.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_gdspm.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_gdspm.FieldControllingOfficeFacilityID,
					exam_applications_gdspm.FieldControllingOfficeName,
					exam_applications_gdspm.FieldRecommendedStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR017", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR018", false, fmt.Errorf("invalid entity type: %s", entityType)
		}

	default:
		log.Printf("Invalid exam code: %d", examCode)
		return nil, 500, " -STR019", false, fmt.Errorf("invalid exam code: %d", examCode)
	}

	// Pivoting the data
	var controllingOffices []ca_reg.CircleOfficeDetailss
	pivotData := make(map[string]map[string]map[string]int)

	if entityType == "CR" {
		for _, result := range results {
			if _, ok := pivotData[result.ControllingOfficeFacilityID]; !ok {
				pivotData[result.ControllingOfficeFacilityID] = make(map[string]map[string]int)
			}
			if _, ok := pivotData[result.ControllingOfficeFacilityID][result.ControllingOfficeName]; !ok {
				pivotData[result.ControllingOfficeFacilityID][result.ControllingOfficeName] = make(map[string]int)
			}
			pivotData[result.ControllingOfficeFacilityID][result.ControllingOfficeName][result.RecommendedStatus] = result.Count
		}

		for circleID, officeMap := range pivotData {
			for officeName, recommendedStatus := range officeMap {
				controllingOffices = append(controllingOffices, ca_reg.CircleOfficeDetailss{
					NodalOfficeFacilityID:         "", // Not used
					NodalOfficeName:               "", // Not used
					FacilityID:                    circleID,
					FacilityName:                  officeName,
					RecommendedCount:              recommendedStatus["Recommended"],
					NotRecommendedCount:           recommendedStatus["Not Recommended"],
					ProvisionallyRecommendedCount: recommendedStatus["Provisionally Recommended"],
				})
				log.Printf("circleID: %s, facilityName: %s, recommendedCount: %d, provisionallyRecommendedCount: %d, notRecommendedCount: %d",
					circleID, officeName, recommendedStatus["Recommended"], recommendedStatus["Provisionally Recommended"], recommendedStatus["Not Recommended"])
			}
		}

	} else if entityType == "DT" {
		for _, result := range results {
			if _, ok := pivotData[result.NodalOfficeFacilityID]; !ok {
				pivotData[result.NodalOfficeFacilityID] = make(map[string]map[string]int)
			}
			if _, ok := pivotData[result.NodalOfficeFacilityID][result.NodalOfficeName]; !ok {
				pivotData[result.NodalOfficeFacilityID][result.NodalOfficeName] = make(map[string]int)
			}
			pivotData[result.NodalOfficeFacilityID][result.NodalOfficeName][result.RecommendedStatus] = result.Count
		}

		for circleID, officeMap := range pivotData {
			for officeName, recommendedStatus := range officeMap {
				controllingOffices = append(controllingOffices, ca_reg.CircleOfficeDetailss{
					NodalOfficeFacilityID:         circleID,
					NodalOfficeName:               officeName,
					FacilityID:                    "",
					FacilityName:                  "",
					RecommendedCount:              recommendedStatus["Recommended"],
					NotRecommendedCount:           recommendedStatus["Not Recommended"],
					ProvisionallyRecommendedCount: recommendedStatus["Provisionally Recommended"],
				})
				log.Printf("nodalOfficeID: %s, nodalOfficeName: %s, recommendedCount: %d, provisionallyRecommendedCount: %d, notRecommendedCount: %d",
					circleID, officeName, recommendedStatus["Recommended"], recommendedStatus["Provisionally Recommended"], recommendedStatus["Not Recommended"])
			}
		}
	}

	return controllingOffices, 200, "", true, nil
}

type Resulttt struct {
	NodalOfficeFacilityID       string `json:"nodal_office_facility_id"`
	NodalOfficeName             string `json:"nodal_office_name"`
	ControllingOfficeFacilityID string `json:"controlling_office_facility_id"`
	ControllingOfficeName       string `json:"controlling_office_name"`
	ApplicationStatus           string `json:"application_status"`
	Count                       int    `json:"count"`
}

func GetApplicationApplicationStatusCountsSummaryForCAandDT(ctx context.Context, client *ent.Client, examCode int32, examYear string, controllingOfficeID string, entityType string) ([]ca_reg.ApplicationStatusDetailss, int32, string, bool, error) {
	log.Printf("Querying application counts for ExamCode: %d, ExamYear: %s, ControllingOfficeID: %s, EntityType: %s", examCode, examYear, controllingOfficeID, entityType)

	var results []Resulttt

	switch examCode {
	case 2:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_IP.
				Query().
				Where(
					exam_applications_ip.ExamCodeEQ(examCode),
					exam_applications_ip.ExamYearEQ(examYear),
					exam_applications_ip.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_ip.FieldNodalOfficeFacilityID,
					exam_applications_ip.FieldNodalOfficeName,
					exam_applications_ip.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR001", false, err
			}

		case "CR":
			err := client.Exam_Applications_IP.
				Query().
				Where(
					exam_applications_ip.ExamCodeEQ(examCode),
					exam_applications_ip.ExamYearEQ(examYear),
					exam_applications_ip.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_ip.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_ip.FieldControllingOfficeFacilityID,
					exam_applications_ip.FieldControllingOfficeName,
					exam_applications_ip.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR002", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR003", false, fmt.Errorf("invalid entity type: %s", entityType)
		}

	case 1:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_PS.
				Query().
				Where(
					exam_applications_ps.ExamCodeEQ(examCode),
					exam_applications_ps.ExamYearEQ(examYear),
					exam_applications_ps.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_ps.FieldNodalOfficeFacilityID,
					exam_applications_ps.FieldNodalOfficeName,
					exam_applications_ps.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR004", false, err
			}

		case "CR":
			err := client.Exam_Applications_PS.
				Query().
				Where(
					exam_applications_ps.ExamCodeEQ(examCode),
					exam_applications_ps.ExamYearEQ(examYear),
					exam_applications_ps.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_ps.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_ip.FieldControllingOfficeFacilityID,
					exam_applications_ip.FieldControllingOfficeName,
					exam_applications_ip.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR005", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR006", false, fmt.Errorf("invalid entity type: %s", entityType)
		}
	case 3:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_PMPA.
				Query().
				Where(
					exam_applications_pmpa.ExamCodeEQ(examCode),
					exam_applications_pmpa.ExamYearEQ(examYear),
					exam_applications_pmpa.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_pmpa.FieldNodalOfficeFacilityID,
					exam_applications_pmpa.FieldNodalOfficeName,
					exam_applications_pmpa.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR007", false, err
			}

		case "CR":
			err := client.Exam_Applications_PMPA.
				Query().
				Where(
					exam_applications_pmpa.ExamCodeEQ(examCode),
					exam_applications_pmpa.ExamYearEQ(examYear),
					exam_applications_pmpa.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_pmpa.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_pmpa.FieldControllingOfficeFacilityID,
					exam_applications_pmpa.FieldControllingOfficeName,
					exam_applications_pmpa.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR008", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR009", false, fmt.Errorf("invalid entity type: %s", entityType)
		}
	case 4:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_GDSPA.
				Query().
				Where(
					exam_applications_gdspa.ExamCodeEQ(examCode),
					exam_applications_gdspa.ExamYearEQ(examYear),
					exam_applications_gdspa.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_gdspa.FieldNodalOfficeFacilityID,
					exam_applications_gdspa.FieldNodalOfficeName,
					exam_applications_gdspa.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR010", false, err
			}

		case "CR":
			err := client.Exam_Applications_GDSPA.
				Query().
				Where(
					exam_applications_gdspa.ExamCodeEQ(examCode),
					exam_applications_gdspa.ExamYearEQ(examYear),
					exam_applications_gdspa.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_gdspa.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_gdspa.FieldControllingOfficeFacilityID,
					exam_applications_gdspa.FieldControllingOfficeName,
					exam_applications_gdspa.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR011", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR012", false, fmt.Errorf("invalid entity type: %s", entityType)
		}

	case 5:
		switch entityType {
		case "DT":
			err := client.Exam_Application_MTSPMMG.
				Query().
				Where(
					exam_application_mtspmmg.ExamCodeEQ(examCode),
					exam_application_mtspmmg.ExamYearEQ(examYear),
					exam_application_mtspmmg.StatusEQ("active"),
				).
				GroupBy(
					exam_application_mtspmmg.FieldNodalOfficeFacilityID,
					exam_application_mtspmmg.FieldNodalOfficeName,
					exam_application_mtspmmg.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR013", false, err
			}

		case "CR":
			err := client.Exam_Application_MTSPMMG.
				Query().
				Where(
					exam_application_mtspmmg.ExamCodeEQ(examCode),
					exam_application_mtspmmg.ExamYearEQ(examYear),
					exam_application_mtspmmg.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_application_mtspmmg.StatusEQ("active"),
				).
				GroupBy(
					exam_application_mtspmmg.FieldControllingOfficeFacilityID,
					exam_application_mtspmmg.FieldControllingOfficeName,
					exam_application_mtspmmg.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR014", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR015", false, fmt.Errorf("invalid entity type: %s", entityType)
		}

	case 6:
		switch entityType {
		case "DT":
			err := client.Exam_Applications_GDSPM.
				Query().
				Where(
					exam_applications_gdspm.ExamCodeEQ(examCode),
					exam_applications_gdspm.ExamYearEQ(examYear),
					exam_applications_gdspm.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_gdspm.FieldNodalOfficeFacilityID,
					exam_applications_gdspm.FieldNodalOfficeName,
					exam_applications_gdspm.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR016", false, err
			}

		case "CR":
			err := client.Exam_Applications_GDSPM.
				Query().
				Where(
					exam_applications_gdspm.ExamCodeEQ(examCode),
					exam_applications_gdspm.ExamYearEQ(examYear),
					exam_applications_gdspm.NodalOfficeFacilityIDEQ(controllingOfficeID),
					exam_applications_gdspm.StatusEQ("active"),
				).
				GroupBy(
					exam_applications_gdspm.FieldControllingOfficeFacilityID,
					exam_applications_gdspm.FieldControllingOfficeName,
					exam_applications_gdspm.FieldApplicationStatus,
				).
				Aggregate(ent.Count()).
				Scan(ctx, &results)
			if err != nil {
				log.Fatalf("failed querying sub-query: %v", err)
				return nil, 500, " -STR017", false, err
			}

		default:
			log.Printf("Invalid entity type: %s", entityType)
			return nil, 422, " -STR018", false, fmt.Errorf("invalid entity type: %s", entityType)
		}

	default:
		log.Printf("Invalid exam code: %d", examCode)
		return nil, 500, " -STR019", false, fmt.Errorf("invalid exam code: %d", examCode)
	}

	// Pivoting the data
	var controllingOffices []ca_reg.ApplicationStatusDetailss
	pivotData := make(map[string]map[string]map[string]int)

	if entityType == "CR" {
		for _, result := range results {
			if _, ok := pivotData[result.ControllingOfficeFacilityID]; !ok {
				pivotData[result.ControllingOfficeFacilityID] = make(map[string]map[string]int)
			}
			if _, ok := pivotData[result.ControllingOfficeFacilityID][result.ControllingOfficeName]; !ok {
				pivotData[result.ControllingOfficeFacilityID][result.ControllingOfficeName] = make(map[string]int)
			}
			pivotData[result.ControllingOfficeFacilityID][result.ControllingOfficeName][result.ApplicationStatus] = result.Count
		}

		for circleID, officeMap := range pivotData {
			for officeName, applicationstatus := range officeMap {

				applicationCount := applicationstatus["CAVerificationPending"] + applicationstatus["ResubmitCAVerificationPending"]
				controllingOffices = append(controllingOffices, ca_reg.ApplicationStatusDetailss{
					NodalOfficeFacilityID: "", // Not used
					NodalOfficeName:       "", // Not used
					FacilityID:            circleID,
					FacilityName:          officeName,
					PendingWithCandidate:  applicationstatus["PendingWithCandidate"],
					PendingWithCA:         applicationCount,
				})

			}
		}

	} else if entityType == "DT" {
		for _, result := range results {
			if _, ok := pivotData[result.NodalOfficeFacilityID]; !ok {
				pivotData[result.NodalOfficeFacilityID] = make(map[string]map[string]int)
			}
			if _, ok := pivotData[result.NodalOfficeFacilityID][result.NodalOfficeName]; !ok {
				pivotData[result.NodalOfficeFacilityID][result.NodalOfficeName] = make(map[string]int)
			}
			pivotData[result.NodalOfficeFacilityID][result.NodalOfficeName][result.ApplicationStatus] = result.Count
		}

		for circleID, officeMap := range pivotData {
			for officeName, applicationstatus := range officeMap {
				applicationCount := applicationstatus["CAVerificationPending"] + applicationstatus["ResubmitCAVerificationPending"]
				controllingOffices = append(controllingOffices, ca_reg.ApplicationStatusDetailss{
					NodalOfficeFacilityID: circleID,
					NodalOfficeName:       officeName,
					FacilityID:            "",
					FacilityName:          "",
					PendingWithCandidate:  applicationstatus["PendingWithCandidate"],
					PendingWithCA:         applicationCount,
				})

			}
		}
	}

	return controllingOffices, 200, "", true, nil
}
func GetCAApplicationCountsNewD(ctx context.Context, client *ent.Client, examCode int32, examYear string, controllingOfficeID string) (int, int, int, error) {
	// Log the parameters to ensure they are as expected
	log.Printf("Querying application counts for ExamCode: %d, ExamYear: %s, ControllingOfficeID: %s", examCode, examYear, controllingOfficeID)

	var recommendedCount, notRecommendedCount, provisionallyRecommendedCount int

	if examCode == 2 {
		// Define a struct to hold the scan resultsvar results []Result

		// Querying the counts grouped by RecommendedStatus
		var results []Result
		err := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(controllingOfficeID),
				exam_applications_ip.StatusEQ("active"),
			).
			GroupBy(exam_applications_ip.FieldRecommendedStatus).
			Aggregate(ent.Count()).
			Scan(ctx, &results)
		if err != nil {
			log.Printf("Error querying counts grouped by recommended status: %v", err)
			return 0, 0, 0, err
		}

		// Parse the results
		for _, result := range results {
			switch result.RecommendedStatus {
			case "Recommended":
				recommendedCount = result.Count
			case "Not Recommended":
				notRecommendedCount = result.Count
			case "Provisionally Recommended":
				provisionallyRecommendedCount = result.Count
			}
		}
	} else {
		log.Printf("Invalid exam code: %d", examCode)
		return 0, 0, 0, fmt.Errorf("invalid exam code: %d", examCode)
	}

	return recommendedCount, notRecommendedCount, provisionallyRecommendedCount, nil
}

func GetPendingApplicationsWithDays(ctx context.Context, client *ent.Client, examCode int32, examYear string, nodalOfficeFacilityID string, daysPending int) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, error) {
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, errors.New("please provide a valid exam code")
	}

	now := time.Now().UTC()
	currentDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	var applicationsIP []*ent.Exam_Applications_IP
	var applicationsPS []*ent.Exam_Applications_PS
	var applicationsGdsPA []*ent.Exam_Applications_GDSPA
	var applicationsPmPa []*ent.Exam_Applications_PMPA
	var applicationsGdsPm []*ent.Exam_Applications_GDSPM
	var applicationsMtsPm []*ent.Exam_Application_MTSPMMG

	var err error

	switch examCode {
	case 1:
		applicationsPS, err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			All(ctx)
	case 2:
		applicationsIP, err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			All(ctx)
	case 3:
		applicationsPmPa, err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			All(ctx)
	case 4:
		applicationsGdsPA, err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			All(ctx)
	case 5:
		applicationsMtsPm, err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			All(ctx)
	case 6:
		applicationsGdsPm, err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			All(ctx)
	default:
		return nil, nil, nil, nil, nil, nil, errors.New("exam code invalid")
	}

	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	filterApplicationsByDays := func(applications interface{}, currentDate time.Time, daysPending int) interface{} {
		switch v := applications.(type) {
		case []*ent.Exam_Applications_IP:
			var filtered []*ent.Exam_Applications_IP
			for _, application := range v {
				submittedDate := time.Date(application.ApplnSubmittedDate.Year(), application.ApplnSubmittedDate.Month(), application.ApplnSubmittedDate.Day(), 0, 0, 0, 0, time.UTC)
				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Applications_PS:
			var filtered []*ent.Exam_Applications_PS
			for _, application := range v {
				submittedDate := time.Date(application.ApplnSubmittedDate.Year(), application.ApplnSubmittedDate.Month(), application.ApplnSubmittedDate.Day(), 0, 0, 0, 0, time.UTC)
				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Applications_GDSPA:
			var filtered []*ent.Exam_Applications_GDSPA
			for _, application := range v {
				submittedDate := time.Date(application.ApplnSubmittedDate.Year(), application.ApplnSubmittedDate.Month(), application.ApplnSubmittedDate.Day(), 0, 0, 0, 0, time.UTC)
				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Applications_PMPA:
			var filtered []*ent.Exam_Applications_PMPA
			for _, application := range v {
				submittedDate := time.Date(application.ApplnSubmittedDate.Year(), application.ApplnSubmittedDate.Month(), application.ApplnSubmittedDate.Day(), 0, 0, 0, 0, time.UTC)
				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Applications_GDSPM:
			var filtered []*ent.Exam_Applications_GDSPM
			for _, application := range v {
				submittedDate := time.Date(application.ApplnSubmittedDate.Year(), application.ApplnSubmittedDate.Month(), application.ApplnSubmittedDate.Day(), 0, 0, 0, 0, time.UTC)
				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Application_MTSPMMG:
			var filtered []*ent.Exam_Application_MTSPMMG
			for _, application := range v {
				submittedDate := time.Date(application.ApplnSubmittedDate.Year(), application.ApplnSubmittedDate.Month(), application.ApplnSubmittedDate.Day(), 0, 0, 0, 0, time.UTC)
				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		}
		return nil
	}

	filteredApplicationsIP := filterApplicationsByDays(applicationsIP, currentDate, daysPending).([]*ent.Exam_Applications_IP)
	filteredApplicationsPS := filterApplicationsByDays(applicationsPS, currentDate, daysPending).([]*ent.Exam_Applications_PS)
	filteredApplicationsGdsPA := filterApplicationsByDays(applicationsGdsPA, currentDate, daysPending).([]*ent.Exam_Applications_GDSPA)
	filteredApplicationsPmPa := filterApplicationsByDays(applicationsPmPa, currentDate, daysPending).([]*ent.Exam_Applications_PMPA)
	filteredApplicationsGdsPm := filterApplicationsByDays(applicationsGdsPm, currentDate, daysPending).([]*ent.Exam_Applications_GDSPM)
	filteredApplicationsMtsPm := filterApplicationsByDays(applicationsMtsPm, currentDate, daysPending).([]*ent.Exam_Application_MTSPMMG)

	if len(filteredApplicationsIP) == 0 && len(filteredApplicationsPS) == 0 && len(filteredApplicationsGdsPA) == 0 && len(filteredApplicationsPmPa) == 0 && len(filteredApplicationsGdsPm) == 0 && len(filteredApplicationsMtsPm) == 0 {
		return nil, nil, nil, nil, nil, nil, errors.New("no applications found for the provided details")
	}

	return filteredApplicationsIP, filteredApplicationsPS, filteredApplicationsGdsPA, filteredApplicationsPmPa, filteredApplicationsGdsPm, filteredApplicationsMtsPm, nil
}

type FacilitySummary struct {
	ControllingFacilityID   string `json:"controlling_authority_facility_id"`
	ControllingFacilityName string `json:"controlling_authority_name"`
	OneDayCount             int    `json:"one_day_count"`
	TwoDayCount             int    `json:"two_day_count"`
	ThreeDayCount           int    `json:"three_day_count"`
	FourDayCount            int    `json:"four_day_count"`
	FiveDayCount            int    `json:"five_day_count"`
	SixDayCount             int    `json:"six_day_count"`
	SevenDayCount           int    `json:"seven_day_count"`
	MoreThanSevenDayCount   int    `json:"more_than_seven_day_count"`
}

type FacilitySubSummary struct {
	ControllingFacilityID   string `json:"controlling_authority_facility_id"`
	ControllingFacilityName string `json:"controlling_authority_name"`
	Count                   int    `json:"count"`
}

func SubGetEmployeeMasterPendingWithCA(ctx context.Context, client *ent.Client, FacilityID string) ([]FacilitySummary, int32, string, bool, error) {

	now := time.Now()

	var oneDayResults []FacilitySubSummary
	var err error

	oneDayAgo := now.Add(-24 * time.Hour)
	twoDaysAgo := now.Add(-2 * 24 * time.Hour)
	threeDaysAgo := now.Add(-3 * 24 * time.Hour)
	fourDaysAgo := now.Add(-4 * 24 * time.Hour)
	fiveDaysAgo := now.Add(-5 * 24 * time.Hour)
	sixDaysAgo := now.Add(-6 * 24 * time.Hour)
	sevenDaysAgo := now.Add(-7 * 24 * time.Hour)
	fmt.Println("oneDayAgo", oneDayAgo)
	fmt.Println("twoDaysAgo", twoDaysAgo)
	fmt.Println("threeDaysAgo", threeDaysAgo)
	fmt.Println("sevenDaysAgo", sevenDaysAgo)
	//tenDaysAgo := now.Add(-10 * 24 * time.Hour)
	//oneMonthAgo := now.AddDate(0, -1, 0)

	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.NodalAuthorityFaciliyIdEQ(FacilityID),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateGTE(oneDayAgo),
		).
		GroupBy(employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &oneDayResults)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	var twoDayResults []FacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.NodalAuthorityFaciliyIdEQ(FacilityID),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(oneDayAgo),
			employeemaster.CreatedDateGTE(twoDaysAgo),
		).
		GroupBy(employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &twoDayResults)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	var threeDayResults []FacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.NodalAuthorityFaciliyIdEQ(FacilityID),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(twoDaysAgo),
			employeemaster.CreatedDateGTE(threeDaysAgo),
		).
		GroupBy(employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &threeDayResults)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	var fourDayResults []FacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.NodalAuthorityFaciliyIdEQ(FacilityID),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(threeDaysAgo),
			employeemaster.CreatedDateGTE(fourDaysAgo),
		).
		GroupBy(employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &fourDayResults)
	if err != nil {
		return nil, 500, " -STR004", false, err
	}

	var fiveDayResults []FacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.NodalAuthorityFaciliyIdEQ(FacilityID),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(fourDaysAgo),
			employeemaster.CreatedDateGTE(fiveDaysAgo),
		).
		GroupBy(employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &fiveDayResults)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	var sixDayResults []FacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.NodalAuthorityFaciliyIdEQ(FacilityID),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(fiveDaysAgo),
			employeemaster.CreatedDateGTE(sixDaysAgo),
		).
		GroupBy(employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &sixDayResults)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	var sevenDaysResults []FacilitySubSummary

	// Fetch data for one week
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.NodalAuthorityFaciliyIdEQ(FacilityID),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(sixDaysAgo),
			employeemaster.CreatedDateGTE(sevenDaysAgo),
		).
		GroupBy(employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &sevenDaysResults)
	if err != nil {
		return nil, 500, " -STR007", false, err
	}

	var moreThanSevenDaysResult []FacilitySubSummary
	// Fetch data for one month
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.NodalAuthorityFaciliyIdEQ(FacilityID),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(sevenDaysAgo),
		).
		GroupBy(employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &moreThanSevenDaysResult)
	if err != nil {
		return nil, 500, " -STR008", false, err
	}

	// Combine results
	summaryMap := make(map[string]*FacilitySummary)

	for _, r := range oneDayResults {
		controllingfacilityid := r.ControllingFacilityID

		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &FacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].OneDayCount = r.Count
	}

	for _, r := range twoDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &FacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].TwoDayCount = r.Count
	}

	for _, r := range threeDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &FacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].ThreeDayCount = r.Count
	}

	for _, r := range fourDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &FacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].FourDayCount = r.Count
	}

	for _, r := range fiveDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &FacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].FiveDayCount = r.Count
	}

	for _, r := range sixDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &FacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].SixDayCount = r.Count
	}

	for _, r := range sevenDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &FacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].SevenDayCount = r.Count
	}

	for _, r := range moreThanSevenDaysResult {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &FacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].MoreThanSevenDayCount = r.Count
	}
	// Convert map to slice
	summaries := make([]FacilitySummary, 0, len(summaryMap))
	for _, summary := range summaryMap {
		summaries = append(summaries, *summary)
	}

	return summaries, 200, "", true, nil

}

type NAFacilitySummary struct {
	NodalFacilityID         string `json:"nodal_authority_faciliy_id"`
	NodalFacilityName       string `json:"nodal_authority_name"`
	ControllingFacilityID   string `json:"controlling_authority_facility_id"`
	ControllingFacilityName string `json:"controlling_authority_name"`
	OneDayCount             int    `json:"one_day_count"`
	TwoDayCount             int    `json:"two_day_count"`
	ThreeDayCount           int    `json:"three_day_count"`
	FourDayCount            int    `json:"four_day_count"`
	FiveDayCount            int    `json:"five_day_count"`
	SixDayCount             int    `json:"six_day_count"`
	SevenDayCount           int    `json:"seven_day_count"`
	MoreThanSevenDayCount   int    `json:"more_than_seven_day_count"`
}

type NAFacilitySubSummary struct {
	NodalFacilityID         string `json:"nodal_authority_faciliy_id"`
	NodalFacilityName       string `json:"nodal_authority_name"`
	ControllingFacilityID   string `json:"controlling_authority_facility_id"`
	ControllingFacilityName string `json:"controlling_authority_name"`
	Count                   int    `json:"count"`
}

func SubGetEmployeeMasterPendingWithCADT(ctx context.Context, client *ent.Client) ([]NAFacilitySummary, int32, string, bool, error) {
	now := time.Now()
	var oneDayResults []NAFacilitySubSummary
	var err error

	oneDayAgo := now.Add(-24 * time.Hour)
	twoDaysAgo := now.Add(-2 * 24 * time.Hour)
	threeDaysAgo := now.Add(-3 * 24 * time.Hour)
	fourDaysAgo := now.Add(-4 * 24 * time.Hour)
	fiveDaysAgo := now.Add(-5 * 24 * time.Hour)
	sixDaysAgo := now.Add(-6 * 24 * time.Hour)
	sevenDaysAgo := now.Add(-7 * 24 * time.Hour)
	fmt.Println("oneDayAgo", oneDayAgo)
	fmt.Println("twoDaysAgo", twoDaysAgo)
	fmt.Println("threeDaysAgo", threeDaysAgo)
	fmt.Println("sevenDaysAgo", sevenDaysAgo)
	//tenDaysAgo := now.Add(-10 * 24 * time.Hour)
	//oneMonthAgo := now.AddDate(0, -1, 0)

	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateGTE(oneDayAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName,
			employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &oneDayResults)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	var twoDayResults []NAFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(oneDayAgo),
			employeemaster.CreatedDateGTE(twoDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName,
			employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &twoDayResults)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	var threeDayResults []NAFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(twoDaysAgo),
			employeemaster.CreatedDateGTE(threeDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName,
			employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &threeDayResults)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	var fourDayResults []NAFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(threeDaysAgo),
			employeemaster.CreatedDateGTE(fourDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName,
			employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &fourDayResults)
	if err != nil {
		return nil, 500, " -STR004", false, err
	}

	var fiveDayResults []NAFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(fourDaysAgo),
			employeemaster.CreatedDateGTE(fiveDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName,
			employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &fiveDayResults)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	var sixDayResults []NAFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(fiveDaysAgo),
			employeemaster.CreatedDateGTE(sixDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName,
			employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &sixDayResults)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	var sevenDaysResults []NAFacilitySubSummary

	// Fetch data for one week
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(sixDaysAgo),
			employeemaster.CreatedDateGTE(sevenDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName,
			employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &sevenDaysResults)
	if err != nil {
		return nil, 500, " -STR007", false, err
	}

	var moreThanSevenDaysResult []NAFacilitySubSummary
	// Fetch data for one month
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(sevenDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName,
			employeemaster.FieldControllingAuthorityFacilityId, employeemaster.FieldControllingAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &moreThanSevenDaysResult)
	if err != nil {
		return nil, 500, " -STR008", false, err
	}

	// Combine results
	summaryMap := make(map[string]*NAFacilitySummary)

	for _, r := range oneDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &NAFacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].OneDayCount = r.Count
	}

	for _, r := range twoDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &NAFacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].TwoDayCount = r.Count
	}

	for _, r := range threeDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &NAFacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].ThreeDayCount = r.Count
	}

	for _, r := range fourDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &NAFacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].FourDayCount = r.Count
	}

	for _, r := range fiveDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &NAFacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].FiveDayCount = r.Count
	}

	for _, r := range sixDayResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &NAFacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].SixDayCount = r.Count
	}

	for _, r := range sevenDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &NAFacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].SevenDayCount = r.Count
	}

	for _, r := range moreThanSevenDaysResult {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &NAFacilitySummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].MoreThanSevenDayCount = r.Count
	}
	// Convert map to slice
	summaries := make([]NAFacilitySummary, 0, len(summaryMap))
	for _, summary := range summaryMap {
		summaries = append(summaries, *summary)
	}

	return summaries, 200, "", true, nil

}

type DTFacilitySummary struct {
	NodalFacilityID       string `json:"nodal_authority_faciliy_id"`
	NodalFacilityName     string `json:"nodal_authority_name"`
	OneDayCount           int    `json:"one_day_count"`
	TwoDayCount           int    `json:"two_day_count"`
	ThreeDayCount         int    `json:"three_day_count"`
	FourDayCount          int    `json:"four_day_count"`
	FiveDayCount          int    `json:"five_day_count"`
	SixDayCount           int    `json:"six_day_count"`
	SevenDayCount         int    `json:"seven_day_count"`
	MoreThanSevenDayCount int    `json:"more_than_seven_day_count"`
}

type DTFacilitySubSummary struct {
	NodalFacilityID   string `json:"nodal_authority_faciliy_id"`
	NodalFacilityName string `json:"nodal_authority_name"`
	Count             int    `json:"count"`
}

func SubGetEmployeeMasterPendingWithNADT(ctx context.Context, client *ent.Client) ([]DTFacilitySummary, int32, string, bool, error) {
	now := time.Now()
	var oneDayResults []DTFacilitySubSummary
	var err error

	oneDayAgo := now.Add(-24 * time.Hour)
	twoDaysAgo := now.Add(-2 * 24 * time.Hour)
	threeDaysAgo := now.Add(-3 * 24 * time.Hour)
	fourDaysAgo := now.Add(-4 * 24 * time.Hour)
	fiveDaysAgo := now.Add(-5 * 24 * time.Hour)
	sixDaysAgo := now.Add(-6 * 24 * time.Hour)
	sevenDaysAgo := now.Add(-7 * 24 * time.Hour)

	//tenDaysAgo := now.Add(-10 * 24 * time.Hour)
	//oneMonthAgo := now.AddDate(0, -1, 0)

	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateGTE(oneDayAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &oneDayResults)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	var twoDayResults []DTFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(oneDayAgo),
			employeemaster.CreatedDateGTE(twoDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &twoDayResults)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	var threeDayResults []DTFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(twoDaysAgo),
			employeemaster.CreatedDateGTE(threeDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &threeDayResults)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	var fourDayResults []DTFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(threeDaysAgo),
			employeemaster.CreatedDateGTE(fourDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &fourDayResults)
	if err != nil {
		return nil, 500, " -STR004", false, err
	}

	var fiveDayResults []DTFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(fourDaysAgo),
			employeemaster.CreatedDateGTE(fiveDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &fiveDayResults)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	var sixDayResults []NAFacilitySubSummary
	// Fetch data for two days
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(fiveDaysAgo),
			employeemaster.CreatedDateGTE(sixDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &sixDayResults)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	var sevenDaysResults []NAFacilitySubSummary

	// Fetch data for one week
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(sixDaysAgo),
			employeemaster.CreatedDateGTE(sevenDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &sevenDaysResults)
	if err != nil {
		return nil, 500, " -STR007", false, err
	}

	var moreThanSevenDaysResult []DTFacilitySubSummary
	// Fetch data for one month
	err = client.EmployeeMaster.Query().
		Where(
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
			employeemaster.CreatedDateLT(sevenDaysAgo),
		).
		GroupBy(
			employeemaster.FieldNodalAuthorityFaciliyId, employeemaster.FieldNodalAuthorityName).
		Aggregate(ent.Count()).
		Scan(ctx, &moreThanSevenDaysResult)
	if err != nil {
		return nil, 500, " -STR008", false, err
	}

	// Combine results
	summaryMap := make(map[string]*DTFacilitySummary)

	for _, r := range oneDayResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTFacilitySummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].OneDayCount = r.Count
	}

	for _, r := range twoDayResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTFacilitySummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].TwoDayCount = r.Count
	}

	for _, r := range threeDayResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTFacilitySummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].ThreeDayCount = r.Count
	}

	for _, r := range fourDayResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTFacilitySummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].FourDayCount = r.Count
	}

	for _, r := range fiveDayResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTFacilitySummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].FiveDayCount = r.Count
	}

	for _, r := range sixDayResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTFacilitySummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].SixDayCount = r.Count
	}

	for _, r := range sevenDaysResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTFacilitySummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].SevenDayCount = r.Count
	}

	for _, r := range moreThanSevenDaysResult {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTFacilitySummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].MoreThanSevenDayCount = r.Count
	}
	// Convert map to slice
	summaries := make([]DTFacilitySummary, 0, len(summaryMap))
	for _, summary := range summaryMap {
		summaries = append(summaries, *summary)
	}

	return summaries, 200, "", true, nil

}

type Summary struct {
	UserID        string `json:"userid"`
	OneDayCount   int    `json:"one_day_count"`
	TwoDayCount   int    `json:"two_day_count"`
	OneWeekCount  int    `json:"one_week_count"`
	OneMonthCount int    `json:"one_month_count"`
}

type SubSummary struct {
	UserID string `json:"userid"`
	Count  int    `json:"count"`
}

func GetSummary(ctx context.Context, client *ent.Client) ([]Summary, error) {
	now := time.Now()

	// 	FieldEQ for equality (equal to).
	// FieldNEQ for inequality (not equal to).
	// FieldGT for greater than.
	// FieldGTE for greater than or equal to.
	// FieldLT for less than.
	// FieldLTE for less than or equal to.
	var oneDayResults []SubSummary
	var err error

	oneDayAgo := now.Add(-24 * time.Hour)
	twoDaysAgo := now.Add(-48 * time.Hour)
	oneWeekAgo := now.Add(-7 * 24 * time.Hour)
	oneMonthAgo := now.AddDate(0, -1, 0)

	// Fetch data for one day
	err = client.ErrorLogs.
		Query().
		Where(errorlogs.EventtimeGT(oneDayAgo)).
		GroupBy(errorlogs.FieldUserid).
		Aggregate(ent.Count()).
		Scan(ctx, &oneDayResults)
	if err != nil {
		return nil, err
	}

	var twoDayResults []SubSummary
	// Fetch data for two days
	err = client.ErrorLogs.
		Query().
		Where(errorlogs.EventtimeGT(twoDaysAgo)).
		GroupBy(errorlogs.FieldUserid).
		Aggregate(ent.Count()).
		Scan(ctx, &twoDayResults)
	if err != nil {
		return nil, err
	}

	var oneWeekResults []SubSummary
	// Fetch data for one week
	err = client.ErrorLogs.
		Query().
		Where(errorlogs.EventtimeGT(oneWeekAgo)).
		GroupBy(errorlogs.FieldUserid).
		Aggregate(ent.Count()).
		Scan(ctx, &oneWeekResults)
	if err != nil {
		return nil, err
	}

	var oneMonthResults []SubSummary
	// Fetch data for one month
	err = client.ErrorLogs.
		Query().
		Where(errorlogs.EventtimeGT(oneMonthAgo)).
		GroupBy(errorlogs.FieldUserid). /*  */
		Aggregate(ent.Count()).
		Scan(ctx, &oneMonthResults)
	if err != nil {
		return nil, err
	}

	// Combine results
	summaryMap := make(map[string]*Summary)

	for _, r := range oneDayResults {
		userid := r.UserID
		if _, exists := summaryMap[userid]; !exists {
			summaryMap[userid] = &Summary{UserID: userid}
		}
		summaryMap[userid].OneDayCount = r.Count
	}

	for _, r := range twoDayResults {
		userid := r.UserID
		//r.ID.(string)
		if _, exists := summaryMap[userid]; !exists {
			summaryMap[userid] = &Summary{UserID: userid}
		}
		summaryMap[userid].TwoDayCount = r.Count
	}

	for _, r := range oneWeekResults {
		userid := r.UserID
		if _, exists := summaryMap[userid]; !exists {
			summaryMap[userid] = &Summary{UserID: userid}
		}
		summaryMap[userid].OneWeekCount = r.Count
	}

	for _, r := range oneMonthResults {
		userid := r.UserID
		if _, exists := summaryMap[userid]; !exists {
			summaryMap[userid] = &Summary{UserID: userid}
		}
		summaryMap[userid].OneMonthCount = r.Count
	}

	// Convert map to slice
	summaries := make([]Summary, 0, len(summaryMap))
	for _, summary := range summaryMap {
		summaries = append(summaries, *summary)
	}

	return summaries, nil
}

func SubGetExamCityCenters(ctx context.Context, client *ent.Client, examYear string, examCode int32) ([]*ent.Center, int32, string, bool, error) {
	if examCode == 0 {
		return nil, 400, " -STR001", false, errors.New("please provide a valid exam code")
	}

	centers, err := client.Center.Query().
		Where(
			center.ExamYearEQ(examYear),
			center.ExamCodeEQ(examCode),
		).
		WithExamscentres().
		All(context.Background())
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	return centers, 200, "", true, nil
}

func SubGetExamCityDivisions(ctx context.Context, client *ent.Client, examYear string, examCode int32, naFacilityId string) ([]*ent.Center, int32, string, bool, error) {
	if examCode == 0 {
		return nil, 400, " -STR001", false, errors.New("please provide a valid exam code")
	}

	centers, err := client.Center.Query().
		Where(
			center.ExamYearEQ(examYear),
			center.ExamCodeEQ(examCode),
			center.NodalOfficeFacilityIdEQ(naFacilityId),
		).
		WithExamscentres().
		All(context.Background())
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	return centers, 200, "", true, nil
}

func GetCandidatePendingApplicationsWithDays(ctx context.Context, client *ent.Client, examCode int32, examYear string, nodalOfficeFacilityID string, daysPending int) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, error) {
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, errors.New("please provide a valid exam code")
	}

	now := time.Now().UTC()
	currentDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	var applicationsIP []*ent.Exam_Applications_IP
	var applicationsPS []*ent.Exam_Applications_PS
	var applicationsGdsPA []*ent.Exam_Applications_GDSPA
	var applicationsPmPa []*ent.Exam_Applications_PMPA
	var applicationsGdsPm []*ent.Exam_Applications_GDSPM
	var applicationsMtsPm []*ent.Exam_Application_MTSPMMG

	var err error

	switch examCode {
	case 1:
		applicationsPS, err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("PendingWithCandidate"),
			).
			All(ctx)
	case 2:
		applicationsIP, err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("PendingWithCandidate"),
			).
			All(ctx)
	case 3:
		applicationsPmPa, err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("PendingWithCandidate"),
			).
			All(ctx)
	case 4:
		applicationsGdsPA, err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("PendingWithCandidate"),
			).
			All(ctx)
	case 5:
		applicationsMtsPm, err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("PendingWithCandidate"),
			).
			All(ctx)
	case 6:
		applicationsGdsPm, err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("PendingWithCandidate"),
			).
			All(ctx)
	default:
		return nil, nil, nil, nil, nil, nil, errors.New("exam code invalid")
	}

	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	filterApplicationsByDays := func(applications interface{}, currentDate time.Time, daysPending int) interface{} {
		switch v := applications.(type) {
		case []*ent.Exam_Applications_IP:
			var filtered []*ent.Exam_Applications_IP
			for _, application := range v {
				submittedDate := time.Date(application.CADate.Year(), application.CADate.Month(), application.CADate.Day(), 0, 0, 0, 0, time.UTC)
				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Applications_PS:
			var filtered []*ent.Exam_Applications_PS
			for _, application := range v {
				submittedDate := time.Date(application.CADate.Year(), application.CADate.Month(), application.CADate.Day(), 0, 0, 0, 0, time.UTC)

				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Applications_GDSPA:
			var filtered []*ent.Exam_Applications_GDSPA
			for _, application := range v {
				submittedDate := time.Date(application.CADate.Year(), application.CADate.Month(), application.CADate.Day(), 0, 0, 0, 0, time.UTC)

				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Applications_PMPA:
			var filtered []*ent.Exam_Applications_PMPA
			for _, application := range v {
				submittedDate := time.Date(application.CADate.Year(), application.CADate.Month(), application.CADate.Day(), 0, 0, 0, 0, time.UTC)

				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Applications_GDSPM:
			var filtered []*ent.Exam_Applications_GDSPM
			for _, application := range v {
				submittedDate := time.Date(application.CADate.Year(), application.CADate.Month(), application.CADate.Day(), 0, 0, 0, 0, time.UTC)

				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		case []*ent.Exam_Application_MTSPMMG:
			var filtered []*ent.Exam_Application_MTSPMMG
			for _, application := range v {
				submittedDate := time.Date(application.CADate.Year(), application.CADate.Month(), application.CADate.Day(), 0, 0, 0, 0, time.UTC)

				daysDifference := int(currentDate.Sub(submittedDate).Hours() / 24)
				if daysDifference == daysPending {
					filtered = append(filtered, application)
				}
			}
			return filtered
		}
		return nil
	}

	filteredApplicationsIP := filterApplicationsByDays(applicationsIP, currentDate, daysPending).([]*ent.Exam_Applications_IP)
	filteredApplicationsPS := filterApplicationsByDays(applicationsPS, currentDate, daysPending).([]*ent.Exam_Applications_PS)
	filteredApplicationsGdsPA := filterApplicationsByDays(applicationsGdsPA, currentDate, daysPending).([]*ent.Exam_Applications_GDSPA)
	filteredApplicationsPmPa := filterApplicationsByDays(applicationsPmPa, currentDate, daysPending).([]*ent.Exam_Applications_PMPA)
	filteredApplicationsGdsPm := filterApplicationsByDays(applicationsGdsPm, currentDate, daysPending).([]*ent.Exam_Applications_GDSPM)
	filteredApplicationsMtsPm := filterApplicationsByDays(applicationsMtsPm, currentDate, daysPending).([]*ent.Exam_Application_MTSPMMG)

	if len(filteredApplicationsIP) == 0 && len(filteredApplicationsPS) == 0 && len(filteredApplicationsGdsPA) == 0 && len(filteredApplicationsPmPa) == 0 && len(filteredApplicationsGdsPm) == 0 && len(filteredApplicationsMtsPm) == 0 {
		return nil, nil, nil, nil, nil, nil, errors.New("no applications found for the provided details")
	}

	return filteredApplicationsIP, filteredApplicationsPS, filteredApplicationsGdsPA, filteredApplicationsPmPa, filteredApplicationsGdsPm, filteredApplicationsMtsPm, nil
}

type ApplicationPendingSummary struct {
	ControllingFacilityID   string `json:"controlling_office_facility_id"`
	ControllingFacilityName string `json:"controlling_office_name"`
	OneDayCount             int    `json:"one_day_count"`
	TwoDayCount             int    `json:"two_day_count"`
	ThreeDayCount           int    `json:"three_day_count"`
	FourDayCount            int    `json:"four_day_count"`
	FiveDayCount            int    `json:"five_day_count"`
	SixDayCount             int    `json:"six_day_count"`
	SevenDayCount           int    `json:"seven_day_count"`
	MoreThanSevenDayCount   int    `json:"more_than_seven_day_count"`
}

type ApplicationPendingSubSummary struct {
	ControllingFacilityID   string `json:"controlling_office_facility_id"`
	ControllingFacilityName string `json:"controlling_office_name"`
	Count                   int    `json:"count"`
}

func SubGetCandidatePendingApplicationsWithCA(ctx context.Context, client *ent.Client, examCode int32, examYear string, nodalOfficeFacilityID string) ([]ApplicationPendingSummary, int32, string, bool, error) {
	if examCode == 0 {
		return nil, 400, " -STR001", false, errors.New("please provide a valid exam code")
	}

	now := time.Now().UTC()

	oneDayAgo := now.Add(-24 * time.Hour)
	twoDaysAgo := now.Add(-2 * 24 * time.Hour)
	threeDaysAgo := now.Add(-3 * 24 * time.Hour)
	fourDaysAgo := now.Add(-4 * 24 * time.Hour)
	fiveDaysAgo := now.Add(-5 * 24 * time.Hour)
	sixDaysAgo := now.Add(-6 * 24 * time.Hour)
	sevenDaysAgo := now.Add(-7 * 24 * time.Hour)

	var err error
	var oneDayResults []ApplicationPendingSubSummary
	var twoDaysResults []ApplicationPendingSubSummary
	var threeDaysResults []ApplicationPendingSubSummary
	var fourDaysResults []ApplicationPendingSubSummary
	var fiveDaysResults []ApplicationPendingSubSummary
	var sixDaysResults []ApplicationPendingSubSummary
	var sevenDaysResults []ApplicationPendingSubSummary
	var moreThanSevenDaysResults []ApplicationPendingSubSummary

	switch examCode {
	case 1:
		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR002", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR003", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR004", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR005", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR006", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR007", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR008", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR009", false, err
		}

	case 2:
		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR010", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR011", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR012", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR013", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR014", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR015", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR016", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR017", false, err
		}
	case 3:
		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR018", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR019", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR020", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR021", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR022", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR023", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR024", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR025", false, err
		}
	case 4:
		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(exam_applications_gdspa.FieldControllingOfficeFacilityID, exam_applications_gdspa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR026", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(exam_applications_gdspa.FieldControllingOfficeFacilityID, exam_applications_gdspa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR027", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(exam_applications_gdspa.FieldControllingOfficeFacilityID, exam_applications_gdspa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR028", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(exam_applications_gdspa.FieldControllingOfficeFacilityID, exam_applications_gdspa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR029", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(exam_applications_gdspa.FieldControllingOfficeFacilityID, exam_applications_gdspa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR030", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(exam_applications_gdspa.FieldControllingOfficeFacilityID, exam_applications_gdspa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR031", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(exam_applications_gdspa.FieldControllingOfficeFacilityID, exam_applications_gdspa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR032", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(exam_applications_gdspa.FieldControllingOfficeFacilityID, exam_applications_gdspa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR033", false, err
		}
	case 5:
		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR034", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(oneDayAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR035", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(twoDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR036", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(threeDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR037", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(fourDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR038", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR039", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(sixDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR040", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR041", false, err
		}
	case 6:
		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR042", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR043", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR044", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR045", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR046", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR047", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR048", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR049", false, err
		}
	default:
		return nil, 400, " -STR050", false, errors.New("exam code invalid")
	}

	// Combine results
	summaryMap := make(map[string]*ApplicationPendingSummary)

	for _, r := range oneDayResults {
		controllingfacilityid := r.ControllingFacilityID

		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &ApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].OneDayCount = r.Count
	}

	for _, r := range twoDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &ApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].TwoDayCount = r.Count
	}

	for _, r := range threeDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &ApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].ThreeDayCount = r.Count
	}

	for _, r := range fourDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &ApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].FourDayCount = r.Count
	}

	for _, r := range fiveDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &ApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].FiveDayCount = r.Count
	}

	for _, r := range sixDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &ApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].SixDayCount = r.Count
	}

	for _, r := range sevenDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &ApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].SevenDayCount = r.Count
	}

	for _, r := range moreThanSevenDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &ApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
		}
		summaryMap[controllingfacilityid].MoreThanSevenDayCount = r.Count
	}
	// Convert map to slice
	summaries := make([]ApplicationPendingSummary, 0, len(summaryMap))
	for _, summary := range summaryMap {
		summaries = append(summaries, *summary)
	}

	return summaries, 200, "", true, nil
}

type DTNAApplicationPendingSummary struct {
	NodalFacilityID         string `json:"nodal_office_facility_id"`
	NodalFacilityName       string `json:"nodal_office_name"`
	ControllingFacilityID   string `json:"controlling_office_facility_id"`
	ControllingFacilityName string `json:"controlling_office_name"`
	OneDayCount             int    `json:"one_day_count"`
	TwoDayCount             int    `json:"two_day_count"`
	ThreeDayCount           int    `json:"three_day_count"`
	FourDayCount            int    `json:"four_day_count"`
	FiveDayCount            int    `json:"five_day_count"`
	SixDayCount             int    `json:"six_day_count"`
	SevenDayCount           int    `json:"seven_day_count"`
	MoreThanSevenDayCount   int    `json:"more_than_seven_day_count"`
}

type DTNAApplicationPendingSubSummary struct {
	NodalFacilityID         string `json:"nodal_office_facility_id"`
	NodalFacilityName       string `json:"nodal_office_name"`
	ControllingFacilityID   string `json:"controlling_office_facility_id"`
	ControllingFacilityName string `json:"controlling_office_name"`
	Count                   int    `json:"count"`
}

func SubGetCandidatePendingApplicationsWithCADT(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]DTNAApplicationPendingSummary, int32, string, bool, error) {
	if examCode == 0 {
		return nil, 400, " -STR001", false, errors.New("please provide a valid exam code")
	}

	now := time.Now().UTC()

	oneDayAgo := now.Add(-24 * time.Hour)
	twoDaysAgo := now.Add(-2 * 24 * time.Hour)
	threeDaysAgo := now.Add(-3 * 24 * time.Hour)
	fourDaysAgo := now.Add(-4 * 24 * time.Hour)
	fiveDaysAgo := now.Add(-5 * 24 * time.Hour)
	sixDaysAgo := now.Add(-6 * 24 * time.Hour)
	sevenDaysAgo := now.Add(-7 * 24 * time.Hour)

	var err error
	var oneDayResults []DTNAApplicationPendingSubSummary
	var twoDaysResults []DTNAApplicationPendingSubSummary
	var threeDaysResults []DTNAApplicationPendingSubSummary
	var fourDaysResults []DTNAApplicationPendingSubSummary
	var fiveDaysResults []DTNAApplicationPendingSubSummary
	var sixDaysResults []DTNAApplicationPendingSubSummary
	var sevenDaysResults []DTNAApplicationPendingSubSummary
	var moreThanSevenDaysResults []DTNAApplicationPendingSubSummary

	switch examCode {
	case 1:
		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR002", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR003", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR004", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR005", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR006", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR007", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR008", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldControllingOfficeFacilityID, exam_applications_ps.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR009", false, err
		}

	case 2:
		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR010", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR011", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR012", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR013", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR014", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR015", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR016", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldControllingOfficeFacilityID, exam_applications_ip.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR017", false, err
		}
	case 3:
		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR018", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR019", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR020", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR021", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR022", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR023", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR024", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR025", false, err
		}
	case 4:
		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeFacilityID, exam_applications_gdspa.FieldNodalOfficeName,
				exam_applications_gdspa.FieldControllingOfficeFacilityID, exam_applications_gdspa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR026", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR027", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR028", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR029", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR030", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR031", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR032", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID, exam_applications_pmpa.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR033", false, err
		}
	case 5:
		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR034", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(oneDayAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR035", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(twoDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR036", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(threeDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR037", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(fourDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR038", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR039", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(sixDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR040", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldControllingOfficeFacilityID, exam_application_mtspmmg.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR041", false, err
		}
	case 6:
		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR042", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR043", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR044", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR045", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR046", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR047", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR048", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldControllingOfficeFacilityID, exam_applications_gdspm.FieldControllingOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR049", false, err
		}
	default:
		return nil, 400, " -STR050", false, errors.New("exam code invalid")
	}

	// Combine results
	summaryMap := make(map[string]*DTNAApplicationPendingSummary)

	for _, r := range oneDayResults {
		controllingfacilityid := r.ControllingFacilityID

		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &DTNAApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].OneDayCount = r.Count
	}
	for _, r := range twoDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &DTNAApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].TwoDayCount = r.Count
	}

	for _, r := range threeDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &DTNAApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].ThreeDayCount = r.Count
	}

	for _, r := range fourDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &DTNAApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].FourDayCount = r.Count
	}

	for _, r := range fiveDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &DTNAApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].FiveDayCount = r.Count
	}

	for _, r := range sixDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &DTNAApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].SixDayCount = r.Count
	}

	for _, r := range sevenDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &DTNAApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].SevenDayCount = r.Count
	}

	for _, r := range moreThanSevenDaysResults {
		controllingfacilityid := r.ControllingFacilityID
		if _, exists := summaryMap[controllingfacilityid]; !exists {
			summaryMap[controllingfacilityid] = &DTNAApplicationPendingSummary{ControllingFacilityID: controllingfacilityid}
			summaryMap[controllingfacilityid].ControllingFacilityName = r.ControllingFacilityName
			summaryMap[controllingfacilityid].NodalFacilityID = r.NodalFacilityID
			summaryMap[controllingfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[controllingfacilityid].MoreThanSevenDayCount = r.Count
	}
	// Convert map to slice
	summaries := make([]DTNAApplicationPendingSummary, 0, len(summaryMap))
	for _, summary := range summaryMap {
		summaries = append(summaries, *summary)
	}

	return summaries, 200, "", true, nil
}

type DTApplicationPendingSummary struct {
	NodalFacilityID       string `json:"nodal_office_facility_id"`
	NodalFacilityName     string `json:"nodal_office_name"`
	OneDayCount           int    `json:"one_day_count"`
	TwoDayCount           int    `json:"two_day_count"`
	ThreeDayCount         int    `json:"three_day_count"`
	FourDayCount          int    `json:"four_day_count"`
	FiveDayCount          int    `json:"five_day_count"`
	SixDayCount           int    `json:"six_day_count"`
	SevenDayCount         int    `json:"seven_day_count"`
	MoreThanSevenDayCount int    `json:"more_than_seven_day_count"`
}

type DTApplicationPendingSubSummary struct {
	NodalFacilityID   string `json:"nodal_office_facility_id"`
	NodalFacilityName string `json:"nodal_office_name"`
	Count             int    `json:"count"`
}

func SubGetCandidatePendingApplicationsWithNADT(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]DTApplicationPendingSummary, int32, string, bool, error) {
	if examCode == 0 {
		return nil, 400, " -STR001", false, errors.New("please provide a valid exam code")
	}

	now := time.Now().UTC()

	oneDayAgo := now.Add(-24 * time.Hour)
	twoDaysAgo := now.Add(-2 * 24 * time.Hour)
	threeDaysAgo := now.Add(-3 * 24 * time.Hour)
	fourDaysAgo := now.Add(-4 * 24 * time.Hour)
	fiveDaysAgo := now.Add(-5 * 24 * time.Hour)
	sixDaysAgo := now.Add(-6 * 24 * time.Hour)
	sevenDaysAgo := now.Add(-7 * 24 * time.Hour)

	var err error
	var oneDayResults []DTApplicationPendingSubSummary
	var twoDaysResults []DTApplicationPendingSubSummary
	var threeDaysResults []DTApplicationPendingSubSummary
	var fourDaysResults []DTApplicationPendingSubSummary
	var fiveDaysResults []DTApplicationPendingSubSummary
	var sixDaysResults []DTApplicationPendingSubSummary
	var sevenDaysResults []DTApplicationPendingSubSummary
	var moreThanSevenDaysResults []DTApplicationPendingSubSummary

	switch examCode {
	case 1:
		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR002", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR003", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR004", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR005", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR006", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR007", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_ps.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR008", false, err
		}

		err = client.Exam_Applications_PS.Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ps.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeFacilityID, exam_applications_ps.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR009", false, err
		}

	case 2:
		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR010", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR011", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR012", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR013", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR014", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR015", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_ip.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR016", false, err
		}

		err = client.Exam_Applications_IP.Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_ip.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeFacilityID, exam_applications_ip.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR017", false, err
		}
	case 3:
		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR018", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR019", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR020", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR021", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR022", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR023", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_pmpa.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR024", false, err
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_pmpa.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeFacilityID, exam_applications_pmpa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR025", false, err
		}
	case 4:
		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeFacilityID, exam_applications_gdspa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR026", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeFacilityID, exam_applications_gdspa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR027", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeFacilityID, exam_applications_gdspa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR028", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeFacilityID, exam_applications_gdspa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR029", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeFacilityID, exam_applications_gdspa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR030", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeFacilityID, exam_applications_gdspa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR031", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_gdspa.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeFacilityID, exam_applications_gdspa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR032", false, err
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspa.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeFacilityID, exam_applications_gdspa.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR033", false, err
		}
	case 5:
		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR034", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(oneDayAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR035", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(twoDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR036", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(threeDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR037", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(fourDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR038", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR039", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(sixDaysAgo),
				exam_application_mtspmmg.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR040", false, err
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_application_mtspmmg.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeFacilityID, exam_application_mtspmmg.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR041", false, err
		}
	case 6:
		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateGTE(oneDayAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &oneDayResults)
		if err != nil {
			return nil, 500, " -STR042", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(oneDayAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(twoDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &twoDaysResults)
		if err != nil {
			return nil, 500, " -STR043", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(twoDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(threeDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &threeDaysResults)
		if err != nil {
			return nil, 500, " -STR044", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(threeDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(fourDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fourDaysResults)
		if err != nil {
			return nil, 500, " -STR045", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(fourDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(fiveDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &fiveDaysResults)
		if err != nil {
			return nil, 500, " -STR046", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(fiveDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(sixDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sixDaysResults)
		if err != nil {
			return nil, 500, " -STR047", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(sixDaysAgo),
				exam_applications_gdspm.ApplnSubmittedDateGTE(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &sevenDaysResults)
		if err != nil {
			return nil, 500, " -STR048", false, err
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				exam_applications_gdspm.ApplnSubmittedDateLT(sevenDaysAgo),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeFacilityID, exam_applications_gdspm.FieldNodalOfficeName).
			Aggregate(ent.Count()).
			Scan(ctx, &moreThanSevenDaysResults)
		if err != nil {
			return nil, 500, " -STR049", false, err
		}
	default:
		return nil, 400, " -STR050", false, errors.New("exam code invalid")
	}

	// Combine results
	summaryMap := make(map[string]*DTApplicationPendingSummary)

	for _, r := range oneDayResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTApplicationPendingSummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].OneDayCount = r.Count
	}

	for _, r := range twoDaysResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTApplicationPendingSummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].TwoDayCount = r.Count
	}

	for _, r := range threeDaysResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTApplicationPendingSummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].ThreeDayCount = r.Count
	}

	for _, r := range fourDaysResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTApplicationPendingSummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].FourDayCount = r.Count
	}

	for _, r := range fiveDaysResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTApplicationPendingSummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].FiveDayCount = r.Count
	}

	for _, r := range sixDaysResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTApplicationPendingSummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].SixDayCount = r.Count
	}

	for _, r := range sevenDaysResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTApplicationPendingSummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].SevenDayCount = r.Count
	}

	for _, r := range moreThanSevenDaysResults {
		nodalfacilityid := r.NodalFacilityID
		if _, exists := summaryMap[nodalfacilityid]; !exists {
			summaryMap[nodalfacilityid] = &DTApplicationPendingSummary{NodalFacilityID: nodalfacilityid}
			summaryMap[nodalfacilityid].NodalFacilityName = r.NodalFacilityName
		}
		summaryMap[nodalfacilityid].MoreThanSevenDayCount = r.Count
	}
	// Convert map to slice
	summaries := make([]DTApplicationPendingSummary, 0, len(summaryMap))
	for _, summary := range summaryMap {
		summaries = append(summaries, *summary)
	}

	return summaries, 200, "", true, nil

}

func SubGetAllApplications(ctx context.Context, client *ent.Client, examCode int32, examYear string, circleFacilityID string) ([]*ent.Exam_Applications_IP, []*ent.Exam_Applications_PS, []*ent.Exam_Applications_GDSPA, []*ent.Exam_Applications_PMPA, []*ent.Exam_Applications_GDSPM, []*ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 500, " -STR001", false, errors.New("please provide a valid exam code")
	}

	switch examCode {
	case 2:
		applications, err := client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
				exam_applications_ip.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			).
			WithCirclePrefRef().
			WithIPApplicationsRef().
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR002", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR003", false, errors.New("no applications found for the provided details")
		}
		return applications, nil, nil, nil, nil, nil, 200, "", true, nil

	case 1:
		applications, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
				exam_applications_ps.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			).
			WithPSApplicationsRef().
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR004", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR005", false, errors.New("no applications found for the provided details")
		}
		return nil, applications, nil, nil, nil, nil, 200, "", true, nil

	case 4:
		applications, err := client.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_gdspa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
				exam_applications_gdspa.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			).
			WithGDSPAApplicationsRef().
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR006", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR007", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, applications, nil, nil, nil, 200, "", true, nil

	case 3:
		applications, err := client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_pmpa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
				exam_applications_pmpa.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			).
			WithPMPAApplicationsRef().
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR008", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR009", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, applications, nil, nil, 200, "", true, nil

	case 6:
		applications, err := client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.NodalOfficeFacilityID(circleFacilityID),
				exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
				exam_applications_gdspm.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			).
			WithGDSPMApplicationsRef().
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR010", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR011", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, applications, nil, 200, "", true, nil

	case 5:
		applications, err := client.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.NodalOfficeFacilityID(circleFacilityID),
				exam_application_mtspmmg.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
				exam_application_mtspmmg.RecommendedStatusIn("Recommended", "Provisionally Recommended"),
			).
			WithMTSPMMGApplicationsRef().
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR012", false, err
		}
		if len(applications) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR013", false, errors.New("no applications found for the provided details")
		}
		return nil, nil, nil, nil, nil, applications, 200, "", true, nil

	default:
		return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, errors.New("invalid exam code")
	}
}
