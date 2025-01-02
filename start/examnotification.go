package start

import (
	"context"
	"errors"
	"fmt"
	"log"

	//"net/http"

	//"net/http"
	"recruit/ent"
	"recruit/ent/adminmaster"
	"recruit/ent/exam"
	"recruit/ent/examnotifications"
	"recruit/ent/facilitymasters"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"

	// "recruit/ent/usermaster"
	"time"
	//"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
)

// Function to create Draft Notification
func CreateExamNotification(client *ent.Client, newNotification *ca_reg.CreateNotificationStruct) (*ent.ExamNotifications, int32, string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examid := newNotification.ExamCode
	username := newNotification.UserName
	facilityid := newNotification.CircleOfficeFacilityId
	officetype := newNotification.IssuedBy

	if examid <= 0 {
		return nil, 422, " -STR001", false, fmt.Errorf("invalid Exam Code: %d", examid)
	}

	if (examid == 1 || examid == 2) && officetype == "Circle" {
		return nil, 422, " -STR002", false, fmt.Errorf(" Circles are not allowed to issue notification of IP and PS Group Exam")
	}

	examdetail, err := client.Exam.
		Query().
		Where(exam.ExamCodeEQ(examid)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR003", false, fmt.Errorf(" no such exam available ")
		} else {
			return nil, 500, " -STR004", false, err
		}
	}

	examname := examdetail.ExamName

	exists, err := client.ExamNotifications.
		Query().
		Where(
			examnotifications.ExamCodeEQ(newNotification.ExamCode),
			examnotifications.ExamYearEQ(newNotification.ExamYear),
			examnotifications.IssuedByEQ(officetype),
			examnotifications.CircleOfficeFacilityIdEQ(facilityid),
			examnotifications.StatusEQ("active"),
			examnotifications.NotificationStatusIn("NotificationIssued", "NotificationSubmitted", "NotificationResubmitted", "NotificationReturned", "NotificationHold")).
		Exist(ctx)

	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	if exists {
		return nil, 422, " -STR006", false, errors.New(" already notification available with   Notification Issued / Notification Submitted / Notification Resubmitted / Notification Returned / Notification Hold is already available for the Exam and exam year")
	}

	currentTime := time.Now().Truncate(time.Second)

	// Start transaction
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

	dn, err := tx.ExamNotifications.Create().
		SetExamCode(newNotification.ExamCode).
		SetExamName(examname).
		SetUserName(username).
		SetCircleOfficeFacilityId(newNotification.CircleOfficeFacilityId).
		SetCircleOfficeName(newNotification.CircleOfficeName).
		SetIssuedBy(newNotification.IssuedBy).
		SetExamYear(newNotification.ExamYear).
		SetEmployeeMasterRequestLastDate(newNotification.EmployeeMasterRequestLastDate).
		SetEmployeeMasterRequestApprovalLastDate(newNotification.EmployeeMasterRequestApprovalLastDate).
		SetExamRegisterLastDate(newNotification.ExamRegisterLastDate).
		SetApplicationStartDate(newNotification.ApplicationStartDate).
		SetApplicationEndDate(newNotification.ApplicationEndDate).
		SetApplicationCorrectionLastDate(newNotification.ApplicationCorrectionLastDate).
		SetApplicationVerificationLastDate(newNotification.ApplicationVerificationLastDate).
		SetNodalOfficerApprovalDate(newNotification.NodalOfficerApprovalDate).
		SetAdmitCardDate(newNotification.AdmitCardDate).
		SetCrucialDate(*newNotification.CrucialDate).
		SetSmsExamShortName(examdetail.SmsExamShortName).
		SetResubmittedApplicationVerificationDate(newNotification.ResubmittedApplicationVerificationDate).
		SetPapers(*newNotification.Papers).
		SetNotificationStatus(newNotification.NotificationStatus).
		SetNotificationOrderNumber(newNotification.NotificationOrderNumber).
		SetOrderDate(newNotification.OrderDate).
		SetCenterAllotmentEndDate(newNotification.CenterAllotmentEndDate).
		SetUpdatedAt(currentTime).
		SetExamShortName(examdetail.ExamShortName).
		SetCreatedById(newNotification.CreatedById).
		SetCreatedBy(newNotification.CreatedBy). // Set the current UTC time
		SetCreatedByName(newNotification.CreatedByName).
		SetCreatedByDesignation(newNotification.CreatedByDesignation).
		SetUpdatedBy(newNotification.UserName).
		SetStatus("active").

		// Set the updated by user (replace with the actual user)
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR007", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR008", false, err
	}
	return dn, 200, "", true, nil
}

// Function to check whether notification is issued based on ExamId and Exam year.
func GetDraftNotification(client *ent.Client, intExamId int32, intExamYear int32, officeType string, facilityId string) ([]*ent.ExamNotifications, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examid := intExamId
	examyear := intExamYear
	//issuedNotification := make([]*ent.ExamNotifications, 0)

	if examid <= 0 {
		return nil, 422, " -STR001", false, errors.New("invalid Exam Code")
	}

	_, err := client.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR002", false, fmt.Errorf(" no exams found for the given exam Code: %v", err)
		} else {
			return nil, 500, " -STR003", false, err
		}
	}

	notifications, err := client.ExamNotifications.
		Query().Where(
		//examnotifications.FlagNEQ(true),
		examnotifications.ExamCodeEQ(examid),
		examnotifications.IssuedBy(officeType),
		examnotifications.CircleOfficeFacilityIdEQ(facilityId),
		examnotifications.StatusEQ("active"), examnotifications.
			ExamYearEQ(examyear),
		examnotifications.NotificationStatusIn("NotificationIssued", "NotificationSubmitted", "NotificationResubmitted", "NotificationReturned", "NotificationHold")).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR004", false, err
	}

	if len(notifications) == 0 {
		return nil, 422, " -STR005", false, errors.New("no notification exists with this criteria ")
	}

	return notifications, 200, "", true, nil
}

// Function to retrieve all the pending draft notifications based on the username
func GetexamNotification(client *ent.Client, un string, newNotification *ent.ExamNotifications) ([]*ent.ExamNotifications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	username := un
	//PendingNotification := make([]*ent.ExamNotifications, 0)

	pn, err := client.ExamNotifications.
		Query().
		Where(
			examnotifications.NotificationStatusIn("NotificationSubmitted", "NotificationResubmitted"),
			examnotifications.UserNameEQ(username),
			//examnotifications.FlagNEQ(true),
			examnotifications.StatusEQ("active")).
		All(ctx)

	if err != nil {
		log.Println("Failed to retrieve exam notifications:", err)
		return nil, fmt.Errorf("failed to retrieve exam notifications: %v", err)
	}

	return append([]*ent.ExamNotifications{}, pn...), nil

}

func GetReexamNotification(client *ent.Client, un string, newNotification *ent.ExamNotifications) ([]*ent.ExamNotifications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	username := un
	//PendingNotification := make([]*ent.ExamNotifications, 0)

	pn, err := client.ExamNotifications.
		Query().
		Where(
			//examnotifications.FlagNEQ(true),
			examnotifications.NotificationStatusEQ("NotificationReturned"),
			examnotifications.UserNameEQ(username),
			examnotifications.StatusEQ("active")).
		All(ctx)

	if err != nil {
		log.Println("Failed to retrieve exam notifications:", err)
		return nil, fmt.Errorf("failed to retrieve exam notifications: %v", err)
	}

	/* for _, pendingnotification := range pn {
		PendingNotification = append(PendingNotification, pendingnotification)
	}

	return PendingNotification, nil */
	return append([]*ent.ExamNotifications{}, pn...), nil
}

// Function to Display details of individual Notification based on Notiffication ID
func GetIndividualDraftNotificationbySuUserName(client *ent.Client, nid int32) ([]*ent.ExamNotifications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	notificationID := nid
	//ipn := make([]*ent.ExamNotifications, 0)

	if notificationID <= 0 {
		return nil, fmt.Errorf(" invalid Notification ID")
	}

	_, err := client.ExamNotifications.
		Query().
		Where(
			examnotifications.IDEQ(notificationID),
			//examnotifications.FlagNEQ(true),
			examnotifications.StatusEQ("active")).
		Only(ctx)

	if err != nil {
		log.Println("Error querying Notification ID :", err)
		return nil, fmt.Errorf(" notification ID not found : %d", notificationID)
	}

	ipn, nerr := client.ExamNotifications.
		Query().
		Where(
			examnotifications.IDEQ(notificationID),
			//examnotifications.FlagNEQ(true),
			examnotifications.StatusEQ("active")).
		All(ctx)

	if nerr != nil {
		log.Println("Failed to retrieve exam notifications:", nerr)
		return nil, fmt.Errorf("failed to retrieve exam notifications: %v", nerr)
	}
	return ipn, nil
}

// function to retrieve remarks of the Notification.
func GetRemarksofNotificationByNotificationID(ctx context.Context, client *ent.Client, nid int32) ([]*ent.ExamNotifications, error) {
	//ctx := context.Background()
	notificationID := nid
	//ipn := make([]*ent.ExamNotifications, 0)

	if notificationID <= 0 {
		return nil, fmt.Errorf(" invalid Notification ID")
	}

	_, err := client.ExamNotifications.
		Query().
		Where(examnotifications.IDEQ(notificationID)).
		Only(ctx)

	if err != nil {
		log.Println("Error querying Notification ID :", err)
		return nil, fmt.Errorf(" notification ID not found : %d", notificationID)
	}

	ipn, nerr := client.ExamNotifications.
		Query().
		Where(
			examnotifications.IDEQ(notificationID),
			examnotifications.StatusEQ("active"),
		//examnotifications.FlagNEQ(true)
		).
		All(ctx)

	if nerr != nil {
		log.Println("Failed to retrieve exam notifications:", nerr)
		return nil, fmt.Errorf("failed to retrieve exam notifications: %v", nerr)
	}
	return ipn, nil
}

