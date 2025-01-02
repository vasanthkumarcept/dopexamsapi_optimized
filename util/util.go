package util

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	cry "crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	//"io/ioutil"
	"log"
	//"math/rand"
	"net/http"
	"os"

	//"regexp"
	"strconv"
	"strings"

	"recruit/ent"
	"recruit/ent/adminmaster"
	"recruit/ent/employeemaster"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/exam_applications_gdspa"
	"recruit/ent/exam_applications_gdspm"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_pmpa"
	"recruit/ent/exam_applications_ps"

	//"recruit/ent/examcategorydisabilitymapping"
	//"recruit/ent/exampostmapping"
	"recruit/ent/recommendationsgdspaapplications"
	"recruit/ent/recommendationsgdspmapplications"
	"recruit/ent/recommendationsipapplications"
	"recruit/ent/recommendationsmtspmmgapplications"
	"recruit/ent/recommendationspmpaapplications"
	"recruit/ent/recommendationspsapplications"
	"recruit/ent/usermaster"
	ca_reg "recruit/payloadstructure/candidate_registration"

	//"recruit/ent/admin_master"

	//"recruit/start"
	"time"

	"github.com/gin-gonic/gin"
)

var Remarks string
var emptyObject = struct{}{}

func getEnvVar() string {
	return os.Getenv("USER_ERROR_REMARKS")
}

var Action string = ""
var errRemark string = ""

// pkcs7Padding adds padding to the plaintext as per the PKCS7 standard
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// pkcs7Unpadding removes padding from the plaintext as per the PKCS7 standard
func pkcs7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("data is empty")
	}
	unpadding := int(data[length-1])
	if unpadding > length || unpadding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding")
	}
	return data[:(length - unpadding)], nil
}
func AdminData(client *ent.Client, userid string) *ent.AdminMaster {
	fmt.Println("userid", userid)
	admin, err := client.AdminMaster.
		Query().
		Where(adminmaster.UserNameEQ(userid), adminmaster.StatussEQ("active")).
		Only(context.Background())
	if err != nil {
		if admin == nil {
			return nil
		}
		return nil
	}
	return admin
}

// encrypt encrypts the plaintext using AES encryption with CBC mode and returns a base64 encoded string
func Encrypt(plainText string) (string, error) {
	key := "a very very very very secret key" // 32 bytes
	iv := "16bytes -iv- for"                  // 16 bytes

	// Ensure key and IV lengths are correct
	keyBytes := []byte(key)
	ivBytes := []byte(iv)
	if len(keyBytes) != 32 {
		fmt.Printf("Encrypt - key length must be 32 bytes")
		return "", fmt.Errorf("system error Contact CEPT")

	}
	if len(ivBytes) != aes.BlockSize {
		fmt.Printf("Encrypt - IV length must be 16 bytes")
		return "", fmt.Errorf("system error Contact CEPT")
	}

	// Create the AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		fmt.Printf("Encrypt - failed to create AES cipher block: %v", err)
		return "", fmt.Errorf("system error Contact CEPT")
	}

	// Apply PKCS7 padding
	plaintextBytes := []byte(plainText)
	paddedText := pkcs7Padding(plaintextBytes, aes.BlockSize)

	// Create CBC mode encrypter
	mode := cipher.NewCBCEncrypter(block, ivBytes)

	// Encrypt the data
	encrypted := make([]byte, len(paddedText))
	mode.CryptBlocks(encrypted, paddedText)

	// Return the encrypted data as base64 string
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// decrypt decrypts the base64 encoded ciphertext using AES decryption with CBC mode
func Decrypt(encryptedText string) (string, error) {
	key := "a very very very very secret key" // 32 bytes
	iv := "16bytes -iv- for"                  // 16 bytes

	// Ensure key and IV lengths are correct
	keyBytes := []byte(key)
	ivBytes := []byte(iv)
	if len(keyBytes) != 32 {
		fmt.Printf("Decrypt - key length must be 32 bytes")
		return "", fmt.Errorf("system error Contact CEPT")

	}
	if len(ivBytes) != aes.BlockSize {
		fmt.Printf("Decrypt - IV length must be 16 bytes")
		return "", fmt.Errorf("system error Contact CEPT")
	}

	// Decode the base64 encoded encrypted text
	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		fmt.Printf("Decrypt -Failed to decode base64: %v\n", err)
		return "", fmt.Errorf("system error Contact CEPT")
	}

	// Create the AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		fmt.Printf("Decrypt - failed to create AES cipher block: %v\n", err)
		return "", fmt.Errorf("system error Contact CEPT")
	}

	// Check if the block size is correct
	if len(cipherText) < aes.BlockSize {
		fmt.Printf("Decrypt - cipherText too short")
		return "", fmt.Errorf("system error contact CEPT")
	}

	// CBC mode always works in whole blocks
	if len(cipherText)%aes.BlockSize != 0 {
		fmt.Printf("Decrypt - cipherText is not a multiple of the block size")
		return "", fmt.Errorf("system error contact CEPT")
	}

	// Create a new CBC decrypter
	mode := cipher.NewCBCDecrypter(block, ivBytes)

	// Decrypt the ciphertext
	mode.CryptBlocks(cipherText, cipherText)

	// Remove padding
	plainText, err := pkcs7Unpadding(cipherText)
	if err != nil {
		fmt.Printf("Decrypt - Failed to remove padding: %v\n", err)
		return "", fmt.Errorf("system error Contact CEPT")

	}

	// Convert decrypted byte array to string
	return string(plainText), nil
}

func GenerateApplicationNumber(client *ent.Client, employeeID int64, examYear string, examGenCode string) (string, error) {
	nextApplicationNumber, err := getNextApplicationNumberFromDatabase(client, examGenCode, examYear, employeeID)
	if err != nil {
		return "", err
	}

	// Get the current year
	currentYear := time.Now().Year()

	// Format the application number as "IPYYYYXXXXXX"
	applicationNumber := fmt.Sprintf("%s%d%06d", examGenCode, currentYear, nextApplicationNumber)

	return applicationNumber, nil
}

