package database

import (
	"context"
	"time"
)

type UserDB interface {
	InsertUser(userID, email, hashedPassword string) (string, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
}

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

func (p *PostgresDB) InsertUser(userID, email, hash string) (string, error) {
	_, err := p.Pool.Exec(context.Background(),
		"INSERT INTO users (id, email, password_hash, created_at) VALUES ($1,$2,$3,$4)",
		userID, email, hash, time.Now(),
	)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (p *PostgresDB) GetUserByEmail(email string) (*User, error) {
	row := p.Pool.QueryRow(
		context.Background(),
		"SELECT id, email, password_hash FROM users WHERE email=$1",
		email,
	)

	var u User
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (p *PostgresDB) GetUserById(id string) (*User, error) {
	row := p.Pool.QueryRow(
		context.Background(),
		"SELECT email FROM users WHERE id=$1",
		id,
	)

	var u User
	err := row.Scan(u.Email)
	if err != nil {
		return nil, err
	}

	return &u, nil

}
