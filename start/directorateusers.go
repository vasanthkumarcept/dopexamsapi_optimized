package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/directorateusers"
	"recruit/util"
)

//    field.Time("HallTicketDownloadDate").Optional(), field.String("NotifyFile").Optional(), field.String("SyllabusFile").Optional(), field.String("VacanciesFile").Optional()}

func CreateDirectorateUsers(client *ent.Client, newDusers *ent.DirectorateUsers) (*ent.DirectorateUsers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.DirectorateUsers.Create().
		SetRole(newDusers.Role).
		SetEmployeedID(newDusers.EmployeedID).
		SetEmployeeName(newDusers.Role).
		SetMobileNumber(newDusers.MobileNumber).
		SetEmailId(newDusers.EmailId).
		SetSequenceNumber(newDusers.SequenceNumber).
		SetStatus(newDusers.Status).
		//SetExamCode(newDusers.ExamCode).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Directorate Users: ", newDusers)
		return nil, fmt.Errorf("failed creating Directorate Users: %w", err)
	}
	log.Println("Directorate Users was created: ", u)

	return u, nil
}

func QueryDUsersByID(ctx context.Context, client *ent.Client, id int32) (*ent.DirectorateUsers, error) {
	//Can use GetX as well

	Dusers, err := client.DirectorateUsers.Get(ctx, id)
	if err != nil {
		log.Println("error at getting Directorate Users ID: ", err)
		return nil, fmt.Errorf("failed querying Directorate Users: %w", err)
	}
	log.Println("Directorate Users details returned: ", Dusers)
	return Dusers, nil
}

func QueryeDirectorateUsers(ctx context.Context, client *ent.Client) ([]*ent.DirectorateUsers, error) {
	//Array of exams
	Dusers, err := client.DirectorateUsers.Query().All(ctx)
	if err != nil {
		log.Println("error at DirectorateUsers: ", err)
		return nil, fmt.Errorf("failed querying DirectorateUsers: %w", err)
	}
	log.Println("Notifications returned: ", Dusers)
	return Dusers, nil
}

func QueryeDirectorateUsersyEmpId(ctx context.Context, client *ent.Client, EmpID int32) ([]*ent.DirectorateUsers, error) {
	//Can use GetX as well
	Dusers, err := client.DirectorateUsers.Query().
		Select(directorateusers.FieldRole, directorateusers.FieldEmployeedID, directorateusers.FieldEmployeeName, directorateusers.FieldMobileNumber, directorateusers.FieldEmailId).
		Where(directorateusers.EmployeedIDEQ(EmpID)).
		All(ctx)

	if err != nil {
		log.Println("error at getting Directorate Users: ", err)
		return nil, fmt.Errorf("failed Directorate Users: %w", err)
	}
	log.Println("Directorate users returned: ", Dusers)

	if len(Dusers) == 0 {
		return nil, fmt.Errorf("no Directorate Users found with EmpID: %d", EmpID)
	}
	for _, users := range Dusers {
		fmt.Printf("ID: %d\n", users.EmployeedID)
		fmt.Printf("EmployeeName: %s\n", users.EmployeeName)
		fmt.Printf("Role: %s\n", users.Role)
		fmt.Printf("Mobile Number: %d\n", users.MobileNumber)
		fmt.Printf("Email ID: %s\n", users.EmailId)
	}
	return Dusers, nil
}
