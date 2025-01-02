package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/exampapers"
	"recruit/util"
)

//    field.Time("HallTicketDownloadDate").Optional(), field.String("NotifyFile").Optional(), field.String("SyllabusFile").Optional(), field.String("VacanciesFile").Optional()}

func CreatePapers(client *ent.Client, newPapers *ent.ExamPapers) (*ent.ExamPapers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.ExamPapers.Create().
		SetPaperDescription(newPapers.PaperDescription).
		SetExamCode(newPapers.ExamCode).
		SetPaperStatus(newPapers.PaperStatus).
		SetCompetitiveQualifying(newPapers.CompetitiveQualifying).
		SetExceptionForDisability(newPapers.ExceptionForDisability).
		SetMaximumMarks(newPapers.MaximumMarks).
		SetDuration(newPapers.Duration).
		SetLocalLanguageAllowedAnswerPaper(newPapers.LocalLanguageAllowedAnswerPaper).
		SetLocalLanguageAllowedQuestionPaper(newPapers.LocalLanguageAllowedQuestionPaper).
		SetOrderNumber(newPapers.OrderNumber).
		SetCreatedDate(newPapers.CreatedDate).
		SetDisabilityTypeID(newPapers.DisabilityTypeID).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Exam Papers: ", newPapers)
		return nil, fmt.Errorf("failed creating Exam Papers: %w", err)
	}
	log.Println("Exam Papers was created: ", u)

	return u, nil
}

func QueryExamPaperByID(ctx context.Context, client *ent.Client, id int32) (*ent.ExamPapers, error) {
	//Can use GetX as well

	Exam_Papers, err := client.ExamPapers.Get(ctx, id)
	if err != nil {
		log.Println("error at getting Exam Paper ID: ", err)
		return nil, fmt.Errorf("failed querying Exam Papers: %w", err)
	}
	log.Println("Exam Paper details returned: ", Exam_Papers)
	return Exam_Papers, nil
}

func QueryExamPapers(ctx context.Context, client *ent.Client) ([]*ent.ExamPapers, error) {
	//Array of exams
	ExamPaper, err := client.ExamPapers.Query().All(ctx)
	if err != nil {
		log.Println("error at ExamPapers: ", err)
		return nil, fmt.Errorf("failed querying ExamPapers: %w", err)
	}
	log.Println("Notifications returned: ", ExamPaper)
	return ExamPaper, nil
}

func QueryExamPapersByExamCode(ctx context.Context, client *ent.Client, examcode int32) ([]*ent.ExamPapers, error) {
	//Can use GetX as well
	exam, err := client.ExamPapers.Query().
		Select(exampapers.FieldID, exampapers.FieldPaperDescription).
		Where(exampapers.ExamCodeEQ(examcode)).
		All(ctx)

	if err != nil {
		log.Println("error at getting examid: ", err)
		return nil, fmt.Errorf("failed querying exam papers: %w", err)
	}
	log.Println("exam returned: ", exam)
	if len(exam) == 0 {
		return nil, fmt.Errorf("no Exam Papers found with Exam Code: %d", examcode)
	}
	for _, paper := range exam {
		fmt.Printf("ID: %d\n", paper.ID)
		fmt.Printf("PaperDescription: %s\n", paper.PaperDescription)
	}
	return exam, nil
}
