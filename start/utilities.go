package start

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"recruit/ent"
	"recruit/ent/errorlogs"
	"recruit/ent/examcategorydisabilitymapping"
	"recruit/ent/exampostmapping"
	"recruit/ent/facilitymasters"
	"recruit/ent/message"
	"recruit/ent/servicerequest"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/sms"
	"recruit/util"
)

func SubMessage(ctx context.Context, client *ent.Client) ([]*ent.Message, int32, string, bool, error) {
	message, err := client.Message.Query().
		Where(
			message.StatusEQ(true),
		).
		Order(ent.Asc(message.FieldPriority)).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	} else {
		if len(message) == 0 {
			return nil, 422, " -STR002", false, errors.New("no message details exists")
		}
	}
	return message, 200, "", true, nil
}

type VersionStruct struct {
	UiVersion        string
	ApiVersion       string
	EnvUiVersion     string
	EnvApiVersion    string
	UiVersionStatus  bool
	APIVersionStatus bool
}

func SubVersion(ctx context.Context, client *ent.Client, uiversion string, apiversion string) ([]VersionStruct, int32, string, bool, error) {
	version, err := client.Version.Query().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	} else {
		if len(version) == 0 {
			return nil, 422, " -STR002", false, errors.New("version details missing")
		}
	}

	//var versionstruct []VersionStruct
	var message string = "current "
	var uiversionstatus, apiversionstatus bool = true, true

	if version[0].ApiVersion != apiversion {
		message += " API Version is " + version[0].ApiVersion + " , "
		apiversionstatus = false
	}

	if version[0].UiVersion != uiversion {
		message += " UI Version is " + version[0].ApiVersion
		uiversionstatus = false
	}

	if message == "current " {
		message = "UI and API Verision matched"
	}
	var versionstruct []VersionStruct
	versionstruct = append(versionstruct, VersionStruct{
		UiVersion:        version[0].UiVersion,
		ApiVersion:       version[0].ApiVersion,
		EnvUiVersion:     uiversion,
		EnvApiVersion:    apiversion,
		UiVersionStatus:  uiversionstatus,
		APIVersionStatus: apiversionstatus,
	})

	return versionstruct, 200, "", true, errors.New(message)
}

func QueryDetailsByExamCode(ctx context.Context, client *ent.Client, id int64) ([]map[string]interface{}, []map[string]interface{}, int32, string, bool, error) {
	categoryData, status, stgError, _, err := QueryCategoryData(ctx, client, id)
	if err != nil {
		return nil, nil, status, " - STR001" + stgError, false, err
	}

	postMappingData, status, stgError, _, err := QueryPostMappingData(ctx, client, id)
	if err != nil {
		return nil, nil, status, " - STR002" + stgError, false, err
	}

	return categoryData, postMappingData, 200, "", true, nil
}

