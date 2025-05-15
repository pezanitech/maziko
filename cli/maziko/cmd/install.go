package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dependencies",
	Long:  "Install both npm dependencies using pnpm and Go dependencies",
	Run:   runInstallCommand,
}

func init() {
	RootCmd.AddCommand(installCmd)
}

// runInstallCommand runs the install script
func runInstallCommand(cmd *cobra.Command, args []string) {
	// First run pnpm install
	fmt.Println("Installing npm dependencies with pnpm...")
	if err := runCommand("pnpm", "i"); err != nil {
		fmt.Fprintf(os.Stderr, "Error installing npm dependencies: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("npm dependencies installed successfully")

	// Then run go mod tidy
	fmt.Println("Installing Go dependencies with go mod tidy...")
	if err := runCommand("go", "mod", "tidy"); err != nil {
		fmt.Fprintf(os.Stderr, "Error installing Go dependencies: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Go dependencies installed successfully")
}