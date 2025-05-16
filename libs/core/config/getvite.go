package config

import "time"

// ---- VITE CONFIG GETTERS ----

// Returns path to Vite manifest file
func GetViteManifestFile() string {
	return AppConfig.Vite.ManifestFile
}

// Returns path to Vite hot file
func GetHotFile() string {
	return AppConfig.Vite.HotFile
}

// Returns maximum number of attempts to detect Vit
func MaxViteDetectionAttempts() int {
	return AppConfig.Vite.DetectionAttempts
}

// Returns interval between Vite detection attempts
func ViteDetectionInterval() time.Duration {
	return time.Duration(AppConfig.Vite.DetectionInterval) * time.Millisecond
}
