package router

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/utils"
	inertia "github.com/romsar/gonertia"
)

type AppRouter struct {
	Router   *chi.Mux
	renderer *inertia.Inertia
}

type InertiaHTTPHandler func(*inertia.Inertia, http.ResponseWriter, *http.Request)

// Global router instance
var appRouter AppRouter

func InitRouter(i *inertia.Inertia) AppRouter {
	appRouter = AppRouter{
		Router:   chi.NewRouter(),
		renderer: i,
	}

	// Serve static files from public directory
	fileServer(appRouter.Router, "/", http.Dir(config.GetPublicDir()))

	// Serve build artifacts
	fileServer(appRouter.Router, config.GetBuildPrefix(), http.Dir(config.GetBuildDir()))

	return appRouter
}

// Helper function to serve files from a directory
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		path += "/"
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	r.Get(path+"*", func(w http.ResponseWriter, r *http.Request) {
		utils.Logger.Debug(
			"Serving static file",
			"path", r.URL.Path,
		)
		fs.ServeHTTP(w, r)
	})
}

// Registers a GET route handler
func GET(handler InertiaHTTPHandler) {
	routePath := determineRoutePath()

	utils.Logger.Info(
		"Registering route",
		"path", routePath,
	)

	appRouter.Router.Get(routePath,
		func(w http.ResponseWriter, r *http.Request) {
			handler(appRouter.renderer, w, r)
		},
	)
}

// Extracts a routes path
func determineRoutePath() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		utils.Logger.Warn(
			"Could not determine route path, using root path",
		)
		return "/"
	}

	callerFunction := runtime.FuncForPC(pc).Name()

	utils.Logger.Debug(
		"Route caller function",
		"function", callerFunction,
	)

	// extract the package path
	parts := strings.Split(callerFunction, "/")

	if len(parts) >= 2 {
		lastPart := parts[len(parts)-1]
		pkgFunc := strings.Split(lastPart, ".")

		if len(pkgFunc) >= 1 {
			pkgName := pkgFunc[0]
			if pkgName == "index" {
				return "/"
			}
			return "/" + pkgName
		}
	}

	utils.Logger.Warn(
		"Could not parse route path from caller, using root path",
	)

	return "/"
}
