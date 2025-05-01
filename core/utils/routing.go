package utils

import (
	"net/http"
	"os"
	"path"

	"github.com/pezanitech/maziko/core/config"
)

func HandleRequest(w http.ResponseWriter, r *http.Request, h http.Handler) {
	h.ServeHTTP(w, r)
}

var StaticFileHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		resourcePath := path.Join(config.GetPublicDir(), r.URL.Path)

		if _, err := os.Stat(resourcePath); err == nil {
			http.ServeFile(w, r, resourcePath) // serve file
			return
		}
		http.NotFound(w, r) // 404 error
	},
)

var BuildDirHandler = http.StripPrefix(
	config.GetBuildPrefix(),
	http.FileServer(http.Dir(config.GetBuildDir())),
)
