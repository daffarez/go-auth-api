package handler

import (
	"net/http"

	"github.com/daffarez/go-auth-api/internal/database"
	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/middleware"
)

// @Summary Get current user
// @Tags User
// @Security BearerAuth
// @Success 200 {object} httpx.UserResponse
// @Router /user [get]
func User(db database.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := middleware.GetUserID(r.Context())

		user, err := db.GetUserById(userID)
		if err != nil || user == nil {
			httpx.Error(w, http.StatusInternalServerError, err.Error(), "failed to fetch user")
			return
		}

		httpx.OK(w, httpx.UserResponse{
			UserID: userID,
			Email:  user.Email,
		})
	}
}
