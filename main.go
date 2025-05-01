package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/router"
	"github.com/pezanitech/maziko/core/utils"
	"github.com/pezanitech/maziko/gen"

	"github.com/pezanitech/maziko/core/cmd"
)

func main() {
	// Initialize configuration
	if err := config.Initialize(); err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) > 1 && os.Args[1] == "dev":
		cmd.RunDev()
	case len(os.Args) > 1 && os.Args[1] == "genroutes":
		cmd.GenerateRoutes()
	default:
		// Initialize logger before use
		utils.InitLogger()

		i := router.InitInertia()

		mux := http.NewServeMux()

		mux.Handle("/", i.Middleware(gen.Routes(i)))

		port := fmt.Sprintf(":%d", config.GetAppPort())
		utils.Logger.Info("Starting server", "address", config.GetAppURL(), "port", port)

		if err := http.ListenAndServe(port, mux); err != nil {
			utils.Logger.Error("Server error", "error", err)
			os.Exit(1)
		}
	}
}
