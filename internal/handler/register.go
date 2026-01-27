package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := httpx.DecodeJSON(r, &input); err != nil {
			httpx.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}

		if err := utils.ValidateRegisterInput(input.Email, input.Password); err != nil {
			httpx.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		hash, err := utils.HashPassword(input.Password)
		if err != nil {
			httpx.JSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to hash password"})
			return
		}

		userID := uuid.New().String()
		_, err = db.Exec(context.Background(),
			"INSERT INTO users (id, email, password_hash, created_at) VALUES ($1,$2,$3,$4)",
			userID, input.Email, hash, time.Now(),
		)
		if err != nil {
			httpx.JSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to insert user"})
			return
		}

		httpx.JSON(w, http.StatusCreated, map[string]string{
			"user_id": userID,
			"email":   input.Email,
		})
	}
}
