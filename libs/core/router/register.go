package router

import (
	"net/http"

	"github.com/pezanitech/maziko/libs/core/logger"
)

// registerRoute registers a route with an HTTP method
func registerRoute(method string, handler InertiaHandlerFunc) {
	routePath := determineRoutePath()
	componentName := determineComponentName()

	logger.Log.Info(
		"Registering route",
		"path", routePath,
		"component", componentName,
		"method", method,
	)

	// store component name for this route
	routeComponents[routePath] = componentName

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

// GET registers a GET request handler for a route
func GET(handler InertiaHandlerFunc) {
	registerRoute("GET", handler)
}

// POST registers a POST request handler for a route
func POST(handler InertiaHandlerFunc) {
	registerRoute("POST", handler)
}

// PUT registers a PUT request handler for a route
func PUT(handler InertiaHandlerFunc) {
	registerRoute("PUT", handler)
}

// DELETE registers a DELETE request handler for a route
func DELETE(handler InertiaHandlerFunc) {
	registerRoute("DELETE", handler)
}

// PATCH registers a PATCH request handler for a route
func PATCH(handler InertiaHandlerFunc) {
	registerRoute("PATCH", handler)
}
