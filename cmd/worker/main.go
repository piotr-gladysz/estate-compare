package main

import (
	"context"
	"fmt"
	_ "github.com/piotr-gladysz/estate-compare/pkg/util/hack/pprof"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/admin"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/crawler"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
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
	conf := crawler.GetConfig()

	fmt.Printf("%+v\n", conf)

	ctx := context.Background()

	d, err := db.NewDB(ctx, conf.DatabaseUrl, conf.DatabaseName)

	if err != nil {
		slog.Error("failed to create db", "error", err.Error())
		return
	}

	defer func() {
		err := d.Close(context.Background())
		if err != nil {
			slog.Error("failed to close db", "error", err.Error())
		}
	}()

	processor := crawler.NewSitesProcessor(
		ctx,
		crawler.NewCrawlerFactoryRegistry(),
		d.GetWatchUrlRepository(),
		d.GetOfferRepository(),
	)

	if conf.ServerEnabled {
		server := admin.NewServer(conf.ServerPort, conf.ServerIp, d.GetWatchUrlRepository(), d.GetOfferRepository(), processor)
		err = server.Run()
		if err != nil {
			slog.Error("failed to run server", "error", err.Error())
			panic(err)
		}

		defer func() {
			err := server.Close()
			if err != nil {
				slog.Error("failed to close server", "error", err.Error())
			}
		}()
	}

	err = processor.Run()
	if err != nil {
		slog.Error("failed to run processor", "error", err.Error())
		panic(err)
	}

}
