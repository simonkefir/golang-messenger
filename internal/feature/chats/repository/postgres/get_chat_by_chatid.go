package chats_repository_postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *ChatRepository) GetChatByChatID(ctx context.Context, chatID int64) (domain.Chat, error) {
	query := `
		SELECT id, created_at
		FROM messenger.chats
		WHERE id = $1
	`

	var chat domain.Chat
	err := r.db.QueryRowContext(ctx, query, chatID).Scan(
		&chat.ID,
		&chat.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Chat{}, core_errors.ErrNotFound
		}
		return domain.Chat{}, err
	}

	return chat, nil
}