// Function to retrieve all the Issued notifications based on the exam year and exam code
func GetexamNotificationByYear(client *ent.Client, eyear, eid int32) ([]*ent.ExamNotifications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examyear := eyear
	examid := eid

	//IssuedNotification := make([]*ent.ExamNotifications, 0)

	if examyear == 0 {
		return nil, fmt.Errorf(" year cannot be Zero")
	}

	if examid <= 0 {
		return nil, fmt.Errorf(" invalid Exam Code")
	}

	_, err := client.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		log.Println("Error querying ExamCode :", err)
		return nil, fmt.Errorf(" no exams found for the given exam Code: %d", eid)
	}

	in, err := client.ExamNotifications.
		Query().
		Where(
			examnotifications.NotificationStatusEQ("NotificationIssued"),
			examnotifications.StatusEQ("active"),
			examnotifications.ExamYear(examyear),
			examnotifications.ExamCodeEQ(examid),
		//examnotifications.FlagNEQ(true)
		).
		All(ctx)

	if err != nil {
		log.Println(" failed to retrieve exam notifications:", err)
		return nil, fmt.Errorf("failed to retrieve exam notifications: %v", err)
	}

	/* for _, issuedNotification := range in {
		IssuedNotification = append(IssuedNotification, issuedNotification)
	}

	return IssuedNotification, nil */
	return append([]*ent.ExamNotifications{}, in...), nil
}

// Function to retrieve all the Issued notifications based on the User name
func GetallIssuedexamNotificationBySuUserName(client *ent.Client, un string) ([]*ent.ExamNotifications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	username := un
	//IssuedNotification := make([]*ent.ExamNotifications, 0)

	if username == "" {
		return nil, fmt.Errorf(" user Name cannot be Blank")
	}

	in, err := client.ExamNotifications.
		Query().
		Where(
			//examnotifications.FlagNEQ(true),
			examnotifications.NotificationStatusEQ("NotificationIssued"),
			examnotifications.UserNameEQ(username),
			examnotifications.StatusEQ("active")).
		All(ctx)

	if err != nil {
		log.Println("Failed to retrieve exam notifications:", err)
		return nil, fmt.Errorf("failed to retrieve exam notifications: %v", err)
	}

	/* for _, issuedNotification := range in {
		IssuedNotification = append(IssuedNotification, issuedNotification)
	}

	return IssuedNotification, nil */
	return append([]*ent.ExamNotifications{}, in...), nil
}

// Function to retrieve individual Issued notifications based on the notification number
func GetIndividualIssuedexamNotificationByNotificationNumber(client *ent.Client, Notinum string) ([]*ent.ExamNotifications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	notificationnumber := Notinum
	//iin := make([]*ent.ExamNotifications, 0)

	if notificationnumber == "" {
		return nil, fmt.Errorf(" invalid Notification Number")
	}

	_, err := client.ExamNotifications.
		Query().
		Where(
			examnotifications.NotificationNumberEQ(notificationnumber),
			//examnotifications.FlagNEQ(true),
			examnotifications.StatusEQ("active")).
		All(ctx)

	if err != nil {
		log.Println("Error querying Notification Number :", err)
		return nil, fmt.Errorf(" notification Number not found : %s", notificationnumber)
	}

	iin, nerr := client.ExamNotifications.
		Query().
		Where(examnotifications.NotificationNumberEQ(notificationnumber), examnotifications.NotificationStatusEQ("NotificationIssued"), examnotifications.StatusEQ("active")).
		All(ctx)

	if nerr != nil {
		log.Println("Failed to retrieve exam notifications:", nerr)
		return nil, fmt.Errorf("failed to retrieve exam notifications: %v", nerr)
	}
	return iin, nil
}

// Function to retrieve all the Issued notifications for a particular year
func GetAllIssuedexamNotificationbyYear(client *ent.Client, year int32) ([]*ent.ExamNotifications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examyear := year

	//IssuedNotification := make([]*ent.ExamNotifications, 0)

	if examyear == 0 {
		return nil, fmt.Errorf(" year cannot be Zero")
	}

	in, err := client.ExamNotifications.
		Query().
		Where(
			examnotifications.NotificationStatusEQ("NotificationIssued"),
			examnotifications.ExamYear(examyear),
			//examnotifications.FlagNEQ(true),
			examnotifications.StatusEQ("active")).
		All(ctx)

	if err != nil {
		log.Println("Failed to retrieve exam notifications:", err)
		return nil, fmt.Errorf("failed to retrieve exam notifications: %v", err)
	}

	/* 	for _, issuedNotification := range in {
	   		IssuedNotification = append(IssuedNotification, issuedNotification)
	   	}

	   	return IssuedNotification, nil */
	return append([]*ent.ExamNotifications{}, in...), nil
}

// GET ALL NOTIFICATION UPDATEDAPI
func GetAllexamNotificationbyYearUpdated(client *ent.Client, examyear int32, circleid string) ([]*ent.ExamNotifications, int32, string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	IssuedNotification := make([]*ent.ExamNotifications, 0)

	if examyear == 0 {
		return nil, 422, "STR001", false, errors.New(" Exam year cannot be Zero")
	}

	in, err := client.ExamNotifications.
		Query().
		Where(examnotifications.ExamYear(examyear),
			examnotifications.CircleOfficeFacilityId(circleid),
			examnotifications.StatusEQ("active")).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	if len(in) == 0 {
		return nil, 422, " -STR003", false, errors.New("no matching data for this Circle and exam year")
	}

	maxNotificationCodes := make(map[int32]int32)

	for _, notification := range in {
		key := notification.ExamCode
		// Update maxNotificationCodes if the current NotificationCode is greater
		if max, exists := maxNotificationCodes[key]; !exists || notification.ID > max {
			maxNotificationCodes[key] = notification.ID
		}
	}

	for _, issuedNotification := range in {
		key := issuedNotification.ExamCode
		if maxID, exists := maxNotificationCodes[key]; exists && issuedNotification.ID == maxID {
			IssuedNotification = append(IssuedNotification, issuedNotification)
		}
	}

	if len(IssuedNotification) <= 0 {
		return nil, 422, " -STR004", false, errors.New("no matching data for this Circle and exam year")

	}
	return IssuedNotification, 200, "", true, nil
}

// Function to Issue Notification
// func IssueExamNotification(client *ent.Client, newNotification *ent.ExamNotifications, ud *ent.AdminMaster) (string, error) {
// 	ctx := context.Background()
// 	examid := newNotification.ExamCode
// 	examyear := newNotification.ExamYear
// 	status := newNotification.NotificationStatus

// 	// editflag := newNotification.EditFlagStatus
// 	username := newNotification.UserName

// 	if examid <= 0 {
// 		return "", fmt.Errorf("invalid Exam Code: %d", examid)
// 	}

// 	examcode, err := client.Exam.
// 		Query().
// 		Where(exam.ExamCodeEQ(examid)).
// 		Only(ctx)

// 	if err != nil {
// 		log.Println("Error querying ExamCode :", err)
// 		return "", fmt.Errorf("No exams found for the given exam Code: %w", err)
// 	}

// 	if examcode == nil {
// 		return "", fmt.Errorf("Exam Details not found with the Exam Code: %d", examid)
// 	}

// 	examdetail, err := client.Exam.Query().Where(exam.ExamCodeEQ(examid)).Only(ctx)
// 	if err != nil {
// 		fmt.Printf("failed to retrieve data: %v", err)
// 		return "", fmt.Errorf("failed to retrieve data: %w", err)
// 	}
// 	examname := examdetail.ExamName

// 	exists, err := client.ExamNotifications.
// 		Query().
// 		Where(
// 			examnotifications.ExamCodeEQ(newNotification.ExamCode),
// 			examnotifications.ExamYearEQ(newNotification.ExamYear),
// 			examnotifications.FlagNEQ(true),
// 			examnotifications.NotificationStatusEQ(status)).Exist(ctx)

// 	if err != nil {
// 		log.Println("Failed to check existing Notification:", err)
// 		return "", fmt.Errorf("failed to check existing Notification: %v", err)
// 	}

// 	if exists {
// 		return "", fmt.Errorf("The Notification is already Issued for the Exam: ")
// 	}

// 	notificationNumber, err := generateNotificationNumber(client, newNotification.ExamCode)
// 	if err != nil {
// 		log.Printf("Failed to generate Notification number: %v", err)
// 		return "", err
// 	}

// 	// Use the generated application number
// 	log.Printf("Generated Notification number: %s", notificationNumber)

// 	currentTime := time.Now().Truncate(time.Second)

// 	exists, nerr := client.AdminMaster.Query().Where(adminmaster.UserName(username), adminmaster.StatussEQ("active")).Exist(ctx)
// 	if nerr != nil {
// 		log.Println("Error checking username existence:", err)
// 		return "", fmt.Errorf("failed to check username existence: %w", err)
// 	}

// 	if !exists {
// 		return "", fmt.Errorf("Invalid username: %s", username)
// 	}

// 	// if editflag == false {
// 	// 	_, nerr := client.ExamNotifications.Update().
// 	// 		SetNotificationStatus(status).
// 	// 		SetExamCode(newNotification.ExamCode).
// 	// 		SetExamName(examname).
// 	// 		SetNotificationRemarks(newNotification.NotificationRemarks).
// 	// 		SetUserName(username).
// 	// 		SetEditFlagStatus(newNotification.EditFlagStatus).
// 	// 		SetNotificationNumber(notificationNumber).
// 	// 		SetUpdatedAt(currentTime).                                                                                                  // Set the current UTC time
// 	// 		SetUpdatedBy(newNotification.UserName).Where(examnotifications.ExamCodeEQ(examid), examnotifications.ExamYearEQ(examyear)). // Set the updated by user (replace with the actual user)
// 	// 		Save(ctx)
// 	// 	if nerr != nil {
// 	// 		log.Println("error at Issuing Notification: ", newNotification)
// 	// 		return "", fmt.Errorf("failed Issuing Notification")
// 	// 	}
// 	// } else if editflag == true {
// 	originalNotification, err := client.ExamNotifications.
// 		Query().
// 		// Where(examnotifications.ExamCodeEQ(examid), examnotifications.ExamYearEQ(examyear), examnotifications.NotificationStatusIn("NotificationSubmitted", "NotificationResubmitted")).Only(ctx)
// 		Where(examnotifications.ExamCodeEQ(examid), examnotifications.ExamYearEQ(examyear), examnotifications.NotificationStatusNEQ("NotificationIssued"), examnotifications.NotificationStatusNEQ("NotificationReturned"), examnotifications.NotificationStatusNEQ("NotificationCancelled"), examnotifications.FlagNEQ(true)).Only(ctx)

// 	if err != nil {
// 		log.Println("Failed to check existing Draft Notification:", err)
// 		return "", fmt.Errorf("failed to check existing Draft Notification: %v", err)
// 	}

// 	if originalNotification == nil {
// 		return "", fmt.Errorf("the Draft Notification is not available")
// 	}

