package core_websocket

type Event struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

const (
	EventMessageCreated = "message_created"
	EventMessageUpdated = "message_updated"
	EventMessageDeleted = "message_deleted"
	EventChatCreated    = "chat_created"
	EventChatDeleted    = "chat_deleted"
)

type MessagePayload struct {
	ID       int64  `json:"id"`
	ChatID   int64  `json:"chat_id"`
	SenderID int64  `json:"sender_id"`
	Content  string `json:"content"`
	SentAt   string `json:"sent_at"`
}

type MessageDeletedPayload struct {
	ChatID    int64 `json:"chat_id"`
	MessageID int64 `json:"message_id"`
}

type ChatPayload struct {
	ID          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	CompanionID int64  `json:"companion_id"`
}

type ChatDeletedPayload struct {
	ChatID int64 `json:"chat_id"`
}
