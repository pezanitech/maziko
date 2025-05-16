package reload

import (
	"os/exec"

	"github.com/pezanitech/maziko/libs/core/logger"
)

// Run starts the application with live reloading enabled.
// It sets up file watching, builds the application, and restarts it when changes are detected.
func Run() {
	logger.Log.Info(
		"Starting application with live reloading...",
	)

	// configure reload settings
	setupConfig := createReloadConfig()

	watcher := createFileWatcher()
	defer watcher.Close()

	// initialize process manager function
	var cmd *exec.Cmd
	processManager := initProcessManager(setupConfig.binPath, &cmd)

	// add directories to watch
	addDirectoriesToWatch(watcher, setupConfig)

	// watch for file changes
	watchForChanges(watcher, setupConfig, processManager)
}
