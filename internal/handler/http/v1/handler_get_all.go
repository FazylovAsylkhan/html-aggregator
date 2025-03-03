package httpRouterV1

import "net/http"

func (h HttpRouterV1) handlerGetAll(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, h.parser.GetSavedLinks())
}