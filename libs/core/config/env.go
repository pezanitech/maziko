package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// Default .env content
const DefaultEnv = `
# ---- Required ----
APP_URL=http://localhost:3080

# ---- Optional ----
# text (default), json, concise
LOGGER_TYPE=concise

# debug, info (default), warn, error
LOG_LEVEL=debug
`

// Checks if .env exists, else creates it
func createEnvFile(envPath string, envExamplePath string, defaultEnv string) error {
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		var envContent []byte

		slog.Debug(
			"No .env file found, creating one",
			"path", envPath,
		)

		// check if .env.example exists
		if _, err := os.Stat(envExamplePath); err == nil {
			slog.Debug(
				"Using .env.example as template",
				"path", envExamplePath,
			)

			envContent, err = os.ReadFile(envExamplePath) // exists, store content
			if err != nil {
				slog.Error(
					"Failed to read .env.example",
					"error", err,
				)
				return err
			}
		} else {
			slog.Debug(
				"No .env.example found, using default template",
			)
			envContent = []byte(defaultEnv) // doesn't exist, use default
		}

		// create .env file with the content
		if err := os.WriteFile(envPath, envContent, 0644); err != nil {
			slog.Error(
				"Failed to create .env file",
				"path", envPath,
				"error", err,
			)
			return err
		}

		slog.Info(
			"Created new .env file",
			"path", envPath,
		)
	} else {
		slog.Debug(
			"Found existing .env file",
			"path", envPath,
		)
	}

	return nil
}

// Initializes .env file, then loads it
func initEnvFile() error {
	slog.Info(
		"Initializing environment configuration",
	)

	if err := createEnvFile( // attempt to create .env
		EnvPath,
		EnvExamplePath,
		DefaultEnv,
	); err != nil {
		slog.Error(
			"Failed to initialize .env file",
			"error", err,
		)
		return err
	}

	// load environment variables from .env file
	if err := godotenv.Load(EnvPath); err != nil {
		slog.Error(
			"Failed to load environment variables",
			"path", EnvPath,
			"error", err,
		)
		return err
	}

	slog.Debug(
		"Successfully loaded environment variables",
		"path", EnvPath,
	)
	return nil
}

// Applies env overrides if present, (overrides maziko.json)
func applyEnvOverrides() {
	slog.Debug(
		"Checking for environment variable overrides",
	)

	if appURL := os.Getenv("APP_URL"); appURL != "" {
		slog.Debug(
			"Applying APP_URL override",
			"value", appURL,
		)
		AppConfig.App.URL = appURL
	}

	if loggerType := os.Getenv("LOGGER_TYPE"); loggerType != "" {
		slog.Debug(
			"Applying LOGGER_TYPE override",
			"value", loggerType,
		)
		AppConfig.Logger.Type = loggerType
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		slog.Debug(
			"Applying LOG_LEVEL override",
			"value", logLevel,
		)
		AppConfig.Logger.Level = logLevel
	}

	slog.Info(
		"Environment configuration complete",
	)
}
