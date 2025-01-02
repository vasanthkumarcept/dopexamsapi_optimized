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

// Eligibility Master
func InsertEligibilityCriteria(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		eligibilities := new(ent.EligibilityMaster)
		if err := gctx.ShouldBindJSON(&eligibilities); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		eligibilities, err := start.CreateEligibilities(client, eligibilities)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Eligibility criteria")
	}
	return gin.HandlerFunc(fn)
}

func GetEligibilityByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		eliID, _ := strconv.ParseInt(id, 10, 32)
		elis, err := start.QueryEligibilitiyMasterByID(ctx, client, int32(eliID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": elis})
	}
	return gin.HandlerFunc(fn)
}

func GetAllEligibilitiyCriteria(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		Dusers, err := start.QueryEligibilitiyMaster(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": Dusers})

	}

	return gin.HandlerFunc(fn)

}

func GetExamWithEligibility(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		examcode, _ := strconv.ParseInt(id, 10, 32)
		exams, err := start.QueryEligibilityWithPapers(ctx, client, int32(examcode))

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exams})

	}
	return gin.HandlerFunc(fn)
}
