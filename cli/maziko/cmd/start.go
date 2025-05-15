package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the production server",
	Long:  "Start the main application server",
	Run:   runStartCommand,
}

func init() {
	RootCmd.AddCommand(startCmd)
}

// Runs the start command
func runStartCommand(cmd *cobra.Command, args []string) {
	if err := runCommand("./build/main"); err != nil {
		os.Exit(1)
	}
}
