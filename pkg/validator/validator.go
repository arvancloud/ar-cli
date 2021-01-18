package validator

import (
	"errors"
	"net"
	"regexp"
	"sort"
	"strings"
)

func HasString(searchTerm string, keys []string) (bool, error) {
	sort.Strings(keys)
	i := sort.SearchStrings(keys, searchTerm)

	ok := i < len(searchTerm) && keys[i] == searchTerm

	if ok {
		return true, nil
	}

	return  false, errors.New(searchTerm + " Is Illegal. Value should be one of " + strings.Join(keys,", "))
}

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

func IsValidIp(ip string) (bool, error) {
	if net.ParseIP(ip) != nil {
		return false, errors.New("not a valid IP address")
	}

	return true, nil
}