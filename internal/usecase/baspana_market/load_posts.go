package baspana

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Post struct {
	Date    string
	Title   string
	Cost    string
	Address string
	Link    string
	Img     string
	Id      string
	Count   string
}

func (b *Baspana) LoadPosts() ([]Post, error) {
	err := b.parser.Start("https://baspana.otbasybank.kz/pool/search", 3)
	if err != nil {
		return nil, fmt.Errorf("starting parser error: %v", err)
	}
	defer b.parser.Cancel()

	text, err := b.parser.GetText(`.pool-templates:last-of-type`)
	if err != nil {
		return nil, fmt.Errorf("starting parser error: %v", err)
	}
	lastPage, err := strconv.Atoi(text)
	if err != nil {
		return nil, fmt.Errorf("error converting lastPage to int: %v", err)
	}
	
	posts := make([]Post, lastPage)
	for i := 1; i <= lastPage; i++ {
		b.log.Infof("Parsing page %v started", i)
		selBtn := fmt.Sprintf(`//div[@class='pool-templates']//a[text()='%v']`, i)
		htmlElement, err := b.parser.ParseNestedPages(i, ".mainContentPool", selBtn)
		if err != nil {
			return nil, fmt.Errorf("error parsing HTML from page %v: %v", i, err)
		}
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlElement))
		if err != nil {
			return nil, fmt.Errorf("error parsing HTML from page %v: %v", i, err)
		}
		posts[i-1] = getPost(doc)
	}
	b.bufferPosts = &posts

	return posts, nil
}

func getPost(doc *goquery.Document) Post {
	post := Post{}
	doc.Find(".card").Each(func(i int, s *goquery.Selection) {
		s.Find(".districtTitle").Each(func(i int, s *goquery.Selection) {
			linkElement := s.Find("a")
			href, exists := linkElement.Attr("href")
			if exists {
				post.Link = href
				post.Title = linkElement.Text()
			}

		})
		s.Find("p.card-text.pb-3.d-flex.justify-content-between span:nth-child(2)").Each(func(i int, s *goquery.Selection) {
			post.Date = s.Text()
		})
		post.Address = s.Find("p.card-text.pool-adress").First().Text()
		post.Cost = s.Find("h3.ob__favorits__squares span.font-weight-bold").Text()
		post.Id = s.Find("div.card-text.pool--mt span.pool-code").First().Text()
		post.Count = s.Find("div.card-text.pool--mt span.pool-code").Eq(1).Text()
		imageSrc, exists := doc.Find("img.d-block").Attr("src")
		if exists {
			post.Img = imageSrc
		}
	})

	return post
}