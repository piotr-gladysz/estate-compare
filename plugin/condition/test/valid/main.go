package main

// #include <stdlib.h>
import "C"

import (
	"fmt"
	"github.com/piotr-gladysz/estate-compare/pkg/util/wasmutil"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//export CheckCondition
func CheckCondition(offerPtr, configPtr uint64) uint64 {

	var offer model.Offer

	err := wasmutil.PtrToObj(offerPtr, &offer)
	if err != nil {
		_log(wasmutil.StrToPtr(err.Error()))
		return 0
	}

	var config map[string]any

	err = wasmutil.PtrToObj(configPtr, &config)
	if err != nil {
		_log(wasmutil.StrToPtr(err.Error()))
		return 0

	}

	now := time.Now()

	notif := model.SentNotification{
		OfferId: offer.ID,
		Created: primitive.NewDateTimeFromTime(now),
		Updated: primitive.NewDateTimeFromTime(now),

		Message: "Offer: " + offer.Name +
			"\nUrl: " + offer.Url +
			"\nHistory len: " + fmt.Sprintf("%d", len(offer.History)) +
			"\nHistory: " + fmt.Sprintf("%v+", offer.History) +
			"\nConfig: " + fmt.Sprintf("%v+", config),
	}

	retPtr, err := wasmutil.ObjToPtr(notif)

	if err != nil {
		_log(wasmutil.StrToPtr(err.Error()))
		return 0

	}

	return retPtr
}

// _log is a WebAssembly import which prints a string (linear memory offset, byteCount) to the console.
//
//go:wasmimport env log
func _log(ptr uint64)

func main() {}
