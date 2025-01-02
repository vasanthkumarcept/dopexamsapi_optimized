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

func GetRoles(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		errRemarks := "500 error DB Connection erorr -DBS01"
		errStatus := start.HandleDBErrorInitial(gctx, client, errRemarks)
		if errStatus == "error" {
			return
		}
		var startFunction string = " main -getRoles - start - QueryRoles "
		//ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		roles, status, stgError, dataStatus, err := start.QueryRoles(ctx, client)
		if err != nil {
			start.StartErrorHandlerWithoutLog(gctx, err, status, stgError, client, startFunction)
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "Admin user created  successfully",
			"data":       roles,
			"dataexists": dataStatus,
		})
	}

	return gin.HandlerFunc(fn)

}

// RoleMaster
func InsertRoles(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		newroles := new(ent.RoleMaster)
		if err := gctx.ShouldBindJSON(&newroles); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newroles, err := start.CreateRoles(client, newroles)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, "Successfully created tnew roles")
	}
	return gin.HandlerFunc(fn)
}

// GetRolesByID godoc
// @Summary Get role by ID
// @Description Get details of a specific role by its ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} start.RollmasterResponse "successfully retreived"
// @Failure 400 {object} start.ErrorResponse "Bad Request"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal Server Error"
// @Router /rect/admin/roles/{id} [get]
func GetRolesByID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		id := gctx.Param("id")
		//var examID int32
		eliID, _ := strconv.ParseInt(id, 10, 32)
		elis, err := start.QueryRolesByID(ctx, client, int32(eliID))
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"data": elis})
	}
	return gin.HandlerFunc(fn)
}

// GetAllRoles godoc
// @Summary Get all roles
// @Description Get a list of all roles
// @Tags Roles
// @Produce  json
// @Success 200 {object} start.RollmasterResponse "List of roles retrieved successfully"
// @Failure 400 {object} start.ErrorResponse "Bad Request"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal Server Error"
// @Router /rect/admin/roles [get]
func GetAllRoles(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), util.GetCtxTimeOut())
		defer cancel()
		Dusers, status, stgError, dataStatus, err := start.QueryRoles(ctx, client)
		if err != nil {
			if status == 422 {
				Action = "422"
				Remarks = "422 error from main - GetAllRolesstart.QueryRoles" + stgError
				start.HandleError(gctx, 422, err.Error()+stgError)
				return
			} else {
				errRemarks := "500 error occured in main - GetAllRolesstart.QueryRoles"
				start.HandleDatabaseErrorWithoutLog(gctx, err, status, stgError, client, errRemarks)
				return
			}
		}
		gctx.JSON(http.StatusOK, gin.H{
			"success":    true,
			"message":    "",
			"data":       Dusers,
			"dataexists": dataStatus,
		})

	}
	return gin.HandlerFunc(fn)
}

// GetPasswordByEmpId godoc
// @Summary Get password by employee ID
// @Description Get the password of a user by their employee ID and send it via SMS
// @Tags Users
// @Accept json
// @Produce json
// @Param id1 path int true "Employee ID"
// @Success 200 {object} start.UserResponse "Your current password is Successfully Senty"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/users/getpassword/{id1} [get]
/* func GetPasswordByEmpId(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := gctx.Request.Context()
		id1 := gctx.Param("id1")
		empid, err := strconv.ParseInt(id1, 10, 64)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get the password for the given employee ID
		password, err := start.QueryPasswordByEmpId(ctx, client, empid)
		fmt.Println(password)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Retrieve user information to send SMS
		user, err := client.UserMaster.Query().
			Where(usermaster.EmployeeIDEQ(empid)).
			Only(ctx)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Send the password through SMS
		err = start.SendtextMSGgetpassword(ctx, client, user)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"Your current password is": "Successfully Sent to Your registered e-mail: " + user.EmailID + " and Mobile: " + user.Mobile,
		})

	}
} */

// Developed By Vasanth

// VerifyAdminLogin godoc
// @Summary Verify admin login
// @Description Verify admin login using OTP
// @Tags Users
// @Accept json
// @Produce json
// @Param requestBody body RequestBody true "Change Password Request"
// @Success 200 {object} start.UserResponse "New Password updated Successfully"
// @Failure 400 {object} start.ErrorResponse "Validation error"
// @Failure 401 {object} start.ErrorResponse "Unauthorized error"
// @Failure 403 {object} start.ErrorResponse "Forbidden error"
// @Failure 404 {object} start.ErrorResponse "Facilities not found"
// @Failure 422 {object} start.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} start.ErrorResponse "Internal server error"
// @Router /rect/adminusers/verifynewUser [post]
/* func VerifyAdminLogin(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {

		newUser := new(ent.UserMaster)
		if err := gctx.ShouldBindJSON(&newUser); err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(newUser.Edges.LogData) <= 0 {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Log data"})
			return
		}

		logdata := newUser.Edges.LogData[0]

		ctx := context.Background()
		_, logerr := client.Logs.Create().
			SetUserid(logdata.Userid).
			SetUsertype(logdata.Usertype).
			SetRemarks("Admin Login with otp Attempted").
			SetAction(logdata.Action).
			SetIpaddress(logdata.Ipaddress).
			SetDevicetype(logdata.Devicetype).
			SetOs(logdata.Os).
			SetBrowser(logdata.Browser).
			SetLatitude(logdata.Latitude).
			SetLongitude(logdata.Longitude).
			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
			Save(ctx)

		// Validate the newUser credentials using the ValidateAdminLogin function
		user, circleOfficeName, err := start.ValidateAdminLogin(client, &ent.UserMaster{
			UserName: newUser.UserName,
			OTP:      newUser.OTP,
		})
		if err != nil {
			ctx := context.Background()
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
				Save(ctx)

			if logerr1 != nil {
				log.Println(logerr.Error())
			}
			gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err1 := client.Logs.Create().
			SetUserid(logdata.Userid).
			SetUsertype(logdata.Usertype).
			SetRemarks("Admin Logged With OTP Sucessfully").
			SetAction(logdata.Action).
			SetIpaddress(logdata.Ipaddress).
			SetDevicetype(logdata.Devicetype).
			SetOs(logdata.Os).
			SetBrowser(logdata.Browser).
			SetLatitude(logdata.Latitude).
			SetLongitude(logdata.Longitude).
			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
			Save(context.Background())
		if err1 != nil {
			log.Println("NO LOG ")
		}

		token := authentication.CreateToken(newUser.EmployeeName, client)
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"RoleUserCode": user.RoleUserCode,
				"UserName":     user.UserName,
				"FacilityID":   user.FacilityID,
				"LoginStatus":  "Verified successfully",
				"OfficeName":   circleOfficeName,
				"EmployeeName": user.EmployeeName,
				"Token":        token,
			},
		})
	}
	return gin.HandlerFunc(fn)
} */

// delete users by employee id
/* func DeleteUserbyEmployeeID(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		ctx := gctx.Request.Context()
		EmployeeID, err := strconv.ParseInt(gctx.Param("empid"), 10, 64)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee ID"})
			return
		}
		_, err = start.DeleteUserbyEmployeeID(ctx, client, EmployeeID)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, gin.H{"error": "error processing"})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"Message":    "User Details deleted successfully",
				"EmployeeID": EmployeeID}})
	}
	return gin.HandlerFunc(fn)
} */
