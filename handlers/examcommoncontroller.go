package handlers

import (
	"context"
	"errors"
	"net/http"
	"recruit/ent"
	"recruit/start"
	"recruit/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExamStatisticsResponse struct {
	Success    bool                     `json:"success"`
	Message    string                   `json:"message"`
	Data       []map[string]interface{} `json:"data"`
	DataExists bool                     `json:"dataexists"`
}

// GetExamApplicationsStatisticsOfficeWise godoc
// @Summary Get Exam Applications Statistics Office Wise
// @Description Get the statistics of exam applications office-wise based on the exam code and facility ID.
// @Tags Exam
// @Produce json
// @Param examcode path int true "Exam Code"
// @Param nofacilityid path string true "Nodal Office Facility ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} ExamStatisticsResponse "ExamApplications Statistics fetched Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/exams/statistics/officewise/get/{examcode}/{nofacilityid}/{examyear} [get]
func GetExamApplicationsStatisticsOfficeWise(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamApplicationsStatisticsOfficeWise- start - GetExamStatisticsDOOfficeWise "

		//:examcode/:nofacilityid/:examyear
		ec := gctx.Param("examcode")
		ExamCodeInt64, _ := strconv.ParseInt(ec, 10, 32)
		ExamCode := int32(ExamCodeInt64)
		NOOfficeID := gctx.Param("nofacilityid")
		Examyear := gctx.Param("examyear")

		var statistics []map[string]interface{}
		statistics, status, stgError, dataStatus, err := start.GetExamStatisticsDOOfficeWise(ctx, client, ExamCode, NOOfficeID, Examyear)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "ExamApplications Statistics fetched Successfully",
			"data":       statistics,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

// Summary circle wise

// GetExamStatisticsForDirectorateCircleWise godoc
// @Summary Get Exam Statistics for Directorate Circle Wise
// @Description Get the statistics of exam applications for directorate circle wise based on the exam code and exam year.
// @Tags Exam
// @Produce json
// @Param examcode path int true "Exam Code"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} ExamStatisticsResponse "Getting Nodal officer statistics done successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/exams/statistics/fordirectorate/circlewise/get/examyear{examcode}/{examyear} [get]
func GetExamStatisticsForDirectorateCircleWise(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		ec := gctx.Param("examcode")
		ExamCodeInt64, _ := strconv.ParseInt(ec, 10, 32)
		ExamCode := int32(ExamCodeInt64)
		ExamYear := gctx.Param("examyear")
		var statistics []map[string]interface{}
		var err error
		var status int32
		var dataStatus bool
		var stgError, startFunction string
		var mainFunction string = " main - GetExamStatisticsForDirectorateCircleWise "

		if ExamCode == 2 {
			startFunction = " - start - GetIPExamStatisticsCircleWise "
			statistics, status, stgError, dataStatus, err = start.GetIPExamStatisticsCircleWise(ctx, client, ExamCode, ExamYear)
		} else if ExamCode == 1 {
			startFunction = " - start - GetPSExamStatisticsCircleWise "
			statistics, status, stgError, dataStatus, err = start.GetPSExamStatisticsCircleWise(ctx, client, ExamCode, ExamYear)
		} else if ExamCode == 4 {
			startFunction = " - start - GetGDSPAExamStatisticsCircleWise "
			statistics, status, stgError, dataStatus, err = start.GetGDSPAExamStatisticsCircleWise(ctx, client, ExamCode, ExamYear)
		} else if ExamCode == 3 {
			startFunction = " - start - GetPMPAExamStatisticsCircleWise "
			statistics, status, stgError, dataStatus, err = start.GetPMPAExamStatisticsCircleWise(ctx, client, ExamCode, ExamYear)
		} else if ExamCode == 6 {
			startFunction = " - start - GetGDSPMExamStatisticsCircleWise "
			statistics, status, stgError, dataStatus, err = start.GetGDSPMExamStatisticsCircleWise(ctx, client, ExamCode, ExamYear)
		} else {
			stgError = "HA-004"
			status = 422
			Remarks = mainFunction + startFunction
			err = errors.New(" invalid ExamCode ")
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		if len(statistics) == 0 {
			stgError = "HA-004"
			status = 422
			Remarks = mainFunction + startFunction
			err = errors.New(" no valid applications are available for the Office")
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Getting Nodal officer statistics done successfully",
			"data":       statistics,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

// For Directorate Statistics .
// Summary..

// GetExamApplicationsStatisticsForDirectorate godoc
// @Summary Get Exam Applications Statistics for Directorate
// @Description Get the statistics of exam applications for directorate based on the exam code and selected year.
// @Tags Exam
// @Produce json
// @Param examcode path int true "Exam Code"
// @Param selectedyear path string true "Selected Year"
// @Success 200 {object} ExamStatisticsResponse "Getting statistics based on application status done successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/exams/statistics/fordirectorate/get/{examcode}/{selectedyear} [get]
func GetExamApplicationsStatisticsForDirectorate(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		ec := gctx.Param("examcode")
		ExamCodeInt64, _ := strconv.ParseInt(ec, 10, 32)
		ExamCode := int32(ExamCodeInt64)
		examYear := gctx.Param("selectedyear")
		var statistics []map[string]interface{}
		var err error
		var status int32
		var dataStatus bool
		var stgError, startFunction string
		var mainFunction string = " main - GetExamApplicationsStatisticsForDirectorate "

		if ExamCode == 2 {
			startFunction = " - start - GetIPExamStatistics "
			statistics, status, stgError, dataStatus, err = start.GetIPExamStatistics(ctx, client, ExamCode, examYear)
		} else if ExamCode == 1 {
			startFunction = " - start - GetPSExamStatistics "
			statistics, status, stgError, dataStatus, err = start.GetPSExamStatistics(ctx, client, ExamCode, examYear)
		} else if ExamCode == 4 {
			startFunction = " - start -GetGDSPAExamStatistics "
			statistics, status, stgError, dataStatus, err = start.GetGDSPAExamStatistics(ctx, client, ExamCode, examYear)
		} else if ExamCode == 3 {
			startFunction = " - start - GetPMPAExamStatistics "
			statistics, status, stgError, dataStatus, err = start.GetPMPAExamStatistics(ctx, client, ExamCode, examYear)
		} else if ExamCode == 6 {
			startFunction = " - start - GetGDSPMExamStatistics "
			statistics, status, stgError, dataStatus, err = start.GetGDSPMExamStatistics(ctx, client, ExamCode, examYear)
		} else {
			stgError = "HA-004"
			status = 422
			Remarks = mainFunction + startFunction
			err = errors.New(" invalid ExamCode ")
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		if len(statistics) == 0 {
			stgError = "HA-004"
			status = 422
			Remarks = mainFunction + startFunction
			err = errors.New(" no valid applications are available for the Office")
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Getting statistics based on application status done successfully",
			"data":       statistics,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}

// Get Statistics of Exam Applications for the NO

// GetExamApplicationsStatisticsForNO godoc
// @Summary Get Exam Applications Statistics for Nodal Officer
// @Description Get the statistics of exam applications for Nodal Officer based on the exam code, facility ID, and exam year.
// @Tags Exam
// @Produce json
// @Param examcode path int true "Exam Code"
// @Param nofacilityid path string true "Facility ID"
// @Param examyear path string true "Exam Year"
// @Success 200 {object} ExamStatisticsResponse "Getting statistics based on application status done successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/exams/statistics/get/{examcode}/{nofacilityid}/{examyear} [get]
func GetExamApplicationsStatisticsForNO(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - GetExamApplicationsStatisticsForNO - start - GetEligibleApplicationsForCircleDetailsTestNew "
		ec := gctx.Param("examcode")
		ExamCodeInt64, _ := strconv.ParseInt(ec, 10, 32)
		ExamCode := int32(ExamCodeInt64)
		NOOfficeID := gctx.Param("nofacilityid")
		Examyear := gctx.Param("examyear")

		var statistics []map[string]interface{}
		var err error
		statistics, status, stgError, dataStatus, err := start.GetEligibleApplicationsForCircleDetails(ctx, client, ExamCode, NOOfficeID, Examyear)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       statistics,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}
