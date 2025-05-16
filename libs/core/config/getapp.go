package config

// ---- APP CONFIG GETTERS ----

// Returns app name
func GetAppName() string {
	return AppConfig.App.Name
}

// Returns app URL
func GetAppURL() string {
	return AppConfig.App.URL
}

// Returns app port
func GetAppPort() int {
	return AppConfig.App.Port
}

// ---- PACKAGE CONFIG GETTERS ----

// Returns package prefix (for generated imports)
func GetPackagePrefix() string {
	return AppConfig.Package.Prefix
}
