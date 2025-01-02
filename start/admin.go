package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	a "recruit/authentication"
	"recruit/ent"
	"recruit/ent/adminmaster"
	"recruit/ent/employeemaster"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"
	"strings"
	"time"
)

func CreateAdmin(ctx context.Context, client *ent.Client, adminRequest ca_reg.AdminCreation) (*ca_reg.AdminMasterResponse, int32, string, bool, error) {
	currentTime := time.Now().Truncate(time.Second)
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}
	defer func() {
		handleTransaction(tx, &err)
	}()

	if err = checkDuplicateAdmin(ctx, tx, adminRequest); err != nil {
		return nil, 422, " -STR003", false, err
	}

	if err = checkExistingAdmin(ctx, tx, adminRequest.UserName); err != nil {
		return nil, 422, " -STR004", false, err
	}

	admin, err := createAdminRecord(ctx, tx, adminRequest, currentTime)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR027", false, err
	}

	response := mapToAdminResponse(admin)
	return response, 200, "", true, nil
}
func generateUniqueNumberAdminUser(client *ent.Client, stgString string) (int64, int32, string, error) {
	ctx := context.TODO()
	lastNumber, err := client.AdminMaster.
		Query().
		Order(ent.Desc(adminmaster.FieldUUID)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// No existing numbers, start from 100001
			return 10000001, 200, stgString + "_01", nil
		}
		return 0, 500, stgString + "_02", fmt.Errorf("failed to retrieve last application: %v", err)
	}
	return int64(lastNumber.UUID) + 1, 200, "", nil
}

