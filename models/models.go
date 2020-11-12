package models



type ORProcess interface {

	HandleAddAcademicSessions(orAcademicSessions []ORAcademicSessions) error

	HandleAddUser(orUser []ORUser) error
	// HandleReomveuser() string
	// HandleEditUser() string

	HandleAddDistrict(orOrg OROrg) (bool, error)
	// HandleReomveDistrict() string
	// HandleEditDistrict() string

	HandleAddSchool(orOrg OROrg) error
	// HandleReomveSchool() string
	// HandleEditSchool() string

	HandleAddClasses(orClass []ORClass) error
	// HandleReomveClass() string
	// HandleEditClass() string

	HandleAddCourses(orCourse []ORCourse) error
	// HandleReomveCourse() string
	// HandleEditCourse() string

	HandleAddEnrollment(orEnrollment []OREnrollment) error

	RollBackOneRoster(orgDistrict []OROrg) error

}


type OrManifest struct {
	PropertyName			string `csv:"propertyName"`
	Value					string `csv:"value"`
}

type ORAcademicSessions struct{
	SourcedId 			string `csv:"sourcedId"`//GUID
	Status 				string `csv:"status"`//Enumeration
	DateLastModified 	string `csv:"dateLastModified"`//DateTime
	Title 				string `csv:"title"`
	SessionType 		string `csv:"type"`//Enumeration
	StartDate 			string `csv:"startDate"`//date
	EndDate 			string `csv:"endDate"`//date
	ParentSourcedId 	string `csv:"parentSourcedId"`//GUID Reference
	SchoolYear 			string `csv:"schoolYear"`//year
}

type ORClass struct{
	SourcedId 			string 			`csv:"sourcedId"`	//GUID
	Status 				string 			`csv:"status"`//Enumeration
	DateLastModified 	string 			`csv:"dateLastModified"`//DateTime
	Title 				string 			`csv:"title"`
	Grades 				string 			`csv:"grades"` //[]string
	CourseSourcedId 	string 			`csv:"courseSourcedId"`//GUID Reference
	ClassCode 			string			`csv:"classCode"`
	ClassType 			string 			`csv:"classType"`//Enumeration
	Location 			string 			`csv:"location"`
	SchoolSourcedId 	string 			`csv:"schoolSourcedId"`//GUID Reference
	TermSourcedIds 		string 		`csv:"termSourcedIds"`//List of GUID Reference
	Subjects 			string 		`csv:"subjects"`
	SubjectCodes 		string		`csv:"subjectCodes"`
	Periods 			string		`csv:"periods"`
}


type ORCourse struct {
	SourcedId 			string  	`csv:"sourcedId"`		//GUID
	Status				string 		`csv:"status"`		//Enumeration
	DateLastModified 	string 		`csv:"dateLastModified"`//DateTime
	SchoolYearSourcedId string 		`csv:"schoolYearSourcedId"`//GUID Reference
	Title 				string 		`csv:"title"`
	CourseCode			string		`csv:"courseCode"`
	// Grades				*[]string	`csv:"grades"`
	OrgSourcedId		string 		`csv:"orgSourcedId"`//GUID Reference
	// Subjects			*[]string	`csv:"subjects"`
	SubjectCodes		string		`csv:"subjectCodes"`
}

type ORDemographics struct{
	SourcedId 			string 						`csv:"sourcedId"`//GUID
	Status				string 						`csv:"status"`//Enumeration
	DateLastModified 	string 						`csv:"dateLastModified"`//DateTime
	BirthDate			string 						`csv:"birthDate"`//date
	Sex					string 						`csv:"sex"`//Enumeration
	AmericanIndianOrAlaskaNative 			string 	`csv:"americanIndianOrAlaskaNative"`//Enumeration
	Asian									string 	`csv:"asian"`//Enumeration
	BlackOrAfricanAmerican					string 	`csv:"blackOrAfricanAmerican"`//Enumeration
	NativeHawaiianOrOtherPacificIslander 	string 	`csv:"nativeHawaiianOrOtherPacificIslander"`//Enumeration
	White									string 	`csv:"white"`//Enumeration
	DemographicRaceTwoOrMoreRaces			string 	`csv:"demographicRaceTwoOrMoreRaces"`//Enumeration
	HispanicOrLatinoEthnicity				string 	`csv:"hispanicOrLatinoEthnicity"`//Enumeration
	CountryOfBirthCode						string	`csv:"countryOfBirthCode"`
	StateOfBirthAbbreviation				string	`csv:"stateOfBirthAbbreviation"`
	CityOfBirth								string	`csv:"cityOfBirth"`
	PublicSchoolResidenceStatus 			string	`csv:"publicSchoolResidenceStatus"`
}

type OREnrollment struct{
	SourcedId 			string			`csv:"sourcedId"` //GUID
	Status				string 			`csv:"status"`//Enumeration
	DateLastModified 	string 			`csv:"dateLastModified"`//DateTime
	ClassSourcedId		string 			`csv:"classSourcedId"`//GUID Reference
	SchoolSourcedId		string 			`csv:"schoolSourcedId"`//GUID Reference
	UserSourcedId 		string 			`csv:"userSourcedId"`//GUID Reference
	Role				string 			`csv:"role"`//Enumeration
	Primary				bool			`csv:"primary"`
	BeginDate			string 			`csv:"beginDate"`//date
	EndDate				string 			`csv:"endDate"`//date
}

type OROrg struct{
	SourcedId 			string `csv:"sourcedId"`	//GUID
	Status 				string `csv:"status"`//Enumeration
	DateLastModified 	string `csv:"dateLastModified"`//DateTime
	Name 				string `csv:"name"`
	OrgType 			string `csv:"type"`// type Enumeration
	Identifier 			string `csv:"identifier"`
	ParentSourcedId 	string `csv:"parentSourcedId"`	//GUID Reference
}


