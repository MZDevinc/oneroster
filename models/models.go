package models

// ORProcess type for platform-dependent functionality
type ORProcess interface {

	// AcademicSessions
	HandleAddAcademicSessions(orAcademicSessions []ORAcademicSessions) error
	HandleDeleteAcademicSessions(orAcademicSessionsIDs []string) error
	HandleEditAcademicSessions(orAcademicSessions []ORAcademicSessions) error
	HandleAddOrEditAcademicSessions(orAcademicSessions []ORAcademicSessions) error
	HandleEstimateNewAcademicSessions(orClass []ORClass, districtIDs []string) (int, error)

	// Users
	HandleAddUsers(orUser []ORUser, districtIDs []string) error
	HandleDeleteUsers(oruserIDs []string, districtIDs []string) error
	HandleEditUsers(orUser []ORUser) error
	HandleAddOrEditUsers(orUser []ORUser, districtIDs []string) error
	HandleEstimateNewUsers(orUsers []ORUser, districtIDs []string) (int, error)

	// Districts
	HandleAddDistrict(orOrg OROrg) (bool, error)
	HandleDeleteDistrict(orOrg OROrg) error
	HandleEditDistrict(orOrg OROrg, districtID string) error
	HandleAddOrEditDistrict(orOrg OROrg) error
	HandleEstimateNewDistricts(orOrgs []OROrg, districtIDs []string) (int, error)

	// Schools
	HandleAddSchool(orOrg OROrg, districtIDs []string) error
	HandleDeleteSchool(orOrg OROrg, districtIDs []string) error
	HandleEditSchool(orOrg OROrg) error
	HandleAddOrEditSchool(orOrg OROrg, districtIDs []string) error
	HandleEstimateNewSchools(orOrgs []OROrg, districtIDs []string) (int, error)

	// Classes
	HandleAddClasses(orClass []ORClass, districtIDs []string) error
	HandleDeleteClasses(orClassIDs []string, districtIDs []string) error
	HandleEditClass(orClass []ORClass) error
	HandleAddOrEditClass(orClass []ORClass, districtIDs []string) error
	HandleEstimateNewClasses(orClasses []ORClass, districtIDs []string) (int, error)

	// Courses
	HandleAddCourses(orCourse []ORCourse, districtIDs []string) error
	HandleDeleteCourses(orCourseIDs []string, districtIDs []string) error
	HandleEditCourse(orCourse []ORCourse) error
	HandleAddOrEditCourse(orCourse []ORCourse, districtIDs []string) error
	HandleEstimateNewCourses(orCourses []ORCourse, districtIDs []string) (int, error)

	// Enrollments
	HandleAddEnrollment(orEnrollment []OREnrollment, districtIDs []string) error
	HandleDeleteEnrollments(orEnrollment []OREnrollment, districtIDs []string) error
	HandleAddOrEditEnrollments(orEnrollment []OREnrollment, districtIDs []string) error
	HandleEstimateNewEnrollments(orEnrollments []OREnrollment, districtIDs []string) (int, error)

	RollBackOneRoster(orgDistrict []OROrg) error

	GetDistrictsIDs(orOrgs []OROrg) ([]string, error)
}

// OrManifest manifest file for oneroster
type OrManifest struct {
	PropertyName string `csv:"propertyName"`
	Value        string `csv:"value"`
}

// ORAcademicSessions academic sessions for oneroster
type ORAcademicSessions struct {
	SourcedID        string    `csv:"sourcedId" json:"sourcedId"`               //GUID
	Status           string    `csv:"status" json:"status"`                     //Enumeration
	DateLastModified string    `csv:"dateLastModified" json:"dateLastModified"` //DateTime
	Title            string    `csv:"title" json:"title"`
	SessionType      string    `csv:"type" json:"type"`           //Enumeration
	StartDate        string    `csv:"startDate" json:"startDate"` //date
	EndDate          string    `csv:"endDate" json:"endDate"`     //date
	ParentSourcedID  string    `csv:"parentSourcedId"`            //GUID Reference
	Parent           GUIDRef   `json:"parent"`
	Children         []GUIDRef `json:"children"`
	SchoolYear       string    `csv:"schoolYear" json:"schoolYear"` //year
}

