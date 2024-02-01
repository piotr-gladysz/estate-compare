package testutils

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ db.OfferRepository = (*OfferRepositoryMock)(nil)

// OfferRepositoryMock is a mock implementation of db.OfferRepository
type OfferRepositoryMock struct {
	Callback     func(method string, args ...any)
	ReturnSingle *db.Offer
	ReturnMany   []*db.Offer
	ReturnError  error
}

func (t *OfferRepositoryMock) Insert(ctx context.Context, offer *db.Offer) error {
	if t.Callback != nil {
		t.Callback("Insert", offer)
	}

	return t.ReturnError
}

func (t *OfferRepositoryMock) Update(ctx context.Context, offer *db.Offer) error {
	if t.Callback != nil {
		t.Callback("Update", offer)
	}
	return t.ReturnError
}

func (t *OfferRepositoryMock) Delete(ctx context.Context, id primitive.ObjectID) error {
	if t.Callback != nil {
		t.Callback("Delete", id)
	}
	return t.ReturnError
}

func (t *OfferRepositoryMock) FindById(ctx context.Context, id primitive.ObjectID) (*db.Offer, error) {
	if t.Callback != nil {
		t.Callback("FindById", id)
	}
	return t.ReturnSingle, t.ReturnError
}

func (t *OfferRepositoryMock) FindBy(ctx context.Context, by primitive.M) ([]*db.Offer, error) {
	if t.Callback != nil {
		t.Callback("FindBy", by)
	}
	return t.ReturnMany, t.ReturnError
}

func (t *OfferRepositoryMock) FindAll(ctx context.Context, limit int64, skip int64) ([]*db.Offer, error) {
	if t.Callback != nil {
		t.Callback("FindAll", limit, skip)
	}
	return t.ReturnMany, t.ReturnError
}