// 	// currentTime := time.Now().Truncate(time.Second)
// 	st, err3 := client.ExamNotifications.Update().SetFlag(true).Where(examnotifications.ExamCodeEQ(newNotification.ExamCode), examnotifications.ExamYearEQ(newNotification.ExamYear)).Save(ctx)
// 	if err3 != nil {
// 		log.Println("Failed to check existing Draft Notification:", err)
// 		return "", fmt.Errorf("failed to check existing Draft Notification: %v", err)
// 	}
// 	fmt.Println(st)

// 	duplicatedNotification := client.ExamNotifications.
// 		Create().
// 		SetExamCode(originalNotification.ExamCode).
// 		SetExamName(originalNotification.ExamName).
// 		SetExamShortName(originalNotification.ExamShortName).
// 		SetUserName(username).
// 		SetExamYear(originalNotification.ExamYear).
// 		SetApplicationStartDate(originalNotification.ApplicationStartDate).
// 		SetApplicationEndDate(originalNotification.ApplicationEndDate).
// 		//SetApplicationCorrectionStartDate(newNotification.ApplicationCorrectionStartDate).
// 		SetApplicationCorrectionLastDate(originalNotification.ApplicationCorrectionLastDate).
// 		SetApplicationVerificationLastDate(originalNotification.ApplicationVerificationLastDate).
// 		SetNodalOfficerApprovalDate(originalNotification.NodalOfficerApprovalDate).
// 		SetAdmitCardDate(originalNotification.AdmitCardDate).
// 		SetCrucialDate(originalNotification.CrucialDate).
// 		SetNotificationRemarks(originalNotification.NotificationRemarks).
// 		SetResubmittedApplicationVerificationDate(originalNotification.ResubmittedApplicationVerificationDate).
// 		SetPapers(originalNotification.Papers).
// 		SetNotificationStatus(status).
// 		SetNotificationReIssueStatus(false).
// 		SetEditFlagStatus(true).
// 		// SetDesignation(newNotification.Designation).
// 		// SetOfficerName(newNotification.OfficerName).
// 		SetApprovedBy(newNotification.ApprovedBy).
// 		SetApprovedByDesignation(newNotification.ApprovedByDesignation).
// 		SetApprovedByName(newNotification.ApprovedByName).
// 		SetApprovedById(ud.ID).
// 		SetCreatedById(ud.ID).
// 		SetCreatedBy(originalNotification.CreatedBy).
// 		SetCreatedByDesignation(originalNotification.CreatedByDesignation).
// 		SetCreatedByName(originalNotification.CreatedByName).
// 		SetNotificationOrderNumber(originalNotification.NotificationOrderNumber).
// 		SetNotesheetScannedCopy(originalNotification.NotesheetScannedCopy).
// 		SetOrderDate(originalNotification.OrderDate).
// 		SetNotificationNumber(notificationNumber).
// 		SetCenterAllotmentEndDate(originalNotification.CenterAllotmentEndDate).
// 		SetUpdatedAt(currentTime). // Set the current UTC time
// 		SetUpdatedBy(username).    // Set the updated by user (replace with the actual user)
// 		SaveX(ctx)

// 	if duplicatedNotification == nil {
// 		return "", fmt.Errorf("Failed to duplicate Original Notification")
// 	}

// 	// }
// 	message := fmt.Sprintf("Notification Issued Successfully for  %s for the year %d", examname, examyear)
// 	return message, nil
// }

// modified
func IssueExamNotification(ctx context.Context, client *ent.Client, newNotification *ca_reg.IssueNotificationStruct, ud *ent.AdminMaster) (string, int32, string, bool, error) {
	//ctx := context.Background()

	examcode := newNotification.ExamCode
	examyear := newNotification.ExamYear
	// Validate exam code
	if newNotification.ExamCode <= 0 {
		return "", 422, " -STR001", false, fmt.Errorf(" invalid Exam Code: %d", newNotification.ExamCode)
	}
	// Start transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		return "", 500, " -STR004", false, err
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

	// Retrieve exam details
	examDetail, err := tx.Exam.
		Query().
		Where(exam.ExamCodeEQ(examcode)).
		Only(ctx)
	if err != nil {
		return "", 500, " -STR001", false, err
	}

	if examDetail == nil {
		return "", 422, " -STR001", false, fmt.Errorf("exam not found with Exam Code: %d", examcode)
	}

	circleFacilities, err := client.FacilityMasters.Query().
		Where(
			facilitymasters.CircleFacilityIDNEQ("AS71000000141"),
			facilitymasters.CircleFacilityIDNEQ("CR34000000000"),
			facilitymasters.CircleFacilityIDNEQ("DT01000000000"),
			facilitymasters.CircleFacilityIDNEQ(" "),
			facilitymasters.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		return "", 500, " -STR001", false, err

	}
	notificationNumber, err := GenerateNotificationNumber(client, examcode)
	if err != nil {
		return "", 500, " -STR001", false, err
	}
	oldAppln, err := client.ExamNotifications.
		Query().
		Where(
			examnotifications.ExamCodeEQ(examcode),
			examnotifications.ExamYearEQ(examyear),
			examnotifications.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return "", 422, " -STR001", false, errors.New(" no active notification found for this Exam ")
		} else {
			return "", 500, " -STR001", false, err
		}
	}
	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err = oldAppln.
		Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return "", 500, " -STR001", false, err
	}

	// Iterate over each circle facility and issue a notification
	for _, facility := range circleFacilities {
		// Create a new notification for each circle facility
		//newNotification := *issueNotification
		newNotification.CircleOfficeFacilityId = facility.CircleFacilityID
		newNotification.CircleOfficeName = facility.CircleFacilityName
		newNotification.NotificationNumber = notificationNumber

		// Check if notification already exists for the exam
		exists, err := tx.ExamNotifications.
			Query().
			Where(
				examnotifications.ExamCodeEQ(examcode),
				examnotifications.ExamYearEQ(examyear),
				examnotifications.StatusEQ("active"),
				examnotifications.NotificationStatusEQ(newNotification.NotificationStatus),
				examnotifications.CircleOfficeFacilityIdEQ(newNotification.CircleOfficeFacilityId)).
			Exist(ctx)
		if err != nil {
			return "", 500, " -STR001", false, err
		}

		if exists {
			return "", 422, " -STR001", false, fmt.Errorf("notification already issued for Exam Code: %d", newNotification.ExamCode)
		}

		currentTime := time.Now().Truncate(time.Second)

		originalNotification, err := tx.ExamNotifications.
			Query().
			Where(examnotifications.ExamCodeEQ(examcode),
				examnotifications.ExamYearEQ(examyear),
				examnotifications.StatusEQ("active"),
				examnotifications.NotificationStatusNotIn("NotificationIssued", "NotificationReturned",
					"NotificationCancelled")).
			Only(ctx)

		if err != nil {
			return "", 500, " -STR001", false, err
		}

		if originalNotification == nil {
			return "", 422, " -STR001", false, fmt.Errorf("the Draft Notification is not available")
		}

		// currentTime := time.Now().Truncate(time.Second)
		_, err3 := tx.ExamNotifications.
			Update().
			SetFlag(true).
			Where(
				examnotifications.ExamCodeEQ(examcode),
				examnotifications.StatusEQ("active"),
				examnotifications.ExamYearEQ(examyear)).
			Save(ctx)
		if err3 != nil {
			return "", 500, " -STR001", false, err
		}

		// Create the new notification
		_, err = tx.ExamNotifications.
			Create().
			SetExamCode(originalNotification.ExamCode).
			SetExamName(originalNotification.ExamName).
			SetExamShortName(originalNotification.ExamShortName).
			SetUserName(newNotification.UserName).
			SetExamYear(originalNotification.ExamYear).
			SetEmployeeMasterRequestLastDate(originalNotification.EmployeeMasterRequestLastDate).
			SetEmployeeMasterRequestApprovalLastDate(originalNotification.EmployeeMasterRequestApprovalLastDate).
			SetExamRegisterLastDate(originalNotification.ExamRegisterLastDate).
			SetApplicationStartDate(originalNotification.ApplicationStartDate).
			SetApplicationEndDate(originalNotification.ApplicationEndDate).
			SetApplicationCorrectionLastDate(originalNotification.ApplicationCorrectionLastDate).
			SetApplicationVerificationLastDate(originalNotification.ApplicationVerificationLastDate).
			SetNodalOfficerApprovalDate(originalNotification.NodalOfficerApprovalDate).
			SetAdmitCardDate(originalNotification.AdmitCardDate).
			SetCrucialDate(originalNotification.CrucialDate).
			SetSmsExamShortName(originalNotification.SmsExamShortName).
			SetNotificationRemarks(originalNotification.NotificationRemarks).
			SetResubmittedApplicationVerificationDate(originalNotification.ResubmittedApplicationVerificationDate).
			SetPapers(originalNotification.Papers).
			SetNotificationStatus(newNotification.NotificationStatus).
			SetNotificationReIssueStatus(false).
			SetEditFlagStatus(true).
			SetStatus("active").
			SetApprovedBy(newNotification.ApprovedBy).
			SetApprovedByDesignation(newNotification.ApprovedByDesignation).
			SetApprovedByName(newNotification.ApprovedByName).
			SetApprovedById(ud.ID).
			SetCreatedById(ud.ID).
			SetCreatedBy(originalNotification.CreatedBy).
			SetCreatedByDesignation(originalNotification.CreatedByDesignation).
			SetCreatedByName(originalNotification.CreatedByName).
			SetNotificationOrderNumber(originalNotification.NotificationOrderNumber).
			SetNotesheetScannedCopy(originalNotification.NotesheetScannedCopy).
			SetOrderDate(originalNotification.OrderDate).
			SetCenterAllotmentEndDate(originalNotification.CenterAllotmentEndDate).
			SetUpdatedAt(currentTime).              // Set the current UTC time
			SetUpdatedBy(newNotification.UserName). // Set the updated by user (replace with the actual user)
			SetCircleOfficeFacilityId(newNotification.CircleOfficeFacilityId).
			SetCircleOfficeName(newNotification.CircleOfficeName).
			SetNotificationNumber(newNotification.NotificationNumber).
			SetNotificationOrderNumber(newNotification.NotificationOrderNumber).
			Save(ctx)
		if err != nil {
			return "", 500, " -STR001", false, err
		}
	}
	// Commit transaction
	if err := tx.Commit(); err != nil {
		return "", 500, " -STR001", false, err
	}

	return "success", 200, "", false, nil
}

