package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func write(w http.ResponseWriter, status int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func Success(w http.ResponseWriter, status int, msg string, data interface{}) {
	write(w, status, APIResponse{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, status int, msg string, err error) {
	write(w, status, APIResponse{
		Success: false,
		Message: msg,
		Errors:  err.Error(),
	})
}

func ValidationError(w http.ResponseWriter, errs []string) {
	write(w, http.StatusBadRequest, APIResponse{
		Success: false,
		Message: "Validation error",
		Errors:  errs,
	})
}