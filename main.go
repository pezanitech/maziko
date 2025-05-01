package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/utils"
	"github.com/pezanitech/maziko/gen"

	"github.com/pezanitech/maziko/core/cmd"
)

func main() {
	// initialize config
	if err := config.Initialize(); err != nil {
		fmt.Printf(
			"Failed to load configuration: %v\n", err,
		)
		os.Exit(1)
	}

	// initialize logger
	utils.InitLogger()

	switch {
	case len(os.Args) > 1 && os.Args[1] == "dev":
		cmd.RunDev()

	case len(os.Args) > 1 && os.Args[1] == "genroutes":
		cmd.GenerateRoutes()

	default:
		RunProd()
	}
}

func RunProd() {
	i := cmd.InitRenderer()
	mux := http.NewServeMux()
	port := fmt.Sprintf("%d", config.GetAppPort())

	// register routes
	mux.Handle("/", i.Middleware(gen.Routes(i)))

	utils.Logger.Info(
		"Starting server",
		"address", config.GetAppURL(),
		"port", port,
	)

	if err := http.ListenAndServe(port, mux); err != nil {
		utils.Logger.Error(
			"Server error",
			"error", err,
		)
		os.Exit(1)
	}
}
