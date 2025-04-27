package cmd

import (
	"net/http"

	"github.com/pezanitech/maziko/backend/handlers"
)

func RunProd() {
	i := initInertia()

	mux := http.NewServeMux()

	mux.Handle("/", i.Middleware(handlers.RootHandler(i)))
	mux.Handle("/build/", handlers.BuildHandler())

	http.ListenAndServe(":3000", mux)
}
