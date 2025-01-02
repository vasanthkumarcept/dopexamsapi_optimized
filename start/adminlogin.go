package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/adminlogin"
	"recruit/util"
	//"strings"
)

func CreateAdminLogin(client *ent.Client, newadminlogin *ent.AdminLogin) (*ent.AdminLogin, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.AdminLogin.
		Create().
		SetRoleUserCode(newadminlogin.RoleUserCode).
		SetRoleName(newadminlogin.RoleName).
		SetCreatedDate(newadminlogin.CreatedDate).
		SetStatus(newadminlogin.Status).
		SetEmployeedID(newadminlogin.EmployeedID).
		SetEmployeeName(newadminlogin.EmployeeName).
		SetEmailid(newadminlogin.Emailid).
		SetMobileNumber(newadminlogin.MobileNumber).
		SetOTP(newadminlogin.OTP).
		//.AddCircleRefIDs().
		Save(ctx)
	if err != nil {
		log.Println("error at Creating RoleMaster Master: ", newadminlogin)
		return nil, fmt.Errorf("failed creating RoleMaster Master: %w", err)
	}
	log.Println("Region was created: ", u)

	return u, nil
}

func QueryAdminLogin(ctx context.Context, client *ent.Client) ([]*ent.AdminLogin, error) {
	//Array of exams
	roles, err := client.AdminLogin.Query().
		All(ctx)
	if err != nil {
		log.Println("error at AdminLogins: ", err)
		return nil, fmt.Errorf("failed querying AdminLogins: %w", err)
	}
	log.Println("AdminLogin returned: ", roles)
	return roles, nil
}

func QueryAdminLoginByID(ctx context.Context, client *ent.Client, id int32) (*ent.AdminLogin, error) {
	//Can use GetX as well

	AdminLogins, err := client.AdminLogin.Get(ctx, id)
	if err != nil {

		log.Println("error at getting AdminLogin ID: ", err)
		return nil, fmt.Errorf("failed querying AdminLogin: %w", err)
	}
	log.Println("AdminLogin details returned: ", AdminLogins)
	return AdminLogins, nil
}

//Validate Login

func ValidateAdminLoginUser(client *ent.Client, newadminlogin *ent.AdminLogin) (*ent.AdminLogin, error) {
	// Check if the username exists
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	username := newadminlogin.Username
	roleName := newadminlogin.RoleName // Check if the username exists
	exists, err := client.AdminLogin.
		Query().
		Where(adminlogin.Username(username)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check username existence: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("no such username: %s", username)
	}
	password := newadminlogin.Password

	// Retrieve the admin login record with the provided username
	adminLogin, err := client.AdminLogin.
		Query().
		Where(adminlogin.Username(username)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve admin login record: %v", err)
	}

	if roleName == "DirectorateUser" {
		// User is a valid Directorate user
		if adminLogin.Username == username && adminLogin.Password == password && adminLogin.RoleName == roleName {
			fmt.Printf("Verified successfully\nUsername: %s\nRoleUserCode: %d\nRoleName: %s\n",
				adminLogin.Username, adminLogin.RoleUserCode, adminLogin.RoleName)
			adminLogin, err = adminLogin.Update().
				SetVerifyRemarks("Verified Successfully").
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to update admin login: %v", err)
			}
			return adminLogin, nil

		} else {
			return nil, fmt.Errorf("incorrect username or password for Directorate user")
		}
	} else if roleName == "DirectorateApprover" {
		// User is a valid Directorate Approver
		if adminLogin.Username == username && adminLogin.Password == password && adminLogin.RoleName == roleName {
			fmt.Printf("Verified successfully\nUsername: %s\nRoleUserCode: %d\nRoleName: %s\n",
				adminLogin.Username, adminLogin.RoleUserCode, adminLogin.RoleName)
			adminLogin, err = adminLogin.Update().
				SetVerifyRemarks("Verified Successfully").
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to update admin login: %v", err)
			}
			return adminLogin, nil
		} else {
			return nil, fmt.Errorf("incorrect username or password for Directorate Approver")
		}
	} else if roleName == "ControllingAuthority" {
		// User is a valid Controlling Authority
		if adminLogin.Password == password {
			fmt.Printf("Verified successfully\nUsername: %s\nRoleUserCode: %d\nRoleName: %s\n",
				adminLogin.Username, adminLogin.RoleUserCode, adminLogin.RoleName)
			adminLogin, err = adminLogin.Update().
				SetVerifyRemarks("Verified Successfully").
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to update admin login: %v", err)
			}
			return adminLogin, nil
		} else {
			return nil, fmt.Errorf("incorrect password for Controlling Authority user")
		}
	} else if roleName == "NodalOfficer" {
		// User is a valid Nodal Officer
		if adminLogin.Password == password {

			fmt.Printf("Verified successfully\nUsername: %s\nRoleUserCode: %d\nRoleName: %s\n",
				adminLogin.Username, adminLogin.RoleUserCode, adminLogin.RoleName)
			adminLogin, err = adminLogin.Update().
				SetVerifyRemarks("Verified Successfully").
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to update admin login: %v", err)
			}
			return adminLogin, nil
		} else {
			return nil, fmt.Errorf("incorrect password for Nodal Officer user")
		}
	} else {
		// User role is not valid
		return nil, fmt.Errorf("not a valid user role")
	}
}
