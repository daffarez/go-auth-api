package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daffarez/go-auth-api/internal/httpx"
	"github.com/daffarez/go-auth-api/internal/mocks"
)

func TestLoginHandler(t *testing.T) {
	mockDB := mocks.NewMockDB()

	input := httpx.LoginRequest{
		Email:    "test@example.com",
		Password: "Password1",
	}
	body, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler := Login(mockDB)
	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var apiResp httpx.APIResponse[httpx.LoginResponse]
	json.NewDecoder(resp.Body).Decode(&apiResp)

	if apiResp.Data.Token == "" {
		t.Fatal("expected token, got empty")
	}
}
