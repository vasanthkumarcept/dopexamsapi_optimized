package candidate_registation

import (
	"time"
)

type EdgesLogdata struct {
	LogData []LogData `json:"LogData"`
}

type EdgesCircleData struct {
	CircleData []CircleData `json:"CirclePrefRef"`
	LogData    []LogData    `json:"LogData"`
}

type EdgesIPApplicatinData struct {
	LogData         []LogData          `json:"LogData"`
	ApplicationData []ApplicationDataV `json:"IPApplicationsRef"`
	CircleData      []CircleData       `json:"CirclePrefRef"`
}
type EdgesIPApplicatinDataNA struct {
	LogData          []LogData          `json:"LogData"`
	ApplicationDataN []ApplicationDataN `json:"IPApplicationsRef"`
	CircleData       []CircleData       `json:"CirclePrefRef"`
}

type EdgesPSApplicationData struct {
	LogData         []LogData          `json:"LogData"`
	ApplicationData []ApplicationDataV `json:"IPApplicationsRef"`
}
type EdgesPSApplicationDataNo struct {
	LogData          []LogData          `json:"LogData"`
	ApplicationDataN []ApplicationDataN `json:"PSApplicationsRef"`
}

type EdgesMTSPMApplicationData struct {
	LogData         []LogData          `json:"LogData"`
	ApplicationData []ApplicationDataV `json:"MTSPMMGApplicationsRef"`
}
type EdgesMTSPMApplicationDataNA struct {
	LogData          []LogData          `json:"LogData"`
	ApplicationDataN []ApplicationDataN `json:"MTSPMMGApplicationsRef"`
}

type EdgesGDSPAApplicationData struct {
	LogData         []LogData          `json:"LogData"`
	ApplicationData []ApplicationDataV `json:"GDSPAApplicationsRef"`
}

type EdgesGDSPAApplicationDataNA struct {
	LogData          []LogData          `json:"LogData"`
	ApplicationDataN []ApplicationDataN `json:"GDSPAApplicationsRef"`
}
type EdgesGDSPMApplicationData struct {
	LogData          []LogData           `json:"LogData"`
	ApplicationDataV []ApplicationDataVV `json:"GDSPMApplicationsRef"`
}
type EdgesGDSPMApplicationDataNAv struct {
	LogData            []LogData           `json:"LogData"`
	ApplicationDataNAv []ApplicationDataNV `json:"GDSPMApplicationsRef"`
}

type EdgesPMPAApplicationData struct {
	LogData         []LogData          `json:"LogData"`
	ApplicationData []ApplicationDataV `json:"PMPAApplicationsRef"`
}
type EdgesPMPAApplicationDataNA struct {
	LogData          []LogData          `json:"LogData"`
	ApplicationDataN []ApplicationDataN `json:"PMPAApplicationsRef"`
}
type EdgesApplicatinData struct {
	LogData         []LogData          `json:"LogData"`
	ApplicationData []ApplicationDataV `json:"IPApplicationsRef"`
	CircleData      []CircleData       `json:"CirclePrefRef"`
}

type CircleprefEdges struct {
	CircleData []CircleData `json:"CircleData"`
}
type IPApplicationsRefEdges struct {
	ApplicationData []ApplicationDataV `json:"ApplicationData"`
}
type PSApplicationsRefEdges struct {
	PsApplicationData []PsApplicationData `json:"PsApplicationData"`
}
type GDSPAApplicationsRefEdge struct {
	GDSPAApplicationData []GDSPAApplicationData `json:"GDSPAApplicationData"`
}
type PMPAApplicationsRefEdge struct {
	PMPAApplicationData []PMPAApplicationData `json:"PMPAApplicationData"`
}
type Gender string

// Constants for Gender values
const (
	Male   Gender = "Male"
	Female Gender = "Female"
	Other  Gender = "Other"
)

// LogData represents individual log entry data
type LogData struct {
	Userid    string  `json:"UserID" binding:"required"`
	Usertype  string  `json:"UserType"`
	Remarks   string  `json:"Remarks" binding:"required"`
	Action    string  `json:"Action"`
	Ipaddress string  `json:"IPAddress" binding:"required"`
	Os        string  `json:"OS"`
	Browser   string  `json:"Browser"`
	Latitude  float64 `json:"Latitude" binding:"required"`
	Longitude float64 `json:"Longitude" binding:"required"`
}
type CircleData struct {
	PlacePrefNo    int64  `json:"PlacePrefNo"`
	PlacePrefValue string `json:"PlacePrefValue"`
}

type ApplicationDataV struct {
	VacancyYear        int32  `json:"VacancyYear" `
	CA_Recommendations string `json:"CA_Recommendations"`
	CA_Remarks         string `json:"CA_Remarks"`
}
type ApplicationDataVV struct {
	VacancyYear        int32  `json:"VacancyYear" `
	CA_Recommendations string `json:"CA_Recommendations"`
	CA_Remarks         string `json:"CA_Remarks"`
	Post               string `json:"Post"`
}

type ApplicationDataN struct {
	VacancyYear        int32  `json:"VacancyYear" `
	NO_Recommendations string `json:"NO_Recommendations"`
	NO_Remarks         string `json:"NO_Remarks"`
}
type ApplicationDataNV struct {
	VacancyYear        int32  `json:"VacancyYear" `
	NO_Recommendations string `json:"NO_Recommendations"`
	NO_Remarks         string `json:"NO_Remarks"`
	Post               string `json:"Post"`
}
type PsApplicationData struct {
	VacancyYear        int32  `json:"VacancyYear" binding:"required"`
	CA_Recommendations string `json:"CA_Recommendations"`
	CA_Remarks         string `json:"CA_Remarks"`
}

type GDSPAApplicationData struct {
	VacancyYear        int32  `json:"VacancyYear" binding:"required"`
	CA_Recommendations string `json:"CA_Recommendations"`
	CA_Remarks         string `json:"CA_Remarks"`
}
type PMPAApplicationData struct {
	VacancyYear        int32  `json:"VacancyYear" binding:"required"`
	CA_Recommendations string `json:"CA_Recommendations"`
	CA_Remarks         string `json:"CA_Remarks"`
}

type Candidate struct {
	UserName string       `json:"UserName" binding:"required"`
	OTP      int          `json:"OTP" binding:"required"`
	Edges    EdgesLogdata `json:"edges" binding:"required"`
}

type CandidateRegistrationGenerateOTP struct {
	UserName string       `json:"UserName" binding:"required"`
	Password string       `json:"Password" binding:"required"`
	Edges    EdgesLogdata `json:"edges" binding:"required"`
}
type UserMasterResponse struct {
	ID                      int64     `json:"ID"`
	EmployeeID              int64     `json:"EmployeeID,omitempty"`
	EmployeeName            string    `json:"EmployeeName,omitempty"`
	Mobile                  string    `json:"Mobile,omitempty"`
	EmailID                 string    `json:"EmailID,omitempty"`
	UserName                string    `json:"UserName,omitempty"`
	Password                string    `json:"Password,omitempty"`
	Status                  bool      `json:"Status,omitempty"`
	Statuss                 string    `json:"Statuss,omitempty"`
	OTP                     int32     `json:"OTP,omitempty"`
	OTPNew                  int32     `json:"OTPNew,omitempty"`
	EmailOTPNew             int32     `json:"EmailOTPNew,omitempty"`
	OTPRemarks              string    `json:"OTPRemarks,omitempty"`
	CreatedAt               time.Time `json:"CreatedAt,omitempty"`
	OTPTriggeredTime        time.Time `json:"OTPTriggeredTime,omitempty"`
	OTPSavedTime            time.Time `json:"OTPSavedTime,omitempty"`
	OTPExpiryTime           time.Time `json:"OTPExpiryTime,omitempty"`
	NewPasswordRequest      bool      `json:"NewPasswordRequest,omitempty"`
	EmailOTP                int32     `json:"EmailOTP,omitempty"`
	EmailOTPRemarks         string    `json:"EmailOTPRemarks,omitempty"`
	EmailCreatedAt          time.Time `json:"EmailCreatedAt,omitempty"`
	EmailOTPTriggeredTime   time.Time `json:"EmailOTPTriggeredTime,omitempty"`
	EmailOTPSavedTime       time.Time `json:"EmailOTPSavedTime,omitempty"`
	EmailOTPExpiryTime      time.Time `json:"EmailOTPExpiryTime,omitempty"`
	EmailNewPasswordRequest bool      `json:"EmailNewPasswordRequest,omitempty"`
	UidToken                string    `json:"UidToken,omitempty"`
	CreatedById             int64     `json:"CreatedById,omitempty"`
	CreatedByEmployeeId     string    `json:"CreatedByEmployeeId,omitempty"`
	CreatedByUserName       string    `json:"CreatedByUserName,omitempty"`
	CreatedByDesignation    string    `json:"CreatedByDesignation,omitempty"`
	CreatedDate             time.Time `json:"CreatedDate,omitempty"`
	DeletedById             int64     `json:"DeletedById,omitempty"`
	DeletedByEmployeeId     string    `json:"DeletedByEmployeeId,omitempty"`
	DeletedByUserName       string    `json:"DeletedByUserName,omitempty"`
	DeletedByDesignation    string    `json:"DeletedByDesignation,omitempty"`
	DeletedDate             time.Time `json:"DeletedDate,omitempty"`
	FacilityID              string    `json:"FacilityID,omitempty"`
	CircleFacilityId        string    `json:"CircleFacilityId,omitempty"`
	CircleFacilityName      string    `json:"CircleFacilityName,omitempty"`
	Designation             string    `json:"Designation,omitempty"`
	RoleUserCode            int32     `json:"RoleUserCode,omitempty"`
	UpdatedBy               string    `json:"UpdatedBy,omitempty"`
	UpdatedDate             time.Time `json:"UpdatedDate,omitempty"`
	ModifiedBy              string    `json:"ModifiedBy,omitempty"`
	ModifiedDate            time.Time `json:"ModifiedDate,omitempty"`
	OperationStatus         string    `json:"OperationStatus,omitempty"`
	ExamCode                int32     `json:"ExamCode,omitempty"`
	ExamCodePS              int32     `json:"ExamCodePS,omitempty"`
	Gender                  string    `json:"Gender,omitempty"`
	DOB                     string    `json:"DOB,omitempty"`
	CreatedBy               string    `json:"CreatedBy,omitempty"`
}

type CandidateChangeEmailMobileGenerateOTP struct {
	NewEmailID  string       `json:"NewEmailID"`
	UserName    string       `json:"UserName" binding:"required"`
	NewMobile   string       `json:"NewMobile" `
	OldEmailID  string       `json:"OldEmailID" binding:"required"`
	OldMobile   string       `json:"OldMobile" binding:"required"`
	MobileEmail string       `json:"MobileEmail" binding:"required"`
	Edges       EdgesLogdata `json:"edges" binding:"required"`
}

type StrucUserResetValidateUserName struct {
	UserName string       `json:"UserName" binding:"required"`
	Edges    EdgesLogdata `json:"edges" binding:"required"`
}

type StrucUserResetValidateOTP struct {
	UserName string       `json:"UserName" binding:"required"`
	OTP      int32        `json:"OTP" binding:"required"`
	Edges    EdgesLogdata `json:"edges" binding:"required"`
}
type StrucUserResetSaveNewPassword struct {
	UserName    string       `json:"UserName" binding:"required"`
	NewPassword string       `json:"NewPassword" binding:"required"`
	OTP         int32        `json:"OTP" binding:"required"`
	Edges       EdgesLogdata `json:"edges" binding:"required"`
}

type StrucAdminResetValidateUserName struct {
	UserName string       `json:"UserName" binding:"required"`
	Edges    EdgesLogdata `json:"edges" binding:"required"`
}

type StrucAdminResetValidateOTP struct {
	UserName string       `json:"UserName" binding:"required"`
	OTP      int32        `json:"OTP" binding:"required"`
	Edges    EdgesLogdata `json:"edges" binding:"required"`
}
type StrucAdminResetSaveNewPassword struct {
	UserName    string       `json:"UserName" binding:"required"`
	NewPassword string       `json:"NewPassword" binding:"required"`
	OTP         int32        `json:"OTP" binding:"required"`
	Edges       EdgesLogdata `json:"edges" binding:"required"`
}

type CandidateRegistrationOTP struct {
	UserName string       `json:"UserName" binding:"required"`
	Password string       `json:"Password" binding:"required"`
	OTP      int32        `json:"OTP" binding:"required"`
	EmailOTP int32        `json:"EmailOTP" binding:"required"`
	Edges    EdgesLogdata `json:"edges" binding:"required"`
}
type CandidateEditProfileOTP struct {
	UserName    string       `json:"UserName" binding:"required"`
	MobileEmail string       `json:"MobileEmail" binding:"required"`
	OldOTP      int32        `json:"OldOTP" binding:"required"`
	NewOTP      int32        `json:"NewOTP"`
	EmailOTPNew int32        `json:"EmailOTPNew"`
	NewEmailID  string       `json:"NewEmailID"`
	NewMobile   string       `json:"NewMobile"`
	Edges       EdgesLogdata `json:"edges" binding:"required"`
}
type AdminLoginGenerateOTP struct {
	UserName string       `json:"UserName" binding:"required"`
	Password string       `json:"Password" binding:"required"`
	Edges    EdgesLogdata `json:"edges" binding:"required"`
}

type AdminLoginValidateOTP struct {
	UserName string       `json:"UserName" binding:"required"`
	OTP      int32        `json:"OTP" binding:"required"`
	Edges    EdgesLogdata `json:"edges" binding:"required"`
}

type AdminCreation struct {
	AuthorityFacilityName       string       `json:"AuthorityFacilityName" binding:"required"`
	CreatedByDesignation        string       `json:"CreatedByDesignation"  binding:"required"`
	CreatedByEmpId              int32        `json:"CreatedByEmpId" binding:"required"`
	CreatedById                 int32        `json:"CreatedById" binding:"required"`
	CreatedByUserName           string       `json:"CreatedByUserName" binding:"required"`
	Designation                 string       `json:"Designation" binding:"required"`
	EmailID                     string       `json:"EmailID" binding:"required"`
	EmployeeId                  int64        `json:"EmployeeId" binding:"required"`
	EmployeeName                string       `json:"EmployeeName" binding:"required"`
	FacilityID                  string       `json:"FacilityID" binding:"required"`
	FacilityIDUniqueid          int64        `json:"FacilityIdUniqueid" binding:"required"`
	FacilityType                string       `json:"FacilityType" binding:"required"`
	Mobile                      string       `json:"Mobile" binding:"required"`
	ReportingOfficeFacilityId   string       `json:"ReportingOfficeFacilityID" binding:"required"`
	ReportingOfficeFacilityName string       `json:"ReportingOfficeFacilityName" binding:"required"`
	RoleUserCode                int32        `json:"RoleUserCode" binding:"required"`
	RoleUserDescription         string       `json:"RoleUserDescription" binding:"required"`
	UserName                    string       `json:"UserName" binding:"required"`
	CircleOfficeFacilityId      string       `json:"CircleOfficeFacilityId" binding:"required"`
	CircleOfficeName            string       `json:"CircleOfficeName" binding:"required"`
	Edges                       EdgesLogdata `json:"edges" binding:"required"`
}

type UpdateAdminMasterStruc struct {
	EmployeeId             int64        `json:"EmployeeId" binding:"required"`
	EmployeeName           string       `json:"EmployeeName" binding:"required"`
	Designation            string       `json:"Designation" binding:"required"`
	RoleUserCode           int32        `json:"RoleUserCode" binding:"required"`
	RoleUserDescription    string       `json:"RoleUserDescription" binding:"required"`
	FacilityID             string       `json:"FacilityID" binding:"required"`
	Mobile                 string       `json:"Mobile" binding:"required"`
	EmailID                string       `json:"EmailID" binding:"required"`
	ModifiedById           int          `json:"ModifiedById" binding:"required"`
	ModifiedByUserName     string       `json:"ModifiedByUserName" binding:"required"`
	ModifiedByEmpId        int64        `json:"ModifiedByEmpId" binding:"required"`
	ModifiedByDesignantion string       `json:"ModifiedByDesignation" binding:"required"`
	AdminUniqueID          int32        `json:"AdminUniqueID" binding:"required"`
	UserName               string       `json:"UserName" binding:"required"`
	Edges                  EdgesLogdata `json:"edges" binding:"required"`
}

