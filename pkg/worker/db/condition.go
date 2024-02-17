package db

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConditionRepository interface {
	Insert(ctx context.Context, condition *model.Condition) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindBy(ctx context.Context, by primitive.M) ([]*model.Condition, error)
	GetWasm(ctx context.Context, id primitive.ObjectID) ([]byte, error)
}

type conditionRepository struct {
	collection *mongo.Collection
}

func (r *conditionRepository) Insert(ctx context.Context, condition *model.Condition) error {
	res, err := r.collection.InsertOne(ctx, condition)
	condition.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func (r *conditionRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, primitive.M{"_id": id})
	return err
}

func (r *conditionRepository) FindBy(ctx context.Context, by primitive.M) ([]*model.Condition, error) {
	var conditions []*model.Condition

	opts := options.Find().SetProjection(primitive.D{{"wasm", 0}})

	cursor, err := r.collection.Find(ctx, by, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(nil, &conditions)
	return conditions, err
}

func (r *conditionRepository) GetWasm(ctx context.Context, id primitive.ObjectID) ([]byte, error) {
	var condition model.Condition

	opts := options.FindOne().SetProjection(primitive.D{{"wasm", 1}})

	err := r.collection.FindOne(ctx, primitive.M{"_id": id}, opts).Decode(&condition)
	return condition.Wasm, err
}
