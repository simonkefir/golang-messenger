package chats_repository_postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *ChatRepository) FindPrivateChat(ctx context.Context, user1, user2 int64) (domain.Chat, error) {
	query := `
	SELECT c.id, c.created_at
	FROM messenger.chats c
	JOIN messenger.chats_participants cp ON cp.chat_id = c.id
	WHERE cp.user_id IN ($1, $2)
	GROUP BY c.id
	HAVING COUNT(DISTINCT cp.user_id) = 2
	`

	var chat domain.Chat
	err := r.db.QueryRowContext(ctx, query, user1, user2).Scan(&chat.ID, &chat.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Chat{}, core_errors.ErrNotFound
		}
		return domain.Chat{}, fmt.Errorf("find private chat: %w", err)
	}

	return chat, nil
}
