package router

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/pezanitech/maziko/core/logger"
	"github.com/pezanitech/maziko/core/utils"
	inertia "github.com/romsar/gonertia"
)

// Renders a page with the provided props
// a different component from the default can be specified
func RenderPage(i *inertia.Inertia, w http.ResponseWriter, r *http.Request, props inertia.Props, component ...string) {
	componentName := ""

	// use provided component or detect from caller's package
	if len(component) > 0 && component[0] != "" {
		componentName = component[0]
	} else {
		// First try to get the component name from the route map based on the URL path
		path := r.URL.Path
		if path == "" {
			path = "/"
		}

		if storedComponent, exists := routeComponents[path]; exists {
			componentName = storedComponent
			logger.Logger.Debug(
				"Using component from route map",
				"path", path,
				"component", componentName,
			)
		} else {
			// Fall back to determining component from caller's package
			componentName = determineComponentName()
		}
	}

	logger.Logger.Info(
		"Rendering page",
		"component", componentName,
	)

	// render the page
	err := i.Render(w, r, componentName, props)
	if err != nil {
		utils.HandleServerErr(w, err)
	}
}

// Extracts the package name from the caller's stack
// skipFrames defines how many stack frames to skip (2 means caller of the caller)
// returns the package name and whether the extraction was successful
func extractPackageName(skipFrames int) (string, bool) {
	pc, _, _, ok := runtime.Caller(skipFrames)
	if !ok {
		return "", false
	}

	callerFunction := runtime.FuncForPC(pc).Name()

	logger.Logger.Debug(
		"Extracting package from caller",
		"function", callerFunction,
		"skipFrames", skipFrames,
	)

	// extract the package path
	parts := strings.Split(callerFunction, "/")

	if len(parts) >= 2 {
		lastPart := parts[len(parts)-1]
		pkgFunc := strings.Split(lastPart, ".")

		if len(pkgFunc) >= 1 {
			return pkgFunc[0], true
		}
	}

	return "", false
}

// Extracts the component name from the caller's package
func determineComponentName() string {
	// Try different stack depths to find the right caller
	// Starting with 2 (direct caller of RenderPage)
	for skipFrames := 2; skipFrames <= 4; skipFrames++ {
		pkgName, ok := extractPackageName(skipFrames)
		if !ok {
			continue
		}

		// Skip our own router package
		if pkgName == "router" || pkgName == "gen" {
			continue
		}

		logger.Logger.Debug(
			"Component determined from package",
			"package", pkgName,
			"skipFrames", skipFrames,
		)

		return pkgName
	}

	logger.Logger.Warn(
		"Could not determine component name, using 'index'",
	)
	return "index"
}

// Extracts the routes path from the caller's package
func determineRoutePath() string {
	// Try different stack depths to find the right caller
	// Starting with 2 (direct caller of GET)
	for skipFrames := 2; skipFrames <= 4; skipFrames++ {
		pkgName, ok := extractPackageName(skipFrames)
		if !ok {
			continue
		}

		// Skip our own router package and gen package
		if pkgName == "router" || pkgName == "gen" {
			continue
		}

		logger.Logger.Debug(
			"Route path determined from package",
			"package", pkgName,
			"skipFrames", skipFrames,
		)

		// handle index package as root path
		if pkgName == "index" {
			return "/"
		}

		return "/" + pkgName
	}

	logger.Logger.Warn(
		"Could not determine route path, using root path",
	)
	return "/"
}
