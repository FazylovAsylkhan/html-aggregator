package htmlParser

import (
	"fmt"

	"github.com/chromedp/chromedp"
)

func (p *ParserPage)GetText(sel string) (string, error){
	var text string
	err := chromedp.Run(p.Ctx,
		chromedp.WaitVisible(sel),
		chromedp.Text(sel, &text),
	)
	if err != nil {
		return "", fmt.Errorf("error on getting %v: %v", sel, err)
	}

	return text, err
}