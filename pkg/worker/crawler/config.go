package crawler

import (
	"errors"
	"github.com/spf13/viper"
	"log/slog"
	"time"
)

// TODO: move to separate package

var (
	ConfigNotLoadedError = errors.New("config not loaded")

	config *Config
)

func LoadConfig() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

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
		slog.Error(ConfigNotLoadedError.Error())
		panic(ConfigNotLoadedError)
	}
	return *config
}

type Config struct {
	DatabaseUrl  string `mapstructure:"DATABASE_URL"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`

	CrawlerPeriod     time.Duration `mapstructure:"CRAWLER_PERIOD"`
	CrawlerPagesCount int           `mapstructure:"CRAWLER_PAGES_COUNT"`

	SeleniumUrl string `mapstructure:"SELENIUM_URL"`

	ServerEnabled bool   `mapstructure:"SERVER_ENABLED"`
	ServerPort    int    `mapstructure:"SERVER_PORT"`
	ServerIp      string `mapstructure:"SERVER_IP"`
}

func getDefaultConfig() *Config {
	dur, _ := time.ParseDuration("6h")

	return &Config{
		DatabaseUrl:  "mongodb://localhost:27017",
		DatabaseName: "estate-compare",

		CrawlerPeriod:     dur,
		CrawlerPagesCount: 3,

		SeleniumUrl: "http://localhost:4444/wd/hub",

		ServerEnabled: true,
		ServerPort:    11080,
		ServerIp:      "0.0.0.0",
	}
}
