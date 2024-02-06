package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SentNotification struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type SentNotificationRepository interface {
}

type sentNotificationRepository struct {
	collection *mongo.Collection
}
