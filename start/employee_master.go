package start

import (
	"context"
	"errors"
	"fmt"

	//"net/http"
	"os"
	"recruit/ent"
	"recruit/ent/employeemaster"
	"recruit/ent/usermaster"
	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"
	"strconv"
	"time"
	//"github.com/gin-gonic/gin"
	//	"strings"
)

func SubVerifyTriggerSMSOTP(ctx context.Context, client *ent.Client, empMasterRequest ca_reg.VerifyTriggerCandidateSMSOTP) (string, int32, string, bool, error) {
	if len(empMasterRequest.EmployeeId) == 0 {
		return "", 422, " -STR001", false, errors.New("enter User Name with eight digit number")
	}
	if empMasterRequest.EmployeeId == "" && len(empMasterRequest.EmployeeId) != 10 {
		return "", 422, " -STR002", false, errors.New("mobile number is missing. ")
	}
	EmployeeIDNum, _ := strconv.ParseInt(empMasterRequest.EmployeeId, 10, 32)

	tx, err := client.Tx(ctx)
	if err != nil {
		return "", 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}
	//var dataExists   bool = false
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
			employeemaster.StatussEQ("active")).
		Exist(ctx)
	if err != nil {
		return "", 500, " -STR003", false, err
	}
	if !employee {
		return "", 422, " -STR004", false, errors.New("no details exists for this employee . ")
	} else {
		employee, err := tx.EmployeeMaster.Query().
			Where(employeemaster.EmployeeIDEQ(EmployeeIDNum),
				employeemaster.StatussEQ("active")).
			Only(ctx)
		if err != nil {
			return "", 500, " -STR004", false, err
		}
		if employee.VerifyStatus {
			return "", 422, " -STR005", false, errors.New("already request approved by your Controlling Authority. you can do First time Registration")
			//add one more validation to check in Usermaster
		} else if employee.FinalSubmitStatus {
			return "", 422, " -STR008", false, errors.New("contact your Controlling authority to approve your request")
		}
		if employee.SmsOtp == empMasterRequest.SmsOTP {
			_, err = tx.EmployeeMaster.
				Update().
				Where(employeemaster.EmployeeID(EmployeeIDNum),
					employeemaster.StatussEQ("active")).
				SetSmsVerifyStatus(true).
				SetCreatedDate(time.Now()).
				Save(ctx)

			if err != nil {
				return "", 500, " -STR010", false, err
			}

			if err = tx.Commit(); err != nil {
				tx.Rollback()
				return "", 500, " -STR011", false, err
			}
			return "SMS OTP verified", 200, "", true, nil
		} else {
			return "Invalid SMS OTP", 422, " -STR009", false, nil
		}
	}

}

func SubVerifyTriggerEmailOTP(ctx context.Context, client *ent.Client, empMasterRequest ca_reg.VerifyTriggerCandidateEmailOTP) (string, int32, string, bool, error) {
	if len(empMasterRequest.EmployeeId) == 0 {
		return "", 422, " -STR001", false, errors.New("enter User Name with eight digit number")
	}
	if empMasterRequest.EmailId == "" {
		return "", 422, " -STR002", false, errors.New("email ID is missing. ")
	}
	EmployeeIDNum, _ := strconv.ParseInt(empMasterRequest.EmployeeId, 10, 32)
	tx, err := client.Tx(ctx)
	if err != nil {
		return "", 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}
	//var dataExists   bool = false
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
			employeemaster.StatussEQ("active")).
		Exist(ctx)
	if err != nil {
		return "", 500, " -STR003", false, err
	}
	if !employee {
		return "", 422, " -STR004", false, errors.New("no details exists for this employee . ")
	} else {
		employee, err := tx.EmployeeMaster.Query().
			Where(employeemaster.EmployeeIDEQ(EmployeeIDNum),
				employeemaster.StatussEQ("active")).
			Only(ctx)
		if err != nil {
			return "", 500, " -STR005", false, err
		}
		if employee.VerifyStatus {
			return "", 422, " -STR006", false, errors.New("already request approved by your Controlling Authority. you can do First time Registration")
			//add one more validation to check in Usermaster
		} else if employee.FinalSubmitStatus {
			return "", 422, " -STR007", false, errors.New("contact your Controlling authority to approve your request")
		} else if !employee.SmsVerifyStatus {
			return "", 422, " -STR009", false, errors.New("mobile number not verified. First complete the Mobile number verification")
		}

		if employee.EmailOtp == empMasterRequest.EmailOTP {
			_, err = tx.EmployeeMaster.
				Update().
				Where(employeemaster.EmployeeID(EmployeeIDNum)).
				SetEmailVerifyStatus(true).
				SetCreatedDate(time.Now()).
				Save(ctx)

			if err != nil {
				return "", 500, " -STR010", false, err
			}

			if err = tx.Commit(); err != nil {
				tx.Rollback()
				return "", 500, " -STR011", false, err
			}

			return "Email OTP verified", 200, "", true, nil
		} else {
			return "Invalid Email OTP", 422, " -STR009", false, nil
		}
	}
}

