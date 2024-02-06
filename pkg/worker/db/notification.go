package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Notification struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type NotificationRepository interface {
	Insert(notification *Notification) error
}

type notificationRepository struct {
	collection *mongo.Collection
}
