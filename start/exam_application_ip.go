package start

import (
	"context"
	"sort"

	"errors"

	"fmt"
	"log"

	"strconv"

	"recruit/ent"
	"recruit/ent/center"
	"recruit/ent/circlesummaryforno"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/exam_applications_gdspa"
	"recruit/ent/exam_applications_gdspm"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_pmpa"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/examcenterhall"
	"recruit/ent/schema"

	"recruit/ent/recommendationsipapplications"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"

	"time"
)

func CreateIPApplications(client *ent.Client, newAppln *ca_reg.ApplicationIp) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	if err := validateTempHallTicket(newAppln.TempHallTicket, newAppln.EmployeeID); err != nil {
		return nil, 422, " -STR001", false, fmt.Errorf("issue for employee %d with temp hall ticket number: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
	}
	if err := validateApplicationData(newAppln); err != nil {
		return nil, 422, " -STR002", false, fmt.Errorf("circle preference values are missing")
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}
	defer func() {
		handleTransaction(tx, &err)
	}()

	statuses := []string{"CAVerificationPending", "ResubmitCAVerificationPending", "PendingWithCandidate", "VerifiedByCA"}
	existing, status, stgError, err := checkIfApplicationExists(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear, newAppln.ExamCode, statuses)
	if status == 500 {
		return nil, 500 + status, " -STR004 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR005 " + stgError, false, err

	}
	// if err != nil {
	// 	return nil, 500 + status, " -STR005 " + stgError, false, err
	// }
	if existing {

		return nil, 422 + status, " -STR006 " + stgError, false, fmt.Errorf("already application submitted for this candidate")
	}
	var currentTime = time.Now().UTC().Truncate(time.Second) // Ensure UTC for consistent comparison
	currentTime = currentTime.Add(5*time.Hour + 30*time.Minute)
	applicationLastDate := newAppln.ApplicationLastDate.UTC().Truncate(time.Second)

	fmt.Print("current time: ", currentTime, " application last date: ", applicationLastDate, "date from payload", newAppln.ApplicationLastDate)
	if currentTime.After(applicationLastDate) {
		return nil, 422, " -STR007", false, fmt.Errorf("application submission deadline has passed as current time is %v", currentTime)
	}
	applicationNumber, err := util.GenerateApplicationNumber(client, newAppln.EmployeeID, newAppln.ExamYear, "IP") //ExamGenCode = IP
	if err != nil {
		return nil, 500, " -STR007", false, err
	}

	createdAppln, status, stgError, err := saveApplication(tx, newAppln, applicationNumber, newAppln.ExamCode, ctx)
	if err != nil {
		return nil, 500 + status, " -STR008 " + stgError, false, err
	}

	return createdAppln, 200, "", true, nil
}

type StrucMappingIdentificationNumberResult struct {
	NodalOfficeFacilityId string `json:"nodal_office_facility_id"`
	StartNo               int32  `json:"start_no"`
	EndNo                 int32  `json:"end_no"`
	UpdateStatus          bool   `json:"update_status"`
}

func convertMappingIdentificationNumbers(
	mappings []ca_reg.StrucMappingIdentificationNumber,
) []schema.StrucMappingIdentificationNumber {
	converted := make([]schema.StrucMappingIdentificationNumber, len(mappings))
	for i, mapping := range mappings {
		converted[i] = schema.StrucMappingIdentificationNumber{
			NodalOfficeFacilityId: mapping.NodalOfficeFacilityId,
			StartNo:               mapping.StartNo,
			EndNo:                 mapping.EndNo,
		}
	}
	return converted
}

func SubCreateExamCenterHall(client *ent.Client, newExamcenterhall *ca_reg.StruExamCenterHall) ([]*StrucMappingIdentificationNumberResult, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	//transaction implementation-------------

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}

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

	existing, err := tx.ExamCenterHall.
		Query().
		Where(
			examcenterhall.ExamYearEQ(newExamcenterhall.ExamYear),
			examcenterhall.ExamCodeEQ(newExamcenterhall.ExamCode),
			examcenterhall.AdminCircleOfficeIDEQ(newExamcenterhall.AdminCircleOfficeID),
			examcenterhall.HallNameEQ(newExamcenterhall.HallName),
			examcenterhall.StatusEQ("active"),
		).
		Exist(ctx)

	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if existing {
		return nil, 422, " -STR003", false, errors.New("already this HallName mapped for candidates")
	}

	// Use the generated application number
	currentTime := time.Now().Truncate(time.Second)

	examCenterHall, err := tx.ExamCenterHall.
		Create().
		SetCenterCode(int32(newExamcenterhall.CenterCode)).
		SetCityID(int32(newExamcenterhall.CityID)).
		SetExamCenterName(newExamcenterhall.ExamCenterName).
		SetExamYear(newExamcenterhall.ExamYear).
		SetExamCode(int32(newExamcenterhall.ExamCode)).
		SetExamName(newExamcenterhall.ExamName).
		SetCenterCityName(newExamcenterhall.CenterCityName).
		SetConductedByFacilityID(newExamcenterhall.ConductedByFacilityID).
		SetConductedBy(newExamcenterhall.ConductedBy).
		SetHallName(newExamcenterhall.HallName).
		SetAdminCircleOfficeID(newExamcenterhall.AdminCircleOfficeID).
		SetMappingIdentificationNumber(convertMappingIdentificationNumbers(newExamcenterhall.MappingIdentificationNumber)).
		SetStatus("active").
		SetCreatedById(newExamcenterhall.CreatedById).
		SetCreatedByUserName(newExamcenterhall.CreatedByUserName).
		SetCreatedByEmpId(newExamcenterhall.CreatedByEmpId).
		SetCreatedByDesignation(newExamcenterhall.CreatedByDesignation).
		SetCreatedDate(currentTime).
		SetNoSeats(newExamcenterhall.NoSeats).
		SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR04", false, err
	}
	var resultExamcenterhall []*StrucMappingIdentificationNumberResult

	// Save the PlaceOfPreferenceIP records.
	//circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(newAppln.Edges.CircleData))
	//for i, circlePrefRef := range newAppln.Edges.CircleData {
	for i, circlePrefRef := range newExamcenterhall.MappingIdentificationNumber {

		if len(circlePrefRef.NodalOfficeFacilityId) == 0 || circlePrefRef.StartNo == 0 || circlePrefRef.EndNo == 0 {
			return nil, 400, " -STR005", false, fmt.Errorf("invalid mapping identification number at index %d", i)
		}

		//fmt.Println(newExamcenterhall.ExamYear, newExamcenterhall.ExamCode, newExamcenterhall.CenterCode, circlePrefRef.NodalOfficeFacilityId, circlePrefRef.StartNo, circlePrefRef.EndNo)
		hallassigned, err := tx.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamYearEQ(newExamcenterhall.ExamYear),
				exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
				exam_applications_ip.HallTicketNumberNEQ(""),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamCodeEQ(newExamcenterhall.ExamCode),
				exam_applications_ip.CenterCodeEQ(newExamcenterhall.CenterCode),
				exam_applications_ip.LienControllingOfficeIDEQ(circlePrefRef.NodalOfficeFacilityId),
				exam_applications_ip.HallTicketGeneratedFlag(true),
				exam_applications_ip.HallIdentificationNumberGTE(circlePrefRef.StartNo),
				exam_applications_ip.HallIdentificationNumberLTE(circlePrefRef.EndNo),
				exam_applications_ip.HallNameEQ(newExamcenterhall.HallName),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR006", false, err
		}

		if hallassigned {
			return nil, 422, " -STR007", false, fmt.Errorf("already this identification numbers mapped for candidates")
		}

		_, err = tx.Exam_Applications_IP.
			Update().
			SetHallName(newExamcenterhall.HallName).
			SetExamCenterHall(examCenterHall.ID).
			Where(
				exam_applications_ip.ExamYearEQ(newExamcenterhall.ExamYear),
				exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
				exam_applications_ip.HallTicketNumberNEQ(""),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamCodeEQ(newExamcenterhall.ExamCode),
				exam_applications_ip.CenterCodeEQ(newExamcenterhall.CenterCode),
				exam_applications_ip.LienControllingOfficeIDEQ(circlePrefRef.NodalOfficeFacilityId),
				exam_applications_ip.HallTicketGeneratedFlag(true),
				exam_applications_ip.HallIdentificationNumberGTE(circlePrefRef.StartNo),
				exam_applications_ip.HallIdentificationNumberLTE(circlePrefRef.EndNo),
			).
			Save(ctx)
		if err != nil {
			return nil, 500, " -STR008", false, err
		}
		resultExamcenterhall = append(resultExamcenterhall, &StrucMappingIdentificationNumberResult{
			NodalOfficeFacilityId: circlePrefRef.NodalOfficeFacilityId,
			StartNo:               circlePrefRef.StartNo,
			EndNo:                 circlePrefRef.EndNo,
			UpdateStatus:          true,
		})
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR009", false, err
	}
	return resultExamcenterhall, 200, "", true, nil
}

// func SubResetExamCenterHall(client *ent.Client, newExamcenterhall *ca_reg.StruExamCenterHall) ([]*StrucMappingIdentificationNumberResult, int32, string, bool, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
// 	defer cancel()

// 	// Start transaction
// 	tx, err := client.Tx(ctx)
// 	if err != nil {
// 		return nil, 500, " -STR001", false, err
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

// 	// Create status string with timestamp
// 	stat := "inactive_" + time.Now().Format("20060102150405")

// 	// Check if the ExamCenterHall already exists
// 	existing, err := tx.ExamCenterHall.
// 		Query().
// 		Where(
// 			examcenterhall.ExamYearEQ(newExamcenterhall.ExamYear),
// 			examcenterhall.ExamCodeEQ(newExamcenterhall.ExamCode),
// 			examcenterhall.AdminCircleOfficeIDEQ(newExamcenterhall.AdminCircleOfficeID),
// 			examcenterhall.HallNameEQ(newExamcenterhall.HallName),
// 			examcenterhall.StatusEQ("active"),
// 		).
// 		Only(ctx)
// 	fmt.Println("im here1")

// 	if err != nil && !ent.IsNotFound(err) {
// 		return nil, 500, " -STR002", false, err
// 	}
// 	fmt.Println("im here2")
// 	if existing != nil {
// 		// Update existing ExamCenterHall record
// 		_, err = tx.ExamCenterHall.
// 			Update().
// 			SetStatus(stat).
// 			SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
// 			Save(ctx)

// 		if err != nil {
// 			return nil, 500, " -STR003", false, err
// 		}
// 	}
// 	fmt.Println("im here3")
// 	var resultExamcenterhall []*StrucMappingIdentificationNumberResult

// 	for i, circlePrefRef := range newExamcenterhall.MappingIdentificationNumber {
// 		if len(circlePrefRef.NodalOfficeFacilityId) == 0 || circlePrefRef.StartNo == 0 || circlePrefRef.EndNo == 0 {
// 			return nil, 400, " -STR005", false, fmt.Errorf("invalid mapping identification number at index %d", i)
// 		}
// 		fmt.Println("im here4")
// 		// Check if there are already assigned halls in the specified range
// 		hallassigned, err := tx.Exam_Applications_IP.
// 			Query().
// 			Where(
// 				exam_applications_ip.ExamYearEQ(newExamcenterhall.ExamYear),
// 				exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 				exam_applications_ip.HallTicketNumberNEQ(""),
// 				exam_applications_ip.StatusEQ("active"),
// 				exam_applications_ip.ExamCodeEQ(newExamcenterhall.ExamCode),
// 				exam_applications_ip.CenterCodeEQ(newExamcenterhall.CenterCode),
// 				exam_applications_ip.LienControllingOfficeIDEQ(circlePrefRef.NodalOfficeFacilityId),
// 				exam_applications_ip.HallTicketGeneratedFlag(true),
// 				exam_applications_ip.HallIdentificationNumberGTE(circlePrefRef.StartNo),
// 				exam_applications_ip.HallIdentificationNumberLTE(circlePrefRef.EndNo),
// 				exam_applications_ip.HallNameEQ(newExamcenterhall.HallName),
// 			).
// 			All(ctx)
// 		fmt.Println("im here5")
// 		if err != nil {
// 			return nil, 500, " -STR006", false, err
// 		}
// 		fmt.Println("im here6")
// 		if hallassigned != nil {
// 			// Update existing Exam_Applications_IP records to remove hall assignments
// 			_, err = tx.Exam_Applications_IP.
// 				Update().
// 				SetHallName("").
// 				SetExamCenterHall(0). // Set to nil or appropriate zero value to clear the foreign key reference
// 				Where(
// 					exam_applications_ip.ExamYearEQ(newExamcenterhall.ExamYear),
// 					exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
// 					exam_applications_ip.HallTicketNumberNEQ(""),
// 					exam_applications_ip.StatusEQ("active"),
// 					exam_applications_ip.ExamCodeEQ(newExamcenterhall.ExamCode),
// 					exam_applications_ip.CenterCodeEQ(newExamcenterhall.CenterCode),
// 					exam_applications_ip.LienControllingOfficeIDEQ(circlePrefRef.NodalOfficeFacilityId),
// 					exam_applications_ip.HallTicketGeneratedFlag(true),
// 					exam_applications_ip.HallIdentificationNumberGTE(circlePrefRef.StartNo),
// 					exam_applications_ip.HallIdentificationNumberLTE(circlePrefRef.EndNo),
// 				).
// 				Save(ctx)
// 			fmt.Println("im here7")

// 			if err != nil {
// 				return nil, 500, " -STR008", false, err
// 			}
// 		}

// 		resultExamcenterhall = append(resultExamcenterhall, &StrucMappingIdentificationNumberResult{
// 			NodalOfficeFacilityId: circlePrefRef.NodalOfficeFacilityId,
// 			StartNo:               circlePrefRef.StartNo,
// 			EndNo:                 circlePrefRef.EndNo,
// 			UpdateStatus:          true,
// 		})
// 	}

// 	if err = tx.Commit(); err != nil {
// 		tx.Rollback()
// 		return nil, 500, " -STR009", false, err
// 	}

// 	return resultExamcenterhall, 200, "", true, nil
// }

