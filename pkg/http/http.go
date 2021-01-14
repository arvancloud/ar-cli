package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var ApiKey = "Apikey 9181cf76-7880-4834-b333-0dba8c4915c9" // todo: read from config file
var ContentType = "application/json"

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
	req.Header.Set("content-type", ContentType)
	req.Header.Set("Authorization", ApiKey)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
