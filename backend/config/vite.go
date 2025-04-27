package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/pezanitech/maziko/backend/utils"
)

func vite(manifestPath, buildDir string) func(path string) (string, error) {
	manifest, err := os.Open(manifestPath)
	if err != nil {
		utils.Logger.Error("cannot open provided vite manifest file", "error", err)
		os.Exit(1)
	}
	defer manifest.Close()

	viteAssets := make(map[string]*struct {
		File   string `json:"file"`
		Source string `json:"src"`
	})

	if err = json.NewDecoder(manifest).Decode(&viteAssets); err != nil {
		utils.Logger.Error("cannot unmarshal vite manifest file to json", "error", err)
		os.Exit(1)
	}

	// print content of viteAssets
	for k, v := range viteAssets {
		utils.Logger.Info("vite asset", "path", k, "file", v.File)
	}

	return func(p string) (string, error) {
		if val, ok := viteAssets[p]; ok {
			return path.Join("/", buildDir, val.File), nil
		}
		return "", fmt.Errorf("asset %q not found", p)
	}
}
