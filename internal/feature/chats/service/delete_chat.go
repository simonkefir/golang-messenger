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

	participants, err := s.chatsRepository.GetChatParticipants(ctx, chatID)
	if err != nil {
		return fmt.Errorf("get chat participants: %w", err)
	}

	if err := s.chatsRepository.DeleteChat(ctx, chatID); err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return core_errors.ErrNotFound
		}
		return fmt.Errorf("delete chat: %w", err)
	}

	for _, p := range participants {
		if p.UserID != userID {
			s.publisher.Publish(p.UserID, core_websocket.Event{
				Type: core_websocket.EventChatDeleted,
				Data: map[string]int64{
					"chat_id": chatID,
				},
			})
		}
	}

	return nil
}
