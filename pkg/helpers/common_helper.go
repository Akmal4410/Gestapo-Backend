package helpers

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Status     bool       `json:"status"`
	StatusCode int        `json:"status_code,omitempty"`
	Message    string     `json:"message,omitempty"`
	Data       any        `json:"data,omitempty"`
	ErrorInfo  *errorInfo `json:"error_info,omitempty"`
}
type errorInfo struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&jsonResponse{
		Status:     true,
		StatusCode: status,
		Data:       data,
	})
}

func ErrorJson(statusCode int, errMsg string) error {
	return nil
}
