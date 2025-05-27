package router

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pezanitech/maziko/libs/core/config"
	"github.com/pezanitech/maziko/libs/core/logger"
)

// FindRouteImports finds routes to import
func FindRouteImports() ([]string, error) {
	var imports []string

	// walk through routes directory
	err := filepath.WalkDir(config.GetRoutesDir(),
		func(path string, dirEntry os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			importStmt, err := buildImportFromPath(path, dirEntry)
			if err != nil {
				return err
			}

			if importStmt != "" {
				imports = append(imports, importStmt)

				logger.Log.Info(
					"Adding route import",
					"path", path,
				)
			}
			return nil
		},
	)
	return imports, err
}

// buildImportFromPath returns import if directory is a route
func buildImportFromPath(path string, dirEntry os.DirEntry) (string, error) {
	if !dirEntry.IsDir() {
		return "", nil
	}

	// don't include routes directory
	if path == config.GetRoutesDir() {
		return "", nil
	}

	importStatement := fmt.Sprintf(
		"\"%s%s\"", config.GetPackagePrefix(), path,
	)

	return importStatement, nil
}

// FindRouteHandlers finds route handlers route directories
func FindRouteHandlers() ([]RouteHandler, error) {
	var handlers []RouteHandler

	// Walk through the routes directory recursively
	err := filepath.WalkDir(config.GetRoutesDir(),
		func(path string, dirEntry os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// skip non-directories and the root routes directory
			if !dirEntry.IsDir() || path == config.GetRoutesDir() {
				return nil
			}

			// create a route handler for this directory
			handler, err := newHandlerFromPath(path)
			if err != nil {
				return err
			}

			handlers = append(handlers, handler)

			logger.Log.Info(
				"Added route handler",
				"path", handler.Path,
				"package", handler.Package,
			)

			return nil
		},
	)

	return handlers, err
}

// newHandlerFromPath creates RouteHandler object from directory path
func newHandlerFromPath(path string) (RouteHandler, error) {
	routesDir := config.GetRoutesDir()
	routePath := strings.TrimPrefix(
		path, strings.TrimPrefix(routesDir, "./"),
	)

	// ensure forward slashes for path
	routePath = strings.ReplaceAll(routePath, "\\", "/")

	if routePath == "/index" {
		routePath = "/"
	}

	// get name end of the path
	packageName := filepath.Base(path)

	// For dynamic routes (directories starting with underscore),
	// create a unique package identifier based on the full path
	if strings.HasPrefix(packageName, "_") {
		// Get the parent directory name to create unique package identifier
		parentDir := filepath.Base(filepath.Dir(path))
		if parentDir != "routes" {
			// Add the parent directory name to make the package name unique
			packageName = parentDir + "." + packageName
		}
	}

	return RouteHandler{
		Path:     routePath,
		Package:  packageName,
		Function: "Route", // main Route() function
	}, nil
}
