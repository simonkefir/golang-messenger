package chats_repository_postgres

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

func (r *ChatRepository) CreateChat(ctx context.Context, userID int64, participantID int64) (domain.Chat, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	var chat domain.Chat
	err = tx.QueryRowContext(ctx,
		`INSERT INTO messenger.chats (created_at) VALUES (NOW()) RETURNING id, created_at`,
	).Scan(&chat.ID, &chat.CreatedAt)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("insert chat: %w", err)
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO messenger.chats_participants (chat_id, user_id) VALUES ($1, $2), ($1, $3)`,
		chat.ID, userID, participantID,
	)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("insert participants: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return domain.Chat{}, fmt.Errorf("commit tx: %w", err)
	}

	return chat, nil
}
