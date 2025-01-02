package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/util"

	"recruit/ent/exam"
	"recruit/ent/exampapers"
)

func CreateExam(client *ent.Client, newExam *ent.Exam) (*ent.Exam, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.Exam.
		Create().
		SetExamName(newExam.ExamName).
		SetExamShortName(newExam.ExamShortName).
		SetExamType(newExam.ExamType).
		SetNotificationBy(newExam.NotificationBy).
		SetNumOfPapers(newExam.NumOfPapers).
		SetNodalOfficerLevel(newExam.NodalOfficerLevel).
		SetConductedBy(newExam.ConductedBy).
		SetTentativeNotificationMandatoryDate(newExam.TentativeNotificationMandatoryDate).
		SetLocalLanguage(newExam.LocalLanguage).
		SetOptionForPost(newExam.OptionForPost).
		SetOptionToWriteExamOtherThanParent(newExam.OptionToWriteExamOtherThanParent).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating exam: ", newExam)
		return nil, fmt.Errorf("failed creating Exam: %w", err)
	}
	log.Println("Exam was created: ", u)

	return u, nil
}

func QueryExamByName(ctx context.Context, client *ent.Client, examname string) (*ent.Exam, error) {
	u, err := client.Exam.Query().Where(exam.ExamName(examname)).Only(ctx)

	if err != nil {
		log.Println("error at gettting examname: ", err)
		return nil, fmt.Errorf("failed querying exam: %w", err)
	}
	log.Println("exam returned by name: ", u)
	return u, nil
}

func QueryExamID(ctx context.Context, client *ent.Client, id int32) (*ent.Exam, error) {
	//Can use GetX as well
	exam, err := client.Exam.Get(ctx, id)
	if err != nil {
		log.Println("error at getting examid: ", err)
		return nil, fmt.Errorf("failed querying exam: %w", err)
	}
	log.Println("exam returned: ", exam)
	return exam, nil
}

func QueryExam(ctx context.Context, client *ent.Client) ([]*ent.Exam, int32, string, bool, error) {
	//Array of exams
	exam, err := client.Exam.Query().
		WithPapers(func(q *ent.ExamPapersQuery) {
			q.Select(exampapers.FieldID, exampapers.FieldPaperDescription)
		}).
		All(ctx)
	if err != nil {
		log.Println("error at exams: ", err)
		return nil, 422, "  -STR001", false, fmt.Errorf("failed querying exams: %w", err)
	}
	//log.Println("exams returned: ", exam)
	return exam, 200, "", true, nil
}

func QueryExamIDAndNames(ctx context.Context, client *ent.Client) ([]*ent.Exam, error) {
	u, err := client.Exam.Query().Select(exam.FieldID, exam.FieldExamName).All(ctx)

	if err != nil {
		log.Println("error at gettting examname: ", err)
		return nil, fmt.Errorf("failed querying exam: %w", err)
	}
	log.Println("exam returned by name: ", u)
	return u, nil
}

func DeleteExamID(client *ent.Client, id int32) error {

	//context not passed for delete dont know why?
	err := client.Exam.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func UpdateExam(client *ent.Client, id int32, newExam *ent.Exam) (*ent.Exam, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	_, err := QueryExamID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	updatedExam, err := client.Exam.UpdateOneID(id).
		SetExamName(newExam.ExamName).
		SetExamShortName(newExam.ExamShortName).
		SetExamType(newExam.ExamType).
		SetNotificationBy(newExam.NotificationBy).
		SetNumOfPapers(newExam.NumOfPapers).
		SetNodalOfficerLevel(newExam.NodalOfficerLevel).
		SetConductedBy(newExam.ConductedBy).
		SetTentativeNotificationMandatoryDate(newExam.TentativeNotificationMandatoryDate).
		SetLocalLanguage(newExam.LocalLanguage).
		SetOptionForPost(newExam.OptionForPost).
		SetOptionToWriteExamOtherThanParent(newExam.OptionToWriteExamOtherThanParent).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedExam, nil
}
