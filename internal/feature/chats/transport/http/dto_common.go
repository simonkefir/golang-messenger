package chats_transport_http

import "github.com/simonkefir/golang-messenger/internal/core/domain"

type CreateChatDTO struct {
	ParticipantID int64 `json:"participant_id" validate:"required"`
}

type ParticipantDTO struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

type ChatResponseDTO struct {
	ID           int64            `json:"id"`
	CreatedAt    string           `json:"created_at"`
	Participants []ParticipantDTO `json:"participants"`
}

type ChatListItemDTO struct {
	ID            int64  `json:"id"`
	CreatedAt     string `json:"created_at"`
	CompanionID   int64  `json:"companion_id"`
	CompanionName string `json:"companion_name"`
}

type ChatCreatedResponseDTO struct {
	ID        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
}

func NewChatCreatedResponseFromDomain(chat domain.Chat) *ChatCreatedResponseDTO {
	return &ChatCreatedResponseDTO{
		ID:        chat.ID,
		CreatedAt: chat.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func NewChatResponseFromDomain(chat domain.ChatWithParticipants) *ChatResponseDTO {
	participants := make([]ParticipantDTO, 0, len(chat.Participants))
	for _, p := range chat.Participants {
		participants = append(participants, ParticipantDTO{
			UserID:   p.UserID,
			Username: p.Username,
		})
	}

	return &ChatResponseDTO{
		ID:           chat.ID,
		CreatedAt:    chat.CreatedAt.Format("2006-01-02 15:04:05"),
		Participants: participants,
	}
}