func QueryCategoryData(ctx context.Context, client *ent.Client, id int64) ([]map[string]interface{}, int32, string, bool, error) {
	mappings, err := client.ExamCategoryDisabilityMapping.Query().
		Where(
			examcategorydisabilitymapping.ExamCodeEQ(id),
			examcategorydisabilitymapping.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -SUB001", false, err
	} else {
		if len(mappings) == 0 {
			return nil, 422, " -SUB002", false, errors.New("no matching category found")
		}
	}

	response := make([]map[string]interface{}, 0) // Response should be a slice of maps

	categories := make([]map[string]interface{}, 0)
	disabilities := make([]map[string]interface{}, 0)

	for _, category := range mappings {
		categoryData := map[string]interface{}{
			"categoryCode":        category.CategoryDisabilityCode,
			"categoryDescription": category.CategoryDisabilityDescription,
			"ageExemption":        category.AgeException,
			"serviceExemption":    category.ServiceException,
			"drivingLicense":      category.DrivingLicense,
		}

		if category.CategoryDisability == "category" {
			categories = append(categories, categoryData)
		} else if category.CategoryDisability == "disability" {
			disabilities = append(disabilities, categoryData)
		}
	}
	response = append(response, map[string]interface{}{
		"data": []map[string]interface{}{
			{
				"category": categories,
			},
			{
				"disability": disabilities,
			},
		},
	})
	return response, 200, "", true, nil
}
func QueryPostMappingData(ctx context.Context, client *ent.Client, id int64) ([]map[string]interface{}, int32, string, bool, error) {
	mappings, err := client.ExamPostMapping.Query().
		Where(
			exampostmapping.ExamCodeEQ(id),
			exampostmapping.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR006", false, err
	} else {
		if len(mappings) == 0 {
			return nil, 422, " -STR006", false, errors.New("no matching post preference found")
		}
	}

	response := make([]map[string]interface{}, 0) // Response should be a slice of maps

	entry := make([]map[string]interface{}, 0)
	present := make([]map[string]interface{}, 0)
	feder := make([]map[string]interface{}, 0)

	for _, category1 := range mappings {
		categoryData1 := map[string]interface{}{
			"PostID":            category1.FromPostCode,
			"PostdDescription":  category1.FromPostDescription,
			"ageCriteria":       category1.AgeCriteria,
			"serviceCriteria":   category1.ServiceCriteria,
			"ToPostDescription": category1.ToPostDescription,
			"ToPostCode":        category1.ToPostCode,
		}

		if category1.PostTypeDescription == "entry cader" {
			entry = append(entry, categoryData1)
		} else if category1.PostTypeDescription == "present cader" {
			present = append(present, categoryData1)
		} else if category1.PostTypeDescription == "feder cader" {
			feder = append(feder, categoryData1)
		}
	}

	response = append(response, map[string]interface{}{
		"data": []map[string]interface{}{
			{
				"entry cader": entry,
			},
			{
				"present cader": present,
			},
			{
				"feder cader": feder,
			},
		},
	})

	return response, 200, "", true, nil
}

func SubGetAllCircles(ctx context.Context, client *ent.Client) ([]*ent.FacilityMasters, int32, string, bool, error) {

	// Fetch all applications for the given facilityID
	circles, err := client.FacilityMasters.
		Query().
		Where(facilitymasters.FacilityTypeEQ("CR"),
			facilitymasters.StatusEQ("active")).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}

	if len(circles) > 0 {
		return circles, 200, "", true, nil
	} else {
		return nil, 422, " -STR002", false, errors.New("no office matched with this PIN Code")
	}
}

func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

/*
	 func SubErrorLogAssignment(client *ent.Client, ID int64, errorLog *ent.ErrorLogs) (*ent.ErrorLogs, int32, string, bool, error) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		if ID <= 0 {
			return nil, 422, "STR001", false, errors.New("id should be greater than zero")
		}

		tx, err := client.Tx(ctx)
		if err != nil {
			return nil, 500, "TX001", false, err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()

		errorLogEntry, err := tx.ErrorLogs.Query().
			Where(errorlogs.IDEQ(ID)).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, 422, "STR002", false, errors.New("this id does not exist")
			}
			return nil, 500, "STR003", false, err
		}
		currentTime := time.Now().Truncate(time.Second)
		var updatedErrorLog *ent.ErrorLogs
		if errorLog.Status == "closed" {
			updatedErrorLog, err = errorLogEntry.Update().
				SetUpdatedBy(errorLog.UpdatedBy).
				SetUpdatedTime(currentTime).
				SetAssignedTo(errorLog.AssignedTo).
				SetRemarksNew(errorLog.RemarksNew).
				SetStatus(errorLog.Status).
				SetClosedOn(currentTime).
				Save(ctx)
			if err != nil {
				return nil, 500, "STR005", false, err
			}
		} else {
			updatedErrorLog, err = errorLogEntry.Update().
				SetUpdatedBy(errorLog.UpdatedBy).
				SetUpdatedTime(currentTime).
				SetAssignedTo(errorLog.AssignedTo).
				Save(ctx)
			if err != nil {
				return nil, 500, "STR006", false, err
			}
		}

		if err = tx.Commit(); err != nil {
			tx.Rollback()
			return nil, 500, "STR007", false, err
		}
		return updatedErrorLog, 200, "", true, nil
	}
*/
type StrucSericeRequest struct {
	ID         int    `json:"id" binding:"required"`
	UpdatedBy  string `json:"updated_by" binding:"required"`
	AssignedTo string `json:"assigned_to"`
	RemarksNew string `json:"remarks_new"`
	Status     string `json:"status"`
}

