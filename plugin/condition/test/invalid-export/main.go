package main

// #include <stdlib.h>
import "C"

//export NotCheckCondition
func NotCheckCondition(offerPtr, configPtr uint64) uint32 {
	return 0
}

// _log is a WebAssembly import which prints a string (linear memory offset, byteCount) to the console.
//
//go:wasmimport env log
func _log(ptr uint64)

func main() {}
