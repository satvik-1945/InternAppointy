// package main
package main

import(
	"log"
	"net/http"
	"time"
	// "reflect"
	"io/ioutil"
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

func Body()


func main() {
    url := "http://api.internship.appointy.com:8000/v1/availability?"

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

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes:", err)
    }
    log.Println(string([]byte(body)))
}