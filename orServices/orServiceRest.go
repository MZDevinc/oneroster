package orServices

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/MZDevinc/oneroster/models"
	"github.com/MZDevinc/oneroster/oauth1"
)

func ProcessAPIs(districtId string, domain string, key, secret string, orProcess models.ORProcess) error {

	// err := ProcessAcademicSessionsAPITest(domain, key, secret, orProcess)
	//call orgs API
	var districtIDs []string
	districtIDs = append(districtIDs, districtId)

	_, err := ProcessOrgsAPI(districtIDs, domain, key, secret, orProcess)
	if err != nil {
		return err
	}

	//call Courses API
	err = ProcessCoursesAPI(domain, key, secret, orProcess, districtIDs)
	if err != nil {
		return err
	}

	//call AcademicSession API
	err = ProcessAcademicSessionsAPI(domain, key, secret, orProcess)
	if err != nil {
		return err
	}

	//call Classes API
	err = ProcessClassesAPI(domain, key, secret, orProcess, districtIDs)
	if err != nil {
		return err
	}

	//call Users API
	err = ProcessUsersAPI(domain, key, secret, orProcess, districtIDs)
	if err != nil {
		return err
	}

	//call Enrollment API
	err = ProcessEntrollmentAPI(domain, key, secret, orProcess, districtIDs)
	if err != nil {
		return err
	}

	return nil
}

// get the AcademicSessions from the API and call the interface methods to Handle the response
func ProcessAcademicSessionsAPI(domain string, key, secret string, orProcess models.ORProcess) error {

	var academicSessions []models.ORAcademicSessions
	// call the api
	limit := 100
	offset := 0
	hasNext := true
	for hasNext == true {
		url := fmt.Sprintf("%s/ims/oneroster/v1p1/academicSessions?limit=%s&offset=%s", domain, strconv.Itoa(limit), strconv.Itoa(offset))

		oneRoster := oauth1.OneRosterNew(key, secret)
		statusCode, response, header := oneRoster.MakeRosterRequest(url)
		totalRowsCount, _ := strconv.Atoi(header.Get("x-total-count"))

		b := []byte(response)

		// If status_code is 200, create array of users from response, otherwise return error
		if statusCode == 200 {

			var academicSessionsResponse models.AcademicSessionsResponse
			json.Unmarshal(b, &academicSessionsResponse)
			academicSessions = academicSessionsResponse.AcademicSessions

		} else if statusCode == 401 {
			return fmt.Errorf("Unauthorized Request: %s", response)
		} else if statusCode == 404 {
			return fmt.Errorf("Not found: %s", response)
		} else if statusCode == 500 {
			return fmt.Errorf("Server Error: %s", response)
		}

		orAcademicSessionToEdit := []models.ORAcademicSessions{}
		orAcademicSessionIDsToDelete := []string{}
		for _, orAcademicSession := range academicSessions {
			orAcademicSession.ParentSourcedID = orAcademicSession.Parent.SourcedID
			if strings.ToLower(orAcademicSession.Status) == strings.ToLower(models.StatusTypeActive) {
				orAcademicSessionToEdit = append(orAcademicSessionToEdit, orAcademicSession)
			} else if strings.ToLower(orAcademicSession.Status) == strings.ToLower(models.StatusTypeToBeDeleted) {
				orAcademicSessionIDsToDelete = append(orAcademicSessionIDsToDelete, orAcademicSession.SourcedID)
			}
		}

		// Add or Edit AcademicSessions
		if len(orAcademicSessionToEdit) > 0 {
			err := orProcess.HandleAddOrEditAcademicSessions(orAcademicSessionToEdit)
			if err != nil {
				return err
			}
		}
		// Delete AcademicSessions
		if len(orAcademicSessionIDsToDelete) > 0 {
			err := orProcess.HandleDeleteAcademicSessions(orAcademicSessionIDsToDelete)
			if err != nil {
				return err
			}
		}

		if totalRowsCount > (offset + limit) {
			offset = offset + 100
		} else {
			hasNext = false
			break
		}

	}
	return nil
}

