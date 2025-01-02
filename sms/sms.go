package sms

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"net/http"
	"net/url"
	"recruit/ent"
	"recruit/ent/adminmaster"

	"github.com/go-resty/resty/v2"

	"recruit/util"
	"strconv"

	//"recruit/ent/smsemaillog"

	//"recruit/ent/exam_applications_ip"
	"recruit/ent/usermaster"
	//"strconv"

	"time"
)

type APIResponse struct {
	Success bool `json:"success"`
	//Message string `json:"message"`
	//ResponseText string `json:"data.response_text"`
}

// Constants for SMS subjects and bodies
const (
	CANDIADATEOTPAUTHSUB   = "DOP Departmental Exam portal Registration - Candidate OTP Authentication"
	CANDIADATEOTPAUTHEADER = "Dear Candidate %s, Employee ID: %s \n"
	CANDIADATEOTPAUTHBODY  = "OTP for DOP Departmental Exam portal registration is %s. This OTP will be valid only for two minutes.\n\n"
	CANDIADATEREGSUB       = "DOP Departmental Exam portal Registration - Candidate Successful Registration"
	CANDIADATEREGHEADER    = "Dear Candidate %s, Employee ID: %s \n"
	CANDIADATEREGBODY      = "For the DOP Departmental Exam portal, you successfully registered and authenticated. \n Now you can log in with this password: %s, and the URL is  utilities.cept.gov.in/dopexam .\n \n Instructions to the Candidates\n 1. Option has been provided to change your password at any time, after login. \n 2. Please keep your Photo and Signature ready before applying, Photo and Signature should be in JPEG format only and should be below (30 to 50 KB).\n 3. Before Final Submit, please check all your information is correct and photo and signature are visible and clear.\n 4. After final submission you can't able to modify the data in the submitted application. Hence before Final submission ensure the correctness of the data.\n 5. Note: Your Application may be rejected if the information is wrong or the photo / Signature is not clear. \n 6. After the Final Submission of your Application ensure the Application number and data fed by you are correctly shown in the printout.\n"
	GETPASSSUB             = "DOP Departmental Exam portal - Candidate - Get Password"
	GETPASSHEADER          = "Dear Candidate %s, Employee ID: %s \n"
	GETPASSBODY            = "Based on your request current password is shared and password is %s.\n"
	CHANGEPASSSUB          = "DOP Departmental Exam portal - Candidate - Reset password"
	CHANGEPASSBODY         = "You used your Profile and reset your Password and a new password is %s \n"
	EXAMSUBMITSUB          = "DOP Departmental Exam portal Application Submission for %s"
	EXAMSUBMITBODY         = "You have successfully submitted your application for %s.\n Your application number is %s and submitted on  %s at %s.\nBased on your selection, your application was forwarded to %s (Controlling Authority)  for %s.\nAt regular intervals, check the status of your application in this portal."
	FOOTER                 = "Caution: Giving false information and unauthorized usage may lead to legal action.\nNote : \nFor any technical issue, kindly send an email to cept@indiapost.gov.in.\nFor any administrative issue/clarification, kindly take it up with your higher authority.\nThis is an automated email don't reply to this email."
)

