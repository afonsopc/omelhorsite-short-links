package storage

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afonsopc/omelhorsite-short-links/utils"
)

type getAllLinksParams struct {
	userID string
}

type getAllLinksResponse struct {
	Links []utils.Link `json:"links"`
}

func GetAllLinksHandler(w http.ResponseWriter, r *http.Request) {
	params, err := getGetAllLinksParams(r)

	if err != nil {
		return
	}

	if params.userID != "" && !utils.IsUserAdmin(r) {
		http.Error(w, utils.ErrorUserNotAllowed, http.StatusForbidden)
		return
	}

	userID := utils.GetUserIdFromHeaders(r)

	links, err := utils.GetAllLinks(userID)

	if err != nil {
		http.Error(w, utils.ErrorGettingAllLinks, http.StatusNotFound)
		return
	}

	response, err := json.Marshal(getAllLinksResponse{Links: links})

	if err != nil {
		http.Error(w, utils.ErrorJSONConvertionError, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(response))
}

func getGetAllLinksParams(r *http.Request) (getAllLinksParams, error) {
	userID, _ := utils.GetUrlQueryParamFromRequest("userId", r)

	return getAllLinksParams{
		userID: userID,
	}, nil
}
