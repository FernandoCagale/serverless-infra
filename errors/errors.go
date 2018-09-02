package errors

import (
	"net/http"
)

//ResponseError struct
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//AddInternalServerError InternalServerError
func AddInternalServerError(message string) ResponseError {
	return ResponseError{http.StatusInternalServerError, message}
}

//AddNotFoundError NotFound
func AddNotFoundError(message string) ResponseError {
	return ResponseError{http.StatusNotFound, message}
}

//AddUnauthorizedError Unauthorized
func AddUnauthorizedError(message string) ResponseError {
	return ResponseError{http.StatusUnauthorized, message}
}

//AddBadRequestError BadRequest
func AddBadRequestError(message string) ResponseError {
	return ResponseError{http.StatusBadRequest, message}
}

//AddMethodNotAllowedError MethodNotAllowed
func AddMethodNotAllowedError(message string) ResponseError {
	return ResponseError{http.StatusMethodNotAllowed, message}
}
