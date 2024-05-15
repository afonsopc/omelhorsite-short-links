package utils

import (
	"fmt"
	"net/http"
)

func GetUrlQueryParamFromRequest(param string, r *http.Request) (string, error) {
	paramValue := r.URL.Query().Get(param)

	if paramValue == "" {
		return "", fmt.Errorf("the \"%s\" url query param needs to be specified", param)
	}

	return paramValue, nil
}
