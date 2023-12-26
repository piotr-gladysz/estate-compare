package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WatchUrl struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Url      string             `json:"url" bson:"url"`
	IsList   bool               `json:"isList" bson:"isList"`
	Created  int64              `json:"created" bson:"created"`
	Updated  int64              `json:"updated" bson:"updated"`
	Disabled bool               `json:"disabled" bson:"disabled"`
}

type WatchUrlRepository struct {
	collection *mongo.Collection
}

func (r *WatchUrlRepository) Insert(ctx context.Context, watchUrl *WatchUrl) error {
	_, err := r.collection.InsertOne(ctx, watchUrl)
	return err
}

func (r *WatchUrlRepository) InsertIfNotExists(ctx context.Context, watchUrl *WatchUrl) error {

	ret := r.collection.FindOneAndUpdate(ctx,
		primitive.M{"url": watchUrl.Url, "isList": watchUrl.IsList},
		primitive.M{"$setOnInsert": watchUrl},
		options.FindOneAndUpdate().SetUpsert(true))

	return ret.Err()
}

func (r *WatchUrlRepository) Update(ctx context.Context, watchUrl *WatchUrl) error {
	_, err := r.collection.UpdateOne(ctx, primitive.M{"_id": watchUrl.ID}, primitive.M{"$set": watchUrl})
	return err
}

func (r *WatchUrlRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, primitive.M{"_id": id})
	return err
}

func (r *WatchUrlRepository) FindById(ctx context.Context, id primitive.ObjectID) (*WatchUrl, error) {
	var watchUrl WatchUrl
	err := r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&watchUrl)
	return &watchUrl, err
}

func (r *WatchUrlRepository) FindBy(ctx context.Context, by primitive.M) ([]*WatchUrl, error) {
	var watchUrls []*WatchUrl
	cursor, err := r.collection.Find(ctx, by)
	if err != nil {
		return nil, err
	}
	err = cursor.All(nil, &watchUrls)
	return watchUrls, err
}

func (r *WatchUrlRepository) FindAll(ctx context.Context) ([]*WatchUrl, error) {
	var watchUrls []*WatchUrl
	cursor, err := r.collection.Find(ctx, primitive.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &watchUrls)
	return watchUrls, err
}
