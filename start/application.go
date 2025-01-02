package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/exampapers"
	"recruit/util"
)

func CreateApplication(client *ent.Client, newApplication *ent.Application) (*ent.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.Application.Create().
		SetEmployeeID(newApplication.EmployeeID).
		SetNotifyCode(newApplication.NotifyCode).
		//SetHallTicketNumber(newApplication.HallTicketNumber).
		//SetCenterCode(newApplication.CenterCode).
		SetAppliedStamp(newApplication.AppliedStamp).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Application: ", newApplication)
		return nil, fmt.Errorf("failed creating Application: %w", err)
	}
	log.Println("Application was created: ", u)

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

func QueryApplicationID(ctx context.Context, client *ent.Client, id int32) (*ent.Application, error) {
	//Can use GetX as well

	u, err := client.Application.Get(ctx, id)
	if err != nil {
		log.Println("error at getting ApplicationID: ", err)
		return nil, fmt.Errorf("failed querying Application: %w", err)
	}
	log.Println("Application returned: ", u)
	return u, nil
}

func QueryApplication(ctx context.Context, client *ent.Client) ([]*ent.Application, error) {
	//Array of exams
	u, err := client.Application.Query().All(ctx)
	if err != nil {
		log.Println("error at Application: ", err)
		return nil, fmt.Errorf("failed querying Application: %w", err)
	}
	log.Println("Application returned: ", u)
	return u, nil
}

func DeleteApplicationID(client *ent.Client, id int32) error {

	//context not passed for delete dont know why?
	err := client.Application.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func UpdateApplicationCenter(client *ent.Client, id int32, newApplication *ent.Application) (*ent.Application, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	_, err := QueryApplicationID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	u, err := client.Application.UpdateOneID(id).
		//SetEmployeeID(newApplication.EmployeeID).
		//SetNotifyCode(newApplication.NotifyCode).
		//SetHallTicketNumber(newApplication.HallTicketNumber).
		SetCenterCode(newApplication.CenterCode).
		//SetAppliedStamp(newApplication.AppliedStamp).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UpdateApplicationHallTicket(client *ent.Client, id int32, newApplication *ent.Application) (*ent.Application, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	_, err := QueryApplicationID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	u, err := client.Application.UpdateOneID(id).
		//SetEmployeeID(newApplication.EmployeeID).
		//SetNotifyCode(newApplication.NotifyCode).
		SetHallTicketNumber(newApplication.HallTicketNumber).
		//SetCenterCode(newApplication.CenterCode).
		//SetAppliedStamp(newApplication.AppliedStamp).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func PaperDetails(client *ent.Client, paperlist *ent.ExamPapers, eid int32) ([]*ent.ExamPapers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	examid := eid

	papers, err := client.ExamPapers.Query().Where(exampapers.ExamCodeEQ(examid)).All(ctx)
	if err != nil {
		log.Println("error at Center cities Master: ", err)
		return nil, fmt.Errorf("failed querying Center cities Master: %w", err)
	}
	return papers, nil
}
