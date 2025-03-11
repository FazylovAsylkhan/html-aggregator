package httpHandler

import (
	"net/http"

	"github.com/FazylovAsylkhan/html-aggregator/internal/database"
	httpRouterV1 "github.com/FazylovAsylkhan/html-aggregator/internal/handler/http/v1"
	"github.com/FazylovAsylkhan/html-aggregator/internal/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func Init(db *database.Queries) *chi.Mux {
	router := chi.NewRouter()
	options := cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}
	router.Use(cors.Handler(options))
	log := logger.New()
	log.SetFormatter(&logger.HandlerFormatter{})
	router.Use(func(next http.Handler) http.Handler {
		return logger.MiddlewareHandler(log, next)
	})
	v1Router := httpRouterV1.Init(db)
	router.Mount("/v1", v1Router)

	return router
}