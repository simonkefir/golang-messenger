package users_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_jwt "github.com/simonkefir/golang-messenger/internal/core/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) Login(ctx context.Context, email *string, username *string, password string) (string, error) {
	var user domain.User
	if username == nil {
		var err error
		user, err = s.usersRepository.GetUserByEmail(ctx, *email)
		if err != nil {
			if errors.Is(err, core_errors.ErrNotFound) {
				return "", core_errors.ErrNotFound
			}

			return "", fmt.Errorf("get user by email: %w", err)
		}
	} else if email == nil {
		var err error
		user, err = s.usersRepository.GetUserByUsername(ctx, *username)
		if err != nil {
			if errors.Is(err, core_errors.ErrNotFound) {
				return "", core_errors.ErrNotFound
			}

			return "", fmt.Errorf("get user by username: %w", err)
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", core_errors.ErrInvalidInput
	}

	token, err := core_jwt.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return token, nil
}
