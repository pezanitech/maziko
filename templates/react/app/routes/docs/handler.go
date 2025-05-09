package docs

import (
	"fmt"
	"net/http"

	"github.com/pezanitech/maziko/libs/core/router"
	inertia "github.com/romsar/gonertia"
)

func Route() {
	router.GET(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
		router.RenderPage(i, w, r, inertia.Props{
			"line1": "Documentation",
			"line2": "Coming Soon",
		})
	})

	router.POST(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
		fmt.Println("POST request received at /docs")
	})
}
