package errors

import (
	"net/http"
	"os"

	"github.com/pezanitech/maziko/libs/core/logger"
)

// Handles server errors with logging and response
func HandleServerErr(w http.ResponseWriter, err error) {
	logger.Log.Error(
		"http error",
		"err", err,
	)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("server error"))
}

// Logs the error and exits the application
func HandleFatalError(message string, err error, fields ...any) {
	logger.Log.Error(
		message,
		append([]any{"error", err}, fields...)...,
	)

	os.Exit(1)
}
