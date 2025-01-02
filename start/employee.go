package start

import (
	"context"
	"errors"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/employeemaster"
	"recruit/util"
)

func CreateEmployeeProfile(client *ent.Client, newEmployeeprofile *ent.Employees) (*ent.Employees, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.Employees.Create().
		SetEmployeedID(newEmployeeprofile.EmployeedID).
		SetIDVerified(newEmployeeprofile.IDVerified).
		SetIDRemStatus(newEmployeeprofile.IDRemStatus).
		SetIDRemarks(newEmployeeprofile.IDRemarks).
		SetEmployeeName(newEmployeeprofile.EmployeeName).
		SetNameVerified(newEmployeeprofile.NameVerified).
		SetNameRemStatus(newEmployeeprofile.NameRemStatus).
		SetNameRemarks(newEmployeeprofile.NameRemarks).
		SetEmployeeFathersName(newEmployeeprofile.EmployeeFathersName).
		SetFathersNameVerified(newEmployeeprofile.FathersNameVerified).
		SetFathersNameRemStatus(newEmployeeprofile.FathersNameRemStatus).
		SetFathersNameRemarks(newEmployeeprofile.FathersNameRemarks).
		SetDOB(newEmployeeprofile.DOB).
		SetDOBVerified(newEmployeeprofile.DOBVerified).
		SetDOBRemStatus(newEmployeeprofile.DOBRemStatus).
		SetDOBRemarks(newEmployeeprofile.DOBRemarks).
		SetGender(newEmployeeprofile.Gender).
		SetGenderVerified(newEmployeeprofile.GenderVerified).
		SetGenderRemStatus(newEmployeeprofile.GenderRemStatus).
		SetGenderRemarks(newEmployeeprofile.GenderRemarks).
		SetMobileNumber(newEmployeeprofile.MobileNumber).
		SetMobileNumberVerified(newEmployeeprofile.MobileNumberVerified).
		SetMobileNumberRemStatus(newEmployeeprofile.MobileNumberRemStatus).
		SetMobileNumberRemarks(newEmployeeprofile.MobileNumberRemarks).
		SetEmailID(newEmployeeprofile.EmailID).
		SetEmailIDVerified(newEmployeeprofile.EmailIDVerified).
		SetEmailIDRemStatus(newEmployeeprofile.EmailIDRemStatus).
		SetEmailIDRemarks(newEmployeeprofile.EmailIDRemarks).
		SetCategoryid(newEmployeeprofile.Categoryid).
		SetEmployeeCategoryCode(newEmployeeprofile.EmployeeCategoryCode).
		SetEmployeeCategory(newEmployeeprofile.EmployeeCategory).
		SetEmployeeCategoryCodeVerified(newEmployeeprofile.EmployeeCategoryCodeVerified).
		SetEmployeeCategoryCodeRemStatus(newEmployeeprofile.EmployeeCategoryCodeRemStatus).
		SetEmployeeCategoryCodeRemarks(newEmployeeprofile.EmployeeCategoryCodeRemarks).
		SetWithDisability(newEmployeeprofile.WithDisability).
		SetWithDisabilityVerified(newEmployeeprofile.WithDisabilityVerified).
		SetWithDisabilityRemStatus(newEmployeeprofile.WithDisabilityRemStatus).
		SetWithDisabilityRemarks(newEmployeeprofile.WithDisabilityRemarks).
		SetDisabilityType(newEmployeeprofile.DisabilityType).
		SetDisabilityTypeVerified(newEmployeeprofile.DisabilityTypeVerified).
		SetDisabilityTypeRemStatus(newEmployeeprofile.DisabilityTypeRemStatus).
		SetDisabilityTypeRemarks(newEmployeeprofile.DisabilityTypeRemarks).
		SetDisabilityPercentage(newEmployeeprofile.DisabilityPercentage).
		SetDisabilityPercentageVerified(newEmployeeprofile.DisabilityPercentageVerified).
		SetDisabilityPercentageRemStatus(newEmployeeprofile.DisabilityPercentageRemStatus).
		SetDisabilityPercentageRemarks(newEmployeeprofile.DisabilityPercentageRemarks).
		//SetEmployeeCadre(newEmployeeprofile.EmployeeCadre).
		SetSignature(newEmployeeprofile.Signature).
		SetSignatureVerified(newEmployeeprofile.SignatureVerified).
		SetSignatureRemStatus(newEmployeeprofile.SignatureRemStatus).
		SetSignatureRemarks(newEmployeeprofile.SignatureRemarks).
		SetPhoto(newEmployeeprofile.Photo).
		SetPhotoVerified(newEmployeeprofile.PhotoVerified).
		SetPhotoRemStatus(newEmployeeprofile.PhotoRemStatus).
		SetPhotoRemarks(newEmployeeprofile.PhotoRemarks).
		SetPostID(newEmployeeprofile.PostID).
		SetPostCode(newEmployeeprofile.PostCode).
		SetEmployeePost(newEmployeeprofile.EmployeePost).
		SetEmployeePostRemarks(newEmployeeprofile.EmployeePostRemarks).
		SetEmployeePostRemStatus(newEmployeeprofile.EmployeePostRemStatus).
		SetEmployeePostVerified(newEmployeeprofile.EmployeePostVerified).
		SetDesignationID(newEmployeeprofile.DesignationID).
		SetEmployeeDesignation(newEmployeeprofile.EmployeeDesignation).
		SetEmployeeDesignationVerified(newEmployeeprofile.EmployeeDesignationVerified).
		SetEmployeeDesignationRemStatus(newEmployeeprofile.EmployeeDesignationRemStatus).
		SetEmployeeDesignationRemarks(newEmployeeprofile.EmployeeDesignationRemarks).
		SetCircleID(newEmployeeprofile.CircleID).
		SetCircleName(newEmployeeprofile.CircleName).
		SetCircleVerified(newEmployeeprofile.CircleVerified).
		SetCircleRemStatus(newEmployeeprofile.CircleRemStatus).
		SetCircleRemarks(newEmployeeprofile.CircleRemarks).
		SetRegionID(newEmployeeprofile.RegionID).
		SetRegionName(newEmployeeprofile.RegionName).
		SetRegionVerified(newEmployeeprofile.RegionVerified).
		SetRegionRemStatus(newEmployeeprofile.RegionRemStatus).
		SetRegionRemarks(newEmployeeprofile.RegionRemarks).
		SetDivisionID(newEmployeeprofile.DivisionID).
		SetDivisionName(newEmployeeprofile.DivisionName).
		SetDivisionVerified(newEmployeeprofile.DivisionVerified).
		SetDivisionRemStatus(newEmployeeprofile.DivisionRemStatus).
		SetDivisionRemarks(newEmployeeprofile.DivisionRemarks).
		SetOfficeID(newEmployeeprofile.OfficeID).
		SetOfficeName(newEmployeeprofile.OfficeName).
		SetOfficeVerified(newEmployeeprofile.OfficeVerified).
		SetOfficeRemStatus(newEmployeeprofile.OfficeRemStatus).
		SetOfficeRemarks(newEmployeeprofile.OfficeRemarks).
		SetRole(newEmployeeprofile.Role).
		SetRoleVerified(newEmployeeprofile.RoleVerified).
		SetRoleRemStatus(newEmployeeprofile.RoleRemStatus).
		SetRoleRemarks(newEmployeeprofile.RoleRemarks).
		SetDCCS(newEmployeeprofile.DCCS).
		SetDCCSVerified(newEmployeeprofile.DCCSVerified).
		SetDCCSRemStatus(newEmployeeprofile.DCCSRemStatus).
		SetDCCSRemarks(newEmployeeprofile.DCCSRemarks).
		SetDCInPresentCadre(newEmployeeprofile.DCInPresentCadre).
		SetDCInPresentCadreVerified(newEmployeeprofile.DCInPresentCadreVerified).
		SetDCInPresentCadreRemStatus(newEmployeeprofile.DCInPresentCadreRemStatus).
		SetDCInPresentCadreRemarks(newEmployeeprofile.DCInPresentCadreRemarks).
		SetAPSWorking(newEmployeeprofile.APSWorking).
		SetAPSWorkingVerified(newEmployeeprofile.APSWorkingVerified).
		SetAPSWorkingRemStatus(newEmployeeprofile.APSWorkingRemStatus).
		SetAPSWorkingRemarks(newEmployeeprofile.APSWorkingRemarks).
		SetProfilestatus(newEmployeeprofile.Profilestatus).
		Save(ctx)

	if err != nil {
		log.Println("error at Creating Employees: ", newEmployeeprofile)
		return nil, fmt.Errorf("failed creating Employees: %w", err)
	}
	log.Println("Employees profile is created: ", u)

	return u, nil
}

