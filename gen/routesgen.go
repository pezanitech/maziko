package gen

import (
	"net/http"

	inertia "github.com/romsar/gonertia"
)

func Routes(i *inertia.Inertia) http.Handler {
	// Return a basic handler that responds with a 404 Not Found
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
}
