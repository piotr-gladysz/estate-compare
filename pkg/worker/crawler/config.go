package crawler

import (
	"errors"
	"github.com/spf13/viper"
	"log/slog"
	"time"
)

var config *Config

func LoadConfig() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	defConfig := getDefaultConfig()

	if err := viper.Unmarshal(defConfig); err != nil {
		slog.Error("failed to unmarshal config", "error", err.Error())
		return err
	}

	config = defConfig
	return nil
}

func GetConfig() Config {
	if config == nil {
		err := errors.New("config not loaded")
		slog.Error(err.Error())
		panic(err)
	}
	return *config
}

type Config struct {
	DatabaseUrl  string `mapstructure:"DATABASE_URL"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`

	CrawlerPeriod time.Duration `mapstructure:"CRAWLER_PERIOD"`

	SeleniumUrl string `mapstructure:"SELENIUM_URL"`
}

func getDefaultConfig() *Config {
	dur, _ := time.ParseDuration("6h")

	return &Config{
		DatabaseUrl:  "mongodb://localhost:27017",
		DatabaseName: "estate-compare",

		CrawlerPeriod: dur,

		SeleniumUrl: "http://localhost:4444/wd/hub",
	}
}
