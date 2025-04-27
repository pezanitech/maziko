package handlers

import (
	"net/http"
)

func BuildHandler() http.Handler {
	return http.StripPrefix("/build/", http.FileServer(http.Dir("./build")))
}
