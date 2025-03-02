package main

import (
	"github.com/FazylovAsylkhan/html-aggregator/cmd/server"
	"github.com/FazylovAsylkhan/html-aggregator/internal/config"
)

func main() {
	cfg := config.Get()
	server.Start(cfg)	
}