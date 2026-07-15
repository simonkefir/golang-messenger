package messages_repository_postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *MsgRepository) PatchMessage(ctx context.Context, messageID int64, senderID int64, content string) (domain.Message, error) {
	query := `
		UPDATE messenger.messages
		SET message = $1
		WHERE id = $2 AND sender_id = $3
		RETURNING id, chat_id, sender_id, message, sent_at
	`

	var msg domain.Message
	err := r.db.QueryRowContext(ctx, query, content, messageID, senderID).Scan(
		&msg.ID, &msg.ChatID, &msg.SenderID, &msg.Content, &msg.SentAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Message{}, core_errors.ErrForbidden
		}
		return domain.Message{}, fmt.Errorf("update message: %w", err)
	}

	return msg, nil
}
