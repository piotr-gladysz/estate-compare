package crawler

import (
	"github.com/tebeka/selenium"
	"time"
)

type MatchType int

const (
	CrawlerNotMatch  MatchType = 0
	CrawlerMatchPage MatchType = 1
	CrawlerMatchList MatchType = 2
)

type PageCrawler interface {
	CrawlOffer(selenium.WebDriver, string) (*Offer, error)
}

type ListCrawler interface {
	GetUrls(selenium.WebDriver, string) ([]string, error)
	NextPage(selenium.WebDriver, string) (string, error)
}

type Factory interface {
	NewPageCrawler() PageCrawler
	NewListCrawler() ListCrawler
	MatchUrl(string) MatchType
}

type Offer struct {
	SiteId     string
	UpdateTime time.Time

	Name           string
	Description    string
	Price          int
	Area           float32
	Rooms          int
	Floor          int
	BuildingFloors int
	Year           int
	Heating        string
	Market         string
	Window         string
	Elevator       bool
	Balcony        bool
	Media          []string
}

type FactoryRegistry struct {
	factories []Factory
}

func NewCrawlerFactoryRegistry() *FactoryRegistry {
	return &FactoryRegistry{}
}

func (r *FactoryRegistry) Register(factory Factory) {
	r.factories = append(r.factories, factory)
}

func (r *FactoryRegistry) GetCrawler(url string) (PageCrawler, ListCrawler) {
	for _, factory := range r.factories {
		switch factory.MatchUrl(url) {
		case CrawlerMatchPage:
			return factory.NewPageCrawler(), nil
		case CrawlerMatchList:
			return nil, factory.NewListCrawler()
		}
	}
	return nil, nil
}
