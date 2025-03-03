package parserPage

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func (p *ParserPage) ParseLinks(url string, waitFor int) ([]Link, error) {
	var arrLinks = []Link{}
	var lastPageStr string
	var lastPage int

	p.log.Infof("Parsing page %v started, wait %v sec.", url, waitFor)
	err := chromedp.Run(p.ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible(`.pool-templates:last-of-type`),
		chromedp.Text(`.pool-templates:last-of-type`, &lastPageStr),
	)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	lastPage, err = strconv.Atoi(lastPageStr)
	if err != nil {
		return nil, fmt.Errorf("error converting lastPage to int: %v", err)
	}

	defer p.cancel()
	if err != nil {
		return nil, fmt.Errorf("error on parsing %v: %v", url, err)
	}

	for i := 1; i <= lastPage; i++ {
		p.log.Infof("Parsing page %v started, wait %v sec.", i, waitFor)
		links, err := p.GetLinks(i, waitFor)
		if err != nil {
			return nil, fmt.Errorf("error getting links on page %v: %v", i, err)
		}
		arrLinks = append(arrLinks, links...)
	}

	return arrLinks, nil
}