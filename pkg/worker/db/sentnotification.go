package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type SentNotificationRepository interface {
}

type sentNotificationRepository struct {
	collection *mongo.Collection
}
