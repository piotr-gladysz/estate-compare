package main

import (
	"github.com/piotr-gladysz/estate-compare/pkg/worker/crawler"
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

}
