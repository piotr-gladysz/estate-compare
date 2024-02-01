package crawler

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"github.com/tebeka/selenium"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type SitesProcessor struct {
	registry     *FactoryRegistry
	watchUrlRepo db.WatchUrlRepository
	offerRepo    db.OfferRepository

	parentCtx context.Context

	procMux    sync.Mutex
	procJob    gocron.Job
	procCtx    context.Context
	procCancel context.CancelFunc
}

func NewSitesProcessor(ctx context.Context, registry *FactoryRegistry, watchUrlRepo db.WatchUrlRepository, offerRepo db.OfferRepository) *SitesProcessor {
	return &SitesProcessor{
		registry:     registry,
		watchUrlRepo: watchUrlRepo,
		offerRepo:    offerRepo,
		parentCtx:    ctx,
	}
}

// Run starts the processor with periodic jobs
func (s *SitesProcessor) Run() error {

	sched, err := gocron.NewScheduler()
	if err != nil {
		slog.Error("failed to create scheduler", "error", err.Error())
		return err
	}

	job, err := sched.NewJob(
		gocron.DurationJob(config.CrawlerPeriod),
		gocron.NewTask(s.Process),
	)

	s.procJob = job

	if err != nil {
		slog.Error("failed to create job", "error", err.Error())
		return err
	}

	sched.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-c:
			fmt.Println("SitesProcessor: shutting down")
			err = sched.Shutdown()
			if err != nil {
				slog.Warn("failed to shutdown scheduler", "error", err.Error())
			}
			return err
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// Process runs the processor once
func (s *SitesProcessor) Process() error {

	s.procMux.Lock()
	ctx, cancel := context.WithCancel(s.parentCtx)

	s.procCtx = ctx
	s.procCancel = cancel

	s.procMux.Unlock()

	defer func() {
		s.procMux.Lock()
		s.procCancel()
		s.procCtx = nil
		s.procCancel = nil
		s.procMux.Unlock()
	}()

	wd, err := GetSelenium()
	defer func(wd selenium.WebDriver) {
		err := wd.Quit()
		if err != nil {
			slog.Warn("failed to quit selenium", "error", err.Error())
		}
	}(wd)
	if err != nil {
		slog.Error("failed to get selenium", "error", err.Error())
		return err
	}

	lists, err := s.watchUrlRepo.FindBy(ctx, primitive.M{"isList": true, "disabled": false})
	if err != nil {
		return err
	}

	for _, list := range lists {
		_ = s.ProcessSiteList(ctx, wd, list.Url)
	}

	sites, err := s.watchUrlRepo.FindBy(ctx, primitive.M{"isList": false, "disabled": false})
	if err != nil {
		return err
	}

	for _, site := range sites {
		_ = s.ProcessSite(ctx, wd, site.Url)
	}

	return nil
}

func (s *SitesProcessor) ProcessSiteList(ctx context.Context, wd selenium.WebDriver, url string) error {
	_, crawler := s.registry.GetCrawler(url)
	if crawler == nil {
		err := errors.New("crawler not found")
		slog.Warn("failed to get crawler", "url", url, "error", err.Error())
		return err
	}

	for i := 0; i < config.CrawlerPagesCount; i++ {

		links, err := crawler.GetUrls(wd, url)
		if err != nil {
			slog.Warn("failed to get urls", "url", url, "error", err.Error())
			return err
		}

		for _, link := range links {
			watchUrl := &db.WatchUrl{
				Url: link,
			}
			err = s.watchUrlRepo.InsertIfNotExists(ctx, watchUrl)
			if err != nil {
				slog.Warn("failed to add site link", "url", url, "error", err.Error())
				if errors.Is(err, context.Canceled) {
					return err
				}
			}
		}

		url, err = crawler.NextPage(wd, url)
		if err != nil {
			slog.Warn("failed to get next page", "url", url, "error", err.Error())
			return err
		}
	}
	return nil
}

func (s *SitesProcessor) ProcessSite(ctx context.Context, wd selenium.WebDriver, url string) error {
	crawler, _ := s.registry.GetCrawler(url)
	if crawler == nil {
		err := errors.New("crawler not found")
		slog.Warn("failed to get crawler", "url", url, "error", err.Error())
		return err
	}

	offer, err := crawler.CrawlOffer(wd, url)
	if err != nil {
		slog.Warn("failed to get offer", "url", url, "error", err.Error())
		return err
	}

	existing, err := s.offerRepo.FindBy(ctx, primitive.M{"site": offer.Site, "siteId": offer.SiteId})
	if err != nil {
		slog.Warn("failed to get offer", "url", url, "error", err.Error())
		return err
	}

	if len(existing) == 0 {
		dbOffer := s.MapOfferToDB(offer, url)
		err = s.offerRepo.Insert(ctx, dbOffer)
		if err != nil {
			slog.Warn("failed to insert offer", "url", url, "error", err.Error())
			return err
		}
	} else {
		existingOffer := existing[0]
		history := existingOffer.History[len(existingOffer.History)-1]
		if history.Price == offer.Price {
			return nil
		}

		existingOffer.Updated = primitive.NewDateTimeFromTime(offer.UpdateTime)
		existingOffer.History = append(existingOffer.History, &db.OfferHistory{
			Updated: primitive.NewDateTimeFromTime(offer.UpdateTime),
			Price:   offer.Price,
		})
		err = s.offerRepo.Update(ctx, existingOffer)
		if err != nil {
			slog.Warn("failed to update offer", "url", url, "error", err.Error())
			return err
		}
	}

	return nil
}

func (s *SitesProcessor) MapOfferToDB(offer *Offer, url string) *db.Offer {
	return &db.Offer{
		SiteId:         offer.SiteId,
		Site:           offer.Site,
		Updated:        primitive.NewDateTimeFromTime(offer.UpdateTime),
		Name:           offer.Name,
		Url:            url,
		Area:           offer.Area,
		Rooms:          offer.Rooms,
		Floor:          offer.Floor,
		BuildingFloors: offer.BuildingFloors,
		Year:           offer.Year,
		Heating:        offer.Heating,
		Market:         offer.Market,
		Window:         offer.Window,
		Elevator:       offer.Elevator,
		Balcony:        offer.Balcony,
		Media:          offer.Media,
		History: []*db.OfferHistory{
			{
				Updated: primitive.NewDateTimeFromTime(offer.UpdateTime),
				Price:   offer.Price,
			},
		},
	}
}
