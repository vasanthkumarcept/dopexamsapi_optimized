package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"recruit/ent"
	"recruit/util"

	"recruit/start"
)

// QueryOfficeByPincodehandle godoc
// @Summary Get offices by pincode
// @Description Get details of offices based on the given pincode
// @Tags Facilities
// @Accept json
// @Produce json
// @Param pincode path int true "Pincode"
// @Success 200 {object} start.EmployeeMasterResponse "Office retrieved based on pincode successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/facilities/byPincode/{pincode} [get]
func QueryOfficeByPincodehandle(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		var startFunction string = " main - QueryOfficeByPincodehandle - start - QueryOfficeByPincode "
		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("pincode")
		PinID, _ := strconv.ParseInt(id, 10, 32)
		offices, status, stgError, dataStatus, err := start.QueryOfficeByPincode(ctx, client, int32(PinID))
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       offices,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetFacilityDetailsByFacilityOfficeID godoc
// @Summary Get facility details by office ID
// @Description Get details of a facility based on the given office ID
// @Tags Facilities
// @Accept json
// @Produce json
// @Param workingofficefacilityid path string true "Office ID"
// @Success 200 {object} start.EmployeeMasterResponse "Facily details extracted"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/facilities/byfacilityofficeid/{workingofficefacilityid} [get]
func GetFacilityDetailsByFacilityOfficeID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		officeID := gctx.Param("workingofficefacilityid")
		var mainFunction string = " main - GetFacilityDetailsByFacilityOfficeID "
		var startFunction string = " - start - QueryFacilityByOfficeID "

		facility, status, stgError, dataStatus, err := start.QueryFacilityByOfficeID(ctx, client, officeID)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Facily details extracted",
			"data":       facility,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// QueryCircleHeadQuartersByExamConductedBy godoc
// @Summary      Query Circle HeadQuarters By Exam Conducted By
// @Description  Get Circle HeadQuarters details based on Exam Conducted By
// @Tags Facilities
// @Accept       json
// @Produce      json
// @Param        id2    path      string  true  "ID of the Exam Conducted By"
// @Success 200 {object} start.EmployeeMasterResponse "CircleHeadQuarters details extracted"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /deptexam/CircleHeadQuarters/{id2} [get]
func QueryCircleHeadQuartersByExamConductedBy(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//cb := gctx.Param("id")
		naid := gctx.Param("id2")
		// circlelist := new(ent.FacilityMasters)
		var mainFunction string = " main - QueryCircleHeadQuartersByExamConductedBy "
		var startFunction string = " - start - QueryCircleHeadQuartersByExamConductedBy "

		circlelists, status, stgError, dataStatus, err := start.QueryCircleHeadQuartersByExamConductedBy(client, naid)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		//gctx.JSON(http.StatusOK, gin.H{"data": circlelists})
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "CircleHeadQuarters details extracted",
			"data":       circlelists,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// QueryCircleHeadQuarters godoc
// @Summary      Query Circle HeadQuarters
// @Description  Get Circle HeadQuarters details
// @Tags Facilities
// @Accept       json
// @Produce      json
// @Success 200 {object} start.EmployeeMasterResponse "CircleHeadQuarters details extracted"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /deptexam/CircleHeadQuarters/Directorate [get]
func QueryCircleHeadQuarters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		var mainFunction string = " main - QueryCircleHeadQuarters "
		var startFunction string = " - start - SubQueryCircleHeadQuarters "

		circlelists, status, stgError, dataStatus, err := start.SubQueryCircleHeadQuarters(client)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		//gctx.JSON(http.StatusOK, gin.H{"data": circlelists})
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "CircleHeadQuarters details extracted",
			"data":       circlelists,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// QueryRegionHeadQuartersByExamConductedBy godoc
// @Summary      Query Region HeadQuarters By Exam Conducted By
// @Description  Get Region HeadQuarters details based on Exam Conducted By
// @Tags         Facilities
// @Accept       json
// @Produce      json
// @Param        id2   path      string  true  "ID of the Exam Conducted By"
// @Success 200 {object} start.EmployeeMasterResponse
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /deptexam/RegionHeadQuarters/{id2} [get]
func QueryRegionHeadQuartersByExamConductedBy(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//cb := gctx.Param("id1")
		naid := gctx.Param("id2")
		var mainFunction string = " main - QueryCircleHeadQuartersByExamConductedBy "
		var startFunction string = " - start - QueryCircleHeadQuartersByExamConductedBy "

		regionlists, status, stgError, dataStatus, err := start.StartQueryRegionHeadQuartersByExamConductedBy(client, naid)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       regionlists,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetFacilitiesByCircleOfficeId godoc
// @Summary      Get Facilities By Circle Office ID
// @Description  Retrieve facilities by Circle Office ID
// @Tags         Facilities
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Circle Office ID"
// @Success 200 {object} start.EmployeeMasterResponse
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router       /deptexam/GetFacilitiesByCircleOfficeId/{id} [get]
func GetFacilitiesByCircleOfficeId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")

		var mainFunction string = " main - GetFacilitiesByCircleOfficeId "
		var startFunction string = " - start - QueryNewFacilityMasterByCircleId "
		//var examID int32
		//FacID, _ := strconv.ParseInt(id, 10, 32)
		facs, status, stgError, dataStatus, err := start.QueryNewFacilityMasterByCircleId(ctx, client, id)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "CircleHeadQuarters details extracted",
			"data":       facs,
			"dataexists": dataStatus,
		})
		//gctx.JSON(http.StatusOK, gin.H{"data": facs})
	}
	return gin.HandlerFunc(fn)
}

// GetFacilitiesByReportingOfficeId godoc
// @Summary Get facilities by reporting office ID
// @Description Get the list of facilities associated with a specific reporting office ID
// @Tags Facilities
// @Accept  json
// @Produce  json
// @Param facilityid path string true "Facility ID"
// @Success 200 {object} start.EmployeeMasterResponse "Facilities retrieved successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/GetFacilitiesByReportingOfficeId/{facilityid} [get]
func GetFacilitiesByReportingOfficeId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var startFunction string = " main - GetFacilitiesByReportingOfficeId - start - QueryNewFacilityMasterByReportingId "
		facilityID := gctx.Param("facilityid")
		reportingOffices, status, stgError, dataStatus, err := start.QueryNewFacilityMasterByReportingId(ctx, client, facilityID)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       reportingOffices,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// // facility new ---vasanth
func CreateFacilitynewMasters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newfac := &ent.FacilityMasters{}
		if err := gctx.ShouldBindJSON(newfac); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		facility, err := start.CreatenewFacility(client, newfac)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"message": "Successfully created the Facility Master", "facility": facility})
	}
	return gin.HandlerFunc(fn)
}
