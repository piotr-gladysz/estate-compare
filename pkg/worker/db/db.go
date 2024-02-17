package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

const (
	offersCollectionName            = "offers"
	watchUrlsCollectionName         = "watch_urls"
	conditionsCollectionName        = "conditions"
	notificationsCollectionName     = "notifications"
	sentNotificationsCollectionName = "sent_notifications"
)

type dB struct {
	db           *mongo.Client
	databaseName string
}

type DB interface {
	GetOfferRepository() OfferRepository
	GetWatchUrlRepository() WatchUrlRepository
	GetConditionRepository() ConditionRepository
	GetNotificationRepository() NotificationRepository
	GetSentNotificationRepository() SentNotificationRepository
	Close(ctx context.Context) error
}

func NewDB(ctx context.Context, dbUrl, dbName string) (DB, error) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		slog.Error("Failed to connect to database", "error", err.Error())
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		slog.Error("Failed to connect to database", "error", err.Error())
		return nil, err
	}

	return &dB{
		db:           client,
		databaseName: dbName,
	}, nil
}

func (d *dB) Close(ctx context.Context) error {
	return d.db.Disconnect(ctx)
}

func (d *dB) GetOfferRepository() OfferRepository {
	return &offerRepository{collection: d.getCollection(offersCollectionName)}
}

func (d *dB) GetWatchUrlRepository() WatchUrlRepository {
	return &watchUrlRepository{collection: d.getCollection(watchUrlsCollectionName)}
}

func (d *dB) getCollection(collectionName string) *mongo.Collection {
	return d.db.Database(d.databaseName).Collection(collectionName)
}

func (d *dB) GetConditionRepository() ConditionRepository {
	return &conditionRepository{collection: d.getCollection(conditionsCollectionName)}
}

func (d *dB) GetNotificationRepository() NotificationRepository {
	return &notificationRepository{collection: d.getCollection(notificationsCollectionName)}
}

func (d *dB) GetSentNotificationRepository() SentNotificationRepository {
	return &sentNotificationRepository{collection: d.getCollection(sentNotificationsCollectionName)}
}
