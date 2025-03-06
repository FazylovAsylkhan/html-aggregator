package server

import (
	"database/sql"
	"net/http"

	"github.com/FazylovAsylkhan/html-aggregator/internal/config"
	"github.com/FazylovAsylkhan/html-aggregator/internal/database"
	httpHandler "github.com/FazylovAsylkhan/html-aggregator/internal/handler/http"
	"github.com/FazylovAsylkhan/html-aggregator/internal/logger"
)

func Start(cfg *config.Config) {
	log := logger.New()
	log.SetFormatter(&logger.GeneralFormatter{})
	
	log.Infof("Connection to DB")
	conn, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("can't connect to database: %v", err)
	}
	db := database.New(conn)
	
	router := httpHandler.Init(db)
	srv := &http.Server{
		Handler: router,
		Addr:    cfg.Address,
	}
	log.Infof("Starting server at address %s with base URL %s", cfg.Address, cfg.BaseURL)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}