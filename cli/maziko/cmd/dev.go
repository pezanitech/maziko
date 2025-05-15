package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start development server",
	Long:  "Generate routes, build Go code temporarily, and start development servers for both Vite and Go",
	Run:   runDevCommand,
}

func init() {
	RootCmd.AddCommand(devCmd)
}

// runDevCommand runs the dev script
func runDevCommand(cmd *cobra.Command, args []string) {
	// First run genroutes
	if err := runCommand("go", "run", "main.go", "genroutes"); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating routes: %v\n", err)
		os.Exit(1)
	}

	// Build Go code temporarily
	if err := buildGoTemp(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building Go code: %v\n", err)
		os.Exit(1)
	}

	// Start concurrent Vite and Go servers
	c1 := exec.Command("pnpm", "vite")
	c2 := exec.Command("./tmp/main", "dev")

	c1.Stdout = os.Stdout
	c1.Stderr = os.Stderr
	c2.Stdout = os.Stdout
	c2.Stderr = os.Stderr

	c1.Start()
	c2.Start()

	// Wait for both to complete
	c1.Wait()
	c2.Wait()
}
