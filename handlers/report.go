package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"recruit/ent"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/facilitymasters"
	"recruit/start"
	"recruit/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetExamDetailsWithExamCodeExamYearCircleFacilityID(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamDetailsWithExamCodeExamYearCircleFacilityID - start.GetExamDetailsWithExamCode "
		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		circleFacilityID := gctx.Param("cirfaclilityid")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, status, stgError, dataStatus, err := start.GetExamDetailsWithExamCode(ctx, client, int32(examCode), examYear, circleFacilityID)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		// response := gin.H{}
		// if ipApplications != nil {
		// 	response["ipApplications"] = ipApplications
		// }
		// if psApplications != nil {
		// 	response["psApplications"] = psApplications
		// }
		// if gdspaApplications != nil {
		// 	response["gdspaApplications"] = gdspaApplications
		// }
		// if pmpaApplications != nil {
		// 	response["pmpaApplications"] = pmpaApplications
		// }
		// if gdspmApplications != nil {
		// 	response["gdspmApplications"] = gdspmApplications
		// }
		// if mtspmApplications != nil {
		// 	response["mtspmApplications"] = mtspmApplications
		// }
		// if len(response) == 0 {
		// 	gctx.JSON(http.StatusOK, gin.H{"message": "No applications found for this inputs"})
		// } else {
		// 	gctx.JSON(http.StatusOK, response)
		// }

		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if gdspaApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if pmpaApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if gdspmApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if mtspmApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": dataStatus})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": dataStatus})
		}
	}
}

func GetExamDetailsWithExamCodeExamYearCA(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamDetailsWithExamCodeExamYearCA - start - SubGetExamDetailsWithExamCodeExamYearCA "
		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		CAFacilityID := gctx.Param("cafacilityid")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, status, stgError, dataStatus, err := start.SubGetExamDetailsWithExamCodeExamYearCA(ctx, client, int32(examCode), examYear, CAFacilityID)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		// response := gin.H{}
		// if ipApplications != nil {
		// 	response["ipApplications"] = ipApplications
		// }
		// if psApplications != nil {
		// 	response["psApplications"] = psApplications
		// }
		// if gdspaApplications != nil {
		// 	response["gdspaApplications"] = gdspaApplications
		// }
		// if pmpaApplications != nil {
		// 	response["pmpaApplications"] = pmpaApplications
		// }
		// if gdspmApplications != nil {
		// 	response["gdspmApplications"] = gdspmApplications
		// }
		// if mtspmApplications != nil {
		// 	response["mtspmApplications"] = mtspmApplications
		// }
		// if len(response) == 0 {
		// 	gctx.JSON(http.StatusOK, gin.H{"message": "No applications found for this inputs"})
		// } else {
		// 	gctx.JSON(http.StatusOK, response)
		// }

		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if gdspaApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if pmpaApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if gdspmApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if mtspmApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": dataStatus})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": dataStatus})
		}
	}
}

