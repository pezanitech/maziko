package router

import (
	"os"
	"strings"
	"time"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/utils"

	inertia "github.com/romsar/gonertia"
)

// Initialize an Inertia instance based on environment
func InitInertia() *inertia.Inertia {
	if isDevMode() {
		return initInertiaDevMode()
	}
	return initInertiaProdMode()
}

// Check if running in dev mode
// laravel-vite-plugin creates a tmp/hot file
func isDevMode() bool {
	utils.Logger.Info("Checking for development environment...")

	// Check for hot file
	if _, err := os.Stat(config.GetHotFile()); err == nil {
		utils.Logger.Info("Development mode detected")
		return true
	}

	// Hot file not found, wait for Vite to start up
	utils.Logger.Info("Waiting for Vite development server to start...")

	for attempt := 1; attempt <= config.MaxViteDetectionAttempts(); attempt++ {
		utils.Logger.Info(
			"Looking for Vite development server",
			"attempt", attempt,
			"of", config.MaxViteDetectionAttempts(),
		)

		time.Sleep(config.ViteDetectionInterval())

		if _, err := os.Stat(config.GetHotFile()); err == nil {
			utils.Logger.Info("Development mode detected")
			return true
		}
	}

	utils.Logger.Info("Vite development server not detected, using production mode")
	return false
}

// Initialize Inertia in development mode with hot reloading
func initInertiaDevMode() *inertia.Inertia {
	i, err := inertia.New(
		RootHTMLTemplate,
		inertia.WithSSR(),
	)
	if err != nil {
		utils.Logger.Error(
			"failed to initialize inertia in dev mode",
			"error", err,
		)
		os.Exit(1)
	}

	i.ShareTemplateFunc(
		"vite",
		func(entry string) (string, error) {
			content, err := os.ReadFile(config.GetHotFile())
			if err != nil {
				return "", err
			}

			url := strings.TrimSpace(string(content))

			if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
				url = url[strings.Index(url, ":")+1:]
			} else {
				url = "//localhost:8080"
			}

			if entry != "" && !strings.HasPrefix(entry, "/") {
				entry = "/" + entry
			}

			return url + entry, nil
		},
	)

	i.ShareTemplateData("hmr", true)
	return i
}

// Initialize Inertia in production mode
func initInertiaProdMode() *inertia.Inertia {
	i, err := inertia.New(
		RootHTMLTemplate,
		inertia.WithVersionFromFile(config.GetViteManifestFile()),
		inertia.WithSSR(),
	)

	if err != nil {
		utils.Logger.Error(
			"failed to initialize inertia in production mode",
			"error", err,
		)
		os.Exit(1)
	}

	i.ShareTemplateFunc(
		"vite",
		vite(config.GetViteManifestFile(), config.GetBuildPrefix()),
	)

	return i
}
