package handler

import (
	"encoding/json"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/security"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			httpx.JSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid request body",
			})
			return
		}

		var storedHash string
		var userID string
		err = db.QueryRow(r.Context(),
			"SELECT id, password_hash FROM users WHERE email=$1",
			req.Email,
		).Scan(&userID, &storedHash)
		if err != nil {
			httpx.JSON(w, http.StatusUnauthorized, map[string]string{
				"error": "invalid credentials",
			})
			return
		}

		err = security.ComparePassword(storedHash, req.Password)
		if err != nil {
			httpx.JSON(w, http.StatusUnauthorized, map[string]string{
				"error": "invalid credentials",
			})
			return
		}

		token, err := security.GenerateToken(userID)
		if err != nil {
			httpx.JSON(w, http.StatusInternalServerError, map[string]string{
				"error": "failed to generate token",
			})
			return
		}

		httpx.JSON(w, http.StatusOK, map[string]string{
			"token": token,
		})
	}
}
