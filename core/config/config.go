package config

import (
	"encoding/json"
	"os"
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
		UseJSON bool `json:"useJson"`
	} `json:"logger"`
}

// Global configuration instance
var AppConfig Config

// Load configuration from JSON file
func Initialize() error {
	godotenv.Load() // load .env
	configPath := "maziko.config.json"

	// Allow overriding config path with .env
	if envPath := os.Getenv("MAZIKO_CONFIG"); envPath != "" {
		configPath = envPath
	}

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode JSON into global AppConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		return err
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

// Returns whether JSON logging is enabled
func UseJSONLogger() bool {
	return AppConfig.Logger.UseJSON
}