func SubResetExamCenterHall(client *ent.Client, newExamcenterhall *ca_reg.StruExamCenterHallReset) ([]*StrucMappingIdentificationNumberResult, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Start transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}

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

	// Create status string with timestamp
	stat := "inactive_" + time.Now().Format("20060102150405")

	// Check if the ExamCenterHall already exists
	existing, err := tx.ExamCenterHall.
		Query().
		Where(
			examcenterhall.ExamYearEQ(newExamcenterhall.ExamYear),
			examcenterhall.ExamCodeEQ(newExamcenterhall.ExamCode),
			examcenterhall.AdminCircleOfficeIDEQ(newExamcenterhall.AdminCircleOfficeID),
			examcenterhall.HallNameEQ(newExamcenterhall.HallName),
			examcenterhall.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, 500, " -STR002", false, err
	}

	if existing != nil {
		// Update existing ExamCenterHall record
		_, err = tx.ExamCenterHall.
			Update().
			SetStatus(stat).
			SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
			Save(ctx)

		if err != nil {
			return nil, 500, " -STR003", false, err
		}
	}

	var resultExamcenterhall []*StrucMappingIdentificationNumberResult

	// for i, circlePrefRef := range newExamcenterhall.MappingIdentificationNumber {
	// 	if len(circlePrefRef.NodalOfficeFacilityId) == 0 || circlePrefRef.StartNo == 0 || circlePrefRef.EndNo == 0 {
	// 		return nil, 400, " -STR004", false, fmt.Errorf("invalid mapping identification number at index %d", i)
	// 	}

	// Fetch existing hall assignments
	hallAssignments, err := tx.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamYearEQ(newExamcenterhall.ExamYear),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.HallTicketNumberNEQ(""),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamCodeEQ(newExamcenterhall.ExamCode),
			exam_applications_ip.CenterCodeEQ(newExamcenterhall.CenterCode),
			//exam_applications_ip.LienControllingOfficeIDEQ(circlePrefRef.NodalOfficeFacilityId),
			exam_applications_ip.HallTicketGeneratedFlag(true),
			//	exam_applications_ip.HallIdentificationNumberGTE(circlePrefRef.StartNo),
			///	exam_applications_ip.HallIdentificationNumberLTE(circlePrefRef.EndNo),
			exam_applications_ip.HallNameEQ(newExamcenterhall.HallName),
		).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	if len(hallAssignments) > 0 {
		for _, assignment := range hallAssignments {
			_, err = tx.Exam_Applications_IP.
				UpdateOneID(assignment.ID).
				SetHallName("").
				SetExamCenterHall(0).
				Save(ctx)

			if err != nil {

				return nil, 500, " -STR006", false, err
			}
		}

		resultExamcenterhall = append(resultExamcenterhall, &StrucMappingIdentificationNumberResult{
			// NodalOfficeFacilityId: circlePrefRef.NodalOfficeFacilityId,
			// StartNo:               circlePrefRef.StartNo,
			// EndNo:                 circlePrefRef.EndNo,
			UpdateStatus: true,
		})
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR009", false, err
	}

	return resultExamcenterhall, 200, "", true, nil
}

func QueryIPExamApplicationsByEmpIDNew(ctx context.Context, client *ent.Client, empid int64, examYear string) (*ent.Exam_Applications_IP, int32, string, bool, error) {

	newAppln, err := client.Exam_Applications_IP.
		Query().
		Where((exam_applications_ip.EmployeeIDEQ(empid)),
			exam_applications_ip.ExamYear(examYear),
			exam_applications_ip.StatusEQ("active")).
		//Order(ent.Desc(exam_applications_ip.FieldID)).
		WithCirclePrefRef().
		WithIPApplicationsRef().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR006", false, errors.New("no application exists ")
		} else {
			return nil, 500, " -STR006", false, err
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

	// Sort circlePrefs by PlacePrefNo in ascending order
	sort.Slice(circlePrefs, func(i, j int) bool {
		return circlePrefs[i].PlacePrefNo < circlePrefs[j].PlacePrefNo
	})
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
	return newAppln, 200, "", true, nil
}

// func UpdateIPApplicationRemarks(client *ent.Client, newAppln *ca_reg.VerifyApplicationIp, nonQualifyService string) (*ent.Exam_Applications_IP, int32, string, bool, error) {
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
// 	oldAppln, err := tx.Exam_Applications_IP.
// 		Query().
// 		Where(
// 			exam_applications_ip.EmployeeIDEQ(newAppln.EmployeeID),
// 			exam_applications_ip.ExamYearEQ(newAppln.ExamYear),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		Only(ctx)
// 	if err != nil {
// 		if ent.IsNotFound(err) {
// 			return nil, 422, " -STR003", false, errors.New("no active application for this candidate ")
// 		} else {
// 			return nil, 500, " -STR004", false, err
// 		}
// 	}

// 	// Define variables for updated application and status.
// 	var updatedAppln *ent.Exam_Applications_IP
// 	var applicationStatus string

// 	currentTime := time.Now().Truncate(time.Second)
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

// 	updatedAppln, err = tx.Exam_Applications_IP.
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
// 		SetCadre(oldAppln.Cadre).
// 		SetCandidateRemarks(oldAppln.CandidateRemarks).
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
// 		SetDCInPresentCadre(oldAppln.DCInPresentCadre).
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
// 		SetExamCityCenterCode(oldAppln.ExamCityCenterCode).
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
// 		SetNodalOfficeFacilityID(oldAppln.NodalOfficeFacilityID).
// 		SetNodalOfficeName(oldAppln.NodalOfficeName).
// 		SetOptionUsed(oldAppln.OptionUsed).
// 		SetPhoto(oldAppln.Photo).
// 		SetPhotoPath(oldAppln.PhotoPath).
// 		SetPresentDesignation(oldAppln.PresentDesignation).
// 		SetPresentPostCode(oldAppln.PresentPostCode).
// 		SetPresentPostDescription(oldAppln.PresentPostDescription).
// 		SetReportingOfficeFacilityID(oldAppln.ReportingOfficeFacilityID).
// 		SetReportingOfficeName(oldAppln.ReportingOfficeName).
// 		SetServiceLength(*newAppln.ServiceLength).
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
// 		SetRecommendedStatus(newAppln.RecommendedStatus).
// 		SetPunishmentStatus(newAppln.PunishmentStatus).
// 		SetDisciplinaryCaseStatus(newAppln.DisciplinaryCaseStatus).
// 		Save(ctx)

// 	if err != nil {
// 		return nil, 500, " -STR010", false, err
// 	}

// 	// Update non-qualifying service if applicable.
// 	if nonQualifyService == "Yes" {
// 		_, err = updatedAppln.
// 			Update().
// 			SetNonQualifyingService(*newAppln.NonQualifyingService).
// 			Save(ctx)
// 		if err != nil {
// 			return nil, 500, " -STR011", false, err
// 		}
// 	}
// 	// Save the PlaceOfPreferenceIP records.
// 	circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(newAppln.Edges.CircleData))
// 	for i, circlePrefRef := range newAppln.Edges.CircleData {
// 		if circlePrefRef.PlacePrefNo == 0 {
// 			return nil, 400, " -STR012", false, fmt.Errorf("circle preference value at index %d is nil", i)
// 		}
// 		circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
// 			Create().
// 			SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
// 			SetApplicationID(updatedAppln.ID).
// 			SetEmployeeID(newAppln.EmployeeID).
// 			SetPlacePrefValue(circlePrefRef.PlacePrefValue).
// 			SetUpdatedAt(currentTime).
// 			Save(ctx)
// 		if err != nil {
// 			return nil, 500, " -STR013", false, err
// 		}
// 		circlePrefRefs[i] = circlePrefRefEntity
// 	}

// 	// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
// 	_, err = updatedAppln.
// 		Update().
// 		AddCirclePrefRef(circlePrefRefs...).
// 		Save(ctx)

// 	if err != nil {
// 		return nil, 500, " -STR014", false, err
// 	}
// 	// Save the Recommendation records.
// 	recommendationsRef := make([]*ent.RecommendationsIPApplications, len(newAppln.Edges.ApplicationData))
// 	for i, recommendation := range newAppln.Edges.ApplicationData {
// 		if recommendation.VacancyYear == 0 {
// 			return nil, 400, " -STR015", false, fmt.Errorf("recommendations value at index %d is nil", i)
// 		}

// 		RecommendationsRefEntity, err := tx.RecommendationsIPApplications.
// 			Create().
// 			SetApplicationID(updatedAppln.ID).
// 			SetApplicationStatus("VerifiedRecommendationsByCA").
// 			SetCARecommendations(recommendation.CA_Recommendations).
// 			SetCARemarks(recommendation.CA_Remarks).
// 			SetCAUpdatedAt(currentTime).
// 			SetCAUserName(newAppln.CA_UserName).
// 			SetEmployeeID(updatedAppln.EmployeeID).
// 			SetExamNameCode(updatedAppln.ExamShortName).
// 			SetExamYear(updatedAppln.ExamYear).
// 			SetNORecommendations(recommendation.CA_Recommendations).
// 			SetNOUpdatedAt(currentTime).
// 			SetVacancyYear(recommendation.VacancyYear).
// 			Save(ctx)

// 		if err != nil {
// 			return nil, 500, " -STR016", false, err
// 		}
// 		recommendationsRef[i] = RecommendationsRefEntity
// 	}

// 	updatedAppln.Update().
// 		ClearIPApplicationsRef().
// 		AddIPApplicationsRef(recommendationsRef...).
// 		Save(ctx)
// 	// Commit the transaction.
// 	if err = tx.Commit(); err != nil {
// 		tx.Rollback()
// 		return nil, 500, " -STR012", false, err
// 	}
// 	return updatedAppln, 200, "", true, nil
// }

func UpdateIPApplicationRemarks(client *ent.Client, newAppln *ca_reg.VerifyApplicationIp, nonQualifyService string) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Validate input.
	if newAppln == nil {
		return nil, 400, " -STR001", false, errors.New("payload received is empty")
	}

	// Start a transaction.
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	defer handleTransaction(tx, &err)

	// Fetch the existing application.
	oldAppln, status, stgError, err := fetchExistingApplication(ctx, tx, newAppln.EmployeeID, newAppln.ExamYear)
	if status == 500 {
		return nil, 500 + status, " -STR003 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR004 " + stgError, false, fmt.Errorf("no active application found for this candidate")

	}
	if err != nil {
		return nil, 500 + status, " -STR005 " + stgError, false, err
	}

	// Determine the new application status.
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

	// Create the updated application.
	updatedAppln, err := createUpdatedApplication(ctx, tx, oldAppln, newAppln, applicationStatus)
	if err != nil {
		return nil, 500, " -STR009", false, err
	}

	// Update non-qualifying service if applicable.
	if nonQualifyService == "Yes" {
		_, err = updatedAppln.
			Update().
			SetNonQualifyingService(*newAppln.NonQualifyingService).
			Save(ctx)
		if err != nil {
			return nil, 500, " -STR010", false, err
		}
	}
	currentTime := time.Now().Truncate(time.Second)

	// Save the PlaceOfPreferenceIP records.
	circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(newAppln.Edges.CircleData))
	for i, circlePrefRef := range newAppln.Edges.CircleData {
		if circlePrefRef.PlacePrefNo == 0 {
			return nil, 400, " -STR011", false, fmt.Errorf("circle preference value at index %d is nil", i)
		}
		circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
			Create().
			SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(newAppln.EmployeeID).
			SetPlacePrefValue(circlePrefRef.PlacePrefValue).
			SetUpdatedAt(currentTime).
			Save(ctx)
		if err != nil {
			return nil, 500, " -STR012", false, err
		}
		circlePrefRefs[i] = circlePrefRefEntity
	}

	// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
	_, err = updatedAppln.
		Update().
		AddCirclePrefRef(circlePrefRefs...).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR013", false, err
	}
	// Save the Recommendation records.

	recommendationsRef := make([]*ent.RecommendationsIPApplications, len(newAppln.Edges.ApplicationData))
	for i, recommendation := range newAppln.Edges.ApplicationData {
		if recommendation.VacancyYear == 0 {
			return nil, 400, " -STR014", false, fmt.Errorf("recommendations value at index %d is nil", i)
		}

		RecommendationsRefEntity, err := tx.RecommendationsIPApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetApplicationStatus("VerifiedRecommendationsByCA").
			SetCARecommendations(recommendation.CA_Recommendations).
			SetCARemarks(recommendation.CA_Remarks).
			SetCAUpdatedAt(currentTime).
			SetCAUserName(newAppln.CA_UserName).
			SetEmployeeID(updatedAppln.EmployeeID).
			SetExamNameCode(updatedAppln.ExamShortName).
			SetExamYear(updatedAppln.ExamYear).
			SetNORecommendations(recommendation.CA_Recommendations).
			SetNOUpdatedAt(currentTime).
			SetVacancyYear(recommendation.VacancyYear).
			Save(ctx)

		if err != nil {
			return nil, 500, " -STR015", false, err
		}
		recommendationsRef[i] = RecommendationsRefEntity
	}

	updatedAppln.Update().
		ClearIPApplicationsRef().
		AddIPApplicationsRef(recommendationsRef...).
		Save(ctx)

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR016", false, err
	}
	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB017", false, err
	}
	return appResponse, 200, "", true, nil
}

