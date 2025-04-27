package cmd

import (
	"net/http"
	"os"

	"github.com/pezanitech/maziko/backend/config"
	"github.com/pezanitech/maziko/backend/handlers"
	"github.com/pezanitech/maziko/backend/utils"
)

func RunProd() {
	utils.InitLogger()

	i := config.InitInertia()

	mux := http.NewServeMux()

	mux.Handle("/", i.Middleware(handlers.RootHandler(i)))
	mux.Handle("/build/", handlers.BuildHandler())

	utils.Logger.Info("Starting server on localhost:3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		utils.Logger.Error("Server error", "error", err)
		os.Exit(1)
	}
}
