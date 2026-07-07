package chats_service

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

func (s *ChatsService) GetChats(ctx context.Context, userID int64) ([]domain.ChatListItem, error) {
	chats, err := s.chatsRepository.GetChatsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get chats: %w", err)
	}
	return chats, nil
}