func SendSmsNew(ctx context.Context, client *ent.Client, userid string, v int, values ...string) string {
	var msg, templateID, entityID, appName string
	var mobile, empname, empid string
	var apiresponse, apiresponsedescription string = "Failed", ""
	var Action, Remarks string
	var vString = fmt.Sprint(v)
	var sType = "SMS"

	if v != 35 {
		// Fetch user details
		ua := util.AdminDatas(client, userid)
		um := util.UserDatas(client, userid)
		em := util.EmployeeData(client, userid)
		// Choose the correct user details
		switch {
		case ua != nil:
			mobile = ua.Mobile
			empname = ua.EmployeeName
			empid = ua.UserName
		case um != nil:
			mobile = um.Mobile
			empname = um.EmployeeName
			empid = um.UserName
		case em != nil:
			mobile = em.MobileNumber
			empname = em.EmployeeName
			empid = userid
		default:
			Action = "422"
			Remarks = "no userd exists for this ID " + userid
			util.SystemLogError(client, Action, Remarks)
			return " No such user exists"
		}

		if mobile == "" {
			Action = "422"
			Remarks = "Mobile number is blank"
			util.SystemLogError(client, Action, Remarks)
			return "Mobile numebr is blank"
		}

	}
	var apitype int32 = 0
	version, err := client.Version.Query().
		All(ctx)
	if err != nil {
		return "unable to fetch API process data"
	} else {
		if len(version) == 0 {
			return "No API process data found"
		}
	}
	apitype = version[0].ApiType
	appName = "DOPExam"
	// Prepare SMS based on the provided type
	switch v {
	case 1:
		otp := values[0]
		msg = "Dear " + empname + "-" + empid + ", OTP for DOP Exam Portal registration is " + otp + ". Valid for two minutes-DOPExam-INDIAPOST"
		templateID = "1007677864473680257"
		entityID = "1001081725895192800"
	case 2:
		//password := values[0]
		msg = "Dear " + empname + "-" + empid + ", User ID was successfully registered. For details check email-DOPExam-INDIAPOST"
		templateID = "1007033293383430282"
		entityID = "1001081725895192800"
	case 3:
		//" CE GDS to PM/MG/MTS ", newPAApppln.ApplicationNumber, newPAApppln.ControllingOfficeName)

		smsShortName := values[0] // SMS Exam Short name
		appNumber := values[1]    // Application Number
		status := values[2]       // Submitted / Resubmitted

		msg = "Dear " + empname + "-" + empid + "," + appNumber + " for " + smsShortName + " was " + status + " successfully. For details check email-DOPExam-INDIAPOST"
		templateID = "1007440965314067008"
		entityID = "1001081725895192800"

	case 4:
		//ctx, client, Smsuserid, 4, "LDCE for MTS to PM/MG", application.ApplicationNumber, "Recommended/Not Recommended"
		smsShortName := values[0] // SMS Exam Short name
		appNumber := values[1]    // Application Number
		status := values[2]       // Recommendations
		//appstat := values[3]
		msg = "Dear " + empname + "-" + empid + ", Apl. No." + appNumber + " for " + smsShortName + " was " + status + ". For details check email-DOPExam-INDIAPOST"
		templateID = "1007453409116473296"
		entityID = "1001081725895192800"

	case 12:
		appNumber := values[0]    // Application Number
		hallticket := values[1]   // Hallticket number
		smsShortName := values[2] // SMS Exam Short name

		msg = "Dear " + empname + "-" + empid + "," + appNumber + ", Hall ticket no. " + hallticket + " " + smsShortName + ". For details check email-DOPExam-INDIAPOST"

		templateID = "1007210429836493992"
		entityID = "1001081725895192800"

	case 13:
		password := values[0] // password

		msg = "Dear " + empname + "-" + empid + ", Current password for the DOP Exam portal is " + password + "-DOPExam-INDIAPOST"

		templateID = "1007896572308218801"
		entityID = "1001081725895192800"

	case 14:
		password := values[0] // password

		msg = "Dear " + empname + "-" + empid + ", New password for the DOP Exam portal is " + password + "-DOPExam-INDIAPOST"

		templateID = "1007247787940708106"
		entityID = "1001081725895192800"

	case 15:
		smsShortName := values[0] // SMS Exam Short name
		appCount := values[1]     // Application count

		msg = "Dear " + empid + ", For " + smsShortName + " exam, " + appCount + " application are pending at your end-DOPExam-INDIAPOST"

		templateID = "1007754017982663829"
		entityID = "1001081725895192800"

	case 31:
		otp := values[0] //OTP

		msg = "Dear " + empid + ", OTP is " + otp + ". Valid for two minutes-DOPExam-INDIAPOST"

		templateID = "1007678962930615116"
		entityID = "1001081725895192800"

	case 32:
		password := values[0] // password

		msg = "Dear " + empid + " Password is " + password + "-DOPExam-INDIAPOST"

		templateID = "1007080791170637735"
		entityID = "1001081725895192800"

	case 34:
		employeeName := values[0] // password
		msg = "Dear " + employeeName + "-" + userid + ", Employee master Creation Request Submitted to CA successfully-DOPExam-INDIAPOST"
		templateID = "1007038777862462117"
		entityID = "1001081725895192800"

	case 35:
		//employeemaster.EmployeeName, 		employeemaster.ModifiedByUserName, "Rejected", employeemaster.MobileNumber
		employeeName := values[0] // employee name
		empIDcaName := userid + ", " + values[1]
		status := values[2] // status
		mobile = values[3]
		msg = "Dear " + employeeName + "-" + empIDcaName + " " + status + " your Employee master creation request -DOPExam-INDIAPOST"
		fmt.Println("msg ", msg)
		templateID = "1007706888030920399"
		entityID = "1001081725895192800"

	case 36:

		employeeName := values[0] // employee name
		roleName := values[1]
		userName := values[2]
		msg = "Dear " + employeeName + ", Admin user created for " + roleName + "  role with username: " + userName + "  password:Cept@123-DOPExam-INDIAPOST"
		templateID = "1007725699258101708"
		entityID = "1001081725895192800"
	case 37:
		fmt.Println("im here")
		otp := values[0]
		msg = "Dear " + empname + "-" + empid + ", OTP for DOP Exam Portal edit profile is " + otp + ". Valid for two minutes-DOPExam-INDIAPOST"
		templateID = "1007677864473680257"
		entityID = "1001081725895192800"

	case 38:
		fmt.Println("im here1")
		applnNumber := values[0]
		hallct := values[1]

		msg = "Dear " + empname + "-" + empid + "-" + applnNumber + ", Hall ticket no.  " + hallct + ".For details, check email-DOPExam-INDIAPOST."
		templateID = "1007210429836493992"
		entityID = "1001081725895192800"

	case 39:
		mobile = values[0]
		otp := values[1]

		msg = "Dear " + empid + ", OTP is " + otp + ". Valid for two minutes-DOPExam-INDIAPOST"

		templateID = "1007678962930615116"
		entityID = "1001081725895192800"

	default:
		Action = "400"
		Remarks = "Invalid SMS Type"
		util.SystemLogError(client, Action, Remarks)
		return " invalid SMS Type"
	}

	var statussms int32 = 0
	var napitype string = ""
	var url string = ""
	switch {
	case apitype == 1:
		statussms, err = SendSMS1(msg, mobile, templateID, entityID, appName)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-" + vString
	case apitype == 2:

		url = "https://uat.cept.gov.in/sms/v1/msgrequest/create"
		statussms, err = SendSMS2(msg, mobile, templateID, entityID, url)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-" + vString
	case apitype == 3:
		url = "https://apiservices.cept.gov.in/bemsggateway/v1/msgrequest/create"
		statussms, err = SendSMS2(msg, mobile, templateID, entityID, url)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-" + vString
	case apitype == 4:
		url = "https://dev.cept.gov.in/bemsggateway/v1/msgrequest/create"
		statussms, err = SendSMS3(msg, mobile, templateID, entityID, url)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-" + vString
	case apitype == 5:
		url = "https://test.cept.gov.in/bemsggateway/v1/msgrequest/create"
		statussms, err = SendSMS3(msg, mobile, templateID, entityID, url)
		napitype = strconv.FormatInt(int64(apitype), 10)
		vString = napitype + "-" + vString

	default:
		Action = "400"
		Remarks = "Invalid SMS API call type"
		util.SystemLogError(client, Action, Remarks)
		return " Invalid SMS  API call type"
	}

	if err != nil {
		if statussms == 400 || statussms == 422 {
			apiresponse = "Failed"
			smsErrorSending(ctx, client, apitype)
			apiresponsedescription = err.Error()
		} else {
			statussms = 500
			apiresponse = "Failed"
			smsErrorSending(ctx, client, apitype)
			apiresponsedescription = err.Error()

		}
	} else if statussms == 200 {
		apiresponse = "Success"
		apiresponsedescription = "SMS sent successfully"
	}

	SmsEmailLog(ctx, client, sType, mobile, empid, vString, templateID, apiresponse, apiresponsedescription)
	return apiresponse
}

