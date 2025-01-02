package start

/* "context"
"errors"
"fmt"
"log"
"recruit/ent"
"recruit/ent/adminmaster"
"recruit/ent/usermaster"
"recruit/util"
"strconv"
"time" */

/*
	 func SendSMSCandidateOTP(ctx context.Context, client *ent.Client, newUser *ent.UserMaster) (int32, int32, error) {
		// Generate OTP
		otp := GenerateOTP()

		// Convert OTP to string
		stringOTP := strconv.Itoa(int(otp))

		// Check if username is not null
		if newUser.UserName == "" {
			return otp, 400, errors.New(" username is null")
		}

		// Retrieve the user from the database based on the username
		user, err := client.UserMaster.Query().Where(usermaster.UserNameEQ(newUser.UserName)).Only(ctx)
		if err != nil {
			return otp, 400, fmt.Errorf("failed to retrieve user data: %v", err)
		}

		// Retrieve the mobile number from the retrieved user
		mobileNumber := user.Mobile
		if mobileNumber == "" {
			return otp, 422, errors.New(" user's mobile number not found")
		}

		// Convert OTP to string
		//stringOTP, _ := strconv.Itoa(otp)

		// Construct the SMS message
		msg := "Dear " + user.EmployeeName + "-" + newUser.UserName + ", OTP for DOP Exam Portal registration is " + stringOTP + ". Valid for two minutes-DOPExam-INDIAPOST"

		// Set other SMS parameters
		phone := mobileNumber
		templateID := "1007677864473680257"
		entityID := "1001081725895192800"
		appName := "DOPExam"

		var apiresponse string
		var apiresponsedescription string

		// Trigger the SMS
		statussms, err := util.SendSMSs(msg, phone, templateID, entityID, appName)
		if err != nil {
			if statussms == 400 {
				apiresponse = "Failed"
				apiresponsedescription = err.Error()
			}
		} else {
			apiresponse = "Success"
			apiresponsedescription = "SMS sent successfully"
		}

		_, _ = client.SmsEmailLog.Create().
			SetType("SMS").
			SetMobileEmail(phone).
			SetUserName(newUser.UserName).
			SetEventCode("1").
			SetEventDescription("Registration Candidate OTP Authentication").
			SetApiResponse(apiresponse).
			SetApiResponseDescription(apiresponsedescription).
			SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
			Save(ctx)
		return otp, 200, nil

}
*/
/* func SendtextMSGgetpassword(ctx context.Context, client *ent.Client, user *ent.UserMaster) error {
	// Check if user is nil
	if user == nil {
		return errors.New(" user is nil")
	}

	// Retrieve the mobile number from the user
	mobileNumber := user.Mobile
	password := user.Password
	empname := user.EmployeeName
	empid := user.EmployeeID
	if mobileNumber == "" {
		return errors.New(" user's mobile number not found")
	}

	// Construct the SMS message
	msg := "Dear " + empname + " - " + strconv.FormatInt(empid, 10) + " Current password for the DOP Exam portal is  " + password + " -DOPExam-INDIAPOST"
	// Set other SMS parameters
	phone := mobileNumber
	templateID := "1007896572308218801"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	statussms, err := util.SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(statussms)
		log.Printf("Failed to send SMS: %v", err)
		return fmt.Errorf("failed to send SMS")
	}

	nerr := SendEMaill(user, password)
	if nerr != nil {
		return fmt.Errorf(" failed to send E-Mail : %v", user.EmailID)
	}
	log.Printf("SMS sent successfully to %s", phone)
	log.Printf("Email sent successfully to %s", user.EmailID)
	return nil
} */

/* func SendtextMSGuserRig(ctx context.Context, client *ent.Client, user *ent.UserMaster) error {
	// Check if user is nil
	if user == nil {
		return errors.New(" user is nil")
	}

	// Retrieve the mobile number from the user
	mobileNumber := user.Mobile
	//password := user.Password
	empname := user.EmployeeName
	empid := user.EmployeeID
	if mobileNumber == "" {
		return errors.New(" user's mobile number not found")
	}

	// Construct the SMS message
	msg := "Dear " + empname + " - " + strconv.FormatInt(empid, 10) + ", User ID was successfully registered. For details check email-DOPExam-INDIAPOST"
	// Set other SMS parameters
	phone := mobileNumber
	templateID := "1007033293383430282"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	statussms, err := util.SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(statussms)
		log.Printf("Failed to send SMS: %v", err)
		return fmt.Errorf(" failed to send SMS")
	}

	nerr := SendEMailuserreg(user, empname, empid)
	if nerr != nil {
		return fmt.Errorf(" failed to send E-Mail : %v", user.EmailID)
	}
	log.Printf("SMS sent successfully to %s", phone)
	log.Printf("Email sent successfully to %s", user.EmailID)
	return nil
} */