func UpdateadminUser(client *ent.Client, newUser ca_reg.UpdateAdminMasterStruc) (*ca_reg.AdminMasterResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	uniqNumber, status, strError, err := generateUniqueNumberAdminUser(client, " -STR001")
	if err != nil {
		return nil, status, strError + strError, false, err
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer handleTransaction(tx, &err)
	//fmt.Println(uniqNumber)
	//id := newUser.UserName
	// Query the user by EmployeeID
	user, err := fetchActiveAdminUser(ctx, tx, newUser.UserName)
	if err != nil {
		return nil, 422, " -STR002", false, err
	}

	stat := "inactive_" + time.Now().Format("20060102150405")

	// Update the user entity with the provided new user data
	if err = updateUserStatus(user, stat, ctx); err != nil {
		return nil, 500, " -STR003", false, err
	}

	duplicateAdmin, err := checkduplicateAdmin(tx, newUser, ctx)
	if err != nil {
		return nil, 500, " -STR004", false, err
	}
	if duplicateAdmin {
		return nil, 422, " -STR005", false, fmt.Errorf("already username exists with this EmployeeId: %v , Facility ID: %v , Role User Code: %v ", newUser.EmployeeId, newUser.FacilityID, newUser.RoleUserCode)
	}
	if user.EmployeeId != newUser.EmployeeId {
		_, err = tx.AdminMaster.
			Create().
			SetEmployeeId(newUser.EmployeeId).
			SetEmployeeName(newUser.EmployeeName).
			SetDesignation(newUser.Designation).
			SetRoleUserCode(user.RoleUserCode).
			SetRoleUserDescription(newUser.RoleUserDescription).
			SetMobile(newUser.Mobile).
			SetEmailID(newUser.EmailID).
			SetUserName(newUser.UserName).
			SetFacilityIDUniqueid(user.FacilityIDUniqueid).
			SetFacilityID(newUser.FacilityID).
			SetAuthorityFacilityName(user.AuthorityFacilityName).
			SetFacilityType(user.FacilityType).
			SetReportingOfficeFacilityId(user.ReportingOfficeFacilityId).
			SetReportingOfficeFacilityName(user.ReportingOfficeFacilityName).
			SetCircleOfficeFacilityId(user.CircleOfficeFacilityId).
			SetCircleOfficeName(user.CircleOfficeName).
			SetCreatedById(user.CreatedById).
			SetCreatedByUserName(user.CreatedByUserName).
			SetCreatedByEmpId(user.CreatedByEmpId).
			SetCreatedByDesignation(user.CreatedByDesignation).
			SetCreatedDate(user.CreatedDate).
			SetModifiedById(newUser.ModifiedById).
			SetModifiedByUserName(newUser.ModifiedByUserName).
			SetModifiedByEmpId(newUser.ModifiedByEmpId).
			SetModifiedByDesignantion(newUser.ModifiedByDesignantion).
			SetModifiedDate(time.Now().Truncate(time.Second)).
			SetUUID(uniqNumber).
			SetStatuss("active").
			Save(ctx)
		if err != nil {
			return nil, 500, " -STR006", false, err
		}
	} else {
		_, err = tx.AdminMaster.
			Create().
			SetEmployeeId(newUser.EmployeeId).
			SetEmployeeName(newUser.EmployeeName).
			SetDesignation(newUser.Designation).
			SetRoleUserCode(user.RoleUserCode).
			SetRoleUserDescription(newUser.RoleUserDescription).
			SetMobile(newUser.Mobile).
			SetEmailID(newUser.EmailID).
			SetUserName(newUser.UserName).
			SetPassword(user.Password).
			SetFacilityIDUniqueid(user.FacilityIDUniqueid).
			SetFacilityID(newUser.FacilityID).
			SetAuthorityFacilityName(user.AuthorityFacilityName).
			SetFacilityType(user.FacilityType).
			SetReportingOfficeFacilityId(user.ReportingOfficeFacilityId).
			SetReportingOfficeFacilityName(user.ReportingOfficeFacilityName).
			SetCircleOfficeFacilityId(user.CircleOfficeFacilityId).
			SetCircleOfficeName(user.CircleOfficeName).
			SetCreatedById(user.CreatedById).
			SetCreatedByUserName(user.CreatedByUserName).
			SetCreatedByEmpId(user.CreatedByEmpId).
			SetCreatedByDesignation(user.CreatedByDesignation).
			SetCreatedDate(user.CreatedDate).
			SetModifiedById(newUser.ModifiedById).
			SetModifiedByUserName(newUser.ModifiedByUserName).
			SetModifiedByEmpId(newUser.ModifiedByEmpId).
			SetModifiedByDesignantion(newUser.ModifiedByDesignantion).
			SetModifiedDate(time.Now().Truncate(time.Second)).
			SetUUID(uniqNumber).
			SetStatuss("active").
			Save(ctx)
		if err != nil {
			return nil, 500, " -STR006", false, err
		}
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR007", false, err
	}
	response := mapToAdminResponse(user)
	return response, 200, "", true, nil
}

func DeleteadminUser(client *ent.Client, id string, newUser ca_reg.DeleteAdminMasterStruc) (*ca_reg.AdminMasterResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer handleTransaction(tx, &err)
	// Query the user by EmployeeID
	user, err := fetchActiveAdminUser(ctx, tx, id)
	if err != nil {
		return nil, 422, "", false, err
	}

	stat := "deleted_" + time.Now().Format("20060102150405")
	// Update the user entity with the provided new user data
	user, err = user.Update().
		SetStatuss(stat).
		SetDeletedByDesignation(newUser.DeletedByDesignation).
		SetDeletedByEmpId(newUser.DeletedByEmpId).
		SetDeletedByUserName(newUser.DeletedByUserName).
		SetDeletedBy(newUser.UserName).
		SetDeletedById(newUser.DeletedById).
		SetDeletedDate(time.Now().UTC().Truncate(24 * time.Hour)).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR003", false, err
	}
	response := mapToAdminMasterResponse(user)
	return response, 200, "", true, nil
}
func QueryAdminUsersByEmpId(ctx context.Context, client *ent.Client, empid int64) ([]*ent.AdminMaster, int32, error) {
	//Can use GetX as well

	users, err := client.AdminMaster.Query().
		Where(adminmaster.EmployeeIdEQ(empid),
			adminmaster.StatussEQ("active"),
		).
		All(ctx)
	if err != nil {
		return nil, 400, fmt.Errorf("failed to fetch data for Admin master: %w", err)
	} else {
		if len(users) == 0 {
			return nil, 422, errors.New("no matching admin users found")
		}
	}
	return users, 200, nil
}

