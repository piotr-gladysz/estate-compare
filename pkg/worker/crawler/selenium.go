package crawler

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
	"log/slog"
)

func GetSelenium() (selenium.WebDriver, error) {
	caps := selenium.Capabilities{
		"browserName": "firefox", "browserVersion": "120.0",
	}
	firefoxCaps := firefox.Capabilities{
		Args: []string{
			"--headless",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",
		},
	}
	caps.AddFirefox(firefoxCaps)

	config := GetConfig()

	wd, err := selenium.NewRemote(caps, config.SeleniumUrl)

	if err != nil {
		slog.Error("failed to create selenium", err.Error())
	}

	return wd, err

}