// func SendTextMessage(ctx context.Context, client *ent.Client, phone, msg string) error {
// 	// Check if phone number is empty
// 	if phone == "" {
// 		return errors.New("Phone number is empty")
// 	}

// 	// Set other SMS parameters
// 	templateID := "1007896572308218801"
// 	entityID := "1001081725895192800"
// 	appName := "IBC"

// 	// Trigger the SMS
// 	err := sendSMSs(msg, phone, templateID, entityID, appName)
// 	if err != nil {
// 		log.Printf("Failed to send SMS: %v", err)
// 		return fmt.Errorf("failed to send SMS")
// 	}

// 	log.Printf("SMS sent successfully to %s", phone)
// 	return nil
// }

// func SendtextMSGhallTicket(ctx context.Context, client *ent.Client, successmsg string) error {
// 	// Check if user is nil
// 	// if user == nil {
// 	// 	return errors.New("User is nil")
// 	// }

// 	// Retrieve the mobile number from the user
// 	mobileNumber := user.Mobile
// 	// //password := user.Password
// 	// empname := user.EmployeeName
// 	// empid := user.EmployeeID
// 	if mobileNumber == "" {
// 		return errors.New("User's mobile number not found")
// 	}

// 	// Construct the SMS message
// 	msg := "Dear "+empname+" - "+strconv.FormatInt(empid, 10)+", User ID was successfully registered. For details check email-DOPExam-INDIAPOST"
// 	// Set other SMS parameters
// 	phone := mobileNumber
// 	templateID := "1007033293383430282"
// 	entityID := "1001081725895192800"
// 	appName := "IBC"

// 	// Trigger the SMS
// 	err := sendSMSs(msg, phone, templateID, entityID, appName)
// 	if err != nil {
// 		log.Printf("Failed to send SMS: %v", err)
// 		return fmt.Errorf("failed to send SMS")
// 	}

// 	nerr := SendEMailuserreg(user, empname,empid)
// 	if nerr != nil {
// 		return fmt.Errorf("Failed to send E-Mail : %v", user.EmailID)
// 	}
// 	log.Printf("SMS sent successfully to %s", phone)
// 	log.Printf("Email sent successfully to %s", user.EmailID)
// 	return nil
// }

/* func SendSMSGetAdminPassword(ctx context.Context, client *ent.Client, user *ent.AdminMaster) (int32, error) {
	// Check if user is nil
	if user == nil {
		return 400, errors.New(" user object is blank")
	}

	// Retrieve the mobile number from the user
	mobileNumber := user.Mobile
	password := user.Password
	usernam := user.UserName
	if mobileNumber == "" {
		return 422, errors.New(" user's mobile number is blank")
	}

	// Construct the SMS message
	msg := "Dear " + usernam + ", Password is " + password + "-DOPExam-INDIAPOST"
	// Set other SMS parameters
	phone := mobileNumber
	templateID := "1007080791170637735"
	entityID := "1001081725895192800"
	appName := "DOP Exam"

	// Trigger the SMS
	smsstatus, err := util.SendSMSs(msg, phone, templateID, entityID, appName)
	var apiresponse string
	var apiresponsedescription string

	if err != nil {
		if smsstatus == 400 {
			apiresponse = "Failed"
			apiresponsedescription = err.Error()
		}
	} else {
		apiresponse = "Success"
		apiresponsedescription = "SMS sent successfully"
	}

	_, err = client.SmsEmailLog.Create().
		SetType("SMS").
		SetMobileEmail(phone).
		SetUserName(usernam).
		SetEventCode("32").
		SetEventDescription("Admin Get/Reset Password").
		SetApiResponse(apiresponse).
		SetApiResponseDescription(apiresponsedescription).
		SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
		Save(ctx)
	if err != nil {
		return 400, fmt.Errorf("failed to create smsEmailLog: %v", err)
	}

	return 200, nil
} */

