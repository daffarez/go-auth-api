package mocks

import (
	"errors"

	"github.com/daffarez/go-auth-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type MockDB struct {
	users map[string]string
}

func NewMockDB() *MockDB {
	return &MockDB{
		users: make(map[string]string),
	}
}

func (m *MockDB) InsertUser(id, email, passwordHash string) (string, error) {
	if email == "exists@email.com" {
		return "", errors.New("duplicate key value violates unique constraint")
	}

	return id, nil
}

func (m *MockDB) GetUserByEmail(email string) (*database.User, error) {
	if email == "test@example.com" {
		hash, _ := bcrypt.GenerateFromPassword([]byte("Password1"), bcrypt.DefaultCost)

		return &database.User{
			ID:           "123",
			Email:        email,
			PasswordHash: string(hash),
		}, nil
	}

	return nil, errors.New("user not found")
}

func (m *MockDB) GetUserById(id string) (*database.User, error) {
	return &database.User{
		ID:    id,
		Email: "exists@email.com",
	}, nil
}
