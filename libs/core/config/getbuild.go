package config

// ---- BUILD CONFIG GETTERS ----

// Returns build prefix
func GetBuildPrefix() string {
	return AppConfig.Build.Prefix
}

// Returns path to build directory
func GetBuildDir() string {
	return AppConfig.Build.Dir
}

// Returns path to temp directory
func GetTempDir() string {
	return AppConfig.Build.TempDir
}

// Returns path to the SSR build directory
func GetSSRDir() string {
	return AppConfig.Build.SSRDir
}