func IssueExamNotificationSingle(client *ent.Client, newNotification ca_reg.IssueNotificationStruct) (*ent.ExamNotifications, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	examcode := newNotification.ExamCode
	examyear := newNotification.ExamYear

	// Start transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, err
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

	// Validate exam code
	// Retrieve exam details
	_, err = tx.Exam.
		Query().
		Where(exam.ExamCodeEQ(examcode)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf(" no such exam available")
		} else {
			return nil, 500, " -STR002", false, err
		}
	}

	// Fetch old application to update its status
	originalNotification, err := tx.ExamNotifications.
		Query().
		Where(
			examnotifications.ExamCodeEQ(examcode),
			examnotifications.ExamYearEQ(examyear),
			examnotifications.StatusEQ("active"),
			examnotifications.NotificationStatusIn("NotificationResubmitted", "NotificationSubmitted"),
			examnotifications.CircleOfficeFacilityIdEQ(newNotification.CircleOfficeFacilityId),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR003", false, errors.New("no notification available in re-submitted state ")
		} else {
			return nil, 500, " -STR004", false, err
		}
	}

	// Generate notification number
	notificationNumber, err := GenerateNotificationNumber(client, newNotification.ExamCode)
	if err != nil {
		return nil, 400, " -STR005", false, fmt.Errorf("failed to generate notification number: %w", err)
	}

	newNotification.NotificationNumber = notificationNumber

	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err = originalNotification.Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	currentTime := time.Now().Truncate(time.Second)

	// Create the new notification
	createdNotification, err := tx.ExamNotifications.
		Create().
		SetIssuedBy(originalNotification.IssuedBy).
		SetExamCode(originalNotification.ExamCode).
		SetExamName(originalNotification.ExamName).
		SetExamShortName(originalNotification.ExamShortName).
		SetUserName(newNotification.UserName).
		SetExamYear(originalNotification.ExamYear).
		SetEmployeeMasterRequestLastDate(originalNotification.EmployeeMasterRequestLastDate).
		SetEmployeeMasterRequestApprovalLastDate(originalNotification.EmployeeMasterRequestApprovalLastDate).
		SetExamRegisterLastDate(originalNotification.ExamRegisterLastDate).
		SetApplicationStartDate(originalNotification.ApplicationStartDate).
		SetApplicationEndDate(originalNotification.ApplicationEndDate).
		SetCircleOfficeFacilityId(originalNotification.CircleOfficeFacilityId).
		SetCircleOfficeName(originalNotification.CircleOfficeName).
		SetApplicationCorrectionLastDate(originalNotification.ApplicationCorrectionLastDate).
		SetApplicationVerificationLastDate(originalNotification.ApplicationVerificationLastDate).
		SetNodalOfficerApprovalDate(originalNotification.NodalOfficerApprovalDate).
		SetAdmitCardDate(originalNotification.AdmitCardDate).
		SetCrucialDate(originalNotification.CrucialDate).
		SetSmsExamShortName(originalNotification.SmsExamShortName).
		SetNotificationRemarks(originalNotification.NotificationRemarks).
		SetResubmittedApplicationVerificationDate(originalNotification.ResubmittedApplicationVerificationDate).
		SetPapers(originalNotification.Papers).
		SetNotificationStatus(newNotification.NotificationStatus).
		SetNotificationReIssueStatus(false).
		SetEditFlagStatus(true).
		SetStatus("active").
		SetCreatedById(originalNotification.CreatedById).
		SetCreatedBy(originalNotification.CreatedBy).
		SetCreatedByDesignation(originalNotification.CreatedByDesignation).
		SetCreatedByName(originalNotification.CreatedByName).
		SetApprovedById(newNotification.ApprovedById).
		SetApprovedBy(newNotification.ApprovedBy).
		SetApprovedByDesignation(newNotification.ApprovedByDesignation).
		SetApprovedByName(newNotification.ApprovedByName).
		SetNotificationOrderNumber(originalNotification.NotificationOrderNumber).
		SetNotesheetScannedCopy(originalNotification.NotesheetScannedCopy).
		SetOrderDate(originalNotification.OrderDate).
		SetCenterAllotmentEndDate(originalNotification.CenterAllotmentEndDate).
		SetUpdatedAt(currentTime).              // Set the current UTC time
		SetUpdatedBy(newNotification.UserName). // Set the updated by user (replace with the actual user)
		SetCircleOfficeFacilityId(newNotification.CircleOfficeFacilityId).
		SetCircleOfficeName(newNotification.CircleOfficeName).
		SetNotificationNumber(newNotification.NotificationNumber).
		SetNotificationOrderNumber(newNotification.NotificationOrderNumber).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR007", false, err
	}
	// Commit transaction
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR008", false, err
	}
	return createdNotification, 200, "", true, nil
}

func GenerateNotificationNumber(client *ent.Client, examcode int32) (string, error) {
	examid := examcode
	nextNotificationNumber, err := getNextNotificationNumberFromDatabase(client)
	if err != nil {
		return "", err
	}

	// Get the current year
	currentYear := time.Now().Year()

	// Format the Notification number as "NOTYYYYCXX"
	notificationNumber := fmt.Sprintf("NOT%d%d%02d", currentYear, examid, nextNotificationNumber)

	return notificationNumber, nil
}

func getNextNotificationNumberFromDatabase(client *ent.Client) (int32, error) {
	ctx := context.TODO()
	lastNotification, err := client.ExamNotifications.
		Query().
		Order(ent.Desc(examnotifications.FieldID)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return 01, nil
		}
		return 0, fmt.Errorf("failed to retrieve last application: %v", err)
	}

	return lastNotification.ID, nil
}

// Function to enable resubmit draft notification to OA user by Supervisor

