package cmd

import (
	"net/http"
	"os"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/utils"
)

func RunProd() {
	// Initialize logger before use
	utils.InitLogger()

	i := config.InitInertia()

	mux := http.NewServeMux()

	mux.Handle("/", i.Middleware(config.Router(i)))

	utils.Logger.Info("Starting server on localhost:3000")

	if err := http.ListenAndServe(":3000", mux); err != nil {
		utils.Logger.Error("Server error", "error", err)
		os.Exit(1)
	}
}
