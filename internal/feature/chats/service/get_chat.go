package chats_service

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *ChatsService) GetChat(ctx context.Context, userID, chatID int64) (domain.ChatWithParticipant, error) {
	isParticipant, err := s.chatsRepository.IsParticipant(ctx, chatID, userID)
	if err != nil {
		return domain.ChatWithParticipant{}, fmt.Errorf("check participant: %w", err)
	}
	if !isParticipant {
		return domain.ChatWithParticipant{}, core_errors.ErrForbidden
	}

	chat, err := s.chatsRepository.GetChatByID(ctx, chatID, userID)
	if err != nil {
		return domain.ChatWithParticipant{}, fmt.Errorf("get chat: %w", err)
	}

	return domain.ChatWithParticipant{
		Chat:        chat.Chat,
		Participant: chat.Participant,
	}, nil
}
