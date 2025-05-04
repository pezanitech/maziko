package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/router"
	"github.com/pezanitech/maziko/core/utils"
)

// GenerateRoutes creates route definitions for the application
// This is the core mechanism behind maziko's routing system
func GenerateRoutes() {
	utils.Logger.Info("Generating routes definitions...")

	// Ensure directories exist
	genDir := prepareGenerationDirectory()

	// Output file for generated routes
	routesgenFile := filepath.Join(genDir, "routesgen.go")

	// Collect required route information
	imports := collectRouteImports()
	routeHandlers := collectRouteHandlers()

	// Generate routes file from template
	generateRoutesFromTemplate(routesgenFile, imports, routeHandlers)

	utils.Logger.Info("Routes generation completed successfully")
}

// prepareGenerationDirectory ensures the generation directory exists
func prepareGenerationDirectory() string {
	genDir := config.GetGenDir()
	if err := os.MkdirAll(genDir, 0755); err != nil {
		utils.Logger.Error(
			"Error creating gen directory",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to create gen directory: %v", err))
	}
	return genDir
}

// collectRouteImports gathers all route packages that need to be imported
func collectRouteImports() []string {
	imports, err := router.FindRouteImports()
	if err != nil {
		utils.Logger.Error(
			"Error collecting route imports",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to collect route imports: %v", err))
	}
	return imports
}

// collectRouteHandlers gathers all route handlers defined in the application
func collectRouteHandlers() []router.RouteHandler {
	routeHandlers, err := router.FindRouteHandlers()
	if err != nil {
		utils.Logger.Error(
			"Error collecting route handlers",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to collect route handlers: %v", err))
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
		utils.Logger.Error(
			"Error parsing template",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to parse template: %v", err))
	}

	// Execute template with data
	var fileContent strings.Builder
	if err := tmpl.Execute(&fileContent, templateData); err != nil {
		utils.Logger.Error(
			"Error executing template",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to execute template: %v", err))
	}

	// Write generated content to file
	if err := os.WriteFile(outputFile, []byte(fileContent.String()), 0644); err != nil {
		utils.Logger.Error(
			"Error writing to routesgen file",
			"path", outputFile,
			"error", err,
		)
		panic(fmt.Sprintf("Failed to write to routesgen: %v", err))
	}

	utils.Logger.Info(
		"Generated routesgen file",
		"path", outputFile,
	)
}
