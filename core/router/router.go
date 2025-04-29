package router

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pezanitech/maziko/app/routes/index"
	"github.com/pezanitech/maziko/app/routes/news"
	"github.com/pezanitech/maziko/core/utils"
	inertia "github.com/romsar/gonertia"
)

func Router(i *inertia.Inertia) http.Handler {
	if err := filepath.Walk("./app/routes", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if path != "./app/routes" {
				fmt.Println("Path:", path)
			}
		}

		return nil
	}); err != nil {
		utils.Logger.Error("Error walking the path", "error", err)
		return nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch true {
		case r.URL.Path == "/" && r.Method == http.MethodGet:
			index.GET(i, w, r)
		case r.URL.Path == "/" && r.Method == http.MethodPost:
			index.POST(i, w, r)
		case r.URL.Path == "/news" && r.Method == http.MethodGet:
			news.GET(i, w, r)
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
