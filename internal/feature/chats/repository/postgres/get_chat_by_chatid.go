package chats_repository_postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (r *ChatRepository) GetChatByID(ctx context.Context, chatID int64, userID int64) (domain.ChatWithParticipant, error) {
	query := `
		SELECT c.id, c.created_at, u.id, u.display_name
		FROM messenger.chats c
		JOIN messenger.chats_participants cp ON cp.chat_id = c.id
		JOIN messenger.users u ON u.id = cp.user_id
		WHERE c.id = $1 AND u.id != $2
	`

	var chat domain.ChatWithParticipant
	err := r.db.QueryRowContext(ctx, query, chatID, userID).Scan(
		&chat.ID,
		&chat.CreatedAt,
		&chat.Participant.UserID,
		&chat.Participant.DisplayName,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ChatWithParticipant{}, core_errors.ErrNotFound
		}
		return domain.ChatWithParticipant{}, fmt.Errorf("get chat by id: %w", err)
	}

	return chat, nil
}