type DeleteAdminMasterStruc struct {
	AdminUniqueID        int32        `json:"AdminUniqueID" binding:"required"`
	UserName             string       `json:"UserName" binding:"required"`
	DeletedById          int64        `json:"DeletedById" binding:"required"`
	DeletedByUserName    string       `json:"DeletedByUserName" binding:"required"`
	DeletedByEmpId       int64        `json:"DeletedByEmpId" binding:"required"`
	DeletedByDesignation string       `json:"DeletedByDesignation" binding:"required"`
	Edges                EdgesLogdata `json:"edges" binding:"required"`
}
type TriggerCandidateSMSOTP struct {
	EmployeeId   string `json:"EmployeeId" binding:"required"`
	MobileNumber string `json:"MobileNumber" binding:"required"`
}
type TriggerCandidateEmailOTP struct {
	EmployeeId string `json:"EmployeeId" binding:"required"`
	EmailId    string `json:"EmailId" binding:"required"`
}
type VerifyTriggerCandidateEmailOTP struct {
	EmployeeId string `json:"EmployeeId" binding:"required"`
	EmailId    string `json:"EmailId" binding:"required"`
	EmailOTP   int64  `json:"EmailOTP" binding:"required"`
}
type VerifyTriggerCandidateSMSOTP struct {
	EmployeeId   string `json:"EmployeeId" binding:"required"`
	MobileNumber string `json:"MobileNumber" binding:"required"`
	SmsOTP       int64  `json:"SmsOTP" binding:"required"`
}
type StrucCreateEmployeeMaster struct {
	EmployeeID                     int64        `json:"EmployeeID" binding:"required"`
	EmployeeName                   string       `json:"EmployeeName" binding:"required"`
	DOB                            string       `json:"DOB" binding:"required"`
	Gender                         string       `json:"Gender" binding:"required"`
	MobileNumber                   string       `json:"MobileNumber" binding:"required"`
	EmailId                        string       `json:"EmailId" binding:"required"`
	EmployeeCategory               string       `json:"EmployeeCategory" binding:"required"`
	EmployeePost                   string       `json:"EmployeePost" binding:"required"`
	FacilityId                     string       `json:"FacilityId" binding:"required"`
	Pincode                        string       `json:"Pincode" binding:"required"`
	OfficeName                     string       `json:"OfficeName" binding:"required"`
	ControllingAuthorityFacilityId string       `json:"ControllingAuthorityFacilityId" binding:"required"`
	ControllingAuthorityName       string       `json:"ControllingAuthorityName" binding:"required"`
	NodalAuthorityFaciliyId        string       `json:"NodalAuthorityFaciliyId" binding:"required"`
	NodalAuthorityName             string       `json:"NodalAuthorityName" binding:"required"`
	CircleFacilityId               string       `json:"CircleFacilityId" binding:"required"`
	CreatedById                    int64        `json:"CreatedById" binding:"required"`
	CreatedByUserName              string       `json:"CreatedByUserName" binding:"required"`
	CreatedByEmpId                 int64        `json:"CreatedByEmpId" binding:"required"`
	CreatedByDesignation           string       `json:"CreatedByDesignation" binding:"required"`
	VerifyStatus                   bool         `json:"VerifyStatus"`
	Statuss                        string       `json:"Statuss"`
	Edges                          EdgesLogdata `json:"edges" binding:"required"`
}
type EmployeeMasterResponse struct {
	ID                             int64     `json:"ID,omitempty"`
	EmployeeID                     int64     `json:"EmployeeID,omitempty"`
	EmployeeName                   string    `json:"EmployeeName,omitempty"`
	DOB                            string    `json:"DOB,omitempty"`
	Gender                         string    `json:"Gender,omitempty"`
	MobileNumber                   string    `json:"MobileNumber,omitempty"`
	EmailID                        string    `json:"EmailID,omitempty"`
	EmployeeCategoryCode           string    `json:"EmployeeCategoryCode,omitempty"`
	EmployeeCategory               string    `json:"EmployeeCategory,omitempty"`
	PostCode                       string    `json:"PostCode,omitempty"`
	EmployeePost                   string    `json:"EmployeePost,omitempty"`
	FacilityID                     string    `json:"FacilityID,omitempty"`
	OfficeName                     string    `json:"OfficeName,omitempty"`
	ControllingAuthorityFacilityId string    `json:"ControllingAuthorityFacilityId,omitempty"`
	ControllingAuthorityName       string    `json:"ControllingAuthorityName,omitempty"`
	NodalAuthorityFacilityId       string    `json:"NodalAuthorityFacilityId,omitempty"`
	NodalAuthorityName             string    `json:"NodalAuthorityName,omitempty"`
	Pincode                        string    `json:"Pincode,omitempty"`
	CircleFacilityID               string    `json:"CircleFacilityID,omitempty"`
	Status                         string    `json:"Status,omitempty"`
	VerifyStatus                   bool      `json:"VerifyStatus,omitempty"`
	UidToken                       string    `json:"UidToken,omitempty"`
	CreatedBy                      string    `json:"CreatedBy,omitempty"`
	DCCS                           string    `json:"DCCS,omitempty"`
	CreatedByID                    int64     `json:"CreatedByID,omitempty"`
	CreatedByUserName              string    `json:"CreatedByUserName,omitempty"`
	CreatedByEmpID                 int64     `json:"CreatedByEmpID,omitempty"`
	CreatedByDesignation           string    `json:"CreatedByDesignation,omitempty"`
	CreatedDate                    time.Time `json:"CreatedDate,omitempty"`
	ModifiedByID                   int64     `json:"ModifiedByID,omitempty"`
	ModifiedByUserName             string    `json:"ModifiedByUserName,omitempty"`
	ModifiedByEmpID                int64     `json:"ModifiedByEmpID,omitempty"`
	ModifiedByDesignation          string    `json:"ModifiedByDesignation,omitempty"`
	ModifiedDate                   time.Time `json:"ModifiedDate,omitempty"`
	DeletedByID                    int64     `json:"DeletedByID,omitempty"`
	DeletedByUserName              string    `json:"DeletedByUserName,omitempty"`
	DeletedByEmpID                 int64     `json:"DeletedByEmpID,omitempty"`
	DeletedByDesignation           string    `json:"DeletedByDesignation,omitempty"`
	DeletedDate                    time.Time `json:"DeletedDate,omitempty"`
	UpdatedAt                      time.Time `json:"UpdatedAt,omitempty"`
	UpdatedBy                      string    `json:"UpdatedBy,omitempty"`
	SmsOtp                         int64     `json:"SmsOtp,omitempty"`
	SmsTriggeredTime               time.Time `json:"SmsTriggeredTime,omitempty"`
	SmsVerifyStatus                bool      `json:"SmsVerifyStatus,omitempty"`
	EmailOtp                       int64     `json:"EmailOtp,omitempty"`
	EmailTriggeredTime             time.Time `json:"EmailTriggeredTime,omitempty"`
	EmailVerifyStatus              bool      `json:"EmailVerifyStatus,omitempty"`
	FinalSubmitStatus              bool      `json:"FinalSubmitStatus,omitempty"`
	DCInPresentCadre               string    `json:"DCInPresentCadre,omitempty"`
	Cadre                          string    `json:"Cadre,omitempty"`
}

type StrucModifyEmployeeMaster struct {
	EmployeeID                     int64        `json:"EmployeeID" binding:"required"`
	EmployeeName                   string       `json:"EmployeeName" binding:"required"`
	DOB                            string       `json:"DOB" binding:"required"`
	Gender                         string       `json:"Gender" binding:"required"`
	MobileNumber                   string       `json:"MobileNumber" binding:"required"`
	EmailId                        string       `json:"EmailId" binding:"required"`
	EmployeeCategory               string       `json:"EmployeeCategory" binding:"required"`
	EmployeePost                   string       `json:"EmployeePost" binding:"required"`
	FacilityId                     string       `json:"FacilityId" binding:"required"`
	Pincode                        string       `json:"Pincode" binding:"required"`
	OfficeName                     string       `json:"OfficeName" binding:"required"`
	ControllingAuthorityFacilityId string       `json:"ControllingAuthorityFacilityId" binding:"required"`
	ControllingAuthorityName       string       `json:"ControllingAuthorityName" binding:"required"`
	NodalAuthorityFaciliyId        string       `json:"NodalAuthorityFaciliyId" binding:"required"`
	NodalAuthorityName             string       `json:"NodalAuthorityName" binding:"required"`
	CircleFacilityId               string       `json:"CircleFacilityId" binding:"required"`
	ModifiedById                   int64        `json:"ModifiedById" binding:"required"`
	ModifiedByUserName             string       `json:"ModifiedByUserName" binding:"required"`
	ModifiedByEmpId                int64        `json:"ModifiedByEmpId" binding:"required"`
	ModifiedByDesignantion         string       `json:"ModifiedByDesignantion" binding:"required"`
	VerifyStatus                   bool         `json:"VerifyStatus"`
	Remarks                        string       `json:"Remarks"`
	Edges                          EdgesLogdata `json:"edges" binding:"required"`
}

type RequestBody struct {
	UserName        string       `json:"UserName" binding:"required"`
	CurrentPassword string       `json:"current_password" binding:"required"`
	NewPassword     string       `json:"new_password" binding:"required"`
	Edges           EdgesLogdata `json:"edges"`
}

type RequestBodyNew struct {
	CurrentPassword string       `json:"current_password" binding:"required"`
	NewPassword     string       `json:"new_password" binding:"required"`
	Edges           EdgesLogdata `json:"edges" binding:"required"`
}

// dev by vk
type CandidateCreation struct {
	CreatedBy    string     `json:"CreatedBy" binding:"required"`
	EmailID      string     `json:"EmailID" binding:"required"`
	EmployeeId   int64      `json:"EmployeeID" binding:"required"`
	EmployeeName string     `json:"EmployeeName" binding:"required"`
	FacilityID   string     `json:"FacilityID" binding:"required"`
	Mobile       string     `json:"MobileNumber" binding:"required"`
	Gender       Gender     `json:"Gender" binding:"required"`
	DOB          *time.Time `json:"DOB" binding:"required"`
	UidToken     string     `json:"UidToken"`

	Edges EdgesLogdata `json:"edges" binding:"required"`
}
type service struct {
	VacancyDate            *time.Time `json:"Vacancy Date" binding:"required"`
	Age                    string     `json:"Age" binding:"required"`
	AgeEligibility         string     `json:"AgeEligibility" binding:"required"`
	Service                string     `json:"Service" binding:"required"`
	ServiceEligibility     string     `json:"ServiceEligibility" binding:"required"`
	ServiceEligibilityYear int64      `json:"ServiceEligibilityYear" binding:"required"`
}

type StruUserGenerateOTP struct {
	UserName    string       `json:"UserName" binding:"required"`
	NewPassword string       `json:"Password" binding:"required"`
	Edges       EdgesLogdata `json:"edges" binding:"required"`
}

type StrucMappingIdentificationNumber struct {
	NodalOfficeFacilityId string `json:"nodal_office_facility_id"`
	StartNo               int32  `json:"start_no"`
	EndNo                 int32  `json:"end_no"`
}

type StruExamCenterHall struct {
	CenterCode            int32  `json:"CenterCode" binding:"required"`
	CityID                int32  `json:"CityID" binding:"required"`
	ExamCenterName        string `json:"ExamCenterName" binding:"required"`
	ExamYear              string `json:"ExamYear" binding:"required"`
	ExamCode              int32  `json:"ExamCode" binding:"required"`
	ExamName              string `json:"ExamName" binding:"required"`
	CenterCityName        string `json:"CenterCityName" binding:"required"`
	ConductedByFacilityID string `json:"ConductedByFacilityID" binding:"required"`
	ConductedBy           string `json:"ConductedBy" binding:"required"`
	HallName              string `json:"HallName" binding:"required"`
	AdminCircleOfficeID   string `json:"AdminCircleOfficeID" binding:"required"`
	//MappingIdentificationNumber *[]interface{}  `json:"MappingIdentificationNumber" binding:"required"`
	MappingIdentificationNumber []StrucMappingIdentificationNumber `json:"MappingIdentificationNumber" binding:"required"`
	CreatedById                 int64                              `json:"CreatedById" `
	CreatedByUserName           string                             `json:"CreatedByUserName" `
	CreatedByEmpId              int64                              `json:"CreatedByEmpId" `
	CreatedByDesignation        string                             `json:"CreatedByDesignation" `
	ModifiedById                int64                              `json:"ModifiedById" `
	ModifiedByUserName          string                             `json:"ModifiedByUserName" `
	ModifiedByEmpId             int64                              `json:"ModifiedByEmpId" `
	ModifiedByDesignantion      string                             `json:"ModifiedByDesignantion" `
	DeletedById                 int64                              `json:"DeletedById" `
	DeletedByUserName           string                             `json:"DeletedByUserName" `
	DeletedByEmpId              int64                              `json:"DeletedByEmpId" `
	DeletedByDesignation        string                             `json:"DeletedByDesignation" `
	NoSeats                     int32                              `json:"NoSeats" binding:"required"`
	Edges                       EdgesCircleData                    `json:"edges" binding:"required"`
}

type StruExamCenterHallReset struct {
	CenterCode            int32  `json:"CenterCode" binding:"required"`
	CityID                int32  `json:"CityID" binding:"required"`
	ExamCenterName        string `json:"ExamCenterName" binding:"required"`
	ExamYear              string `json:"ExamYear" binding:"required"`
	ExamCode              int32  `json:"ExamCode" binding:"required"`
	ExamName              string `json:"ExamName" binding:"required"`
	CenterCityName        string `json:"CenterCityName" binding:"required"`
	ConductedByFacilityID string `json:"ConductedByFacilityID" binding:"required"`
	ConductedBy           string `json:"ConductedBy" binding:"required"`
	HallName              string `json:"HallName" binding:"required"`
	AdminCircleOfficeID   string `json:"AdminCircleOfficeID" binding:"required"`

	Edges EdgesCircleData `json:"edges" binding:"required"`
}

type ApplicationIp struct {
	CandidateRemarks                string          `json:"CandidateRemarks"`
	CategoryCode                    string          `json:"CategoryCode" binding:"required"`
	CategoryDescription             string          `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string          `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32           `json:"CenterId" binding:"required"`
	CentrePreference                string          `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string          `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string          `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string          `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string          `json:"DCCS"`
	DOB                             string          `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string          `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string          `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string          `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string          `json:"DeputationOfficeName"`
	DeputationOfficePincode         string          `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string          `json:"DeputationOfficeUniqueId"`
	InDeputation                    string          `json:"InDeputation" `
	DeputationType                  string          `json:"DeputationType"`
	DesignationID                   string          `json:"DesignationID"`
	DisabilityPercentage            int32           `json:"DisabilityPercentage"`
	DisabilityTypeCode              string          `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string          `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string          `json:"DisabilityTypeID"`
	Edges                           EdgesCircleData `json:"edges" binding:"required"`
	EducationCode                   string          `json:"EducationCode"`
	EducationDescription            string          `json:"EducationDescription"`
	EmailID                         string          `json:"EmailID" binding:"required"`
	EmployeeID                      int64           `json:"EmployeeID" binding:"required"`
	EmployeeName                    string          `json:"EmployeeName" binding:"required"`
	EntryPostCode                   string          `json:"EntryPostCode"`
	EntryPostDescription            string          `json:"EntryPostDescription"`
	ExamCode                        int32           `json:"ExamCode" binding:"required"`
	ExamName                        string          `json:"ExamName" binding:"required"`
	ExamShortName                   string          `json:"ExamShortName" binding:"required"`
	ExamYear                        string          `json:"ExamYear" binding:"required"`
	FacilityUniqueID                string          `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string          `json:"FeederPostCode" binding:"required"`
	FeederPostDescription           string          `json:"FeederPostDescription" binding:"required"`
	FeederPostJoiningDate           string          `json:"FeederPostJoiningDate" binding:"required"`
	Gender                          string          `json:"Gender" binding:"required"`
	LienControllingOfficeID         string          `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string          `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string          `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string          `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string          `json:"NodalOfficeName" binding:"required"`
	Photo                           string          `json:"Photo" binding:"required"`
	PhotoPath                       string          `json:"PhotoPath" binding:"required"`
	PresentDesignation              string          `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string          `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string          `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string          `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string          `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{}  `json:"ServiceLength" binding:"required"`
	Signature                       string          `json:"Signature" binding:"required"`
	SignaturePath                   string          `json:"SignaturePath" binding:"required"`
	TempHallTicket                  string          `json:"TempHallTicket" binding:"required"`
	UserID                          int32           `json:"UserID"`
	WorkingOfficeCircleFacilityID   string          `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string          `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string          `json:"WorkingOfficeDivisionFacilityID"`
	WorkingOfficeDivisionName       string          `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string          `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string          `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32           `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string          `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string          `json:"WorkingOfficeRegionName" `
	Cadre                           string          `json:"Cadre"`
	ApplicationLastDate             time.Time       `json:"ApplicationLastDate"`
}
type UpdateExamCentersInIP struct {
	CandidateRemarks                string          `json:"CandidateRemarks"`
	CategoryCode                    string          `json:"CategoryCode"`
	CategoryDescription             string          `json:"CategoryDescription"`
	CenterFacilityId                string          `json:"CenterFacilityId" `
	CenterId                        int32           `json:"CenterId"`
	CentrePreference                string          `json:"CentrePreference"`
	ClaimingQualifyingService       string          `json:"ClaimingQualifyingService"`
	ControllingOfficeFacilityID     string          `json:"ControllingOfficeFacilityID" `
	ControllingOfficeName           string          `json:"ControllingOfficeName"`
	DCCS                            string          `json:"DCCS"`
	DOB                             string          `json:"DOB"`
	DeputationControllingOfficeID   string          `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string          `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string          `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string          `json:"DeputationOfficeName"`
	DeputationOfficePincode         string          `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string          `json:"DeputationOfficeUniqueId"`
	InDeputation                    string          `json:"InDeputation" `
	DeputationType                  string          `json:"DeputationType"`
	DesignationID                   string          `json:"DesignationID"`
	DisabilityPercentage            int32           `json:"DisabilityPercentage"`
	DisabilityTypeCode              string          `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string          `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string          `json:"DisabilityTypeID"`
	Edges                           EdgesCircleData `json:"edges"`
	EducationCode                   string          `json:"EducationCode"`
	EducationDescription            string          `json:"EducationDescription"`
	EmailID                         string          `json:"EmailID"`
	EmployeeID                      int64           `json:"EmployeeID"`
	EmployeeName                    string          `json:"EmployeeName"`
	EntryPostCode                   string          `json:"EntryPostCode"`
	EntryPostDescription            string          `json:"EntryPostDescription"`
	ExamCode                        int32           `json:"ExamCode" `
	ExamName                        string          `json:"ExamName" `
	ExamShortName                   string          `json:"ExamShortName" `
	ExamYear                        string          `json:"ExamYear"`
	FacilityUniqueID                string          `json:"FacilityUniqueID" `
	FeederPostCode                  string          `json:"FeederPostCode"`
	FeederPostDescription           string          `json:"FeederPostDescription" `
	FeederPostJoiningDate           string          `json:"FeederPostJoiningDate" `
	Gender                          string          `json:"Gender" `
	LienControllingOfficeID         string          `json:"LienControllingOfficeID" `
	LienControllingOfficeName       string          `json:"LienControllingOfficeName" `
	MobileNumber                    string          `json:"MobileNumber" `
	NodalOfficeFacilityID           string          `json:"NodalOfficeFacilityID" `
	NodalOfficeName                 string          `json:"NodalOfficeName" `
	Photo                           string          `json:"Photo" `
	PhotoPath                       string          `json:"PhotoPath" `
	PresentDesignation              string          `json:"PresentDesignation" `
	PresentPostCode                 string          `json:"PresentPostCode"`
	PresentPostDescription          string          `json:"PresentPostDescription" `
	ReportingOfficeFacilityID       string          `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string          `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{}  `json:"ServiceLength" `
	Signature                       string          `json:"Signature" `
	SignaturePath                   string          `json:"SignaturePath" `
	TempHallTicket                  string          `json:"TempHallTicket"`
	UserID                          int32           `json:"UserID"`
	WorkingOfficeCircleFacilityID   string          `json:"WorkingOfficeCircleFacilityID" `
	WorkingOfficeCircleName         string          `json:"WorkingOfficeCircleName" `
	WorkingOfficeDivisionFacilityID string          `json:"WorkingOfficeDivisionFacilityID"`
	WorkingOfficeDivisionName       string          `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string          `json:"WorkingOfficeFacilityID" `
	WorkingOfficeName               string          `json:"WorkingOfficeName" `
	WorkingOfficePincode            int32           `json:"WorkingOfficePincode"`
	WorkingOfficeRegionFacilityID   string          `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string          `json:"WorkingOfficeRegionName" `
	Cadre                           string          `json:"Cadre"`
}

