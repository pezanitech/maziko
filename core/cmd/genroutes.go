package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/errors"
	"github.com/pezanitech/maziko/core/logger"
	"github.com/pezanitech/maziko/core/router"
)

// GenerateRoutes creates route definitions for the application
// This is the core mechanism behind maziko's routing system
func GenerateRoutes() {
	logger.Log.Info("Generating routes definitions...")

	// Ensure directories exist
	genDir := prepareGenerationDirectory()

	// Output file for generated routes
	routesgenFile := filepath.Join(genDir, "routesgen.go")

	// Collect required route information
	imports := collectRouteImports()
	routeHandlers := collectRouteHandlers()

	// Generate routes file from template
	generateRoutesFromTemplate(routesgenFile, imports, routeHandlers)

	logger.Log.Info("Routes generation completed successfully")
}

// prepareGenerationDirectory ensures the generation directory exists
func prepareGenerationDirectory() string {
	genDir := config.GetGenDir()
	if err := os.MkdirAll(genDir, 0755); err != nil {
		errors.HandleFatalError("Error creating gen directory", err)
	}
	return genDir
}

// collectRouteImports gathers all route packages that need to be imported
func collectRouteImports() []string {
	imports, err := router.FindRouteImports()
	if err != nil {
		errors.HandleFatalError("Error collecting route imports", err)
	}
	return imports
}

// collectRouteHandlers gathers all route handlers defined in the application
func collectRouteHandlers() []router.RouteHandler {
	routeHandlers, err := router.FindRouteHandlers()
	if err != nil {
		errors.HandleFatalError("Error collecting route handlers", err)
	}
	return routeHandlers
}

// generateRoutesFromTemplate creates the routes file using the template
func generateRoutesFromTemplate(outputFile string, imports []string, routeHandlers []router.RouteHandler) {
	// Prepare template data
	templateData := router.RouteTemplateData{
		BuildPrefix:   config.GetBuildPrefix(),
		Imports:       imports,
		RouteHandlers: routeHandlers,
	}

	// Parse routes template
	tmpl, err := template.New("routes").Parse(router.RoutesTemplate)
	if err != nil {
		errors.HandleFatalError("Error parsing template", err)
	}

	// Execute template with data
	var fileContent strings.Builder
	if err := tmpl.Execute(&fileContent, templateData); err != nil {
		errors.HandleFatalError("Error executing template", err)
	}

	// Write generated content to file
	if err := os.WriteFile(outputFile, []byte(fileContent.String()), 0644); err != nil {
		errors.HandleFatalError(
			"Error writing to routesgen file", err,
			"path", outputFile,
		)
	}

	logger.Log.Info(
		"Generated routesgen file",
		"path", outputFile,
	)
}
