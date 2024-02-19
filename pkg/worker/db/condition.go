package db

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ ConditionRepository = (*conditionRepository)(nil)

type ConditionRepository interface {
	Insert(ctx context.Context, condition *model.Condition) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindBy(ctx context.Context, by primitive.M) ([]*model.Condition, error)
	FindById(ctx context.Context, id primitive.ObjectID) (*model.Condition, error)
	FindAll(ctx context.Context, limit int64, skip int64) ([]*model.Condition, int64, error)
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
	defer cursor.Close(ctx)

	return conditions, err
}

func (r *conditionRepository) FindById(ctx context.Context, id primitive.ObjectID) (*model.Condition, error) {
	var condition model.Condition

	err := r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&condition)
	return &condition, err

}

func (r *conditionRepository) FindAll(ctx context.Context, limit int64, skip int64) ([]*model.Condition, int64, error) {
	var conditions []*model.Condition

	opts := options.Find().SetProjection(primitive.D{{"wasm", 0}}).SetLimit(limit).SetSkip(skip)

	cursor, err := r.collection.Find(ctx, primitive.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(nil, &conditions)
	if err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, primitive.M{})

	if err != nil {
		return nil, 0, err
	}

	return conditions, total, err
}

func (r *conditionRepository) GetWasm(ctx context.Context, id primitive.ObjectID) ([]byte, error) {
	var condition model.Condition

	opts := options.FindOne().SetProjection(primitive.D{{"wasm", 1}})

	err := r.collection.FindOne(ctx, primitive.M{"_id": id}, opts).Decode(&condition)
	return condition.Wasm, err
}
