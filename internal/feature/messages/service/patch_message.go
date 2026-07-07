package messages_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *MessagesService) PatchMessage(ctx context.Context, senderID int64, chatID int64, messageID int64, content string) (domain.Message, error) {
	patched, err := s.messagesRepository.PatchMessage(ctx, chatID, messageID, content)
	if err != nil {
		if errors.Is(err, core_errors.ErrAlreadyExists) {
			return domain.Message{}, core_errors.ErrAlreadyExists
		}
		return domain.Message{}, fmt.Errorf("patch message: %w", err)
	}

	return patched, nil
}
