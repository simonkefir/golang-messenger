package messages_repository_postgres

import (
	"context"
	"fmt"

	core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"
)

func (r *MsgRepository) DeleteMessage(ctx context.Context, chatID int64, messageID int64) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	query := `
		DELETE FROM messenger.messages
		WHERE id = $1 AND chat_id = $2
	`

	cmdTag, err := r.pool.Exec(ctx, query, messageID, chatID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return core_postgres_pool.ErrNoRows
	}

	return nil
}
