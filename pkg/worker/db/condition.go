package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Condition struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Wasm []byte             `json:"wasm" bson:"wasm"`
}

type ConditionRepository interface {
	Insert(ctx context.Context, condition *Condition) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindBy(ctx context.Context, by primitive.M) ([]*Condition, error)
}

type conditionRepository struct {
	collection *mongo.Collection
}