/*
	 func PutResubmitDraftNotification(client *ent.Client, newNotification *ent.ExamNotifications, ud *ent.AdminMaster) (string, error) {
		ctx := context.Background()
		examid := newNotification.ExamCode
		status := newNotification.NotificationStatus
		username := newNotification.UserName
		//notificationID := newNotification.ID
		examyear := newNotification.ExamYear
		remarks := newNotification.NotificationRemarks
		//All Approved 3 fields

		if examid <= 0 {
			return "", fmt.Errorf("invalid Exam Code: %d", examid)
		}

		examcode, err := client.Exam.
			Query().
			Where(exam.ID(examid)).
			Only(ctx)

		if err != nil {
			log.Println("Error querying ExamCode :", err)
			return "", fmt.Errorf(" no exams found for the given exam Code: %w", err)
		}

		if examcode == nil {
			return "", fmt.Errorf(" exam Details not found with the Exam Code: %d", examid)
		}

		// examdetail, err := client.Exam.Query().Where(exam.IDEQ(examid)).Only(ctx)
		// if err != nil {
		// 	fmt.Printf("failed to retrieve data: %v", err)
		// 	return nil, fmt.Errorf("failed to retrieve data: %w", err)
		// }
		// examname := examdetail.ExamName

		exists, nerr := client.AdminMaster.Query().Where(adminmaster.UserName(username), adminmaster.StatussEQ("active")).Exist(ctx)
		if nerr != nil {
			log.Println("Error checking username existence:", err)
			return "", fmt.Errorf("failed to check username existence: %w", err)
		}

		if !exists {
			return "", fmt.Errorf(" invalid username: %s", username)
		}

		exists, err1 := client.ExamNotifications.
			Query().
			Where(examnotifications.FlagNEQ(true), examnotifications.ExamCodeEQ(examid), examnotifications.ExamYearEQ(examyear)).Exist(ctx)

		if err1 != nil {
			log.Println("Failed to check existing Notification:", err)
			return "", fmt.Errorf("failed to check existing Notification: %v", err)
		}

		if !exists {
			return "", fmt.Errorf(" the Draft Notification is not available")
		}

		data, err2 := client.ExamNotifications.Query().Where(examnotifications.FlagNEQ(true), examnotifications.ExamCodeEQ(examid), examnotifications.ExamYearEQ(examyear), examnotifications.StatusEQ("active")).Only(ctx)
		if err2 != nil {
			log.Println("Failed to check existing Notification:", err)
			return "", fmt.Errorf("failed to check existing Notification: %v", err)
		}

		currentTime := time.Now().Truncate(time.Second)

		// _, nerr1 := client.ExamNotifications.Update().
		// 	SetUserName(username).
		// 	SetNotificationStatus(status).
		// 	SetNotificationRemarks(newNotification.NotificationRemarks).
		// 	SetUpdatedAt(currentTime).                                            // Set the current UTC time
		// 	SetUpdatedBy(username).Where(examnotifications.IDEQ(notificationID)). // Set the updated by user (replace with the actual user)
		// 	Save(ctx)

		st, err3 := client.ExamNotifications.Update().SetFlag(true).Where(examnotifications.ExamCodeEQ(newNotification.ExamCode), examnotifications.ExamYearEQ(newNotification.ExamYear)).Save(ctx)
		if err3 != nil {
			log.Println("Failed to check existing Draft Notification:", err)
			return "", fmt.Errorf("failed to check existing Draft Notification: %v", err)
		}
		fmt.Println(st)
		oldAppln, err := client.ExamNotifications.
			Query().
			Where(
				examnotifications.ExamCodeEQ(newNotification.ExamCode),
				examnotifications.ExamYearEQ(newNotification.ExamYear),
				examnotifications.StatusEQ("active"),
			).
			All(ctx)

		if err != nil {
			return "", fmt.Errorf("error in fetching data from the Exam table: %v", err)
		}

		if len(oldAppln) == 0 {
			return "", fmt.Errorf("no active notification found for this Exam")
		}
		stat := "inactive_" + time.Now().Format("20060102150405")
		for _, notif := range oldAppln {
			_, err = notif.Update().
				SetStatus(stat).
				Save(ctx)
			if err != nil {
				return "", fmt.Errorf("failed to update application: %v", err)
			}
		}
		dn, nerr1 := client.ExamNotifications.Create().
			SetExamCode(data.ExamCode).
			SetExamName(data.ExamName).
			SetUserName(username).
			SetExamYear(data.ExamYear).
			SetApplicationStartDate(data.ApplicationStartDate).
			SetApplicationEndDate(data.ApplicationEndDate).
			// SetApplicationCorrectionStartDate(data.ApplicationCorrectionStartDate).
			SetApplicationCorrectionLastDate(data.ApplicationCorrectionLastDate).
			SetApplicationVerificationLastDate(data.ApplicationVerificationLastDate).
			SetNodalOfficerApprovalDate(data.NodalOfficerApprovalDate).
			SetAdmitCardDate(data.AdmitCardDate).
			SetCrucialDate(data.CrucialDate).
			SetResubmittedApplicationVerificationDate(data.ResubmittedApplicationVerificationDate).
			SetPapers(data.Papers).
			SetNotificationStatus(status).
			SetNotificationReIssueStatus(data.NotificationReIssueStatus).
			SetEditFlagStatus(data.EditFlagStatus).
			// SetDesignation(data.Designation).
			// SetOfficerName(data.OfficerName).
			SetNotificationOrderNumber(data.NotificationOrderNumber).
			SetNotesheetScannedCopy(data.NotesheetScannedCopy).
			SetOrderDate(data.OrderDate).
			SetCenterAllotmentEndDate(data.CenterAllotmentEndDate).
			SetNotificationRemarks(remarks).
			SetExamShortName(data.ExamShortName).
			SetIssuedBy(data.IssuedBy).
			SetCircleOfficeFacilityId(data.CircleOfficeFacilityId).
			SetCircleOfficeName(data.CircleOfficeName).
			SetApprovedBy(newNotification.ApprovedBy).
			SetApprovedByDesignation(newNotification.ApprovedByDesignation).
			SetApprovedByName(newNotification.ApprovedByName).
			SetApprovedById(ud.ID).
			SetStatus("active").
			//SetNotificationNumber(newNotification.NotificationNumber).
			SetUpdatedAt(currentTime). // Set the current UTC time
			SetUpdatedBy(username).    // Set the updated by user (replace with the actual user)
			Save(ctx)
		if nerr1 != nil {
			log.Println("error at Updating Draft Notification: ", newNotification)
			return "", fmt.Errorf("failed Updating Draft Notification: %w", err)
		}
		fmt.Println(dn)

		message := "Draft Notification Returned successfully to OA user"
		return message, nil
	}
*/
func PutResubmitDraftNotification(client *ent.Client, newNotification *ca_reg.ResubitNotificationStruct) (string, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	examid := newNotification.ExamCode
	status := newNotification.NotificationStatus
	username := newNotification.UserName
	examyear := newNotification.ExamYear
	//	remarks := newNotification.NotificationRemarks
	circleid := newNotification.CircleOfficeFacilityId

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

	if examid <= 0 {
		return "", 400, " -STR001", false, fmt.Errorf("invalid Exam Code: %d", examid)
	}

	_, err = tx.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return "", 422, " -STR002", false, fmt.Errorf(" no such exam available ")
		} else {
			return "", 500, " -STR003", false, err
		}
	}

	currentTime := time.Now().Truncate(time.Second)

	data, err := client.ExamNotifications.
		Query().
		Where(
			examnotifications.ExamCodeEQ(examid),
			examnotifications.StatusEQ("active"),
			examnotifications.CircleOfficeFacilityIdEQ(circleid),
			examnotifications.NotificationStatusIn("NotificationSubmitted", "NotificationResubmitted"),
			examnotifications.ExamYearEQ(examyear)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return "", 422, " -STR004", false, fmt.Errorf(" no notification found for this exam and year with Notification Submitted / Notification Resubmitted status ")
		} else {
			return "", 500, " -STR005", false, err
		}
	}
	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err = data.Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return "", 500, " -STR006", false, err
	}

	_, err = client.ExamNotifications.Create().
		SetExamCode(data.ExamCode).
		SetExamName(data.ExamName).
		SetUserName(username).
		SetExamYear(data.ExamYear).
		SetExamShortName(data.ExamShortName).
		SetEmployeeMasterRequestLastDate(newNotification.EmployeeMasterRequestLastDate).
		SetEmployeeMasterRequestApprovalLastDate(newNotification.EmployeeMasterRequestApprovalLastDate).
		SetExamRegisterLastDate(newNotification.ExamRegisterLastDate).
		SetApplicationStartDate(newNotification.ApplicationStartDate).
		SetApplicationEndDate(newNotification.ApplicationEndDate).
		SetApplicationCorrectionLastDate(newNotification.ApplicationCorrectionLastDate).
		SetApplicationVerificationLastDate(newNotification.ApplicationVerificationLastDate).
		SetNodalOfficerApprovalDate(newNotification.NodalOfficerApprovalDate).
		SetAdmitCardDate(newNotification.AdmitCardDate).
		SetCrucialDate(*newNotification.CrucialDate).
		SetSmsExamShortName(data.SmsExamShortName).
		SetResubmittedApplicationVerificationDate(data.ResubmittedApplicationVerificationDate).
		SetPapers(*newNotification.Papers).
		SetNotificationStatus(status).
		SetNotificationOrderNumber(newNotification.NotificationOrderNumber).
		SetOrderDate(newNotification.OrderDate).
		SetCenterAllotmentEndDate(newNotification.CenterAllotmentEndDate).
		SetNotificationRemarks(newNotification.NotificationRemarks).
		SetExamShortName(newNotification.ExamShortName).
		SetIssuedBy(data.IssuedBy).
		SetCircleOfficeFacilityId(data.CircleOfficeFacilityId).
		SetCircleOfficeName(data.CircleOfficeName).
		SetCreatedById(data.CreatedById).
		SetCreatedBy(data.CreatedBy).
		SetCreatedByName(data.CreatedByName).
		SetCreatedByDesignation(data.CreatedByDesignation).
		SetApprovedById(newNotification.CreatedById).
		SetApprovedBy(newNotification.CreatedBy).
		SetApprovedByName(newNotification.CreatedByName).
		SetApprovedByDesignation(newNotification.CreatedByDesignation).
		SetStatus("active").
		SetUpdatedAt(currentTime). // Set the current UTC time
		SetUpdatedBy(username).    // Set the updated by user (replace with the actual user)
		Save(ctx)
	if err != nil {
		return "", 500, " -STR007", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return "", 500, " -STR008", false, err
	}

	message := "Draft Notification Returned successfully to OA user"
	return message, 200, "", true, nil
}

// Function to Update Resubmitted Draft

func UpdateResubmitDraftNotificationNew(client *ent.Client, newNotification ca_reg.UpdateNotificationStruct) (string, int32, string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examid := newNotification.ExamCode
	username := newNotification.UserName
	examyear := newNotification.ExamYear
	circleid := newNotification.CircleOfficeFacilityId
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

	_, err = tx.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return "", 422, " -STR001", false, fmt.Errorf(" no such exam available ")
		} else {
			return "", 500, " -STR002", false, err
		}
	}

	oldAppln, err := tx.ExamNotifications.
		Query().
		Where(
			examnotifications.ExamCodeEQ(examid),
			examnotifications.ExamYearEQ(examyear),
			examnotifications.CircleOfficeFacilityIdEQ(circleid),
			examnotifications.NotificationStatusEQ("NotificationReturned"),
			examnotifications.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return "", 422, " -STR003", false, errors.New(" no active notification found for this Exam ")
		} else {
			return "", 500, " -STR004", false, err
		}
	}
	stat := "inactive_" + time.Now().Format("20060102150405")
	_, err = oldAppln.Update().
		SetStatus(stat).
		Save(ctx)
	if err != nil {
		return "", 500, " -STR005", false, err
	}
	_, err = tx.ExamNotifications.Create().
		SetNotificationStatus(newNotification.NotificationStatus).
		SetExamCode(newNotification.ExamCode).
		SetExamName(newNotification.ExamName).
		SetExamShortName(newNotification.ExamShortName).
		SetIssuedBy(newNotification.IssuedBy).
		SetCircleOfficeFacilityId(newNotification.CircleOfficeFacilityId).
		SetCircleOfficeName(newNotification.CircleOfficeName).
		SetNotificationRemarks(newNotification.NotificationRemarks).
		SetUserName(username).
		SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
		SetUpdatedBy(newNotification.UserName).
		SetOrderDate(newNotification.OrderDate).
		SetExamYear(newNotification.ExamYear).
		SetPapers(*newNotification.Papers).
		SetEmployeeMasterRequestLastDate(newNotification.EmployeeMasterRequestLastDate).
		SetEmployeeMasterRequestApprovalLastDate(newNotification.EmployeeMasterRequestApprovalLastDate).
		SetExamRegisterLastDate(newNotification.ExamRegisterLastDate).
		SetApplicationStartDate(newNotification.ApplicationStartDate).
		SetApplicationEndDate(newNotification.ApplicationEndDate).
		SetApplicationVerificationLastDate(newNotification.ApplicationVerificationLastDate).
		SetApplicationCorrectionLastDate(newNotification.ApplicationCorrectionLastDate).
		SetResubmittedApplicationVerificationDate(newNotification.ResubmittedApplicationVerificationDate).
		SetNodalOfficerApprovalDate(newNotification.NodalOfficerApprovalDate).
		SetAdmitCardDate(newNotification.AdmitCardDate).
		SetCreatedById(newNotification.CreatedById).
		SetCreatedBy(newNotification.CreatedBy).
		SetCreatedByDesignation(newNotification.CreatedByDesignation).
		SetCreatedByName(newNotification.CreatedByName).
		SetCrucialDate(*newNotification.CrucialDate).
		SetSmsExamShortName(oldAppln.SmsExamShortName).
		SetNotificationOrderNumber(newNotification.NotificationOrderNumber).
		SetCenterAllotmentEndDate(newNotification.CenterAllotmentEndDate).
		SetStatus("active").
		Save(ctx)

	if err != nil {
		return "", 500, " -STR006", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return "", 500, " -STR007", false, err
	}

	message := "Notification Updated Sucessfully"
	return message, 200, "", true, nil
}

