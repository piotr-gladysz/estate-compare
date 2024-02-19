//go:build pprof
// +build pprof

package pprof

import (
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
}
