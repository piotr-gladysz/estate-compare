package main

import (
	"fmt"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/crawler"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/site"
	"log/slog"
	"os"
)

func init() {
	err := crawler.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", err.Error())
		panic(err)
	}

	logOpts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handler := slog.NewTextHandler(os.Stdout, &logOpts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func main() {

	// TODO: dev version, to change
	config := crawler.GetConfig()

	fmt.Printf("%+v\n", config)

	wd, err := crawler.GetSelenium()

	if err != nil {
		panic(err)
	}

	defer wd.Quit()

	//err := db.InitDB(config.DatabaseUrl, config.DatabaseName)
	//if err != nil {
	//	panic(err)
	//}
	//
	///// test
	//
	fact := site.OtodomFactory{}
	crawl := fact.NewListCrawler()

	url := "https://www.otodom.pl/pl/wyniki/sprzedaz/mieszkanie/podkarpackie/rzeszow/rzeszow/rzeszow?viewType=listing"

	//urls, err := crawl.GetUrls(wd, url)
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", urls)

	next, err := crawl.NextPage(wd, url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", next)

}
