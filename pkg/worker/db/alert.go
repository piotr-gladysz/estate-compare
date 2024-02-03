package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Alert struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type AlertRepository interface {
	Insert(alert *Alert) error
}
