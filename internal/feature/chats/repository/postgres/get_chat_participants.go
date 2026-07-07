package chats_repository_postgres

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

func (r *ChatRepository) GetChatParticipants(ctx context.Context, chatID int64) ([]domain.ChatParticipant, error) {
	query := `
		SELECT u.id, u.username
		FROM messenger.chats_participants cp
		JOIN messenger.users u ON u.id = cp.user_id
		WHERE cp.chat_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, chatID)
	if err != nil {
		return nil, fmt.Errorf("get chat participants: %w", err)
	}
	defer rows.Close()

	var participants []domain.ChatParticipant
	for rows.Next() {
		var p domain.ChatParticipant
		if err := rows.Scan(&p.UserID, &p.Username); err != nil {
			return nil, fmt.Errorf("scan participant: %w", err)
		}
		participants = append(participants, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}

	return participants, nil
}
