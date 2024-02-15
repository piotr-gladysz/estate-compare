package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Notification struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	ConditionId primitive.ObjectID `json:"conditionId" bson:"conditionId"`
	Config      map[string]any     `json:"config" bson:"config"`
}

type NotificationRepository interface {
	Insert(ctx context.Context, notification *Notification) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindBy(ctx context.Context, by primitive.M) ([]*Notification, error)
}

type notificationRepository struct {
	collection *mongo.Collection
}

func (r *notificationRepository) Insert(ctx context.Context, notification *Notification) error {
	res, err := r.collection.InsertOne(ctx, notification)
	notification.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func (r *notificationRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, primitive.M{"_id": id})
	return err
}

func (r *notificationRepository) FindBy(ctx context.Context, by primitive.M) ([]*Notification, error) {
	var notifications []*Notification
	cursor, err := r.collection.Find(ctx, by)
	if err != nil {
		return nil, err
	}

	err = cursor.All(nil, &notifications)
	return notifications, err
}
