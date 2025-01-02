package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"recruit/authentication"
	"recruit/ent"
	"recruit/ent/adminmaster"
	"recruit/util"

	"recruit/ent/employeemaster"
	"recruit/ent/usermaster"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"strconv"
	"strings"
	"time"
)

type UserExistence struct {
	NewUser bool `json:"newuser"`
}

// "strings"
func QueryUserMasterByUserName(ctx context.Context, client *ent.Client, username string) (*ent.UserMaster, int32, string, bool, error) {
	//Can use GetX as well
	if username == "" {
		return nil, 422, " -STR001", false, errors.New(" the username is empty")
	}
	user, err := client.UserMaster.Query().
		Where(usermaster.UserNameEQ(username)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR002", false, fmt.Errorf(" no such username: %s", username)
		}
		return nil, 500, " -STR003", false, err
	}
	return user, 200, "", true, nil
}

func QueryUserMasterByEmpId(ctx context.Context, client *ent.Client, empid int64) (bool, error) {
	if empid == 0 {
		return false, errors.New(" please pass Emp ID as a non-zero parameter")
	}
	exists, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empid)).
		Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to check empid existence: %v", err)
	}
	return !exists, nil
}

func CreateUserByEmpId(ctx context.Context, client *ent.Client, empid int64) (*ent.UserMaster, error) {
	if empid <= 0 {
		return nil, errors.New(" please input a nonzero Emp ID as an input parameter")
	}

	exists, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empid)).Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check empid existence: %v", err)
	}
	if exists {
		user, err := client.UserMaster.Query().
			Where(usermaster.EmployeeIDEQ(empid)).
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query existing user: %v", err)
		}

		if user.Password != "" && user.Status {
			user.NewPasswordRequest = false
			user, err = user.Update().SetNewPasswordRequest(false).Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to update existing user: %v", err)
			}
			return user, nil
		}

		return user, nil
	}
	employee, err := client.EmployeeMaster.Query().
		Where(employeemaster.EmployeeIDEQ(empid), employeemaster.StatussEQ("active")).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf(" the Employee Id is not available in Employee Master: %v", err)
	}

	if employee.FacilityID == "" {
		return nil, errors.New(" please update your Office ID with your parent office")
	}

	if employee.MobileNumber == "" {
		return nil, errors.New(" please get your mobile number updated with your parent Office")
	}
	if employee.EmailID == "" {
		return nil, errors.New(" please get your Email ID updated with your parent Office")
	}

	// Trigger the SMS OTP.
	if employee.FacilityID != "" && employee.MobileNumber != "" && employee.EmailID != "" && employee.EmployeeName != "" {
		user, err := client.UserMaster.
			Create().
			SetEmployeeID(empid).
			SetRoleUserCode(1).
			SetUserName(strconv.Itoa(int(empid))).
			SetStatus(false).
			SetMobile(employee.MobileNumber).
			SetEmailID(employee.EmailID).
			SetEmployeeName(employee.EmployeeName).
			SetFacilityID(employee.FacilityID).
			SetNewPasswordRequest(true).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %v", err)
		}

		return user, nil
	}

	return nil, nil
}

func QueryUsersByEmpId(ctx context.Context, client *ent.Client, empid int64) (*ent.UserMaster, error) {
	//Can use GetX as well

	user, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empid)).
		Only(ctx)

	if err != nil {
		log.Println("error at gettting users with emp id: ", err)
		return nil, fmt.Errorf("failed at users in user master: %w", err)
	}
	log.Println("user returned by empid : ", user)
	return user, nil
}

// Generate OTP
func GenerateOTP() int32 {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	otp := rand.Intn(max-min+1) + min
	return int32(otp)
}

// Send SMS

// Update Password
/* func UpdateUserByEmpID(client *ent.Client, empID int64) (*ent.UserMaster, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if empID <= 0 {
		return nil, errors.New(" please input a nonzero Emp ID as an input parameter")
	}

	exists, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empID)).
		//Where(usermaster.PasswordIsNil()).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check empID existence: %v", err)
	}
	if exists {
		userMaster, err := client.UserMaster.Query().
			Where(usermaster.EmployeeIDEQ(empID)).
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve userMaster: %v", err)
		}

		otp := GenerateOTP() // Assuming you have a function to generate the OTP

		updatedUser, err := userMaster.Update().
			SetOTP(otp).
			SetPassword(userMaster.Password).
			SetStatus(true).
			SetNewPasswordRequest(false).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update user: %v", err)
		}

		return updatedUser, nil
	}

	return nil, errors.New(" the Employee ID is unavailable ")
} */

