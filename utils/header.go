package utils

import "net/http"

func SetContentTypeJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
