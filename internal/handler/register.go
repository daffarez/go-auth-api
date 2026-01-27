package handler

import (
	"encoding/json"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/security"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			httpx.JSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid request body",
			})
			return
		}

		if req.Email == "" || req.Password == "" {
			httpx.JSON(w, http.StatusBadRequest, map[string]string{
				"error": "email and password are required",
			})
			return
		}

		hashedPassword, err := security.HashPassword(req.Password)
		if err != nil {
			httpx.JSON(w, http.StatusInternalServerError, map[string]string{
				"error": "failed to hash password",
			})
			return
		}

		id := uuid.New()

		_, err = db.Exec(
			r.Context(),
			`INSERT INTO users (id, email, password_hash)
			 VALUES ($1, $2, $3)`,
			id,
			req.Email,
			hashedPassword,
		)
		if err != nil {
			httpx.JSON(w, http.StatusInternalServerError, map[string]string{
				"error": "failed to insert user",
			})
			return
		}

		httpx.JSON(w, http.StatusCreated, map[string]string{
			"id": id.String(),
		})
	}
}
