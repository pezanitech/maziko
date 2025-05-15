package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Long:  "Build Vite assets, generate routes, and build Go code",
	Run:   runBuildCommand,
}

func init() {
	RootCmd.AddCommand(buildCmd)
}

// runBuildCommand runs the build script
func runBuildCommand(cmd *cobra.Command, args []string) {
	// First build Vite
	if err := buildVite(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building Vite: %v\n", err)
		os.Exit(1)
	}

	// Generate routes
	if err := runCommand("go", "run", "main.go", "genroutes"); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating routes: %v\n", err)
		os.Exit(1)
	}

	// Build Go code
	if err := buildGo(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building Go code: %v\n", err)
		os.Exit(1)
	}
}