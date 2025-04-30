package main

import (
	"os"

	"github.com/pezanitech/maziko/core/cmd"
)

func main() {
	switch {
	case len(os.Args) > 1 && os.Args[1] == "dev":
		cmd.RunDev()
	case len(os.Args) > 1 && os.Args[1] == "genroutes":
		cmd.GenerateRoutes()
	default:
		cmd.RunProd()
	}
}
