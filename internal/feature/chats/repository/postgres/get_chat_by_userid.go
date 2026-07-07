package chats_repository_postgres

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

func (r *ChatRepository) GetChatsByUserID(ctx context.Context, userID int64) ([]domain.ChatListItem, error) {
	query := `
		SELECT c.id, c.created_at, u.id, u.username
		FROM messenger.chats c
		JOIN messenger.chats_participants cp ON cp.chat_id = c.id
		JOIN messenger.users u ON u.id = cp.user_id
		WHERE c.id IN (
			SELECT chat_id FROM messenger.chats_participants WHERE user_id = $1
		)
		AND u.id != $1
		ORDER BY c.created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("get chats by user id: %w", err)
	}
	defer rows.Close()

	var chats []domain.ChatListItem
	for rows.Next() {
		var item domain.ChatListItem
		if err := rows.Scan(&item.ID, &item.CreatedAt, &item.CompanionID, &item.CompanionName); err != nil {
			return nil, fmt.Errorf("scan chat list item: %w", err)
		}
		chats = append(chats, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}

	return chats, nil
}