type ORUser struct {
	SourcedId 			string 	`csv:"sourcedId"`	//GUID
	Status 				string 	`csv:"status"`		//Enumeration
	DateLastModified 	string 	`csv:"dateLastModified"`//DateTime
	EnabledUser 		bool	`csv:"enabledUser"`
	OrgSourcedIds 		string 	`csv:"orgSourcedIds"`//List of GUID References.
	Role 				string 	`csv:"role"`		//Enumeration
	Username 			string	`csv:"username"`
	UserIds 			string  `csv:"userIds"` //[] string
	GivenName 			string	`csv:"givenName"`
	FamilyName 			string	`csv:"familyName"`
	MiddleName 			string	`csv:"middleName"`
	Identifier 			string 	`csv:"identifier"`
	Email 				string	`csv:"email"`
	Sms 				string	`csv:"sms"`
	Phone 				string	`csv:"phone"`
	AgentSourcedIds 	string `csv:"agentSourcedIds"`//List of GUID References
	Grades 				string 	`csv:"grades"`
	Password 			string	`csv:"password"`

}

type ORCategory struct{
	SourcedId 			string 	//GUID
	Status 				string 	//Enumeration
	DateLastModified 	string  //DateTime
	Title 				string 
}

type ORClassResources struct{
	SourcedId 			string 		//GUID
	Status				string 			//Enumeration
	DateLastModified 	string //DateTime
	Title 				string 
	ClassSourcedId 		string //GUID Reference
	ResourceSourcedId 	string //GUID Reference
}

type ORCourseResources struct{
	SourcedId 			string 		//GUID
	Status				string 			//Enumeration
	DateLastModified 	string //DateTime
	Title 				string 
	CourseSourcedId 	string //GUID Reference
	ResourceSourcedId 	string //GUID Reference
}

type ORResource struct{
	SourcedId 			string 		//GUID
	Status 				string 			//Enumeration
	DateLastModified 	string //DateTime
	VendorResourceId 	string //id
	Title				string 
	Roles				[]string //Enumeration List
	Importance			string 
	VendorId			string //id
	ApplicationId		string //id
}

type ORResult struct {
	SourcedId string 		//GUID
	Status string 			//Enumeration
	DateLastModified string //DateTime
	LineItemSourcedId string//GUID Reference
	StudentSourcedId string //GUID Reference
	ScoreStatus string 		//Enumeration
	Score float64 			//float
	ScoreDate string 		//date
	Comment string

}

type ORLineItems struct {
	SourcedId 			string 			//GUID
	Status				string 			//Enumeration
	DateLastModified 	string //DateTime
	Title 				string 
	Description			string
	SssignDate			string //date
	DueDate				string //date
	ClassSourcedId		string // GUID References
	CategorySourcedId	string // GUID References
	GradingPeriodSourcedId	string // GUID References
	ResultValueMin 		float64
	ResultValueMax		float64
}

// import type 
const (
	IMPORT_TYPE_BULK = "bulk"
	IMPORT_TYPE_DELTA = "delta"
	IMPORT_TYPE_ABSENT = "absent"
)

// manifest property names 
const (
	MANIFEST_PRO_VERSION = "manifest.version"
	MANIFEST_PRO_ONEROSTER_VERSION = "oneroster.version"
	MANIFEST_PRO_FILE_ACADEMICSESSIONS = "file.academicSessions"
	MANIFEST_PRO_FILE_CATEGORIES = "file.categories"
	MANIFEST_PRO_FILE_CLASSES = "file.classes"
	MANIFEST_PRO_FILE_CLASSRESOURCES = "file.classResources"
	MANIFEST_PRO_FILE_COURSES = "file.courses"
	MANIFEST_PRO_FILE_COURSERESOURCES = "file.courseResources"
	MANIFEST_PRO_FILE_DEMOGRAPHICS = "file.demographics"
	MANIFEST_PRO_FILE_ENROLLMENTS = "file.enrollments"
	MANIFEST_PRO_FILE_LINEITEMS = "file.lineItems"
	MANIFEST_PRO_FILE_ORGS = "file.orgs"
	MANIFEST_PRO_FILE_RESOURCES = "file.resources"
	MANIFEST_PRO_FILE_RESULTS = "file.results"
	MANIFEST_PRO_FILE_USERS = "file.users"
	MANIFEST_PRO_SOURCE_SYSTEMNAME = "source.systemName"
	MANIFEST_PRO_SOURCE_SYSTEMCODE = "source.systemCode"
)

//csv files name 
const (
	CSV_NAME_MANIFEST = "manifest.csv"
	CSV_NAME_ACADEMICSESSIONS = "academicSessions.csv"
	CSV_NAME_CATEGORIES = "categories.csv"
	CSV_NAME_CLASSES = "classes.csv"
	CSV_NAME_COURSES = "courses.csv"
	CSV_NAME_CLASSRESOURCES = "classResources.csv"
	CSV_NAME_DEMOGRAPHICS = "demographics.csv"
	CSV_NAME_ENROLLMENTS = "enrollments.csv"
	CSV_NAME_ORGS = "orgs.csv"
	CSV_NAME_RESOURCES = "resources.csv"
	CSV_NAME_LINEITEMS = "lineItems.csv"
	CSV_NAME_RESULTS = "results.csv"
	CSV_NAME_USERS = "users.csv"
)

//orgs types
const (
	ORG_TYPE_DISTRICT = "district"
	ORG_TYPE_SCHOOL = "school"

)