// Function to Cancel Draft
func UpdateCancelDraftNotification(client *ent.Client, newNotification *ent.ExamNotifications, ud *ent.AdminMaster) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examid := newNotification.ExamCode
	//status := newNotification.NotificationStatus
	username := newNotification.UserName
	// notificationID := newNotification.ID
	examyear := newNotification.ExamYear

	if examid <= 0 {
		return "", fmt.Errorf("invalid Exam Code: %d", examid)
	}

	examcode, err := client.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		log.Println("Error querying ExamCode :", err)
		return "", fmt.Errorf(" no exams found for the given exam Code: %w", err)
	}

	if examcode == nil {
		return "", fmt.Errorf(" exam Details not found with the Exam Code: %d", examid)
	}

	exists, nerr := client.AdminMaster.Query().Where(adminmaster.UserName(username)).Exist(ctx)
	if nerr != nil {
		log.Println("Error checking username existence:", err)
		return "", fmt.Errorf("failed to check username existence: %w", err)
	}

	if !exists {
		return "", fmt.Errorf(" invalid username: %s", username)
	}

	exists, err1 := client.ExamNotifications.
		Query().
		Where(examnotifications.FlagNEQ(true), examnotifications.ExamCodeEQ(examid), examnotifications.ExamYearEQ(examyear)).Exist(ctx)

	if err1 != nil {
		log.Println("Failed to check existing Notification:", err)
		return "", fmt.Errorf("failed to check existing Notification: %v", err)
	}

	if !exists {
		return "", fmt.Errorf("the Draft Notification is not available")
	}

	if newNotification.ExamCode == 0 ||
		newNotification.ExamName == "" ||
		newNotification.ExamYear == 0 ||
		newNotification.UserName == "" ||
		newNotification.OrderDate.IsZero() ||
		newNotification.ApplicationStartDate.IsZero() ||
		newNotification.ApplicationEndDate.IsZero() ||
		newNotification.ApplicationVerificationLastDate.IsZero() ||
		// newNotification.ApplicationCorrectionStartDate.IsZero() ||
		newNotification.ApplicationCorrectionLastDate.IsZero() ||
		newNotification.ResubmittedApplicationVerificationDate.IsZero() ||
		newNotification.NodalOfficerApprovalDate.IsZero() ||
		newNotification.AdmitCardDate.IsZero() ||
		newNotification.NotificationOrderNumber == "" ||

		// newNotification.Designation == "" ||
		newNotification.CenterAllotmentEndDate.IsZero() ||
		newNotification.UpdatedBy == "" || (len(newNotification.Papers) == 0) || (len(newNotification.CrucialDate) == 0) || newNotification.NotificationStatus == "" {
		// Handle the error, e.g., return an error or log a message
		return "", fmt.Errorf(" ensure All Fileds are not empty")
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		log.Println("Failed to start transaction:", err)
		return "", fmt.Errorf(" failed to start transaction: %v", err)
	}

	// Defer rollback if there's an error
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Println("Failed to rollback transaction:", rerr)
			}
		}
	}()

	st, err4 := tx.ExamNotifications.Update().SetFlag(true).Where(examnotifications.ExamCodeEQ(newNotification.ExamCode), examnotifications.ExamYearEQ(newNotification.ExamYear)).Save(ctx)
	if err4 != nil {
		log.Println("Failed to check existing Draft Notification:", err)
		return "", fmt.Errorf("failed to check existing Draft Notification: %v", err)
	}
	fmt.Println("here", st)

	_, err3 := tx.ExamNotifications.Create().
		SetNotificationStatus(newNotification.NotificationStatus).
		SetExamCode(newNotification.ExamCode).
		SetExamName(newNotification.ExamName).
		SetNotificationRemarks(newNotification.NotificationRemarks).
		SetUserName(username).
		SetExamShortName(newNotification.ExamShortName).
		// SetNotificationNumber(newNotification.NotificationNumber).
		SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
		SetUpdatedBy(newNotification.UserName).
		SetOrderDate(newNotification.OrderDate).
		SetExamYear(newNotification.ExamYear).
		SetPapers(newNotification.Papers).
		SetEmployeeMasterRequestLastDate(newNotification.EmployeeMasterRequestLastDate).
		SetEmployeeMasterRequestApprovalLastDate(newNotification.EmployeeMasterRequestApprovalLastDate).
		SetExamRegisterLastDate(newNotification.ExamRegisterLastDate).
		SetApplicationStartDate(newNotification.ApplicationStartDate).
		SetApplicationEndDate(newNotification.ApplicationEndDate).
		SetApplicationVerificationLastDate(newNotification.ApplicationVerificationLastDate).
		// SetApplicationCorrectionStartDate(newNotification.ApplicationCorrectionStartDate).
		SetApplicationCorrectionLastDate(newNotification.ApplicationCorrectionLastDate).
		SetResubmittedApplicationVerificationDate(newNotification.ResubmittedApplicationVerificationDate).
		SetNodalOfficerApprovalDate(newNotification.NodalOfficerApprovalDate).
		SetAdmitCardDate(newNotification.AdmitCardDate).
		// SetOfficerName(newNotification.OfficerName).
		// SetDesignation(newNotification.Designation).
		SetApprovedBy(newNotification.ApprovedBy).
		SetApprovedById(ud.ID).
		SetApprovedByDesignation(newNotification.ApprovedByDesignation).
		SetApprovedByName(newNotification.ApprovedByName).
		SetCrucialDate(newNotification.CrucialDate).
		SetNotificationOrderNumber(newNotification.NotificationOrderNumber).
		SetCenterAllotmentEndDate(newNotification.CenterAllotmentEndDate).
		SetCircleOfficeFacilityId(newNotification.CircleOfficeFacilityId).
		SetCircleOfficeName(newNotification.CircleOfficeName).
		SetStatus("active").
		Save(ctx)

	if err3 != nil {
		log.Println("Error at Cancelling Notification: ", newNotification)
		return "", fmt.Errorf(" failed to Cancel the Notification")
	}

	if err := tx.Commit(); err != nil {
		log.Println("Failed to commit transaction:", err)
		return "", fmt.Errorf(" failed to commit transaction: %v", err)
	}

	message := "Notification Cancelled"
	return message, nil
}

// Function to Cancel Draft

func SubUpdateCancelDraftNotificationSingle(client *ent.Client, newNotification *ent.ExamNotifications, ud *ent.AdminMaster) (string, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examid := newNotification.ExamCode
	//status := newNotification.NotificationStatus
	username := newNotification.UserName
	// notificationID := newNotification.ID
	examyear := newNotification.ExamYear

	if examid <= 0 {
		return "", 422, " -STR001", false, fmt.Errorf("invalid Exam Code: %d", examid)
	}

	examcode, err := client.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		return "", 422, " -STR002", false, fmt.Errorf(" no exams found for the given exam Code: %w", err)
	}

	if examcode == nil {
		return "", 422, " -STR003", false, fmt.Errorf(" exam Details not found with the Exam Code: %d", examid)
	}

	exists, nerr := client.AdminMaster.Query().Where(adminmaster.UserName(username)).Exist(ctx)
	if nerr != nil {
		return "", 500, " -STR004", false, err
	}

	if !exists {
		return "", 422, " -STR005", false, fmt.Errorf(" invalid username: %s", username)
	}

	exists, err1 := client.ExamNotifications.
		Query().
		Where(examnotifications.FlagNEQ(true), examnotifications.ExamCodeEQ(examid), examnotifications.ExamYearEQ(examyear)).Exist(ctx)

	if err1 != nil {
		return "", 500, " -STR006", false, err
	}

	if !exists {
		return "", 422, " -STR007", false, fmt.Errorf("the Draft Notification is not available")
	}

	if newNotification.ExamCode == 0 ||
		newNotification.ExamName == "" ||
		newNotification.ExamYear == 0 ||
		newNotification.UserName == "" ||
		newNotification.OrderDate.IsZero() ||
		newNotification.ApplicationStartDate.IsZero() ||
		newNotification.ApplicationEndDate.IsZero() ||
		newNotification.ApplicationVerificationLastDate.IsZero() ||
		// newNotification.ApplicationCorrectionStartDate.IsZero() ||
		newNotification.ApplicationCorrectionLastDate.IsZero() ||
		newNotification.ResubmittedApplicationVerificationDate.IsZero() ||
		newNotification.NodalOfficerApprovalDate.IsZero() ||
		newNotification.AdmitCardDate.IsZero() ||
		newNotification.NotificationOrderNumber == "" ||

		// newNotification.Designation == "" ||
		newNotification.CenterAllotmentEndDate.IsZero() ||
		newNotification.UpdatedBy == "" || (len(newNotification.Papers) == 0) || (len(newNotification.CrucialDate) == 0) || newNotification.NotificationStatus == "" {
		// Handle the error, e.g., return an error or log a message
		return "", 422, " -STR008", false, fmt.Errorf(" ensure All Fileds are not empty")
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		return "", 500, " -STR006", false, err
	}

	// Defer rollback if there's an error
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Println("Failed to rollback transaction:", rerr)
			}
		}
	}()

	_, err = tx.ExamNotifications.Update().SetFlag(true).Where(examnotifications.ExamCodeEQ(newNotification.ExamCode), examnotifications.ExamYearEQ(newNotification.ExamYear)).Save(ctx)
	if err != nil {
		return "", 500, " -STR006", false, err
	}

	_, err = tx.ExamNotifications.Create().
		SetNotificationStatus(newNotification.NotificationStatus).
		SetExamCode(newNotification.ExamCode).
		SetExamName(newNotification.ExamName).
		SetNotificationRemarks(newNotification.NotificationRemarks).
		SetUserName(username).
		SetExamShortName(newNotification.ExamShortName).
		// SetNotificationNumber(newNotification.NotificationNumber).
		SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
		SetUpdatedBy(newNotification.UserName).
		SetOrderDate(newNotification.OrderDate).
		SetExamYear(newNotification.ExamYear).
		SetPapers(newNotification.Papers).
		SetEmployeeMasterRequestLastDate(newNotification.EmployeeMasterRequestLastDate).
		SetEmployeeMasterRequestApprovalLastDate(newNotification.EmployeeMasterRequestApprovalLastDate).
		SetExamRegisterLastDate(newNotification.ExamRegisterLastDate).
		SetApplicationStartDate(newNotification.ApplicationStartDate).
		SetApplicationEndDate(newNotification.ApplicationEndDate).
		SetApplicationVerificationLastDate(newNotification.ApplicationVerificationLastDate).
		// SetApplicationCorrectionStartDate(newNotification.ApplicationCorrectionStartDate).
		SetApplicationCorrectionLastDate(newNotification.ApplicationCorrectionLastDate).
		SetResubmittedApplicationVerificationDate(newNotification.ResubmittedApplicationVerificationDate).
		SetNodalOfficerApprovalDate(newNotification.NodalOfficerApprovalDate).
		SetAdmitCardDate(newNotification.AdmitCardDate).
		// SetOfficerName(newNotification.OfficerName).
		// SetDesignation(newNotification.Designation).
		SetApprovedBy(newNotification.ApprovedBy).
		SetApprovedById(ud.ID).
		SetApprovedByDesignation(newNotification.ApprovedByDesignation).
		SetApprovedByName(newNotification.ApprovedByName).
		SetCrucialDate(newNotification.CrucialDate).
		SetNotificationOrderNumber(newNotification.NotificationOrderNumber).
		SetCenterAllotmentEndDate(newNotification.CenterAllotmentEndDate).
		SetCircleOfficeFacilityId(newNotification.CircleOfficeFacilityId).
		SetCircleOfficeName(newNotification.CircleOfficeName).
		SetStatus("active").
		Save(ctx)

	if err != nil {

		return "", 500, " -STR006", false, err
	}

	if err := tx.Commit(); err != nil {

		return "", 500, " -STR006", false, err
	}

	message := "Notification Cancelled"
	return message, 200, "", true, nil
}