func QueryEmployeesWithID(ctx context.Context, client *ent.Client, id int32) (*ent.Employees, error) {
	//Can use GetX as well

	employees, err := client.Employees.Get(ctx, id)
	if err != nil {
		log.Println("error at getting EmployeeID: ", err)
		return nil, fmt.Errorf("failed querying Employees: %w", err)
	}
	log.Println("Employee returned: ", employees)
	return employees, nil
}

func QueryEmployees(ctx context.Context, client *ent.Client) ([]*ent.Employees, error) {
	//Array of exams
	newemployee, err := client.Employees.Query().
		All(ctx)
	if err != nil {
		log.Println("error at ExamCalendarID: ", err)
		return nil, fmt.Errorf("failed querying Employees: %w", err)
	}
	log.Println(" Employees data returned: ", newemployee)
	return newemployee, nil
}

//Update fields by CA

func UpdateVerificationDetails(client *ent.Client, id int32, newEmployeeprofile *ent.Employees) (*ent.Employees, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	_, err := QueryEmployeesWithID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	updatedEmployee, err := client.Employees.UpdateOneID(id).
		SetIDVerified(newEmployeeprofile.IDVerified).
		SetIDRemarks(newEmployeeprofile.IDRemarks).
		SetNameVerified(newEmployeeprofile.NameVerified).
		SetNameRemarks(newEmployeeprofile.NameRemarks).
		SetFathersNameVerified(newEmployeeprofile.FathersNameVerified).
		SetFathersNameRemarks(newEmployeeprofile.FathersNameRemarks).
		SetDOBVerified(newEmployeeprofile.DOBVerified).
		SetDOBRemarks(newEmployeeprofile.DOBRemarks).
		SetGenderVerified(newEmployeeprofile.GenderVerified).
		SetGenderRemarks(newEmployeeprofile.GenderRemarks).
		SetMobileNumber(newEmployeeprofile.MobileNumber).
		SetEmailID(newEmployeeprofile.EmailID).
		SetEmployeeCategoryCodeVerified(newEmployeeprofile.EmployeeCategoryCodeVerified).
		SetEmployeeCategoryCodeRemarks(newEmployeeprofile.EmployeeCategoryCodeRemarks).
		SetWithDisabilityVerified(newEmployeeprofile.WithDisabilityVerified).
		SetWithDisabilityRemarks(newEmployeeprofile.WithDisabilityRemarks).
		SetDisabilityTypeVerified(newEmployeeprofile.DisabilityTypeVerified).
		SetDisabilityTypeRemarks(newEmployeeprofile.DisabilityTypeRemarks).
		SetDisabilityPercentageVerified(newEmployeeprofile.DisabilityPercentageVerified).
		SetDisabilityPercentageRemarks(newEmployeeprofile.DisabilityPercentageRemarks).
		SetSignatureVerified(newEmployeeprofile.SignatureVerified).
		SetSignatureRemarks(newEmployeeprofile.SignatureRemarks).
		SetPhotoVerified(newEmployeeprofile.PhotoVerified).
		SetPhotoRemarks(newEmployeeprofile.PhotoRemarks).
		//SetPostID(newEmployeeprofile.PostID).
		//SetPostCode(newEmployeeprofile.PostCode).
		SetEmployeePostVerified(newEmployeeprofile.EmployeePostVerified).
		SetEmployeePostRemarks(newEmployeeprofile.EmployeePostRemarks).

		//SetPostCodeRemStatus(newEmployeeprofile.PostCodeRemStatus).
		//SetPostCodeRemarks(newEmployeeprofile.PostCodeRemarks).
		SetEmployeeDesignationVerified(newEmployeeprofile.EmployeeDesignationVerified).
		SetEmployeeDesignationRemarks(newEmployeeprofile.EmployeeDesignationRemarks).
		SetCircleVerified(newEmployeeprofile.CircleVerified).
		SetCircleRemarks(newEmployeeprofile.CircleRemarks).
		SetRegionVerified(newEmployeeprofile.RegionVerified).
		SetRegionRemarks(newEmployeeprofile.RegionRemarks).
		SetDivisionVerified(newEmployeeprofile.DivisionVerified).
		SetDivisionRemarks(newEmployeeprofile.DivisionRemarks).
		SetOfficeVerified(newEmployeeprofile.OfficeVerified).
		SetOfficeRemarks(newEmployeeprofile.OfficeRemarks).
		SetRoleVerified(newEmployeeprofile.RoleVerified).
		SetRoleRemarks(newEmployeeprofile.RoleRemarks).
		SetDCCSVerified(newEmployeeprofile.DCCSVerified).
		SetDCCSRemarks(newEmployeeprofile.DCCSRemarks).
		SetDCInPresentCadreVerified(newEmployeeprofile.DCInPresentCadreVerified).
		SetDCInPresentCadreRemarks(newEmployeeprofile.DCInPresentCadreRemarks).
		SetAPSWorkingVerified(newEmployeeprofile.APSWorkingVerified).
		SetAPSWorkingRemarks(newEmployeeprofile.APSWorkingRemarks).
		SetProfilestatus(newEmployeeprofile.Profilestatus).
		Save(context.Background())

	//.Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedEmployee, nil
}

