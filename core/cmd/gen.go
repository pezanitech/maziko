package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pezanitech/maziko/core/utils"
)

var (
	genDir        = "./core/gen"
	buildDir      = "./build"
	routesDir     = "./app/routes"
	publicDir     = "./app/public"
	buildPrefix   = "/build/"
	packagePrefix = "github.com/pezanitech/maziko/"
)

// Package declaration and common imports
var packageDeclaration = `
package gen

import "net/http"
import "os"
import "path"
import "strings"
import inertia "github.com/romsar/gonertia"
`

// Route handler definitions
var routeHandler = fmt.Sprintf(`
func DefineRoutes(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch true {
		case r.URL.Path == "/" && r.Method == http.MethodGet:
			index.GET(i, w, r)
		case r.URL.Path == "/" && r.Method == http.MethodPost:
			index.POST(i, w, r)
		case r.URL.Path == "/news" && r.Method == http.MethodGet:
			news.GET(i, w, r)
		case strings.HasPrefix(r.URL.Path, "%s"):
			handleRequest(w, r, buildDirHandler)
		default:
			handleRequest(w, r, staticFileHandler)
		}
	})
}`, buildPrefix)

var requestHandler = `
func handleRequest(w http.ResponseWriter, r *http.Request, f func() http.Handler) {
	f().ServeHTTP(w, r)
}`

var staticFileHandler = fmt.Sprintf(`
func staticFileHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourcePath := path.Join("%s", r.URL.Path)

		if _, err := os.Stat(resourcePath); err == nil {
			// file exists, serve it
			http.ServeFile(w, r, resourcePath)
			return
		}

		// no file exists, respond with a 404
		http.NotFound(w, r)
	})
}`, publicDir)

var buildDirHandler = fmt.Sprintf(`
func buildDirHandler() http.Handler {
	return http.StripPrefix(
		"%s",
		http.FileServer(http.Dir("%s")),
	)
}`, buildPrefix, buildDir)

// collectRouteImports gathers all route imports while walking the routes directory
func collectRouteImports(path string, dirEntry os.DirEntry, err error) (string, error) {
	if err != nil {
		return "", err
	}

	if dirEntry.IsDir() {
		// don't register routes directory
		if path == routesDir {
			return "", nil
		}

		// Build import statement for this directory
		return fmt.Sprintf("import \"%s%s\"\n", packagePrefix, path), nil
	}

	return "", nil
}

// collectAllRouteImports walks through the routes directory and collects all route imports
func collectAllRouteImports() (string, error) {
	var imports strings.Builder
	err := filepath.WalkDir(routesDir, func(path string, dirEntry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		importStatement, err := collectRouteImports(path, dirEntry, err)
		if err != nil {
			return err
		}

		if importStatement != "" {
			imports.WriteString(importStatement)
			utils.Logger.Info("Adding route", "path", path)
		}

		return nil
	})

	return imports.String(), err
}

func GenerateRoutes() {
	// Initialize logger before use
	utils.InitLogger()

	utils.Logger.Info("Generating routes definitions...")

	// Create directory if it doesn't exist
	if err := os.MkdirAll(genDir, 0755); err != nil {
		utils.Logger.Error(
			"Error creating gen directory",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to create gen directory: %v", err))
	}

	routesgenPath := filepath.Join(genDir, "routesgen.go")

	// Build the file content in memory
	var fileContent strings.Builder

	// Write package declaration
	fileContent.WriteString(packageDeclaration)

	// Collect all route imports
	imports, err := collectAllRouteImports()
	if err != nil {
		utils.Logger.Error(
			"Error collecting route imports",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to collect route imports: %v", err))
	}

	// Add imports to file content
	fileContent.WriteString(imports)

	// Write routing definitions
	fileContent.WriteString(routeHandler)
	fileContent.WriteString(requestHandler)
	fileContent.WriteString(staticFileHandler)
	fileContent.WriteString(buildDirHandler)

	// Log the generated content
	utils.Logger.Info("Generated routes file content", "path", routesgenPath)

	// Write the full content to the file at once
	if err := os.WriteFile(routesgenPath, []byte(fileContent.String()), 0644); err != nil {
		utils.Logger.Error(
			"Error writing to routesgen.go file",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to write to routesgen.go: %v", err))
	}

	utils.Logger.Info("Routes generation completed successfully")
}