func GetExamDetailsWithExamCodeExamYearCircleFacilityIDBystatus(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamDetailsWithExamCodeExamYearCircleFacilityIDBystatus - start.GetExamDetailsWithExamCodeAndStatus "

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		circleFacilityID := gctx.Param("cirfaclilityid")
		recommendationStatus := gctx.Param("recommendationcode")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		statusCode, err := strconv.Atoi(recommendationStatus)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recommendation code"})
			return
		}
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, status, stgError, dataStatus, err := start.GetExamDetailsWithExamCodeAndStatusNew(ctx, client, int32(examCode), examYear, circleFacilityID, int32(statusCode))
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}

		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if gdspaApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if pmpaApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if gdspmApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if mtspmApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": dataStatus})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": dataStatus})
		}
	}
}
func GetExamDetailsWithExamCodeExamYearCircleFacilityIDByApplnstatus(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamDetailsWithExamCodeExamYearCircleFacilityIDByApplnstatus - start.GetExamDetailsWithExamCodeAndApplnStatus "

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		circleFacilityID := gctx.Param("cirfaclilityid")
		applnStatus := gctx.Param("applnStatuscode")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		statusCode, err := strconv.Atoi(applnStatus)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recommendation code"})
			return
		}
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, status, stgError, dataStatus, err := start.GetExamDetailsWithExamCodeAndApplnStatus(ctx, client, int32(examCode), examYear, circleFacilityID, int32(statusCode))
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}

		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if gdspaApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if pmpaApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if gdspmApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if mtspmApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": dataStatus})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": dataStatus})
		}
	}
}
func GetExamDetailsHallticketview(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamDetailsHallticketview - start.GetExamDetailsWithHallticktdet "
		//:examcode/:examyear/:cirfaclilityid/:divfaclilityid
		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		circleFacilityID := gctx.Param("cirfaclilityid")
		divisionFacilityID := gctx.Param("divfaclilityid")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, status, stgError, dataStatus, err := start.GetExamDetailsWithHallticktdet(ctx, client, int32(examCode), examYear, circleFacilityID, divisionFacilityID)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}

		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if gdspaApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if pmpaApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if gdspmApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if mtspmApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": dataStatus})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": dataStatus})
		}
	}
}

func GetExamDetailsHallticketviewDR(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamDetailsHallticketviewDR - start - GetExamDetailsWithHallticktdet "
		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, status, stgError, dataStatus, err := start.SubGetExamDetailsHallticketviewDR(ctx, client, int32(examCode), examYear)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}

		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if gdspaApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if pmpaApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if gdspmApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if mtspmApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": dataStatus})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": dataStatus})
		}
	}
}

func GetExamDetailsAttendanceView(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamDetailsAttendanceView - start - SubGetExamDetailsAttendanceView "
		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		circleFacilityID := gctx.Param("cirfaclilityid")
		divisionFacilityID := gctx.Param("divfaclilityid")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, status, stgError, dataStatus, err := start.SubGetExamDetailsAttendanceView(ctx, client, int32(examCode), examYear, circleFacilityID, divisionFacilityID)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}

		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if gdspaApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if pmpaApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if gdspmApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if mtspmApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": dataStatus})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": dataStatus})
		}
	}
}

func GetExamSummaryByRecommendedStatus(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		errRemarks := "500 error DB Connection error -DBS01"
		errStatus := start.HandleDBErrorInitial(gctx, client, errRemarks)
		if errStatus == "error" {
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = "main - GetExamDetailsWithExamCodeExamYearCircleFacilityIDBystatus - start.GetExamDetailsWithExamCodeAndStatus"

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		facilityID := gctx.Param("facilityid")
		requestType := gctx.Param("type") // "directorate" or "circle"

		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}

		if strings.HasPrefix(facilityID, "DR") {
			// Directorate
			switch requestType {
			case "directorate":
				circles, err := start.GetCirclesDetails(ctx, client, int32(examCode), examYear, facilityID)
				if err != nil {
					start.StartErrorHandlerWithoutLog(gctx, err, 500, err.Error(), client, startFunction)
					return
				}
				gctx.JSON(http.StatusOK, gin.H{"success": true, "data": circles})
			default:
				gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request type for directorate"})
			}
		} else if strings.HasPrefix(facilityID, "CR") {
			// Circle
			switch requestType {
			case "circle":
				circleDetails, err := start.GetCAFacilityDetails(ctx, client, int32(examCode), examYear, facilityID)
				if err != nil {
					start.StartErrorHandlerWithoutLog(gctx, err, 500, err.Error(), client, startFunction)
					return
				}
				gctx.JSON(http.StatusOK, gin.H{"success": true, "data": circleDetails})
			default:
				gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request type for circle"})
			}
		} else {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid facility ID"})
		}
	}
}

