package db

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository interface {
	Insert(ctx context.Context, notification *model.Notification) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindBy(ctx context.Context, by primitive.M) ([]*model.Notification, error)
}

type notificationRepository struct {
	collection *mongo.Collection
}

func (r *notificationRepository) Insert(ctx context.Context, notification *model.Notification) error {
	res, err := r.collection.InsertOne(ctx, notification)
	notification.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func (r *notificationRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, primitive.M{"_id": id})
	return err
}

func (r *notificationRepository) FindBy(ctx context.Context, by primitive.M) ([]*model.Notification, error) {
	var notifications []*model.Notification
	cursor, err := r.collection.Find(ctx, by)
	if err != nil {
		return nil, err
	}

	err = cursor.All(nil, &notifications)
	return notifications, err
}
