package config

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/pezanitech/maziko/app/routes/index"
	inertia "github.com/romsar/gonertia"
)

func Router(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch true {
		case r.URL.Path == "/" && r.Method == http.MethodGet:
			index.GET(i, w, r)
		case r.URL.Path == "/" && r.Method == http.MethodPost:
			index.POST(i, w, r)
		case r.URL.Path == "/news" && r.Method == http.MethodGet:
			index.GET(i, w, r)
		case strings.HasPrefix(r.URL.Path, "/build/"):
			buildDirHandler().ServeHTTP(w, r)
		default:
			staticFileHandler().ServeHTTP(w, r)
		}
	})
}

func staticFileHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(path.Join("app/public", r.URL.Path)); err == nil {
			http.ServeFile(w, r, path.Join("app/public", r.URL.Path))
			return
		}

		// If none exist, respond with a 404 error
		http.NotFound(w, r)
	})
}

func buildDirHandler() http.Handler {
	return http.StripPrefix(
		"/build/",
		http.FileServer(http.Dir("./build")),
	)
}
