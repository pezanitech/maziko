package utils

import (
	"net/http"
	"os"

	"github.com/pezanitech/maziko/core/logger"
)

// Handles HTTP server errors with logging and response
func HandleServerErr(w http.ResponseWriter, err error) {
	logger.Logger.Error("http error", "err", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("server error"))
}

// Logs the error and exits the application
func HandleFatalError(message string, err error, fields ...any) {
	logger.Logger.Error(message, append([]any{"error", err}, fields...)...)
	os.Exit(1)
}
