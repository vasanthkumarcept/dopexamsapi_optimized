package handlers

import (
	"context"

	"net/http"
	"recruit/ent"
	"recruit/start"
	"recruit/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExamCalendarsResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	DataExists bool        `json:"dataexists"`
}

// GetExamCalendars godoc
// @Summary Get all exam calendars
// @Description Get a list of all exam calendars
// @Tags ExamCalendars
// @Accept  json
// @Produce  json
// @Success 200 {object} ExamCalendarsResponse "Exam calender fetched successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/examcalendars [get]
func GetExamCalendars(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetExamCalendars "
		var startFunction string = " - start - QueryExamCalendars "

		examcalendars, status, stgError, dataStatus, err := start.QueryExamCalendars(ctx, client)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		//gctx.JSON(http.StatusOK, gin.H{"data": examcalendars})
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Exam calender fetched successfully",
			"data":       examcalendars,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)

}

// For Exam Calendars ...!
func CreateExamCalendar(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newExamCalendar := new(ent.ExamCalendar)
		if err := gctx.ShouldBindJSON(&newExamCalendar); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newExamCalendar, err := start.CreateExamCalendar(client, newExamCalendar)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Exam Calendar")
	}
	return gin.HandlerFunc(fn)
}

// GetExamCalendarID godoc
// @Summary Get all exam calendars by ID
// @Description Get a list of all exam calendars by ID
// @Tags ExamCalendars
// @Accept  json
// @Produce  json
// @Param id path int true "Exam Calendar ID"
// @Success 200 {object} ExamCalendarsResponse "Exam calender fetched successfully by ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/examcalendars/{id} [get]
func GetExamCalendarID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetExamCalendarID "
		var startFunction string = " - start - QueryExamCalendarID "

		id := gctx.Param("id")
		//var examID int32
		calendarID, _ := strconv.ParseInt(id, 10, 32)

		calendars, status, stgError, dataStatus, err := start.QueryExamCalendarID(ctx, client, int32(calendarID))
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		//gctx.JSON(http.StatusOK, gin.H{"data": calendars})
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Exam calender fetched successfully by ID",
			"data":       calendars,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)
}