func getNextApplicationNumberFromDatabase(client *ent.Client, examGenCode string, examYear string, employeeID int64) (int64, error) {
	ctx := context.TODO()
	if examGenCode == "IP" {
		lastApplication, err := client.Exam_Applications_IP.
			Query().
			Order(ent.Desc(exam_applications_ip.FieldID)).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return 100001, nil
			}
			return 0, fmt.Errorf("failed to retrieve last application: %v", err)
		}
		return lastApplication.ID + 1, nil
	} else if examGenCode == "GDSPA" {
		lastApplication, err := client.Exam_Applications_GDSPA.
			Query().
			Order(ent.Desc(exam_applications_gdspa.FieldID)).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return 100001, nil
			}
			return 0, fmt.Errorf("failed to retrieve last application: %v", err)
		}
		return lastApplication.ID + 1, nil
	} else if examGenCode == "GDSPM" {
		lastApplication, err := client.Exam_Applications_GDSPM.
			Query().
			Order(ent.Desc(exam_applications_gdspm.FieldID)).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return 100001, nil
			}
			return 0, fmt.Errorf("failed to retrieve last application: %v", err)
		}
		return lastApplication.ID + 1, nil
	} else if examGenCode == "MTSPM" {
		lastApplication, err := client.Exam_Application_MTSPMMG.
			Query().
			Order(ent.Desc(exam_application_mtspmmg.FieldID)).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return 100001, nil
			}
			return 0, fmt.Errorf("failed to retrieve last application: %v", err)
		}
		return lastApplication.ID + 1, nil
	} else if examGenCode == "PMPA" {
		lastApplication, err := client.Exam_Applications_PMPA.
			Query().
			Order(ent.Desc(exam_applications_pmpa.FieldID)).
			First(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				return 100001, nil
			}
			return 0, fmt.Errorf("failed to retrieve last application: %v", err)
		}
		return lastApplication.ID + 1, nil
	} else if examGenCode == "PS" {
		lastApplication, err := client.Exam_Applications_PS.
			Query().
			Order(ent.Desc(exam_applications_ps.FieldID)).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return 100001, nil
			}
			return 0, fmt.Errorf("failed to retrieve last application: %v", err)
		}
		return lastApplication.ID + 1, nil
	} else {
		return 0, fmt.Errorf("invalid examGenCode")
	}

}
func GetCtxTimeOut() time.Duration {
	timeOut, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	return time.Duration(timeOut) * time.Second
}

// Middleware to check DB connection
func CheckDatabaseConnection(client *ent.Client) error {
	// Attempt to query something simple from the database to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), GetCtxTimeOut())
	defer cancel()
	err := client.Debug().Schema.Create(ctx)
	if err != nil {
		return err
	}
	return nil
}

/*
	 func GenerateOTP() int32 {
		rand.Seed(time.Now().UnixNano())
		min := 100000
		max := 999999
		otp := rand.Intn(max-min+1) + min
		return int32(otp)
	}
*/
func GenerateOTP() int32 {
	var n uint32
	binary.Read(cry.Reader, binary.LittleEndian, &n)
	otp := 100000 + (n % 900000)
	return int32(otp)
}

//	func GenerateNewOTP() int32 {
//		rand.Seed(time.Now().UnixNano())
//		min := 100000
//		max := 999999
//		otp := rand.Intn(max-min+1) + min
//		return int32(otp)
//	}
func GenerateNewOTP() int32 {
	var n uint32
	binary.Read(cry.Reader, binary.LittleEndian, &n)
	otp := 100000 + (n % 900000)
	return int32(otp)
}

func GenerateEmailOTP() int32 {
	var n uint32
	binary.Read(cry.Reader, binary.LittleEndian, &n)
	otp := 100000 + (n % 900000)
	return int32(otp)
}

/* func GenerateEmailOTP() int32 {
	rand.Seed(time.Now().UnixNano())
	min := 500001
	max := 999999
	otp := rand.Intn(max-min+1) + min
	return int32(otp)
} */

func GenerateEmailNewOTP() int32 {
	var n uint32
	binary.Read(cry.Reader, binary.LittleEndian, &n)
	otp := 100000 + (n % 900000)
	return int32(otp)
}

/*
	 func GenerateEmailNewOTP() int32 {
		rand.Seed(time.Now().UnixNano())
		min := 100000
		max := 499999
		otp := rand.Intn(max-min+1) + min
		return int32(otp)
	}
*/
func SystemLogErrorNew(client *ent.Client, action string, err string, userName string) error {
	username, _ := Decrypt(userName)

	_, logerr1 := client.ErrorLogs.Create().
		SetUserdetails(username).
		SetRemarks(err).
		SetAction(action).
		Save(context.Background())

	if logerr1 != nil {
		fmt.Println("Error in updating ErrorLogs table", logerr1.Error())
		return logerr1
	}

	return nil
}

func SystemLogError(client *ent.Client, action string, err string) error {
	_, logerr1 := client.ErrorLogs.Create().
		SetRemarks(err).
		SetAction(action).
		Save(context.Background())

	if logerr1 != nil {
		fmt.Println("Error in updating ErrorLogs table", logerr1.Error())
		return logerr1
	}

	return nil
}
func LogErrorNew(client *ent.Client, logdata ca_reg.LogData, err error) error {
	fmt.Println("Inside Log")
	ua := AdminData(client, logdata.Userid)
	um := UserData(client, logdata.Userid)
	var userdetails string
	var useruniqueid int64
	if ua == nil && um == nil {
		fmt.Println("NO USER FOUND")
		userdetails = "INVALID USER SENT IN LOG DATA"
		useruniqueid = 404
		//return fmt.Errorf("NO USER FOUND ")
	}

	if ua != nil {
		userdetails = ua.UserName
		useruniqueid = ua.ID
	}
	if um != nil {
		userdetails = um.UserName
		useruniqueid = um.ID
	}
	log.Println("userdetails", userdetails, useruniqueid)
	fmt.Println(logdata)
	_, logerr1 := client.ErrorLogs.Create().
		SetUserid(logdata.Userid).
		SetUsertype(logdata.Usertype).
		SetRemarks(logdata.Remarks + err.Error()).
		SetAction(logdata.Action).
		SetIpaddress(logdata.Ipaddress).
		//SetDevicetype(logdata.Devicetype).
		SetOs(logdata.Os).
		SetBrowser(logdata.Browser).
		SetLatitude(logdata.Latitude).
		SetLongitude(logdata.Longitude).
		SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
		SetUserdetails(userdetails).
		SetUniqueid(useruniqueid).
		Save(context.Background())

	if logerr1 != nil {
		fmt.Println("Error in updating ErrorLogs table", logerr1.Error())
		return logerr1
	}

	return nil
}
func LoggerNew(client *ent.Client, logdata ca_reg.LogData) error {
	ua := AdminData(client, logdata.Userid)
	um := UserData(client, logdata.Userid)
	//var userdetails string
	var useruniqueid int64

	if ua != nil {
		//userdetails = ua.UserName
		useruniqueid = ua.ID
	}
	if um != nil {
		//userdetails = um.UserName
		useruniqueid = um.ID
	}
	_, logerr1 := client.Logs.Create().
		SetUserid(logdata.Userid).
		SetUniqueid(useruniqueid).
		SetUsertype(logdata.Usertype).
		SetRemarks(logdata.Action + " successful").
		SetAction(logdata.Action).
		SetIpaddress(logdata.Ipaddress).
		//SetDevicetype(logdata.Devicetype).
		SetOs(logdata.Os).
		SetBrowser(logdata.Browser).
		SetLatitude(logdata.Latitude).
		SetLongitude(logdata.Longitude).
		SetEventtime(time.Now().UTC().Truncate(24 * time.Hour)).
		SetUniqueid(useruniqueid).
		Save(context.Background())

	if logerr1 != nil {
		return logerr1
	}
	return nil
}