// nodalofficefacilityid
type ControllingOfficeDetails struct {
	FacilityID                    string `json:"facility_id"`
	FacilityName                  string `json:"facility_name"`
	RecommendedCount              int    `json:"recommended_count"`
	NotRecommendedCount           int    `json:"not_recommended_count"`
	ProvisionallyRecommendedCount int    `json:"provisionally_recommended_count"`
}

type CircleOfficeDetails struct {
	FacilityID         string                     `json:"facility_id"`
	FacilityName       string                     `json:"facility_name"`
	ControllingOffices []ControllingOfficeDetails `json:"controlling_offices"`
}
type Facility struct {
	NodalOfficerFacilityID           string `json:"nodal_officer_facility_id" ent:"field:nodal_officer_facility_id"`
	NodalOfficerFacilityName         string `json:"nodal_officer_facility_name" ent:"field:nodal_officer_facility_name"`
	ControllingAuthorityFacilityID   string `json:"controlling_authority_facility_id" ent:"field:controlling_authority_facility_id"`
	ControllingAuthorityFacilityName string `json:"controlling_authority_facility_name" ent:"field:controlling_authority_facility_name"`
}

func GetExamSummaryByRecommendedStatuss(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		nodalOfficerFacilityID := gctx.Param("facilityid") //Directorate/nodal officer facility id

		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}

		var facilities []Facility
		err = client.FacilityMasters.
			Query().
			Where(facilitymasters.NodalOfficerFacilityIDEQ(nodalOfficerFacilityID)).
			GroupBy(
				facilitymasters.FieldNodalOfficerFacilityID,
				facilitymasters.FieldNodalOfficerFacilityName,
				facilitymasters.FieldControllingAuthorityFacilityID,
				facilitymasters.FieldControllingAuthorityFacilityName,
			).
			Scan(ctx, &facilities)
		if err != nil {
			log.Printf("Failed to query facilities: %v", err)
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query facilities"})
			return
		}

		if len(facilities) == 0 {
			gctx.JSON(http.StatusNotFound, gin.H{"error": "No facilities found"})
			return
		}

		var controllingOffices []ControllingOfficeDetails
		for _, facility := range facilities {
			controllingOfficeID := facility.ControllingAuthorityFacilityID
			controllingOfficeName := facility.ControllingAuthorityFacilityName

			recommendedCount, notRecommendedCount, provisionallyRecommendedCount, err := start.GetCAApplicationCounts(ctx, client, int32(examCode), examYear, controllingOfficeID)
			if err != nil {
				log.Printf("Failed to get application counts for facility ID %s: %v", controllingOfficeID, err)
				gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get application counts for facility ID %s: %v", controllingOfficeID, err)})
				return
			}

			controllingOffices = append(controllingOffices, ControllingOfficeDetails{
				FacilityID:                    controllingOfficeID,
				FacilityName:                  controllingOfficeName,
				RecommendedCount:              recommendedCount,
				NotRecommendedCount:           notRecommendedCount,
				ProvisionallyRecommendedCount: provisionallyRecommendedCount,
			})
		}

		circleOfficeDetails := CircleOfficeDetails{
			FacilityID:         nodalOfficerFacilityID,
			FacilityName:       facilities[0].NodalOfficerFacilityName, // assuming all rows have the same nodal officer facility name
			ControllingOffices: controllingOffices,
		}

		gctx.JSON(http.StatusOK, gin.H{"success": true, "data": circleOfficeDetails})
	}
}

