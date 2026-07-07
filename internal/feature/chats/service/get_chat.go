package chats_service

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *ChatsService) GetChat(ctx context.Context, userID, chatID int64) (domain.ChatWithParticipants, error) {
	isParticipant, err := s.chatsRepository.IsParticipant(ctx, chatID, userID)
	if err != nil {
		return domain.ChatWithParticipants{}, fmt.Errorf("check participant: %w", err)
	}
	if !isParticipant {
		return domain.ChatWithParticipants{}, core_errors.ErrForbidden
	}

	chat, err := s.chatsRepository.GetChatByChatID(ctx, chatID)
	if err != nil {
		return domain.ChatWithParticipants{}, fmt.Errorf("get chat: %w", err)
	}

	participants, err := s.chatsRepository.GetChatParticipants(ctx, chatID)
	if err != nil {
		return domain.ChatWithParticipants{}, fmt.Errorf("get participants: %w", err)
	}

	return domain.ChatWithParticipants{
		Chat:         chat,
		Participants: participants,
	}, nil
}
