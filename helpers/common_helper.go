package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
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

var validate = validator.New()

func ValidateBody(r *http.Request, data any) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return err
	}
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
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

func ErrorJson(w http.ResponseWriter, status int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(jsonResponse{
		Status: false,
		ErrorInfo: &errorInfo{
			StatusCode: status,
			Message:    err.Error(),
		},
	})
}
