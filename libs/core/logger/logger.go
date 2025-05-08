package logger

import (
	"log/slog"
	"sync"

	"github.com/joho/godotenv"
	"github.com/pezanitech/maziko/libs/core/config"
)

// App wide logger instance
var Log *slog.Logger
var loggerOnce sync.Once

// Initialize the logger if it hasn't been
func InitLogger() *slog.Logger {
	loggerOnce.Do(func() {
		loadEnvVariables() // load .env
		logLevel := config.GetLogLevel()
		loggerType := config.GetLoggerType()
		Log = createLogger(loggerType, logLevel)
	})

	return Log
}

// Loads variables from .env file if available
func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		slog.Info(
			"No .env file found, using environment variables",
		)
	}
}