type ApplicationGroupB struct {
	CandidateRemarks                string                 `json:"CandidateRemarks"`
	CategoryCode                    string                 `json:"CategoryCode" binding:"required"`
	CategoryDescription             string                 `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string                 `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32                  `json:"CenterId" binding:"required"`
	CentrePreference                string                 `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string                 `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string                 `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string                 `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string                 `json:"DCCS" binding:"required"`
	DOB                             string                 `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string                 `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string                 `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string                 `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string                 `json:"DeputationOfficeName"`
	DeputationOfficePincode         string                 `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string                 `json:"DeputationOfficeUniqueId"`
	InDeputation                    string                 `json:"InDeputation" `
	DeputationType                  string                 `json:"DeputationType"`
	DesignationID                   string                 `json:"DesignationID"`
	DisabilityPercentage            int32                  `json:"DisabilityPercentage" `
	DisabilityTypeCode              string                 `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string                 `json:"DisabilityTypeDescription" binding:"required"`
	DisabilityTypeID                string                 `json:"DisabilityTypeID"`
	Edges                           EdgesPSApplicationData `json:"edges" binding:"required"`
	EducationCode                   string                 `json:"EducationCode"`
	EducationDescription            string                 `json:"EducationDescription"`
	EmailID                         string                 `json:"EmailID" binding:"required"`
	EmployeeID                      int64                  `json:"EmployeeID" binding:"required"`
	EmployeeName                    string                 `json:"EmployeeName" binding:"required"`
	EmployeePost                    string                 `json:"EmployeePost" `
	EntryPostCode                   string                 `json:"EntryPostCode" `
	EntryPostDescription            string                 `json:"EntryPostDescription" binding:"required"`
	ExamCode                        int32                  `json:"ExamCode" binding:"required"`
	ExamName                        string                 `json:"ExamName" binding:"required"`
	ExamShortName                   string                 `json:"ExamShortName" binding:"required"`
	ExamYear                        string                 `json:"ExamYear" binding:"required"`
	FacilityName                    string                 `json:"FacilityName" `
	FacilityUniqueID                string                 `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string                 `json:"FeederPostCode" binding:"required"`
	FeederPostDescription           string                 `json:"FeederPostDescription" binding:"required"`
	FeederPostJoiningDate           string                 `json:"FeederPostJoiningDate" binding:"required"`
	Gender                          string                 `json:"Gender" binding:"required"`
	LienControllingOfficeID         string                 `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string                 `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string                 `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string                 `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string                 `json:"NodalOfficeName" binding:"required"`
	Photo                           string                 `json:"Photo" binding:"required"`
	PhotoPath                       string                 `json:"PhotoPath" binding:"required"`
	PresentDesignation              string                 `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string                 `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string                 `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string                 `json:"ReportingOfficeFacilityID" `
	ReportingOfficeName             string                 `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{}         `json:"ServiceLength" binding:"required"`
	Signature                       string                 `json:"Signature" binding:"required"`
	SignaturePath                   string                 `json:"SignaturePath" binding:"required"`
	TempHallTicket                  string                 `json:"TempHallTicket" binding:"required"`
	UserID                          int32                  `json:"UserID" binding:"required"`
	WorkingOfficeCircleFacilityID   string                 `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string                 `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string                 `json:"WorkingOfficeDivisionFacilityID"`
	WorkingOfficeDivisionName       string                 `json:"WorkingOfficeDivisionName"`
	WorkingOfficeFacilityID         string                 `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string                 `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32                  `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string                 `json:"WorkingOfficeRegionFacilityID"`
	WorkingOfficeRegionName         string                 `json:"WorkingOfficeRegionName"`
	ApplicationLastDate             time.Time              `json:"ApplicationLastDate" binding:"required"`
}

type ResubmitApplicationIp struct {
	ApplicationID                   int32           `json:"id"`
	ApplicationNumber               string          `json:"ApplicationNumber"`
	CandidateRemarks                string          `json:"CandidateRemarks"`
	CategoryCode                    string          `json:"CategoryCode" binding:"required"`
	CategoryDescription             string          `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string          `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32           `json:"CenterId" binding:"required"`
	CentrePreference                string          `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string          `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string          `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string          `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string          `json:"DCCS"`
	DOB                             string          `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string          `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string          `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string          `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string          `json:"DeputationOfficeName"`
	DeputationOfficePincode         string          `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string          `json:"DeputationOfficeUniqueId"`
	InDeputation                    string          `json:"InDeputation" `
	DeputationType                  string          `json:"DeputationType"`
	DesignationID                   string          `json:"DesignationID"`
	DisabilityPercentage            int32           `json:"DisabilityPercentage"`
	DisabilityTypeCode              string          `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string          `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string          `json:"DisabilityTypeID"`
	Edges                           EdgesCircleData `json:"edges" binding:"required"`
	EducationCode                   string          `json:"EducationCode"`
	EducationDescription            string          `json:"EducationDescription"`
	EmailID                         string          `json:"EmailID" binding:"required"`
	EmployeeID                      int64           `json:"EmployeeID" binding:"required"`
	EmployeeName                    string          `json:"EmployeeName" binding:"required"`
	EntryPostCode                   string          `json:"EntryPostCode"`
	EntryPostDescription            string          `json:"EntryPostDescription"`
	ExamCode                        int32           `json:"ExamCode" binding:"required"`
	ExamName                        string          `json:"ExamName" binding:"required"`
	ExamShortName                   string          `json:"ExamShortName" binding:"required"`
	ExamYear                        string          `json:"ExamYear" binding:"required"`
	FacilityUniqueID                string          `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string          `json:"FeederPostCode" binding:"required"`
	FeederPostDescription           string          `json:"FeederPostDescription" binding:"required"`
	FeederPostJoiningDate           string          `json:"FeederPostJoiningDate" binding:"required"`
	Gender                          string          `json:"Gender" binding:"required"`
	LienControllingOfficeID         string          `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string          `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string          `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string          `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string          `json:"NodalOfficeName" binding:"required"`
	Photo                           string          `json:"Photo" binding:"required"`
	PhotoPath                       string          `json:"PhotoPath" binding:"required"`
	PresentDesignation              string          `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string          `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string          `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string          `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string          `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{}  `json:"ServiceLength" binding:"required"`
	Signature                       string          `json:"Signature" binding:"required"`
	SignaturePath                   string          `json:"SignaturePath" binding:"required"`
	TempHallTicket                  string          `json:"TempHallTicket" binding:"required"`
	UserID                          int32           `json:"UserID"`
	WorkingOfficeCircleFacilityID   string          `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string          `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string          `json:"WorkingOfficeDivisionFacilityID"`
	WorkingOfficeDivisionName       string          `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string          `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string          `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32           `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string          `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string          `json:"WorkingOfficeRegionName" `
	Cadre                           string          `json:"Cadre"`
	ApplicationCorrectionLastDate   time.Time       `json:"ApplicationCorrectionLastDate"`
}

type ReApplicationGroupB struct {
	ApplicationID                   int64                  `json:"id"`
	ApplicationNumber               string                 `json:"ApplicationNumber"`
	CandidateRemarks                string                 `json:"CandidateRemarks"`
	CategoryCode                    string                 `json:"CategoryCode" binding:"required"`
	CategoryDescription             string                 `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string                 `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32                  `json:"CenterId" binding:"required"`
	CentrePreference                string                 `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string                 `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string                 `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string                 `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string                 `json:"DCCS" binding:"required"`
	DOB                             string                 `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string                 `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string                 `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string                 `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string                 `json:"DeputationOfficeName"`
	DeputationOfficePincode         string                 `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string                 `json:"DeputationOfficeUniqueId"`
	InDeputation                    string                 `json:"InDeputation" `
	DeputationType                  string                 `json:"DeputationType"`
	DesignationID                   string                 `json:"DesignationID"`
	DisabilityPercentage            int32                  `json:"DisabilityPercentage" `
	DisabilityTypeCode              string                 `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string                 `json:"DisabilityTypeDescription" binding:"required"`
	DisabilityTypeID                string                 `json:"DisabilityTypeID"`
	Edges                           EdgesPSApplicationData `json:"edges" binding:"required"`
	EducationCode                   string                 `json:"EducationCode"`
	EducationDescription            string                 `json:"EducationDescription"`
	EmailID                         string                 `json:"EmailID" binding:"required"`
	EmployeeID                      int64                  `json:"EmployeeID" binding:"required"`
	EmployeeName                    string                 `json:"EmployeeName" binding:"required"`
	EmployeePost                    string                 `json:"EmployeePost" `
	EntryPostCode                   string                 `json:"EntryPostCode" `
	EntryPostDescription            string                 `json:"EntryPostDescription" binding:"required"`
	ExamCode                        int32                  `json:"ExamCode" binding:"required"`
	ExamName                        string                 `json:"ExamName" binding:"required"`
	ExamShortName                   string                 `json:"ExamShortName" binding:"required"`
	ExamYear                        string                 `json:"ExamYear" binding:"required"`
	FacilityName                    string                 `json:"FacilityName" `
	FacilityUniqueID                string                 `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string                 `json:"FeederPostCode" binding:"required"`
	FeederPostDescription           string                 `json:"FeederPostDescription" binding:"required"`
	FeederPostJoiningDate           string                 `json:"FeederPostJoiningDate" binding:"required"`
	Gender                          string                 `json:"Gender" binding:"required"`
	LienControllingOfficeID         string                 `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string                 `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string                 `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string                 `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string                 `json:"NodalOfficeName" binding:"required"`
	Photo                           string                 `json:"Photo" binding:"required"`
	PhotoPath                       string                 `json:"PhotoPath" binding:"required"`
	PresentDesignation              string                 `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string                 `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string                 `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string                 `json:"ReportingOfficeFacilityID" `
	ReportingOfficeName             string                 `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{}         `json:"ServiceLength" binding:"required"`
	Signature                       string                 `json:"Signature" binding:"required"`
	SignaturePath                   string                 `json:"SignaturePath" binding:"required"`
	TempHallTicket                  string                 `json:"TempHallTicket" binding:"required"`
	UserID                          int32                  `json:"UserID" binding:"required"`
	WorkingOfficeCircleFacilityID   string                 `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string                 `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string                 `json:"WorkingOfficeDivisionFacilityID"`
	WorkingOfficeDivisionName       string                 `json:"WorkingOfficeDivisionName"`
	WorkingOfficeFacilityID         string                 `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string                 `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32                  `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string                 `json:"WorkingOfficeRegionFacilityID"`
	WorkingOfficeRegionName         string                 `json:"WorkingOfficeRegionName"`
	ApplicationCorrectionLastDate   time.Time              `json:"ApplicationCorrectionLastDate"`
}

type VerifyApplicationGroupB struct {
	AppliactionRemarks          string                 `json:"AppliactionRemarks" `
	ApplicationNumber           string                 `json:"ApplicationNumber" binding:"required"`
	CA_EmployeeDesignation      string                 `json:"CA_EmployeeDesignation"  `
	CA_EmployeeID               string                 `json:"CA_EmployeeId" binding:"required"`
	CA_GeneralRemarks           string                 `json:"CA_GeneralRemarks" `
	CA_Remarks                  string                 `json:"CA_Remarks" `
	CA_UserName                 string                 `json:"CA_UserName" binding:"required"`
	ControllingOfficeFacilityID string                 `json:"ControllingOfficeFacilityID" binding:"required"`
	Edges                       EdgesPSApplicationData `json:"edges"`
	EmployeeID                  int64                  `json:"EmployeeID" binding:"required"`
	ExamCode                    int32                  `json:"ExamCode"`
	ExamName                    string                 `json:"ExamName" `
	ExamShortName               string                 `json:"ExamShortName"`
	ExamYear                    string                 `json:"ExamYear"`
	GenerateHallTicketFlag      bool                   `json:"GenerateHallTicketFlag" `
	ID                          int64                  `json:"id" binding:"required"`
	NonQualifyingService        *[]interface{}         `json:"NonQualifyingService" `
	ServiceLength               *[]interface{}         `json:"ServiceLength" binding:"required"`
	UserID                      int32                  `json:"UserID" `
	RecommendedStatus           string                 `json:"RecommendedStatus"`
	PunishmentStatus            bool                   `json:"PunishmentStatus" `
	DisciplinaryCaseStatus      bool                   `json:"DisciplinaryCaseStatus" `
}
type NAVerifyApplicationGroupB struct {
	AppliactionRemarks          string                   `json:"AppliactionRemarks" `
	ApplicationNumber           string                   `json:"ApplicationNumber" binding:"required"`
	NA_EmployeeDesignation      string                   `json:"NA_EmployeeDesignation"  `
	NA_EmployeeID               string                   `json:"NA_EmployeeId" binding:"required"`
	NA_GeneralRemarks           string                   `json:"NA_GeneralRemarks" `
	NA_Remarks                  string                   `json:"NA_Remarks" `
	NA_UserName                 string                   `json:"NA_UserName" binding:"required"`
	ControllingOfficeFacilityID string                   `json:"ControllingOfficeFacilityID" binding:"required"`
	Edges                       EdgesPSApplicationDataNo `json:"edges"`
	EmployeeID                  int64                    `json:"EmployeeID" binding:"required"`
	ExamCode                    int32                    `json:"ExamCode"`
	ExamName                    string                   `json:"ExamName"`
	ExamShortName               string                   `json:"ExamShortName"`
	ExamYear                    string                   `json:"ExamYear"`
	GenerateHallTicketFlag      bool                     `json:"GenerateHallTicketFlag"`
	ID                          int64                    `json:"id" binding:"required"`
	UserID                      int32                    `json:"UserID" `
	RecommendedStatus           string                   `json:"RecommendedStatus"`
}

/*
	 type UnitPreference struct {
		PostalAssistant []UnitPref `json:"Postal Assistant"`
	}

	type UnitPref struct {
		UnitPrefNo    int64  `json:"UnitPrefNo" binding:"required"`
		UnitPrefValue string `json:"UnitPrefValue" binding:"required"`
	}

	type PostPreference struct {
		PostalAssistant []PostPref `json:"Postal Assistant"`
	}

	type PostPref struct {
		Cadre         string `json:"Cadre" binding:"required"`
		PostPrefNo    int64  `json:"PostPrefNo" binding:"required"`
		PostPrefValue string `json:"PostPrefValue" binding:"required"`
	}
*/
type ApplicationGDStoPA struct {
	Cadre                           string         `json:"Cadre"`
	CandidateRemarks                string         `json:"CandidateRemarks"`
	CategoryCode                    string         `json:"CategoryCode" binding:"required"`
	CategoryDescription             string         `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string         `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32          `json:"CenterId" binding:"required"`
	CentrePreference                string         `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string         `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string         `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string         `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string         `json:"DCCS"`
	DOB                             string         `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string         `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string         `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string         `json:"DeputationOfficeFacilityID"`
	DeputationOfficeName            string         `json:"DeputationOfficeName"`
	DeputationOfficePincode         string         `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string         `json:"DeputationOfficeUniqueId"`
	InDeputation                    string         `json:"InDeputation" `
	DeputationType                  string         `json:"DeputationType"`
	DesignationID                   string         `json:"DesignationID"`
	DisabilityPercentage            int32          `json:"DisabilityPercentage"`
	DisabilityTypeCode              string         `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string         `json:"DisabilityTypeDescription" binding:"required"`
	DisabilityTypeID                string         `json:"DisabilityTypeID"`
	Edges                           EdgesLogdata   `json:"edges" binding:"required"`
	EducationCode                   string         `json:"EducationCode"`
	EducationDescription            string         `json:"EducationDescription" binding:"required"`
	EmailID                         string         `json:"EmailID" binding:"required"`
	EmployeeID                      int64          `json:"EmployeeID" binding:"required"`
	EmployeeName                    string         `json:"EmployeeName" binding:"required"`
	EmployeePost                    string         `json:"EmployeePost" `
	EntryPostCode                   string         `json:"EntryPostCode"`
	EntryPostDescription            string         `json:"EntryPostDescription"`
	ExamCode                        int32          `json:"ExamCode" binding:"required"`
	ExamName                        string         `json:"ExamName" binding:"required"`
	ExamShortName                   string         `json:"ExamShortName" binding:"required"`
	ExamYear                        string         `json:"ExamYear" binding:"required"`
	FacilityName                    string         `json:"FacilityName"`
	FacilityUniqueID                string         `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string         `json:"FeederPostCode" binding:"required"`
	FeederPostDescription           string         `json:"FeederPostDescription" binding:"required"`
	FeederPostJoiningDate           string         `json:"FeederPostJoiningDate" binding:"required"`
	Gender                          string         `json:"Gender" binding:"required"`
	LienControllingOfficeID         string         `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string         `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string         `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string         `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string         `json:"NodalOfficeName" binding:"required"`
	Photo                           string         `json:"Photo" binding:"required"`
	PhotoPath                       string         `json:"PhotoPath" binding:"required"`
	PostPreferences                 *[]interface{} `json:"PostPreferences" binding:"required"`
	PresentDesignation              string         `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string         `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string         `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string         `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string         `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{} `json:"ServiceLength" binding:"required"`
	Signature                       string         `json:"Signature" binding:"required"`
	SignaturePath                   string         `json:"SignaturePath" binding:"required"`
	SubdivisionOfficeFacilityID     string         `json:"SubdivisionOfficeFacilityID" `
	SubdivisionOfficeName           string         `json:"SubdivisionOfficeName" `
	TempHallTicket                  string         `json:"TempHallTicket" binding:"required"`
	UnitPreferences                 *[]interface{} `json:"UnitPreferences" binding:"required"`
	UserID                          int32          `json:"UserID" binding:"required"`
	WorkingOfficeCircleFacilityID   string         `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string         `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string         `json:"WorkingOfficeDivisionFacilityID" `
	WorkingOfficeDivisionName       string         `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string         `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string         `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32          `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string         `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string         `json:"WorkingOfficeRegionName"`
}

