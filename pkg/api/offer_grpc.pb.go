// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: offer.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OfferService_GetOffer_FullMethodName  = "/api.OfferService/GetOffer"
	OfferService_GetOffers_FullMethodName = "/api.OfferService/GetOffers"
)

// OfferServiceClient is the client API for OfferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OfferServiceClient interface {
	GetOffer(ctx context.Context, in *GetOfferRequest, opts ...grpc.CallOption) (*OfferResponse, error)
	GetOffers(ctx context.Context, in *GetOffersRequest, opts ...grpc.CallOption) (*OfferListResponse, error)
}

type offerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOfferServiceClient(cc grpc.ClientConnInterface) OfferServiceClient {
	return &offerServiceClient{cc}
}

func (c *offerServiceClient) GetOffer(ctx context.Context, in *GetOfferRequest, opts ...grpc.CallOption) (*OfferResponse, error) {
	out := new(OfferResponse)
	err := c.cc.Invoke(ctx, OfferService_GetOffer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offerServiceClient) GetOffers(ctx context.Context, in *GetOffersRequest, opts ...grpc.CallOption) (*OfferListResponse, error) {
	out := new(OfferListResponse)
	err := c.cc.Invoke(ctx, OfferService_GetOffers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OfferServiceServer is the server API for OfferService service.
// All implementations must embed UnimplementedOfferServiceServer
// for forward compatibility
type OfferServiceServer interface {
	GetOffer(context.Context, *GetOfferRequest) (*OfferResponse, error)
	GetOffers(context.Context, *GetOffersRequest) (*OfferListResponse, error)
	mustEmbedUnimplementedOfferServiceServer()
}

// UnimplementedOfferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOfferServiceServer struct {
}

func (UnimplementedOfferServiceServer) GetOffer(context.Context, *GetOfferRequest) (*OfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffer not implemented")
}
func (UnimplementedOfferServiceServer) GetOffers(context.Context, *GetOffersRequest) (*OfferListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffers not implemented")
}
func (UnimplementedOfferServiceServer) mustEmbedUnimplementedOfferServiceServer() {}

// UnsafeOfferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OfferServiceServer will
// result in compilation errors.
type UnsafeOfferServiceServer interface {
	mustEmbedUnimplementedOfferServiceServer()
}

func RegisterOfferServiceServer(s grpc.ServiceRegistrar, srv OfferServiceServer) {
	s.RegisterService(&OfferService_ServiceDesc, srv)
}

func _OfferService_GetOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfferServiceServer).GetOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OfferService_GetOffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfferServiceServer).GetOffer(ctx, req.(*GetOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfferService_GetOffers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOffersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfferServiceServer).GetOffers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OfferService_GetOffers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfferServiceServer).GetOffers(ctx, req.(*GetOffersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OfferService_ServiceDesc is the grpc.ServiceDesc for OfferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OfferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.OfferService",
	HandlerType: (*OfferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOffer",
			Handler:    _OfferService_GetOffer_Handler,
		},
		{
			MethodName: "GetOffers",
			Handler:    _OfferService_GetOffers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "offer.proto",
}