func GetExamSummaryByRecommendedStatussForCAandDT(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamSummaryByRecommendedStatussForCAandDT - start.GetApplicationRecomenddedCountsSummaryForCAandDT "
		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		nodalOfficerFacilityID := gctx.Param("facilityid") // Directorate/nodal officer facility id
		entityType := gctx.Param("entitytype")             // New variable to differentiate between circle and directorate
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}

		//capture one variable to differentiate between circle and directorate ,this variable should be pass variable to sub function

		recommendedstatussummary, status, stgError, dataStatus, err := start.GetApplicationRecomenddedCountsSummaryForCAandDT(ctx, client, int32(examCode), examYear, nodalOfficerFacilityID, entityType)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": recommendedstatussummary, "dataexists": dataStatus})
	}

}
func GetPendingWithCaNaApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		circleFacilityID := gctx.Param("facilityid")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		daysPendingStr := gctx.Param("daysPending")
		daysPending, err := strconv.Atoi(daysPendingStr)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid days pending"})
			return
		}
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, err := start.GetPendingApplicationsWithDays(ctx, client, int32(examCode), examYear, circleFacilityID, daysPending)
		if err != nil {
			gctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if psApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if psApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if psApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if psApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": false})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": true})
		}
	}

	return gin.HandlerFunc(fn)
}
func GetEmployeeMasterPendingWithCA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetEmployeeMasterPendingWithCA "
		var startFunction string = " - start - SubGetEmployeeMasterPendingWithCA "

		circleFacilityID := gctx.Param("facilityid")

		summaries, status, stgError, _, err := start.SubGetEmployeeMasterPendingWithCA(ctx, client, circleFacilityID)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		if len(summaries) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": false})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": summaries, "dataexists": true})
		}
	}

	return gin.HandlerFunc(fn)
}

func GetEmployeeMasterPendingWithCADT(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetEmployeeMasterPendingWithCADT "
		var startFunction string = " - start - SubGetEmployeeMasterPendingWithCADT "

		summaries, status, stgError, dataStatus, err := start.SubGetEmployeeMasterPendingWithCADT(ctx, client)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       summaries,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

func GetEmployeeMasterPendingWithNADT(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetEmployeeMasterPendingWithNADT "
		var startFunction string = " - start - SubGetEmployeeMasterPendingWithNADT "
		summaries, status, stgError, _, err := start.SubGetEmployeeMasterPendingWithNADT(ctx, client)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		if len(summaries) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": false})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": summaries, "dataexists": true})
		}
	}

	return gin.HandlerFunc(fn)
}

func GetPendingWithCandidateApplications(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		circleFacilityID := gctx.Param("facilityid")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		daysPendingStr := gctx.Param("daysPending")
		daysPending, err := strconv.Atoi(daysPendingStr)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid days pending"})
			return
		}
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, err := start.GetCandidatePendingApplicationsWithDays(ctx, client, int32(examCode), examYear, circleFacilityID, daysPending)
		if err != nil {
			gctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if psApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if psApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if psApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if psApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": false})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": true})
		}
	}

	return gin.HandlerFunc(fn)
}

func GetPendingWithCandidateApplicationsWithCA(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetPendingWithCandidateApplicationsWithCA "
		var startFunction string = " - start - SubGetCandidatePendingApplicationsWithCA "

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		nodalFacilityID := gctx.Param("nodalfacilityid")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}

		summaries, status, stgError, _, err := start.SubGetCandidatePendingApplicationsWithCA(ctx, client, int32(examCode), examYear, nodalFacilityID)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		if len(summaries) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": false})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": summaries, "dataexists": true})
		}

	}

	return gin.HandlerFunc(fn)
}

func GetPendingWithCandidateApplicationsWithCADT(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetPendingWithCandidateApplicationsWithCADT "
		var startFunction string = " - start - SubGetCandidatePendingApplicationsWithCADT "

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}

		summaries, status, stgError, _, err := start.SubGetCandidatePendingApplicationsWithCADT(ctx, client, int32(examCode), examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		if len(summaries) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": false})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": summaries, "dataexists": true})
		}

	}

	return gin.HandlerFunc(fn)
}

func GetPendingWithCandidateApplicationsWithNADT(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetPendingWithCandidateApplicationsWithNADT "
		var startFunction string = " - start - SubGetCandidatePendingApplicationsWithNADT "

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}

		summaries, status, stgError, _, err := start.SubGetCandidatePendingApplicationsWithNADT(ctx, client, int32(examCode), examYear)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		if len(summaries) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": false})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": summaries, "dataexists": true})
		}

	}

	return gin.HandlerFunc(fn)
}

