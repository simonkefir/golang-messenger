package chats_repository_postgres

import (
	"context"
	"fmt"

	core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"
)

func (r *ChatRepository) DeleteChat(ctx context.Context, chatID int64) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	query := `DELETE FROM messenger.chats WHERE id = $1`

	result, err := r.pool.Exec(ctx, query, chatID)
	if err != nil {
		return fmt.Errorf("delete chat: %w", err)
	}

	if result.RowsAffected() == 0 {
		return core_postgres_pool.ErrNoRows
	}

	return nil
}
