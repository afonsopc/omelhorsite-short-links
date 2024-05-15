package root

import (
	"fmt"
	"net/http"

	"github.com/afonsopc/omelhorsite-short-links/utils"
	"github.com/go-chi/chi/v5"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Eis, oh Rei Excelso. Os votos sagrados. Q'os Lusos honrados. Vêm livres, vêm livres fazer. Vêm livres fazer")
}

func RootLinkHandler(w http.ResponseWriter, r *http.Request) {
	linkID := chi.URLParam(r, "linkID")

	link, err := utils.GetLinkInfo(linkID)

	if err != nil {
		http.Error(w, utils.ErrorLinkNotFound, http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link.ForwardUrl, http.StatusMovedPermanently)
}