// Validate Login

func ValidateLoginUser(client *ent.Client, newuser *ent.UserMaster) (*ca_reg.ResponseVerifyCandidateUserLogin, int32, string, bool, error) {
	var userDetails ca_reg.ResponseVerifyCandidateUserLogin
	// Check if the username exists

	if newuser == nil {
		return nil, 422, " -STR001", false, errors.New(" userDetails cannot be blank")
	}
	newuser.UserName = strings.TrimSpace(newuser.UserName)
	newuser.Password = strings.TrimSpace(newuser.Password)

	stg, _ := strconv.Atoi(newuser.UserName)
	stgUserName := int64(stg)
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	username := newuser.UserName

	if newuser.UserName == "" {
		return nil, 422, " -STR002", false, errors.New(" username cannot be empty")
	}
	if newuser.Password == "" {
		return nil, 422, " -STR003", false, fmt.Errorf(" for %s Password is empty", username)
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

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

	employee, err := tx.EmployeeMaster.Query().
		Where(employeemaster.EmployeeIDEQ(stgUserName), employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(true)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR004", false, errors.New("for you no active employee details found")
		} else {
			return nil, 500, " -STR005", false, err
		}
	}

	// Retrieve the user record with the provided username and password
	user, err := tx.UserMaster.
		Query().
		Where(
			usermaster.UserNameEQ(username),
			usermaster.StatussEQ("active"),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR006", false, errors.New("no such username exists ")
		} else {
			return nil, 500, " -STR007", false, err
		}
	}
	fmt.Println("before OTP", user.OTP)
	if !user.Status {
		return nil, 422, " -STR008", false, errors.New("registration not done for this user")
	}

	if user.Password != newuser.Password {
		return nil, 422, " -STR009", false, fmt.Errorf("for %s incorrect password entered", username)
	}

	token := authentication.CreateToken(username, client)
	_, err = tx.UserMaster.
		Update().
		SetUidToken(token).
		Where(usermaster.UserNameEQ(username), usermaster.StatussEQ("active"),
			usermaster.StatusEQ(true)).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR010", false, err
	}
	fmt.Println("after OTP", user.OTP)
	userDetails = ca_reg.ResponseVerifyCandidateUserLogin{
		UserID:           user.ID,
		EmployeeID:       user.EmployeeID,
		EmployeeName:     user.EmployeeName,
		RoleUserCode:     user.RoleUserCode,
		UserName:         user.UserName,
		Mobile:           user.Mobile,
		Email:            user.EmailID,
		DOB:              employee.DOB,
		Gender:           employee.Gender,
		EmployeeCategory: employee.EmployeeCategory,
		EmployeePost:     employee.EmployeePost,
		FacilityId:       employee.FacilityID,
		CircleFacilityId: employee.CircleFacilityID,
		Token:            token,
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR011", false, err
	}
	return &userDetails, 200, "", true, nil
}

//sennd sms

//Update Password for an user with EmpID. in main

/* func UpdateUserPasswordByEmpID(client *ent.Client, empID int64, newUser *ent.UserMaster) (*ent.UserMaster, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if empID <= 0 {
		return nil, 422, " -STR001", false, errors.New(" enter valid employee id")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, err
	}

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

	userMaster, err := tx.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR002", false, errors.New(" the Employee ID does not exist")
		}
		return nil, 500, " -STR003", false, err
	}

	// Check if the newuser.OTP matches the OTP in the database
	if newUser.OTP != userMaster.OTP {
		//log.Printf("Mismatched OTP for Employee ID %d. Expected OTP: %s, Provided OTP: %s", empID, userMaster.OTP, newUser.OTP)
		return nil, 422, " -STR004", false, errors.New(" invalid OTP")
	}

	// Check if the newuser.OTP = payload.OTP, OTP expires+6mins <=currentime and time. (OTP and OTP expires time i)
	if userMaster.Password == "" {
		userMaster, err = userMaster.Update().
			SetPassword(newUser.Password).
			SetOTP(newUser.OTP).
			SetNewPasswordRequest(false).
			SetStatus(true).
			Save(ctx)
		if err != nil {
			return nil, 500, " -STR005", false, err
		}
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR006", false, err
	}
	return userMaster, 200, "", true, nil
} */

