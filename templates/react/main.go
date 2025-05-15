package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pezanitech/maziko/apps/react/gen"
	"github.com/pezanitech/maziko/libs/core/config"
	"github.com/pezanitech/maziko/libs/core/logger"
	"github.com/pezanitech/maziko/libs/core/router"

	"github.com/pezanitech/maziko/libs/core/cmd"
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

	logger.Log.Info(
		"Starting server",
		"address", config.GetAppURL(),
		"port", port,
	)

	if err := http.ListenAndServe(":"+port, appRouter.Router); err != nil {
		logger.Log.Error(
			"Server error",
			"error", err,
		)
		os.Exit(1)
	}
}
