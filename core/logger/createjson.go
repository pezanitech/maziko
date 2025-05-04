package logger

import (
	"log/slog"
	"os"
)

// Creates a JSON formatted logger
func createJSONLogger(logLevel slog.Level) (*slog.Logger, string) {
	handler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: logLevel,
		},
	)
	return slog.New(handler), "JSON"
}
