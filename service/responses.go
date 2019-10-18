package service

import (
	"encoding/json"
	"net/http"
)

// JSONResponse formats and sends back data as JSON
func JSONResponse(w http.ResponseWriter, data interface{}) {
	payload, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// WriteBuildQueued responds with a JSON payload
// indicating the build entered the queue
func WriteBuildQueued(w http.ResponseWriter, details string) {
	w.WriteHeader(http.StatusOK)

	JSONResponse(w, map[string]string{
		"status":  "ok",
		"details": details,
	})
}

// WriteError sends an error response in JSON format
func WriteError(w http.ResponseWriter, status int, details string) {
	w.WriteHeader(status)

	JSONResponse(w, map[string]string{
		"status":  "error",
		"details": details,
	})
}
