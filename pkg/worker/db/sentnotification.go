package db

import (
	"context"
	"errors"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ SentNotificationRepository = (*sentNotificationRepository)(nil)

var InvalidNotificationError = errors.New("invalid notification")

type SentNotificationRepository interface {
}

type sentNotificationRepository struct {
	collection *mongo.Collection
}

func (r *sentNotificationRepository) Insert(ctx context.Context, sent *model.SentNotification, notification *model.Notification) error {

	if notification == nil || notification.ID.IsZero() {
		return InvalidNotificationError
	}

	sent.NotificationId = notification.ID

	res, err := r.collection.InsertOne(ctx, sent)
	sent.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func (r *sentNotificationRepository) FindBy(ctx context.Context, by primitive.M) ([]*model.SentNotification, error) {
	var sentNotifications []*model.SentNotification
	cursor, err := r.collection.Find(ctx, by)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(nil, &sentNotifications)
	return sentNotifications, err
}
