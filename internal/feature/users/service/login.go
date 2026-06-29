package users_service

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	"github.com/simonkefir/golang-messenger/internal/core/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.usersRepository.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return "", fmt.Errorf("get user by email: %w", err)
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", core_errors.ErrInvalidInput
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return token, nil
	// 1. находим юзера по email
	// 2. bcrypt.CompareHashAndPassword
	// 3. генерируем JWT
	// 4. возвращаем токен
}
