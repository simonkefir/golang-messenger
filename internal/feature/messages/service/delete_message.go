package messages_service

import (
	"context"
	"fmt"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_websocket "github.com/simonkefir/golang-messenger/internal/core/websocket"
)

func (s *MessagesService) DeleteMessage(ctx context.Context, senderID int64, chatID int64, messageID int64) error {
	ok, err := s.chatsChecker.IsParticipant(ctx, chatID, senderID)
	if err != nil {
		return fmt.Errorf("check participant: %w", err)
	}
	if !ok {
		return core_errors.ErrForbidden
	}

	if err := s.messagesRepository.DeleteMessage(ctx, chatID, messageID); err != nil {
		return fmt.Errorf("delete message: %w", err)
	}

	participant, err := s.chatsChecker.GetChatParticipant(ctx, chatID, senderID)
	if err == nil {
		s.publisher.Publish(participant.UserID, core_websocket.Event{
			Type: core_websocket.EventMessageDeleted,
			Data: core_websocket.MessageDeletedPayload{
				ChatID:    chatID,
				MessageID: messageID,
			},
		})
	}
	return nil
}
