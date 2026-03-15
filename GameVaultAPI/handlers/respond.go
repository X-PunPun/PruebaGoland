package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func RespondError(w http.ResponseWriter, status int, code, message string) {
	RespondJSON(w, status, ErrorResponse{
		Code:  code,
		Error: message,
	})
}
