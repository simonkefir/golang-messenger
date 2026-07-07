package chats_transport_http

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_server "github.com/simonkefir/golang-messenger/internal/core/transport/http/server"
)

type ChatsService interface {
	CreateChat(
		ctx context.Context,
		userID int64,
		chat_participant int64,
	) (domain.Chat, error)
	DeleteChat(
		ctx context.Context,
		userID int64,
		chatID int64,
	) error
	GetChat(
		ctx context.Context,
		userID int64,
		chatID int64,
	) (domain.ChatWithParticipants, error)
	GetChats(
		ctx context.Context,
		userID int64,
	) ([]domain.ChatListItem, error)
}

type ChatsHTTPHandler struct {
	svc      ChatsService
	validate *validator.Validate
}

func NewChatsHTTPHandler(svc ChatsService) *ChatsHTTPHandler {
	return &ChatsHTTPHandler{
		svc:      svc,
		validate: validator.New(),
	}
}

func (h *ChatsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodPost,
			Path:       "/chats",
			Handler:    h.CreateChat,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWTMiddleware},
		},
		{
			Method:     http.MethodDelete,
			Path:       "/chats/{id}",
			Handler:    h.DeleteChat,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWTMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/chats/{id}",
			Handler:    h.GetChat,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWTMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/chats",
			Handler:    h.GetChats,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWTMiddleware},
		},
	}

}
