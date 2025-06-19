package utils

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, status int, mensaje string) {
	JSON(w, status, map[string]string{"error": mensaje})
}

func ParseJSON(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}
