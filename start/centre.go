package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/center"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"
	"strconv"

	//"recruit/ent/divisionmaster"
	"recruit/ent/exam"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/exam_applications_gdspm"
	"recruit/ent/exam_applications_pmpa"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/examcitycenter"

	//"recruit/ent/regionmaster"

	//"strings"
	"time"
)

func CreateCenter(client *ent.Client, newCenter *ca_reg.CenterReq) (*ent.Center, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	currentTime := time.Now().Truncate(time.Second)
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

	examdetail, err := tx.Exam.
		Query().
		Where(
			exam.IDEQ(newCenter.ExamCode)).
		Only(ctx)
	if err != nil {

		return nil, 422, " -STR001", false, fmt.Errorf("failed to retrieve data: %w", err)
	}

	examname := examdetail.ExamShortName

	u, err := tx.Center.Create().
		SetExamCenterName(newCenter.ExamCenterName).
		//	SetNodalOfficerCode(newCenter.NodalOfficerCode).
		SetExamCode(newCenter.ExamCode).
		//	SetRegionID(newCenter.RegionID).
		//	SetCircleID(newCenter.CircleID).
		//	SetDivisionID(newCenter.DivisionID).
		//SetFacilityID(newCenter.FacilityID).
		//	SetExamNameCode(newCenter.ExamNameCode).
		SetExamName(examname).
		SetConductedBy(newCenter.ConductedBy).
		SetConductedByFacilityID(newCenter.ConductedByFacilityID).
		SetExamYear(newCenter.ExamYear).
		SetNAUserName(newCenter.NAUserName).
		SetNodalOfficeFacilityId(newCenter.NodalOfficeFacilityID).
		//SetAdminCircleOfficeID(newCenter.AdminCircleOfficeID).
		SetAddress(newCenter.Address).
		SetLandmark(newCenter.Landmark).
		SetPincode(newCenter.Pincode).
		SetMaxSeats(newCenter.MaxSeats).
		//SetNoAlloted(newCenter.NoAlloted).
		SetPendingSeats(newCenter.MaxSeats).
		SetStatus(newCenter.Status).
		SetUpdatedAt(currentTime).
		SetCityID(newCenter.CityID).
		//SetUpdatedBy(newCenter.UpdatedBy).
		SetCenterCityName(newCenter.CenterCityName).
		//SetPapers(newCenter.Papers).
		Save(ctx)
	if err != nil {
		log.Println("error creating newCenter:", err)
		return nil, 422, " -STR002", false, fmt.Errorf("failed creating newCenter: %w", err)
	}
	log.Println("newCenter was created:", u)
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR003", false, err
	}

	return u, 200, "", true, nil
}

/*func QueryExamByName(ctx context.Context, client *ent.Client, examname string) (*ent.Exam, error) {
	u, err := client.Exam.Query().Where(exam.ExamName(examname)).Only(ctx)

	if err != nil {
		log.Println("error at gettting examname: ", err)
		return nil, fmt.Errorf("failed querying exam: %w", err)
	}
	log.Println("exam returned by name: ", u)
	return u, nil
}*/

func QueryCenterID(ctx context.Context, client *ent.Client, id int32) (*ent.Center, int32, string, bool, error) {
	// Check if id is zero
	if id == 0 {
		return nil, 422, "  -STR001", false, fmt.Errorf(" input parameter is missing")
	}

	u, err := client.Center.Get(ctx, id)
	if err != nil {

		return nil, 422, "  -STR002", false, fmt.Errorf("failed querying Center: %w", err)
	}
	log.Println("Center returned: ", u)
	return u, 200, "", true, nil
}

func QueryVersion(ctx context.Context, client *ent.Client) ([]*ent.Version, error) {
	// Perform query to fetch the version data (replace with your actual query logic)
	version, err := client.Version.Query().
		All(ctx)
	if err != nil {
		log.Printf("Error querying version: %v\n", err)
		return nil, fmt.Errorf("failed querying version: %w", err)
	}

	log.Printf("Version fetched: %+v\n", version)
	return version, nil
}

func QueryCenter(ctx context.Context, client *ent.Client) ([]*ent.Center, error) {
	//Array of exams
	u, err := client.Center.Query().All(ctx)
	if err != nil {
		log.Println("error at Center: ", err)
		return nil, fmt.Errorf("failed querying Center: %w", err)
	}
	log.Println("Center returned: ", u)
	return u, nil
}

