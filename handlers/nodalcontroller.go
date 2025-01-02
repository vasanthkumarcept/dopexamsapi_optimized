package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"recruit/ent"

	"recruit/start"

	ca_reg "recruit/payloadstructure/candidate_registration"
	"recruit/util"

	_ "github.com/lib/pq"
)

// Create Exam Center ...

// CreateExamCenter godoc
// @Summary Create Exam Center
// @Description Create a new exam center
// @Tags Nodal
// @Accept  json
// @Produce  json
// @Param examCenter body ca_reg.CenterReq true "Exam Center Data"
// @Success 200 {object} CenterrResponse ""New Exam Centre Created successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/center/submit [post]
func CreateExamCenter(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var newCenter ca_reg.CenterReq
		//newCenter := new(ent.Center)
		// ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		// defer cancel()
		var mainFunction string = " main - CreateExamCenter "
		var startFunction string = " - start CreateCenter "
		if err := gctx.ShouldBindJSON(&newCenter); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newCenter.Edges.LogData[0]
		//newCenter,status, stgError, dataStatus, err := start.CreateCenter(client, newCenter)
		newCenterr, status, stgError, dataStatus, err := start.CreateCenter(client, &newCenter)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "New Exam Centre Created",
			"data": gin.H{
				"ExamCenterName":   newCenterr.ExamCenterName,
				"Address":          newCenterr.Address,
				"Pincode":          newCenterr.Pincode,
				"Landmark":         newCenterr.Landmark,
				"Center City name": newCenterr.CenterCityName,
			},
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// Fetch Exam Centre With Center ID
type CenterrResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	DataExists bool        `json:"dataexists"`
}

// GetCenterID godoc
// @Summary Get Center ID
// @Description Get details of a specific exam center by ID
// @Tags Nodal
// @Accept  json
// @Produce  json
// @Param id path int true "Center ID"
// @Success 200 {object} CenterrResponse "Active Exam fetched successfully by Exam year"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/center/get/{id} [get]
func GetCenterID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetCenterID "
		var startFunction string = " - start - QueryCenterID "

		id := gctx.Param("id")
		//var examID int32
		CenterID, _ := strconv.ParseInt(id, 10, 32)

		newCenter, status, stgError, dataStatus, err := start.QueryCenterID(ctx, client, int32(CenterID))
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{

			"success": true,
			"message": "",
			"data": gin.H{
				"ID":               newCenter.ID,
				"ExamCenterName":   newCenter.ExamCenterName,
				"Address":          newCenter.Address,
				"Pincode":          newCenter.Pincode,
				"Landmark":         newCenter.Landmark,
				"Center City name": newCenter.CenterCityName,
				"Seats Capacity":   newCenter.MaxSeats,
				"Seats Alloted":    newCenter.NoAlloted,
			},
			"dataexists": dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// GetExamCentresBynodalOfficeIDExamCode godoc
// @Summary Get Exam Centres by Nodal Office ID and Exam Code
// @Description Get exam centres based on Nodal Office ID, Exam Code, and City ID
// @Tags Nodal
// @Accept  json
// @Produce  json
// @Param id1 path int true "Exam Code"
// @Param id2 path string true "Nodal Office ID"
// @Param id3 path string true "Additional Parameter"
// @Param id4 path int true "City ID"
// @Success 200 {object} CenterrResponse "Exam Center fetched successfully by NodalOffice ID "
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/centers/{id1}/{id2}/{id3}/{id4} [get]
func GetExamCentresBynodalOfficeIDExamCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		ec := gctx.Param("id1")
		ExamCodeInt64, _ := strconv.ParseInt(ec, 10, 32)
		ExamCode := int32(ExamCodeInt64)
		NOOfficeID := gctx.Param("id2")
		id3 := gctx.Param("id3")
		id4 := gctx.Param("id4")
		Cityt64, _ := strconv.ParseInt(id4, 10, 32)
		Cityid := int32(Cityt64)
		var mainFunction string = " main - GetExamCentresBynodalOfficeIDExamCode "
		var startFunction string = " - start - GetExamCentersByExamCodeNOOfficeID "

		examcenters, status, stgError, dataStatus, err := start.GetExamCentersByExamCodeNOOfficeID(ctx, client, ExamCode, Cityid, NOOfficeID, id3)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		gctx.JSON(http.StatusOK, gin.H{

			"success":    true,
			"message":    "",
			"data":       examcenters,
			"dataexists": dataStatus,
		})

	}

	return gin.HandlerFunc(fn)
}

// func GetVersions(client *ent.Client) gin.HandlerFunc {
// 	return func(gctx *gin.Context) {
// 		ctx := gctx.Request.Context() // Use request context

// 		// Call QueryVersion function to fetch version data
// 		version, err := start.QueryVersion(ctx, client)
// 		if err != nil {
// 			gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

//			// Prepare JSON response with version data
//			gctx.JSON(http.StatusOK, gin.H{
//				"data": gin.H{
//					"ID":          version.ID,
//					"UI Version":  version.UiVersion,
//					"API Version": version.ApiVersion,
//				},
//			})
//		}
//	}
func GetVersions(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := gctx.Request.Context() // Use request context

		// Call QueryVersion function to fetch all version data
		versions, err := start.QueryVersion(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Prepare JSON response with version data
		var versionsData []gin.H
		for _, version := range versions {
			versionsData = append(versionsData, gin.H{
				"ID":          version.ID,
				"UI Version":  version.UiVersion,
				"API Version": version.ApiVersion,
			})
		}

		gctx.JSON(http.StatusOK, gin.H{
			"data": versionsData,
		})
	}
}

// Get Nodal officer details with UserName.
func GetNodalOfficerByUsername(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		username := gctx.Param("id")

		nodalofficer, err := start.QueryNodalOfficerByUsername(ctx, client, username)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": nodalofficer})

	}

	return gin.HandlerFunc(fn)
}
