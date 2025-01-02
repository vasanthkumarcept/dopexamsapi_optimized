package handlers

import (
	"context"
	"fmt"
	"net/http"
	"recruit/ent"
	"recruit/ent/adminmaster"
	"recruit/ent/usermaster"
	"recruit/start"
	"recruit/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*func SomeHandler(db *sql.DB) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        // Your handler code goes in here - e.g.
        rows, err := db.Query(...)

        c.String(200, results)
    }

    return gin.HandlerFunc(fn)
}*/

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

/*func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}

	return "Unknown error"
}*/

func GetExamID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()

		id := gctx.Param("id")
		//var examID int32
		examID, _ := strconv.ParseInt(id, 10, 32)

		exam, err := start.QueryExamID(ctx, client, int32(examID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": exam})

	}

	return gin.HandlerFunc(fn)
}

// CA Verified with Emp ID
/* func GetMTSPMMGCAVerifiedDetailsByEmpId(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := context.Background()
		id := gctx.Param("id")
		//var examID int32
		empid, _ := strconv.ParseInt(id, 10, 64)
		circles, err := start.QueryPMPAApplicationsByCAVerifiedByEmpID(ctx, client, int64(empid))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": circles})
	}
	return gin.HandlerFunc(fn)
} */

// func AdminData(client *ent.Client, userid string) *ent.AdminMaster {
// 	admin, err := client.AdminMaster.
// 		Query().
// 		Where(adminmaster.UserNameEQ(userid), adminmaster.StatussEQ("active")).
// 		Only(context.Background())
// 	if err != nil {
// 		return nil
// 	}
// 	return admin
// }

// func Logger(client *ent.Client, logdata *ent.Logs) error {

// 	u := AdminData(client, logdata.Userid)
// 	_, logerr1 := client.Logs.Create().
// 		SetUserid(logdata.Userid).
// 		SetUsertype(logdata.Usertype).
// 		SetRemarks(logdata.Remarks).
// 		SetAction(logdata.Action).
// 		SetIpaddress(logdata.Ipaddress).
// 		SetDevicetype(logdata.Devicetype).
// 		SetOs(logdata.Os).
// 		SetBrowser(logdata.Browser).
// 		SetLatitude(logdata.Latitude).
// 		SetLongitude(logdata.Longitude).
// 		SetEventtime(time.Now()).
// 		SetUserdetails(u.EmployeeName).
// 		SetUniqueid(u.ID).
// 		Save(context.Background())

// 	if logerr1 != nil {
// 		return logerr1
// 	}

// 	return nil
// }

func AdminData(client *ent.Client, userid string) *ent.AdminMaster {
	admin, err := client.AdminMaster.
		Query().
		Where(adminmaster.UserNameEQ(userid), adminmaster.StatussEQ("active")).
		Only(context.Background())
	if err != nil {
		return nil
	}
	return admin
}

func UserData(client *ent.Client, userid string) *ent.UserMaster {
	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(userid), usermaster.StatussEQ("active")).
		Only(context.Background())
	if err != nil {
		return nil
	}
	return user
}

func LogError(client *ent.Client, logdata *ent.Logs, err error) error {
	fmt.Println("Inside Log")
	ua := AdminData(client, logdata.Userid)
	um := UserData(client, logdata.Userid)

	if ua == nil && um == nil {
		fmt.Println("NO USER FOUND")
		return fmt.Errorf("NO USER FOUND ")
	}
	var userdetails string
	var useruniqueid int64

	if ua != nil {
		userdetails = ua.UserName
		useruniqueid = ua.ID
	}
	if um != nil {
		userdetails = um.UserName
		useruniqueid = um.ID
	}
	_, logerr1 := client.Logs.Create().
		SetUserid(logdata.Userid).
		SetUsertype(logdata.Usertype).
		SetRemarks(err.Error()).
		SetAction(logdata.Action).
		SetIpaddress(logdata.Ipaddress).
		SetDevicetype(logdata.Devicetype).
		SetOs(logdata.Os).
		SetBrowser(logdata.Browser).
		SetLatitude(logdata.Latitude).
		SetLongitude(logdata.Longitude).
		SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
		SetUserdetails(userdetails).
		SetUniqueid(useruniqueid).
		Save(context.Background())

	if logerr1 != nil {
		fmt.Println(logerr1.Error())
		return logerr1
	}

	return nil
}

func Logger(client *ent.Client, logdata *ent.Logs) error {
	ua := AdminData(client, logdata.Userid)
	um := UserData(client, logdata.Userid)

	if ua == nil && um == nil {
		return fmt.Errorf("NO USER FOUND ")
	}
	var userdetails string
	var useruniqueid int64
	if ua != nil {
		userdetails = ua.UserName
		useruniqueid = ua.ID
	}
	if um != nil {
		userdetails = um.UserName
		useruniqueid = um.ID
	}
	_, logerr1 := client.Logs.Create().
		SetUserid(logdata.Userid).
		SetUsertype(logdata.Usertype).
		SetRemarks(logdata.Action + " successful").
		SetAction(logdata.Action).
		SetIpaddress(logdata.Ipaddress).
		SetDevicetype(logdata.Devicetype).
		SetOs(logdata.Os).
		SetBrowser(logdata.Browser).
		SetLatitude(logdata.Latitude).
		SetLongitude(logdata.Longitude).
		SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
		SetUserdetails(userdetails).
		SetUniqueid(useruniqueid).
		Save(context.Background())

	if logerr1 != nil {
		return logerr1
	}
	return nil
}
