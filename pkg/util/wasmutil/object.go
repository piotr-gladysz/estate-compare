package wasmutil

// #include <stdlib.h>
import "C"

import (
	"unsafe"

	"github.com/shamaton/msgpack/v2"
)

// PtrToObj converts a single uint64 with format (ptr << 32) | size to an object
// This method is intended to be used in WASM exports
func PtrToObj(strPtr uint64, obj any) error {

	ptr := strPtr >> 32
	size := strPtr & 0xffffffff

	bytes := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), int(size))

	err := msgpack.Unmarshal(bytes, obj)
	if err != nil {
		return err
	}

	return nil

}

// ObjToPtr copies an object to a newly allocated memory and returns a single uint64 with format (ptr << 32) | size
// This method is intended to be used in WASM exports
func ObjToPtr(obj any) (uint64, error) {

	bytes, err := msgpack.Marshal(obj)
	if err != nil {
		return 0, err
	}

	size := uint64(len(bytes))
	ptr := C.malloc(C.size_t(size))

	copy(unsafe.Slice((*byte)(ptr), size), bytes)

	return uint64(uintptr(ptr))<<32 | size, nil
}
