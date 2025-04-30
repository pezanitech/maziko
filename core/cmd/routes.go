package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/router"
	"github.com/pezanitech/maziko/core/utils"
)

// collectRouteImports gathers all route imports while walking the routes directory
func collectRouteImports(path string, dirEntry os.DirEntry, err error) (string, error) {
	if err != nil {
		return "", err
	}

	if dirEntry.IsDir() {
		// don't register routes directory
		if path == config.RoutesDir {
			return "", nil
		}

		// Build import statement for this directory
		return fmt.Sprintf("import \"%s%s\"\n", config.PackagePrefix, path), nil
	}

	return "", nil
}

// collectAllRouteImports walks through the routes directory and collects all route imports
func collectAllRouteImports() ([]string, error) {
	var imports []string
	err := filepath.WalkDir(config.RoutesDir, func(path string, dirEntry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		importStatement, err := collectRouteImports(path, dirEntry, err)
		if err != nil {
			return err
		}

		if importStatement != "" {
			imports = append(imports, fmt.Sprintf("\"%s%s\"", config.PackagePrefix, path))
			utils.Logger.Info("Adding route", "path", path)
		}

		return nil
	})

	return imports, err
}

// collectRouteHandlers generates route handlers for each discovered route directory
func collectRouteHandlers() ([]router.RouteHandler, error) {
	var handlers []router.RouteHandler
	httpMethods := []string{"http.MethodGet", "http.MethodPost", "http.MethodPut", "http.MethodDelete"}

	err := filepath.WalkDir(config.RoutesDir, func(path string, dirEntry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root routes directory and non-directories
		if path == config.RoutesDir || !dirEntry.IsDir() {
			return nil
		}

		// Get the route path by removing the routes directory prefix
		routePath := strings.TrimPrefix(path, strings.TrimPrefix(config.RoutesDir, "./"))
		// Convert to URL path format
		routePath = strings.ReplaceAll(routePath, "\\", "/")
		// If it's an index route (app/routes/index), make it the root path
		if routePath == "/index" {
			routePath = "/"
		}

		// Get the package name from the last part of the path
		packageName := filepath.Base(path)

		// Add handlers for all HTTP methods
		for _, method := range httpMethods {
			// Extract method name and convert to uppercase
			methodName := strings.ToUpper(strings.TrimPrefix(method, "http.Method"))
			handlers = append(handlers, router.RouteHandler{
				Path:     routePath,
				Method:   method,
				Package:  packageName,
				Function: methodName,
			})
		}

		utils.Logger.Info(
			"Added route handlers",
			"path", routePath,
			"package", packageName,
		)
		return nil
	})

	return handlers, err
}
