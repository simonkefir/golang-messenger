package users_service

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	user.PasswordHash = string(hashed)

	created, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, core_errors.ErrAlreadyExists) {
			return domain.User{}, core_errors.ErrAlreadyExists
		}
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return created, nil
}
