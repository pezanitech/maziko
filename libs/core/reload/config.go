package reload

import (
	"os"
	"path/filepath"
	"time"

	"github.com/pezanitech/maziko/libs/core/config"
	"github.com/pezanitech/maziko/libs/core/errors"
)

// Config parameters for live reloading
type reloadConfig struct {
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

// createReloadConfig creates an instance of reloadConfig with global config values
func createReloadConfig() reloadConfig {
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

	return reloadConfig{
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
