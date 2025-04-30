package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pezanitech/maziko/core/utils"
)

var (
	genDir        = "./core/gen"
	buildDir      = "./build"
	routesDir     = "./app/routes"
	publicDir     = "./app/public"
	buildPrefix   = "/build/"
	packagePrefix = "github.com/pezanitech/maziko/"
)

func registerRoutes(path string, dirEntry os.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if dirEntry.IsDir() {
		// don't register routes directory
		if path == routesDir {
			return nil
		}

		fmt.Println("Path:", path)

		// Write path to routesgen.go file
		routesgenPath := filepath.Join(genDir, "routesgen.go")

		// Open file in append mode
		f, err := os.OpenFile(routesgenPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			utils.Logger.Error(
				"Error opening routesgen.go file",
				"error", err,
			)
			return err
		}
		defer f.Close()

		// Write path to file
		if _, err := f.WriteString(fmt.Sprintf("import \"%s%s\"\n", packagePrefix, path)); err != nil {
			utils.Logger.Error(
				"Error writing to routesgen.go file",
				"error", err,
			)
			return err
		}
	}

	return nil
}

func GenerateRoutes() {
	// Initialize logger before use
	utils.InitLogger()

	// Initialize a fresh routesgen.go file before walking the routes
	routesgenPath := filepath.Join(genDir, "routesgen.go")

	utils.Logger.Info("Generating routes definitions...")

	// Create directory if it doesn't exist
	if err := os.MkdirAll(genDir, 0755); err != nil {
		utils.Logger.Error(
			"Error creating gen directory",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to create gen directory: %v", err))
	}

	// Create a new empty file (overwriting any existing one)
	f, err := os.Create(routesgenPath)
	if err != nil {
		utils.Logger.Error(
			"Error creating routesgen.go file",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to create routesgen.go: %v", err))
	}

	// Write package declaration
	if _, err := f.WriteString(`
package gen

import "net/http"
import "os"
import "path"
import "strings"
import inertia "github.com/romsar/gonertia"

`); err != nil {
		utils.Logger.Error(
			"Error writing package declaration to routesgen.go",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to write to routesgen.go: %v", err))
	}
	f.Close()

	if err := filepath.WalkDir(routesDir, registerRoutes); err != nil {
		utils.Logger.Error(
			"Error registering routes",
			"error", err,
		)

		panic(fmt.Sprintf("Failed to register routes: %v", err))
	}

	// Open file in append mode
	f, err = os.OpenFile(routesgenPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		utils.Logger.Error(
			"Error opening routesgen.go file",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to open routesgen.go: %v", err))
	}
	defer f.Close()

	// Write routing definitions
	if _, err := f.WriteString(fmt.Sprintf(`
func DefineRoutes(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch true {
		case r.URL.Path == "/" && r.Method == http.MethodGet:
			index.GET(i, w, r)
		case r.URL.Path == "/" && r.Method == http.MethodPost:
			index.POST(i, w, r)
		case r.URL.Path == "/news" && r.Method == http.MethodGet:
			news.GET(i, w, r)
		case strings.HasPrefix(r.URL.Path, "%s"):
			handleRequest(w, r, buildDirHandler)
		default:
			handleRequest(w, r, staticFileHandler)
		}
	})
}
	
func handleRequest(w http.ResponseWriter, r *http.Request, f func() http.Handler) {
	f().ServeHTTP(w, r)
}

func staticFileHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourcePath := path.Join("%s", r.URL.Path)

		if _, err := os.Stat(resourcePath); err == nil {
			// file exists, serve it
			http.ServeFile(w, r, resourcePath)
			return
		}

		// no file exists, respond with a 404
		http.NotFound(w, r)
	})
}

func buildDirHandler() http.Handler {
	return http.StripPrefix(
		"%s",
		http.FileServer(http.Dir("%s")),
	)
}
`, buildPrefix, publicDir, buildPrefix, buildDir)); err != nil {
		utils.Logger.Error(
			"Error writing package declaration to routesgen.go",
			"error", err,
		)
		panic(fmt.Sprintf("Failed to write to routesgen.go: %v", err))
	}

	utils.Logger.Info("Routes generation completed successfully")
}
