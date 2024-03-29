package db

import (
	"context"
	"errors"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ NotificationRepository = (*notificationRepository)(nil)

var InvalidConditionError = errors.New("invalid condition")

type NotificationRepository interface {
	Insert(ctx context.Context, notification *model.Notification, condition *model.Condition) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindBy(ctx context.Context, by primitive.M) ([]*model.Notification, error)
	FindAll(ctx context.Context, limit int64, skip int64) ([]*model.Notification, int64, error)
}

type notificationRepository struct {
	collection *mongo.Collection
}

func (r *notificationRepository) Insert(ctx context.Context, notification *model.Notification, condition *model.Condition) error {

	if condition == nil || condition.ID.IsZero() {
		return InvalidConditionError
	}

	notification.ConditionId = condition.ID

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
	defer cursor.Close(ctx)

	err = cursor.All(nil, &notifications)
	return notifications, err
}

func (r *notificationRepository) FindAll(ctx context.Context, limit int64, skip int64) ([]*model.Notification, int64, error) {
	var notifications []*model.Notification

	cursor, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetLimit(limit).SetSkip(skip))
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(nil, &notifications)
	if err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, primitive.M{})
	if err != nil {
		return nil, 0, err
	}

	return notifications, total, err
}
