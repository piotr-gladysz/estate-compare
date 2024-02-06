package condition

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/tetratelabs/wazero/api"
)

var MallocNotFoundError = fmt.Errorf("malloc function not found")
var FreeNotFoundError = fmt.Errorf("free function not found")
var WriteMemoryError = fmt.Errorf("failed to write to memory")
var ReadMemoryError = fmt.Errorf("failed to read from memory")

// StrToPtr converts a string to a pointer in the module's memory. It returns the pointer in format (ptr << 32) | size
// This method is intended to be used in host code
func StrToPtr(ctx context.Context, mod api.Module, str string) (uint64, error) {
	malloc := mod.ExportedFunction("malloc")

	if malloc == nil {
		return 0, MallocNotFoundError
	}

	strLen := uint64(len(str))

	ptr, err := malloc.Call(ctx, strLen)
	if err != nil {
		return 0, err
	}

	if !mod.Memory().WriteString(uint32(ptr[0]), str) {
		return 0, WriteMemoryError
	}

	return ptr[0]<<32 | strLen, nil
}

// PtrToStr converts a pointer in format (ptr << 32) | size to a string
// This method is intended to be used in host code
func PtrToStr(ctx context.Context, mod api.Module, strPtr uint64) (string, error) {
	ptr := uint32(strPtr >> 32)
	strLen := uint32(strPtr & 0xffffffff)

	bytes, ok := mod.Memory().Read(ptr, strLen)

	if !ok {
		return "", ReadMemoryError
	}

	err := Free(ctx, mod, strPtr)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Free frees the memory allocated in the module
// This method is intended to be used in host code
func Free(ctx context.Context, mod api.Module, strPtr uint64) error {

	free := mod.ExportedFunction("free")

	if free == nil {
		return FreeNotFoundError
	}

	ptr := uint32(strPtr >> 32)

	_, err := free.Call(ctx, uint64(ptr))

	return err
}

// Alloc allocates memory in the module
// This method is intended to be used in host code
func Alloc(ctx context.Context, mod api.Module, size uint64) (uint64, error) {
	malloc := mod.ExportedFunction("malloc")

	if malloc == nil {
		return 0, MallocNotFoundError
	}

	ptr, err := malloc.Call(ctx, size)
	if err != nil {
		return 0, err
	}

	return ptr[0], nil
}

// Log logs a string
// This method is intended to be exported to WASM
func Log(ctx context.Context, m api.Module, strPtr uint64) {
	str, err := PtrToStr(ctx, m, strPtr)

	if err != nil {
		slog.Error("WASM", "error", err)
		return
	}

	slog.Info("WASM", "log", str)
}
