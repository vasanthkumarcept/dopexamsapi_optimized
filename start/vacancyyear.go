package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/util"
)

func CreateVacancyYears(client *ent.Client, newVacYr *ent.VacancyYear) (*ent.VacancyYear, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	u, err := client.VacancyYear.Create().
		SetVacancyYear(newVacYr.VacancyYear).
		SetFromDate(newVacYr.FromDate).
		SetToDate(newVacYr.ToDate).
		SetNotifyCode(newVacYr.NotifyCode).
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Vacancy Years: ", newVacYr)
		return nil, fmt.Errorf("failed creating Vacancy Years: %w", err)
	}
	log.Println("Vacancy Year was added to Vacancy years: ", u)
	return u, nil
}

func QueryVacancyYearID(ctx context.Context, client *ent.Client, id int32) (*ent.VacancyYear, error) {
	//Can use GetX as well

	newVacYr, err := client.VacancyYear.Get(ctx, id)
	if err != nil {
		log.Println("error at getting Vacancy Years: ", err)
		return nil, fmt.Errorf("failed querying Vacancy Years: %w", err)
	}
	log.Println("newVacYr returned: ", newVacYr)
	return newVacYr, nil
}

func QueryVacancyYear(ctx context.Context, client *ent.Client) ([]*ent.VacancyYear, error) {
	//Array of exams
	newVacYr, err := client.VacancyYear.Query().All(ctx)
	if err != nil {
		log.Println("error at vacancy Year ID: ", err)
		return nil, fmt.Errorf("failed querying Vacancy Years: %w", err)
	}
	log.Println("Vacancy Year data returned: ", newVacYr)
	return newVacYr, nil
}

func DeleteVacancyYearID(client *ent.Client, id int32) error {

	//context not passed for delete dont know why?
	err := client.VacancyYear.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func UpdateVacancyYearID(client *ent.Client, id int32, newVacYr *ent.VacancyYear) (*ent.VacancyYear, error) {

	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
	_, err := QueryVacancyYearID(ctx, client, id)
	if err != nil {
		return nil, err
	}
	UpdateVY, err := client.VacancyYear.UpdateOneID(id).
		SetVacancyYear(newVacYr.VacancyYear).
		SetFromDate(newVacYr.FromDate).
		SetToDate(newVacYr.ToDate).
		SetNotifyCode(newVacYr.NotifyCode).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return UpdateVY, nil
}
