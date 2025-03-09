package httpRouterV1

import (
	"fmt"
	"net/http"
	"strings"
)

func (h HttpRouterV1) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	number := strings.Split(path, "/")[3]
	fmt.Println(number)
	posts, err := h.baspana.CreatePost(r, number)
	if err != nil {
		RespondWithError(w, 500, err.Error())
	}
	RespondWithJSON(w, 200, posts)
}
