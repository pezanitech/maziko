package config

// ---- PATH CONFIG GETTERS ----

// Returns path to generation directory
func GetGenDir() string {
	return AppConfig.Paths.Gen
}

// Returns path to routes directory
func GetRoutesDir() string {
	return AppConfig.Paths.Routes
}

// Returns path to public directory
func GetPublicDir() string {
	return AppConfig.Paths.Public
}