func SubEmailSmsTriggeringPenidngWithCandidate(ctx context.Context, client *ent.Client, examYear string, nodalOfficeId string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {

	// Check if the EmployeeID exists.
	Appln, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeId),
			exam_applications_ip.ApplicationStatusEQ("PendingWithCandidate"),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}
	if len(Appln) == 0 {
		return nil, 422, " -STR006", false, errors.New("no records found to send reminder for candidates")
	}

	// Define variables for updated application and status.

	return Appln, 200, "", true, nil
}
func SubResetCenterIPApplicationNA(client *ent.Client, applicationRecord *ca_reg.NAApplicationIpCenterChange) (*ent.Exam_Applications_IP, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if applicationRecord == nil {
		return nil, 400, " -STR001", false, errors.New("payload received in empty")
	}
	empID := applicationRecord.EmployeeID
	id1 := applicationRecord.ExamYear
	if empID == 0 {
		return nil, 422, " -STR002", false, errors.New("employee id should not be empty")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR017", false, err
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
	exists, err := tx.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(empID),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamYearEQ(id1),
		).
		Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR003", false, fmt.Errorf("no active application avaiable for this emmployee id: %d in verified state ", empID)
		} else {
			return nil, 500, " -STR004", false, err
		}
	}
	fmt.Println("Exists: ", exists)
	if !exists {
		return nil, 422, " -STR003", false, fmt.Errorf("no active application avaiable for this emmployee id: %d in verified state ", empID)
	}

	records, err := tx.RecommendationsIPApplications.
		Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		All(ctx)

	// Retrieve all records for the employee ID from RecommendationsIPApplications
	fmt.Println("records ", records)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR006", false, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}

	currentTime := time.Now().Truncate(time.Second)
	stat := "inactive_" + time.Now().Format("20060102150405")

	// Update the retrieved record with the provided values

	updatedRecord, err := tx.Exam_Applications_IP.
		Query().
		Where(
			(exam_applications_ip.EmployeeIDEQ(empID)),
			exam_applications_ip.ExamYear(id1),
			exam_applications_ip.StatusEQ("active")).
		WithCirclePrefRef().
		WithIPApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR007", false, errors.New("no application exists ")
		} else {
			return nil, 500, " -STR008", false, err
		}
	}
	fmt.Println("updatedRecord", updatedRecord)
	if updatedRecord == nil {
		return nil, 422, " -STR007", false, errors.New("no application exists ")
	}

	if updatedRecord.HallName != "" {
		return nil, 422, " -STR007", false, errors.New("already Exam Hall name assigned hence Exam city change not allowed")
	}
	_, err = updatedRecord.
		Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR009", false, err
	}
	// Hall Ticket Generated Flag
	var flagNO bool
	if updatedRecord.GenerateHallTicketFlagByNO != nil {
		flagNO = *updatedRecord.GenerateHallTicketFlagByNO
	} else {
		// Set a default value or handle the nil case
		flagNO = false
	}

	var flag1 bool
	if updatedRecord.GenerateHallTicketFlag != nil {
		flag1 = *updatedRecord.GenerateHallTicketFlag
	} else {
		// Set a default value or handle the nil case
		flag1 = false
	}
	updatedAppln := tx.Exam_Applications_IP.
		Create().
		SetApplicationNumber(updatedRecord.ApplicationNumber).
		SetApplicationStatus(updatedRecord.ApplicationStatus).
		SetApplnSubmittedDate(updatedRecord.ApplnSubmittedDate).
		SetCAEmployeeDesignation(updatedRecord.CAEmployeeDesignation).
		SetCADate(updatedRecord.CADate).
		SetCAEmployeeID(updatedRecord.CAEmployeeID).
		SetCAGeneralRemarks(updatedRecord.CAGeneralRemarks).
		SetCAUserId(updatedRecord.CAUserId).
		SetCAUserName(updatedRecord.CAUserName).
		SetCadre(updatedRecord.Cadre).
		SetCandidateRemarks(updatedRecord.CandidateRemarks).
		SetCategoryDescription(updatedRecord.CategoryDescription).
		SetCenterFacilityId(applicationRecord.CenterFacilityId).
		SetCenterId(applicationRecord.CenterId).
		SetCentrePreference(applicationRecord.CentrePreference).
		SetClaimingQualifyingService(updatedRecord.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(updatedRecord.ControllingOfficeFacilityID).
		SetControllingOfficeName(updatedRecord.ControllingOfficeName).
		SetDCCS(updatedRecord.DCCS).
		SetDOB(updatedRecord.DOB).
		SetDeputationControllingOfficeID(updatedRecord.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(updatedRecord.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(updatedRecord.DeputationOfficeFacilityID).
		SetDeputationOfficeName(updatedRecord.DeputationOfficeName).
		SetDeputationOfficePincode(updatedRecord.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(updatedRecord.DeputationOfficeUniqueId).
		SetInDeputation(updatedRecord.InDeputation).
		SetDeputationType(updatedRecord.DeputationType).
		SetDesignationID(updatedRecord.DesignationID).
		SetDisabilityPercentage(updatedRecord.DisabilityPercentage).
		SetDisabilityTypeCode(updatedRecord.DisabilityTypeCode).
		SetDisabilityTypeDescription(updatedRecord.DisabilityTypeDescription).
		SetDisabilityTypeID(updatedRecord.DisabilityTypeID).
		SetEducationCode(updatedRecord.EducationCode).
		SetEducationDescription(updatedRecord.EducationDescription).
		SetEmailID(updatedRecord.EmailID).
		SetEmployeeID(updatedRecord.EmployeeID).
		SetEmployeeName(updatedRecord.EmployeeName).
		SetEntryPostCode(updatedRecord.EntryPostCode).
		SetEntryPostDescription(updatedRecord.EntryPostDescription).
		SetExamCode(updatedRecord.ExamCode).
		SetExamName(updatedRecord.ExamName).
		SetExamShortName(updatedRecord.ExamShortName).
		SetExamYear(updatedRecord.ExamYear).
		SetExamCityCenterCode(applicationRecord.CenterId).
		SetFacilityUniqueID(updatedRecord.FacilityUniqueID).
		SetFeederPostCode(updatedRecord.FeederPostCode).
		SetFeederPostDescription(updatedRecord.FeederPostDescription).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetGender(updatedRecord.Gender).
		SetLienControllingOfficeID(updatedRecord.LienControllingOfficeID).
		SetLienControllingOfficeName(updatedRecord.LienControllingOfficeName).
		SetMobileNumber(updatedRecord.MobileNumber).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetNodalOfficeName(updatedRecord.NodalOfficeName).
		SetPhoto(updatedRecord.Photo).
		SetPhotoPath(updatedRecord.PhotoPath).
		SetPresentDesignation(updatedRecord.PresentDesignation).
		SetPresentPostCode(updatedRecord.PresentPostCode).
		SetPresentPostDescription(updatedRecord.PresentPostDescription).
		SetReportingOfficeFacilityID(updatedRecord.ReportingOfficeFacilityID).
		SetReportingOfficeName(updatedRecord.ReportingOfficeName).
		SetServiceLength(updatedRecord.ServiceLength).
		SetSignature(updatedRecord.Signature).
		SetSignaturePath(updatedRecord.SignaturePath).
		SetTempHallTicket(updatedRecord.TempHallTicket).
		SetUserID(updatedRecord.UserID).
		SetWorkingOfficeCircleFacilityID(updatedRecord.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(updatedRecord.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(updatedRecord.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(updatedRecord.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(updatedRecord.WorkingOfficeFacilityID).
		SetWorkingOfficeName(updatedRecord.WorkingOfficeName).
		SetWorkingOfficePincode(updatedRecord.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(updatedRecord.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(updatedRecord.WorkingOfficeRegionName).
		SetHallTicketNumber(updatedRecord.HallTicketNumber).
		SetHallTicketGeneratedFlag(updatedRecord.HallTicketGeneratedFlag).
		SetHallTicketGeneratedDate(updatedRecord.HallTicketGeneratedDate).
		SetHallIdentificationNumber(updatedRecord.HallIdentificationNumber).
		SetGenerateHallTicketFlag(flag1).
		SetGenerateHallTicketFlagByNO(flagNO).
		SetCACorrected("NA ExamCity Change").
		SetCACorrectedDate(currentTime).
		SetNAEmployeeDesignation(applicationRecord.NA_EmployeeDesignation).
		SetNAEmployeeID(applicationRecord.NA_EmployeeID).
		SetNAUserName(applicationRecord.NA_UserName).
		SetNARemarks(applicationRecord.NA_Remarks).
		SetNADate(currentTime).
		SaveX(ctx)
	if err != nil {
		return nil, 500, " -STR012", false, err
	}
	fmt.Println(updatedAppln.ID)

	//new-------

	circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(updatedRecord.Edges.CirclePrefRef))
	for i, circlePrefRef := range updatedRecord.Edges.CirclePrefRef {
		if circlePrefRef.PlacePrefNo == 0 {
			return nil, 422, " -STR011", false, fmt.Errorf("circle preference value at index %d is nil", i)
		}

		circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
			Create().
			SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(applicationRecord.EmployeeID).
			SetPlacePrefValue(circlePrefRef.PlacePrefValue).
			SetUpdatedAt(currentTime).
			Save(ctx)

		if err != nil {
			return nil, 500, " -STR012", false, err
		}
		circlePrefRefs[i] = circlePrefRefEntity
	}
	fmt.Println("circlePrefRefs", circlePrefRefs)
	// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
	_, err = updatedAppln.
		Update().
		AddCirclePrefRef(circlePrefRefs...).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR013", false, err
	}

	recommendationsRef := make([]*ent.RecommendationsIPApplications, len(updatedRecord.Edges.IPApplicationsRef))
	for i, recommendation := range updatedRecord.Edges.IPApplicationsRef {
		if recommendation.VacancyYear == 0 {
			return nil, 400, " -STR015", false, fmt.Errorf("recommendations value at index %d is nil", i)
		}

		RecommendationsRefEntity, err := tx.RecommendationsIPApplications.
			Create().
			SetApplicationID(updatedAppln.ID).
			SetApplicationStatus("VerifiedRecommendationsByCA").
			SetCARecommendations(recommendation.CARecommendations).
			SetCARemarks(recommendation.CARemarks).
			SetCAUpdatedAt(recommendation.CAUpdatedAt).
			SetCAUserName(recommendation.CAUserName).
			SetEmployeeID(recommendation.EmployeeID).
			SetExamNameCode(recommendation.ExamNameCode).
			SetExamYear(recommendation.ExamYear).
			SetNORecommendations(recommendation.NORecommendations).
			SetNOUpdatedAt(recommendation.NOUpdatedAt).
			SetVacancyYear(recommendation.VacancyYear).
			Save(ctx)

		if err != nil {
			return nil, 500, " -STR016", false, err
		}
		recommendationsRef[i] = RecommendationsRefEntity
	}
	fmt.Println("recommendationsRef", recommendationsRef)
	updatedAppln.Update().
		ClearIPApplicationsRef().
		AddIPApplicationsRef(recommendationsRef...).
		Save(ctx)

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR019", false, err
	}
	return updatedAppln, 200, "", true, nil
}
func SubCheckIPExamcenterhall(client *ent.Client, examyear string, examcode int32, naFacilityId string, coFacilityID string, examHall string) ([]*ent.ExamCenterHall, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	// Check if empID is null

	// Retrieve all records for the employee ID
	records, err := client.ExamCenterHall.Query().
		Where(examcenterhall.AdminCircleOfficeIDEQ(naFacilityId),
			examcenterhall.ExamYearEQ(examyear),
			examcenterhall.ExamCodeEQ(examcode),
			examcenterhall.ConductedByFacilityIDEQ(coFacilityID),
			examcenterhall.HallNameEQ(examHall),
			examcenterhall.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no examcenterhall found for this city ID: %s", naFacilityId)
	}
	return records, 200, "", true, nil
}
func SubIPApplicationCAEdit(client *ent.Client, applicationEdit *ca_reg.ApplicationIpEditCA) (*ent.Exam_Applications_IP, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if !isNumeric(applicationEdit.TempHallTicket) {
		return nil, 400, " -STR100", false, fmt.Errorf("issue for employee %d with temp hall ticket number: %s", applicationEdit.EmployeeID, applicationEdit.TempHallTicket)
	}
	if len(applicationEdit.TempHallTicket) != 8 {
		return nil, 400, " -STR101", false, fmt.Errorf("issue for employee %d with temp hall ticket number length issue: %s", applicationEdit.EmployeeID, applicationEdit.TempHallTicket)
	}
	// Check if newAppln is not nil.
	if applicationEdit.EmployeeID == 0 {
		return nil, 400, " -STR102", false, errors.New("payload is empty")
	}

	if applicationEdit == nil {
		return nil, 400, " -STR001", false, errors.New("payload received in empty")
	}
	empID := applicationEdit.EmployeeID
	id1 := applicationEdit.ExamYear
	if empID == 0 {
		return nil, 422, " -STR002", false, errors.New("employee id should not be empty")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR017", false, err
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
	oldAppln, err := tx.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(empID),
			exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamYearEQ(id1),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR003", false, errors.New("no active application found for this candidate with this status")
		} else {
			return nil, 500, " -STR004", false, err
		}
	}
	if oldAppln == nil {
		return nil, 422, " -STR005", false, errors.New("no active application found for this candidate")
	}
	currentTime := time.Now().Truncate(time.Second)
	stat := "inactive_" + time.Now().Format("20060102150405")

	if applicationEdit.Edges.CircleData == nil || len(applicationEdit.Edges.CircleData) == 0 {
		return nil, 400, " -STR008", false, fmt.Errorf("circle preference values are missing! Please provide Circle preferences")
	}

	if len(applicationEdit.Edges.CircleData) != 23 {
		return nil, 400, " -STR009", false, fmt.Errorf("invalid number of Circle preferences. Must provide preferences for all 23 circles")
	}

	// Update old application status to inactive.
	_, err = oldAppln.
		Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR011", false, err
	}

	// Hall Ticket Generated Flag

	// Create new application record.
	updatedAppln, err := tx.Exam_Applications_IP.
		Create().
		SetApplicationNumber(oldAppln.ApplicationNumber).
		SetApplicationStatus("ResubmitCAVerificationPending").
		SetApplnSubmittedDate(oldAppln.ApplnSubmittedDate).
		SetCAEmployeeDesignation(oldAppln.CAEmployeeDesignation).
		SetCADate(currentTime).
		SetCAEmployeeID(applicationEdit.CA_EmployeeID).
		SetCAGeneralRemarks(applicationEdit.CA_GeneralRemarks).
		SetCAUserId(applicationEdit.CA_UserID).
		SetCAUserName(applicationEdit.CA_UserName).
		SetCadre(applicationEdit.Cadre).
		SetCandidateRemarks(applicationEdit.CandidateRemarks).
		SetCategoryCode(applicationEdit.CategoryCode).
		SetCategoryDescription(applicationEdit.CategoryDescription).
		SetCenterFacilityId(applicationEdit.CenterFacilityId).
		SetCenterId(applicationEdit.CenterId).
		SetCentrePreference(applicationEdit.CentrePreference).
		SetClaimingQualifyingService(applicationEdit.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(applicationEdit.ControllingOfficeFacilityID).
		SetControllingOfficeName(applicationEdit.ControllingOfficeName).
		SetDCCS(applicationEdit.DCCS).
		SetDOB(applicationEdit.DOB).
		SetDeputationControllingOfficeID(applicationEdit.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(applicationEdit.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(applicationEdit.DeputationOfficeFacilityID).
		SetDeputationOfficeName(applicationEdit.DeputationOfficeName).
		SetDeputationOfficePincode(applicationEdit.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(applicationEdit.DeputationOfficeUniqueId).
		SetInDeputation(applicationEdit.InDeputation).
		SetDeputationType(applicationEdit.DeputationType).
		SetDesignationID(applicationEdit.DesignationID).
		SetDisabilityPercentage(applicationEdit.DisabilityPercentage).
		SetDisabilityTypeCode(applicationEdit.DisabilityTypeCode).
		SetDisabilityTypeDescription(applicationEdit.DisabilityTypeDescription).
		SetDisabilityTypeID(applicationEdit.DisabilityTypeID).
		SetEducationCode(applicationEdit.EducationCode).
		SetEducationDescription(applicationEdit.EducationDescription).
		SetEmailID(applicationEdit.EmailID).
		SetEmployeeID(applicationEdit.EmployeeID).
		SetEmployeeName(applicationEdit.EmployeeName).
		SetEntryPostCode(applicationEdit.EntryPostCode).
		SetEntryPostDescription(applicationEdit.EntryPostDescription).
		SetExamCode(applicationEdit.ExamCode).
		SetExamName(applicationEdit.ExamName).
		SetExamShortName(applicationEdit.ExamShortName).
		SetExamYear(applicationEdit.ExamYear).
		SetExamCityCenterCode(applicationEdit.CenterId).
		SetFacilityUniqueID(applicationEdit.FacilityUniqueID).
		SetFeederPostCode(applicationEdit.FeederPostCode).
		SetFeederPostDescription(applicationEdit.FeederPostDescription).
		SetFeederPostJoiningDate(applicationEdit.FeederPostJoiningDate).
		SetGender(applicationEdit.Gender).
		SetLienControllingOfficeID(applicationEdit.LienControllingOfficeID).
		SetLienControllingOfficeName(applicationEdit.LienControllingOfficeName).
		SetMobileNumber(applicationEdit.MobileNumber).
		SetNodalOfficeFacilityID(applicationEdit.NodalOfficeFacilityID).
		SetNodalOfficeName(applicationEdit.NodalOfficeName).
		SetPhoto(applicationEdit.Photo).
		SetPhotoPath(applicationEdit.PhotoPath).
		SetPresentDesignation(applicationEdit.PresentDesignation).
		SetPresentPostCode(applicationEdit.PresentPostCode).
		SetPresentPostDescription(applicationEdit.PresentPostDescription).
		SetReportingOfficeFacilityID(applicationEdit.ReportingOfficeFacilityID).
		SetReportingOfficeName(applicationEdit.ReportingOfficeName).
		SetServiceLength(*applicationEdit.ServiceLength).
		SetSignature(applicationEdit.Signature).
		SetSignaturePath(applicationEdit.SignaturePath).
		SetTempHallTicket(applicationEdit.TempHallTicket).
		SetUserID(applicationEdit.UserID).
		SetWorkingOfficeCircleFacilityID(applicationEdit.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(applicationEdit.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(applicationEdit.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(applicationEdit.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(applicationEdit.WorkingOfficeFacilityID).
		SetWorkingOfficeName(applicationEdit.WorkingOfficeName).
		SetWorkingOfficePincode(applicationEdit.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(applicationEdit.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(applicationEdit.WorkingOfficeRegionName).
		SetCACorrected("CA Corrected").
		SetCACorrectedDate(currentTime).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR012", false, err
	}

	// Save the PlaceOfPreferenceIP records.
	circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(applicationEdit.Edges.CircleData))
	for i, circlePrefRef := range applicationEdit.Edges.CircleData {
		if circlePrefRef.PlacePrefNo == 0 {
			return nil, 400, " -STR013", false, fmt.Errorf("circle preference value at index %d is nil", i)
		}

		circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
			Create().
			SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(applicationEdit.EmployeeID).
			SetPlacePrefValue(circlePrefRef.PlacePrefValue).
			SetUpdatedAt(currentTime).
			Save(ctx)

		if err != nil {
			return nil, 500, " -STR014", false, err
		}
		circlePrefRefs[i] = circlePrefRefEntity
	}

	// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
	_, err = updatedAppln.
		Update().
		AddCirclePrefRef(circlePrefRefs...).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR015", false, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR016", false, err
	}
	return updatedAppln, 200, "", true, nil

}

func SubResetIPApplicationNA(client *ent.Client, applicationRecord *ca_reg.NAVerifyApplicationIp) (*ent.Exam_Applications_IP, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if applicationRecord == nil {
		return nil, 400, " -STR001", false, errors.New("payload received in empty")
	}
	empID := applicationRecord.EmployeeID
	id1 := applicationRecord.ExamYear
	if empID == 0 {
		return nil, 422, " -STR002", false, errors.New("employee id should not be empty")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR017", false, err
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
	exists, err := tx.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(empID),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamYearEQ(id1),
		).
		Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR003", false, fmt.Errorf("no active application avaiable for this emmployee id: %d in verified state ", empID)
		} else {
			return nil, 500, " -STR004", false, err
		}
	}
	if !exists {
		return nil, 422, " -STR003", false, fmt.Errorf("no active application avaiable for this emmployee id: %d in verified state ", empID)
	}

	records, err := tx.RecommendationsIPApplications.
		Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		All(ctx)

	// Retrieve all records for the employee ID from RecommendationsIPApplications

	if err != nil {
		return nil, 500, " -STR005", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR006", false, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}

	currentTime := time.Now().Truncate(time.Second)
	stat := "inactive_" + time.Now().Format("20060102150405")

	// Update the retrieved record with the provided values

	updatedRecord, err := tx.Exam_Applications_IP.
		Query().
		Where(
			(exam_applications_ip.EmployeeIDEQ(empID)),
			exam_applications_ip.ExamYear(id1),
			exam_applications_ip.StatusEQ("active")).
		WithCirclePrefRef().
		WithIPApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR007", false, errors.New("no application exists ")
		} else {
			return nil, 500, " -STR008", false, err
		}
	}

	if updatedRecord == nil {
		return nil, 422, " -STR007", false, errors.New("no application exists ")
	}

	if updatedRecord.HallTicketGeneratedFlag {
		return nil, 422, " -STR007", false, errors.New("already Hall ticket generated hence reset not allowed")
	}
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

	var recomondPref []*ent.RecommendationsIPApplications
	for _, edge := range updatedRecord.Edges.IPApplicationsRef {
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
	updatedRecord.Edges.IPApplicationsRef = recomondPref
	updatedRecord.UpdatedAt = updatedRecord.UpdatedAt.Truncate(24 * time.Hour)

	_, err = updatedRecord.
		Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR009", false, err
	}
	// Hall Ticket Generated Flag

	updatedAppln := tx.Exam_Applications_IP.
		Create().
		SetApplicationNumber(updatedRecord.ApplicationNumber).
		SetApplicationStatus("ResubmitCAVerificationPending").
		SetApplnSubmittedDate(updatedRecord.ApplnSubmittedDate).
		SetCAEmployeeDesignation(updatedRecord.CAEmployeeDesignation).
		SetCADate(updatedRecord.CADate).
		SetCAEmployeeID(updatedRecord.CAEmployeeID).
		SetCAGeneralRemarks(updatedRecord.CAGeneralRemarks).
		SetCAUserId(updatedRecord.CAUserId).
		SetCAUserName(updatedRecord.CAUserName).
		SetCadre(updatedRecord.Cadre).
		SetCandidateRemarks(updatedRecord.CandidateRemarks).
		SetCategoryCode(updatedRecord.CategoryCode).
		SetCategoryDescription(updatedRecord.CategoryDescription).
		SetCenterFacilityId(updatedRecord.CenterFacilityId).
		SetCenterId(updatedRecord.CenterId).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetCentrePreference(updatedRecord.CentrePreference).
		SetClaimingQualifyingService(updatedRecord.ClaimingQualifyingService).
		SetControllingOfficeFacilityID(updatedRecord.ControllingOfficeFacilityID).
		SetControllingOfficeName(updatedRecord.ControllingOfficeName).
		SetDCCS(updatedRecord.DCCS).
		SetDOB(updatedRecord.DOB).
		SetDeputationControllingOfficeID(updatedRecord.DeputationControllingOfficeID).
		SetDeputationControllingOfficeName(updatedRecord.DeputationControllingOfficeName).
		SetDeputationOfficeFacilityID(updatedRecord.DeputationOfficeFacilityID).
		SetDeputationOfficeName(updatedRecord.DeputationOfficeName).
		SetDeputationOfficePincode(updatedRecord.DeputationOfficePincode).
		SetDeputationOfficeUniqueId(updatedRecord.DeputationOfficeUniqueId).
		SetInDeputation(updatedRecord.InDeputation).
		SetDeputationType(updatedRecord.DeputationType).
		SetDesignationID(updatedRecord.DesignationID).
		SetDisabilityPercentage(updatedRecord.DisabilityPercentage).
		SetDisabilityTypeCode(updatedRecord.DisabilityTypeCode).
		SetDisabilityTypeDescription(updatedRecord.DisabilityTypeDescription).
		SetDisabilityTypeID(updatedRecord.DisabilityTypeID).
		SetEducationCode(updatedRecord.EducationCode).
		SetEducationDescription(updatedRecord.EducationDescription).
		SetEmailID(updatedRecord.EmailID).
		SetEmployeeID(updatedRecord.EmployeeID).
		SetEmployeeName(updatedRecord.EmployeeName).
		SetEntryPostCode(updatedRecord.EntryPostCode).
		SetEntryPostDescription(updatedRecord.EntryPostDescription).
		SetExamCode(updatedRecord.ExamCode).
		SetExamName(updatedRecord.ExamName).
		SetExamShortName(updatedRecord.ExamShortName).
		SetExamYear(updatedRecord.ExamYear).
		SetExamCityCenterCode(updatedRecord.CenterId).
		SetFacilityUniqueID(updatedRecord.FacilityUniqueID).
		SetFeederPostCode(updatedRecord.FeederPostCode).
		SetFeederPostDescription(updatedRecord.FeederPostDescription).
		SetFeederPostJoiningDate(updatedRecord.FeederPostJoiningDate).
		SetGender(updatedRecord.Gender).
		SetLienControllingOfficeID(updatedRecord.LienControllingOfficeID).
		SetLienControllingOfficeName(updatedRecord.LienControllingOfficeName).
		SetMobileNumber(updatedRecord.MobileNumber).
		SetNodalOfficeFacilityID(updatedRecord.NodalOfficeFacilityID).
		SetNodalOfficeName(updatedRecord.NodalOfficeName).
		SetPhoto(updatedRecord.Photo).
		SetPhotoPath(updatedRecord.PhotoPath).
		SetPresentDesignation(updatedRecord.PresentDesignation).
		SetPresentPostCode(updatedRecord.PresentPostCode).
		SetPresentPostDescription(updatedRecord.PresentPostDescription).
		SetReportingOfficeFacilityID(updatedRecord.ReportingOfficeFacilityID).
		SetReportingOfficeName(updatedRecord.ReportingOfficeName).
		SetServiceLength(updatedRecord.ServiceLength).
		SetSignature(updatedRecord.Signature).
		SetSignaturePath(updatedRecord.SignaturePath).
		SetTempHallTicket(updatedRecord.TempHallTicket).
		SetUserID(updatedRecord.UserID).
		SetWorkingOfficeCircleFacilityID(updatedRecord.WorkingOfficeCircleFacilityID).
		SetWorkingOfficeCircleName(updatedRecord.WorkingOfficeCircleName).
		SetWorkingOfficeDivisionFacilityID(updatedRecord.WorkingOfficeDivisionFacilityID).
		SetWorkingOfficeDivisionName(updatedRecord.WorkingOfficeDivisionName).
		SetWorkingOfficeFacilityID(updatedRecord.WorkingOfficeFacilityID).
		SetWorkingOfficeName(updatedRecord.WorkingOfficeName).
		SetWorkingOfficePincode(updatedRecord.WorkingOfficePincode).
		SetWorkingOfficeRegionFacilityID(updatedRecord.WorkingOfficeRegionFacilityID).
		SetWorkingOfficeRegionName(updatedRecord.WorkingOfficeRegionName).
		SaveX(ctx)
	fmt.Println(updatedAppln.ID)
	//new-------
	circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(applicationRecord.Edges.CircleData))
	for i, circlePrefRef := range applicationRecord.Edges.CircleData {
		if circlePrefRef.PlacePrefNo == 0 {
			return nil, 422, " -STR011", false, fmt.Errorf("circle preference value at index %d is nil", i)
		}

		circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
			Create().
			SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(applicationRecord.EmployeeID).
			SetPlacePrefValue(circlePrefRef.PlacePrefValue).
			SetUpdatedAt(currentTime).
			Save(ctx)

		if err != nil {
			return nil, 500, " -STR012", false, err
		}
		circlePrefRefs[i] = circlePrefRefEntity
	}

	// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
	_, err = updatedAppln.
		Update().
		AddCirclePrefRef(circlePrefRefs...).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR013", false, err
	}

	// For Resubmission

	// Insert into recommendations.
	// Save the Recommendation records.

	/* 	recommendationsRef := make([]*ent.RecommendationsIPApplications, len(applicationRecord.Edges.ApplicationDataN))
	   	for i, recommendation := range applicationRecord.Edges.ApplicationDataN {
	   		if recommendation.VacancyYear == 0 {
	   			return nil, 422, " -STR014", false, fmt.Errorf(" recommendations value at index %d is nil", i)
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
	   			return nil, 500, " -STR018", false, err
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
	   			SetCAUserName(prevRecommendation.CAUserName). // Use newAppln.CAUserName instead of updatedAppln.CAUserName
	   			SetCARemarks(prevRecommendation.CARemarks).   // Use newAppln.CARemarks instead of updatedAppln.CARemarks
	   			SetCAUpdatedAt(prevRecommendation.CAUpdatedAt).
	   			SetNOUpdatedAt(currentTime).
	   			SetNOUserName(applicationRecord.NA_UserName). // Use newAppln.CAUserName instead of updatedAppln.CAUserName
	   			SetNORemarks(recommendation.NO_Remarks).      // Use newAppln.CARemarks instead of updatedAppln.CARemarks

	   			SetApplicationStatus("VerifiedRecommendationsByNA").
	   			Save(ctx)
	   		if err != nil {
	   			return nil, 500, " -STR015", false, err
	   		}

	   		recommendationsRef[i] = RecommendationsRefEntity
	   	}

	updatedAppln.Update().
		//ClearIPApplicationsRef().
		AddIPApplicationsRef(recommendationsRef...).
		Save(ctx)*/
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR019", false, err
	}
	return updatedAppln, 200, "", true, nil
}

func ResubmitApplicationn(client *ent.Client, newAppln *ca_reg.ResubmitApplicationIp) (*ent.Exam_Applications_IP, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if !isNumeric(newAppln.TempHallTicket) {
		return nil, 400, " -STR100", false, fmt.Errorf("issue for employee %d with temp hall ticket number: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
	}
	if len(newAppln.TempHallTicket) != 8 {
		return nil, 400, " -STR101", false, fmt.Errorf("issue for employee %d with temp hall ticket number length issue: %s", newAppln.EmployeeID, newAppln.TempHallTicket)
	}
	// Check if newAppln is not nil.
	if newAppln.EmployeeID == 0 {
		return nil, 400, " -STR001", false, errors.New("payload is empty")
	}
	//transaction implementation--------------
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
			return nil, 422, " -STR003", false, errors.New("no active application found for this candidate")
		} else {
			return nil, 500, " -STR004", false, err
		}
	}
	existing, err := tx.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(newAppln.EmployeeID),
			exam_applications_ip.ApplicationStatusIn("ResubmitCAVerificationPending"),
			exam_applications_ip.ExamYearEQ(newAppln.ExamYear),
			exam_applications_ip.StatusEQ("active"),
		).
		Exist(ctx)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	if existing {
		return nil, 422, " -STR006", false, errors.New("already application resubmitted for this candidate")
	}
	// Insert a new record with the specified conditions.

	//var currentTime = time.Now().Truncate(time.Second)
	var currentTime = time.Now().UTC().Truncate(time.Second) // Ensure UTC for consistent comparison
	currentTime = currentTime.Add(5*time.Hour + 30*time.Minute)
	correctionLastDate := newAppln.ApplicationCorrectionLastDate.UTC().Truncate(time.Second)

	fmt.Print("current time: ", currentTime, " correction last date: ", correctionLastDate, "date from payload", newAppln.ApplicationCorrectionLastDate)
	if currentTime.After(correctionLastDate) {
		return nil, 422, " -STR106", false, fmt.Errorf("application correction deadline has passed as current time is %v", currentTime)
	}
	stat := "inactive_" + time.Now().Format("20060102150405")

	if oldAppln == nil {
		return nil, 422, " -STR005", false, errors.New("no active application found for this candidate")
	} else {
		// Update the existing record.
		if oldAppln.ApplicationStatus == "VerifiedByNA" || oldAppln.ApplicationStatus == "VerifiedByCA" {
			return nil, 422, " -STR006", false, errors.New("the Application was already verified By Nodal Authority/ Controlling Authority")
		}

		if oldAppln.ApplicationStatus != "PendingWithCandidate" {
			return nil, 422, " -STR007", false, errors.New("this application was not in pending with candidate status")
		} else {
			if newAppln.Edges.CircleData == nil || len(newAppln.Edges.CircleData) == 0 {
				return nil, 400, " -STR008", false, fmt.Errorf("circle preference values are missing! Please provide Circle preferences")
			}

			if len(newAppln.Edges.CircleData) != 23 {
				return nil, 400, " -STR009", false, fmt.Errorf("invalid number of Circle preferences. Must provide preferences for all 23 circles")
			}

			applicationNumber1, err := util.GenerateApplicationNumber(client, newAppln.EmployeeID, newAppln.ExamYear, "IP")
			//applicationNumber1, err := generateApplicationNumber(client, newAppln.EmployeeID)

			if err != nil {
				return nil, 422, " -STR010", false, fmt.Errorf("failed to generate application number: %v", err)
			}

			// Update old application status to inactive.
			_, err = oldAppln.
				Update().
				SetStatus(stat).
				Save(ctx)
			if err != nil {
				return nil, 500, " -STR011", false, err
			}

			// Create new application record.
			updatedAppln, err := tx.Exam_Applications_IP.
				Create().
				SetApplicationNumber(applicationNumber1).
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
				return nil, 500, " -STR012", false, err
			}

			fmt.Println("application id", updatedAppln.ID)
			// Save the PlaceOfPreferenceIP records.
			circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(newAppln.Edges.CircleData))
			for i, circlePrefRef := range newAppln.Edges.CircleData {
				if circlePrefRef.PlacePrefNo == 0 {
					return nil, 400, " -STR013", false, fmt.Errorf("circle preference value at index %d is nil", i)
				}

				circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
					Create().
					SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
					SetApplicationID(updatedAppln.ID).
					SetEmployeeID(newAppln.EmployeeID).
					SetPlacePrefValue(circlePrefRef.PlacePrefValue).
					SetUpdatedAt(currentTime).
					Save(ctx)

				if err != nil {
					return nil, 500, " -STR014", false, err
				}
				circlePrefRefs[i] = circlePrefRefEntity
			}

			// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
			_, err = updatedAppln.
				Update().
				AddCirclePrefRef(circlePrefRefs...).
				Save(ctx)
			if err != nil {
				return nil, 500, " -STR015", false, err
			}
			if err = tx.Commit(); err != nil {
				tx.Rollback()
				return nil, 500, " -STR016", false, err
			}
			return updatedAppln, 200, "", true, nil
		}
	}
}
func ResubmitApplication(client *ent.Client, newAppln *ca_reg.ResubmitApplicationIp) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Validate input parameters
	if err := validateApplicationInputs(newAppln); err != nil {
		return nil, 400, " -STR001", false, err
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	defer func() {
		handleTransaction(tx, &err)
	}()
	// Fetch old application if it exists
	oldAppln, status, stgError, err := fetchOldApplication(ctx, tx, newAppln)
	if status == 500 {
		return nil, 500 + status, " -STR004 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR005 " + stgError, false, fmt.Errorf("no active application found for this candidate")

	}

	if err != nil {
		return nil, 500 + status, " -STR005 " + stgError, false, err
	}

	// Check if a resubmission already exists
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

func UpdateNodalRecommendationsByEmpIDold(client *ent.Client, applicationRecord *ca_reg.NAVerifyApplicationIp) (*ent.Exam_Applications_IP, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if applicationRecord == nil {
		return nil, 400, " -STR001", false, errors.New("payload received in empty")
	}
	empID := applicationRecord.EmployeeID
	id1 := applicationRecord.ExamYear
	if empID == 0 {
		return nil, 422, " -STR002", false, errors.New("employee id should not be empty")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR017", false, err
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
	// Check if empID exists in exam_applications_ip and the status is "VerifiedByCA" or "VerifiedByNA"
	exists, err := tx.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(empID),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamYearEQ(id1),
		).
		Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR003", false, fmt.Errorf("no active application avaiable for this emmployee id: %d in verified state ", empID)
		} else {
			return nil, 500, " -STR004", false, err
		}
	}
	if !exists {
		return nil, 422, " -STR003", false, fmt.Errorf("no active application avaiable for this emmployee id: %d in verified state ", empID)
	}

	records, err := tx.RecommendationsIPApplications.
		Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		All(ctx)

	// Retrieve all records for the employee ID from RecommendationsIPApplications

	if err != nil {
		return nil, 500, " -STR005", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR006", false, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}

	currentTime := time.Now().Truncate(time.Second)
	stat := "inactive_" + time.Now().Format("20060102150405")

	// Update the retrieved record with the provided values

	updatedRecord, err := tx.Exam_Applications_IP.
		Query().
		Where(
			(exam_applications_ip.EmployeeIDEQ(empID)),
			exam_applications_ip.ExamYear(id1),
			exam_applications_ip.StatusEQ("active")).
		WithCirclePrefRef().
		WithIPApplicationsRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR007", false, errors.New("no application exists ")
		} else {
			return nil, 500, " -STR008", false, err
		}
	}

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

	var recomondPref []*ent.RecommendationsIPApplications
	for _, edge := range updatedRecord.Edges.IPApplicationsRef {
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
	updatedRecord.Edges.IPApplicationsRef = recomondPref
	updatedRecord.UpdatedAt = updatedRecord.UpdatedAt.Truncate(24 * time.Hour)

	_, err = updatedRecord.
		Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR009", false, err
	}
	// Hall Ticket Generated Flag
	var hallticketgeneratedflag bool

	if applicationRecord.GenerateHallTicketFlag {
		if updatedRecord.HallTicketNumber != "" {
			hallticketgeneratedflag = true
		} else {
			hallticketgeneratedflag = false
		}
	} else {
		hallticketgeneratedflag = false
	}
	updatedAppln := tx.Exam_Applications_IP.
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
		//SetGenerateHallTicketFlagByNO(applicationRecord.GenerateHallTicketFlagByNO).
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
		//SetNAUserName(applicationRecord.NAUserName).
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
		//SetGenerateHallTicketFlag(*updatedRecord.GenerateHallTicketFlag).
		SaveX(ctx)

	//new-------
	circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(applicationRecord.Edges.CircleData))
	for i, circlePrefRef := range applicationRecord.Edges.CircleData {
		if circlePrefRef.PlacePrefNo == 0 {
			return nil, 422, " -STR011", false, fmt.Errorf("circle preference value at index %d is nil", i)
		}

		circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
			Create().
			SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(applicationRecord.EmployeeID).
			SetPlacePrefValue(circlePrefRef.PlacePrefValue).
			SetUpdatedAt(currentTime).
			Save(ctx)

		if err != nil {
			return nil, 500, " -STR012", false, err
		}
		circlePrefRefs[i] = circlePrefRefEntity
	}

	// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
	_, err = updatedAppln.
		Update().
		AddCirclePrefRef(circlePrefRefs...).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR013", false, err
	}

	// For Resubmission

	// Insert into recommendations.
	// Save the Recommendation records.

	recommendationsRef := make([]*ent.RecommendationsIPApplications, len(applicationRecord.Edges.ApplicationDataN))
	for i, recommendation := range applicationRecord.Edges.ApplicationDataN {
		if recommendation.VacancyYear == 0 {
			return nil, 422, " -STR014", false, fmt.Errorf(" recommendations value at index %d is nil", i)
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
			return nil, 500, " -STR018", false, err
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
			SetCAUserName(prevRecommendation.CAUserName). // Use newAppln.CAUserName instead of updatedAppln.CAUserName
			SetCARemarks(prevRecommendation.CARemarks).   // Use newAppln.CARemarks instead of updatedAppln.CARemarks
			SetCAUpdatedAt(prevRecommendation.CAUpdatedAt).
			SetNOUpdatedAt(currentTime).
			SetNOUserName(applicationRecord.NA_UserName). // Use newAppln.CAUserName instead of updatedAppln.CAUserName
			SetNORemarks(recommendation.NO_Remarks).      // Use newAppln.CARemarks instead of updatedAppln.CARemarks

			SetApplicationStatus("VerifiedRecommendationsByNA").
			Save(ctx)
		if err != nil {
			return nil, 500, " -STR015", false, err
		}

		recommendationsRef[i] = RecommendationsRefEntity
	}

	updatedAppln.Update().
		//ClearIPApplicationsRef().
		AddIPApplicationsRef(recommendationsRef...).
		Save(ctx)
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR019", false, err
	}
	return updatedAppln, 200, "", true, nil

}
func UpdateNodalRecommendationsByEmpID(client *ent.Client, applicationRecord *ca_reg.NAVerifyApplicationIp) (*ca_reg.ApplicationsResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if err := validateInput(applicationRecord); err != nil {
		return nil, 422, " -STR001", false, errors.New("employee id should not be empty")
	}
	empID := applicationRecord.EmployeeID
	id1 := applicationRecord.ExamYear

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	defer handleTransaction(tx, &err)

	// Check if empID exists in exam_applications_ip and the status is "VerifiedByCA" or "VerifiedByNA"
	// Check if application exists
	exists, status, stgError, err := checkApplicationExists(tx, ctx, applicationRecord)

	if status == 500 {
		return nil, 500, " -STR003 " + stgError, false, err
	}

	if status == 422 {
		return nil, 422, " -STR004 " + stgError, false, fmt.Errorf("no active application found for this candidate")
	}

	if err != nil {
		return nil, 500, " -STR005 " + stgError, false, err
	}

	if !exists {
		return nil, 422, " -STR004 no active application found", false, fmt.Errorf("no active application found for this candidate")
	}

	records, err := getRecommendationsByEmpID(ctx, tx, empID)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR006", false, fmt.Errorf("no records found in recommendations for employee with ID %d", empID)
	}

	currentTime := time.Now().Truncate(time.Second)
	stat := "inactive_" + time.Now().Format("20060102150405")

	// Update the retrieved record with the provided values
	updatedRecord, status, stgError, err := getActiveExamApplicationIP(ctx, tx, empID, id1)
	if status == 500 {
		return nil, 500 + status, " -STR003 " + stgError, false, err
	}
	if status == 422 {
		return nil, 422 + status, " -STR004 " + stgError, false, fmt.Errorf("no active application exists")

	}
	if err != nil {
		return nil, 500 + status, " -STR005 " + stgError, false, err
	}
	// Extract only the desired fields from the CirclePrefRef edge
	updatedRecord, err = updateRecordEdges(updatedRecord)
	if err != nil {
		return nil, 500, " -STR009", false, err
	}
	_, err = updatedRecord.
		Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR009", false, err
	}
	// Hall Ticket Generated Flag
	hallticketgeneratedflag := checkHallTicketGenerated(applicationRecord, updatedRecord)

	updatedAppln, err := createUpdatedAppln(tx, applicationRecord, updatedRecord, hallticketgeneratedflag, ctx)
	if err != nil {
		return nil, 500, " -STR010", false, err
	}

	//new-------
	circlePrefRefs := make([]*ent.PlaceOfPreferenceIP, len(applicationRecord.Edges.CircleData))
	for i, circlePrefRef := range applicationRecord.Edges.CircleData {
		if circlePrefRef.PlacePrefNo == 0 {
			return nil, 422, " -STR011", false, fmt.Errorf("circle preference value at index %d is nil", i)
		}

		circlePrefRefEntity, err := tx.PlaceOfPreferenceIP.
			Create().
			SetPlacePrefNo(int32(circlePrefRef.PlacePrefNo)).
			SetApplicationID(updatedAppln.ID).
			SetEmployeeID(applicationRecord.EmployeeID).
			SetPlacePrefValue(circlePrefRef.PlacePrefValue).
			SetUpdatedAt(currentTime).
			Save(ctx)

		if err != nil {
			return nil, 500, " -STR012", false, err
		}
		circlePrefRefs[i] = circlePrefRefEntity
	}

	// Add the PlaceOfPreferenceIP references to the Exam_Applications_IP entity.
	_, err = updatedAppln.
		Update().
		AddCirclePrefRef(circlePrefRefs...).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR013", false, err
	}

	// Save the Recommendation records.

	recommendationsRef, err := createRecommendationsRef(ctx, tx, applicationRecord, updatedAppln)
	if err != nil {
		return nil, 500, " -STR015", false, err
	}

	updatedAppln.Update().
		//ClearIPApplicationsRef().
		AddIPApplicationsRef(recommendationsRef...).
		Save(ctx)
	appResponse, err := MapExamApplicationsToResponse(updatedAppln)
	if err != nil {
		return nil, 500, " -SUB017", false, err
	}
	return appResponse, 200, "", true, nil

}

