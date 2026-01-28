package handler

import (
	"encoding/json"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/database"
	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/security"
	"github.com/daffarez/go-auth-api/internal/utils"
)

// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body httpx.LoginRequest true "Login payload"
// @Success 200 {object} httpx.LoginSuccessResponse
// @Failure 401 {object} httpx.ErrorResponse
// @Router /auth/login [post]
func Login(db database.UserDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req httpx.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, err.Error(), "invalid request body")
			return
		}

		user, err := db.GetUserByEmail(req.Email)
		if err != nil || user == nil {
			httpx.Error(w, http.StatusUnauthorized, err.Error(), "invalid credentials")
			return
		}

		err = utils.ComparePassword(user.PasswordHash, req.Password)
		if err != nil {
			httpx.Error(w, http.StatusUnauthorized, err.Error(), "invalid credentials")
			return
		}

		token, err := security.GenerateToken(user.ID)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, err.Error(), "failed to generate token")
			return
		}

		httpx.OK(w, httpx.LoginResponse{
			Token: token,
		})
	}
}
