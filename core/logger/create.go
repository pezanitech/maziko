package logger

import "log/slog"

// Creates a new logger based on loggerType and logLevel
func createLogger(loggerType string, logLevel slog.Level) *slog.Logger {
	var logger *slog.Logger
	var handlerType string

	switch loggerType {
	case "json":
		logger, handlerType = createJSONLogger(logLevel)
	case "concise":
		logger, handlerType = createConciseLogger(logLevel)
	default: // "text" or any unrecognized type
		logger, handlerType = createTextLogger(logLevel)
	}

	logger.Info(
		"Using "+handlerType+" logger",
		"level", logLevel,
	)

	return logger
}
