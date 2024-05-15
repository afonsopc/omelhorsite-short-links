package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"unicode/utf8"
)

type Response struct {
	ID string `json:"id"`
}

func getUserIdFromToken(token string) string {
	accountsServiceEndpoint := GetAccountsServiceConfiguration().Endpoint

	url := accountsServiceEndpoint + "/account?info_to_get[id]=true"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	ThrowIfError(err)

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	ThrowIfError(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	ThrowIfError(err)

	var response Response
	err = json.Unmarshal(body, &response)
	ThrowIfError(err)

	return response.ID
}

func getIsAdminFromToken(token string) bool {
	accountsServiceEndpoint := GetAccountsServiceConfiguration().Endpoint

	url := accountsServiceEndpoint + "/admin"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	ThrowIfError(err)

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	ThrowIfError(err)

	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func getTokenFromHeaders(r *http.Request) string {
	header := r.Header.Get("Authorization")

	if utf8.RuneCountInString(header) > utf8.RuneCountInString("Bearer ") {
		return header[utf8.RuneCountInString("Bearer "):]
	}

	return ""
}

func GetUserIdFromHeaders(r *http.Request) string {
	token := getTokenFromHeaders(r)

	if token == "" {
		return ""
	}

	return getUserIdFromToken(token)
}

func IsUserAdmin(r *http.Request) bool {
	token := getTokenFromHeaders(r)

	if token == "" {
		return false
	}

	return getIsAdminFromToken(token)
}
