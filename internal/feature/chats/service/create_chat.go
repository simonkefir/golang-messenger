package chats_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *ChatsService) CreateChat(ctx context.Context, userID int64, chat_participant int64) (domain.Chat, error) {
	if userID == chat_participant {
		return domain.Chat{}, core_errors.ErrInvalidInput
	}

	existingChat, err := s.chatsRepository.FindPrivateChat(ctx, userID, chat_participant)
	if err == nil {
		return existingChat, nil
	}
	if !errors.Is(err, core_errors.ErrNotFound) {
		return domain.Chat{}, fmt.Errorf("find private chat: %w", err)
	}

	chat, err := s.chatsRepository.CreateChat(ctx, userID, chat_participant)
	if err != nil {
		return domain.Chat{}, fmt.Errorf("create chat: %w", err)
	}

	return chat, nil
}
