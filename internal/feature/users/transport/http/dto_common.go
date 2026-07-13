package users_transport_http

import (
	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=40"   example:"alex"`
	Email    string `json:"email"    validate:"required,email"          example:"alex@mail.com"`
	Password string `json:"password" validate:"required,min=8"          example:"wasd1234"`
}

type LoginUserDTO struct {
	Email    string `json:"email"    validate:"required,email"          example:"alex@mail.com"`
	Password string `json:"password" validate:"required,min=8"          example:"wasd1234"`
}

type PatchUserDTO struct {
	Username *string `json:"username" validate:"omitempty,min=3,max=40" example:"alexg"`
	Password *string `json:"password" validate:"omitempty,min=8"        example:"wasd1212"`
}

type LoginResponse struct {
	Token string `json:"token"                                          example:"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOjEsImV4cCI6NDEwMjQ0NDgwMH0.c36q-l2h78e5VqZQ3tZPbMZXZupO4g3M8hnudYMbdOE"`
}

type UserDTOResponse struct {
	ID        int64  `json:"id"                                         example:"1"`
	Username  string `json:"username"                                   example:"alex"`
	Email     string `json:"email"                                      example:"alex@mail.com"`
	CreatedAt string `json:"created_at"                                 example:"2006-01-02 15:04:05"`
}

func NewUserResponseFromDomain(user *domain.User) *UserDTOResponse {
	return &UserDTOResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
