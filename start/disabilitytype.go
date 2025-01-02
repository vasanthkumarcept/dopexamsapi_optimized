package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/util"
)

func CreateEmployeeDisability(client *ent.Client, newDisability *ent.Disability) (*ent.Disability, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.Disability.
		Create().
		SetDisabilityTypeCode(newDisability.DisabilityTypeCode).
		SetDisabilityTypeDescription(newDisability.DisabilityTypeDescription).
		SetDisabilityPercentage(newDisability.DisabilityPercentage).
		SetDisabilityFlag(newDisability.DisabilityFlag).
		//SetPaperCode(newDisability.PaperCode).

		Save(ctx)
	if err != nil {
		log.Println("error at Creating Employee Disability: ", newDisability)
		return nil, fmt.Errorf("failed creating employee Disability: %w", err)
	}
	log.Println("Employee Disability was created: ", u)

	return u, nil
}

func QueryEmployeeDisabilityID(ctx context.Context, client *ent.Client, id int32) (*ent.Disability, error) {
	//Can use GetX as well
	empdisability, err := client.Disability.Get(ctx, id)
	if err != nil {
		log.Println("error at getting Disability: ", err)
		return nil, fmt.Errorf("failed querying employee Disability: %w", err)
	}
	log.Println("Employee Disability returned: ", empdisability)
	return empdisability, nil
}

func QueryEmployeeDisabilities(ctx context.Context, client *ent.Client) ([]*ent.Disability, error) {
	//Array of exams
	employeeDisability, err := client.Disability.Query().
		All(ctx)
	if err != nil {
		log.Println("error at exams: ", err)
		return nil, fmt.Errorf("failed querying Employee Disability: %w", err)
	}
	log.Println("Employee Disability returned: ", employeeDisability)
	return employeeDisability, nil
}
