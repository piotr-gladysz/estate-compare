package main

// Test plugin which exports valid function but with invalid arguments

//export CheckCondition
func CheckCondition(offerPtr uint32) uint64 {
	return 0
}

// _log is a WebAssembly import which prints a string (linear memory offset, byteCount) to the console.
//
//go:wasmimport env log
func _log(ptr uint64)

func main() {}
