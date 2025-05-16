package reload

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/pezanitech/maziko/libs/core/errors"
	"github.com/pezanitech/maziko/libs/core/logger"
)

// createFileWatcher creates a new file watcher
func createFileWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		errors.HandleFatalError(
			"Failed to create file watcher", err,
		)
	}
	return watcher
}

// addDirectoriesToWatch walks the project and adds directories to the watcher
func addDirectoriesToWatch(watcher *fsnotify.Watcher, config reloadConfig) {
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
		},
	)

	if err != nil {
		errors.HandleFatalError(
			"Failed to add directories to watcher", err,
		)
	}
}

// shouldSkipDir checks if a directory should be excluded from watching
func shouldSkipDir(path string, excludeDirs []string) bool {
	for _, excludeDir := range excludeDirs {
		if strings.Contains(path, excludeDir) {
			return true
		}
	}
	return false
}

// watchForChanges monitors file changes and triggers rebuilds when needed
func watchForChanges(watcher *fsnotify.Watcher, config reloadConfig, buildAndRun func()) {
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

// shouldSkipFile checks if file should be ignored based on extension and patterns
func shouldSkipFile(filename string, config reloadConfig) bool {
	hasIncludedExt := false

	// check if file has included extension
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
