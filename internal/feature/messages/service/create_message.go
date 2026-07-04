package messages_service

import (
	"context"

	"github.com/simonkefir/golang-messenger/internal/core/domain"
)

func (s *MessagesService) CreateMessage(ctx context.Context, msg domain.Message) error
