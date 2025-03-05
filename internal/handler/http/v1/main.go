package httpRouterV1

import (
	parserPage "github.com/FazylovAsylkhan/html-aggregator/internal/parser/html"
	baspana "github.com/FazylovAsylkhan/html-aggregator/internal/usecase/baspana_market"
	"github.com/go-chi/chi"
)

type HttpRouterV1 struct {
	parser *parserPage.ParserPage
	baspana *baspana.Baspana
}

func Init() *chi.Mux {	
	var h = HttpRouterV1 {
		parser: parserPage.Init(),
		baspana: baspana.Init(),
	}
	router := chi.NewRouter()

	router.Get("/load", h.handlerLoadPosts)

	return router
}
