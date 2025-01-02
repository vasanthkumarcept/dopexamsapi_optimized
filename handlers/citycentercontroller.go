package handlers

import (
	"context"
	"net/http"
	"recruit/ent"
	"recruit/start"
	"recruit/util"
	"strconv"
	"time"

	ca_reg "recruit/payloadstructure/candidate_registration"

	"github.com/gin-gonic/gin"
)

type ExamCityCenterCreateRequest struct {
	ExamCode              int32     `json:"ExamCode" binding:"required"`
	ExamName              string    `json:"ExamName" binding:"required"`
	ExamShortName         string    `json:"ExamShortName"`
	ExamYear              int32     `json:"ExamYear" binding:"required"`
	ConductedBy           string    `json:"ConductedBy"`
	NodalOfficeFacilityID string    `json:"NodalOfficeFacilityID"`
	NodalOfficeName       string    `json:"NodalOfficeName"`
	NotificationCode      int32     `json:"NotificationCode"`
	NotificationNumber    string    `json:"NotificationNumber"`
	CenterCityName        string    `json:"CenterCityName"`
	CreatedById           int64     `json:"CreatedById"`
	CreatedByUserName     string    `json:"CreatedByUserName"`
	CreatedByEmpId        int64     `json:"CreatedByEmpId"`
	CreatedByDesignation  string    `json:"CreatedByDesignation"`
	Status                string    `json:"Status"`
	CircleCityName        string    `json:"CircleCityName"`
	DivisionCode          int32     `json:"DivisionCode"`
	RegionCode            int32     `json:"RegionCode"`
	DivisionName          string    `json:"DivisionName"`
	RegionID              int32     `json:"RegionID"`
	RegionName            string    `json:"RegionName"`
	RegionCityName        string    `json:"RegionCityName"`
	CentreCityName        string    `json:"CentreCityName"`
	Remarks               string    `json:"Remarks"`
	UpdatedAt             time.Time `json:"UpdatedAt"`
	UpdatedBy             string    `json:"UpdatedBy"`
	CentreCode            int32     `json:"CentreCode"`
	CircleID              int32     `json:"CircleID"`
}
type ExamCityCenterCreateResponse struct {
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

// CreateExamCityCenters godoc
// @Summary Create a new exam city center
// @Description Create a new exam city center with the input payload
// @Tags ExamCityCenters
// @Accept json
// @Produce json
// @Param examCityCenter body ExamCityCenterCreateRequest true "Exam City Center"
// @Success 200 {object} ExamCityCenterCreateResponse "Exam City Center created successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/CreateExamCityCenters [post]
func CreateExamCityCenters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var mainFunction string = " main - CreateExamCityCenters "
		var startFunction string = " - start CreateExamCityCenters "
		var newExamCity ca_reg.ExamCityCenterRequest

		if err := gctx.ShouldBindJSON(&newExamCity); err != nil {
			Remarks := "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		logdata := newExamCity.Edges.LogData[0]
		newExamcity, status, stgError, dataStatus, err := start.CreateExamCityCenters(client, &newExamCity)

		if err != nil {
			Remarks := mainFunction + startFunction
			start.StartErrorHandlerWithLog(gctx, err, status, stgError, logdata, client, Remarks)
			return
		}
		util.LoggerNew(client, logdata)
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Exam City Center created successfully",
			"data":       newExamcity,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// GetCentersByCity godoc
// @Summary Get Exam City Centers by City
// @Description Retrieve all exam city centers for a specified city
// @Tags ExamCityCenters
// @Accept json
// @Produce json
// @Param id path string true "City ID"
// @Success 200 {object} ExamCityCenterCreateResponse "Exam City Center retrieved  successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/centersBycity/{id} [get]
func GetCentersByCity(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		city := gctx.Param("id")
		centers, status, err := start.QueryCenterByCity(ctx, client, city)
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.QueryCenterByCity " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "400"
				Remarks = "start.QueryCenterByCity " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return
			} else {
				Action = "500"
				Remarks = "start.QueryCenterByCity " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": centers})
	}
	return gin.HandlerFunc(fn)
}

// GetCentersByCircleID godoc
// @Summary Get Exam City Centers by Circle ID
// @Description Retrieve all exam city centers for a specified circle ID
// @Tags ExamCityCenters
// @Accept json
// @Produce json
// @Param id path int true "Circle ID"
// @Success 200 {object} ExamCityCenterCreateResponse "Exam City Center retrieved successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/centersByCircleID/{id} [get]
func GetCentersByCircleID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		cid := gctx.Param("id")
		circleid, _ := strconv.ParseInt(cid, 10, 32)

		centers, status, err := start.QueryCenterByCircleID(ctx, client, int32(circleid))
		if err != nil {
			if status == 400 {
				Action = "400"
				Remarks = "start.QueryCenterByCircleID " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			} else if status == 422 {
				Action = "400"
				Remarks = "start.QueryCenterByCircleID " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return
			} else {
				Action = "500"
				Remarks = "start.QueryCenterByCircleID " + err.Error()
				util.SystemLogError(client, Action, Remarks)
				gctx.JSON(http.StatusBadRequest, gin.H{"error": UserErrorRemarks})
				return
			}
		}
		gctx.JSON(http.StatusOK, gin.H{"data": centers})
	}
	return gin.HandlerFunc(fn)
}

func GetCentersByConductingAuthority(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetCentersByConductingAuthority "
		var startFunction string = " - start - SubGetCentersByConductingAuthority "
		//:examyear/:examcode/:conductingid
		conductingId := gctx.Param("conductingid")
		examYear := gctx.Param("examyear")
		examCode := gctx.Param("examcode")
		centers, status, stgError, dataStatus, err := start.SubGetCentersByConductingAuthority(ctx, client, examYear, examCode, conductingId)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"dataexists": dataStatus,
			"data":       centers,
		})
	}
	return gin.HandlerFunc(fn)
}

/* func GetCenters(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		centers, err := start.QueryCenter(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": centers})

	}

	return gin.HandlerFunc(fn)

} */
