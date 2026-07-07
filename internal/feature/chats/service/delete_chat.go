package chats_service

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *ChatsService) DeleteChat(ctx context.Context, userID int64, chatID int64) error {
	isParticipant, err := s.chatsRepository.IsParticipant(ctx, chatID, userID)
	if err != nil {
		return fmt.Errorf("check participant: %w", err)
	}

	if !isParticipant {
		return core_errors.ErrForbidden
	}

	if err := s.chatsRepository.DeleteChat(ctx, chatID); err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return core_errors.ErrNotFound
		}
		return fmt.Errorf("delete chat: %w", err)
	}

	return nil
}
