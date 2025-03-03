package parserPage

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func (p *ParserPage) GetLinks(pageNumber int, waitFor int) ([]Link, error) {
	var htmlElement string
	links := []Link{}
	if pageNumber == 1 {
		err := chromedp.Run(p.ctx,
			chromedp.OuterHTML(".mainContentPool", &htmlElement),
		)
		if err != nil {
			return nil, fmt.Errorf("error getting outer HTML on page %v: %v", pageNumber, err)
		}
	} else {
		err := chromedp.Run(p.ctx,
			chromedp.Click(fmt.Sprintf(`//div[@class='pool-templates']//a[text()='%v']`, pageNumber), chromedp.NodeVisible),
			chromedp.WaitVisible(".mainContentPool", chromedp.BySearch),
			chromedp.OuterHTML(".mainContentPool", &htmlElement),
		)
		if err != nil {
			return nil, fmt.Errorf("error clicking on page number %v: %v", pageNumber, err)
		}
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlElement))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML from page %v: %v", pageNumber, err)
	}
	links = p.FindLinks(doc)
	return links, nil
}