// ORClass classes for oneroster
type ORClass struct {
	SourcedID        string    `csv:"sourcedId" json:"sourcedId"`               //GUID
	Status           string    `csv:"status" json:"status"`                     //Enumeration
	DateLastModified string    `csv:"dateLastModified" json:"dateLastModified"` //DateTime
	Title            string    `csv:"title" json:"title"`
	Grades           string    `csv:"grades" json:""`  //[]string
	CourseSourcedID  string    `csv:"courseSourcedId"` //GUID Reference
	Course           GUIDRef   `json:"course"`
	ClassCode        string    `csv:"classCode" json:"classCode"`
	ClassType        string    `csv:"classType" json:"classType"` //Enumeration
	Location         string    `csv:"location" json:"location"`
	SchoolSourcedID  string    `csv:"schoolSourcedId" ` //GUID Reference
	School           GUIDRef   `json:"school"`
	TermSourcedIds   string    `csv:"termSourcedIds" ` //List of GUID Reference
	Terms            []GUIDRef `json:"terms"`
	Subjects         string    `csv:"subjects" json:"subjects"`
	SubjectCodes     string    `csv:"subjectCodes" json:"subjectCodes"`
	Periods          string    `csv:"periods" json:"periods"`
	Resources        []GUIDRef `json:"resources"`
}

// ORCourse courses for oneroster
type ORCourse struct {
	SourcedID           string  `csv:"sourcedId" json:"sourcedId"`               //GUID
	Status              string  `csv:"status" json:"status"`                     //Enumeration
	DateLastModified    string  `csv:"dateLastModified" json:"dateLastModified"` //DateTime
	SchoolYearSourcedID string  `csv:"schoolYearSourcedId"`                      //GUID Reference
	SchoolYear          GUIDRef `json:"schoolYear"`
	Title               string  `csv:"title" json:"title"`
	CourseCode          string  `csv:"courseCode" json:"courseCode"`
	// Grades				*[]string	`csv:"grades"`
	OrgSourcedID string  `csv:"orgSourcedId"` //GUID Reference
	Org          GUIDRef `json:"org"`
	Subjects     string  `csv:"subjects" json:"subjects"`
	SubjectCodes string  `csv:"subjectCodes" json:"subjectCodes"`
}

// ORDemographics demographics for one roster
type ORDemographics struct {
	SourcedID                            string `csv:"sourcedId" json:"sourcedId"`                                                       //GUID
	Status                               string `csv:"status" json:"status"`                                                             //Enumeration
	DateLastModified                     string `csv:"dateLastModified" json:"dateLastModified"`                                         //DateTime
	BirthDate                            string `csv:"birthDate" json:"birthDate"`                                                       //date
	Sex                                  string `csv:"sex" json:"sex"`                                                                   //Enumeration
	AmericanIndianOrAlaskaNative         string `csv:"americanIndianOrAlaskaNative" json:"americanIndianOrAlaskaNative"`                 //Enumeration
	Asian                                string `csv:"asian" json:"asian"`                                                               //Enumeration
	BlackOrAfricanAmerican               string `csv:"blackOrAfricanAmerican" json:"blackOrAfricanAmerican"`                             //Enumeration
	NativeHawaiianOrOtherPacificIslander string `csv:"nativeHawaiianOrOtherPacificIslander" json:"nativeHawaiianOrOtherPacificIslander"` //Enumeration
	White                                string `csv:"white" json:"white"`                                                               //Enumeration
	DemographicRaceTwoOrMoreRaces        string `csv:"demographicRaceTwoOrMoreRaces" json:"demographicRaceTwoOrMoreRaces"`               //Enumeration
	HispanicOrLatinoEthnicity            string `csv:"hispanicOrLatinoEthnicity" json:"hispanicOrLatinoEthnicity"`                       //Enumeration
	CountryOfBirthCode                   string `csv:"countryOfBirthCode" json:"countryOfBirthCode"`
	StateOfBirthAbbreviation             string `csv:"stateOfBirthAbbreviation" json:"stateOfBirthAbbreviation"`
	CityOfBirth                          string `csv:"cityOfBirth" json:"cityOfBirth"`
	PublicSchoolResidenceStatus          string `csv:"publicSchoolResidenceStatus" json:"publicSchoolResidenceStatus"`
}

