package validator

import (
	"errors"
	"regexp"
)

func IsApiKey(apiKey string) (bool, error) {
	var validApiKey = regexp.MustCompile(`^Apikey [a-z0-9\-]+$$`)
	if !validApiKey.MatchString(apiKey) {
		return false, errors.New("API token should be in format: 'Apikey xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx'")
	}
	return true, nil
}

func IsDomain(domain string) (bool, error) {
	var validApiKey = regexp.MustCompile(`(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`)
	if !validApiKey.MatchString(domain) {
		return false, errors.New("domain name should be in format: 'example.com'")
	}
	return true, nil
}
