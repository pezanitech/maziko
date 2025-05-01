package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/router"
	"github.com/pezanitech/maziko/core/utils"

	inertia "github.com/romsar/gonertia"
)

// Initialize a renderer instance based on environment
func InitRenderer() *inertia.Inertia {
	if isDevMode() {
		return initDevRenderer()
	}
	return initProdRenderer()
}

// Check if running in dev mode
// laravel-vite-plugin creates a tmp/hot file
func isDevMode() bool {
	utils.Logger.Info("Checking for development environment...")

	// check for hot file
	if _, err := os.Stat(config.GetHotFile()); err == nil {
		utils.Logger.Info("Development mode detected")
		return true
	}

	// hot file not found, wait for vite to start up
	utils.Logger.Info(
		"Waiting for Vite development server to start...",
	)

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

	utils.Logger.Info(
		"Vite development server not detected, using production mode",
	)
	return false
}

// Initialize renderer in development mode with hot reloading
func initDevRenderer() *inertia.Inertia {
	i, err := inertia.New(
		router.RootHTMLTemplate,
		inertia.WithSSR(),
	)
	if err != nil {
		utils.Logger.Error(
			"failed to initialize renderer in dev mode",
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

// Initialize renderer in production mode
func initProdRenderer() *inertia.Inertia {
	i, err := inertia.New(
		router.RootHTMLTemplate,
		inertia.WithVersionFromFile(config.GetViteManifestFile()),
		inertia.WithSSR(),
	)

	if err != nil {
		utils.Logger.Error(
			"failed to initialize renderer in production mode",
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

// Creates a function that resolves vite asset paths from manifest
func vite(manifestPath, buildDir string) func(path string) (string, error) {
	manifest, err := os.Open(manifestPath)
	if err != nil {
		utils.Logger.Error(
			"cannot open provided vite manifest file",
			"error", err,
		)
		os.Exit(1)
	}
	defer manifest.Close()

	viteAssets := make(map[string]*struct {
		File   string `json:"file"`
		Source string `json:"src"`
	})

	if err = json.NewDecoder(manifest).Decode(&viteAssets); err != nil {
		utils.Logger.Error(
			"cannot unmarshal vite manifest file to json",
			"error", err,
		)
		os.Exit(1)
	}

	// print content of viteAssets
	for k, v := range viteAssets {
		utils.Logger.Info(
			"vite asset",
			"path", k,
			"file", v.File,
		)
	}

	return func(p string) (string, error) {
		if val, ok := viteAssets[p]; ok {
			return path.Join("/", buildDir, val.File), nil
		}
		return "", fmt.Errorf("asset %q not found", p)
	}
}
