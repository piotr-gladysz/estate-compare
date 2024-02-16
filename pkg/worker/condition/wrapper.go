package condition

import (
	"context"
	"fmt"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"io"
	"log/slog"
)

var FunctionNotFoundError = fmt.Errorf("function not found")

type Wrapper struct {
	module    api.Module
	checkFunc api.Function
}

func NewWrapper(ctx context.Context, reader io.Reader) (*Wrapper, error) {

	data, err := io.ReadAll(reader)

	if err != nil {
		slog.Error("failed to read module data", "error", err)
		return nil, err
	}

	runtime := wazero.NewRuntime(ctx)

	_, err = runtime.NewHostModuleBuilder("env").
		NewFunctionBuilder().WithFunc(Log).Export("log").
		Instantiate(ctx)

	if err != nil {
		slog.Error("failed to instantiate host module", "error", err)
		return nil, err
	}

	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

	module, err := runtime.Instantiate(ctx, data)

	if err != nil {
		slog.Error("failed to instantiate module", "error", err)
		return nil, err
	}

	checkFunc := module.ExportedFunction("CheckCondition")

	if checkFunc == nil {
		return nil, FunctionNotFoundError
	}

	//TODO: check parameters type of checkFunc

	return &Wrapper{
		module:    module,
		checkFunc: checkFunc,
	}, nil
}

func (w *Wrapper) Close(ctx context.Context) error {
	return w.module.Close(ctx)
}
