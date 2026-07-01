package users_service

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
)

func (s *UsersService) DeleteMe(ctx context.Context, userID int64) error {
	if err := s.usersRepository.DeleteMe(ctx, userID); err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return fmt.Errorf("delete user by id: %w", err)
		}
	}

	return nil
}
