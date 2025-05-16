package reload

import (
	"time"

	"github.com/pezanitech/maziko/libs/core/logger"
)

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
