package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var ApiKey string = "" // todo: read from config file
var ContentType = "application/json"

var potentialSuccessStatusCodes = map[int]string{
	201: "Resource created Successfully",
	200: "ok",
}

var potentialErrorStatusCodes = map[int]string{
	401: "Access token is missing or invalid",
	404: "Resource not found",
	422: "The given data was invalid",
}

func Post(url string, payload map[string]string) {
	postBody, _ := json.Marshal(payload)
	req, newRequestErr := http.NewRequest("POST", url, bytes.NewBuffer(postBody))

	if newRequestErr != nil {
		log.Fatal(newRequestErr)
	}

	response, responseErr := do(req)

	if responseErr != nil {
		log.Fatal(responseErr)
	}else {
		handleResponse(response)
	}
}

func handleResponse(response *http.Response){
	hasHttpError := potentialErrorStatusCodes[response.StatusCode] != ""

	if  hasHttpError{
		log.Fatal(response.Status)
	}else {
		println(potentialSuccessStatusCodes[response.StatusCode])
	}
}

func do(req *http.Request) (*http.Response, error){
	req.Header.Set("content-type", ContentType)
	req.Header.Set("Authorization", ApiKey)

	return http.DefaultClient.Do(req)
}