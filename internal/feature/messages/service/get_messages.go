package messages_service

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *MessagesService) GetMessages(ctx context.Context, userID int64, chatID int64) ([]domain.Message, error) {
	ok, err := s.chatsChecker.IsParticipant(ctx, chatID, userID)
	if err != nil {
		return nil, fmt.Errorf("check participant: %w", err)
	}
	if !ok {
		return nil, core_errors.ErrForbidden
	}

	messages, err := s.messagesRepository.GetChatMessages(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("get messages: %w", err)
	}

	return messages, nil
}
