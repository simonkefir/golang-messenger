package chats_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"
)

func (r *ChatRepository) GetChatByID(ctx context.Context, chatID int64, userID int64) (domain.ChatWithParticipant, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	query := `
		SELECT c.id, c.created_at, u.id, u.display_name
		FROM messenger.chats c
		JOIN messenger.chats_participants cp ON cp.chat_id = c.id
		JOIN messenger.users u ON u.id = cp.user_id
		WHERE c.id = $1 AND u.id != $2
	`

	var chat domain.ChatWithParticipant
	row := r.pool.QueryRow(ctx, query, chatID, userID)
	if err := row.Scan(
		&chat.ID,
		&chat.CreatedAt,
		&chat.Participant.UserID,
		&chat.Participant.DisplayName,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.ChatWithParticipant{}, core_errors.ErrNotFound
		}

		return domain.ChatWithParticipant{}, fmt.Errorf("get chat by id: %w", err)
	}

	return chat, nil
}
