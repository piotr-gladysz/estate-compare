package admin

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
	"time"
)

type WatchUrlServer struct {
	api.UnimplementedWatchUrlServiceServer
	repo db.WatchUrlRepository
}

func NewWatchUrlServer(repo db.WatchUrlRepository) *WatchUrlServer {
	return &WatchUrlServer{repo: repo}
}

func (w *WatchUrlServer) AddUrl(ctx context.Context, request *api.AddUrlRequest) (*api.UrlResponse, error) {

	now := primitive.NewDateTimeFromTime(time.Now())
	url := &model.WatchUrl{
		Url:      request.Url,
		IsList:   request.IsList,
		Created:  now,
		Updated:  now,
		Disabled: false,
	}

	err := w.repo.Insert(ctx, url)

	if err != nil {
		return nil, err
	}

	return w.watchUrlToResponse(url), nil
}

func (w *WatchUrlServer) SetState(ctx context.Context, request *api.SetStateRequest) (*api.UrlResponse, error) {
	now := primitive.NewDateTimeFromTime(time.Now())
	id, err := primitive.ObjectIDFromHex(request.Id)

	if err != nil {
		slog.Error("failed to parse id", "error", err.Error())
		return nil, err
	}

	url, err := w.repo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	url.Updated = now
	url.Disabled = request.IsDisabled

	err = w.repo.Update(ctx, url)

	if err != nil {
		return nil, err
	}

	return w.watchUrlToResponse(url), nil
}

func (w *WatchUrlServer) GetUrls(ctx context.Context, request *api.GetUrlsRequest) (*api.UrlListResponse, error) {
	limit := request.PageSize
	skip := (request.Page - 1) * request.PageSize

	urls, total, err := w.repo.FindAll(ctx, int64(skip), int64(limit))

	if err != nil {
		return nil, err
	}

	urlResponses := make([]*api.UrlResponse, len(urls))

	for i, url := range urls {
		urlResponses[i] = w.watchUrlToResponse(url)
	}

	return &api.UrlListResponse{
		Urls:  urlResponses,
		Total: total,
	}, nil
}

func (w *WatchUrlServer) watchUrlToResponse(url *model.WatchUrl) *api.UrlResponse {
	return &api.UrlResponse{
		Url:        url.Url,
		IsList:     url.IsList,
		Id:         url.ID.Hex(),
		IsDisabled: url.Disabled,
		Created:    url.Created.Time().Unix(),
		Updated:    url.Updated.Time().Unix(),
	}
}
