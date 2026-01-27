package handler

import (
	"context"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func User(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := middleware.GetUserID(r.Context())

		var email string
		err := db.QueryRow(context.Background(), "SELECT email FROM users WHERE id=$1", userID).Scan(&email)
		if err != nil {
			httpx.JSON(w, http.StatusInternalServerError, map[string]string{
				"error": "failed to fetch user",
			})
			return
		}

		httpx.JSON(w, http.StatusOK, map[string]string{
			"user_id": userID,
			"email":   email,
		})
	}
}
