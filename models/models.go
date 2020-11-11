package models



type ORProcess interface {
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
	// manifestVersion 		string `csv:"manifest.version"`
	// onerosterVersion 		string `csv:"oneroster.version"`
	// fileAcademicSessions 	string `csv:"file.academicSessions"`//Enumeration
	// fileCategories 			string `csv:"file.categories"`//Enumeration
	// fileClasses 			string `csv:"file.classes"`//Enumeration
	// fileClassResources 		string `csv:"file.classResources"`//Enumeration
	// fileCourses 			string `csv:"file.courses"`//Enumeration
	// fileCourseResources 	string `csv:"file.courseResources"`//Enumeration
	// fileDemographics 		string `csv:"file.demographics"`//Enumeration
	// fileEnrollments 		string `csv:"file.enrollments"`//Enumeration
	// fileLineItems 			string `csv:"file.lineItems"`//Enumeration
	// fileOrgs 				string `csv:"file.orgs"`//Enumeration
	// fileResources 			string `csv:"file.resources"`//Enumeration
	// fileResults 			string `csv:"file.results"`//Enumeration
	// fileUsers 				string `csv:"file.users"`//Enumeration
	// sourceSystemName		string `csv:"source.systemName"`
	// sourceSystemCode		string `csv:"source.systemCode"`
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

type ORCategory struct{
	sourcedId 			string 	//GUID
	status 				string 	//Enumeration
	dateLastModified 	string  //DateTime
	title 				string 
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

type ORClassResources struct{
	sourcedId 			string 		//GUID
	status				string 			//Enumeration
	dateLastModified 	string //DateTime
	title 				string 
	classSourcedId 		string //GUID Reference
	resourceSourcedId 	string //GUID Reference
}

type ORCourseResources struct{
	sourcedId 			string 		//GUID
	status				string 			//Enumeration
	dateLastModified 	string //DateTime
	title 				string 
	courseSourcedId 	string //GUID Reference
	resourceSourcedId 	string //GUID Reference
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
	sourcedId 			string 		//GUID
	status				string 			//Enumeration
	dateLastModified 	string //DateTime
	birthDate			string //date
	sex					string //Enumeration
	americanIndianOrAlaskaNative 			string //Enumeration
	asian									string //Enumeration
	blackOrAfricanAmerican					string //Enumeration
	nativeHawaiianOrOtherPacificIslander 	string //Enumeration
	white									string //Enumeration
	demographicRaceTwoOrMoreRaces			string //Enumeration
	hispanicOrLatinoEthnicity				string //Enumeration
	countryOfBirthCode						string
	stateOfBirthAbbreviation				string
	cityOfBirth								string
	publicSchoolResidenceStatus 			string
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

type ORLineItems struct {
	sourcedId 			string 		//GUID
	status				string 			//Enumeration
	dateLastModified 	string //DateTime
	title 				string 
	description			string
	assignDate			string //date
	dueDate				string //date
	classSourcedId		string // GUID References
	categorySourcedId	string // GUID References
	gradingPeriodSourcedId	string // GUID References
	resultValueMin 		float64
	resultValueMax		float64
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

type ORResource struct{
	sourcedId 			string 		//GUID
	status 				string 			//Enumeration
	dateLastModified 	string //DateTime
	vendorResourceId 	string //id
	title				string 
	roles				[]string //Enumeration List
	importance			string 
	vendorId			string //id
	applicationId		string //id
}


type ORResult struct {
	sourcedId string 		//GUID
	status string 			//Enumeration
	dateLastModified string //DateTime
	lineItemSourcedId string//GUID Reference
	studentSourcedId string //GUID Reference
	scoreStatus string 		//Enumeration
	score float64 			//float
	scoreDate string 		//date
	comment string

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

