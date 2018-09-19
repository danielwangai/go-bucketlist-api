package controllers

import (
	"encoding/json"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJson(w, statusCode, map[string]string{"error": message})
}
