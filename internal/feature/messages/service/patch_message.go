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

	patched, err := s.messagesRepository.PatchMessage(ctx, messageID, senderID, content)
	if err != nil {
		if errors.Is(err, core_errors.ErrAlreadyExists) {
			return domain.Message{}, core_errors.ErrAlreadyExists
		}
		return domain.Message{}, fmt.Errorf("patch message: %w", err)
	}

	participant, err := s.chatsChecker.GetChatParticipant(ctx, chatID, senderID)
	if err == nil {
		s.publisher.Publish(participant.UserID, core_websocket.Event{
			Type: core_websocket.EventMessageUpdated,
			Data: core_websocket.MessagePayload{
				ID:       patched.ID,
				ChatID:   patched.ChatID,
				SenderID: patched.SenderID,
				Content:  patched.Content,
				SentAt:   patched.SentAt.Format("2006-01-02 15:04:05"),
			},
		})
	}

	return patched, nil
}
