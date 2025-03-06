package config

import (
	"flag"
	"sync"
)

type Config struct {
	Address    string
	BaseURL string
	DbURL string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		address := flag.String("address", "localhost:8080", "Address to listen on for the web server")
		isHttps := flag.String("https", "true", "Protocol of base url")
		baseUrl := flag.String("baseURL", "baspana.otbasybank.kz", "Base url for parsing")
		db := flag.String("db", "postgres://postgres:password@localhost:5432/html_aggregator?sslmode=disable", "link to DB")
		flag.Parse()
	
		if *isHttps == "true" {
			*baseUrl = "https://" + *baseUrl
		} else {
			*baseUrl = "http://" + *baseUrl
		}
	
		cfg = Config{
			Address:    *address,
			BaseURL: *baseUrl,
			DbURL: *db,
		}
	})

	return &cfg
}