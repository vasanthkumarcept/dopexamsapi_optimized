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

// For Vacancy Years ...!
func CreateVacarncyYears(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newVacancyYear := new(ent.VacancyYear)

		if err := gctx.ShouldBindJSON(&newVacancyYear); err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newVacancyYear, err := start.CreateVacancyYears(client, newVacancyYear)

		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Vacancy Year")

	}
	return gin.HandlerFunc(fn)
}

func GetVacancyYearID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		id := gctx.Param("id")
		//var examID int32
		VacancyYearID, _ := strconv.ParseInt(id, 10, 32)

		vys, err := start.QueryVacancyYearID(ctx, client, int32(VacancyYearID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": vys})

	}

	return gin.HandlerFunc(fn)
}

func UpdateVacancyYearID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newVacYr := new(ent.VacancyYear)
		id := gctx.Param("id")
		VacancyYearID, _ := strconv.ParseInt(id, 10, 32)
		if err := gctx.ShouldBindJSON(&newVacYr); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		vcy, err := start.UpdateVacancyYearID(client, int32(VacancyYearID), newVacYr)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": vcy})
	}
	return gin.HandlerFunc(fn)
}

func GetVacancyYears(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		vcs, err := start.QueryVacancyYear(ctx, client)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": vcs})

	}

	return gin.HandlerFunc(fn)

}
