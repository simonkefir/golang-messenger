package chats_repository_postgres

import (
	"context"
	"fmt"
)

func (r *ChatRepository) IsParticipant(ctx context.Context, chatID int64, userID int64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	var exists bool
	row := r.pool.QueryRow(
		ctx,
		`SELECT EXISTS(SELECT 1 FROM messenger.chats_participants WHERE chat_id = $1 AND user_id = $2)`,
		chatID,
		userID,
	)
	

	if err := row.Scan(&exists); err != nil {
		return false, fmt.Errorf("check participant: %w", err)
	}

	return exists, nil
}
