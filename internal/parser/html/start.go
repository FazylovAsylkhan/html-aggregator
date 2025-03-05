package htmlParser

import (
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func (p *ParserPage) Start(url string, waitFor int) error {
	p.log.Infof("Parsing page %v started, wait %v sec.", url, waitFor)
	err := chromedp.Run(p.Ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return fmt.Errorf("error on parsing %v: %v", url, err)
	}
	return  nil
}