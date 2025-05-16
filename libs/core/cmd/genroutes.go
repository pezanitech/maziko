package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pezanitech/maziko/libs/core/config"
	"github.com/pezanitech/maziko/libs/core/errors"
	"github.com/pezanitech/maziko/libs/core/logger"
	"github.com/pezanitech/maziko/libs/core/router"
)

// GenerateRoutes generates route definitions for the application.
// The core mechanism behind Maziko's routing system
func GenerateRoutes() {
	logger.Log.Info("Generating routes definitions...")

	// ensure directories exist
	genDir := prepareGenerationDirectory()

	// output file for generated routes
	outputFile := filepath.Join(genDir, "routesgen.go")

	// collect required route information
	imports := collectRouteImports()
	routeHandlers := collectRouteHandlers()

	// generate routes file from template
	generateRoutesFromTemplate(outputFile, imports, routeHandlers)

	logger.Log.Info(
		"Routes generation completed successfully",
	)
}

// prepareGenerationDirectory ensures the generation directory exists
func prepareGenerationDirectory() string {
	genDir := config.GetGenDir()
	if err := os.MkdirAll(genDir, 0755); err != nil {
		errors.HandleFatalError(
			"Error creating gen directory", err,
		)
	}
	return genDir
}

// collectRouteImports collects all route packages that need to be imported
func collectRouteImports() []string {
	imports, err := router.FindRouteImports()
	if err != nil {
		errors.HandleFatalError(
			"Error collecting route imports", err,
		)
	}
	return imports
}

// collectRouteHandlers collects all route handlers defined in the application
func collectRouteHandlers() []router.RouteHandler {
	routeHandlers, err := router.FindRouteHandlers()
	if err != nil {
		errors.HandleFatalError(
			"Error collecting route handlers", err,
		)
	}
	return routeHandlers
}

// generateRoutesFromTemplate generates a routesgen file using a template
func generateRoutesFromTemplate(outputFile string, imports []string, routeHandlers []router.RouteHandler) {
	// prepare template data
	templateData := router.RouteTemplateData{
		BuildPrefix:   config.GetBuildPrefix(),
		Imports:       imports,
		RouteHandlers: routeHandlers,
	}

	// parse routes template
	tmpl, err := template.New("routes").Parse(router.RoutesTemplate)
	if err != nil {
		errors.HandleFatalError(
			"Error parsing template", err,
		)
	}

	// execute template with data
	var fileContent strings.Builder
	if err := tmpl.Execute(&fileContent, templateData); err != nil {
		errors.HandleFatalError(
			"Error executing template", err,
		)
	}

	// write generated content to file
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
