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

// Generates route definitions for the application
// The core mechanism behind maziko's routing system
func GenerateRoutes() {
	utils.Logger.Info(
		"Generating routes definitions...",
	)

	// create gen directory if it doesn't exist
	genDir := config.GetGenDir()
	if err := os.MkdirAll(genDir, 0755); err != nil {
		utils.Logger.Error(
			"Error creating gen directory",
			"error", err,
		)

		panic(fmt.Sprintf(
			"Failed to create gen directory: %v", err,
		))
	}

	routesgenFile := filepath.Join(genDir, "routesgen.go")

	// collect all routes to import
	imports, err := router.CollectAllRouteImports()
	if err != nil {
		utils.Logger.Error(
			"Error collecting route imports",
			"error", err,
		)

		panic(fmt.Sprintf(
			"Failed to collect route imports: %v", err,
		))
	}

	// collect route handlers for each route
	routeHandlers, err := router.CollectRouteHandlers()
	if err != nil {
		utils.Logger.Error(
			"Error collecting route handlers",
			"error", err,
		)

		panic(fmt.Sprintf(
			"Failed to collect route handlers: %v", err,
		))
	}

	// create template data struct
	templateData := router.RouteTemplateData{
		BuildPrefix:   config.GetBuildPrefix(),
		Imports:       imports,
		RouteHandlers: routeHandlers,
	}

	// parse and execute routes template
	tmpl, err := template.New("routes").Parse(
		router.RoutesTemplate,
	)
	if err != nil {
		utils.Logger.Error(
			"Error parsing template",
			"error", err,
		)

		panic(fmt.Sprintf(
			"Failed to parse template: %v", err,
		))
	}

	var fileContent strings.Builder
	if err := tmpl.Execute(&fileContent, templateData); err != nil {
		utils.Logger.Error(
			"Error executing template",
			"error", err,
		)

		panic(fmt.Sprintf(
			"Failed to execute template: %v", err,
		))
	}

	// write routesgen file
	if err := os.WriteFile(routesgenFile, []byte(fileContent.String()), 0644); err != nil {
		utils.Logger.Error(
			"Error writing to routesgen file",
			"path", routesgenFile,
			"error", err,
		)

		panic(fmt.Sprintf(
			"Failed to write to routesgen: %v", err,
		))
	}

	utils.Logger.Info(
		"Generated routesgen file",
		"path", routesgenFile,
	)

	utils.Logger.Info(
		"Routes generation completed successfully",
	)
}
