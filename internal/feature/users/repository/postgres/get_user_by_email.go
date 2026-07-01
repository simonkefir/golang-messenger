package users_repository_postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	query := `
	SELECT id, username, email, password_hash, created_at 
	FROM messenger.users
	WHERE email = $1
	`

	var user domain.User

	err := r.db.QueryRowContext(ctx, query, email).Scan(
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