// func AdminData(client *ent.Client, userid string) *ent.AdminMaster {
// 	admin, err := client.AdminMaster.
// 		Query().
// 		Where(adminmaster.UserNameEQ(userid), adminmaster.StatussEQ("active")).
// 		Only(context.Background())
// 	if err != nil {
// 		if admin == nil {
// 			return nil
// 		}
// 		return nil
// 	}
// 	return admin
// }
// func UserData(client *ent.Client, userid string) *ent.UserMaster {
// 	user, err := client.UserMaster.
// 		Query().
// 		Where(usermaster.UserNameEQ(userid), usermaster.StatussEQ("active")).
// 		Only(context.Background())
// 	if err != nil {
// 		return nil
// 	}
// 	return user
// }

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
func AdminDatas(client *ent.Client, userid string) *ent.AdminMaster {
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

// ip exams func
func GetApplicationRemarksByEmployeeID(ctx context.Context, client *ent.Client, employeeID int64) (string, error) {
	application, err := client.Exam_Applications_IP.
		Query().
		Where(
			exam_applications_ip.EmployeeIDEQ(employeeID),
			exam_applications_ip.ApplicationStatusEQ("PendingWithCandidate"),
		).
		Order(ent.Desc(exam_applications_ip.FieldID)).
		First(ctx)

	if err != nil {
		return "", fmt.Errorf("failed to retrieve the PS Group B Application: %v", err)
	}

	return application.AppliactionRemarks, nil
}
func getInputRecordByVacancyYear(inputRecords []*ent.RecommendationsIPApplications, vacancyYear int32) *ent.RecommendationsIPApplications {
	// Find the corresponding input record based on vacancy year
	for _, record := range inputRecords {
		if record.VacancyYear == vacancyYear {
			return record
		}
	}
	return nil
}
func GetIPDivisionsByCircleOfficeID(ctx context.Context, client *ent.Client, circleOfficeID string) ([]*ent.Exam_Applications_IP, error) {
	// Check if the circle office ID exists in the exam_application_ip table.
	exists, err := client.Exam_Applications_IP.
		Query().
		Where(exam_applications_ip.NodalOfficeFacilityIDEQ(circleOfficeID)).
		Exist(ctx)
	if err != nil {
		log.Printf("Failed to query exam_application_ip: %v\n", err)
		return nil, fmt.Errorf("failed to query exam_application_ip: %v", err)
	}
	if !exists {
		log.Printf("Circle office ID does not exist: %s\n", circleOfficeID)
		return nil, fmt.Errorf("circle office ID does not exist")
	}

	// Query the exam_application_ip table for unique records based on the provided conditions.
	applications, err := client.Exam_Applications_IP.
		Query().
		Select(
			exam_applications_ip.FieldReportingOfficeFacilityID,
			exam_applications_ip.FieldReportingOfficeName,
		).
		Where(
			exam_applications_ip.NodalOfficeFacilityIDEQ(circleOfficeID),
			exam_applications_ip.Not(exam_applications_ip.GenerateHallTicketFlag(true)),
			exam_applications_ip.ExamCityCenterCodeIsNil(),
		).
		All(ctx)
	if err != nil {
		log.Printf("Failed to query exam_application_ip: %v\n", err)
		return nil, fmt.Errorf("failed to query exam_application_ip: %v", err)
	}

	// Filter and return distinct records based on reporting office ID and name.
	distinctApplications := make(map[string]*ent.Exam_Applications_IP)
	for _, app := range applications {
		key := app.ReportingOfficeFacilityID
		distinctApplications[key] = app
	}

	result := make([]*ent.Exam_Applications_IP, 0, len(distinctApplications))
	for _, app := range distinctApplications {
		result = append(result, app)
	}

	log.Printf("Retrieved %d distinct divisions for Circle Office ID: %s\n", len(result), circleOfficeID)

	// Log the applications as an array of strings
	appStrings := make([]string, len(result))
	for i, app := range result {
		appStrings[i] = fmt.Sprintf("Reporting Office ID: %s, Reporting Office Name: %s", app.ReportingOfficeFacilityID, app.ReportingOfficeName)
	}
	log.Printf("Applications: %+v\n", appStrings)

	return result, nil
}

func SendEmaill(ctx context.Context, client *ent.Client, email, subject, body string) error {
	// Check if email address is empty
	if email == "" {
		return errors.New("email address is empty")
	}

	// Set other email parameters
	// For example, you may set the email template ID, entity ID, and application name here

	// Trigger the email sending process
	err := sendEmaill(email, subject, body)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return fmt.Errorf("failed to send email")
	}

	log.Printf("Email sent successfully to %s", email)
	return nil
}
func sendEmaill(email, subject, body string) error {
	url := "https://dopverysecure.in/services/mail/send-email"

	requestBody := map[string]string{
		"body":    body,
		"subject": subject,
		"to":      email,
	}

	// Ignore SSL certificate validation
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second}

	fmt.Println("Sending OTP through eMail")
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer((jsonBody)))
	if err != nil {
		fmt.Println("EMAIL IS NOT SENT ERROR IN EMAIL API")
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status Code:", resp.Status)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	fmt.Println("Response Body:", string(respBody))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	if string(respBody) != "Message sent to next instance" {
		return fmt.Errorf("message not sent successfully")
	}
	//fmt.Println("OTP Sent to eMail")
	return nil
}
func SendSMSs(msg, phone, templateID, entityID, appName string) (int32, error) {
	url := "https://api.cept.gov.in/sendsms/api/values/sendsms"

	payload := `{
		"Msg": "` + msg + `",
		"Phone": "` + phone + `",
		"TemplateID": "` + templateID + `",
		"EntityID": "` + entityID + `",
		"AppName": "` + appName + `"
	}`
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {

		return 400, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	//client := http.DefaultClient
	client := &http.Client{
		Timeout: 10 * time.Second, // Ensure you set a reasonable timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		return 400, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 400, fmt.Errorf("failed to read response body: %v", err)
	}
	log.Println("Response body:", string(body))

	if resp.StatusCode != http.StatusOK {
		return 400, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	response := string(body)
	if !strings.Contains(response, "SMS Pushed to NIC Successfully") {
		return 400, fmt.Errorf("failed to send SMS")
	}

	return 200, nil
}
func SendBulkSMS(ctx context.Context, client *ent.Client, phoneNumbers []string, smsTemplates string, candidateData map[string]string) error {
	for _, phone := range phoneNumbers {
		// Replace placeholders in the SMS template with actual values
		message := fmt.Sprintf("Dear  %s-%s-%s, Hall ticket no. %s .For details, check email-DOPExam-INDIAPOST.", candidateData["name"], candidateData["id"], candidateData["appln"], candidateData["hallticket"])
		fmt.Println(message + "The value ")

		// Send SMS message to the current phone number
		err := SendTextMessage(ctx, client, phone, message)
		if err != nil {
			// Handle error
			return err
		}
	}
	return nil
}

func SendTextMessage(ctx context.Context, client *ent.Client, phone, msg string) error {
	// Check if phone number is empty
	if phone == "" {
		return errors.New("phone number is empty")
	}

	// Set other SMS parameters
	templateID := "1007210429836493992"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	status, err := SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(status)
		log.Printf("Failed to send SMS: %v", err)
		return fmt.Errorf("failed to send SMS")
	}

	log.Printf("SMS sent successfully to %s", phone)
	return nil
}

// ps
func SendBulkSMSps(ctx context.Context, client *ent.Client, phoneNumbers []string, smsTemplates string, candidateData map[string]string) error {
	for _, phone := range phoneNumbers {
		// Replace placeholders in the SMS template with actual values
		message := fmt.Sprintf("Dear  %s-%s, Hall ticket no. %s .For details, check email-DOPExam-INDIAPOST.", candidateData["name"], candidateData["id"], candidateData["hallticket"])
		fmt.Println(message + "The value ")

		// Send SMS message to the current phone number
		err := SendTextMessageps(ctx, client, phone, message)
		if err != nil {
			// Handle error
			return err
		}
	}
	return nil
}

func SendTextMessageps(ctx context.Context, client *ent.Client, phone, msg string) error {
	// Check if phone number is empty
	if phone == "" {
		return errors.New("phone number is empty")
	}

	// Set other SMS parameters
	templateID := "1007210429836493992"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	status, err := SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(status)
		log.Printf("Failed to send SMS: %v", err)
		return fmt.Errorf("failed to send SMS")
	}

	log.Printf("SMS sent successfully to %s", phone)
	return nil
}
func GetDivisionID(circleIDStr string) int32 {
	circleID, err := strconv.ParseInt(circleIDStr, 10, 32)
	if err != nil {
		// Handle the error if the conversion fails
		log.Printf("Failed to convert circleID: %v\n", err)
		return 0 // Or whatever default value you prefer
	}
	return int32(circleID)
}
func GetCircleID(circleIDStr string) int32 {
	circleID, err := strconv.ParseInt(circleIDStr, 10, 32)
	if err != nil {
		// Handle the error if the conversion fails
		log.Printf("Failed to convert circleID: %v\n", err)
		return 0 // Or whatever default value you prefer
	}
	return int32(circleID)
}

