// @title Departmental Examination API
// @version 1.0
// @description This is a departmental examination API with Swagger documentation
// @host localhost:8080
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"recruit/ent"
	"recruit/mail"

	h "recruit/handlers"
	"recruit/util"
	"strconv"
	"syscall"
	"time"

	_ "recruit/docs" // Import Swagger docs

	_ "entgo.io/ent"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//var clidb *ent.Client

const UserErrorRemarks = "Try after sometime if still error persists contact CEPT"

var Action string
var Remarks string

func init() {
	time.Local = time.FixedZone("Asia/Kolkata", 5*60*60+30*60) // UTC+5:30

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed loading env file: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	// Configure the connection pool
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v\n", err)
	}

	//new modified
	// Convert string values to appropriate types
	maxConns, err := strconv.Atoi(os.Getenv("DB_MAX_CONNS"))
	if err != nil {
		log.Fatalf("Invalid DB_MAX_CONNS value: %v\n", err)
	}
	minConns, err := strconv.Atoi(os.Getenv("DB_MIN_CONNS"))
	if err != nil {
		log.Fatalf("Invalid DB_MIN_CONNS value: %v\n", err)
	}
	maxConnLifetime, err := time.ParseDuration(os.Getenv("DB_MAX_CONN_LIFETIME"))
	if err != nil {
		log.Fatalf("Invalid DB_MAX_CONN_LIFETIME value: %v\n", err)
	}
	maxConnIdleTime, err := time.ParseDuration(os.Getenv("DB_MAX_CONN_IDLE_TIME"))
	if err != nil {
		log.Fatalf("Invalid DB_MAX_CONN_IDLE_TIME value: %v\n", err)
	}

	config.MaxConns = int32(maxConns)
	config.MinConns = int32(minConns)
	config.MaxConnLifetime = maxConnLifetime
	config.MaxConnIdleTime = maxConnIdleTime

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	// Register the pgx driver with database/sql
	sqlDB, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to open sql.DB: %v\n", err)
	}
	defer sqlDB.Close()

	// Use the connection pool with Ent
	drv := entsql.OpenDB(dialect.Postgres, sqlDB)
	client := ent.NewClient(ent.Driver(drv))
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("Failed to close client: %v\n", err)
		}
	}()

	//client, err := ent.Open("postgres", "host=172.24.19.181 port=2000 user=postgres dbname=recruitment password=secret sslmode=disable")

	//client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=secret sslmode=disable")
	//202 development

	// Handle termination signals to close the pool gracefully
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Shutting down...")
		client.Close()
		os.Exit(0)
	}()

	//Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	initRoutes(client)
	client.Close()

}

