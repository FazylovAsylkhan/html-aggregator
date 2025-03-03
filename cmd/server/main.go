package server

import (
	"net/http"

	"github.com/FazylovAsylkhan/html-aggregator/internal/config"
	httpHandler "github.com/FazylovAsylkhan/html-aggregator/internal/handler/http"
	"github.com/FazylovAsylkhan/html-aggregator/internal/logger"
)

func Start(cfg *config.Config) error{
	log := logger.New()
	log.SetFormatter(&logger.GeneralFormatter{})
	
	router := httpHandler.Init()
	srv := &http.Server{
		Handler: router,
		Addr:    cfg.Address,
	}
	log.Infof("Starting server at address %s with base URL %s", cfg.Address, cfg.BaseURL)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	return nil
}