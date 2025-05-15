package index

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pezanitech/maziko/libs/core/router"
	inertia "github.com/romsar/gonertia"
)

type Feature struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type HomePageData struct {
	Description      string    `json:"description"`
	FeaturesTitle    string    `json:"featuresTitle"`
	FeaturesSubtitle string    `json:"featuresSubtitle"`
	Features         []Feature `json:"features"`
}

func Route() {
	router.GET(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("Failed to get current file path")
		}

		dataPath := filepath.Join(filepath.Dir(filename), "data.json")

		fileData, err := os.ReadFile(dataPath)
		if err != nil {
			log.Printf("Error reading data file: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		var pageData HomePageData
		if err := json.Unmarshal(fileData, &pageData); err != nil {
			log.Printf("Error parsing JSON data: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// pass the data as props directly to the frontend
		router.RenderPage(i, w, r, inertia.Props{
			"description":      pageData.Description,
			"featuresTitle":    pageData.FeaturesTitle,
			"featuresSubtitle": pageData.FeaturesSubtitle,
			"features":         pageData.Features,
		})
	})
}
