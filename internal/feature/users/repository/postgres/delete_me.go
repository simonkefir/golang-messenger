package users_repository_postgres

import (
	"context"
	"fmt"
)

func (r *UserRepository) DeleteMe(ctx context.Context, userID int64) error {
	query := `
	DELETE FROM messenger.users
	WHERE ID=$1;
	`

	if err := r.db.QueryRowContext(ctx, query, userID); err != nil {
		return fmt.Errorf("exec query: %v", err)
	}

	return nil
}
