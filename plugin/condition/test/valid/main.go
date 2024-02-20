package main

import (
	"fmt"
	"github.com/piotr-gladysz/estate-compare/pkg/util/wasmutil"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

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
//
//export CheckCondition
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
//
//go:wasmimport env log
func _log(ptr uint64)

// main is required for the plugin to build
func main() {}
