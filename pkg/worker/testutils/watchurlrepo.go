package testutils

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WatchUrlRepositoryMock struct {
	Callback     func(this *WatchUrlRepositoryMock, method string, args ...any)
	ReturnSingle *db.WatchUrl
	ReturnMany   []*db.WatchUrl
	ReturnError  error
}

func (w *WatchUrlRepositoryMock) Insert(ctx context.Context, watchUrl *db.WatchUrl) error {
	if w.Callback != nil {
		w.Callback(w, "Insert", watchUrl)
	}

	return w.ReturnError
}

func (w *WatchUrlRepositoryMock) InsertIfNotExists(ctx context.Context, watchUrl *db.WatchUrl) error {
	if w.Callback != nil {
		w.Callback(w, "InsertIfNotExists", watchUrl)
	}

	return w.ReturnError
}

func (w *WatchUrlRepositoryMock) Update(ctx context.Context, watchUrl *db.WatchUrl) error {
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

func (w *WatchUrlRepositoryMock) FindById(ctx context.Context, id primitive.ObjectID) (*db.WatchUrl, error) {
	if w.Callback != nil {
		w.Callback(w, "FindById", id)
	}

	return w.ReturnSingle, w.ReturnError
}

func (w *WatchUrlRepositoryMock) FindBy(ctx context.Context, by primitive.M) ([]*db.WatchUrl, error) {
	if w.Callback != nil {
		w.Callback(w, "FindBy", by)
	}

	return w.ReturnMany, w.ReturnError
}

func (w *WatchUrlRepositoryMock) FindAll(ctx context.Context, limit int64, skip int64) ([]*db.WatchUrl, error) {
	if w.Callback != nil {
		w.Callback(w, "FindAll", limit, skip)
	}

	return w.ReturnMany, w.ReturnError
}
