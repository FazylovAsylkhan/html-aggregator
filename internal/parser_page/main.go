package parserPage

import (
	"context"

	"github.com/FazylovAsylkhan/html-aggregator/internal/logger"
	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
)

type Link struct {
	Date string
	Title string
	Cost string
	Address string
	Link string
	Img string
	Id string
	Count string
}

type Store struct {
	unsortedLinks []Link
}

type ParserPage struct {
	ctx context.Context
	cancel context.CancelFunc
	log *logrus.Logger
	store Store
}

func Init() *ParserPage {
	ctx, cancel := chromedp.NewContext(context.Background())
	log := logger.New()
	log.SetFormatter(&logger.GeneralFormatter{})

	parser := ParserPage{
		ctx: ctx,
		cancel: cancel,
		log: log,
	}
	return &parser
}



func (p *ParserPage) SaveLinks(links []Link) {
	p.store.unsortedLinks = links
}
func (p *ParserPage) GetSavedLinks() []Link{
	return p.store.unsortedLinks
}
