package messages_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_websocket "github.com/simonkefir/golang-messenger/internal/core/websocket"
)

func (s *MessagesService) PatchMessage(ctx context.Context, senderID int64, chatID int64, messageID int64, content string) (domain.Message, error) {
	ok, err := s.chatsChecker.IsParticipant(ctx, chatID, senderID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("check participant: %w", err)
	}
	if !ok {
		return domain.Message{}, core_errors.ErrForbidden
	}

	patched, err := s.messagesRepository.PatchMessage(ctx, chatID, messageID, content)
	if err != nil {
		if errors.Is(err, core_errors.ErrAlreadyExists) {
			return domain.Message{}, core_errors.ErrAlreadyExists
		}
		return domain.Message{}, fmt.Errorf("patch message: %w", err)
	}

	participants, err := s.chatsChecker.GetChatParticipants(ctx, chatID)
	if err == nil {
		for _, p := range participants {
			s.publisher.Publish(p.UserID, core_websocket.Event{
				Type: core_websocket.EventMessageUpdated,
				Data: patched,
			})
		}
	}

	return patched, nil
}