func GetExamCityCenters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var mainFunction string = " main - GetExamCityCenters "
		var startFunction string = " - start - SubGetExamCityCenters "

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		summaries, status, stgError, _, err := start.SubGetExamCityCenters(context.Background(), client, examYear, int32(examCode))
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		if len(summaries) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No exam city and Exam center mapping found", "data": nil, "dataexists": false})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": summaries, "dataexists": true})
		}
	}
	return gin.HandlerFunc(fn)
}

func GetExamCityDivisions(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var mainFunction string = " main - GetExamCityDivisions "
		var startFunction string = " - start - SubGetExamCityDivisions "
		//:examcode/:examyear/:nofacilityid
		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		naFacilityId := gctx.Param("nofacilityid")
		summaries, status, stgError, _, err := start.SubGetExamCityDivisions(context.Background(), client, examYear, int32(examCode), naFacilityId)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		if len(summaries) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No exam city and Exam center mapping found", "data": nil, "dataexists": false})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": summaries, "dataexists": true})
		}
	}
	return gin.HandlerFunc(fn)
}

func Gettingresultbasedondates(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		summaries, err := start.GetSummary(context.Background(), client)
		if err != nil {
			gctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(200, summaries)

	}

	return gin.HandlerFunc(fn)
}

func GetExamSummaryByApplicationsStatussForCAandDT(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamSummaryByApplicationsStatussForCAandDT - start.GetApplicationApplicationStatusCountsSummaryForCAandDT "

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		nodalOfficerFacilityID := gctx.Param("facilityid") // Directorate/nodal officer facility id
		entityType := gctx.Param("entitytype")             // New variable to differentiate between circle and directorate
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}

		//capture one variable to differentiate between circle and directorate ,this variable should be pass variable to sub function

		applicationstatusSummarry, status, stgError, dataStatus, err := start.GetApplicationApplicationStatusCountsSummaryForCAandDT(ctx, client, int32(examCode), examYear, nodalOfficerFacilityID, entityType)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": applicationstatusSummarry, "dataexists": dataStatus})
	}

}

// type CircleOfficeDetails struct {
//     FacilityID         string                     `json:"facility_id"`
//     FacilityName       string                     `json:"facility_name"`
//     ControllingOffices []ControllingOfficeDetails `json:"controlling_offices"`
// }

type DirectorateDetails struct {
	DirectorateFacilityID   string                `json:"directorate_facility_id"`
	DirectorateFacilityName string                `json:"directorate_facility_name"`
	CircleOffices           []CircleOfficeDetails `json:"circle_offices"`
}

// func GetExamSummaryByDirectorate(client *ent.Client) gin.HandlerFunc {
// 	return func(gctx *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
// 		defer cancel()

// 		examCodeParam := gctx.Param("examcode")
// 		examYear := gctx.Param("examyear")
// 		directorateFacilityID := gctx.Param("facilityid")
// 		fmt.Println("facilityid", directorateFacilityID)

// 		examCode, err := strconv.Atoi(examCodeParam)
// 		if err != nil {
// 			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
// 			return
// 		}

// 		// Query to get all circles for the directorate
// 		var facilities []Facility
// 		err = client.FacilityMasters.
// 			Query().
// 			Where(facilitymasters.FacilityIDHasPrefix("DT")).
// 			Where(facilitymasters.FacilityID(directorateFacilityID)).
// 			Scan(ctx, &facilities)
// 		if err != nil {
// 			log.Printf("Failed to query facilities: %v", err)
// 			gctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query facilities"})
// 			return
// 		}

// 		if len(facilities) == 0 {
// 			gctx.JSON(http.StatusNotFound, gin.H{"error": "No facilities found"})
// 			return
// 		}

// 		circleMap := make(map[string]*CircleOfficeDetails)

// 		for _, facility := range facilities {
// 			circleFacilityID := facility.NodalOfficerFacilityID
// 			circleFacilityName := facility.NodalOfficerFacilityName

