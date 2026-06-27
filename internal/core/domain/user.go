package domain

import "time"

type User struct {
	ID           int64
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(
	id int64,
	username string,
	email string,
	password_hash string,
	created_at time.Time,
) User {
	return User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: password_hash,
		CreatedAt:    created_at,
	}
}