// Get Admit card details .

/* func GetApplicationsWithHallTicket(client *ent.Client, examCode int32, employeeID int64, exmyear string) (*ent.Exam_Applications_IP, *ent.Exam_Applications_PS, *ent.Exam_Applications_GDSPA, *ent.Exam_Applications_PMPA, *ent.Exam_Applications_GDSPM, *ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	ctx := context.Background()

	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	// Check if exam code is valid
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 400, " -STR001", false, errors.New("please provide a valid exam code")
	}

	if examCode == 2 {
		application, err := tx.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.EmployeeIDEQ(employeeID),
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(exmyear),

				exam_applications_ip.HallTicketNumberNEQ(""),
				exam_applications_ip.CenterIdNEQ(0),
				exam_applications_ip.StatusEQ("active"),
			).
			WithExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR002", false, fmt.Errorf("no application  found for this employee %d in IP Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR003", false, err
			}
		}
		// if recommendedstatus = "not recommended" show as "Candidate not recommended for this exam "
		// else , if HallTicketNumberEQ "" and centideq (null/0 ) "show as "Hall ticket not yet generated "
		// remove RecommendationsIPApplications
		// Fetch the associated RecommendationsIP records matching the employee ID
		recommendations, err := tx.RecommendationsIPApplications.Query().
			Where(recommendationsipapplications.EmployeeIDEQ(employeeID)).
			// put condition to get only details based on exam.applicationid and    reccomm.application_id
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR004", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR005", false, errors.New("no recommendations found for this exam")
		}
		// Assign the fetched recommendations to the application entity
		application.Edges.IPApplicationsRef = recommendations

		return application, nil, nil, nil, nil, nil, 200, "", true, nil
	} else if examCode == 1 {
		// Check if the employee_ID exists in the Exam_Applications_PS table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationps, err := tx.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.EmployeeIDEQ(employeeID),
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(exmyear),

				exam_applications_ps.HallTicketNumberNEQ(""),
				exam_applications_ps.ExamCityCenterCodeNEQ(0),
				exam_applications_ps.StatusEQ("active"),
			).
			WithPSExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR006", false, fmt.Errorf("no application  found for this employee %d in PSGRB Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR007", false, err
			}
		}

		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsPSApplications.Query().
			Where(recommendationspsapplications.EmployeeIDEQ(employeeID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR008", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR009", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationps.Edges.PSApplicationsRef = recommendations

		return nil, applicationps, nil, nil, nil, nil, 200, "", true, nil
	} else if examCode == 4 {
		// Check if the employee_ID exists in the Exam_Applications_GDSPA table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationgdspa, err := tx.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.EmployeeIDEQ(employeeID),
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(exmyear),

				exam_applications_gdspa.HallTicketNumberNEQ(""),
				exam_applications_gdspa.ExamCityCenterCodeNEQ(0),
				exam_applications_gdspa.StatusEQ("active"),
			).
			WithGDSPAExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR010", false, fmt.Errorf("no application  found for this employee %d in GDSPA Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR011", false, err
			}
		}

		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsGDSPAApplications.Query().
			Where(recommendationsgdspaapplications.EmployeeIDEQ(employeeID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR012", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR013", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationgdspa.Edges.GDSPAApplicationsRef = recommendations

		return nil, nil, applicationgdspa, nil, nil, nil, 200, "", true, nil
	} else if examCode == 3 {
		// Check if the employee_ID exists in the Exam_Applications_GDSPA table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationpmpa, err := tx.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.EmployeeIDEQ(employeeID),
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(exmyear),

				exam_applications_pmpa.HallTicketNumberNEQ(""),
				exam_applications_pmpa.ExamCityCenterCodeNEQ(0),
				exam_applications_pmpa.StatusEQ("active"),
			).
			WithPMPAExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, fmt.Errorf("no application  found for this employee %d in PMPA Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR015", false, err
			}
		}

		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsPMPAApplications.Query().
			Where(recommendationspmpaapplications.EmployeeIDEQ(employeeID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR016", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR017", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationpmpa.Edges.PMPAApplicationsRef = recommendations

		return nil, nil, nil, applicationpmpa, nil, nil, 200, "", true, nil
	} else if examCode == 5 {
		// Check if the employee_ID exists in the Exam_Applications_GDSPA table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationmtspmmg, err := tx.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.EmployeeIDEQ(employeeID),
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(exmyear),

				exam_application_mtspmmg.HallTicketNumberNEQ(""),
				exam_application_mtspmmg.ExamCityCenterCodeNEQ(0),
				exam_application_mtspmmg.StatusEQ("active"),
			).
			WithMTSPMMGExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR018", false, fmt.Errorf("no application  found for this employee %d in MTSPMMG Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR019", false, err
			}
		}

		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsMTSPMMGApplications.Query().
			Where(recommendationsmtspmmgapplications.EmployeeIDEQ(employeeID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR020", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR021", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationmtspmmg.Edges.MTSPMMGApplicationsRef = recommendations

		return nil, nil, nil, nil, nil, applicationmtspmmg, 200, "", true, nil
	} else if examCode == 6 {
		// Check if the employee_ID exists in the Exam_Applications_GDSPA table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationgdspm, err := tx.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.EmployeeIDEQ(employeeID),
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(exmyear),

				exam_applications_gdspm.HallTicketNumberNEQ(""),
				exam_applications_gdspm.ExamCityCenterCodeNEQ(0),
				exam_applications_gdspm.StatusEQ("active"),
			).
			WithGDSPMExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR022", false, fmt.Errorf("no application  found for this employee %d in GDSPM Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR023", false, err
			}
		}

		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsGDSPMApplications.Query().
			Where(recommendationsgdspmapplications.EmployeeIDEQ(employeeID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR024", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR025", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationgdspm.Edges.GDSPMApplicationsRef = recommendations

		return nil, nil, nil, nil, applicationgdspm, nil, 200, "", true, nil
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, nil, nil, nil, nil, nil, 500, " -STR027", false, err
	}

	return nil, nil, nil, nil, nil, nil, 422, " -STR026", false, errors.New("invalid exam code")
} */

