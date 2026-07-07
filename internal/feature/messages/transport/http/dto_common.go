package messages_transport_http

import "github.com/simonkefir/golang-messenger/internal/core/domain"

type CreateMessageDTO struct {
	Content string `json:"content" validate:"required,min=1,max=4000"`
}

type DeleteMessageDTO struct {
	ID int64 `json:"id" validate:"required,min=1"`
}

type PatchMessageDTO struct {
	ID      int64  `json:"id" validate:"required,min=1"`
	Content string `json:"content" validate:"required,min=1,max=4000"`
}

type MessageDTOResponse struct {
	ID       int64  `json:"id"`
	ChatID   int64  `json:"chat_id"`
	SenderID int64  `json:"sender_id"`
	Content  string `json:"content"`
	SentAt   string `json:"sent_at"`
}

func NewMessageResponseFromDomain(msg domain.Message) *MessageDTOResponse {
	return &MessageDTOResponse{
		ID:       msg.ID,
		ChatID:   msg.ChatID,
		SenderID: msg.SenderID,
		Content:  msg.Content,
		SentAt:   msg.SentAt.Format("2006-01-02 15:04:05"),
	}
}
