package users_transport_http

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_http_server "github.com/simonkefir/golang-messenger/internal/core/transport/http/server"
)

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	Login(
		ctx context.Context,
		email string,
		password string,
	) (string, error)
}

type UsersHTTPHandler struct {
	svc      UsersService
	validate *validator.Validate
}

func NewUsersHTTPHandler(svc UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		svc:      svc,
		validate: validator.New(),
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users/register",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users/login",
			Handler: h.Login,
		},
	}
}
