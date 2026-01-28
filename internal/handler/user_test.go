package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daffarez/go-auth-api/internal/mocks"
)

func TestUserHandler(t *testing.T) {
	mockDB := mocks.NewMockDB()

	req := httptest.NewRequest("GET", "/user", nil)
	req.Header.Set("Authorization", "Bearer mocked-jwt-token")

	w := httptest.NewRecorder()
	handler := User(mockDB)

	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}
