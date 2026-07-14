package chats_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_websocket "github.com/simonkefir/golang-messenger/internal/core/websocket"
)

func (s *ChatsService) CreateChat(ctx context.Context, userID int64, chat_participant int64) (domain.ChatWithParticipant, error) {
	if userID == chat_participant {
		return domain.ChatWithParticipant{}, core_errors.ErrInvalidInput
	}

	existingChat, err := s.chatsRepository.FindPrivateChat(ctx, userID, chat_participant)
	if err == nil {
		return existingChat, nil
	}
	if !errors.Is(err, core_errors.ErrNotFound) {
		return domain.ChatWithParticipant{}, fmt.Errorf("find private chat: %w", err)
	}

	chat, err := s.chatsRepository.CreateChat(ctx, userID, chat_participant)
	if err != nil {
		return domain.ChatWithParticipant{}, fmt.Errorf("create chat: %w", err)
	}

	participant, err := s.chatsRepository.GetChatParticipant(ctx, chat.ID, userID)
	if err != nil {
		return domain.ChatWithParticipant{Chat: chat}, fmt.Errorf("get chat participant: %w", err)
	}

	s.publisher.Publish(chat_participant, core_websocket.Event{
		Type: core_websocket.EventChatCreated,
		Data: chat,
	})

	return domain.ChatWithParticipant{Chat: chat, Participant: participant}, nil
}
