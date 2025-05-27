package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version of the CLI
const Version = "0.1.6"

// Represents the base command when called without subcommands
var RootCmd = &cobra.Command{
	Use: "maziko",
	Short: fmt.Sprintf(
		"%s\n%s\n",
		"Maziko - A full-stack framework",
		"built with Inertia.js and Go",
	),
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		// no subcommand, print help
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func InitRootCmd() {
	// define global flags and options.
	RootCmd.PersistentFlags().BoolP(
		"verbose", "v", false,
		"Enable verbose output (not implemented)",
	)
}
