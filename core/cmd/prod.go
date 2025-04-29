package cmd

import (
	"net/http"
	"os"

	"github.com/pezanitech/maziko/core/router"
	"github.com/pezanitech/maziko/core/utils"
)

func RunProd() {
	// Initialize logger before use
	utils.InitLogger()

	i := router.InitInertia()

	mux := http.NewServeMux()

	mux.Handle("/", i.Middleware(router.Router(i)))

	utils.Logger.Info("Starting server on localhost:3000")

	if err := http.ListenAndServe(":3000", mux); err != nil {
		utils.Logger.Error("Server error", "error", err)
		os.Exit(1)
	}
}
