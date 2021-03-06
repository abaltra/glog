package responsehandler

import (
	"encoding/json"
	"net/http"
)


// Error model info
// @Description Error message with HTTP status code and error message
type Error struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"Message"`
}

func EncodeJSONError(w http.ResponseWriter, err error, status int) {
	response := &Error{
		StatusCode: status,
		Message:    err.Error(),
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode((response))
}
