package utils

import (
	"net/http"
)

func HandleServerErr(w http.ResponseWriter, err error) {
	Logger.Error("http error", "err", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("server error"))
}
