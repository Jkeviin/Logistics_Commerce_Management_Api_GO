package utils

import (
	"github.com/bootcamp-go/web/response"
	"net/http"
)

// HandleError handles errors based on a mapping of error types to HTTP status codes
func HandleError(w http.ResponseWriter, err error, errorMapping map[error]int, defaultStatus int) {
	if status, exists := errorMapping[err]; exists {
		response.Error(w, status, err.Error())
	} else {
		response.Error(w, defaultStatus, err.Error())
	}
}
