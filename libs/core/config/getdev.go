package config

import "time"

// ---- DEV MODE CONFIG GETTERS ----

// Returns the root directory for file watching
func GetDevRootDir() string {
	return AppConfig.Dev.RootDir
}

// Returns list of regex patterns to exclude from file watching
func GetDevExcludeRegexes() []string {
	return AppConfig.Dev.ExcludeRegexes
}

// Returns list of directories to exclude from file watching
func GetDevExcludeDirs() []string {
	return AppConfig.Dev.ExcludeDirs
}

// Returns list of file extensions to include in file watching
func GetDevIncludeExts() []string {
	return AppConfig.Dev.IncludeExts
}

// Returns the build delay in milliseconds
func GetDevBuildDelay() time.Duration {
	return time.Duration(AppConfig.Dev.BuildDelay) * time.Millisecond
}