// Update First Time username password with OTP expiry validation
func NewValidateLoginUser(client *ent.Client, newuser *ent.UserMaster) (*ca_reg.UserMasterResponse, int32, string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Check if the username exists
	username := newuser.UserName
	if username == "" {
		return nil, 422, " -STR001", false, errors.New(" username cannot be empty")
	}
	if newuser.Password == "" {
		return nil, 422, " -STR002", false, errors.New(" password cannot be empty")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer handleTransaction(tx, &err)

	// Retrieve the user record with the provided username
	user, err := fetchUser(tx, ctx, newuser.UserName)
	if err != nil {
		return nil, 422, " -STR003", false, err
	}

	if err := validateUserStatus(user); err != nil {
		return nil, 422, err.Error(), false, err
	}

	// Validate the input OTP with the stored OTP
	if user.OTP != newuser.OTP {
		return nil, 422, " -STR006", false, errors.New(" invalid Mobile OTP")
	} else if user.EmailOTP != newuser.EmailOTP {
		return nil, 422, " -STR007", false, errors.New(" invalid Email OTP")
	} else if time.Now().After(user.OTPExpiryTime) {
		response := createUserResponse(user)
		return response, 422, " -STR008", false, errors.New(" otp Expired, Kinldy regenerate the OTP")

	}

	if err := updateUser(tx, ctx, user, newuser.Password); err != nil {
		return nil, 500, " -STR009", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR010", false, err
	}
	response := createUserResponse(user)
	return response, 200, "", true, nil
}
func ModifyEmployeeMasterUserMaster(client *ent.Client, empMasterRequest ca_reg.StrucProfileEmployeeMaster) (string, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	id := empMasterRequest.EmployeeID

	tx, err := client.Tx(ctx)
	if err != nil {
		return "", 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

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

	// Query the user by EmployeeID
	user, err := tx.EmployeeMaster.Query().
		Where(
			employeemaster.EmployeeIDEQ(id),
			employeemaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return "", 500, " -STR003", false, err
		}
	}

	// Format the current time to "yyyymmddhhmmss"
	stat := "inactive_" + time.Now().Format("20060102150405")

	// Update the user entity with the provided new user data
	_, err = user.Update().
		SetStatuss(stat).
		Save(ctx)
	if err != nil {
		return "", 500, " -STR005", false, err
	}

	_, err = tx.EmployeeMaster.Create().
		SetCreatedById(user.CreatedById).
		SetCreatedByUserName(user.CreatedByUserName).
		SetCreatedByEmpId(user.CreatedByEmpId).
		SetCreatedByDesignation(user.CreatedByDesignation).
		SetCreatedDate(user.CreatedDate).
		SetModifiedDate(time.Now()).
		SetVerifyStatus(user.VerifyStatus).
		SetEmployeeID(empMasterRequest.EmployeeID).
		SetEmployeeName(empMasterRequest.EmployeeName).
		SetDOB(empMasterRequest.DOB).
		SetGender(empMasterRequest.Gender).
		SetMobileNumber(empMasterRequest.MobileNumber).
		SetEmailID(empMasterRequest.EmailID).
		SetEmployeeCategory(empMasterRequest.EmployeeCategory).
		SetEmployeePost(empMasterRequest.EmployeePost).
		SetFacilityID(empMasterRequest.FacilityID).
		SetPincode(empMasterRequest.Pincode).
		SetOfficeName(empMasterRequest.OfficeName).
		SetControllingAuthorityFacilityId(empMasterRequest.ControllingAuthorityFacilityID).
		SetControllingAuthorityName(empMasterRequest.ControllingAuthorityName).
		SetNodalAuthorityFaciliyId(empMasterRequest.NodalAuthorityFacilityID).
		SetNodalAuthorityName(empMasterRequest.NodalAuthorityName).
		SetCircleFacilityID(empMasterRequest.CircleFacilityID).
		SetModifiedById(empMasterRequest.ModifiedByID).
		SetModifiedByUserName(empMasterRequest.ModifiedByUserName).
		SetModifiedByEmpId(empMasterRequest.ModifiedByEmpID).
		SetUpdatedBy("API").
		Save(ctx)
	if err != nil {
		return "", 500, " -STR006", false, err
	}

	existingUser, err := tx.UserMaster.Query().
		Where(usermaster.UserNameEQ(empMasterRequest.UserName),
			usermaster.StatussEQ("active")).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return "", 500, " -STR010", false, err
	}

	_, err = existingUser.Update().
		SetStatuss(stat).
		Save(ctx)
	if err != nil {
		return "", 500, " -STR005", false, err
	}

	// Handle nil case for existingUser
	var password string
	var createdAt time.Time
	if existingUser != nil {
		password = existingUser.Password
		createdAt = existingUser.CreatedAt
	}
	modifiedByEmpIDStr := strconv.FormatInt(empMasterRequest.ModifiedByEmpID, 10)
	_, err = tx.UserMaster.Create().
		SetEmployeeID(empMasterRequest.EmployeeID).
		SetEmployeeName(empMasterRequest.EmployeeName).
		SetMobile(empMasterRequest.MobileNumber).
		SetEmailID(empMasterRequest.EmailID).
		SetUserName(empMasterRequest.UserName).
		SetPassword(password).
		SetCreatedAt(createdAt).
		SetNewPasswordRequest(true).
		SetCreatedById(empMasterRequest.ModifiedByID).
		SetCreatedByUserName(empMasterRequest.ModifiedByUserName).
		SetFacilityID(empMasterRequest.FacilityID).
		SetCircleFacilityId(empMasterRequest.NodalAuthorityFacilityID).
		SetCircleFacilityName(empMasterRequest.NodalAuthorityName).
		SetUpdatedDate(time.Now()).
		SetModifiedDate(time.Now()).
		SetUidToken(existingUser.UidToken).
		SetCreatedByEmployeeId(modifiedByEmpIDStr).
		SetCircleFacilityId(empMasterRequest.CircleFacilityID).
		SetStatus(true).
		SetCreatedBy("API").
		SetRoleUserCode(1).
		Save(ctx)
	if err != nil {
		return "", 500, " -STR012", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return "", 500, " -STR013", false, err
	}
	return "UserDetails Updated Successfully", 200, "", true, nil
}

