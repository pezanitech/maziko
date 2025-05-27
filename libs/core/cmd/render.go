package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pezanitech/maziko/libs/core/config"
	"github.com/pezanitech/maziko/libs/core/errors"
	"github.com/pezanitech/maziko/libs/core/logger"
	"github.com/pezanitech/maziko/libs/core/router"
)

// Initializes an Inertia instance based on environment
func InitRenderer() router.Inertia {
	if isDevMode() {
		return initDevRenderer()
	}
	return initProdRenderer()
}

// Checks whether application is running in dev mode
// In dev mode, Vite creates a tmp/hot file
func isDevMode() bool {
	logger.Log.Info(
		"Checking for development environment...",
	)

	// check for hot file - immediate indicator of dev mode
	if _, err := os.Stat(config.GetHotFile()); err == nil {
		logger.Log.Info("Development mode detected")
		return true
	}

	// hot file not found - wait for Vite server to start
	return waitForViteServer()
}

// Attempts to detect Vite development server
func waitForViteServer() bool {
	logger.Log.Info(
		"Waiting for Vite development server to start...",
	)

	// retry multiple times with a delay
	for attempt := 1; attempt <= config.MaxViteDetectionAttempts(); attempt++ {
		logger.Log.Info(
			"Looking for Vite development server",
			"attempt", attempt,
			"of", config.MaxViteDetectionAttempts(),
		)

		time.Sleep(config.ViteDetectionInterval())

		// check again for hot file
		if _, err := os.Stat(config.GetHotFile()); err == nil {
			logger.Log.Info(
				"Development mode detected",
			)
			return true
		}
	}

	logger.Log.Info(
		"Vite development server not detected, using production mode",
	)
	return false
}

// Creates an Inertia instance configured for development
func initDevRenderer() router.Inertia {
	i, err := router.NewInertia(
		router.RootHTMLTemplate,
		router.InertiaOptions.WithSSR(),
	)
	if err != nil {
		errors.HandleFatalError(
			"Failed to initialize renderer in dev mode", err,
		)
	}

	// add vite template function for dev mode
	i.ShareTemplateFunc(
		"vite",
		createDevViteFunction(),
	)

	// enable hot module reloading
	i.ShareTemplateData("hmr", true)

	// share default metadata with the template
	i.ShareTemplateData("meta", router.DefaultMetaData)

	return i
}

// Creates a template function that resolves asset paths in development
func createDevViteFunction() func(string) (string, error) {
	return func(entry string) (string, error) {
		// read Vite hot file to get dev server URL
		content, err := os.ReadFile(config.GetHotFile())
		if err != nil {
			return "", err
		}

		// parse the URL from hot file content
		url := strings.TrimSpace(string(content))

		// format URL consistently
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			url = url[strings.Index(url, ":")+1:]
		} else {
			url = "//localhost:8080"
		}

		// ensure path format is consistent
		if entry != "" && !strings.HasPrefix(entry, "/") {
			entry = "/" + entry
		}

		return url + entry, nil
	}
}

// Creates an Inertia instance configured for production
func initProdRenderer() router.Inertia {
	manifestPath := config.GetViteManifestFile()
	var serverVersion string
	versionBytes, err := os.ReadFile(manifestPath)

	if err != nil {
		logger.Log.Error(
			"Failed to read Vite manifest file for versioning. Inertia asset versioning will be disabled. This can lead to issues with stale browser tabs after deployments.",
			"path", manifestPath,
			"error", err,
		)
		// serverVersion will remain "", which effectively disables version checks in gonertia if i.version == ""
	} else {
		serverVersion = strings.TrimSpace(string(versionBytes))
		if serverVersion == "" {
			logger.Log.Warn(
				"Vite manifest file is empty or only whitespace. Inertia asset versioning may be ineffective.",
				"path", manifestPath,
			)
		} else {
			logger.Log.Info("Successfully loaded version from manifest file.", "path", manifestPath)
		}
	}

	i, err := router.NewInertia(
		router.RootHTMLTemplate,
		router.InertiaOptions.WithVersion(serverVersion), // Use the explicitly loaded version
		router.InertiaOptions.WithSSR(),
	)

	if err != nil {
		errors.HandleFatalError(
			"Failed to initialize renderer in production mode", err,
		)
	}

	// add the vite template function for production mode
	i.ShareTemplateFunc(
		"vite",
		createProdViteFunction(
			config.GetViteManifestFile(),
			config.GetBuildPrefix(),
		),
	)

	// share default metadata with the template
	i.ShareTemplateData("meta", router.DefaultMetaData)

	return i
}

// Create a template function that resolves
// vite asset paths from manifest for production
func createProdViteFunction(manifestPath, buildDir string) func(string) (string, error) {
	// load and parse the Vite manifest file
	viteAssets := loadViteManifest(manifestPath)

	// template function that uses manifest to resolve asset paths
	return func(p string) (string, error) {
		if val, ok := viteAssets[p]; ok {
			return path.Join("/", buildDir, val.File), nil
		}
		return "", fmt.Errorf("asset %q not found", p)
	}
}

// Loads the Vite manifest file and parses it
func loadViteManifest(manifestPath string) map[string]*struct {
	File   string `json:"file"`
	Source string `json:"src"`
} {
	// open manifest file
	manifest, err := os.Open(manifestPath)
	if err != nil {
		errors.HandleFatalError(
			"Cannot open provided vite manifest file", err,
		)
	}
	defer manifest.Close()

	// parse manifest JSON
	viteAssets := make(map[string]*struct {
		File   string `json:"file"`
		Source string `json:"src"`
	})

	if err = json.NewDecoder(manifest).Decode(&viteAssets); err != nil {
		errors.HandleFatalError(
			"Cannot unmarshal vite manifest file to JSON", err,
		)
	}

	// log available assets for debugging
	for k, v := range viteAssets {
		logger.Log.Debug(
			"Vite asset",
			"path", k,
			"file", v.File,
		)
	}

	return viteAssets
}
