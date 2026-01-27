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

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, ErrorResponse{
		Status:  status,
		Message: message,
	})
}
