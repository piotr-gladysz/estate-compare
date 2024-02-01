package crawler

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/testutils"
	"testing"
)

func MockProcessor() (*SitesProcessor, *testutils.WatchUrlRepositoryMock, *testutils.OfferRepositoryMock) {
	watchUrlRepo := &testutils.WatchUrlRepositoryMock{}
	offerRepo := &testutils.OfferRepositoryMock{}

	registry := NewCrawlerFactoryRegistry()

	processor := NewSitesProcessor(context.TODO(), registry, watchUrlRepo, offerRepo)

	return processor, watchUrlRepo, offerRepo
}

func TestSitesProcessor_ProcessSite(t *testing.T) {
	//processor, watchUrlRepo, offerRepo := MockProcessor()
	//
	//factory := testutils.NewFactoryMock()
	//processor.registry.Register(factory)

}
