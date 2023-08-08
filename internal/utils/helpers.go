package utils

import "net/http"

func ReturnJsonResponse(w http.ResponseWriter, httpCode int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(msg)
}
