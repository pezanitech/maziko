package handlers

import (
	"net/http"
	"os"
	"path"

	"github.com/pezanitech/maziko/backend/utils"
	inertia "github.com/romsar/gonertia"
)

func RootHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			err := i.Render(w, r, "index", inertia.Props{
				"line1": "A full-stack framework",
				"line2": "built with Inertia.js and Go! 💚",
			})

			if err != nil {
				utils.HandleServerErr(w, err)
			}
		default:
			// if file exists in public serve it
			if _, err := os.Stat(path.Join("public", r.URL.Path)); err == nil {
				http.ServeFile(w, r, path.Join("public", r.URL.Path))
				return
			}

			// otherwise return 404
			http.NotFound(w, r)
		}
	})
}
