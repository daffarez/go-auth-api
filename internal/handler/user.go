package handler

import (
	"context"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

// @Summary Get current user
// @Tags User
// @Security BearerAuth
// @Success 200 {object} httpx.UserResponse
// @Router /user [get]
func User(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := middleware.GetUserID(r.Context())

		var email string
		err := db.QueryRow(context.Background(), "SELECT email FROM users WHERE id=$1", userID).Scan(&email)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "failed to fetch user")
			return
		}

		httpx.OK(w, httpx.UserResponse{
			UserID: userID,
			Email:  email,
		})
	}
}
