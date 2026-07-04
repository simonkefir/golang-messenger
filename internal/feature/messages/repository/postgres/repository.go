package messages_repository_postgres

import (
	"database/sql"
	"errors"
)

type MsgRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *MsgRepository {
	return &MsgRepository{
		db: db,
	}
}

func isPgUniqueViolation(err error) bool {
	var pgErr interface{ SQLState() string }
	if errors.As(err, &pgErr) {
		return pgErr.SQLState() == "23505"
	}
	return false
}
