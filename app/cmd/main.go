package main

import (
	"github.com/FazylovAsylkhan/html-aggregator/cmd/server"
	"github.com/FazylovAsylkhan/html-aggregator/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Get()
	server.Start(cfg)	
}