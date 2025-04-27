package config

import (
	"os"
	"strings"
	"time"

	"github.com/pezanitech/maziko/backend/utils"

	inertia "github.com/romsar/gonertia"
)

func InitInertia() *inertia.Inertia {
	viteHotFile := ".tmp/hot"
	rootViewFile := "frontend/root.html"

	// check if laravel-vite-plugin is running in dev mode
	// it puts a "hot" file in the .tmp directory
	_, err := os.Stat(viteHotFile)

	if err != nil {
		// retry after 3 seconds, 3 attempts
		for i := range 3 {
			_, err = os.Stat(viteHotFile)
			if err == nil {
				break
			}

			utils.Logger.Info(
				"waiting for laravel-vite-plugin to start...",
				"attempt", i+1,
			)

			time.Sleep(3 * time.Second)
		}
	}

	if err == nil {
		i, err := inertia.NewFromFile(
			rootViewFile,
			inertia.WithSSR(),
		)
		if err != nil {
			utils.Logger.Error("failed to initialize inertia", "error", err)
		}

		i.ShareTemplateFunc("vite", func(entry string) (string, error) {
			content, err := os.ReadFile(viteHotFile)
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
		})

		i.ShareTemplateData("hmr", true)
		return i
	}

	manifestPath := "./build/.vite/manifest.json"

	i, err := inertia.NewFromFile(
		rootViewFile,
		inertia.WithVersionFromFile(manifestPath),
		inertia.WithSSR(),
	)

	if err != nil {
		utils.Logger.Error("failed to initialize inertia", "error", err)
		os.Exit(1)
	}

	i.ShareTemplateFunc("vite", vite(manifestPath, "/build/"))

	return i
}
