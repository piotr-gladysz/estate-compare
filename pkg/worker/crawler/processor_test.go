package crawler

import (
	"context"
	"errors"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/testutils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var testTime = time.Now()

func NewTestOffer() *Offer {
	return &Offer{
		SiteId:         "siteId",
		Site:           "site",
		UpdateTime:     testTime,
		Name:           "name",
		Description:    "description",
		Price:          1,
		Area:           2,
		Rooms:          3,
		Floor:          4,
		BuildingFloors: 5,
		Year:           1900,
		Heating:        "heating",
		Market:         "market",
		Window:         "window",
		Elevator:       true,
		Balcony:        true,
		Media:          []string{"media"},
	}
}

func MockProcessor() (*SitesProcessor, *testutils.WatchUrlRepositoryMock, *testutils.OfferRepositoryMock) {
	watchUrlRepo := &testutils.WatchUrlRepositoryMock{}
	offerRepo := &testutils.OfferRepositoryMock{}

	registry := NewCrawlerFactoryRegistry()

	processor := NewSitesProcessor(context.TODO(), registry, watchUrlRepo, offerRepo)

	return processor, watchUrlRepo, offerRepo
}

func TestSitesProcessor_MapOfferToDB(t *testing.T) {
	processor, _, _ := MockProcessor()

	offer := NewTestOffer()
	url := "url"

	dbOffer := processor.MapOfferToDB(offer, url)

	if dbOffer.SiteId != offer.SiteId {
		t.Error("SiteId should be equal")
	}

	if dbOffer.Site != offer.Site {
		t.Error("Site should be equal")
	}

	if dbOffer.Updated != primitive.NewDateTimeFromTime(offer.UpdateTime) {
		t.Error("Updated should be equal")
	}

	if dbOffer.Name != offer.Name {
		t.Error("Name should be equal")
	}

	if dbOffer.Url != url {
		t.Error("Url should be equal")
	}

	if dbOffer.Area != offer.Area {
		t.Error("Area should be equal")
	}

	if dbOffer.Rooms != offer.Rooms {
		t.Error("Rooms should be equal")
	}

	if dbOffer.Floor != offer.Floor {
		t.Error("Floor should be equal")
	}

	if dbOffer.BuildingFloors != offer.BuildingFloors {
		t.Error("BuildingFloors should be equal")
	}

	if dbOffer.Year != offer.Year {
		t.Error("Year should be equal")
	}

	if dbOffer.Heating != offer.Heating {
		t.Error("Heating should be equal")
	}

	if dbOffer.Market != offer.Market {
		t.Error("Market should be equal")
	}

	if dbOffer.Window != offer.Window {
		t.Error("Window should be equal")
	}

	if dbOffer.Elevator != offer.Elevator {
		t.Error("Elevator should be equal")
	}

	if dbOffer.Balcony != offer.Balcony {
		t.Error("Balcony should be equal")
	}

	if len(dbOffer.Media) != len(offer.Media) {
		t.Error("Media should be equal")
	}
}

func TestSitesProcessor_ProcessSite(t *testing.T) {
	processor, _, offerRepo := MockProcessor()

	factory := NewFactoryMock()
	processor.registry.Register(factory)

	matchCalled := false
	crawlCalled := false

	findByCalled := false
	insertCalled := false
	updateCalled := false

	var dbOffer *db.Offer

	factory.Callback = func(this *FactoryMock, method string, args ...any) {
		switch method {
		case "MatchUrl":
			matchCalled = true
		}
	}

	factory.ReturnPageCrawler.Callback = func(this *PageCrawlerMock, method string, args ...any) {
		switch method {
		case "CrawlOffer":
			crawlCalled = true
		}
	}

	offerRepo.Callback = func(this *testutils.OfferRepositoryMock, method string, args ...any) {
		switch method {
		case "FindBy":
			findByCalled = true
		case "Insert":
			insertCalled = true
			dbOffer = args[0].(*db.Offer)
		case "Update":
			updateCalled = true
			dbOffer = args[0].(*db.Offer)
		}
	}

	// Test 1 - Crawler not match

	factory.ReturnMatchType = CrawlerNotMatch
	err := processor.ProcessSite(context.TODO(), nil, "test")

	if !matchCalled {
		t.Error("MatchUrl should be called")
	}

	if err == nil {
		t.Error("Error should be nil")
	} else if !errors.Is(err, CrawlerNotFoundError) {
		t.Error("Error should be ErrCrawlerNotFound")
	}

	// Test 2 - Crawler match page, page not found

	matchCalled = false
	crawlCalled = false

	findByCalled = false
	insertCalled = false
	updateCalled = false

	testOffer := NewTestOffer()

	factory.ReturnMatchType = CrawlerMatchPage
	factory.ReturnPageCrawler.ReturnOffer = testOffer

	offerRepo.ReturnMany = make([]*db.Offer, 0)

	err = processor.ProcessSite(context.TODO(), nil, "test")

	if err != nil {
		t.Error("Error should be nil")
	}

	if !findByCalled {
		t.Error("FindBy should be called")
	}

	if !crawlCalled {
		t.Error("CrawlOffer should be called")
	}

	if !insertCalled {
		t.Error("Insert should be called")
	}

	if updateCalled {
		t.Error("Update should not be called")
	}

	if dbOffer == nil {
		t.Error("Offer should not be nil")
	}

	// Test 3 - Crawler match page, page found, price not changed

	matchCalled = false
	crawlCalled = false

	findByCalled = false
	insertCalled = false
	updateCalled = false

	testOffer = NewTestOffer()

	factory.ReturnMatchType = CrawlerMatchPage
	factory.ReturnPageCrawler.ReturnOffer = testOffer

	offerRepo.ReturnMany = []*db.Offer{{
		History: []*db.OfferHistory{{
			Price: testOffer.Price + 10,
		}},
	}}

	err = processor.ProcessSite(context.TODO(), nil, "test")

	if err != nil {
		t.Error("Error should be nil")
	}

	if !updateCalled {
		t.Error("Update should be called")
	}

	if insertCalled {
		t.Error("Insert should not be called")
	}

	if dbOffer == nil {
		t.Error("Offer should not be nil")
	}

	if len(dbOffer.History) != 2 {
		t.Error("History should have 2 elements")
	}

}