// get the Orgs from the API and call the interface methods to Handle the response
func ProcessOrgsAPI(districtIDs []string, domain string, key, secret string, orProcess models.ORProcess) ([]models.OROrg, error) {

	var orgs []models.OROrg
	// // call the api
	limit := 100
	offset := 0
	hasNext := true
	districts := []models.OROrg{}

	for hasNext == true {
		url := fmt.Sprintf("%s/ims/oneroster/v1p1/orgs?limit=%s&offset=%s", domain, strconv.Itoa(limit), strconv.Itoa(offset))

		oneRoster := oauth1.OneRosterNew(key, secret)
		statusCode, response, header := oneRoster.MakeRosterRequest(url)
		totalRowsCount, _ := strconv.Atoi(header.Get("x-total-count"))

		b := []byte(response)

		// If status_code is 200, create array of users from response, otherwise return error
		if statusCode == 200 {

			var orgsResponse models.OrgsResponse
			json.Unmarshal(b, &orgsResponse)
			orgs = orgsResponse.Orgs

		} else if statusCode == 401 {
			return districts, fmt.Errorf("Unauthorized Request: %s", response)
		} else if statusCode == 404 {
			return districts, fmt.Errorf("Not found: %s", response)
		} else if statusCode == 500 {
			return districts, fmt.Errorf("Server Error: %s", response)
		}

		// Handle the districts
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.OrgTypeDistrict {

				if strings.ToLower(org.Status) == strings.ToLower(models.StatusTypeActive) {
					err = orProcess.HandleEditDistrict(org, districtIDs[0])
				} else if strings.ToLower(org.Status) == strings.ToLower(models.StatusTypeToBeDeleted) {
					err = orProcess.HandleDeleteDistrict(org)
				}

				if err != nil {
					return districts, err
				}
				districts = append(districts, org)
			}
		}

		// Handle the schools
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.OrgTypeSchool {
				org.ParentSourcedID = org.Parent.SourcedID
				if org.ParentSourcedID == "" {
					org.ParentSourcedID = districts[0].SourcedID
				}

				if strings.ToLower(org.Status) == strings.ToLower(models.StatusTypeActive) {
					err = orProcess.HandleAddOrEditSchool(org, districtIDs)
				} else if strings.ToLower(org.Status) == strings.ToLower(models.StatusTypeToBeDeleted) {
					err = orProcess.HandleDeleteSchool(org, districtIDs)
				}

				if err != nil {
					return districts, err
				}
			}
		}
		if totalRowsCount > (offset + limit) {
			offset = offset + 100
		} else {
			hasNext = false
			break
		}

	}
	return districts, nil
}

// get the Courses from the API and call the interface methods to Handle the response
func ProcessCoursesAPI(domain string, key, secret string, orProcess models.ORProcess, districtIDs []string) error {

	var orCourses []models.ORCourse
	// call the api
	limit := 100
	offset := 0
	hasNext := true
	for hasNext == true {
		url := fmt.Sprintf("%s/ims/oneroster/v1p1/courses?limit=%s&offset=%s", domain, strconv.Itoa(limit), strconv.Itoa(offset))

		oneRoster := oauth1.OneRosterNew(key, secret)
		statusCode, response, header := oneRoster.MakeRosterRequest(url)
		totalRowsCount, _ := strconv.Atoi(header.Get("x-total-count"))
		b := []byte(response)

		// If status_code is 200, create array of users from response, otherwise return error
		if statusCode == 200 {

			var coursesResponse models.CoursesResponse
			json.Unmarshal(b, &coursesResponse)
			orCourses = coursesResponse.Courses

		} else if statusCode == 401 {
			return fmt.Errorf("Unauthorized Request: %s", response)
		} else if statusCode == 404 {
			return fmt.Errorf("Not found: %s", response)
		} else if statusCode == 500 {
			return fmt.Errorf("Server Error: %s", response)
		}

		orCourseToEdit := []models.ORCourse{}
		orCoursesIDsToDelete := []string{}
		for _, orCourse := range orCourses {
			orCourse.OrgSourcedID = orCourse.Org.SourcedID
			if strings.ToLower(orCourse.Status) == strings.ToLower(models.StatusTypeActive) {
				orCourseToEdit = append(orCourseToEdit, orCourse)
			} else if strings.ToLower(orCourse.Status) == strings.ToLower(models.StatusTypeToBeDeleted) {
				orCoursesIDsToDelete = append(orCoursesIDsToDelete, orCourse.SourcedID)
			}

		}

		// Add or Edit Courses
		if len(orCourseToEdit) > 0 {
			err := orProcess.HandleAddOrEditCourse(orCourseToEdit, districtIDs)
			if err != nil {
				return err
			}
		}

		// delete Courses
		if len(orCoursesIDsToDelete) > 0 {
			err := orProcess.HandleDeleteCourses(orCoursesIDsToDelete, districtIDs)
			if err != nil {
				return err
			}
		}
		if totalRowsCount > (offset + limit) {
			offset = offset + 100
		} else {
			hasNext = false
			break
		}

	}

	return nil
}

