package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/util"
)

//    field.Time("HallTicketDownloadDate").Optional(), field.String("NotifyFile").Optional(), field.String("SyllabusFile").Optional(), field.String("VacanciesFile").Optional()}

func CreateNotification(client *ent.Client, newNotification *ent.Notification) (*ent.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.Notification.Create().
		SetExamCode(newNotification.ExamCode).
		SetExamYear(newNotification.ExamYear).
		SetApplicationStartDate(newNotification.ApplicationStartDate).
		SetApplicationEndDate(newNotification.ApplicationEndDate).
		SetVerificationDateByController(newNotification.VerificationDateByController).
		SetCorrectionDateByCandidate(newNotification.CorrectionDateByCandidate).
		SetCorrectionVeriyDateByController(newNotification.CorrectionVeriyDateByController).
		SetHallTicketAllotmentDateByNodalOfficer(newNotification.HallTicketAllotmentDateByNodalOfficer).
		SetHallTicketDownloadDate(newNotification.HallTicketDownloadDate).
		SetNotifyFile(newNotification.NotifyFile).
		SetSyllabusFile(newNotification.SyllabusFile).
		SetVacanciesFile(newNotification.VacanciesFile).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Notification: ", newNotification)
		return nil, fmt.Errorf("failed creating Notification: %w", err)
	}
	log.Println("Notification was created: ", u)

	return u, nil
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

func QueryNotificationID(ctx context.Context, client *ent.Client, id int32) (*ent.Notification, error) {
	//Can use GetX as well

	Notification, err := client.Notification.Get(ctx, id)
	if err != nil {
		log.Println("error at getting NotificationID: ", err)
		return nil, fmt.Errorf("failed querying Notification: %w", err)
	}
	log.Println("Notification returned: ", Notification)
	return Notification, nil
}

func QueryNotification(ctx context.Context, client *ent.Client) ([]*ent.Notification, error) {
	//Array of exams
	notification, err := client.Notification.Query().All(ctx)
	if err != nil {
		log.Println("error at Notification: ", err)
		return nil, fmt.Errorf("failed querying Notifications: %w", err)
	}
	log.Println("Notifications returned: ", notification)
	return notification, nil
}

func DeleteNotificationID(client *ent.Client, id int32) error {

	//context not passed for delete dont know why?
	err := client.Notification.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func UpdateNotification(client *ent.Client, id int32, newNotification *ent.Notification) (*ent.Notification, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	_, err := QueryNotificationID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	updatedNotification, err := client.Notification.UpdateOneID(id).
		SetExamCode(newNotification.ExamCode).
		SetExamYear(newNotification.ExamYear).
		SetApplicationStartDate(newNotification.ApplicationStartDate).
		SetApplicationEndDate(newNotification.ApplicationEndDate).
		SetVerificationDateByController(newNotification.VerificationDateByController).
		SetCorrectionDateByCandidate(newNotification.CorrectionDateByCandidate).
		SetCorrectionVeriyDateByController(newNotification.CorrectionVeriyDateByController).
		SetHallTicketAllotmentDateByNodalOfficer(newNotification.HallTicketAllotmentDateByNodalOfficer).
		SetHallTicketDownloadDate(newNotification.HallTicketDownloadDate).
		SetNotifyFile(newNotification.NotifyFile).
		SetSyllabusFile(newNotification.SyllabusFile).
		SetVacanciesFile(newNotification.VacanciesFile).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedNotification, nil
}
