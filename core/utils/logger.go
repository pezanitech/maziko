package utils

import (
	"log/slog"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/pezanitech/maziko/core/config"
)

// App wide logger instance
var Logger *slog.Logger
var loggerOnce sync.Once

// Initialize the logger if it hasn't been already
func InitLogger() *slog.Logger {
	loggerOnce.Do(func() {
		// load variables from .env
		if err := godotenv.Load(); err != nil {
			slog.Info(
				"No .env file found, using environment variables",
			)
		}

		// check if JSON logger is enabled
		if config.UseJSONLogger() {
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
	})

	return Logger
}
