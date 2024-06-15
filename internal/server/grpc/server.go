package grpc_server

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"online-store-server/internal/service"
	product_grpc "online-store-server/internal/transport/grpc/product"
)

type GrpcServer struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       int
}

func New(log *slog.Logger, services *service.Service, port int) *GrpcServer {

	grpcServer := grpc.NewServer()

	product_grpc.Register(grpcServer, *services)

	return &GrpcServer{
		log:        log,
		grpcServer: grpcServer,
		port:       port,
	}
}

// MustRun runs gRPC server and panics if any error occurs.
func (s *GrpcServer) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *GrpcServer) Run() error {
	const op = "grpc_server.Run"

	log := s.log.With(slog.String("op", op), slog.Int("port", s.port))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", listener.Addr().String()))

	if err := s.grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop stops gRPC server.
func (s *GrpcServer) Stop() {
	const op = "grpc_server.Stop"

	s.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", s.port))

	s.grpcServer.GracefulStop()
}
