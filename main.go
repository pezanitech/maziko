package main

import (
	"net/http"
	"os"

	"github.com/pezanitech/maziko/core/router"
	"github.com/pezanitech/maziko/core/utils"
	"github.com/pezanitech/maziko/gen"

	"github.com/pezanitech/maziko/core/cmd"
)

func main() {
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

		utils.Logger.Info("Starting server on localhost:3000")

		if err := http.ListenAndServe(":3000", mux); err != nil {
			utils.Logger.Error("Server error", "error", err)
			os.Exit(1)
		}

	}
}
