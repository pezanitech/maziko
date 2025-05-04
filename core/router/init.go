package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/logger"
	inertia "github.com/romsar/gonertia"
)

type AppRouter struct {
	Router   *chi.Mux
	renderer *inertia.Inertia
}

type InertiaHTTPHandler func(*inertia.Inertia, http.ResponseWriter, *http.Request)

// Global router instance
var appRouter AppRouter

// Stores route component names for handlers
var routeComponents map[string]string

func InitRouter(i *inertia.Inertia) AppRouter {
	appRouter = AppRouter{
		Router:   chi.NewRouter(),
		renderer: i,
	}

	// Initialize the route components map
	routeComponents = make(map[string]string)

	fileServer( // Serve static files (app/public/)
		appRouter.Router,
		"/",
		http.Dir(config.GetPublicDir()),
	)

	fileServer( // Serve build artifacts (build/)
		appRouter.Router,
		config.GetBuildPrefix(),
		http.Dir(config.GetBuildDir()),
	)

	return appRouter
}

// Helper function to serve files from a directory
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		path += "/"
	}

	fileServer := http.StripPrefix(path, http.FileServer(root))

	r.Get(path+"*", func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Debug(
			"Serving static file",
			"path", r.URL.Path,
		)
		fileServer.ServeHTTP(w, r)
	})
}

// Registers a GET route handler
func GET(handler InertiaHTTPHandler) {
	routePath := determineRoutePath()
	componentName := determineComponentName()

	logger.Logger.Info(
		"Registering route",
		"path", routePath,
		"component", componentName,
	)

	// store component name for this route
	routeComponents[routePath] = componentName

	appRouter.Router.Get(routePath,
		func(w http.ResponseWriter, r *http.Request) {
			handler(appRouter.renderer, w, r)
		},
	)
}