// get the Classes from the API and call the interface methods to Handle the response
func ProcessClassesAPI(domain string, key, secret string, orProcess models.ORProcess, districtIDs []string) error {

	var orClasses []models.ORClass
	// call the api
	limit := 100
	offset := 0
	hasNext := true
	for hasNext == true {
		url := fmt.Sprintf("%s/ims/oneroster/v1p1/classes?limit=%s&offset=%s", domain, strconv.Itoa(limit), strconv.Itoa(offset))

		oneRoster := oauth1.OneRosterNew(key, secret)
		statusCode, response, header := oneRoster.MakeRosterRequest(url)
		totalRowsCount, _ := strconv.Atoi(header.Get("x-total-count"))
		b := []byte(response)

		// If status_code is 200, create array of users from response, otherwise return error
		if statusCode == 200 {

			var classesResponse models.ClassesResponse
			json.Unmarshal(b, &classesResponse)
			orClasses = classesResponse.Classes

		} else if statusCode == 401 {
			return fmt.Errorf("Unauthorized Request: %s", response)
		} else if statusCode == 404 {
			return fmt.Errorf("Not found: %s", response)
		} else if statusCode == 500 {
			return fmt.Errorf("Server Error: %s", response)
		}

		orClassesToEdit := []models.ORClass{}
		orClassIDsToDelete := []string{}
		for _, orClass := range orClasses {
			orClass.SchoolSourcedID = orClass.School.SourcedID
			orClass.CourseSourcedID = orClass.Course.SourcedID
			termsIds := []string{}
			for _, term := range orClass.Terms {
				termsIds = append(termsIds, term.SourcedID)
			}
			termsIdsString := strings.Join(termsIds, ",")
			orClass.TermSourcedIds = termsIdsString

			if strings.ToLower(orClass.Status) == strings.ToLower(models.StatusTypeActive) {
				orClassesToEdit = append(orClassesToEdit, orClass)
			} else if strings.ToLower(orClass.Status) == strings.ToLower(models.StatusTypeToBeDeleted) {
				orClassIDsToDelete = append(orClassIDsToDelete, orClass.SourcedID)
			}

		}

		// Add or Edit  Classes
		if len(orClassesToEdit) > 0 {
			err := orProcess.HandleAddOrEditClass(orClassesToEdit, districtIDs)
			if err != nil {
				return err
			}
		}

		// Delete Classes
		if len(orClassIDsToDelete) > 0 {
			err := orProcess.HandleDeleteClasses(orClassIDsToDelete, districtIDs)
			if err != nil {
				return err
			}
		}
		if totalRowsCount > (offset + limit) {
			offset = offset + 100
		} else {
			hasNext = false
			break
		}

	}

	return nil
}

// get the Users from the API and call the interface methods to Handle the response
func ProcessUsersAPI(domain string, key, secret string, orProcess models.ORProcess, districtIDs []string) error {

	var orUsers []models.ORUser
	// call the api
	limit := 100
	offset := 0
	hasNext := true
	for hasNext == true {
		url := fmt.Sprintf("%s/ims/oneroster/v1p1/users?limit=%s&offset=%s", domain, strconv.Itoa(limit), strconv.Itoa(offset))

		oneRoster := oauth1.OneRosterNew(key, secret)
		statusCode, response, header := oneRoster.MakeRosterRequest(url)
		totalRowsCount, _ := strconv.Atoi(header.Get("x-total-count"))
		b := []byte(response)

		// If status_code is 200, create array of users from response, otherwise return error
		if statusCode == 200 {

			var usersResponse models.UsersResponse
			json.Unmarshal(b, &usersResponse)
			orUsers = usersResponse.Users

		} else if statusCode == 401 {
			return fmt.Errorf("Unauthorized Request: %s", response)
		} else if statusCode == 404 {
			return fmt.Errorf("Not found: %s", response)
		} else if statusCode == 500 {
			return fmt.Errorf("Server Error: %s", response)
		}

		orUsersToEdit := []models.ORUser{}
		orUsersIDsToDelete := []string{}
		for _, orUser := range orUsers {

			// collect orgsId and add it in orUser.OrgSourcedIds
			orgsIds := []string{}
			for _, org := range orUser.Orgs {
				orgsIds = append(orgsIds, org.SourcedID)
			}
			orgsIdsString := strings.Join(orgsIds, ",")
			orUser.OrgSourcedIds = orgsIdsString

			// collect usersids and add it in orUser.UserIds
			userIds := []string{}
			for _, iden := range orUser.UserIdsIdentifer {
				userIds = append(userIds, iden.Identifier)
			}
			userIdsString := strings.Join(userIds, ",")
			orUser.UserIds = userIdsString

			// collect agentSourcedIds and add it in orUser.AgentSourcedIds
			agentSourcedIds := []string{}
			for _, agent := range orUser.Agents {
				agentSourcedIds = append(agentSourcedIds, agent.SourcedID)
			}
			agentSourcedIdsString := strings.Join(agentSourcedIds, ",")
			orUser.AgentSourcedIds = agentSourcedIdsString

			if strings.ToLower(orUser.Status) == strings.ToLower(models.StatusTypeActive) {
				orUsersToEdit = append(orUsersToEdit, orUser)
			} else if strings.ToLower(orUser.Status) == strings.ToLower(models.StatusTypeToBeDeleted) {
				orUsersIDsToDelete = append(orUsersIDsToDelete, orUser.SourcedID)
			}
		}

		// Add or Edit Users
		if len(orUsersToEdit) > 0 {
			err := orProcess.HandleAddOrEditUsers(orUsersToEdit, districtIDs)
			if err != nil {
				return err
			}
		}

		//Delete Users
		if len(orUsersIDsToDelete) > 0 {
			err := orProcess.HandleDeleteUsers(orUsersIDsToDelete, districtIDs)
			if err != nil {
				return err
			}
		}
		if totalRowsCount > (offset + limit) {
			offset = offset + 100
		} else {
			hasNext = false
			break
		}

	}

	return nil
}

