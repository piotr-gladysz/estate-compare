package admin

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
)

var _ api.WatchUrlServiceServer = (*WatchUrlServer)(nil)

type WatchUrlServer struct {
	api.UnimplementedWatchUrlServiceServer
	repo db.WatchUrlRepository
}

func NewWatchUrlServer(repo db.WatchUrlRepository) *WatchUrlServer {
	return &WatchUrlServer{repo: repo}
}

func (w WatchUrlServer) AddUrl(ctx context.Context, request *api.AddUrlRequest) (*api.UrlResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w WatchUrlServer) SetState(ctx context.Context, request *api.SetStateRequest) (*api.UrlResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w WatchUrlServer) GetUrls(request *api.GetUrlsRequest, server api.WatchUrlService_GetUrlsServer) error {
	//TODO implement me
	panic("implement me")
}
