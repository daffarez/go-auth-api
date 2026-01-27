package handler

import (
	"encoding/json"
	"net/http"

	"github.com/daffarez/go-auth-api/internal/httpx"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
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

	httpx.JSON(w, http.StatusCreated, map[string]string{
		"message": "user registered (dummy)",
	})
}
