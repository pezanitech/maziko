package news

import (
	"fmt"
	"net/http"

	"github.com/pezanitech/maziko/core/utils"
	inertia "github.com/romsar/gonertia"
)

func GET(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
	err := i.Render(w, r, "news", inertia.Props{
		"line1": "A full-stack framework",
		"line2": "built with Inertia.js and Go! 💚",
	})

	if err != nil {
		utils.HandleServerErr(w, err)
	}
}

func POST(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request received")
}
