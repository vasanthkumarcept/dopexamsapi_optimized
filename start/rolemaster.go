package start

import (
	"context"
	"errors"
	"fmt"
	"log"

	//"net/http"
	"recruit/ent"
	"recruit/util"
	//"strconv"
	//"github.com/gin-gonic/gin"
)

func CreateRoles(client *ent.Client, newroles *ent.RoleMaster) (*ent.RoleMaster, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.RoleMaster.
		Create().
		SetRoleName(newroles.RoleName).
		SetStatus(newroles.Status).
		SetCreatedDate(newroles.CreatedDate).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating roles in role master: ", newroles)
		return nil, fmt.Errorf("failed creating roles : %w", err)
	}
	log.Println("New role is created: ", u)

	return u, nil
}

func QueryRolesByID(ctx context.Context, client *ent.Client, roleid int32) (*ent.RoleMaster, error) {
	//Can use GetX as well
	roles, err := client.RoleMaster.Get(ctx, roleid)
	if err != nil {
		log.Println("error at getting roles: ", err)
		return nil, fmt.Errorf("failed querying Role Masterr %w", err)
	}
	log.Println("Roles returned: ", roles)
	return roles, nil
}

func QueryRoles(ctx context.Context, client *ent.Client) ([]*ent.RoleMaster, int32, string, bool, error) {
	//Array of exams
	roles, err := client.RoleMaster.Query().
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	if len(roles) == 0 {
		return nil, 422, " -STR002", false, errors.New("no roles found")
	}
	return roles, 200, "", true, nil
}
