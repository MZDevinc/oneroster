package orServices

import (
    "fmt"
	"github.com/MZDevinc/oneroster/models"
	// "net/http"
	// "io/ioutil"
	// "errors"
	"encoding/json"

	// "net/url"
	"github.com/MZDevinc/oneroster/oauth1"
	// "strings"

	// "strconv"
	// "sync/atomic"
	// "time"
	// "crypto/rand"
	// "encoding/binary"
) 



func ProcessAPIs(domain string, key, secret string, orProcess models.ORProcess) error {

	err := ProcessAcademicSessionsAPITest(domain, key, secret, orProcess)
	// err := ProcessOrgsAPI(domain, token, orProcess)
	// err = ProcessCoursesAPI(domain, token, orProcess)
	// err = ProcessAcademicSessionsAPI(domain, key, secret, orProcess)
	// err = ProcessClassesAPI(domain, token, orProcess)
	// err = ProcessUsersAPI(domain, token, orProcess)
	// err = ProcessEntrollmentAPI(domain, token, orProcess)

	if err != nil {
		return err
	}

	return nil;
}




func ProcessAcademicSessionsAPITest(domain string, key, secret string, orProcess models.ORProcess) error {


	var academicSessions []models.ORAcademicSessions
	// call the api 
	url := fmt.Sprintf("%s/ims/oneroster/v1p1/academicSessions", domain)

///////***** add oauth 1 params to request *****/////	
	
	// consumer := NewConsumer(key, secret, url, "GET")
	// signature, err := consumer.Sign()
	// fmt.Println(" >>> signature: ", signature,"   err: ",err)

	// using oauth1 file
	respByte, err := Createrequest(key, secret, "GET", url)
	if err != nil {
		return err
	}
	json.Unmarshal(respByte, &academicSessions)
    return nil
}


func ProcessAcademicSessionsAPI(domain string, key, secret string, orProcess models.ORProcess) error {


	var academicSessions []models.ORAcademicSessions
	// call the api 
	url := fmt.Sprintf("%s/ims/oneroster/v1p1/academicSessions", domain)
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
	respBytes, err := Createrequest(key, secret, "GET", url)
	if err != nil {
		return err
	}
	json.Unmarshal(respBytes, &academicSessions)


	orAcademicSessionToEdit := []models.ORAcademicSessions{}
	orAcademicSessionIDsToDelete := []string{}
	for _,orAcademicSession := range academicSessions {
		orAcademicSession.ParentSourcedId = orAcademicSession.Parent.SourcedId
		if orAcademicSession.Status == models.STATUS_TYPE_ACTIVE{
			orAcademicSessionToEdit = append(orAcademicSessionToEdit, orAcademicSession)
		}else if orAcademicSession.Status == models.STATUS_TYPE_TOBEDELETED{
			orAcademicSessionIDsToDelete = append(orAcademicSessionIDsToDelete, orAcademicSession.SourcedId)
		}
	}
	
	// Add or Edit AcademicSessions
	if len(orAcademicSessionToEdit) >0 {
		err := orProcess.HandleAddOrEditAcademicSessions(orAcademicSessionToEdit)
		if err != nil {
			return err
		}
	}	
	// Delete AcademicSessions
	if len(orAcademicSessionIDsToDelete) >0 {
		err := orProcess.HandleDeleteAcademicSessions(orAcademicSessionIDsToDelete)
		if err != nil {
			return err
		}
	}	
    return nil
}



func ProcessOrgsAPI(domain string, key, secret string, orProcess models.ORProcess)  error {

	var orgs []models.OROrg
	// call the api 
	url := fmt.Sprintf("%s/ims/oneroster/v1p1/orgs", domain)
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

	respBytes, err := Createrequest(key, secret, "GET", url)
	if err != nil {
		return err
	}
	json.Unmarshal(respBytes, &orgs)


	for _, org := range orgs {
		
		var err error = nil
		if org.OrgType == models.ORG_TYPE_DISTRICT {
		
			if org.Status == models.STATUS_TYPE_ACTIVE{
				err = orProcess.HandleAddOrEditDistrict(org)
			}else if org.Status == models.STATUS_TYPE_TOBEDELETED{
				err = orProcess.HandleDeleteDistrict(org)
			}

			if err != nil {
				return err
			}
			
		} else if org.OrgType == models.ORG_TYPE_SCHOOL {
			org.ParentSourcedId = org.Parent.SourcedId
			if org.Status == models.STATUS_TYPE_ACTIVE{
				err = orProcess.HandleAddOrEditSchool(org)
			}else if org.Status == models.STATUS_TYPE_TOBEDELETED{
				err = orProcess.HandleDeleteSchool(org)
			}

			if err != nil {
				return  err
			}
		}
	}

    return  nil
}



func ProcessCoursesAPI(domain string, key, secret string, orProcess models.ORProcess)  error {

	
	var orCourses []models.ORCourse
	// call the api 
	url := fmt.Sprintf("%s/ims/oneroster/v1p1/courses", domain)
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

	respBytes, err := Createrequest(key, secret, "GET", url)
	if err != nil {
		return err
	}
	json.Unmarshal(respBytes, &orCourses)

	orCourseToEdit := []models.ORCourse{}
	orCoursesIDsToDelete := []string{}
	for _,orCourse := range orCourses {
		if orCourse.Status == models.STATUS_TYPE_ACTIVE{
			// err = orProcess.HandleEditCourse(orCourse)
			orCourseToEdit = append(orCourseToEdit, orCourse)
		}else if orCourse.Status == models.STATUS_TYPE_TOBEDELETED{
			// err = orProcess.HandleDeleteCourse(orCourse)
			orCoursesIDsToDelete = append(orCoursesIDsToDelete, orCourse.SourcedId)
		}
		if err != nil {
			return err
		}
	}
	err = orProcess.HandleAddOrEditCourse(orCourseToEdit)
	err = orProcess.HandleDeleteCourses(orCoursesIDsToDelete)
	if err != nil {
		return err
	}
	
    return nil
}


