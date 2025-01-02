package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/util"
)

func CreateEmployeeCategory(client *ent.Client, newcategory *ent.EmployeeCategory) (*ent.EmployeeCategory, int32, string, bool, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.EmployeeCategory.
		Create().
		//SetCategrycode(newcategory.Categrycode).
		SetCategoryDescription(newcategory.CategoryDescription).
		SetMinimumMarks(newcategory.MinimumMarks).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	//log.Println("Employee category was created: ", u)

	return u, 200, "", true, nil
}

func QueryEmployeeCategoryID(ctx context.Context, client *ent.Client, id int32) (*ent.EmployeeCategory, error) {
	//Can use GetX as well
	empcat, err := client.EmployeeCategory.Get(ctx, id)
	if err != nil {
		log.Println("error at getting examid: ", err)
		return nil, fmt.Errorf("failed querying employee category: %w", err)
	}
	//log.Println("EmployeeCategory returned: ", empcat)
	return empcat, nil
}

func QueryEmployeeCategories(ctx context.Context, client *ent.Client) ([]*ent.EmployeeCategory, int32, string, bool, error) {
	//Array of exams
	employeecategory, err := client.EmployeeCategory.Query().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	} else {
		if len(employeecategory) == 0 {
			return nil, 422, " -STR002", false, errors.New("no employee category found")
		}
	}
	return employeecategory, 200, "", true, nil
}
