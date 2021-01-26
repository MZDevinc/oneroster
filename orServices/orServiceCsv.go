package orservices

import (
	"fmt"
	"os"

	"github.com/MZDevinc/oneroster/models"

	// "github.com/jszwec/csvutil"
	"strings"

	"github.com/gocarina/gocsv"
)

func ProcessFiles(dirPath string, orProcess models.ORProcess) error {

	// read the manifest csv file
	manifestPath := fmt.Sprintf("%s/manifest.csv", dirPath)
	manifestRows, err := ReadManifestCSV(manifestPath)
	if err != nil {
		fmt.Println(">> err ReadManifestCsv: ", err)
		return err
	}

	var orgDistrict []models.OROrg
	var districtIDs []string

	// put the manifest data into   --- map[propertyName] = propertyValue---
	mainfestTable := make(map[string]string)
	for _, manifestRow := range manifestRows {
		switch manifestRow.PropertyName {

		// class sourceName
		case models.ManifestProSourceSystemName:
			mainfestTable[models.ManifestProSourceSystemName] = manifestRow.Value

		// Acadimic Sessions
		case models.ManifestProFileAcademicSessions:
			mainfestTable[models.ManifestProFileAcademicSessions] = manifestRow.Value

		// classes
		case models.ManifestProFileClasses:
			mainfestTable[models.ManifestProFileClasses] = manifestRow.Value

		// courses
		case models.ManifestProFileCourses:
			mainfestTable[models.ManifestProFileCourses] = manifestRow.Value

		// Enrollments
		case models.ManifestProFileEnrollments:
			mainfestTable[models.ManifestProFileEnrollments] = manifestRow.Value

		// Orgs
		case models.ManifestProFileOrgs:
			mainfestTable[models.ManifestProFileOrgs] = manifestRow.Value

		// Users
		case models.ManifestProFileUsers:
			mainfestTable[models.ManifestProFileUsers] = manifestRow.Value

		// Demographics ( we don't save it in Edgems, maybe we will need it later )
		case models.ManifestProFileDemographics:
			mainfestTable[models.ManifestProFileDemographics] = manifestRow.Value

			// we don't read the resources and results, we don't need it until now

			// case models.MANIFEST_PRO_FILE_RESULTS:
			// mainfestTable[models.MANIFEST_PRO_FILE_RESULTS] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_RESOURCES:
			// mainfestTable[models.MANIFEST_PRO_FILE_RESOURCES] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_LINEITEMS:
			// mainfestTable[models.MANIFEST_PRO_FILE_LINEITEMS] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_COURSERESOURCES:
			// mainfestTable[models.MANIFEST_PRO_FILE_COURSERESOURCES] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_CLASSRESOURCES:
			// mainfestTable[models.MANIFEST_PRO_FILE_CLASSRESOURCES] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_CATEGORIES:
			// mainfestTable[models.MANIFEST_PRO_FILE_CATEGORIES] = manifestRow.Value
		}

	}

	// the files should be readed in order
	//process Ditricts and schools
	if mainfestTable[models.ManifestProFileOrgs] != models.ImportTypeAbsent {
		doRollback := false
		if strings.Contains(strings.ToLower(mainfestTable[models.ManifestProSourceSystemName]), "classlink") {
			orgDistrict, districtIDs, doRollback, err = ProcessOrgsClassLinkCSV(dirPath, orProcess, mainfestTable[models.ManifestProFileOrgs])
		} else {
			orgDistrict, districtIDs, doRollback, err = ProcessOrgsCSV(dirPath, orProcess, mainfestTable[models.ManifestProFileOrgs])
		}

		if err != nil {
			if doRollback {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}
	//process Courses
	if mainfestTable[models.ManifestProFileCourses] != models.ImportTypeAbsent {
		err = ProcessCoursesCSV(dirPath, orProcess, mainfestTable[models.ManifestProFileCourses], districtIDs)
		if err != nil {
			if mainfestTable[models.ManifestProFileCourses] != models.ImportTypeBulk {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}

	//process Academic Session
	if mainfestTable[models.ManifestProFileAcademicSessions] != models.ImportTypeAbsent {
		err = ProcessAcademicSessionsCSV(dirPath, orProcess, mainfestTable[models.ManifestProFileAcademicSessions])
		if err != nil {
			if mainfestTable[models.ManifestProFileAcademicSessions] != models.ImportTypeBulk {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}
	//process Classes
	if mainfestTable[models.ManifestProFileClasses] != models.ImportTypeAbsent {
		err = ProcessClassesCSV(dirPath, orProcess, mainfestTable[models.ManifestProFileClasses], districtIDs)
		if err != nil {
			if mainfestTable[models.ManifestProFileClasses] != models.ImportTypeBulk {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}

	//process Users
	if mainfestTable[models.ManifestProFileUsers] != models.ImportTypeAbsent {
		err = ProcessUsersCSV(dirPath, orProcess, mainfestTable[models.ManifestProFileUsers], districtIDs)
		if err != nil {
			if mainfestTable[models.ManifestProFileUsers] != models.ImportTypeBulk {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}

	//process User Entrollments
	if mainfestTable[models.ManifestProFileEnrollments] != models.ImportTypeAbsent {
		err = ProcessEnrollmentCSV(dirPath, orProcess, mainfestTable[models.ManifestProFileEnrollments], districtIDs)
		if err != nil {
			if mainfestTable[models.ManifestProFileEnrollments] != models.ImportTypeBulk {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}

	//process Demographics  (we don't use it right now)
	// if mainfestTable[models.ManifestProFileDemographics] != models.ImportTypeAbsent {
	// 	err = ProcessDemographics(dirPath, orProcess, mainfestTable[models.ManifestProFileDemographics])
	// 	if err != nil {
	// 		fmt.Println(">>> (rollback) errer happen when ProcessDemographics err -> ",err)
	// 		err = orProcess.RollBackOneRoster(orgDistrict)
	// 		return err
	// 	}
	// }

	return nil
}

func ReadManifestCSV(filename string) ([]models.OrManifest, error) {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var manifestValues []models.OrManifest

	err = gocsv.UnmarshalFile(f, &manifestValues)
	if err != nil {
		return nil, err
	}
	return manifestValues, nil
}

func ProcessAcademicSessionsCSV(dirPath string, orProcess models.ORProcess, importType string) error {

	academicSessionsPath := fmt.Sprintf("%s/%s", dirPath, models.CsvNameAcademicSessions)

	f, err := os.Open(academicSessionsPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var academicSessions []models.ORAcademicSessions

	err = gocsv.UnmarshalFile(f, &academicSessions)
	if err != nil {
		return err
	}
	if importType == models.ImportTypeBulk {

		err := orProcess.HandleAddAcademicSessions(academicSessions)
		if err != nil {
			return err
		}
	} else if importType == models.ImportTypeDelta {
		orAcademicSessionToEdit := []models.ORAcademicSessions{}
		orAcademicSessionIDsToDelete := []string{}
		for _, orAcademicSession := range academicSessions {
			if orAcademicSession.Status == models.StatusTypeActive {
				orAcademicSessionToEdit = append(orAcademicSessionToEdit, orAcademicSession)
				// err = orProcess.HandleEditClass(orClass)
			} else if orAcademicSession.Status == models.StatusTypeToBeDeleted {
				orAcademicSessionIDsToDelete = append(orAcademicSessionIDsToDelete, orAcademicSession.SourcedID)
			}
			if err != nil {
				return err
			}
		}
		err = orProcess.HandleEditAcademicSessions(orAcademicSessionToEdit)
		err = orProcess.HandleDeleteAcademicSessions(orAcademicSessionIDsToDelete)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessOrgsCSV(dirPath string, orProcess models.ORProcess, importType string) ([]models.OROrg, []string, bool, error) {

	var orgDistricts []models.OROrg
	var districtIDs []string
	// do rollback for all district or not
	rollback := true
	orgsPath := fmt.Sprintf("%s/%s", dirPath, models.CsvNameOrgs)

	f, err := os.Open(orgsPath)
	if err != nil {
		return orgDistricts, districtIDs, rollback, err
	}
	defer f.Close()
	var orgs []models.OROrg

	err = gocsv.UnmarshalFile(f, &orgs)
	if err != nil {
		fmt.Println(err)
		return orgDistricts, districtIDs, rollback, err
	}

	if importType == models.ImportTypeBulk {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.OrgTypeDistrict {
				// collect all district
				orgDistricts = append(orgDistricts, org)

				rollback, err = orProcess.HandleAddDistrict(org)
				if err != nil {
					return orgDistricts, districtIDs, rollback, err
				}
			}
		}

		// get the mongo IDs for the district to use for edit and delete other files data
		districtIDs, err := orProcess.GetDistrictsIDs(orgDistricts)
		if err != nil {
			return orgDistricts, districtIDs, true, err
		}

		for _, org := range orgs {

			if org.OrgType == models.OrgTypeSchool {
				err = orProcess.HandleAddSchool(org, districtIDs)
				if err != nil {
					return orgDistricts, districtIDs, true, err
				}
			}
		}
	} else if importType == models.ImportTypeDelta {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.OrgTypeDistrict {
				// // collect all district
				orgDistricts = append(orgDistricts, org)

				if org.Status == models.StatusTypeActive {
					// err = orProcess.HandleEditDistrict(org)
					err = orProcess.HandleAddOrEditDistrict(org)
				} else if org.Status == models.StatusTypeToBeDeleted {
					err = orProcess.HandleDeleteDistrict(org)
				}

				if err != nil {
					return orgDistricts, districtIDs, false, err
				}

			}
		}

		// get the mongo IDs for the district to use for edit and delete other files data
		districtIDs, err := orProcess.GetDistrictsIDs(orgDistricts)
		if err != nil {
			return orgDistricts, districtIDs, true, err
		}

		for _, org := range orgs {
			if org.OrgType == models.OrgTypeSchool {
				if org.Status == models.StatusTypeActive {
					err = orProcess.HandleAddOrEditSchool(org, districtIDs)
				} else if org.Status == models.StatusTypeToBeDeleted {
					err = orProcess.HandleDeleteSchool(org, districtIDs)
				}

				if err != nil {
					return orgDistricts, districtIDs, false, err
				}
			}
		}

	}

	return orgDistricts, districtIDs, false, nil
}

func ProcessOrgsClassLinkCSV(dirPath string, orProcess models.ORProcess, importType string) ([]models.OROrg, []string, bool, error) {

	var orgDistricts []models.OROrg
	var districtIDs []string
	// do rollback for all district or not
	rollback := true
	orgsPath := fmt.Sprintf("%s/%s", dirPath, models.CsvNameOrgs)

	f, err := os.Open(orgsPath)
	if err != nil {
		return orgDistricts, districtIDs, rollback, err
	}
	defer f.Close()
	var orgs []models.OROrg

	err = gocsv.UnmarshalFile(f, &orgs)
	if err != nil {
		fmt.Println(err)
		return orgDistricts, districtIDs, rollback, err
	}

	if importType == models.ImportTypeBulk {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.OrgTypeDistrict {
				// collect all district
				orgDistricts = append(orgDistricts, org)

				rollback, err = orProcess.HandleAddDistrict(org)
				if err != nil {
					return orgDistricts, districtIDs, rollback, err
				}
			}
		}

		districtIDs, err := orProcess.GetDistrictsIDs(orgDistricts)
		if err != nil {
			return orgDistricts, districtIDs, true, err
		}

		for _, org := range orgs {
			if org.OrgType == models.OrgTypeSchool {
				if len(orgDistricts) == 1 && org.ParentSourcedID == "" {
					org.ParentSourcedID = orgDistricts[0].SourcedID
				}
				err = orProcess.HandleAddSchool(org, districtIDs)
				if err != nil {
					return orgDistricts, districtIDs, true, err
				}
			}
		}
	} else if importType == models.ImportTypeDelta {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.OrgTypeDistrict {
				// collect all district
				orgDistricts = append(orgDistricts, org)

				if org.Status == models.StatusTypeActive {
					// err = orProcess.HandleEditDistrict(org)
					err = orProcess.HandleAddOrEditDistrict(org)
				} else if org.Status == models.StatusTypeToBeDeleted {
					err = orProcess.HandleDeleteDistrict(org)
				}

				if err != nil {
					return orgDistricts, districtIDs, false, err
				}

			}
		}

		districtIDs, err := orProcess.GetDistrictsIDs(orgDistricts)
		if err != nil {
			return orgDistricts, districtIDs, true, err
		}

		for _, org := range orgs {
			if org.OrgType == models.OrgTypeSchool {
				if org.Status == models.StatusTypeActive {
					err = orProcess.HandleAddOrEditSchool(org, districtIDs)
				} else if org.Status == models.StatusTypeToBeDeleted {
					err = orProcess.HandleDeleteSchool(org, districtIDs)
				}

				if err != nil {
					return orgDistricts, districtIDs, false, err
				}
			}
		}
	}

	return orgDistricts, districtIDs, false, nil
}

func ProcessCoursesCSV(dirPath string, orProcess models.ORProcess, importType string, districtIDs []string) error {

	coursesPath := fmt.Sprintf("%s/%s", dirPath, models.CsvNameCourses)

	f, err := os.Open(coursesPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orCourses []models.ORCourse
	err = gocsv.UnmarshalFile(f, &orCourses)
	if err != nil {
		return err
	}

	if importType == models.ImportTypeBulk {
		err := orProcess.HandleAddCourses(orCourses, districtIDs)
		if err != nil {
			return err
		}
	} else if importType == models.ImportTypeDelta {
		orCourseToEdit := []models.ORCourse{}
		orCoursesIDsToDelete := []string{}
		for _, orCourse := range orCourses {
			if orCourse.Status == models.StatusTypeActive {
				// err = orProcess.HandleEditCourse(orCourse)
				orCourseToEdit = append(orCourseToEdit, orCourse)
			} else if orCourse.Status == models.StatusTypeToBeDeleted {
				// err = orProcess.HandleDeleteCourse(orCourse)
				orCoursesIDsToDelete = append(orCoursesIDsToDelete, orCourse.SourcedID)
			}
			if err != nil {
				return err
			}
		}
		err = orProcess.HandleAddOrEditCourse(orCourseToEdit, districtIDs)
		err = orProcess.HandleDeleteCourses(orCoursesIDsToDelete, districtIDs)
		if err != nil {
			return err
		}

	}
	return nil
}

// UnmarshalClassesString unmarshals oneroster classes csv string
func UnmarshalClassesString(csvStr string) ([]models.ORClass, error) {
	var orClasses []models.ORClass
	if err := gocsv.UnmarshalString(csvStr, &orClasses); err != nil {
		return orClasses, err
	}
	return orClasses, nil
}

// ProcessClassesString process classes from CSV string
func ProcessClassesString(csvStr string, orProcess models.ORProcess, importType string, districtIDs []string) error {
	orClasses, err := UnmarshalClassesString(csvStr)
	if err != nil {
		return err
	}
	return ProcessClasses(orProcess, orClasses, importType, districtIDs)
}

// ProcessClassesCSV processes oneroster csv file for classes
func ProcessClassesCSV(dirPath string, orProcess models.ORProcess, importType string, districtIDs []string) error {
	classesPath := fmt.Sprintf("%s/%s", dirPath, models.CsvNameClasses)

	f, err := os.Open(classesPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orClasses []models.ORClass
	err = gocsv.UnmarshalFile(f, &orClasses)
	if err != nil {
		return err
	}

	return ProcessClasses(orProcess, orClasses, importType, districtIDs)
}

// ProcessClasses processes class data
func ProcessClasses(orProcess models.ORProcess, orClasses []models.ORClass, importType string, districtIDs []string) error {
	if importType == models.ImportTypeBulk {
		err := orProcess.HandleAddClasses(orClasses, districtIDs)
		if err != nil {
			return err
		}
	} else if importType == models.ImportTypeDelta {
		orClassesToEdit := []models.ORClass{}
		orClassIDsToDelete := []string{}
		for _, orClass := range orClasses {
			if orClass.Status == models.StatusTypeActive {
				orClassesToEdit = append(orClassesToEdit, orClass)
				// err = orProcess.HandleEditClass(orClass)
			} else if orClass.Status == models.StatusTypeToBeDeleted {
				orClassIDsToDelete = append(orClassIDsToDelete, orClass.SourcedID)
			}
		}
		if err := orProcess.HandleAddOrEditClass(orClassesToEdit, districtIDs); err != nil {
			return err
		}
		if err := orProcess.HandleDeleteClasses(orClassIDsToDelete, districtIDs); err != nil {
			return err
		}
	}
	return nil
}

// ProcessUsersCSV process users from CSV file
func ProcessUsersCSV(dirPath string, orProcess models.ORProcess, importType string, districtIDs []string) error {

	usersPath := fmt.Sprintf("%s/%s", dirPath, models.CsvNameUsers)

	f, err := os.Open(usersPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orUsers []models.ORUser

	if err := gocsv.UnmarshalFile(f, &orUsers); err != nil {
		return err
	}

	return ProcessUsers(orProcess, orUsers, importType, districtIDs)
}

// UnmarshalUsersString unmarshals oneroster users csv string
func UnmarshalUsersString(csvStr string) ([]models.ORUser, error) {
	var orUsers []models.ORUser
	if err := gocsv.UnmarshalString(csvStr, &orUsers); err != nil {
		return orUsers, err
	}
	return orUsers, nil
}

// ProcessUsersString process users from CSV string
func ProcessUsersString(csvStr string, orProcess models.ORProcess, importType string, districtIDs []string) error {
	orUsers, err := UnmarshalUsersString(csvStr)
	if err != nil {
		return err
	}
	return ProcessUsers(orProcess, orUsers, importType, districtIDs)
}

// ProcessUsers process oneroster users using oneroster process
func ProcessUsers(orProcess models.ORProcess, orUsers []models.ORUser, importType string, districtIDs []string) error {
	if importType == models.ImportTypeBulk {
		err := orProcess.HandleAddUsers(orUsers, districtIDs)
		if err != nil {
			return err
		}
	} else if importType == models.ImportTypeDelta {
		orUsersToEdit := []models.ORUser{}
		orUsersIDsToDelete := []string{}
		for _, orUser := range orUsers {
			if orUser.Status == models.StatusTypeActive {
				orUsersToEdit = append(orUsersToEdit, orUser)
			} else if orUser.Status == models.StatusTypeToBeDeleted {
				orUsersIDsToDelete = append(orUsersIDsToDelete, orUser.SourcedID)
			}
		}
		if err := orProcess.HandleAddOrEditUsers(orUsersToEdit, districtIDs); err != nil {
			return err
		}
		if err := orProcess.HandleDeleteUsers(orUsersIDsToDelete, districtIDs); err != nil {

			return err
		}
	}
	return nil
}

// Enrollments

// UnmarshalEnrollmentsString unmarshals oneroster enrollments csv string
func UnmarshalEnrollmentsString(csvStr string) ([]models.OREnrollment, error) {
	var orEnrollments []models.OREnrollment
	if err := gocsv.UnmarshalString(csvStr, &orEnrollments); err != nil {
		return orEnrollments, err
	}
	return orEnrollments, nil
}

// ProcessEnrollmentsString process enrollments from CSV string
func ProcessEnrollmentsString(csvStr string, orProcess models.ORProcess, importType string, districtIDs []string) error {
	orEnrollments, err := UnmarshalEnrollmentsString(csvStr)
	if err != nil {
		return err
	}
	return ProcessEnrollments(orProcess, orEnrollments, importType, districtIDs)
}

// ProcessEnrollments process oneroster enrollments
func ProcessEnrollments(orProcess models.ORProcess, orEnrollments []models.OREnrollment, importType string, districtIDs []string) error {
	if importType == models.ImportTypeBulk {
		err := orProcess.HandleAddEnrollment(orEnrollments, districtIDs)
		if err != nil {
			return err
		}
	} else if importType == models.ImportTypeDelta {

		orEntrollmentsToEdit := []models.OREnrollment{}
		orEntrollmentsIDsToDelete := []models.OREnrollment{}
		for _, orEntrollment := range orEnrollments {
			if orEntrollment.Status == models.StatusTypeActive {
				orEntrollmentsToEdit = append(orEntrollmentsToEdit, orEntrollment)
			} else if orEntrollment.Status == models.StatusTypeToBeDeleted {
				orEntrollmentsIDsToDelete = append(orEntrollmentsIDsToDelete, orEntrollment)
			}
		}
		if err := orProcess.HandleDeleteEnrollments(orEntrollmentsIDsToDelete, districtIDs); err != nil {
			return err
		}
		if err := orProcess.HandleAddOrEditEnrollments(orEntrollmentsToEdit, districtIDs); err != nil {
			return err
		}
	}
	return nil
}

// ProcessEnrollmentCSV process oneroster enrollment CSV
func ProcessEnrollmentCSV(dirPath string, orProcess models.ORProcess, importType string, districtIDs []string) error {

	entrollmentsPath := fmt.Sprintf("%s/%s", dirPath, models.CsvNameEnrollments)

	f, err := os.Open(entrollmentsPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orEntrollments []models.OREnrollment

	err = gocsv.UnmarshalFile(f, &orEntrollments)
	if err != nil {
		return err
	}
	return ProcessEnrollments(orProcess, orEntrollments, importType, districtIDs)
}

func ProcessDemographicsCSV(dirPath string, orProcess models.ORProcess, importType string) error {

	demographicsPath := fmt.Sprintf("%s/%s", dirPath, models.CsvNameDemographics)

	f, err := os.Open(demographicsPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orDemographics []models.ORDemographics

	err = gocsv.UnmarshalFile(f, &orDemographics)
	if err != nil {
		return err
	}
	if importType == models.ImportTypeBulk {
		// err := orProcess.HandleAddorDemographics(orEntrollments)
		// if err != nil {
		// 	fmt.Println(">>> ProcessEntrollments error ",err)
		// 	return err
		// }
	} else if importType == models.ImportTypeDelta {
		fmt.Println(">> *** Delta *** ProcessDemographics")
	}

	return nil
}
