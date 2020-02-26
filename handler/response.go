package handler

import (
	"encoding/json"
	"net/http"
)

// SuccessBody holds data for success response
type SuccessBody struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    interface{} `json:"meta"`
}

// MultipleErrorBody holds data for more than one error response
type MultipleErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   interface{} `json:"meta"`
}

// ErrorBody holds data for error response
type ErrorBody struct {
	Error ErrorInfo   `json:"error"`
	Meta  interface{} `json:"meta"`
}

// MetaInfo holds meta data
type MetaInfo struct {
	HTTPStatus int   `json:"http_status"`
	Limit      int64 `json:"limit,omitempty"`
	Offset     int64 `json:"offset,omitempty"`
	Total      int64 `json:"total,omitempty"`
}

// ErrorInfo holds error detail
type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Field   string `json:"field,omitempty"`
}

func (ei ErrorInfo) String() string {
	return ei.Message
}

// Write is a function to write data in json format
func Write(w http.ResponseWriter, result interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(result)
}

// OK is a function to write an OK response
func OK(w http.ResponseWriter, data interface{}, message string) {
	body := SuccessBody{
		Data:    data,
		Message: message,
		Meta: MetaInfo{
			HTTPStatus: http.StatusOK,
		},
	}

	Write(w, body, http.StatusOK)
}

// OKWithMeta is a function to write an OK response with offset and limit
func OKWithMeta(w http.ResponseWriter, data interface{}, message string, offset, limit int64) {
	body := SuccessBody{
		Data:    data,
		Message: message,
		Meta: MetaInfo{
			HTTPStatus: http.StatusOK,
			Offset:     offset,
			Limit:      limit,
		},
	}

	Write(w, body, http.StatusOK)
}

// Error is a function to write an error response
func Error(w http.ResponseWriter, err error) {
	// TODO: create function to define error code and http status
	body := ErrorBody{
		Error: ErrorInfo{
			Message: err.Error(),
			Code:    500,
		},
		Meta: MetaInfo{
			HTTPStatus: http.StatusInternalServerError,
		},
	}

	Write(w, body, http.StatusInternalServerError)
}
