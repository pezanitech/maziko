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

// RunDev starts the application in development mode with hot reloading
func RunDev() {
	utils.Logger.Info("Starting development mode...")

	// Configure development environment
	setupConfig := newDevConfig()

	// Set up file watcher
	watcher := setupFileWatcher()
	defer watcher.Close()

	// Start initial build and set up watch directories
	var cmd *exec.Cmd
	buildAndRun := createBuildAndRunFunc(setupConfig.binPath, &cmd)

	// Initial build and run
	buildAndRun()

	// Add directories to watch
	addDirectoriesToWatch(watcher, setupConfig)

	// Watch for file changes
	watchForChanges(watcher, setupConfig, buildAndRun)
}

// devConfig holds all configuration parameters for dev mode
type devConfig struct {
	rootDir        string
	excludeRegexes []string
	excludeDirs    []string
	includeExts    []string
	buildDelay     time.Duration
	tmpDir         string
	buildDir       string
	binPath        string
	buildCmd       string
}

// newDevConfig loads all configuration settings for development mode
func newDevConfig() devConfig {
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

	// Create temp directory if it doesn't exist
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		utils.Logger.Error("Failed to create tmp directory", "error", err)
		os.Exit(1)
	}

	return devConfig{
		rootDir:        rootDir,
		excludeRegexes: excludeRegexes,
		excludeDirs:    excludeDirs,
		includeExts:    includeExts,
		buildDelay:     buildDelay,
		tmpDir:         tmpDir,
		buildDir:       buildDir,
		binPath:        binPath,
		buildCmd:       buildCmd,
	}
}

// setupFileWatcher creates and configures a new file watcher
func setupFileWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		utils.Logger.Error("Failed to create file watcher", "error", err)
		os.Exit(1)
	}
	return watcher
}

// createBuildAndRunFunc returns a function that builds and runs the application
func createBuildAndRunFunc(binPath string, cmdRef **exec.Cmd) func() {
	return func() {
		// Kill any running process
		cmd := *cmdRef
		if cmd != nil && cmd.Process != nil {
			utils.Logger.Info("Stopping process...")

			if err := cmd.Process.Kill(); err != nil {
				utils.Logger.Error("Failed to kill process", "error", err)
			}
			cmd.Wait() // Wait for process to exit
		}

		// Build the application
		utils.Logger.Info("Building application...")
		buildCmdExec := exec.Command("sh", "-c", "go build -o "+binPath+" .")
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
		*cmdRef = cmd
	}
}

// addDirectoriesToWatch walks through the project and adds directories to the watcher
func addDirectoriesToWatch(watcher *fsnotify.Watcher, config devConfig) {
	err := filepath.Walk(
		config.rootDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip excluded directories
			if info.IsDir() {
				for _, excludeDir := range config.excludeDirs {
					if strings.Contains(path, excludeDir) {
						return filepath.SkipDir
					}
				}
				return watcher.Add(path)
			}
			return nil
		})

	if err != nil {
		utils.Logger.Error("Failed to add directories to watcher", "error", err)
		os.Exit(1)
	}
}

// watchForChanges monitors file changes and triggers rebuilds when needed
func watchForChanges(watcher *fsnotify.Watcher, config devConfig, buildAndRun func()) {
	// Timer for debouncing
	var debounceTimer *time.Timer

	utils.Logger.Info("Watching for file changes...")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if shouldSkipFile(event.Name, config) {
				continue
			}

			// Check if the event is a file modification
			if event.Op&fsnotify.Write == fsnotify.Write {
				utils.Logger.Info("File modified", "file", event.Name)

				// Reset the debounce timer
				if debounceTimer != nil {
					debounceTimer.Stop()
				}
				debounceTimer = time.AfterFunc(config.buildDelay, buildAndRun)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			utils.Logger.Error("Watcher error", "error", err)
		}
	}
}

// shouldSkipFile determines if a file should be ignored based on extension and patterns
func shouldSkipFile(filename string, config devConfig) bool {
	// Check if the file has one of the included extensions
	hasIncludedExt := false
	for _, ext := range config.includeExts {
		if strings.HasSuffix(filename, ext) {
			hasIncludedExt = true
			break
		}
	}

	// Skip files that don't have included extensions
	if !hasIncludedExt {
		return true
	}

	// Skip files that match exclude regexes
	for _, excludeRegex := range config.excludeRegexes {
		if strings.Contains(filename, excludeRegex) {
			return true
		}
	}

	return false
}
