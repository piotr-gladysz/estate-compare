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

// PageCrawler is an interface for crawling single offer page
type PageCrawler interface {

	// CrawlOffer should return Offer struct with all the data from the given url or error
	CrawlOffer(wd selenium.WebDriver, url string) (*Offer, error)
}

// ListCrawler is an interface for crawling list of offers
type ListCrawler interface {

	// GetUrls should return list of urls from the given url or error
	GetUrls(wd selenium.WebDriver, url string) ([]string, error)

	// NextPage should return next page url or error
	NextPage(wd selenium.WebDriver, url string) (string, error)
}

// Factory is an interface for creating new crawlers and determining if the given url is supported
type Factory interface {

	// NewPageCrawler should return struct implementing PageCrawler interface
	NewPageCrawler() PageCrawler

	// NewListCrawler should return struct implementing ListCrawler interface
	NewListCrawler() ListCrawler

	// MatchUrl should return MatchType for the given url
	// CrawlerMatchPage if the given url is supported by NewPageCrawler
	// CrawlerMatchList if the given url is supported by NewListCrawler
	// CrawlerNotMatch if the given url is not supported
	MatchUrl(url string) MatchType
}

type Offer struct {
	SiteId     string
	Site       string
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