// OREnrollment enrollments for oneroster
type OREnrollment struct {
	SourcedID        string  `csv:"sourcedId" json:"sourcedId"`               //GUID
	Status           string  `csv:"status" json:"status"`                     //Enumeration
	DateLastModified string  `csv:"dateLastModified" json:"dateLastModified"` //DateTime
	ClassSourcedID   string  `csv:"classSourcedId"`                           //GUID Reference
	Class            GUIDRef `json:"class"`
	SchoolSourcedID  string  `csv:"schoolSourcedId"` //GUID Reference
	School           GUIDRef `json:"school"`
	UserSourcedID    string  `csv:"userSourcedId"` //GUID Reference
	User             GUIDRef `json:"user"`
	Role             string  `csv:"role" json:"role"` //Enumeration
	Primary          bool    `csv:"primary" json:"primary"`
	BeginDate        string  `csv:"beginDate" json:"beginDate"` //date
	EndDate          string  `csv:"endDate" json:"endDate"`     //date
}

// OROrg orgs for oneroster
type OROrg struct {
	SourcedID        string    `csv:"sourcedId" json:"sourcedId"`               //GUID
	Status           string    `csv:"status" json:"status"`                     //Enumeration
	DateLastModified string    `csv:"dateLastModified" json:"dateLastModified"` //DateTime
	Name             string    `csv:"name" json:"name"`
	OrgType          string    `csv:"type" json:"type"` // type Enumeration
	Identifier       string    `csv:"identifier" json:"identifier"`
	ParentSourcedID  string    `csv:"parentSourcedId"` //GUID Reference
	Parent           GUIDRef   `json:"parent"`
	Children         []GUIDRef `json:"children"`
}

// ORUser users for oneroster
type ORUser struct {
	SourcedID        string          `csv:"sourcedId" json:"sourcedId"`               //GUID
	Status           string          `csv:"status" json:"status"`                     //Enumeration
	DateLastModified string          `csv:"dateLastModified" json:"dateLastModified"` //DateTime
	EnabledUser      bool            `csv:"enabledUser" json:"enabledUser"`
	OrgSourcedIds    string          `csv:"orgSourcedIds"` //List of GUID References.
	Orgs             []GUIDRef       `json:"orgs"`
	Role             string          `csv:"role" json:"role"` //Enumeration
	Username         string          `csv:"username" json:"username"`
	UserIds          string          `csv:"userIds"` //[] string
	UserIdsIdentifer []UserIdentifer `json:"userIds"`
	GivenName        string          `csv:"givenName" json:"givenName"`
	FamilyName       string          `csv:"familyName" json:"familyName"`
	MiddleName       string          `csv:"middleName" json:"middleName"`
	Identifier       string          `csv:"identifier" json:"identifier"`
	Email            string          `csv:"email" json:"email"`
	Sms              string          `csv:"sms" json:"sms"`
	Phone            string          `csv:"phone" json:"phone"`
	AgentSourcedIds  string          `csv:"agentSourcedIds"` //List of GUID References
	Agents           []GUIDRef       `json:"agents"`
	Grades           string          `csv:"grades" json:"grades"`
	Password         string          `csv:"password" json:"password"`
}

// ORCategory categories for oneroster
type ORCategory struct {
	SourcedID        string //GUID
	Status           string //Enumeration
	DateLastModified string //DateTime
	Title            string
}

// ORClassResources class resources for oneroster
type ORClassResources struct {
	SourcedID         string //GUID
	Status            string //Enumeration
	DateLastModified  string //DateTime
	Title             string
	ClassSourcedID    string //GUID Reference
	ResourceSourcedID string //GUID Reference
}

// ORCourseResources course resources for oneroster
type ORCourseResources struct {
	SourcedID         string //GUID
	Status            string //Enumeration
	DateLastModified  string //DateTime
	Title             string
	CourseSourcedID   string //GUID Reference
	ResourceSourcedID string //GUID Reference
}

// ORResource resource for oneroster
type ORResource struct {
	SourcedID        string //GUID
	Status           string //Enumeration
	DateLastModified string //DateTime
	VendorResourceID string //id
	Title            string
	Roles            []string //Enumeration List
	Importance       string
	VendorID         string //id
	ApplicationID    string //id
}

// ORResult results for oneroster
type ORResult struct {
	SourcedID         string  //GUID
	Status            string  //Enumeration
	DateLastModified  string  //DateTime
	LineItemSourcedID string  //GUID Reference
	StudentSourcedID  string  //GUID Reference
	ScoreStatus       string  //Enumeration
	Score             float64 //float
	ScoreDate         string  //date
	Comment           string
}

