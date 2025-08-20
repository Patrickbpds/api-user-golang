package rest_errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Causes []Causes `json:"causes ,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestError) Errors() string {
	return r.Message
}

func NewRestError(message string, status int, error string, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
		Causes:  causes,
	}
}

func NewBadRequestError(message string, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not Found",
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusForbidden,
		Error:   "Forbidden",
	}
}
