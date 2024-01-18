package auth

import (
	"errors"
	"net/http"
	"strings"
)

//extracts an API key from the headers of a HTTP request
//format of API key viz: `Authorisation: ApiKey <string>`
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val=="" {
		return "", errors.New("could not find auth info")
	}

	vals := strings.Split(val, "")
	if len(vals)!=2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0]!="ApiKey" {
		return "", errors.New("first part of auth header is malformed")
	}

	return vals[1], nil
}