package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pezanitech/maziko/libs/core/config"
	"github.com/pezanitech/maziko/libs/core/logger"
)

type AppRouter struct {
	Router   *chi.Mux
	renderer Inertia
}

// Global appRouter instance
var appRouter AppRouter

// Stores route component names for handlers
var routeComponents map[string]string

// InitRouter initializes and returns a global AppRouter instance
func InitRouter(i Inertia) AppRouter {
	appRouter = AppRouter{
		Router:   chi.NewRouter(),
		renderer: i,
	}

	// initialize route components map
	routeComponents = make(map[string]string)

	fileServer( // serve static files (app/public/)
		appRouter.Router,
		"/",
		http.Dir(config.GetPublicDir()),
	)

	fileServer( // serve build artifacts (build/)
		appRouter.Router,
		config.GetBuildPrefix(),
		http.Dir(config.GetBuildDir()),
	)

	return appRouter
}

// fileServer serves files from a directory
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		path += "/"
	}

	server := http.StripPrefix(path, http.FileServer(root))

	r.Get(path+"*", func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Debug(
			"Serving static file",
			"path", r.URL.Path,
		)

		server.ServeHTTP(w, r)
	})
}
