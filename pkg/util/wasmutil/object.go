package wasmutil

// #include <stdlib.h>
import "C"

import (
	"unsafe"

	"github.com/shamaton/msgpack/v2"
)

func ptrToObj(strPtr uint64, obj any) error {

	ptr := strPtr >> 32
	size := strPtr & 0xffffffff

	bytes := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), int(size))

	err := msgpack.Unmarshal(bytes, obj)
	if err != nil {
		return err
	}

	return nil

}

func objToPtr(obj any) (uint64, error) {

	bytes, err := msgpack.Marshal(obj)
	if err != nil {
		return 0, err
	}

	size := uint64(len(bytes))
	ptr := C.malloc(C.size_t(size))

	copy(unsafe.Slice((*byte)(ptr), size), bytes)

	return uint64(uintptr(ptr))<<32 | size, nil
}
