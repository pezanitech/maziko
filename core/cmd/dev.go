package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/pezanitech/maziko/core/config"
	"github.com/pezanitech/maziko/core/utils"

	"github.com/fsnotify/fsnotify"
)

func RunDev() {
	utils.Logger.Info("Starting development mode...")

	// Get config values
	rootDir := config.GetDevRootDir()
	excludeRegexes := config.GetDevExcludeRegexes()
	excludeDirs := config.GetDevExcludeDirs()
	includeExts := config.GetDevIncludeExts()
	buildDelay := config.GetDevBuildDelay()

	// Get build directory values
	tmpDir := config.GetTempDir()
	buildDir := config.GetBuildDir()
	ssrDir := config.GetSSRDir()

	// Add the build directories to excludeDirs
	excludeDirs = append(excludeDirs, tmpDir, buildDir, ssrDir)

	// Set binary path to be inside the temp directory
	binPath := filepath.Join(tmpDir, "main")
	buildCmd := "go build -o " + binPath + " ."

	// create temp directory if it doesn't exist
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		utils.Logger.Error("Failed to create tmp directory", "error", err)
		os.Exit(1)
	}

	// create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		utils.Logger.Error("Failed to create file watcher", "error", err)
		os.Exit(1)
	}
	defer watcher.Close()

	// start build and run process
	var cmd *exec.Cmd

	buildAndRun := func() {
		// kill any running process
		if cmd != nil && cmd.Process != nil {
			utils.Logger.Info("stopping process...")

			if err := cmd.Process.Kill(); err != nil {
				utils.Logger.Error("Failed to kill process", "error", err)
			}
			cmd.Wait() // Wait for process to exit
		}

		// build the application
		utils.Logger.Info("Building application...")
		buildCmdExec := exec.Command("sh", "-c", buildCmd)
		buildCmdExec.Stdout = os.Stdout
		buildCmdExec.Stderr = os.Stderr
		if err := buildCmdExec.Run(); err != nil {
			utils.Logger.Error("Build failed", "error", err)
			return
		}

		// Run the application
		utils.Logger.Info("Starting application...")
		cmd = exec.Command(binPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			utils.Logger.Error("Failed to start application", "error", err)
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
		utils.Logger.Error("Failed to add directories to watcher", "error", err)
		os.Exit(1)
	}

	// Timer for debouncing
	var debounceTimer *time.Timer

	// Watch for file changes
	utils.Logger.Info("Watching for file changes...")
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
				utils.Logger.Info("File modified", "file", event.Name)

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
			utils.Logger.Error("Watcher error", "error", err)
		}
	}
}