func ValidateOTPUserEmailorMobile(client *ent.Client, newuser ca_reg.CandidateEditProfileOTP) (string, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	username := newuser.UserName
	if username == "" {
		return "", 422, " -STR001", false, errors.New("username cannot be empty")
	}

	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserName(username), usermaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", 422, " -STR003", false, errors.New("no such username exists")
		}
		return "", 500, " -STR004", false, err
	}

	if time.Now().After(user.OTPExpiryTime) {
		return "OTP expired, kindly regenerate the OTP", 422, " -STR008", false, errors.New("OTP expired, kindly regenerate the OTP")
	}

	if newuser.MobileEmail == "Mobile" {
		if user.OTP != newuser.OldOTP {
			return "", 422, " -STR006", false, errors.New("invalid mobile OTP")
		}
		if user.OTPNew != newuser.NewOTP {
			return "", 422, " -STR007", false, errors.New("invalid new mobile OTP")
		}
	} else if newuser.MobileEmail == "EmailID" {
		if user.EmailOTP != newuser.OldOTP {
			return "", 422, " -STR006", false, errors.New("invalid email OTP")
		}
		if user.EmailOTPNew != newuser.EmailOTPNew {
			return "", 422, " -STR007", false, errors.New("invalid new email OTP")
		}
	} else {
		return "", 422, " -STR009", false, errors.New("invalid MobileEmail value")
	}

	return "OTP verified successfully", 200, "", true, nil
}

func ChangeUserPasswordByUsername(client *ent.Client, username string, currentPassword, newPassword string) (*ent.UserMaster, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

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

	userMaster, err := tx.UserMaster.Query().
		Where(usermaster.UserNameEQ(username), usermaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, fmt.Errorf(" user %s not exists ", username)
		}
		return nil, 500, " -STR002", false, err
	}

	if currentPassword != userMaster.Password {
		return nil, 422, " -STR003", false, errors.New("current password is incorrect ")
	}

	userMaster, err = userMaster.Update().
		SetPassword(string(newPassword)).
		SetUpdatedDate(time.Now().Truncate(time.Second)).
		Save(context.Background())
	if err != nil {
		return nil, 500, " -STR004", false, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR005", false, err
	}
	return userMaster, 200, "", false, nil

}