/* func SendSMSGetUserPassword(ctx context.Context, client *ent.Client, user *ent.UserMaster) (int32, error) {
	// Check if user is nil
	if user == nil {
		return 400, errors.New(" user object is blank")
	}

	// Retrieve the mobile number from the user
	mobileNumber := user.Mobile
	password := user.Password
	usernam := user.UserName
	if mobileNumber == "" {
		return 422, errors.New(" user's mobile number is blank")
	}

	// Construct the SMS message
	msg := "Dear " + usernam + ", Password is " + password + "-DOPExam-INDIAPOST"
	// Set other SMS parameters
	phone := mobileNumber
	templateID := "1007080791170637735"
	entityID := "1001081725895192800"
	appName := "DOP Exam"

	// Trigger the SMS
	smsstatus, err := util.SendSMSs(msg, phone, templateID, entityID, appName)
	var apiresponse string
	var apiresponsedescription string

	if err != nil {
		if smsstatus == 400 {
			apiresponse = "Failed"
			apiresponsedescription = err.Error()
		}
	} else {
		apiresponse = "Success"
		apiresponsedescription = "SMS sent successfully"
	}

	_, err = client.SmsEmailLog.Create().
		SetType("SMS").
		SetMobileEmail(phone).
		SetUserName(usernam).
		SetEventCode("32").
		SetEventDescription("Admin Get/Reset Password").
		SetApiResponse(apiresponse).
		SetApiResponseDescription(apiresponsedescription).
		SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
		Save(ctx)
	if err != nil {
		return 400, fmt.Errorf("failed to create smsEmailLog: %v", err)
	}

	return 200, nil
} */

/* func SendtextMSGApplicationSubmit(ctx context.Context, client *ent.Client, user *ent.AdminMaster) error {
	// Check if user is nil
	if user == nil {
		return errors.New(" user is nil")
	}

	// Retrieve the mobile number from the user
	mobileNumber := user.Mobile
	password := user.Password
	if mobileNumber == "" {
		return errors.New(" user's mobile number not found")
	}

	// Construct the SMS message
	msg := "Dear Customer, OTP for IBC verification is " + password + " please do not share it with anyone - INDPOST"
	// Set other SMS parameters
	phone := mobileNumber
	templateID := "1007344609998507114"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	statussms, err := util.SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(statussms)
		log.Printf("Failed to send SMS: %v", err)
		return fmt.Errorf("failed to send SMS")
	}

	nerr := SendEMailladmin(user, password)
	if nerr != nil {
		return fmt.Errorf(" failed to send E-Mail : %v", user.EmailID)
	}
	log.Printf("SMS sent successfully to %s", phone)
	log.Printf("Email sent successfully to %s", user.EmailID)
	return nil
} */

/* func SendtextMSGchangePassword(ctx context.Context, client *ent.Client, newUser *ent.UserMaster) (string, error) {
	// Check if username is not null
	if newUser.UserName == "" {
		return "", errors.New(" username is null")
	}

	// Retrieve the user from the database based on the username
	user, err := client.UserMaster.Query().Where(usermaster.UserNameEQ(newUser.UserName)).Only(ctx)
	if err != nil {
		return "", fmt.Errorf(" failed to retrieve user: %v", err)
	}

	// Retrieve the mobile number from the retrieved user
	mobileNumber := user.Mobile
	empname := user.EmployeeName
	empid := user.EmployeeID
	password := user.Password
	if mobileNumber == "" {
		return "", errors.New(" user's mobile number not found")
	}

	// Construct the SMS message
	msg := "Dear " + empname + " - " + strconv.FormatInt(empid, 10) + ", New password for the DOP Exam portal is " + password + " - DOPExam-INDIAPOST"

	// Set other SMS parameters
	phone := mobileNumber
	templateID := "1007247787940708106" // Modify with your actual template ID
	entityID := "1001081725895192800"   // Modify with your actual entity ID
	appName := "IBC"                    // Modify with your actual application name

	// Trigger the SMS
	statussms, err := util.SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(statussms)
		log.Printf("Failed to send SMS: %v", err)
		return "", fmt.Errorf(" failed to send SMS")
	}

	log.Printf(" sms sent successfully to %s", phone)

	return "SMS sent successfully.", nil
} */
