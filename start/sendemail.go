package start

// "context"
// "errors"
// "fmt"
// "log"
// "recruit/ent"
// "strconv"
// "time"
//"recruit/ent/usermaster"
//"strconv"

/*
	 func SendEMaill(newUser *ent.UserMaster, password string) error {

		// Retrieve the e-Mail ID from the retrieved user
		emailid := newUser.EmailID
		if emailid == "" {
			return errors.New(" user's e-Mail ID not found")
		}

		// save OTP into new variable
		// Construct the e-mail message and parameters
		epass := password
		body := "Dear Customer,your password is : " + password + ", please do not share it with anyone - INDPOST"
		emailsubject := "Password for Online Departmental Examination"

		// Trigger the mail
		_, err := sendMail(emailid, emailsubject, body)
		if err != nil {
			log.Printf("Failed to send e-Mail: %v", err)
			return fmt.Errorf("failed to send password to email")
		}

		log.Printf("Password sent successfully. User: %s, password: %s", newUser.UserName, epass)

		return nil
	}
*/
/* func SendEMailuserreg(newUser *ent.UserMaster, empname string, empid int64) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New(" user's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	epass := empname
	epid := strconv.FormatInt(empid, 10)
	body := "Dear " + epass + " - " + epid + ", User ID was successfully registered. For details check email-DOPExam-INDIAPOST"
	emailsubject := "Password for Online Departmental Examination"

	// Trigger the mail
	_, err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send password to email")
	}

	log.Printf(" User ID was successfully. User: %s, employeeid: %s", newUser.UserName, epid)

	return nil
} */

/* func SendEMailladmin(newUser *ent.AdminMaster, password string) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New(" user's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	epass := password
	body := "Dear Customer,your password is : " + password + ", please do not share it with anyone - INDPOST"
	emailsubject := "Password for Online Departmental Examination"

	// Trigger the mail
	_, err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send password to email")
	}

	log.Printf("Password sent successfully. User: %s, password: %s", newUser.UserName, epass)

	return nil
}
*/
/*
	 func SendEMailll(newUser *ent.UserMaster, newPassword string) error {

		// Retrieve the e-Mail ID from the retrieved user
		emailid := newUser.EmailID
		if emailid == "" {
			return errors.New(" user's e-Mail ID not found")
		}

		// save OTP into new variable
		// Construct the e-mail message and parameters
		epass := newPassword
		body := "Dear Customer,your new password is : " + newPassword + ", please do not share it with anyone - INDPOST"
		emailsubject := " New Password for Online Departmental Examination"

		// Trigger the mail
		_, err := sendMail(emailid, emailsubject, body)
		if err != nil {
			log.Printf("Failed to send e-Mail: %v", err)
			return fmt.Errorf("failed to send new password to email")
		}

		log.Printf("Password sent successfully. User: %s, password: %s", newUser.UserName, epass)

		return nil
	}
*/
/* func SendEMailAdminOTP(newUser *ent.AdminMaster, otp string) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New(" user's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	eotp := otp
	body := "Dear " + newUser.EmployeeName + " for user name: " + newUser.UserName + ", OTP for Online Departmental Verification is " + eotp + ", valid for five Mins, please do not share it with anyone - INDIAPOST"
	emailsubject := "Admin Login - OTP for Online Departmental Examination"

	// Trigger the mail
	_, err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send OTPs to email")
	}

	log.Printf("OTP sent successfully. User: %s, OTP: %s", newUser.UserName, eotp)

	return nil
} */

/* func SendEMailUserOTP(newUser *ent.AdminMaster, otp string) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New(" user's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	eotp := otp
	body := "Dear " + newUser.EmployeeName + " for user name: " + newUser.UserName + ", OTP for Online Departmental Exam Portal Registration  is " + eotp + ", valid for five Mins, please do not share it with anyone - INDIAPOST	Caution: Giving false information and unauthorised usage may lead to legal action.	Note : 	For any technical issue, kindly send an email to cept@indiapost.gov.in. 	For any administrative issue/clarification, kindly take it up with your higher authority. 	This is an automated email don't reply to this email."

	emailsubject := "Candidate Login - OTP for Online Departmental Examination"

	// Trigger the mail
	_, err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send OTPs to email")
	}

	log.Printf("OTP sent successfully. User: %s, OTP: %s", newUser.UserName, eotp)

	return nil
} */

/* func SendEmailCandidateOTP(ctx context.Context, client *ent.Client, newUser *ent.UserMaster) (int32, int32, error) {
	// Generate OTP
	otpEmail := GenerateOTP()

	// Convert OTP to string
	stringOTPEmail := strconv.Itoa(int(otpEmail))

	emailid := newUser.EmailID
	if emailid == "" {
		return otpEmail, 422, errors.New(" user's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters

	body := "Dear Candidate " + newUser.EmployeeName + "  , Employee ID: " + newUser.UserName + " .\n OTP for DOP Departmental Exam portal registration is " + stringOTPEmail + ".\n  This OTP will be valid only for four minutes.\n\n	Caution: \n Giving false information and unauthorised usage may lead to legal action.\n\n 	Note :\n 	For any technical issue, kindly send an email to cept@indiapost.gov.in. \n For any administrative issue/clarification, kindly take it up with your higher authority.\n 	This is an automated email don't reply to this email."

	emailsubject := "OTP for Candidate Online Departmental Examination portal Registration"

	var apiresponse string
	var apiresponsedescription string
	var status int32
	// Trigger the mail
	_, err := sendMail(emailid, emailsubject, body)
	//fmt.Println("emailstatus", emailstatus)
	if err != nil {
		status = 400
		apiresponse = "Failed"
		apiresponsedescription = err.Error()
	} else {
		status = 200
		apiresponse = "Success"
		apiresponsedescription = "SMS sent successfully"
	}

	_, err1 := client.SmsEmailLog.Create().
		SetType("Email").
		SetMobileEmail(emailid).
		SetUserName(newUser.UserName).
		SetEventCode("1").
		SetEventDescription("Registration Candidate OTP Authentication").
		SetApiResponse(apiresponse).
		SetApiResponseDescription(apiresponsedescription).
		SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
		Save(ctx)
	if err1 != nil {
		return 0, 400, err1
	}
	if status == 400 {
		return 0, 400, err
	} else {
		return otpEmail, 200, nil
	}
} */

/* func SendeAdminMailll(newUser *ent.AdminMaster, newPassword string) error {

	// Retrieve the e-Mail ID from the retrieved user
	emailid := newUser.EmailID
	if emailid == "" {
		return errors.New(" user's e-Mail ID not found")
	}

	// save OTP into new variable
	// Construct the e-mail message and parameters
	epass := newPassword
	body := "Dear Customer,your new password is : " + newPassword + ", please do not share it with anyone - INDPOST"
	emailsubject := " New Password for Online Departmental Examination"

	// Trigger the mail
	_, err := sendMail(emailid, emailsubject, body)
	if err != nil {
		log.Printf("Failed to send e-Mail: %v", err)
		return fmt.Errorf("failed to send new password to email")
	}

	log.Printf("Password sent successfully. User: %s, password: %s", newUser.UserName, epass)

	return nil
} */
