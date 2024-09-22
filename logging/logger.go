package logging

import (
	"log/slog"
	"os"
)

func GetLogger() *slog.Logger {

	appEnv := os.Getenv("APP_ENV")

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)

	if appEnv == "production" {
		opts := &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		}
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)

	logger.Info("environem set APP_ENV " + appEnv)

	return logger

}
