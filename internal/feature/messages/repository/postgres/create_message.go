package messages_repository_postgres

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

func (r *MsgRepository) CreateMessage(ctx context.Context, senderID int64, chatID int64, content string) (domain.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	query := `
		INSERT INTO messenger.messages (chat_id, sender_id, message)
		VALUES ($1, $2, $3)
		RETURNING id, chat_id, sender_id, message, sent_at
	`

	var msg domain.Message

	row := r.pool.QueryRow(ctx, query, chatID, senderID, content)
	if err := row.Scan(
		&msg.ID,
		&msg.ChatID,
		&msg.SenderID,
		&msg.Content,
		&msg.SentAt,
	); err != nil {
		return domain.Message{}, fmt.Errorf("create message: %w", err)
	}

	return msg, nil
}