func GetApplicationsWithHallTicket(client *ent.Client, examCode int32, employeeID int64, examYear string) (*ent.Exam_Applications_IP, *ent.Exam_Applications_PS, *ent.Exam_Applications_GDSPA, *ent.Exam_Applications_PMPA, *ent.Exam_Applications_GDSPM, *ent.Exam_Application_MTSPMMG, int32, string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), GetCtxTimeOut())
	defer cancel()
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, 500, " -TX001", false, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	// Check if exam code is valid
	if examCode == 0 {
		return nil, nil, nil, nil, nil, nil, 400, " -STR001", false, errors.New("please provide a valid exam code")
	}

	if examCode == 2 {
		application, err := tx.Exam_Applications_IP.
			Query().
			Where(
				exam_applications_ip.EmployeeIDEQ(employeeID),
				exam_applications_ip.ExamCodeEQ(examCode),
				exam_applications_ip.ExamYearEQ(examYear),
				exam_applications_ip.StatusEQ("active"),
				exam_applications_ip.HallTicketNumberNEQ(""),
				exam_applications_ip.CenterCodeNotNil(),
			).
			WithExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR002", false, fmt.Errorf("no application  found for this employee %d for IP Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR003", false, err
			}
		}
		fmt.Println("application", application)
		// if recommendedstatus = "not recommended" show as "Candidate not recommended for this exam "
		// else , if HallTicketNumberEQ "" and centideq (null/0 ) "show as "Hall ticket not yet generated "
		// remove RecommendationsIPApplications
		// Fetch the associated RecommendationsIP records matching the employee ID
		if application.RecommendedStatus == "not recommended" {
			return nil, nil, nil, nil, nil, nil, 422, " -STR004", false, errors.New("candidate not recommended for this exam ")

		} else if application.HallTicketNumber == "" && application.ExamCityCenterCode == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR005", false, errors.New("hall ticket not yet generated ")
		}
		recommendations, err := tx.RecommendationsIPApplications.Query().
			Where(recommendationsipapplications.EmployeeIDEQ(employeeID),
				recommendationsipapplications.ApplicationIDEQ(application.ID)).
			// put condition to get only details based on exam.applicationid and    reccomm.application_id
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR006", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR007", false, errors.New("no recommendations found for this exam")
		}
		// Assign the fetched recommendations to the application entity
		application.Edges.IPApplicationsRef = recommendations

		return application, nil, nil, nil, nil, nil, 200, "", true, nil
	} else if examCode == 1 {
		// Check if the employee_ID exists in the Exam_Applications_PS table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationps, err := tx.Exam_Applications_PS.
			Query().
			Where(
				exam_applications_ps.EmployeeIDEQ(employeeID),
				exam_applications_ps.ExamCodeEQ(examCode),
				exam_applications_ps.ExamYearEQ(examYear),
				exam_applications_ps.StatusEQ("active"),
				exam_applications_ps.HallTicketNumberNEQ(""),
				exam_applications_ps.CenterCodeNotNil(),
			).
			WithPSExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR008", false, fmt.Errorf("no application  found for this employee %d For  PS Group B Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR009", false, err
			}
		}
		if applicationps.RecommendedStatus == "not recommended" {
			return nil, nil, nil, nil, nil, nil, 422, " -STR010", false, errors.New("candidate not recommended for this exam ")

		} else if applicationps.HallTicketNumber == "" && applicationps.ExamCityCenterCode == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR011", false, errors.New("hall ticket not yet generated ")
		}
		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsPSApplications.Query().
			Where(
				recommendationspsapplications.EmployeeIDEQ(employeeID),
				recommendationspsapplications.ApplicationIDEQ(applicationps.ID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR012", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR013", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationps.Edges.PSApplicationsRef = recommendations

		return nil, applicationps, nil, nil, nil, nil, 200, "", true, nil
	} else if examCode == 4 {
		// Check if the employee_ID exists in the Exam_Applications_GDSPA table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationgdspa, err := tx.Exam_Applications_GDSPA.
			Query().
			Where(
				exam_applications_gdspa.EmployeeIDEQ(employeeID),
				exam_applications_gdspa.ExamCodeEQ(examCode),
				exam_applications_gdspa.ExamYearEQ(examYear),
				exam_applications_gdspa.StatusEQ("active"),
				exam_applications_gdspa.HallTicketNumberNEQ(""),
				exam_applications_gdspa.CenterCodeNotNil(),
			).
			WithGDSPAExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR014", false, fmt.Errorf("no application  found for this employee %d in GDS to PA Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR015", false, err
			}
		}
		if applicationgdspa.RecommendedStatus == "not recommended" {
			return nil, nil, nil, nil, nil, nil, 422, " -STR016", false, errors.New("candidate not recommended for this exam ")

		} else if applicationgdspa.HallTicketNumber == "" && applicationgdspa.ExamCityCenterCode == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR017", false, errors.New("hall ticket not yet generated ")
		}

		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsGDSPAApplications.Query().
			Where(recommendationsgdspaapplications.EmployeeIDEQ(employeeID),
				recommendationsgdspaapplications.ApplicationIDEQ(applicationgdspa.ID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR018", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR019", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationgdspa.Edges.GDSPAApplicationsRef = recommendations

		return nil, nil, applicationgdspa, nil, nil, nil, 200, "", true, nil
	} else if examCode == 3 {
		// Check if the employee_ID exists in the Exam_Applications_GDSPA table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationpmpa, err := tx.Exam_Applications_PMPA.
			Query().
			Where(
				exam_applications_pmpa.EmployeeIDEQ(employeeID),
				exam_applications_pmpa.ExamCodeEQ(examCode),
				exam_applications_pmpa.ExamYearEQ(examYear),
				exam_applications_pmpa.StatusEQ("active"),
				exam_applications_pmpa.HallTicketNumberNEQ(""),
				exam_applications_pmpa.CenterCodeNotNil(),
			).
			WithPMPAExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR020", false, fmt.Errorf("no application  found for this employee %d for Pm to PA Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR021", false, err
			}
		}
		if applicationpmpa.RecommendedStatus == "not recommended" {
			return nil, nil, nil, nil, nil, nil, 422, " -STR022", false, errors.New("candidate not recommended for this exam ")

		} else if applicationpmpa.HallTicketNumber == "" && applicationpmpa.ExamCityCenterCode == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR023", false, errors.New("hall ticket not yet generated ")
		}
		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsPMPAApplications.Query().
			Where(recommendationspmpaapplications.EmployeeIDEQ(employeeID),
				recommendationspmpaapplications.ApplicationIDEQ(applicationpmpa.ID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR024", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR025", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationpmpa.Edges.PMPAApplicationsRef = recommendations

		return nil, nil, nil, applicationpmpa, nil, nil, 200, "", true, nil
	} else if examCode == 5 {
		// Check if the employee_ID exists in the Exam_Applications_GDSPA table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationmtspmmg, err := tx.Exam_Application_MTSPMMG.
			Query().
			Where(
				exam_application_mtspmmg.EmployeeIDEQ(employeeID),
				exam_application_mtspmmg.ExamCodeEQ(examCode),
				exam_application_mtspmmg.ExamYearEQ(examYear),
				exam_application_mtspmmg.StatusEQ("active"),
				exam_application_mtspmmg.HallTicketNumberNEQ(""),
				exam_application_mtspmmg.CenterCodeNotNil(),
			).
			WithMTSPMMGExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR026", false, fmt.Errorf("no application  found for this employee %d for MTS to PM/MG Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR027", false, err
			}
		}
		if applicationmtspmmg.RecommendedStatus == "not recommended" {
			return nil, nil, nil, nil, nil, nil, 422, " -STR028", false, errors.New("candidate not recommended for this exam ")

		} else if applicationmtspmmg.HallTicketNumber == "" && applicationmtspmmg.ExamCityCenterCode == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR029", false, errors.New("hall ticket not yet generated ")
		}
		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsMTSPMMGApplications.Query().
			Where(recommendationsmtspmmgapplications.EmployeeIDEQ(employeeID),
				recommendationsmtspmmgapplications.ApplicationIDEQ(applicationmtspmmg.ID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR030", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR031", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationmtspmmg.Edges.MTSPMMGApplicationsRef = recommendations

		return nil, nil, nil, nil, nil, applicationmtspmmg, 200, "", true, nil
	} else if examCode == 6 {
		// Check if the employee_ID exists in the Exam_Applications_GDSPA table

		// Query the Exam_Applications_PS table to retrieve the applicant details
		applicationgdspm, err := tx.Exam_Applications_GDSPM.
			Query().
			Where(
				exam_applications_gdspm.EmployeeIDEQ(employeeID),
				exam_applications_gdspm.ExamCodeEQ(examCode),
				exam_applications_gdspm.ExamYearEQ(examYear),
				exam_applications_gdspm.StatusEQ("active"),
				exam_applications_gdspm.HallTicketNumberNEQ(""),
				exam_applications_gdspm.CenterCodeNotNil(),
			).
			WithGDSPMExamCentres().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, nil, nil, nil, nil, nil, 422, " -STR032", false, fmt.Errorf("no application  found for this employee %d for GDS to PM/MG Applications ", employeeID)
			} else {
				return nil, nil, nil, nil, nil, nil, 500, " -STR033", false, err
			}
		}

		if applicationgdspm.RecommendedStatus == "not recommended" {
			return nil, nil, nil, nil, nil, nil, 422, " -STR034", false, errors.New("candidate not recommended for this exam ")

		} else if applicationgdspm.HallTicketNumber == "" && applicationgdspm.ExamCityCenterCode == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR035", false, errors.New("hall ticket not yet generated ")
		}
		// Fetch the associated RecommendationsPS records matching the employee ID
		recommendations, err := tx.RecommendationsGDSPMApplications.Query().
			Where(recommendationsgdspmapplications.EmployeeIDEQ(employeeID),
				recommendationsgdspmapplications.ApplicationIDEQ(applicationgdspm.ID)).
			All(ctx)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, 500, " -STR036", false, err
		}

		if len(recommendations) == 0 {
			return nil, nil, nil, nil, nil, nil, 422, " -STR037", false, errors.New("no recommendations found for this exam")
		}

		// Assign the fetched recommendations to the application entity
		applicationgdspm.Edges.GDSPMApplicationsRef = recommendations

		return nil, nil, nil, nil, applicationgdspm, nil, 200, "", true, nil
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, nil, nil, nil, nil, nil, 500, " -STR038", false, err
	}

	return nil, nil, nil, nil, nil, nil, 422, " -STR039", false, errors.New("invalid exam code")
}
func GenerateHallticketNumber(examCode int32, examYear string, categoryCode string, circleID int32 /*regionID int32,*/, divisionID int32, identificationNo int) string {
	// Generate the Hallticket Number based on the provided formatfmt.Sprintf("%d%s%d%d%d%d%04d", examCode, examYear, getFormattedCode(circleID), regionID, getFormattedCode(divisionID), categoryCode, identificationNo)
	hallticketNumber := fmt.Sprintf("%d%s%s%s%d%04d", examCode, examYear, getFormattedCode(circleID) /*regionID,*/, getFormattedCode(divisionID), categoryCode, identificationNo)
	return hallticketNumber
}

