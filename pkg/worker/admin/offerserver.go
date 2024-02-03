package admin

import (
	"context"
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

var _ api.OfferServiceServer = (*OfferServer)(nil)

type OfferServer struct {
	api.UnimplementedOfferServiceServer
	repo db.OfferRepository
}

func NewOfferServer(repo db.OfferRepository) *OfferServer {
	return &OfferServer{repo: repo}
}

func (o *OfferServer) GetOffer(ctx context.Context, request *api.GetOfferRequest) (*api.OfferResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		slog.Error("failed to parse id", "error", err.Error())
		return nil, err
	}

	ret, err := o.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return o.offerToResponse(ret, true), nil

}
func (o *OfferServer) GetOffers(ctx context.Context, request *api.GetOffersRequest) (*api.OfferListResponse, error) {
	limit := request.PageSize
	skip := (request.Page - 1) * request.PageSize

	ret, err := o.repo.FindAll(ctx, int64(skip), int64(limit))

	if err != nil {
		return nil, err
	}

	offerResponses := make([]*api.OfferResponse, len(ret))

	for i, offer := range ret {
		offerResponses[i] = o.offerToResponse(offer, false)
	}

	return &api.OfferListResponse{
		Offers: offerResponses,
	}, nil

}

func (o *OfferServer) offerToResponse(offer *db.Offer, withHistory bool) *api.OfferResponse {
	ret := &api.OfferResponse{
		Id:             offer.ID.Hex(),
		SiteId:         offer.SiteId,
		Site:           offer.Site,
		Created:        offer.Created.Time().Unix(),
		Updated:        offer.Updated.Time().Unix(),
		Name:           offer.Name,
		Url:            offer.Url,
		Area:           offer.Area,
		Rooms:          int32(offer.Rooms),
		Floor:          int32(offer.Floor),
		BuildingFloors: int32(offer.BuildingFloors),
		Year:           int32(offer.Year),
		Heating:        offer.Heating,
		Market:         offer.Market,
		Window:         offer.Window,
		Elevator:       offer.Elevator,
		Balcony:        offer.Balcony,
		Media:          offer.Media,
		History:        []*api.OfferHistory{},
	}

	if withHistory {
		for _, h := range offer.History {
			ret.History = append(ret.History, &api.OfferHistory{
				Updated: h.Updated.Time().Unix(),
				Price:   int32(h.Price),
			})
		}
	}

	return ret
}
