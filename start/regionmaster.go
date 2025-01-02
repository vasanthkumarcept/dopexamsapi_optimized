package start

import (
	//"context"
	//"fmt"
	//"log"
	//"recruit/ent"
	//"recruit/ent/divisionmaster"
	//"recruit/ent/regionmaster"
)

/* func CreateRegionMaster(client *ent.Client, newRegion *ent.RegionMaster) (*ent.RegionMaster, error) {
	//fmt.Println("Num of Papers: ", newExam.NumOfPapers)

	ctx := context.Background()
	u, err := client.RegionMaster.
		Create().
		SetRegionCode(newRegion.RegionCode).
		SetRegionOfficeId(newRegion.RegionOfficeId).
		SetRegionOfficeName(newRegion.RegionOfficeName).
		SetOfficeType(newRegion.OfficeType).
		SetReportingOfficeType(newRegion.ReportingOfficeType).
		SetReportingOfficeCode(newRegion.ReportingOfficeCode).
		SetEmailID(newRegion.EmailID).
		SetMobileNumber(newRegion.MobileNumber).
		SetCircleCode(newRegion.CircleCode).
		//.AddCircleRefIDs().
		Save(ctx)
	if err != nil {
		log.Println("error at Creating Region Master: ", newRegion)
		return nil, fmt.Errorf("failed creating Region Master: %w", err)
	}
	//log.Println("Region was created: ", u)

	return u, nil
}

func QueryRegionMaster(ctx context.Context, client *ent.Client) ([]*ent.RegionMaster, error) {
	//Array of exams
	regions, err := client.RegionMaster.Query().
		All(ctx)
	if err != nil {
		log.Println("error at Circle Master: ", err)
		return nil, fmt.Errorf("failed querying Region Master: %w", err)
	}
	//log.Println("Region returned: ", regions)
	return regions, nil
}

func QueryRegionMasterByCircleCode(ctx context.Context, client *ent.Client, Ccode int32) ([]*ent.RegionMaster, error) {
	//Can use GetX as well

	Region_Master, err := client.RegionMaster.Query().
		Where(regionmaster.CircleCodeEQ(Ccode)).
		All(ctx)

	if err != nil {
		log.Println("error at gettting region master: ", err)
		return nil, fmt.Errorf("failed region master: %w", err)
	}
	//log.Println("Region returned by Circlecode : ", Region_Master)
	return Region_Master, nil
}
func QueryRegionMasterByID(ctx context.Context, client *ent.Client, id int32) (*ent.RegionMaster, error) {
	//Can use GetX as well

	Region_Master, err := client.RegionMaster.Get(ctx, id)
	if err != nil {

		log.Println("error at getting RegionMaster ID: ", err)
		return nil, fmt.Errorf("failed querying RegionMaster: %w", err)
	}
	//log.Println("RegionMaster details returned: ", Region_Master)
	return Region_Master, nil
}

func QueryRegionHeadQuartersByExamConductedBy(client *ent.Client, regionlists *ent.RegionMaster, naid int32) ([]*ent.RegionMaster, error) {
	ctx := context.Background()
	nid := naid

	regions, err := client.RegionMaster.Query().Where(regionmaster.CircleCodeEQ(nid)).
		All(ctx)
	if err != nil {
		log.Println("error at Region Master: ", err)
		return nil, fmt.Errorf("failed querying Region Master: %w", err)
	}
	//log.Println("Regions returned: ", Regions)
	return regions, nil
}
*/
/* func QueryDivisionHeadQuartersByExamConductedBy(client *ent.Client, divisionlists *ent.DivisionMaster, regionid int32) ([]*ent.DivisionMaster, error) {
	ctx := context.Background()
	rid := regionid

	divisions, err := client.DivisionMaster.Query().Where(divisionmaster.RegionCodeEQ(rid)).
		All(ctx)
	if err != nil {
		log.Println("error at Division Master: ", err)
		return nil, fmt.Errorf("failed querying Division Master: %w", err)
	}
	//log.Println("Divisions returned: ", divisions)
	return divisions, nil
}
*/