// Get all applications based on facility id  and year
func SubGetPASAIPApplicationsFacilityIDYear(ctx context.Context, client *ent.Client, facilityID string, year string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {
	// Array of exams

	if facilityID == "" || year == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID and Examyear cannot be blank/null")
	}
	records, err := client.Exam_Applications_IP.Query().
		Where(exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
			exam_applications_ip.ExamYearEQ(year),
			exam_applications_ip.StatusEQ("active"),
		).
		Order(ent.Desc(exam_applications_ip.FieldID)). // Order by descending updated_at timestamp
		WithCirclePrefRef().                           // Add the Where clause with multiple statuses using Or
		All(ctx)
	if err != nil {
		return nil, 400, " -STR002", false, fmt.Errorf(" failed querying IP exams Applications: %w", err)
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf(" nil Applications for the Year %s and facility ID  %s", year, facilityID)
	}

	return records, 200, "", true, nil
}

// Get All CA Pending records ...
func QueryIPApplicationsByCAVerificationsPending(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {
	// Array of exams

	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID and Examyear cannot be null")
	}
	log.Println("Input Facility ID:", facilityID, "Examyear:", id1) // Log the facility ID and Examyear

	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.And(

				exam_applications_ip.Or(
					exam_applications_ip.ApplicationStatusIn("CAVerificationPending", "ResubmitCAVerificationPending"),
				),
				exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
				exam_applications_ip.ExamYearEQ(id1),
				exam_applications_ip.StatusEQ("active"),
			),
		).
		Order(ent.Desc(exam_applications_ip.FieldID)). // Order by descending updated_at timestamp
		//Limit(1).                                      // Limit to 1 record (the latest)
		WithCirclePrefRef(). // Add the Where clause with multiple statuses using Or
		All(ctx)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR006", false, fmt.Errorf(" no Applications is pending for CA pending verification for the Office ID %s", facilityID)
	}

	return records, 200, "", true, nil
}