func smsErrorSending(ctx context.Context, client *ent.Client, apitype int32) {
	var napitype string = ""
	var appName string = "DOPExam"
	var statussms int32 = 0
	var statussms1 int32 = 0
	var err error
	var err1 error
	var url string = ""
	var apiresponse string = ""
	var apiresponsedescription string = ""
	napitype = strconv.FormatInt(int64(apitype), 10)
	msg := "Dear OTPFailed-" + napitype + ", OTP for DOP Exam Portal registration is 123456. Valid for two minutes-DOPExam-INDIAPOST"
	templateID := "1007677864473680257"
	entityID := "1001081725895192800"
	vString := napitype
	mobile := "7299174555"

	switch {

	case apitype == 1:
		url = "https://uat.cept.gov.in/sms/v1/msgrequest/create"
		statussms1, err1 = SendSMS2(msg, mobile, templateID, entityID, url)
		if err1 != nil {
			url = "https://apiservices.cept.gov.in/bemsggateway/v1/msgrequest/create"
			statussms, err = SendSMS2(msg, mobile, templateID, entityID, url)
		}
	case apitype == 2:
		statussms1, err1 = SendSMS1(msg, mobile, templateID, entityID, appName)
		if err1 != nil {
			url = "https://apiservices.cept.gov.in/bemsggateway/v1/msgrequest/create"
			statussms, err = SendSMS2(msg, mobile, templateID, entityID, url)
		}
	case apitype == 3:
		statussms1, err1 = SendSMS1(msg, mobile, templateID, entityID, appName)
		if err1 != nil {
			url = "https://uat.cept.gov.in/sms/v1/msgrequest/create"
			statussms, err = SendSMS2(msg, mobile, templateID, entityID, url)
		}
	case apitype == 4:
		statussms1, err1 = SendSMS1(msg, mobile, templateID, entityID, appName)
		if err1 != nil {
			url = "https://uat.cept.gov.in/sms/v1/msgrequest/create"
			statussms, err = SendSMS2(msg, mobile, templateID, entityID, url)
		}
	case apitype == 5:
		statussms1, err1 = SendSMS1(msg, mobile, templateID, entityID, appName)
		if err1 != nil {
			url = "https://uat.cept.gov.in/sms/v1/msgrequest/create"
			statussms, err = SendSMS2(msg, mobile, templateID, entityID, url)
		}
	default:

		util.SystemLogError(client, "400", "Invalid SMS Type")

	}
	if err != nil && err1 != nil {
		if (statussms == 400 || statussms == 422) && (statussms1 == 400 || statussms1 == 422) {
			apiresponse = "Failed"
			apiresponsedescription = err.Error()
		} else {
			statussms = 500
			apiresponse = "Failed"
			apiresponsedescription = err.Error()
		}
	} else if statussms == 200 || statussms1 == 200 {
		apiresponse = "Success"
		apiresponsedescription = "SMS sent successfully"
	}
	SmsEmailLog(ctx, client, "SMS", mobile, "10024225", vString, templateID, apiresponse, apiresponsedescription)
}
func SmsEmailLog(ctx context.Context, client *ent.Client,
	sType string, mobile string, empid string, vString string, templateID string, apiresponse string, apiresponsedescription string) {
	_, err := client.SmsEmailLog.Create().
		SetType(sType).
		SetMobileEmail(mobile).
		SetUserName(empid).
		SetEventCode(vString).
		SetEventDescription(templateID).
		SetApiResponse(apiresponse).
		SetApiResponseDescription(apiresponsedescription).
		SetEventtime(time.Now().Truncate(time.Second)).
		Save(ctx)
	if err != nil {
		fmt.Println("smslog update error", err.Error())
		util.SystemLogError(client, "400", "SMS - SmsEmailLog- updation failed")
	}
}

