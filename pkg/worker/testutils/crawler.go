package testutils

import (
	"github.com/piotr-gladysz/estate-compare/pkg/worker/crawler"
	"github.com/tebeka/selenium"
)

// NewFactoryMock creates a new PageCrawlerMock with preinitialized crawlers
func NewFactoryMock() *FactoryMock {
	return &FactoryMock{
		ReturnPageCrawler:     &PageCrawlerMock{},
		ReturnPageListCrawler: &ListCrawlerMock{},
	}
}

type PageCrawlerMock struct {
	Callback    func(this *PageCrawlerMock, method string, args ...any)
	ReturnOffer *crawler.Offer
	ReturnError error
}

type ListCrawlerMock struct {
	Callback func(this *ListCrawlerMock, method string, args ...any)

	ReturnUrls     []string
	ReturnNextPage string
	ReturnError    error
}

type FactoryMock struct {
	Callback func(this *FactoryMock, method string, args ...any)

	ReturnPageCrawler     crawler.PageCrawler
	ReturnPageListCrawler crawler.ListCrawler
	ReturnMatchType       crawler.MatchType
}

// PageCrawlerMock implementation
// -----------------------------

func (p *PageCrawlerMock) CrawlOffer(driver selenium.WebDriver, s string) (*crawler.Offer, error) {
	if p.Callback != nil {
		p.Callback(p, "CrawlOffer", driver, s)
	}

	return p.ReturnOffer, p.ReturnError
}

// ListCrawlerMock implementation
// ----------------------------------

func (l *ListCrawlerMock) GetUrls(driver selenium.WebDriver, s string) ([]string, error) {
	if l.Callback != nil {
		l.Callback(l, "GetUrls", driver, s)
	}

	return l.ReturnUrls, l.ReturnError
}

func (l *ListCrawlerMock) NextPage(driver selenium.WebDriver, s string) (string, error) {
	if l.Callback != nil {
		l.Callback(l, "NextPage", driver, s)
	}

	return l.ReturnNextPage, l.ReturnError
}

// FactoryMock implementation
// -----------------------------

func (f FactoryMock) NewPageCrawler() crawler.PageCrawler {
	if f.Callback != nil {
		f.Callback(&f, "NewPageCrawler")
	}

	return f.ReturnPageCrawler
}

func (f FactoryMock) NewListCrawler() crawler.ListCrawler {
	if f.Callback != nil {
		f.Callback(&f, "NewListCrawler")
	}

	return f.ReturnPageListCrawler
}

func (f FactoryMock) MatchUrl(s string) crawler.MatchType {
	if f.Callback != nil {
		f.Callback(&f, "MatchUrl", s)
	}

	return f.ReturnMatchType
}
