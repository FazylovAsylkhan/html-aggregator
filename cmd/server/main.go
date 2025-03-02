package server

import (
	"github.com/FazylovAsylkhan/html-aggregator/internal/config"
	"github.com/FazylovAsylkhan/html-aggregator/internal/logger"
)

func Start(cfg *config.Config) error{
	log := logger.New()
	log.SetFormatter(&logger.ServerFormatter{})
	log.Infof("Starting server at address %s with base URL %s", cfg.Address, cfg.BaseURL)

	return nil
}