func SubGetAllPendingApplicationsOnDeputation(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {
	// Array of exams

	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID and Examyear cannot be null")
	}

	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.LienControllingOfficeIDEQ(facilityID),
			exam_applications_ip.DeputationControllingOfficeIDNEQ(""),
			exam_applications_ip.ExamYearEQ(id1),
			exam_applications_ip.StatusEQ("active"),
		).
		Order(ent.Desc(exam_applications_ip.FieldID)). // Order by descending updated_at timestamp
		//Limit(1).                                      // Limit to 1 record (the latest)
		WithCirclePrefRef(). // Add the Where clause with multiple statuses using Or
		All(ctx)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	if len(records) == 0 {
		return nil, 422, " -STR006", false, fmt.Errorf(" no application found deputed officials for the Office ID %s", facilityID)
	}

	return records, 200, "", true, nil
}

// Get All CA verified records
func QueryIPApplicationsByCAVerified(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {

	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.And(
				exam_applications_ip.ApplicationStatusEQ("VerifiedByCA"),
				exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(id1),
			),
		).
		//WithIPApplicationsRef().
		WithCirclePrefRef().
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
func QueryIPApplicationsByCAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_IP, int32, string, bool, error) {

	record, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ApplicationStatusEQ("VerifiedByCA"), // Check for "CAVerified" status
			exam_applications_ip.EmployeeIDEQ(employeeID),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active"),
		).
		WithIPApplicationsRef().
		WithCirclePrefRef().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf("no application found for this employee ID: %d with 'CAVerified' status", employeeID)
		}
		return nil, 500, " -STR002", false, err
	}
	return record, 200, "", true, nil
}

