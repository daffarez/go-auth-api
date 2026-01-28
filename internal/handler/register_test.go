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

func TestRegisterHandler(t *testing.T) {
	mockDB := mocks.NewMockDB()

	tests := []struct {
		name       string
		email      string
		password   string
		wantStatus int
		wantMsg    string
	}{
		{
			name:       "success",
			email:      "newuser@example.com",
			password:   "Password2",
			wantStatus: http.StatusCreated,
			wantMsg:    "",
		},
		{
			name:       "email exists",
			email:      "exists@email.com",
			password:   "Password3",
			wantStatus: http.StatusConflict,
			wantMsg:    "email already registered",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := httpx.RegisterRequest{
				Email:    tt.email,
				Password: tt.password,
			}
			body, _ := json.Marshal(input)
			req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handler := Register(mockDB)
			handler.ServeHTTP(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Fatalf("expected status %d, got %d", tt.wantStatus, resp.StatusCode)
			}

			switch tt.wantStatus {
			case http.StatusCreated:
				var apiResp httpx.APIResponse[httpx.UserResponse]
				json.NewDecoder(resp.Body).Decode(&apiResp)
				if apiResp.Data.UserID == "" {
					t.Fatal("expected user ID, got empty")
				}
			case http.StatusConflict:
				var apiResp httpx.APIResponse[httpx.ErrorResponse]
				json.NewDecoder(resp.Body).Decode(&apiResp)
				if apiResp.Message != tt.wantMsg {
					t.Fatalf("expected message %q, got %q", tt.wantMsg, apiResp.Message)
				}
			default:
				t.Fatalf("unhandled status %d", resp.StatusCode)
			}
		})
	}

}