// 			recommendedCount, notRecommendedCount, provisionallyRecommendedCount, err := GetCAApplicationCounts(ctx, client, int32(examCode), examYear, circleFacilityID)
// 			if err != nil {
// 				log.Printf("Failed to get application counts for facility ID %s: %v", circleFacilityID, err)
// 				gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get application counts for facility ID %s: %v", circleFacilityID, err)})
// 				return
// 			}

// 			circleOffice, exists := circleMap[circleFacilityID]
// 			if !exists {
// 				circleOffice = &CircleOfficeDetails{
// 					FacilityID:         circleFacilityID,
// 					FacilityName:       circleFacilityName,
// 					ControllingOffices: []ControllingOfficeDetails{},
// 				}
// 				circleMap[circleFacilityID] = circleOffice
// 			}

// 			circleOffice.ControllingOffices = append(circleOffice.ControllingOffices, ControllingOfficeDetails{
// 				FacilityID:                    facility.ControllingAuthorityFacilityID,
// 				FacilityName:                  facility.ControllingAuthorityFacilityName,
// 				RecommendedCount:              recommendedCount,
// 				NotRecommendedCount:           notRecommendedCount,
// 				ProvisionallyRecommendedCount: provisionallyRecommendedCount,
// 			})
// 		}

// 		circleOffices := make([]CircleOfficeDetails, 0, len(circleMap))
// 		for _, circle := range circleMap {
// 			circleOffices = append(circleOffices, *circle)
// 		}

// 		directorateDetails := DirectorateDetails{
// 			DirectorateFacilityID:   directorateFacilityID,
// 			DirectorateFacilityName: facilities[0].NodalOfficerFacilityName, // assuming all rows have the same directorate facility name
// 			CircleOffices:           circleOffices,
// 		}

// 		gctx.JSON(http.StatusOK, gin.H{"success": true, "data": directorateDetails})
// 	}
// }

func GetExamSummaryByDirectorate(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := gctx.Request.Context()

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		directorateFacilityID := gctx.Param("facilityid")
		fmt.Println("facilityid", directorateFacilityID)

		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}

		// Query using ent's query builder to get circle facilities excluding specified ones
		facilities, err := getDistinctCircleFacilities(client)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("facilities are ", facilities)

		if len(facilities) == 0 {
			gctx.JSON(http.StatusNotFound, gin.H{"error": "No facilities found"})
			return
		}

		circleMap := make(map[string]*CircleOfficeDetails)

		for _, facility := range facilities {
			circleFacilityID := facility.CircleFacilityID
			circleFacilityName := facility.CircleFacilityName

			recommendedCount, notRecommendedCount, provisionallyRecommendedCount, err := GetCAApplicationCounts(ctx, client, int32(examCode), examYear, circleFacilityID)
			if err != nil {
				log.Printf("Failed to get application counts for facility ID %s: %v", circleFacilityID, err)
				gctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get application counts for facility ID %s: %v", circleFacilityID, err)})
				return
			}

			circleOffice, exists := circleMap[circleFacilityID]
			if !exists {
				circleOffice = &CircleOfficeDetails{
					FacilityID:         circleFacilityID,
					FacilityName:       circleFacilityName,
					ControllingOffices: []ControllingOfficeDetails{},
				}
				circleMap[circleFacilityID] = circleOffice
			}

			circleOffice.ControllingOffices = append(circleOffice.ControllingOffices, ControllingOfficeDetails{
				FacilityID:                    facility.ControllingAuthorityFacilityID,
				FacilityName:                  facility.ControllingAuthorityFacilityName,
				RecommendedCount:              recommendedCount,
				NotRecommendedCount:           notRecommendedCount,
				ProvisionallyRecommendedCount: provisionallyRecommendedCount,
			})
		}

		circleOffices := make([]CircleOfficeDetails, 0, len(circleMap))
		for _, circle := range circleMap {
			circleOffices = append(circleOffices, *circle)
		}

		directorateDetails := DirectorateDetails{
			DirectorateFacilityID:   directorateFacilityID,
			DirectorateFacilityName: facilities[0].NodalOfficerFacilityName, // assuming all rows have the same directorate facility name
			CircleOffices:           circleOffices,
		}

		gctx.JSON(http.StatusOK, gin.H{"success": true, "data": directorateDetails})
	}
}