func QueryIPApplicationsByCAPendingByEmpID(ctx context.Context, client *ent.Client, empID int64, examYear string) (*ent.Exam_Applications_IP, int32, string, bool, error) {

	// Retrieve the latest record based on UpdatedAt timestamp
	record, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(empID),
			exam_applications_ip.Or(
				exam_applications_ip.ApplicationStatusEQ("CAVerificationPending"),
				exam_applications_ip.ApplicationStatusEQ("ResubmitCAVerificationPending"),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
			),
		).
		Order(ent.Desc("updated_at")). // Order by UpdatedAt in descending order
		WithIPApplicationsRef().
		WithCirclePrefRef().
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
func GetOldCAApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_IP, int32, string, bool, error) {

	application, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(employeeID),
			exam_applications_ip.ApplicationStatusEQ("PendingWithCandidate"),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active"),
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
func GetRecommendationsByEmpID(client *ent.Client, empID int64) ([]*ent.RecommendationsIPApplications, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	// Check if empID is null
	if empID == 0 {
		return nil, 422, " -STR001", false, fmt.Errorf("no employee ID provided to process")
	}

	// Retrieve all records for the employee ID
	records, err := client.RecommendationsIPApplications.Query().
		Where(recommendationsipapplications.EmployeeIDEQ(empID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no recommendations found for this employee ID: %d", empID)
	}

	return records, 200, "", true, nil
}

// Get Recommendations/ Remarks with Emp ID
func GetExamCenterHallbyCityID(client *ent.Client, cityID int32, examyear string, examcode int32, centerid int32) ([]*ent.ExamCenterHall, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	// Check if empID is null
	if cityID == 0 {
		return nil, 422, " -STR001", false, fmt.Errorf("no city ID provided to process")
	}

	// Retrieve all records for the employee ID
	records, err := client.ExamCenterHall.Query().
		Where(examcenterhall.CityIDEQ(cityID),
			examcenterhall.ExamYearEQ(examyear),
			examcenterhall.ExamCodeEQ(examcode),
			examcenterhall.CenterCodeEQ(centerid),
			examcenterhall.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no examcenterhall found for this city ID: %d", cityID)
	}

	return records, 200, "", true, nil
}

func GetCandidateExamCenterHallbyCityID(client *ent.Client, cityID int32, examyear string, examcode int32, centerid int32, hallname string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	// Check if empID is null
	if cityID == 0 {
		return nil, 422, " -STR001", false, fmt.Errorf("no city ID provided to process")
	}

	// Retrieve all records for the employee ID
	records, err := client.Exam_Applications_IP.Query().
		Where(exam_applications_ip.ExamCityCenterCodeEQ(cityID),
			exam_applications_ip.ExamYearEQ(examyear),
			exam_applications_ip.ExamCodeEQ(examcode),
			exam_applications_ip.CenterCodeEQ(centerid),
			exam_applications_ip.HallNameEQ(hallname),
			exam_applications_ip.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no candidates in examcenterhall found for this city ID: %d", cityID)
	}

	return records, 200, "", true, nil
}

// Get All NA Verified Records
func QueryIPApplicationsByNAVerified(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.And(
				exam_applications_ip.ApplicationStatusEQ("VerifiedByNA"),
				exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
				exam_applications_ip.ExamYearEQ(id1),
				exam_applications_ip.StatusEQ("active"),
			),
		).
		WithCirclePrefRef().
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
func QueryIPApplicationsByNAVerifiedByEmpID(ctx context.Context, client *ent.Client, employeeID int64, examYear string) (*ent.Exam_Applications_IP, int32, string, bool, error) {

	record, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ApplicationStatusEQ("VerifiedByNA"), // Check for "CAVerified" status
			exam_applications_ip.EmployeeIDEQ(employeeID),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active")).
		WithIPApplicationsRef().
		WithCirclePrefRef().
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

func QueryIPApplicationsByNAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID string, facilityID1 string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {
	// Array of exams
	if facilityID == "" || facilityID1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.And(
				exam_applications_ip.ApplicationStatusEQ("VerifiedByNA"),
				exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_ip.ExamYearEQ(facilityID1),
				exam_applications_ip.StatusEQ("active"),
			),
		).
		WithCirclePrefRef().
		WithIPApplicationsRef().
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
func QueryIPApplicationsByCAVerifiedForNA(ctx context.Context, client *ent.Client, facilityID, facilityID1 string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {
	if facilityID == "" || facilityID1 == "" {
		return nil, 422, " -STR001", false, fmt.Errorf("facility ID  and ExamYear cannot be null")
	}
	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.And(
				exam_applications_ip.ApplicationStatusEQ("VerifiedByCA"),
				exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID), // Add the Where clause
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(facilityID1),
			),
		).
		WithIPApplicationsRef().
		WithCirclePrefRef().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(records) == 0 {
		return nil, 422, " -STR003", false, fmt.Errorf("no application verified by CA/NA for the Office ID %s", facilityID)
	}
	return records, 200, "", true, nil
}

// Generate ht with centercode
func GenerateHallticketNumberIP(ctx context.Context, client *ent.Client, examYear string) (string, int32, string, bool, error) {

	currentTime := time.Now().Truncate(time.Second)
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

	applications, err := tx.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.ExamCodeNEQ(0),
			//exam_applications_ip.ExamYearNEQ(""),
			exam_applications_ip.CategoryCodeNEQ(""),
			exam_applications_ip.ExamCityCenterCode(0),
			exam_applications_ip.EmployeeIDNEQ(0),
			exam_applications_ip.HallTicketGeneratedFlagNEQ(true),
			exam_applications_ip.ExamYearEQ(examYear),
		).
		Order(ent.Asc(exam_applications_ip.FieldID)).
		All(ctx)
	if err != nil {
		return "", 500, " -STR001", false, err
	} else {
		if len(applications) == 0 {
			return "", 422, " -STR002", false, errors.New("no application found for generating hallticket number")
		}
	}

	circleStats := make(map[string]int)
	for _, application := range applications {
		key := application.WorkingOfficeCircleFacilityID
		circleStats[key]++

		identificationNo := circleStats[key]
		examYear := application.ExamYear
		if len(examYear) >= 2 {
			examYear = examYear[len(examYear)-2:]
		}

		hallticketNumber := util.GenerateHallticketNumber(
			application.ExamCode,
			examYear,
			application.CategoryCode,
			util.GetCircleID(application.WorkingOfficeCircleFacilityID),

			util.GetDivisionID(application.WorkingOfficeDivisionFacilityID),
			identificationNo)

		// Validate if the hallticket number is of 12 digits
		if len(hallticketNumber) != 12 {
			continue
		}

		_, err = application.Update().
			SetHallTicketNumber(hallticketNumber).
			SetHallTicketGeneratedFlag(true).
			SetHallTicketGeneratedDate(currentTime).
			Save(ctx)
		if err != nil {
			return "", 500, " -STR001", false, err
		}
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return "", 500, " -STR002", false, err
	}

	// Return the success message with the count of eligible candidates
	successMessage := fmt.Sprintf("Hall Ticket generated successfully for %d eligible candidates", len(applications))
	return successMessage, 200, "", true, nil
}

func GenerateHallticketNumberrIP(ctx context.Context, client *ent.Client, year string, examCode int32, nodalOfficerFacilityID string) (string, int32, string, bool, error) {
	// Retrieve the last hall ticket number and extract its last four digits

	//lastFourDigitsMap := make(map[int]bool)
	lastHallTicketNumber, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamYearEQ(year),
			//exam_applications_ip.ExamCodeEQ(examCode),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			//exam_applications_ip.GenerateHallTicketFlagEQ(true),
			exam_applications_ip.HallTicketNumberNEQ(""),
			exam_applications_ip.StatusEQ("active"),
		).
		//Order(ent.Desc(exam_applications_ip.FieldHallTicketNumber)).
		Order(ent.Desc(exam_applications_ip.FieldHallIdentificationNumber)).
		First(ctx)

	// Check for error, handle not found errors explicitly
	if err != nil {
		if ent.IsNotFound(err) {
			lastHallTicketNumber = nil
		} else {
			return "", 500, " -STR001", false, err
		}
	}
	startNumber := 1
	lastFourDigits := 0
	if lastHallTicketNumber != nil {
		lastFourDigits = int(lastHallTicketNumber.HallIdentificationNumber)
		startNumber = lastFourDigits + 1
	}

	/* 	// Initialize lastHallTicketNumber if not found
	   	if lastHallTicketNumber == nil {
	   		lastHallTicketNumber = &ent.Exam_Applications_IP{HallTicketNumber: "100000000"} // Assuming ExamApplicationsIP struct
	   	}

	   	// Extract last four digits
	   	//lastFourDigitsStr := lastHallTicketNumber.HallTicketNumber[len(lastHallTicketNumber.HallTicketNumber)-4:]
	   	//lastFourDigits, err := strconv.Atoi(lastFourDigitsStr)
	   	lastFourDigits := lastHallTicketNumber.RoleUserCode
	   	if err != nil {
	   		return "", 400, " -STR002", false, errors.New("unable to get last hall ticket number")
	   	}
	   	lastFourDigitsMap[lastFourDigits] = true
	*/
	// Retrieve all eligible applications
	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamYearEQ(year),
			exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficerFacilityID),
			exam_applications_ip.GenerateHallTicketFlagEQ(true),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.Or(
				exam_applications_ip.HallTicketNumberEQ(""),
			),
		).
		Order(ent.Asc(exam_applications_ip.FieldTempHallTicket)).
		All(ctx)
	if err != nil {
		return "", 500, " -STR003", false, err
	} else {
		if len(applications) == 0 {
			return "", 422, " -STR004", false, errors.New("no application pending for hallticket generation")
		}
	}

	// If no data, set the start number to 1, else set it to the maximum number found + 1
	/* startNumber := 1
	if len(lastFourDigitsMap) > 0 {
		for lastFourDigits := range lastFourDigitsMap {
			startNumber = lastFourDigits + 1
		}
	} */

	currentTime := time.Now().Truncate(time.Second)
	// Generate hall tickets
	var successCount int
	for _, application := range applications {
		hallTicketNumber := fmt.Sprintf("%s%04d", application.TempHallTicket, startNumber)
		_, err := application.Update().
			SetHallTicketNumber(hallTicketNumber).
			SetHallTicketGeneratedFlag(true).
			SetHallTicketGeneratedDate(currentTime).
			SetHallIdentificationNumber(int32(startNumber)).
			Save(ctx)
		application.HallTicketNumber = hallTicketNumber
		if err != nil {
			return "", 500, " -STR005", false, err
		}
		startNumber++
		successCount++

	}

	// Return success message
	return fmt.Sprintf("Generated hall tickets successfully for %d eligible candidates", successCount), 200, "", true, nil
}

func UpdateExamCentresIPExamsreturnarray(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_IP) ([]UpdateResult, int32, string, bool, error) {
	var updateResults []UpdateResult

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
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
	for _, newappl := range newappls {
		// Check if ReportingOfficeID exists in the table
		exists, err := tx.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ReportingOfficeFacilityIDEQ(newappl.ReportingOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(newappl.ExamYear),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR001" + newappl.ReportingOfficeFacilityID, false, err
		}

		if !exists {
			updateResult := UpdateResult{
				ReportingOfficeID: newappl.ReportingOfficeFacilityID,
				CenterCode:        newappl.ExamCityCenterCode,
				Message:           "Reporting Office ID Does not exist and skipped",
				//RecordCount:       0,
			}
			updateResults = append(updateResults, updateResult)
			continue
		}

		// Check if CenterCode exists in the center.master table
		centerExists, err := tx.Center.
			Query().
			Where(center.IDEQ(newappl.ExamCityCenterCode)).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR002", false, err
		}
		if !centerExists {
			log.Printf("The CenterCode %d does not exist in the center.master table. Skipping to the next value in the loop.\n", newappl.ExamCityCenterCode)
			updateResult := UpdateResult{
				ReportingOfficeID: newappl.ReportingOfficeFacilityID,
				CenterCode:        newappl.ExamCityCenterCode,
				Message:           "CenterCode does not exist and skipped",
				//RecordCount:       0,
			}
			updateResults = append(updateResults, updateResult)
			continue
		}

		applications, err := tx.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ReportingOfficeFacilityIDEQ(newappl.ReportingOfficeFacilityID),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ExamYearEQ(newappl.ExamYear),
			).
			Order(ent.Desc(exam_applications_ip.FieldID)).
			All(ctx)
		if err != nil {
			return nil, 500, " -STR003", false, err
		}

		count := len(applications)
		if count > 0 {
			// Update the CenterCode for each application record
			for _, application := range applications {
				application.ExamCityCenterCode = newappl.ExamCityCenterCode
				_, err = application.Update().SetExamCityCenterCode(newappl.ExamCityCenterCode).Save(ctx)
				if err != nil {
					return nil, 500, " -STR004", false, err
				}
				// You can access the updated application values using the 'updatedApplication' variable
				// For example: updatedApplication.EmployeeID, updatedApplication.ApplicationNumber, etc.
			}

			updateResult := UpdateResult{
				ReportingOfficeID: newappl.ReportingOfficeFacilityID,
				CenterCode:        newappl.ExamCityCenterCode,
				Message:           "Updated Successfully",
				//RecordCount:       count,
			}
			updateResults = append(updateResults, updateResult)
		}
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR005", false, err
	}

	return updateResults, 200, "", true, nil
}