// validate login  status after pswd matcches , with otpexpiry validation
func ValidateAdminLogin(client *ent.Client, newUser *ent.UserMaster) (*ent.UserMaster, string, error) {
	// Check if the newUser and password are not nil
	if newUser == nil {
		return nil, "", errors.New(" userMaster cannot be nil")
	}

	if newUser.OTP <= 0 {
		return nil, "", errors.New(" otp cannot be nil")
	}

	// Trim the username and password
	newUser.UserName = strings.TrimSpace(newUser.UserName)

	// Check if the username exists
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	exists, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(newUser.UserName)).
		Exist(ctx)

	if err != nil {
		return nil, "", fmt.Errorf("failed to check username existence: %v", err)
	}

	if !exists {
		return nil, "", fmt.Errorf(" invalid Username or username not found: %s", newUser.UserName)
	}

	// Retrieve the user record with the provided username
	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(newUser.UserName)).
		// WithCircleUsersRef().
		Only(ctx)

	if err != nil {
		return nil, "", fmt.Errorf("failed to retrieve user: %v", err)
	}
	dbOTP := user.OTP

	// Compare the input OTP with the OTP from the database
	if newUser.OTP != dbOTP {
		// Log the mismatched OTPs
		log.Printf("Input OTP: %d, Database OTP: %d, Mobile Number: %s, Username: %s", newUser.OTP, dbOTP, user.Mobile, user.UserName)
		return nil, "", errors.New(" incorrect OTP")
	} else if time.Now().After(user.OTPExpiryTime) {
		return user, "", errors.New(" otp Expired, Kinldy regenerate the OTP")
	}

	// Check if FacilityID is null
	if user.FacilityID == "" {
		return nil, "", errors.New(" no valid associated OfficeID for the AdminUser: " + newUser.UserName)
	}
	var officeName string

	return user, officeName, nil
}

