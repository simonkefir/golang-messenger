package chats_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"
)

func (r *ChatRepository) FindPrivateChat(ctx context.Context, user1, user2 int64) (domain.ChatWithParticipant, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	query := `
	SELECT c.id, c.created_at
	FROM messenger.chats c
	JOIN messenger.chats_participants cp ON cp.chat_id = c.id
	WHERE cp.user_id IN ($1, $2)
	GROUP BY c.id
	HAVING COUNT(DISTINCT cp.user_id) = 2
	`

	var chat domain.ChatWithParticipant
	rows := r.pool.QueryRow(ctx, query, user1, user2)
	if err := rows.Scan(&chat.ID, &chat.CreatedAt); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.ChatWithParticipant{}, core_errors.ErrNotFound
		}
		return domain.ChatWithParticipant{}, fmt.Errorf("find private chat: %w", err)
	}

	return chat, nil
}
