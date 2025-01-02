package start

import (
	"context"
	"fmt"
	"log"
	"recruit/ent"
	"recruit/ent/facilitymasters"
	"recruit/util"
)

func QueryCircleHeadQuartersByExamConductedBy(client *ent.Client, nid string) ([]string, int32, string, bool, error) {
	fmt.Println("entering to start part: ")
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	circles, err := client.FacilityMasters.
		Query().
		Where(
			facilitymasters.FacilityTypeEQ("CR"),
			facilitymasters.CircleFacilityIDEQ(nid),
			facilitymasters.StatusEQ("active"),
		).
		GroupBy(facilitymasters.FieldCityName). // Replace FieldCityName with your actual column name
		Strings(ctx)

	//.All(ctx)
	fmt.Println("start querying Circle Master: ", circles)
	if err != nil {
		log.Println("error at Circle Master: ", err)
		return nil, 422, " -STR001", false, fmt.Errorf("failed querying Circle Master: %w", err)
	}
	return circles, 200, "", true, nil
}

func SubQueryCircleHeadQuarters(client *ent.Client) ([]*ent.FacilityMasters, int32, string, bool, error) {
	fmt.Println("entering to start part: ")
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	circles, err := client.FacilityMasters.
		Query().
		Where(
			facilitymasters.FacilityTypeEQ("CR"),
			facilitymasters.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		log.Println("error at Circle Master: ", err)
		return nil, 422, " -STR001", false, fmt.Errorf("failed querying Circle Master: %w", err)
	}
	return circles, 200, "", true, nil
}

func StartQueryRegionHeadQuartersByExamConductedBy(client *ent.Client, naid string) ([]string, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	nid := naid
	regions, err := client.FacilityMasters.Query().
		Where(
			facilitymasters.FacilityTypeEQ("DV"),
			facilitymasters.CircleFacilityIDEQ(nid),
			facilitymasters.StatusEQ("active"),
		).
		GroupBy(facilitymasters.FieldCityName). // Replace FieldRegion with your actual column name
		Strings(ctx)

	if err != nil {
		log.Println("error at Facility Master: ", err)
		return nil, 422, " -STR001", false, fmt.Errorf("failed querying Region Master: %w", err)
	}
	//log.Println("Regions returned: ", Regions)
	return regions, 200, "", true, nil
}