func SendSMS1(msg, phone, templateID, entityID, appName string) (int32, error) {
	url := "https://api.cept.gov.in/sendsms/api/values/sendsms"

	payload := fmt.Sprintf(`{
		"Msg": "%s",
		"Phone": "%s",
		"TemplateID": "%s",
		"EntityID": "%s",
		"AppName": "%s"
	}`, msg, phone, templateID, entityID, appName)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		return 400, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return 400, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 400, fmt.Errorf("failed to read response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return 400, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	response := string(body)
	if !strings.Contains(response, "SMS Pushed to NIC Successfully") {
		return 400, fmt.Errorf("failed to send SMS: %s", response)
	}
	return 200, nil
}

type ResponseSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		CommunicationID  string `json:"communication_id"`
		CompleteResponse string `json:"complete_response"`
		ReferenceID      string `json:"ReferenceID"`
		Status           string `json:"status"`
		ResponseText     string `json:"response_text"`
	} `json:"data"`
}
type ResponseFailure struct {
	Success bool     `json:"success"`
	Message []string `json:"message"`
	ErrorNo []string `json:"errorno"`
}

// func CallAPI(url string, method string, headers map[string]string, params map[string]interface{}) (map[string]interface{}, error) {
func SendSMS2(msg, phone, templateID, entityID string, url string) (int32, error) {
	//url := "https://uat.cept.gov.in/sms/v1/msgrequest/create"

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]interface{}{
		"application_id": "28",
		"facility_id":    "CE00308100000",
		"priority":       1,
		"message_text":   msg,
		"sender_id":      "INPOST",
		"mobile_numbers": phone,
		"entity_id":      entityID,
		"template_id":    templateID,
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: false,                       // Set to true only for testing; not recommended for production.
			Renegotiation:      tls.RenegotiateOnceAsClient, // Adjust renegotiation settings.
		},
		DisableKeepAlives: true, // Disable keep-alive
	}
	client := resty.New().SetTimeout(30 * time.Second)
	client.SetTransport(tr)
	request := client.R()
	request.SetHeaders(headers)

	method := "POST"
	request.SetBody(params)
	response, err := request.Execute(method, url)
	if err != nil {
		return 400, err
	}
	var errorValue string = ""
	// Unmarshal the JSON data into the Response struct
	var responseSuccess ResponseSuccess
	err = json.Unmarshal(response.Body(), &responseSuccess)
	if err != nil {
		var responseFailure ResponseFailure
		err = json.Unmarshal(response.Body(), &responseFailure)
		if err != nil {
			return 400, err
		}
		// Extract the errorno value
		if len(responseFailure.ErrorNo) > 0 {
			errorValue = responseFailure.ErrorNo[0]
		} else {
			errorValue = "unknown error"
		}
		return 400, errors.New(errorValue)
	}
	if responseSuccess.Success {
		//fmt.Println("success", responseSuccess.Message)
		//errorValue = responseSuccess.Message
		return 200, nil
	} else {
		return 400, errors.New(errorValue)
	}
}

