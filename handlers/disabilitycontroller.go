package handlers

import (
	"context"
	"net/http"
	"recruit/ent"
	"recruit/start"
	"recruit/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DisabilityRequest struct {
	DisabilityTypeCode        string    `json:"DisabilityTypeCode"`
	DisabilityTypeDescription string    `json:"DisabilityTypeDescription"`
	DisabilityPercentage      int32     `json:"DisabilityPercentage"`
	DisabilityFlag            string    `json:"DisabilityFlag"`
	HallCategoryCode          string    `json:"HallCategoryCode"`
	OrderNumber               string    `json:"OrderNumber"`
	CreatedById               int64     `json:"CreatedById"`
	CreatedByUserName         string    `json:"CreatedByUserName"`
	CreatedByEmployeeId       string    `json:"CreatedByEmployeeId"`
	CreatedByDesignation      string    `json:"CreatedByDesignation"`
	CreatedDate               time.Time `json:"CreatedDate"`
	VerifiedById              int64     `json:"VerifiedById"`
	VerifiedByUserName        string    `json:"VerifiedByUserName"`
	VerifiedByEmployeeId      string    `json:"VerifiedByEmployeeId"`
	VerifiedByDesignation     string    `json:"VerifiedByDesignation"`
	VerifiedDate              time.Time `json:"VerifiedDate"`
	Statuss                   string    `json:"Statuss"`
	DeletedById               int64     `json:"DeletedById"`
	DeletedByUserName         string    `json:"DeletedByUserName"`
	DeletedByEmployeeId       string    `json:"DeletedByEmployeeId"`
	DeletedByDesignation      string    `json:"DeletedByDesignation"`
	DeletedDate               time.Time `json:"DeletedDate"`
}
type DisabilityResponse struct {
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

// / Disabilities

// CreateDisabilityTypes creates a new Disability type.
// @Summary Create a new Disability type
// @Description Create a new Disability type with given parameters
// @Tags Disability
// @Accept json
// @Produce json
// @Param request body DisabilityRequest true "Disability data"
// @Success 200 {object} string "Successfully created the Disability Types"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/employeedisbilitytypes [post]
func CreateDisabilityTypes(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newdisability := new(ent.Disability)
		if err := gctx.ShouldBindJSON(&newdisability); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newdisability, err := start.CreateEmployeeDisability(client, newdisability)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created the Disability Types")
	}
	return gin.HandlerFunc(fn)
}

// GetAllDisabilityTypes godoc
// @Summary Get all disability types
// @Description Get all disability types from the database
// @Tags Disability
// @Accept json
// @Produce json
// @Success 200 {object} DisabilityResponse "Successfully all disability types Retrieved"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "No Articles are invoiced to the beat"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /deptexam/employeedisbilitytypes [get]
func GetAllDisabilityTypes(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		disability, err := start.QueryEmployeeDisabilities(ctx, client)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": disability})

	}
	return gin.HandlerFunc(fn)
}

func GetDisabilitiesByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		id := gctx.Param("id")
		//var examID int32
		disabilityID, _ := strconv.ParseInt(id, 10, 32)

		disability, err := start.QueryEmployeeDisabilityID(ctx, client, int32(disabilityID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": disability})

	}

	return gin.HandlerFunc(fn)
}
