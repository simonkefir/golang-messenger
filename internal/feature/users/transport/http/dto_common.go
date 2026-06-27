package users_transport_http

import (
	"time"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=40"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserDTOResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserResponseFromDomain(user *domain.User) *UserDTOResponse {
	return &UserDTOResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
