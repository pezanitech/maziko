package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// genroutesCmd represents the genroutes command
var genroutesCmd = &cobra.Command{
	Use:   "genroutes",
	Short: "Generate routes",
	Long:  "Generate routes for the application",
	Run:   runGenroutesCommand,
}

func init() {
	RootCmd.AddCommand(genroutesCmd)
}

// runGenroutesCommand runs the genroutes script
func runGenroutesCommand(cmd *cobra.Command, args []string) {
	// Find current directory's main.go file or in parent directory
	mainPath := "main.go"
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		// Try parent directory
		parentMain := filepath.Join("..", "main.go")
		if _, err := os.Stat(parentMain); err == nil {
			mainPath = parentMain
		} else {
			fmt.Fprintf(os.Stderr, "main.go not found\n")
			os.Exit(1)
		}
	}

	if err := runCommand("go", "run", mainPath, "genroutes"); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating routes: %v\n", err)
		os.Exit(1)
	}
}