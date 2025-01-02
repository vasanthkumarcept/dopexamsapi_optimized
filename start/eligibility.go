package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/exam"
	"recruit/ent/exampapers"
	"recruit/util"
)

func CreateEligibilities(client *ent.Client, newEligibility *ent.EligibilityMaster) (*ent.EligibilityMaster, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.EligibilityMaster.
		Create().
		SetExamCode(newEligibility.ExamCode).
		SetNotifyCode(newEligibility.NotifyCode).
		SetCategoryCode(newEligibility.CategoryCode).
		SetPostCode(newEligibility.PostCode).
		SetExamName(newEligibility.ExamName).
		SetGdsService(newEligibility.GdsService).
		SetAgeCriteria(newEligibility.AgeCriteria).
		SetServiceCriteria(newEligibility.ServiceCriteria).
		SetDrivingLicenseCriteria(newEligibility.ComputerKnowledge).
		SetLevelOfPayMatrixEligibility(newEligibility.LevelOfPayMatrixEligibility).
		SetEducation(newEligibility.Education).
		SetPaperCode(newEligibility.PaperCode).
		SetPaperDescription(newEligibility.PaperDescription).
		SetMinimumMarks(newEligibility.MinimumMarks).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating eligibilities: ", newEligibility)
		return nil, fmt.Errorf("failed creating eligibilities: %w", err)
	}
	log.Println("Eligibility was created: ", u)

	return u, nil
}

/*
func QueryEligibilitiyMasterByName(ctx context.Context, client *ent.Client, eligibilities string) (*ent.EligibilityMaster, error) {
	u, err := client.Exam.Query().Where(exam.ExamName(examname)).Only(ctx)

	if err != nil {
		log.Println("error at gettting examname: ", err)
		return nil, fmt.Errorf("failed querying exam: %w", err)
	}
	log.Println("exam returned by name: ", u)
	return u, nil
}*/

func QueryEligibilitiyMasterByID(ctx context.Context, client *ent.Client, id int32) (*ent.EligibilityMaster, error) {
	//Can use GetX as well
	eligibilities, err := client.EligibilityMaster.Get(ctx, id)
	if err != nil {
		log.Println("error at getting Eligibility Master: ", err)
		return nil, fmt.Errorf("failed querying EligibilityMaster: %w", err)
	}
	log.Println("EligibilityMaster returned: ", eligibilities)
	return eligibilities, nil
}

func QueryEligibilitiyMaster(ctx context.Context, client *ent.Client) ([]*ent.EligibilityMaster, error) {
	//Array of exams
	eligibilities, err := client.EligibilityMaster.Query().All(ctx)
	if err != nil {
		log.Println("error at Eligibility Master: ", err)
		return nil, fmt.Errorf("failed querying Eligibility Master: %w", err)
	}
	log.Println("Eligibilities returned: ", eligibilities)
	return eligibilities, nil
}

/*
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

	ctx := context.Background()
	_, err := QueryExamID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	updatedExam, err := client.Exam.UpdateOneID(id).
		SetExamName(newExam.ExamName).
		SetNotificationBy(newExam.NotificationBy).
		SetNumOfPapers(newExam.NumOfPapers).
		SetNodalOfficerLevel(newExam.NodalOfficerLevel).
		SetConductedBy(newExam.ConductedBy).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedExam, nil
}
*/

func QueryEligibilityWithPapers(ctx context.Context, client *ent.Client, examcode int32) ([]*ent.Exam, error) {
	//Array of exams
	exam, err := client.Exam.Query().
		Where(exam.IDEQ(examcode)).
		WithExamEligibility(func(q *ent.EligibilityMasterQuery) {
			q.WithExamPaperEligibility(func(q *ent.ExamPapersQuery) {
				q.Select(exampapers.FieldID, exampapers.FieldPaperDescription, exampapers.FieldCompetitiveQualifying)
			})
		}).
		All(ctx)
	if err != nil {
		log.Println("error at exams: ", err)
		return nil, fmt.Errorf("failed querying exams: %w", err)
	}
	log.Println("exams returned: ", exam)
	//if len(exam) == 0 {
	//	return nil, fmt.Errorf("no Exam Papers found with Exam Code: %d", examcode)
	//}

	return exam, nil
}
