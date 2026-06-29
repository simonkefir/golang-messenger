package users_service

import (
	"context"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

type UsersService struct {
	usersRepository usersRepository
}

type usersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUserByEmail(
		ctx context.Context,
		email string,
	) (domain.User, error)
}

func NewUsersService(
	usersRepository usersRepository,
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}
