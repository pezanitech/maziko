// Hompage handler

package index

import (
	"fmt"
	"net/http"

	"github.com/pezanitech/maziko/libs/core/router"
	inertia "github.com/romsar/gonertia"
)

func Route() {
	router.GET(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
		router.RenderPage(i, w, r, inertia.Props{
			"headline":    "A full-stack framework for modern web development",
			"description": "Build apps with a Go backend and a frontend library of your choice",
			"features": []map[string]string{
				{
					"title":       "Full-Stack Go & Frontend",
					"description": "Develop with Go on the backend and your choice of modern frontend technology",
					"icon":        "rocket",
				},
				{
					"title":       "SPA Experience",
					"description": "Enjoy SPA-like navigation without API boilerplate using Inertia.js",
					"icon":        "sparkles",
				},
				{
					"title":       "File-Based Routing",
					"description": "Automatic route generation based on directory structure",
					"icon":        "folder",
				},
				{
					"title":       "Server-Side Rendering",
					"description": "Built-in SSR support for improved SEO and initial load performance",
					"icon":        "server",
				},
				{
					"title":       "Hot Module Replacement",
					"description": "Fast development feedback with HMR for both Go and frontend code",
					"icon":        "bolt",
				},
				{
					"title":       "Modern Tooling",
					"description": "Includes CLI tools, Tailwind CSS, Vite, and Nx monorepo support",
					"icon":        "wrench",
				},
			},
			"codeExample": fmt.Sprintf(
				"%s\n%s\n\n%s\n%s\n%s\n%s\n%s",
				"# Install the Maziko CLI globally:",
				"go install github.com/pezanitech/maziko/libs/maziko@latest",
				"# Create a new maziko app",
				"maziko create react -n my-app",
				"cd my-app",
				"pnpm i",
				"pnpm dev",
			),
		})
	})
}
