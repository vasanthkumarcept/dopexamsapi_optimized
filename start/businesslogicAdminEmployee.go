package start

import (
	"context"
	"errors"
	"fmt"
	"os"
	"recruit/ent"
	"recruit/ent/adminmaster"
	"recruit/ent/employeemaster"
	"recruit/ent/usermaster"
	"time"

	ca_reg "recruit/payloadstructure/candidate_registration"
)

type customError struct {
	Code       string
	Message    string
	StatusCode int32
}

func (e *customError) Error() string {
	return e.Message
}

func checkDuplicateAdmin(ctx context.Context, tx *ent.Tx, adminRequest ca_reg.AdminCreation) error {
	exists, err := tx.AdminMaster.
		Query().
		Where(
			adminmaster.EmployeeIdEQ(adminRequest.EmployeeId),
			adminmaster.RoleUserCodeEQ(adminRequest.RoleUserCode),
			adminmaster.FacilityIDEQ(adminRequest.FacilityID),
			adminmaster.StatussEQ("active"),
		).Exist(ctx)
	if err != nil {
		return fmt.Errorf("failed to check duplicate admin: %w", err)
	}
	if exists {
		return fmt.Errorf("already username exists for this EmployeeId: %v , Facility ID: %v , Role User Code: %v",
			adminRequest.EmployeeId, adminRequest.FacilityID, adminRequest.RoleUserCode)
	}
	return nil
}
func checkExistingAdmin(ctx context.Context, tx *ent.Tx, userName string) error {
	exists, err := tx.AdminMaster.
		Query().
		Where(
			adminmaster.UserNameEQ(userName),
			adminmaster.StatussEQ("active"),
		).Exist(ctx)
	if err != nil {
		return fmt.Errorf("failed to check existing admin: %w", err)
	}
	if exists {
		return fmt.Errorf("admin user: %s already exists", userName)
	}
	return nil
}
func createAdminRecord(ctx context.Context, tx *ent.Tx, adminRequest ca_reg.AdminCreation, currentTime time.Time) (*ent.AdminMaster, error) {
	admin, err := tx.AdminMaster.
		Create().
		SetEmployeeId(adminRequest.EmployeeId).
		SetEmployeeName(adminRequest.EmployeeName).
		SetDesignation(adminRequest.Designation).
		SetFacilityID(adminRequest.FacilityID).
		SetFacilityType(adminRequest.FacilityType).
		SetUserName(adminRequest.UserName).
		SetRoleUserCode(adminRequest.RoleUserCode).
		SetRoleUserDescription(adminRequest.RoleUserDescription).
		SetFacilityIDUniqueid(adminRequest.FacilityIDUniqueid).
		SetReportingOfficeFacilityId(adminRequest.ReportingOfficeFacilityId).
		SetReportingOfficeFacilityName(adminRequest.ReportingOfficeFacilityName).
		SetCircleOfficeFacilityId(adminRequest.CircleOfficeFacilityId).
		SetCircleOfficeName(adminRequest.CircleOfficeName).
		SetMobile(adminRequest.Mobile).
		SetEmailID(adminRequest.EmailID).
		SetAuthorityFacilityName(adminRequest.AuthorityFacilityName).
		SetCreatedById(int64(adminRequest.CreatedById)).
		SetCreatedByUserName(adminRequest.CreatedByUserName).
		SetCreatedByEmpId(int64(adminRequest.CreatedByEmpId)).
		SetCreatedByDesignation(adminRequest.CreatedByDesignation).
		SetCreatedDate(currentTime).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create admin record: %w", err)
	}
	return admin, nil
}
func mapToAdminResponse(admin *ent.AdminMaster) *ca_reg.AdminMasterResponse {
	return &ca_reg.AdminMasterResponse{
		ID:                          admin.ID,
		EmployeeId:                  admin.EmployeeId,
		EmployeeName:                admin.EmployeeName,
		Designation:                 admin.Designation,
		RoleUserCode:                admin.RoleUserCode,
		RoleUserDescription:         admin.RoleUserDescription,
		Mobile:                      admin.Mobile,
		EmailID:                     admin.EmailID,
		UserName:                    admin.UserName,
		FacilityIDUniqueID:          admin.FacilityIDUniqueid,
		FacilityID:                  admin.FacilityID,
		AuthorityFacilityName:       admin.AuthorityFacilityName,
		FacilityType:                admin.FacilityType,
		ReportingOfficeFacilityID:   admin.ReportingOfficeFacilityId,
		ReportingOfficeFacilityName: admin.ReportingOfficeFacilityName,
		CircleOfficeFacilityID:      admin.CircleOfficeFacilityId,
		CircleOfficeName:            admin.CircleOfficeName,
		Status:                      "active",
		CreatedByID:                 admin.CreatedById,
		CreatedByUserName:           admin.CreatedByUserName,
		CreatedByEmpID:              admin.CreatedByEmpId,
		CreatedByDesignation:        admin.CreatedByDesignation,
		CreatedDate:                 admin.CreatedDate,
	}
}
func mapToAdminMasterResponse(admin *ent.AdminMaster) *ca_reg.AdminMasterResponse {
	return &ca_reg.AdminMasterResponse{
		ID:                          admin.ID,
		EmployeeId:                  admin.EmployeeId,
		EmployeeName:                admin.EmployeeName,
		Designation:                 admin.Designation,
		RoleUserCode:                admin.RoleUserCode,
		RoleUserDescription:         admin.RoleUserDescription,
		Mobile:                      admin.Mobile,
		EmailID:                     admin.EmailID,
		UserName:                    admin.UserName,
		
	}
}
func determineVerifyStatus(requestedStatus bool) bool {
	envmode := os.Getenv("ENV_MODE")
	if envmode == "production" || envmode == "uatdev" {
		return requestedStatus
	}
	return true
}
func checkEmployeeExists(ctx context.Context, tx *ent.Tx, employeeID int64) (bool, error) {
	return tx.EmployeeMaster.
		Query().
		Where(
			employeemaster.EmployeeIDEQ(employeeID),
			employeemaster.StatussEQ("active"),
		).
		Exist(ctx)
}
func fetchActiveEmployee(ctx context.Context, tx *ent.Tx, employeeID int64) (*ent.EmployeeMaster, error) {
	return tx.EmployeeMaster.
		Query().
		Where(
			employeemaster.EmployeeIDEQ(employeeID),
			employeemaster.StatussEQ("active"),
		).
		Only(ctx)
}

