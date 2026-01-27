package handler

import (
	"encoding/json"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/security"
	"github.com/daffarez/go-auth-api/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body httpx.LoginRequest true "Login payload"
// @Success 200 {object} httpx.LoginSuccessResponse
// @Failure 401 {object} httpx.ErrorResponse
// @Router /auth/login [post]
func Login(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req httpx.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}

		var storedHash string
		var userID string
		err = db.QueryRow(r.Context(),
			"SELECT id, password_hash FROM users WHERE email=$1",
			req.Email,
		).Scan(&userID, &storedHash)
		if err != nil {
			httpx.Error(w, http.StatusUnauthorized, "invalid credentials")
			return
		}

		err = utils.ComparePassword(storedHash, req.Password)
		if err != nil {
			httpx.Error(w, http.StatusUnauthorized, "invalid credentials")
			return
		}

		token, err := security.GenerateToken(userID)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "failed to generate token")
			return
		}

		httpx.OK(w, httpx.LoginResponse{
			Token: token,
		})
	}
}
