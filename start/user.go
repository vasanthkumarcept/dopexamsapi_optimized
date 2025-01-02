package start

//"recruit/ent/user"

/*func QueryExamByName(ctx context.Context, client *ent.Client, examname string) (*ent.Exam, error) {
	u, err := client.Exam.Query().Where(exam.ExamName(examname)).Only(ctx)

	if err != nil {
		log.Println("error at gettting examname: ", err)
		return nil, fmt.Errorf("failed querying exam: %w", err)
	}
	log.Println("exam returned by name: ", u)
	return u, nil
}*/

/* func QueryUserID(ctx context.Context, client *ent.Client, id string) (*ent.User, error) {
	//Need to conver it to ID

	u, err := client.User.Query().Where(user.EmployeedID(id)).Only(ctx)

	//u, err := client.User.Get(ctx, id)
	if err != nil {
		log.Println("error at getting UserID: ", err)
		return nil, fmt.Errorf("failed querying UserID: %w", err)
	}
	log.Println("NodalOfficer returned: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	//Needs to provide based on divisional id
	u, err := client.User.Query().All(ctx)
	if err != nil {
		log.Println("error at User: ", err)
		return nil, fmt.Errorf("failed querying User: %w", err)
	}
	log.Println("NodalOfficer returned: ", u)
	return u, nil
} */

// func QueryRoles(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
// 	//Needs to provide based on divisional id
// 	u, err := client.User.Query().All(ctx)
// 	if err != nil {
// 		log.Println("error at User: ", err)
// 		return nil, fmt.Errorf("failed querying User: %w", err)
// 	}
// 	log.Println("NodalOfficer returned: ", u)
// 	return u, nil
// }

/* func DeleteUserID(ctx context.Context, client *ent.Client, id string) error {

	//Need to change it ID instead of string...
	empid, err1 := client.User.Query().Where(user.EmployeedID(id)).Only(ctx)

	if err1 != nil {
		return err1
	}
	//context not passed for delete dont know why?
	err := client.NodalOfficer.DeleteOneID(int32(empid.ID)).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}
*/
/*
	 func UpdateUser(client *ent.Client, id string, newUser *ent.User) (*ent.User, error) {

		ctx := context.Background()
		//This needs to be optimised
		empid, err := client.User.Query().Where(user.EmployeedID(id)).Only(ctx)
		//	_, err := QueryNodalOfficerID(ctx, client, emplo)
		if err != nil {
			return nil, err
		}
		u, err := client.User.UpdateOneID(empid.ID).
			SetEmployeedID(newUser.EmployeedID).
			SetIDVerified(newUser.IDVerified).
			SetIDRemStatus(newUser.IDRemStatus).
			SetIDRemarks(newUser.IDRemarks).
			SetEmployeedName(newUser.EmployeedName).
			SetNameVerified(newUser.NameVerified).
			SetNameRemStatus(newUser.NameRemStatus).
			SetNameRemarks(newUser.NameRemarks).
			SetDOB(newUser.DOB).
			SetDOBVerified(newUser.DOBVerified).
			SetDOBRemStatus(newUser.DOBRemStatus).
			SetDOBRemarks(newUser.DOBRemarks).
			SetGender(newUser.Gender).
			SetGenderVerified(newUser.GenderVerified).
			SetGenderRemStatus(newUser.GenderRemStatus).
			SetGenderRemarks(newUser.GenderRemarks).
			SetCadreid(newUser.Cadreid).
			SetCadreidVerified(newUser.CadreidVerified).
			SetCadreidRemStatus(newUser.CadreidRemStatus).
			SetCadreidRemarks(newUser.CadreidRemarks).
			SetOfficeID(newUser.OfficeID).
			SetOfficeIDVerified(newUser.OfficeIDVerified).
			SetOfficeIDRemarks(newUser.OfficeIDRemarks).
			SetOfficeIDRemStatus(newUser.OfficeIDRemStatus).
			SetPH(newUser.PH).
			SetPHRemStatus(newUser.PHRemStatus).
			SetPHVerified(newUser.PHVerified).
			SetPHRemarks(newUser.PHRemarks).
			SetPHDetails(newUser.PHDetails).
			SetPHDetailsVerified(newUser.PHDetailsVerified).
			SetPHDetailsRemStatus(newUser.PHDetailsRemStatus).
			SetPHDetailsRemarks(newUser.PHDetailsRemarks).
			SetAPSWorking(newUser.APSWorking).
			SetAPSWorkingVerified(newUser.APSWorkingVerified).
			SetAPSWorkingRemarks(newUser.APSWorkingRemarks).
			SetAPSWorkingRemStatus(newUser.APSWorkingRemStatus).
			SetProfilestatus(newUser.Profilestatus).
			Save(context.Background())
		if err != nil {
			return nil, err
		}
		return u, nil
	}
*/
/* func UpdateUser1(client *ent.Client, id string, newUser *ent.User) (*ent.User, error) {

	//As of now both updates are same

	ctx := context.Background()

	empid, err := client.User.Query().Where(user.EmployeedID(id)).Only(ctx)
	//_, err := QueryNodalOfficerID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	u, err := client.User.UpdateOneID(empid.ID).
		SetEmployeedID(newUser.EmployeedID).
		SetIDVerified(newUser.IDVerified).
		SetIDRemStatus(newUser.IDRemStatus).
		SetIDRemarks(newUser.IDRemarks).
		SetEmployeedName(newUser.EmployeedName).
		SetNameVerified(newUser.NameVerified).
		SetNameRemStatus(newUser.NameRemStatus).
		SetNameRemarks(newUser.NameRemarks).
		SetDOB(newUser.DOB).
		SetDOBVerified(newUser.DOBVerified).
		SetDOBRemStatus(newUser.DOBRemStatus).
		SetDOBRemarks(newUser.DOBRemarks).
		SetGender(newUser.Gender).
		SetGenderVerified(newUser.GenderVerified).
		SetGenderRemStatus(newUser.GenderRemStatus).
		SetGenderRemarks(newUser.GenderRemarks).
		SetCadreid(newUser.Cadreid).
		SetCadreidVerified(newUser.CadreidVerified).
		SetCadreidRemStatus(newUser.CadreidRemStatus).
		SetCadreidRemarks(newUser.CadreidRemarks).
		SetOfficeID(newUser.OfficeID).
		SetOfficeIDVerified(newUser.OfficeIDVerified).
		SetOfficeIDRemarks(newUser.OfficeIDRemarks).
		SetOfficeIDRemStatus(newUser.OfficeIDRemStatus).
		SetPH(newUser.PH).
		SetPHRemStatus(newUser.PHRemStatus).
		SetPHVerified(newUser.PHVerified).
		SetPHRemarks(newUser.PHRemarks).
		SetPHDetails(newUser.PHDetails).
		SetPHDetailsVerified(newUser.PHDetailsVerified).
		SetPHDetailsRemStatus(newUser.PHDetailsRemStatus).
		SetPHDetailsRemarks(newUser.PHDetailsRemarks).
		SetAPSWorking(newUser.APSWorking).
		SetAPSWorkingVerified(newUser.APSWorkingVerified).
		SetAPSWorkingRemarks(newUser.APSWorkingRemarks).
		SetAPSWorkingRemStatus(newUser.APSWorkingRemStatus).
		SetProfilestatus(newUser.Profilestatus).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return u, nil
} */
