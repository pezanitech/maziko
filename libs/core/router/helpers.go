package router

import (
	"net/http"
	"strings"

	"github.com/pezanitech/maziko/libs/core/errors"
	"github.com/pezanitech/maziko/libs/core/logger"
)

// RenderPage renders a page with the provided props,
// can optionally be specified with component parameter
func RenderPage(i Inertia, w http.ResponseWriter, r *http.Request, props Props, component ...string) {
	componentName := ""

	if len(component) > 0 && component[0] != "" {
		componentName = component[0] // use provided component name
	} else {
		path := r.URL.Path
		if path == "" {
			path = "/"
		}

		// get name from routeComponents map using URL path
		if storedComponent, exists := routeComponents[path]; exists {
			componentName = storedComponent

			logger.Log.Debug(
				"Using component from route map",
				"path", path,
				"component", componentName,
			)
			// check if there is dynamic segment path matching the URL path
		} else if dynamicComponent, found := findDynamicRouteComponent(path); found {
			componentName = dynamicComponent

			logger.Log.Debug(
				"Using component from dynamic route match",
				"path", path,
				"component", componentName,
			)
		} else {
			msg := "Component not found in routeComponents map"

			logger.Log.Error(msg, "path", path)
			panic(msg)
		}
	}

	logger.Log.Info(
		"Rendering page",
		"component", componentName,
	)

	// render the page
	err := RenderInertiaPage(i, w, r, componentName, props)
	if err != nil {
		errors.HandleServerErr(w, err)
	}
}

// RenderPageWithMeta renders a page with the provided props and metadata,
// can optionally be specified with component parameter
func RenderPageWithMeta(i Inertia, w http.ResponseWriter, r *http.Request, props Props, meta MetaData, component ...string) {
	componentName := ""

	if len(component) > 0 && component[0] != "" {
		componentName = component[0] // use provided component name
	} else {
		path := r.URL.Path
		if path == "" {
			path = "/"
		}

		// get name from routeComponents map using URL path
		if storedComponent, exists := routeComponents[path]; exists {
			componentName = storedComponent

			logger.Log.Debug(
				"Using component from route map",
				"path", path,
				"component", componentName,
			)
			// check if there is dynamic segment path matching the URL path
		} else if dynamicComponent, found := findDynamicRouteComponent(path); found {
			componentName = dynamicComponent

			logger.Log.Debug(
				"Using component from dynamic route match",
				"path", path,
				"component", componentName,
			)
		} else {
			msg := "Component not found in routeComponents map"

			logger.Log.Error(msg, "path", path)
			panic(msg)
		}
	}

	logger.Log.Info(
		"Rendering page with metadata",
		"component", componentName,
		"title", meta.Title,
	)

	// render the page with metadata
	err := RenderInertiaPageWithMeta(i, w, r, componentName, props, meta)
	if err != nil {
		errors.HandleServerErr(w, err)
	}
}

// findDynamicRouteComponent looks for a matching dynamic route in routeComponents
// It checks if any route with URL parameters matches the given path
func findDynamicRouteComponent(path string) (string, bool) {
	// If path is empty, return not found
	if path == "" {
		return "", false
	}

	// Split the current path into segments
	pathSegments := strings.Split(strings.Trim(path, "/"), "/")

	// Iterate through all registered routes
	for routePath, component := range routeComponents {
		// Skip if this isn't a dynamic route (doesn't contain '{')
		if !strings.Contains(routePath, "{") {
			continue
		}

		// Split the route path into segments
		routeSegments := strings.Split(strings.Trim(routePath, "/"), "/")

		// Skip if segment count doesn't match
		if len(routeSegments) != len(pathSegments) {
			continue
		}

		// Check if this route matches the path
		matches := true
		for i, routeSegment := range routeSegments {
			// If this segment is a parameter (wrapped in {}), it matches any value
			if strings.HasPrefix(routeSegment, "{") && strings.HasSuffix(routeSegment, "}") {
				// Parameter segments always match
				continue
			}

			// For regular segments, they must match exactly
			if routeSegment != pathSegments[i] {
				matches = false
				break
			}
		}

		// If all segments matched, return the component
		if matches {
			logger.Log.Debug(
				"Found matching dynamic route",
				"path", path,
				"route", routePath,
				"component", component,
			)
			return component, true
		}
	}

	// No matching dynamic route found
	return "", false
}
