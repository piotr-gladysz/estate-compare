# Estate Compare

## Overview
The Estate Price Notifier is a tool designed to monitor and notify users of real estate price changes. With customizable crawlers, notification senders, and the ability to add complex logic for checking notification conditions through WebAssembly (WASM) plugins..

## Features
- **Custom Crawlers**: Define your own crawlers to fetch real estate prices from various sources.
- **Notification Channels**: Set up custom notification mechanisms (e.g., email, SMS, push notifications) to alert you of price changes.
- **WASM Plugins**: Extend the project's functionality with WASM plugins to implement complex logic for when notifications should be triggered.
- **CLI Interface**: Use the command-line interface to manage your links, check results, dynamically load plugins, and more.

## TODO
- [ ] More examples(crawlers, senders, plugins)
- [ ] Refactor config
- [ ] Basic notification channels 
- [ ] Resend notifications if failed
- [ ] Add more tests, especially benchmarks
- [ ] Add more documentation


## Getting Started

### Prerequisites
- Go 1.21 or newer([link](https://golang.org/dl/))
- Docker([link](https://www.docker.com/products/docker-desktop))
- MongoDB, recommended to use docker image `mongo:7`
- Selenium, recommended to use docker image `selenium/standalone-firefox:120.0`
- Tinygo for building WASM plugins([link](https://tinygo.org/getting-started/))

### Usage


#### WASM plugins
To add custom logic for checking notification conditions, you can create a WebAssembly plugin.
The plugin must meet the following conditions.
- Because WASM cannot receive and return complex types, the plugin must use the `wasmutil` package to convert complex types to and from pointers.
- The plugin must export a function called `CheckCondition` with the following signature(parameters and explained in the example):
```go
CheckCondition(offerPtr, configPtr, action uint64) uint64
```
- The plugin must have a main function with an empty body.
- The plugin can use the `_log` function to print to the console. This function receives a pointer to a string the same as `CheckCondition` function.
Example:
```go
package main
import (
	"fmt"
	"github.com/piotr-gladysz/estate-compare/pkg/util/wasmutil"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//export CheckCondition
// CheckCondition is a function that checks if the given offer meets the conditions for sending a notification
// offerPtr is a pointer to the Offer struct
// configPtr is a pointer to the config map[string]any
// action is a value of OfferAction
//
// OfferActionAdd         OfferAction = 1 // Offer is created
// OfferActionUpdate      OfferAction = 2 // Unused
// OfferActionPriceChange OfferAction = 3 // Offer price has changed
// OfferActionSame        OfferAction = 4 // Offer was crawled but nothing has changed
//
// returns pointer to the SentNotification struct or 0 if notification should not be sent
func CheckCondition(offerPtr, configPtr, action uint64) uint64 {

	var offer model.Offer

	// Convert pointer to Offer struct
	if err := wasmutil.PtrToObj(offerPtr, &offer); err != nil {
		_log(wasmutil.StrToPtr(err.Error()))
		return 0
	}

	var config map[string]any

	// Convert pointer to config map
	if err := wasmutil.PtrToObj(configPtr, &config); err != nil {
		_log(wasmutil.StrToPtr(err.Error()))
		return 0
	}

	now := time.Now()

	// Example notification
	notif := model.SentNotification{
		OfferId: offer.ID,
		Created: primitive.NewDateTimeFromTime(now),
		Updated: primitive.NewDateTimeFromTime(now),

		Message: "Offer: " + offer.Name +
			"\nUrl: " + offer.Url +
			"\nHistory len: " + fmt.Sprintf("%d", len(offer.History)) +
			"\nHistory: " + fmt.Sprintf("%v+", offer.History) +
			"\nConfig: " + fmt.Sprintf("%v+", config) +
			"\nAction: " + fmt.Sprintf("%d", action),
	}

	// Convert SentNotification struct to pointer
	retPtr, err := wasmutil.ObjToPtr(notif)
	if err != nil {
		_log(wasmutil.StrToPtr(err.Error()))
		return 0

	}

	return retPtr
}

// _log is a WebAssembly import which prints a string to the console.
// ptr must be in the form of (ptr << 32) | size
//go:wasmimport env log
func _log(ptr uint64)

// main is required for the plugin to build
func main() {}

```
Such plugin can be built using command:
```bash
tinygo build -o bin/plugins/condition/my-plugin.wasm -scheduler=none -target=wasi ./plugin/condition/my-plugin/main.go
````

If successful, the plugin will be placed in the `bin/plugins/condition` directory. Now can be added to database using CLI with command:
```bash
./bin/cli condition add --path bin/plugins/condition/my-plugin.wasm --name my-plugin
```


#### Crawler factory
To add crawler logic, the following interfaces must be implemented:

```go
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
```
Then register the factory in FactoryRegistry which is passed to site processor

```go
factoryRegistry := crawler.NewCrawlerFactoryRegistry()
factoryRegistry.Register(NewMyCrawlerFactory())
processor := crawler.NewSitesProcessor(ctx, factoryRegistry, ...)
```
#### Notification channels
To add notification sender logic, the following interface must be implemented:

```go
// NotificationChannel is an interface for sending notifications
type NotificationChannel interface {
	
    // SendNotification sends notification to the user and returns error if any
    SendNotification(ctx context.Context, sentNotif *model.SentNotification, offer *model.Offer) error
    
    // GetName returns name of the sender for identification
    GetName() string
}
```
Then register the channel in SenderRegistry which is passed to notification processor

```go
channelRegistry := notification.NewChannelRegistry()
channelRegistry.Register(NewMyNotificationChannel())
```

ChannelRegistry is passed to `NotificationSender` which is responsible for sending notifications using registered channels.
You can use simple implementation called `Notifier` which sends notifications to all registered channels or create your own.
`NotificationSender` is passed to site processor

```go
notifier := notification.NewNotifier(db, conditionRegistry, channelRegistry)
processor := crawler.NewSitesProcessor(ctx, factoryRegistry, notifier, ...)

```



### Project structure
TODO

## License
This project is licensed under the MIT License