package reload

import (
	"os"
	"os/exec"

	"github.com/pezanitech/maziko/libs/core/logger"
)

// buildApp builds the application with the specified binary output path
func buildApp(binPath string) bool {
	logger.Log.Info(
		"Building application...",
	)

	buildCmd := "go build -o " + binPath + " ."
	buildCmdExec := exec.Command("sh", "-c", buildCmd)
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