func SubTriggerSMSOTP(ctx context.Context, client *ent.Client, empMasterRequest ca_reg.TriggerCandidateSMSOTP) (*ent.EmployeeMaster, int32, string, bool, error) {
	if len(empMasterRequest.EmployeeId) == 0 {
		return nil, 422, " -STR001", false, errors.New("enter User Name with eight digit number")
	}
	if empMasterRequest.EmployeeId == "" && len(empMasterRequest.EmployeeId) != 10 {
		return nil, 422, " -STR002", false, errors.New("mobile number is missing. ")
	}
	EmployeeIDNum, _ := strconv.ParseInt(empMasterRequest.EmployeeId, 10, 32)
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}
	//var dataExists   bool = false
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
			employeemaster.StatussEQ("active")).
		Exist(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}
	if employee {
		employee, err := tx.EmployeeMaster.Query().
			Where(employeemaster.EmployeeIDEQ(EmployeeIDNum),
				employeemaster.StatussEQ("active")).
			Only(ctx)
		if err != nil {
			return nil, 500, " -STR004", false, err
		}
		if employee.VerifyStatus {
			userExists, err := tx.UserMaster.Query().
				Where(usermaster.UserNameEQ(empMasterRequest.EmployeeId),
					usermaster.StatussEQ("active")).
				Exist(ctx)
			if err != nil {
				return nil, 500, " -STR005", false, err
			}
			if userExists {
				return nil, 422, " -STR006", false, errors.New("already you submitted employee master request and completed first time registration. Now you can proceed for login ")
			} else {
				return nil, 422, " -STR007", false, errors.New("already you submitted employee master request and approved by your Controlling Authority. Now you can proceed for First time Registration")
			}
			//add one more validation to check in Usermaster
		} else if employee.FinalSubmitStatus {
			return nil, 422, " -STR008", false, errors.New("already you submitted you employee master creation request. Contact your Controlling authority to approve your request")
		} else {
			_, err = tx.EmployeeMaster.Delete().Where(employeemaster.EmployeeIDEQ(EmployeeIDNum),
				employeemaster.StatussEQ("active")).Exec(ctx)
			if err != nil {
				return nil, 500, " -STR009", false, err
			}
		}
	}
	smsotp := util.GenerateOTP()

	otpGeneratedTime := time.Now()

	employeeUpdate, err := tx.EmployeeMaster.
		Create().
		SetEmployeeID(EmployeeIDNum).
		SetMobileNumber(empMasterRequest.MobileNumber).
		SetSmsOtp(int64(smsotp)).
		SetSmsTriggeredTime(otpGeneratedTime).
		SetSmsVerifyStatus(false).
		SetCreatedDate(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR010", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR011", false, err
	}

	return employeeUpdate, 200, "", true, nil

}

