package start

import (
	"context"
	"errors"

	//"errors"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/usermaster"
	caa_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"
)

func CreateCandidate(ctx context.Context, client *ent.Client, candidateRequest caa_reg.CandidateCreation) (*ent.EmployeeMaster, int32, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, 500, fmt.Errorf("failed to start transaction: %w", err)
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
	admin, err := tx.EmployeeMaster.
		Create().
		SetEmployeeID(candidateRequest.EmployeeId).
		SetEmployeeName(candidateRequest.EmployeeName).
		SetFacilityID(candidateRequest.FacilityID).
		SetMobileNumber(candidateRequest.Mobile).
		SetEmailID(candidateRequest.EmailID).
		SetCreatedby(candidateRequest.CreatedBy).
		SetUidToken(candidateRequest.UidToken).
		Save(ctx)
	if err != nil {
		log.Printf("Error creating Candidate: %v\n", err)
		return nil, 400, fmt.Errorf("failed creating Candidate: %w", err)
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, 500, err
	}
	log.Printf("Candidate was created: %v\n", admin)
	return admin, 200, nil
}
func QueryCandidateUsersByEmpId(ctx context.Context, client *ent.Client, empid int64) (*ent.UserMaster, int32, string, bool, error) {
	//Can use GetX as well

	user, err := client.UserMaster.Query().
		Where(usermaster.EmployeeIDEQ(empid),
			usermaster.StatussEQ("active"),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR001", false, errors.New("no such employee id found")
		} else {
			return nil, 500, " -STR002", false, err
		}
	}
	return user, 200, "", true, nil
}
func UpdateCandidateUser(client *ent.Client, id int64, newUser *ent.UserMaster) (*ent.UserMaster, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()

	// Query the user by EmployeeID
	user, err := client.UserMaster.Query().
		Where(usermaster.EmployeeID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// Update the user entity with the provided new user data
	user, err = user.Update().
		SetEmployeeID(newUser.EmployeeID).
		SetEmployeeName(newUser.EmployeeName).
		SetFacilityID(newUser.FacilityID).
		SetGender(newUser.Gender).
		SetMobile(newUser.Mobile).
		SetEmailID(newUser.EmailID).
		SetDOB(newUser.DOB).
		SetCreatedBy(newUser.CreatedBy).
		SetUidToken(newUser.UidToken).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
