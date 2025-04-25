package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	inertia "github.com/romsar/gonertia"
)

func Execute() {
	i := initInertia()

	mux := http.NewServeMux()

	mux.Handle("/", i.Middleware(rootHandler(i)))
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./build"))))

	http.ListenAndServe(":8000", mux)
}

func initInertia() *inertia.Inertia {
	viteHotFile := ".tmp/hot"
	rootViewFile := "frontend/views/root.html"

	// check if laravel-vite-plugin is running in dev mode
	// it puts a "hot" file in the .tmp directory
	_, err := os.Stat(viteHotFile)

	if err != nil {
		// retry after 3 seconds, 3 attempts
		for i := range 3 {
			_, err = os.Stat(viteHotFile)
			if err == nil {
				break
			}

			log.Printf(
				"waiting for laravel-vite-plugin to start... attempt %d\n", i+1,
			)

			time.Sleep(3 * time.Second)
		}
	}

	if err == nil {
		i, err := inertia.NewFromFile(
			rootViewFile,
			inertia.WithSSR(),
		)
		if err != nil {
			log.Fatal(err)
		}
		i.ShareTemplateFunc("vite", func(entry string) (string, error) {
			content, err := os.ReadFile(viteHotFile)
			if err != nil {
				return "", err
			}
			url := strings.TrimSpace(string(content))
			if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
				url = url[strings.Index(url, ":")+1:]
			} else {
				url = "//localhost:8080"
			}
			if entry != "" && !strings.HasPrefix(entry, "/") {
				entry = "/" + entry
			}
			return url + entry, nil
		})

		i.ShareTemplateData("hmr", true)
		return i
	}

	manifestPath := "./build/.vite/manifest.json"

	i, err := inertia.NewFromFile(
		rootViewFile,
		inertia.WithVersionFromFile(manifestPath),
		inertia.WithSSR(),
	)
	if err != nil {
		log.Fatal(err)
	}

	i.ShareTemplateFunc("vite", vite(manifestPath, "/build/"))

	return i
}

func vite(manifestPath, buildDir string) func(path string) (string, error) {
	f, err := os.Open(manifestPath)
	if err != nil {
		log.Fatalf("cannot open provided vite manifest file: %s", err)
	}
	defer f.Close()

	viteAssets := make(map[string]*struct {
		File   string `json:"file"`
		Source string `json:"src"`
	})
	err = json.NewDecoder(f).Decode(&viteAssets)
	// print content of viteAssets
	for k, v := range viteAssets {
		log.Printf("%s: %s\n", k, v.File)
	}

	if err != nil {
		log.Fatalf("cannot unmarshal vite manifest file to json: %s", err)
	}

	return func(p string) (string, error) {
		if val, ok := viteAssets[p]; ok {
			return path.Join("/", buildDir, val.File), nil
		}
		return "", fmt.Errorf("asset %q not found", p)
	}
}

func rootHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			err := i.Render(w, r, "Home/Index", inertia.Props{
				"text": "Inertia.js with React and Go! 💚",
			})
			if err != nil {
				handleServerErr(w, err)
			}
		default:
			// if file exists in public serve it
			_, err := os.Stat(path.Join("public", r.URL.Path))
			if err == nil {
				http.ServeFile(w, r, path.Join("public", r.URL.Path))
				return
			}

			// otherwise return 404
			http.NotFound(w, r)
		}
	})
}

func handleServerErr(w http.ResponseWriter, err error) {
	log.Printf("http error: %s\n", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("server error"))
}
