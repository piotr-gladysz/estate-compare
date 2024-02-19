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
var InvalidOfferError = errors.New("invalid offer")

type SentNotificationRepository interface {
	Insert(ctx context.Context, sent *model.SentNotification, notification *model.Notification, offer *model.Offer) error
	FindBy(ctx context.Context, by primitive.M) ([]*model.SentNotification, error)
	Update(ctx context.Context, sent *model.SentNotification) error
}

type sentNotificationRepository struct {
	collection *mongo.Collection
}

func (r *sentNotificationRepository) Insert(ctx context.Context, sent *model.SentNotification, notification *model.Notification, offer *model.Offer) error {

	if notification == nil || notification.ID.IsZero() {
		return InvalidNotificationError
	}

	if offer == nil || offer.ID.IsZero() {
		return InvalidOfferError
	}

	sent.NotificationId = notification.ID
	sent.OfferId = offer.ID

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

func (r *sentNotificationRepository) Update(ctx context.Context, sent *model.SentNotification) error {
	_, err := r.collection.UpdateOne(ctx, primitive.M{"_id": sent.ID}, primitive.M{"$set": sent})
	return err
}
