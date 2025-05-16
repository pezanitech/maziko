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

// InitLogger initializes the logger if it hasn't been
func InitLogger() *slog.Logger {
	loggerOnce.Do(func() {
		loadEnvVariables() // load .env
		logLevel := config.GetLogLevel()
		loggerType := config.GetLoggerType()
		Log = createLogger(loggerType, logLevel)
	})

	return Log
}

// InitLoggerWithOptions initializes a logger with logger type and level
func InitLoggerWithOptions(loggerType string, logLevelStr string) *slog.Logger {
	loadEnvVariables() // load .env

	// convert string level to slog.Level
	var logLevel slog.Level
	switch logLevelStr {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo // Default to info if unrecognized
	}

	return createLogger(loggerType, logLevel)
}

// loadEnvVariables loads variables from .env file if available
func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		slog.Info(
			"No .env file found, using environment variables",
		)
	}
}
