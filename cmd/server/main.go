package main

import (
	"log"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/health", handler.Health)
	router.Post("/register", handler.Register)

	log.Println("Server running on :8088")
	log.Fatal(http.ListenAndServe(":8088", router))
}
