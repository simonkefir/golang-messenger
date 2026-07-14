package users_repository_postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	query := `
	SELECT id, username, display_name, email, password_hash, created_at 
	FROM messenger.users
	WHERE username = $1
	`

	var user domain.User

	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.DisplayName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, core_errors.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("query user by username: %w", err)
	}

	return user, nil
}
