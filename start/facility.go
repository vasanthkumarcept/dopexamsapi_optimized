package start

//"context"
//"fmt"
//"log"
//"recruit/ent"
//"recruit/ent/regionmaster"

type OfficeInfo struct {
	FacilityOfficeID int32  `json:"FacilityOfficeID"`
	Pincode          int32  `json:"Pincode"`
	FacilityName     string `json:"FacilityName"`
}

/* func QueryFacilityMasterByRegionCode(ctx context.Context, client *ent.Client, Rcode string) ([]*ent.RegionMaster, error) {
	//Can use GetX as well

	Facility_Master, err := client.RegionMaster.Query().
		Where(regionmaster.RegionOfficeIdEQ(Rcode)).
		All(ctx)

	if err != nil {
		log.Println("error at gettting Facility master: ", err)
		return nil, fmt.Errorf("failed Facility master: %w", err)
	}
	//log.Println("Facility returned by Region code : ", Facility_Master)
	return Facility_Master, nil
}
*/