// get the Enrollments from the API and call the interface methods to Handle the response
func ProcessEntrollmentAPI(domain string, key, secret string, orProcess models.ORProcess, districtIDs []string) error {

	var orEnrollments []models.OREnrollment
	// call the api
	limit := 100
	offset := 0
	hasNext := true
	for hasNext == true {
		url := fmt.Sprintf("%s/ims/oneroster/v1p1/enrollments?limit=%s&offset=%s", domain, strconv.Itoa(limit), strconv.Itoa(offset))

		oneRoster := oauth1.OneRosterNew(key, secret)
		statusCode, response, header := oneRoster.MakeRosterRequest(url)
		totalRowsCount, _ := strconv.Atoi(header.Get("x-total-count"))

		b := []byte(response)

		// If status_code is 200, create array of users from response, otherwise return error
		if statusCode == 200 {

			var enrollmentsResponse models.EnrollmentsResponse
			json.Unmarshal(b, &enrollmentsResponse)
			orEnrollments = enrollmentsResponse.Enrollments

		} else if statusCode == 401 {
			return fmt.Errorf("Unauthorized Request: %s", response)
		} else if statusCode == 404 {
			return fmt.Errorf("Not found: %s", response)
		} else if statusCode == 500 {
			return fmt.Errorf("Server Error: %s", response)
		}

		orEntrollmentsToEdit := []models.OREnrollment{}
		orEntrollmentsIDsToDelete := []models.OREnrollment{}
		for _, orEnrollment := range orEnrollments {
			orEnrollment.ClassSourcedID = orEnrollment.Class.SourcedID
			orEnrollment.SchoolSourcedID = orEnrollment.School.SourcedID
			orEnrollment.UserSourcedID = orEnrollment.User.SourcedID

			if strings.ToLower(orEnrollment.Status) == strings.ToLower(models.StatusTypeActive) {
				orEntrollmentsToEdit = append(orEntrollmentsToEdit, orEnrollment)
			} else if strings.ToLower(orEnrollment.Status) == strings.ToLower(models.StatusTypeToBeDeleted) {
				orEntrollmentsIDsToDelete = append(orEntrollmentsIDsToDelete, orEnrollment)
			}

		}

		//Add or Edit Enrollments
		if len(orEntrollmentsToEdit) > 0 {
			err := orProcess.HandleDeleteEnrollments(orEntrollmentsIDsToDelete, districtIDs)
			if err != nil {
				return err
			}
		}

		// Delete Enrollments
		if len(orEntrollmentsToEdit) > 0 {
			err := orProcess.HandleAddOrEditEnrollments(orEntrollmentsToEdit, districtIDs)
			if err != nil {
				return err
			}
		}

		if totalRowsCount > (offset + limit) {
			offset = offset + 100
		} else {
			hasNext = false
			break
		}
	}

	return nil
}

func ProcessDemographicsAPI(domain string, key, secret string, orProcess models.ORProcess) error {

	// var orDemographics []models.ORDemographics
	// call the api
	// url := fmt.Sprintf("%s/ims/oneroster/v1p1/demographics", domain)
	// client := http.DefaultClient
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	return  err
	// }
	// req.Header.Add("Content-Type", "application/json")
	// // req.Header.Add("Authorization", "Bearer "+token)

	// resp, err := client.Do(req)
	// if err != nil {
	// 	return err
	// }
	// if resp.StatusCode != 200 {
	// 	b, _ := ioutil.ReadAll(resp.Body)
	// 	fmt.Println(string(b))
	// 	return  errors.New("Status:" + resp.Status)
	// }
	// respBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	// respBytes, err := Createrequest(key, secret, "GET", url)
	// if err != nil {
	// 	return err
	// }

	// json.Unmarshal(respBytes, &orDemographics)

	// add, edit and delete orDemographics code  (we don't use it right now, maybe later we'll need it )
	// fmt.Println(">> ProcessDemographics: ", ur)

	return nil
}
