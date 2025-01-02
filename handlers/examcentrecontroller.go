package handlers

import (
	"net/http"
	"recruit/ent"
	"recruit/start"
	"recruit/util"
	"strconv"

	ca_reg "recruit/payloadstructure/candidate_registration"

	"github.com/gin-gonic/gin"
)

type CenterResponse struct {
	ID             int32  `json:"ID"`
	ExamCenterName string `json:"ExamCenterName"`
	Address        string `json:"Address"`
	Pincode        int32  `json:"Pincode"`
	Landmark       string `json:"Landmark"`
	CenterCityName string `json:"CenterCityName"`
	Message        string `json:"Message"`
	Status         bool   `json:"Status"`
}

// Update Exam Centers ...

// UpdateExamCenter godoc
// @Summary Update an Exam Center
// @Description Update an existing exam center by ID
// @Tags ExamCenters
// @Accept json
// @Produce json
// @Param id path int true "Center ID"
// @Param center body ca_reg.CenterRequest true "Exam Center Data"
// @Success 200 {object} CenterResponse "Exam Centre Updated"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/center/update/{id} [put]
func UpdateExamCenter(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		var newCenter ca_reg.CenterRequest
		id := gctx.Param("id")
		CenterID, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			Remarks := "400 error - invalid CenterID: " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		var mainFunction string = " main - UpdateExamCenter "
		var startFunction string = " - start - UpdateCenter "

		if err := gctx.ShouldBindJSON(&newCenter); err != nil {
			Remarks := "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		logdata := newCenter.Edges.LogData[0]

		updatedCenter, status, stgError, dataStatus, err := start.UpdateCenter(client, int32(CenterID), newCenter)
		if err != nil {
			Remarks := mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}

		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Exam Centre Updated",
			"data":       updatedCenter,
			"dataexists": dataStatus,
		})
	}
}
