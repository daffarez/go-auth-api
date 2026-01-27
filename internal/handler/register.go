package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Register godoc
// @Summary Register new user
// @Description Create new user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body httpx.RegisterRequest true "Register payload"
// @Success 201 {object} httpx.RegisterSuccessResponse
// @Failure 400 {object} httpx.ErrorResponse
// @Failure 500 {object} httpx.ErrorResponse
// @Router /auth/register [post]
func Register(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input httpx.RegisterRequest
		if err := httpx.DecodeJSON(r, &input); err != nil {
			LogError("invalid JSON", err)
			httpx.Error(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if err := utils.ValidateRegisterInput(input.Email, input.Password); err != nil {
			LogError(fmt.Sprintf("invalid email or password %s", input.Email), err)
			httpx.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		hash, err := utils.HashPassword(input.Password)
		if err != nil {
			LogError(fmt.Sprintf("failed to hash password for email %s", input.Email), err)
			httpx.Error(w, http.StatusInternalServerError, "failed to hash password")
			return
		}

		userID := uuid.New().String()
		_, err = db.Exec(context.Background(),
			"INSERT INTO users (id, email, password_hash, created_at) VALUES ($1,$2,$3,$4)",
			userID, input.Email, hash, time.Now(),
		)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				LogError(fmt.Sprintf("email %s already exist", input.Email), err)
				httpx.Error(w, http.StatusConflict, "email already registered")
				return
			}

			LogError(fmt.Sprintf("failed to insert user %s", input.Email), err)
			httpx.Error(w, http.StatusInternalServerError, "failed to create user")
			return
		}

		httpx.Created(w, httpx.UserResponse{
			UserID: userID,
			Email:  input.Email,
		})
	}
}
