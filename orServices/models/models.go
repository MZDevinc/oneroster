package models



type ORProcess interface {
	HandleAddUser() string
	HandleReomveuser() string
	HandleEditUser() string

	HandleAddDistrict() string
	HandleReomveDistrict() string
	HandleEditDistrict() string

	HandleAddSchool() string
	HandleReomveSchool() string
	HandleEditSchool() string

	HandleAddClass() string
	HandleReomveClass() string
	HandleEditClass() string

	HandleAddCourse() string
	HandleReomveCourse() string
	HandleEditCourse() string

}

type OrManifest struct {
	manifestVersion 		string
	onerosterVersion 		string 
	fileAcademicSessions 	string //Enumeration
	fileCategories 			string //Enumeration
	fileClasses 			string //Enumeration
	fileClassResources 		string //Enumeration
	fileCourses 			string //Enumeration
	fileCourseResources 	string //Enumeration
	fileDemographics 		string //Enumeration
	fileEnrollments 		string //Enumeration
	fileLineItems 			string //Enumeration
	fileOrgs 				string //Enumeration
	fileResources 			string //Enumeration
	fileResults 			string //Enumeration
	fileUsers 				string //Enumeration
	sourceSystemName		string 
	sourceSystemCode		string
}

type ORAcademicSessions struct{
	sourcedId 			string //GUID
	status 				string //Enumeration
	dateLastModified 	string //DateTime
	title 				string
	sessionType 		string //Enumeration
	startDate 			string //date
	endDate 			string //date
	parentSourcedId 	string //GUID Reference
	schoolYear 			string //year
}

type ORCategory struct{
	sourcedId 			string 	//GUID
	status 				string 	//Enumeration
	dateLastModified 	string  //DateTime
	title 				string 
}

type ORClass struct{
	sourcedId 			string 		//GUID
	status 				string 			//Enumeration
	dateLastModified 	string //DateTime
	title 				string 
	grades 				[]string 
	courseSourcedId 	string //GUID Reference
	classCode 			string
	classType 			string //Enumeration
	location 			string 
	schoolSourcedId 	string //GUID Reference
	termSourcedIds 		[]string //List of GUID Reference
	subjects 			[]string 
	subjectCodes 		[]string
	periods 			[]string

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
	sourcedId 			string 		//GUID
	status				string 			//Enumeration
	dateLastModified 	string //DateTime
	schoolYearSourcedId string //GUID Reference
	title 				string 
	courseCode			string
	grades				[]string
	orgSourcedId		string //GUID Reference
	subjects			[]string
	subjectCodes		string
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
	sourcedId 			string 		//GUID
	status				string 			//Enumeration
	dateLastModified 	string //DateTime
	classSourcedId		string //GUID Reference
	schoolSourcedId		string //GUID Reference
	userSourcedId 		string //GUID Reference
	role				string //Enumeration
	primary				bool
	beginDate			string //date
	endDate				string //date
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
	sourcedId 			string 		//GUID
	status 				string 			//Enumeration
	dateLastModified 	string //DateTime
	name 				string
	orgType 			string			// type Enumeration
	identifier 			string
	parentSourcedId 	string 	//GUID Reference
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
	sourcedId 			string 		//GUID
	status 				string 			//Enumeration
	dateLastModified 	string //DateTime
	enabledUser 		bool
	orgSourcedIds 		string 	//List of GUID References.
	role 				string 			//Enumeration
	username 			string
	userIds 			[]string
	givenName 			string
	familyName 			string
	middleName 			string
	identifier 			string 
	email 				string
	sms 				string
	phone 				string
	agentSourcedIds 	[]string //List of GUID References
	grades 				string 
	password 			string

}