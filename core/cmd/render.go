package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/logger"
	"github.com/pezanitech/maziko/core/router"
	"github.com/pezanitech/maziko/core/utils"

	inertia "github.com/romsar/gonertia"
)

// InitRenderer creates and configures an Inertia renderer instance based on environment
func InitRenderer() *inertia.Inertia {
	if isDevMode() {
		return initDevRenderer()
	}
	return initProdRenderer()
}

// isDevMode determines if the application is running in development mode
// In development, Vite creates a tmp/hot file that we can detect
func isDevMode() bool {
	logger.Log.Info("Checking for development environment...")

	// Check for hot file - immediate indicator of dev mode
	if _, err := os.Stat(config.GetHotFile()); err == nil {
		logger.Log.Info("Development mode detected")
		return true
	}

	// Hot file not found - wait for Vite server to start
	return waitForViteServer()
}

// waitForViteServer attempts to detect a starting Vite development server
func waitForViteServer() bool {
	logger.Log.Info("Waiting for Vite development server to start...")

	// Retry multiple times with a delay
	for attempt := 1; attempt <= config.MaxViteDetectionAttempts(); attempt++ {
		logger.Log.Info(
			"Looking for Vite development server",
			"attempt", attempt,
			"of", config.MaxViteDetectionAttempts(),
		)

		time.Sleep(config.ViteDetectionInterval())

		// Check again for hot file
		if _, err := os.Stat(config.GetHotFile()); err == nil {
			logger.Log.Info("Development mode detected")
			return true
		}
	}

	logger.Log.Info("Vite development server not detected, using production mode")
	return false
}

// initDevRenderer creates an Inertia renderer configured for development
func initDevRenderer() *inertia.Inertia {
	i, err := inertia.New(
		router.RootHTMLTemplate,
		inertia.WithSSR(),
	)

	if err != nil {
		utils.HandleFatalError("Failed to initialize renderer in dev mode", err)
	}

	// Add the vite template function for dev mode
	i.ShareTemplateFunc(
		"vite",
		createDevViteFunction(),
	)

	// Enable hot module reloading
	i.ShareTemplateData("hmr", true)

	return i
}

// createDevViteFunction creates a template function that resolves asset paths in development
func createDevViteFunction() func(string) (string, error) {
	return func(entry string) (string, error) {
		// Read the Vite hot file to get the development server URL
		content, err := os.ReadFile(config.GetHotFile())
		if err != nil {
			return "", err
		}

		// Parse the URL from hot file content
		url := strings.TrimSpace(string(content))

		// Format URL consistently
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			url = url[strings.Index(url, ":")+1:]
		} else {
			url = "//localhost:8080"
		}

		// Ensure path format is consistent
		if entry != "" && !strings.HasPrefix(entry, "/") {
			entry = "/" + entry
		}

		return url + entry, nil
	}
}

// initProdRenderer creates an Inertia renderer configured for production
func initProdRenderer() *inertia.Inertia {
	i, err := inertia.New(
		router.RootHTMLTemplate,
		inertia.WithVersionFromFile(config.GetViteManifestFile()),
		inertia.WithSSR(),
	)

	if err != nil {
		utils.HandleFatalError(
			"Failed to initialize renderer in production mode", err,
		)
	}

	// Add the vite template function for production mode
	i.ShareTemplateFunc(
		"vite",
		createProdViteFunction(config.GetViteManifestFile(), config.GetBuildPrefix()),
	)

	return i
}

// createProdViteFunction creates a template function that resolves
// vite asset paths from manifest for production
func createProdViteFunction(manifestPath, buildDir string) func(string) (string, error) {
	// Load and parse the Vite manifest file
	viteAssets := loadViteManifest(manifestPath)

	// Return template function that uses the manifest to resolve asset paths
	return func(p string) (string, error) {
		if val, ok := viteAssets[p]; ok {
			return path.Join("/", buildDir, val.File), nil
		}
		return "", fmt.Errorf("asset %q not found", p)
	}
}

// loadViteManifest loads the Vite manifest file and parses it
func loadViteManifest(manifestPath string) map[string]*struct {
	File   string `json:"file"`
	Source string `json:"src"`
} {
	// Open manifest file
	manifest, err := os.Open(manifestPath)
	if err != nil {
		utils.HandleFatalError("Cannot open provided vite manifest file", err)
	}
	defer manifest.Close()

	// Parse manifest JSON
	viteAssets := make(map[string]*struct {
		File   string `json:"file"`
		Source string `json:"src"`
	})

	if err = json.NewDecoder(manifest).Decode(&viteAssets); err != nil {
		utils.HandleFatalError("Cannot unmarshal vite manifest file to JSON", err)
	}

	// Log available assets for debugging
	for k, v := range viteAssets {
		logger.Log.Info(
			"Vite asset",
			"path", k,
			"file", v.File,
		)
	}

	return viteAssets
}
