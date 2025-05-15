package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean build artifacts",
	Long:  "Remove temporary files, build directories, and clean Go cache",
	Run:   runCleanCommand,
}

func init() {
	RootCmd.AddCommand(cleanCmd)
}

// promptForRemoval asks user for confirmation before removing a file/directory
func promptForRemoval(path string) bool {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting absolute path for %s: %v\n", path, err)
		absPath = path
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Remove %s? (y/n): ", absPath)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

// runCleanCommand runs the clean script
func runCleanCommand(cmd *cobra.Command, args []string) {
	// Remove directories
	dirs := []string{"tmp", "build", "ssrBuild", "node_modules", ".pnpm-store"}
	for _, dir := range dirs {
		// Skip if directory doesn't exist
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}

		if promptForRemoval(dir) {
			if err := os.RemoveAll(dir); err != nil {
				fmt.Fprintf(os.Stderr, "Error removing %s: %v\n", dir, err)
			} else {
				fmt.Printf("Removed %s\n", dir)
			}
		} else {
			fmt.Printf("Skipped removing %s\n", dir)
		}
	}

	// Clean Go cache
	fmt.Print("Clean Go cache? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	} else {
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			if err := runCommand("go", "clean", "-cache", "-testcache"); err != nil {
				fmt.Fprintf(os.Stderr, "Error cleaning Go cache: %v\n", err)
			} else {
				fmt.Println("Cleaned Go cache")
			}
		} else {
			fmt.Println("Skipped cleaning Go cache")
		}
	}
}