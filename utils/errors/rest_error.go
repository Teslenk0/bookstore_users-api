package errors

import (
	"net/http"
)

//RestError - Struct for handling errors and having like a template
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

//NewBadRequestError - Function to create bad request errors like factory
func NewBadRequestError(message string) *RestError {
	var restErr = RestError{
		Message: `"` + message + `"`,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
	return &restErr
}
