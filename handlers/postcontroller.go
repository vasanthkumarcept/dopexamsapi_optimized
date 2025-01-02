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

// / Employee posts
func CreateEmployeePost(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newposts := new(ent.EmployeePosts)
		if err := gctx.ShouldBindJSON(&newposts); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newposts, err := start.CreateEmployeePosts(client, newposts)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Employee Posts")
	}
	return gin.HandlerFunc(fn)
}

func GetAllEmployeePosts(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		posts, err := start.QueryEmployeePosts(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": posts})

	}
	return gin.HandlerFunc(fn)
}

func GetEmpPostsByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		id := gctx.Param("id")
		//var examID int32
		postID, _ := strconv.ParseInt(id, 10, 32)

		posts, err := start.QueryEmployeePostsID(ctx, client, int32(postID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": posts})

	}

	return gin.HandlerFunc(fn)
}