// Get Circle details summary ofExam Applications for the Nodal Officer Office ID. - For IP ALone
func GetEligibleApplicationsForCircleDetailsOld(ctx context.Context, client *ent.Client, examCode int32, circleOfficeID string, Examyear string) ([]map[string]interface{}, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		log.Println("No such valid exam code exists")
		return nil, fmt.Errorf(" no such valid exam code exists")
	}

	if circleOfficeID == "" {
		return nil, fmt.Errorf("please provide Nodal Officer's office ID")
	}

	// Check if circleOfficeID exists in Exam_Applications_PS table
	count, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_ip.ExamYearEQ(Examyear),
			exam_applications_ip.StatusEQ("active"),
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
	applications, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_ip.ExamYearEQ(Examyear),
			exam_applications_ip.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_ip.FieldID)).
		All(ctx)
	if err != nil {
		log.Printf("Failed to retrieve applications: %v", err)
		return nil, fmt.Errorf(" failed to retrieve applications: %v", err)
	}

	uniqueEmployees := make(map[int64]struct{})
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_IP) // Map to store the latest application for each employee
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

type EligibleApplications struct {
	ApplicationStatus       string `json:"application_status"`
	GenerateHallTicketFlag  *bool  `json:"generate_hall_ticket_flag"` // Match the database column name
	Count                   int    `json:"count"`
	EmployeeID              int64  `json:"employee_id"`
	WorkingOfficeCircleName string `json:"working_office_circle_name"`
	HallTicketGeneratedFlag *bool  `json:"hall_ticket_generated_flag"`
}

type EligibleApplicationss struct {
	ApplicationStatus       string `json:"application_status"`
	RecommendedStatus       string `json:"recommended_status"`
	HallTicketGeneratedFlag *bool  `json:"hall_ticket_generated_flag"`
	NodalOfficeName         string `json:"nodal_office_name"`
	Count                   int    `json:"count"`
}

type EligibleApplicationsss struct {
	ApplicationStatus           string `json:"application_status"`
	ControllingOfficeFacilityID string `json:"controlling_office_facility_id"`
	ControllingOfficeName       string `json:"controlling_office_name"`
	RecommendedStatus           string `json:"recommended_status"`
	HallTicketGeneratedFlag     *bool  `json:"hall_ticket_generated_flag"`
	NodalOfficeName             string `json:"nodal_office_name"`
	Count                       int    `json:"count"`
}

