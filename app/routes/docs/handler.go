package docs

import (
	"net/http"

	"github.com/pezanitech/maziko/core/router"
	"github.com/pezanitech/maziko/core/utils"
	inertia "github.com/romsar/gonertia"
)

func Route() {
	router.GET(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "docs", inertia.Props{
			"line1": "Documentation",
			"line2": "Coming Soon",
		})

		if err != nil {
			utils.HandleServerErr(w, err)
		}
	})
}
