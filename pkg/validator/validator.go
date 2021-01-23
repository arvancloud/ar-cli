package validator

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	IPv6len = 16
)

func HasString(searchTerm string, keys []string) (bool, error) {
	sort.Strings(keys)
	i := sort.SearchStrings(keys, searchTerm)

	ok := i < len(searchTerm) && keys[i] == searchTerm

	if ok {
		return true, nil
	}

	return false, errors.New(searchTerm + " Is Illegal. Value should be one of " + strings.Join(keys, ", "))
}

func HasInt(searchTerm int, keys []int) (bool, error) {
	sort.Ints(keys)
	i := sort.SearchInts(keys, searchTerm)

	ok := i < searchTerm && keys[i] == searchTerm

	if ok {
		return true, nil
	}

	return false, errors.New(strconv.Itoa(searchTerm) + " Is Illegal. Value should be one of " + strings.Trim(strings.Replace(fmt.Sprint(keys), " ", ", ", -1), "[]"))
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

func IsValidIpv4(ip string) (bool, error) {
	addr, _ := net.ResolveTCPAddr("tcp", ip)

	if addr.IP.To4() == nil {
		return false, errors.New("not a valid IP address")
	}

	return true, nil
}

func IsValidIpv6(ip string) (bool, error) {

	addr, _ := net.ResolveTCPAddr("tcp", ip)
	isIpv6 := len(ip) == IPv6len && addr.IP.To4() == nil

	if !isIpv6 {
		return false, errors.New("not a valid IPv6 address")
	}

	return true, nil
}
