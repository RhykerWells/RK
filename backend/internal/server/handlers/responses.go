package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type SuccessResponse struct {
	Status string `json:"status"`
}

func RespondError(w http.ResponseWriter, status int, err error) {
	RespondErrorMessage(w, status, err.Error())
}

func RespondErrorMessage(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Status: "error",
		Error:  message,
	})
}

func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}