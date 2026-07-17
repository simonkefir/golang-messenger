package users_repository_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *UserRepository) DeleteUser(ctx context.Context, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	query := `
	DELETE FROM messenger.users
	WHERE ID=$1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%d' : %w", userID, core_errors.ErrNotFound)
	}

	return nil
}
