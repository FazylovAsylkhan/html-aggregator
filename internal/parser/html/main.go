package htmlParser

import (
	"context"

	"github.com/FazylovAsylkhan/html-aggregator/internal/logger"
	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
)

type ParserPage struct {
	Ctx context.Context
	Cancel context.CancelFunc
	log *logrus.Logger
}

func Init() *ParserPage {
	ctx, cancel := chromedp.NewContext(context.Background())
	log := logger.New()
	log.SetFormatter(&logger.GeneralFormatter{})

	parser := ParserPage{
		Ctx: ctx,
		Cancel: cancel,
		log: log,
	}
	return &parser
}