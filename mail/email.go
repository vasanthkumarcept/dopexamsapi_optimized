package mail

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"

	//"errors"
	"fmt"
	"io"

	//"log"
	"net/http"
	"recruit/ent"
	"recruit/ent/adminmaster"
	"recruit/ent/usermaster"
	"recruit/util"
	"time"

	"github.com/gin-gonic/gin"
	//"recruit/ent/usermaster"
	//"strconv"
)

const (
	FOOTER = "Caution: Giving false information and unauthorized usage may lead to legal action.\nNote : \nFor any technical issue, kindly send an email to cept@indiapost.gov.in.\nFor any administrative issue/clarification, kindly take it up with your higher authority.\nThis is an automated email don't reply to this email."
	// value = 1
	// Candidate OTP Authentication
	//<EmployeeName> ,<Employe ID> .
	// <OTP>
	CANDIADATEOTPAUTHSUB   = "DOP Departmental Exam portal Registration - Candidate OTP Authentication"
	CANDIADATEOTPAUTHEADER = "Dear Candidate %s, Employee ID: %s \n"
	CANDIADATEOTPAUTHBODY  = "OTP for DOP Departmental Exam portal Registration is %s. This OTP will be valid only for two minutes.\n"

	CANDIADATEOTPAUTHSUBE   = "DOP Departmental Exam portal EditProfile - Candidate OTP Authentication"
	CANDIADATEOTPAUTHEADERE = "Dear Candidate %s, Employee ID: %s \n"
	CANDIADATEOTPAUTHBODYE  = "OTP for DOP Departmental Exam portal Profile Edit is %s. This OTP will be valid only for two minutes.\n"
	// value = 2
	// CANDIDATE SUCESSFULL REGISTRATION
	//<EmployeeName> ,<Employe ID> .
	// <password>, <URL>
	CANDIADATEREGSUB    = "DOP Departmental Exam portal Registration - Candidate Successful Registration"
	CANDIADATEREGHEADER = "Dear Candidate %s, Employee ID: %s \n"
	CANDIADATEREGBODY   = "For the DOP Departmental Exam portal login, you successfully registered and authenticated. \n Now you can log in with this password: %s, and the URL is  %s .\n \n Instructions to the Candidates\n 1. Option has been provided to change your password at any time, after login. \n 2. Please keep your Photo and Signature ready before applying, Photo and Signature should be in JPEG format only and should be below (30 to 50 KB).\n 3. Before Final Submit, please check all your information is correct and photo and signature are visible and clear.\n 4. After final submission you can't able to modify the data in the submitted application. Hence before Final submission ensure the correctness of the data.\n 5. Note: Your Application may be rejected if the information is wrong or the photo / Signature is not clear. \n 6. After the Final Submission of your Application ensure the Application number and data fed by you are correctly shown in the printout.\n"

	// value = 3
	// Application Submission
	// <Exam Short Name>
	// <EmployeeName> ,<Employe ID> .
	// <Exam Short Name>, <Application Number> , <Application Submitted date and time>,  <Controlling authority Office Name> <Exam Short Name>
	EXAMSUBMITSUB  = "DOP Departmental Exam portal Application Submission for %s"
	EXAMSUBMITBODY = "You have successfully submitted your application for %s.\n Your application number is %s and submitted on  %s.\n Based on your selection, your application was forwarded to %s (Controlling Authority)  for %s.\nAt regular intervals, check the status of your application in this portal."

	// value =  4  and 5
	//Controlling authority Verification / Comments - Recommended / not recommended
	// <Exam Short Name>
	// <EmployeeName> ,<Employe ID> .
	// <Application Number>,  <Application Submitted Date and time>, <Exam Short Name>, <recommended/not recommended> ,<controlling authority Office Name>, <Commnets of Controlling Authorirty>.

	CAVERIFYRENOTRESUB  = "DOP Departmental Exam portal for %s -  Controlling authority Verification / Comments- %s"
	CAVERIFYRENOTREBODY = "Your application number is %s, submitted  on  %s for %s, got scrutinized and %s to write this Exam. \n 	Remarks of Controlling Authority %s: %s. \n  This status may change at any time based on the Nodal officer observation. Hence, at regular intervals, check the status of your application in this portal.\n	  "

	// value = 6
	//Controlling authority Verification / Comments - for resubmit
	// <Exam Short Name>
	// <EmployeeName> ,<Employe ID> .
	// <Application Number>, <Application Submitted on Date and Time> , <Exam Short Name>, <Controlling Authority Office Name>	, <Comments of Controlling Authority>, <Controlling Authority Office Name>
	CAVERIFYRESUBSUB = "DOP Departmental Exam portal for %s -  Controlling authority Verification / Comments-Returned for Re-Submission"
	// fmt.Sprintf(CAVERIFYRESUBBODY, appno, submittedDate, examshortname, ca, comments, ca)
	CAVERIFYRESUBBODY = "Your application number is %s, submitted on  %s for %s, got scrutinized by the controlling authority and returned for supply of omission.\n 	Remarks of Controlling Authority %s: %s.\n If you found the observation noticed by your controlling authority is controversial to you, before re-submission, get it clarified by your controlling authority.\n For any clarification, you are requested to contact your %s (Controlling authority) \n This status may change at any time based on the Nodal officer observation. Hence, at regular intervals, check the status of your application in this portal. \n 	Instructions to the Candidates\n 	 	1. Login using your credentials\n 2. You can view your Submitted Application in the View and Print Application 	Menu\n 	3. In your application you can find the Comments of your Reporting Authority at the start of the application. \n 	4. Option to modify information only for the items marked as incorrect. \n 	5. Submit your application with correct information as marked by your \ncontrolling authority. \n 	6. Note: Your application may be rejected if your information or images are wrong/not clear. \n 	7. After the Final Submission of your Application, you can take Print out of your application and send it to your Reporting Authority Again and also keep it for your future reference.\n"
	//You can Re-Submit your Application on or before %s\n (get the details from notification)

	// value = 7
	// <Exam Short Name>
	// Application Re-Submission
	//	<EmployeeName> ,<Employe ID> .
	//  <Exam Short Name>,  <Application Number> ,  <Application Submitted date and time>, <Controlling authority Office Name>
	CANDIDATERESUBSUB = "DOP Departmental Exam portal Application Re-Submission for %s"

	//fmt.Sprintf(CANDIDATERESUBBODY, examshortname, appno, submittedDate, ca)
	CANDIDATERESUBBODY = "You have successfully re-submitted your application for %s.\n Your application number is %s and re-submitted on  %s. \n 	Based on your selection, your application was forwarded to %s (Controlling Authority)  . 	At regular intervals, check the status of your application in this portal.\n"

	// value = 8, 9, 10,11
	// Nodal Officer Verification / Comments
	// <Exam Short Name>
	//	<EmployeeName> ,<Employe ID> .
	//  <Application Number>, <Application submitted date> , <Exam Short Name>,  <Nodal office Name> , <new recommendations>, <Nodal ofifcer remarks> <controlling authority office name>
	NORECSUB  = "DOP Departmental Exam portal for %s - Nodal Officer - Change of Recommendations to  %s"
	NORECBODY = "Your application number is %s, submitted on  %s  for %s got scrutinized by the %s (Nodal Officer) and changed the status of the recommendations to %s. \n 	Remarks of Nodal Officer: %s.\n For further clarification, you are requested to contact your %s (Controlling authority). \n  This status may change at any time based on the Nodal officer observation. Hence, at regular intervals, check the status of your application in this portal.\n  "

	// value = 12
	// Hall ticket Intimation
	// <Exam Short Name>
	//	<EmployeeName> ,<Employe ID> .
	//  <Application Number>, <Application Submitted date and time> , <Exam Short Name>,  <Nodal Officer Officer Name> , <Hall ticket Number> , <Hall ticket generated date and time>, <URL> , <Controlling authority Office Name>
	HALLTICKETSUB = "DOP Departmental Exam portal for %s- Hall ticket Intimation"
	//fmt.Sprintf(HALLTICKETBODY, appno, submittedDate, examshortname, no, hallTicket, hallTicketDate, url, ca)
	HALLTICKETBODY = "Your application number is %s, submitted on  %s for %s. \n 	For this application Hall ticket was generated by the %s (Nodal Officer)  and the Hall ticket number is %s on %s. \n 	Download your Hall ticket from this URL: %s using your login credentials. \n	For further clarification, you are requested to contact your %s (Controlling authority). \n 	This status may change at any time based on the Nodal officer observation. Hence, at regular intervals, check the status of your application/Hallticket in this portal.\n"

	// value = 3
	// change to 13
	// Candidate - Get Password
	//	<EmployeeName> ,<Employe ID> .
	// <password>
	GETPASSSUB    = "DOP Departmental Exam portal -  Get Password"
	GETPASSHEADER = "Dear Candidate %s, Employee ID: %s \n"
	GETPASSBODY   = "Based on your request current password is shared and password is %s.\n"

	// value = 4
	// change to 14
	//Candidate - Reset password
	//	<EmployeeName> ,<Employe ID> .
	//<password>
	CHANGEPASSSUB  = "DOP Departmental Exam portal -  Reset password"
	CHANGEPASSBODY = "You used your Profile and resetted your Password and a new password is %s \n"

	CANDIADATEEMPLOYEEREQUESTSUB  = "DOP Departmental Exam portal - Candidate - Employee master Creation Request"
	CANDIADATEEMPLOYEEREQUESTBODY = "Your Employee master creation request submitted to your Controlling authority (%s) Successfully. Without approval of Controlling authority you cannot able to do Registration for this portal. Please wait for futher update\n"

	CANDIADATEEMPLOYEEREQUESTSTATUSSUB  = "DOP Departmental Exam portal - Candidate - Employee master Creation Request response"
	CANDIADATEEMPLOYEEREQUESTSTATUSBODY = "Your Employee master creation request was processed by the %s Controlling authority and status is %s.\n with comments: %s .\n If approved kindly proceed for registration, otherwise contact your controlling authority and resubmit with correct details once again\n"

	ADMINUSERSUB    = "DOP Departmental Exam portal - Amin User Creation "
	ADMINUSERHEADER = "Dear Officer %s, Employee ID: %s \n"
	ADMINUSERBODY   = "Based on your request Admin user credentials created for Role : %s and Username is :%s.\n Default password is sent to your registered mobile number.\n"
)

