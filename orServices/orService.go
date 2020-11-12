package orServices

import (
	// "encoding/csv"
    "fmt"
	"os"
	// "strconv"
	"github.com/MZDevinc/oneroster/models"
	// "github.com/jszwec/csvutil"
	"github.com/gocarina/gocsv"
	"strings"
) 



func ProcessFiles(dirPath string, orProcess models.ORProcess)  error{
	// fmt.Println(" >> inside ProccessFiles -> filePath: ",filePath)
	// orUser := models.ORUser{}
	// fmt.Println(" >> inside call interface func<<<  ")
	// inResult := orProcess.HandleAddUser(orUser)
	// fmt.Println(">>>  orProcess inResult: ", inResult)

		// read the manifest csv file 
	manifestPath := fmt.Sprintf("%s/manifest.csv", dirPath)
	manifestRows, err := ReadManifestCsv(manifestPath)
	if err != nil {
		fmt.Println(">> err ReadManifestCsv: ",err)
		return err
	}
	fmt.Println(">> manifestRows: ",manifestRows)
	var orgDistrict []models.OROrg

	// put the manifest data into   --- map[propertyName] = propertyValue---
	mainfestTable := make(map[string]string)
	for _, manifestRow := range manifestRows {
		switch manifestRow.PropertyName {

		// class sourceName
		case models.MANIFEST_PRO_SOURCE_SYSTEMNAME:
			mainfestTable[models.MANIFEST_PRO_SOURCE_SYSTEMNAME] = manifestRow.Value

		// Acadimic Sessions		
		case models.MANIFEST_PRO_FILE_ACADEMICSESSIONS:
			mainfestTable[models.MANIFEST_PRO_FILE_ACADEMICSESSIONS] = manifestRow.Value
		
		// classes
		case models.MANIFEST_PRO_FILE_CLASSES:
			mainfestTable[models.MANIFEST_PRO_FILE_CLASSES] = manifestRow.Value
		
		// courses	
		case models.MANIFEST_PRO_FILE_COURSES:
			mainfestTable[models.MANIFEST_PRO_FILE_COURSES] = manifestRow.Value
		
		// Enrollments
		case models.MANIFEST_PRO_FILE_ENROLLMENTS:
			mainfestTable[models.MANIFEST_PRO_FILE_ENROLLMENTS] = manifestRow.Value
	
		// Orgs
		case models.MANIFEST_PRO_FILE_ORGS:
			mainfestTable[models.MANIFEST_PRO_FILE_ORGS] = manifestRow.Value
		
		// Users	
		case models.MANIFEST_PRO_FILE_USERS:
			mainfestTable[models.MANIFEST_PRO_FILE_USERS] = manifestRow.Value

		// Demographics ( we don't save it in Edgems, maybe we will need it later )
		case models.MANIFEST_PRO_FILE_DEMOGRAPHICS:
			mainfestTable[models.MANIFEST_PRO_FILE_DEMOGRAPHICS] = manifestRow.Value

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
	if mainfestTable[models.MANIFEST_PRO_FILE_ORGS] != models.IMPORT_TYPE_ABSENT {
		doRollback := false
		if strings.Contains(strings.ToLower(mainfestTable[models.MANIFEST_PRO_SOURCE_SYSTEMNAME]),"classlink"){
			orgDistrict, doRollback, err = ProcessOrgsClassLink(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_ORGS])
		} else {
			orgDistrict, doRollback, err = ProcessOrgs(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_ORGS])
		}
		
		if err != nil {
			if doRollback {
				fmt.Println(">>> (rollback) errer happen when processOrgs err -> ",err)
				err = orProcess.RollBackOneRoster(orgDistrict)
			} else {
				fmt.Println(">>> (no rollback) errer happen when processOrgs err -> ",err)
			}
			return err
		}
	}
	//process Courses
	if mainfestTable[models.MANIFEST_PRO_FILE_COURSES] != models.IMPORT_TYPE_ABSENT {
		err = ProcessCourses(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_COURSES])
		if err != nil {
			fmt.Println(">>> (rollback) errer happen when ProcessCourses err -> ",err)
			err = orProcess.RollBackOneRoster(orgDistrict)
			return err
		}
	}

	//process Academic Session
	if mainfestTable[models.MANIFEST_PRO_FILE_ACADEMICSESSIONS] != models.IMPORT_TYPE_ABSENT {
		err = ProcessAcademicSessions(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_ACADEMICSESSIONS])
		if err != nil {
			fmt.Println(">>> (rollback) errer happen when ProcessAcademicSessions err -> ",err)
			err = orProcess.RollBackOneRoster(orgDistrict)
			return err
		}
	}
	//process Classes
	if mainfestTable[models.MANIFEST_PRO_FILE_CLASSES] != models.IMPORT_TYPE_ABSENT {
		err = ProcessClasses(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_CLASSES])
		if err != nil {
			fmt.Println(">>> (rollback) errer happen when ProcessClasses err -> ",err)
			err = orProcess.RollBackOneRoster(orgDistrict)
			return err
		}
	}

	//process Users
	if mainfestTable[models.MANIFEST_PRO_FILE_USERS] != models.IMPORT_TYPE_ABSENT {
		err = ProcessUsers(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_USERS])
		if err != nil {
			fmt.Println(">>> (rollback) errer happen when ProcessUsers err -> ",err)
			err = orProcess.RollBackOneRoster(orgDistrict)
			return err
		}
	}

	//process User Entrollments
	if mainfestTable[models.MANIFEST_PRO_FILE_ENROLLMENTS] != models.IMPORT_TYPE_ABSENT {
		err = ProcessEntrollment(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_ENROLLMENTS])
		if err != nil {
			fmt.Println(">>> (rollback) errer happen when ProcessEntrollments err -> ",err)
			err = orProcess.RollBackOneRoster(orgDistrict)
			return err
		}
	}

	//process Demographics  (we don't use it right now)
	// if mainfestTable[models.MANIFEST_PRO_FILE_DEMOGRAPHICS] != models.IMPORT_TYPE_ABSENT {
	// 	err = ProcessDemographics(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_DEMOGRAPHICS])
	// 	if err != nil {
	// 		fmt.Println(">>> (rollback) errer happen when ProcessDemographics err -> ",err)
	// 		err = orProcess.RollBackOneRoster(orgDistrict)
	// 		return err
	// 	}
	// }

	return nil;
}