func GetCAApplicationCounts(ctx context.Context, client *ent.Client, examCode int32, examYear string, nodalOfficeFacilityID string) (int, int, int, error) {
	log.Printf("Querying application counts for ExamCode: %d, ExamYear: %s, NodalOfficeFacilityID: %s", examCode, examYear, nodalOfficeFacilityID)
	fmt.Println("nodalOfficeFacilityID", nodalOfficeFacilityID)
	var recommendedCount, notRecommendedCount, provisionallyRecommendedCount int
	var err error

	switch examCode {
	case 2:
		recommendedCount, err = client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.RecommendedStatusEQ("Recommended"),
				exam_applications_ip.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Recommended count: %d", recommendedCount)

		notRecommendedCount, err = client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.RecommendedStatusEQ("Not Recommended"),
				exam_applications_ip.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying not recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Not recommended count: %d", notRecommendedCount)

		provisionallyRecommendedCount, err = client.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ip.RecommendedStatusEQ("Provisionally Recommended"),
				exam_applications_ip.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying provisionally recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Provisionally recommended count: %d", provisionallyRecommendedCount)

	case 1:
		recommendedCount, err = client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.RecommendedStatusEQ("Recommended"),
				exam_applications_ps.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Recommended count: %d", recommendedCount)

		notRecommendedCount, err = client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.RecommendedStatusEQ("Not Recommended"),
				exam_applications_ps.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying not recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Not recommended count: %d", notRecommendedCount)

		provisionallyRecommendedCount, err = client.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.NodalOfficeFacilityIDEQ(nodalOfficeFacilityID),
				exam_applications_ps.RecommendedStatusEQ("Provisionally Recommended"),
				exam_applications_ps.StatusEQ("active"),
			).
			Count(ctx)
		if err != nil {
			log.Printf("Error querying provisionally recommended count: %v", err)
			return 0, 0, 0, err
		}
		log.Printf("Provisionally recommended count: %d", provisionallyRecommendedCount)
	}

	return recommendedCount, notRecommendedCount, provisionallyRecommendedCount, nil
}

func GetAllApplications(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main -etAllApplications - start.GetExamDetailsWithExamCodeAndApplnStatus "

		examCodeParam := gctx.Param("examcode")
		examYear := gctx.Param("examyear")
		circleFacilityID := gctx.Param("cirfaclilityid")
		//applnStatus := gctx.Param("applnStatuscode")
		examCode, err := strconv.Atoi(examCodeParam)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam code"})
			return
		}
		//statusCode, err := strconv.Atoi(applnStatus)
		/* 		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recommendation code"})
			return
		} */
		ipApplications, psApplications, gdspaApplications, pmpaApplications, gdspmApplications, mtspmApplications, status, stgError, dataStatus, err := start.SubGetAllApplications(ctx, client, int32(examCode), examYear, circleFacilityID)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}

		data := gin.H{}
		if ipApplications != nil {
			data["ipApplications"] = ipApplications
		}
		if psApplications != nil {
			data["psApplications"] = psApplications
		}
		if gdspaApplications != nil {
			data["gdspaApplications"] = gdspaApplications
		}
		if pmpaApplications != nil {
			data["pmpaApplications"] = pmpaApplications
		}
		if gdspmApplications != nil {
			data["gdspmApplications"] = gdspmApplications
		}
		if mtspmApplications != nil {
			data["mtspmApplications"] = mtspmApplications
		}

		if len(data) == 0 {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "No applications found for this inputs", "data": nil, "dataexists": dataStatus})
		} else {
			gctx.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": data, "dataexists": dataStatus})
		}
	}
}