func SendEMailNew(ctx context.Context, client *ent.Client, userid string, v int, values ...string) string {
	var apiresponse, apiresponsedescription string = "Failed", ""
	var email, empname, empid, vString, sType string = "", "", "", "", ""
	var subject, body, otp, password, comments string = "", "", "", "", ""
	var ca, no, examshortname, appno, submittedDate, recommended string = "", "", "", "", "", ""
	var hallTicket, hallTicketDate string = "", ""
	var url string = "https://k8sapi.cept.gov.in/deptexam"
	var Action, Remarks string
	sType = "Email"

	if v != 35 {
		ua := AdminData(client, userid)
		um := UserData(client, userid)
		em := util.EmployeeData(client, userid)

		switch {
		case ua != nil:
			email = ua.EmailID
			empname = ua.EmployeeName
			empid = ua.UserName
		case um != nil:
			email = um.EmailID
			empname = um.EmployeeName
			empid = um.UserName
		case em != nil:
			email = em.EmailID
			empname = em.EmployeeName
		default:
			Action = "422"
			Remarks = "no userd exists for this ID " + userid
			util.SystemLogError(client, Action, Remarks)
			return " No such user exists"
		}
	}

	/* 	date1 := time.Now().Format("02-01-2006")
	   	time1 := time.Now().Format("03:04:05 PM")
	*/
	switch v {
	// Candidate OTP Authentication
	//<EmployeeName> ,<Employe ID> .
	// <OTP>
	case 1:
		for index, values := range values {
			if index == 0 {
				otp = values
			}
		}
		subject = CANDIADATEOTPAUTHSUB
		body = fmt.Sprintf(CANDIADATEOTPAUTHEADER, empname, empid) + fmt.Sprintf(CANDIADATEOTPAUTHBODY, otp) + FOOTER

	//Candidate Successful Registration
	//<EmployeeName> ,<Employe ID> .
	// <password>, <URL>
	case 2:
		for index, values := range values {
			if index == 0 {
				password = values
			}
		}
		subject = CANDIADATEREGSUB
		body = fmt.Sprintf(CANDIADATEREGHEADER, empname, empid) + fmt.Sprintf(CANDIADATEREGBODY, password, url) + FOOTER

	//Application Submission
	// <Exam Short Name>
	// <EmployeeName> ,<Employe ID> .
	// <Exam Short Name>, <Application Number> , <Application Submitted date and time>,  <Controlling authority Office Name> <Exam Short Name>
	case 3:
		for index, values := range values {
			if index == 0 {
				examshortname = values
			}
			if index == 1 {
				appno = values
			}
			if index == 2 {
				ca = values
			}
			if index == 3 {
				submittedDate = values
			}
		}
		subject = fmt.Sprintf(EXAMSUBMITSUB, examshortname)
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(EXAMSUBMITBODY, examshortname, appno, submittedDate, ca, examshortname) + "\n" + FOOTER
		// <Exam Short Name>, <Application Number> , <Application Submitted date and time>,  <Controlling authority Office Name> <Exam Short Name>

	// 4 and 5
	//Controlling authority Verification / Comments -  Recommended/ not recommended
	// <Exam Short Name>
	// <EmployeeName> ,<Employe ID> .
	// <Application Number>,  <Application Submitted Date and time>, <Exam Short Name>, <recommended/not recommended> ,<controlling authority Office Name>, <Commnets of Controlling Authorirty>.
	case 4:
		//case 11:
		for index, values := range values {
			if index == 0 {
				examshortname = values
			}
			if index == 1 {
				appno = values
			}
			if index == 2 {
				ca = values
			}
			if index == 3 {
				submittedDate = values
			}
			if index == 4 {
				recommended = values
			}
			if index == 5 {
				comments = values
			}
		}
		subject = fmt.Sprintf(CAVERIFYRENOTRESUB, examshortname, recommended)
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(CAVERIFYRENOTREBODY, appno, submittedDate, examshortname, recommended, ca, comments) + "\n" + FOOTER
		// <Application Number>,  <Application Submitted Date and time>, <Exam Short Name>, <recommended/not recommended> ,<controlling authority Office Name>, <Commnets of Controlling Authorirty>.

	// Controlling authority Verification / Comments - returned for re-submission
	// <Exam Short Name>
	// <EmployeeName> ,<Employe ID> .
	// <Application Number>, <Application Submitted on Date and Time> , <Exam Short Name>, <Controlling Authority Office Name>	, <Comments of Controlling Authority>, <Controlling Authority Office Name>
	case 6:
		//case 11:
		for index, values := range values {
			if index == 0 {
				examshortname = values
			}
			if index == 1 {
				appno = values
			}
			if index == 2 {
				ca = values
			}
			if index == 3 {
				submittedDate = values
			}
			if index == 4 {
				comments = values
			}
		}
		subject = fmt.Sprintf(CAVERIFYRESUBSUB, examshortname)
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(CAVERIFYRESUBBODY, appno, submittedDate, examshortname, ca, comments, ca) + "\n" + FOOTER
		// <Application Number>, <Application Submitted on Date and Time> , <Exam Short Name>, <Controlling Authority Office Name>	, <Comments of Controlling Authority>, <Controlling Authority Office Name>

	// Candidate - Application Re-Submission
	// <Exam Short Name>
	// Application Re-Submission
	//	<EmployeeName> ,<Employe ID> .
	//  <Exam Short Name>,  <Application Number> ,  <Application Submitted date and time>, <Controlling authority Office Name>
	case 7:
		for index, values := range values {
			if index == 0 {
				examshortname = values
			}
			if index == 1 {
				appno = values
			}
			if index == 2 {
				ca = values
			}
			if index == 3 {
				submittedDate = values
			}
		}
		subject = fmt.Sprintf(CANDIDATERESUBSUB, examshortname)
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(CANDIDATERESUBBODY, examshortname, appno, submittedDate, ca) + "\n" + FOOTER
		//  <Exam Short Name>,  <Application Number> ,  <Application Submitted date and time>, <Controlling authority Office Name>

	//8,9,10,11
	// Nodal Officer Verification / Comments
	// <Exam Short Name>
	//	<EmployeeName> ,<Employe ID> .
	//  <Application Number>, <Application submitted date> , <Exam Short Name>,  <Nodal office Name> , <new recommendations>, <Nodal ofifcer remarks> <controlling authority office name>

	case 8:
		for index, values := range values {
			if index == 0 {
				examshortname = values
			}
			if index == 1 {
				appno = values
			}
			if index == 2 {
				ca = values
			}
			if index == 3 {
				submittedDate = values
			}
			if index == 4 {
				no = values
			}
			if index == 5 {
				recommended = values
			}
			if index == 6 {
				comments = values
			}
		}
		subject = fmt.Sprintf(NORECSUB, examshortname, recommended)
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(NORECBODY, appno, submittedDate, examshortname, no, recommended, comments, ca) + "\n" + FOOTER
		//  <Application Number>, <Application submitted date> , <Exam Short Name>,  <Nodal office Name> , <new recommendations>, <Nodal ofifcer remarks> <controlling authority office name>

	//  Hall ticket Intimation
	// <Exam Short Name>
	//	<EmployeeName> ,<Employe ID> .
	//  <Application Number>, <Application Submitted date and time> , <Exam Short Name>,  <Nodal Officer Officer Name> , <Hall ticket Number> , <Hall ticket generated date and time>, <URL> , <Controlling authority Office Name>

	case 12:
		for index, values := range values {
			if index == 0 {
				examshortname = values
			}
			if index == 1 {
				appno = values
			}
			if index == 2 {
				ca = values
			}
			if index == 3 {
				submittedDate = values
			}
			if index == 4 {
				no = values
			}
			if index == 5 {
				hallTicket = values
			}
			if index == 6 {
				hallTicketDate = values
			}
		}
		subject = fmt.Sprintf(HALLTICKETSUB, examshortname)
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(HALLTICKETBODY, appno, submittedDate, examshortname, no, hallTicket, hallTicketDate, url, ca) + "\n" + FOOTER
		//  <Application Number>, <Application Submitted date and time> , <Exam Short Name>,  <Nodal Officer Officer Name> , <Hall ticket Number> , <Hall ticket generated date and time>, <URL> , <Controlling authority Office Name>

	// Candidate - Get Password
	//	<EmployeeName> ,<Employe ID> .
	// <password>
	case 13:
		for index, values := range values {
			if index == 0 {
				password = values
			}
		}
		subject = GETPASSSUB
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(GETPASSBODY, password) + FOOTER

	//Candidate - Reset passowrd
	//	<EmployeeName> ,<Employe ID> .
	//<password>
	case 14:
		for index, values := range values {
			if index == 0 {
				password = values
			}
		}
		subject = CHANGEPASSSUB
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(CHANGEPASSBODY, password) + FOOTER

	case 34:
		for index, values := range values {
			if index == 0 {
				empname = values
			}
			if index == 1 {
				ca = values
			}
		}
		subject = CANDIADATEEMPLOYEEREQUESTSUB
		body = fmt.Sprintf(CANDIADATEREGHEADER, empname, userid) + fmt.Sprintf(CANDIADATEEMPLOYEEREQUESTBODY, ca) + FOOTER

	case 35:
		for index, values := range values {
			if index == 0 {
				empname = values
			}
			if index == 1 {
				ca = values
			}
			if index == 2 {
				recommended = values
			}
			if index == 3 {
				comments = values
			}
			if index == 4 {
				email = values
			}
		}
		subject = CANDIADATEEMPLOYEEREQUESTSTATUSSUB
		body = fmt.Sprintf(CANDIADATEREGHEADER, empname, userid) + fmt.Sprintf(CANDIADATEEMPLOYEEREQUESTSTATUSBODY, ca, recommended, comments) + FOOTER
	case 36:
		for index, values := range values {
			if index == 0 {
				empname = values
			}
			if index == 1 {
				ca = values // role name
			}
			if index == 2 {
				recommended = values // admin username
			}
		}
		subject = ADMINUSERSUB
		body = fmt.Sprintf(ADMINUSERHEADER, empname, userid) + fmt.Sprintf(ADMINUSERBODY, ca, recommended) + FOOTER
	case 37:
		for index, values := range values {
			if index == 0 {
				email = values
			}
			if index == 1 {
				otp = values
			}
		}
		subject = CANDIADATEOTPAUTHSUBE
		body = fmt.Sprintf(CANDIADATEOTPAUTHEADERE, empname, empid) + fmt.Sprintf(CANDIADATEOTPAUTHBODYE, otp) + FOOTER
	case 38:
		for index, values := range values {
			if index == 0 {
				otp = values
			}
		}
		subject = CANDIADATEOTPAUTHSUBE
		body = fmt.Sprintf(CANDIADATEOTPAUTHEADERE, empname, empid) + fmt.Sprintf(CANDIADATEOTPAUTHBODYE, otp) + FOOTER

	}

	// Trigger the mail
	statusemail, err := sendMailNew(email, subject, body)
	//statussms, err := util.SendSMSs(msg, phone, templateID, entityID, appName)

	if err != nil {
		if statusemail == 400 || statusemail == 422 {
			apiresponse = "Failed"
			apiresponsedescription = err.Error()
		} else {
			statusemail = 500
			apiresponse = "Failed"
			apiresponsedescription = err.Error()
		}
	} else if statusemail == 200 {
		apiresponse = "Success"
		apiresponsedescription = "Email sent successfully"
	}

	SmsEmailLog(ctx, client, sType, email, empid, vString, subject, apiresponse, apiresponsedescription)

	return apiresponse
}

