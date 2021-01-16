package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ebrahimahmadi/ar-cli/pkg/config"
	"github.com/ebrahimahmadi/ar-cli/pkg/helpers"
	"net/http"
)

var Config = config.GetConfigInfo()

var potentialErrMessages = map[int]string{
	401: "Access Token is missing or invalid",
	422: "The given data is invalid",
	404: "Resource not found",
}

type RequestBag struct {
	BodyPayload map[string]string
	URLQueries  map[string]string
	URL         string
	Method      string
}

type Test struct {
	Id string `json:"id"`
}

func (r *RequestBag) Do() (*http.Response, error) {
	b := new(bytes.Buffer)
	for key, value := range r.BodyPayload {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	body, err := json.Marshal(r.BodyPayload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	for key, value := range r.URLQueries {
		addQueryToUrl(*req, key, value)
	}

	return do(req)
}

func HandleResponseErr(response *http.Response) {
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300

	if !statusOK {
		err := helpers.ToBeColored{Expression: potentialErrMessages[response.StatusCode]}
		err.StdoutError().StopExecution()
	}
}

func addQueryToUrl(request http.Request, key string, value string) {
	query := request.URL.Query()

	query.Add(key, value)

	request.URL.RawQuery = query.Encode()
}

func do(req *http.Request) (*http.Response, error) {
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", Config.GetApiKey())

	return http.DefaultClient.Do(req)
}