func UpdateCancelDraftNotificationForSpecificCircle(client *ent.Client, newNotification *ent.ExamNotifications, ud *ent.AdminMaster) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examid := newNotification.ExamCode
	username := newNotification.UserName
	examyear := newNotification.ExamYear

	if examid <= 0 {
		return "", fmt.Errorf("invalid Exam Code: %d", examid)
	}

	examcode, err := client.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		log.Println("Error querying ExamCode :", err)
		return "", fmt.Errorf(" no exams found for the given exam Code: %w", err)
	}

	if examcode == nil {
		return "", fmt.Errorf(" exam Details not found with the Exam Code: %d", examid)
	}

	exists, nerr := client.AdminMaster.Query().Where(adminmaster.UserName(username)).Exist(ctx)
	if nerr != nil {
		log.Println("Error checking username existence:", err)
		return "", fmt.Errorf("failed to check username existence: %w", err)
	}

	if !exists {
		return "", fmt.Errorf(" invalid username: %s", username)
	}

	exists, err1 := client.ExamNotifications.
		Query().
		Where(examnotifications.ExamCodeEQ(examid), examnotifications.ExamYearEQ(examyear), examnotifications.CircleOfficeFacilityIdEQ(newNotification.CircleOfficeFacilityId)).Exist(ctx)

	if err1 != nil {
		log.Println("Failed to check existing Notification:", err)
		return "", fmt.Errorf("failed to check existing Notification: %v", err)
	}

	if !exists {
		return "", fmt.Errorf("the Draft Notification is not available")
	}

	if newNotification.ExamCode == 0 ||
		newNotification.ExamName == "" ||
		newNotification.ExamYear == 0 ||
		newNotification.UserName == "" ||
		newNotification.OrderDate.IsZero() ||
		newNotification.ApplicationStartDate.IsZero() ||
		newNotification.ApplicationEndDate.IsZero() ||
		newNotification.ApplicationVerificationLastDate.IsZero() ||
		newNotification.ApplicationCorrectionLastDate.IsZero() ||
		newNotification.ResubmittedApplicationVerificationDate.IsZero() ||
		newNotification.NodalOfficerApprovalDate.IsZero() ||
		newNotification.AdmitCardDate.IsZero() ||
		newNotification.NotificationOrderNumber == "" ||
		newNotification.CenterAllotmentEndDate.IsZero() ||
		newNotification.UpdatedBy == "" || (len(newNotification.Papers) == 0) || (len(newNotification.CrucialDate) == 0) || newNotification.NotificationStatus == "" {
		return "", fmt.Errorf("ensure All Fields are not empty")
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		log.Println("Failed to start transaction:", err)
		return "", fmt.Errorf("failed to start transaction: %v", err)
	}

	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Println("Failed to rollback transaction:", rerr)
			}
		}
	}()

	st, err4 := tx.ExamNotifications.Update().
		SetFlag(true).
		Where(examnotifications.ExamCodeEQ(newNotification.ExamCode), examnotifications.ExamYearEQ(newNotification.ExamYear), examnotifications.CircleOfficeFacilityIdEQ(newNotification.CircleOfficeFacilityId)).
		Save(ctx)
	if err4 != nil {
		log.Println("Failed to check existing Draft Notification:", err)
		return "", fmt.Errorf("failed to check existing Draft Notification: %v", err)
	}
	fmt.Println("here", st)

	_, err3 := tx.ExamNotifications.Create().
		SetNotificationStatus(newNotification.NotificationStatus).
		SetExamCode(newNotification.ExamCode).
		SetExamName(newNotification.ExamName).
		SetNotificationRemarks(newNotification.NotificationRemarks).
		SetUserName(username).
		SetExamShortName(newNotification.ExamShortName).
		SetUpdatedAt(time.Now().UTC().Truncate(24 * time.Hour)).
		SetUpdatedBy(newNotification.UserName).
		SetOrderDate(newNotification.OrderDate).
		SetExamYear(newNotification.ExamYear).
		SetPapers(newNotification.Papers).
		SetEmployeeMasterRequestLastDate(newNotification.EmployeeMasterRequestLastDate).
		SetEmployeeMasterRequestApprovalLastDate(newNotification.EmployeeMasterRequestApprovalLastDate).
		SetExamRegisterLastDate(newNotification.ExamRegisterLastDate).
		SetApplicationStartDate(newNotification.ApplicationStartDate).
		SetApplicationEndDate(newNotification.ApplicationEndDate).
		SetApplicationVerificationLastDate(newNotification.ApplicationVerificationLastDate).
		SetApplicationCorrectionLastDate(newNotification.ApplicationCorrectionLastDate).
		SetResubmittedApplicationVerificationDate(newNotification.ResubmittedApplicationVerificationDate).
		SetNodalOfficerApprovalDate(newNotification.NodalOfficerApprovalDate).
		SetAdmitCardDate(newNotification.AdmitCardDate).
		SetApprovedBy(newNotification.ApprovedBy).
		SetApprovedById(ud.ID).
		SetApprovedByDesignation(newNotification.ApprovedByDesignation).
		SetApprovedByName(newNotification.ApprovedByName).
		SetCrucialDate(newNotification.CrucialDate).
		SetSmsExamShortName(examcode.SmsExamShortName).
		SetNotificationOrderNumber(newNotification.NotificationOrderNumber).
		SetCenterAllotmentEndDate(newNotification.CenterAllotmentEndDate).
		SetCircleOfficeFacilityId(newNotification.CircleOfficeFacilityId).
		SetCircleOfficeName(newNotification.CircleOfficeName).
		SetStatus("active").
		Save(ctx)

	if err3 != nil {
		log.Println("Error at Cancelling Notification: ", newNotification)
		return "", fmt.Errorf(" failed to Cancel the Notification")
	}

	if err := tx.Commit(); err != nil {
		log.Println("Failed to commit transaction:", err)
		return "", fmt.Errorf(" failed to commit transaction: %v", err)
	}

	message := "Notification Cancelled"
	return message, nil
}

// NEW API FOR GETTING LIVE NOTIFICATION DATA
func GetAllexamNotificationUpdated(client *ent.Client, circleid string) ([]*ent.ExamNotifications, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	in, err := client.ExamNotifications.
		Query().
		Where(
			//examnotifications.FlagNEQ(true),
			examnotifications.NotificationStatusNEQ("NotificationIssued"),
			examnotifications.CircleOfficeFacilityIdEQ(circleid),
			examnotifications.StatusEQ("active")).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	if len(in) == 0 {
		return nil, 422, " -STR002", false, errors.New("no matching data for this Circle with given exam year in notification issued status")
	}

	return in, 200, "", true, nil
}

// Function to re-issue notification
func PutReIssueNotification(client *ent.Client, newNotification *ent.ExamNotifications) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examid := newNotification.ExamCode
	status := newNotification.NotificationStatus
	renotstatus := newNotification.NotificationReIssueStatus
	username := newNotification.UserName
	examyear := newNotification.ExamYear
	notificationNumber := newNotification.NotificationNumber

	if examid <= 0 {
		return "", fmt.Errorf("invalid Exam Code: %d", examid)
	}
	// Start transaction
	tx, err := client.Tx(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Printf("Failed to rollback transaction: %v", rerr)
			}
		}
	}()

	examcode, err := tx.Exam.
		Query().
		Where(exam.ID(examid)).
		Only(ctx)

	if err != nil {
		log.Println("Error querying ExamCode :", err)
		return "", fmt.Errorf(" no exams found for the given exam Code: %w", err)
	}

	if examcode == nil {
		return "", fmt.Errorf(" exam Details not found with the Exam Code: %d", examid)
	}

	// examdetail, err := client.Exam.Query().Where(exam.IDEQ(examid)).Only(ctx)
	// if err != nil {
	// 	fmt.Printf("failed to retrieve data: %v", err)
	// 	return nil, fmt.Errorf("failed to retrieve data: %w", err)
	// }
	// examname := examdetail.ExamName

	exists, nerr := tx.AdminMaster.Query().Where(adminmaster.UserName(username)).Exist(ctx)
	if nerr != nil {
		log.Println("Error checking username existence:", err)
		return "", fmt.Errorf("failed to check username existence: %w", err)
	}

	if !exists {
		return "", fmt.Errorf(" invalid username: %s", username)
	}

	originalNotification, err := tx.ExamNotifications.
		Query().
		Where(
			//examnotifications.FlagNEQ(true),
			examnotifications.NotificationNumberEQ(notificationNumber),
			examnotifications.ExamCodeEQ(examid),
			examnotifications.ExamYearEQ(examyear),
			examnotifications.NotificationStatusEQ("NotificationIssued")).
		Only(ctx)

	if err != nil {
		log.Println("Failed to check existing Notification:", err)
		return "", fmt.Errorf("failed to check existing Notification: %v", err)
	}

	if originalNotification == nil {
		return "", fmt.Errorf(" the Issued Notification is not available")
	}

	currentTime := time.Now().Truncate(time.Second)
	st, err3 := tx.ExamNotifications.Update().SetFlag(true).Where(examnotifications.ExamCodeEQ(newNotification.ExamCode), examnotifications.ExamYearEQ(newNotification.ExamYear)).Save(ctx)
	if err3 != nil {
		log.Println("Failed to check existing Draft Notification:", err)
		return "", fmt.Errorf("failed to check existing Draft Notification: %v", err)
	}
	fmt.Println(st)
	duplicatedNotification := tx.ExamNotifications.
		Create().
		SetExamCode(originalNotification.ExamCode).
		SetExamName(originalNotification.ExamName).
		SetUserName(username).
		SetExamYear(originalNotification.ExamYear).
		SetNotificationRemarks(originalNotification.NotificationRemarks).
		SetEmployeeMasterRequestLastDate(originalNotification.EmployeeMasterRequestLastDate).
		SetEmployeeMasterRequestApprovalLastDate(originalNotification.EmployeeMasterRequestApprovalLastDate).
		SetExamRegisterLastDate(originalNotification.ExamRegisterLastDate).
		SetApplicationStartDate(originalNotification.ApplicationStartDate).
		SetApplicationEndDate(originalNotification.ApplicationEndDate).
		SetApplicationCorrectionStartDate(originalNotification.ApplicationCorrectionStartDate).
		SetApplicationCorrectionLastDate(originalNotification.ApplicationCorrectionLastDate).
		SetApplicationVerificationLastDate(originalNotification.ApplicationVerificationLastDate).
		SetNodalOfficerApprovalDate(originalNotification.NodalOfficerApprovalDate).
		SetAdmitCardDate(originalNotification.AdmitCardDate).
		SetResubmittedApplicationVerificationDate(originalNotification.ResubmittedApplicationVerificationDate).
		SetPapers(originalNotification.Papers).
		SetNotificationStatus(status).
		SetNotificationReIssueStatus(false).
		SetEditFlagStatus(originalNotification.EditFlagStatus).
		SetDesignation(originalNotification.Designation).
		SetOfficerName(originalNotification.OfficerName).
		SetNotificationOrderNumber(originalNotification.NotificationOrderNumber).
		SetNotesheetScannedCopy(originalNotification.NotesheetScannedCopy).
		SetOrderDate(originalNotification.OrderDate).
		SetNotificationNumber(originalNotification.NotificationNumber).
		SetUpdatedAt(currentTime). // Set the current UTC time
		SetUpdatedBy(username).    // Set the updated by user (replace with the actual user)
		SetStatus("active").
		SetCircleOfficeFacilityId(newNotification.CircleOfficeFacilityId).
		SetCircleOfficeName(newNotification.CircleOfficeName).
		SetNotificationNumber(newNotification.NotificationNumber).
		SetNotificationOrderNumber(newNotification.NotificationOrderNumber).
		SaveX(ctx)

	if duplicatedNotification == nil {
		return "", fmt.Errorf(" failed to duplicate Original Notification")
	}

	_, nerr1 := tx.ExamNotifications.Update().
		Where(
			examnotifications.NotificationNumberEQ(notificationNumber),
			examnotifications.StatusEQ("active"),
		//examnotifications.FlagNEQ(true)
		).
		SetUserName(username).
		SetNotificationReIssueStatus(renotstatus).
		SetUpdatedAt(currentTime). // Set the current UTC time
		SetUpdatedBy(username).    // Set the updated by user (replace with the actual user)
		Save(ctx)
	if nerr1 != nil {
		log.Println("error at Updating original Notification: ", newNotification)
		return "", fmt.Errorf("failed Updating original Notification: %w", err)
	}
	// Commit transaction
	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}
	message := "notification Re-Issued"
	return message, nil
}

