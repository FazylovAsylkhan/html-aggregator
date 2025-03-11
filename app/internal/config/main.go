package config

import (
	"log"
	"os"
	"sync"
)

type Config struct {
	Address string
	BaseURL string
	DbURL   string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		dbURL := os.Getenv("DB_URL")
		serverAddress := os.Getenv("SERVER_ADDRESS")
		parsingBaseURL := os.Getenv("PARSING_BASE_URL")

		if dbURL == "" {
			log.Fatal("DB_URL is not set")
		}
		if serverAddress == "" {
			log.Fatal("SERVER_ADDRESS is not set")
		}
		if parsingBaseURL == "" {
			log.Fatal("PARSING_BASE_URL is not set")
		}

		cfg = Config{
			Address: serverAddress,
			BaseURL: parsingBaseURL,
			DbURL:   dbURL,
		}
	})

	return &cfg
}