func SubTriggerEmailOTP(ctx context.Context, client *ent.Client, empMasterRequest ca_reg.TriggerCandidateEmailOTP) (*ent.EmployeeMaster, int32, string, bool, error) {
	if len(empMasterRequest.EmployeeId) == 0 {
		return nil, 422, " -STR001", false, errors.New("enter User Name with eight digit number")
	}
	if empMasterRequest.EmailId == "" {
		return nil, 422, " -STR002", false, errors.New("email ID is missing. ")
	}
	EmployeeIDNum, _ := strconv.ParseInt(empMasterRequest.EmployeeId, 10, 32)
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}
	//var dataExists   bool = false
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
			employeemaster.StatussEQ("active")).
		Exist(ctx)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}
	if employee {
		employee, err := tx.EmployeeMaster.Query().
			Where(employeemaster.EmployeeIDEQ(EmployeeIDNum),
				employeemaster.StatussEQ("active")).
			Only(ctx)
		if err != nil {
			return nil, 500, " -STR004", false, err
		}
		if employee.VerifyStatus {
			userExists, err := tx.UserMaster.Query().
				Where(usermaster.UserNameEQ(empMasterRequest.EmployeeId),
					usermaster.StatussEQ("active")).
				Exist(ctx)
			if err != nil {
				return nil, 500, " -STR005", false, err
			}
			if userExists {
				return nil, 422, " -STR006", false, errors.New("already you submitted employee master request and completed first time registration. You can login now")
			} else {
				return nil, 422, " -STR007", false, errors.New("already you submitted employee master request and approved by your Controlling Authority. you can do First time Registration")
			}
			//add one more validation to check in Usermaster
		} else if employee.FinalSubmitStatus {
			return nil, 422, " -STR008", false, errors.New("already you submitted you employee master creation request. Contact your Controlling authority to approve your request")
		} else if !employee.SmsVerifyStatus {
			return nil, 422, " -STR008", false, errors.New("mobile number not verified. First complete the Mobile number verification")
		}
	}

	emailotp := util.GenerateEmailOTP()

	emailOtpGeneratedTime := time.Now()

	_, err = tx.EmployeeMaster.
		Update().
		Where(employeemaster.EmployeeID(EmployeeIDNum),
			employeemaster.StatussEQ("active")).
		SetEmailID(empMasterRequest.EmailId).
		SetEmailVerifyStatus(false).
		SetEmailOtp(int64(emailotp)).
		SetEmailTriggeredTime(emailOtpGeneratedTime).
		SetCreatedDate(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR010", false, err
	}
	employeeUpdate, err := tx.EmployeeMaster.Query().
		Where(employeemaster.EmployeeIDEQ(EmployeeIDNum),
			employeemaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		return nil, 500, " -STR011", false, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR012", false, err
	}
	return employeeUpdate, 200, "", true, nil
}

// func SubCreateEmployeeMaster(ctx context.Context, client *ent.Client, empMasterRequest ca_reg.StrucCreateEmployeeMaster, z *ent.AdminMaster) (*ent.EmployeeMaster, int32, error) {
func SubCreateEmployeeMaster(ctx context.Context, client *ent.Client, empMasterRequest ca_reg.StrucCreateEmployeeMaster) (*ca_reg.EmployeeMasterResponse, int32, string, bool, error) {
	employeeIDStr := strconv.FormatInt(empMasterRequest.EmployeeID, 10)
	if len(employeeIDStr) != 8 {
		return nil, 422, " -STR001", false, errors.New("employee id should be 8 digit")
	}
	verifystatus := determineVerifyStatus(empMasterRequest.VerifyStatus)

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	defer handleTransaction(tx, &err)

	exists, err := checkEmployeeExists(ctx, tx, empMasterRequest.EmployeeID)
	if err != nil {
		return nil, 500, " -STR003", false, err
	}

	if !exists {
		return nil, 422, " -STR004", false, errors.New("no details exist for this employee")
	}

	employee, err := fetchActiveEmployee(ctx, tx, empMasterRequest.EmployeeID)
	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	err = validateEmployeeStatus(ctx, employee, employeeIDStr, tx)
	if err != nil {
		return nil, 422, "", false, err
	}

	status := determineStatus(empMasterRequest.Statuss)

	err = updateEmployee(ctx, tx, empMasterRequest, verifystatus, status)
	if err != nil {
		return nil, 500, " -STR006", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR007", false, err
	}

	response, err := buildEmployeeResponse(ctx, client, empMasterRequest.EmployeeID)
	if err != nil {
		return nil, 500, " -STR008", false, err
	}

	return response, 200, "", true, nil
}

