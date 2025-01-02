package start

import (
	//"context"
	//"fmt"
	//"log"
	//"recruit/ent"
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/adminmaster"
	//"recruit/ent/exam_ip"
)

func QueryNodalOfficerByUsername(ctx context.Context, client *ent.Client, username string) (*ent.AdminMaster, error) {
	// Check if username is empty
	if username == "" {
		return nil, fmt.Errorf("username is mandatory")
	}

	// Can use GetX as well
	u, err := client.AdminMaster.Query().
		Where(adminmaster.UserName(username),
			adminmaster.StatussEQ("active")).
		Only(ctx)
	if err != nil {
		log.Println("error at getting NodalOfficer: ", err)
		return nil, fmt.Errorf("failed querying NodalOfficer: %w", err)
	}
	log.Println("NodalOfficer details returned: ", u)
	return u, nil
}

/* func CreateNodalOfficer(client *ent.Client, newNodalOfficer *ent.NodalOfficer) (*ent.NodalOfficer, error) {
	ctx := context.Background()
	u, err := client.NodalOfficer.Create().
		SetNodalOfficerName(newNodalOfficer.NodalOfficerName).
		SetDesignationID(newNodalOfficer.DesignationID).
		SetMobileNumber(newNodalOfficer.MobileNumber).
		SetNotifyCode(newNodalOfficer.NotifyCode).
		SetExamCode(newNodalOfficer.ExamCode).
		SetEmailID(newNodalOfficer.EmailID).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Application: ", newNodalOfficer)
		return nil, fmt.Errorf("failed creating Application: %w", err)
	}
	log.Println("Application was created: ", u)

	return u, nil
}
*/
/*func QueryExamByName(ctx context.Context, client *ent.Client, examname string) (*ent.Exam, error) {
	u, err := client.Exam.Query().Where(exam.ExamName(examname)).Only(ctx)

	if err != nil {
		log.Println("error at gettting examname: ", err)
		return nil, fmt.Errorf("failed querying exam: %w", err)
	}
	log.Println("exam returned by name: ", u)
	return u, nil
}*/

/* func QueryNodalOfficerID(ctx context.Context, client *ent.Client, id int32) (*ent.NodalOfficer, error) {
	//Can use GetX as well

	u, err := client.NodalOfficer.Get(ctx, id)
	if err != nil {
		log.Println("error at getting NodalOfficerID: ", err)
		return nil, fmt.Errorf("failed querying NodalOfficer: %w", err)
	}
	log.Println("NodalOfficer returned: ", u)
	return u, nil
}

func QueryNodalOfficer(ctx context.Context, client *ent.Client) ([]*ent.NodalOfficer, error) {
	//Array of exams
	u, err := client.NodalOfficer.Query().All(ctx)
	if err != nil {
		log.Println("error at NodalOfficer: ", err)
		return nil, fmt.Errorf("failed querying NodalOfficer: %w", err)
	}
	log.Println("NodalOfficer returned: ", u)
	return u, nil
}

func DeleteNodalOfficerID(client *ent.Client, id int32) error {

	//context not passed for delete dont know why?
	err := client.NodalOfficer.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func UpdateNodalOfficer(client *ent.Client, id int32, newNodalOfficer *ent.NodalOfficer) (*ent.NodalOfficer, error) {

	ctx := context.Background()
	_, err := QueryNodalOfficerID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	u, err := client.NodalOfficer.UpdateOneID(id).
		SetNodalOfficerName(newNodalOfficer.NodalOfficerName).
		SetDesignationID(newNodalOfficer.DesignationID).
		SetMobileNumber(newNodalOfficer.MobileNumber).
		SetNotifyCode(newNodalOfficer.NotifyCode).
		SetExamCode(newNodalOfficer.ExamCode).
		SetEmailID(newNodalOfficer.EmailID).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UpdateNodalOfficerHallTicketApproval(client *ent.Client, id int32, newNodalOfficer *ent.NodalOfficer) (*ent.NodalOfficer, error) {

	ctx := context.Background()
	_, err := QueryNodalOfficerID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	u, err := client.NodalOfficer.UpdateOneID(id).
		SetHallTicketApproved(newNodalOfficer.HallTicketApproved).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return u, nil
}



// Query Exams with Exam Code
func QueryExamIPByExamNameCode(ctx context.Context, client *ent.Client, examNameCode string) (*ent.Exam_IP, error) {
	// Check if examNameCode is empty
	if examNameCode == "" {
		return nil, fmt.Errorf("Please provide exam name code")
	}

	u, err := client.Exam_IP.Query().
		Where(exam_ip.ExamNameCode(examNameCode)).
		Only(ctx)
	if err != nil {
		log.Println("error at getting Exam_IP: ", err)
		return nil, fmt.Errorf("failed querying Exam_IP: %w", err)
	}
	log.Println("Exam_IP details returned: ", u)
	return u, nil
}
*/
//
