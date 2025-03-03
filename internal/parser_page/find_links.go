package parserPage

import "github.com/PuerkitoBio/goquery"

func (p *ParserPage) FindLinks(doc *goquery.Document) []Link {
	var links []Link
	link := Link{}
	doc.Find(".card").Each(func(i int, s *goquery.Selection) {
		s.Find(".districtTitle").Each(func(i int, s *goquery.Selection) {
			linkElement := s.Find("a")
			href, exists := linkElement.Attr("href")
			if exists {
				link.Link = href
				link.Title = linkElement.Text()
			}

		})
		s.Find("p.card-text.pb-3.d-flex.justify-content-between span:nth-child(2)").Each(func(i int, s *goquery.Selection) {
			link.Date = s.Text()
		})
		link.Address = s.Find("p.card-text.pool-adress").First().Text()
		link.Cost = s.Find("h3.ob__favorits__squares span.font-weight-bold").Text()
		link.Id = s.Find("div.card-text.pool--mt span.pool-code").First().Text()
		link.Count = s.Find("div.card-text.pool--mt span.pool-code").Eq(1).Text()
		imageSrc, exists := doc.Find("img.d-block").Attr("src")
		if exists {
			link.Img = imageSrc
		}
	})
	links = append(links, link)

	return links
}