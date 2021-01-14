package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ebrahimahmadi/ar-cli/pkg/config"
	"log"
	"net/http"
)

var Config = config.GetConfigInfo()

func Get(url string, queries map[string]string) (*http.Response, error) {
	req := newRequest("GET", url, nil)

	if queries != nil {
		for key, value := range queries {
			addQueryToUrl(*req, key, value)
		}
	}

	response, responseErr := do(req)

	if responseErr != nil {
		return nil, responseErr
	} else {
		return handleResponse(response)
	}
}

func Delete(url string, payload map[string]string, queries map[string]string) (*http.Response, error) {
	req := newRequest("DELETE", url, payload)

	if queries != nil {
		for key, value := range queries {
			addQueryToUrl(*req, key, value)
		}
	}

	response, responseErr := do(req)

	if responseErr != nil {
		return nil, responseErr
	} else {
		return handleResponse(response)
	}
}

func Post(url string, payload map[string]string) (*http.Response, error) {
	req := newRequest("POST", url, payload)

	response, responseErr := do(req)

	if responseErr != nil {
		return nil, responseErr
	} else {
		return handleResponse(response)
	}
}


func Put(url string, payload map[string]string) (*http.Response, error) {
	req := newRequest("PUT", url, payload)

	response, responseErr := do(req)

	if responseErr != nil {
		return nil, responseErr
	} else {
		return handleResponse(response)
	}
}


func addQueryToUrl(request http.Request, key string, value string) {
	query := request.URL.Query()

	query.Add(key, value)

	request.URL.RawQuery = query.Encode()
}

func newRequest(method string, url string, payload map[string]string) *http.Request {
	postBody, _ := json.Marshal(payload)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(postBody))

	if err != nil {
		log.Fatal(err)
	}

	return req
}

func handleResponse(response *http.Response) (*http.Response, error) {
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300

	if !statusOK {
		return nil, errors.New(response.Status)
	} else {
		return response, nil
	}
}

func do(req *http.Request) (*http.Response, error) {
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", Config.GetApiKey())

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
