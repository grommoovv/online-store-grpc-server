package logger

import (
	"log/slog"
	"online-store-server/internal/lib/logger/custom/slogpretty"
	"os"
)

const (
	envLocal = "local"
	endDev   = "dev"
	endProd  = "prod"
)

func Init(envMode string) *slog.Logger {
	var logger *slog.Logger

	switch envMode {
	case envLocal:
		logger = InitPretty()
	case endDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case endProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}

func InitPretty() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