func QueryCenterByCity(ctx context.Context, client *ent.Client, cities string) ([]*ent.Center, int32, error) {
	//Array of exams
	//cities := city
	if cities == " " || len(cities) == 0 {
		return nil, 422, errors.New("enter valid City Name")
	}

	examCenters, err := client.Center.Query().Where(
		center.CenterCityName(cities)).
		All(ctx)
	if err != nil {
		return nil, 400, fmt.Errorf("failed to fetch Center data: %w", err)
	} else {
		if len(cities) == 0 {
			return nil, 422, errors.New("no matching Exam Center found")
		}
	}
	return examCenters, 200, nil
}

func QueryCenterByCircleID(ctx context.Context, client *ent.Client, circle int32) ([]*ent.Center, int32, error) {

	examCenters, err := client.Center.Query().
		Where(
			center.CircleIDEQ(circle)).
		All(ctx)
	if err != nil {
		return nil, 400, fmt.Errorf("failed to fetch data for Exam Center: %w", err)
	}
	if len(examCenters) == 0 {
		return nil, 422, errors.New("no matching found for Exam Center")
	}

	return examCenters, 200, nil
}

func SubGetCentersByConductingAuthority(ctx context.Context, client *ent.Client, examYear string, examCode string, conductingId string) ([]*ent.Center, int32, string, bool, error) {

	StrExamcode, err2 := strconv.ParseInt(examCode, 10, 32)
	if err2 != nil {
		return nil, 422, " -STR004", false, errors.New("invalid Exam Code")
	}

	// Convert int64 to int32
	Examcode := int32(StrExamcode)

	examCenters, err := client.Center.Query().
		Where(
			center.ConductedByFacilityIDEQ(conductingId),
			center.ExamYearEQ(examYear),
			center.ExamCodeEQ(Examcode),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	if len(examCenters) == 0 {
		return nil, 422, " -STR002", false, errors.New("no matching found for Exam Center")
	}

	return examCenters, 200, "", true, nil
}

func DeleteCenterID(client *ent.Client, id int32) error {

	//context not passed for delete dont know why?
	err := client.Center.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

// Working one to return array of updated records/ skipped records.
type UpdateResult struct {
	ReportingOfficeID string `json:"ReportingOfficeID"`
	CenterCode        int32  `json:"CenterCode"`
	Message           string `json:"Message"`
	//RecordCount       int    `json:"RecordCount"`
}

// Get Exam Centers with ExamCode & NO Office ID
func GetExamCentersByExamCodeNOOfficeID(ctx context.Context, client *ent.Client, examCode int32, Cityid int32, noOfficeID string, id3 string) ([]*ent.Center, int32, string, bool, error) {
	// Validate the parameters
	if noOfficeID == "" {
		return nil, 422, " -STR001", false, fmt.Errorf("NOOfficeID is mandatory")
	}
	if examCode <= 0 {
		return nil, 422, " -STR002", false, fmt.Errorf("please provide a valid ExamCode")
	}

	// Retrieve the exam centers that match the NOOfficeID and ExamCode
	centers, err := client.Center.
		Query().
		Where(
			center.ExamCodeEQ(examCode),
			center.NodalOfficeFacilityIdEQ(noOfficeID),
			center.ExamYearEQ(id3),
			center.CityID(Cityid),
			center.Status(true),
		).
		All(ctx)
	if err != nil {

		return nil, 422, " -STR003", false, fmt.Errorf("failed to query exam centers: %v", err)
	}

	if len(centers) == 0 {
		return nil, 422, " -STR004", false, fmt.Errorf("no valid Exam Centers for the Nodal Officer Office ID/ExamCode")
	}

	return centers, 200, "", true, nil
}

func UpdateExamCentresPSExams(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_PS) ([]string, error) {
	var updatedRecords []string

	for _, newappl := range newappls {
		applications, err := client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ReportingOfficeFacilityIDEQ(newappl.ReportingOfficeFacilityID),
				exam_applications_ps.GenerateHallTicketFlag(false),
				exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Order(ent.Desc(exam_applications_ps.FieldID)).
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

// PS Grp B Exams
func UpdateExamCentresPSExamsreturnarray(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_PS) ([]UpdateResult, int32, string, bool, error) {
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
		exists, err := tx.Exam_Applications_PS.
			Query().
			Where(exam_applications_ps.ReportingOfficeFacilityIDEQ(newappl.ReportingOfficeFacilityID)).
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

		applications, err := tx.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ReportingOfficeFacilityIDEQ(newappl.ReportingOfficeFacilityID),
				exam_applications_ps.GenerateHallTicketFlag(true),
				exam_applications_ps.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Order(ent.Desc(exam_applications_ps.FieldID)).
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

func UpdateExamCentresPMPAExams(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_PMPA) ([]string, error) {
	var updatedRecords []string

	for _, newappl := range newappls {
		applications, err := client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ReportingOfficeIDEQ(newappl.ReportingOfficeID),
				exam_applications_pmpa.GenerateHallTicketFlag(false),
				exam_applications_pmpa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Order(ent.Desc(exam_applications_pmpa.FieldID)).
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

func UpdateExamCentresPMPAExamsreturnarray(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_PMPA) ([]UpdateResult, error) {
	var updateResults []UpdateResult

	for _, newappl := range newappls {
		// Check if ReportingOfficeID exists in the table
		exists, err := client.Exam_Applications_PMPA.
			Query().
			Where(exam_applications_pmpa.ReportingOfficeIDEQ(newappl.ReportingOfficeID)).
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

		applications, err := client.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.ReportingOfficeIDEQ(newappl.ReportingOfficeID),
				//exam_applications_pmpa.GenerateHallTicketFlag(true),
				//exam_applications_pmpa.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
			).
			Order(ent.Desc(exam_applications_pmpa.FieldID)).
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

func UpdateExamCentresGDSPMExams(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_GDSPM) ([]string, error) {
	var updatedRecords []string

	for _, newappl := range newappls {
		applications, err := client.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ReportingOfficeIDEQ(newappl.ReportingOfficeID),
				exam_applications_gdspm.GenerateHallTicketFlag(false),
				exam_applications_gdspm.ApplicationStatusIn("VerifiedByCA", "VerifiedByNA"),
				exam_applications_gdspm.Status("active"),
				exam_applications_gdspm.ExamYearEQ(newappl.ExamYear),
			).
			Order(ent.Desc(exam_applications_gdspm.FieldID)).
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

func UpdateExamCentresGDSPMExamsreturnarray(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Applications_GDSPM) ([]UpdateResult, int32, string, bool, error) {
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
		exists, err := tx.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ReportingOfficeIDEQ(newappl.ReportingOfficeID),
				exam_applications_gdspm.Status("active"),
				exam_applications_gdspm.ExamYearEQ(newappl.ExamYear),
			).
			Exist(ctx)
		if err != nil {
			//Failed to check ReportingOfficeID existence
			return nil, 500, " -STR001 " + newappl.ReportingOfficeFacilityID, false, err
		}
		if !exists {
			//The ReportingOfficeID %s does not exist in the Applications table. Skipping to the next value in the loop.
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
		centerExists, err := tx.Center.
			Query().
			Where(center.IDEQ(newappl.ExamCityCenterCode)).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR002", false, err
		}
		if !centerExists {
			//The CenterCode %d does not exist in the center.master table. Skipping to the next value in the loop
			updateResult := UpdateResult{
				ReportingOfficeID: newappl.ReportingOfficeID,
				CenterCode:        newappl.ExamCityCenterCode,
				Message:           "CenterCode does not exist and skipped",
				//RecordCount:       0,
			}
			updateResults = append(updateResults, updateResult)
			continue
		}

		applications, err := tx.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.ReportingOfficeIDEQ(newappl.ReportingOfficeID),
				exam_applications_gdspm.Status("active"),
				exam_applications_gdspm.ExamYearEQ(newappl.ExamYear),
			).
			Order(ent.Desc(exam_applications_gdspm.FieldID)).
			All(ctx)
		if err != nil {
			return nil, 500, " -STR003", false, err
		}
		count := len(applications)
		if count == 0 {
			return nil, 400, " -STR004", false, errors.New(" no reporting office found")
		}

		if count > 0 {
			// Update the CenterCode for each application record
			for _, application := range applications {
				application.ExamCityCenterCode = newappl.ExamCityCenterCode
				_, err = application.Update().SetExamCityCenterCode(newappl.ExamCityCenterCode).Save(ctx)
				if err != nil {
					return nil, 500, " -STR005", false, err
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
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR006", false, err
	}

	return updateResults, 200, "", true, nil
}

func QueryCityNamesByExamIDFromExamCenter(client *ent.Client, citylists *ent.ExamCityCenter, nn string, cid int32, code int32) ([]*ent.ExamCityCenter, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	notiNo := nn
	crid := cid
	codeid := code

	cities, err := client.ExamCityCenter.Query().Where(examcitycenter.NodalOfficeFacilityIDEQ(notiNo), examcitycenter.ExamYearEQ(crid), examcitycenter.ExamCodeEQ(codeid), examcitycenter.StatusEQ("active")).
		All(ctx)
	if err != nil {
		log.Println("error at Center cities Master: ", err)
		return nil, 422, "  -STR001", false, fmt.Errorf("failed querying Center cities Master: %w", err)
	}
	//log.Println("Circles returned: ", circles)
	return cities, 200, "", true, nil
}
func QueryExamCityNamesForIPExam(client *ent.Client, ctx context.Context, nn string, cid int32, code int32) ([]*ent.ExamCityCenter, int32, string, bool, error) {

	notiNo := nn
	crid := cid
	codeid := code

	cities, err := client.ExamCityCenter.
		Query().
		Where(
			examcitycenter.NotificationNumberEQ(notiNo),
			examcitycenter.ExamYearEQ(crid),
			examcitycenter.ExamCodeEQ(codeid),
			examcitycenter.StatusEQ("active")).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	} else {
		if len(cities) == 0 {
			return nil, 422, " -STR002", false, errors.New(" no active Exam City center found")
		}
	}
	return cities, 200, "", true, nil
}

/* func PostPreferences(client *ent.Client, newExamCity *ent.ExamCityCenter) (*ent.ExamCityCenter, error) {
	ctx := context.Background()
	examid := newExamCity.ExamCode
	examyear := newExamCity.ExamYear
	regioncode := newExamCity.RegionCode
	//divisioncode := newExamCity.DivisionCode
	notificationNumber := newExamCity.NotificationNumber

	if examid <= 0 {
		return nil, fmt.Errorf("invalid Exam Code: %d", examid)
	}

	examcode, err := client.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		log.Println("Error querying ExamCode :", err)
		return nil, fmt.Errorf("No exams found for the given exam Code: %w", err)
	}

	if examcode == nil {
		return nil, fmt.Errorf("Exam Details not found with the Exam Code: %d", examid)
	}

	examdetail, err := client.Exam.Query().Where(exam.IDEQ(examid)).Only(ctx)
	if err != nil {
		fmt.Printf("failed to retrieve data: %v", err)
		return nil, fmt.Errorf("failed to retrieve data: %w", err)
	}

	//var regiondet, divisiondet int32
	var regiondet int32
	if regioncode > 0 {
		regiondet, err = client.RegionMaster.Query().Where(regionmaster.RegionCodeEQ(regioncode)).OnlyID(ctx)
		if err != nil {
			fmt.Printf("failed to retrieve data: %v", err)
			return nil, fmt.Errorf("failed to retrieve data: %w", err)
		}
	}
	/* else if divisioncode > 0 {

		divisiondet, err = client.DivisionMaster.Query().Where(divisionmaster.DivisionCodeEQ(divisioncode)).OnlyID(ctx)
		if err != nil {
			fmt.Printf("failed to retrieve data: %v", err)
			return nil, fmt.Errorf("failed to retrieve data: %w", err)
		}
	}
	regionid := regiondet
	//divisionid := divisiondet
	examname := examdetail.ExamName

	if newExamCity.ConductedBy == "Directorate" || newExamCity.ConductedBy == "directorate" {
		//newExamCity.ExamCode == 1 || newExamCity.ExamCode == 2 {

		exists, err := client.ExamCityCenter.
			Query().
			Where(
				examcitycenter.ExamCodeEQ(newExamCity.ExamCode),
				examcitycenter.NotificationNumberEQ(newExamCity.NotificationNumber),
				examcitycenter.CircleIDEQ(newExamCity.CircleID),
				examcitycenter.ConductedByEQ(newExamCity.ConductedBy)).Exist(ctx)

		if err != nil {
			log.Println("Failed to check existing city names:", err)
			return nil, fmt.Errorf("failed to check existing city names: %v", err)
		}

		if exists {
			return nil, fmt.Errorf("The City name is already added for the Exam")
		}
	} else if newExamCity.ConductedBy == "Circle" || newExamCity.ConductedBy == "circle" {
		//newExamCity.ExamCode == 3 || newExamCity.ExamCode == 4 || newExamCity.ExamCode == 5 || newExamCity.ExamCode == 6 {

		exists, err := client.ExamCityCenter.
			Query().
			Where(
				examcitycenter.ExamCodeEQ(newExamCity.ExamCode),
				examcitycenter.NotificationNumberEQ(newExamCity.NotificationNumber),
				examcitycenter.CircleIDEQ(newExamCity.CircleID),
				examcitycenter.RegionCodeEQ(regionid)).Exist(ctx)

		if err != nil {
			log.Println("Failed to check existing city names:", err)
			return nil, fmt.Errorf("failed to check existing city names: %v", err)
		}

		if exists {
			return nil, fmt.Errorf("The City name is already added for the Exam")
		}

		exists1, err := client.ExamCityCenter.
			Query().
			Where(
				examcitycenter.ExamCodeEQ(newExamCity.ExamCode),
				examcitycenter.NotificationNumberEQ(newExamCity.NotificationNumber),
				examcitycenter.CircleIDEQ(newExamCity.CircleID),
			).Exist(ctx)
		//examcitycenter.DivisionCodeEQ(divisionid)

		if err != nil {
			log.Println("Failed to check existing city names:", err)
			return nil, fmt.Errorf("failed to check existing city names: %v", err)
		}

		if exists1 {
			return nil, fmt.Errorf("The City name is already added for the Exam")
		}

	}

	//currentTime := time.Now().Truncate(time.Second)

	if newExamCity.ConductedBy == "Directorate" || newExamCity.ConductedBy == "directorate" {
		//newExamCity.ExamCode == 1 || newExamCity.ExamCode == 2 {

		currentTime := time.Now().Truncate(time.Second)

		cp, err := client.ExamCityCenter.Create().
			SetExamCode(examid).
			SetExamName(examname).
			SetExamYear(examyear).
			SetNotificationNumber(notificationNumber).
			SetUpdatedAt(currentTime).
			SetConductedBy("Directorate").
			SetCircleID(newExamCity.CircleID).
			//SetRegionCode(newExamCity.RegionCode).
			//SetDivisionCode(newExamCity.DivisionCode).
			//SetCircleName(newExamCity.CircleName).
			SetCenterCityName(newExamCity.CenterCityName).
			//SetRegionName(newExamCity.RegionName).
			//SetDivisionName(newExamCity.DivisionName).
			SetStatus("inactive").
			Save(ctx)
		if err != nil {
			log.Println("error at Creating city preference: ", newExamCity)
			return nil, fmt.Errorf("failed creating city preference: %w", err)
		}

		return cp, nil

	} else if newExamCity.ConductedBy == "Circle" || newExamCity.ConductedBy == "circle" {
		//newExamCity.ExamCode == 3 || newExamCity.ExamCode == 4 || newExamCity.ExamCode == 5 || newExamCity.ExamCode == 6 {

		currentTime := time.Now().Truncate(time.Second)

		if newExamCity.DivisionCode > 0 && newExamCity.DivisionName != "" {
			cp, err := client.ExamCityCenter.Create().
				SetExamCode(examid).
				SetExamName(examname).
				SetExamYear(examyear).
				SetNotificationNumber(notificationNumber).
				SetUpdatedAt(currentTime).
				SetConductedBy("Circle").
				SetCircleID(newExamCity.CircleID).
				//SetRegionCode(newExamCity.RegionCode).
				//SetDivisionCode(divisionid).
				//SetCircleName(newExamCity.CircleName).
				//SetRegionName(newExamCity.RegionName).
				SetDivisionName(newExamCity.DivisionName).
				SetCenterCityName(newExamCity.DivisionName).
				SetStatus("inactive").Save(ctx)
			if err != nil {
				log.Println("error at Creating city preference: ", newExamCity)
				return nil, fmt.Errorf("failed creating city preference: %w", err)
			}
			return cp, nil
		} else if newExamCity.RegionCode > 0 && newExamCity.RegionName != "" {
			cp, err := client.ExamCityCenter.Create().
				SetExamCode(examid).
				SetExamName(examname).
				SetExamYear(examyear).
				SetNotificationNumber(notificationNumber).
				SetUpdatedAt(currentTime).
				SetConductedBy("Circle").
				SetCircleID(newExamCity.CircleID).
				SetRegionCode(regionid).
				//SetDivisionCode(newExamCity.DivisionCode).
				//SetCircleName(newExamCity.CircleName).
				SetRegionName(newExamCity.RegionName).
				SetCenterCityName(newExamCity.RegionName).
				//SetDivisionName(newExamCity.DivisionName).
				SetStatus("inactive").Save(ctx)
			if err != nil {
				log.Println("error at Creating city preference: ", newExamCity)
				return nil, fmt.Errorf("failed creating city preference: %w", err)
			}
			return cp, nil
		}
	}
	return nil, fmt.Errorf("failed creating city preference")
}
*/
/* func CreateExamCityCenters(client *ent.Client, newExamCity *ent.ExamCityCenter) (*ent.ExamCityCenter, int32, error) {
	ctx := context.Background()
	examid := newExamCity.ExamCode
	examyear := newExamCity.ExamYear
	examname := newExamCity.ExamName
	notificationNumber := newExamCity.NotificationNumber
	conductedby := newExamCity.ConductedBy
	centerCityName := newExamCity.CenterCityName
	nodalofficefacilityid := newExamCity.NodalOfficeFacilityID
	nodalofficename := newExamCity.NodalOfficeName
	createdbyusername := newExamCity.CreatedByUserName
	examshortname := newExamCity.ExamShortName

	if examid <= 0 {
		return nil, 422, fmt.Errorf("invalid Exam Code: %d", examid)
	}

	examcode, err := client.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, errors.New("for you no active application  available for this exam ")
		} else {
			return nil, 400, fmt.Errorf("error in fetching application: %v", err)
		}
	}

	if err != nil {
		log.Println("Error querying ExamCode :", err)
		return nil, fmt.Errorf("No exams found for the given exam Code: %w", err)
	}

	if examcode == nil {
		return nil, fmt.Errorf("Exam Details not found with the Exam Code: %d", examid)
	}

	examdetail, err := client.Exam.Query().Where(exam.IDEQ(examid)).Only(ctx)
	if err != nil {
		fmt.Printf("failed to retrieve data: %v", err)
		return nil, fmt.Errorf("failed to retrieve data: %w", err)
	}

	fmt.Println(examdetail)

	if newExamCity.ConductedBy == "Directorate" || newExamCity.ConductedBy == "directorate" {
		//newExamCity.ExamCode == 1 || newExamCity.ExamCode == 2 {
		currentTime := time.Now().Truncate(time.Second)

	exists, err := client.ExamCityCenter.
		Query().
		Where(
			examcitycenter.CenterCityNameEQ(newExamCity.CenterCityName),
			examcitycenter.ExamCodeEQ(newExamCity.ExamCode),
			examcitycenter.NotificationNumberEQ(newExamCity.NotificationNumber),
			examcitycenter.NodalOfficeFacilityIDEQ(newExamCity.NodalOfficeFacilityID),
			examcitycenter.ConductedByEQ(newExamCity.ConductedBy)).
		Exist(ctx)
	if err != nil {
		return nil, 400, fmt.Errorf("failed to check existance of  city names: %v", err)
	}
	if exists {
		return nil, fmt.Errorf("this city name is already added for the Exam")
	}
 	} else if newExamCity.ConductedBy == "Circle" || newExamCity.ConductedBy == "circle" {
		exists, err := client.ExamCityCenter.
			Query().
			Where(
				examcitycenter.CenterCityNameEQ(newExamCity.CenterCityName),
				examcitycenter.ExamCodeEQ(newExamCity.ExamCode),
				examcitycenter.NotificationNumberEQ(newExamCity.NotificationNumber),
				examcitycenter.NodalOfficeFacilityIDEQ(newExamCity.NodalOfficeFacilityID),
				examcitycenter.ConductedByEQ(newExamCity.ConductedBy),
			).Exist(ctx)

		if err != nil {
			log.Println("Failed to check existing city names:", err)
			return nil, fmt.Errorf("failed to check existing city names: %v", err)
		}

		if exists {
			return nil, fmt.Errorf("The City name is already added for the Exam")
		}

	}

	exists, err := client.AdminMaster.
		Query().
		Where(adminmaster.UserNameEQ(createdbyusername)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check username existence: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("no such username: %s", createdbyusername)
	}

 	// Fetch the existing AdminMaster entity
	adminMaster, err := client.AdminMaster.
		Query().
		Where(adminmaster.UserNameEQ(createdbyusername)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch admin master: %v", err)
	}

	cp, err := client.ExamCityCenter.Create().
		SetExamYear(newExamCity.ExamYear).
		SetExamName(newExamCity.ExamName).
		SetExamCode(newExamCity.ExamCode).
		SetExamShortName(newExamCity.ExamShortName).
		SetNotificationNumber(newExamCity.NotificationNumber).
		SetConductedBy(newExamCity.ConductedBy).
		SetCenterCityName(newExamCity.CenterCityName).
		SetNodalOfficeFacilityID(newExamCity.NodalOfficeFacilityID).
		SetNodalOfficeName(newExamCity.NodalOfficeName).
		SetCreatedByUserName(newExamCity.CreatedByUserName).
		SetUpdatedAt(currentTime).
		SetCreatedById(newExamCity.CreatedById).
		SetCreatedByEmpId(newExamCity.CreatedByEmpId).
		SetCreatedByDesignation(newExamCity.Designation).
		SetStatus("active").
		Save(ctx)
	if err != nil {
		return nil, 400, fmt.Errorf("failed creating city preference: %w", err)
	}

	return cp, 200, nil

}
*/

func CreateExamCityCenters(client *ent.Client, newExamCity *ent.ExamCityCenter) (*ent.ExamCityCenter, int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	currentTime := time.Now().Truncate(time.Second)
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, fmt.Errorf("failed to start transaction: %w", err)
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

	exists, err := tx.ExamCityCenter.
		Query().
		Where(
			examcitycenter.CenterCityNameEQ(newExamCity.CenterCityName),
			examcitycenter.ExamCodeEQ(newExamCity.ExamCode),
			examcitycenter.NotificationNumberEQ(newExamCity.NotificationNumber),
			examcitycenter.NodalOfficeFacilityIDEQ(newExamCity.NodalOfficeFacilityID),
			examcitycenter.ConductedByEQ(newExamCity.ConductedBy)).
		Exist(ctx)
	if err != nil {
		return nil, 400, fmt.Errorf("failed to check existance of  city names: %v", err)
	}
	if exists {
		return nil, 422, errors.New("this city name is already added for the Exam")
	}

	cp, err := tx.ExamCityCenter.Create().
		SetExamYear(newExamCity.ExamYear).
		SetExamName(newExamCity.ExamName).
		SetExamCode(newExamCity.ExamCode).
		SetExamShortName(newExamCity.ExamShortName).
		SetNotificationNumber(newExamCity.NotificationNumber).
		SetConductedBy(newExamCity.ConductedBy).
		SetCenterCityName(newExamCity.CenterCityName).
		SetNodalOfficeFacilityID(newExamCity.NodalOfficeFacilityID).
		SetNodalOfficeName(newExamCity.NodalOfficeName).
		SetCreatedByUserName(newExamCity.CreatedByUserName).
		SetUpdatedAt(currentTime).
		SetCreatedById(newExamCity.CreatedById).
		SetCreatedByEmpId(newExamCity.CreatedByEmpId).
		SetCreatedByDesignation(newExamCity.CreatedByDesignation).
		SetStatus("active").
		Save(ctx)
	if err != nil {
		return nil, 400, fmt.Errorf("failed creating city preference: %w", err)
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, err
	}
	return cp, 200, nil
}

func UpdateExamCentresMTSPMMGExamsreturnarray(ctx context.Context, client *ent.Client, newappls []*ent.Exam_Application_MTSPMMG) ([]UpdateResult, int32, string, bool, error) {
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
		exists, err := tx.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ReportingOfficeIDEQ(newappl.ReportingOfficeID),
				exam_application_mtspmmg.Status("active"),
				exam_application_mtspmmg.ExamYearEQ(newappl.ExamYear),
			).
			Exist(ctx)
		if err != nil {
			//Failed to check ReportingOfficeID existence
			return nil, 500, " -STR001 " + newappl.ReportingOfficeFacilityID, false, err
		}
		if !exists {
			//The ReportingOfficeID %s does not exist in the Applications table. Skipping to the next value in the loop.loop.\n", newappl.ReportingOfficeID)
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
		centerExists, err := tx.Center.
			Query().
			Where(center.IDEQ(newappl.ExamCityCenterCode)).
			Exist(ctx)
		if err != nil {
			return nil, 500, " -STR002", false, err
		}
		if !centerExists {
			//The CenterCode %d does not exist in the center.master table. Skipping to the next value in the loop
			updateResult := UpdateResult{
				ReportingOfficeID: newappl.ReportingOfficeID,
				CenterCode:        newappl.ExamCityCenterCode,
				Message:           "CenterCode does not exist and skipped",
				//RecordCount:       0,
			}
			updateResults = append(updateResults, updateResult)
			continue
		}

		applications, err := tx.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.ReportingOfficeIDEQ(newappl.ReportingOfficeID),
				exam_application_mtspmmg.StatusEQ("active"),
			).
			Order(ent.Desc(exam_application_mtspmmg.FieldID)).
			All(ctx)
		if err != nil {
			return nil, 500, " -STR003", false, err
		}
		count := len(applications)
		if count == 0 {
			return nil, 400, " -STR004", false, errors.New(" no reporting office found")
		}
		if count > 0 {
			// Update the CenterCode for each application record
			for _, application := range applications {
				application.ExamCityCenterCode = newappl.ExamCityCenterCode
				_, err = application.Update().SetExamCityCenterCode(newappl.ExamCityCenterCode).Save(ctx)
				if err != nil {
					return nil, 500, " -STR005", false, err
				}
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
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR006", false, err
	}
	return updateResults, 200, "", true, nil
}

func UpdateCenter(client *ent.Client, id int32, newCenter ca_reg.CenterRequest) (*ent.Center, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Check if id is zero
	if id == 0 {
		return nil, 422, " -STR001", false, fmt.Errorf("CenterID is mandatory for update")
	}
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

	examdetail, err := tx.Exam.Query().Where(exam.IDEQ(newCenter.ExamCode)).Only(ctx)
	if err != nil {
		fmt.Printf("failed to retrieve data: %v", err)
		return nil, 422, " -STR002", false, fmt.Errorf("failed to retrieve data: %w", err)
	}

	examname := examdetail.ExamName

	// Check if the center exists and its status is true/*
	existingCenter, _, _, _, err := QueryCenterID(ctx, client, id)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}
	if existingCenter == nil {
		return nil, 422, " -STR004", false, fmt.Errorf(" no such exam center exists")
	}

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		// Handle the error, such as logging or returning an error
		log.Printf("Error loading location: %v", err)
		return nil, 500, " -STR004", false, err
	}

	currentTime := time.Now().In(loc).Truncate(time.Second)

	updatedCenter, err := tx.Center.UpdateOneID(id).
		//SetNotifyCode(newCenter.NotifyCode).
		//SetNodalOfficerCode(newCenter.NodalOfficerCode).
		SetExamCenterName(newCenter.ExamCenterName).
		SetNodalOfficerCode(newCenter.NodalOfficerCode).
		SetExamCode(newCenter.ExamCode).
		SetRegionID(newCenter.RegionID).
		SetCircleID(newCenter.CircleID).
		SetDivisionID(newCenter.DivisionID).
		SetFacilityID(newCenter.FacilityID).
		//SetExamNameCode(newCenter.ExamNameCode).
		SetExamName(examname).
		SetNAUserName(newCenter.NAUserName).
		SetNodalOfficeFacilityId(newCenter.NodalOfficeFacilityId).
		SetAdminCircleOfficeID(newCenter.AdminCircleOfficeID).
		SetAddress(newCenter.Address).
		SetLandmark(newCenter.Landmark).
		SetPincode(newCenter.Pincode).
		SetMaxSeats(newCenter.MaxSeats).
		SetNoAlloted(newCenter.NoAlloted).
		SetPendingSeats(newCenter.PendingSeats).
		SetStatus(newCenter.Status).
		SetUpdatedAt(currentTime).
		SetUpdatedBy(newCenter.UpdatedBy).
		SetCenterCityName(newCenter.CenterCityName).
		SetPapers(*newCenter.Papers).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR027", false, err
	}

	return updatedCenter, 200, "", true, nil
}
