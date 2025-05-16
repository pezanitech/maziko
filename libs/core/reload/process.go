package reload

import (
	"os"
	"os/exec"

	"github.com/pezanitech/maziko/libs/core/logger"
)

// initProcessManager creates a function that manages, building, and running the app
func initProcessManager(binPath string, cmdRef **exec.Cmd) func() {
	processManager := func() {
		// kill any running process
		stopProcess(cmdRef)

		// build and run the application
		if buildApp(binPath) {
			startProcess(binPath, cmdRef)
		}
	}

	// initial build and run
	processManager()

	return processManager
}

// stopProcess attempts to kill the currently running process
func stopProcess(cmdRef **exec.Cmd) {
	cmd := *cmdRef

	// only attempt to kill if process is valid
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

// startProcess runs the built application
func startProcess(binPath string, cmdRef **exec.Cmd) {
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
