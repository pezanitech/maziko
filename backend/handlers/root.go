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
			handleRoot(i, w, r)
		default:
			handleFiles(w, r)
		}
	})
}

// handleRoot handles requests to the root endpoint ("/").
// It renders the "index" page with provided props.
//
// Parameters:
//   - i: The Inertia instance used for rendering the page
//   - w: The HTTP response writer to write the response to
//   - r: The HTTP request containing client information
//
// If rendering fails, it delegates error handling to the utils.HandleServerErr function.
func handleRoot(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
	err := i.Render(w, r, "index", inertia.Props{
		"line1": "A full-stack framework",
		"line2": "built with Inertia.js and Go! 💚",
	})

	if err != nil {
		utils.HandleServerErr(w, err)
	}
}

// handleFiles serves static files from the "public" directory to HTTP clients.
// If the requested file exists in the public directory, it serves it directly.
// Otherwise, it returns a 404 Not Found error to the client.
//
// Parameters:
//   - w: The HTTP response writer to write the response to
//   - r: The HTTP request containing the path to the requested file
func handleFiles(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat(path.Join("public", r.URL.Path)); err == nil {
		http.ServeFile(w, r, path.Join("public", r.URL.Path))
		return
	}

	http.NotFound(w, r)
}
