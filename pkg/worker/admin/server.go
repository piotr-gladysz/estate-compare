package admin

import (
	"errors"
	"fmt"
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type Server struct {
	watchUrlRepo db.WatchUrlRepository
	offerRepo    db.OfferRepository

	control ProcessorControl

	port int
	ip   string

	server   *grpc.Server
	listener net.Listener
}

func NewServer(port int, ip string, watchUrlRepo db.WatchUrlRepository, offerRepo db.OfferRepository, control ProcessorControl) *Server {
	return &Server{watchUrlRepo: watchUrlRepo, offerRepo: offerRepo, control: control, port: port, ip: ip}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	api.RegisterWatchUrlServiceServer(grpcServer, NewWatchUrlServer(s.watchUrlRepo))
	api.RegisterOfferServiceServer(grpcServer, NewOfferServer(s.offerRepo))
	api.RegisterProcessorServiceServer(grpcServer, NewProcessorServer(s.control))

	s.listener = lis
	s.server = grpcServer

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			slog.Error("failed to serve grpc", "error", err.Error())
		}
	}()
	return nil
}

func (s *Server) Close() error {
	if s.server == nil {
		err := errors.New("server not running")
		slog.Error(err.Error())
		return err
	}

	s.server.GracefulStop()
	return nil
}
