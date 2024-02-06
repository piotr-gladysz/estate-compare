package wasmutil

// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

// StrToLeakedPtr copies a string to a newly allocated memory and returns a single uint64 with format (ptr << 32) | size
// This method is intended to be used in WASM exports
func StrToLeakedPtr(str string) uint64 {
	size := C.ulong(len(str))
	ptr := unsafe.Pointer(C.malloc(size))
	copy(unsafe.Slice((*byte)(ptr), size), str)

	return uint64(uintptr(ptr))<<32 | uint64(size)

}

// StrToPtr converts a string to an uint64 pointer in the module's memory. It returns the pointer in format (ptr << 32) | size
// This method is intended to be used in WASM exports
func StrToPtr(str string) uint64 {
	ptr := unsafe.Pointer(unsafe.StringData(str))

	return uint64(uintptr(ptr))<<32 | uint64(len(str))
}

// PtrToStr converts a single uint64 with format (ptr << 32) | size to a string
// This method is intended to be used in WASM exports
func PtrToStr(strPtr uint64) string {
	ptr := strPtr >> 32
	size := strPtr & 0xffffffff

	return string(unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), int(size)))
}