func ReadManifestCsv(filename string) ([]models.OrManifest, error) {
    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
	var manifestValues []models.OrManifest

	err = gocsv.UnmarshalFile(f, &manifestValues)
	if err != nil { 
		fmt.Println(">>>> error UnmarshalFile <<< using gocsv: ")
		fmt.Println(err)
	}
    return manifestValues, nil
}

func ProcessAcademicSessions(dirPath string, orProcess models.ORProcess, importType string) error {

	academicSessionsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_ACADEMICSESSIONS)

	f, err := os.Open(academicSessionsPath)
    if err != nil {
        return  err
    }
    defer f.Close()
	var academicSessions []models.ORAcademicSessions

	err = gocsv.UnmarshalFile(f, &academicSessions)
	if err != nil { 
		fmt.Println(">>>> error UnmarshalFile academicSessions <<< using gocsv: ")
		fmt.Println(err)
		return err
	}
	if importType == models.IMPORT_TYPE_BULK {
	
		err := orProcess.HandleAddAcademicSessions(academicSessions)
			if err != nil {
				fmt.Println(">>> ProcessAcademicSessions error ",err)
				return err
			}
	}

    return nil
}



func ProcessOrgs(dirPath string, orProcess models.ORProcess, importType string) ([]models.OROrg, bool, error) {

	var orgDistricts []models.OROrg
	// do rollback for all district or not
	rollback := true
	orgsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_ORGS)

	f, err := os.Open(orgsPath)
    if err != nil {
        return  orgDistricts, rollback, err
    }
    defer f.Close()
	var orgs []models.OROrg

	err = gocsv.UnmarshalFile(f, &orgs)
	if err != nil { 
		fmt.Println(err)
		return orgDistricts, rollback, err
	}
	
	if importType == models.IMPORT_TYPE_BULK {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.ORG_TYPE_DISTRICT {
				// collect all district 
				orgDistricts = append(orgDistricts, org)
				
				rollback, err = orProcess.HandleAddDistrict(org)
				if err != nil {
					return orgDistricts,rollback, err
				}
			} else if org.OrgType == models.ORG_TYPE_SCHOOL {
				err = orProcess.HandleAddSchool(org)
				if err != nil {
					return orgDistricts, true, err
				}
			}
		}
	}

    return orgDistricts, false, nil
}


