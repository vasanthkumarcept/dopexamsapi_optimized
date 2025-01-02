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

// / Exam Papertypes
func InsertExamPaperTypes(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newExamPapertypes := new(ent.PaperTypes)
		if err := gctx.ShouldBindJSON(&newExamPapertypes); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newExamPapertypes, err := start.CreatePaperType(client, newExamPapertypes)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Exam Paper Types")
	}
	return gin.HandlerFunc(fn)
}

func GetExamPaperTypesByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		ExamPaperTypeID, _ := strconv.ParseInt(id, 10, 32)
		exampapertypes, err := start.QueryExamPaperTypeByID(ctx, client, int32(ExamPaperTypeID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exampapertypes})
	}
	return gin.HandlerFunc(fn)
}

func GetAllExamPaperTypes(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		papertypes, err := start.QueryExamPaperTypes(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": papertypes})

	}

	return gin.HandlerFunc(fn)

}

// QueryExamPapersByExamCode
// QueryCircleMasterWithRegions
func GetExamPaperTypesWithPaperCode(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		PaperID, _ := strconv.ParseInt(id, 10, 32)
		papertypes, err := start.QueryExamPaperTypesByPaperCode(ctx, client, int32(PaperID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Get ExamPapers by ExamCode and retrieve only ID and PaperDescription.

		// Create an array to store the PaperID and PaperDescription.
		var papertypeData []map[string]interface{}
		for _, paper := range papertypes {
			papertypeData = append(papertypeData, map[string]interface{}{
				"PaperID":          paper.ID,
				"PaperDescription": paper.PaperTypeDescription,
			})
		}
		gctx.JSON(http.StatusOK, gin.H{"data": papertypeData})
	}
	return gin.HandlerFunc(fn)
}

// For ExamPapers
func InsertExamPapers(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newExamPapers := new(ent.ExamPapers)
		if err := gctx.ShouldBindJSON(&newExamPapers); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newExamPapers, err := start.CreatePapers(client, newExamPapers)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Exam Papers")
	}
	return gin.HandlerFunc(fn)
}

func GetExamPapersByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		ExamPaperID, _ := strconv.ParseInt(id, 10, 32)
		exampapers, err := start.QueryExamPaperByID(ctx, client, int32(ExamPaperID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exampapers})
	}
	return gin.HandlerFunc(fn)
}
