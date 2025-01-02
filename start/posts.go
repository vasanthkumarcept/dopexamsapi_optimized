package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/util"
)

func CreateEmployeePosts(client *ent.Client, newposts *ent.EmployeePosts) (*ent.EmployeePosts, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.EmployeePosts.
		Create().
		SetPostCode(newposts.PostCode).
		SetPostDescription(newposts.PostDescription).
		SetGroup(newposts.Group).
		SetPayLevel(newposts.PayLevel).
		SetScale(newposts.Scale).
		SetBaseCadreFlag(newposts.BaseCadreFlag).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Employee Posts: ", newposts)
		return nil, fmt.Errorf("failed creating employee Posts: %w", err)
	}
	log.Println("Employee Posts was created: ", u)

	return u, nil
}

func QueryEmployeePostsID(ctx context.Context, client *ent.Client, id int32) (*ent.EmployeePosts, error) {
	//Can use GetX as well
	empposts, err := client.EmployeePosts.Get(ctx, id)
	if err != nil {
		log.Println("error at getting posts: ", err)
		return nil, fmt.Errorf("failed querying employee posts: %w", err)
	}
	log.Println("Employee posts returned: ", empposts)
	return empposts, nil
}

func QueryEmployeePosts(ctx context.Context, client *ent.Client) ([]*ent.EmployeePosts, error) {
	//Array of exams
	employeePosts, err := client.EmployeePosts.Query().
		All(ctx)
	if err != nil {
		log.Println("error at exams: ", err)
		return nil, fmt.Errorf("failed querying  Employee Posts: %w", err)
	}
	log.Println("Employee Designation returned: ", employeePosts)
	return employeePosts, nil
}