func QueryEmpUsersByEmpId(ctx context.Context, client *ent.Client, empid int64) ([]*ent.EmployeeMaster, int32, string, bool, error) {
	//Can use GetX as well

	user, err := client.EmployeeMaster.Query().
		Where(employeemaster.EmployeeIDEQ(empid),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatus(true)).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	if len(user) == 0 {
		// If no matching entities are found, return an error
		return nil, 422, " -STR002", false, errors.New("no data found in employee master")
	}

	return user, 200, "", true, nil
}
func QueryAdminUsersByusername(ctx context.Context, client *ent.Client, empid string) ([]*ent.AdminMaster, int32, error) {
	//Can use GetX as well

	user, err := client.AdminMaster.Query().
		Where(adminmaster.UserNameEQ(empid),
			adminmaster.StatussEQ("active"),
		).
		All(ctx)

	if err != nil {
		log.Println("error at gettting users with emp id: ", err)
		return nil, 400, fmt.Errorf("unable to fetch data try after some time : %v,", err)
	}
	if len(user) == 0 {
		// If no matching entities are found, return an error
		return nil, 422, errors.New("no matching users found")
	}

	return user, 200, nil
}
func QueryAdminUsersByfacility(ctx context.Context, client *ent.Client, empid string, role int32) ([]*ent.AdminMaster, int32, string, bool, error) {
	//Can use GetX as well

	adminuser, err := client.AdminMaster.Query().
		Where(
			adminmaster.FacilityIDEQ(empid),
			adminmaster.RoleUserCodeEQ(role),
			adminmaster.Or(
				adminmaster.StatussEQ("active"),
				adminmaster.StatussHasPrefix("deleted_"),
			),
		).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	/* if len(adminuser) == 0 {
		return nil, 422, " -STR002", false, errors.New("no matching role found")
	} */
	// nil result validation not made because this was handled by UI for exist and not exists.
	return adminuser, 200, "", true, nil
}

type StrucAdminRole struct {
	ID                          int    `json:"id"`
	EmployeeId                  int    `json:"employee_id"`
	EmployeeName                string `json:"employee_name"`
	Designation                 string `json:"designation"`
	RoleUserCode                int    `json:"role_user_code"`
	RoleUserDescription         string `json:"role_user_description"`
	Mobile                      string `json:"mobile"`
	EmailID                     string `json:"email_id"`
	UserName                    string `json:"user_name"`
	FacilityID                  string `json:"facility_id"`
	AuthorityFacilityName       string `json:"authority_facility_name"`
	FacilityType                string `json:"facility_type"`
	ReportingOfficeFacilityId   string `json:"reporting_office_facility_id"`
	ReportingOfficeFacilityName string `json:"reporting_office_facility_name"`
	CircleOfficeFacilityId      string `json:"circle_office_facility_id"`
	CircleOfficeName            string `json:"circle_office_name"`
}

func QueryAdminUsersByRole(ctx context.Context, client *ent.Client, role int32) ([]StrucAdminRole, int32, string, bool, error) {
	var adminResults []StrucAdminRole

	err := client.AdminMaster.Query().
		Where(
			adminmaster.RoleUserCodeEQ(role),
			adminmaster.StatussEQ("active"),
		).
		Select(
			adminmaster.FieldID,
			adminmaster.FieldEmployeeId,
			adminmaster.FieldEmployeeName,
			adminmaster.FieldDesignation,
			adminmaster.FieldRoleUserCode,
			adminmaster.FieldRoleUserDescription,
			adminmaster.FieldMobile,
			adminmaster.FieldEmailID,
			adminmaster.FieldUserName,
			adminmaster.FieldFacilityID,
			adminmaster.FieldAuthorityFacilityName,
			adminmaster.FieldFacilityType,
			adminmaster.FieldReportingOfficeFacilityId,
			adminmaster.FieldReportingOfficeFacilityName,
			adminmaster.FieldCircleOfficeFacilityId,
			adminmaster.FieldCircleOfficeName,
		).
		Scan(ctx, &adminResults)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}

	return adminResults, 200, "", true, nil
}

