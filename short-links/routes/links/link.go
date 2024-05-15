package storage

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afonsopc/omelhorsite-short-links/utils"
)

type createLinkResponse struct {
	Link utils.Link `json:"link"`
}

type createLinkPayload struct {
	ForwardUrl string `json:"forwardUrl"`
}

type getLinkParams struct {
	LinkID string `json:"id"`
}

type getLinkResponse struct {
	Link utils.Link `json:"link"`
}

type deleteLinkParams struct {
	LinkID string `json:"id"`
}

func CreateLinkHandler(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetUserIdFromHeaders(r)

	var requestBody createLinkPayload

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, utils.ErrorInvalidJSONPayload, http.StatusBadRequest)
		return
	}

	if !utils.IsUrl(requestBody.ForwardUrl) {
		http.Error(w, utils.ErrorInvalidJSONPayload, http.StatusBadRequest)
		return
	}

	link, err := utils.CreateLink(requestBody.ForwardUrl, userID)

	if err != nil {
		http.Error(w, utils.ErrorCreatingLink, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(createLinkResponse{Link: link})

	if err != nil {
		http.Error(w, utils.ErrorJSONConvertionError, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(response))
}

func GetLinkHandler(w http.ResponseWriter, r *http.Request) {
	params, err := getGetLinkParams(r)

	if err != nil {
		return
	}

	link, err := utils.GetLinkInfo(params.LinkID)

	if err != nil {
		http.Error(w, utils.ErrorLinkNotFound, http.StatusNotFound)
		return
	}

	response, err := json.Marshal(getLinkResponse{Link: link})

	if err != nil {
		http.Error(w, utils.ErrorJSONConvertionError, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(response))
}

func DeleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	params, err := getDeleteLinkParams(r)

	if err != nil {
		return
	}

	userID := utils.GetUserIdFromHeaders(r)

	link, err := utils.GetLinkInfo(params.LinkID)

	if err != nil {
		http.Error(w, utils.ErrorLinkNotFound, http.StatusNotFound)
		return
	}

	if link.UserID != userID && !utils.IsUserAdmin(r) {
		http.Error(w, utils.ErrorUserNotAllowed, http.StatusForbidden)
		return
	}

	err = utils.DeleteLink(params.LinkID)

	if err != nil {
		http.Error(w, utils.ErrorDeletingLink, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, http.StatusNoContent)
}

func getDeleteLinkParams(r *http.Request) (deleteLinkParams, error) {
	linkID, _ := utils.GetUrlQueryParamFromRequest("id", r)

	return deleteLinkParams{
		LinkID: linkID,
	}, nil
}

func getGetLinkParams(r *http.Request) (getLinkParams, error) {
	linkID, _ := utils.GetUrlQueryParamFromRequest("id", r)

	return getLinkParams{
		LinkID: linkID,
	}, nil
}
