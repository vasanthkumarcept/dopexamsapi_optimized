package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/util"
)

func CreateEmployeeDesignation(client *ent.Client, newdesignation *ent.EmployeeDesignation) (*ent.EmployeeDesignation, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.EmployeeDesignation.
		Create().
		SetDesignationCode(newdesignation.DesignationCode).
		SetDesignationDescription(newdesignation.DesignationDescription).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Employee designation: ", newdesignation)
		return nil, fmt.Errorf("failed creating employee designation: %w", err)
	}
	log.Println("Employee designation was created: ", u)

	return u, nil
}

func QueryEmployeeDesignationID(ctx context.Context, client *ent.Client, id int32) (*ent.EmployeeDesignation, error) {
	//Can use GetX as well
	empdesgn, err := client.EmployeeDesignation.Get(ctx, id)
	if err != nil {
		log.Println("error at getting designation: ", err)
		return nil, fmt.Errorf("failed querying employee designation: %w", err)
	}
	log.Println("Employee designation returned: ", empdesgn)
	return empdesgn, nil
}

func QueryEmployeeDesignations(ctx context.Context, client *ent.Client) ([]*ent.EmployeeDesignation, error) {
	//Array of exams
	employeeDesignation, err := client.EmployeeDesignation.Query().
		All(ctx)
	if err != nil {
		log.Println("error at exams: ", err)
		return nil, fmt.Errorf("failed querying Employee Designation: %w", err)
	}
	log.Println("Employee Designation returned: ", employeeDesignation)
	return employeeDesignation, nil
}