func SubUserResetValidateUserName(ctx context.Context, client *ent.Client, newUsers ca_reg.StrucUserResetValidateUserName) (*ca_reg.UserMasterResponse, int32, string, bool, error) {
	if len(newUsers.UserName) != 8 {
		return nil, 422, " -STR001", false, errors.New("enter User Name with eight digit number")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer handleTransaction(tx, &err)
	user, err := fetchActiveUser(tx, ctx, newUsers.UserName)
	if err != nil {
		return nil, 422, " -STR002", false, err
	}

	if !user.Status {
		return nil, 422, " -STR003", false, fmt.Errorf("user : %s was not verified", newUsers.UserName)
	}
	otpGeneratedTime := time.Now()
	otpExpiryTime := otpGeneratedTime.Add(time.Minute * 2)
	otp := util.GenerateOTP()

	// Save the OTP in the UserMaster entity
	usernew, err := updateUserOTP(user, ctx, otp, otpExpiryTime)
	if err != nil {
		return nil, 500, " -STR004", false, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR005", false, err
	}
	ResponseData := &ca_reg.UserMasterResponse{
		ID:           usernew.ID,
		EmployeeID:   usernew.EmployeeID,
		UserName:     usernew.UserName,
		EmailID:      usernew.EmailID,
		Mobile:       usernew.Mobile,
		RoleUserCode: usernew.RoleUserCode,
	}
	// Trigger the SMS OTP.
	return ResponseData, 200, "", true, nil
}

func SubUserResetValidateOTP(ctx context.Context, client *ent.Client, newUsers ca_reg.StrucUserResetValidateOTP) (*ca_reg.UserMasterResponse, int32, string, bool, error) {
	if len(newUsers.UserName) != 8 {
		return nil, 422, " -STR001", false, errors.New("enter User Name with eight digit number")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer handleTransaction(tx, &err)

	users, err := client.UserMaster.Query().
		Where(usermaster.UserNameEQ(newUsers.UserName),
			usermaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR002", false, errors.New("no such user exists in active state ")
		} else {
			return nil, 500, " -STR003", false, err
		}
	}
	if !users.Status {
		return nil, 422, " -STR004", false, fmt.Errorf("registration not done for this user : %s ", newUsers.UserName)
	}
	if users.OTP != newUsers.OTP {
		return nil, 422, " -STR005", false, errors.New("invalid OTP")
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR006", false, err
	}
	response := &ca_reg.UserMasterResponse{
		ID:                 users.ID,
		UserName:           users.UserName,
		Status:             users.Status,
		Mobile:             users.Mobile,
		DOB:                users.DOB,
		Statuss:            users.Statuss,
		EmailID:            users.EmailID,
		EmployeeID:         users.EmployeeID,
		EmployeeName:       users.EmployeeName,
		FacilityID:         users.FacilityID,
		CircleFacilityId:   users.CircleFacilityId,
		CircleFacilityName: users.CircleFacilityName,
		RoleUserCode:       users.RoleUserCode,
	}
	// Trigger the SMS OTP.
	return response, 200, "", true, nil
}

func SubUserResetSaveNewPassword(ctx context.Context, client *ent.Client, newUsers ca_reg.StrucUserResetSaveNewPassword) (*ca_reg.UserMasterResponse, int32, string, bool, error) {
	if len(newUsers.UserName) != 8 {
		return nil, 422, " -STR001", false, errors.New("enter User Name with eight digit number")
	}
	if newUsers.OTP == 0 {
		return nil, 422, " -STR002", false, errors.New("otp not received with six digit number")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer handleTransaction(tx, &err)

	getUser, err := fetchActiveUser(tx, ctx, newUsers.UserName)
	if err != nil {
		return nil, 422, " -STR003", false, err
	}

	if !getUser.Status {
		return nil, 422, " -STR005", false, fmt.Errorf("registration not done for this user : %s ", newUsers.UserName)
	}

	if getUser.OTP != newUsers.OTP {
		return nil, 422, " -STR006", false, errors.New("invalid OTP")
	}

	if err := updateUserPassword(tx, ctx, newUsers); err != nil {
		return nil, 500, " -STR007", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR008", false, err
	}
	return nil, 200, "", true, nil
}

func SubAdminResetValidateUserName(ctx context.Context, client *ent.Client, newUsers ca_reg.StrucAdminResetValidateUserName) (*ent.AdminMaster, int32, string, bool, int32, error) {
	if len(newUsers.UserName) == 0 {
		return nil, 422, " -STR001", false, 0, errors.New("admin user name is empty")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, 0, fmt.Errorf("failed to start transaction: %w", err)
	}

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

	user, err := tx.AdminMaster.Query().
		Where(adminmaster.UserNameEQ(newUsers.UserName),
			adminmaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR002", false, 0, errors.New("no such admin user exists")
		} else {
			return nil, 500, " -STR003", false, 0, err
		}
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
		return nil, 500, " -STR005", false, 0, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR027", false, 0, err
	}
	return user, 200, "", true, otp, nil

}

func SubAdminResetValidateOTP(ctx context.Context, client *ent.Client, newUsers ca_reg.StrucAdminResetValidateOTP) (*ent.AdminMaster, int32, string, bool, error) {
	if len(newUsers.UserName) == 0 {
		return nil, 422, " -STR001", false, errors.New("admin user name should be empty")
	}
	if newUsers.OTP == 0 {
		return nil, 422, " -STR002", false, errors.New("otp received with zero value")
	}
	user, err := client.AdminMaster.Query().
		Where(adminmaster.UserNameEQ(newUsers.UserName),
			adminmaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR003", false, fmt.Errorf("admin user : %s not exists", newUsers.UserName)
		} else {
			return nil, 500, " -STR004", false, err
		}
	}

	if user.OTP != newUsers.OTP {
		return nil, 422, " -STR005", false, errors.New("invalid OTP")
	}

	return user, 200, "", true, nil
}

