package users_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) PatchMe(ctx context.Context, userID int64, username *string, password *string) (domain.User, error) {
	var passwordHash *string

	if password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			return domain.User{}, fmt.Errorf("hash password: %w", err)
		}

		hashedStr := string(hashed)
		passwordHash = &hashedStr
	}

	patched, err := s.usersRepository.PatchMe(ctx, userID, username, passwordHash)
	if err != nil {
		if errors.Is(err, core_errors.ErrAlreadyExists) {
			return domain.User{}, core_errors.ErrAlreadyExists
		}
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return patched, nil
}
