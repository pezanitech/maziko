// Documentation page handler

package docs

import (
	"net/http"

	"github.com/pezanitech/maziko/libs/core/router"
	inertia "github.com/romsar/gonertia"
)

func Route() {
	router.GET(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
		router.RenderPage(i, w, r, inertia.Props{
			"title":       "Documentation",
			"description": "Learn how to build full-stack applications with Maziko",
			"sections": []map[string]interface{}{
				{
					"title":   "Getting Started",
					"id":      "getting-started",
					"content": "Get up and running with Maziko quickly",
					"subsections": []map[string]string{
						{
							"title":   "Installation",
							"id":      "installation",
							"content": "Install the Maziko CLI globally:\n\n```bash\ngo install github.com/pezanitech/maziko/libs/maziko@latest```",
						},
						{
							"title":   "Creating a Project",
							"id":      "creating-a-project",
							"content": "Use the CLI to create your first project:\n\n```bash\nmaziko create react -n my-app\ncd my-app\npnpm i\npnpm dev```",
						},
					},
				},
				{
					"title":   "Key Concepts",
					"id":      "key-concepts",
					"content": "Understanding Maziko's architecture",
					"subsections": []map[string]string{
						{
							"title":   "File-Based Routing",
							"id":      "file-based-routing",
							"content": "Maziko uses a file-based routing system where each route is a directory containing at least a handler.go and page.tsx file.",
						},
						{
							"title":   "Inertia Integration",
							"id":      "inertia-integration",
							"content": "Inertia.js bridges your Go backend with your frontend framework, providing a seamless SPA-like experience without the API boilerplate.",
						},
					},
				},
				{
					"title":   "Frontend Development",
					"id":      "frontend-development",
					"content": "Building your UI with Maziko",
					"subsections": []map[string]string{
						{
							"title":   "Component Structure",
							"id":      "component-structure",
							"content": "Create React components that receive data from your Go handlers via the Inertia protocol.",
						},
						{
							"title":   "Accessing Properties",
							"id":      "accessing-properties",
							"content": "Use the `usePage` hook to access data passed from your Go handler:\n\n```tsx\nconst { props } = usePage()\nconsole.log(props.title) // Access data sent from handler.go```",
						},
					},
				},
				{
					"title":   "Backend Development",
					"id":      "backend-development",
					"content": "Working with Go handlers and routes",
					"subsections": []map[string]string{
						{
							"title":   "Creating Routes",
							"id":      "creating-routes",
							"content": "Define routes in your handler.go file using the router package:\n\n```go\nfunc Route() {\n  router.GET(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {\n    router.RenderPage(i, w, r, inertia.Props{\n      \"key\": \"value\",\n    })\n  })\n}```",
						},
					},
				},
				{
					"title":   "Deployment",
					"id":      "deployment",
					"content": "Taking your Maziko app to production",
					"subsections": []map[string]string{
						{
							"title":   "Building for Production",
							"id":      "building-for-production",
							"content": "Build your app for production:\n\n```bash\npnpm build```",
						},
						{
							"title":   "Running in Production",
							"id":      "running-in-production",
							"content": "Start your production server:\n\n```bash\npnpm start```",
						},
					},
				},
			},
		})
	})
}
