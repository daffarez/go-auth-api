package handler

import (
	"net/http"

	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/middleware"
)

func User() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := middleware.GetUserID(r.Context())

		httpx.JSON(w, http.StatusOK, map[string]string{
			"user_id": userID,
		})
	}
}
