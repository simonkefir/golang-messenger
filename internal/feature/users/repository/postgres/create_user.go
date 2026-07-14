package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *UserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	query := `
		INSERT INTO messenger.users (username, display_name, email, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.DisplayName,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		if isPgUniqueViolation(err) {
			return domain.User{}, core_errors.ErrAlreadyExists
		}
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