func ProcessClassesAPI(domain string, key, secret string, orProcess models.ORProcess)  error {

	var orClasses []models.ORClass
	// call the api 
	url := fmt.Sprintf("%s/ims/oneroster/v1p1/classes", domain)
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

	respBytes, err := Createrequest(key, secret, "GET", url)
	if err != nil {
		return err
	}

	json.Unmarshal(respBytes, &orClasses)
	
	orClassesToEdit := []models.ORClass{}
	orClassIDsToDelete := []string{}
	for _,orClass := range orClasses {
		if orClass.Status == models.STATUS_TYPE_ACTIVE{
			orClassesToEdit = append(orClassesToEdit, orClass)
		}else if orClass.Status == models.STATUS_TYPE_TOBEDELETED{
			orClassIDsToDelete = append(orClassIDsToDelete, orClass.SourcedId)
		}
		if err != nil {
			return err
		}
	}
	err = orProcess.HandleAddOrEditClass(orClassesToEdit)
	err = orProcess.HandleDeleteClasses(orClassIDsToDelete)
	if err != nil {
		return err
	}
	
    return nil
}

func ProcessUsersAPI(domain string, key, secret string, orProcess models.ORProcess)  error {

	var orUsers []models.ORUser
	// call the api 
	url := fmt.Sprintf("%s/ims/oneroster/v1p1/users", domain)
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

	respBytes, err := Createrequest(key, secret, "GET", url)
	if err != nil {
		return err
	}

	json.Unmarshal(respBytes, &orUsers)


	orUsersToEdit := []models.ORUser{}
	orUsersIDsToDelete := []string{}
	for _,orUser := range orUsers {
		if orUser.Status == models.STATUS_TYPE_ACTIVE{
			orUsersToEdit = append(orUsersToEdit, orUser)
		}else if orUser.Status == models.STATUS_TYPE_TOBEDELETED{
			orUsersIDsToDelete = append(orUsersIDsToDelete, orUser.SourcedId)
		}
		if err != nil {
			return err
		}
	}
	err = orProcess.HandleAddOrEditUsers(orUsersToEdit)
	err = orProcess.HandleDeleteUsers(orUsersIDsToDelete)
	if err != nil {
		return err
	}
	

    return nil
}


func ProcessEntrollmentAPI(domain string, key, secret string, orProcess models.ORProcess)  error {

	var orEnrollments []models.OREnrollment
	// call the api 
	url := fmt.Sprintf("%s/ims/oneroster/v1p1/enrollments", domain)
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

	respBytes, err := Createrequest(key, secret, "GET", url)
	if err != nil {
		return err
	}

	json.Unmarshal(respBytes, &orEnrollments)

		
	orEntrollmentsToEdit := []models.OREnrollment{}
	orEntrollmentsIDsToDelete := []models.OREnrollment{}
	for _,orEnrollment := range orEnrollments {
		if orEnrollment.Status == models.STATUS_TYPE_ACTIVE{
			orEntrollmentsToEdit = append(orEntrollmentsToEdit, orEnrollment)
		}else if orEnrollment.Status == models.STATUS_TYPE_TOBEDELETED{
			orEntrollmentsIDsToDelete = append(orEntrollmentsIDsToDelete, orEnrollment)
		}
		if err != nil {
			return err
		}
	}
	err = orProcess.HandleDeleteEnrollments(orEntrollmentsIDsToDelete)
	err = orProcess.HandleAddOrEditEnrollments(orEntrollmentsToEdit)
	if err != nil {
		return err
	}

    return nil
}

func ProcessDemographicsAPI(domain string, key, secret string, orProcess models.ORProcess)  error {

	var orDemographics []models.ORDemographics
	// call the api 
	url := fmt.Sprintf("%s/ims/oneroster/v1p1/demographics", domain)
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

	respBytes, err := Createrequest(key, secret, "GET", url)
	if err != nil {
		return err
	}

	json.Unmarshal(respBytes, &orDemographics)

	// add, edit and delete orDemographics code  (we don't use it right now, maybe later we'll need it )
	fmt.Println(">> ProcessDemographics")
	

    return nil
}


func Createrequest(key, secret, method, url string) ([]byte, error){
	oauthParams := oauth1.OAuthParameters{}
	oauthParams.ConsumerKey = &key
	oauthParams.ConsumerSecret = &secret
	v := "1.0"
	oauthParams.Version = &v
	sig := oauth1.GetHMACSigner(secret, "")
	oauthParams.Signer = sig
	signerMethod := oauthParams.Signer.GetMethod()
	oauthParams.Method = &signerMethod
	oauthParams.Build()
	// signature, err := oauthParams.GetOAuthSignature("GET", url, nil)
	// fmt.Println(" >>> signature2: ", signature,"   err2: ",err)

	return oauthParams.DoOauthRequestTest(method, url, nil)
}

/////// Oauth 1 /////////


