package users_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *UsersService) GetUserByID(ctx context.Context, userID int64) (domain.User, error) {
	user, err := s.usersRepository.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return domain.User{}, core_errors.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("get user by id: %w", err)
	}

	return user, nil
}

func (s *UsersService) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	user, err := s.usersRepository.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return domain.User{}, core_errors.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("get user by username: %w", err)
	}

	return user, nil
}
