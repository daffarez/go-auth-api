package httpx

import (
	"net/http"
)

func OK[T any](w http.ResponseWriter, data T) {
	JSON(w, http.StatusOK, APIResponse[T]{
		Data: data,
	})
}

func Created[T any](w http.ResponseWriter, data T) {
	JSON(w, http.StatusCreated, APIResponse[T]{
		Data: data,
	})
}

func Error(w http.ResponseWriter, status int, errorMessage string, message string) {
	JSON(w, status, APIResponse[ErrorResponse]{
		Data: ErrorResponse{
			Status:       status,
			ErrorMessage: errorMessage,
		},
		Message: message,
	})
}
