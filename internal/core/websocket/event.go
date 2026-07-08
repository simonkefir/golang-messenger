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
