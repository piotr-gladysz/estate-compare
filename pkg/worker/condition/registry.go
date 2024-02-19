package condition

import (
	"bytes"
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Registry struct {
	conditions        map[string]*Wrapper
	conditionProvider ConditionProvider
}

type ConditionProvider interface {
	GetWasm(ctx context.Context, id primitive.ObjectID) ([]byte, error)
}

func NewConditionRegistry(provider ConditionProvider) *Registry {
	return &Registry{conditions: make(map[string]*Wrapper), conditionProvider: provider}
}

func (r *Registry) Register(condition *model.Condition, wrapper *Wrapper) {

	if _, ok := r.conditions[condition.ID.Hex()]; ok {
		return
	}

	r.conditions[condition.ID.Hex()] = wrapper
}

func (r *Registry) Get(ctx context.Context, condition *model.Condition) (*Wrapper, error) {

	if wrapper, ok := r.conditions[condition.ID.Hex()]; ok {
		return wrapper, nil
	}

	wasm, err := r.conditionProvider.GetWasm(context.Background(), condition.ID)
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(wasm)

	wrapper, err := NewWrapper(ctx, buff)
	if err != nil {
		return nil, err
	}
	r.conditions[condition.ID.Hex()] = wrapper

	return wrapper, nil
}
