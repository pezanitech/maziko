package router

import (
	"net/http"

	"github.com/pezanitech/maziko/libs/core/logger"
)

// Registers a route with an HTTP method
func registerRoute(method string, handler InertiaHTTPHandler) {
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
		appRouter.Router.Get(
			routePath,
			func(w http.ResponseWriter, r *http.Request) {
				handler(appRouter.renderer, w, r)
			},
		)
	case "POST":
		appRouter.Router.Post(
			routePath,
			func(w http.ResponseWriter, r *http.Request) {
				handler(appRouter.renderer, w, r)
			},
		)
	case "PUT":
		appRouter.Router.Put(
			routePath,
			func(w http.ResponseWriter, r *http.Request) {
				handler(appRouter.renderer, w, r)
			},
		)
	case "DELETE":
		appRouter.Router.Delete(
			routePath,
			func(w http.ResponseWriter, r *http.Request) {
				handler(appRouter.renderer, w, r)
			},
		)
	case "PATCH":
		appRouter.Router.Patch(
			routePath,
			func(w http.ResponseWriter, r *http.Request) {
				handler(appRouter.renderer, w, r)
			},
		)
	}
}

// Registers a GET request handler for a route
func GET(handler InertiaHTTPHandler) {
	registerRoute("GET", handler)
}

// Registers a POST request handler for a route
func POST(handler InertiaHTTPHandler) {
	registerRoute("POST", handler)
}

// Registers a PUT request handler for a route
func PUT(handler InertiaHTTPHandler) {
	registerRoute("PUT", handler)
}

// Registers a DELETE request handler for a route
func DELETE(handler InertiaHTTPHandler) {
	registerRoute("DELETE", handler)
}

// Registers a PATCH request handler for a route
func PATCH(handler InertiaHTTPHandler) {
	registerRoute("PATCH", handler)
}
