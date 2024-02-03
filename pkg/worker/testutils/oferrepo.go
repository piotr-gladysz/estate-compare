package testutils

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ db.OfferRepository = (*OfferRepositoryMock)(nil)

// OfferRepositoryMock is a mock implementation of db.OfferRepository
type OfferRepositoryMock struct {
	Callback     func(this *OfferRepositoryMock, method string, args ...any)
	ReturnSingle *db.Offer
	ReturnMany   []*db.Offer
	ReturnError  error
}

func (m *OfferRepositoryMock) Insert(ctx context.Context, offer *db.Offer) error {
	if m.Callback != nil {
		m.Callback(m, "Insert", offer)
	}

	return m.ReturnError
}

func (m *OfferRepositoryMock) Update(ctx context.Context, offer *db.Offer) error {
	if m.Callback != nil {
		m.Callback(m, "Update", offer)
	}
	return m.ReturnError
}

func (m *OfferRepositoryMock) Delete(ctx context.Context, id primitive.ObjectID) error {
	if m.Callback != nil {
		m.Callback(m, "Delete", id)
	}
	return m.ReturnError
}

func (m *OfferRepositoryMock) FindById(ctx context.Context, id primitive.ObjectID) (*db.Offer, error) {
	if m.Callback != nil {
		m.Callback(m, "FindById", id)
	}
	return m.ReturnSingle, m.ReturnError
}

func (m *OfferRepositoryMock) FindBy(ctx context.Context, by primitive.M) ([]*db.Offer, error) {
	if m.Callback != nil {
		m.Callback(m, "FindBy", by)
	}
	return m.ReturnMany, m.ReturnError
}

func (m *OfferRepositoryMock) FindAll(ctx context.Context, limit int64, skip int64) ([]*db.Offer, error) {
	if m.Callback != nil {
		m.Callback(m, "FindAll", limit, skip)
	}
	return m.ReturnMany, m.ReturnError
}