func SendSMS3(msg, phone, templateID, entityID string, apiURL string) (int32, error) {
	//url := "https://uat.cept.gov.in/sms/v1/msgrequest/create"

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]interface{}{
		"application_id": "28",
		"facility_id":    "CE00308100000",
		"priority":       1,
		"message_text":   msg,
		"sender_id":      "INPOST",
		"mobile_numbers": phone,
		"entity_id":      entityID,
		"template_id":    templateID,
	}

	// Proxy configuration
	proxyURL, _ := url.Parse("http://172.28.12.2:3128")
	fmt.Println("proxyURL", proxyURL)
	tr := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			domainsToUseProxy := map[string]bool{
				"dev.cept.gov.in":  true,
				"test.cept.gov.in": true,
			}

			if domainsToUseProxy[req.URL.Host] {
				return proxyURL, nil
			}
			return nil, nil
		},
		TLSClientConfig: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: false,
			Renegotiation:      tls.RenegotiateOnceAsClient,
		},
		DisableKeepAlives: true,
	}
	fmt.Println("tr", tr)
	client := resty.New().
		SetTransport(tr).
		SetTimeout(30 * time.Second)
	fmt.Println("client", client)

	request := client.R().
		SetHeaders(headers).
		SetBody(params)
	fmt.Println("request", request)
	response, err := request.Post(apiURL)
	fmt.Println("response", response)
	fmt.Println("response err", err)
	if err != nil {
		return 400, err
	}

	var responseSuccess ResponseSuccess
	err = json.Unmarshal(response.Body(), &responseSuccess)
	fmt.Println("responseSuccess err", err)
	if err != nil {
		var responseFailure ResponseFailure
		if err = json.Unmarshal(response.Body(), &responseFailure); err != nil {
			return 400, err
		}

		errorValue := "unknown error"
		if len(responseFailure.ErrorNo) > 0 {
			errorValue = responseFailure.ErrorNo[0]
			fmt.Println("responseFailure.ErrorNo", errorValue)
		}
		return 400, errors.New(errorValue)
	}

	if responseSuccess.Success {
		return 200, nil
	}

	return 400, errors.New("request failed")
}

