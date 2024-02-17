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

var ServerNotRunningError = errors.New("server not running")

type Server struct {
	db      db.DB
	control ProcessorControl

	port int
	ip   string

	server   *grpc.Server
	listener net.Listener
}

func NewServer(port int, ip string, db db.DB, control ProcessorControl) *Server {
	return &Server{db: db, control: control, port: port, ip: ip}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	api.RegisterWatchUrlServiceServer(grpcServer, NewWatchUrlServer(s.db.GetWatchUrlRepository()))
	api.RegisterOfferServiceServer(grpcServer, NewOfferServer(s.db.GetOfferRepository()))
	api.RegisterProcessorServiceServer(grpcServer, NewProcessorServer(s.control))
	api.RegisterConditionServiceServer(grpcServer, NewConditionServer(s.db.GetConditionRepository(), s.db.GetNotificationRepository()))

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

		slog.Error(ServerNotRunningError.Error())
		return ServerNotRunningError
	}

	s.server.GracefulStop()
	return nil
}
