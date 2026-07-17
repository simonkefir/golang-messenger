package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

func (r *UserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	query := `
		INSERT INTO messenger.users (username, display_name, email, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	row := r.pool.QueryRow(
		ctx,
		query,
		user.Username,
		user.DisplayName,
		user.Email,
		user.PasswordHash,
	)
	err := row.Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return user, nil
}
