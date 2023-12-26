package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Offer struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	SiteId     string    `json:"siteId" bson:"siteId"`
	Site       string    `json:"site" bson:"site"`
	UpdateTime time.Time `json:"updateTime" bson:"updateTime"`

	Name           string   `json:"name" bson:"name"`
	Url            string   `json:"url" bson:"url"`
	Area           float32  `json:"area" bson:"area"`
	Rooms          int      `json:"rooms" bson:"rooms"`
	Floor          int      `json:"floor" bson:"floor"`
	BuildingFloors int      `json:"buildingFloors" bson:"buildingFloors"`
	Year           int      `json:"year" bson:"year"`
	Heating        string   `json:"heating" bson:"heating"`
	Market         string   `json:"market" bson:"market"`
	Window         string   `json:"window" bson:"window"`
	Elevator       bool     `json:"elevator" bson:"elevator"`
	Balcony        bool     `json:"balcony" bson:"balcony"`
	Media          []string `json:"media" bson:"media"`

	History []*OfferHistory `json:"history" bson:"history"`
}

type OfferHistory struct {
	UpdateTime time.Time `json:"updateTime" bson:"updateTime"`
	Price      int       `json:"price" bson:"price"`
}

type OfferRepository struct {
	collection *mongo.Collection
}

func (r *OfferRepository) Insert(ctx context.Context, offer *Offer) error {
	_, err := r.collection.InsertOne(ctx, offer)
	return err
}

func (r *OfferRepository) Update(ctx context.Context, offer *Offer) error {
	_, err := r.collection.UpdateOne(ctx, primitive.M{"_id": offer.ID}, primitive.M{"$set": offer})
	return err
}

func (r *OfferRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, primitive.M{"_id": id})
	return err
}

func (r *OfferRepository) FindById(ctx context.Context, id primitive.ObjectID) (*Offer, error) {
	var offer Offer
	err := r.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&offer)
	return &offer, err
}

func (r *OfferRepository) FindBy(ctx context.Context, by primitive.M) ([]*Offer, error) {
	var offers []*Offer
	cursor, err := r.collection.Find(ctx, by)
	if err != nil {
		return nil, err
	}

	err = cursor.All(nil, &offers)
	return offers, err
}
