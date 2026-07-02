package chats_service

import (
	"context"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

type ChatsService struct {
	chatsRepository chatsRepository
}

type chatsRepository interface {
	CreateChat(
		ctx context.Context,
		userID int64,
		chat_participant int64,
	) (domain.Chat, error)
	FindPrivateChat(
		ctx context.Context,
		user1 int64,
		user2 int64,
	) (domain.Chat, error)
	IsParticipant(
		ctx context.Context,
		userID int64,
		chatID int64,
	) (bool, error)
	DeleteChat(
		ctx context.Context,
		chatID int64,
	) error
	GetChatByChatID(
		ctx context.Context,
		chatID int64,
	) (domain.Chat, error)
	GetChatsByUserID(
		ctx context.Context,
		userID int64,
	) ([]domain.ChatListItem, error)
	GetChatParticipants(
		ctx context.Context,
		chatID int64,
	) ([]domain.ChatParticipant, error)
}

func NewChatsService(
	chatsRepository chatsRepository,
) *ChatsService {
	return &ChatsService{
		chatsRepository: chatsRepository,
	}
}
