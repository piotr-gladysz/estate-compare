package admin

import (
	"bytes"
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/condition"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

var _ api.ConditionServiceServer = (*ConditionServer)(nil)

type ConditionServer struct {
	api.UnimplementedConditionServiceServer
	repo db.ConditionRepository
}

var ConditionAlreadyExists = errors.Errorf("condition already exists")

func (c *ConditionServer) AddCondition(ctx context.Context, in *api.AddConditionRequest) (*api.ConditionResponse, error) {

	existing, err := c.repo.FindBy(ctx, primitive.M{"name": in.Name})

	if err != nil {
		return nil, err
	}

	if len(existing) > 0 {
		slog.Error("condition already exists", "name", in.Name)
		return nil, ConditionAlreadyExists
	}

	buf := bytes.NewBuffer(in.Wasm)

	_, err = condition.NewWrapper(ctx, buf)

	if err != nil {
		return nil, err
	}

	return nil, nil

}

func (c *ConditionServer) GetCondition(ctx context.Context, in *api.GetConditionRequest) (*api.ConditionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ConditionServer) GetConditions(ctx context.Context, in *api.GetConditionsRequest) (*api.ConditionsListResponse, error) {
	//TODO implement me
	panic("implement me")
}