func getFormattedCode(code int32) string {
	// Format the code as a string with the required number of digits
	lastTwoDigits := code % 100
	return fmt.Sprintf("%02d", lastTwoDigits)
}

// ps group b
func getPSInputRecordByVacancyYear(inputRecords []*ent.RecommendationsPSApplications, vacancyYear int32) *ent.RecommendationsPSApplications {
	// Find the corresponding input record based on vacancy year
	for _, record := range inputRecords {
		if record.VacancyYear == vacancyYear {
			return record
		}
	}
	return nil
}
func GenerateHallticketNumberPS(examCode int32, examYear string, categoryCode string, circleID string, divisionID string, identificationNo int) string {
	hallticketNumber := fmt.Sprintf("%d%s%s%d%04d", examCode, examYear, getFormattedCodePS(circleID), getFormattedCodePS(divisionID), identificationNo)
	return hallticketNumber
}
func getFormattedCodePS(code string) string {
	// Parse the code as an integer
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		// Handle the error if the code cannot be converted to an integer
		log.Printf("Failed to parse code '%s' as an integer: %v", code, err)
		return "" // Return an empty string or handle the error accordingly
	}

	// Extract the last two digits from the integer
	lastTwoDigits := codeInt % 100

	// Format the last two digits with leading zeros if necessary
	return fmt.Sprintf("%02d", lastTwoDigits)
}