// func QueryActiveExamsByExamYear(client *ent.Client, examyear int32) ([]*ent.ExamNotifications, error) {
// 	ctx := context.Background()
// 	eyear := examyear

// 	ActiveNotification := make([]*ent.ExamNotifications, 0)

// 	if examyear == 0 {
// 		return nil, fmt.Errorf("Year cannot be Zero")
// 	}

// 	in, err := client.ExamNotifications.
// 		Query().
// 		Where(examnotifications.NotificationStatusEQ("NotificationIssued"), examnotifications.ExamYear(eyear)).
// 		All(ctx)

// 	if err != nil {
// 		log.Println("Failed to retrieve Active exams:", err)
// 		return nil, fmt.Errorf("failed to retrieve Active exams : %v", err)
// 	}

// 	for _, issuedNotification := range in {
// 		ActiveNotification = append(ActiveNotification, issuedNotification)
// 	}

// 	return ActiveNotification, nil
// }

func QueryActiveExamsByExamYear(client *ent.Client, examyear int32, circleOfficeId string) ([]map[string]interface{}, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	eyear := examyear

	ActiveNotification := make([]map[string]interface{}, 0)
	if examyear == 0 {
		return nil, 422, "  -STR001", false, fmt.Errorf(" year cannot be Zero")
	}

	// Query active exam notifications
	activeNotifications, err := client.ExamNotifications.
		Query().
		Where(
			examnotifications.NotificationStatusEQ("NotificationIssued"),
			examnotifications.ExamYearEQ(eyear),
			examnotifications.Or(
				examnotifications.And(
					examnotifications.CircleOfficeFacilityIdEQ(circleOfficeId),
					examnotifications.IssuedByEQ("Circle"),
				),
				examnotifications.IssuedByEQ("Directorate"),
			),
			examnotifications.StatusEQ("active"),
		).
		All(ctx)

	if err != nil {

		return nil, 422, "  -STR002", false, fmt.Errorf("failed to retrieve Active exams: %v", err)

	}
	if len(activeNotifications) == 0 {
		return nil, 422, "  -STR003", false, errors.New("no active exam notifications found")
	}

	for _, issuedNotification := range activeNotifications {
		// Validate exam code against the exam schema
		examCode := issuedNotification.ExamCode
		exam, err := client.Exam.
			Query().
			Where(exam.ExamCodeEQ(examCode)).
			Select(exam.FieldCalenderIssuedBy, exam.FieldNotificationBy, exam.FieldConductedBy).
			Only(ctx)
		if err != nil {
			// If exam code not found, skip and continue
			log.Printf("Exam with code %d not found: %v\n", examCode, err)
			continue
		}

		// Create a map to store notification and exam data
		notificationData := map[string]interface{}{
			"Exam Code":                          issuedNotification.ExamCode,
			"Exam Name":                          issuedNotification.ExamName,
			"Notification Number":                issuedNotification.NotificationNumber,
			"User Name":                          issuedNotification.UserName,
			"Exam Year":                          issuedNotification.ExamYear,
			"Application Start Date":             issuedNotification.ApplicationStartDate,
			"Application End Date":               issuedNotification.ApplicationEndDate,
			"Application Correction Start Date":  issuedNotification.ApplicationCorrectionStartDate,
			"Application Correction End Date":    issuedNotification.ApplicationCorrectionLastDate,
			"Application Verification Last Date": issuedNotification.ApplicationVerificationLastDate,
			"Center Allotment End Date":          issuedNotification.CenterAllotmentEndDate,
			"Nodal Officer Approval Date":        issuedNotification.NodalOfficerApprovalDate,
			"Admit Card Date":                    issuedNotification.AdmitCardDate,
			"Crucial Date":                       issuedNotification.CrucialDate,
			"Notification Order Number":          issuedNotification.NotificationOrderNumber,
			"Order Date":                         issuedNotification.OrderDate,
			"Exam Short Name":                    issuedNotification.ExamShortName,
			"Calender Issued By":                 exam.CalenderIssuedBy,
			"Notification By":                    exam.NotificationBy,
			"Conducted By":                       exam.ConductedBy,
			"Papers":                             issuedNotification.Papers,
		}

		// Append the combined data to the result
		ActiveNotification = append(ActiveNotification, notificationData)
	}

	return ActiveNotification, 200, "", true, nil
}

func QueryActiveExamsByExamYearWithoutCAFacilityID(client *ent.Client, examyear int32) ([]map[string]interface{}, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	eyear := examyear

	ActiveNotification := make([]map[string]interface{}, 0)
	if examyear == 0 {
		return nil, 422, "  -STR001", false, fmt.Errorf("year cannot be Zero")
	}

	// Query active exam notifications
	activeNotifications, err := client.ExamNotifications.
		Query().
		Where(examnotifications.NotificationStatusEQ("NotificationIssued"), examnotifications.ExamYear(eyear), examnotifications.StatusEQ("active")).
		All(ctx)

	if err != nil {
		return nil, 422, "  -STR002", false, fmt.Errorf("failed to retrieve Active exams: %v", err)
	}

	for _, issuedNotification := range activeNotifications {
		// Validate exam code against the exam schema
		examCode := issuedNotification.ExamCode
		exam, err := client.Exam.
			Query().
			Where(exam.ExamCodeEQ(examCode)).
			Select(exam.FieldCalenderIssuedBy, exam.FieldNotificationBy, exam.FieldConductedBy).
			Only(ctx)
		if err != nil {
			// If exam code not found, skip and continue
			log.Printf("Exam with code %d not found: %v\n", examCode, err)
			continue
		}

		// Create a map to store notification and exam data
		notificationData := map[string]interface{}{
			"CircleOfficeID":                     issuedNotification.CircleOfficeFacilityId,
			"CircleOfficeName":                   issuedNotification.CircleOfficeName,
			"Exam Code":                          issuedNotification.ExamCode,
			"Exam Name":                          issuedNotification.ExamName,
			"Notification Number":                issuedNotification.NotificationNumber,
			"User Name":                          issuedNotification.UserName,
			"Exam Year":                          issuedNotification.ExamYear,
			"Application Start Date":             issuedNotification.ApplicationStartDate,
			"Application End Date":               issuedNotification.ApplicationEndDate,
			"Application Correction Start Date":  issuedNotification.ApplicationCorrectionStartDate,
			"Application Correction End Date":    issuedNotification.ApplicationCorrectionLastDate,
			"Application Verification Last Date": issuedNotification.ApplicationVerificationLastDate,
			"Center Allotment End Date":          issuedNotification.CenterAllotmentEndDate,
			"Nodal Officer Approval Date":        issuedNotification.NodalOfficerApprovalDate,
			"Admit Card Date":                    issuedNotification.AdmitCardDate,
			"Crucial Date":                       issuedNotification.CrucialDate,
			"Notification Order Number":          issuedNotification.NotificationOrderNumber,
			"Order Date":                         issuedNotification.OrderDate,
			"Exam Short Name":                    issuedNotification.ExamShortName,
			"Calender Issued By":                 exam.CalenderIssuedBy,
			"Notification By":                    exam.NotificationBy,
			"Conducted By":                       exam.ConductedBy,
			"Papers":                             issuedNotification.Papers,
		}

		// Append the combined data to the result
		ActiveNotification = append(ActiveNotification, notificationData)
	}

	return ActiveNotification, 200, "", true, nil
}

func QueryExamByCode(client *ent.Client, examCode int32) (*ent.Exam, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Query Exam by Exam Code
	exam, err := client.Exam.
		Query().
		Where(exam.ExamCodeEQ(examCode)).
		First(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve exam: %v", err)
	}

	return exam, nil
}