func SubErrorLogAssignment(client *ent.Client, serviceRequest *ca_reg.StrucSericeRequest) (*ent.ErrorLogs, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	if serviceRequest.ID <= 0 {
		return nil, 422, "STR001", false, errors.New("id should be greater than zero")
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, "TX001", false, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	currentTime := time.Now().Truncate(time.Second)

	exists, err := tx.ServiceRequest.Query().
		Where(
			servicerequest.ID(serviceRequest.ID),
			servicerequest.StatusNEQ("closed"),
		).
		Exist(ctx)
	if err != nil {
		return nil, 500, "STR003", false, err
	}

	if exists {
		// Insert into ServiceRequest if not already present
		if serviceRequest.Status == "closed" {
			_, err = tx.ServiceRequest.Update().
				SetUpdatedBy(serviceRequest.UpdatedBy).
				SetUpdatedTime(currentTime).
				SetRemarksNew(serviceRequest.RemarksNew).
				SetStatus(serviceRequest.Status).
				SetClosedOn(currentTime).
				Save(ctx)
			if err != nil {
				return nil, 500, "STR005", false, err
			}
		} else {
			_, err = tx.ServiceRequest.Update().
				SetUpdatedBy(serviceRequest.UpdatedBy).
				SetUpdatedTime(currentTime).
				SetAssignedTo(serviceRequest.AssignedTo).
				Save(ctx)
			if err != nil {
				return nil, 500, "STR006", false, err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, "STR007", false, err
	}
	return nil, 200, "", true, nil
}

type StrucGetServiceRequestLog struct {
	ID          int       `json:"id"`
	Remarks     string    `json:"remarks"`
	Action      string    `json:"action"`
	PushedTime  time.Time `json:"pushed_time"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedTime time.Time `json:"updated_time"`
	AssignedTo  string    `json:"assigned_to"`
	RemarksNew  string    `json:"remarks_new"`
	Status      string    `json:"status"`
	ClosedOn    time.Time `json:"closed_on"`
}

type InsertServiceRequest struct {
	Remarks string `json:"remarks"`
	Action  string `json:"action"`
}

func SubGetErrorLogs(ctx context.Context, client *ent.Client, adminUserName string) ([]StrucGetServiceRequestLog, int32, string, bool, error) {
	var serviceRequestResults []StrucGetServiceRequestLog

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, "TX001", false, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var insertServiceRequest []InsertServiceRequest
	err = tx.ErrorLogs.Query().
		Where(
			errorlogs.StatusIsNil(),
			errorlogs.ActionNEQ("422"),
		).
		GroupBy(
			errorlogs.FieldRemarks, errorlogs.FieldAction).
		Scan(ctx, &insertServiceRequest)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	fmt.Println(insertServiceRequest)
	currentTime := time.Now().Truncate(time.Second)
	var checkError bool = false
	for _, logEntry := range insertServiceRequest {
		// Check if the remark from ErrorLogsTable is not present in ServiceRequest
		exists, err := tx.ServiceRequest.Query().
			Where(servicerequest.Remarks(logEntry.Remarks),
				servicerequest.StatusNEQ("closed")).
			Exist(ctx)
		if err != nil {
			checkError = true
			break // Exit the loop if an error occurs
		}
		fmt.Println("Exists", exists)
		if !exists {
			// Insert into ServiceRequest if not already present
			_, err := tx.ServiceRequest.Create().
				SetRemarks(logEntry.Remarks).
				SetAction(logEntry.Action).
				SetPushedTime(currentTime).
				Save(ctx)
			if err != nil {
				checkError = true
				break // Exit the loop if an error occurs
			}
		}
		fmt.Println("Service request update")
		_, err = tx.ErrorLogs.Update().
			SetUpdatedBy(adminUserName).
			SetUpdatedTime(currentTime).
			SetStatus("transferred").
			SetClosedOn(currentTime).
			Save(ctx)
		if err != nil {
			checkError = true
			break // Exit the loop if an error occurs
		}
		fmt.Println("Error log update")

	}
	if checkError {
		return nil, 500, "STR003", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, "STR007", false, err
	}
	//return updatedErrorLog, 200, "", true, nil

	err = client.ServiceRequest.Query().
		Where(
			servicerequest.StatusIsNil(),
		).
		Select(
			servicerequest.FieldID,
			servicerequest.FieldRemarks,
			servicerequest.FieldAction,
			servicerequest.FieldPushedTime,
			servicerequest.FieldUpdatedBy,
			servicerequest.FieldUpdatedTime,
			servicerequest.FieldAssignedTo,
			servicerequest.FieldRemarksNew,
			servicerequest.FieldStatus,
			servicerequest.FieldClosedOn,
		).
		Scan(ctx, &serviceRequestResults)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	fmt.Println("serviceRequestResults", serviceRequestResults)
	return serviceRequestResults, 200, "", true, nil
}

func SubTestSMS(ctx context.Context, client *ent.Client, napitype string) (string, string) {
	//var examID int32
	apitype, _ := strconv.ParseInt(napitype, 10, 64)
	var apiresponse string = ""
	var apiresponsedescription string = ""
	var statussms int32 = 0
	//var napitype string = ""
	var url string = ""
	var vString string = ""
	appName := "DOPExam"
	msg := "Dear Testing-" + napitype + ", User ID was successfully registered. For details check email-DOPExam-INDIAPOST"
	templateID := "1007033293383430282"
	entityID := "1001081725895192800"
	mobile := "7299174555"
	var err error
	switch {
	case apitype == 1:
		statussms, err = sms.SendSMS1(msg, mobile, templateID, entityID, appName)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-2"
	case apitype == 2:

		url = "https://uat.cept.gov.in/sms/v1/msgrequest/create"
		statussms, err = sms.SendSMS2(msg, mobile, templateID, entityID, url)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-2"
	case apitype == 3:
		url = "https://apiservices.cept.gov.in/bemsggateway/v1/msgrequest/create"
		statussms, err = sms.SendSMS2(msg, mobile, templateID, entityID, url)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-2"
	case apitype == 4:
		url = "https://dev.cept.gov.in/bemsggateway/v1/msgrequest/create"
		statussms, err = sms.SendSMS2(msg, mobile, templateID, entityID, url)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-2"
	case apitype == 5:
		url = "https://test.cept.gov.in/bemsggateway/v1/msgrequest/create"
		statussms, err = sms.SendSMS2(msg, mobile, templateID, entityID, url)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-2"

	default:
		Action = "400"
		Remarks = "Invalid SMS API call type"
		util.SystemLogError(client, Action, Remarks)
		return "Failed", " Invalid SMS  API call type"
	}

	if err != nil {
		if statussms == 400 || statussms == 422 {
			apiresponse = "Failed"
			apiresponsedescription = err.Error()
		} else {
			statussms = 500
			apiresponse = "Failed"
			apiresponsedescription = err.Error()

		}
	} else if statussms == 200 {
		apiresponse = "Success"
		apiresponsedescription = "SMS sent successfully"
	}

	sms.SmsEmailLog(ctx, client, "SMS", mobile, "Testing", vString, templateID, apiresponse, apiresponsedescription)
	return apiresponse, apiresponsedescription

}
