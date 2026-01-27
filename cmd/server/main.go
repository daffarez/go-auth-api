package main

import (
	"log"
	"net/http"
	"os"

	"github.com/daffarez/go-auth-api/internal/database"
	"github.com/daffarez/go-auth-api/internal/handler"
	"github.com/daffarez/go-auth-api/internal/middleware"
	"github.com/daffarez/go-auth-api/internal/security"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	security.InitJWTSecret()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	db := database.NewPostgres(databaseURL)
	defer db.Close()

	router := chi.NewRouter()

	router.Use(middleware.LoggerJSON)

	mountAuthRoutes(router, db)
	mountUserRoutes(router, db)

	log.Printf("Server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func mountAuthRoutes(r chi.Router, db *pgxpool.Pool) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register(db))
		r.Post("/login", handler.Login(db))
	})
}

func mountUserRoutes(r chi.Router, db *pgxpool.Pool) {
	r.Route("/user", func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Get("/", handler.User(db))
	})
}
