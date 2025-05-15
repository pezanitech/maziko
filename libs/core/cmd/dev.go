package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/pezanitech/maziko/libs/core/config"
	"github.com/pezanitech/maziko/libs/core/errors"
	"github.com/pezanitech/maziko/libs/core/logger"

	"github.com/fsnotify/fsnotify"
)

// Holds all configuration parameters for dev mode
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

// Starts the app in dev mode with hot reloading
func RunDev() {
	logger.Log.Info("Starting development mode...")

	// configure development environment
	setupConfig := newDevConfig()

	// set up file watcher
	watcher := setupFileWatcher()
	defer watcher.Close()

	// start initial build and set up watch directories
	var cmd *exec.Cmd
	buildAndRun := createBuildAndRunFunc(setupConfig.binPath, &cmd)

	// initial build and run
	buildAndRun()

	// add directories to watch
	addDirectoriesToWatch(watcher, setupConfig)

	// watch for file changes
	watchForChanges(watcher, setupConfig, buildAndRun)
}

// Creates an instance of devConfig with global config values
func newDevConfig() devConfig {
	// get config values
	rootDir := config.GetDevRootDir()
	excludeRegexes := config.GetDevExcludeRegexes()
	excludeDirs := config.GetDevExcludeDirs()
	includeExts := config.GetDevIncludeExts()
	buildDelay := config.GetDevBuildDelay()

	// get build directory values
	tmpDir := config.GetTempDir()
	buildDir := config.GetBuildDir()
	ssrDir := config.GetSSRDir()

	// add the build directories to excludeDirs
	excludeDirs = append(excludeDirs, tmpDir, buildDir, ssrDir)

	// set binary output path to temp directory
	binPath := filepath.Join(tmpDir, "main")
	buildCmd := "go build -o " + binPath + " ."

	// create temp directory if it doesn't exist
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		errors.HandleFatalError(
			"Failed to create tmp directory", err,
		)
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

// Createss a new file watcher
func setupFileWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		errors.HandleFatalError(
			"Failed to create file watcher", err,
		)
	}
	return watcher
}

// Creates a function to build and run the application
func createBuildAndRunFunc(binPath string, cmdRef **exec.Cmd) func() {
	return func() {
		// kill any running process
		stopProcess(cmdRef)

		// build and run the application
		if buildApp(binPath) {
			startApp(binPath, cmdRef)
		}
	}
}

// Kills the currently running process if it exists
func stopProcess(cmdRef **exec.Cmd) {
	cmd := *cmdRef

	if cmd != nil && cmd.Process != nil {
		logger.Log.Info("Stopping process...")

		if err := cmd.Process.Kill(); err != nil {
			logger.Log.Error(
				"Failed to kill process",
				"error", err,
			)
		}
		cmd.Wait() // Wait for process to exit
	}
}

// Executes the build command for the app
func buildApp(binPath string) bool {
	logger.Log.Info("Building application...")

	buildCmdExec := exec.Command("sh", "-c", "go build -o "+binPath+" .")
	buildCmdExec.Stdout = os.Stdout
	buildCmdExec.Stderr = os.Stderr

	if err := buildCmdExec.Run(); err != nil {
		logger.Log.Error(
			"Build failed",
			"error", err,
		)
		return false
	}
	return true
}

// Runs the built application
func startApp(binPath string, cmdRef **exec.Cmd) {
	logger.Log.Info("Starting application...")

	cmd := exec.Command(binPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		logger.Log.Error(
			"Failed to start application",
			"error", err,
		)
		return
	}
	*cmdRef = cmd
}

// Walks the project and adds directories to the watcher
func addDirectoriesToWatch(watcher *fsnotify.Watcher, config devConfig) {
	err := filepath.Walk(
		config.rootDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip excluded directories
			if info.IsDir() {
				if shouldSkipDir(path, config.excludeDirs) {
					return filepath.SkipDir
				}
				return watcher.Add(path)
			}
			return nil
		})

	if err != nil {
		errors.HandleFatalError(
			"Failed to add directories to watcher", err,
		)
	}
}

// Checks if a directory should be excluded from watching
func shouldSkipDir(path string, excludeDirs []string) bool {
	for _, excludeDir := range excludeDirs {
		if strings.Contains(path, excludeDir) {
			return true
		}
	}
	return false
}

// Monitors file changes and triggers rebuilds when needed
func watchForChanges(watcher *fsnotify.Watcher, config devConfig, buildAndRun func()) {
	// timer for debouncing
	var debounceTimer *time.Timer

	logger.Log.Info("Watching for file changes...")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if shouldSkipFile(event.Name, config) {
				continue
			}

			// check if the event is a file modification
			if event.Op&fsnotify.Write == fsnotify.Write {
				logger.Log.Info(
					"File modified",
					"file", event.Name,
				)

				debounceRebuild(
					&debounceTimer,
					config.buildDelay,
					event.Name,
					buildAndRun,
				)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Log.Error(
				"Watcher error",
				"error", err,
			)
		}
	}
}

// debounceRebuild delays rebuilding to avoid multiple rapid rebuilds
func debounceRebuild(
	timerRef **time.Timer,
	delay time.Duration,
	filename string,
	buildFunc func(),
) {
	// log which file triggered the rebuild
	logger.Log.Debug(
		"Scheduling rebuild",
		"trigger", filename,
		"delay", delay,
	)

	// reset the debounce timer
	if *timerRef != nil {
		(*timerRef).Stop()
	}
	*timerRef = time.AfterFunc(delay, buildFunc)
}

// Checks if file should be ignored based on extension and patterns
func shouldSkipFile(filename string, config devConfig) bool {
	// check if file has an included extension
	hasIncludedExt := false

	for _, ext := range config.includeExts {
		if strings.HasSuffix(filename, ext) {
			hasIncludedExt = true
			break
		}
	}

	// skip files without included extensions
	if !hasIncludedExt {
		return true
	}

	// skip files that match exclude regexes
	for _, excludeRegex := range config.excludeRegexes {
		if strings.Contains(filename, excludeRegex) {
			return true
		}
	}

	return false
}
