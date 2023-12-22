package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type WatchUrl struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}
