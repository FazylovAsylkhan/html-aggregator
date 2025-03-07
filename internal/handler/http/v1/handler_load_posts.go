package httpRouterV1

import (
	"net/http"
)

func (h HttpRouterV1) handlerLoadPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.baspana.LoadPosts(r.Context())
	if err != nil {
		RespondWithError(w, 500, err.Error())
	}
	RespondWithJSON(w, 200, posts)
}