func SubAdminResetSaveNewPassword(ctx context.Context, client *ent.Client, newUsers ca_reg.StrucAdminResetSaveNewPassword) (*ca_reg.AdminMasterResponse, int32, string, bool, error) {
	if len(newUsers.UserName) == 0 {
		return nil, 422, " -STR001", false, errors.New("enter valid admin user id")
	}
	if newUsers.OTP == 0 {
		return nil, 422, " -STR002", false, errors.New("otp received with nil value")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer handleTransaction(tx, &err)

	getUser, err := fetchActiveAdminUser(ctx, tx, newUsers.UserName)
	if err != nil {
		return nil, 422, "", false, err
	}

	if getUser.OTP != newUsers.OTP {
		return nil, 422, " -STR005", false, errors.New("invalid OTP")
	}

	if err := updateAdminPassword(ctx, tx, newUsers.UserName, newUsers.NewPassword); err != nil {
		return nil, 500, "", false, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR007", false, err
	}
	return nil, 200, "", true, nil
}

func NewCreateUserByEmpId(ctx context.Context, client *ent.Client, newUsers ca_reg.CandidateRegistrationGenerateOTP) (*ca_reg.UserMasterResponse, int32, string, bool, error) {
	if len(newUsers.UserName) == 0 {
		return nil, 422, " -STR001", false, errors.New("enter User Name with eight digit number")
	}
	EmployeeIDNum, _ := strconv.ParseInt(newUsers.UserName, 10, 32)
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer handleTransaction(tx, &err)

	employee, err := tx.EmployeeMaster.Query().
		Where(employeemaster.EmployeeIDEQ(EmployeeIDNum),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(true)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR002", false, errors.New("contact your Controlling authority to update you employee data")
		} else {
			return nil, 500, " -STR003", false, err
		}
	}

	if employee.FacilityID == "" {
		return nil, 422, " -STR004", false, errors.New("office Name is missing. Contact your Controlling authority to update missing data")
	}

	if employee.MobileNumber == "" && len(employee.MobileNumber) != 10 {
		return nil, 422, " -STR005", false, errors.New("mobile number is missing. Contact your Controlling authority to update missing data")
	}
	if employee.EmailID == "" {
		return nil, 422, " -STR006", false, errors.New("email ID is missing. Contact your Controlling authority to update missing data")
	}
	//email string, Remarks string, logdata ca_reg.LogData, gctx *gin.Context, client *ent.Client
	var valEamil string = ValidateEmail(employee.EmailID)
	if valEamil != "" {
		return nil, 422, " -STR007", false, errors.New(valEamil)
	}
	if employee.CircleFacilityID == "" {
		return nil, 422, " -STR007", false, errors.New(" nodal office details missing. Contact your Controlling authority to update missing data")
	}

	user, err := tx.UserMaster.Query().
		Where(usermaster.UserNameEQ(newUsers.UserName),
			usermaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, 500, " -STR010", false, err
		}
	} else {
		if !user.Status {
			_, err = tx.UserMaster.Delete().Where(usermaster.UserNameEQ(newUsers.UserName)).Exec(ctx)
			if err != nil {
				return nil, 500, " -STR009", false, err
			}
		}
		if user.Status {
			return nil, 422, " -STR011", false, errors.New("already user name exists in active status")
		}
	}

	smsotp := util.GenerateOTP()

	otpGeneratedTime := time.Now()
	otpExpiryTime := otpGeneratedTime.Add(time.Minute * 5)
	otp := util.GenerateEmailOTP()
	users, err := tx.UserMaster.
		Create().
		SetEmployeeID(EmployeeIDNum).
		SetRoleUserCode(1).
		SetUserName(newUsers.UserName).
		SetStatus(false).
		SetMobile(employee.MobileNumber).
		SetEmailID(employee.EmailID).
		SetEmployeeName(employee.EmployeeName).
		SetFacilityID(employee.FacilityID).
		SetCircleFacilityId(employee.NodalAuthorityFaciliyId).
		SetCircleFacilityName(employee.NodalAuthorityName).
		SetNewPasswordRequest(true).
		SetPassword(newUsers.Password).
		SetOTP(smsotp).
		SetEmailOTP(otp).
		SetOTPSavedTime(otpGeneratedTime).
		SetOTPExpiryTime(otpExpiryTime).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR012", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR013", false, err
	}
	response := &ca_reg.UserMasterResponse{
		ID:                 users.ID,
		UserName:           users.UserName,
		Status:             users.Status,
		Mobile:             users.Mobile,
		DOB:                users.DOB,
		Statuss:            users.Statuss,
		EmailID:            users.EmailID,
		EmployeeID:         users.EmployeeID,
		EmployeeName:       users.EmployeeName,
		FacilityID:         users.FacilityID,
		CircleFacilityId:   users.CircleFacilityId,
		CircleFacilityName: users.CircleFacilityName,
		RoleUserCode:       users.RoleUserCode,
	}
	return response, 200, "", true, nil

}

