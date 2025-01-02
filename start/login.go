package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/login"
)

func QueryLoginByEmpID(ctx context.Context, client *ent.Client, empid int32) ([]*ent.Login, error) {
	//Can use GetX as well

	newlogin, err := client.Login.Query().
		Where(login.EmployeedIDEQ(empid)).
		All(ctx)

	if err != nil {
		log.Println("error at gettting login details: ", err)
		return nil, fmt.Errorf("failed login details: %w", err)
	}
	log.Println("details returned by login : ", newlogin)
	return newlogin, nil
}
