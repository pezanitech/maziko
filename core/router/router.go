package router

import (
	"net/http"

	"github.com/pezanitech/maziko/core/gen"
	inertia "github.com/romsar/gonertia"
)

func Router(i *inertia.Inertia) http.Handler {
	return gen.DefineRoutes(i)
}
