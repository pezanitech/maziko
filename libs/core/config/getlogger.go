package config

import (
	"log/slog"
	"strings"
)

// ---- LOGGER CONFIG GETTERS ----

// Returns the logger type (text, json, concise)
func GetLoggerType() string {
	logType := strings.ToLower(AppConfig.Logger.Type)

	// Default to text if not specified or invalid
	if logType != "text" && logType != "json" && logType != "concise" {
		return "text"
	}

	return logType
}

// Returns the configured log level
func GetLogLevel() slog.Level {
	level := strings.ToLower(AppConfig.Logger.Level)

	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		// use info if unspecified or invalid
		return slog.LevelInfo
	}
}
