package condition

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"github.com/pkg/errors"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"io"
	"log/slog"
)

var FunctionNotFoundError = errors.New("function not found")
var InvalidFunctionDefinitionError = errors.New("invalid function definition")

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

	funcParams := checkFunc.Definition().ParamTypes()

	if len(funcParams) != 3 || funcParams[0] != api.ValueTypeI64 || funcParams[1] != api.ValueTypeI64 || funcParams[2] != api.ValueTypeI64 {
		return nil, InvalidFunctionDefinitionError
	}

	funcResults := checkFunc.Definition().ResultTypes()

	if len(funcResults) != 1 || funcResults[0] != api.ValueTypeI64 {
		return nil, InvalidFunctionDefinitionError
	}

	return &Wrapper{
		module:    module,
		checkFunc: checkFunc,
	}, nil
}

func (w *Wrapper) CheckOffer(ctx context.Context, offer *model.Offer, action model.OfferAction, config map[string]any) (*model.SentNotification, error) {

	offerPtr, err := ObjToPointer(ctx, w.module, offer)
	if err != nil {
		return nil, err
	}

	configPtr, err := ObjToPointer(ctx, w.module, config)
	if err != nil {
		return nil, err
	}

	retPtr, err := w.checkFunc.Call(ctx, offerPtr, configPtr, uint64(action))
	if err != nil {
		return nil, err
	}

	if retPtr[0] == 0 {
		return nil, nil
	}

	var ret model.SentNotification
	err = PointerToObj(ctx, w.module, retPtr[0], &ret)

	if err != nil {
		return nil, err

	}

	ret.SendingStatus = make(map[string]any)

	return &ret, nil
}

func (w *Wrapper) Close(ctx context.Context) error {
	return w.module.Close(ctx)
}
