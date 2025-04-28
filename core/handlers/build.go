package handlers

import (
	"net/http"
)

// Creates an HTTP handler that serves static
// build files from the "./build" directory.
//
// It strips the "/build/" prefix from the URL path before looking for the file.
//
// Returns:
//   - http.Handler: A handler that serves files from the build directory
func BuildHandler() http.Handler {
	return http.StripPrefix("/build/", http.FileServer(http.Dir("./build")))
}
