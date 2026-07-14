package users_repository_postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *UserRepository) GetUserByID(ctx context.Context, userID int64) (domain.User, error) {
	query := `
	SELECT id, username, display_name, email, password_hash, created_at 
	FROM messenger.users
	WHERE id = $1
	`

	var user domain.User

	err := r.db.QueryRowContext(ctx, query, userID).Scan(
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
		return domain.User{}, fmt.Errorf("query user by id: %w", err)
	}

	return user, nil
}
