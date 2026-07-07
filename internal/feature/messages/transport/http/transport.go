package messages_transport_http

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_server "github.com/simonkefir/golang-messenger/internal/core/transport/http/server"
)

type MessagesService interface {
	CreateMessage(
		ctx context.Context,
		userID int64,
		chatID int64,
		content string,
	) (domain.Message, error)
	GetMessages(
		ctx context.Context,
		userID int64,
		chatID int64,
	) ([]domain.Message, error)
	DeleteMessage(
		ctx context.Context,
		userID int64,
		chatID int64,
		messageID int64,
	) error
	PatchMessage(
		ctx context.Context,
		userID int64,
		chatID int64,
		messageID int64,
		content string,
	) (domain.Message, error)
}

type MessagesHTTPHandler struct {
	svc      MessagesService
	validate *validator.Validate
}

func NewMessagesHTTPHandler(svc MessagesService) *MessagesHTTPHandler {
	return &MessagesHTTPHandler{
		svc:      svc,
		validate: validator.New(),
	}
}

func (h *MessagesHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodPost,
			Path:       "/chats/{chat_id}/messages",
			Handler:    h.CreateMessage,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWTMiddleware},
		},
		{
			Method:     http.MethodGet,
			Path:       "/chats/{chat_id}/messages",
			Handler:    h.GetMessages,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWTMiddleware},
		},
		{
			Method:     http.MethodDelete,
			Path:       "/chats/{chat_id}/messages",
			Handler:    h.DeleteMessage,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWTMiddleware},
		},
		{
			Method:     http.MethodPatch,
			Path:       "/chats/{chat_id}/messages",
			Handler:    h.PatchMessage,
			Middleware: []core_http_middleware.Middleware{core_http_middleware.JWTMiddleware},
		},
	}
}