func ProcessOrgsClassLink(dirPath string, orProcess models.ORProcess, importType string) ([]models.OROrg, bool, error) {

	var orgDistricts []models.OROrg
	// do rollback for all district or not
	rollback := true
	orgsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_ORGS)

	f, err := os.Open(orgsPath)
    if err != nil {
        return  orgDistricts, rollback, err
    }
    defer f.Close()
	var orgs []models.OROrg

	err = gocsv.UnmarshalFile(f, &orgs)
	if err != nil { 
		fmt.Println(err)
		return orgDistricts, rollback, err
	}
	
	if importType == models.IMPORT_TYPE_BULK {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.ORG_TYPE_DISTRICT {
				// collect all district 
				orgDistricts = append(orgDistricts, org)
				
				rollback, err = orProcess.HandleAddDistrict(org)
				if err != nil {
					return orgDistricts,rollback, err
				}
			} 
		}

		for _, org := range orgs{
			if org.OrgType == models.ORG_TYPE_SCHOOL {
				if len(orgDistricts) ==1 && org.ParentSourcedId ==""{
					org.ParentSourcedId = orgDistricts[0].SourcedId
				}
				err = orProcess.HandleAddSchool(org)
				if err != nil {
					return orgDistricts, true, err
				}
			}
		}
	}

    return orgDistricts, false, nil
}


func ProcessCourses(dirPath string, orProcess models.ORProcess, importType string) error {

	fmt.Println(">>> ProcessCourses <<<")
	coursesPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_COURSES)

	f, err := os.Open(coursesPath)
    if err != nil {
        return err
    }
    defer f.Close()
	var orCourses []models.ORCourse
	fmt.Println(">>> ProcessCourses before UnmarshalFile <<<")
	err = gocsv.UnmarshalFile(f, &orCourses)
	if err != nil { 
		fmt.Println(">>> ProcessCourses error UnmarshalFile",err)
		return err
	}
	

	fmt.Println(">>> ProcessCourses  orCourses length: ",len(orCourses))
	if importType == models.IMPORT_TYPE_BULK {
			err := orProcess.HandleAddCourses(orCourses)
			if err != nil {
				fmt.Println(">>> ProcessCourses error ",err)
				return err
			}
	}
    return nil
}


func ProcessClasses(dirPath string, orProcess models.ORProcess, importType string) error {

	fmt.Println(">>> ProcessClass <<<")
	classesPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_CLASSES)

	f, err := os.Open(classesPath)
    if err != nil {
        return err
    }
    defer f.Close()
	var orClasses []models.ORClass
	fmt.Println(">>> ProcessClasses before UnmarshalFile <<<")
	err = gocsv.UnmarshalFile(f, &orClasses)
	if err != nil { 
		fmt.Println(">>> ProcessClasses error UnmarshalFile",err)
		return err
	}
	
	if importType == models.IMPORT_TYPE_BULK {
			err := orProcess.HandleAddClasses(orClasses)
			if err != nil {
				fmt.Println(">>> ProcessClasses error ",err)
				return err
			}
	}
    return nil
}

func ProcessUsers(dirPath string, orProcess models.ORProcess, importType string) error {

	usersPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_USERS)

	f, err := os.Open(usersPath)
    if err != nil {
        return  err
    }
    defer f.Close()
	var orUsers []models.ORUser

	err = gocsv.UnmarshalFile(f, &orUsers)
	if err != nil { 
		fmt.Println(">>> ProcessUsers error UnmarshalFile",err)
		return err
	}
	if importType == models.IMPORT_TYPE_BULK {
		err := orProcess.HandleAddUser(orUsers)
		if err != nil {
			fmt.Println(">>> ProcessUsers error ",err)
			return err
		}
	}

    return nil
}


func ProcessEntrollment(dirPath string, orProcess models.ORProcess, importType string) error {

	entrollmentsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_ENROLLMENTS)

	f, err := os.Open(entrollmentsPath)
    if err != nil {
        return  err
    }
    defer f.Close()
	var orEntrollments []models.OREnrollment

	err = gocsv.UnmarshalFile(f, &orEntrollments)
	if err != nil { 
		fmt.Println(">>> ProcessEntrollment error UnmarshalFile",err)
		return err
	}
	if importType == models.IMPORT_TYPE_BULK {
		err := orProcess.HandleAddEnrollment(orEntrollments)
		if err != nil {
			fmt.Println(">>> ProcessEntrollments error ",err)
			return err
		}
	}

    return nil
}

func ProcessDemographics(dirPath string, orProcess models.ORProcess, importType string) error {

	demographicsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_DEMOGRAPHICS)

	f, err := os.Open(demographicsPath)
    if err != nil {
        return  err
    }
    defer f.Close()
	var orDemographics []models.ORDemographics

	err = gocsv.UnmarshalFile(f, &orDemographics)
	if err != nil { 
		fmt.Println(">>> ProcessDemographics error UnmarshalFile",err)
		return err
	}
	if importType == models.IMPORT_TYPE_BULK {
		// err := orProcess.HandleAddorDemographics(orEntrollments)
		// if err != nil {
		// 	fmt.Println(">>> ProcessEntrollments error ",err)
		// 	return err
		// }
	}

    return nil
}

