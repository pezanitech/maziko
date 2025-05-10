// Hompage handler

package index

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/pezanitech/maziko/libs/core/router"
	inertia "github.com/romsar/gonertia"
)

type CodeStep struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Type        string `json:"type"`
	Filename    string `json:"filename,omitempty"`
}

type CodeExample struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Code        string     `json:"code"`
	Type        string     `json:"type"`
	Filename    string     `json:"filename,omitempty"`
	Steps       []CodeStep `json:"steps,omitempty"`
}

type Feature struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type HomePageData struct {
	Description      string        `json:"description"`
	FeaturesTitle    string        `json:"featuresTitle"`
	FeaturesSubtitle string        `json:"featuresSubtitle"`
	Features         []Feature     `json:"features"`
	CodeExamples     []CodeExample `json:"codeExamples"`
}

func Route() {
	router.GET(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
		// Get the current file's path
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("Failed to get current file path")
		}

		// Construct path to data.json
		dataPath := filepath.Join(filepath.Dir(filename), "data.json")

		// Read and parse the JSON file
		fileData, err := ioutil.ReadFile(dataPath)
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

		// Pass the data as props directly to the frontend
		router.RenderPage(i, w, r, inertia.Props{
			"description":      pageData.Description,
			"featuresTitle":    pageData.FeaturesTitle,
			"featuresSubtitle": pageData.FeaturesSubtitle,
			"features":         pageData.Features,
			"codeExamples":     pageData.CodeExamples,
		})
	})
}
