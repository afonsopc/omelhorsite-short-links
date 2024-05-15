package utils

import "log"

func ThrowIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

const (
	ErrorJSONConvertionError = "There has been a problem converting the response to JSON, we're sory 😔"
	ErrorInvalidJSONPayload  = "Invalid JSON payload"
	ErrorCreatingLink        = "Error creating link"
	ErrorLinkNotFound        = "Link not found"
	ErrorUserNotAllowed      = "User not allowed"
	ErrorDeletingLink        = "Problem deleting link"
	ErrorGettingAllLinks     = "Problem getting all links"
)
