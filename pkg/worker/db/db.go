package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

const (
	sentNotificationsCollectionName = "sent_notifications"
	notificationsCollectionName     = "notifications"
	offersCollectionName            = "offers"
	watchUrlsCollectionName         = "watch_urls"
	conditionsCollectionName        = "conditions"
)

type dB struct {
	db           *mongo.Client
	databaseName string
}

type DB interface {
	GetAlertRepository() NotificationRepository
	GetOfferRepository() OfferRepository
	GetWatchUrlRepository() WatchUrlRepository
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

func (d *dB) GetAlertRepository() NotificationRepository {
	// TODO: implement
	return nil
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
