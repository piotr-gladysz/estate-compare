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

type DB struct {
	db           *mongo.Client
	databaseName string
}

func NewDB(dbUrl, dbName string) (*DB, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUrl))
	if err != nil {
		slog.Error("Failed to connect to database", err.Error())
		return nil, err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		slog.Error("Failed to connect to database", err.Error())
		return nil, err
	}

	return &DB{
		db:           client,
		databaseName: dbName,
	}, nil
}
