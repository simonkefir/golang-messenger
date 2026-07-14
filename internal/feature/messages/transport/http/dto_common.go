package messages_transport_http

import "github.com/simonkefir/golang-messenger/internal/core/domain"

type CreateMessageDTO struct {
	Content string `json:"content" validate:"required,min=1,max=4000" example:"mary is the girl, who i wanna f.."`
}

type DeleteMessageDTO struct {
	ID int64 `json:"id" validate:"required,min=1"    example:"71670"`
}

type PatchMessageDTO struct {
	ID      int64  `json:"id" validate:"required,min=1"               example:"17174"`
	Content string `json:"content" validate:"required,min=1,max=4000" example:"mary is the girl, who i wanna kiss"`
}

type MessageDTOResponse struct {
	ID       int64  `json:"id"        example:"777"`
	ChatID   int64  `json:"chat_id"   example:"3"`
	SenderID int64  `json:"sender_id" example:"333" `
	Content  string `json:"content"   example:"its not what u r, its just what u did, dont hang up the phone, i lov3 u to death, eternal return..."`
	SentAt   string `json:"sent_at"   example:"2006-01-02 15:04:05"`
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
