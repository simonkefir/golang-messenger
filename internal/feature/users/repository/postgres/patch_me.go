package users_repository_postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *UserRepository) PatchMe(ctx context.Context, userID int64, username *string, password *string) (domain.User, error) {
	query := `
		UPDATE messenger.users
		SET
			username = COALESCE($1, username),
			password_hash = COALESCE($2, password_hash)
		WHERE id = $3
		RETURNING id, username, email, password_hash, created_at
	`

	var user domain.User

	err := r.db.QueryRowContext(ctx, query, username, password, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, core_errors.ErrNotFound
		}
	}

	return user, nil
}
