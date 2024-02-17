package db

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SentNotificationRepository interface {
}

type sentNotificationRepository struct {
	collection *mongo.Collection
}

func (r *sentNotificationRepository) Insert(ctx context.Context, sent *model.SentNotification) error {

	res, err := r.collection.InsertOne(ctx, sent)
	sent.ID = res.InsertedID.(primitive.ObjectID)
	return err
}
