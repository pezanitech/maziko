package cmd

import (
	"net/http"

	"github.com/pezanitech/maziko/backend/config"
	"github.com/pezanitech/maziko/backend/handlers"
)

func RunProd() {
	i := config.InitInertia()

	mux := http.NewServeMux()

	mux.Handle("/", i.Middleware(handlers.RootHandler(i)))
	mux.Handle("/build/", handlers.BuildHandler())

	http.ListenAndServe(":3000", mux)
}
