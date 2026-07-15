package chats_service

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_websocket "github.com/simonkefir/golang-messenger/internal/core/websocket"
)

func (s *ChatsService) DeleteChat(ctx context.Context, userID int64, chatID int64) error {
	isParticipant, err := s.chatsRepository.IsParticipant(ctx, chatID, userID)
	if err != nil {
		return fmt.Errorf("check participant: %w", err)
	}

	if !isParticipant {
		return core_errors.ErrForbidden
	}

	participant, err := s.chatsRepository.GetChatParticipant(ctx, chatID, userID)
	if err != nil {
		return fmt.Errorf("get chat participant: %w", err)
	}

	if err := s.chatsRepository.DeleteChat(ctx, chatID); err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return core_errors.ErrNotFound
		}
		return fmt.Errorf("delete chat: %w", err)
	}

	s.publisher.Publish(participant.UserID, core_websocket.Event{
		Type: core_websocket.EventChatDeleted,
		Data: core_websocket.ChatDeletedPayload{
			ChatID: chatID,
		},
	})

	return nil
}
