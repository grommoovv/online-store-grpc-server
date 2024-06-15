package server

import (
	_ "github.com/lib/pq"
	"log/slog"
	"online-store-server/internal/config"
	"online-store-server/internal/database/postgres"
	"online-store-server/internal/repository"
	grpc_server "online-store-server/internal/server/grpc"
	"online-store-server/internal/service"
)

type Server struct {
	GRPCServer *grpc_server.GrpcServer
}

func New(log *slog.Logger, grpcPort int, psqlCfg *config.PsqlConfig) *Server {

	psql := postgres.MustConnect(log, postgres.PSQL{
		Host:     psqlCfg.Host,
		Port:     psqlCfg.Port,
		Username: psqlCfg.Username,
		Password: psqlCfg.Password,
		DBName:   psqlCfg.DBName,
		SSLMode:  psqlCfg.SSLMode,
	})

	// TODO:
	repo := repository.New(log, psql)
	services := service.New(log, repo)

	grpcServer := grpc_server.New(log, services, grpcPort)

	return &Server{
		GRPCServer: grpcServer,
	}
}
