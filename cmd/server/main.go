package main

import (
	"log"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/database"
	"github.com/daffarez/go-auth-api/internal/handler"
	"github.com/daffarez/go-auth-api/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	db := database.NewPostgres(
		"postgres://daffarez:password123@localhost:5432/authdb?sslmode=disable",
	)
	defer db.Close()

	router := chi.NewRouter()

	mountAuthRoutes(router, db)
	mountUserRoutes(router)

	log.Println("Server running on :8088")
	log.Fatal(http.ListenAndServe(":8088", router))
}

func mountAuthRoutes(r chi.Router, db *pgxpool.Pool) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register(db))
		r.Post("/login", handler.Login(db))
	})
}

func mountUserRoutes(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Get("/", handler.User())
	})
}
