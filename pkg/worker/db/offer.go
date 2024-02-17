package db

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ OfferRepository = (*offerRepository)(nil)

type OfferRepository interface {
	Insert(ctx context.Context, offer *model.Offer) error
	Update(ctx context.Context, offer *model.Offer) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindById(ctx context.Context, id primitive.ObjectID) (*model.Offer, error)
	FindBy(ctx context.Context, by primitive.M) ([]*model.Offer, error)
	FindAll(ctx context.Context, limit int64, skip int64) ([]*model.Offer, error)
}

type offerRepository struct {
	collection *mongo.Collection
}

func (r *offerRepository) Insert(ctx context.Context, offer *model.Offer) error {
	res, err := r.collection.InsertOne(ctx, offer)
	offer.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func (r *offerRepository) Update(ctx context.Context, offer *model.Offer) error {
	_, err := r.collection.UpdateOne(ctx, primitive.M{"_id": offer.ID}, primitive.M{"$set": offer})
	return err
}

func (r *offerRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, primitive.M{"_id": id})
	return err
}

func (r *offerRepository) FindById(ctx context.Context, id primitive.ObjectID) (*model.Offer, error) {
	var offer model.Offer
	err := r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&offer)
	return &offer, err
}

func (r *offerRepository) FindBy(ctx context.Context, by primitive.M) ([]*model.Offer, error) {
	var offers []*model.Offer
	cursor, err := r.collection.Find(ctx, by)
	if err != nil {
		return nil, err
	}

	err = cursor.All(nil, &offers)
	return offers, err
}

func (r *offerRepository) FindAll(ctx context.Context, limit int64, skip int64) ([]*model.Offer, error) {
	var offers []*model.Offer
	cursor, err := r.collection.Find(ctx, primitive.M{}, options.Find().SetLimit(limit).SetSkip(skip))
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(nil, &offers)
	return offers, err
}
