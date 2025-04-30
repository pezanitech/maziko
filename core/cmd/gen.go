package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

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

// Define template structure
type RouteTemplateData struct {
	BuildPrefix   string
	PublicDir     string
	BuildDir      string
	Imports       []string
	RouteHandlers []RouteHandler
}

type RouteHandler struct {
	Path     string
	Method   string
	Package  string
	Function string
}

// Routes template
var routesTemplate = `
package gen

import "net/http"
import "os"
import "path"
import "strings"
import inertia "github.com/romsar/gonertia"
{{range .Imports}}
import {{.}}
{{end}}

func DefineRoutes(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch true {
		{{range .RouteHandlers}}
		case r.URL.Path == "{{.Path}}" && r.Method == {{.Method}}:
			{{.Package}}.{{.Function}}(i, w, r)
		{{end}}
		case strings.HasPrefix(r.URL.Path, "{{.BuildPrefix}}"):
			handleRequest(w, r, buildDirHandler)
		default:
			handleRequest(w, r, staticFileHandler)
		}
	})
}

func handleRequest(w http.ResponseWriter, r *http.Request, f func() http.Handler) {
	f().ServeHTTP(w, r)
}

func staticFileHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourcePath := path.Join("{{.PublicDir}}", r.URL.Path)

		if _, err := os.Stat(resourcePath); err == nil {
			// file exists, serve it
			http.ServeFile(w, r, resourcePath)
			return
		}

		// no file exists, respond with a 404
		http.NotFound(w, r)
	})
}

func buildDirHandler() http.Handler {
	return http.StripPrefix(
		"{{.BuildPrefix}}",
		http.FileServer(http.Dir("{{.BuildDir}}")),
	)
}
`

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
func collectAllRouteImports() ([]string, error) {
	var imports []string
	err := filepath.WalkDir(routesDir, func(path string, dirEntry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		importStatement, err := collectRouteImports(path, dirEntry, err)
		if err != nil {
			return err
		}

		if importStatement != "" {
			imports = append(imports, fmt.Sprintf("\"%s%s\"", packagePrefix, path))
			utils.Logger.Info("Adding route", "path", path)
		}

		return nil
	})

	return imports, err
}

// collectRouteHandlers generates route handlers for each discovered route directory
func collectRouteHandlers() ([]RouteHandler, error) {
	var handlers []RouteHandler
	httpMethods := []string{"http.MethodGet", "http.MethodPost", "http.MethodPut", "http.MethodDelete"}

	err := filepath.WalkDir(routesDir, func(path string, dirEntry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root routes directory and non-directories
		if path == routesDir || !dirEntry.IsDir() {
			return nil
		}

		// Get the route path by removing the routes directory prefix
		routePath := strings.TrimPrefix(path, strings.TrimPrefix(routesDir, "./"))
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
			handlers = append(handlers, RouteHandler{
				Path:     routePath,
				Method:   method,
				Package:  packageName,
				Function: methodName,
			})
		}

		utils.Logger.Info("Added route handlers", "path", routePath, "package", packageName)
		return nil
	})

	return handlers, err
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

	// Collect all route imports
	imports, err := collectAllRouteImports()
	if err != nil {
		utils.Logger.Error(
			"Error collecting route imports",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to collect route imports: %v", err))
	}

	// Collect route handlers dynamically instead of hardcoding them
	routeHandlers, err := collectRouteHandlers()
	if err != nil {
		utils.Logger.Error(
			"Error collecting route handlers",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to collect route handlers: %v", err))
	}

	// Create template data
	data := RouteTemplateData{
		BuildPrefix:   buildPrefix,
		PublicDir:     publicDir,
		BuildDir:      buildDir,
		Imports:       imports,
		RouteHandlers: routeHandlers,
	}

	// Parse and execute the template
	tmpl, err := template.New("routes").Parse(routesTemplate)
	if err != nil {
		utils.Logger.Error(
			"Error parsing template",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to parse template: %v", err))
	}

	var fileContent strings.Builder
	if err := tmpl.Execute(&fileContent, data); err != nil {
		utils.Logger.Error(
			"Error executing template",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to execute template: %v", err))
	}

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
