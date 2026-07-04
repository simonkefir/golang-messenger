package messages_service

import (
	"context"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

type MessagesService struct {
	messagesRepository messagesRepository
}

type messagesRepository interface {
	CreateMessage(
		ctx context.Context,
		msg domain.Message,
	) (domain.Message, error)
}

func NewMessagesService(
	messagesRepository messagesRepository,
) *MessagesService {
	return &MessagesService{
		messagesRepository: messagesRepository,
	}
}
