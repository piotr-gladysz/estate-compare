package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Offer struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type OfferHistory struct {
}

type OfferRepository struct {
	collection *mongo.Collection
}