func MobileEmailchangeUserByUsername(ctx context.Context, client *ent.Client, newUser ca_reg.CandidateChangeEmailMobileGenerateOTP) (*ent.UserMaster, int32, string, bool, error) {

	EmployeeIDNum, _ := strconv.ParseInt(newUser.UserName, 10, 32)
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

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

	employee, err := tx.EmployeeMaster.Query().
		Where(employeemaster.EmployeeIDEQ(EmployeeIDNum),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(true)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR002", false, errors.New("contact your Controlling authority to update your employee data")
		}
		return nil, 500, " -STR003", false, err
	}

	if employee.FacilityID == "" {
		return nil, 422, " -STR004", false, errors.New("office Name is missing. Contact your Controlling authority to update missing data")
	}

	if employee.MobileNumber == "" || len(employee.MobileNumber) != 10 {
		return nil, 422, " -STR005", false, errors.New("mobile number is missing or invalid. Contact your Controlling authority to update missing data")
	}
	if employee.EmailID == "" {
		return nil, 422, " -STR006", false, errors.New("email ID is missing. Contact your Controlling authority to update missing data")
	}

	var valEmail string = ValidateEmail(employee.EmailID)
	if valEmail != "" {
		return nil, 422, " -STR007", false, errors.New(valEmail)
	}
	if employee.CircleFacilityID == "" {
		return nil, 422, " -STR008", false, errors.New("nodal office details missing. Contact your Controlling authority to update missing data")
	}

	user, err := tx.UserMaster.Query().
		Where(usermaster.UserNameEQ(newUser.UserName),
			usermaster.StatussEQ("active")).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, 500, " -STR011", false, err
	}
	fmt.Println("username is ", newUser.UserName)

	if user == nil {
		return nil, 422, " -STR012", false, errors.New("user name does not exist")
	}
	smsotp := util.GenerateOTP()
	otpGeneratedTime := time.Now()
	otpExpiryTime := otpGeneratedTime.Add(time.Minute * 5)

	emailOTP := util.GenerateEmailOTP()

	var newEmailOTP, newMobileOTP int32
	if newUser.MobileEmail == "EmailID" {
		newEmailOTP = util.GenerateEmailOTP()
		for newEmailOTP == emailOTP {
			newEmailOTP = util.GenerateEmailNewOTP() // Ensure newMobileOTP is different from smsotp
		}
		fmt.Println("oldEmailOTP", emailOTP)
		fmt.Println("newEmailOTP", newEmailOTP)
	} else if newUser.MobileEmail == "Mobile" {
		newMobileOTP = util.GenerateOTP()
		for newMobileOTP == smsotp {
			newMobileOTP = util.GenerateNewOTP() // Ensure newMobileOTP is different from smsotp
		}

	}
	user, err = user.Update().
		SetOTP(smsotp).
		SetEmailOTPNew(newEmailOTP).
		SetOTPNew(newMobileOTP).
		SetEmailOTP(emailOTP).
		SetOTPSavedTime(otpGeneratedTime).
		SetOTPExpiryTime(otpExpiryTime).
		Save(ctx)
	if err != nil {
		return nil, 500, " -STR013", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR014", false, err
	}
	return user, 200, "", true, nil
}

// Delete a user with user name
func DeleteUserByUserName(client *ent.Client, username string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Check if the user with the given username exists
	existingUser, err := client.UserMaster.Query().Where(usermaster.UserNameEQ(username)).First(ctx)
	if err != nil {
		return "", err
	}

	// Delete the user if it exists
	err = client.UserMaster.DeleteOne(existingUser).Exec(ctx)
	if err != nil {
		return "", err
	}

	return username + " is deleted successfully.", nil
}

// sending email

func QueryPasswordByEmpId(ctx context.Context, client *ent.Client, empid int64) (string, error) {
	user, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empid)).
		Only(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user.Password, nil
}

func ChangeUserPasswordByEmpID(client *ent.Client, empID int64, currentPassword, newPassword string) (*ent.UserMaster, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	if empID <= 0 {
		return nil, 400, " -STR001", false, errors.New(" please input a nonzero Emp ID as an input parameter")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			if err := tx.Commit(); err != nil {
				tx.Rollback()
			}
		}
	}()
	userMaster, err := tx.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR002", false, errors.New(" the Employee ID does not exist")
		}
		return nil, 500, " -STR003", false, err
	}

	if currentPassword != userMaster.Password {
		return nil, 422, " -STR004", false, fmt.Errorf(" wrong Password")
	}

	userMaster, err = userMaster.Update().
		SetPassword(string(newPassword)).
		SetModifiedDate(time.Now().Truncate(time.Second)).
		Save(context.Background())
	if err != nil {
		return nil, 500, " -STR005", false, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR006", false, err
	}

	return userMaster, 200, "", true, nil
}
