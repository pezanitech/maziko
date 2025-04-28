package main

import (
	"os"

	"github.com/pezanitech/maziko/core/cmd"
)

func main() {
	switch true {
	case len(os.Args) > 1 && os.Args[1] == "dev":
		cmd.RunDev()
	default:
		cmd.RunProd()
	}
}
