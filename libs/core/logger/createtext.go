package logger

import (
	"log/slog"
	"os"
)

// Creates a text formatted logger
func createTextLogger(logLevel slog.Level) (*slog.Logger, string) {
	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: logLevel,
		},
	)
	return slog.New(handler), "text"
}