func validateEmployeeStatus(ctx context.Context, employee *ent.EmployeeMaster, employeeIDStr string, tx *ent.Tx) error {
	if employee.VerifyStatus {
		exists, err := tx.UserMaster.
			Query().
			Where(
				usermaster.UserNameEQ(employeeIDStr),
				usermaster.StatussEQ("active"),
			).
			Exist(ctx)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("already submitted employee master request and completed first-time registration. You can log in now")
		}
		return errors.New("already submitted employee master request and approved by your Controlling Authority. You can do First-time Registration")
	} else if employee.FinalSubmitStatus {
		return errors.New("contact your Controlling Authority to approve your request")
	} else if !employee.SmsVerifyStatus {
		return errors.New("mobile number not verified. Complete Mobile number verification first")
	} else if !employee.EmailVerifyStatus {
		return errors.New("email ID not verified. Complete email ID verification first")
	}
	return nil
}

func determineStatus(requestedStatus string) string {
	if requestedStatus == "" {
		return "active"
	}
	if requestedStatus != "rejected" {
		return requestedStatus + time.Now().Format("20060102150405")
	}
	return requestedStatus
}

func updateEmployee(ctx context.Context, tx *ent.Tx, empMasterRequest ca_reg.StrucCreateEmployeeMaster, verifystatus bool, status string) error {
	_, err := tx.EmployeeMaster.
		Update().
		Where(
			employeemaster.EmployeeIDEQ(empMasterRequest.EmployeeID),
			employeemaster.StatussEQ("active"),
		).
		SetEmployeeName(empMasterRequest.EmployeeName).
		SetDOB(empMasterRequest.DOB).
		SetGender(empMasterRequest.Gender).
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
		SetStatuss(status).
		SetFinalSubmitStatus(true).
		SetCreatedDate(time.Now()).
		Save(ctx)
	return err
}

