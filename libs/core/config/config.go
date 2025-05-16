package config

import (
	"encoding/json"
	"log/slog"
	"os"
)

type Config struct {
	App struct {
		Name string `json:"name"`
		URL  string `json:"url"`
		Port int    `json:"port"`
	} `json:"app"`

	Build struct {
		Prefix  string `json:"prefix"`
		Dir     string `json:"dir"`
		TempDir string `json:"tempDir"`
		SSRDir  string `json:"ssrDir"`
	} `json:"build"`

	Vite struct {
		ManifestFile      string `json:"manifestFile"`
		HotFile           string `json:"hotFile"`
		DetectionAttempts int    `json:"detectionAttempts"`
		DetectionInterval int    `json:"detectionInterval"`
	} `json:"vite"`

	Paths struct {
		Routes string `json:"routes"`
		Public string `json:"public"`
		Gen    string `json:"gen"`
	} `json:"paths"`

	Package struct {
		Prefix string `json:"prefix"`
	} `json:"package"`

	Logger struct {
		Type  string `json:"type"` // "text","json","concise"
		Level string `json:"level"`
	} `json:"logger"`

	Dev struct {
		RootDir        string   `json:"rootDir"`
		ExcludeRegexes []string `json:"excludeRegexes"`
		ExcludeDirs    []string `json:"excludeDirs"`
		IncludeExts    []string `json:"includeExts"`
		BuildDelay     int      `json:"buildDelay"`
	} `json:"dev"`
}

// Global configuration instance
var AppConfig Config

// Constants
const ConfigPath = "maziko.json"
const EnvPath = ".env"
const EnvExamplePath = ".env.example"

// Load configuration from JSON file
func Initialize() error {
	if err := initEnvFile(); err != nil {
		return err
	}

	if err := loadConfigFile(); err != nil {
		return err
	}

	applyEnvOverrides()

	return nil
}

// Loads and parses the configuration file
func loadConfigFile() error {
	slog.Info(
		"Loading configuration file",
		"path", ConfigPath,
	)

	file, err := os.Open(ConfigPath)
	if err != nil {
		slog.Error(
			"Failed to open config file",
			"path", ConfigPath,
			"error", err,
		)
		return err
	}
	defer file.Close()

	// decode JSON into global AppConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		slog.Error(
			"Failed to parse JSON config",
			"path", ConfigPath,
			"error", err,
		)
		return err
	}

	slog.Debug(
		"Configuration loaded successfully",
		"path", ConfigPath,
	)

	// log key configuration values
	slog.Debug(
		"App configuration",
		"name", AppConfig.App.Name,
		"url", AppConfig.App.URL,
		"port", AppConfig.App.Port,
	)

	slog.Debug(
		"Logger configuration",
		"type", AppConfig.Logger.Type,
		"level", AppConfig.Logger.Level,
	)

	return nil
}
