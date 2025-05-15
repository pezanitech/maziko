package router

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/pezanitech/maziko/libs/core/errors"
	"github.com/pezanitech/maziko/libs/core/logger"
)

// Renders a page with the provided props,
// can optionally be specified with component parameter
func RenderPage(i Inertia, w http.ResponseWriter, r *http.Request, props Props, component ...string) {
	componentName := ""

	if len(component) > 0 && component[0] != "" {
		componentName = component[0] // use provided component name
	} else {
		path := r.URL.Path
		if path == "" {
			path = "/"
		}

		// get name from routeComponents map using URL path
		if storedComponent, exists := routeComponents[path]; exists {
			componentName = storedComponent

			logger.Log.Debug(
				"Using component from route map",
				"path", path,
				"component", componentName,
			)
		} else {
			msg := "Component not found in routeComponents map"

			logger.Log.Error(msg, "path", path)
			panic(msg)
		}
	}

	logger.Log.Info(
		"Rendering page",
		"component", componentName,
	)

	// render the page
	err := RenderInertiaPage(i, w, r, componentName, props)
	if err != nil {
		errors.HandleServerErr(w, err)
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

	logger.Log.Debug(
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

// Finds the caller's package name by searching up the stack
// Returns the package name and a boolean indicating success
func findCallerPackage() (string, bool) {
	// try different stack depths to find the right caller
	// starting with 3 (skip findCallerPackage and it's parent)
	for skipFrames := 3; skipFrames <= 5; skipFrames++ {
		pkgName, ok := extractPackageName(skipFrames)
		if !ok {
			continue
		}

		// Skip router and gen package
		if pkgName == "router" || pkgName == "gen" {
			continue
		}

		logger.Log.Debug(
			"Caller package found",
			"package", pkgName,
			"skipFrames", skipFrames,
		)

		return pkgName, true
	}

	logger.Log.Warn(
		"Could not determine caller package",
		"frames_checked", "3-5",
	)

	return "", false
}

// Determines the component name from the caller's package
func determineComponentName() string {
	pkgName, ok := findCallerPackage()
	if !ok {
		logger.Log.Warn(
			"Could not determine component name, using 'index'",
		)
		return "index"
	}

	logger.Log.Debug(
		"Component determined from package",
		"package", pkgName,
	)

	return pkgName
}

// Determines the routes path from the caller's package
func determineRoutePath() string {
	pkgName, ok := findCallerPackage()
	if !ok {
		logger.Log.Warn(
			"Could not determine route path, using root path",
		)
		return "/"
	}

	logger.Log.Debug(
		"Route path determined from package",
		"package", pkgName,
	)

	// handle index package as root path
	if pkgName == "index" {
		return "/"
	}

	return "/" + pkgName
}
