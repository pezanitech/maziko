package index

import (
	"net/http"

	"github.com/pezanitech/maziko/core/router"
	"github.com/pezanitech/maziko/core/utils"
	inertia "github.com/romsar/gonertia"
)

func Route() {
	router.GET(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "index", inertia.Props{
			"line1": "A full-stack framework",
			"line2": "built with Inertia.js and Go! 💚",
		})

		if err != nil {
			utils.HandleServerErr(w, err)
		}
	})
}
