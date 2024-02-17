package db

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ WatchUrlRepository = (*watchUrlRepository)(nil)

type WatchUrlRepository interface {
	Insert(ctx context.Context, watchUrl *model.WatchUrl) error
	InsertIfNotExists(ctx context.Context, watchUrl *model.WatchUrl) error
	Update(ctx context.Context, watchUrl *model.WatchUrl) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindById(ctx context.Context, id primitive.ObjectID) (*model.WatchUrl, error)
	FindBy(ctx context.Context, by primitive.M) ([]*model.WatchUrl, error)
	FindAll(ctx context.Context, limit int64, skip int64) ([]*model.WatchUrl, int64, error)
}

type watchUrlRepository struct {
	collection *mongo.Collection
}

func (r *watchUrlRepository) Insert(ctx context.Context, watchUrl *model.WatchUrl) error {
	res, err := r.collection.InsertOne(ctx, watchUrl)
	watchUrl.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func (r *watchUrlRepository) InsertIfNotExists(ctx context.Context, watchUrl *model.WatchUrl) error {

	ret := r.collection.FindOneAndUpdate(ctx,
		primitive.M{"url": watchUrl.Url, "isList": watchUrl.IsList},
		primitive.M{"$setOnInsert": watchUrl},
		options.FindOneAndUpdate().SetUpsert(true))

	return ret.Err()
}

func (r *watchUrlRepository) Update(ctx context.Context, watchUrl *model.WatchUrl) error {
	_, err := r.collection.UpdateOne(ctx, primitive.M{"_id": watchUrl.ID}, primitive.M{"$set": watchUrl})
	return err
}

func (r *watchUrlRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, primitive.M{"_id": id})
	return err
}

func (r *watchUrlRepository) FindById(ctx context.Context, id primitive.ObjectID) (*model.WatchUrl, error) {
	var watchUrl model.WatchUrl

	err := r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&watchUrl)
	return &watchUrl, err
}

func (r *watchUrlRepository) FindBy(ctx context.Context, by primitive.M) ([]*model.WatchUrl, error) {
	var watchUrls []*model.WatchUrl
	cursor, err := r.collection.Find(ctx, by)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(nil, &watchUrls)
	return watchUrls, err
}

func (r *watchUrlRepository) FindAll(ctx context.Context, limit int64, skip int64) ([]*model.WatchUrl, int64, error) {
	var watchUrls []*model.WatchUrl

	opts := options.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := r.collection.Find(ctx, primitive.M{}, opts)

	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &watchUrls)
	if err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, primitive.M{})

	if err != nil {
		return nil, 0, err
	}

	return watchUrls, total, err

}
