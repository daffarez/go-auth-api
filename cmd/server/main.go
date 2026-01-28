// @title Go Auth API
// @version 1.0
// @description Simple authentication API using Go, JWT, and PostgreSQL
// @termsOfService https://example.com/terms/

// @contact.name Daffarez
// @contact.email daffarez@email.com

// @host localhost:8088
// @BasePath /

package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/daffarez/go-auth-api/docs"
	"github.com/daffarez/go-auth-api/internal/database"
	"github.com/daffarez/go-auth-api/internal/handler"
	"github.com/daffarez/go-auth-api/internal/middleware"
	"github.com/daffarez/go-auth-api/internal/security"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	security.InitJWTSecret()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	db := database.Postgres(databaseURL)
	defer db.Close()

	router := chi.NewRouter()

	router.Use(middleware.LoggerJSON)

	mountAuthRoutes(router, db)
	mountUserRoutes(router, db)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8088/swagger/doc.json"),
	))

	log.Printf("Server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func mountAuthRoutes(r chi.Router, db database.UserDB) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register(db))
		r.Post("/login", handler.Login(db))
	})
}

func mountUserRoutes(r chi.Router, db database.UserDB) {
	r.Route("/user", func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Get("/", handler.User(db))
	})
}