// ORLineItems line items for oneroster
type ORLineItems struct {
	SourcedID              string //GUID
	Status                 string //Enumeration
	DateLastModified       string //DateTime
	Title                  string
	Description            string
	SssignDate             string //date
	DueDate                string //date
	ClassSourcedID         string // GUID References
	CategorySourcedID      string // GUID References
	GradingPeriodSourcedID string // GUID References
	ResultValueMin         float64
	ResultValueMax         float64
}

// import type
const (
	ImportTypeBulk   = "bulk"
	ImportTypeDelta  = "delta"
	ImportTypeAbsent = "absent"
	// Custom extension
	ImportTypeNewEstimate = "new_estimate"
)

// manifest property names
const (
	ManifestProVersion              = "manifest.version"
	ManifestProOnerosterVersion     = "oneroster.version"
	ManifestProFileAcademicSessions = "file.academicSessions"
	ManifestProFileCategories       = "file.categories"
	ManifestProFileClasses          = "file.classes"
	ManifestProFileClassResources   = "file.classResources"
	ManifestProFileCourses          = "file.courses"
	ManifestProFileCourseResources  = "file.courseResources"
	ManifestProFileDemographics     = "file.demographics"
	ManifestProFileEnrollments      = "file.enrollments"
	ManifestProFileLineItems        = "file.lineItems"
	ManifestProFileOrgs             = "file.orgs"
	ManifestProFileResources        = "file.resources"
	ManifestProFileResults          = "file.results"
	ManifestProFileUsers            = "file.users"
	ManifestProSourceSystemName     = "source.systemName"
	ManifestProSourceSystemCode     = "source.systemCode"
)

// csv files name
const (
	CsvNameManifest         = "manifest.csv"
	CsvNameAcademicSessions = "academicSessions.csv"
	CsvNameCategories       = "categories.csv"
	CsvNameClasses          = "classes.csv"
	CsvNameCourses          = "courses.csv"
	CsvNameClassResources   = "classResources.csv"
	CsvNameDemographics     = "demographics.csv"
	CsvNameEnrollments      = "enrollments.csv"
	CsvNameOrgs             = "orgs.csv"
	CsvNameResources        = "resources.csv"
	CsvNameLineItems        = "lineItems.csv"
	CsvNameResults          = "results.csv"
	CsvNameUsers            = "users.csv"
)

// orgs types
const (
	OrgTypeDistrict = "district"
	OrgTypeSchool   = "school"
)

//Status types
const (
	StatusTypeActive      = "Active"
	StatusTypeToBeDeleted = "ToBeDeleted"
)

////// JSON ///////

// GUIDRef reference to GUID
type GUIDRef struct {
	Href      string `json:"href"`
	SourcedID string `json:"sourcedId"`
	GUIDType  string `json:"type"`
}

// UserIdentifer holds identity for user
type UserIdentifer struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

// GUID constants
const (
	GUIDTypeAcademicSession = "academicSession"
	GUIDTypeCategory        = "category"
	GUIDTypeClass           = "class"
	GUIDTypeCourse          = "course"
	GUIDTypeDemographics    = "demographics"
	GUIDTypeEnrollment      = "enrollment"
	GUIDTypeOrg             = "org"
	GUIDTypeResource        = "resource"
	GUIDTypeLineItem        = "lineItem"
	GUIDTypeResult          = "result"
	GUIDTypeUser            = "user"
	GUIDTypeStudent         = "student"
	GUIDTypeTeacher         = "teacher"
	GUIDTypeTerm            = "term"
	GUIDTypeGradingPeriod   = "gradingPeriod"
)

//// Rest API Responses ////

// OrgsResponse for org responses
type OrgsResponse struct {
	Orgs []OROrg `json:"orgs"`
}

// AcademicSessionsResponse for academic session responses
type AcademicSessionsResponse struct {
	AcademicSessions []ORAcademicSessions `json:"academicSessions"`
}

// ClassesResponse for class responses
type ClassesResponse struct {
	Classes []ORClass `json:"classes"`
}

// CoursesResponse for course responses
type CoursesResponse struct {
	Courses []ORCourse `json:"courses"`
}

// EnrollmentsResponse for enrollment responses
type EnrollmentsResponse struct {
	Enrollments []OREnrollment `json:"enrollments"`
}

// UsersResponse for user responses
type UsersResponse struct {
	Users []ORUser `json:"users"`
}