func ConvertMapToStringMap(params map[string]interface{}) map[string]string {
	stringParams := make(map[string]string)
	for key, value := range params {
		stringParams[key] = interfaceToString(value)
	}
	return stringParams
}

/* func sendSMS3(msg, phone, templateID, entityID string) (int32, error) {
	url := "https://apiservices.cept.gov.in/bemsggateway/v1/msgrequest/create"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]interface{}{
		"application_id": "28",
		"facility_id":    "CE00308100000",
		"priority":       1,
		"message_text":   msg,
		"sender_id":      "INPOST",
		"mobile_numbers": phone,
		"entity_id":      entityID,
		"template_id":    templateID,
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: false,                       // Set to true only for testing; not recommended for production.
			Renegotiation:      tls.RenegotiateOnceAsClient, // Adjust renegotiation settings.
		},
		DisableKeepAlives: true, // Disable keep-alive
	}
	client := resty.New().SetTimeout(30 * time.Second)
	client.SetTransport(tr)
	request := client.R()
	request.SetHeaders(headers)

	method := "POST"
	request.SetBody(params)
	response, err := request.Execute(method, url)
	if err != nil {
		return 400, err
	}
	var errorValue string = ""
	// Unmarshal the JSON data into the Response struct
	var responseSuccess ResponseSuccess
	err = json.Unmarshal(response.Body(), &responseSuccess)
	if err != nil {
		var responseFailure ResponseFailure
		err = json.Unmarshal(response.Body(), &responseFailure)
		if err != nil {
			return 400, err
		}
		// Extract the errorno value
		if len(responseFailure.ErrorNo) > 0 {
			errorValue = responseFailure.ErrorNo[0]
		} else {
			errorValue = "unknown error"
		}
		return 400, errors.New(errorValue)
	}
	if responseSuccess.Success {
		fmt.Println("success", responseSuccess.Message)
		errorValue = responseSuccess.Message
		return 200, nil
	} else {
		return 400, errors.New(errorValue)
	}
} */

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
		Where(usermaster.UserNameEQ(userid)).
		Only(context.Background())
	if err != nil {
		return nil
	}
	return user
}

func interfaceToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
