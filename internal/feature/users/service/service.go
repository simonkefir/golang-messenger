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
	GetUserByID(
		ctx context.Context,
		userID int64,
	) (domain.User, error)
	GetUserByUsername(
		ctx context.Context,
		username string,
	) (domain.User, error)
	PatchMe(
		ctx context.Context,
		userID int64,
		username *string,
		display_name *string,
		email *string,
		password *string,
	) (domain.User, error)
	DeleteMe(
		ctx context.Context,
		userID int64,
	) error
}

func NewUsersService(
	usersRepository usersRepository,
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}
