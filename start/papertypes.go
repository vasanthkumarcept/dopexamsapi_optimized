package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/papertypes"
	"recruit/util"
)

//    field.Time("HallTicketDownloadDate").Optional(), field.String("NotifyFile").Optional(), field.String("SyllabusFile").Optional(), field.String("VacanciesFile").Optional()}

func CreatePaperType(client *ent.Client, newPapers *ent.PaperTypes) (*ent.PaperTypes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.PaperTypes.Create().
		SetPaperCode(newPapers.PaperCode).
		SetPaperTypeDescription(newPapers.PaperTypeDescription).
		SetOrderNumber(newPapers.OrderNumber).
		SetSequenceNumber(newPapers.SequenceNumber).
		SetCreatedDate(newPapers.CreatedDate).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Exam Paper Types: ", newPapers)
		return nil, fmt.Errorf("failed creating Exam Paper Types: %w", err)
	}
	log.Println("Exam Papers was created: ", u)

	return u, nil
}

func QueryExamPaperTypeByID(ctx context.Context, client *ent.Client, id int32) (*ent.PaperTypes, error) {
	//Can use GetX as well

	Exam_Papertypes, err := client.PaperTypes.Get(ctx, id)
	if err != nil {
		log.Println("error at getting Exam Paper Type ID: ", err)
		return nil, fmt.Errorf("failed querying Exam Paper Type: %w", err)
	}
	log.Println("Exam Paper Type details returned: ", Exam_Papertypes)
	return Exam_Papertypes, nil
}

func QueryExamPaperTypes(ctx context.Context, client *ent.Client) ([]*ent.PaperTypes, error) {
	//Array of exams
	PaperTypes, err := client.PaperTypes.Query().All(ctx)
	if err != nil {
		log.Println("error at PaperTypes: ", err)
		return nil, fmt.Errorf("failed querying PaperTypes: %w", err)
	}
	log.Println("Notifications returned: ", PaperTypes)
	return PaperTypes, nil
}

func QueryExamPaperTypesByPaperCode(ctx context.Context, client *ent.Client, PaperCode int32) ([]*ent.PaperTypes, error) {
	//Can use GetX as well
	papertypes, err := client.PaperTypes.Query().
		Select(papertypes.FieldID, papertypes.FieldPaperTypeDescription).
		Where(papertypes.PaperCodeEQ(PaperCode)).
		All(ctx)

	if err != nil {
		log.Println("error at getting paper code: ", err)
		return nil, fmt.Errorf("failed querying exam paper types: %w", err)
	}
	log.Println("paper types returned: ", papertypes)
	if len(papertypes) == 0 {
		return nil, fmt.Errorf("no Exam Paper Types found with Paper Code: %d", PaperCode)
	}
	for _, paper := range papertypes {
		fmt.Printf("ID: %d\n", paper.ID)
		fmt.Printf("PaperTypeDescription: %s\n", paper.PaperTypeDescription)
	}
	return papertypes, nil
}
