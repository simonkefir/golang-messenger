package users_repository_postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"
)

func (r *UserRepository) PatchUser(ctx context.Context, userID int64, username, displayName, email, password *string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	setClauses := []string{}
	args := []interface{}{}
	argIdx := 1

	if username != nil {
		setClauses = append(setClauses, fmt.Sprintf("username = $%d", argIdx))
		args = append(args, *username)
		argIdx++
	}
	if displayName != nil {
		setClauses = append(setClauses, fmt.Sprintf("display_name = $%d", argIdx))
		args = append(args, *displayName)
		argIdx++
	}
	if email != nil {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argIdx))
		args = append(args, *email)
		argIdx++
	}
	if password != nil {
		setClauses = append(setClauses, fmt.Sprintf("password_hash = $%d", argIdx))
		args = append(args, *password)
		argIdx++
	}

	if len(setClauses) == 0 {
		return domain.User{}, fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf(`
        UPDATE messenger.users
        SET %s
        WHERE id = $%d
        RETURNING id, username, display_name, email, password_hash, created_at
    `, strings.Join(setClauses, ", "), argIdx)
	args = append(args, userID)

	var user domain.User
	row := r.pool.QueryRow(
		ctx,
		query,
		args...,
	)
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.DisplayName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d' concurrently accessed: %w",
				userID,
				core_errors.ErrAlreadyExists,
			)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return user, nil
}
