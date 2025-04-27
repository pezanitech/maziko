package utils

import (
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// Shared Logger instance used across the application
var Logger *slog.Logger

func InitLogger() {
	loadEnvFile()

	useJSONLogger := os.Getenv("JSON_LOGGER") == "true"

	if useJSONLogger {
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})

		Logger = slog.New(handler)
		Logger.Info("Using JSON logger")
	} else {
		Logger = slog.Default()
		Logger.Info("Using default logger")
	}
}

// Loads environment variables from .env file
func loadEnvFile() {
	// determine the project root directory
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "../..")

	envPath := filepath.Join(projectRoot, ".env")
	err := godotenv.Load(envPath)

	if err != nil {
		// use slog since logger is not initialized
		slog.Error("Failed to load .env file", "error", err)
	}
}
