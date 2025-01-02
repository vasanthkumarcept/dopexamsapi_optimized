package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"recruit/ent"
	"recruit/start"
	"recruit/util"

	_ "github.com/lib/pq"

	"strconv"
)

// GetExams godoc
// @Summary Get Exams
// @Description Fetch all available exams.
// @Tags Exam
// @Produce json
// @Success 200 {object} ExamResponse "Active Exam fetched successfully by Exam year"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/exams [get]
func GetExams(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - GetExams "
		var startFunction string = " - start - QueryExam "
		exams, status, stgError, dataStatus, err := start.QueryExam(ctx, client)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		//gctx.JSON(http.StatusOK, gin.H{"data": exams})
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Exam fetched successfully",
			"data":       exams,
			"dataexists": dataStatus,
		})

	}

	return gin.HandlerFunc(fn)

}

func Updateexam(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newExam := new(ent.Exam)

		id := gctx.Param("id")

		examID, _ := strconv.ParseInt(id, 10, 32)

		if err := gctx.ShouldBindJSON(&newExam); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		logdata := newExam.Edges.LogData[0]

		exam, err := start.UpdateExam(client, int32(examID), newExam)
		if err != nil {
			util.LogError(client, logdata, err)
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exam})

		util.Logger(client, logdata)
	}

	return gin.HandlerFunc(fn)
}

type ExamResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	DataExists bool        `json:"dataexists"`
}

