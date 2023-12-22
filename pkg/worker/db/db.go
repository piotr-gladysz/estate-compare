package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

const (
	alertsCollectionName    = "alerts"
	offersCollectionName    = "offers"
	watchUrlsCollectionName = "watch_urls"
)

var (
	db           *mongo.Client
	databaseName string
)

func InitDB(dbUrl, dbName string) error {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUrl))
	if err != nil {
		slog.Error("Failed to connect to database", err.Error())
		return err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		slog.Error("Failed to connect to database", err.Error())
		return err
	}

	db = client
	databaseName = dbName
	return nil
}

func OffersCollection() *mongo.Collection {
	return db.Database(databaseName).Collection(offersCollectionName)
}

func AlertsCollection() *mongo.Collection {
	return db.Database(databaseName).Collection(alertsCollectionName)
}

func WatchUrlsCollection() *mongo.Collection {
	return db.Database(databaseName).Collection(watchUrlsCollectionName)
}
