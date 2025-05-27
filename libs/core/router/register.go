package router

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/pezanitech/maziko/libs/core/logger"
)

// registerRoute registers a route with an HTTP method
func registerRoute(method string, routePath string, handler InertiaHandlerFunc) {
	// componentName = routePath without leading and trailing slashes
	componentName := strings.Trim(routePath, "/")

	// set "/index" routePath to "/"
	if routePath == "/index" {
		routePath = "/"
	}

	// convert paths with underscores (like articles/_slug) to chi router pattern
	if strings.Contains(routePath, "_") {
		segments := strings.Split(routePath, "/")
		for i, segment := range segments {
			if strings.HasPrefix(segment, "_") {
				// Convert _paramName to {paramName}
				paramName := segment[1:] // Remove the underscore
				segments[i] = "{" + paramName + "}"
			}
		}
		routePath = strings.Join(segments, "/")
	}

	logger.Log.Info(
		"Registering route",
		"path", routePath,
		"component", componentName,
		"method", method,
	)

	// store component name for this route
	routeComponents[routePath] = componentName

	// register the route on the router
	switch method {
	case "GET":
		appRouter.Router.Get(routePath, newRouteHandler(handler))
	case "POST":
		appRouter.Router.Post(routePath, newRouteHandler(handler))
	case "PUT":
		appRouter.Router.Put(routePath, newRouteHandler(handler))
	case "DELETE":
		appRouter.Router.Delete(routePath, newRouteHandler(handler))
	case "PATCH":
		appRouter.Router.Patch(routePath, newRouteHandler(handler))
	}
}

// newRouteHandler creates an standard http.HandlerFunc from an InertiaHTTPHandler
func newRouteHandler(handler InertiaHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(appRouter.renderer, w, r)
	}
}

// extractPath extracts the route path from a file path or gets it from the caller's file path.
// When given a file path as input, it identifies paths by looking for "/routes/" and removes "/handler.go" if present.
// If no input is provided, it gets the caller's file path and applies the same logic.
func extractPath(input ...string) string {
	var filePath string

	if len(input) > 0 && input[0] != "" {
		filePath = input[0]
	} else {
		// Get caller's file name (skip 1 additional frame to get the actual caller)
		_, file, _, ok := runtime.Caller(2)
		if !ok {
			panic("Failed to get caller information")
		}
		filePath = file
	}

	const marker = "/routes/"
	const suffix = "/handler.go"

	// find position of "/routes/"
	i := strings.Index(filePath, marker)
	if i == -1 {
		return ""
	}

	// extract everything after the marker
	path := filePath[i+len(marker):]

	// Check if the path ends with "/handler.go" and remove it if it does
	j := strings.LastIndex(path, suffix)
	if j != -1 && j == len(path)-len(suffix) {
		// store everything before the index of the "/handler.go"
		path = "/" + path[:j]
	}

	return path
}

// GET registers a GET request handler for a route
func GET(handler InertiaHandlerFunc) {
	// Extract the path automatically from the caller
	routePath := extractPath()
	if routePath == "" {
		panic("Failed to extract route path")
	}

	registerRoute("GET", routePath, handler)
}

// POST registers a POST request handler for a route
func POST(handler InertiaHandlerFunc) {
	// Extract the path automatically from the caller
	routePath := extractPath()
	if routePath == "" {
		panic("Failed to extract route path")
	}

	registerRoute("POST", routePath, handler)
}

// PUT registers a PUT request handler for a route
func PUT(handler InertiaHandlerFunc) {
	// Extract the path automatically from the caller
	routePath := extractPath()
	if routePath == "" {
		panic("Failed to extract route path")
	}

	registerRoute("PUT", routePath, handler)
}

// DELETE registers a DELETE request handler for a route
func DELETE(handler InertiaHandlerFunc) {
	// Extract the path automatically from the caller
	routePath := extractPath()
	if routePath == "" {
		panic("Failed to extract route path")
	}

	registerRoute("DELETE", routePath, handler)
}

// PATCH registers a PATCH request handler for a route
func PATCH(handler InertiaHandlerFunc) {
	// Extract the path automatically from the caller
	routePath := extractPath()
	if routePath == "" {
		panic("Failed to extract route path")
	}

	registerRoute("PATCH", routePath, handler)
}
