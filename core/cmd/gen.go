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

// GenerateRoutes is the main function that orchestrates route generation
func GenerateRoutes() {
	// Initialize logger before use
	utils.InitLogger()

	utils.Logger.Info("Generating routes definitions...")

	// Create directory if it doesn't exist
	genDir := config.GetGenDir()
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
	data := router.RouteTemplateData{
		BuildPrefix:   config.GetBuildPrefix(),
		Imports:       imports,
		RouteHandlers: routeHandlers,
	}

	// Parse and execute the template
	tmpl, err := template.New("routes").Parse(router.RoutesTemplate)
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
