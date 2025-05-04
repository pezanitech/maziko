package logger

import (
	"log/slog"
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
		loadEnvVariables()

		// get logger configuration
		logLevel := config.GetLogLevel()
		loggerType := config.GetLoggerType()

		// create and configure logger based on type
		Logger = createLogger(loggerType, logLevel)
	})

	return Logger
}

// Loads variables from .env file if available
func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found, using environment variables")
	}
}