type ReApplicationGDStoPA struct {
	ApplicationID                   int64          `json:"id" binding:"required"`
	ApplicationNumber               string         `json:"ApplicationNumber" binding:"required"`
	Cadre                           string         `json:"Cadre"`
	CandidateRemarks                string         `json:"CandidateRemarks"`
	CategoryCode                    string         `json:"CategoryCode" binding:"required"`
	CategoryDescription             string         `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string         `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32          `json:"CenterId" binding:"required"`
	CentrePreference                string         `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string         `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string         `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string         `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string         `json:"DCCS"`
	DOB                             string         `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string         `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string         `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string         `json:"DeputationOfficeFacilityID"`
	DeputationOfficeName            string         `json:"DeputationOfficeName"`
	DeputationOfficePincode         string         `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string         `json:"DeputationOfficeUniqueId"`
	InDeputation                    string         `json:"InDeputation" `
	DeputationType                  string         `json:"DeputationType"`
	DesignationID                   string         `json:"DesignationID"`
	DisabilityPercentage            int32          `json:"DisabilityPercentage"`
	DisabilityTypeCode              string         `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string         `json:"DisabilityTypeDescription" binding:"required"`
	DisabilityTypeID                string         `json:"DisabilityTypeID"`
	Edges                           EdgesLogdata   `json:"edges" binding:"required"`
	EducationCode                   string         `json:"EducationCode"`
	EducationDescription            string         `json:"EducationDescription" binding:"required"`
	EmailID                         string         `json:"EmailID" binding:"required"`
	EmployeeID                      int64          `json:"EmployeeID" binding:"required"`
	EmployeeName                    string         `json:"EmployeeName" binding:"required"`
	EmployeePost                    string         `json:"EmployeePost" `
	EntryPostCode                   string         `json:"EntryPostCode"`
	EntryPostDescription            string         `json:"EntryPostDescription"`
	ExamCode                        int32          `json:"ExamCode" binding:"required"`
	ExamName                        string         `json:"ExamName" binding:"required"`
	ExamShortName                   string         `json:"ExamShortName" binding:"required"`
	ExamYear                        string         `json:"ExamYear" binding:"required"`
	FacilityName                    string         `json:"FacilityName"`
	FacilityUniqueID                string         `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string         `json:"FeederPostCode" binding:"required"`
	FeederPostDescription           string         `json:"FeederPostDescription" binding:"required"`
	FeederPostJoiningDate           string         `json:"FeederPostJoiningDate" binding:"required"`
	Gender                          string         `json:"Gender" binding:"required"`
	LienControllingOfficeID         string         `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string         `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string         `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string         `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string         `json:"NodalOfficeName" binding:"required"`
	Photo                           string         `json:"Photo" binding:"required"`
	PhotoPath                       string         `json:"PhotoPath" binding:"required"`
	PostPreferences                 *[]interface{} `json:"PostPreferences" binding:"required"`
	PresentDesignation              string         `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string         `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string         `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string         `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string         `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{} `json:"ServiceLength" binding:"required"`
	Signature                       string         `json:"Signature" binding:"required"`
	SignaturePath                   string         `json:"SignaturePath" binding:"required"`
	SubdivisionOfficeFacilityID     string         `json:"SubdivisionOfficeFacilityID" `
	SubdivisionOfficeName           string         `json:"SubdivisionOfficeName" `
	TempHallTicket                  string         `json:"TempHallTicket" binding:"required"`
	UnitPreferences                 *[]interface{} `json:"UnitPreferences" binding:"required"`
	UserID                          int32          `json:"UserID" binding:"required"`
	WorkingOfficeCircleFacilityID   string         `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string         `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string         `json:"WorkingOfficeDivisionFacilityID" `
	WorkingOfficeDivisionName       string         `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string         `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string         `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32          `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string         `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string         `json:"WorkingOfficeRegionName"`
	ApplicationCorrectionLastDate   time.Time      `json:"ApplicationCorrectionLastDate" binding:"required"`
}

type VerifyApplicationGDStoPA struct {
	AppliactionRemarks          string                    `json:"AppliactionRemarks" `
	ApplicationNumber           string                    `json:"ApplicationNumber" binding:"required"`
	CA_EmployeeDesignation      string                    `json:"CA_EmployeeDesignation"  `
	CA_EmployeeID               string                    `json:"CA_EmployeeId" binding:"required"`
	CA_GeneralRemarks           string                    `json:"CA_GeneralRemarks" `
	CA_Remarks                  string                    `json:"CA_Remarks" `
	CA_UserName                 string                    `json:"CA_UserName" binding:"required"`
	ControllingOfficeFacilityID string                    `json:"ControllingOfficeFacilityID" binding:"required"`
	Edges                       EdgesGDSPAApplicationData `json:"edges"`
	EmployeeID                  int64                     `json:"EmployeeID" binding:"required"`
	ExamCode                    int32                     `json:"ExamCode"`
	ExamName                    string                    `json:"ExamName" `
	ExamShortName               string                    `json:"ExamShortName"`
	ExamYear                    string                    `json:"ExamYear"`
	GenerateHallTicketFlag      bool                      `json:"GenerateHallTicketFlag" `
	ID                          int64                     `json:"id" binding:"required"`
	NonQualifyingService        *[]interface{}            `json:"NonQualifyingService" `
	ServiceLength               *[]interface{}            `json:"ServiceLength" binding:"required"`
	UserID                      int32                     `json:"UserID" `
	RecommendedStatus           string                    `json:"RecommendedStatus"`
	PunishmentStatus            bool                      `json:"PunishmentStatus" `       //new coloumn
	DisciplinaryCaseStatus      bool                      `json:"DisciplinaryCaseStatus" ` //new coloumn
}
type NAVerifyApplicationGDStoPA struct {
	AppliactionRemarks          string                      `json:"AppliactionRemarks" `
	ApplicationNumber           string                      `json:"ApplicationNumber" binding:"required"`
	NA_EmployeeDesignation      string                      `json:"NA_EmployeeDesignation"  `
	NA_EmployeeID               string                      `json:"NA_EmployeeId" binding:"required"`
	NA_GeneralRemarks           string                      `json:"NA_GeneralRemarks" `
	NA_Remarks                  string                      `json:"NA_Remarks" `
	NA_UserName                 string                      `json:"NA_UserName" binding:"required"`
	ControllingOfficeFacilityID string                      `json:"ControllingOfficeFacilityID"`
	Edges                       EdgesGDSPAApplicationDataNA `json:"edges"`
	EmployeeID                  int64                       `json:"EmployeeID" binding:"required"`
	ExamCode                    int32                       `json:"ExamCode"`
	ExamName                    string                      `json:"ExamName" `
	ExamShortName               string                      `json:"ExamShortName"`
	ExamYear                    string                      `json:"ExamYear"`
	GenerateHallTicketFlag      bool                        `json:"GenerateHallTicketFlag" `
	ID                          int64                       `json:"id" binding:"required"`
	UserID                      int32                       `json:"UserID" `
	RecommendedStatus           string                      `json:"RecommendedStatus"`
}
type ApplicationPMPA struct {
	Cadre                           string         `json:"Cadre"`
	CandidateRemarks                string         `json:"CandidateRemarks" `
	CategoryCode                    string         `json:"CategoryCode" binding:"required"`
	CategoryDescription             string         `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string         `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32          `json:"CenterId" binding:"required"`
	CentrePreference                string         `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string         `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string         `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string         `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string         `json:"DCCS"`
	DOB                             string         `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string         `json:"DeputationControllingOfficeID" `
	DeputationControllingOfficeName string         `json:"DeputationControllingOfficeName" `
	DeputationOfficeFacilityID      string         `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string         `json:"DeputationOfficeName" `
	DeputationOfficePincode         string         `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string         `json:"DeputationOfficeUniqueId" `
	InDeputation                    string         `json:"InDeputation" `
	DeputationType                  string         `json:"DeputationType" `
	DesignationID                   string         `json:"DesignationID" `
	DisabilityPercentage            int32          `json:"DisabilityPercentage"`
	DisabilityTypeCode              string         `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string         `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string         `json:"DisabilityTypeID"`
	Edges                           EdgesLogdata   `json:"edges" binding:"required"`
	EducationCode                   string         `json:"EducationCode" `
	EducationDescription            string         `json:"EducationDescription" `
	EmailID                         string         `json:"EmailID" binding:"required"`
	EmployeeID                      int64          `json:"EmployeeID" binding:"required"`
	EmployeeName                    string         `json:"EmployeeName" binding:"required"`
	EmployeePost                    string         `json:"EmployeePost"`
	EntryPostCode                   string         `json:"EntryPostCode" binding:"required"`
	EntryPostDescription            string         `json:"EntryPostDescription" binding:"required"`
	ExamCode                        int32          `json:"ExamCode" binding:"required"`
	ExamName                        string         `json:"ExamName" binding:"required"`
	ExamShortName                   string         `json:"ExamShortName" binding:"required"`
	ExamYear                        string         `json:"ExamYear" binding:"required"`
	FacilityName                    string         `json:"FacilityName" `
	FacilityUniqueID                string         `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string         `json:"FeederPostCode" `
	FeederPostDescription           string         `json:"FeederPostDescription" `
	FeederPostJoiningDate           string         `json:"FeederPostJoiningDate" `
	Gender                          string         `json:"Gender" binding:"required"`
	LienControllingOfficeID         string         `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string         `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string         `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string         `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string         `json:"NodalOfficeName" binding:"required"`
	Photo                           string         `json:"Photo" binding:"required"`
	PhotoPath                       string         `json:"PhotoPath" binding:"required"`
	PMMailGuardMTSEngagement        *[]interface{} `json:"PMMailGuardMTSEngagement" binding:"required"`
	PostPreferences                 *[]interface{} `json:"PostPreferences" binding:"required"`
	PresentDesignation              string         `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string         `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string         `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string         `json:"ReportingOfficeFacilityID" binding:"required"`
	ReportingOfficeName             string         `json:"ReportingOfficeName" binding:"required"`
	ServiceLength                   *[]interface{} `json:"ServiceLength" binding:"required"`
	Signature                       string         `json:"Signature" binding:"required"`
	SignaturePath                   string         `json:"SignaturePath" binding:"required"`
	TempHallTicket                  string         `json:"TempHallTicket" binding:"required"`
	UnitPreferences                 *[]interface{} `json:"UnitPreferences" binding:"required"`
	UserID                          int32          `json:"UserID" binding:"required"`
	WorkingOfficeCircleFacilityID   string         `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string         `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string         `json:"WorkingOfficeDivisionFacilityID" `
	WorkingOfficeDivisionName       string         `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string         `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string         `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32          `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string         `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string         `json:"WorkingOfficeRegionName" `
}

type ReApplicationPMPA struct {
	ApplicationID                   int32          `json:"id" binding:"required"`
	ApplicationNumber               string         `json:"ApplicationNumber" binding:"required"`
	Cadre                           string         `json:"Cadre"`
	CandidateRemarks                string         `json:"CandidateRemarks" `
	CategoryCode                    string         `json:"CategoryCode" binding:"required"`
	CategoryDescription             string         `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string         `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32          `json:"CenterId" binding:"required"`
	CentrePreference                string         `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string         `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string         `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string         `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string         `json:"DCCS"`
	DOB                             string         `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string         `json:"DeputationControllingOfficeID" `
	DeputationControllingOfficeName string         `json:"DeputationControllingOfficeName" `
	DeputationOfficeFacilityID      string         `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string         `json:"DeputationOfficeName" `
	DeputationOfficePincode         string         `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string         `json:"DeputationOfficeUniqueId" `
	InDeputation                    string         `json:"InDeputation" `
	DeputationType                  string         `json:"DeputationType" `
	DesignationID                   string         `json:"DesignationID" `
	DisabilityPercentage            int32          `json:"DisabilityPercentage"`
	DisabilityTypeCode              string         `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string         `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string         `json:"DisabilityTypeID"`
	Edges                           EdgesLogdata   `json:"edges" binding:"required"`
	EducationCode                   string         `json:"EducationCode" `
	EducationDescription            string         `json:"EducationDescription" `
	EmailID                         string         `json:"EmailID" binding:"required"`
	EmployeeID                      int64          `json:"EmployeeID" binding:"required"`
	EmployeeName                    string         `json:"EmployeeName" binding:"required"`
	EmployeePost                    string         `json:"EmployeePost" `
	EntryPostCode                   string         `json:"EntryPostCode" binding:"required"`
	EntryPostDescription            string         `json:"EntryPostDescription" binding:"required"`
	ExamCode                        int32          `json:"ExamCode" binding:"required"`
	ExamName                        string         `json:"ExamName" binding:"required"`
	ExamShortName                   string         `json:"ExamShortName" binding:"required"`
	ExamYear                        string         `json:"ExamYear" binding:"required"`
	FacilityName                    string         `json:"FacilityName" `
	FacilityUniqueID                string         `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string         `json:"FeederPostCode" `
	FeederPostDescription           string         `json:"FeederPostDescription" `
	FeederPostJoiningDate           string         `json:"FeederPostJoiningDate" `
	Gender                          string         `json:"Gender" binding:"required"`
	LienControllingOfficeID         string         `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string         `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string         `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string         `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string         `json:"NodalOfficeName" binding:"required"`
	Photo                           string         `json:"Photo" binding:"required"`
	PhotoPath                       string         `json:"PhotoPath" binding:"required"`
	PMMailGuardMTSEngagement        *[]interface{} `json:"PMMailGuardMTSEngagement" binding:"required"`
	PostPreferences                 *[]interface{} `json:"PostPreferences" binding:"required"`
	PresentDesignation              string         `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string         `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string         `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string         `json:"ReportingOfficeFacilityID" binding:"required"`
	ReportingOfficeName             string         `json:"ReportingOfficeName" binding:"required"`
	ServiceLength                   *[]interface{} `json:"ServiceLength" binding:"required"`
	Signature                       string         `json:"Signature" binding:"required"`
	SignaturePath                   string         `json:"SignaturePath" binding:"required"`
	TempHallTicket                  string         `json:"TempHallTicket" binding:"required"`
	UnitPreferences                 *[]interface{} `json:"UnitPreferences" binding:"required"`
	UserID                          int32          `json:"UserID" binding:"required"`
	WorkingOfficeCircleFacilityID   string         `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string         `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string         `json:"WorkingOfficeDivisionFacilityID" `
	WorkingOfficeDivisionName       string         `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string         `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string         `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32          `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string         `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string         `json:"WorkingOfficeRegionName"`
	ApplicationCorrectionLastDate   time.Time      `json:"ApplicationCorrectionLastDate"`
}

type PMPAExamEdges struct {
	PMPAApplicationData []PMPAApplicationData `json:"PMPAApplicationData"`
	LogData             []LogData             `json:"LogData"`
}

