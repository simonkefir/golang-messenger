package chats_repository_postgres

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

func (r *ChatRepository) GetChatParticipant(ctx context.Context, chatID int64, excludeUserID int64) (domain.ChatParticipant, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeOut())
	defer cancel()

	query := `
		SELECT u.id, u.username
		FROM messenger.chats_participants cp
		JOIN messenger.users u ON u.id = cp.user_id
		WHERE cp.chat_id = $1 AND u.id != $2
	`

	var participant domain.ChatParticipant
	row := r.pool.QueryRow(
		ctx,
		query,
		chatID,
		excludeUserID,
	)

	if err := row.Scan(&participant.UserID, &participant.DisplayName); err != nil {
		return domain.ChatParticipant{}, fmt.Errorf("get chat participant: %w", err)
	}

	return participant, nil
}
