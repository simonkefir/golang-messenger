package domain

import "time"

type User struct {
	ID           int64
	Username     string
	DisplayName  string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