type VerifyApplicationPMPA struct {
	AppliactionRemarks          string                   `json:"AppliactionRemarks" `
	ApplicationNumber           string                   `json:"ApplicationNumber" binding:"required"`
	CA_EmployeeDesignation      string                   `json:"CA_EmployeeDesignation"  `
	CA_EmployeeID               string                   `json:"CA_EmployeeId" binding:"required"`
	CA_GeneralRemarks           string                   `json:"CA_GeneralRemarks" `
	CA_Remarks                  string                   `json:"CA_Remarks" `
	CA_UserName                 string                   `json:"CA_UserName" binding:"required"`
	ControllingOfficeFacilityID string                   `json:"ControllingOfficeFacilityID" binding:"required"`
	Edges                       EdgesPMPAApplicationData `json:"edges"`
	EmployeeID                  int64                    `json:"EmployeeID" binding:"required"`
	ExamCode                    int32                    `json:"ExamCode"`
	ExamName                    string                   `json:"ExamName" `
	ExamShortName               string                   `json:"ExamShortName"`
	ExamYear                    string                   `json:"ExamYear"`
	GenerateHallTicketFlag      bool                     `json:"GenerateHallTicketFlag" `
	ID                          int64                    `json:"id" binding:"required"`
	NonQualifyingService        *[]interface{}           `json:"NonQualifyingService" `
	ServiceLength               *[]interface{}           `json:"ServiceLength" binding:"required"`
	UserID                      int32                    `json:"UserID" `
	RecommendedStatus           string                   `json:"RecommendedStatus"`
	PunishmentStatus            bool                     `json:"PunishmentStatus" `       //new coloumn
	DisciplinaryCaseStatus      bool                     `json:"DisciplinaryCaseStatus" ` //new coloumn
}
type NAVerifyApplicationPMPA struct {
	AppliactionRemarks          string                     `json:"AppliactionRemarks" `
	ApplicationNumber           string                     `json:"ApplicationNumber" binding:"required"`
	NA_EmployeeDesignation      string                     `json:"NA_EmployeeDesignation"  `
	NA_EmployeeID               string                     `json:"NA_EmployeeId" binding:"required"`
	NA_GeneralRemarks           string                     `json:"NA_GeneralRemarks" `
	NA_Remarks                  string                     `json:"NA_Remarks" `
	NA_UserName                 string                     `json:"NA_UserName" binding:"required"`
	ControllingOfficeFacilityID string                     `json:"ControllingOfficeFacilityID" `
	Edges                       EdgesPMPAApplicationDataNA `json:"edges"`
	EmployeeID                  int64                      `json:"EmployeeID" binding:"required"`
	ExamCode                    int32                      `json:"ExamCode"`
	ExamName                    string                     `json:"ExamName" `
	ExamShortName               string                     `json:"ExamShortName"`
	ExamYear                    string                     `json:"ExamYear"`
	GenerateHallTicketFlag      bool                       `json:"GenerateHallTicketFlag" `
	ID                          int64                      `json:"id" binding:"required"`
	UserID                      int32                      `json:"UserID" `
	RecommendedStatus           string                     `json:"RecommendedStatus"`
}
type ApplicationGDSPM struct {
	Cadre                           string         `json:"Cadre"`
	CandidateRemarks                string         `json:"CandidateRemarks" `
	CategoryCode                    string         `json:"CategoryCode" binding:"required"`
	CategoryDescription             string         `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string         `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32          `json:"CenterId" binding:"required"`
	CentrePreference                string         `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string         `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string         `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string         `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string         `json:"DCCS"`
	DOB                             string         `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string         `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string         `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string         `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string         `json:"DeputationOfficeName" `
	DeputationOfficePincode         string         `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string         `json:"DeputationOfficeUniqueId"`
	InDeputation                    string         `json:"InDeputation" `
	DeputationType                  string         `json:"DeputationType" `
	DesignationID                   string         `json:"DesignationID" `
	DisabilityPercentage            int32          `json:"DisabilityPercentage" `
	DisabilityTypeCode              string         `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string         `json:"DisabilityTypeDescription" `
	DisabilityTypeID                string         `json:"DisabilityTypeID"`
	Edges                           EdgesLogdata   `json:"edges" binding:"required"`
	EducationCode                   string         `json:"EducationCode"`
	EducationDescription            string         `json:"EducationDescription" binding:"required"`
	EmailID                         string         `json:"EmailID" binding:"required"`
	EmployeeID                      int64          `json:"EmployeeID" binding:"required"`
	EmployeeName                    string         `json:"EmployeeName" binding:"required"`
	EmployeePost                    string         `json:"EmployeePost" `
	EntryPostCode                   string         `json:"EntryPostCode" `
	EntryPostDescription            string         `json:"EntryPostDescription" `
	ExamCode                        int32          `json:"ExamCode" binding:"required"`
	ExamName                        string         `json:"ExamName" binding:"required"`
	ExamShortName                   string         `json:"ExamShortName" binding:"required"`
	ExamYear                        string         `json:"ExamYear" binding:"required"`
	FacilityName                    string         `json:"FacilityName" `
	FacilityUniqueID                string         `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string         `json:"FeederPostCode" `
	FeederPostDescription           string         `json:"FeederPostDescription"`
	FeederPostJoiningDate           string         `json:"FeederPostJoiningDate" binding:"required"`
	Gender                          string         `json:"Gender" binding:"required"`
	LienControllingOfficeID         string         `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string         `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string         `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string         `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string         `json:"NodalOfficeName" binding:"required"`
	Photo                           string         `json:"Photo" binding:"required"`
	PhotoPath                       string         `json:"PhotoPath" binding:"required"`
	PostPreferences                 *[]interface{} `json:"PostPreferences" binding:"required"`
	PresentDesignation              string         `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string         `json:"PresentPostCode" `
	PresentPostDescription          string         `json:"PresentPostDescription" `
	ReportingOfficeFacilityID       string         `json:"ReportingOfficeFacilityID" `
	ReportingOfficeName             string         `json:"ReportingOfficeName" `
	ServiceLength                   *[]interface{} `json:"ServiceLength" binding:"required"`
	SubdivisionOfficeFacilityID     string         `json:"SubdivisionOfficeFacilityID" `
	SubdivisionOfficeName           string         `json:"SubdivisionOfficeName" `
	Signature                       string         `json:"Signature" binding:"required"`
	SignaturePath                   string         `json:"SignaturePath" binding:"required"`
	TempHallTicket                  string         `json:"TempHallTicket" binding:"required"`
	UnitPreferences                 *[]interface{} `json:"UnitPreferences" binding:"required"`
	UserID                          int32          `json:"UserID" binding:"required"`
	WorkingOfficeCircleFacilityID   string         `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string         `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string         `json:"WorkingOfficeDivisionFacilityID" `
	WorkingOfficeDivisionName       string         `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string         `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string         `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32          `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string         `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string         `json:"WorkingOfficeRegionName" `
	GDSEngagement                   *[]interface{} `json:"GDSEngagement" `
	PMMailGuardMTSEngagement        *[]interface{} `json:"PMMailGuardMTSEngagement"`
	ApplicationLastDate             time.Time      `json:"ApplicationLastDate" binding:"required"`
}
type ApplicationGDSPMforUpdateExamCenters struct {
	Cadre                           string         `json:"Cadre"`
	CandidateRemarks                string         `json:"CandidateRemarks"`
	CategoryCode                    string         `json:"CategoryCode"`
	CategoryDescription             string         `json:"CategoryDescription"`
	CenterFacilityId                string         `json:"CenterFacilityId"`
	CenterId                        int32          `json:"CenterId"`
	CentrePreference                string         `json:"CentrePreference"`
	ClaimingQualifyingService       string         `json:"ClaimingQualifyingService"`
	ControllingOfficeFacilityID     string         `json:"ControllingOfficeFacilityID"`
	ControllingOfficeName           string         `json:"ControllingOfficeName"`
	DCCS                            string         `json:"DCCS"`
	DOB                             string         `json:"DOB"`
	DeputationControllingOfficeID   string         `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string         `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string         `json:"DeputationOfficeFacilityID"`
	DeputationOfficeName            string         `json:"DeputationOfficeName"`
	DeputationOfficePincode         string         `json:"DeputationOfficePincode"`
	DeputationOfficeUniqueId        string         `json:"DeputationOfficeUniqueId"`
	InDeputation                    string         `json:"InDeputation"`
	DeputationType                  string         `json:"DeputationType"`
	DesignationID                   string         `json:"DesignationID"`
	DisabilityPercentage            int32          `json:"DisabilityPercentage"`
	DisabilityTypeCode              string         `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string         `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string         `json:"DisabilityTypeID"`
	Edges                           EdgesLogdata   `json:"edges"`
	EducationCode                   string         `json:"EducationCode"`
	EducationDescription            string         `json:"EducationDescription"`
	EmailID                         string         `json:"EmailID"`
	EmployeeID                      int64          `json:"EmployeeID"`
	EmployeeName                    string         `json:"EmployeeName"`
	EmployeePost                    string         `json:"EmployeePost"`
	EntryPostCode                   string         `json:"EntryPostCode"`
	EntryPostDescription            string         `json:"EntryPostDescription"`
	ExamCode                        int32          `json:"ExamCode"`
	ExamName                        string         `json:"ExamName"`
	ExamShortName                   string         `json:"ExamShortName"`
	ExamYear                        string         `json:"ExamYear"`
	FacilityName                    string         `json:"FacilityName"`
	FacilityUniqueID                string         `json:"FacilityUniqueID"`
	FeederPostCode                  string         `json:"FeederPostCode"`
	FeederPostDescription           string         `json:"FeederPostDescription"`
	FeederPostJoiningDate           string         `json:"FeederPostJoiningDate"`
	Gender                          string         `json:"Gender"`
	LienControllingOfficeID         string         `json:"LienControllingOfficeID"`
	LienControllingOfficeName       string         `json:"LienControllingOfficeName"`
	MobileNumber                    string         `json:"MobileNumber"`
	NodalOfficeFacilityID           string         `json:"NodalOfficeFacilityID"`
	NodalOfficeName                 string         `json:"NodalOfficeName"`
	Photo                           string         `json:"Photo"`
	PhotoPath                       string         `json:"PhotoPath"`
	PostPreferences                 *[]interface{} `json:"PostPreferences"`
	PresentDesignation              string         `json:"PresentDesignation"`
	PresentPostCode                 string         `json:"PresentPostCode"`
	PresentPostDescription          string         `json:"PresentPostDescription"`
	ReportingOfficeFacilityID       string         `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string         `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{} `json:"ServiceLength"`
	SubdivisionOfficeFacilityID     string         `json:"SubdivisionOfficeFacilityID"`
	SubdivisionOfficeName           string         `json:"SubdivisionOfficeName"`
	Signature                       string         `json:"Signature"`
	SignaturePath                   string         `json:"SignaturePath"`
	TempHallTicket                  string         `json:"TempHallTicket"`
	UnitPreferences                 *[]interface{} `json:"UnitPreferences"`
	UserID                          int32          `json:"UserID"`
	WorkingOfficeCircleFacilityID   string         `json:"WorkingOfficeCircleFacilityID"`
	WorkingOfficeCircleName         string         `json:"WorkingOfficeCircleName"`
	WorkingOfficeDivisionFacilityID string         `json:"WorkingOfficeDivisionFacilityID"`
	WorkingOfficeDivisionName       string         `json:"WorkingOfficeDivisionName"`
	WorkingOfficeFacilityID         string         `json:"WorkingOfficeFacilityID"`
	WorkingOfficeName               string         `json:"WorkingOfficeName"`
	WorkingOfficePincode            int32          `json:"WorkingOfficePincode"`
	WorkingOfficeRegionFacilityID   string         `json:"WorkingOfficeRegionFacilityID"`
	WorkingOfficeRegionName         string         `json:"WorkingOfficeRegionName"`
}

type ReApplicationGDSPM struct {
	ApplicationID                   int64          `json:"id" binding:"required"`
	ApplicationNumber               string         `json:"ApplicationNumber" binding:"required"`
	ExamName                        string         `json:"ExamName" binding:"required"`
	ExamCode                        int32          `json:"ExamCode" binding:"required"`
	ExamShortName                   string         `json:"ExamShortName" binding:"required"`
	ExamYear                        string         `json:"ExamYear" binding:"required"`
	UserID                          int32          `json:"UserID" binding:"required"`
	EmployeeID                      int64          `json:"EmployeeID" binding:"required"`
	EmployeeName                    string         `json:"EmployeeName" binding:"required"`
	MobileNumber                    string         `json:"MobileNumber" binding:"required"`
	EmailID                         string         `json:"EmailID" binding:"required"`
	Gender                          string         `json:"Gender" binding:"required"`
	GDSEngagement                   *[]interface{} `json:"GDSEngagement" `
	CandidateRemarks                string         `json:"CandidateRemarks" `
	CategoryDescription             string         `json:"CategoryDescription" binding:"required"`
	CategoryCode                    string         `json:"CategoryCode" binding:"required"`
	DOB                             string         `json:"DOB" binding:"required"`
	DisabilityTypeID                string         `json:"DisabilityTypeID"`
	DisabilityTypeCode              string         `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string         `json:"DisabilityTypeDescription" binding:"required"`
	DisabilityPercentage            int32          `json:"DisabilityPercentage" `
	DCCS                            string         `json:"DCCS"`
	PresentPostCode                 string         `json:"PresentPostCode" `
	PresentPostDescription          string         `json:"PresentPostDescription" `
	FeederPostCode                  string         `json:"FeederPostCode" `
	FeederPostDescription           string         `json:"FeederPostDescription"`
	DesignationID                   string         `json:"DesignationID" `
	ClaimingQualifyingService       string         `json:"ClaimingQualifyingService" binding:"required"`
	CenterId                        int32          `json:"CenterId" binding:"required"`
	PresentDesignation              string         `json:"PresentDesignation" binding:"required"`
	EducationCode                   string         `json:"EducationCode"`
	EducationDescription            string         `json:"EducationDescription" `
	FacilityUniqueID                string         `json:"FacilityUniqueID" binding:"required"`
	WorkingOfficePincode            int32          `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeFacilityID         string         `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string         `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficeCircleFacilityID   string         `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string         `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeRegionFacilityID   string         `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string         `json:"WorkingOfficeRegionName" `
	WorkingOfficeDivisionFacilityID string         `json:"WorkingOfficeDivisionFacilityID" `
	WorkingOfficeDivisionName       string         `json:"WorkingOfficeDivisionName" `
	ReportingOfficeFacilityID       string         `json:"ReportingOfficeFacilityID" `
	ReportingOfficeName             string         `json:"ReportingOfficeName" `
	LienControllingOfficeID         string         `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string         `json:"LienControllingOfficeName" binding:"required"`
	InDeputation                    string         `json:"InDeputation" `
	DeputationType                  string         `json:"DeputationType" `
	DeputationOfficeUniqueId        string         `json:"DeputationOfficeUniqueId"`
	DeputationOfficeFacilityID      string         `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string         `json:"DeputationOfficeName" `
	DeputationControllingOfficeID   string         `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string         `json:"DeputationControllingOfficeName"`
	ControllingOfficeFacilityID     string         `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string         `json:"ControllingOfficeName" binding:"required"`
	NodalOfficeFacilityID           string         `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string         `json:"NodalOfficeName" binding:"required"`
	CenterFacilityId                string         `json:"CenterFacilityId" binding:"required"`
	CentrePreference                string         `json:"CentrePreference" binding:"required"`
	Photo                           string         `json:"Photo" binding:"required"`
	PhotoPath                       string         `json:"PhotoPath" binding:"required"`
	SignaturePath                   string         `json:"SignaturePath" binding:"required"`
	Signature                       string         `json:"Signature" binding:"required"`
	TempHallTicket                  string         `json:"TempHallTicket" binding:"required"`
	Remarks                         string         `json:"Remarks" `
	ServiceLength                   *[]interface{} `json:"ServiceLength" binding:"required"`
	Edges                           EdgesLogdata   `json:"edges" binding:"required"`
	DeputationOfficePincode         string         `json:"DeputationOfficePincode" `
	Cadre                           string         `json:"Cadre"`
	EmployeePost                    string         `json:"EmployeePost" `
	FacilityName                    string         `json:"FacilityName" `
	EntryPostDescription            string         `json:"EntryPostDescription" `
	FeederPostJoiningDate           string         `json:"FeederPostJoiningDate" binding:"required"`
	DCInPresentCadre                string         `json:"DCInPresentCadre"`
	EntryPostCode                   string         `json:"EntryPostCode" `
	PostPreferences                 *[]interface{} `json:"PostPreferences" binding:"required"`
	UnitPreferences                 *[]interface{} `json:"UnitPreferences" binding:"required"`
	ApplicationCorrectionLastDate   time.Time      `json:"ApplicationCorrectionLastDate" binding:"required"`
}

type VerifyApplicationGDSPM struct {
	AppliactionRemarks     string                    `json:"AppliactionRemarks" `
	ApplicationNumber      string                    `json:"ApplicationNumber" binding:"required"`
	CA_EmployeeDesignation string                    `json:"CA_EmployeeDesignation"  `
	CA_EmployeeID          string                    `json:"CA_EmployeeId" binding:"required"`
	CA_GeneralRemarks      string                    `json:"CA_GeneralRemarks" `
	CA_Remarks             string                    `json:"CA_Remarks" `
	CA_UserName            string                    `json:"CA_UserName" binding:"required"`
	Post                   string                    `json:"Post" `
	Edges                  EdgesGDSPMApplicationData `json:"edges"`
	EmployeeID             int64                     `json:"EmployeeID" binding:"required"`
	ExamCode               int32                     `json:"ExamCode"`
	ExamName               string                    `json:"ExamName"`
	ExamShortName          string                    `json:"ExamShortName"`
	ExamYear               string                    `json:"ExamYear"`
	GenerateHallTicketFlag bool                      `json:"GenerateHallTicketFlag" `
	ID                     int64                     `json:"id" binding:"required"`
	NonQualifyingService   *[]interface{}            `json:"NonQualifyingService" `
	ServiceLength          *[]interface{}            `json:"ServiceLength" binding:"required"`
	UserID                 int32                     `json:"UserID"`
	RecommendedStatus      string                    `json:"RecommendedStatus"`
	PunishmentStatus       bool                      `json:"PunishmentStatus" `       //new coloumn
	DisciplinaryCaseStatus bool                      `json:"DisciplinaryCaseStatus" ` //new coloumn
}
type NAVerifyApplicationGDSPM struct {
	AppliactionRemarks     string                       `json:"AppliactionRemarks" `
	ApplicationNumber      string                       `json:"ApplicationNumber" binding:"required"`
	NA_EmployeeDesignation string                       `json:"NA_EmployeeDesignation"  `
	NA_EmployeeID          string                       `json:"NA_EmployeeId" binding:"required"`
	NA_GeneralRemarks      string                       `json:"NA_GeneralRemarks"`
	NA_Remarks             string                       `json:"NA_Remarks" `
	NA_UserName            string                       `json:"NA_UserName" binding:"required"`
	Edges                  EdgesGDSPMApplicationDataNAv `json:"edges"`
	EmployeeID             int64                        `json:"EmployeeID" binding:"required"`
	ExamCode               int32                        `json:"ExamCode"`
	ExamName               string                       `json:"ExamName"`
	ExamShortName          string                       `json:"ExamShortName"`
	ExamYear               string                       `json:"ExamYear"`
	GenerateHallTicketFlag bool                         `json:"GenerateHallTicketFlag" `
	ID                     int64                        `json:"id" binding:"required"`
	UserID                 int32                        `json:"UserID"`
	RecommendedStatus      string                       `json:"RecommendedStatus"`
}
type ApplicationMTSPM struct {
	Cadre                           string         `json:"Cadre"`
	CandidateRemarks                string         `json:"CandidateRemarks" `
	CategoryCode                    string         `json:"CategoryCode" binding:"required"`
	CategoryDescription             string         `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string         `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32          `json:"CenterId" binding:"required"`
	CentrePreference                string         `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string         `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string         `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string         `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string         `json:"DCCS"`
	DOB                             string         `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string         `json:"DeputationControllingOfficeID" `
	DeputationControllingOfficeName string         `json:"DeputationControllingOfficeName" `
	DeputationOfficeFacilityID      string         `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string         `json:"DeputationOfficeName" `
	DeputationOfficePincode         string         `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string         `json:"DeputationOfficeUniqueId" `
	InDeputation                    string         `json:"InDeputation" `
	DeputationType                  string         `json:"DeputationType"`
	DesignationID                   string         `json:"DesignationID"`
	DisabilityPercentage            int32          `json:"DisabilityPercentage"`
	DisabilityTypeCode              string         `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string         `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string         `json:"DisabilityTypeID"`
	Edges                           EdgesLogdata   `json:"edges" binding:"required"`
	EducationCode                   string         `json:"EducationCode"`
	EducationDescription            string         `json:"EducationDescription"`
	EmailID                         string         `json:"EmailID" binding:"required"`
	EmployeeID                      int64          `json:"EmployeeID" binding:"required"`
	EmployeeName                    string         `json:"EmployeeName" binding:"required"`
	EmployeePost                    string         `json:"EmployeePost" `
	EntryPostCode                   string         `json:"EntryPostCode" binding:"required" `
	EntryPostDescription            string         `json:"EntryPostDescription" binding:"required"`
	ExamCode                        int32          `json:"ExamCode" binding:"required"`
	ExamName                        string         `json:"ExamName" binding:"required"`
	ExamShortName                   string         `json:"ExamShortName" binding:"required"`
	ExamYear                        string         `json:"ExamYear" binding:"required"`
	FacilityName                    string         `json:"FacilityName"`
	FacilityUniqueID                string         `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string         `json:"FeederPostCode" binding:"required"`
	FeederPostDescription           string         `json:"FeederPostDescription" binding:"required"`
	FeederPostJoiningDate           string         `json:"FeederPostJoiningDate" binding:"required"`
	Gender                          string         `json:"Gender" binding:"required"`
	GDSEngagement                   *[]interface{} `json:"GDSEngagement" `
	LienControllingOfficeID         string         `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string         `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string         `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string         `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string         `json:"NodalOfficeName" binding:"required"`
	Photo                           string         `json:"Photo" binding:"required"`
	PhotoPath                       string         `json:"PhotoPath" binding:"required"`
	PostPreferences                 *[]interface{} `json:"PostPreferences" binding:"required"`
	PresentDesignation              string         `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string         `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string         `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string         `json:"ReportingOfficeFacilityID" binding:"required"`
	ReportingOfficeName             string         `json:"ReportingOfficeName" binding:"required"`
	ServiceLength                   *[]interface{} `json:"ServiceLength" binding:"required"`
	Signature                       string         `json:"Signature" binding:"required"`
	SignaturePath                   string         `json:"SignaturePath" binding:"required"`
	TempHallTicket                  string         `json:"TempHallTicket" binding:"required"`
	UnitPreferences                 *[]interface{} `json:"UnitPreferences" binding:"required"`
	UserID                          int32          `json:"UserID" binding:"required"`
	WorkingOfficeCircleFacilityID   string         `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string         `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string         `json:"WorkingOfficeDivisionFacilityID" `
	WorkingOfficeDivisionName       string         `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string         `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string         `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32          `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string         `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string         `json:"WorkingOfficeRegionName" `
}
type ApplicationMTSPMUpdatecenter struct {
	Cadre                           string         `json:"Cadre"`
	CandidateRemarks                string         `json:"CandidateRemarks" `
	CategoryCode                    string         `json:"CategoryCode"`
	CategoryDescription             string         `json:"CategoryDescription"`
	CenterFacilityId                string         `json:"CenterFacilityId" `
	CenterId                        int32          `json:"CenterId" `
	CentrePreference                string         `json:"CentrePreference"`
	ClaimingQualifyingService       string         `json:"ClaimingQualifyingService"`
	ControllingOfficeFacilityID     string         `json:"ControllingOfficeFacilityID" `
	ControllingOfficeName           string         `json:"ControllingOfficeName" `
	DCCS                            string         `json:"DCCS"`
	DOB                             string         `json:"DOB" `
	DeputationControllingOfficeID   string         `json:"DeputationControllingOfficeID" `
	DeputationControllingOfficeName string         `json:"DeputationControllingOfficeName" `
	DeputationOfficeFacilityID      string         `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string         `json:"DeputationOfficeName" `
	DeputationOfficePincode         string         `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string         `json:"DeputationOfficeUniqueId" `
	InDeputation                    string         `json:"InDeputation" `
	DeputationType                  string         `json:"DeputationType"`
	DesignationID                   string         `json:"DesignationID"`
	DisabilityPercentage            int32          `json:"DisabilityPercentage"`
	DisabilityTypeCode              string         `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string         `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string         `json:"DisabilityTypeID"`
	Edges                           EdgesLogdata   `json:"edges" `
	EducationCode                   string         `json:"EducationCode"`
	EducationDescription            string         `json:"EducationDescription"`
	EmailID                         string         `json:"EmailID"`
	EmployeeID                      int64          `json:"EmployeeID" `
	EmployeeName                    string         `json:"EmployeeName"`
	EmployeePost                    string         `json:"EmployeePost" `
	EntryPostCode                   string         `json:"EntryPostCode"`
	EntryPostDescription            string         `json:"EntryPostDescription"`
	ExamCode                        int32          `json:"ExamCode"`
	ExamName                        string         `json:"ExamName"`
	ExamShortName                   string         `json:"ExamShortName"`
	ExamYear                        string         `json:"ExamYear"`
	FacilityName                    string         `json:"FacilityName"`
	FacilityUniqueID                string         `json:"FacilityUniqueID"`
	FeederPostCode                  string         `json:"FeederPostCode"`
	FeederPostDescription           string         `json:"FeederPostDescription"`
	FeederPostJoiningDate           string         `json:"FeederPostJoiningDate"`
	Gender                          string         `json:"Gender" `
	GDSEngagement                   *[]interface{} `json:"GDSEngagement" `
	LienControllingOfficeID         string         `json:"LienControllingOfficeID"`
	LienControllingOfficeName       string         `json:"LienControllingOfficeName"`
	MobileNumber                    string         `json:"MobileNumber" `
	NodalOfficeFacilityID           string         `json:"NodalOfficeFacilityID"`
	NodalOfficeName                 string         `json:"NodalOfficeName"`
	Photo                           string         `json:"Photo"`
	PhotoPath                       string         `json:"PhotoPath"`
	PostPreferences                 *[]interface{} `json:"PostPreferences"`
	PresentDesignation              string         `json:"PresentDesignation"`
	PresentPostCode                 string         `json:"PresentPostCode"`
	PresentPostDescription          string         `json:"PresentPostDescription"`
	ReportingOfficeFacilityID       string         `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string         `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{} `json:"ServiceLength"`
	Signature                       string         `json:"Signature"`
	SignaturePath                   string         `json:"SignaturePath"`
	TempHallTicket                  string         `json:"TempHallTicket"`
	UnitPreferences                 *[]interface{} `json:"UnitPreferences"`
	UserID                          int32          `json:"UserID"`
	WorkingOfficeCircleFacilityID   string         `json:"WorkingOfficeCircleFacilityID"`
	WorkingOfficeCircleName         string         `json:"WorkingOfficeCircleName"`
	WorkingOfficeDivisionFacilityID string         `json:"WorkingOfficeDivisionFacilityID" `
	WorkingOfficeDivisionName       string         `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string         `json:"WorkingOfficeFacilityID"`
	WorkingOfficeName               string         `json:"WorkingOfficeName"`
	WorkingOfficePincode            int32          `json:"WorkingOfficePincode"`
	WorkingOfficeRegionFacilityID   string         `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string         `json:"WorkingOfficeRegionName" `
}

type VerifyApplicationMTSPM struct {
	AppliactionRemarks     string                    `json:"AppliactionRemarks" `
	ApplicationNumber      string                    `json:"ApplicationNumber" binding:"required"`
	CA_EmployeeDesignation string                    `json:"CA_EmployeeDesignation"  `
	CA_EmployeeID          string                    `json:"CA_EmployeeId" binding:"required"`
	CA_GeneralRemarks      string                    `json:"CA_GeneralRemarks" `
	CA_Remarks             string                    `json:"CA_Remarks" `
	CA_UserName            string                    `json:"CA_UserName" binding:"required"`
	Edges                  EdgesMTSPMApplicationData `json:"edges"`
	EmployeeID             int64                     `json:"EmployeeID" binding:"required"`
	ExamCode               int32                     `json:"ExamCode"`
	ExamName               string                    `json:"ExamName" `
	ExamShortName          string                    `json:"ExamShortName"`
	ExamYear               string                    `json:"ExamYear"`
	GenerateHallTicketFlag bool                      `json:"GenerateHallTicketFlag" `
	ID                     int64                     `json:"id" binding:"required"`
	NonQualifyingService   *[]interface{}            `json:"NonQualifyingService" `
	ServiceLength          *[]interface{}            `json:"ServiceLength" binding:"required"`
	UserID                 int32                     `json:"UserID" `
	RecommendedStatus      string                    `json:"RecommendedStatus"`
	PunishmentStatus       bool                      `json:"PunishmentStatus" `       //new coloumn
	DisciplinaryCaseStatus bool                      `json:"DisciplinaryCaseStatus" ` //new coloumn
}
type NAVerifyApplicationMTSPM struct {
	AppliactionRemarks     string                      `json:"AppliactionRemarks" `
	ApplicationNumber      string                      `json:"ApplicationNumber" binding:"required"`
	NA_EmployeeDesignation string                      `json:"NA_EmployeeDesignation"  `
	NA_EmployeeID          string                      `json:"NA_EmployeeId" binding:"required"`
	NA_GeneralRemarks      string                      `json:"NA_GeneralRemarks" `
	NA_Remarks             string                      `json:"NA_Remarks" `
	NA_UserName            string                      `json:"NA_UserName" binding:"required"`
	Edges                  EdgesMTSPMApplicationDataNA `json:"edges"`
	EmployeeID             int64                       `json:"EmployeeID" binding:"required"`
	ExamCode               int32                       `json:"ExamCode"`
	ExamName               string                      `json:"ExamName" `
	ExamShortName          string                      `json:"ExamShortName"`
	ExamYear               string                      `json:"ExamYear"`
	GenerateHallTicketFlag bool                        `json:"GenerateHallTicketFlag" `
	ID                     int64                       `json:"id" binding:"required"`
	UserID                 int32                       `json:"UserID" `
	RecommendedStatus      string                      `json:"RecommendedStatus"`
}

type VerifyApplicationVAGDSPM struct {
	ApplicationNumber      string       `json:"ApplicationNumber" binding:"required"`
	VA_EmployeeDesignation string       `json:"VA_EmployeeDesignation"`
	VA_EmployeeID          string       `json:"VA_EmployeeId" binding:"required"`
	VA_GeneralRemarks      string       `json:"VA_GeneralRemarks" binding:"required"`
	VA_UserName            string       `json:"VA_UserName" binding:"required"`
	Edges                  EdgesLogdata `json:"edges"`
	EmployeeID             int64        `json:"EmployeeID" binding:"required"`
	ExamCode               int32        `json:"ExamCode" binding:"required"`
	ExamName               string       `json:"ExamName" binding:"required"`
	ExamShortName          string       `json:"ExamShortName" binding:"required"`
	ExamYear               string       `json:"ExamYear" binding:"required"`
	ID                     int64        `json:"id" binding:"required"`
}

type VerifyApplicationIp struct {
	AppliactionRemarks     string                `json:"AppliactionRemarks"`
	ApplicationNumber      string                `json:"ApplicationNumber" binding:"required"`
	CA_EmployeeDesignation string                `json:"CA_EmployeeDesignation"`
	CA_EmployeeID          string                `json:"CA_EmployeeId" binding:"required"`
	CA_GeneralRemarks      string                `json:"CA_GeneralRemarks"`
	CA_Remarks             string                `json:"CA_Remarks" `
	CA_UserName            string                `json:"CA_UserName" binding:"required"`
	Edges                  EdgesIPApplicatinData `json:"edges"`
	EmployeeID             int64                 `json:"EmployeeID" binding:"required"`
	ExamCode               int32                 `json:"ExamCode"`
	ExamName               string                `json:"ExamName" `
	ExamShortName          string                `json:"ExamShortName"`
	ExamYear               string                `json:"ExamYear"`
	GenerateHallTicketFlag bool                  `json:"GenerateHallTicketFlag" `
	ID                     int64                 `json:"id" binding:"required"`
	NonQualifyingService   *[]interface{}        `json:"NonQualifyingService" `
	ServiceLength          *[]interface{}        `json:"ServiceLength" binding:"required"`
	UserID                 int32                 `json:"UserID" `
	RecommendedStatus      string                `json:"RecommendedStatus"`
	PunishmentStatus       bool                  `json:"PunishmentStatus"`       //new coloumn
	DisciplinaryCaseStatus bool                  `json:"DisciplinaryCaseStatus"` //new coloumn
}
type NAApplicationIpCenterChange struct {
	AppliactionRemarks     string                  `json:"AppliactionRemarks" `
	ApplicationNumber      string                  `json:"ApplicationNumber" binding:"required"`
	CenterFacilityId       string                  `json:"CenterFacilityId" binding:"required"`
	NA_EmployeeDesignation string                  `json:"NA_EmployeeDesignation"  `
	NA_EmployeeID          string                  `json:"NA_EmployeeId" binding:"required"`
	NA_GeneralRemarks      string                  `json:"NA_GeneralRemarks" `
	NA_Remarks             string                  `json:"NA_Remarks" `
	NA_UserName            string                  `json:"NA_UserName" binding:"required"`
	Edges                  EdgesIPApplicatinDataNA `json:"edges"`
	EmployeeID             int64                   `json:"EmployeeID" binding:"required"`
	ExamCode               int32                   `json:"ExamCode"`
	ExamName               string                  `json:"ExamName" `
	ExamShortName          string                  `json:"ExamShortName"`
	ExamYear               string                  `json:"ExamYear"`
	GenerateHallTicketFlag bool                    `json:"GenerateHallTicketFlag" `
	ID                     int64                   `json:"id" binding:"required"`
	UserID                 int32                   `json:"UserID" `
	CenterId               int32                   `json:"CenterId" binding:"required"`
	CentrePreference       string                  `json:"CentrePreference" binding:"required"`
}
type ApplicationIpEditCA struct {
	ApplicationID                   int32           `json:"id"`
	ApplicationNumber               string          `json:"ApplicationNumber"`
	CandidateRemarks                string          `json:"CandidateRemarks"`
	CategoryCode                    string          `json:"CategoryCode" binding:"required"`
	CategoryDescription             string          `json:"CategoryDescription" binding:"required"`
	CenterFacilityId                string          `json:"CenterFacilityId" binding:"required"`
	CenterId                        int32           `json:"CenterId" binding:"required"`
	CentrePreference                string          `json:"CentrePreference" binding:"required"`
	ClaimingQualifyingService       string          `json:"ClaimingQualifyingService" binding:"required"`
	ControllingOfficeFacilityID     string          `json:"ControllingOfficeFacilityID" binding:"required"`
	ControllingOfficeName           string          `json:"ControllingOfficeName" binding:"required"`
	DCCS                            string          `json:"DCCS"`
	DOB                             string          `json:"DOB" binding:"required"`
	DeputationControllingOfficeID   string          `json:"DeputationControllingOfficeID"`
	DeputationControllingOfficeName string          `json:"DeputationControllingOfficeName"`
	DeputationOfficeFacilityID      string          `json:"DeputationOfficeFacilityID" `
	DeputationOfficeName            string          `json:"DeputationOfficeName"`
	DeputationOfficePincode         string          `json:"DeputationOfficePincode" `
	DeputationOfficeUniqueId        string          `json:"DeputationOfficeUniqueId"`
	InDeputation                    string          `json:"InDeputation" `
	DeputationType                  string          `json:"DeputationType"`
	DesignationID                   string          `json:"DesignationID"`
	DisabilityPercentage            int32           `json:"DisabilityPercentage"`
	DisabilityTypeCode              string          `json:"DisabilityTypeCode"`
	DisabilityTypeDescription       string          `json:"DisabilityTypeDescription"`
	DisabilityTypeID                string          `json:"DisabilityTypeID"`
	Edges                           EdgesCircleData `json:"edges" binding:"required"`
	EducationCode                   string          `json:"EducationCode"`
	EducationDescription            string          `json:"EducationDescription"`
	EmailID                         string          `json:"EmailID" binding:"required"`
	EmployeeID                      int64           `json:"EmployeeID" binding:"required"`
	EmployeeName                    string          `json:"EmployeeName" binding:"required"`
	EntryPostCode                   string          `json:"EntryPostCode"`
	EntryPostDescription            string          `json:"EntryPostDescription"`
	ExamCode                        int32           `json:"ExamCode" binding:"required"`
	ExamName                        string          `json:"ExamName" binding:"required"`
	ExamShortName                   string          `json:"ExamShortName" binding:"required"`
	ExamYear                        string          `json:"ExamYear" binding:"required"`
	FacilityUniqueID                string          `json:"FacilityUniqueID" binding:"required"`
	FeederPostCode                  string          `json:"FeederPostCode" binding:"required"`
	FeederPostDescription           string          `json:"FeederPostDescription" binding:"required"`
	FeederPostJoiningDate           string          `json:"FeederPostJoiningDate" binding:"required"`
	Gender                          string          `json:"Gender" binding:"required"`
	LienControllingOfficeID         string          `json:"LienControllingOfficeID" binding:"required"`
	LienControllingOfficeName       string          `json:"LienControllingOfficeName" binding:"required"`
	MobileNumber                    string          `json:"MobileNumber" binding:"required"`
	NodalOfficeFacilityID           string          `json:"NodalOfficeFacilityID" binding:"required"`
	NodalOfficeName                 string          `json:"NodalOfficeName" binding:"required"`
	Photo                           string          `json:"Photo" binding:"required"`
	PhotoPath                       string          `json:"PhotoPath" binding:"required"`
	PresentDesignation              string          `json:"PresentDesignation" binding:"required"`
	PresentPostCode                 string          `json:"PresentPostCode" binding:"required"`
	PresentPostDescription          string          `json:"PresentPostDescription" binding:"required"`
	ReportingOfficeFacilityID       string          `json:"ReportingOfficeFacilityID"`
	ReportingOfficeName             string          `json:"ReportingOfficeName"`
	ServiceLength                   *[]interface{}  `json:"ServiceLength" binding:"required"`
	Signature                       string          `json:"Signature" binding:"required"`
	SignaturePath                   string          `json:"SignaturePath" binding:"required"`
	TempHallTicket                  string          `json:"TempHallTicket" binding:"required"`
	UserID                          int32           `json:"UserID"`
	WorkingOfficeCircleFacilityID   string          `json:"WorkingOfficeCircleFacilityID" binding:"required"`
	WorkingOfficeCircleName         string          `json:"WorkingOfficeCircleName" binding:"required"`
	WorkingOfficeDivisionFacilityID string          `json:"WorkingOfficeDivisionFacilityID"`
	WorkingOfficeDivisionName       string          `json:"WorkingOfficeDivisionName" `
	WorkingOfficeFacilityID         string          `json:"WorkingOfficeFacilityID" binding:"required"`
	WorkingOfficeName               string          `json:"WorkingOfficeName" binding:"required"`
	WorkingOfficePincode            int32           `json:"WorkingOfficePincode" binding:"required"`
	WorkingOfficeRegionFacilityID   string          `json:"WorkingOfficeRegionFacilityID" `
	WorkingOfficeRegionName         string          `json:"WorkingOfficeRegionName" `
	Cadre                           string          `json:"Cadre"`
	CA_EmployeeID                   string          `json:"CA_EmployeeId" binding:"required"`
	CA_EmployeeDesignation          string          `json:"CA_EmployeeDesignation"  `
	CA_GeneralRemarks               string          `json:"CA_GeneralRemarks" `
	CA_UserName                     string          `json:"CA_UserName" binding:"required"`
	CA_UserID                       int32           `json:"CA_UserID" binding:"required"`
}
type NAVerifyApplicationIp struct {
	AppliactionRemarks     string                  `json:"AppliactionRemarks" `
	ApplicationNumber      string                  `json:"ApplicationNumber" binding:"required"`
	NA_EmployeeDesignation string                  `json:"NA_EmployeeDesignation"  `
	NA_EmployeeID          string                  `json:"NA_EmployeeId" binding:"required"`
	NA_GeneralRemarks      string                  `json:"NA_GeneralRemarks" `
	NA_Remarks             string                  `json:"NA_Remarks" `
	NA_UserName            string                  `json:"NA_UserName" binding:"required"`
	Edges                  EdgesIPApplicatinDataNA `json:"edges"`
	EmployeeID             int64                   `json:"EmployeeID" binding:"required"`
	ExamCode               int32                   `json:"ExamCode"`
	ExamName               string                  `json:"ExamName" `
	ExamShortName          string                  `json:"ExamShortName"`
	ExamYear               string                  `json:"ExamYear"`
	GenerateHallTicketFlag bool                    `json:"GenerateHallTicketFlag" `
	ID                     int64                   `json:"id" binding:"required"`
	UserID                 int32                   `json:"UserID" `
	RecommendedStatus      string                  `json:"RecommendedStatus"`
}

type ResponseVerifyCandidateUserLogin struct {
	UserID           int64  `json:"UserID" `
	EmployeeID       int64  `json:"EmployeeID" `
	EmployeeName     string `json:"EmployeeName" `
	RoleUserCode     int32  `json:"RoleUserCode" `
	UserName         string `json:"UserName" `
	Mobile           string `json:"Mobile" `
	Email            string `json:"EmailID" `
	DOB              string `json:"DOB" `
	Gender           string `json:"Gender" `
	EmployeeCategory string `json:"EmployeeCategory" `
	EmployeePost     string `json:"EmployeePost" `
	FacilityId       string `json:"FacilityID" `
	CircleFacilityId string `json:"CircleFacilityID" `
	Token            string `json:"Token" `
}

type CreateUserRequest struct {
	EmployeeID              int64     `json:"EmployeeID" binding:"required"`
	EmployeeName            string    `json:"EmployeeName" binding:"required"`
	Mobile                  string    `json:"Mobile" binding:"required"`
	EmailID                 string    `json:"EmailID" binding:"required"`
	UserName                string    `json:"UserName" binding:"required"`
	Password                string    `json:"Password" binding:"required"`
	Status                  bool      `json:"Status"`
	Statuss                 string    `json:"Statuss"`
	OTP                     int32     `json:"OTP"`
	OTPRemarks              string    `json:"OTPRemarks"`
	CreatedAt               time.Time `json:"CreatedAt"`
	OTPTriggeredTime        time.Time `json:"OTPTriggeredTime"`
	OTPSavedTime            time.Time `json:"OTPSavedTime"`
	OTPExpiryTime           time.Time `json:"OTPExpiryTime"`
	NewPasswordRequest      bool      `json:"NewPasswordRequest"`
	EmailOTP                int32     `json:"EmailOTP"`
	EmailOTPRemarks         string    `json:"EmailOTPRemarks"`
	EmailCreatedAt          time.Time `json:"EmailCreatedAt"`
	EmailOTPTriggeredTime   time.Time `json:"EmailOTPTriggeredTime"`
	EmailOTPSavedTime       time.Time `json:"EmailOTPSavedTime"`
	EmailOTPExpiryTime      time.Time `json:"EmailOTPExpiryTime"`
	EmailNewPasswordRequest bool      `json:"EmailNewPasswordRequest"`
	UidToken                string    `json:"UidToken"`
	CreatedById             int64     `json:"CreatedById"`
	CreatedByEmployeeId     string    `json:"CreatedByEmployeeId"`
	CreatedByUserName       string    `json:"CreatedByUserName"`
	CreatedByDesignation    string    `json:"CreatedByDesignation"`
	CreatedDate             time.Time `json:"CreatedDate"`
	DeletedById             int64     `json:"DeletedById"`
	DeletedByEmployeeId     string    `json:"DeletedByEmployeeId"`
	DeletedByUserName       string    `json:"DeletedByUserName"`
	DeletedByDesignation    string    `json:"DeletedByDesignation"`
	DeletedDate             time.Time `json:"DeletedDate"`
	FacilityID              string    `json:"FacilityID"`
	CircleFacilityId        string    `json:"CircleFacilityId"`
	CircleFacilityName      string    `json:"CircleFacilityName"`
	Designation             string    `json:"Designation"`
	RoleUserCode            int32     `json:"RoleUserCode"`
	Updatedby               string    `json:"Updatedby"`
	UpdatedDate             time.Time `json:"UpdatedDate"`
	Modifiedby              string    `json:"Modifiedby"`
	ModifiedDate            time.Time `json:"ModifiedDate"`
	Operationstatus         string    `json:"Operationstatus"`
	ExamCode                int32     `json:"ExamCode"`
	ExamCodePS              int32     `json:"ExamCodePS"`
	Gender                  string    `json:"Gender"`
	DOB                     string    `json:"DOB"`
	CreatedBy               string    `json:"CreatedBy"`
}

type UserRequest struct {
	ID                  int       `json:"id"`
	EmployeedID         string    `json:"EmployeedID"`
	IDVerified          bool      `json:"IDVerified"`
	IDRemStatus         bool      `json:"IDRemStatus"`
	IDRemarks           string    `json:"IDRemarks"`
	EmployeedName       string    `json:"EmployeedName"`
	NameVerified        bool      `json:"NameVerified"`
	NameRemStatus       bool      `json:"NameRemStatus"`
	NameRemarks         string    `json:"NameRemarks"`
	DOB                 time.Time `json:"DOB"`
	DOBVerified         bool      `json:"DOBVerified"`
	DOBRemStatus        bool      `json:"DOBRemStatus"`
	DOBRemarks          string    `json:"DOBRemarks"`
	Gender              string    `json:"Gender"`
	GenderVerified      bool      `json:"GenderVerified"`
	GenderRemStatus     bool      `json:"GenderRemStatus"`
	GenderRemarks       string    `json:"GenderRemarks"`
	Cadreid             int32     `json:"Cadreid"`
	CadreidVerified     bool      `json:"CadreidVerified"`
	CadreidRemStatus    bool      `json:"CadreidRemStatus"`
	CadreidRemarks      string    `json:"CadreidRemarks"`
	OfficeID            int32     `json:"OfficeID"`
	OfficeIDVerified    bool      `json:"OfficeIDVerified"`
	OfficeIDRemStatus   bool      `json:"OfficeIDRemStatus"`
	OfficeIDRemarks     string    `json:"OfficeIDRemarks"`
	PH                  bool      `json:"PH"`
	PHVerified          bool      `json:"PHVerified"`
	PHRemStatus         bool      `json:"PHRemStatus"`
	PHRemarks           string    `json:"PHRemarks"`
	PHDetails           string    `json:"PHDetails"`
	PHDetailsVerified   bool      `json:"PHDetailsVerified"`
	PHDetailsRemStatus  bool      `json:"PHDetailsRemStatus"`
	PHDetailsRemarks    string    `json:"PHDetailsRemarks"`
	APSWorking          bool      `json:"APSWorking"`
	APSWorkingVerified  bool      `json:"APSWorkingVerified"`
	APSWorkingRemStatus bool      `json:"APSWorkingRemStatus"`
	APSWorkingRemarks   string    `json:"APSWorkingRemarks"`
	Profilestatus       bool      `json:"Profilestatus"`
}

type CreateNotificationStruct struct {
	ExamName                               string         `json:"ExamName" binding:"required"`
	ExamCode                               int32          `json:"ExamCode" binding:"required"`
	UserName                               string         `json:"UserName" binding:"required"`
	IssuedBy                               string         `json:"IssuedBy"`
	CircleOfficeFacilityId                 string         `json:"CircleOfficeFacilityId"`
	CircleOfficeName                       string         `json:"CircleOfficeName"`
	NotificationOrderNumber                string         `json:"NotificationOrderNumber"`
	OrderDate                              time.Time      `json:"OrderDate" binding:"required"`
	ExamYear                               int32          `json:"ExamYear" binding:"required"`
	Papers                                 *[]interface{} `json:"Papers" binding:"required"`
	EmployeeMasterRequestLastDate          time.Time      `json:"EmployeeMasterRequestLastDate"`
	EmployeeMasterRequestApprovalLastDate  time.Time      `json:"EmployeeMasterRequestApprovalLastDate"`
	ExamRegisterLastDate                   time.Time      `json:"ExamRegisterLastDate"`
	ApplicationStartDate                   time.Time      `json:"ApplicationStartDate"`
	ApplicationEndDate                     time.Time      `json:"ApplicationEndDate" binding:"required"`
	ApplicationVerificationLastDate        time.Time      `json:"ApplicationVerificationLastDate" binding:"required"`
	ApplicationCorrectionLastDate          time.Time      `json:"ApplicationCorrectionLastDate"`
	ResubmittedApplicationVerificationDate time.Time      `json:"ResubmittedApplicationVerificationDate" binding:"required"`
	NodalOfficerApprovalDate               time.Time      `json:"NodalOfficerApprovalDate" binding:"required"`
	CenterAllotmentEndDate                 time.Time      `json:"CenterAllotmentEndDate" binding:"required"`
	AdmitCardDate                          time.Time      `json:"AdmitCardDate" binding:"required"`
	CrucialDate                            *[]interface{} `json:"CrucialDate" binding:"required"`
	NotificationStatus                     string         `json:"NotificationStatus"`
	UpdatedBy                              string         `json:"UpdatedBy"`
	CreatedById                            int64          `json:"CreatedById"`
	CreatedBy                              string         `json:"CreatedBy"`
	CreatedByName                          string         `json:"CreatedByName"`
	CreatedByDesignation                   string         `json:"CreatedByDesignation"`
	Edges                                  EdgesLogdata   `json:"edges"`
}

type ResubitNotificationStruct struct {
	ExamName                               string         `json:"ExamName" binding:"required"`
	ExamCode                               int32          `json:"ExamCode" binding:"required"`
	UserName                               string         `json:"UserName" binding:"required"`
	IssuedBy                               string         `json:"IssuedBy"`
	CircleOfficeFacilityId                 string         `json:"CircleOfficeFacilityId"`
	CircleOfficeName                       string         `json:"CircleOfficeName"`
	NotificationOrderNumber                string         `json:"NotificationOrderNumber"`
	OrderDate                              time.Time      `json:"OrderDate" binding:"required"`
	ExamYear                               int32          `json:"ExamYear" binding:"required"`
	Id                                     int32          `json:"Id"`
	Papers                                 *[]interface{} `json:"Papers" binding:"required"`
	EmployeeMasterRequestLastDate          time.Time      `json:"EmployeeMasterRequestLastDate"`
	EmployeeMasterRequestApprovalLastDate  time.Time      `json:"EmployeeMasterRequestApprovalLastDate"`
	ExamRegisterLastDate                   time.Time      `json:"ExamRegisterLastDate"`
	ApplicationStartDate                   time.Time      `json:"ApplicationStartDate"`
	ApplicationEndDate                     time.Time      `json:"ApplicationEndDate" binding:"required"`
	ApplicationVerificationLastDate        time.Time      `json:"ApplicationVerificationLastDate" binding:"required"`
	ApplicationCorrectionLastDate          time.Time      `json:"ApplicationCorrectionLastDate"`
	ResubmittedApplicationVerificationDate time.Time      `json:"ResubmittedApplicationVerificationDate" binding:"required"`
	NodalOfficerApprovalDate               time.Time      `json:"NodalOfficerApprovalDate" binding:"required"`
	CenterAllotmentEndDate                 time.Time      `json:"CenterAllotmentEndDate" binding:"required"`
	AdmitCardDate                          time.Time      `json:"AdmitCardDate" binding:"required"`
	CrucialDate                            *[]interface{} `json:"CrucialDate" binding:"required"`
	NotificationStatus                     string         `json:"NotificationStatus"`
	NotificationRemarks                    string         `json:"NotificationRemarks"`
	UpdatedBy                              string         `json:"UpdatedBy"`
	CreatedById                            int64          `json:"CreatedById"`
	CreatedBy                              string         `json:"CreatedBy"`
	CreatedByName                          string         `json:"CreatedByName"`
	CreatedByDesignation                   string         `json:"CreatedByDesignation"`
	ExamShortName                          string         `json:"ExamShortName"`
	Edges                                  EdgesLogdata   `json:"edges"`
}

type UpdateNotificationStruct struct {
	ExamName                               string         `json:"ExamName" binding:"required"`
	ExamCode                               int32          `json:"ExamCode" binding:"required"`
	UserName                               string         `json:"UserName" binding:"required"`
	IssuedBy                               string         `json:"IssuedBy"`
	CircleOfficeFacilityId                 string         `json:"CircleOfficeFacilityId"`
	CircleOfficeName                       string         `json:"CircleOfficeName"`
	NotificationOrderNumber                string         `json:"NotificationOrderNumber"`
	OrderDate                              time.Time      `json:"OrderDate" binding:"required"`
	ExamYear                               int32          `json:"ExamYear" binding:"required"`
	Id                                     int32          `json:"Id"`
	Papers                                 *[]interface{} `json:"Papers" binding:"required"`
	EmployeeMasterRequestLastDate          time.Time      `json:"EmployeeMasterRequestLastDate"`
	EmployeeMasterRequestApprovalLastDate  time.Time      `json:"EmployeeMasterRequestApprovalLastDate"`
	ExamRegisterLastDate                   time.Time      `json:"ExamRegisterLastDate"`
	ApplicationStartDate                   time.Time      `json:"ApplicationStartDate"`
	ApplicationEndDate                     time.Time      `json:"ApplicationEndDate" binding:"required"`
	ApplicationVerificationLastDate        time.Time      `json:"ApplicationVerificationLastDate" binding:"required"`
	ApplicationCorrectionStartDate         time.Time      `json:"ApplicationCorrectionStartDate" binding:"required"`
	ApplicationCorrectionLastDate          time.Time      `json:"ApplicationCorrectionLastDate"`
	ResubmittedApplicationVerificationDate time.Time      `json:"ResubmittedApplicationVerificationDate" binding:"required"`
	NodalOfficerApprovalDate               time.Time      `json:"NodalOfficerApprovalDate" binding:"required"`
	CenterAllotmentEndDate                 time.Time      `json:"CenterAllotmentEndDate" binding:"required"`
	AdmitCardDate                          time.Time      `json:"AdmitCardDate" binding:"required"`
	CrucialDate                            *[]interface{} `json:"CrucialDate" binding:"required"`
	NotificationStatus                     string         `json:"NotificationStatus"`
	NotificationRemarks                    string         `json:"NotificationRemarks"`
	UpdatedBy                              string         `json:"UpdatedBy"`
	CreatedById                            int64          `json:"CreatedById"`
	CreatedBy                              string         `json:"CreatedBy"`
	CreatedByName                          string         `json:"CreatedByName"`
	CreatedByDesignation                   string         `json:"CreatedByDesignation"`
	ExamShortName                          string         `json:"ExamShortName"`
	Edges                                  EdgesLogdata   `json:"edges"`
}

type IssueNotificationStruct struct {
	ApprovedBy                             string         `json:"ApprovedBy"`
	ApprovedById                           int64          `json:"ApprovedById"`
	ApprovedByName                         string         `json:"ApprovedByName"`
	ApprovedByDesignation                  string         `json:"ApprovedByDesignation"`
	Id                                     int32          `json:"Id"`
	UserName                               string         `json:"UserName"`
	ExamYear                               int32          `json:"ExamYear"`
	ApplicationStartDate                   time.Time      `json:"ApplicationStartDate" binding:"required"`
	ApplicationEndDate                     time.Time      `json:"ApplicationEndDate" binding:"required"`
	ApplicationCorrectionLastDate          time.Time      `json:"ApplicationCorrectionLastDate" binding:"required"`
	ApplicationVerificationLastDate        time.Time      `json:"ApplicationVerificationLastDate" binding:"required"`
	CenterAllotmentEndDate                 time.Time      `json:"CenterAllotmentEndDate" binding:"required"`
	NodalOfficerApprovalDate               time.Time      `json:"NodalOfficerApprovalDate" binding:"required"`
	AdmitCardDate                          time.Time      `json:"AdmitCardDate" binding:"required"`
	CrucialDate                            *[]interface{} `json:"CrucialDate" binding:"required"`
	NotificationOrderNumber                string         `json:"NotificationOrderNumber"`
	ExamShortName                          string         `json:"ExamShortName"`
	CircleOfficeFacilityId                 string         `json:"CircleOfficeFacilityId"`
	CircleOfficeName                       string         `json:"CircleOfficeName"`
	IssuedBy                               string         `json:"IssuedBy"`
	OrderDate                              time.Time      `json:"OrderDate" binding:"required"`
	CreatedBy                              string         `json:"CreatedBy"`
	CreatedById                            int64          `json:"CreatedById"`
	CreatedByName                          string         `json:"CreatedByName"`
	CreatedByDesignation                   string         `json:"CreatedByDesignation"`
	ResubmittedApplicationVerificationDate time.Time      `json:"ResubmittedApplicationVerificationDate" binding:"required"`
	Papers                                 *[]interface{} `json:"Papers" binding:"required"`
	NotificationStatus                     string         `json:"NotificationStatus"`
	Status                                 string         `json:"Status"`
	ExamCode                               int32          `json:"ExamCode"`
	ExamName                               string         `json:"ExamName"`
	NotificationRemarks                    string         `json:"NotificationRemarks"`
	Edges                                  EdgesLogdata   `json:"edges"`
	NotificationNumber                     string         `json:"NotificationNumber"`
}
type CenterRequest struct {
	CityID                int32          `json:"CityID" binding:"required"`
	NodalOfficerCode      int32          `json:"NodalOfficerCode" binding:"required"`
	ExamCenterName        string         `json:"ExamCenterName" binding:"required"`
	ExamCode              int32          `json:"ExamCode" binding:"required"`
	RegionID              int32          `json:"RegionID" binding:"required"`
	CircleID              int32          `json:"CircleID" binding:"required"`
	DivisionID            int32          `json:"DivisionID" binding:"required"`
	FacilityID            int32          `json:"FacilityID" binding:"required"`
	ExamNameCode          string         `json:"ExamNameCode" binding:"required"`
	ExamName              string         `json:"ExamName" binding:"required"`
	NAUserName            string         `json:"NAUserName"`
	NodalOfficeFacilityId string         `json:"NodalOfficeFacilityId"`
	AdminCircleOfficeID   string         `json:"AdminCircleOfficeID"`
	Address               string         `json:"Address" binding:"required"`
	Landmark              string         `json:"Landmark"`
	CenterCityName        string         `json:"CenterCityName" binding:"required"`
	Pincode               int32          `json:"Pincode" binding:"required"`
	MaxSeats              int32          `json:"MaxSeats"`
	NoAlloted             int32          `json:"NoAlloted"`
	PendingSeats          int32          `json:"PendingSeats"`
	Status                bool           `json:"Status"`
	ExamYear              string         `json:"ExamYear"`
	ConductedBy           string         `json:"ConductedBy"`
	ConductedByFacilityID string         `json:"ConductedByFacilityID"`
	UpdatedAt             time.Time      `json:"UpdatedAt"`
	UpdatedBy             string         `json:"UpdatedBy"`
	Papers                *[]interface{} `json:"Papers"` // Assuming Papers is a list of strings
	Edges                 EdgesLogdata   `json:"edges"`
}

type CenterReq struct {
	ExamCode              int32        `json:"ExamCode"`
	ExamCenterName        string       `json:"ExamCenterName" binding:"required"`
	Address               string       `json:"Address"`
	Landmark              string       `json:"Landmark"`
	Pincode               int32        `json:"Pincode"`
	MaxSeats              int32        `json:"MaxSeats" binding:"required"`
	NOUserName            string       `json:"NOUserName" binding:"required"`
	CityID                int32        `json:"CityID"`
	Status                bool         `json:"Status" binding:"required"`
	NodalOfficeFacilityID string       `json:"NodalOfficeFacilityID" binding:"required"`
	CenterCityName        string       `json:"CenterCityName"`
	ConductedBy           string       `json:"ConductedBy"`
	ConductedByFacilityID string       `json:"ConductedByFacilityID"`
	NAUserName            string       `json:"NAUserName"`
	ExamYear              string       `json:"ExamYear"`
	Edges                 EdgesLogdata `json:"edges"`
}
type StrucProfileEmployeeMaster struct {
	UserName                       string       `json:"UserName" binding:"required"`
	EmployeeID                     int64        `json:"EmployeeID" binding:"required"`
	EmployeeName                   string       `json:"EmployeeName" binding:"required"`
	DOB                            string       `json:"DOB" ` // Assuming DOB is in string format, change to time.Time if it's a date
	Gender                         string       `json:"Gender"`
	MobileNumber                   string       `json:"MobileNumber" binding:"required"`
	EmailID                        string       `json:"EmailID" binding:"required"`
	EmployeeCategory               string       `json:"EmployeeCategory"`
	EmployeePost                   string       `json:"EmployeePost"`
	FacilityID                     string       `json:"FacilityID"`
	Pincode                        string       `json:"Pincode"`
	OfficeName                     string       `json:"OfficeName"`
	ControllingAuthorityFacilityID string       `json:"ControllingAuthorityFacilityID"`
	ControllingAuthorityName       string       `json:"ControllingAuthorityName"`
	NodalAuthorityFacilityID       string       `json:"NodalAuthorityFaciliyId"`
	NodalAuthorityName             string       `json:"NodalAuthorityName"`
	CircleFacilityID               string       `json:"CircleFacilityID"`
	ModifiedByID                   int64        `json:"ModifiedByID"`
	ModifiedByUserName             string       `json:"ModifiedByUserName"`
	ModifiedByEmpID                int64        `json:"ModifiedByEmpID"`
	Edges                          EdgesLogdata `json:"edges"`
}
type CircleOfficeDetailss struct {
	NodalOfficeFacilityID         string `json:"nodalOffice_facility_id"`
	NodalOfficeName               string `json:"nodal_office_name"`
	FacilityID                    string `json:"facility_id"`
	FacilityName                  string `json:"facility_name"`
	RecommendedCount              int    `json:"recommended_count"`
	NotRecommendedCount           int    `json:"not_recommended_count"`
	ProvisionallyRecommendedCount int    `json:"provisionally_recommended_count"`
}

type ApplicationStatusDetailss struct {
	NodalOfficeFacilityID string `json:"nodalOffice_facility_id"`
	NodalOfficeName       string `json:"nodal_office_name"`
	FacilityID            string `json:"facility_id"`
	FacilityName          string `json:"facility_name"`
	PendingWithCandidate  int    `json:"pending_with_candidate"`
	PendingWithCA         int    `json:"pending_with_ca"`
}

type StrucSericeRequest struct {
	ID         int64  `json:"id" binding:"required"`
	UpdatedBy  string `json:"updated_by" binding:"required"`
	AssignedTo string `json:"assigned_to"`
	RemarksNew string `json:"remarks_new"`
	Status     string `json:"status"`
}
type ApplicationsResponse struct {
	EmployeeID            int64     `json:"EmployeeID,omitempty"`
	ApplicationNumber     string    `json:"ApplicationNumber,omitempty"`
	ApplicationStatus     string    `json:"ApplicationStatus,omitempty"`
	MobileNumber          string    `json:"MobileNumber,omitempty"`
	EmailID               string    `json:"EmailID,omitempty"`
	RoleUserCode          int32     `json:"RoleUserCode,omitempty"`
	ApplnSubmittedDate    time.Time `json:"ApplnSubmittedDate,omitempty"`
	ControllingOfficeName string    `json:"ControllingOfficeName,omitempty"`
	ApplicationRemarks    string    `json:"ApplicationRemarks,omitempty"`
	RecommendedStatus     string    `json:"RecommendedStatus,omitempty"`
	CAGeneralRemarks      string    `json:"CAGeneralRemarks,omitempty"`
	NARemarks             string    `json:"NARemarks,omitempty"`
}
type AdminMasterResponse struct {
	ID                          int64     `json:"id,omitempty"`
	EmployeeId                  int64     `json:"EmployeeId,omitempty"`
	EmployeeName                string    `json:"EmployeeName,omitempty"`
	Designation                 string    `json:"Designation,omitempty"`
	RoleUserCode                int32     `json:"RoleUserCode,omitempty"`
	RoleUserDescription         string    `json:"RoleUserDescription,omitempty"`
	Mobile                      string    `json:"Mobile,omitempty"`
	EmailID                     string    `json:"EmailId,omitempty"`
	UserName                    string    `json:"UserName,omitempty"`
	FacilityIDUniqueID          int64     `json:"FacilityIdUniqueId,omitempty"`
	FacilityID                  string    `json:"FacilityId,omitempty"`
	AuthorityFacilityName       string    `json:"AuthorityFacilityName,omitempty"`
	FacilityType                string    `json:"FacilityType,omitempty"`
	ReportingOfficeFacilityID   string    `json:"ReportingOfficeFacilityId,omitempty"`
	ReportingOfficeFacilityName string    `json:"ReportingOfficeFacilityName,omitempty"`
	CircleOfficeFacilityID      string    `json:"CircleOfficeFacilityId,omitempty"`
	CircleOfficeName            string    `json:"CircleOfficeName,omitempty"`
	Status                      string    `json:"Status,omitempty"`
	CreatedByID                 int64     `json:"CreatedById,omitempty"`
	CreatedByUserName           string    `json:"CreatedByUserName,omitempty"`
	CreatedByEmpID              int64     `json:"CreatedByEmpId,omitempty"`
	CreatedByDesignation        string    `json:"CreatedByDesignation,omitempty"`
	CreatedDate                 time.Time `json:"CreatedDate,omitempty"`
	ModifiedByID                int64     `json:"ModifiedById,omitempty"`
	ModifiedByUserName          string    `json:"ModifiedByUserName,omitempty"`
	ModifiedByEmpID             int64     `json:"ModifiedByEmpId,omitempty"`
	ModifiedByDesignation       string    `json:"ModifiedByDesignation,omitempty"`
	ModifiedDate                time.Time `json:"ModifiedDate,omitempty"`
	DeletedByID                 int64     `json:"DeletedById,omitempty"`
	DeletedByUserName           string    `json:"DeletedByUserName,omitempty"`
	DeletedByEmpID              int64     `json:"DeletedByEmpId,omitempty"`
	DeletedByDesignation        string    `json:"DeletedByDesignation,omitempty"`
	DeletedDate                 time.Time `json:"DeletedDate,omitempty"`
	NewPasswordRequest          bool      `json:"NewPasswordRequest,omitempty"`
	EventTime                   time.Time `json:"EventTime,omitempty"`
	UUID                        int64     `json:"Uuid,omitempty"`
	UpdatedBy                   string    `json:"UpdatedBy,omitempty"`
	UpdatedDate                 time.Time `json:"UpdatedDate,omitempty"`
}
type ExamCityCenterRequest struct {
	ExamCode              int32        `json:"ExamCode" binding:"required"`
	ExamName              string       `json:"ExamName" binding:"required"`
	ExamShortName         string       `json:"ExamShortName" binding:"required"`
	ExamYear              int32        `json:"ExamYear"  binding:"required"`
	ConductedBy           string       `json:"ConductedBy"  binding:"required"`
	NodalOfficeFacilityID string       `json:"NodalOfficeFacilityID"  binding:"required"`
	NodalOfficeName       string       `json:"NodalOfficeName"`
	NotificationCode      int32        `json:"NotificationCode"`
	NotificationNumber    string       `json:"NotificationNumber"  binding:"required"`
	CenterCityName        string       `json:"CenterCityName" binding:"required"`
	CreatedById           int64        `json:"CreatedById"  binding:"required"`
	CreatedByUserName     string       `json:"CreatedByUserName"  binding:"required"`
	CreatedByEmpId        int64        `json:"CreatedByEmpId"  binding:"required"`
	CreatedByDesignation  string       `json:"CreatedByDesignation"  binding:"required"`
	CreatedDate           time.Time    `json:"CreatedDate"  binding:"required"`
	Status                string       `json:"Status"`
	DeletedById           int64        `json:"DeletedById,omitempty"`
	DeletedByUserName     string       `json:"DeletedByUserName,omitempty"`
	DeletedByEmployeeId   int64        `json:"DeletedByEmployeeId,omitempty"`
	DeletedByDesignation  string       `json:"DeletedByDesignation,omitempty"`
	DeletedDate           time.Time    `json:"DeletedDate,omitempty"`
	CircleCityName        string       `json:"CircleCityName"`
	DivisionCode          int32        `json:"DivisionCode"`
	RegionCode            int32        `json:"RegionCode"`
	DivisionName          string       `json:"DivisionName"`
	RegionID              int32        `json:"RegionID"`
	RegionName            string       `json:"RegionName"`
	RegionCityName        string       `json:"RegionCityName"`
	CentreCityName        string       `json:"CentreCityName"`
	Remarks               string       `json:"Remarks"`
	UpdatedAt             time.Time    `json:"UpdatedAt"`
	UpdatedBy             string       `json:"UpdatedBy"`
	CentreCode            int32        `json:"CentreCode"`
	CircleID              int32        `json:"CircleID"`
	Edges                 EdgesLogdata `json:"edges" binding:"required"`
}
type ExamCityCenterResponse struct {
	ID                    int32     `json:"ID,omitempty"`
	ExamCode              int32     `json:"ExamCode,omitempty"`
	ExamName              string    `json:"ExamName,omitempty"`
	ExamShortName         string    `json:"ExamShortName,omitempty"`
	ExamYear              int32     `json:"ExamYear,omitempty"`
	ConductedBy           string    `json:"ConductedBy,omitempty"`
	NodalOfficeFacilityID string    `json:"NodalOfficeFacilityID,omitempty"`
	NodalOfficeName       string    `json:"NodalOfficeName,omitempty"`
	NotificationCode      int32     `json:"NotificationCode,omitempty"`
	NotificationNumber    string    `json:"NotificationNumber,omitempty"`
	CenterCityName        string    `json:"CenterCityName,omitempty"`
	CreatedById           int64     `json:"CreatedById,omitempty"`
	CreatedByUserName     string    `json:"CreatedByUserName,omitempty"`
	CreatedByEmpId        int64     `json:"CreatedByEmpId,omitempty"`
	CreatedByDesignation  string    `json:"CreatedByDesignation,omitempty"`
	CreatedDate           time.Time `json:"CreatedDate,omitempty"`
	Status                string    `json:"Status,omitempty"`
	DeletedById           int64     `json:"DeletedById,omitempty"`
	DeletedByUserName     string    `json:"DeletedByUserName,omitempty"`
	DeletedByEmployeeId   int64     `json:"DeletedByEmployeeId,omitempty"`
	DeletedByDesignation  string    `json:"DeletedByDesignation,omitempty"`
	DeletedDate           time.Time `json:"DeletedDate,omitempty"`
	CircleCityName        string    `json:"CircleCityName,omitempty"`
	DivisionCode          int32     `json:"DivisionCode,omitempty"`
	RegionCode            int32     `json:"RegionCode,omitempty"`
	DivisionName          string    `json:"DivisionName,omitempty"`
	RegionID              int32     `json:"RegionID,omitempty"`
	RegionName            string    `json:"RegionName,omitempty"`
	RegionCityName        string    `json:"RegionCityName,omitempty"`
	CentreCityName        string    `json:"CentreCityName,omitempty"`
	Remarks               string    `json:"Remarks,omitempty"`
	UpdatedAt             time.Time `json:"UpdatedAt,omitempty"`
	UpdatedBy             string    `json:"UpdatedBy,omitempty"`
	CentreCode            int32     `json:"CentreCode,omitempty"`
	CircleID              int32     `json:"CircleID,omitempty"`
}