func sendMailNew(emailid, emailsubject, body string) (int32, error) {
	url := "https://dopverysecure.in/services/mail/send-email"

	requestBody := map[string]string{
		"body":    body,
		"subject": emailsubject,
		"to":      emailid,
	}

	// Ignore SSL certificate validation
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   30 * time.Second}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return 400, fmt.Errorf("failed in json marshal: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer((jsonBody)))
	if err != nil {
		//fmt.Println("EMAIL IS NOT SENT ERROR IN EMAIL API")
		return 400, fmt.Errorf("failed in newRequest newbuffer API: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		//return fmt.Errorf("failed to send request: %v", err)
		return 400, fmt.Errorf("failed in client- Do request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status Code:", resp.Status)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 400, fmt.Errorf("failed to read response body: %v", err)
	}
	fmt.Println("Response Body:", string(respBody))

	if resp.StatusCode != http.StatusOK {
		return 400, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	if string(respBody) != "Message sent to next instance" {
		return 422, fmt.Errorf("Message not sent successfully")
	}
	fmt.Println("OTP Sent to eMail")
	return 200, nil
}

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

func SmsEmailLog(ctx context.Context, client *ent.Client,
	sType string, email string, empid string, vString string, subject string, apiresponse string, apiresponsedescription string) {
	_, err := client.SmsEmailLog.Create().
		SetType(sType).
		SetMobileEmail(email).
		SetUserName(empid).
		SetEventCode(vString).
		SetEventDescription(subject).
		SetApiResponse(apiresponse).
		SetApiResponseDescription(apiresponsedescription).
		SetEventtime(time.Now().Truncate(time.Second)).
		Save(ctx)
	if err != nil {
		util.SystemLogError(client, "400", "SmsEmailLog- updation failed")
	}
}

func SendEMailTest() (int32, error) {
	var subject string
	var body string
	var email string

	subject = "Test email"
	body = "To check email is working or not"
	email = "mohandoss28@gmail.com"

	_, err := sendMail(email, subject, body)
	if err != nil {
		return 422, err
	}
	return 200, nil
}

func SendTestEmail(client *ent.Client) gin.HandlerFunc {
	fn := func(gctx *gin.Context) {
		_, err1 := SendEMailTest()
		if err1 != nil {
			gctx.JSON(http.StatusOK, gin.H{"message": err1.Error()})
			return
		}
		gctx.JSON(http.StatusOK, gin.H{"message": "Email successfully sent"})
	}
	return gin.HandlerFunc(fn)
}

func sendMail(emailid, emailsubject, body string) (int32, error) {
	url := "https://dopverysecure.in/services/mail/send-email"

	requestBody := map[string]string{
		"body":    body,
		"subject": emailsubject,
		"to":      emailid,
	}
	// Ignore SSL certificate validation
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr,
		Timeout: 10 * time.Second, // Ensure a timeout is set
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return 400, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer((jsonBody)))
	if err != nil {
		return 400, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return 400, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 400, fmt.Errorf("failed to read response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return 400, fmt.Errorf("unexpected response status: %s", resp.Status)

	}
	if string(respBody) != "Message sent to next instance" {
		return 400, fmt.Errorf("message not sent successfully")
	}
	return 200, nil
}

/* func SendEMailNewedit(ctx context.Context, client *ent.Client, userid string, v int, newemail string, values ...string) string {
	var apiresponse, apiresponsedescription string = "Failed", ""
	var email, empname, empid, vString, sType string
	var subject, body, otp string
	var Action, Remarks string
	vString = fmt.Sprint(v)
	sType = "Email"
	ua := AdminData(client, userid)
	um := UserData(client, userid)
	em := start.EmployeeData(client, userid)

	switch {
	case ua != nil:
		email = newemail
		empname = ua.EmployeeName
		empid = ua.UserName
	case um != nil:
		email = newemail
		empname = um.EmployeeName
		empid = um.UserName

	case em != nil:
		email = newemail
		empname = em.EmployeeName

	default:
		Action = "422"
		Remarks = "no userd exists for this ID " + userid
		util.SystemLogError(client, Action, Remarks)
		return " No such user exists"
	}

	switch v {
	// Candidate OTP Authentication
	//<EmployeeName> ,<Employe ID> .
	// <OTP>
	case 1:
		for index, values := range values {
			if index == 0 {
				otp = values
			}
		}
		subject = CANDIADATEOTPAUTHSUBE
		body = fmt.Sprintf(CANDIADATEOTPAUTHEADERE, empname, empid) + fmt.Sprintf(CANDIADATEOTPAUTHBODYE, otp) + FOOTER

	}

	// Trigger the mail
	statusemail, err := sendMailNew(email, subject, body)
	//statussms, err := util.SendSMSs(msg, phone, templateID, entityID, appName)

	if err != nil {
		if statusemail == 400 || statusemail == 422 {
			apiresponse = "Failed"
			apiresponsedescription = err.Error()
		} else {
			statusemail = 500
			apiresponse = "Failed"
			apiresponsedescription = err.Error()
		}
	} else if statusemail == 200 {
		apiresponse = "Success"
		apiresponsedescription = "Email sent successfully"
	}

	SmsEmailLog(ctx, client, sType, email, empid, vString, subject, apiresponse, apiresponsedescription)

	return apiresponse
} */

/*
	 func SendEMail(client *ent.Client, userid string, v int, values ...string) error {
		ua := AdminData(client, userid)
		um := UserData(client, userid)
		if ua == nil && um == nil {
			fmt.Println("NO USER Found")
			return fmt.Errorf("NO USER FOUND ")
		}
		var email string
		var empname string
		var empid string
		if ua != nil {
			email = ua.EmailID
			empname = ua.EmployeeName
			empid = fmt.Sprintf("%d", ua.EmployeeId)
		}
		if um != nil {
			email = um.EmailID
			empname = um.EmployeeName
			empid = fmt.Sprintf("%d", um.EmployeeID)
		}

		if email == "" {
			return fmt.Errorf("No Email Found")
		}
		fmt.Println(email)
		var subject string
		var body string
		var otp string
		var password string

		var examshortname string
		var appno string
		date1 := time.Now().Format("02-01-2006")
		time1 := time.Now().Format("03:04:05 PM")
		var ca string

		switch v {
		case 1:
			for index, values := range values {
				if index == 0 {
					otp = values
				}
			}
			subject = CANDIADATEOTPAUTHSUB
			body = fmt.Sprintf(CANDIADATEOTPAUTHEADER, empname, empid) + fmt.Sprintf(CANDIADATEOTPAUTHBODY, otp) + FOOTER

		case 2:
			for index, values := range values {
				if index == 0 {
					password = values
				}
			}
			subject = CANDIADATEREGSUB
			body = fmt.Sprintf(CANDIADATEREGHEADER, empname, empid) + fmt.Sprintf(CANDIADATEREGBODY, password) + FOOTER

		case 3:
			for index, values := range values {
				if index == 0 {
					password = values
				}
			}
			subject = GETPASSSUB
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(GETPASSBODY, password) + FOOTER

		case 4:
			for index, values := range values {
				if index == 0 {
					password = values
				}
			}
			subject = CHANGEPASSSUB
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(CHANGEPASSBODY, password) + FOOTER

		case 11:
			for index, values := range values {
				if index == 0 {
					examshortname = values
				}
				if index == 1 {
					appno = values
				}
				if index == 2 {
					ca = values
				}
			}
			subject = fmt.Sprintf(EXAMSUBMITSUB, examshortname)
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(EXAMSUBMITBODY, examshortname, appno, date1, time1, ca, examshortname) + "\n" + FOOTER

		case 12:
			for index, values := range values {
				if index == 0 {
					examshortname = values
				}
				if index == 1 {
					appno = values
				}
				if index == 2 {
					ca = values
				}
			}
			subject = fmt.Sprintf(EXAMSUBMITSUBB, examshortname)
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(EXAMSUBMITBODYY, examshortname, appno, date1, time1, ca, examshortname) + "\n" + FOOTER
		}

		fmt.Println(subject, body)
		// Trigger the mail
		err := sendMail(email, subject, body)
		if err != nil {
			log.Printf("Failed to send e-Mail: %v", err)
			return fmt.Errorf("failed to send OTPs to email")
		}

		//	log.Printf("OTP sent successfully. User: %s, OTP: %s", newUser.UserName, eotp)

		return nil
	}

	func SendEMail11(client *ent.Client, userid string, v int, values ...string) (int32, error) {
		fmt.Println("Inside SendMail ", userid)
		ua := AdminData(client, userid)
		um := UserData(client, userid)
		if ua == nil && um == nil {
			fmt.Println("NO USER Found")
			return 400, fmt.Errorf("NO USER FOUND ")
		}
		var email string
		var empname string
		var empid string
		if ua != nil {
			email = ua.EmailID
			empname = ua.EmployeeName
			empid = fmt.Sprintf("%d", ua.EmployeeId)
		}
		if um != nil {
			email = um.EmailID
			empname = um.EmployeeName
			empid = fmt.Sprintf("%d", um.EmployeeID)
		}

		if email == "" {
			return 400, fmt.Errorf("No Email Found")
		}
		fmt.Println(email)
		var subject string
		var body string
		var otp string
		var password string

		var examshortname string
		var appno string
		date1 := time.Now().Format("02-01-2006")
		time1 := time.Now().Format("03:04:05 PM")
		var ca string

		switch v {
		case 1:
			for index, values := range values {
				if index == 0 {
					otp = values
				}
			}
			subject = CANDIADATEOTPAUTHSUB
			body = fmt.Sprintf(CANDIADATEOTPAUTHEADER, empname, empid) + fmt.Sprintf(CANDIADATEOTPAUTHBODY, otp) + FOOTER

		case 2:
			for index, values := range values {
				if index == 0 {
					password = values
				}
			}
			subject = CANDIADATEREGSUB
			body = fmt.Sprintf(CANDIADATEREGHEADER, empname, empid) + fmt.Sprintf(CANDIADATEREGBODY, password) + FOOTER

		case 3:
			for index, values := range values {
				if index == 0 {
					password = values
				}
			}
			subject = GETPASSSUB
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(GETPASSBODY, password) + FOOTER

		case 4:
			for index, values := range values {
				if index == 0 {
					password = values
				}
			}
			subject = CHANGEPASSSUB
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(CHANGEPASSBODY, password) + FOOTER

		case 11:
			for index, values := range values {
				if index == 0 {
					examshortname = values
				}
				if index == 1 {
					appno = values
				}
				if index == 2 {
					ca = values
				}
			}
			subject = fmt.Sprintf(EXAMSUBMITSUB, examshortname)
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(EXAMSUBMITBODY, examshortname, appno, date1, time1, ca, examshortname) + "\n" + FOOTER

		case 13:
			for index, values := range values {
				if index == 0 {
					examshortname = values
				}
				if index == 1 {
					appno = values
				}
				if index == 2 {
					ca = values
				}
			}
			subject = fmt.Sprintf(EXAMSUBMITSUBB, examshortname)
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(EXAMSUBMITBODYY, examshortname, appno, date1, time1, ca, examshortname) + "\n" + FOOTER
		}

		fmt.Println(subject, body)
		// Trigger the mail
		err := sendMail(email, subject, body)
		if err != nil {
			log.Printf("Failed to send e-Mail: %v", err)
			return 422, fmt.Errorf("failed to send OTPs to email")
		}

		//	log.Printf("OTP sent successfully. User: %s, OTP: %s", newUser.UserName, eotp)

		return 200, nil
	}

	func SendEMail1(client *ent.Client, userid string, v int, values ...string) (int32, error) {
		fmt.Println("Inside SendMail ", userid)
		ua := AdminData(client, userid)
		um := UserData(client, userid)
		if ua == nil && um == nil {
			fmt.Println("NO USER Found")
			return 400, fmt.Errorf("NO USER FOUND ")
		}
		var email string
		var empname string
		var empid string
		if ua != nil {
			email = ua.EmailID
			empname = ua.EmployeeName
			empid = fmt.Sprintf("%d", ua.EmployeeId)
		}
		if um != nil {
			email = um.EmailID
			empname = um.EmployeeName
			empid = fmt.Sprintf("%d", um.EmployeeID)
		}

		if email == "" {
			return 400, fmt.Errorf("No Email Found")
		}
		fmt.Println(email)
		var subject string
		var body string
		var otp string
		var password string

		var examshortname string
		var appno string
		date1 := time.Now().Format("02-01-2006")
		time1 := time.Now().Format("03:04:05 PM")
		var ca string

		switch v {
		case 1:
			for index, values := range values {
				if index == 0 {
					otp = values
				}
			}
			subject = CANDIADATEOTPAUTHSUB
			body = fmt.Sprintf(CANDIADATEOTPAUTHEADER, empname, empid) + fmt.Sprintf(CANDIADATEOTPAUTHBODY, otp) + FOOTER

		case 2:
			for index, values := range values {
				if index == 0 {
					password = values
				}
			}
			subject = CANDIADATEREGSUB
			body = fmt.Sprintf(CANDIADATEREGHEADER, empname, empid) + fmt.Sprintf(CANDIADATEREGBODY, password) + FOOTER

		case 3:
			for index, values := range values {
				if index == 0 {
					password = values
				}
			}
			subject = GETPASSSUB
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(GETPASSBODY, password) + FOOTER

		case 4:
			for index, values := range values {
				if index == 0 {
					password = values
				}
			}
			subject = CHANGEPASSSUB
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(CHANGEPASSBODY, password) + FOOTER

		case 12:
			for index, values := range values {
				if index == 0 {
					examshortname = values
				}
				if index == 1 {
					appno = values
				}
				if index == 2 {
					ca = values
				}
			}
			subject = fmt.Sprintf(EXAMSUBMITSUB, examshortname)
			body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(EXAMSUBMITBODY, examshortname, appno, date1, time1, ca, examshortname) + "\n" + FOOTER

		}

		fmt.Println(subject, body)
		// Trigger the mail
		err := sendMail(email, subject, body)
		if err != nil {
			log.Printf("Failed to send e-Mail: %v", err)
			return 422, fmt.Errorf("failed to send OTPs to email")
		}

		//	log.Printf("OTP sent successfully. User: %s, OTP: %s", newUser.UserName, eotp)

		return 2000, nil
	}
*/

/* func SendEMailmtspm(client *ent.Client, userid string, v int, values ...string) (int32, error) {
	fmt.Println("Inside SendMail ", userid)
	ua := AdminData(client, userid)
	um := UserData(client, userid)
	if ua == nil && um == nil {
		fmt.Println("NO USER Found")
		return 400, fmt.Errorf("NO USER FOUND ")
	}
	var email string
	var empname string
	var empid string
	if ua != nil {
		email = ua.EmailID
		empname = ua.EmployeeName
		empid = fmt.Sprintf("%d", ua.EmployeeId)
	}
	if um != nil {
		email = um.EmailID
		empname = um.EmployeeName
		empid = fmt.Sprintf("%d", um.EmployeeID)
	}

	if email == "" {
		return 400, fmt.Errorf("No Email Found")
	}
	fmt.Println(email)
	var subject string
	var body string
	var otp string
	var password string

	var examshortname string
	var appno string
	date1 := time.Now().Format("02-01-2006")
	time1 := time.Now().Format("03:04:05 PM")
	var ca string

	switch v {
	case 1:
		for index, values := range values {
			if index == 0 {
				otp = values
			}
		}
		subject = CANDIADATEOTPAUTHSUB
		body = fmt.Sprintf(CANDIADATEOTPAUTHEADER, empname, empid) + fmt.Sprintf(CANDIADATEOTPAUTHBODY, otp) + FOOTER

	case 2:
		for index, values := range values {
			if index == 0 {
				password = values
			}
		}
		subject = CANDIADATEREGSUB
		body = fmt.Sprintf(CANDIADATEREGHEADER, empname, empid) + fmt.Sprintf(CANDIADATEREGBODY, password) + FOOTER

	case 3:
		for index, values := range values {
			if index == 0 {
				password = values
			}
		}
		subject = GETPASSSUB
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(GETPASSBODY, password) + FOOTER

	case 4:
		for index, values := range values {
			if index == 0 {
				password = values
			}
		}
		subject = CHANGEPASSSUB
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(CHANGEPASSBODY, password) + FOOTER

	case 14:
		for index, values := range values {
			if index == 0 {
				examshortname = values
			}
			if index == 1 {
				appno = values
			}
			if index == 2 {
				ca = values
			}
		}
		subject = fmt.Sprintf(EXAMSUBMITSUBBB, examshortname)
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(EXAMSUBMITBODYYY, examshortname, appno, date1, time1, ca, examshortname) + "\n" + FOOTER

	}

	fmt.Println(subject, body)
	// Trigger the mail
	err := sendMail(email, subject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return 422, fmt.Errorf("failed to send OTPs to email")
	}

	//	log.Printf("OTP sent successfully. User: %s, OTP: %s", newUser.UserName, eotp)

	return 200, nil
}
func SendEMailpa(client *ent.Client, userid string, v int, values ...string) (int32, error) {
	fmt.Println("Inside SendMail ", userid)
	ua := AdminData(client, userid)
	um := UserData(client, userid)
	if ua == nil && um == nil {
		fmt.Println("NO USER Found")
		return 400, fmt.Errorf("NO USER FOUND ")
	}
	fmt.Println("im here11")

	var email string
	var empname string
	var empid string
	if ua != nil {
		email = ua.EmailID
		empname = ua.EmployeeName
		empid = fmt.Sprintf("%d", ua.EmployeeId)
	}
	if um != nil {
		email = um.EmailID
		empname = um.EmployeeName
		empid = fmt.Sprintf("%d", um.EmployeeID)
	}

	if email == "" {
		return 422, fmt.Errorf("No Email Found")
	}
	fmt.Println(email)
	var subject string
	var body string
	var otp string
	var password string

	var examshortname string
	var appno string
	date1 := time.Now().Format("02-01-2006")
	time1 := time.Now().Format("03:04:05 PM")
	var ca string

	switch v {
	case 1:
		for index, values := range values {
			if index == 0 {
				otp = values
			}
		}
		subject = CANDIADATEOTPAUTHSUB
		body = fmt.Sprintf(CANDIADATEOTPAUTHEADER, empname, empid) + fmt.Sprintf(CANDIADATEOTPAUTHBODY, otp) + FOOTER

	case 2:
		for index, values := range values {
			if index == 0 {
				password = values
			}
		}
		subject = CANDIADATEREGSUB
		body = fmt.Sprintf(CANDIADATEREGHEADER, empname, empid) + fmt.Sprintf(CANDIADATEREGBODY, password) + FOOTER

	case 3:
		for index, values := range values {
			if index == 0 {
				password = values
			}
		}
		subject = GETPASSSUB
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(GETPASSBODY, password) + FOOTER

	case 4:
		for index, values := range values {
			if index == 0 {
				password = values
			}
		}
		subject = CHANGEPASSSUB
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(CHANGEPASSBODY, password) + FOOTER

	case 14:
		for index, values := range values {
			if index == 0 {
				examshortname = values
			}
			if index == 1 {
				appno = values
			}
			if index == 2 {
				ca = values
			}
		}
		subject = fmt.Sprintf(EXAMSUBMITSUBBB, examshortname)
		body = fmt.Sprintf(GETPASSHEADER, empname, empid) + fmt.Sprintf(EXAMSUBMITBODYYY, examshortname, appno, date1, time1, ca, examshortname) + "\n" + FOOTER

	}

	fmt.Println(subject, body)
	// Trigger the mail
	err := sendMail(email, subject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return 422, fmt.Errorf("failed to send OTPs to email")
	}

	//	log.Printf("OTP sent successfully. User: %s, OTP: %s", newUser.UserName, eotp)

	return 200, nil
}
func SendEMaill(newUser *ent.UserMaster, password string) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New("User's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	epass := password
	body := "Dear Customer,your password is : " + password + ", please do not share it with anyone - INDPOST"
	emailsubject := "Password for Online Departmental Examination"

	// Trigger the mail
	err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send password to email")
	}

	log.Printf("Password sent successfully. User: %s, password: %s", newUser.UserName, epass)

	return nil
}

func SendEMailladmin(newUser *ent.AdminMaster, password string) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New("User's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	epass := password
	body := "Dear Customer,your password is : " + password + ", please do not share it with anyone - INDPOST"
	emailsubject := "Password for Online Departmental Examination"

	// Trigger the mail
	err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send password to email")
	}

	log.Printf("Password sent successfully. User: %s, password: %s", newUser.UserName, epass)

	return nil
}

func SendEMailll(newUser *ent.UserMaster, newPassword string) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New("User's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	epass := newPassword
	body := "Dear Customer,your new password is : " + newPassword + ", please do not share it with anyone - INDPOST"
	emailsubject := " New Password for Online Departmental Examination"

	// Trigger the mail
	err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send new password to email")
	}

	log.Printf("Password sent successfully. User: %s, password: %s", newUser.UserName, epass)

	return nil
}
func SendEMaillll(newUser *ent.AdminMaster, otp string) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New("User's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	eotp := otp
	body := "Dear Customer, OTP for Verification is " + eotp + ", valid for 5 Mins, please do not share it with anyone - INDPOST"
	emailsubject := "OTP for Online Departmental Examination"

	// Trigger the mail
	err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send OTPs to email")
	}

	log.Printf("OTP sent successfully. User: %s, OTP: %s", newUser.UserName, eotp)

	return nil
}

func SendeAdminMailll(newUser *ent.AdminMaster, newPassword string) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New("User's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	epass := newPassword
	body := "Dear Customer,your new password is : " + newPassword + ", please do not share it with anyone - INDPOST"
	emailsubject := " New Password for Online Departmental Examination"

	// Trigger the mail
	err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send new password to email")
	}

	log.Printf("Password sent successfully. User: %s, password: %s", newUser.UserName, epass)

	return nil
} */
