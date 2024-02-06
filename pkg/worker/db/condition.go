package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Condition struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type conditionRepository struct {
	collection *mongo.Collection
}
