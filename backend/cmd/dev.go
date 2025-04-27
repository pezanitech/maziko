package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	rootDir     = "."
	tmpDir      = ".tmp"
	binPath     = "./.tmp/main"
	buildCmd    = "go build -o ./.tmp/main ."
	buildDelay  = 1000 * time.Millisecond
	excludeDirs = []string{
		"assets",
		".tmp",
		"vendor",
		"testdata",
		"node_modules",
		"frontend",
		"bin",
		"public",
		"ssrBuild",
		"build",
	}
	excludeRegexes = []string{"_test.go"}
	includeExts    = []string{
		".go",
		".tpl",
		".tmpl",
		".html",
	}
)

func Dev() {
	// create tmp directory if it doesn't exist

	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		log.Fatalf(
			"Failed to create tmp directory: %v",
			err,
		)
	}

	// create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf(
			"Failed to create file watcher: %v",
			err,
		)
	}
	defer watcher.Close()

	// start build and run process
	var cmd *exec.Cmd

	buildAndRun := func() {
		// kill any running process
		if cmd != nil && cmd.Process != nil {
			fmt.Println("\nstopping process...")

			if err := cmd.Process.Kill(); err != nil {
				log.Printf("Failed to kill process: %v", err)
			}
			cmd.Wait() // Wait for process to exit
		}

		// build the application
		fmt.Println("Building application...")
		buildCmdExec := exec.Command("sh", "-c", buildCmd)
		buildCmdExec.Stdout = os.Stdout
		buildCmdExec.Stderr = os.Stderr
		if err := buildCmdExec.Run(); err != nil {
			log.Printf("Build failed: %v", err)
			return
		}

		// Run the application
		fmt.Println("Starting application...")
		cmd = exec.Command(binPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			log.Printf("Failed to start application: %v", err)
			return
		}
	}

	// Initial build and run
	buildAndRun()

	// Add directories to watch
	if err := filepath.Walk(
		rootDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip excluded directories
			if info.IsDir() {
				for _, excludeDir := range excludeDirs {
					if strings.Contains(path, excludeDir) {
						return filepath.SkipDir
					}
				}

				return watcher.Add(path)
			}

			return nil
		}); err != nil {
		log.Fatalf("Failed to add directories to watcher: %v", err)
	}

	// Timer for debouncing
	var debounceTimer *time.Timer

	// Watch for file changes
	fmt.Println("Watching for file changes...")
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			// Check if the file has one of the included extensions
			hasIncludedExt := false
			for _, ext := range includeExts {
				if strings.HasSuffix(event.Name, ext) {
					hasIncludedExt = true
					break
				}
			}

			// Skip files that don't have included extensions
			if !hasIncludedExt {
				continue
			}

			// Skip files that match exclude regexes
			skip := false
			for _, excludeRegex := range excludeRegexes {
				if strings.Contains(event.Name, excludeRegex) {
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			// Check if the event is a file modification
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("File modified: %s\n", event.Name)

				// Reset the debounce timer
				if debounceTimer != nil {
					debounceTimer.Stop()
				}
				debounceTimer = time.AfterFunc(buildDelay, buildAndRun)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Watcher error: %v", err)
		}
	}
}
