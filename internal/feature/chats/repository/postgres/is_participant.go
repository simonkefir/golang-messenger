package chats_repository_postgres

import (
	"context"
	"fmt"
)

func (r *ChatRepository) IsParticipant(ctx context.Context, chatID, userID int64) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx,
		`SELECT EXISTS(SELECT 1 FROM messenger.chats_participants WHERE chat_id = $1 AND user_id = $2)`,
		chatID, userID,
	).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("check participant: %w", err)
	}

	return exists, nil
}