func ValidateAdminLoginn(client *ent.Client, newUser *ent.AdminMaster) (*ca_reg.AdminMasterResponse, int32, string, string, bool, error) {
	if err := validateAdminLoginInputs(newUser); err != nil {
		return nil, 422, "", " -STR001", false, err
	}
	newUser.UserName = strings.TrimSpace(newUser.UserName)
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, "", " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer handleTransaction(tx, &err)

	user, err := fetchAdminUser(tx, newUser.UserName, ctx)
	if err != nil {
		return nil, 422, "", " -STR003", false, err
	}

	dbOTP := user.OTP

	if newUser.OTP != dbOTP {
		return nil, 422, "", " -STR005", false, errors.New("incorrect OTP")
	} else if time.Now().After(user.OTPExpiryTime) {
		return nil, 422, "", " -STR006", false, errors.New(" otp expired, kindly regenerate the OTP")
	}

	if user.FacilityID == "" {
		return nil, 422, "", " -STR006", false, fmt.Errorf("facility ID for this Admin user was not mapped yet. Kindly contact your immediate supervising office to update details: %s", newUser.UserName)
	}

	// var officeName string
	token := a.CreateToken(newUser.UserName, client)

	if err := updateAdminToken(tx, newUser.UserName, token, ctx); err != nil {
		return nil, 500, "", " -STR007", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, "", " -STR008", false, err
	}
	response := mapToAdminResponse(user)
	return response, 200, token, "", true, nil
}

func ValidateAdminUserLoginn(client *ent.Client, newUser *ent.AdminMaster) (*ca_reg.AdminMasterResponse, int32, string, bool, int32, error) {
	// Check if the newUser and password are not nil

	if newUser == nil {
		return nil, 400, " -STR001", false, 0, errors.New("for login data fetched with blank")
	}

	// // Trim the username and password
	newUser.UserName = strings.TrimSpace(newUser.UserName)
	newUser.Password = strings.TrimSpace(newUser.Password)

	// Check if the username exists
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, 0, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer handleTransaction(tx, &err)

	// if adminmaster.StatussEQ("active") {
	user, err := fetchAdminUserr(tx, newUser.UserName)
	if err != nil {
		return nil, 422, " -STR002", false, 0, err
	}

	// Compare the password from the input with the user's password stored in the database
	if user.Password != newUser.Password {
		return nil, 422, " -STR003", false, 0, errors.New("incorrect Password")
	}

	otpGeneratedTime := time.Now()
	otpExpiryTime := otpGeneratedTime.Add(time.Minute * 2)
	otp := util.GenerateOTP()

	// Save the OTP in the UserMaster entity
	_, err = user.Update().
		SetOTP(otp).
		SetOTPExpiryTime(otpExpiryTime).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR004", false, 0, err
	}
	//SendSMSAndSaveOTPp
	/// Send SMS and save OTP

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR005", false, 0, err
	}
	response := mapToAdminResponse(user)
	return response, 200, "", true, otp, nil

}

func QueryAdminPasswordByEmpId(ctx context.Context, client *ent.Client, usern string) (string, error) {
	user, err := client.AdminMaster.Query().
		Where(adminmaster.UserNameEQ(usern), adminmaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user.Password, nil
}

func ChangeAdminUserPasswordByUsername(client *ent.Client, username string, currentPassword, newPassword string) (*ent.AdminMaster, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, "  -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer cancel()

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	adminMaster, err := tx.AdminMaster.Query().
		Where(adminmaster.UserNameEQ(username), adminmaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf("admin user %s not exists ", username)
		}
		return nil, 500, " -STR002", false, err
	}

	if currentPassword != adminMaster.Password {
		return nil, 422, " -STR003", false, errors.New("current password is incorrect ")
	}

	adminMaster, err = adminMaster.Update().
		SetPassword(string(newPassword)).
		SetEventTime(time.Now().Truncate(time.Second)).
		Save(context.Background())
	if err != nil {
		return nil, 500, " -STR004", false, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR027", false, err
	}

	//return adminMaster, 200, nil
	return adminMaster, 200, "", false, nil

}

func AdminDataNew(client *ent.Client, userid string) (*ent.AdminMaster, int32, error) {
	admin, err := client.AdminMaster.
		Query().
		Where(adminmaster.UserNameEQ(userid), adminmaster.StatussEQ("active")).
		Only(context.Background())
	if err != nil {
		if admin == nil {
			return nil, 422, fmt.Errorf("no data exists for this Admin user %s", userid)
		}
		return nil, 400, fmt.Errorf("unable to fetch data from Admin Master for this user %s , %v", userid, err)
	}
	return admin, 200, nil
}