func buildEmployeeResponse(ctx context.Context, client *ent.Client, employeeID int64) (*ca_reg.EmployeeMasterResponse, error) {
	empuser, err := client.EmployeeMaster.Query().
		Where(
			employeemaster.EmployeeIDEQ(employeeID),
			employeemaster.StatussEQ("active"),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &ca_reg.EmployeeMasterResponse{
		ID:                             empuser.ID,
		EmployeeID:                     empuser.EmployeeID,
		EmployeeName:                   empuser.EmployeeName,
		DOB:                            empuser.DOB,
		Gender:                         empuser.Gender,
		MobileNumber:                   empuser.MobileNumber,
		EmailID:                        empuser.EmailID,
		EmployeeCategoryCode:           empuser.EmployeeCategoryCode,
		EmployeeCategory:               empuser.EmployeeCategory,
		PostCode:                       empuser.PostCode,
		EmployeePost:                   empuser.EmployeePost,
		FacilityID:                     empuser.FacilityID,
		OfficeName:                     empuser.OfficeName,
		Pincode:                        empuser.Pincode,
		ControllingAuthorityFacilityId: empuser.ControllingAuthorityFacilityId,
		ControllingAuthorityName:       empuser.ControllingAuthorityName,
		NodalAuthorityFacilityId:       empuser.NodalAuthorityFaciliyId,
		NodalAuthorityName:             empuser.NodalAuthorityName,
		Status:                         empuser.Statuss,
		VerifyStatus:                   empuser.VerifyStatus,
		FinalSubmitStatus:              empuser.FinalSubmitStatus,
	}, nil
}
func processVerifyStatus(ctx context.Context, tx *ent.Tx, user *ent.EmployeeMaster, empMasterRequest ca_reg.StrucModifyEmployeeMaster) (*ca_reg.EmployeeMasterResponse, error) {
	stat := "inactive_" + time.Now().Format("20060102150405")

	_, err := user.Update().
		SetStatuss(stat).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.EmployeeMaster.Create().
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
		SetCreatedById(user.CreatedById).
		SetCreatedByUserName(user.CreatedByUserName).
		SetCreatedByEmpId(user.CreatedByEmpId).
		SetCreatedByDesignation(user.CreatedByDesignation).
		SetCreatedDate(user.CreatedDate).
		SetModifiedById(empMasterRequest.ModifiedById).
		SetModifiedByUserName(empMasterRequest.ModifiedByUserName).
		SetModifiedByEmpId(empMasterRequest.ModifiedByEmpId).
		SetModifiedByDesignantion(empMasterRequest.ModifiedByDesignantion).
		SetModifiedDate(time.Now().UTC().Truncate(24 * time.Hour)).
		SetVerifyStatus(empMasterRequest.VerifyStatus).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return fetchEmployeeResponse(ctx, tx, empMasterRequest.EmployeeID, "active")
}

func processDeletionStatus(ctx context.Context, tx *ent.Tx, user *ent.EmployeeMaster, empMasterRequest ca_reg.StrucModifyEmployeeMaster) (*ca_reg.EmployeeMasterResponse, error) {
	stat := "deleted_" + time.Now().Format("20060102150405")

	_, err := user.Update().
		SetStatuss(stat).
		SetModifiedById(empMasterRequest.ModifiedById).
		SetModifiedByUserName(empMasterRequest.ModifiedByUserName).
		SetModifiedByEmpId(empMasterRequest.ModifiedByEmpId).
		SetModifiedByDesignantion(empMasterRequest.ModifiedByDesignantion).
		SetModifiedDate(time.Now().UTC().Truncate(24 * time.Hour)).
		SetCadre(empMasterRequest.Remarks).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return fetchEmployeeResponse(ctx, tx, empMasterRequest.EmployeeID, stat)
}

func fetchEmployeeResponse(ctx context.Context, tx *ent.Tx, employeeID int64, status string) (*ca_reg.EmployeeMasterResponse, error) {
	empuser, err := tx.EmployeeMaster.Query().
		Where(
			employeemaster.EmployeeIDEQ(employeeID),
			employeemaster.StatussEQ(status),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &ca_reg.EmployeeMasterResponse{
		ID:                             empuser.ID,
		EmployeeID:                     empuser.EmployeeID,
		EmployeeName:                   empuser.EmployeeName,
		DOB:                            empuser.DOB,
		Gender:                         empuser.Gender,
		MobileNumber:                   empuser.MobileNumber,
		EmailID:                        empuser.EmailID,
		EmployeeCategoryCode:           empuser.EmployeeCategoryCode,
		EmployeeCategory:               empuser.EmployeeCategory,
		PostCode:                       empuser.PostCode,
		EmployeePost:                   empuser.EmployeePost,
		FacilityID:                     empuser.FacilityID,
		OfficeName:                     empuser.OfficeName,
		Pincode:                        empuser.Pincode,
		ControllingAuthorityFacilityId: empuser.ControllingAuthorityFacilityId,
		ControllingAuthorityName:       empuser.ControllingAuthorityName,
		NodalAuthorityFacilityId:       empuser.NodalAuthorityFaciliyId,
		NodalAuthorityName:             empuser.NodalAuthorityName,
		Status:                         empuser.Statuss,
		VerifyStatus:                   empuser.VerifyStatus,
		FinalSubmitStatus:              empuser.FinalSubmitStatus,
	}, nil
}
func checkIfEmployeeActive(ctx context.Context, tx *ent.Tx, employeeID int64) (bool, error) {
	return tx.UserMaster.Query().
		Where(
			usermaster.EmployeeIDEQ(employeeID),
			usermaster.StatussEQ("active"),
			usermaster.StatusEQ(true),
		).
		Exist(ctx)
}
func createUserResponse(user *ent.UserMaster) *ca_reg.UserMasterResponse {
	return &ca_reg.UserMasterResponse{
		ID:                 user.ID,
		UserName:           user.UserName,
		Status:             user.Status,
		Mobile:             user.Mobile,
		DOB:                user.DOB,
		Statuss:            user.Statuss,
		EmailID:            user.EmailID,
		EmployeeID:         user.EmployeeID,
		EmployeeName:       user.EmployeeName,
		FacilityID:         user.FacilityID,
		CircleFacilityId:   user.CircleFacilityId,
		CircleFacilityName: user.CircleFacilityName,
		RoleUserCode:       user.RoleUserCode,
	}
}
func fetchUser(tx *ent.Tx, ctx context.Context, username string) (*ent.UserMaster, error) {
	user, err := tx.UserMaster.
		Query().
		Where(usermaster.UserName(username),
			usermaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("no such username exists")
		}
		return nil, err
	}
	return user, nil
}
func validateUserStatus(user *ent.UserMaster) error {
	if user.Status {
		return errors.New(" -STR005: already registration done and in active status")
	}
	return nil
}
func updateUser(tx *ent.Tx, ctx context.Context, user *ent.UserMaster, password string) error {
	user.Password = password
	user.Status = true
	user.CreatedAt = time.Now()
	user.OTPSavedTime = time.Now()

	_, err := tx.UserMaster.
		Update().
		Where(usermaster.EmployeeIDEQ(user.EmployeeID)).
		SetPassword(user.Password).
		SetStatus(user.Status).
		SetCreatedAt(user.CreatedAt).
		SetOTPSavedTime(user.OTPSavedTime).
		Save(ctx)
	return err
}
func fetchActiveUser(tx *ent.Tx, ctx context.Context, userName string) (*ent.UserMaster, error) {
	user, err := tx.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(userName),
			usermaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("no such username exists in active")
		}
		return nil, err
	}
	return user, nil
}
func updateUserOTP(user *ent.UserMaster, ctx context.Context, otp int32, otpExpiryTime time.Time) (*ent.UserMaster, error) {
	return user.Update().
		SetOTP(otp).
		SetOTPExpiryTime(otpExpiryTime).
		Save(ctx)
}
func updateUserPassword(tx *ent.Tx, ctx context.Context, newUsers ca_reg.StrucUserResetSaveNewPassword) error {
	_, err := tx.UserMaster.
		Update().
		Where(
			usermaster.UserNameEQ(newUsers.UserName),
			usermaster.StatussEQ("active")).
		SetPassword(newUsers.NewPassword).
		Save(ctx)
	return err
}
func validateAdminLoginInputs(newUser *ent.AdminMaster) error {
	if newUser == nil {
		return errors.New("UserMaster cannot be nil")
	}
	if newUser.OTP <= 0 {
		return errors.New("OTP cannot be nil")
	}
	
	return nil
}
func fetchAdminUser(tx *ent.Tx, userName string, ctx context.Context) (*ent.AdminMaster, error) {
	user, err := tx.AdminMaster.
		Query().
		Where(adminmaster.UserNameEQ(userName), adminmaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("no Admin user found for this user %s", userName)
		}
		return nil, err
	}
	return user, nil
}
func updateAdminToken(tx *ent.Tx, userName, token string, ctx context.Context) error {
	_, err := tx.AdminMaster.
		Update().
		SetUidToken(token).
		Where(adminmaster.UserNameEQ(userName)).
		Save(ctx)
	return err
}
func fetchAdminUserr(tx *ent.Tx, userName string) (*ent.AdminMaster, error) {
	user, err := tx.AdminMaster.
		Query().
		Where(
			adminmaster.UserNameEQ(userName),
			adminmaster.RoleUserCodeIn(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17),
			adminmaster.StatussEQ("active"),
		).
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("no active Admin user found for this user %s", userName)
		}
		return nil, err
	}
	return user, nil
}
func fetchActiveAdminUser(ctx context.Context, tx *ent.Tx, userName string) (*ent.AdminMaster, error) {
	user, err := tx.AdminMaster.Query().Where(adminmaster.UserNameEQ(userName), adminmaster.StatussEQ("active")).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("admin user : %s not exists", userName)
		}
		return nil, err
	}
	return user, nil
}
func updateAdminPassword(ctx context.Context, tx *ent.Tx, userName, newPassword string) error {
	_, err := tx.AdminMaster.
		Update().
		Where(adminmaster.UserNameEQ(userName), adminmaster.StatussEQ("active")).
		SetPassword(newPassword).
		Save(ctx)
	return err
}
func updateUserStatus(user *ent.AdminMaster, status string, ctx context.Context) error {
	_, err := user.Update().SetStatuss(status).Save(ctx)
	return err
}
func checkduplicateAdmin(tx *ent.Tx, newUser ca_reg.UpdateAdminMasterStruc, ctx context.Context) (bool, error) {
	return tx.AdminMaster.Query().
		Where(
			adminmaster.EmployeeIdEQ(newUser.EmployeeId),
			adminmaster.RoleUserCodeEQ(newUser.RoleUserCode),
			adminmaster.FacilityIDEQ(newUser.FacilityID),
			adminmaster.StatussEQ("active"),
		).Exist(ctx)
}
