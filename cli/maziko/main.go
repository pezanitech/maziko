package main

import (
	"fmt"
	"os"

	"github.com/pezanitech/maziko/cli/maziko/cmd"
)

func main() {
	// initialize root command
	cmd.InitRootCmd()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
