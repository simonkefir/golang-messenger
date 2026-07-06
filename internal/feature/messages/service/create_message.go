package messages_service

import (
	"context"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *MessagesService) CreateMessage(ctx context.Context, senderID int64, chatID int64, content string) (domain.Message, error) {
	ok, err := s.chatsChecker.IsParticipant(ctx, chatID, senderID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("check participant: %w", err)
	}
	if !ok {
		return domain.Message{}, core_errors.ErrForbidden
	}

	msg, err := s.messagesRepository.CreateMessage(ctx, senderID, chatID, content)
	if err != nil {
		return domain.Message{}, fmt.Errorf("create message: %w", err)
	}

	return msg, nil
}
