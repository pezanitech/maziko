package config

import (
	"encoding/json"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Application configuration variables
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

const ConfigPath = "maziko.json"
const EnvPath = ".env"
const EnvExamplePath = ".env.example"

// Default content for the .env file
const defaultEnvContent = `# Required
APP_URL=http://localhost:3080

# Optional
# Logger type options: text (default), json, concise
LOGGER_TYPE=concise
# Log level options: debug, info, warn, error (default: info)
LOG_LEVEL=debug
`

// Load configuration from JSON file
func Initialize() error {
	// check if .env file exists, if not create one
	if _, err := os.Stat(EnvPath); os.IsNotExist(err) {
		var envContent []byte

		// check if .env.example exists
		if _, err := os.Stat(EnvExamplePath); err == nil {
			// store .env.example content
			envContent, err = os.ReadFile(EnvExamplePath)
			if err != nil {
				return err
			}
		} else {
			// use default content
			envContent = []byte(defaultEnvContent)
		}

		// create .env with the content
		err = os.WriteFile(EnvPath, envContent, 0644)
		if err != nil {
			return err
		}
	}

	godotenv.Load() // load .env

	file, err := os.Open(ConfigPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// decode JSON into global AppConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		return err
	}

	// apply environment variable overrides if present
	if appURL := os.Getenv("APP_URL"); appURL != "" {
		AppConfig.App.URL = appURL
	}

	if loggerType := os.Getenv("LOGGER_TYPE"); loggerType != "" {
		AppConfig.Logger.Type = loggerType
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		AppConfig.Logger.Level = logLevel
	}

	return nil
}

// ---- App config getters ----

// Returns app name
func GetAppName() string {
	return AppConfig.App.Name
}

// Returns app URL
func GetAppURL() string {
	return AppConfig.App.URL
}

// Returns app port
func GetAppPort() int {
	return AppConfig.App.Port
}

// ---- Build config getters ----

// Returns build prefix
func GetBuildPrefix() string {
	return AppConfig.Build.Prefix
}

// Returns path to build directory
func GetBuildDir() string {
	return AppConfig.Build.Dir
}

// Returns path to temp directory
func GetTempDir() string {
	return AppConfig.Build.TempDir
}

// Returns path to the SSR build directory
func GetSSRDir() string {
	return AppConfig.Build.SSRDir
}

// ---- Vite config getters ----

// Returns path to Vite manifest file
func GetViteManifestFile() string {
	return AppConfig.Vite.ManifestFile
}

// Returns path to Vite hot file
func GetHotFile() string {
	return AppConfig.Vite.HotFile
}

// Returns maximum number of attempts to detect Vit
func MaxViteDetectionAttempts() int {
	return AppConfig.Vite.DetectionAttempts
}

// Returns interval between Vite detection attempts
func ViteDetectionInterval() time.Duration {
	return time.Duration(AppConfig.Vite.DetectionInterval) * time.Millisecond
}

// ---- Path config getters ----

// Returns path to generation directory
func GetGenDir() string {
	return AppConfig.Paths.Gen
}

// Returns path to routes directory
func GetRoutesDir() string {
	return AppConfig.Paths.Routes
}

// Returns path to public directory
func GetPublicDir() string {
	return AppConfig.Paths.Public
}

// ---- Package config getters ----

// Returns package prefix (for generated imports)
func GetPackagePrefix() string {
	return AppConfig.Package.Prefix
}

// ---- Logger config getters ----

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

// ---- Dev mode config getters ----

// Returns the root directory for file watching
func GetDevRootDir() string {
	return AppConfig.Dev.RootDir
}

// Returns list of regex patterns to exclude from file watching
func GetDevExcludeRegexes() []string {
	return AppConfig.Dev.ExcludeRegexes
}

// Returns list of directories to exclude from file watching
func GetDevExcludeDirs() []string {
	return AppConfig.Dev.ExcludeDirs
}

// Returns list of file extensions to include in file watching
func GetDevIncludeExts() []string {
	return AppConfig.Dev.IncludeExts
}

// Returns the build delay in milliseconds
func GetDevBuildDelay() time.Duration {
	return time.Duration(AppConfig.Dev.BuildDelay) * time.Millisecond
}
