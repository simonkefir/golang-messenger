package users_transport_http

import (
	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

type CreateUserDTO struct {
	Username    string `json:"username"     validate:"required,min=1,max=30"     example:"alex"`
	DisplayName string `json:"display_name" validate:"required,min=1,max=40"     example:"alex g"`
	Email       string `json:"email"        validate:"required,email,max=100"    example:"alex@mail.com"`
	Password    string `json:"password"     validate:"required,min=8"            example:"wasd1234"`
}

type LoginUserDTO struct {
	Email    *string `json:"email"    validate:"omitempty,email,max=100"           example:"alex@mail.com"`
	Username *string `json:"username" validate:"omitempty,min=1,max=30"            example:"alex"`
	Password string  `json:"password" validate:"required,min=8"                    example:"wasd1234"`
}

type PatchUserDTO struct {
	Username    *string `json:"username"     validate:"omitempty,max=30"       example:"alex"`
	DisplayName *string `json:"display_name" validate:"omitempty,max=40"       example:"alex g"`
	Email       *string `json:"email"        validate:"omitempty,max=100"      example:"alexg@example.com"`
	Password    *string `json:"password"     validate:"omitempty,min=8"        example:"wasd1212"`
}

type LoginResponse struct {
	Token string `json:"token"                                            example:"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOjEsImV4cCI6NDEwMjQ0NDgwMH0.c36q-l2h78e5VqZQ3tZPbMZXZupO4g3M8hnudYMbdOE"`
}

type UserDTOResponse struct {
	ID          int64  `json:"id"                                         example:"1"`
	Username    string `json:"username"                                   example:"alex"`
	DisplayName string `json:"display_name"                               example:"alex g"`
	Email       string `json:"email"                                      example:"alex@mail.com"`
	CreatedAt   string `json:"created_at"                                 example:"2006-01-02 15:04:05"`
}

func NewUserResponseFromDomain(user *domain.User) *UserDTOResponse {
	return &UserDTOResponse{
		ID:          user.ID,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
