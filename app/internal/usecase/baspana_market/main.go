package baspana

import (
	"github.com/FazylovAsylkhan/html-aggregator/internal/database"
	"github.com/FazylovAsylkhan/html-aggregator/internal/logger"
	parserPage "github.com/FazylovAsylkhan/html-aggregator/internal/parser/html"
	"github.com/sirupsen/logrus"
)

type Baspana struct {
	parser *parserPage.ParserPage
	log *logrus.Logger
	bufferPosts *[]Post
	DB *database.Queries
}

func Init(db *database.Queries) *Baspana {
	log := logger.New()
	log.SetFormatter(&logger.GeneralFormatter{})
	var b = Baspana{
		parser: parserPage.Init(),
		log: log,
		DB: db,
	}

	return &b
}