// mtspmmg
func GenerateHallticketNumberMTSPMMG(examCode int32, examYear string, categoryCode string, circleID int32 /*regionID int32,*/, divisionID int32, identificationNo int) string {
	// Generate the Hallticket Number based on the provided formatfmt.Sprintf("%d%s%d%d%d%d%04d", examCode, examYear, getFormattedCode(circleID), regionID, getFormattedCode(divisionID), categoryCode, identificationNo)
	hallticketNumber := fmt.Sprintf("%d%s%s%s%d%04d", examCode, examYear, getFormattedCodeMTSPMMG(circleID) /*regionID,*/, getFormattedCodeMTSPMMG(divisionID), categoryCode, identificationNo)
	return hallticketNumber
}

func getFormattedCodeMTSPMMG(code int32) string {
	// Format the code as a string with the required number of digits
	lastTwoDigits := code % 100
	return fmt.Sprintf("%02d", lastTwoDigits)
}
func SendBulkSMSMtsPm(ctx context.Context, client *ent.Client, phoneNumbers []string, smsTemplates string, candidateData map[string]string) error {
	for _, phone := range phoneNumbers {
		// Replace placeholders in the SMS template with actual values
		message := fmt.Sprintf("Dear  %s-%s, Hall ticket no. %s .For details, check email-DOPExam-INDIAPOST.", candidateData["name"], candidateData["id"], candidateData["hallticket"])
		fmt.Println(message + "The value ")

		// Send SMS message to the current phone number
		err := SendTextMessageMtsPm(ctx, client, phone, message)
		if err != nil {
			// Handle error
			return err
		}
	}
	return nil
}

func SendTextMessageMtsPm(ctx context.Context, client *ent.Client, phone, msg string) error {
	// Check if phone number is empty
	if phone == "" {
		return errors.New("Phone number is empty")
	}

	// Set other SMS parameters
	templateID := "1007210429836493992"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	status, err := SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(status)
		log.Printf("Failed to send SMS: %v", err)
		return fmt.Errorf("failed to send SMS")
	}

	log.Printf("SMS sent successfully to %s", phone)
	return nil
}
func SendBulkSMSgdsPm(ctx context.Context, client *ent.Client, phoneNumbers []string, smsTemplates string, candidateData map[string]string) error {
	for _, phone := range phoneNumbers {
		// Replace placeholders in the SMS template with actual values
		message := fmt.Sprintf("Dear  %s-%s, Hall ticket no. %s .For details, check email-DOPExam-INDIAPOST.", candidateData["name"], candidateData["id"], candidateData["hallticket"])
		fmt.Println(message + "The value ")

		// Send SMS message to the current phone number
		err := SendTextMessagegdsPm(ctx, client, phone, message)
		if err != nil {
			// Handle error
			return err
		}
	}
	return nil
}
func SendTextMessagegdsPm(ctx context.Context, client *ent.Client, phone, msg string) error {
	// Check if phone number is empty
	if phone == "" {
		return errors.New("Phone number is empty")
	}

	// Set other SMS parameters
	templateID := "1007210429836493992"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	statussms, err := SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(statussms)
		log.Printf("Failed to send SMS: %v", err)
		return fmt.Errorf("failed to send SMS")
	}

	log.Printf("SMS sent successfully to %s", phone)
	return nil
}
func SendBulkSMSgdsPa(ctx context.Context, client *ent.Client, phoneNumbers []string, smsTemplates string, candidateData map[string]string) error {
	for _, phone := range phoneNumbers {
		// Replace placeholders in the SMS template with actual values
		message := fmt.Sprintf("Dear  %s-%s, Hall ticket no. %s .For details, check email-DOPExam-INDIAPOST.", candidateData["name"], candidateData["id"], candidateData["hallticket"])
		fmt.Println(message + "The value ")

		// Send SMS message to the current phone number
		err := SendTextMessagegdsPa(ctx, client, phone, message)
		if err != nil {
			// Handle error
			return err
		}
	}
	return nil
}

func SendTextMessagegdsPa(ctx context.Context, client *ent.Client, phone, msg string) error {
	// Check if phone number is empty
	if phone == "" {
		return errors.New("phone number is empty")
	}

	// Set other SMS parameters
	templateID := "1007210429836493992"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	status, err := SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(status)
		log.Printf("Failed to send SMS: %v", err)
		return fmt.Errorf("failed to send SMS")
	}

	log.Printf("SMS sent successfully to %s", phone)
	return nil
}
func GenerateHallticketNumberGDSPA(examCode int32, examYear string, categoryCode string, circleID int32 /*regionID int32,*/, divisionID int32, identificationNo int) string {
	// Generate the Hallticket Number based on the provided formatfmt.Sprintf("%d%s%d%d%d%d%04d", examCode, examYear, getFormattedCode(circleID), regionID, getFormattedCode(divisionID), categoryCode, identificationNo)
	hallticketNumber := fmt.Sprintf("%d%s%s%s%d%04d", examCode, examYear, getFormattedCodeGDSPA(circleID) /*regionID,*/, getFormattedCodeGDSPA(divisionID), categoryCode, identificationNo)
	return hallticketNumber
}

func getFormattedCodeGDSPA(code int32) string {
	// Format the code as a string with the required number of digits
	lastTwoDigits := code % 100
	return fmt.Sprintf("%02d", lastTwoDigits)
}

