package httpRouterV1

import (
	parserPage "github.com/FazylovAsylkhan/html-aggregator/internal/parser_page"
	"github.com/go-chi/chi"
)

type HttpRouterV1 struct {
	parser *parserPage.ParserPage
}

func Init() *chi.Mux {	
	var h = HttpRouterV1 {
		parser: parserPage.Init(),
	}
	router := chi.NewRouter()

	router.Get("/load", h.handlerLoad)
	router.Get("/all", h.handlerGetAll)
	router.Get("/region/{region}", h.handlerGetAll)

	return router
}