// ResubmitDetails by Candidate
func UpdateEmpDetailsByEmpID(client *ent.Client, empid int32, newEmployeeprofile *ent.Employees) (*ent.Employees, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	emps, err := QueryEmployeesWithID(ctx, client, empid)
	/*emps, err := client.Employees.Query().
	Where(employees.EmployeedIDEQ(empid)).
	All(ctx)*/

	if err != nil {
		return nil, err
	}

	updatedEmployee, err := client.Employees.UpdateOneID(emps.EmployeedID).
		//.UpdateOneID(empid).
		SetIDVerified(newEmployeeprofile.IDVerified).
		SetIDRemarks(newEmployeeprofile.IDRemarks).
		SetNameVerified(newEmployeeprofile.NameVerified).
		SetNameRemarks(newEmployeeprofile.NameRemarks).
		SetFathersNameVerified(newEmployeeprofile.FathersNameVerified).
		SetFathersNameRemarks(newEmployeeprofile.FathersNameRemarks).
		SetDOBVerified(newEmployeeprofile.DOBVerified).
		SetDOBRemarks(newEmployeeprofile.DOBRemarks).
		SetGenderVerified(newEmployeeprofile.GenderVerified).
		SetGenderRemarks(newEmployeeprofile.GenderRemarks).
		SetEmployeeCategoryCodeVerified(newEmployeeprofile.EmployeeCategoryCodeVerified).
		SetEmployeeCategoryCodeRemarks(newEmployeeprofile.EmployeeCategoryCodeRemarks).
		SetWithDisabilityVerified(newEmployeeprofile.WithDisabilityVerified).
		SetWithDisabilityRemarks(newEmployeeprofile.WithDisabilityRemarks).
		SetDisabilityTypeVerified(newEmployeeprofile.DisabilityTypeVerified).
		SetDisabilityTypeRemarks(newEmployeeprofile.DisabilityTypeRemarks).
		SetDisabilityPercentageVerified(newEmployeeprofile.DisabilityPercentageVerified).
		SetDisabilityPercentageRemarks(newEmployeeprofile.DisabilityPercentageRemarks).
		SetSignatureVerified(newEmployeeprofile.SignatureVerified).
		SetSignatureRemarks(newEmployeeprofile.SignatureRemarks).
		SetPhotoVerified(newEmployeeprofile.PhotoVerified).
		SetPhotoRemarks(newEmployeeprofile.PhotoRemarks).SetPostID(newEmployeeprofile.PostID).
		//SetPostCode(newEmployeeprofile.PostCode).
		//SetPostCodeVerified(newEmployeeprofile.PostCodeVerified).
		SetEmployeePost(newEmployeeprofile.EmployeePost).
		//SetPostCodeRemStatus(newEmployeeprofile.PostCodeRemStatus).
		//SetPostCodeRemarks(newEmployeeprofile.PostCodeRemarks).
		SetEmployeeDesignationVerified(newEmployeeprofile.EmployeeDesignationVerified).
		SetEmployeeDesignationRemarks(newEmployeeprofile.EmployeeDesignationRemarks).
		SetCircleVerified(newEmployeeprofile.CircleVerified).
		SetCircleRemarks(newEmployeeprofile.CircleRemarks).
		SetRegionVerified(newEmployeeprofile.RegionVerified).
		SetRegionRemarks(newEmployeeprofile.RegionRemarks).
		SetDivisionVerified(newEmployeeprofile.DivisionVerified).
		SetDivisionRemarks(newEmployeeprofile.DivisionRemarks).
		SetOfficeVerified(newEmployeeprofile.OfficeVerified).
		SetOfficeRemarks(newEmployeeprofile.OfficeRemarks).
		SetRoleVerified(newEmployeeprofile.RoleVerified).
		SetRoleRemarks(newEmployeeprofile.RoleRemarks).
		SetDCCSVerified(newEmployeeprofile.DCCSVerified).
		SetDCCSRemarks(newEmployeeprofile.DCCSRemarks).
		SetDCInPresentCadreVerified(newEmployeeprofile.DCInPresentCadreVerified).
		SetDCInPresentCadreRemarks(newEmployeeprofile.DCInPresentCadreRemarks).
		SetAPSWorkingVerified(newEmployeeprofile.APSWorkingVerified).
		SetAPSWorkingRemarks(newEmployeeprofile.APSWorkingRemarks).
		SetProfilestatus(newEmployeeprofile.Profilestatus).
		Save(context.Background())

	//.Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedEmployee, nil
}

/*
func QueryEmployeeMasterByEmpID(ctx context.Context, client *ent.Client, empID int32) ([]*ent.Employees, error) {
	//Can use GetX as well

	emps, err := client.Employees.Query().
		Where(employees.EmployeedIDEQ(empID)).
		All(ctx)

	if err != nil {
		log.Println("error at gettting employees from employee master: ", err)
		return nil, fmt.Errorf("failed at employee  master: %w", err)
	}
	log.Println("Employee details returned by Employee master : ", emps)
	return emps, nil
}*/

// This is added fro Querying EmpID from Employee Master
func QueryEmployeeMasterByEmpID(ctx context.Context, client *ent.Client, empID int64) (*ent.EmployeeMaster, int32, string, bool, error) {
	//Can use GetX as well

	emps, err := client.EmployeeMaster.
		Query().
		Where(employeemaster.EmployeeIDEQ(empID),
			employeemaster.StatussEQ("active"),
			employeemaster.VerifyStatusEQ(true)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, errors.New("no employee exists or in inactive state ")
		} else {
			return nil, 500, " -STR002", false, err
		}
	}
	return emps, 200, "", true, nil
}
