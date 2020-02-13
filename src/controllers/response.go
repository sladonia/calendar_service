package controllers

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func RespondError(w http.ResponseWriter, err ApiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.GetStatusCode())
	json.NewEncoder(w).Encode(err)
}
