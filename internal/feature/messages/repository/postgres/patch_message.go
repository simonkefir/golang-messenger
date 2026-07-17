package messages_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"
)

func (r *MsgRepository) PatchMessage(ctx context.Context, messageID int64, senderID int64, content string) (domain.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	query := `
		UPDATE messenger.messages
		SET message = $1
		WHERE id = $2 AND sender_id = $3
		RETURNING id, chat_id, sender_id, message, sent_at
	`

	var msg domain.Message

	row := r.pool.QueryRow(
		ctx,
		query,
		content,
		messageID,
		senderID,
	)
	if err := row.Scan(&msg.ID, &msg.ChatID, &msg.SenderID, &msg.Content, &msg.SentAt); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Message{}, fmt.Errorf(
				"message with id='%d' concurrently accessed: %w",
				messageID,
				core_errors.ErrAlreadyExists,
			)
		}

		return domain.Message{}, fmt.Errorf("scan error: %w", err)
	}

	return msg, nil
}