// QueryActiveExamsByExamYear godoc
// @Summary Query Active Exams by Exam Year
// @Description Get active exams for a given exam year and facility ID.
// @Tags Exam
// @Produce json
// @Param selectedyear path int true "Selected Year"
// @Param circleidcontext path string true "Circle ID Context"
// @Success 200 {object} ExamResponse "Active Exam fetched successfully by Exam year"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/ActiveExamsByYear/{selectedyear}/{circleidcontext} [get]
func QueryActiveExamsByExamYear(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		//selectedyear/:circleidcontext

		var mainFunction string = " main - QueryActiveExamsByExamYear "
		var startFunction string = " - start - QueryActiveExamsByExamYear "
		eyear64, _ := strconv.ParseInt(gctx.Param("selectedyear"), 10, 32)
		year := int32(eyear64)
		facilityid := gctx.Param("circleidcontext")
		//fmt.Println("active exams", year, facilityid)
		activeExams, status, stgError, dataStatus, err := start.QueryActiveExamsByExamYear(client, year, facilityid)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		//	gctx.JSON(http.StatusOK, gin.H{"data": activeExams})

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Active Exam fetched successfully by Exam year",
			"data":       activeExams,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// QueryActiveExamsByExamYearWithoutCAFacilityID godoc
// @Summary Query Active Exams by Exam Year without CA Facility ID
// @Description Get active exams for a given exam year without specifying CA Facility ID.
// @Tags Exam
// @Produce json
// @Param selectedyear path int true "Selected Year"
// @Success 200 {object} ExamResponse "Active Exam fetched successfully by Exam year WithoutCAFacilityID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/ActiveExamsByYear/{selectedyear} [get]
func QueryActiveExamsByExamYearWithoutCAFacilityID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var mainFunction string = " main - QueryActiveExamsByExamYearWithoutCAFacilityID "
		var startFunction string = " - start - QueryActiveExamsByExamYearWithoutCAFacilityID "
		eyear64, _ := strconv.ParseInt(gctx.Param("selectedyear"), 10, 32)
		year := int32(eyear64)
		activeExams, status, stgError, dataStatus, err := start.QueryActiveExamsByExamYearWithoutCAFacilityID(client, year)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		// if len(activeExams) <= 0 {
		// 	gctx.JSON(http.StatusOK, gin.H{"Message": "No Active exams for the year"})
		// 	return
		// }

		//	gctx.JSON(http.StatusOK, gin.H{"data": activeExams})
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Active Exam fetched successfully by Exam year WithoutCAFacilityID",
			"data":       activeExams,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// QueryCityNamesByExamIDFromExamCenter godoc
// @Summary Query City Names by Exam ID from Exam Center
// @Description Get city names associated with a specific exam ID from the exam center.
// @Tags Exam
// @Produce json
// @Param id1 path string true "ID1"
// @Param id2 path int true "ID2"
// @Success 200 {object} ExamResponse  "City Names fetched successfully by Exam year"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/QueryCityNamesByNotiIDFromExamCenter/{id1}/{id2} [get]
func QueryCityNamesByExamIDFromExamCenter(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var mainFunction string = " main - QueryCityNamesByExamIDFromExamCenter "
		var startFunction string = " - start -QueryCityNamesByExamIDFromExamCenter "
		eid, _ := strconv.ParseInt(gctx.Param("id2"), 10, 32)
		nn := gctx.Param("id1")

		citylist := new(ent.ExamCityCenter)
		citylists, status, stgError, dataStatus, err := start.QueryCityNamesByExamIDFromExamCenter(client, citylist, nn, int32(eid), int32(eid))

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		// citynames := make([]string, 0)
		// for _, cities := range citylists {
		// 	citynames = append(citynames, cities.CenterCityName)
		// }

		citynames := make(map[int]string)
		for _, cities := range citylists {
			citynames[int(cities.ID)] = cities.CenterCityName
		}

		// gctx.JSON(http.StatusOK, gin.H{
		// 	"data": gin.H{
		// 		"Center City Names": citynames,
		// 	}})

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "City Names fetched successfully by Exam year",
			"data":       citynames,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// QueryCityNamesByNotificationIDFromExamCenter godoc
// @Summary Query City Names by Notification ID from Exam Center
// @Description Get city names associated with a specific notification ID, exam ID, and code from the exam center.
// @Tags Exam
// @Produce json
// @Param id1 path string true "Notification ID"
// @Param id2 path int true "Exam ID"
// @Param id3 path int true "Code"
// @Success 200 {object} ExamResponse   "Active Exam fetched successfully by Notification ID"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/QueryCityNamesByNotificationIDFromExamCenter/{id1}/{id2}/{id3} [get]
func QueryCityNamesByNotificationIDFromExamCenter(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		nn := gctx.Param("id1")
		eid, _ := strconv.ParseInt(gctx.Param("id2"), 10, 32)
		code, _ := strconv.ParseInt(gctx.Param("id3"), 10, 32)
		var mainFunction string = " main - QueryCityNamesByNotificationIDFromExamCenter "
		var startFunction string = " - start - QueryCityNamesByExamIDFromExamCenter "
		citylist := new(ent.ExamCityCenter)
		citylists, status, stgError, dataStatus, err := start.QueryCityNamesByExamIDFromExamCenter(client, citylist, nn, int32(eid), int32(code))

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		// gctx.JSON(http.StatusOK, gin.H{
		// 	"data": citylists})

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Active Exam fetched successfully by Notification ID",
			"data":       citylists,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// QueryExamCityNamesForIPExam godoc
// @Summary Query Exam City Names for IP Exam
// @Description Get city names associated with a specific notification number, exam year, and exam code for IP Exam.
// @Tags Exam
// @Produce json
// @Param notificationnumber path string true "Notification Number"
// @Param examyear path int true "Exam Year"
// @Param examcode path int true "Exam Code"
// @Success 200 {object} ExamResponse "Exam city extracted successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/QueryExamCityNamesForIPExam/{notificationnumber}/{examyear}/{examcode} [get]
func QueryExamCityNamesForIPExam(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		// :notificationnumber/:examyear/:examcode
		nn := gctx.Param("notificationnumber")
		eid, _ := strconv.ParseInt(gctx.Param("examyear"), 10, 32)
		code, _ := strconv.ParseInt(gctx.Param("examcode"), 10, 32)
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var startFunction string = " main - QueryExamCityNamesForIPExam - start -QueryExamCityNamesForIPExam "
		//citylist := new(ent.ExamCityCenter)
		citylists, status, stgError, dataStatus, err := start.QueryExamCityNamesForIPExam(client, ctx, nn, int32(eid), int32(code))

		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Exam city extracted successfully",
			"data":       citylists,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}
