package httpx

type APIResponse[T any] struct {
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Status  int
	Message string `json:"error" example:"invalid email or password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserID string `json:"user_id" example:"271ca6d0-b901-41e0-920b-0e759627b09b"`
	Email  string `json:"email" example:"user@email.com"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token" example:"271ca6d0-b901-41e0-920b-0e759627b09b"`
}

type LoginSuccessResponse struct {
	Data LoginResponse `json:"data"`
}

type RegisterSuccessResponse struct {
	Data UserResponse `json:"data"`
}
