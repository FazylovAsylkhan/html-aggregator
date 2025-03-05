package htmlParser

import (
	"fmt"

	"github.com/chromedp/chromedp"
)

func (p *ParserPage) ParseNestedPages(pageNumber int, selWrapper, selBtn string) (string, error) {
	var htmlElement string
	if pageNumber == 1 {
		err := chromedp.Run(p.Ctx,
			chromedp.OuterHTML(selWrapper, &htmlElement),
		)
		if err != nil {
			return "", fmt.Errorf("error getting outer HTML on page %v: %v", pageNumber, err)
		}
	} else {
		err := chromedp.Run(p.Ctx,
			chromedp.Click(selBtn, chromedp.NodeVisible),
			chromedp.WaitVisible(selWrapper, chromedp.BySearch),
			chromedp.OuterHTML(selWrapper, &htmlElement),
		)
		if err != nil {
			return "", fmt.Errorf("error clicking on page number %v: %v", pageNumber, err)
		}
	}
	return htmlElement, nil
}