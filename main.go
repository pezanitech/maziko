package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/logger"
	"github.com/pezanitech/maziko/core/router"
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
	logger.InitLogger()

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
	port := fmt.Sprintf("%d", config.GetAppPort())

	appRouter := router.InitRouter(i)

	gen.Routes()

	logger.Logger.Info(
		"Starting server",
		"address", config.GetAppURL(),
		"port", port,
	)

	if err := http.ListenAndServe(":"+port, appRouter.Router); err != nil {
		logger.Logger.Error(
			"Server error",
			"error", err,
		)
		os.Exit(1)
	}
}
