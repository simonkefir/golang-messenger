package messages_service

import (
	"context"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

type MessagesService struct {
	messagesRepository messagesRepository
	chatsChecker       chatsChecker
}

type messagesRepository interface {
	CreateMessage(
		ctx context.Context,
		senderID int64,
		chatID int64,
		content string,
	) (domain.Message, error)
	GetChatMessages(
		ctx context.Context,
		chatID int64,
	) ([]domain.Message, error)
	DeleteMessage(
		ctx context.Context,
		chatID int64,
		messageID int64,
	) error
	PatchMessage(
		ctx context.Context,
		chatID int64,
		messageID int64,
		content string,
	) (domain.Message, error)
}

type chatsChecker interface {
	IsParticipant(
		ctx context.Context,
		chatID int64,
		userID int64,
	) (bool, error)
}

func NewMessagesService(
	messagesRepository messagesRepository,
	chatsChecker chatsChecker,
) *MessagesService {
	return &MessagesService{
		messagesRepository: messagesRepository,
		chatsChecker:       chatsChecker,
	}
}
