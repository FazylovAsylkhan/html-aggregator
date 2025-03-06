package httpRouterV1

import (
	"github.com/FazylovAsylkhan/html-aggregator/internal/database"
	parserPage "github.com/FazylovAsylkhan/html-aggregator/internal/parser/html"
	baspana "github.com/FazylovAsylkhan/html-aggregator/internal/usecase/baspana_market"
	"github.com/go-chi/chi"
)

type HttpRouterV1 struct {
	parser *parserPage.ParserPage
	baspana *baspana.Baspana
}

func Init(db *database.Queries) *chi.Mux {	
	var h = HttpRouterV1 {
		parser: parserPage.Init(),
		baspana: baspana.Init(db),
	}
	router := chi.NewRouter()

	router.Get("/load", h.handlerLoadPosts)
	router.Get("/create-post/{number}", h.handlerCreatePost)

	return router
}
