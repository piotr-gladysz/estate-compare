package admin

import (
	"bytes"
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/condition"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

var _ api.ConditionServiceServer = (*ConditionServer)(nil)

var ConditionAlreadyExists = errors.Errorf("condition already exists")

type ConditionServer struct {
	api.UnimplementedConditionServiceServer

	condRepo  db.ConditionRepository
	notifRepo db.NotificationRepository
}

func NewConditionServer(condRepo db.ConditionRepository, notifRepo db.NotificationRepository) *ConditionServer {
	return &ConditionServer{condRepo: condRepo, notifRepo: notifRepo}
}

func (c *ConditionServer) AddCondition(ctx context.Context, request *api.AddConditionRequest) (*api.ConditionResponse, error) {

	existing, err := c.condRepo.FindBy(ctx, primitive.M{"name": request.Name})

	if err != nil {
		return nil, err
	}

	if len(existing) > 0 {
		slog.Error("condition already exists", "name", request.Name)
		return nil, ConditionAlreadyExists
	}

	buf := bytes.NewBuffer(request.Wasm)

	w, err := condition.NewWrapper(ctx, buf)

	if err != nil {
		return nil, err
	}

	defer w.Close(ctx)

	cond := &model.Condition{
		Name: request.Name,
		Wasm: request.Wasm,
	}

	err = c.condRepo.Insert(ctx, cond)

	if err != nil {
		return nil, err
	}

	return c.mapConditionToResponse(cond), nil

}

func (c *ConditionServer) GetCondition(ctx context.Context, request *api.GetConditionRequest) (*api.ConditionResponse, error) {

	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	ret, err := c.condRepo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return c.mapConditionToResponse(ret), nil
}

func (c *ConditionServer) GetConditions(ctx context.Context, request *api.GetConditionsRequest) (*api.ConditionsListResponse, error) {
	limit := request.PageSize
	skip := (request.Page - 1) * request.PageSize

	ret, total, err := c.condRepo.FindAll(ctx, int64(skip), int64(limit))

	if err != nil {
		return nil, err
	}

	conditionResponses := make([]*api.ConditionResponse, len(ret))

	for i, cond := range ret {
		conditionResponses[i] = c.mapConditionToResponse(cond)
	}

	return &api.ConditionsListResponse{
		Conditions: conditionResponses,
		Total:      total,
	}, nil
}

func (c *ConditionServer) DeleteCondition(ctx context.Context, request *api.GetConditionRequest) (*api.ConditionResponse, error) {

	id, err := primitive.ObjectIDFromHex(request.Id)

	if err != nil {
		return nil, err
	}

	ret, err := c.condRepo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	err = c.condRepo.Delete(ctx, ret.ID)

	if err != nil {
		return nil, err
	}

	return c.mapConditionToResponse(ret), nil
}

func (c *ConditionServer) mapConditionToResponse(condition *model.Condition) *api.ConditionResponse {
	return &api.ConditionResponse{
		Id:   condition.ID.Hex(),
		Name: condition.Name,
	}
}
