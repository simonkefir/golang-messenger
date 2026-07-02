package chats_repository_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *ChatRepository) DeleteChat(ctx context.Context, chatID int64) error {
	query := `DELETE FROM messenger.chats WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, chatID)
	if err != nil {
		return fmt.Errorf("delete chat: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
