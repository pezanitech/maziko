package router

import (
	"strings"
	"text/template"
)

// Template functions for use in the route generation template
var templateFuncs = template.FuncMap{
	"contains":      strings.Contains,
	"before":        before,
	"after":         after,
	"extractParent": extractParent,
	"extractName":   extractName,
}

// before returns the substring before the first occurrence of sep
func before(s, sep string) string {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i]
	}
	return s
}

// after returns the substring after the first occurrence of sep
func after(s, sep string) string {
	if i := strings.Index(s, sep); i >= 0 {
		return s[i+len(sep):]
	}
	return s
}

// extractParent extracts the parent directory name from a route path
// e.g., "github.com/.../routes/articles/_slug" returns "articles"
func extractParent(path string) string {
	// Find "routes/" in the path
	routesIdx := strings.Index(path, "routes/")
	if routesIdx == -1 {
		return ""
	}

	// Get everything after "routes/"
	routePath := path[routesIdx+7:]

	// Split by "/"
	parts := strings.Split(routePath, "/")

	// The parent is the first part if there are at least 2 parts
	if len(parts) >= 2 {
		return parts[0]
	}

	return ""
}

// extractName extracts the route name from a route path
// e.g., "github.com/.../routes/articles/_slug" returns "slug" (without underscore)
func extractName(path string) string {
	// Split the path by "/"
	parts := strings.Split(path, "/")

	// Get the last part
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]

		 // Remove trailing quote if present
		if strings.HasSuffix(lastPart, "\"") {
			lastPart = lastPart[:len(lastPart)-1]
		}

		// If it starts with underscore, remove it for display purposes (the _ will be added in the template)
		if strings.HasPrefix(lastPart, "_") {
			return lastPart[1:] // Return without the underscore since we add it in the template
		}

		return lastPart
	}

	return ""
}

// GetTemplateFuncs returns the template functions for route generation
func GetTemplateFuncs() template.FuncMap {
	return templateFuncs
}