func initRoutes(client *ent.Client) {
	//var apiSubPath string = "/deptexam/"
	apiSubPath := os.Getenv("API_SUB_PATH")
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "PUT", "GET", "OPTIONS", "DELETE"}
	config.MaxAge = 48 * time.Hour
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Cache-Control", "Access-Control-Allow-Origin", "UidToken", "UserName"}

	r.Use(cors.New(config))
	r.GET(apiSubPath+"/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Recruitment APIs"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	r.GET(apiSubPath+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.POST(apiSubPath+"sendsms", GenerateOTPAndSendSMSAndSave(client))
	admin := r.Group("/api/admin")
	admin.Use(h.ValidateToken)
	admin.Use(h.Totopverify)
	//admin.Use(gettoken)
	fileHandler := h.FileHandler{}

	//r.DELETE(apiSubPath+"employees/delete/:empid", util.TokenValidationMiddlewareNew(client), h.DeleteUserbyEmployeeID(client))

	r.GET(apiSubPath+"ActiveExamsByYear/:selectedyear", util.TokenValidationMiddlewareNew(client), h.QueryActiveExamsByExamYearWithoutCAFacilityID(client))
	r.GET(apiSubPath+"ActiveExamsByYear/:selectedyear/:circleidcontext", util.TokenValidationMiddlewareNew(client), h.QueryActiveExamsByExamYear(client))
	r.GET(apiSubPath+"CircleHeadQuarters/:id1/:id2", util.TokenValidationMiddlewareNew(client), h.QueryCircleHeadQuartersByExamConductedBy(client))
	r.GET(apiSubPath+"CircleHeadQuarters/Directorate", util.TokenValidationMiddlewareNew(client), h.QueryCircleHeadQuarters(client))
	r.GET(apiSubPath+"GetFacilitiesByCircleOfficeId/:id", util.TokenValidationMiddlewareNew(client), h.GetFacilitiesByCircleOfficeId(client))               //---Hari dev
	r.GET(apiSubPath+"GetFacilitiesByReportingOfficeId/:facilityid", util.TokenValidationMiddlewareNew(client), h.GetFacilitiesByReportingOfficeId(client)) //---vasanth dev

	r.GET(apiSubPath+"PaperDetails/:id", util.TokenValidationMiddlewareNew(client), h.PaperDetails(client))
	r.GET(apiSubPath+"QueryCityNamesByNotiIDFromExamCenter/:id1/:id2", util.TokenValidationMiddlewareNew(client), h.QueryCityNamesByExamIDFromExamCenter(client))
	r.GET(apiSubPath+"QueryCityNamesByNotificationIDFromExamCenter/:id1/:id2/:id3", util.TokenValidationMiddlewareNew(client), h.QueryCityNamesByNotificationIDFromExamCenter(client)) //Hari
	r.GET(apiSubPath+"QueryExamCityNamesForIPExam/:notificationnumber/:examyear/:examcode", util.TokenValidationMiddlewareNew(client), h.QueryExamCityNamesForIPExam(client))          //hari
	r.GET(apiSubPath+"RegionHeadQuarters/Circle/:id2", h.QueryRegionHeadQuartersByExamConductedBy(client))

	r.GET(apiSubPath+"candidate/getcandidateuserbyempid/:employeeid", util.TokenValidationMiddlewareNew(client), h.GetCandidateUsersByEmpId(client)) //--vasanth dev
	r.GET(apiSubPath+"center/get/:id", util.TokenValidationMiddlewareNew(client), h.GetCenterID(client))
	r.GET(apiSubPath+"centers/:id1/:id2/:id3/:id4", util.TokenValidationMiddlewareNew(client), h.GetExamCentresBynodalOfficeIDExamCode(client))
	r.GET(apiSubPath+"centersByCircleID/:id", util.TokenValidationMiddlewareNew(client), h.GetCentersByCircleID(client))
	r.GET(apiSubPath+"centersByConductingAuthority/:examyear/:examcode/:conductingid", util.TokenValidationMiddlewareNew(client), h.GetCentersByConductingAuthority(client))
	r.GET(apiSubPath+"centersBycity/:id", util.TokenValidationMiddlewareNew(client), h.GetCentersByCity(client))
	r.GET(apiSubPath+"employeecategories", util.TokenValidationMiddlewareNew(client), h.GetAllEmployeeCategories(client))
	r.GET(apiSubPath+"employeedesignation", util.TokenValidationMiddlewareNew(client), h.GetAllEmployeeDesignations(client))
	r.GET(apiSubPath+"employeedposts", util.TokenValidationMiddlewareNew(client), h.GetAllEmployeePosts(client))
	r.GET(apiSubPath+"employeedposts/:id", util.TokenValidationMiddlewareNew(client), h.GetEmpPostsByID(client))
	r.GET(apiSubPath+"employees", util.TokenValidationMiddlewareNew(client), h.GetAllEmployees(client))
	r.GET(apiSubPath+"employees/:id", util.TokenValidationMiddlewareNew(client), h.GetEmployeesByID(client))
	r.GET(apiSubPath+"employees/search/byempid/:employeeid", util.TokenValidationMiddlewareNew(client), h.GetEmployeeDetailsByEmpId(client))
	r.GET(apiSubPath+"empmaster/GetEmployeesBasedOnCA/:cafacilityid", util.TokenValidationMiddlewareNew(client), h.GetEmployeesBasedOnCA(client))
	r.GET(apiSubPath+"empmaster/viewemployee/:employeeid", util.TokenValidationMiddlewareNew(client), h.ViewEmployeeMaster(client))
	r.GET(apiSubPath+"examcalendars", util.TokenValidationMiddlewareNew(client), h.GetExamCalendars(client))
	r.GET(apiSubPath+"examcalendars/:id", util.TokenValidationMiddlewareNew(client), h.GetExamCalendarID(client))
	r.GET(apiSubPath+"exampapers", util.TokenValidationMiddlewareNew(client), h.GetExamPapers(client))
	r.GET(apiSubPath+"exampapers/byexamcode/:id", util.TokenValidationMiddlewareNew(client), h.GetExamPaperTypesWithPaperCode(client))
	r.GET(apiSubPath+"exampapertypes", util.TokenValidationMiddlewareNew(client), h.GetAllExamPaperTypes(client))
	r.GET(apiSubPath+"exampapertypes/:id", util.TokenValidationMiddlewareNew(client), h.GetExamPaperTypesByID(client))
	r.GET(apiSubPath+"exampapertypes/bypapercode/:id", util.TokenValidationMiddlewareNew(client), h.GetExamPaperTypesWithPaperCode(client))

	r.GET(apiSubPath+"exams/statistics/fordirectorate/circlewise/get/:examcode/:examyear", util.TokenValidationMiddlewareNew(client), h.GetExamStatisticsForDirectorateCircleWise(client))
	r.GET(apiSubPath+"exams/statistics/fordirectorate/get/:examcode/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetExamApplicationsStatisticsForDirectorate(client))
	r.GET(apiSubPath+"exams/statistics/get/:examcode/:nofacilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetExamApplicationsStatisticsForNO(client))
	r.GET(apiSubPath+"exams/statistics/officewise/get/:examcode/:nofacilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetExamApplicationsStatisticsOfficeWise(client))
	r.GET(apiSubPath+"facilities/byPincode/:pincode", util.TokenValidationMiddlewareNew(client), h.QueryOfficeByPincodehandle(client))
	r.GET(apiSubPath+"facilities/byPincodeNoUID/:pincode", util.TokenValidationMiddlewareUID(client), h.QueryOfficeByPincodehandle(client))
	r.GET(apiSubPath+"facilities/byfacilityofficeid/:workingofficefacilityid", util.TokenValidationMiddlewareNew(client), h.GetFacilityDetailsByFacilityOfficeID(client))

	r.GET(apiSubPath+"getExamApplicationsByCityPref/:examYear/:examcode/:cityid", util.TokenValidationMiddlewareNew(client), h.GetExamApplicatonsPreferenenceCityWiseStats(client))
	r.GET(apiSubPath+"getExamApplicationsByCityPrefGDSPA/:examYear/:examcode/:cityid", util.TokenValidationMiddlewareNew(client), h.GetExamApplicatonsPreferenenceCityWiseStatsGDSPA(client))
	r.GET(apiSubPath+"getExamApplicationsByCityPrefGDSPM/:examYear/:examcode/:cityid", util.TokenValidationMiddlewareNew(client), h.GetExamApplicatonsPreferenenceCityWiseStatsGDSPM(client))
	r.GET(apiSubPath+"getExamApplicationsByCityPrefMTSPMMG/:examYear/:examcode/:cityid", util.TokenValidationMiddlewareNew(client), h.GetExamApplicatonsPreferenenceCityWiseStatsMTSPMMG(client))
	r.GET(apiSubPath+"getExamApplicationsByCityPrefPMPA/:examYear/:examcode/:cityid", util.TokenValidationMiddlewareNew(client), h.GetExamApplicatonsPreferenenceCityWiseStatsPMPA(client))
	r.GET(apiSubPath+"getExamApplicationsByCityPrefPS/:examYear/:examcode/:cityid", util.TokenValidationMiddlewareNew(client), h.GetExamApplicatonsPreferenenceCityWiseStatsPS(client))
	r.GET(apiSubPath+"getallpdfs/:code/:year", util.TokenValidationMiddlewareNew(client), h.GetPdf(client))
	r.GET(apiSubPath+"nodalofficerbyusername/:id", util.TokenValidationMiddlewareNew(client), h.GetNodalOfficerByUsername(client))

	r.GET(apiSubPath+"getExamApplicationsByCenterIP/:examYear/:examcode/:centerid/:startno/:endno", util.TokenValidationMiddlewareNew(client), h.GetExamApplicationsByCenterIP(client))

	r.GET(apiSubPath+"roles", util.TokenValidationMiddlewareNew(client), h.GetRoles(client)) //dev by vasanth

	r.GET(apiSubPath+"vacancyyears", util.TokenValidationMiddlewareNew(client), h.GetVacancyYears(client))
	r.GET(apiSubPath+"vacancyyears/:id", util.TokenValidationMiddlewareNew(client), h.GetVacancyYearID(client))

	r.POST(apiSubPath+"CreateExamCityCenters", util.TokenValidationMiddlewareNew(client), h.CreateExamCityCenters(client)) //--created by vasanth

	r.POST(apiSubPath+"candidatecreate", util.TokenValidationMiddlewareNew(client), h.CreatecandidateAUser(client)) //
	r.POST(apiSubPath+"center/submit", util.TokenValidationMiddlewareNew(client), h.CreateExamCenter(client))

	r.POST(apiSubPath+"employee/updateprofile", util.TokenValidationMiddlewareNew(client), h.UpdateUserDetails(client))
	r.POST(apiSubPath+"employeecategories", util.TokenValidationMiddlewareNew(client), h.CreateEmployeeCategories(client))
	r.POST(apiSubPath+"employeedesignation", util.TokenValidationMiddlewareNew(client), h.CreateEmployeeDesignations(client))
	r.POST(apiSubPath+"employeedposts", util.TokenValidationMiddlewareNew(client), h.CreateEmployeePost(client))
	r.POST(apiSubPath+"employeeprofile/verify/:id", util.TokenValidationMiddlewareNew(client), h.Updateemployeeverifydetails(client))
	r.POST(apiSubPath+"employees", util.TokenValidationMiddlewareNew(client), h.CreateEmployee(client))
	r.POST(apiSubPath+"empmaster/createemployee", util.TokenValidationMiddlewareNew(client), h.CreateEmployeeMaster(client)) //---K.Mohandoss

	r.POST(apiSubPath+"examcalendars", util.TokenValidationMiddlewareNew(client), h.CreateExamCalendar(client))
	r.POST(apiSubPath+"exampapers", util.TokenValidationMiddlewareNew(client), h.InsertExamPapers(client))
	r.POST(apiSubPath+"exampapertypes", util.TokenValidationMiddlewareNew(client), h.InsertExamPaperTypes(client))

	//r.POST(apiSubPath+"profile", util.TokenValidationMiddlewareNew(client), h.CreateUser(client))           //NL
	r.POST(apiSubPath+"profileadmin", util.TokenValidationMiddlewareNew(client), h.CreateAdminUser(client)) //---Vasanth

	//r.POST(apiSubPath+"users/bulkinsert/", util.TokenValidationMiddlewareNew(client), h.InsertBulkUsersHandler(client))
	r.GET(apiSubPath+"getAllCircles", util.TokenValidationMiddlewareUID(client), h.GetAllCircles(client))
	r.POST(apiSubPath+"empmaster/submitemployee", util.TokenValidationMiddlewareUID(client), h.CreateEmployeeMaster(client)) //---K.Mohandoss
	r.POST(apiSubPath+"empmaster/smsotp", util.TokenValidationMiddlewareUID(client), h.TriggerSMSOTP(client))                //---K.Mohandoss
	r.PUT(apiSubPath+"empmaster/emailotp", util.TokenValidationMiddlewareUID(client), h.TriggerEmailOTP(client))             //---K.Mohandoss
	r.PUT(apiSubPath+"empmaster/verifysmsotp", util.TokenValidationMiddlewareUID(client), h.VerifyTriggerSMSOTP(client))     //---K.Mohandoss
	r.PUT(apiSubPath+"empmaster/verifyemailotp", util.TokenValidationMiddlewareUID(client), h.VerifyTriggerEmailOTP(client)) //---K.Mohandoss
	r.POST(apiSubPath+"users/login", util.TokenValidationMiddlewareUID(client), h.VerifyUserLogin(client))
	r.POST(apiSubPath+"users/new/submit", util.TokenValidationMiddlewareUID(client), h.FirstTimeUserCreation(client))
	r.POST(apiSubPath+"users/new/update", util.TokenValidationMiddlewareUID(client), h.UpdateFirstTimeUserDetails(client))
	r.POST(apiSubPath+"users/reset/savenewpassword", util.TokenValidationMiddlewareUID(client), h.UserResetSaveNewPassword(client))
	r.POST(apiSubPath+"users/reset/validateOTP", util.TokenValidationMiddlewareUID(client), h.UserResetValidateOTP(client))
	r.POST(apiSubPath+"users/reset/validateusername", util.TokenValidationMiddlewareUID(client), h.UserResetValidateUserName(client))
	r.POST(apiSubPath+"users/trigotp/change/emailmobile", util.TokenValidationMiddlewareNew(client), h.EditMobileorEmailCredentials(client))
	r.POST(apiSubPath+"users/validate/otp", util.TokenValidationMiddlewareNew(client), h.ValidateOTPUserEmailMobile(client))
	r.POST(apiSubPath+"vacancyyears", util.TokenValidationMiddlewareNew(client), h.CreateVacarncyYears(client))

	r.GET(apiSubPath+"GDSPAexams/applications/getbyempid/:empid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPAApplicationsByEmpId(client))
	r.GET(apiSubPath+"GDSPAexams/caprevremarks/:empid", util.TokenValidationMiddlewareNew(client), h.GetGDSPACAPendingOldRemarksByEmpId(client))
	r.GET(apiSubPath+"GDSPAexams/getAllPendingWithCandidate/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetAllGDSPAPendingWithCandidate(client))
	r.GET(apiSubPath+"GDSPAexams/getallcapending/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPACAPendingDetailsByEmpId(client))
	r.GET(apiSubPath+"GDSPAexams/getallcapendingapplications/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetAllGDSPACAPendingVerifications(client))
	r.GET(apiSubPath+"GDSPAexams/getallcaverified/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPACAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"GDSPAexams/getallcaverifiedapplications/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetAllGDSPACAVerified(client))
	r.GET(apiSubPath+"GDSPAexams/getallcaverifiedapplicationsforna/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPAAllCAVerifiedForNA(client))
	r.GET(apiSubPath+"GDSPAexams/getallnaverified/:empid", util.TokenValidationMiddlewareNew(client), h.GetGDSPANAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"GDSPAexams/getallnaverifiedapplications/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPAAllNAVerified(client))
	r.GET(apiSubPath+"GDSPAexams/getallnaverifiedapplicationsforna/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPAAllNAVerifiedForNA(client))
	r.GET(apiSubPath+"GDSPAexams/hallticket/get/:examcode/:empid", util.TokenValidationMiddlewareNew(client), h.GetGDSPAHallTicketWithExamCodeEmpID(client))
	r.GET(apiSubPath+"GDSPAexams/recommendations/:empid", util.TokenValidationMiddlewareNew(client), h.GetGDSPAExamRecommendationsByEmpId(client))
	r.POST(apiSubPath+"GDSPAexams/applications/submit", util.TokenValidationMiddlewareNew(client), h.CreateNewGDSPAApplications(client))

	r.PUT(apiSubPath+"GDSPAexams/applications/resubmit", util.TokenValidationMiddlewareNew(client), h.ResubmitGDSPAApplication(client))
	r.PUT(apiSubPath+"GDSPAexams/applications/verify", util.TokenValidationMiddlewareNew(client), h.VerifyGDSPAApplication(client))
	r.PUT(apiSubPath+"GDSPAexams/center/updatecenters", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInGDSPAApplsreturnstring(client))
	r.PUT(apiSubPath+"GDSPAexams/noverify", util.TokenValidationMiddlewareNew(client), h.UpdateNodalRecommendationsGDSPAByEmpID(client))
	r.PUT(apiSubPath+"GdsPaexams/Halltickets", util.TokenValidationMiddlewareNew(client), h.GenerateHallticketNumbersGdspa(client)) // vasanth

	r.GET(apiSubPath+"GDSPMexams/applications/getbyempid/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPMApplicationsByEmpId(client))
	r.GET(apiSubPath+"GDSPMexams/caprevremarks/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPMCAPendingOldRemarksByEmpId(client))
	r.GET(apiSubPath+"GDSPMexams/getAllPendingWithCandidate/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllGDSPMPendingWithCandidate(client))
	r.GET(apiSubPath+"GDSPMexams/getallcapending/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPMCAPendingDetailsByEmpId(client))
	r.GET(apiSubPath+"GDSPMexams/getallcapendingapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllGDSPMCAPendingVerifications(client))
	r.GET(apiSubPath+"GDSPMexams/getallcaverified/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPMCAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"GDSPMexams/getallcaverifiedapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllGDSPMCAVerified(client))
	r.GET(apiSubPath+"GDSPMexams/getallcaverifiedapplicationsforna/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPMAllCAVerifiedForNA(client))
	r.GET(apiSubPath+"GDSPMexams/getallnaverified/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPMNAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"GDSPMexams/getallnaverifiedapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPMAllNAVerified(client))
	r.GET(apiSubPath+"GDSPMexams/getallnaverifiedapplicationsforna/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPMAllNAVerifiedForNA(client))
	r.GET(apiSubPath+"GDSPMexams/getallvapendingapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllGDSPMVAPendingVerifications(client))
	r.GET(apiSubPath+"GDSPMexams/hallticket/get/:examcode/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetGDSPMHallTicketWithExamCodeEmpID(client))
	r.GET(apiSubPath+"GDSPMexams/recommendations/:employeeid", util.TokenValidationMiddlewareNew(client), h.GetGDSPMExamRecommendationsByEmpId(client))

	r.POST(apiSubPath+"GDSPMexams/Applications/Submit", util.TokenValidationMiddlewareNew(client), h.CreateNewGDSPMApplicationss(client)) //dev by vasanth
	r.PUT(apiSubPath+"GDSPMexams/Applications/resubmit", util.TokenValidationMiddlewareNew(client), h.ResubmitGDSPMApplications(client))  //dev by vasanth
	r.PUT(apiSubPath+"GDSPMexams/Halltickets", util.TokenValidationMiddlewareNew(client), h.GenerateHallticketNumbersGDSPM(client))       // vasanth
	r.PUT(apiSubPath+"GDSPMexams/applications/Verify", util.TokenValidationMiddlewareNew(client), h.VerifyGDSPMApplications(client))      //dev by vasanth
	r.PUT(apiSubPath+"GDSPMexams/applications/vaVerify", util.TokenValidationMiddlewareNew(client), h.VerifyGDSPMVAApplications(client))  //dev by vasanth
	r.PUT(apiSubPath+"GDSPMexams/center/updatecenters", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInGDSPMApplsreturnstring(client))
	r.PUT(apiSubPath+"GDSPMexams/noverify", util.TokenValidationMiddlewareNew(client), h.UpdateNodalRecommendationsGDSPMByEmpID(client))

	r.GET(apiSubPath+"MTSPMMGexams/applications/getbyempid/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetMTSPMMGApplicationsByEmpId(client))
	r.GET(apiSubPath+"MTSPMMGexams/caprevremarks/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetMTSPMMGCAPendingOldRemarksByEmpId(client)) //NEEDD TO BE DONE
	r.GET(apiSubPath+"MTSPMMGexams/getAllPendingWithCandidate/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllMTSPMMGPendingWithCandidate(client))
	r.GET(apiSubPath+"MTSPMMGexams/getallcapending/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetMTSPMMGCAPendingDetailsByEmpId(client))
	r.GET(apiSubPath+"MTSPMMGexams/getallcapendingapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllMTSPMMGCAPendingVerifications(client))
	r.GET(apiSubPath+"MTSPMMGexams/getallcaverified/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetMTSPMMGCAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"MTSPMMGexams/getallcaverifiedapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllMTSPMMGCAVerified(client))
	r.GET(apiSubPath+"MTSPMMGexams/getallcaverifiedapplicationsforna/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.MtspmAllCAVerifiedForNA(client))
	r.GET(apiSubPath+"MTSPMMGexams/getallnaverifiedapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetMTSPMMGAllNAVerified(client))
	r.GET(apiSubPath+"MTSPMMGexams/getallnaverifiedapplicationsforna/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetMTSPMMGAllNAVerifiedForNA(client))
	r.GET(apiSubPath+"MTSPMMGexams/hallticket/get/:examcode/:exmployeeid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetMTSPMMGHallTicketWithExamCodeEmpID(client))
	r.GET(apiSubPath+"MTSPMMGexams/recommendations/:employeeid", util.TokenValidationMiddlewareNew(client), h.GetMTSPMMGExamRecommendationsByEmpId(client))

	r.POST(apiSubPath+"MTSPMMGexams/applications/submit", util.TokenValidationMiddlewareNew(client), h.CreateNewMTSPMMGApplications(client))
	r.PUT(apiSubPath+"MTSPMMGexams/applications/resubmit", util.TokenValidationMiddlewareNew(client), h.ResubmitMTSPMMGApplication(client))
	r.PUT(apiSubPath+"MTSPMMGexams/applications/verify", util.TokenValidationMiddlewareNew(client), h.VerifyMTSPMMGApplication(client))
	r.PUT(apiSubPath+"MTSPMMGexams/center/updatecenters", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInMTSPMMGApplsreturnstring(client)) //LOG NEED TO BE ADDED
	r.PUT(apiSubPath+"MTSPMMGexams/noverify", util.TokenValidationMiddlewareNew(client), h.UpdateNodalRecommendationsMTSPMMGByEmpID(client))
	r.PUT(apiSubPath+"MTSPMMGexams/Halltickets", util.TokenValidationMiddlewareNew(client), h.GenerateHallticketNumbersMtsPm(client)) // vasanth

	r.GET(apiSubPath+"PMPAexams/applications/getbyempid/:id/:id1", util.TokenValidationMiddlewareNew(client), h.GetPMPAApplicationsByEmpId(client))
	r.GET(apiSubPath+"PMPAexams/caprevremarks/:id", util.TokenValidationMiddlewareNew(client), h.GetPMPACAPendingOldRemarksByEmpId(client))
	r.GET(apiSubPath+"PMPAexams/getAllPendingWithCandidate/:id/:id1", util.TokenValidationMiddlewareNew(client), h.GetAllPMPAPendingWithCandidate(client))
	r.GET(apiSubPath+"PMPAexams/getallcapending/:id", util.TokenValidationMiddlewareNew(client), h.GetPMPACAPendingDetailsByEmpId(client))
	r.GET(apiSubPath+"PMPAexams/getallcapendingapplications/:id/:id1", util.TokenValidationMiddlewareNew(client), h.GetAllPMPACAPendingVerifications(client))
	r.GET(apiSubPath+"PMPAexams/getallcaverified/:id", util.TokenValidationMiddlewareNew(client), h.GetPMPACAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"PMPAexams/getallcaverifiedapplications/:id/:id1", util.TokenValidationMiddlewareNew(client), h.GetAllPMPACAVerified(client))
	r.GET(apiSubPath+"PMPAexams/getallcaverifiedapplicationsforna/:id/:id1", util.TokenValidationMiddlewareNew(client), h.GetPMPAAllCAVerifiedForNA(client))
	r.GET(apiSubPath+"PMPAexams/getallnaverified/:id", util.TokenValidationMiddlewareNew(client), h.GetPMPANAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"PMPAexams/getallnaverifiedapplications/:id/:id1", util.TokenValidationMiddlewareNew(client), h.GetPMPAAllNAVerified(client))
	r.GET(apiSubPath+"PMPAexams/getallnaverifiedapplicationsforna/:id/:id1", util.TokenValidationMiddlewareNew(client), h.GetPMPAAllNAVerifiedForNA(client))
	r.GET(apiSubPath+"PMPAexams/hallticket/get/:id1/:id2", util.TokenValidationMiddlewareNew(client), h.GetPMPAHallTicketWithExamCodeEmpID(client))
	r.GET(apiSubPath+"PMPAexams/recommendations/:id", util.TokenValidationMiddlewareNew(client), h.GetPMPAExamRecommendationsByEmpId(client))

	r.POST(apiSubPath+"PMPAexams/applications/submit", util.TokenValidationMiddlewareNew(client), h.CreateNewPMPAApplications(client))

	r.PUT(apiSubPath+"PMPAexams/applications/resubmit", util.TokenValidationMiddlewareNew(client), h.ResubmitPMPAApplication(client))
	r.PUT(apiSubPath+"PMPAexams/applications/verify", util.TokenValidationMiddlewareNew(client), h.VerifyPMPAApplication(client))
	r.PUT(apiSubPath+"PMPAexams/center/updatecenters", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInPMPAApplsreturnstring(client))
	r.PUT(apiSubPath+"PMPAexams/halltickets", util.TokenValidationMiddlewareNew(client), h.GetPMPAHallticketNumberscenter(client)) // just returns a string message generated successfully.
	r.PUT(apiSubPath+"PMPAexams/noverify", util.TokenValidationMiddlewareNew(client), h.UpdateNodalRecommendationsPMPAByEmpID(client))
	r.PUT(apiSubPath+"PMPAexams/Halltickets", util.TokenValidationMiddlewareNew(client), h.GenerateHallticketNumbersPmPa(client)) // vasanth

	r.DELETE(apiSubPath+"admin/deleteadminuserbyempid/:username", util.TokenValidationMiddlewareNew(client), h.DeleteAdminUsersByUsername(client)) //--dev by vassanth

	r.GET(apiSubPath+"admin/getadminuserbyempid/:id", util.TokenValidationMiddlewareNew(client), h.GetAdminUsersByEmpId(client))                                   //--vasanth
	r.GET(apiSubPath+"admin/getadminuserbyfacilityrole/:facilityid/:role", util.TokenValidationMiddlewareNew(client), h.GetAllAdminUsersByFacilityandRole(client)) //--vasanth
	r.GET(apiSubPath+"admin/getadminuserbyusername/:id", util.TokenValidationMiddlewareNew(client), h.GetAdminUsersByUsername(client))
	r.GET(apiSubPath+"admin/getempmasteruserbyempid/:employeeid", util.TokenValidationMiddlewareNew(client), h.GetEmpMasterUsersByEmpId(client)) //--vasanth
	r.GET(apiSubPath+"admin/roles", util.TokenValidationMiddlewareNew(client), h.GetAllRoles(client))
	r.GET(apiSubPath+"admin/roles/:id", util.TokenValidationMiddlewareNew(client), h.GetRolesByID(client))

	r.POST(apiSubPath+"admins/reset/savenewpassword", util.TokenValidationMiddlewareUID(client), h.AdminResetSaveNewPassword(client))
	r.POST(apiSubPath+"admins/reset/validateOTP", util.TokenValidationMiddlewareUID(client), h.AdminResetValidateOTP(client))
	r.POST(apiSubPath+"admins/reset/validateusername", util.TokenValidationMiddlewareUID(client), h.AdminResetValidateUserName(client))
	r.POST(apiSubPath+"adminusers/verifyloginn", util.TokenValidationMiddlewareUID(client), h.VerifyAdminLoginn(client)) //--vasanth dev

	r.POST(apiSubPath+"adminusers/verifyuserr", util.TokenValidationMiddlewareUID(client), h.VerifyAdminUserLoginn(client)) //--vasanth dev

	r.PUT(apiSubPath+"admin/ChangeAdminPassword", util.TokenValidationMiddlewareNew(client), h.ChangeAdminPassword(client))  //----vasanth dev
	r.PUT(apiSubPath+"admin/geterrorlogs/:adminusername", util.TokenValidationMiddlewareNew(client), h.GetErrorLogs(client)) //--vasanth

	r.PUT(apiSubPath+"candidate/updatecandidateuserbyempid/:id", util.TokenValidationMiddlewareNew(client), h.UpdateCandidateUsersByEmpId(client)) //--vasanth dev
	r.PUT(apiSubPath+"center/update/:id", util.TokenValidationMiddlewareNew(client), h.UpdateExamCenter(client))
	r.PUT(apiSubPath+"center/updatecenters", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInIPApplsreturnstring(client))            //NL
	r.PUT(apiSubPath+"center/updatecenters/GDSPA", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInGDSPAApplsreturnstring(client))   //NL
	r.PUT(apiSubPath+"center/updatecenters/GDSPM", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInGDSPMApplsreturnstring(client))   //NL
	r.PUT(apiSubPath+"center/updatecenters/IP", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInIPApplsreturnstring(client))         //NL
	r.PUT(apiSubPath+"center/updatecenters/MTSPMMG", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInGDSPMApplsreturnstring(client)) //NL
	r.PUT(apiSubPath+"center/updatecenters/PMPA", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInPMPAApplsreturnstring(client))     //NL
	r.PUT(apiSubPath+"center/updatecenters/PS", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInPSApplsreturnstring(client))         //NL
	r.PUT(apiSubPath+"empmaster/modifyemployee", util.TokenValidationMiddlewareNew(client), h.ModifyEmployeeMaster(client))

	r.GET(apiSubPath+"IpExams/GetDetailsByOnExamCode/:Exam_Code", util.TokenValidationMiddlewareNew(client), h.GetDetailsBasedOnExamCode(client)) //----vasanth
	r.GET(apiSubPath+"ipexams/applications/getbyempid/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetIPApplicationsByEmpId(client))
	r.GET(apiSubPath+"ipexams/caprevremarks/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetCAPendingOldRemarksByEmpId(client))
	r.GET(apiSubPath+"ipexams/getAllPendingWithCandidate/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllPendingWithCandidate(client))
	r.GET(apiSubPath+"ipexams/getallcapending/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetCAPendingDetailsByEmpId(client))
	r.GET(apiSubPath+"ipexams/getallcapendingapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllCAPendingVerifications(client))
	r.GET(apiSubPath+"ipexams/getallapplicationsondeputation/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllPendingApplicationsOnDeputation(client))
	r.GET(apiSubPath+"ipexams/getallcaverified/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetCAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"ipexams/getallcaverifiedapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllCAVerified(client))
	r.GET(apiSubPath+"ipexams/getallcaverifiedapplicationsforna/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllCAVerifiedForNA(client))
	r.GET(apiSubPath+"ipexams/getallnaverified/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetNAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"ipexams/getallnaverifiedapplications/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllNAVerified(client))
	r.GET(apiSubPath+"ipexams/getallnaverifiedapplicationsforna/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllNAVerifiedForNA(client))
	r.GET(apiSubPath+"ipexams/hallticket/get/:examname/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetHallTicketWithExamCodeEmpID(client))
	r.GET(apiSubPath+"ipexams/recommendations/:employeeid", util.TokenValidationMiddlewareNew(client), h.GetIPExamRecommendationsByEmpId(client))
	r.GET(apiSubPath+"ipexams/examcenterhall/:cityid/:examyear/:examcode/:centerid", util.TokenValidationMiddlewareNew(client), h.GetIPExamcenterhallBycityid(client))                    //---vasanth
	r.GET(apiSubPath+"ipexams/checkexamcenterhall/:examyear/:examcode/:nafacilityid/:cofacilityid/:hallname", util.TokenValidationMiddlewareNew(client), h.CheckIPExamcenterhall(client)) //---vasanth

	r.GET(apiSubPath+"ipexams/candiadatesexamcenterhall/:cityid/:examyear/:examcode/:centerid/:hallname", util.TokenValidationMiddlewareNew(client), h.GetIPCandiadatesExamcenterhallBycityid(client)) //---vasanth
	r.POST(apiSubPath+"ipexams/ExamCenterHall/create", util.TokenValidationMiddlewareNew(client), h.CreateExamCenterHall(client))                                                                      //---vasanth
	r.POST(apiSubPath+"ipexams/applications/submit", util.TokenValidationMiddlewareNew(client), h.CreateNewIPApplications(client))                                                                     //---vasanth
	r.PUT(apiSubPath+"ipexams/applications/resubmit", util.TokenValidationMiddlewareNew(client), h.ResubmitApplication(client))
	r.PUT(apiSubPath+"ipexams/ExamCenterHall/reset", util.TokenValidationMiddlewareNew(client), h.ResetExamCenterHall(client)) //---vasanth
	r.PUT(apiSubPath+"ipexams/applications/verify", util.TokenValidationMiddlewareNew(client), h.VerifyIPApplication(client))
	r.PUT(apiSubPath+"ipexams/applications/nareset", util.TokenValidationMiddlewareNew(client), h.ResetIPApplicationNA(client))
	r.PUT(apiSubPath+"ipexams/applications/caedit", util.TokenValidationMiddlewareNew(client), h.IPApplicationCAEdit(client))
	r.PUT(apiSubPath+"ipexams/applications/naresetcenter", util.TokenValidationMiddlewareNew(client), h.ResetCenterIPApplicationNA(client))

	r.PUT(apiSubPath+"ipexams/halltickets/:examyear", util.TokenValidationMiddlewareNew(client), h.GenerateHallticketNumberscenter(client)) // just returns a string message generated successfully.
	r.PUT(apiSubPath+"ipexams/noverify", util.TokenValidationMiddlewareNew(client), h.UpdateNodalRecommendationsIPByEmpID(client))
	r.PUT(apiSubPath+"Ipexams/Halltickets", util.TokenValidationMiddlewareNew(client), h.GenerateHallticketNumbers(client)) // vasanth
	r.PUT(apiSubPath+"ipexams/pendingwithcandiate/:nodalofficeid/:examyear", util.TokenValidationMiddlewareNew(client), h.EmailSmsTriggeringPenidngWithCandidate(client))

	r.GET(apiSubPath+"notification/AllINByUserName/:id", util.TokenValidationMiddlewareNew(client), h.GetallIssuedexamNotificationBySuUserName(client))
	r.GET(apiSubPath+"notification/GetAllIssuednotificationByYear/:id", util.TokenValidationMiddlewareNew(client), h.GetAllIssuedexamNotificationbyYear(client))
	r.GET(apiSubPath+"notification/GetDraftNotification/:examid/:examyear/:officetype/:facilityid", util.TokenValidationMiddlewareNew(client), h.GetDraftNotification(client))
	r.GET(apiSubPath+"notification/GetNotificationRemarks/:id", util.TokenValidationMiddlewareNew(client), h.GetRemarksofNotificationByNotificationID(client))
	r.GET(apiSubPath+"notification/GetnotificationByYear/:id1/:id2", util.TokenValidationMiddlewareNew(client), h.GetexamNotificationbyYear(client))
	r.GET(apiSubPath+"notification/IndividualINByNotificationNumber/:id", util.TokenValidationMiddlewareNew(client), h.GetIndividualIssuedexamNotificationByNotificationNumber(client))
	r.GET(apiSubPath+"notification/IndividulaDNByUserName/:id", util.TokenValidationMiddlewareNew(client), h.GetIndividualDraftNotificationbySuUserName(client))
	r.GET(apiSubPath+"notification/PendingDNByUserName/:id", util.TokenValidationMiddlewareNew(client), h.GetDraftexamNotificationBySuUserName(client))
	r.GET(apiSubPath+"notification/PendingReDNByUserName/:id", util.TokenValidationMiddlewareNew(client), h.GetReDraftexamNotificationBySuUserName(client))
	r.GET(apiSubPath+"notification/getPNDnotifications/:circleid", util.TokenValidationMiddlewareNew(client), h.GetPNDNotifications(client))                 //4
	r.GET(apiSubPath+"notification/getallnotificationsmax/:examyear/:circleid", util.TokenValidationMiddlewareNew(client), h.GetAllNotificationsMax(client)) //2
	r.POST(apiSubPath+"notification/create", util.TokenValidationMiddlewareNew(client), h.CreateexamNotification(client))

	r.PUT(apiSubPath+"notification/PutIssueNotification", util.TokenValidationMiddlewareNew(client), h.PutIssueNotification(client))
	r.PUT(apiSubPath+"notification/PutIssueNotificationSingle", util.TokenValidationMiddlewareNew(client), h.PutIssueNotificationSingle(client))
	r.PUT(apiSubPath+"notification/PutIssueNotificationWithEditing", util.TokenValidationMiddlewareNew(client), h.PutIssueNotification(client))
	r.PUT(apiSubPath+"notification/PutReIssueNotification", util.TokenValidationMiddlewareNew(client), h.PutReIssueNotification(client))
	r.PUT(apiSubPath+"notification/PutResubmitDraftNotification", util.TokenValidationMiddlewareNew(client), h.PutResubmitDraftNotification(client))
	r.PUT(apiSubPath+"notification/UpdateResubmitDraftNotification", util.TokenValidationMiddlewareNew(client), h.UpdateResubmitDraftNotificationNew(client)) // 1
	r.PUT(apiSubPath+"notification/cancel", util.TokenValidationMiddlewareNew(client), h.CancelDraftNotification(client))
	r.PUT(apiSubPath+"notification/cancel/specific/circle", util.TokenValidationMiddlewareNew(client), h.CancelDraftNotificationForSpecificCircle(client)) //3  //3

	r.GET(apiSubPath+"psexams/applications/getbyempid/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetPSApplicationsByEmpId(client))
	r.GET(apiSubPath+"psexams/caprevremarks/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetPSCAPendingOldRemarksByEmpId(client))
	r.GET(apiSubPath+"psexams/getAllPSPendingWithCandidate/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetAllPSPendingWithCandidate(client))
	r.GET(apiSubPath+"psexams/getallcapending/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetPSCAPendingDetailsByEmpId(client))
	r.GET(apiSubPath+"psexams/getallcapendingapplications/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetAllPSCAPendingVerifications(client))
	r.GET(apiSubPath+"psexams/getallcaverified/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetPSCAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"psexams/getallcaverifiedapplications/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetAllPSCAVerified(client))
	r.GET(apiSubPath+"psexams/getallcaverifiedapplicationsforna/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetPSAllCAVerifiedForNA(client))
	r.GET(apiSubPath+"psexams/getallnaverified/:employeeid/examyear", util.TokenValidationMiddlewareNew(client), h.GetPSNAVerifiedDetailsByEmpId(client))
	r.GET(apiSubPath+"psexams/getAllPendingWithCandidate/:facilityid/:selectedyear", util.TokenValidationMiddlewareNew(client), h.GetAllPendingWithCandidate(client))
	r.GET(apiSubPath+"psexams/getallnaverifiedapplications/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetPSAllNAVerified(client))
	r.GET(apiSubPath+"psexams/getallnaverifiedapplicationsforna/:facilityid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetPSAllNAVerifiedForNA(client))
	r.GET(apiSubPath+"psexams/hallticket/get/:examname/:employeeid/:examyear", util.TokenValidationMiddlewareNew(client), h.GetHallTicketWithExamCodeEmpID(client))
	r.GET(apiSubPath+"psexams/recommendations/:employeeid", util.TokenValidationMiddlewareNew(client), h.GetPSExamRecommendationsByEmpId(client))

	r.POST(apiSubPath+"psexams/applications/submit", util.TokenValidationMiddlewareNew(client), h.CreateNewPSApplications(client))

	r.PUT(apiSubPath+"psexams/applications/resubmit", util.TokenValidationMiddlewareNew(client), h.ResubmitPSApplication(client))
	r.PUT(apiSubPath+"psexams/applications/verify", util.TokenValidationMiddlewareNew(client), h.VerifyPSApplication(client))
	r.PUT(apiSubPath+"psexams/center/updatecenters", util.TokenValidationMiddlewareNew(client), h.UpdateExamCentersInPSApplsreturnstring(client))
	r.PUT(apiSubPath+"psexams/halltickets", util.TokenValidationMiddlewareNew(client), h.GetPSHallticketNumberscenter(client)) // just returns a string message generated successfully.
	r.PUT(apiSubPath+"psexams/noverify", util.TokenValidationMiddlewareNew(client), h.UpdateNodalRecommendationsPSByEmpID(client))

	r.PUT(apiSubPath+"PSexams/Halltickets", util.TokenValidationMiddlewareNew(client), h.GenerateHallticketNumbersPs(client)) // vasanth

	r.PUT(apiSubPath+"updateCenterCodeForApplications", util.TokenValidationMiddlewareNew(client), h.UpdateCenterCodeForApplications(client))
	r.PUT(apiSubPath+"updateCenterCodeForApplicationsGDSPA", util.TokenValidationMiddlewareNew(client), h.UpdateCenterCodeForApplicationsGDSPA(client))
	r.PUT(apiSubPath+"updateCenterCodeForApplicationsGDSPM", util.TokenValidationMiddlewareNew(client), h.UpdateCenterCodeForApplicationsGDSPM(client))     //NL
	r.PUT(apiSubPath+"updateCenterCodeForApplicationsMTSPMMG", util.TokenValidationMiddlewareNew(client), h.UpdateCenterCodeForApplicationsMTSPMMG(client)) //NL
	r.PUT(apiSubPath+"updateCenterCodeForApplicationsPMPA", util.TokenValidationMiddlewareNew(client), h.UpdateCenterCodeForApplicationsPMPA(client))       //NL
	r.PUT(apiSubPath+"updateCenterCodeForApplicationsPS", util.TokenValidationMiddlewareNew(client), h.UpdateCenterCodeForApplicationsPS(client))           //NL

	r.GET(apiSubPath+"report/getAllExamDetails/:examcode/:examyear/:cirfaclilityid", util.TokenValidationMiddlewareNew(client), h.GetExamDetailsWithExamCodeExamYearCircleFacilityID(client))
	r.GET(apiSubPath+"report/getAllExamDetailsCA/:examcode/:examyear/:cafacilityid", util.TokenValidationMiddlewareNew(client), h.GetExamDetailsWithExamCodeExamYearCA(client))
	r.GET(apiSubPath+"report/getAllExamDetailsByApplicationStatus/:examcode/:examyear/:cirfaclilityid/:applnStatuscode", util.TokenValidationMiddlewareNew(client), h.GetExamDetailsWithExamCodeExamYearCircleFacilityIDByApplnstatus(client))
	r.GET(apiSubPath+"report/getAllExamDetailsByRecommendedStatus/:examcode/:examyear/:cirfaclilityid/:recommendationcode", util.TokenValidationMiddlewareNew(client), h.GetExamDetailsWithExamCodeExamYearCircleFacilityIDBystatus(client))
	r.GET(apiSubPath+"report/getAllExamDetailsHallticketView/:examcode/:examyear/:cirfaclilityid/:divfaclilityid", util.TokenValidationMiddlewareNew(client), h.GetExamDetailsHallticketview(client))
	r.GET(apiSubPath+"report/getAllExamDetailsHallticketViewDR/:examcode/:examyear", util.TokenValidationMiddlewareNew(client), h.GetExamDetailsHallticketviewDR(client))
	r.GET(apiSubPath+"report/getAllExamDetailsAttendanceView/:examcode/:examyear/:cirfaclilityid/:divfaclilityid", util.TokenValidationMiddlewareNew(client), h.GetExamDetailsAttendanceView(client))
	r.GET(apiSubPath+"report/getAllExamSummaryByApplicationsStatus/:examcode/:examyear/:facilityid/:entitytype", util.TokenValidationMiddlewareNew(client), h.GetExamSummaryByApplicationsStatussForCAandDT(client))
	r.GET(apiSubPath+"report/getAllExamSummaryByRecommendedStatus/:examcode/:examyear/:facilityid/:entitytype", util.TokenValidationMiddlewareNew(client), h.GetExamSummaryByRecommendedStatussForCAandDT(client))
	r.GET(apiSubPath+"report/getPendingApplicationsBasedOnDays/:examcode/:examyear/:facilityid/:daysPending", util.TokenValidationMiddlewareNew(client), h.GetPendingWithCaNaApplications(client))
	r.GET(apiSubPath+"report/getCandidatePendingApplicationsBasedOnDays/:examcode/:examyear/:facilityid/:daysPending", util.TokenValidationMiddlewareNew(client), h.GetPendingWithCandidateApplications(client))
	r.GET(apiSubPath+"report/getEmployeeMasterPendingWithCA/Circle/:facilityid", util.TokenValidationMiddlewareNew(client), h.GetEmployeeMasterPendingWithCA(client))
	r.GET(apiSubPath+"report/getEmployeeMasterPendingWithCA/Directorate", util.TokenValidationMiddlewareNew(client), h.GetEmployeeMasterPendingWithCADT(client))
	r.GET(apiSubPath+"report/getEmployeeMasterPendingWithNA/Directorate", util.TokenValidationMiddlewareNew(client), h.GetEmployeeMasterPendingWithNADT(client))
	r.GET(apiSubPath+"report/getCandidatePendingApplicationswithCA/Circle/:examcode/:examyear/:nodalfacilityid", util.TokenValidationMiddlewareNew(client), h.GetPendingWithCandidateApplicationsWithCA(client))
	r.GET(apiSubPath+"report/getCandidatePendingApplicationswithCA/Directorate/:examcode/:examyear", util.TokenValidationMiddlewareNew(client), h.GetPendingWithCandidateApplicationsWithCADT(client))
	r.GET(apiSubPath+"report/getCandidatePendingApplicationswithNA/Directorate/:examcode/:examyear", util.TokenValidationMiddlewareNew(client), h.GetPendingWithCandidateApplicationsWithNADT(client))

	r.GET(apiSubPath+"report/gettingresultbasedondates", util.TokenValidationMiddlewareNew(client), h.Gettingresultbasedondates(client))
	r.GET(apiSubPath+"report/getexamcitycenters/:examcode/:examyear", util.TokenValidationMiddlewareNew(client), h.GetExamCityCenters(client))
	r.GET(apiSubPath+"report/getexamcitydivisions/:examcode/:examyear/:nofacilityid", util.TokenValidationMiddlewareNew(client), h.GetExamCityDivisions(client))
	r.PUT(apiSubPath+"profileadminupdate", util.TokenValidationMiddlewareNew(client), h.UpdateAdminUser(client)) //---vasanth

	r.PUT(apiSubPath+"users/ChangeUserPassword", util.TokenValidationMiddlewareNew(client), h.ChangeUserPasswordNew(client)) //----vasanth dev

	r.PUT(apiSubPath+"errorlogassignment/:id", util.TokenValidationMiddlewareNew(client), h.ErrorLogAssignment(client))
	r.GET(apiSubPath+"admin/getadminuserbyrole/:role", util.TokenValidationMiddlewareNew(client), h.GetAllAdminUsersByRole(client)) //--vasanth

	r.GET("/totp", h.Topgenerate)
	r.GET(apiSubPath+"/api/admin/gettoken/:id", h.Gettoken(client))
	r.GET(apiSubPath+"/totp", h.Topgenerate)
	r.GET(apiSubPath+"/version/:uiversion", h.GetVersion(client))
	r.GET(apiSubPath+"/message", h.GetMessage(client))
	r.POST(apiSubPath+"/refreshtoken", h.ValidateRefreshToken)
	r.POST(apiSubPath+"/totopverify", h.Totopverify)
	r.GET(apiSubPath+"exams", h.GetExams(client))

	r.POST(apiSubPath+"download", fileHandler.Download)        //NL
	r.POST(apiSubPath+"pdf/download", fileHandler.PDFDownload) //NL
	r.POST(apiSubPath+"adminusers/verifynewUser", h.VerifyAdminLogin(client))
	r.GET(apiSubPath+"report/getallapplications/:examcode/:examyear/:cirfaclilityid", h.GetAllApplications(client))
	r.PUT(apiSubPath+"users/changepassword/:id", h.ChangeUserPassword(client)) ////----vasanth

	r.POST(apiSubPath+"upload", func(c *gin.Context) {
		fileHandler.Upload(c, client)
	})

	r.POST(apiSubPath+"testsms/:apitypes", h.TestSMS(client))
	r.GET(apiSubPath+"testing/email", mail.SendTestEmail(client))
	r.Run(":8080")

}

// Edges represents the edges data containing log information

//END OF MAIN
