package chats_transport_http

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_server "github.com/simonkefir/golang-messenger/internal/core/transport/http/server"
	core_websocket "github.com/simonkefir/golang-messenger/internal/core/websocket"
)

type ChatsService interface {
	CreateChat(
		ctx context.Context,
		userID int64,
		chat_participantID int64,
	) (domain.ChatWithParticipant, error)
	DeleteChat(
		ctx context.Context,
		userID int64,
		chatID int64,
	) error
	GetChat(
		ctx context.Context,
		userID int64,
		chatID int64,
	) (domain.ChatWithParticipant, error)
	GetChats(
		ctx context.Context,
		userID int64,
	) ([]domain.ChatListItem, error)
}

type ChatsHTTPHandler struct {
	svc       ChatsService
	validate  *validator.Validate
	publisher core_websocket.EventPublisher
}

func NewChatsHTTPHandler(svc ChatsService, publisher core_websocket.EventPublisher) *ChatsHTTPHandler {
	return &ChatsHTTPHandler{
		svc:       svc,
		validate:  validator.New(),
		publisher: publisher,
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
