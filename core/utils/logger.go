package utils

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// application wide logger instance
var Logger *slog.Logger

func InitLogger() {
	// load variables from .env
	if err := godotenv.Load(); err != nil {
		slog.Error("Failed to load .env file", "error", err)
	}

	if os.Getenv("JSON_LOGGER") == "true" {
		handler := slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		)

		Logger = slog.New(handler) // use JSON logger
		Logger.Info("Using JSON logger")
	} else {
		Logger = slog.Default() // use default logger
		Logger.Info("Using default logger")
	}
}
