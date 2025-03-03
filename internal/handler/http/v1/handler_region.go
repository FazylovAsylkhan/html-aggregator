package httpRouterV1

import (
	"net/http"
)

func (h HttpRouterV1) handlerLoad(w http.ResponseWriter, r *http.Request) {
	links, err := h.parser.ParseLinks("https://baspana.otbasybank.kz/pool/search", 3)
	if err != nil {
		RespondWithError(w, 500, err.Error())
	}
	h.parser.SaveLinks(links)
	RespondWithJSON(w, 200, links)
}
