package start

import (
	"context"
	"errors"
	"fmt"
	"log"

	//"net/http"

	//"fmt"
	"time"

	//"log"
	"recruit/ent"
	"recruit/util"

	//"recruit/ent/circlemaster"
	//"recruit/ent/divisionmaster"
	"recruit/ent/facilitymasters"
	//"github.com/gin-gonic/gin"
	//"recruit/ent/regionmaster"
)

type OfficeInfoo struct {
	FacilityOfficeID int32  `json:"FacilityOfficeID"`
	Pincode          int32  `json:"Pincode"`
	FacilityName     string `json:"FacilityName"`
}

func CreatenewFacility(client *ent.Client, newfacility *ent.FacilityMasters) (*ent.FacilityMasters, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
	defer cancel()
	uniqNumber, err := generateUniqNumber(client)
	if err != nil {
		return nil, err
	}
	fmt.Println(uniqNumber)
	currentTime := time.Now().Truncate(time.Second)
	fmt.Println(currentTime)
	u, err := client.FacilityMasters.
		Create().
		SetFacilityID(newfacility.FacilityID).
		SetFacilityType(newfacility.FacilityType).
		SetFacilityIDDescription(newfacility.FacilityIDDescription).
		SetReportingOfficeFacilityID(newfacility.ReportingOfficeFacilityID).
		SetReportingOfficeFacilityName(newfacility.ReportingOfficeFacilityName).
		SetHOFacilityID(newfacility.HOFacilityID).
		SetHOFacilityName(newfacility.HOFacilityName).
		SetSubDivisionFacilityID(newfacility.SubDivisionFacilityID).
		SetSubDivisionFacilityName(newfacility.SubDivisionFacilityName).
		SetDivisionFacilityID(newfacility.DivisionFacilityID).
		SetDivisionFacilityName(newfacility.DivisionFacilityName).
		SetRegionFacilityID(newfacility.RegionFacilityID).
		SetRegionFacilityName(newfacility.RegionFacilityName).
		SetCircleFacilityID(newfacility.CircleFacilityID).
		SetCircleFacilityName(newfacility.CircleFacilityName).
		SetPincode(newfacility.Pincode).
		SetControllingAuthorityFacilityID(newfacility.ControllingAuthorityFacilityID).
		SetControllingAuthorityFacilityName(newfacility.ControllingAuthorityFacilityName).
		SetNodalOfficerFacilityID(newfacility.NodalOfficerFacilityID).
		SetNodalOfficerFacilityName(newfacility.NodalOfficerFacilityName).
		SetDeliveryNonDeliveryOffice(newfacility.DeliveryNonDeliveryOffice).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Facility: %w", err)
	}
	return u, nil
}

func generateUniqNumber(client *ent.Client) (int32, error) {
	ctx := context.TODO()
	lastNumber, err := client.FacilityMasters.
		Query().
		Order(ent.Desc(facilitymasters.FieldUUID)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// No existing numbers, start from 100001
			return 1000000001, nil
		}
		return 0, fmt.Errorf("failed to retrieve last application: %v", err)
	}
	return int32(lastNumber.UUID) + 1, nil
}

func QueryFacilitiesNewMaster(ctx context.Context, client *ent.Client) ([]*ent.FacilityMasters, error) {
	//Array of exams
	Facilities, err := client.FacilityMasters.Query().
		All(ctx)
	if err != nil {
		log.Println("error at Facility Master: ", err)
		return nil, fmt.Errorf("failed querying Facility Master: %w", err)
	}
	//log.Println("Facility returned: ", Facilities)
	return Facilities, nil
}
func QueryFacilityNewMasterByID(ctx context.Context, client *ent.Client, id int32) (*ent.FacilityMasters, error) {
	//Can use GetX as well

	Facility_Master, err := client.FacilityMasters.Get(ctx, id)
	if err != nil {

		log.Println("error at getting Facility ID: ", err)
		return nil, fmt.Errorf("failed querying Facility Master: %w", err)
	}
	log.Println("Facility Master details returned: ", Facility_Master)
	return Facility_Master, nil
}

func QueryNewFacilityMasterByReportingId(ctx context.Context, client *ent.Client, Rcode string) ([]*ent.FacilityMasters, int32, string, bool, error) {

	reportingOffices, err := client.FacilityMasters.Query().
		Where(facilitymasters.ReportingOfficeFacilityIDEQ(Rcode),
			facilitymasters.StatusEQ("active")).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR001", false, err
	}
	if len(reportingOffices) == 0 {
		return nil, 422, " -STR002", false, errors.New("no matching reporting Offices found")
	}
	return reportingOffices, 200, "", true, nil
}

func QueryNewFacilityMasterByCircleId(ctx context.Context, client *ent.Client, Rcode string) ([]*ent.FacilityMasters, int32, string, bool, error) {
	//Can use GetX as well

	Facility_Master, err := client.FacilityMasters.Query().
		Where(
			facilitymasters.CircleFacilityIDEQ(Rcode),
			facilitymasters.StatusEQ("active"),
			facilitymasters.Not(
				facilitymasters.FacilityTypeIn("BN", "CB", "DB", "EP", "EX", "HO", "HR", "IC", "LH", "LP", "MM", "MO", "PC", "PH", "PO", "RL", "RP", "SP", "SR", "TM", "BO"),
			),
		).
		All(ctx)

	if err != nil {
		log.Println("error at gettting Facility master: ", err)
		return nil, 422, " -STR001", false, fmt.Errorf("failed Facility master: %w", err)
	}
	//log.Println("Facility returned by Region code : ", Facility_Master)
	return Facility_Master, 200, "", true, nil
}

func QueryOfficeByPincode(ctx context.Context, client *ent.Client, id int32) ([]*ent.FacilityMasters, int32, string, bool, error) {
	pin := id

	if pin <= 0 {
		return nil, 422, " -STR001", false, errors.New("enter pincode with valid six digit number")
	}
	// Fetch all applications for the given facilityID
	pincode := fmt.Sprint(pin)
	officeInfos, err := client.FacilityMasters.
		Query().
		Where(facilitymasters.PincodeEQ(pincode),
			facilitymasters.StatusEQ("active")).
		All(ctx)
	if err != nil {
		return nil, 500, " -STR002", false, err
	}

	if len(officeInfos) > 0 {
		return officeInfos, 200, "", true, nil
	} else {
		return nil, 422, " -STR003", false, errors.New("no office matched with this PIN Code")
	}
}
func QueryFacilityByOfficeID(ctx context.Context, client *ent.Client, officeID string) (*ent.FacilityMasters, int32, string, bool, error) {
	if officeID == "" {
		return nil, 422, " -STR001", false, fmt.Errorf("officeID cannot be empty")
	}

	facility, err := client.FacilityMasters.
		Query().
		Where(facilitymasters.FacilityIDEQ(officeID),
			facilitymasters.StatusEQ("active"),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, 422, " -STR002", false, errors.New("no active facility details found ")
		} else {
			return nil, 500, " -STR003", false, err
		}
	}

	return facility, 200, "", true, nil
}
