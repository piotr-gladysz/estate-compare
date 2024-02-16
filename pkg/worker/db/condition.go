package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Condition struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Updated primitive.DateTime `json:"updated" bson:"updated"`
	Created primitive.DateTime `json:"created" bson:"created"`
	Name    string             `json:"name" bson:"name"`
	Wasm    []byte             `json:"wasm" bson:"wasm"`
}

type ConditionRepository interface {
	Insert(ctx context.Context, condition *Condition) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindBy(ctx context.Context, by primitive.M) ([]*Condition, error)
	GetWasm(ctx context.Context, id primitive.ObjectID) ([]byte, error)
}

type conditionRepository struct {
	collection *mongo.Collection
}

func (r *conditionRepository) Insert(ctx context.Context, condition *Condition) error {
	res, err := r.collection.InsertOne(ctx, condition)
	condition.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func (r *conditionRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, primitive.M{"_id": id})
	return err
}

func (r *conditionRepository) FindBy(ctx context.Context, by primitive.M) ([]*Condition, error) {
	var conditions []*Condition

	opts := options.Find().SetProjection(primitive.D{{"wasm", 0}})

	cursor, err := r.collection.Find(ctx, by, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(nil, &conditions)
	return conditions, err
}

func (r *conditionRepository) GetWasm(ctx context.Context, id primitive.ObjectID) ([]byte, error) {
	var condition Condition

	opts := options.FindOne().SetProjection(primitive.D{{"wasm", 1}})

	err := r.collection.FindOne(ctx, primitive.M{"_id": id}, opts).Decode(&condition)
	return condition.Wasm, err
}
