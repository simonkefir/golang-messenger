package users_transport_http

import (
	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=40"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginUserDTO struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type PatchUserDto struct {
	Username *string `json:"username" validate:"omitempty,min=3,max=40"`
	Password *string `json:"password" validate:"omitempty,min=8"`
}

type UserDTOResponse struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func NewUserResponseFromDomain(user *domain.User) *UserDTOResponse {
	return &UserDTOResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