func SendBulkSMSgdsPmPa(ctx context.Context, client *ent.Client, phoneNumbers []string, smsTemplates string, candidateData map[string]string) error {
	for _, phone := range phoneNumbers {
		// Replace placeholders in the SMS template with actual values
		message := fmt.Sprintf("Dear  %s-%s, Hall ticket no. %s .For details, check email-DOPExam-INDIAPOST.", candidateData["name"], candidateData["id"], candidateData["hallticket"])
		fmt.Println(message + "The value ")

		// Send SMS message to the current phone number
		err := SendTextMessagegdsPmPa(ctx, client, phone, message)
		if err != nil {
			// Handle error
			return err
		}
	}
	return nil
}

func SendTextMessagegdsPmPa(ctx context.Context, client *ent.Client, phone, msg string) error {
	// Check if phone number is empty
	if phone == "" {
		return errors.New("Phone number is empty")
	}

	// Set other SMS parameters
	templateID := "1007210429836493992"
	entityID := "1001081725895192800"
	appName := "IBC"

	// Trigger the SMS
	status, err := SendSMSs(msg, phone, templateID, entityID, appName)
	if err != nil {
		log.Println(status)
		log.Printf("Failed to send SMS: %v", err)
		return fmt.Errorf("failed to send SMS")
	}

	log.Printf("SMS sent successfully to %s", phone)
	return nil
}

func GenerateHallticketNumberPMPA(examCode int32, examYear string, categoryCode string, circleID int32 /*regionID int32,*/, divisionID int32, identificationNo int) string {
	// Generate the Hallticket Number based on the provided formatfmt.Sprintf("%d%s%d%d%d%d%04d", examCode, examYear, getFormattedCode(circleID), regionID, getFormattedCode(divisionID), categoryCode, identificationNo)
	hallticketNumber := fmt.Sprintf("%d%s%s%s%d%04d", examCode, examYear, getFormattedCodePMPA(circleID) /*regionID,*/, getFormattedCodePMPA(divisionID), categoryCode, identificationNo)
	return hallticketNumber
}

func getFormattedCodePMPA(code int32) string {
	// Format the code as a string with the required number of digits
	lastTwoDigits := code % 100
	return fmt.Sprintf("%02d", lastTwoDigits)
}

// Middle Ware

func TokenValidationMiddlewareNew(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		token, err := Decrypt(gctx.Request.Header.Get("UidToken"))
		if err != nil {
			Remarks := "400  error main - TokenValidationMiddleware - Token  tampered " + err.Error()
			MainHandleError(gctx, client, Remarks, " -TO01", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}

		if token == "" {
			Remarks := "500 main - TokenValidationMiddleware - Token  is blank "
			MainHandleError(gctx, client, Remarks, " -TO02", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}
		username, err := Decrypt(gctx.Request.Header.Get("UserName"))
		if err != nil {
			Remarks := "400 main - TokenValidationMiddleware - UserName  tampered "
			MainHandleError(gctx, client, Remarks, " -TO03", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}
		if username == "" {
			Remarks := "422 main - TokenValidationMiddleware - UserName  is blank"
			MainHandleError(gctx, client, Remarks, " -TO04", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}

		var usermasteruserdetails *ent.UserMaster
		var UID string
		var err2 error
		var dbadminerror, dbusererror bool = false, false
		adminuserdetails, err1 := client.AdminMaster.Query().Where(adminmaster.UserNameEQ(username), adminmaster.StatussEQ("active")).Only(context.Background())
		if err1 != nil {
			if ent.IsNotFound(err1) {
				usermasteruserdetails, err2 = client.UserMaster.Query().Where(usermaster.UserNameEQ(username), usermaster.StatussEQ("active"), usermaster.StatusEQ(true)).Only(context.Background())
				if err2 != nil {
					if !ent.IsNotFound(err2) {
						dbusererror = true
					}
				}
			} else {
				dbadminerror = true
			}
		}
		if dbadminerror || dbusererror {
			Remarks := "500 DB error " + err1.Error() + err2.Error()
			MainHandleError(gctx, client, Remarks, " -TO05", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}
		if err1 != nil && err2 != nil {
			Remarks := "401 no active admin/user found"
			MainHandleError(gctx, client, Remarks, " -TO06", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}
		if adminuserdetails != nil {
			UID = adminuserdetails.UidToken
		}
		if usermasteruserdetails != nil {
			UID = usermasteruserdetails.UidToken
		}

		if token != UID {
			Remarks := "422 main - TokenValidationMiddlewareUID - Invalid Token"
			MainHandleError(gctx, client, Remarks, " -TO07", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}
		gctx.Next()
	}
}

func TokenValidationMiddlewareUID(client *ent.Client) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		token, err := Decrypt(gctx.Request.Header.Get("UserName"))
		if err != nil {
			Remarks := "400  error main - TokenValidationMiddlewareUID - Token  tampered " + err.Error()
			MainHandleError(gctx, client, Remarks, " -TO01", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}

		if token == "" {
			Remarks := "500 main - TokenValidationMiddlewareUID - Token  is blank "
			MainHandleError(gctx, client, Remarks, " -TO02", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}

		if token != "TestingDone" {
			Remarks := "422 main - TokenValidationMiddlewareUID - Invalid Token"
			MainHandleError(gctx, client, Remarks, " -TO07", gctx.GetHeader("UserName"))
			gctx.Abort()
			return
		}
		gctx.Next()
	}
}

func MainHandleError(ctx *gin.Context, client *ent.Client, Remarks string, handleError string, username string) {
	UserErrorRemarks := getEnvVar()
	Action = "400"
	errRemark = Remarks + handleError
	SystemLogErrorNew(client, Action, errRemark, username)
	handleString := UserErrorRemarks + handleError
	HandleError(ctx, 400, handleString)
	return
}

func HandleError(ctx *gin.Context, status int, handleString string) {
	emptyObject := struct{}{}
	{
		ctx.JSON(status, gin.H{
			"success":    true,
			"message":    handleString,
			"data":       emptyObject,
			"dataexists": false,
		})
	}
}
func UserDatas(client *ent.Client, userid string) *ent.UserMaster {
	user, err := client.UserMaster.
		Query().
		Where(usermaster.UserNameEQ(userid), usermaster.StatussEQ("active")).
		Only(context.Background())
	if err != nil {
		return nil
	}
	return user
}
func EmployeeData(client *ent.Client, userid string) *ent.EmployeeMaster {
	empid, _ := strconv.ParseInt(userid, 10, 64)
	user, err := client.EmployeeMaster.
		Query().
		Where(employeemaster.EmployeeIDEQ(empid), employeemaster.StatussEQ("active")).
		Only(context.Background())
	if err != nil {
		return nil
	}
	return user
}
