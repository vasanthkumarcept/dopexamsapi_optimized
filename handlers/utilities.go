package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"recruit/ent"
	"recruit/ent/pdf"
	"recruit/start"
	"recruit/util"
	"strconv"

	ca_reg "recruit/payloadstructure/candidate_registration"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

/* func GetDetailsBasedOnExamCode(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		errRemarks := "500 error DB Connection erorr -DBS01"
		errStatus := start.HandleDBErrorInitial(gctx, client, errRemarks)
		if errStatus == "error" {
			return
		}

		id := gctx.Param("Exam_Code")
		id1, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Exam_Code provided"})
			return
		}
		var mainFunction string = " main - GetDetailsBasedOnExamCode "

		// Handle the case when the Exam_Code is not found
		if id1 == 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Exam_Code is missing or invalid"})
			return
		}

		// Get the details for the given exam code
		//categoryData, postMappingData, Err := util.QueryDetailsByExamCode(ctx, client, id1)
		var startFunction string = " - start - QueryDetailsByExamCode "
		categoryData, status, stgError, _, err := start.QueryCategoryData(ctx, client, id1)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		startFunction = " - start - QueryCategoryData "
		postMappingData, status, stgError, dataStatus, err := start.QueryPostMappingData(ctx, client, id1)

		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Category and Post mapping data  successfully",
			"data": gin.H{
				"examCategoryDisabilityMapping": categoryData,
				"examPostMapping":               postMappingData,
			},
			"dataexists": dataStatus,
		})
	}
} */

func CreateEmployeeCategories(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		errRemarks := "500 error DB Connection erorr -DBS01"
		errStatus := start.HandleDBErrorInitial(gctx, client, errRemarks)
		if errStatus == "error" {
			return
		}
		var mainFunction string = " main - CreateEmployeeCategories "
		var startFunction string = " - start - CreateEmployeeCategory "
		newEmployeecategories := new(ent.EmployeeCategory)
		if err := gctx.ShouldBindJSON(&newEmployeecategories); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}
		newEmployeecategories, status, stgError, dataStatus, err := start.CreateEmployeeCategory(client, newEmployeecategories)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Employee category fetched successfully",
			"data":       newEmployeecategories,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

func GetAllEmployeeCategories(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		errRemarks := "500 error DB Connection erorr -DBS01"
		errStatus := start.HandleDBErrorInitial(gctx, client, errRemarks)
		if errStatus == "error" {
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		var mainFunction string = " main - getAllEmployeeCategories "
		var startFunction string = " - start - QueryEmployeeCategories "

		newEmployeCategories, status, stgError, dataStatus, err := start.QueryEmployeeCategories(ctx, client)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Employee category data extracted successfully",
			"data":       newEmployeCategories,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)

}

func GetAllCircles(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		errRemarks := "500 error DB Connection erorr -DBS01"
		errStatus := start.HandleDBErrorInitial(gctx, client, errRemarks)
		if errStatus == "error" {
			return
		}
		var startFunction string = " main - QueryOfficeByPincodehandle - start - QueryOfficeByPincode "
		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		circles, status, stgError, dataStatus, err := start.SubGetAllCircles(ctx, client)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       circles,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

func PaperDetails(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		examid, _ := strconv.ParseInt(gctx.Param("id"), 10, 32)

		paperlist := new(ent.ExamPapers)
		paperlists, err := start.PaperDetails(client, paperlist, int32(examid))
		if err != nil {

			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"data": paperlists,
		})

	}
	return gin.HandlerFunc(fn)
}

func GetDetailsBasedOnExamCode(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		var mainFunction string = " main - GetDetailsBasedOnExamCode "
		var startFunction string = " - start - QueryDetailsByExamCode "
		ctx := gctx.Request.Context()
		id := gctx.Param("Exam_Code")
		id1, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Exam_Code provided"})
			return
		}

		// Handle the case when the Exam_Code is not found
		if id1 == 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Exam_Code is missing or invalid"})
			return
		}

		// Get the details for the given exam code
		categoryData, postMappingData, status, stgError, _, err := start.QueryDetailsByExamCode(ctx, client, id1)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}

		response := gin.H{
			"data": gin.H{
				"examCategoryDisabilityMapping": categoryData,
				"examPostMapping":               postMappingData,
			},
		}

		gctx.JSON(http.StatusOK, response)
	}
}

func Totopverify(c *gin.Context) {
	passcode := c.GetHeader("totp")

	//passcode := c.Request.Header["totp"][0]
	if passcode != "" {
		//c.String(http.StatusUnauthorized, "unauthorised")

		//valid := totp.Validate(passcode, key.Secret())

		valid := totp.Validate(passcode, "MFSWCNBRMVSGGY3FMZTGCZBQGY2WMYZYME4WIZDGGY3DMMJYGUZWMMRZG43TMMBYGVRGKNTFGAYGMMRTGVRWMZDDHAZDAZTGGM4WC")
		if valid {
			//c.String(http.StatusOK, "Valid")
			//os.Exit(0)
		} else {
			c.String(http.StatusUnauthorized, "unauthorised")
			c.Abort()
			return
			//os.Exit(1)
		}

	} else {

		c.String(http.StatusUnauthorized, "unauthorised")
		c.Abort()
		return
	}

	c.Next()

}

func Topgenerate(c *gin.Context) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "gateway.cept.gov.in",
		AccountName: "Vams",
		Secret:      []byte("aea41edcceffad065fc8a9ddf6661853f29776085be6e00f235cfdc820ff39a"),
		Period:      90,
		Algorithm:   otp.AlgorithmSHA256,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"key": key.Secret()})

}

