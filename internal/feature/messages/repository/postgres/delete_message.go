package messages_repository_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *MsgRepository) DeleteMessage(ctx context.Context, chatID int64, messageID int64) error {
	query := `
		DELETE FROM messenger.messages
		WHERE id = $1 AND chat_id = $2
	`

	result, err := r.db.ExecContext(ctx, query, messageID, chatID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
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
