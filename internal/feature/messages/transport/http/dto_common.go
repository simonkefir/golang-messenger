package messages_transport_http

type CreateMessageDTO struct {
	Content string `json:"content" validate:"required,min=1,max=4000"`
}
