package testutils

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ db.WatchUrlRepository = (*WatchUrlRepositoryMock)(nil)

type WatchUrlRepositoryMock struct {
	Callback     func(this *WatchUrlRepositoryMock, method string, args ...any)
	ReturnSingle *model.WatchUrl
	ReturnMany   []*model.WatchUrl
	ReturnError  error
}

func (w *WatchUrlRepositoryMock) Insert(ctx context.Context, watchUrl *model.WatchUrl) error {
	if w.Callback != nil {
		w.Callback(w, "Insert", watchUrl)
	}

	return w.ReturnError
}

func (w *WatchUrlRepositoryMock) InsertIfNotExists(ctx context.Context, watchUrl *model.WatchUrl) error {
	if w.Callback != nil {
		w.Callback(w, "InsertIfNotExists", watchUrl)
	}

	return w.ReturnError
}

func (w *WatchUrlRepositoryMock) Update(ctx context.Context, watchUrl *model.WatchUrl) error {
	if w.Callback != nil {
		w.Callback(w, "Update", watchUrl)
	}

	return w.ReturnError
}

func (w *WatchUrlRepositoryMock) Delete(ctx context.Context, id primitive.ObjectID) error {
	if w.Callback != nil {
		w.Callback(w, "Delete", id)
	}

	return w.ReturnError
}

func (w *WatchUrlRepositoryMock) FindById(ctx context.Context, id primitive.ObjectID) (*model.WatchUrl, error) {
	if w.Callback != nil {
		w.Callback(w, "FindById", id)
	}

	return w.ReturnSingle, w.ReturnError
}

func (w *WatchUrlRepositoryMock) FindBy(ctx context.Context, by primitive.M) ([]*model.WatchUrl, error) {
	if w.Callback != nil {
		w.Callback(w, "FindBy", by)
	}

	return w.ReturnMany, w.ReturnError
}

func (w *WatchUrlRepositoryMock) FindAll(ctx context.Context, limit int64, skip int64) ([]*model.WatchUrl, int64, error) {
	if w.Callback != nil {
		w.Callback(w, "FindAll", limit, skip)
	}

	return w.ReturnMany, int64(len(w.ReturnMany)), w.ReturnError
}