func SubCreateEmployeeMasterbackup(ctx context.Context, client *ent.Client, empMasterRequest ca_reg.StrucCreateEmployeeMaster) (*ent.EmployeeMaster, int32, string, bool, error) {
	employeeIDStr := strconv.FormatInt(empMasterRequest.EmployeeID, 10)

	envmode := os.Getenv("ENV_MODE")
	var verifystatus bool
	if envmode == "production" || envmode == "uatdev" {
		verifystatus = empMasterRequest.VerifyStatus
	} else {
		verifystatus = true
	}

	if len(employeeIDStr) != 8 {
		return nil, 422, " -STR001", false, errors.New("employee id should be 8 digit")
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
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

	exists, err := tx.EmployeeMaster.
		Query().
		Where(employeemaster.EmployeeIDEQ(empMasterRequest.EmployeeID),
			employeemaster.StatussEQ("active"),
		).
		Exist(ctx)
	if err != nil {
		// Some error occurred during the query.
		return nil, 500, " -STR003", false, err
	}

	if exists {

		return nil, 422, " -STR004", false, fmt.Errorf("this employee ID %d exists in the employee master ", empMasterRequest.EmployeeID)
	}
	currentTime := time.Now().Truncate(time.Second)
	fmt.Println(currentTime)

	var statuss string
	if empMasterRequest.Statuss == "" {
		statuss = "active"
	} else {
		if empMasterRequest.Statuss != "rejected" {
			statuss += time.Now().Format("20060102150405")
		}
		statuss = empMasterRequest.Statuss
	}

	_, err = tx.EmployeeMaster.
		Create().
		SetEmployeeID(empMasterRequest.EmployeeID).
		SetEmployeeName(empMasterRequest.EmployeeName).
		SetDOB(empMasterRequest.DOB).
		SetGender(empMasterRequest.Gender).
		SetMobileNumber(empMasterRequest.MobileNumber).
		SetEmailID(empMasterRequest.EmailId).
		SetEmployeeCategory(empMasterRequest.EmployeeCategory).
		SetEmployeePost(empMasterRequest.EmployeePost).
		SetFacilityID(empMasterRequest.FacilityId).
		SetPincode(empMasterRequest.Pincode).
		SetOfficeName(empMasterRequest.OfficeName).
		SetControllingAuthorityFacilityId(empMasterRequest.ControllingAuthorityFacilityId).
		SetControllingAuthorityName(empMasterRequest.ControllingAuthorityName).
		SetNodalAuthorityFaciliyId(empMasterRequest.NodalAuthorityFaciliyId).
		SetNodalAuthorityName(empMasterRequest.NodalAuthorityName).
		SetCircleFacilityID(empMasterRequest.CircleFacilityId).
		SetCreatedById(empMasterRequest.CreatedById).
		SetCreatedByUserName(empMasterRequest.CreatedByUserName).
		SetCreatedByEmpId(empMasterRequest.CreatedByEmpId).
		SetCreatedByDesignation(empMasterRequest.CreatedByDesignation).
		SetVerifyStatus(verifystatus).
		SetStatuss(statuss).
		SetCreatedDate(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, 500, " -STR005", false, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, " -STR006", false, err
	}
	empuser, err := client.EmployeeMaster.Query().
		Where(
			employeemaster.EmployeeIDEQ(empMasterRequest.EmployeeID),
			employeemaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, 500, " -STR009", false, err
		}
	}
	return empuser, 200, "", true, nil

}

func SubModifyEmployeeMaster(client *ent.Client, empMasterRequest ca_reg.StrucModifyEmployeeMaster) (*ca_reg.EmployeeMasterResponse, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		handleTransaction(tx, &err)
	}()

	exists, err := checkIfEmployeeActive(ctx, tx, empMasterRequest.EmployeeID)
	if err != nil {
		return nil, 500, "-STR001", false, err
	}
	if exists {
		return nil, 422, "-STR002", false, fmt.Errorf("this employee ID %d already in active state, hence modification not allowed", empMasterRequest.EmployeeID)
	}

	// Query the user by EmployeeID

	user, err := fetchActiveEmployee(ctx, tx, empMasterRequest.EmployeeID)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, 500, "-STR003", false, err
		}
	}
	/* 	if !user.VerifyStatus {
		return nil, 422, " -STR004", false, fmt.Errorf("kindly use 'Approve Candidate' option to process this employee ID details")
	} */

	if empMasterRequest.VerifyStatus {
		response, err := processVerifyStatus(ctx, tx, user, empMasterRequest)
		return response, 200, "", true, err
	} else {
		response, err := processDeletionStatus(ctx, tx, user, empMasterRequest)
		return response, 200, "", true, err
	}
}
func SubGetEmployeesBasedOnCA(ctx context.Context, client *ent.Client, caFacilityId string) ([]*ent.EmployeeMaster, int32, string, bool, error) {
	//Can use GetX as well
	employees, err := client.EmployeeMaster.Query().
		Where(employeemaster.ControllingAuthorityFacilityIdEQ(caFacilityId),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(false),
		).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	if len(employees) == 0 {
		return nil, 422, " -STR002", false, fmt.Errorf("no pending employees found for th given CA facility ID")
	}
	// user is not nil and the query was successful
	return employees, 200, "", true, nil
}

func SubViewEmployeeMaster(ctx context.Context, client *ent.Client, empid int64) ([]*ent.EmployeeMaster, int32, string, bool, error) {
	//Can use GetX as well
	user, err := client.EmployeeMaster.Query().
		Where(employeemaster.EmployeeIDEQ(empid),
			employeemaster.StatussEQ("active"),
		).
		All(ctx)

	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	if len(user) == 0 {
		return nil, 422, " -STR002", false, fmt.Errorf("no active employee found with given employee ID")
	}
	// user is not nil and the query was successful
	return user, 200, "", true, nil
}
