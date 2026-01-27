package main

import (
	"log"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/database"
	"github.com/daffarez/go-auth-api/internal/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	db := database.NewPostgres(
		"postgres://daffarez:password123@localhost:5432/authdb?sslmode=disable",
	)
	defer db.Close()

	router := chi.NewRouter()

	router.Get("/health", handler.Health)
	router.Post("/register", handler.Register(db))
	router.Post("/login", handler.Login(db))

	log.Println("Server running on :8088")
	log.Fatal(http.ListenAndServe(":8088", router))
}