func GetEligibleApplicationsForCircleDetails(ctx context.Context, client *ent.Client, examCode int32, circleOfficeID string, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, fmt.Errorf("no such valid exam code exists")
	}

	if circleOfficeID == "" {
		return nil, 422, " -STR002", false, fmt.Errorf("please provide Nodal Officer's office ID")
	}

	var results []EligibleApplicationss
	// Building the query
	switch examCode {
	case 2:

		officeexist, err := client.Exam_Applications_IP.Query().
			Where(exam_applications_ip.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR003", false, err
		}
		if officeexist {
			return nil, 422, " -STR004", false, fmt.Errorf("invalid application status found")
		}

		err = client.Exam_Applications_IP.Query().
			Where(exam_applications_ip.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldApplicationStatus,
				exam_applications_ip.FieldRecommendedStatus,
				exam_applications_ip.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)

		if err != nil {
			return nil, 422, " -STR005", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	case 1:
		officeexist, err := client.Exam_Applications_PS.Query().
			Where(exam_applications_ps.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR006", false, err
		}
		if officeexist {
			return nil, 422, " -STR007", false, fmt.Errorf("invalid application status found")
		}
		err = client.Exam_Applications_PS.Query().
			Where(exam_applications_ps.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldApplicationStatus,
				exam_applications_ps.FieldRecommendedStatus,
				exam_applications_ps.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)

		if err != nil {
			return nil, 422, " -STR008", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	case 3:
		officeexist, err := client.Exam_Applications_PMPA.Query().
			Where(exam_applications_pmpa.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR009", false, err
		}
		if officeexist {
			return nil, 422, " -STR010", false, fmt.Errorf("invalid application status found")
		}
		err = client.Exam_Applications_PMPA.Query().
			Where(exam_applications_pmpa.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldApplicationStatus,
				exam_applications_pmpa.FieldRecommendedStatus,
				exam_applications_pmpa.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)

		if err != nil {
			return nil, 422, " -STR011", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	case 4:

		officeexist, err := client.Exam_Applications_GDSPA.Query().
			Where(exam_applications_gdspa.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR012", false, err
		}
		if officeexist {
			return nil, 422, " -STR013", false, fmt.Errorf("invalid application status found")
		}
		err = client.Exam_Applications_GDSPA.Query().
			Where(exam_applications_gdspa.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeName,
				exam_applications_gdspa.FieldApplicationStatus,
				exam_applications_gdspa.FieldRecommendedStatus,
				exam_applications_gdspa.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)

		if err != nil {
			return nil, 422, " -STR014", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	case 5:
		officeexist, err := client.Exam_Application_MTSPMMG.Query().
			Where(exam_application_mtspmmg.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR015", false, err
		}
		if officeexist {
			return nil, 422, " -STR016", false, fmt.Errorf("invalid application status found")
		}
		err = client.Exam_Application_MTSPMMG.Query().
			Where(exam_application_mtspmmg.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldApplicationStatus,
				exam_application_mtspmmg.FieldRecommendedStatus,
				exam_application_mtspmmg.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)

		if err != nil {
			return nil, 422, " -STR017", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	case 6:

		officeexist, err := client.Exam_Applications_GDSPM.Query().
			Where(exam_applications_gdspm.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR018", false, err
		}
		if officeexist {
			return nil, 422, " -STR019", false, fmt.Errorf("invalid application status found")
		}
		err = client.Exam_Applications_GDSPM.Query().
			Where(exam_applications_gdspm.NodalOfficeFacilityIDEQ(circleOfficeID),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldApplicationStatus,
				exam_applications_gdspm.FieldRecommendedStatus,
				exam_applications_gdspm.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)

		if err != nil {

			return nil, 422, " -STR020", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	default:
		return nil, 422, " -STR021", false, fmt.Errorf("invalid exam code")
	}
	if len(results) == 0 {
		return nil, 422, " -STR022", false, fmt.Errorf("no data exists for the given parameters")
	}
	circleName := ""

	// Variables to hold counts
	permittedCount := 0
	notPermittedCount := 0
	pendingCount := 0
	pendingWithCandidateCount := 0
	hallticketGeneratedCount := 0
	hallticketNotGeneratedCount := 0
	employeeCount := 0

	for _, result := range results {
		circleName = result.NodalOfficeName
		if result.ApplicationStatus == "VerifiedByCA" || result.ApplicationStatus == "VerifiedByNA" {
			if result.RecommendedStatus == "Recommended" || result.RecommendedStatus == "Provisionally Recommended" {
				permittedCount += result.Count
				employeeCount += result.Count
			} else if result.RecommendedStatus == "Not Recommended" {
				notPermittedCount += result.Count
				employeeCount += result.Count
			}
		} else if result.ApplicationStatus == "PendingWithCandidate" {
			pendingWithCandidateCount += result.Count
			employeeCount += result.Count
		} else {
			pendingCount += result.Count
			employeeCount += result.Count
		}

		if (result.ApplicationStatus == "VerifiedByCA" || result.ApplicationStatus == "VerifiedByNA") &&
			(result.RecommendedStatus == "Recommended" || result.RecommendedStatus == "Provisionally Recommended") {
			if *result.HallTicketGeneratedFlag {
				hallticketGeneratedCount += result.Count
			} else {
				hallticketNotGeneratedCount += result.Count
			}
		}
	}
	result := []map[string]interface{}{
		{
			"CircleID":                      circleOfficeID,
			"CircleName":                    circleName,
			"NoOfApplicationsReceived":      employeeCount,
			"NoPermitted":                   permittedCount,
			"NoNotPermitted":                notPermittedCount,
			"NoPending":                     pendingCount,
			"NoPendingWithCandidate":        pendingWithCandidateCount,
			"NoHallticketGeneratedCount":    hallticketGeneratedCount,
			"NoHallticketNotGeneratedCount": hallticketNotGeneratedCount,
		},
	}

	return result, 200, "", true, nil
}

// 	// Calculate the number of unique employees
// 	employeeCount, err := client.Exam_Applications_IP.Query().
// 		Where(exam_applications_ip.NodalOfficeFacilityIDEQ(circleOfficeID),
// 			exam_applications_ip.ExamYearEQ(examYear),
// 			exam_applications_ip.StatusEQ("active"),
// 		).
// 		GroupBy(exam_applications_ip.FieldEmployeeID).
// 		Count(ctx)
// 	if err != nil {
// 		log.Printf("Failed to count unique employees: %v", err)
// 		return nil, fmt.Errorf("failed to count unique employees: %v", err)
// 	}

// 	// Get the circle name (assuming there's a way to get it from the circleOfficeID)
// 	circleName := "Example Circle Name" // Replace with actual logic to get the circle name

// 	result := []map[string]interface{}{
// 		{
// 			"CircleID":                      circleOfficeID,
// 			"CircleName":                    circleName,
// 			"NoOfApplicationsReceived":      employeeCount,
// 			"NoPermitted":                   permittedCount,
// 			"NoNotPermitted":                notPermittedCount,
// 			"NoPending":                     pendingCount,
// 			"NoPendingWithCandidate":        pendingWithCandidateCount,
// 			"NoHallTicketGeneratedCount":    hallTicketGeneratedCount,
// 			"NoHallTicketNotGeneratedCount": hallTicketNotGeneratedCount,
// 		},
// 	}

// 	return result, nil
// }

// Get IP Exam statistics
func GetIPExamStatistics(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Query to get the applications from Exam_Applications_IP table
	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamCodeEQ(examCode),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_ip.FieldEmployeeID), ent.Desc(exam_applications_ip.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	// Create a map to store the latest application for each employee
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_IP)

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
type CircleWiseSummary struct {
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

func GetIPExamStatisticsCircleWise(ctx context.Context, client *ent.Client, examCode int32, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Query to get the applications from Exam_Applications_IP table
	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamCodeEQ(examCode),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active")).
		Order(ent.Asc(exam_applications_ip.FieldEmployeeID), ent.Desc(exam_applications_ip.FieldID)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	// Create a map to store employee-wise latest applications
	employeeLatestApplication := make(map[int64]*ent.Exam_Applications_IP)

	// Loop through the applications and store the latest application for each employee
	for _, app := range applications {
		if _, found := employeeLatestApplication[app.EmployeeID]; !found {
			employeeLatestApplication[app.EmployeeID] = app
		}
	}

	// Create a map to store circle-wise summaries
	circleSummaries := make(map[string]*CircleWiseSummary)

	// Loop through the latest applications to update counts
	for _, app := range employeeLatestApplication {
		circleOfficeID := app.NodalOfficeFacilityID
		if circleSummaries[circleOfficeID] == nil {
			approvalFlag, err := getApprovalFlagForHallTicket(client, circleOfficeID)
			if err != nil {
				log.Printf("Failed to get ApprovalFlagForHallTicket for CircleOfficeID %v: %v", circleOfficeID, err)
				continue
			}
			circleSummaries[circleOfficeID] = &CircleWiseSummary{
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

func getApprovalFlagForHallTicket(client *ent.Client, circleOfficeID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	circleMaster, err := client.CircleSummaryForNO.
		Query().
		Where(circlesummaryforno.CircleOfficeIdEQ(circleOfficeID)).
		Only(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get CircleMaster for CircleOfficeID %v: %v", circleOfficeID, err)
	}

	return circleMaster.ApproveHallTicketGenrationIP, nil
}

// IP office wise stats for NO
type DOOfficeWiseSummary struct {
	CircleName                  string
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

// // Get All Pending with Candidate
// Assuming Exam_Applications_IP has a field named EmployeeID, you might adapt the code like this:
func QueryIPApplicationsByPendingWithCandidate(ctx context.Context, client *ent.Client, facilityID string, id1 string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {
	if facilityID == "" || id1 == "" {
		return nil, 422, " -STR001", false, errors.New("facility ID  and Exam Year cannot be empty")
	}

	// Fetch all applications matching the criteria
	records, err := client.Exam_Applications_IP.Query().
		Where(
			exam_applications_ip.ControllingOfficeFacilityIDEQ(facilityID),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamYearEQ(id1),
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
	latestApplications := make(map[int64]*ent.Exam_Applications_IP)

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
	var result []*ent.Exam_Applications_IP
	for _, application := range latestApplications {
		result = append(result, application)
	}

	if len(result) == 0 {
		return nil, 422, " -STR004", false, fmt.Errorf("no Applicationsf found pending with candiadte under this Office ID %s", facilityID)
	}

	return result, 200, "", true, nil
}

func GetIPExamStatisticsDOOfficeWiseLOld(ctx context.Context, client *ent.Client, examCode int32, facilityID string, Examyear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		return nil, 422, " -STR002", false, errors.New("facility ID cannot be null")
	}

	// Fetch all applications for the given facilityID
	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamYearEQ(Examyear),
		).
		Order(ent.Asc(exam_applications_ip.FieldEmployeeID), ent.Desc(exam_applications_ip.FieldUpdatedAt)).
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
				reportingOfficeID := app.ControllingOfficeFacilityID //-----Here reportingOfficeID is a variable name it is related to ControllingOffice only

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

func GetIPExamStatisticsDOOfficeWiseLNew(ctx context.Context, client *ent.Client, examCode int32, facilityID string, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		return nil, 422, " -STR002", false, errors.New("facility ID cannot be null")
	}
	var results []EligibleApplicationsss

	err := client.Exam_Applications_IP.Query().
		Where(exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
			exam_applications_ip.ExamYearEQ(examYear),
			exam_applications_ip.StatusEQ("active"),
		).
		GroupBy(
			exam_applications_ip.FieldNodalOfficeName,
			exam_applications_ip.FieldControllingOfficeFacilityID,
			exam_applications_ip.FieldControllingOfficeName,
			exam_applications_ip.FieldApplicationStatus,
			exam_applications_ip.FieldRecommendedStatus,
			exam_applications_ip.FieldHallTicketGeneratedFlag).
		Aggregate(ent.Count()).
		Scan(ctx, &results)

	if err != nil {
		return nil, 500, " -STR003", false, err
	}
	if len(results) == 0 {
		return nil, 422, " -STR004", false, fmt.Errorf("no data exists for the given parameters")
	}

	// Create a map to store the greatest ApplicationID for each unique employee_id
	//greatestAppIDs := make(map[int64]int32)
	circleName := ""
	controllingOfficeName := ""
	controllingOfficeFacilityId := ""
	permittedCount := 0
	notPermittedCount := 0
	pendingCount := 0
	pendingWithCandidateCount := 0
	hallticketGeneratedCount := 0
	hallticketNotGeneratedCount := 0
	employeeCount := 0
	// Loop through applications to find the greatest ApplicationID for each employee_id
	for _, result := range results {
		circleName = result.NodalOfficeName
		controllingOfficeName = result.ControllingOfficeName
		controllingOfficeFacilityId = result.ControllingOfficeFacilityID
		if result.ApplicationStatus == "VerifiedByCA" || result.ApplicationStatus == "VerifiedByNA" {
			if result.RecommendedStatus == "Recommended" || result.RecommendedStatus == "Provisionally Recommended" {
				permittedCount += result.Count
				employeeCount += result.Count
			} else if result.RecommendedStatus == "Not Recommended" {
				notPermittedCount += result.Count
				employeeCount += result.Count
			}
		} else if result.ApplicationStatus == "PendingWithCandidate" {
			pendingWithCandidateCount += result.Count
			employeeCount += result.Count
		} else {
			pendingCount += result.Count
			employeeCount += result.Count
		}

		if (result.ApplicationStatus == "VerifiedByCA" || result.ApplicationStatus == "VerifiedByNA") &&
			(result.RecommendedStatus == "Recommended" || result.RecommendedStatus == "Provisionally Recommended") {
			if *result.HallTicketGeneratedFlag {
				hallticketGeneratedCount += result.Count
			} else {
				hallticketNotGeneratedCount += result.Count
			}
		}
	}
	result := []map[string]interface{}{
		{
			"CircleID":                      facilityID,
			"CircleName":                    circleName,
			"ControllingOfficeFacilityID":   controllingOfficeFacilityId,
			"ControllingOfficeName":         controllingOfficeName,
			"NoOfApplicationsReceived":      employeeCount,
			"NoPermitted":                   permittedCount,
			"NoNotPermitted":                notPermittedCount,
			"NoPending":                     pendingCount,
			"NoPendingWithCandidate":        pendingWithCandidateCount,
			"NoHallticketGeneratedCount":    hallticketGeneratedCount,
			"NoHallticketNotGeneratedCount": hallticketNotGeneratedCount,
		},
	}

	return result, 200, "", true, nil
}
func GetExamStatisticsDOOfficeWise(ctx context.Context, client *ent.Client, examCode int32, facilityID string, examYear string) ([]map[string]interface{}, int32, string, bool, error) {
	// Check if exam code is valid
	if examCode <= 0 {
		return nil, 422, " -STR001", false, errors.New("no such valid exam code exists")
	}

	// Check if facilityID is provided
	if facilityID == "" {
		return nil, 422, " -STR002", false, errors.New("facility ID cannot be null")
	}
	var results []EligibleApplicationsss

	switch examCode {
	case 2:
		officeexist, err := client.Exam_Applications_IP.Query().
			Where(exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR003", false, err
		}
		if officeexist {
			return nil, 422, " -STR004", false, fmt.Errorf("invalid application status found")
		}

		err = client.Exam_Applications_IP.Query().
			Where(exam_applications_ip.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_ip.FieldNodalOfficeName,
				exam_applications_ip.FieldControllingOfficeFacilityID,
				exam_applications_ip.FieldControllingOfficeName,
				exam_applications_ip.FieldApplicationStatus,
				exam_applications_ip.FieldRecommendedStatus,
				exam_applications_ip.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)
		if err != nil {

			return nil, 422, " -STR005", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	case 1:
		officeexist, err := client.Exam_Applications_PS.Query().
			Where(exam_applications_ps.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR006", false, err
		}
		if officeexist {
			return nil, 422, " -STR007", false, fmt.Errorf("invalid application status found")
		}

		err = client.Exam_Applications_PS.Query().
			Where(exam_applications_ps.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_ps.FieldNodalOfficeName,
				exam_applications_ps.FieldControllingOfficeFacilityID,
				exam_applications_ps.FieldControllingOfficeName,
				exam_applications_ps.FieldApplicationStatus,
				exam_applications_ps.FieldRecommendedStatus,
				exam_applications_ps.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)
		if err != nil {

			return nil, 422, " -STR008", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	case 3:
		officeexist, err := client.Exam_Applications_PMPA.Query().
			Where(exam_applications_pmpa.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR009", false, err
		}
		if officeexist {
			return nil, 422, " -STR010", false, fmt.Errorf("invalid application status found")
		}

		err = client.Exam_Applications_PMPA.Query().
			Where(exam_applications_pmpa.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_pmpa.FieldNodalOfficeName,
				exam_applications_pmpa.FieldControllingOfficeFacilityID,
				exam_applications_pmpa.FieldControllingOfficeName,
				exam_applications_pmpa.FieldApplicationStatus,
				exam_applications_pmpa.FieldRecommendedStatus,
				exam_applications_pmpa.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)
		if err != nil {

			return nil, 422, " -STR011", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	case 4:
		officeexist, err := client.Exam_Applications_GDSPA.Query().
			Where(exam_applications_gdspa.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR012", false, err
		}
		if officeexist {
			return nil, 422, " -STR013", false, fmt.Errorf("invalid application status found")
		}

		err = client.Exam_Applications_GDSPA.Query().
			Where(exam_applications_gdspa.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_gdspa.FieldNodalOfficeName,
				exam_applications_gdspa.FieldControllingOfficeFacilityID,
				exam_applications_gdspa.FieldControllingOfficeName,
				exam_applications_gdspa.FieldApplicationStatus,
				exam_applications_gdspa.FieldRecommendedStatus,
				exam_applications_gdspa.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)
		if err != nil {

			return nil, 422, " -STR014", false, fmt.Errorf("failed querying sub-query: %v", err)
		}

	case 5:
		officeexist, err := client.Exam_Application_MTSPMMG.Query().
			Where(exam_application_mtspmmg.NodalOfficeFacilityIDEQ(facilityID),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR015", false, err
		}
		if officeexist {
			return nil, 422, " -STR016", false, fmt.Errorf("invalid application status found")
		}

		err = client.Exam_Application_MTSPMMG.Query().
			Where(exam_application_mtspmmg.NodalOfficeFacilityIDEQ(facilityID),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
			).
			GroupBy(
				exam_application_mtspmmg.FieldNodalOfficeName,
				exam_application_mtspmmg.FieldControllingOfficeFacilityID,
				exam_application_mtspmmg.FieldControllingOfficeName,
				exam_application_mtspmmg.FieldApplicationStatus,
				exam_application_mtspmmg.FieldRecommendedStatus,
				exam_application_mtspmmg.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)
		if err != nil {

			return nil, 422, " -STR017", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	case 6:
		officeexist, err := client.Exam_Applications_GDSPM.Query().
			Where(exam_applications_gdspm.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.ApplicationStatusNotIn("VerifiedByCA", "VerifiedByNA", "PendingWithCandidate",
					"CAVerificationPending", "ResubmitCAVerificationPending"),
			).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR018", false, err
		}
		if officeexist {
			return nil, 422, " -STR019", false, fmt.Errorf("invalid application status found")
		}

		err = client.Exam_Applications_GDSPM.Query().
			Where(exam_applications_gdspm.NodalOfficeFacilityIDEQ(facilityID),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
			).
			GroupBy(
				exam_applications_gdspm.FieldNodalOfficeName,
				exam_applications_gdspm.FieldControllingOfficeFacilityID,
				exam_applications_gdspm.FieldControllingOfficeName,
				exam_applications_gdspm.FieldApplicationStatus,
				exam_applications_gdspm.FieldRecommendedStatus,
				exam_applications_gdspm.FieldHallTicketGeneratedFlag).
			Aggregate(ent.Count()).
			Scan(ctx, &results)
		if err != nil {

			return nil, 422, " -STR020", false, fmt.Errorf("failed querying sub-query: %v", err)
		}
	default:
		return nil, 422, " -STR021", false, fmt.Errorf("unsupported exam code")
	}
	if len(results) == 0 {
		return nil, 422, " -STR022", false, fmt.Errorf("no data exists for the given parameters")
	}

	// Create a map to store reporting office-wise summaries
	doOfficeSummaries := make(map[string]*DOOfficeWiseSummary)

	// Loop through the grouped results to build the summaries
	for _, result := range results {
		controllingOfficeID := result.ControllingOfficeFacilityID

		if doOfficeSummaries[controllingOfficeID] == nil {
			doOfficeSummaries[controllingOfficeID] = &DOOfficeWiseSummary{
				CircleName:                  result.NodalOfficeName,
				ControllingOfficeFacilityID: controllingOfficeID,
				ControllingOfficeName:       result.ControllingOfficeName,
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

		// Update the summary based on the application status and hall ticket flags
		doOfficeSummary := doOfficeSummaries[controllingOfficeID]
		doOfficeSummary.Received += result.Count

		if result.ApplicationStatus == "VerifiedByCA" || result.ApplicationStatus == "VerifiedByNA" {
			if result.RecommendedStatus == "Recommended" || result.RecommendedStatus == "Provisionally Recommended" {
				doOfficeSummary.Permitted += result.Count
			} else if result.RecommendedStatus == "Not Recommended" {
				doOfficeSummary.NotPermitted += result.Count
			}
		} else if result.ApplicationStatus == "PendingWithCandidate" {
			doOfficeSummary.PendingWithCandidate += result.Count
		} else {
			doOfficeSummary.PendingWithCA += result.Count
		}

		if (result.ApplicationStatus == "VerifiedByCA" || result.ApplicationStatus == "VerifiedByNA") &&
			(result.RecommendedStatus == "Recommended" || result.RecommendedStatus == "Provisionally Recommended") {
			if *result.HallTicketGeneratedFlag {
				doOfficeSummary.HallTicketGenerated += result.Count
			} else {
				doOfficeSummary.HallTicketNotGenerated += result.Count
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
			"SNo":                           serialNumber,
			"CircleID":                      facilityID,
			"CircleName":                    summary.CircleName,
			"ControllingOfficeFacilityID":   summary.ControllingOfficeFacilityID,
			"ControllingOfficeName":         summary.ControllingOfficeName,
			"NoOfApplicationsReceived":      summary.Received,
			"NoPermitted":                   summary.Permitted,
			"NoNotPermitted":                summary.NotPermitted,
			"NoPending":                     summary.PendingWithCA,
			"NoPendingWithCandidate":        summary.PendingWithCandidate,
			"NoHallticketGeneratedCount":    summary.HallTicketGenerated,
			"NoHallticketNotGeneratedCount": summary.HallTicketNotGenerated,
		})
	}

	return result, 200, "", true, nil
}

type ExamStats struct {
	CircleName                  string
	ControllingOfficeName       string
	ControllingOfficeFacilityID string
	NodalOfficeFacilityID       string
	NoOfCandidatesChosenCity    int
	NoOfCandidatesAlloted       int
}

func SubGetExamApplicationsByCenterIP(ctx context.Context, client *ent.Client, ExamYear string, SubExamcode string, CenterId string, StartNo string, EndNo string) ([]*ent.Exam_Applications_IP, int32, string, bool, error) {

	// Convert strings to int64
	StrExamcode, err1 := strconv.ParseInt(SubExamcode, 10, 32)
	if err1 != nil {
		return nil, 422, " -STR001", false, errors.New("invalid exam code")
	}

	StrCenterId, err2 := strconv.ParseInt(CenterId, 10, 32)
	if err2 != nil {
		return nil, 422, " -STR002", false, errors.New("invalid Center Id")
	}

	StrStartno, err2 := strconv.ParseInt(StartNo, 10, 32)
	if err2 != nil {
		return nil, 422, " -STR003", false, errors.New("invalid Start number")
	}
	StrEndno, err2 := strconv.ParseInt(EndNo, 10, 32)
	if err2 != nil {
		return nil, 422, " -STR004", false, errors.New("invalid End Number")
	}

	// Convert int64 to int32
	Examcode := int32(StrExamcode)
	Centerid := int32(StrCenterId)
	Startno := int32(StrStartno)
	Endno := int32(StrEndno)
	if ExamYear == " " {
		return nil, 422, " -STR005", false, errors.New("exam year should not null")
	}

	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamYearEQ(ExamYear),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.HallTicketNumberNEQ(""),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamCodeEQ(Examcode),
			exam_applications_ip.CenterCodeEQ(Centerid),
			exam_applications_ip.HallTicketGeneratedFlag(true),
			exam_applications_ip.HallIdentificationNumberGTE(Startno),
			exam_applications_ip.HallIdentificationNumberLTE(Endno),
		).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR006", false, err
	} else {
		if len(applications) == 0 {
			return nil, 422, " -STR007", false, errors.New("no matching data with Exam center")
		}
	}

	return applications, 200, "", true, nil
}

func GetExamApplicatonsPreferenenceCityWiseStats(ctx context.Context, client *ent.Client, ExamYear string, SubExamcode string, SubCityid string) ([]ExamStats, int32, string, bool, error) {
	var result []ExamStats

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

	applications, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.ExamYearEQ(ExamYear),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.HallTicketNumberNEQ(""),
			exam_applications_ip.StatusEQ("active"),
			exam_applications_ip.ExamCodeEQ(Examcode),
			exam_applications_ip.CenterIdEQ(Cityid),
			exam_applications_ip.HallTicketGeneratedFlag(true),
		).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR004", false, err
	} else {
		if len(applications) == 0 {
			return nil, 422, " -STR005", false, errors.New("no matching data with Exam City")
		}
	}

	groupedApplications := make(map[string]ExamStats)

	for _, examApplication := range applications {
		nodalOfficeFacilityID := examApplication.NodalOfficeFacilityID
		nodalOfficeName := examApplication.NodalOfficeName

		controllingOfficeFacilityID := examApplication.ControllingOfficeFacilityID
		controllingOfficeName := examApplication.ControllingOfficeName

		centerCode := examApplication.CenterCode

		// Check if ExamStats for the reporting office already exists
		stats, ok := groupedApplications[controllingOfficeName]
		if !ok {
			stats = ExamStats{
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

func UpdateCenterCodeForApplications(ctx context.Context, client *ent.Client, controllingOfficeFacilityID string, examCenterID, seatsToAllot, examCityID int32) (int, []*ent.Exam_Applications_IP, int32, string, bool, error) {
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
	applications, err := tx.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.CenterCodeIsNil(),
			exam_applications_ip.CenterIdEQ(examCityID),
			exam_applications_ip.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			exam_applications_ip.ControllingOfficeFacilityIDEQ(controllingOfficeFacilityID),
			exam_applications_ip.HallTicketNumberNEQ(""),
			exam_applications_ip.HallTicketGeneratedFlag(true),
			exam_applications_ip.StatusEQ("active"),
		).
		Order(ent.Asc(exam_applications_ip.FieldApplnSubmittedDate)).
		Limit(int(seatsToAllot)). // Limit the number of records to be updated
		All(ctx)

	if err != nil {
		return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR002", false, err
	}
	var updatedCount int

	for _, application := range applications {
		_, err := tx.Exam_Applications_IP.
			UpdateOne(application).
			SetCenterCode(examCenterID).
			Save(ctx)

		if err != nil {
			return 0, nil, 500, controllingOfficeFacilityID + "-" + strExamCenterID + " -STR003", false, err
		}

		updatedCount++

	}

	// Counting Updated Applications
	updateCount, err := tx.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.CenterCodeEQ(examCenterID),
			exam_applications_ip.StatusEQ("active"),
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

	// Success
	return updatedCount, applications, 200, "", true, nil
}