func TestSMS(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		napitype := gctx.Param("apitypes")
		apiresponse, apiresponsedescription := start.SubTestSMS(ctx, client, napitype)
		gctx.JSON(http.StatusOK, gin.H{
			"ApiResponse":    apiresponse,
			"APIDescription": apiresponsedescription,
		})
	}
	return gin.HandlerFunc(fn)
}

func GetVersion(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		uiversion := gctx.Param("uiversion")
		apiversion := os.Getenv("API_VERSION")

		var mainFunction string = " main - GetVersion "
		var startFunction string = " - start - SubVersion "

		// Get the details for the given exam code
		versiondata, status, stgError, dataStatus, err := start.SubVersion(ctx, client, uiversion, apiversion)
		fmt.Println(versiondata, status, stgError, dataStatus, err)
		if err != nil {
			if !dataStatus {
				Remarks = mainFunction + startFunction
				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
				return
			}
		}

		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    err.Error(),
			"data":       versiondata,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

func GetMessage(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		var mainFunction string = " main - GetMessage "
		var startFunction string = " - start - SubMessage "

		// Get the details for the given exam code
		messagedata, status, stgError, dataStatus, err := start.SubMessage(ctx, client)
		fmt.Println(messagedata, status, stgError, dataStatus, err)
		if err != nil {
			if !dataStatus {
				Remarks = mainFunction + startFunction
				start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
				return
			}
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       messagedata,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

// Paths of PDF
func GetPdf(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		code1 := gctx.Param("code")
		year1 := gctx.Param("year")
		year, _ := strconv.Atoi(year1)
		code, _ := strconv.Atoi(code1)

		pdfs, err := client.PDF.Query().Where(pdf.ExamcodeEQ(code), pdf.YearEQ(year)).All(ctx)

		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": pdfs})

	}

	return gin.HandlerFunc(fn)
}

func ErrorLogAssignment(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		var serviceRequest ca_reg.StrucSericeRequest

		var mainFunction string = " main - ErrorLogAssignment "
		var startFunction string = " - start - SubErrorLogAssignment "

		if err := gctx.ShouldBindJSON(&serviceRequest); err != nil {
			Remarks = "400 error from " + mainFunction + " - ShouldBindJSON " + err.Error()
			start.MainHandleError(gctx, client, Remarks, " -HA01", gctx.GetHeader("UserName"))
			return
		}

		_, status, stgError, _, err := start.SubErrorLogAssignment(client, &serviceRequest)
		if err != nil {
			Remarks = mainFunction + startFunction
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, Remarks)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"message": "Error log updated successfully",
			},
		})
	}
	return gin.HandlerFunc(fn)
}

func GetErrorLogs(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		adminUserName := gctx.Param("adminusername")
		var startFunction string = " main - GetErrorLogs - start.SubGetErrorLogs "
		errorLogs, status, stgError, dataStatus, err := start.SubGetErrorLogs(ctx, client, adminUserName)
		if err != nil {
			// err error, status int32, stgError string, client *ent.Client, Remarks string
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       errorLogs,
			"dataexists": dataStatus,
		})
	}
	return gin.HandlerFunc(fn)
}

/* func QueryCategoryData(ctx context.Context, client *ent.Client, id int64) ([]map[string]interface{}, error) {
	mappings, err := client.ExamCategoryDisabilityMapping.Query().
		Where(
			examcategorydisabilitymapping.ExamCodeEQ(id),
			examcategorydisabilitymapping.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve category mappings: %w", err)
	}

	response := make([]map[string]interface{}, 0) // Response should be a slice of maps

	categories := make([]map[string]interface{}, 0)
	disabilities := make([]map[string]interface{}, 0)

	for _, category := range mappings {
		categoryData := map[string]interface{}{
			"categoryCode":        category.CategoryDisabilityCode,
			"categoryDescription": category.CategoryDisabilityDescription,
			"ageExemption":        category.AgeException,
			"serviceExemption":    category.ServiceException,
			"drivingLicense":      category.DrivingLicense,
		}

		if category.CategoryDisability == "category" {
			categories = append(categories, categoryData)
		} else if category.CategoryDisability == "disability" {
			disabilities = append(disabilities, categoryData)
		}
	}

	response = append(response, map[string]interface{}{
		"data": []map[string]interface{}{
			{
				"category": categories,
			},
			{
				"disability": disabilities,
			},
		},
	})

	return response, nil
}
func QueryPostMappingData(ctx context.Context, client *ent.Client, id int64) ([]map[string]interface{}, error) {
	mappings, err := client.ExamPostMapping.Query().
		Where(
			exampostmapping.ExamCodeEQ(id),
			exampostmapping.StatusEQ("active"),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve post mapping data: %w", err)
	}

	response := make([]map[string]interface{}, 0) // Response should be a slice of maps

	entry := make([]map[string]interface{}, 0)
	present := make([]map[string]interface{}, 0)
	feder := make([]map[string]interface{}, 0)

	for _, category1 := range mappings {
		categoryData1 := map[string]interface{}{
			"PostID":            category1.FromPostCode,
			"PostdDescription":  category1.FromPostDescription,
			"ageCriteria":       category1.AgeCriteria,
			"serviceCriteria":   category1.ServiceCriteria,
			"ToPostDescription": category1.ToPostDescription,
			"ToPostCode":        category1.ToPostCode,
		}

		if category1.PostTypeDescription == "entry cader" {
			entry = append(entry, categoryData1)
		} else if category1.PostTypeDescription == "present cader" {
			present = append(present, categoryData1)
		} else if category1.PostTypeDescription == "feder cader" {
			feder = append(feder, categoryData1)
		}
	}

	response = append(response, map[string]interface{}{
		"data": []map[string]interface{}{
			{
				"entry cader": entry,
			},
			{
				"present cader": present,
			},
			{
				"feder cader": feder,
			},
		},
	})

	return response, nil
} */
