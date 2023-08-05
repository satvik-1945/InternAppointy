// package main
package main

import(
	"log"
	"net/http"
	"time"
	// "reflect"
	"io"
// entities
)
type Resource struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BusinessHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type BlockHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type Appointment struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type Duration struct {
	Seconds int64 `json:"seconds"`
}

// endpoint request structs

type ListBusinessHoursRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListBlockHoursRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListAppointmentRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

// helper functions

func TimeToString(tm time.Time) string {
	return tm.Format(time.RFC3339)
}

func StringToTime(timeStr string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

// func Body()


//fetch the data from the api
func helper1() {

    url := "http://api.internship.appointy.com:8000/v1/durations?"

    // Create a Bearer string by appending string access token
    var bearer =  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5fQ.b_yAlhv8FMhXexY_zXdth3OFlESUVnp505kOlVgFmRg"

    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		// if u want to read the body many time
		// u need to restore 
		// reader := io.NopCloser(bytes.NewReader(bodyBytes))
		if err != nil {
			log.Fatal(err)
		}
		ans := string(bodyBytes)
		// log.Info(bodyString)
		return ans

	}
}
func helper2() {

    url := "http://api.internship.appointy.com:8000/v1/resources?"

    // Create a Bearer string by appending string access token
    var bearer =  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5fQ.b_yAlhv8FMhXexY_zXdth3OFlESUVnp505kOlVgFmRg"

    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		// if u want to read the body many time
		// u need to restore 
		// reader := io.NopCloser(bytes.NewReader(bodyBytes))
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		// log.Info(bodyString)
		println(bodyString)
	}
}
func helper3() {

    url := "http://api.internship.appointy.com:8000/v1/business-hours?"

    // Create a Bearer string by appending string access token
    var bearer =  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5fQ.b_yAlhv8FMhXexY_zXdth3OFlESUVnp505kOlVgFmRg"

    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		// if u want to read the body many time
		// u need to restore 
		// reader := io.NopCloser(bytes.NewReader(bodyBytes))
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		// log.Info(bodyString)
		println(bodyString)
	}
}
func helper4() {

    url := "http://api.internship.appointy.com:8000/v1/block-hours?"

    // Create a Bearer string by appending string access token
    var bearer =  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5fQ.b_yAlhv8FMhXexY_zXdth3OFlESUVnp505kOlVgFmRg"

    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		// if u want to read the body many time
		// u need to restore 
		// reader := io.NopCloser(bytes.NewReader(bodyBytes))
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		// log.Info(bodyString)
		println(bodyString)
	}
}
func helper5() {

    url := "http://api.internship.appointy.com:8000/v1/appointments?"

    // Create a Bearer string by appending string access token
    var bearer =  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5fQ.b_yAlhv8FMhXexY_zXdth3OFlESUVnp505kOlVgFmRg"

    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		// if u want to read the body many time
		// u need to restore 
		// reader := io.NopCloser(bytes.NewReader(bodyBytes))
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		// log.Info(bodyString)
		println(bodyString)
	}
}

func main(){
	

}