package main

import (
	"log/slog"
	"online-store-server/internal/config"
	"online-store-server/internal/lib/logger"
	"online-store-server/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoadConfig()

	log := logger.Init(cfg.Env)

	log.Info("starting application...", slog.Any("config", cfg))

	srv := server.New(log, cfg.Grpc.Port, &cfg.Psql)

	go func() {
		srv.GRPCServer.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-stop

	log.Info("gracefully stopping application...", slog.String("signal", sign.String()))

	srv.GRPCServer.Stop()

	log.Info("application gracefully stopped")
}
