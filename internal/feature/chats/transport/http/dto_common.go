package chats_transport_http

import "github.com/simonkefir/golang-messenger/internal/core/domain"

type CreateChatDTO struct {
	ParticipantID int64 `json:"participant_id" validate:"required" example:"333"`
}

type ParticipantDTO struct {
	UserID      int64  `json:"user_id"          example:"333"`
	DisplayName string `json:"display_name"     example:"alex g"`
}

type ChatDTOResponse struct {
	ID          int64          `json:"id"           example:"333"`
	CreatedAt   string         `json:"created_at"   example:"2006-01-02 15:04:05"`
	Participant ParticipantDTO `json:"participant"`
}

type ChatListItemDTO struct {
	ID                   int64  `json:"id"              example:"333"`
	CreatedAt            string `json:"created_at"      example:"2006-01-02 15:04:05"`
	CompanionID          int64  `json:"companion_id"    example:"141"`
	CompanionDisplayName string `json:"companion_name"  example:"alex g"`
}

func NewChatResponseFromDomain(chat domain.Chat, participant ParticipantDTO) *ChatDTOResponse {
	return &ChatDTOResponse{
		ID:          chat.ID,
		CreatedAt:   chat.CreatedAt.Format("2006-01-02 15:04:05"),
		Participant: participant,
	}
}
